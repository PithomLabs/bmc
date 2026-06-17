package nullspec

import (
	"fmt"
	"strings"

	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// ValidateNullModelSpecReport performs all strict EBP boundary checks and schema validations.
func ValidateNullModelSpecReport(r *NullModelSpecReport) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// 1. Schema, kind, and scope
	if r.SchemaVersion != "bmc0a-nullmodel-spec-v0.1" {
		fail("schema_version", fmt.Sprintf("unsupported schema version: %s (expected 'bmc0a-nullmodel-spec-v0.1')", r.SchemaVersion))
	}
	if r.SpecKind != "null_model_scaffold" {
		fail("spec_kind", fmt.Sprintf("unsupported spec_kind: %s (expected 'null_model_scaffold')", r.SpecKind))
	}
	if r.SpecScope != SpecScopeNullModelOnly {
		fail("spec_scope", fmt.Sprintf("unsupported spec_scope: %s (expected '%s')", r.SpecScope, SpecScopeNullModelOnly))
	}

	// 2. Toy analysis & final truth claim
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}

	// 3. Hard safety invariants
	if r.ResidualComputed {
		fail("residual_computed", "residual_computed must be false (hard invariant)")
	}
	if r.NullComparisonComputed {
		fail("null_comparison_computed", "null_comparison_computed must be false (hard invariant)")
	}
	if r.FriedmannRecoveryClaim {
		fail("friedmann_recovery_claim", "friedmann_recovery_claim must be false (hard invariant)")
	}

	// 4. Null Model registry checks
	requiredNullModels := map[string]bool{
		"constant_phase_control":                             true,
		"randomized_phase_control":                           true,
		"matched_amplitude_randomized_phase_control":         true,
		"classical_frw_reference_trajectory":                 true,
		"same_branch_segmentation_under_null_wavefunctions":  true,
		"node_neighborhood_stress_case":                      true,
		"clock_choice_alternative_branch_diagnostic":         true,
	}

	registeredCount := make(map[string]int)
	for idx, nm := range r.NullModels {
		prefix := fmt.Sprintf("null_models[%d]", idx)
		registeredCount[nm.NullModelID]++

		// Status checks
		if nm.Status != StatusPlanned && nm.Status != StatusDeferred && nm.Status != StatusBlocked {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s (must be planned, deferred, or blocked)", nm.Status))
		}
		if nm.Status == "passed" || nm.Status == "failed" || nm.Status == "validated" || nm.Status == "recovered" || nm.Status == "proved" {
			fail(prefix+".status", "null model status cannot be passed, failed, validated, recovered, or proved")
		}

		// Required flag check
		if !nm.RequiredBeforeResidualPromotion {
			fail(prefix+".required_before_residual_promotion", "required_before_residual_promotion must be true")
		}

		// Reject unknown/extra null models
		if !requiredNullModels[nm.NullModelID] {
			fail(prefix+".null_model_id", fmt.Sprintf("unknown/extra null model ID: %s", nm.NullModelID))
		}
	}

	for id := range requiredNullModels {
		count := registeredCount[id]
		if count == 0 {
			fail("null_models", fmt.Sprintf("missing required null model: %s", id))
		} else if count > 1 {
			fail("null_models", fmt.Sprintf("null model %s registered %d times (expected exactly once)", id, count))
		}
	}

	// 5. Input requirements check
	if len(r.InputRequirements) == 0 {
		fail("input_requirements", "input_requirements cannot be empty")
	}
	for idx, ir := range r.InputRequirements {
		prefix := fmt.Sprintf("input_requirements[%d]", idx)
		if ir.AvailabilityStatus != AvailabilityAvailable &&
			ir.AvailabilityStatus != AvailabilityPlanned &&
			ir.AvailabilityStatus != AvailabilityDeferred &&
			ir.AvailabilityStatus != AvailabilityBlocked {
			fail(prefix+".availability_status", fmt.Sprintf("invalid status: %s", ir.AvailabilityStatus))
		}
	}

	// 6. Metric contracts check
	if len(r.MetricContracts) == 0 {
		fail("metric_contracts", "metric_contracts cannot be empty")
	}
	for idx, mc := range r.MetricContracts {
		prefix := fmt.Sprintf("metric_contracts[%d]", idx)
		if mc.Status != StatusPlanned && mc.Status != StatusDeferred && mc.Status != StatusBlocked {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", mc.Status))
		}
		if !mc.RequiredBeforeResidualPromotion {
			fail(prefix+".required_before_residual_promotion", "required_before_residual_promotion must be true for all metric contracts")
		}
	}

	// 7. Comparison contracts check
	if len(r.FutureComparisonContracts) == 0 {
		fail("future_comparison_contracts", "future_comparison_contracts cannot be empty")
	}
	for idx, cc := range r.FutureComparisonContracts {
		prefix := fmt.Sprintf("future_comparison_contracts[%d]", idx)
		if cc.ComparisonComputed {
			fail(prefix+".comparison_computed", "comparison_computed must be false")
		}
		if cc.Status != StatusPlanned && cc.Status != StatusDeferred && cc.Status != StatusBlocked {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", cc.Status))
		}
	}

	// 8. Gates validation (exact gate cardinality and status)
	requiredGateNames := []string{
		"toy_analysis_only_gate",
		"no_final_truth_claim_gate",
		"no_residual_computation_gate",
		"no_null_comparison_result_gate",
		"null_model_registry_complete_gate",
		"required_before_residual_promotion_gate",
		"friedmann_recovery_claim_blocked_gate",
		"full_bmc_blocked_gate",
		"clock_choice_debt_active_gate",
		"faithfulness_contested_gate",
	}

	gateCounts := make(map[string]int)
	for idx, g := range r.Gates {
		prefix := fmt.Sprintf("gates[%d]", idx)
		if g.Status != "pass" {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s (must be pass)", g.Status))
		}
		gateCounts[g.Name]++
	}

	for _, name := range requiredGateNames {
		count := gateCounts[name]
		if count == 0 {
			fail("gates", fmt.Sprintf("missing required gate: %s", name))
		} else if count > 1 {
			fail("gates", fmt.Sprintf("gate %s appears %d times (expected exactly once)", name, count))
		}
	}

	// 9. EBP debt check
	if r.PromotionGate.Status != report.StatusBlocked {
		fail("promotion_gate.status", "promotion gate status must remain blocked")
	}
	if r.EbpDebt.ClockChoiceDebt != "active" {
		fail("ebp_debt.clock_choice_debt", "clock_choice_debt must remain active")
	}
	if r.EbpDebt.ClassicalTargetDebt != "active" {
		fail("ebp_debt.classical_target_debt", "classical_target_debt must remain active")
	}
	if r.EbpDebt.UnitConventionDebt != "active" {
		fail("ebp_debt.unit_convention_debt", "unit_convention_debt must remain active")
	}
	if r.EbpDebt.SignConventionDebt != "active" {
		fail("ebp_debt.sign_convention_debt", "sign_convention_debt must remain active")
	}
	if r.EbpDebt.NormalizationDebt != "active" {
		fail("ebp_debt.normalization_debt", "normalization_debt must remain active")
	}
	if r.EbpDebt.NeedNullModel != "active" {
		fail("ebp_debt.needNullModel", "needNullModel must remain active")
	}
	if r.EbpDebt.NeedFaithfulnessReview != "contested" {
		fail("ebp_debt.needFaithfulnessReview", "needFaithfulnessReview must remain contested")
	}
	if r.EbpDebt.PromotionStatus != "planned_candidate_only" {
		fail("ebp_debt.promotion_status", "promotion_status must be planned_candidate_only")
	}

	// 10. Warnings validation
	hasNoComputeWarn := false
	hasNoNullResultWarn := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No Friedmann residual was computed") {
			hasNoComputeWarn = true
		}
		if strings.Contains(w, "No null-model comparison result was computed") {
			hasNoNullResultWarn = true
		}
	}
	if !hasNoComputeWarn {
		fail("warnings", "must include warning containing: 'No Friedmann residual was computed'")
	}
	if !hasNoNullResultWarn {
		fail("warnings", "must include warning containing: 'No null-model comparison result was computed'")
	}

	return errs
}
