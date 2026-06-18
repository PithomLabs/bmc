package convergence

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// ConvergenceRun represents the results of a single trajectory integration run.
type ConvergenceRun struct {
	RunID                           string   `json:"run_id"`
	Stepper                         string   `json:"stepper"`
	DeltaLambda                     float64  `json:"delta_lambda"`
	Steps                           int      `json:"steps"`
	TrajectoryFinite                bool     `json:"trajectory_finite"`
	NodeContactDetected             bool     `json:"node_contact_detected"`
	NodeContactStep                 *int     `json:"node_contact_step,omitempty"`
	FinalAlpha                      *float64 `json:"final_alpha"`
	FinalPhi                        *float64 `json:"final_phi"`
	EndpointDistanceToReference     *float64 `json:"endpoint_distance_to_reference,omitempty"`
	MaxPointwiseDistanceToReference *float64 `json:"max_pointwise_distance_to_reference,omitempty"`
	Status                          string   `json:"status"`
	Notes                           []string `json:"notes"`
}

// ConvergenceAudit holds the complete diagnostic audit of stepper and dt convergence.
type ConvergenceAudit struct {
	ProfileID            string           `json:"profile_id"`
	ToyAnalysisOnly      bool             `json:"toy_analysis_only"`
	PhysicsClaim         string           `json:"physics_claim"`
	ReferenceRunID       string           `json:"reference_run_id"`
	Runs                 []ConvergenceRun `json:"runs"`
	InterpretationStatus string           `json:"interpretation_status"`
	Warnings             []string         `json:"warnings"`
}

// Allowed run statuses
const (
	StatusComputed              = "computed"
	StatusBlockedNodeContact    = "blocked_by_node_contact"
	StatusBlockedNonfinite      = "blocked_by_nonfinite_trajectory"
	StatusBlockedRefUnavailable = "blocked_by_reference_unavailable"
	StatusInvalidInput          = "invalid_input"
)

// Allowed interpretation statuses
const (
	InterpretationComparisonOnly            = "numerical_comparison_only"
	InterpretationMixed                     = "convergence_mixed"
	InterpretationUnstable                  = "convergence_unstable"
	InterpretationCandidateStableUnpromoted = "convergence_candidate_stable_unpromoted"
	InterpretationBlockedNodeContact        = "blocked_by_node_contact"
	InterpretationBlockedNonfinite          = "blocked_by_nonfinite_trajectory"
)

// LambdaAlignmentTolerance is the tolerance allowed for checking that lambda values match between coarse and ref runs.
const LambdaAlignmentTolerance = 1e-9

