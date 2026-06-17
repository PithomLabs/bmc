package nullrun

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
	"assumed success",
	"fabricated",
	"inferred success",
	"assumed",
	"inferred_success",
	"validated_physics",
}

// ValidateReport performs all strict Sprint 9 validation checks.
func ValidateReport(r *NullRunReport, rawJSON string) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// Raw JSON presence check for optional diagnostics keys
	if rawJSON != "" && rawJSON != "{}" && strings.Contains(rawJSON, "null_model_runs") {
		var rawMap map[string]interface{}
		if err := json.Unmarshal([]byte(rawJSON), &rawMap); err == nil {
			if runs, ok := rawMap["null_model_runs"].([]interface{}); ok {
				for idx, runVal := range runs {
					if runMap, ok := runVal.(map[string]interface{}); ok {
						if diagVal, ok := runMap["diagnostics"]; ok {
							if diagMap, ok := diagVal.(map[string]interface{}); ok {
								requiredKeys := []string{"min_amplitude_r", "max_abs_q_away_from_nodes", "max_phase_gradient"}
								for _, key := range requiredKeys {
									if _, present := diagMap[key]; !present {
										fail("null_model_runs", fmt.Sprintf("missing optional metric key %q in diagnostics at null_model_runs[%d]", key, idx))
									}
								}
							} else {
								fail("null_model_runs", fmt.Sprintf("invalid diagnostics object at null_model_runs[%d]", idx))
							}
						} else {
							fail("null_model_runs", fmt.Sprintf("missing diagnostics object at null_model_runs[%d]", idx))
						}
					}
				}
			}
		}
	}

	// 1. Schema, kind, scope, and vocabulary
	if r.SchemaVersion != "bmc0a-nullrun-v0.1" {
		fail("schema_version", "unsupported schema version")
	}
	if r.EbpDebtVocabulary != "ptw_adversarial_review_debt_status_v0.1" {
		fail("ebp_debt_vocabulary", "invalid ebp_debt_vocabulary detected")
	}
	if r.ArtifactKind != "null_model_runner_report" {
		fail("artifact_kind", "unsupported artifact_kind")
	}
	if r.Scope != "bmc0a_only" {
		fail("scope", "unsupported scope")
	}

	// 2. Toy analysis & final truth claim
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}

	// 3. Simulation flags and recovery/novelty prohibitions
	if r.ResidualComputed {
		fail("residual_computed", "residual_computed must be false")
	}
	if !r.NullDiagnosticsComputed {
		fail("null_diagnostics_computed", "null_diagnostics_computed must be true")
	}
	if r.RecoveryClaim {
		fail("recovery_claim", "recovery_claim must be false")
	}
	if r.ScientificNoveltyClaimMade {
		fail("scientific_novelty_claim_made", "scientific_novelty_claim_made must be false")
	}
	if r.FullBmcToyGate != "blocked" {
		fail("full_bmc_toy_gate", "full_bmc_toy_gate must be blocked")
	}

	// 4. Comparison consistency rules
	hasComparableNullDiags := false
	for _, run := range r.NullModelRuns {
		if run.RunStatus == RunStatusDiagnosticsGenerated {
			hasComparableNullDiags = true
			break
		}
	}
	if !hasComparableNullDiags {
		if r.TargetNullComparisonComputed {
			fail("target_null_comparison_computed", "target_null_comparison_computed must be false when no comparable null diagnostics exist")
		}
		if r.InterpretationStatus != InterpretBlockedByNoComparableDiagnostics {
			fail("interpretation_status", "interpretation status must be blocked_by_no_comparable_null_diagnostics when no comparable null diagnostics exist")
		}
	}

	if len(r.TargetNullDiagnosticComparisons) == 0 {
		if r.TargetNullComparisonComputed {
			fail("target_null_comparison_computed", "target_null_comparison_computed must be false when no comparisons exist")
		}
		if r.InterpretationStatus != InterpretBlockedByNoComparableDiagnostics {
			fail("interpretation_status", "interpretation status must be blocked_by_no_comparable_null_diagnostics when no comparisons exist")
		}
	} else {
		if !r.TargetNullComparisonComputed {
			fail("target_null_comparison_computed", "target_null_comparison_computed must be true when comparisons exist")
		}
		// Valid interpretation status check (phrase-safe)
		is := r.InterpretationStatus
		if is != InterpretDiagComparisonOnly && is != InterpretMixedDiagnostics &&
			is != InterpretInsufficientSeparation && is != InterpretTargetNullSeparationCandidate &&
			is != InterpretBlockedByClockFragility && is != InterpretBlockedByNodeObstruction &&
			is != InterpretBlockedByNoComparableDiagnostics {
			fail("interpretation_status", "invalid interpretation_status detected")
		}

		// Strictly validate each comparison record
		hasValidComputedComparison := false
		for idx, comp := range r.TargetNullDiagnosticComparisons {
			if comp.ComparisonID == "" {
				fail("target_null_diagnostic_comparisons", fmt.Sprintf("comparison_id cannot be empty at target_null_diagnostic_comparisons[%d]", idx))
			}
			if r.TargetNullComparisonComputed && !comp.ComparisonComputed {
				fail("target_null_diagnostic_comparisons", fmt.Sprintf("comparison_computed must be true when report-level comparison is true at target_null_diagnostic_comparisons[%d]", idx))
			}
			if len(comp.MetricsCompared) == 0 {
				fail("target_null_diagnostic_comparisons", fmt.Sprintf("metrics_compared cannot be empty at target_null_diagnostic_comparisons[%d]", idx))
			}
			if len(comp.NullModelIDs) == 0 {
				fail("target_null_diagnostic_comparisons", fmt.Sprintf("null_model_ids cannot be empty at target_null_diagnostic_comparisons[%d]", idx))
			}

			// Validate each null model ID in the comparison
			for _, mid := range comp.NullModelIDs {
				foundRun := false
				var referencedRun NullModelRun
				for _, run := range r.NullModelRuns {
					if run.NullModelID == mid {
						foundRun = true
						referencedRun = run
						break
					}
				}
				if !foundRun {
					fail("target_null_diagnostic_comparisons", fmt.Sprintf("referenced null_model_id %q does not exist at target_null_diagnostic_comparisons[%d]", mid, idx))
				} else {
					if referencedRun.RunStatus != RunStatusDiagnosticsGenerated {
						fail("target_null_diagnostic_comparisons", fmt.Sprintf("referenced null_model_id %q has invalid run_status at target_null_diagnostic_comparisons[%d]", mid, idx))
					}
				}
			}

			// Valid interpretation status check
			isComp := comp.InterpretationStatus
			if isComp != InterpretDiagComparisonOnly && isComp != InterpretMixedDiagnostics &&
				isComp != InterpretInsufficientSeparation && isComp != InterpretTargetNullSeparationCandidate &&
				isComp != InterpretBlockedByClockFragility && isComp != InterpretBlockedByNodeObstruction &&
				isComp != InterpretBlockedByNoComparableDiagnostics {
				fail("target_null_diagnostic_comparisons", fmt.Sprintf("invalid interpretation_status detected at target_null_diagnostic_comparisons[%d]", idx))
			}

			// Reject victory/winner language in reason and status
			checkVictoryLanguage := func(val, path string) {
				valLower := strings.ToLower(val)
				victoryTerms := []string{"winner", "outperformed", "validated", "recovered", "proved"}
				for _, term := range victoryTerms {
					if strings.Contains(valLower, term) {
						fail(path, "victory or superiority language detected")
					}
				}
			}
			checkVictoryLanguage(comp.Reason, fmt.Sprintf("target_null_diagnostic_comparisons[%d].reason", idx))
			checkVictoryLanguage(comp.InterpretationStatus, fmt.Sprintf("target_null_diagnostic_comparisons[%d].interpretation_status", idx))

			if comp.ComparisonComputed && len(comp.MetricsCompared) > 0 && len(comp.NullModelIDs) > 0 {
				// verify that referenced runs are valid
				allReferencedRunsValid := true
				for _, mid := range comp.NullModelIDs {
					foundRun := false
					var referencedRun NullModelRun
					for _, run := range r.NullModelRuns {
						if run.NullModelID == mid {
							foundRun = true
							referencedRun = run
							break
						}
					}
					if !foundRun || referencedRun.RunStatus != RunStatusDiagnosticsGenerated {
						allReferencedRunsValid = false
					}
				}
				if allReferencedRunsValid {
					hasValidComputedComparison = true
				}
			}
		}

		if r.TargetNullComparisonComputed && !hasValidComputedComparison {
			fail("target_null_diagnostic_comparisons", "at least one TargetNullDiagnosticComparison must satisfy honest comparison constraints")
		}
	}

	// 5. Null model run record validation
	requiredNullModels := map[string]bool{
		"constant_phase_control":                             true,
		"randomized_phase_control":                           true,
		"matched_amplitude_randomized_phase_control":         true,
		"classical_frw_reference_trajectory":                 true,
		"same_branch_segmentation_under_null_wavefunctions":  true,
		"node_neighborhood_stress_case":                      true,
		"clock_choice_alternative_branch_diagnostic":         true,
	}

	seenNullModelIDs := make(map[string]bool)
	for idx, run := range r.NullModelRuns {
		if run.NullModelID == "" {
			fail("null_model_runs", "null_model_id cannot be empty")
		}
		if seenNullModelIDs[run.NullModelID] {
			fail("null_model_runs", "duplicate null_model_id detected")
		}
		seenNullModelIDs[run.NullModelID] = true

		if !requiredNullModels[run.NullModelID] {
			fail("null_model_runs", fmt.Sprintf("unknown null_model_id detected at null_model_runs[%d]", idx))
		}

		// Run status validation (phrase-safe)
		if run.RunStatus != RunStatusDiagnosticsGenerated &&
			run.RunStatus != RunStatusBlocked &&
			run.RunStatus != RunStatusDeferred {
			fail("null_model_runs", fmt.Sprintf("invalid run_status detected at null_model_runs[%d]", idx))
		}

		// Diagnostic provenance validation (phrase-safe)
		prov := run.DiagnosticProvenance
		if prov != ProvenanceComputed && prov != ProvenanceDeterministicFixture &&
			prov != ProvenanceSourceArtifact && prov != ProvenanceBlocked &&
			prov != ProvenanceDeferred {
			fail("null_model_runs", fmt.Sprintf("invalid diagnostic_provenance detected at null_model_runs[%d]", idx))
		}

		// Diagnostic status validation (phrase-safe)
		ds := run.DiagnosticStatus
		if ds != DiagStatusFinite && ds != DiagStatusNonfinite &&
			ds != DiagStatusNodeBlocked && ds != DiagStatusClockFragile &&
			ds != DiagStatusLocalOnly && ds != DiagStatusNotAvailable {
			fail("null_model_runs", fmt.Sprintf("invalid diagnostic_status detected at null_model_runs[%d]", idx))
		}

		// Metrics validation: check negative sentinels
		diag := run.Diagnostics
		if diag.NumValidTrajectoryPoints < 0 || diag.NumClockSegments < 0 || diag.NumTurningPoints < 0 {
			fail("null_model_runs", fmt.Sprintf("invalid sentinel numeric metric detected at null_model_runs[%d]", idx))
		}

		// Validate non-finite and negative/sentinel floats
		checkFloat := func(f *float64, fieldName string) {
			if f != nil {
				if math.IsNaN(*f) || math.IsInf(*f, 0) {
					fail("null_model_runs", fmt.Sprintf("nonfinite metric detected in %s at null_model_runs[%d]", fieldName, idx))
				}
				if *f < 0 {
					fail("null_model_runs", fmt.Sprintf("negative or sentinel metric value detected in %s at null_model_runs[%d]", fieldName, idx))
				}
			}
		}
		checkFloat(diag.MinAmplitudeR, "min_amplitude_r")
		checkFloat(diag.MaxAbsQAwayFromNodes, "max_abs_q_away_from_nodes")
		checkFloat(diag.MaxPhaseGradient, "max_phase_gradient")

		if run.NullModelID == "classical_frw_reference_trajectory" {
			if run.RunStatus != RunStatusDiagnosticsGenerated &&
				run.RunStatus != RunStatusBlocked &&
				run.RunStatus != RunStatusDeferred {
				fail("null_model_runs", "invalid run_status for classical_frw_reference_trajectory")
			}
			if run.DiagnosticProvenance != ProvenanceDeterministicFixture &&
				run.DiagnosticProvenance != ProvenanceSourceArtifact &&
				run.DiagnosticProvenance != ProvenanceBlocked {
				fail("null_model_runs", "invalid diagnostic_provenance for classical_frw_reference_trajectory")
			}
			if !strings.Contains(run.Notes, "reference comparator only; no residual or recovery interpretation") {
				fail("null_model_runs", "invalid notes for classical_frw_reference_trajectory")
			}
		}
	}

	for id := range requiredNullModels {
		if !seenNullModelIDs[id] {
			fail("null_model_runs", "missing required null model run")
		}
	}

	// 6. Target-Null comparisons validation
	for idx, comp := range r.TargetNullDiagnosticComparisons {
		if comp.ComparisonID == "" {
			fail("target_null_diagnostic_comparisons", "comparison_id cannot be empty")
		}
		if !comp.ComparisonComputed {
			fail("target_null_diagnostic_comparisons", "comparison_computed must be true")
		}
		// Interpretation status validation (phrase-safe)
		is := comp.InterpretationStatus
		if is != InterpretDiagComparisonOnly && is != InterpretMixedDiagnostics &&
			is != InterpretInsufficientSeparation && is != InterpretTargetNullSeparationCandidate &&
			is != InterpretBlockedByClockFragility && is != InterpretBlockedByNodeObstruction &&
			is != InterpretBlockedByNoComparableDiagnostics {
			fail("target_null_diagnostic_comparisons", fmt.Sprintf("invalid interpretation_status detected at target_null_diagnostic_comparisons[%d]", idx))
		}
	}

	// 7. Gates validation
	requiredGates := map[string]bool{
		"toy_analysis_only_gate":               true,
		"no_final_truth_claim_gate":            true,
		"no_residual_computation_gate":         true,
		"null_diagnostics_computed_gate":       true,
		"target_null_comparison_computed_gate": true,
		"no_winner_claim_gate":                 true,
		"no_recovery_claim_gate":               true,
		"no_scientific_novelty_claim_gate":     true,
		"full_bmc_blocked_gate":                true,
		"faithfulness_contested_gate":          true,
	}
	gateCounts := make(map[string]int)
	for idx, g := range r.Gates {
		if g.Name == "" {
			fail("gates", fmt.Sprintf("invalid gate name detected at gates[%d]", idx))
		}
		gateCounts[g.Name]++
		if !requiredGates[g.Name] {
			fail("gates", fmt.Sprintf("unknown gate detected at gates[%d]", idx))
		}
		if g.Status != "pass" {
			fail("gates", fmt.Sprintf("invalid gate status detected at gates[%d]", idx))
		}
	}

	for gName := range requiredGates {
		count := gateCounts[gName]
		if count == 0 {
			fail("gates", "missing required gate")
		} else if count > 1 {
			fail("gates", "duplicated gate detected")
		}
	}

	// 8. EBP Debt validation
	debtVal := r.EbpDebt
	validDebtStatuses := map[string]bool{
		"unpaid":                           true,
		"partial":                          true,
		"retired":                          true,
		"contested":                        true,
		"overclaimed":                      true,
		"absent":                           true,
		"null_model_runner_candidate_only": true,
	}

	checkDebtField := func(field, val string) {
		if !validDebtStatuses[val] {
			fail(field, fmt.Sprintf("invalid debt status detected in %s", field))
		}
	}

	checkDebtField("ebp_debt.needLiteratureAudit", debtVal.NeedLiteratureAudit)
	checkDebtField("ebp_debt.needMap", debtVal.NeedMap)
	checkDebtField("ebp_debt.needInvariant", debtVal.NeedInvariant)
	checkDebtField("ebp_debt.needToyCheck", debtVal.NeedToyCheck)
	checkDebtField("ebp_debt.needNullModel", debtVal.NeedNullModel)
	checkDebtField("ebp_debt.needObstruction", debtVal.NeedObstruction)
	checkDebtField("ebp_debt.needFaithfulnessReview", debtVal.NeedFaithfulnessReview)
	checkDebtField("ebp_debt.clock_choice_debt", debtVal.ClockChoiceDebt)
	checkDebtField("ebp_debt.classical_target_debt", debtVal.ClassicalTargetDebt)
	checkDebtField("ebp_debt.unit_convention_debt", debtVal.UnitConventionDebt)
	checkDebtField("ebp_debt.sign_convention_debt", debtVal.SignConventionDebt)
	checkDebtField("ebp_debt.normalization_debt", debtVal.NormalizationDebt)
	checkDebtField("ebp_debt.containsFinalTruthClaim", debtVal.ContainsFinalTruthClaim)
	checkDebtField("ebp_debt.promotion_status", debtVal.PromotionStatus)

	if r.EbpDebt.ContainsFinalTruthClaim != "absent" {
		fail("ebp_debt.containsFinalTruthClaim", "containsFinalTruthClaim must be absent")
	}
	if r.EbpDebt.NeedFaithfulnessReview != "contested" {
		fail("ebp_debt.needFaithfulnessReview", "needFaithfulnessReview must be contested")
	}
	if r.EbpDebt.PromotionStatus != "null_model_runner_candidate_only" {
		fail("ebp_debt.promotion_status", "promotion_status must be null_model_runner_candidate_only")
	}

	// 9. Warnings validation
	hasNoResidualWarn := false
	hasBlockedWarn := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No residual was computed.") {
			hasNoResidualWarn = true
		}
		if strings.Contains(w, "Full BMC remains blocked.") {
			hasBlockedWarn = true
		}
	}
	if !hasNoResidualWarn {
		fail("warnings", "missing warning: “No residual was computed.”")
	}
	if !hasBlockedWarn {
		fail("warnings", "missing warning: “Full BMC remains blocked.”")
	}

	// 10. Case-insensitive forbidden phrase scanning (phrase-safe errors)
	checkForbidden := func(val string, loc string) {
		valLower := strings.ToLower(val)
		for _, phrase := range forbiddenPhrases {
			if strings.Contains(valLower, phrase) {
				fail(loc, fmt.Sprintf("forbidden phrase detected at %s", loc))
			}
		}
	}

	for idx, w := range r.Warnings {
		checkForbidden(w, fmt.Sprintf("warning[%d]", idx))
	}
	for idx, run := range r.NullModelRuns {
		checkForbidden(run.BlockedReason, fmt.Sprintf("null_model_runs[%d].blocked_reason", idx))
		checkForbidden(run.Notes, fmt.Sprintf("null_model_runs[%d].notes", idx))
		for wIdx, dw := range run.Diagnostics.DiagnosticWarnings {
			checkForbidden(dw, fmt.Sprintf("null_model_runs[%d].diagnostics.diagnostic_warnings[%d]", idx, wIdx))
		}
	}
	for idx, comp := range r.TargetNullDiagnosticComparisons {
		checkForbidden(comp.Reason, fmt.Sprintf("target_null_diagnostic_comparisons[%d].reason", idx))
	}

	// Double check raw JSON content for safety (excluding allowed gate metadata containing "winner")
	rawJSONLower := strings.ToLower(rawJSON)
	rawJSONClean := strings.ReplaceAll(rawJSONLower, "no_winner_claim_gate", "no_xxxxxx_claim_gate")
	rawJSONClean = strings.ReplaceAll(rawJSONClean, "confirms no winner claims are made.", "confirms no xxxxxx claims are made.")
	for _, phrase := range forbiddenPhrases {
		if strings.Contains(rawJSONClean, phrase) {
			fail("json_content", "forbidden phrase detected in JSON content")
		}
	}

	return errs
}
