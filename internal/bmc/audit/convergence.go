package audit

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/obstruction"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

// RunStepSizeSweep runs the safe profile with step sizes: 0.1, 0.05, 0.025, 0.0125
// and corresponding steps: 100, 200, 400, 800 to keep the total interval fixed at T=10.0.
func RunStepSizeSweep(params model.SuperpositionParams) ([]StepSizeResult, error) {
	stepSizes := []float64{0.1, 0.05, 0.025, 0.0125}
	stepCounts := []int{100, 200, 400, 800}

	results := make([]StepSizeResult, len(stepSizes))
	trajectories := make([]model.Trajectory, len(stepSizes))

	sw := wave.NewSuperpositionWave(
		params.C1Real, params.C1Imag, params.K1, params.Omega1,
		params.C2Real, params.C2Imag, params.K2, params.Omega2,
	)

	// First, integrate all trajectories
	for i, dt := range stepSizes {
		steps := stepCounts[i]
		stepper := guidance.NewRK4Stepper()
		initialState := model.MiniState{Alpha: params.Alpha0, Phi: params.Phi0}
		trajectories[i] = guidance.Integrate(sw, initialState, stepper, dt, steps, params.NodeThresh)
	}

	// Finest trajectory is the last one (index 3: dt = 0.0125)
	finestTraj := trajectories[len(trajectories)-1]
	var finestFinalAlpha, finestFinalPhi float64
	if len(finestTraj.Points) > 0 {
		finestFinalPoint := finestTraj.Points[len(finestTraj.Points)-1]
		finestFinalAlpha = finestFinalPoint.State.Alpha
		finestFinalPhi = finestFinalPoint.State.Phi
	}

	for i, dt := range stepSizes {
		steps := stepCounts[i]
		traj := trajectories[i]

		// Final state of current trajectory
		var finalAlpha, finalPhi float64
		if len(traj.Points) > 0 {
			finalPoint := traj.Points[len(traj.Points)-1]
			finalAlpha = finalPoint.State.Alpha
			finalPhi = finalPoint.State.Phi
		}

		var driftAlphaPtr, driftPhiPtr *float64
		var driftAlphaStatus, driftAlphaReason string
		var driftPhiStatus, driftPhiReason string

		if len(traj.Points) > 0 && len(finestTraj.Points) > 0 {
			da := finalAlpha - finestFinalAlpha
			dp := finalPhi - finestFinalPhi
			driftAlphaPtr = &da
			driftPhiPtr = &dp
		} else {
			driftAlphaStatus = "not_applicable"
			driftAlphaReason = "no trajectory points available for drift calculation"
			driftPhiStatus = "not_applicable"
			driftPhiReason = "no trajectory points available for drift calculation"
		}

		// Calculate QAlongTrajectory
		qValues := qpotential.QAlongTrajectory(sw, traj, report.DefaultQFiniteDifferenceStep)

		// Filter Q values near nodes to find maxAbsQ
		var validQValues []float64
		for j, q := range qValues {
			if j < len(traj.Points) {
				p := traj.Points[j]
				if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
					continue
				}
				rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, sw)
				if rVal >= params.NodeThresh {
					validQValues = append(validQValues, q)
				}
			}
		}
		var maxAbsQPtr *float64
		var maxAbsQStatus, maxAbsQReason string
		if len(validQValues) > 0 {
			mq := qpotential.MaxAbsQ(validQValues)
			maxAbsQPtr = &mq
		} else {
			maxAbsQStatus = "not_applicable"
			maxAbsQReason = "no away-from-node points available"
		}

		// Minimum amplitude
		minR := math.MaxFloat64
		for _, p := range traj.Points {
			if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
				continue
			}
			r := wave.AmplitudeField(p.State.Alpha, p.State.Phi, sw)
			if r < minR {
				minR = r
			}
		}
		var minRPtr *float64
		var minRStatus, minRReason string
		if minR != math.MaxFloat64 {
			minRPtr = &minR
		} else {
			minRStatus = "not_applicable"
			minRReason = "no valid trajectory points available"
		}

		// Maximum phase gradient magnitude
		maxGradMag := 0.0
		hasAwayFromNode := false
		for _, p := range traj.Points {
			if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
				continue
			}
			r := wave.AmplitudeField(p.State.Alpha, p.State.Phi, sw)
			if r < params.NodeThresh {
				continue
			}
			dSa, dSp := wave.PhaseGradient(p.State.Alpha, p.State.Phi, sw)
			mag := math.Sqrt(dSa*dSa + dSp*dSp)
			if mag > maxGradMag {
				maxGradMag = mag
			}
			hasAwayFromNode = true
		}
		var maxGradMagPtr *float64
		var maxGradMagStatus, maxGradMagReason string
		if hasAwayFromNode {
			maxGradMagPtr = &maxGradMag
		} else {
			maxGradMagStatus = "not_applicable"
			maxGradMagReason = "no away-from-node points available"
		}

		// Clock Monotonicity check
		increasing := true
		decreasing := true
		for j := 1; j < len(traj.Points); j++ {
			if math.IsNaN(traj.Points[j].State.Alpha) {
				continue
			}
			diff := traj.Points[j].State.Phi - traj.Points[j-1].State.Phi
			if diff <= 0 {
				increasing = false
			}
			if diff >= 0 {
				decreasing = false
			}
		}
		isClockMonotonic := (increasing || decreasing) && len(traj.Points) > 1

		// Evaluate safe superposition technical gate checks
		// (without Friedmann which is deferred)
		compRes := wdw.ComponentResidualsSuperposition(params.K1, params.Omega1, params.K2, params.Omega2)
		allCompPass := true
		for _, c := range compRes {
			if math.Abs(c.Residual) > params.Tolerance {
				allCompPass = false
			}
		}

		trajectoryCheck := guidance.CheckTrajectory(traj)
		nodeContact, _ := obstruction.DetectNodeContact(traj, sw, params.NodeThresh)
		gradientFinite, _ := obstruction.DetectPhaseGradientFinite(traj, sw, params.NodeThresh, params.MaxPhaseGrad)
		qFinite, _ := obstruction.DetectQFiniteAwayFromNodes(traj, qValues, sw, params.NodeThresh)

		pass := allCompPass &&
			trajectoryCheck.Status == model.StatusPass &&
			isClockMonotonic &&
			(maxAbsQPtr == nil || (!math.IsNaN(*maxAbsQPtr) && !math.IsInf(*maxAbsQPtr, 0))) &&
			!nodeContact &&
			gradientFinite &&
			qFinite

		var techStatus model.CheckStatus = model.StatusPass
		if !pass {
			techStatus = model.StatusFail
		}

		results[i] = StepSizeResult{
			StepSize:                 dt,
			Steps:                    steps,
			EndpointDriftAlpha:       driftAlphaPtr,
			EndpointDriftAlphaStatus: driftAlphaStatus,
			EndpointDriftAlphaReason: driftAlphaReason,
			EndpointDriftPhi:         driftPhiPtr,
			EndpointDriftPhiStatus:   driftPhiStatus,
			EndpointDriftPhiReason:   driftPhiReason,
			MaxAbsQ:                  maxAbsQPtr,
			MaxAbsQStatus:            maxAbsQStatus,
			MaxAbsQReason:            maxAbsQReason,
			MinAmplitudeR:            minRPtr,
			MinAmplitudeRStatus:      minRStatus,
			MinAmplitudeRReason:      minRReason,
			MaxPhaseGrad:             maxGradMagPtr,
			MaxPhaseGradStatus:       maxGradMagStatus,
			MaxPhaseGradReason:       maxGradMagReason,
			ClockMonotonic:           isClockMonotonic,
			TechnicalGateStatus:      techStatus,
		}
	}

	return results, nil
}
