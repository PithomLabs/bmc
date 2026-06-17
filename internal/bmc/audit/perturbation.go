package audit

import (
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// RunPerturbationSweep runs a grid of parameters around the safe profile:
// c2_real = [0.45, 0.50, 0.55], k2 = [1.9, 2.0, 2.1] with omega2 = -k2.
// Nested-loop order: outer loop c2, inner loop k2.
func RunPerturbationSweep(safeParams model.SuperpositionParams) ([]ParameterPerturbationResult, error) {
	c2Vals := []float64{0.45, 0.50, 0.55}
	k2Vals := []float64{1.9, 2.0, 2.1}

	results := make([]ParameterPerturbationResult, 0, len(c2Vals)*len(k2Vals))

	for _, c2Real := range c2Vals {
		for _, k2 := range k2Vals {
			p := safeParams
			p.C2Real = c2Real
			p.K2 = k2
			p.Omega2 = -k2

			rep, err := report.GenerateSuperposition(p, false)
			if err != nil {
				return nil, err
			}

			// We need to check if node contact appeared
			nodeContact := rep.Checks["node_contact_free"].Status != model.StatusPass

			results = append(results, ParameterPerturbationResult{
				C2Real:               c2Real,
				K2:                   k2,
				Omega2:               -k2,
				NodeContact:          nodeContact,
				QFiniteAwayFromNodes: rep.Checks["q_finite_away_from_nodes"].Status,
				ClockMonotonic:       rep.Checks["clock_monotonicity"].Status,
				SafeGateStatus:       rep.TechnicalGate.Status,
			})
		}
	}

	return results, nil
}

// RunNodeProbeOffsetSweep runs small offsets around the exact node:
// (0, 0), (1e-8, 0), (1e-6, 0), (1e-4, 0).
func RunNodeProbeOffsetSweep(nodeProbeParams model.SuperpositionParams) ([]NodeProbeOffsetResult, error) {
	offsets := []struct {
		alpha float64
		phi   float64
	}{
		{0.0, 0.0},
		{1e-8, 0.0},
		{1e-6, 0.0},
		{1e-4, 0.0},
	}

	results := make([]NodeProbeOffsetResult, len(offsets))

	sw := wave.NewSuperpositionWave(
		nodeProbeParams.C1Real, nodeProbeParams.C1Imag, nodeProbeParams.K1, nodeProbeParams.Omega1,
		nodeProbeParams.C2Real, nodeProbeParams.C2Imag, nodeProbeParams.K2, nodeProbeParams.Omega2,
	)

	for i, off := range offsets {
		p := nodeProbeParams
		p.Alpha0 = off.alpha
		p.Phi0 = off.phi

		rep, err := report.GenerateSuperposition(p, false)
		if err != nil {
			return nil, err
		}

		initialR := wave.AmplitudeField(p.Alpha0, p.Phi0, sw)
		shortCircuit := initialR < p.NodeThresh
		integrated := !shortCircuit

		results[i] = NodeProbeOffsetResult{
			OffsetAlpha:           off.alpha,
			OffsetPhi:             off.phi,
			InitialAmplitude:      &initialR,
			ShortCircuitTriggered: shortCircuit,
			Integrated:            integrated,
			NodeContactFree:       rep.Checks["node_contact_free"].Status,
			PhaseGradientFinite:   rep.Checks["phase_gradient_finite"].Status,
			QFiniteAwayFromNodes:  rep.Checks["q_finite_away_from_nodes"].Status,
		}
	}

	return results, nil
}
