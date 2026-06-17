package friedmannspec

import (
	"fmt"
	"math"
	"strings"

	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// ValidateFriedmannSpecReport performs all strict EBP boundary checks and schema validations.
func ValidateFriedmannSpecReport(r *FriedmannSpecReport) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// 1. Schema and kind
	if r.SchemaVersion != "bmc0a-friedmann-spec-v0.1" {
		fail("schema_version", fmt.Sprintf("unsupported schema version: %s (expected 'bmc0a-friedmann-spec-v0.1')", r.SchemaVersion))
	}
	if r.SpecKind != "friedmann_residual_specification" {
		fail("spec_kind", fmt.Sprintf("unsupported spec_kind: %s (expected 'friedmann_residual_specification')", r.SpecKind))
	}

	// 2. Toy analysis & final truth claim
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}

	// 3. Spec scope allowed values
	validScopes := map[string]bool{
		SpecScopeCandidateOnly: true,
		SpecScopeBlocked:       true,
		SpecScopeContested:     true,
	}
	if !validScopes[r.SpecScope] {
		fail("spec_scope", fmt.Sprintf("invalid spec_scope: %s", r.SpecScope))
	}

	// 4. Hard safety invariants (no computation or recovery claims)
	if r.ResidualComputed {
		fail("residual_computed", "residual_computed must be false (hard invariant)")
	}
	if r.FriedmannRecoveryClaim {
		fail("friedmann_recovery_claim", "friedmann_recovery_claim must be false (hard invariant)")
	}

	// 5. Candidate maps validation
	for idx, m := range r.CandidateMaps {
		prefix := fmt.Sprintf("candidate_maps[%d]", idx)
		if m.Status != StatusCandidateOnly && m.Status != StatusBlocked && m.Status != StatusContested {
			fail(prefix+".status", fmt.Sprintf("invalid candidate map status: %s", m.Status))
		}
		if m.Status == "validated" || m.Status == "recovered" || m.Status == "proved" {
			fail(prefix+".status", "candidate map status cannot be validated, recovered, or proved")
		}
		// Validate ClassicalTargetDebt in map
		if m.ClassicalTargetDebt != "active" && m.ClassicalTargetDebt != "contested" && m.ClassicalTargetDebt != "blocked" {
			fail(prefix+".classical_target_debt", fmt.Sprintf("invalid classical target debt status: %q (must be active, contested, or blocked)", m.ClassicalTargetDebt))
		}
	}

	// 6. Branch requirements validation
	for idx, b := range r.BranchRequirements {
		prefix := fmt.Sprintf("branch_requirements[%d]", idx)
		if b.BranchResidualReadiness != StatusCandidateOnly && b.BranchResidualReadiness != StatusBlocked && b.BranchResidualReadiness != StatusContested {
			fail(prefix+".branch_residual_readiness", fmt.Sprintf("invalid branch readiness status: %s", b.BranchResidualReadiness))
		}
		if b.BranchResidualReadiness == "ready" || b.BranchResidualReadiness == "pass" || b.BranchResidualReadiness == "recovered" {
			fail(prefix+".branch_residual_readiness", "branch readiness status cannot be ready, pass, or recovered")
		}
		if math.IsNaN(b.ClockRange) || math.IsInf(b.ClockRange, 0) || b.ClockRange < 0 {
			fail(prefix+".clock_range", "clock_range must be finite and nonnegative")
		}
		if math.IsNaN(b.LambdaRange) || math.IsInf(b.LambdaRange, 0) || b.LambdaRange < 0 {
			fail(prefix+".lambda_range", "lambda_range must be finite and nonnegative")
		}
	}

	// 7. Derivative checks validation
	for idx, d := range r.DerivativeReadinessChecks {
		prefix := fmt.Sprintf("derivative_readiness_checks[%d]", idx)
		if d.Status != StatusCandidateOnly && d.Status != StatusBlocked && d.Status != StatusContested {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", d.Status))
		}
	}

	// 8. Formula candidate validation
	for idx, f := range r.ResidualFormulaCandidates {
		prefix := fmt.Sprintf("residual_formula_candidates[%d]", idx)
		if f.Status != StatusCandidateOnly && f.Status != StatusBlocked && f.Status != StatusContested {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", f.Status))
		}
	}

	// 9. Null model requirements cannot be empty
	if len(r.NullModelRequirements) == 0 {
		fail("null_model_requirements", "null_model_requirements cannot be empty")
	}
	for idx, n := range r.NullModelRequirements {
		prefix := fmt.Sprintf("null_model_requirements[%d]", idx)
		if n.Status != "planned" && n.Status != "deferred" && n.Status != "blocked" {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", n.Status))
		}
		if !n.RequiredBeforeResidualPromotion {
			fail(prefix+".required_before_residual_promotion", "required_before_residual_promotion must be true for all null model requirements")
		}
	}

	// 10. Gates check (exact gate cardinality and status)
	requiredGateNames := []string{
		"toy_analysis_only_gate",
		"no_final_truth_claim_gate",
		"local_branch_only_gate",
		"clock_choice_debt_active_gate",
		"classical_target_candidate_only_gate",
		"unit_convention_debt_gate",
		"null_model_debt_gate",
		"faithfulness_contested_gate",
		"no_residual_computation_gate",
		"full_bmc_blocked_gate",
	}

	gateCounts := make(map[string]int)
	for idx, g := range r.Gates {
		prefix := fmt.Sprintf("gates[%d]", idx)
		if g.Status != "pass" && g.Status != "blocked" && g.Status != "contested" {
			fail(prefix+".status", fmt.Sprintf("invalid status: %s", g.Status))
		}
		gateCounts[g.Name]++

		// Status requirements
		if g.Name == "no_residual_computation_gate" && g.Status != "pass" {
			fail(prefix+".status", "no_residual_computation_gate must be pass")
		}
		if g.Name == "full_bmc_blocked_gate" && g.Status != "pass" {
			fail(prefix+".status", "full_bmc_blocked_gate must be pass")
		}
		if g.Name == "faithfulness_contested_gate" && g.Status != "pass" {
			fail(prefix+".status", "faithfulness_contested_gate must be pass")
		}
	}

	for _, name := range requiredGateNames {
		count := gateCounts[name]
		if count == 0 {
			fail("gates", fmt.Sprintf("missing required gate: %s", name))
		} else if count > 1 {
			fail("gates", fmt.Sprintf("gate %s appears %d times (expected exactly once)", name, count))
		}
	}

	// 11. EBP debt check
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

	// 12. Warnings check for required warnings
	hasNoComputeWarn := false
	hasNoClaimWarn := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No Friedmann residual was computed") {
			hasNoComputeWarn = true
		}
		if strings.Contains(w, "No Friedmann recovery is claimed") {
			hasNoClaimWarn = true
		}
	}
	if !hasNoComputeWarn {
		fail("warnings", "must include warning containing: 'No Friedmann residual was computed'")
	}
	if !hasNoClaimWarn {
		fail("warnings", "must include warning containing: 'No Friedmann recovery is claimed'")
	}

	return errs
}
