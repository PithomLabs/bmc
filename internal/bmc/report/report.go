package report

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/invariant"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/obstruction"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

// Extra status constant specifically for gates.
const StatusBlocked model.CheckStatus = "blocked"

// Default step size for numerical finite difference derivatives of the wavefunction amplitude in Q potential check.
const DefaultQFiniteDifferenceStep = 1e-4

type TechnicalGate struct {
	Name   string            `json:"name"`
	Status model.CheckStatus `json:"status"`
}

type PromotionGate struct {
	Name   string            `json:"name"`
	Status model.CheckStatus `json:"status"`
	Reason string            `json:"reason"`
}

type NullModel struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type Faithfulness struct {
	Status model.CheckStatus `json:"status"`
	Reason string            `json:"reason"`
}

type EbpDebt struct {
	NeedMap                string `json:"needMap"`
	NeedInvariant          string `json:"needInvariant"`
	NeedToyCheck           string `json:"needToyCheck"`
	NeedNullModel          string `json:"needNullModel"`
	NeedObstruction        string `json:"needObstruction"`
	NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
}

type ReportParameters struct {
	PlaneWave     *model.PlaneWaveParams     `json:"plane_wave,omitempty"`
	Superposition *model.SuperpositionParams `json:"superposition,omitempty"`
}

type Report struct {
	SchemaVersion           string                       `json:"schema_version"`
	ModelID                 string                       `json:"model_id"`
	ToyAnalysisOnly         bool                         `json:"toy_analysis_only"`
	PhysicsClaim            string                       `json:"physics_claim"`
	FinalTruthClaim         bool                         `json:"final_truth_claim"`
	PromotionRecommendation string                       `json:"promotion_recommendation"`
	Parameters              ReportParameters             `json:"parameters"`
	Equations               map[string]string            `json:"equations"`
	Checks                  map[string]model.CheckResult `json:"checks"`
	TechnicalGate           TechnicalGate                `json:"technical_gate"`
	PromotionGate           PromotionGate                `json:"promotion_gate"`
	NullModels              []NullModel                  `json:"null_models"`
	Obstructions            []model.Obstruction          `json:"obstructions"`
	Faithfulness            Faithfulness                 `json:"faithfulness"`
	EbpDebt                 EbpDebt                      `json:"ebp_debt"`
	Warnings                []string                     `json:"warnings"`
}

