package residualaudit

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/residualrun"
)

type ResidualStabilityAudit struct {
	StabilityID           string   `json:"stability_id"`
	BranchID              string   `json:"branch_id"`
	StabilityComputed     bool     `json:"stability_computed"`
	PerturbationKind      string   `json:"perturbation_kind"`
	PerturbationMagnitude float64  `json:"perturbation_magnitude"`
	BaselineMetric        string   `json:"baseline_metric"`
	BaselineValue         *float64 `json:"baseline_value"`
	PerturbedValue        *float64 `json:"perturbed_value"`
	AbsoluteDelta         *float64 `json:"absolute_delta"`
	RelativeDelta         *float64 `json:"relative_delta"`
	StabilityStatus       string   `json:"stability_status"`
	Notes                 string   `json:"notes"`
}

func buildStabilityAudits(diags []residualrun.CandidateResidualDiagnostic) []ResidualStabilityAudit {
	audits := []ResidualStabilityAudit{}
	for _, diag := range diags {
		if !diag.ResidualComputed {
			continue
		}
		metrics := []string{"mean_abs_residual", "max_abs_residual", "rms_residual"}
		kinds := []string{PerturbAlphaPoint, PerturbPhiPoint, PerturbLambdaSpacing}
		for _, metric := range metrics {
			for _, kind := range kinds {
				audits = append(audits, computeStabilityAudit(diag, kind, 1e-6, metric))
			}
			if len(diag.ResidualInputPoints) >= 4 {
				audits = append(audits, computeStabilityAudit(diag, PerturbBranchSubset, 0, metric))
			} else {
				audits = append(audits, ResidualStabilityAudit{
					StabilityID:           "stability_" + diag.BranchID + "_" + metric + "_branch_subset_resampling",
					BranchID:              diag.BranchID,
					StabilityComputed:     false,
					PerturbationKind:      PerturbBranchSubset,
					PerturbationMagnitude: 0,
					BaselineMetric:        metric,
					StabilityStatus:       StabilityMissingInput,
					Notes:                 "Branch subset resampling requires at least four residual input points.",
				})
			}
		}
	}
	return audits
}

func computeStabilityAudit(diag residualrun.CandidateResidualDiagnostic, kind string, mag float64, metric string) ResidualStabilityAudit {
	id := "stability_" + diag.BranchID + "_" + metric + "_" + kind
	base, ok := recomputeMetricFromResidualInputPoints(diag.ResidualInputPoints, metric)
	if !ok {
		return ResidualStabilityAudit{
			StabilityID:           id,
			BranchID:              diag.BranchID,
			StabilityComputed:     false,
			PerturbationKind:      kind,
			PerturbationMagnitude: mag,
			BaselineMetric:        metric,
			StabilityStatus:       StabilityMissingInput,
			Notes:                 "Stability audit blocked by missing residual input points.",
		}
	}

	copied := copyResidualInputPoints(diag.ResidualInputPoints)
	switch kind {
	case PerturbAlphaPoint:
		ok = perturbAlphaPoint(copied, mag)
	case PerturbPhiPoint:
		ok = perturbPhiPoint(copied, mag)
	case PerturbLambdaSpacing:
		ok = perturbLambdaSpacing(copied, mag)
	case PerturbBranchSubset:
		copied, ok = subsetResidualInputPoints(copied)
	}
	if !ok {
		return ResidualStabilityAudit{StabilityID: id, BranchID: diag.BranchID, StabilityComputed: false, PerturbationKind: kind, PerturbationMagnitude: mag, BaselineMetric: metric, StabilityStatus: StabilityMissingInput, Notes: "Stability audit blocked by insufficient interval proxy inputs."}
	}
	perturbed, ok := recomputeMetricFromResidualInputPoints(copied, metric)
	if !ok {
		return ResidualStabilityAudit{StabilityID: id, BranchID: diag.BranchID, StabilityComputed: false, PerturbationKind: kind, PerturbationMagnitude: mag, BaselineMetric: metric, BaselineValue: &base, StabilityStatus: StabilityNonfinite, Notes: "Perturbed stability metric was nonfinite."}
	}
	if math.IsNaN(perturbed) || math.IsInf(perturbed, 0) {
		return ResidualStabilityAudit{StabilityID: id, BranchID: diag.BranchID, StabilityComputed: false, PerturbationKind: kind, PerturbationMagnitude: mag, BaselineMetric: metric, BaselineValue: &base, StabilityStatus: StabilityNonfinite, Notes: "Perturbed stability metric was nonfinite."}
	}

	absDelta := math.Abs(perturbed - base)
	var relDelta *float64
	if base != 0 {
		rel := absDelta / math.Abs(base)
		relDelta = &rel
	}
	status := StabilityStable
	if absDelta > 1e-12 {
		if relDelta == nil || *relDelta > 1e-6 {
			status = StabilitySensitive
		}
	}
	return ResidualStabilityAudit{
		StabilityID:           id,
		BranchID:              diag.BranchID,
		StabilityComputed:     true,
		PerturbationKind:      kind,
		PerturbationMagnitude: mag,
		BaselineMetric:        metric,
		BaselineValue:         &base,
		PerturbedValue:        &perturbed,
		AbsoluteDelta:         &absDelta,
		RelativeDelta:         relDelta,
		StabilityStatus:       status,
		Notes:                 "Deterministic candidate-only interval proxy stability diagnostic recomputed from copied residual input records.",
	}
}

