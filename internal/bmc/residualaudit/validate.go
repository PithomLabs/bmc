package residualaudit

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/PithomLabs/bmc/internal/bmc/report"
)

var forbiddenPhrases = []string{
	"bmc beats null models",
	"bmc outperforms controls",
	"null models passed",
	"null models failed",
	"winner",
	"validated physics",
	"confirmed recovery",
	"friedmann recovery",
	"recovery of friedmann",
	"ready for recovery",
	"full bmc validated",
	"scientifically novel",
	"breakthrough",
	"problem of time solved",
	"quantum gravity validated",
	"classical limit achieved",
	"classical cosmology recovered",
}

func ValidateReport(r *ResidualAuditReport, rawJSON string) []report.ValidationError {
	var errs []report.ValidationError
	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{Field: field, Message: msg, Severity: report.ValidationFail})
	}

	if r.SchemaVersion != SchemaVersion {
		fail("schema_version", "unsupported schema version")
	}
	if r.ArtifactKind != "residual_null_comparison_audit" {
		fail("artifact_kind", "invalid artifact kind")
	}
	if r.Scope != "bmc0a_only" {
		fail("scope", "invalid scope")
	}
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}
	if r.RecoveryClaim {
		fail("recovery_claim", "recovery_claim must be false")
	}
	if r.ScientificNoveltyClaimMade {
		fail("scientific_novelty_claim_made", "scientific_novelty_claim_made must be false")
	}
	if r.BmcBeatsNullModelsClaim {
		fail("bmc_beats_null_models_claim", "bmc_beats_null_models_claim must be false")
	}
	if r.FullBmcToyGate != "blocked" {
		fail("full_bmc_toy_gate", "full_bmc_toy_gate must be blocked")
	}
	if !r.LocalBranchOnly {
		fail("local_branch_only", "local_branch_only must be true")
	}
	if r.GlobalCosmologyClaim {
		fail("global_cosmology_claim", "global_cosmology_claim must be false")
	}
	if r.ResidualAuditComputed {
		if len(r.ComparisonAudits) == 0 {
			fail("comparison_audits", "comparison_audits cannot be empty when residual_audit_computed is true")
		}
		if len(r.StabilityAudits) == 0 {
			fail("stability_audits", "stability_audits cannot be empty when residual_audit_computed is true")
		}
	} else if len(r.ComparisonAudits) > 0 || len(r.StabilityAudits) > 0 {
		fail("residual_audit_computed", "audit records must be empty when residual_audit_computed is false")
	}

	validateSources(r, fail)
	validateComparisonAudits(r, fail)
	validateStabilityAudits(r, fail)
	validateGates(r, fail)
	validateDebt(r, fail)
	validateWarnings(r, fail)
	validateForbiddenText(r, rawJSON, fail)

	return errs
}

func validateSources(r *ResidualAuditReport, fail func(string, string)) {
	required := map[string]bool{
		"bmc0a_clock_readiness":    true,
		"bmc0a_nullrun":            true,
		"bmc0a_local_residual":     true,
		"bmc0a_friedmann_spec":     true,
		"bmc0a_prior_art_boundary": true,
	}
	seen := map[string]int{}
	validProv := map[string]bool{ProvenanceFileRead: true, ProvenanceSourceSummary: true, ProvenanceNotAvailable: true, ProvenanceBlocked: true}
	for idx, src := range r.SourceArtifacts {
		if !required[src.ArtifactID] {
			fail("source_artifacts", fmt.Sprintf("unknown source artifact at source_artifacts[%d]", idx))
		}
		seen[src.ArtifactID]++
		if !validProv[src.Provenance] {
			fail("source_artifacts", fmt.Sprintf("invalid source provenance at source_artifacts[%d]", idx))
		}
		if src.Provenance == ProvenanceFileRead {
			if src.Path == "" {
				fail("source_artifacts", fmt.Sprintf("file_read source path cannot be empty at source_artifacts[%d]", idx))
			}
			if src.Status != "available" && src.Status != "read_success" {
				fail("source_artifacts", fmt.Sprintf("file_read source status must be available at source_artifacts[%d]", idx))
			}
		}
	}
	for id := range required {
		if seen[id] == 0 {
			fail("source_artifacts", "missing required source artifact")
		}
		if seen[id] > 1 {
			fail("source_artifacts", "duplicate source artifact")
		}
	}
}

