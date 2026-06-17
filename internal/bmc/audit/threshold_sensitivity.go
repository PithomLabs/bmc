package audit

import (
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// RunThresholdSweep runs safe and node-probe profiles with thresholds: 1e-4, 1e-5, 1e-6.
// The array is returned in a deterministic order:
// Sorted alphabetically by profile, then by threshold descending.
func RunThresholdSweep(safeParams, nodeProbeParams model.SuperpositionParams) ([]ThresholdSensitivityResult, error) {
	thresholds := []float64{1e-4, 1e-5, 1e-6}
	profiles := []struct {
		name   string
		params model.SuperpositionParams
	}{
		{"node-probe", nodeProbeParams},
		{"safe", safeParams},
	}

	results := make([]ThresholdSensitivityResult, 0, len(profiles)*len(thresholds))

	for _, prof := range profiles {
		for _, thresh := range thresholds {
			p := prof.params
			p.NodeThresh = thresh

			rep, err := report.GenerateSuperposition(p, false)
			if err != nil {
				return nil, err
			}

			results = append(results, ThresholdSensitivityResult{
				Threshold:            thresh,
				Profile:              prof.name,
				NodeContactFree:      rep.Checks["node_contact_free"].Status,
				QFiniteAwayFromNodes: rep.Checks["q_finite_away_from_nodes"].Status,
				PhaseGradientFinite:  rep.Checks["phase_gradient_finite"].Status,
				TechnicalGateName:    rep.TechnicalGate.Name,
				TechnicalGateStatus:  rep.TechnicalGate.Status,
			})
		}
	}

	return results, nil
}