func recomputeMetricFromResidualInputPoints(points []residualrun.ResidualInputPoint, metricName string) (float64, bool) {
	if len(points) == 0 {
		return 0, false
	}
	sumAbs := 0.0
	sumSq := 0.0
	maxAbs := 0.0
	for _, pt := range points {
		if pt.CandidateLeftHandSide == nil || pt.CandidateRightHandSide == nil {
			return 0, false
		}
		res := *pt.CandidateLeftHandSide - *pt.CandidateRightHandSide
		if math.IsNaN(res) || math.IsInf(res, 0) {
			return 0, false
		}
		abs := math.Abs(res)
		sumAbs += abs
		sumSq += res * res
		if abs > maxAbs {
			maxAbs = abs
		}
	}
	switch metricName {
	case "mean_abs_residual":
		return sumAbs / float64(len(points)), true
	case "max_abs_residual":
		return maxAbs, true
	case "rms_residual":
		return math.Sqrt(sumSq / float64(len(points))), true
	default:
		return 0, false
	}
}

func copyResidualInputPoints(points []residualrun.ResidualInputPoint) []residualrun.ResidualInputPoint {
	copied := make([]residualrun.ResidualInputPoint, len(points))
	for idx, pt := range points {
		copied[idx] = pt
		if pt.Lambda != nil {
			v := *pt.Lambda
			copied[idx].Lambda = &v
		}
		if pt.Alpha != nil {
			v := *pt.Alpha
			copied[idx].Alpha = &v
		}
		if pt.Phi != nil {
			v := *pt.Phi
			copied[idx].Phi = &v
		}
		if pt.CandidateLeftHandSide != nil {
			v := *pt.CandidateLeftHandSide
			copied[idx].CandidateLeftHandSide = &v
		}
		if pt.CandidateRightHandSide != nil {
			v := *pt.CandidateRightHandSide
			copied[idx].CandidateRightHandSide = &v
		}
	}
	return copied
}

func perturbAlphaPoint(points []residualrun.ResidualInputPoint, magnitude float64) bool {
	for idx := range points {
		if points[idx].Alpha == nil || points[idx].CandidateLeftHandSide == nil {
			return false
		}
		if math.IsNaN(*points[idx].Alpha) || math.IsInf(*points[idx].Alpha, 0) ||
			math.IsNaN(*points[idx].CandidateLeftHandSide) || math.IsInf(*points[idx].CandidateLeftHandSide, 0) {
			return false
		}
		*points[idx].Alpha += magnitude
		velocityProxy := math.Sqrt(math.Abs(*points[idx].CandidateLeftHandSide)) + magnitude
		lhs := velocityProxy * velocityProxy
		points[idx].CandidateLeftHandSide = &lhs
		return true
	}
	return false
}

func perturbPhiPoint(points []residualrun.ResidualInputPoint, magnitude float64) bool {
	for idx := range points {
		if points[idx].Phi == nil || points[idx].CandidateRightHandSide == nil {
			return false
		}
		if math.IsNaN(*points[idx].Phi) || math.IsInf(*points[idx].Phi, 0) ||
			math.IsNaN(*points[idx].CandidateRightHandSide) || math.IsInf(*points[idx].CandidateRightHandSide, 0) {
			return false
		}
		*points[idx].Phi += magnitude
		velocityProxy := math.Sqrt(math.Abs(*points[idx].CandidateRightHandSide)) + magnitude
		rhs := velocityProxy * velocityProxy
		points[idx].CandidateRightHandSide = &rhs
		return true
	}
	return false
}

func perturbLambdaSpacing(points []residualrun.ResidualInputPoint, magnitude float64) bool {
	if len(points) == 0 || magnitude < 0 {
		return false
	}
	scale := 1 + magnitude
	if scale <= 0 || math.IsNaN(scale) || math.IsInf(scale, 0) {
		return false
	}
	factor := 1 / (scale * scale)
	for idx := range points {
		if points[idx].CandidateLeftHandSide == nil || points[idx].CandidateRightHandSide == nil {
			return false
		}
		if math.IsNaN(*points[idx].CandidateLeftHandSide) || math.IsInf(*points[idx].CandidateLeftHandSide, 0) ||
			math.IsNaN(*points[idx].CandidateRightHandSide) || math.IsInf(*points[idx].CandidateRightHandSide, 0) {
			return false
		}
		lhs := *points[idx].CandidateLeftHandSide * factor
		rhs := *points[idx].CandidateRightHandSide * factor
		points[idx].CandidateLeftHandSide = &lhs
		points[idx].CandidateRightHandSide = &rhs
	}
	return true
}

func subsetResidualInputPoints(points []residualrun.ResidualInputPoint) ([]residualrun.ResidualInputPoint, bool) {
	if len(points) < 4 {
		return nil, false
	}
	subset := []residualrun.ResidualInputPoint{}
	for idx, pt := range points {
		if idx%2 == 0 {
			subset = append(subset, pt)
		}
	}
	return subset, true
}
