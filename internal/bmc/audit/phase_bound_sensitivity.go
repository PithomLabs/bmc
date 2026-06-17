package audit

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// RunPhaseBoundSweep runs the safe profile with bounds: 25, 50, 100, 200.
// Emits the results in fixed ascending order of bounds.
func RunPhaseBoundSweep(params model.SuperpositionParams) ([]PhaseGradientBoundResult, error) {
	bounds := []float64{25, 50, 100, 200}
	results := make([]PhaseGradientBoundResult, len(bounds))

	sw := wave.NewSuperpositionWave(
		params.C1Real, params.C1Imag, params.K1, params.Omega1,
		params.C2Real, params.C2Imag, params.K2, params.Omega2,
	)

	for i, bound := range bounds {
		p := params
		p.MaxPhaseGrad = bound

		rep, err := report.GenerateSuperposition(p, false)
		if err != nil {
			return nil, err
		}

		// Re-run the safe profile trajectory to find max observed phase gradient magnitude
		stepper := guidance.NewRK4Stepper()
		initialState := model.MiniState{Alpha: p.Alpha0, Phi: p.Phi0}
		traj := guidance.Integrate(sw, initialState, stepper, p.LambdaStep, p.Steps, p.NodeThresh)

		maxObs := 0.0
		hasPoints := false
		for _, pt := range traj.Points {
			if math.IsNaN(pt.State.Alpha) || math.IsNaN(pt.State.Phi) {
				continue
			}
			r := wave.AmplitudeField(pt.State.Alpha, pt.State.Phi, sw)
			if r < p.NodeThresh {
				continue
			}
			dSa, dSp := wave.PhaseGradient(pt.State.Alpha, pt.State.Phi, sw)
			mag := math.Sqrt(dSa*dSa + dSp*dSp)
			if mag > maxObs {
				maxObs = mag
			}
			hasPoints = true
		}

		var maxObsPtr *float64
		var maxObsStatus, maxObsReason string
		if hasPoints {
			maxObsPtr = &maxObs
		} else {
			maxObsStatus = "not_applicable"
			maxObsReason = "no away-from-node points available"
		}

		isBinding := maxObs > bound

		results[i] = PhaseGradientBoundResult{
			Bound:                      bound,
			MaxObservedPhaseGrad:       maxObsPtr,
			MaxObservedPhaseGradStatus: maxObsStatus,
			MaxObservedPhaseGradReason: maxObsReason,
			IsBinding:                  isBinding,
			PhaseGradientFinite:        rep.Checks["phase_gradient_finite"].Status,
		}
	}

	return results, nil
}
