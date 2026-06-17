package clockseg

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// ComputeClockIndependentDiagnostics computes physical diagnostics of the trajectory
// that do not depend on global monotonicity or choice of relational clock phi.
func ComputeClockIndependentDiagnostics(
	traj model.Trajectory,
	wf wave.WaveFunction,
	nodeThresh float64,
	segments []ClockSegment,
	turningPoints []ClockTurningPoint,
) ClockIndependentDiagnostic {
	var pathLength float64
	var minAmplitudeR = math.MaxFloat64
	var maxAbsQ = 0.0
	var maxPhaseGrad = 0.0
	var validPointsCount = 0
	var finiteness = true
	var nodeContactFree = true

	var firstValidLambda = -1.0
	var lastValidLambda = -1.0

	for i, p := range traj.Points {
		isFinitePoint := !math.IsNaN(p.State.Alpha) && !math.IsInf(p.State.Alpha, 0) &&
			!math.IsNaN(p.State.Phi) && !math.IsInf(p.State.Phi, 0) &&
			!math.IsNaN(p.Lambda) && !math.IsInf(p.Lambda, 0)

		if !isFinitePoint {
			finiteness = false
			continue
		}

		validPointsCount++
		if firstValidLambda < 0 {
			firstValidLambda = p.Lambda
		}
		lastValidLambda = p.Lambda

		if i > 0 {
			pPrev := traj.Points[i-1]
			isPrevFinite := !math.IsNaN(pPrev.State.Alpha) && !math.IsInf(pPrev.State.Alpha, 0) &&
				!math.IsNaN(pPrev.State.Phi) && !math.IsInf(pPrev.State.Phi, 0)
			if isPrevFinite {
				dAlpha := p.State.Alpha - pPrev.State.Alpha
				dPhi := p.State.Phi - pPrev.State.Phi
				pathLength += math.Sqrt(dAlpha*dAlpha + dPhi*dPhi)
			}
		}

		rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, wf)
		if rVal < minAmplitudeR {
			minAmplitudeR = rVal
		}

		if rVal < nodeThresh {
			nodeContactFree = false
		} else {
			qVal := qpotential.Q(wf, p.State.Alpha, p.State.Phi, 1e-4)
			if math.Abs(qVal) > maxAbsQ && !math.IsNaN(qVal) && !math.IsInf(qVal, 0) {
				maxAbsQ = math.Abs(qVal)
			}

			dSdAlpha, dSdPhi := wave.PhaseGradient(p.State.Alpha, p.State.Phi, wf)
			phaseGrad := math.Sqrt(dSdAlpha*dSdAlpha + dSdPhi*dSdPhi)
			if phaseGrad > maxPhaseGrad && !math.IsNaN(phaseGrad) && !math.IsInf(phaseGrad, 0) {
				maxPhaseGrad = phaseGrad
			}
		}
	}

	var totalLambdaInterval float64
	if firstValidLambda >= 0 && lastValidLambda >= 0 {
		totalLambdaInterval = lastValidLambda - firstValidLambda
	}

	if minAmplitudeR == math.MaxFloat64 {
		minAmplitudeR = 0.0
	}

	return ClockIndependentDiagnostic{
		PathLengthInConfigurationSpace: pathLength,
		TotalLambdaInterval:            totalLambdaInterval,
		NumValidTrajectoryPoints:       validPointsCount,
		NumClockSegments:               len(segments),
		NumTurningPoints:               len(turningPoints),
		MinAmplitudeR:                  minAmplitudeR,
		MaxAbsQAwayFromNodes:           maxAbsQ,
		MaxPhaseGradient:               maxPhaseGrad,
		NodeContactFree:                nodeContactFree,
		TrajectoryFiniteness:           finiteness,
	}
}