func validateComparisonAudits(r *ResidualAuditReport, fail func(string, string)) {
	validStatus := map[string]bool{
		AuditStatusComparisonAudited: true, AuditStatusComparisonBlocked: true, AuditStatusComparisonMissing: true,
		AuditStatusComparisonDecorative: true, AuditStatusComparisonUnstable: true, AuditStatusComparisonMixed: true,
		AuditStatusSourceUnavailable: true,
	}
	validProv := map[string]bool{ProvenanceFileRead: true, ProvenanceDerivedFromFileRead: true, ProvenanceSourceSummary: true, ProvenanceBlocked: true}
	validInterp := validInterpretationStatuses()
	for idx, ca := range r.ComparisonAudits {
		if ca.AuditID == "" {
			fail("comparison_audits", fmt.Sprintf("audit_id cannot be empty at comparison_audits[%d]", idx))
		}
		if !validStatus[ca.AuditStatus] {
			fail("comparison_audits", fmt.Sprintf("invalid audit_status at comparison_audits[%d]", idx))
		}
		if !validProv[ca.AuditProvenance] {
			fail("comparison_audits", fmt.Sprintf("invalid audit_provenance at comparison_audits[%d]", idx))
		}
		if !validInterp[ca.InterpretationStatus] {
			fail("comparison_audits", fmt.Sprintf("invalid interpretation_status at comparison_audits[%d]", idx))
		}
		if ca.AuditComputed {
			if len(ca.MetricsAudited) == 0 {
				fail("comparison_audits", fmt.Sprintf("metrics_audited cannot be empty at comparison_audits[%d]", idx))
			}
			if len(ca.TargetResidualIDs) == 0 {
				fail("comparison_audits", fmt.Sprintf("target_residual_ids cannot be empty at comparison_audits[%d]", idx))
			}
			if len(ca.NullModelIDs) == 0 {
				fail("comparison_audits", fmt.Sprintf("null_model_ids cannot be empty at comparison_audits[%d]", idx))
			}
		}
	}
	if !validInterp[r.InterpretationStatus] {
		fail("interpretation_status", "invalid interpretation_status")
	}
}

func validateStabilityAudits(r *ResidualAuditReport, fail func(string, string)) {
	validKind := map[string]bool{PerturbAlphaPoint: true, PerturbPhiPoint: true, PerturbLambdaSpacing: true, PerturbBranchSubset: true, PerturbNone: true}
	validStatus := map[string]bool{StabilityStable: true, StabilitySensitive: true, StabilityUnstable: true, StabilityMissingInput: true, StabilityNonfinite: true, StabilityNotComputed: true}
	for idx, sa := range r.StabilityAudits {
		if sa.StabilityID == "" {
			fail("stability_audits", fmt.Sprintf("stability_id cannot be empty at stability_audits[%d]", idx))
		}
		if !validKind[sa.PerturbationKind] {
			fail("stability_audits", fmt.Sprintf("invalid perturbation_kind at stability_audits[%d]", idx))
		}
		if math.IsNaN(sa.PerturbationMagnitude) || math.IsInf(sa.PerturbationMagnitude, 0) {
			fail("stability_audits", fmt.Sprintf("nonfinite perturbation_magnitude at stability_audits[%d]", idx))
		}
		if sa.PerturbationMagnitude < 0 {
			fail("stability_audits", fmt.Sprintf("negative perturbation_magnitude at stability_audits[%d]", idx))
		}
		if !validStatus[sa.StabilityStatus] {
			fail("stability_audits", fmt.Sprintf("invalid stability_status at stability_audits[%d]", idx))
		}
		if sa.StabilityComputed {
			if sa.StabilityStatus == StabilityNotComputed {
				fail("stability_audits", fmt.Sprintf("computed stability audit cannot have not_computed status at stability_audits[%d]", idx))
			}
			checkRequiredFloat(sa.BaselineValue, "baseline_value", idx, fail)
			checkRequiredFloat(sa.PerturbedValue, "perturbed_value", idx, fail)
			checkRequiredFloat(sa.AbsoluteDelta, "absolute_delta", idx, fail)
			if sa.RelativeDelta != nil {
				checkRequiredFloat(sa.RelativeDelta, "relative_delta", idx, fail)
				if *sa.RelativeDelta < 0 {
					fail("stability_audits", fmt.Sprintf("relative_delta cannot be negative at stability_audits[%d]", idx))
				}
			}
			if sa.AbsoluteDelta != nil && *sa.AbsoluteDelta < 0 {
				fail("stability_audits", fmt.Sprintf("absolute_delta cannot be negative at stability_audits[%d]", idx))
			}
		} else {
			if sa.StabilityStatus == StabilityStable || sa.StabilityStatus == StabilitySensitive {
				fail("stability_audits", fmt.Sprintf("noncomputed stability audit cannot have computed stability status at stability_audits[%d]", idx))
			}
		}
	}
}

func checkRequiredFloat(v *float64, field string, idx int, fail func(string, string)) {
	if v == nil {
		fail("stability_audits", fmt.Sprintf("%s cannot be null for computed stability audit at stability_audits[%d]", field, idx))
		return
	}
	if math.IsNaN(*v) || math.IsInf(*v, 0) {
		fail("stability_audits", fmt.Sprintf("nonfinite %s at stability_audits[%d]", field, idx))
	}
}