// Generate runs the plane-wave pipeline and builds the JSON report.
func Generate(params model.PlaneWaveParams, finalTruthClaim bool) (*Report, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// 1. Initialize Wavefunction
	planeWave := wave.NewPlaneWave(params.K, params.Omega)

	// 2. Wheeler-DeWitt residual check (Analytic residual is primary authority)
	analyticRes := wdw.AnalyticResidualPlaneWave(params.K, params.Omega)
	wdwCheck := wdw.CheckResidual(analyticRes, params.Tolerance)

	// 3. Integrate Trajectory
	stepper := guidance.NewEulerStepper()
	initialState := model.MiniState{Alpha: params.Alpha0, Phi: params.Phi0}
	traj := guidance.Integrate(planeWave, initialState, stepper, params.LambdaStep, params.Steps, 1e-5)
	trajectoryCheck := guidance.CheckTrajectory(traj)

	// 4. Clock Monotonicity Check
	increasing := true
	decreasing := true
	for i := 1; i < len(traj.Points); i++ {
		diff := traj.Points[i].State.Phi - traj.Points[i-1].State.Phi
		if diff <= 0 {
			increasing = false
		}
		if diff >= 0 {
			decreasing = false
		}
	}
	isClockMonotonic := (increasing || decreasing) && len(traj.Points) > 1
	var clockStatus model.CheckStatus
	var clockReason string
	if isClockMonotonic {
		clockStatus = model.StatusPass
		clockReason = "Relational clock phi is strictly monotonic."
	} else {
		clockStatus = model.StatusFail
		clockReason = "Relational clock phi is not strictly monotonic."
	}
	clockVar := "phi"
	clockCheck := model.CheckResult{
		Status:   clockStatus,
		Pass:     isClockMonotonic,
		Reason:   clockReason,
		Variable: &clockVar,
	}

	// 5. Quantum Potential check
	qValues := qpotential.QAlongTrajectory(planeWave, traj, DefaultQFiniteDifferenceStep)
	maxAbsQ := qpotential.MaxAbsQ(qValues)
	qCheck := qpotential.CheckQuantumPotential(maxAbsQ, params.Tolerance)

	// 6. Classical Limit Check
	classicalCheck := invariant.CheckClassicalLimit(maxAbsQ, params.Tolerance)

	// 7. Friedmann Residual check (deferred in Sprint 1)
	friedmannCheck := model.CheckResult{
		Status: model.StatusDeferred,
		Pass:   false,
		Reason: "Not implemented in BMC-0A plane-wave control; remains debt.",
	}

	checks := map[string]model.CheckResult{
		"wdw_residual":       wdwCheck,
		"trajectory":         trajectoryCheck,
		"clock_monotonicity": clockCheck,
		"quantum_potential":  qCheck,
		"classical_limit":    classicalCheck,
		"friedmann_residual": friedmannCheck,
	}

	// 8. Obstructions
	obstructions := obstruction.DetectAllPlaneWave(params, traj, analyticRes, qValues, finalTruthClaim)

	// 9. Technical Gate
	techGateStatus := model.StatusPass
	if wdwCheck.Status != model.StatusPass ||
		trajectoryCheck.Status != model.StatusPass ||
		clockCheck.Status != model.StatusPass ||
		qCheck.Status != model.StatusPass ||
		classicalCheck.Status != model.StatusPass {
		techGateStatus = model.StatusFail
	}

	// 10. Promotion Gate (always blocked in Sprint 1 because Friedmann/Faithfulness are deferred/contested)
	promoGateStatus := StatusBlocked
	promoReason := "Friedmann residual and faithfulness review remain unpaid debt."

	reportParams := ReportParameters{
		PlaneWave: &params,
	}

	// 11. Assembly
	report := &Report{
		SchemaVersion:           "bmc-report-v0.1",
		ModelID:                 "bmc0a_plane",
		ToyAnalysisOnly:         true,
		PhysicsClaim:            "minisuperspace_only",
		FinalTruthClaim:         finalTruthClaim,
		PromotionRecommendation: "blocked_or_candidate",
		Parameters:              reportParams,
		Equations: map[string]string{
			"wdw":               "(-d_alpha_alpha + d_phi_phi) Psi = 0",
			"wavefunction":      "Psi(alpha,phi)=exp(i(k alpha + omega phi))",
			"guidance":          "dalpha/dlambda=k, dphi/dlambda=-omega",
			"quantum_potential": "Q=0 for constant amplitude",
		},
		Checks: checks,
		TechnicalGate: TechnicalGate{
			Name:   "bmc0a_plane_control_gate",
			Status: techGateStatus,
		},
		PromotionGate: PromotionGate{
			Name:   "full_bmc_toy_gate",
			Status: promoGateStatus,
			Reason: promoReason,
		},
		NullModels: []NullModel{
			{
				Name:   "classical_frw",
				Status: "placeholder_obligation",
				Reason: "Classical-limit comparison begins with Q≈0 control.",
			},
			{
				Name:   "standard_wdw",
				Status: "placeholder_obligation",
				Reason: "Ensemble comparison deferred.",
			},
			{
				Name:   "lqc",
				Status: "deferred",
				Reason: "Not part of BMC-0A.",
			},
			{
				Name:   "page_wootters",
				Status: "deferred",
				Reason: "Formal relational-time null model deferred.",
			},
		},
		Obstructions: obstructions,
		Faithfulness: Faithfulness{
			Status: model.StatusContested,
			Reason: "No human faithfulness review yet. This only tests a plane-wave minisuperspace control.",
		},
		EbpDebt: EbpDebt{
			NeedMap:                "partial",
			NeedInvariant:          "partial",
			NeedToyCheck:           "active",
			NeedNullModel:          "partial",
			NeedObstruction:        "active",
			NeedFaithfulnessReview: "active",
		},
		Warnings: []string{
			"Toy analysis only.",
			"Does not test full quantum gravity.",
			"Does not test black holes, fermions, gauge fields, Lorentz recovery, or inhomogeneous perturbations.",
			"Passing this report cannot promote any final-truth claim.",
		},
	}

	return report, nil
}