// RunAudit performs the stepper and time-step refinement audit on the given superposition parameters.
func RunAudit(params model.SuperpositionParams) (*ConvergenceAudit, error) {

	// Reject invalid steps or step size
	if math.IsNaN(params.LambdaStep) || math.IsInf(params.LambdaStep, 0) || params.LambdaStep <= 0 {
		return nil, fmt.Errorf("invalid step size: lambda_step must be finite and positive, got %f", params.LambdaStep)
	}

	if params.Steps <= 0 {
		return nil, fmt.Errorf("invalid steps: steps must be greater than zero, got %d", params.Steps)
	}

	warnings := []string{
		"Toy audit analysis only.",
		"Does not test full quantum gravity.",
		"This audit compares numerical self-consistency of toy trajectory integration under stepper and dt refinement; it does not test whether the trajectory is physically correct.",
		"A finest RK4 run is a local numerical reference, not physical ground truth.",
		"Passing this audit cannot promote any final-truth claim.",
	}

	audit := &ConvergenceAudit{
		ProfileID:            "bmc0a-superposition-safe",
		ToyAnalysisOnly:      true,
		PhysicsClaim:         "minisuperspace_only",
		ReferenceRunID:       "rk4_dt_4",
		InterpretationStatus: InterpretationComparisonOnly,
		Warnings:             warnings,
	}

	wf := wave.NewSuperpositionWave(
		params.C1Real, params.C1Imag, params.K1, params.Omega1,
		params.C2Real, params.C2Imag, params.K2, params.Omega2,
	)

	initial := model.MiniState{Alpha: params.Alpha0, Phi: params.Phi0}
	nodeThresh := params.NodeThresh

	dt := params.LambdaStep
	steps := params.Steps

	runConfigs := []struct {
		runID       string
		stepperName string
		stepper     guidance.Stepper
		stepFactor  int
	}{
		{"euler_dt", "euler", guidance.NewEulerStepper(), 1},
		{"rk4_dt", "rk4", guidance.NewRK4Stepper(), 1},
		{"rk4_dt_2", "rk4", guidance.NewRK4Stepper(), 2},
		{"rk4_dt_4", "rk4", guidance.NewRK4Stepper(), 4},
	}

	trajectories := make(map[string]model.Trajectory)
	runs := make(map[string]*ConvergenceRun)
	var runList []ConvergenceRun

	// Integrate trajectories
	for _, config := range runConfigs {
		runDt := dt / float64(config.stepFactor)
		runSteps := steps * config.stepFactor

		traj := guidance.Integrate(wf, initial, config.stepper, runDt, runSteps, nodeThresh)
		trajectories[config.runID] = traj

		nodeContact, contactIdx := detectNodeContact(wf, traj, nodeThresh)
		finite := isFinite(traj)

		run := &ConvergenceRun{
			RunID:               config.runID,
			Stepper:             config.stepperName,
			DeltaLambda:         runDt,
			Steps:               runSteps,
			TrajectoryFinite:    finite,
			NodeContactDetected: nodeContact,
			Notes:               []string{},
		}

		if nodeContact {
			run.Status = StatusBlockedNodeContact
			stepNum := contactIdx
			run.NodeContactStep = &stepNum
			run.Notes = append(run.Notes, fmt.Sprintf("Node contact detected at step %d", contactIdx))

			// Record partial coordinate if finite (index before node contact if index > 0)
			if contactIdx > 0 && contactIdx < len(traj.Points) {
				prevPt := traj.Points[contactIdx-1]
				aCopy := prevPt.State.Alpha
				pCopy := prevPt.State.Phi
				run.FinalAlpha = &aCopy
				run.FinalPhi = &pCopy
				run.Notes = append(run.Notes, fmt.Sprintf("Last finite coordinate recorded: alpha=%f, phi=%f", aCopy, pCopy))
			}
		} else if !finite {
			run.Status = StatusBlockedNonfinite
			run.Notes = append(run.Notes, "Trajectory contains non-finite values (NaN/Inf) without node contact.")
		} else {
			run.Status = StatusComputed
			lastIdx := len(traj.Points) - 1
			aCopy := traj.Points[lastIdx].State.Alpha
			pCopy := traj.Points[lastIdx].State.Phi
			run.FinalAlpha = &aCopy
			run.FinalPhi = &pCopy
		}

		runs[config.runID] = run
	}

	// Check reference run availability
	refRun, refOk := runs["rk4_dt_4"]
	refTraj, refTrajOk := trajectories["rk4_dt_4"]
	refAvailable := refOk && refTrajOk && refRun.Status == StatusComputed

	// Process comparisons
	for _, config := range runConfigs {
		run := runs[config.runID]
		traj := trajectories[config.runID]

		if run.Status == StatusComputed {
			if !refAvailable {
				run.Status = StatusBlockedRefUnavailable
				run.Notes = append(run.Notes, "Reference run is unavailable.")
			} else {
				// Compute endpoint distance to reference
				lastIdx := len(traj.Points) - 1
				refLastIdx := len(refTraj.Points) - 1

				alphaRun := traj.Points[lastIdx].State.Alpha
				phiRun := traj.Points[lastIdx].State.Phi

				alphaRef := refTraj.Points[refLastIdx].State.Alpha
				phiRef := refTraj.Points[refLastIdx].State.Phi

				dist := math.Sqrt((alphaRun-alphaRef)*(alphaRun-alphaRef) + (phiRun-phiRef)*(phiRun-phiRef))
				run.EndpointDistanceToReference = &dist

				// Pointwise comparison aligned by lambda
				stepRatio := 4 / config.stepFactor
				maxPointwise := 0.0
				alignmentPassed := true

				for i := 0; i <= lastIdx; i++ {
					refIdx := i * stepRatio
					if refIdx >= len(refTraj.Points) {
						alignmentPassed = false
						run.Notes = append(run.Notes, fmt.Sprintf("Pointwise index alignment index out of bounds at i=%d", i))
						break
					}

					coarsePt := traj.Points[i]
					refPt := refTraj.Points[refIdx]

					// Explicitly assert lambda values match within tolerance
					if math.Abs(coarsePt.Lambda-refPt.Lambda) > LambdaAlignmentTolerance {

						alignmentPassed = false
						run.Notes = append(run.Notes, fmt.Sprintf("Lambda mismatch at step i=%d: coarse lambda=%f, ref lambda=%f", i, coarsePt.Lambda, refPt.Lambda))
						break
					}

					dAlpha := coarsePt.State.Alpha - refPt.State.Alpha
					dPhi := coarsePt.State.Phi - refPt.State.Phi
					ptDist := math.Sqrt(dAlpha*dAlpha + dPhi*dPhi)
					if ptDist > maxPointwise {
						maxPointwise = ptDist
					}
				}

				if alignmentPassed {
					run.MaxPointwiseDistanceToReference = &maxPointwise
				} else {
					run.Notes = append(run.Notes, "Pointwise alignment failed; max pointwise distance omitted.")
				}
			}
		} else {
			run.EndpointDistanceToReference = nil
			run.MaxPointwiseDistanceToReference = nil
		}

		runList = append(runList, *run)
	}

	audit.Runs = runList

	// Evaluate Interpretation Status
	if refRun.Status == StatusBlockedNodeContact {
		audit.InterpretationStatus = InterpretationBlockedNodeContact
	} else if refRun.Status == StatusBlockedNonfinite {
		audit.InterpretationStatus = InterpretationBlockedNonfinite
	} else {
		hasNodeContact := false
		hasNonfinite := false
		allComputed := true

		for _, r := range runList {
			if r.Status == StatusBlockedNodeContact {
				hasNodeContact = true
				allComputed = false
			}
			if r.Status == StatusBlockedNonfinite {
				hasNonfinite = true
				allComputed = false
			}
			if r.Status == StatusBlockedRefUnavailable || r.Status == StatusInvalidInput {
				allComputed = false
			}
		}

		if hasNodeContact {
			audit.InterpretationStatus = InterpretationBlockedNodeContact
		} else if hasNonfinite {
			audit.InterpretationStatus = InterpretationBlockedNonfinite
		} else if allComputed {
			rk4_dt, ok1 := runs["rk4_dt"]
			rk4_dt_2, ok2 := runs["rk4_dt_2"]
			if ok1 && ok2 && rk4_dt.EndpointDistanceToReference != nil && rk4_dt_2.EndpointDistanceToReference != nil {
				drift1 := *rk4_dt.EndpointDistanceToReference
				drift2 := *rk4_dt_2.EndpointDistanceToReference
				if drift2 <= drift1 {
					audit.InterpretationStatus = InterpretationCandidateStableUnpromoted
				} else {
					audit.InterpretationStatus = InterpretationUnstable
				}
			} else {
				audit.InterpretationStatus = InterpretationComparisonOnly
			}
		} else {
			audit.InterpretationStatus = InterpretationMixed
		}
	}

	return audit, nil
}

// detectNodeContact checks if a trajectory contacts a node at any step.
func detectNodeContact(wf wave.WaveFunction, traj model.Trajectory, nodeThresh float64) (bool, int) {
	for i, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
			if i > 0 {
				prior := traj.Points[i-1].State
				if !math.IsNaN(prior.Alpha) && !math.IsNaN(prior.Phi) {
					rVal := wave.AmplitudeField(prior.Alpha, prior.Phi, wf)
					if rVal < nodeThresh {
						return true, i
					}
				}
			} else {
				return true, 0
			}
		} else {
			rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, wf)
			if rVal < nodeThresh {
				return true, i
			}
		}
	}
	return false, -1
}

// isFinite checks if all coordinates in the trajectory are finite.
func isFinite(traj model.Trajectory) bool {
	for _, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsInf(p.State.Alpha, 0) ||
			math.IsNaN(p.State.Phi) || math.IsInf(p.State.Phi, 0) {
			return false
		}
	}
	return true
}