func validateGates(r *ResidualAuditReport, fail func(string, string)) {
	required := map[string]bool{
		"toy_analysis_only_gate": true, "no_final_truth_claim_gate": true, "residual_audit_scope_gate": true,
		"comparison_integrity_gate": true, "stability_audit_gate": true, "no_recovery_claim_gate": true,
		"no_scientific_novelty_claim_gate": true, "no_bmc_beats_null_models_claim_gate": true,
		"full_bmc_blocked_gate": true, "faithfulness_contested_gate": true,
	}
	counts := map[string]int{}
	for idx, g := range r.Gates {
		if !required[g.Name] {
			fail("gates", fmt.Sprintf("unknown gate at gates[%d]", idx))
		}
		if g.Status != "pass" {
			fail("gates", fmt.Sprintf("gate status must be pass at gates[%d]", idx))
		}
		counts[g.Name]++
	}
	for name := range required {
		if counts[name] == 0 {
			fail("gates", "missing required gate")
		}
		if counts[name] > 1 {
			fail("gates", "duplicated gate")
		}
	}
}

func validateDebt(r *ResidualAuditReport, fail func(string, string)) {
	if r.EbpDebtVocabulary != "ptw_adversarial_review_debt_status_v0.1" {
		fail("ebp_debt_vocabulary", "invalid ebp_debt_vocabulary")
	}
	if r.EbpDebt.ContainsFinalTruthClaim != "absent" {
		fail("ebp_debt.containsFinalTruthClaim", "containsFinalTruthClaim must be absent")
	}
	if r.EbpDebt.NeedFaithfulnessReview != "contested" {
		fail("ebp_debt.needFaithfulnessReview", "needFaithfulnessReview must be contested")
	}
	if r.EbpDebt.ClockChoiceDebt != "unpaid" || r.EbpDebt.ClassicalTargetDebt != "unpaid" ||
		r.EbpDebt.UnitConventionDebt != "unpaid" || r.EbpDebt.SignConventionDebt != "unpaid" || r.EbpDebt.NormalizationDebt != "unpaid" {
		fail("ebp_debt", "core convention debts must remain unpaid")
	}
	if r.EbpDebt.PromotionStatus != "residual_audit_candidate_only" {
		fail("ebp_debt.promotion_status", "promotion_status must be residual_audit_candidate_only")
	}
}

func validateWarnings(r *ResidualAuditReport, fail func(string, string)) {
	hasNoRecovery := false
	hasFullBlocked := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No recovery claim is made.") {
			hasNoRecovery = true
		}
		if strings.Contains(w, "Full BMC remains blocked.") {
			hasFullBlocked = true
		}
	}
	if !hasNoRecovery {
		fail("warnings", "missing required no-recovery warning")
	}
	if !hasFullBlocked {
		fail("warnings", "missing required full-BMC-blocked warning")
	}
}

func validateForbiddenText(r *ResidualAuditReport, rawJSON string, fail func(string, string)) {
	check := func(value, field string) {
		clean := sanitizeAllowedNegations(strings.ToLower(value))
		for _, phrase := range forbiddenPhrases {
			if strings.Contains(clean, phrase) {
				fail(field, "forbidden phrase detected")
			}
		}
	}
	for idx, w := range r.Warnings {
		check(w, fmt.Sprintf("warnings[%d]", idx))
	}
	for idx, ca := range r.ComparisonAudits {
		check(ca.Notes, fmt.Sprintf("comparison_audits[%d].notes", idx))
		for fidx, finding := range ca.Findings {
			check(finding, fmt.Sprintf("comparison_audits[%d].findings[%d]", idx, fidx))
		}
	}
	for idx, sa := range r.StabilityAudits {
		check(sa.Notes, fmt.Sprintf("stability_audits[%d].notes", idx))
	}
	if rawJSON != "" && rawJSON != "{}" && json.Valid([]byte(rawJSON)) {
		check(rawJSON, "json_content")
	}
}

func sanitizeAllowedNegations(s string) string {
	replacements := map[string]string{
		"no bmc-beats-null-models claim is made.": "no superiority claim is made.",
		"no bmc beats null models claim is made.": "no superiority claim is made.",
		"no_recovery_claim_gate":                  "no_x_claim_gate",
		"no_bmc_beats_null_models_claim_gate":     "no_superiority_claim_gate",
	}
	for old, repl := range replacements {
		s = strings.ReplaceAll(s, old, repl)
	}
	return s
}

func validInterpretationStatuses() map[string]bool {
	return map[string]bool{
		InterpretDiagnosticAuditOnly: true, InterpretComparisonStructurallyHonest: true,
		InterpretComparisonIntegrityFailed: true, InterpretComparisonStabilityMixed: true,
		InterpretComparisonUnstable: true, InterpretInsufficientTargetNullSeparation: true,
		InterpretTargetNullSeparationCandidate: true, InterpretBlockedByMissingResidualInputs: true,
		InterpretBlockedByMissingNullInputs: true, InterpretBlockedBySourceUnavailable: true,
	}
}