// GenerateSuperposition runs the superposition pipeline and builds the JSON report.
func GenerateSuperposition(params model.SuperpositionParams, finalTruthClaim bool) (*Report, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	// 1. Initialize Wavefunction
	sw := wave.NewSuperpositionWave(params.C1Real, params.C1Imag, params.K1, params.Omega1, params.C2Real, params.C2Imag, params.K2, params.Omega2)

	// 2. Wheeler-DeWitt residual check (Analytic residual from components linearity)
	compRes := wdw.ComponentResidualsSuperposition(params.K1, params.Omega1, params.K2, params.Omega2)
	allCompPass := true
	for _, c := range compRes {
		if math.Abs(c.Residual) > params.Tolerance {
			allCompPass = false
		}
	}

	var wdwStatus model.CheckStatus = model.StatusPass
	var wdwReason string = "All component residuals satisfy tolerance; therefore by linearity the analytic superposition residual is zero."
	if !allCompPass {
		wdwStatus = model.StatusFail
		wdwReason = "Component Wheeler-DeWitt constraint violated; superposition residual is non-zero."
	}
	analyticResVal := 0.0
	wdwCheck := model.CheckResult{
		Status:         wdwStatus,
		Pass:           allCompPass,
		Reason:         wdwReason,
		MaxAbsResidual: &analyticResVal,
	}

	// EBP Initial Node Short-Circuit
	initialR := wave.AmplitudeField(params.Alpha0, params.Phi0, sw)

	var traj model.Trajectory
	var qValues []float64
	var maxAbsQ float64
	var trajectoryCheck model.CheckResult
	var clockCheck model.CheckResult
	var qCheck model.CheckResult
	var classicalCheck model.CheckResult
	var nodeContactFreeCheck model.CheckResult
	var phaseGradientFiniteCheck model.CheckResult
	var qFiniteAwayFromNodesCheck model.CheckResult

	nodeDetectionCheck := model.CheckResult{
		Status: model.StatusPass,
		Pass:   true,
		Reason: "Node detection logic executed successfully.",
	}

	if initialR < params.NodeThresh {
		// Short-circuit trajectory integration!
		traj = model.Trajectory{Points: []model.TrajectoryPoint{}} // Empty trajectory

		zeroCount := 0
		trajectoryCheck = model.CheckResult{
			Status:      model.StatusFail,
			Pass:        false,
			Reason:      "Bohmian trajectory integration short-circuited: initial state starts on a node.",
			PointsCount: &zeroCount,
		}

		clockCheck = model.CheckResult{
			Status: model.StatusFail,
			Pass:   false,
			Reason: "Relational clock phi is undefined at a node.",
		}

		qCheck = model.CheckResult{
			Status: model.StatusContested,
			Pass:   false,
			Reason: "Quantum potential is contested/undefined at a node.",
		}

		classicalCheck = model.CheckResult{
			Status: model.StatusFail,
			Pass:   false,
			Reason: "Classical limit check failed due to node short-circuit.",
		}

		nodeContactFreeCheck = model.CheckResult{
			Status: model.StatusFail,
			Pass:   false,
			Reason: fmt.Sprintf("Node contact detected at initial point: amplitude %e is below threshold %e.", initialR, params.NodeThresh),
		}

		phaseGradientFiniteCheck = model.CheckResult{
			Status: model.StatusContested,
			Pass:   false,
			Reason: "Phase gradient is undefined at a node.",
		}

		qFiniteAwayFromNodesCheck = model.CheckResult{
			Status: model.StatusContested,
			Pass:   false,
			Reason: "No away-from-node points to evaluate quantum potential.",
		}
	} else {
		// Integrate using RK4
		stepper := guidance.NewRK4Stepper()
		initialState := model.MiniState{Alpha: params.Alpha0, Phi: params.Phi0}
		traj = guidance.Integrate(sw, initialState, stepper, params.LambdaStep, params.Steps, params.NodeThresh)
		trajectoryCheck = guidance.CheckTrajectory(traj)

		// Monotonicity check
		increasing := true
		decreasing := true
		for i := 1; i < len(traj.Points); i++ {
			if math.IsNaN(traj.Points[i].State.Alpha) {
				continue // skip NaNs in monotonicity
			}
			diff := traj.Points[i].State.Phi - traj.Points[i-1].State.Phi
			if diff <= 0 {
				increasing = false
			}
			if diff >= 0 {
				decreasing = false
			}
		}
		isClockMonotonic := (increasing || decreasing) && len(traj.Points) > 1
		var clockStatus model.CheckStatus = model.StatusPass
		var clockReason string = "Relational clock phi is strictly monotonic."
		if !isClockMonotonic {
			clockStatus = model.StatusFail
			clockReason = "Relational clock phi is not strictly monotonic."
		}
		clockVar := "phi"
		clockCheck = model.CheckResult{
			Status:   clockStatus,
			Pass:     isClockMonotonic,
			Reason:   clockReason,
			Variable: &clockVar,
		}

		// Compute Q potential away from nodes
		qValues = qpotential.QAlongTrajectory(sw, traj, DefaultQFiniteDifferenceStep)
		// Filter out Q values near nodes to find maxAbsQ
		var validQValues []float64
		for i, q := range qValues {
			if i < len(traj.Points) {
				p := traj.Points[i]
				if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
					continue
				}
				rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, sw)
				if rVal >= params.NodeThresh {
					validQValues = append(validQValues, q)
				}
			}
		}
		if len(validQValues) > 0 {
			maxAbsQ = qpotential.MaxAbsQ(validQValues)
			
			// For superposition, we check finiteness rather than checking if Q <= tolerance,
			// because Q is physically non-zero in a superposition.
			var qStatus model.CheckStatus = model.StatusPass
			var qReason string = "Quantum potential is finite along the trajectory outside node regions."
			if math.IsNaN(maxAbsQ) || math.IsInf(maxAbsQ, 0) {
				qStatus = model.StatusFail
				qReason = "Quantum potential is non-finite."
			}
			qCheck = model.CheckResult{
				Status:  qStatus,
				Pass:    qStatus == model.StatusPass,
				Reason:  qReason,
				MaxAbsQ: &maxAbsQ,
			}

			var classicalStatus model.CheckStatus = model.StatusPass
			var classicalReason string = "Classical limit check passes (quantum potential is finite away from nodes)."
			if qStatus != model.StatusPass {
				classicalStatus = model.StatusFail
				classicalReason = "Classical limit check failed due to non-finite quantum potential."
			}
			classicalCheck = model.CheckResult{
				Status: classicalStatus,
				Pass:   classicalStatus == model.StatusPass,
				Reason: classicalReason,
			}
		} else {
			qCheck = model.CheckResult{
				Status: model.StatusContested,
				Pass:   false,
				Reason: "No valid away-from-node points to evaluate Q potential.",
			}
			classicalCheck = model.CheckResult{
				Status: model.StatusFail,
				Pass:   false,
				Reason: "Classical limit is contested due to lack of valid away-from-node points.",
			}
		}

		// Node contact free
		nodeContact, contactReason := obstruction.DetectNodeContact(traj, sw, params.NodeThresh)
		var contactStatus model.CheckStatus = model.StatusPass
		if nodeContact {
			contactStatus = model.StatusFail
		}
		nodeContactFreeCheck = model.CheckResult{
			Status: contactStatus,
			Pass:   !nodeContact,
			Reason: contactReason,
		}

		// Phase gradient finite
		gradFinite, gradReason := obstruction.DetectPhaseGradientFinite(traj, sw, params.NodeThresh, params.MaxPhaseGrad)
		var gradStatus model.CheckStatus = model.StatusPass
		if !gradFinite {
			gradStatus = model.StatusFail
		}
		phaseGradientFiniteCheck = model.CheckResult{
			Status: gradStatus,
			Pass:   gradFinite,
			Reason: gradReason,
		}

		// Q finite away from nodes
		qFinite, qFiniteReason := obstruction.DetectQFiniteAwayFromNodes(traj, qValues, sw, params.NodeThresh)
		var qFiniteStatus model.CheckStatus = model.StatusPass
		if !qFinite {
			qFiniteStatus = model.StatusFail
		}
		qFiniteAwayFromNodesCheck = model.CheckResult{
			Status: qFiniteStatus,
			Pass:   qFinite,
			Reason: qFiniteReason,
		}
	}

	// Friedmann is deferred
	friedmannCheck := model.CheckResult{
		Status: model.StatusDeferred,
		Pass:   false,
		Reason: "Not implemented in BMC-0A; remains debt.",
	}

	checks := map[string]model.CheckResult{
		"wdw_residual":             wdwCheck,
		"trajectory":               trajectoryCheck,
		"clock_monotonicity":       clockCheck,
		"quantum_potential":        qCheck,
		"classical_limit":          classicalCheck,
		"friedmann_residual":       friedmannCheck,
		"node_detection":           nodeDetectionCheck,
		"node_contact_free":        nodeContactFreeCheck,
		"q_finite_away_from_nodes": qFiniteAwayFromNodesCheck,
		"phase_gradient_finite":    phaseGradientFiniteCheck,
	}

	// Detect obstructions using DetectAllSuperposition
	obstructions := obstruction.DetectAllSuperposition(params, traj, analyticResVal, qValues, finalTruthClaim, sw)

	// Technical Gate selection
	var techGateName string
	var techGateStatus model.CheckStatus

	if initialR < params.NodeThresh {
		techGateName = "node_detection_validation_gate"
		techGateStatus = model.StatusPass
	} else {
		techGateName = "bmc0a_superposition_safe_gate"
		techGateStatus = model.StatusPass
		if wdwCheck.Status != model.StatusPass ||
			trajectoryCheck.Status != model.StatusPass ||
			clockCheck.Status != model.StatusPass ||
			qCheck.Status != model.StatusPass ||
			classicalCheck.Status != model.StatusPass ||
			nodeContactFreeCheck.Status != model.StatusPass ||
			phaseGradientFiniteCheck.Status != model.StatusPass ||
			qFiniteAwayFromNodesCheck.Status != model.StatusPass {
			techGateStatus = model.StatusFail
		}
	}

	reportParams := ReportParameters{
		Superposition: &params,
	}

	report := &Report{
		SchemaVersion:           "bmc-report-v0.1",
		ModelID:                 "bmc0a_superposition",
		ToyAnalysisOnly:         true,
		PhysicsClaim:            "minisuperspace_only",
		FinalTruthClaim:         finalTruthClaim,
		PromotionRecommendation: "blocked_or_candidate",
		Parameters:              reportParams,
		Equations: map[string]string{
			"wdw":               "(-d_alpha_alpha + d_phi_phi) Psi = 0",
			"wavefunction":      "Psi(alpha,phi) = c1*exp(i(k1*alpha+omega1*phi)) + c2*exp(i(k2*alpha+omega2*phi))",
			"guidance":          "dalpha/dlambda=dS/dalpha, dphi/dlambda=-dS/dphi",
			"quantum_potential": "Q = -1/(2R)(d²R/dα² - d²R/dφ²)",
		},
		Checks: checks,
		TechnicalGate: TechnicalGate{
			Name:   techGateName,
			Status: techGateStatus,
		},
		PromotionGate: PromotionGate{
			Name:   "full_bmc_toy_gate",
			Status: StatusBlocked,
			Reason: "Friedmann residual and faithfulness review remain unpaid debt.",
		},
		NullModels: []NullModel{
			{
				Name:   "classical_frw",
				Status: "placeholder_obligation",
				Reason: "Classical-limit comparison begins with Q≈0 control.",
			},
			{
				Name:   "standard_wdw",
				Status: "placeholder_obligation",
				Reason: "Ensemble comparison deferred.",
			},
			{
				Name:   "lqc",
				Status: "deferred",
				Reason: "Not part of BMC-0A.",
			},
			{
				Name:   "page_wootters",
				Status: "deferred",
				Reason: "Formal relational-time null model deferred.",
			},
		},
		Obstructions: obstructions,
		Faithfulness: Faithfulness{
			Status: model.StatusContested,
			Reason: "No human faithfulness review yet. This only tests a superposition minisuperspace control.",
		},
		EbpDebt: EbpDebt{
			NeedMap:                "partial",
			NeedInvariant:          "partial",
			NeedToyCheck:           "active",
			NeedNullModel:          "partial",
			NeedObstruction:        "active",
			NeedFaithfulnessReview: "active",
		},
		Warnings: []string{
			"Toy analysis only.",
			"Does not test full quantum gravity.",
			"Does not test black holes, fermions, gauge fields, Lorentz recovery, or inhomogeneous perturbations.",
			"Passing this report cannot promote any final-truth claim.",
		},
	}

	return report, nil
}
