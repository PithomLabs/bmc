package residualrun

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
	"classical limit achieved",
	"classical cosmology",
}

// ValidateReport performs all strict Sprint 10.1 validation checks.
func ValidateReport(r *ResidualRunReport, rawJSON string) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// Raw JSON presence check for optional diagnostics keys inside metrics
	if rawJSON != "" && rawJSON != "{}" && strings.Contains(rawJSON, "candidate_residual_diagnostics") {
		var rawMap map[string]interface{}
		if err := json.Unmarshal([]byte(rawJSON), &rawMap); err == nil {
			if runs, ok := rawMap["candidate_residual_diagnostics"].([]interface{}); ok {
				for idx, runVal := range runs {
					if runMap, ok := runVal.(map[string]interface{}); ok {
						if metricsVal, ok := runMap["metrics"]; ok {
							if metricsMap, ok := metricsVal.(map[string]interface{}); ok {
								requiredKeys := []string{"mean_abs_residual", "max_abs_residual", "rms_residual"}
								for _, key := range requiredKeys {
									if _, present := metricsMap[key]; !present {
										fail("candidate_residual_diagnostics", fmt.Sprintf("missing optional metric key %q in metrics at candidate_residual_diagnostics[%d]", key, idx))
									}
								}
							} else {
								fail("candidate_residual_diagnostics", fmt.Sprintf("invalid metrics object at candidate_residual_diagnostics[%d]", idx))
							}
						} else {
							fail("candidate_residual_diagnostics", fmt.Sprintf("missing metrics object at candidate_residual_diagnostics[%d]", idx))
						}
					}
				}
			}
		}
	}

	// 1. Schema, kind, scope, and vocabulary
	if r.SchemaVersion != "bmc0a-local-residual-v0.1" {
		fail("schema_version", "unsupported schema version")
	}
	if r.EbpDebtVocabulary != "ptw_adversarial_review_debt_status_v0.1" {
		fail("ebp_debt_vocabulary", "invalid ebp_debt_vocabulary detected")
	}
	if r.ArtifactKind != "candidate_local_branch_residual_runner" {
		fail("artifact_kind", "unsupported artifact_kind")
	}
	if r.Scope != "bmc0a_only" {
		fail("scope", "unsupported scope")
	}

	// 2. Toy analysis, final truth claim, and local boundary
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}
	if !r.LocalBranchOnly {
		fail("local_branch_only", "local_branch_only must be true")
	}
	if r.GlobalCosmologyClaim {
		fail("global_cosmology_claim", "global_cosmology_claim must be false")
	}

	// 3. Simulation flags and recovery/novelty prohibitions
	if r.ResidualRecoveryClaim {
		fail("residual_recovery_claim", "residual_recovery_claim must be false")
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

	// 4. Source Artifacts Validation
	if len(r.SourceArtifacts) == 0 {
		fail("source_artifacts", "source_artifacts cannot be empty")
	}
	validProvenances := map[string]bool{
		"file_read":               true,
		"source_artifact_summary": true,
		"not_available":           true,
	}

	expectedSources := map[string]bool{
		"bmc0a_clock_readiness":    true,
		"bmc0a_friedmann_spec":     true,
		"bmc0a_nullrun":            true,
		"bmc0a_prior_art_boundary": true,
	}
	seenSources := make(map[string]int)

	for idx, src := range r.SourceArtifacts {
		if src.ArtifactID == "" {
			fail("source_artifacts", fmt.Sprintf("artifact_id cannot be empty at source_artifacts[%d]", idx))
		}
		if !expectedSources[src.ArtifactID] {
			fail("source_artifacts", fmt.Sprintf("unknown source artifact ID %q at source_artifacts[%d]", src.ArtifactID, idx))
		}
		seenSources[src.ArtifactID]++

		if !validProvenances[src.Provenance] {
			fail("source_artifacts", fmt.Sprintf("invalid provenance at source_artifacts[%d]", idx))
		}

		if src.Provenance == "file_read" {
			if src.Path == "" {
				fail("source_artifacts", fmt.Sprintf("path cannot be empty for file_read at source_artifacts[%d]", idx))
			}
			if src.Status != "available" && src.Status != "read_success" {
				fail("source_artifacts", fmt.Sprintf("status must be available or read_success for file_read at source_artifacts[%d]", idx))
			}
		}

		if r.CandidateResidualComputed && src.Provenance != "file_read" {
			fail("source_artifacts", fmt.Sprintf("provenance must be file_read when candidate_residual_computed is true at source_artifacts[%d]", idx))
		}
	}

	for sID := range expectedSources {
		if seenSources[sID] == 0 {
			fail("source_artifacts", fmt.Sprintf("missing expected source artifact: %s", sID))
		} else if seenSources[sID] > 1 {
			fail("source_artifacts", fmt.Sprintf("duplicate source artifact: %s", sID))
		}
	}

	// Find clock readiness source artifact
	var clockSrc *SourceArtifactRef
	for i := range r.SourceArtifacts {
		if r.SourceArtifacts[i].ArtifactID == "bmc0a_clock_readiness" {
			clockSrc = &r.SourceArtifacts[i]
			break
		}
	}

	if clockSrc != nil && clockSrc.Provenance == "file_read" {
		registryMap := make(map[string]bool)
		for _, bid := range r.SourceBranchRegistry.BranchIDs {
			registryMap[bid] = true
		}

		for idx, b := range r.LocalBranchEligibility {
			if !registryMap[b.BranchID] {
				fail("local_branch_eligibility", fmt.Sprintf("branch_id %q is not present in source_branch_registry at local_branch_eligibility[%d]", b.BranchID, idx))
			}
		}

		for idx, rd := range r.CandidateResidualDiagnostics {
			if rd.ResidualProvenance == ProvenanceComputed {
				if !registryMap[rd.BranchID] {
					fail("candidate_residual_diagnostics", fmt.Sprintf("computed diagnostic branch %q is not present in source_branch_registry at candidate_residual_diagnostics[%d]", rd.BranchID, idx))
				}
			}
		}

		for idx, cl := range r.CalculationLedger {
			if cl.CalculationStatus == "computed_from_local_branch" {
				if !registryMap[cl.BranchID] {
					fail("calculation_ledger", fmt.Sprintf("calculation ledger branch %q is not present in source_branch_registry at calculation_ledger[%d]", cl.BranchID, idx))
				}
			}
		}
	}

	// 5. Local Branch Eligibility Validation
	if len(r.LocalBranchEligibility) == 0 {
		fail("local_branch_eligibility", "local_branch_eligibility cannot be empty")
	}
	validEligibilityStatuses := map[string]bool{
		EligibilityEligibleLocalBranch:            true,
		EligibilityBlockedByNodeObstruction:       true,
		EligibilityBlockedByClockFragility:        true,
		EligibilityBlockedByNonfiniteTrajectory:   true,
		EligibilityBlockedByDerivativeUnreadiness: true,
		EligibilitySourceUnavailable:              true,
	}
	eligibleMap := make(map[string]LocalBranchEligibility)
	for idx, b := range r.LocalBranchEligibility {
		if b.BranchID == "" {
			fail("local_branch_eligibility", fmt.Sprintf("branch_id cannot be empty at local_branch_eligibility[%d]", idx))
		}
		if !validEligibilityStatuses[b.EligibilityStatus] {
			fail("local_branch_eligibility", fmt.Sprintf("invalid eligibility_status at local_branch_eligibility[%d]", idx))
		}
		if b.Eligible && b.EligibilityStatus != EligibilityEligibleLocalBranch {
			fail("local_branch_eligibility", fmt.Sprintf("inconsistent eligible status at local_branch_eligibility[%d]", idx))
		}
		eligibleMap[b.BranchID] = b
	}

	// 6. Convention Ledger Validation
	if len(r.ConventionLedger) == 0 {
		fail("convention_ledger", "convention_ledger cannot be empty")
	}
	requiredDebts := map[string]bool{
		"clock_choice_debt":        true,
		"classical_target_debt":    true,
		"unit_convention_debt":     true,
		"sign_convention_debt":     true,
		"normalization_debt":       true,
		"faithfulness_review_debt": true,
	}
	seenDebts := make(map[string]int)
	validLedgerStatuses := map[string]bool{
		"unpaid":    true,
		"partial":   true,
		"contested": true,
		"blocked":   true,
	}
	for idx, ledger := range r.ConventionLedger {
		if !requiredDebts[ledger.ConventionID] {
			fail("convention_ledger", fmt.Sprintf("unknown convention_id at convention_ledger[%d]", idx))
		}
		seenDebts[ledger.ConventionID]++
		if !validLedgerStatuses[ledger.Status] {
			fail("convention_ledger", fmt.Sprintf("invalid ledger status (cannot be retired) at convention_ledger[%d]", idx))
		}
		if ledger.Status == "retired" {
			fail("convention_ledger", fmt.Sprintf("convention debt %q cannot be retired at convention_ledger[%d]", ledger.ConventionID, idx))
		}
		if ledger.Description == "" {
			fail("convention_ledger", fmt.Sprintf("description cannot be empty at convention_ledger[%d]", idx))
		}
		if !ledger.HumanReviewRequired {
			fail("convention_ledger", fmt.Sprintf("human_review_required must be true for unresolved debt %q at convention_ledger[%d]", ledger.ConventionID, idx))
		}
		if ledger.ConventionID != "faithfulness_review_debt" && ledger.Status == "partial" {
			fail("convention_ledger", fmt.Sprintf("core convention debt %q cannot have status partial at convention_ledger[%d]", ledger.ConventionID, idx))
		}
	}
	for d := range requiredDebts {
		if seenDebts[d] == 0 {
			fail("convention_ledger", fmt.Sprintf("missing required convention debt: %s", d))
		} else if seenDebts[d] > 1 {
			fail("convention_ledger", fmt.Sprintf("duplicated convention debt: %s", d))
		}
	}

	// 7. Candidate Residual Diagnostics Validation
	if r.CandidateResidualComputed && len(r.CandidateResidualDiagnostics) == 0 {
		fail("candidate_residual_computed", "candidate_residual_diagnostics cannot be empty when candidate_residual_computed is true")
	}
	hasEligibleBranchDiagnostic := false
	validResidualStatuses := map[string]bool{
		ResidualStatusGenerated:                true,
		ResidualStatusInputBlocked:             true,
		ResidualStatusNonfinite:                true,
		ResidualStatusBlockedByClockFragility:  true,
		ResidualStatusBlockedByNodeObstruction: true,
		ResidualStatusBlockedByConventionDebt:  true,
		ResidualStatusSourceUnavailable:        true,
	}
	validResidualProvenances := map[string]bool{
		ProvenanceComputed:             true,
		ProvenanceDeterministicFixture: true,
		ProvenanceSourceArtifact:       true,
		ProvenanceBlocked:              true,
	}

	for idx, rd := range r.CandidateResidualDiagnostics {
		if rd.BranchID == "" {
			fail("candidate_residual_diagnostics", fmt.Sprintf("branch_id cannot be empty at candidate_residual_diagnostics[%d]", idx))
		}
		b, foundBranch := eligibleMap[rd.BranchID]
		if !foundBranch {
			fail("candidate_residual_diagnostics", fmt.Sprintf("referenced branch_id %q does not exist at candidate_residual_diagnostics[%d]", rd.BranchID, idx))
		}

		isEligible := foundBranch && b.Eligible

		if !isEligible && rd.ResidualComputed {
			fail("candidate_residual_diagnostics", fmt.Sprintf("residual_computed cannot be true for ineligible branch at candidate_residual_diagnostics[%d]", idx))
		}

		if rd.ResidualComputed {
			// Enforce full branch eligibility criteria
			if !b.Eligible ||
				b.EligibilityStatus != EligibilityEligibleLocalBranch ||
				!b.NodeContactFree ||
				!b.TrajectoryFinite ||
				b.DerivativeReadinessStatus != "ready" {
				fail("candidate_residual_diagnostics", fmt.Sprintf("residual_computed is true but branch %q does not satisfy all eligibility criteria at candidate_residual_diagnostics[%d]", rd.BranchID, idx))
			}

			// Residual status and provenance consistency
			if rd.ResidualStatus != ResidualStatusGenerated {
				fail("candidate_residual_diagnostics", fmt.Sprintf("residual_computed is true but status is not generated at candidate_residual_diagnostics[%d]", idx))
			}
			if rd.ResidualProvenance == ProvenanceBlocked {
				fail("candidate_residual_diagnostics", fmt.Sprintf("residual_computed is true but provenance is blocked at candidate_residual_diagnostics[%d]", idx))
			}
		} else {
			if rd.ResidualStatus == ResidualStatusGenerated {
				fail("candidate_residual_diagnostics", fmt.Sprintf("residual_computed is false but status is generated at candidate_residual_diagnostics[%d]", idx))
			}
		}

		if !validResidualStatuses[rd.ResidualStatus] {
			fail("candidate_residual_diagnostics", fmt.Sprintf("invalid residual_status at candidate_residual_diagnostics[%d]", idx))
		}
		if !validResidualProvenances[rd.ResidualProvenance] {
			fail("candidate_residual_diagnostics", fmt.Sprintf("invalid residual_provenance at candidate_residual_diagnostics[%d]", idx))
		}

		// Enforce computed_from_bmc0a_local_branch requires auditable calculation provenance
		if rd.ResidualProvenance == ProvenanceComputed {
			foundLedger := false
			for _, ledger := range r.CalculationLedger {
				if ledger.BranchID == rd.BranchID && ledger.CalculationStatus == "computed_from_local_branch" {
					foundLedger = true
					break
				}
			}
			if !foundLedger {
				fail("calculation_ledger", fmt.Sprintf("provenance computed_from_bmc0a_local_branch requires matching calculation ledger entry for branch %q", rd.BranchID))
			}
		}

		// Tighten deterministic_fixture rule
		if rd.ResidualProvenance == ProvenanceDeterministicFixture {
			if r.InterpretationStatus == "target_null_residual_separation_candidate_unpromoted" {
				fail("interpretation_status", "cannot use target_null_residual_separation_candidate_unpromoted interpretation when using deterministic_fixture")
			}
		}

		if rd.ResidualComputed {
			if len(rd.ResidualInputPoints) == 0 {
				fail("candidate_residual_diagnostics", fmt.Sprintf("computed diagnostic must have non-empty residual_input_points at candidate_residual_diagnostics[%d]", idx))
			}
			if len(rd.ResidualInputPoints) != rd.Metrics.NumEvaluationPoints {
				fail("candidate_residual_diagnostics", fmt.Sprintf("residual_input_points length must match num_evaluation_points at candidate_residual_diagnostics[%d]", idx))
			}
			var prevLambda float64
			for pIdx, pt := range rd.ResidualInputPoints {
				if pt.BranchID != rd.BranchID {
					fail("candidate_residual_diagnostics", fmt.Sprintf("residual input point branch_id %q does not match diagnostic branch_id %q at candidate_residual_diagnostics[%d].points[%d]", pt.BranchID, rd.BranchID, idx, pIdx))
				}
				if pt.PointIndex != pIdx {
					fail("candidate_residual_diagnostics", fmt.Sprintf("residual input point index %d does not match slice index %d at candidate_residual_diagnostics[%d].points[%d]", pt.PointIndex, pIdx, idx, pIdx))
				}
				if pt.Lambda == nil || pt.Alpha == nil || pt.Phi == nil || pt.CandidateLeftHandSide == nil || pt.CandidateRightHandSide == nil {
					fail("candidate_residual_diagnostics", fmt.Sprintf("residual input point fields cannot be nil at candidate_residual_diagnostics[%d].points[%d]", idx, pIdx))
				} else {
					pointValues := []struct {
						name string
						val  float64
					}{
						{"lambda", *pt.Lambda},
						{"alpha", *pt.Alpha},
						{"phi", *pt.Phi},
						{"candidate_left_hand_side", *pt.CandidateLeftHandSide},
						{"candidate_right_hand_side", *pt.CandidateRightHandSide},
					}
					for _, pv := range pointValues {
						if math.IsNaN(pv.val) || math.IsInf(pv.val, 0) {
							fail("candidate_residual_diagnostics", fmt.Sprintf("nonfinite residual input point field %s at candidate_residual_diagnostics[%d].points[%d]", pv.name, idx, pIdx))
						}
					}
					if pIdx > 0 && *pt.Lambda <= prevLambda {
						fail("candidate_residual_diagnostics", fmt.Sprintf("lambda must be strictly increasing within candidate_residual_diagnostics[%d].points", idx))
					}
					prevLambda = *pt.Lambda
				}
				if pt.InputProvenance != "file_read" && pt.InputProvenance != "derived_from_file_read" {
					fail("candidate_residual_diagnostics", fmt.Sprintf("invalid residual input point provenance at candidate_residual_diagnostics[%d].points[%d]", idx, pIdx))
				}
			}
		} else {
			if len(rd.ResidualInputPoints) > 0 {
				fail("candidate_residual_diagnostics", fmt.Sprintf("blocked diagnostic must have empty residual_input_points at candidate_residual_diagnostics[%d]", idx))
			}
		}

		metrics := rd.Metrics
		if !isEligible {
			if metrics.MeanAbsResidual != nil || metrics.MaxAbsResidual != nil || metrics.RmsResidual != nil {
				fail("candidate_residual_diagnostics", fmt.Sprintf("metrics must be null for ineligible branch at candidate_residual_diagnostics[%d]", idx))
			}
		} else {
			if rd.ResidualComputed {
				if metrics.MeanAbsResidual == nil || metrics.MaxAbsResidual == nil || metrics.RmsResidual == nil {
					fail("candidate_residual_diagnostics", fmt.Sprintf("computed residual diagnostic cannot have nil metrics at candidate_residual_diagnostics[%d]", idx))
				}
			}
		}

		// Metric point counts check
		if metrics.NumEvaluationPoints < 0 {
			fail("candidate_residual_diagnostics", fmt.Sprintf("num_evaluation_points cannot be negative at candidate_residual_diagnostics[%d]", idx))
		}
		if metrics.NumFiniteResidualPoints < 0 {
			fail("candidate_residual_diagnostics", fmt.Sprintf("num_finite_residual_points cannot be negative at candidate_residual_diagnostics[%d]", idx))
		}
		if metrics.NumFiniteResidualPoints > metrics.NumEvaluationPoints {
			fail("candidate_residual_diagnostics", fmt.Sprintf("num_finite_residual_points cannot exceed num_evaluation_points at candidate_residual_diagnostics[%d]", idx))
		}
		if metrics.ResidualFinite && metrics.NumFiniteResidualPoints == 0 {
			fail("candidate_residual_diagnostics", fmt.Sprintf("residual_finite cannot be true when num_finite_residual_points is zero at candidate_residual_diagnostics[%d]", idx))
		}

		checkFloat := func(f *float64, fieldName string) {
			if f != nil {
				if math.IsNaN(*f) || math.IsInf(*f, 0) {
					fail("candidate_residual_diagnostics", fmt.Sprintf("nonfinite metric detected in %s at candidate_residual_diagnostics[%d]", fieldName, idx))
				}
				if *f < 0 {
					fail("candidate_residual_diagnostics", fmt.Sprintf("negative or sentinel metric value detected in %s at candidate_residual_diagnostics[%d]", fieldName, idx))
				}
			}
		}
		checkFloat(metrics.MeanAbsResidual, "mean_abs_residual")
		checkFloat(metrics.MaxAbsResidual, "max_abs_residual")
		checkFloat(metrics.RmsResidual, "rms_residual")

		if rd.ResidualComputed && isEligible {
			hasEligibleBranchDiagnostic = true
		}
	}

	if r.CandidateResidualComputed && !hasEligibleBranchDiagnostic {
		fail("candidate_residual_computed", "candidate_residual_computed is true but no eligible local branch has computed diagnostics")
	}

	if r.CandidateResidualComputed {
		is := r.InterpretationStatus
		if is != InterpretDiagComparisonOnly && is != InterpretMixedResidualDiagnostics &&
			is != InterpretInsufficientTargetNullSeparation && is != InterpretTargetNullResidualSeparationCandidate &&
			is != InterpretBlockedByNoComparableNullDiagnostics && is != InterpretBlockedByConventionDebt &&
			is != InterpretBlockedByClockFragility && is != InterpretBlockedByNodeObstruction &&
			is != InterpretBlockedByNoEligibleLocalBranch && is != InterpretBlockedByMissingResidualInputs {
			fail("interpretation_status", "invalid interpretation_status detected")
		}
	}

	if !r.CandidateResidualComputed {
		// Enforce candidate_residual_computed = false while any diagnostic has residual_computed = true
		for idx, rd := range r.CandidateResidualDiagnostics {
			if rd.ResidualComputed {
				fail("candidate_residual_computed", fmt.Sprintf("candidate_residual_computed is false but diagnostic at index %d has residual_computed = true", idx))
			}
		}

		// Interpretation status must be blocked
		is := r.InterpretationStatus
		if is != InterpretBlockedByNoEligibleLocalBranch && is != InterpretBlockedByConventionDebt &&
			is != InterpretBlockedByClockFragility && is != InterpretBlockedByNodeObstruction &&
			is != InterpretBlockedByNoComparableNullDiagnostics && is != InterpretBlockedByMissingResidualInputs {
			fail("interpretation_status", "interpretation_status must be blocked when candidate_residual_computed is false")
		}
		// All metric values in all runs must be nil
		for idx, rd := range r.CandidateResidualDiagnostics {
			if rd.Metrics.MeanAbsResidual != nil || rd.Metrics.MaxAbsResidual != nil || rd.Metrics.RmsResidual != nil {
				fail("candidate_residual_diagnostics", fmt.Sprintf("metrics must be null when candidate_residual_computed is false at candidate_residual_diagnostics[%d]", idx))
			}
		}
	}

	// 8. Target/Null Comparison Validation
	blockedOrDeferredNulls := map[string]bool{
		"same_branch_segmentation_under_null_wavefunctions": true,
		"node_neighborhood_stress_case":                     true,
		"clock_choice_alternative_branch_diagnostic":        true,
	}

	if !r.CandidateResidualComputed {
		if len(r.ResidualNullComparisons) > 0 {
			fail("residual_null_comparisons", "residual_null_comparisons must be empty when candidate_residual_computed is false")
		}
		for idx, cl := range r.CalculationLedger {
			if cl.CalculationStatus == "computed_from_local_branch" {
				fail("calculation_ledger", fmt.Sprintf("cannot have computed_from_local_branch calculation status when candidate_residual_computed is false at calculation_ledger[%d]", idx))
			}
		}
	} else {
		for idx, comp := range r.ResidualNullComparisons {
			if !comp.ComparisonComputed {
				isComp := comp.InterpretationStatus
				if isComp != InterpretBlockedByNoComparableNullDiagnostics &&
					isComp != InterpretBlockedByConventionDebt &&
					isComp != InterpretBlockedByClockFragility &&
					isComp != InterpretBlockedByNodeObstruction &&
					isComp != InterpretBlockedByNoEligibleLocalBranch &&
					isComp != InterpretBlockedByMissingResidualInputs {
					fail("residual_null_comparisons", fmt.Sprintf("interpretation_status must be a blocked status when comparison_computed is false at residual_null_comparisons[%d]", idx))
				}
			}
		}
	}

	// Validation of Calculation Ledger fields for transparency
	for idx, cl := range r.CalculationLedger {
		if cl.CalculationID == "" {
			fail("calculation_ledger", fmt.Sprintf("calculation_id cannot be empty at calculation_ledger[%d]", idx))
		}
		if cl.BranchID == "" {
			fail("calculation_ledger", fmt.Sprintf("branch_id cannot be empty at calculation_ledger[%d]", idx))
		}
		if cl.FormulaID == "" {
			fail("calculation_ledger", fmt.Sprintf("formula_id cannot be empty at calculation_ledger[%d]", idx))
		}
		forbiddenFormulaIDs := map[string]bool{
			"friedmann_residual":          true,
			"classical_residual":          true,
			"recovery_residual":           true,
			"cosmology_recovery_residual": true,
		}
		if forbiddenFormulaIDs[strings.ToLower(cl.FormulaID)] {
			fail("calculation_ledger", fmt.Sprintf("forbidden formula_id at calculation_ledger[%d]", idx))
		}
		if cl.FormulaDescription == "" {
			fail("calculation_ledger", fmt.Sprintf("formula_description cannot be empty at calculation_ledger[%d]", idx))
		}
		if cl.ConventionProfile == "" {
			fail("calculation_ledger", fmt.Sprintf("convention_profile cannot be empty at calculation_ledger[%d]", idx))
		}
		if len(cl.InputFields) == 0 {
			fail("calculation_ledger", fmt.Sprintf("input_fields cannot be empty at calculation_ledger[%d]", idx))
		}
		if cl.CalculationStatus == "" {
			fail("calculation_ledger", fmt.Sprintf("calculation_status cannot be empty at calculation_ledger[%d]", idx))
		}

		if cl.CalculationStatus == "computed_from_local_branch" {
			if cl.InputProvenance != "file_read" && cl.InputProvenance != "derived_from_file_read" {
				fail("calculation_ledger", fmt.Sprintf("invalid input_provenance for computed entry at calculation_ledger[%d]", idx))
			}
			if cl.NumInputPoints <= 0 {
				fail("calculation_ledger", fmt.Sprintf("num_input_points must be greater than zero for computed entry at calculation_ledger[%d]", idx))
			}
			if cl.FormulaSource == "" {
				fail("calculation_ledger", fmt.Sprintf("formula_source cannot be empty for computed entry at calculation_ledger[%d]", idx))
			}
			notesLower := strings.ToLower(cl.Notes)
			terms := []string{"winner", "outperform", "validated", "recovered", "proved", "superiority"}
			for _, t := range terms {
				if strings.Contains(notesLower, t) {
					fail("calculation_ledger", fmt.Sprintf("victory or recovery language detected in calculation ledger notes at calculation_ledger[%d]", idx))
				}
			}
		}
	}

	for idx, comp := range r.ResidualNullComparisons {
		if comp.ComparisonID == "" {
			fail("residual_null_comparisons", fmt.Sprintf("comparison_id cannot be empty at residual_null_comparisons[%d]", idx))
		}
		if comp.ComparisonComputed {
			if len(comp.TargetResidualIDs) == 0 {
				fail("residual_null_comparisons", fmt.Sprintf("target_residual_ids cannot be empty at residual_null_comparisons[%d]", idx))
			}
			if len(comp.NullModelIDs) == 0 {
				fail("residual_null_comparisons", fmt.Sprintf("null_model_ids cannot be empty at residual_null_comparisons[%d]", idx))
			}
			if len(comp.MetricsCompared) == 0 {
				fail("residual_null_comparisons", fmt.Sprintf("metrics_compared cannot be empty at residual_null_comparisons[%d]", idx))
			}

			// Validate target residuals integrity (each must exist and be computed)
			for _, trid := range comp.TargetResidualIDs {
				foundTarget := false
				for _, rd := range r.CandidateResidualDiagnostics {
					if rd.ResidualID == trid {
						foundTarget = true
						if !rd.ResidualComputed {
							fail("residual_null_comparisons", fmt.Sprintf("referenced target residual %q must have residual_computed = true at residual_null_comparisons[%d]", trid, idx))
						}
					}
				}
				if !foundTarget {
					fail("residual_null_comparisons", fmt.Sprintf("referenced target residual %q does not exist at residual_null_comparisons[%d]", trid, idx))
				}
			}

			// Reject blocked/deferred null references
			for _, nmid := range comp.NullModelIDs {
				if blockedOrDeferredNulls[nmid] {
					fail("residual_null_comparisons", fmt.Sprintf("cannot compare against blocked/deferred null model %q at residual_null_comparisons[%d]", nmid, idx))
				}
			}
		}

		// Valid interpretation status check (phrase-safe)
		isComp := comp.InterpretationStatus
		if isComp != InterpretDiagComparisonOnly && isComp != InterpretMixedResidualDiagnostics &&
			isComp != InterpretInsufficientTargetNullSeparation && isComp != InterpretTargetNullResidualSeparationCandidate &&
			isComp != InterpretBlockedByNoComparableNullDiagnostics && isComp != InterpretBlockedByConventionDebt &&
			isComp != InterpretBlockedByClockFragility && isComp != InterpretBlockedByNodeObstruction &&
			isComp != InterpretBlockedByNoEligibleLocalBranch {
			fail("residual_null_comparisons", fmt.Sprintf("invalid interpretation_status at residual_null_comparisons[%d]", idx))
		}

		// Reject victory language in comparison text fields
		checkVictory := func(val, path string) {
			valLower := strings.ToLower(val)
			terms := []string{"winner", "outperform", "validated", "recovered", "proved"}
			for _, t := range terms {
				if strings.Contains(valLower, t) {
					fail(path, "victory or superiority language detected")
				}
			}
		}
		checkVictory(comp.Reason, fmt.Sprintf("residual_null_comparisons[%d].reason", idx))
		checkVictory(comp.InterpretationStatus, fmt.Sprintf("residual_null_comparisons[%d].interpretation_status", idx))
	}

	// 9. Required Gates Validation
	requiredGates := map[string]bool{
		"toy_analysis_only_gate":              true,
		"no_final_truth_claim_gate":           true,
		"candidate_residual_diagnostics_gate": true,
		"no_recovery_claim_gate":              true,
		"no_scientific_novelty_claim_gate":    true,
		"no_bmc_beats_null_models_claim_gate": true,
		"full_bmc_blocked_gate":               true,
		"convention_debts_visible_gate":       true,
		"faithfulness_contested_gate":         true,
		"no_full_bmc_promotion_gate":          true,
	}
	gateCounts := make(map[string]int)
	for idx, g := range r.Gates {
		if g.Name == "" {
			fail("gates", fmt.Sprintf("invalid gate name at gates[%d]", idx))
		}
		gateCounts[g.Name]++
		if !requiredGates[g.Name] {
			fail("gates", fmt.Sprintf("unknown gate detected at gates[%d]", idx))
		}
		if g.Status != "pass" {
			fail("gates", fmt.Sprintf("gate status must be pass at gates[%d]", idx))
		}
	}
	for gName := range requiredGates {
		count := gateCounts[gName]
		if count == 0 {
			fail("gates", fmt.Sprintf("missing required gate: %s", gName))
		} else if count > 1 {
			fail("gates", fmt.Sprintf("duplicated gate detected: %s", gName))
		}
	}

	// 10. EBP Debt Status Validation
	debtVal := r.EbpDebt
	validDebtStatuses := map[string]bool{
		"unpaid":      true,
		"partial":     true,
		"retired":     true,
		"contested":   true,
		"overclaimed": true,
		"absent":      true,
		"candidate_residual_runner_candidate_only": true,
	}
	checkDebtField := func(field, val string) {
		if !validDebtStatuses[val] {
			fail(field, fmt.Sprintf("invalid status in %s", field))
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

	if debtVal.ContainsFinalTruthClaim != "absent" {
		fail("ebp_debt.containsFinalTruthClaim", "containsFinalTruthClaim must be absent")
	}
	if debtVal.NeedFaithfulnessReview != "contested" {
		fail("ebp_debt.needFaithfulnessReview", "needFaithfulnessReview must be contested")
	}
	if debtVal.ClockChoiceDebt != "unpaid" {
		fail("ebp_debt.clock_choice_debt", "clock_choice_debt must be unpaid")
	}
	if debtVal.ClassicalTargetDebt != "unpaid" {
		fail("ebp_debt.classical_target_debt", "classical_target_debt must be unpaid")
	}
	if debtVal.UnitConventionDebt != "unpaid" {
		fail("ebp_debt.unit_convention_debt", "unit_convention_debt must be unpaid")
	}
	if debtVal.SignConventionDebt != "unpaid" {
		fail("ebp_debt.sign_convention_debt", "sign_convention_debt must be unpaid")
	}
	if debtVal.NormalizationDebt != "unpaid" {
		fail("ebp_debt.normalization_debt", "normalization_debt must be unpaid")
	}
	if debtVal.PromotionStatus != "candidate_residual_runner_candidate_only" {
		fail("ebp_debt.promotion_status", "promotion_status must be candidate_residual_runner_candidate_only")
	}

	// 11. Warnings Validation
	hasNoRecoveryWarn := false
	hasFullBmcBlockedWarn := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No recovery claim is made.") {
			hasNoRecoveryWarn = true
		}
		if strings.Contains(w, "Full BMC remains blocked.") {
			hasFullBmcBlockedWarn = true
		}
	}
	if !hasNoRecoveryWarn {
		fail("warnings", "missing warning: “No recovery claim is made.”")
	}
	if !hasFullBmcBlockedWarn {
		fail("warnings", "missing warning: “Full BMC remains blocked.”")
	}

	// 12. Case-Insensitive Forbidden Phrase Scan (Phrase-Safe)
	checkForbidden := func(val, loc string) {
		valLower := strings.ToLower(val)
		for _, phrase := range forbiddenPhrases {
			if strings.Contains(valLower, phrase) {
				fail(loc, "forbidden phrase detected in report metadata")
			}
		}
	}

	for idx, w := range r.Warnings {
		checkForbidden(w, fmt.Sprintf("warning[%d]", idx))
	}
	for idx, rd := range r.CandidateResidualDiagnostics {
		checkForbidden(rd.BlockedReason, fmt.Sprintf("candidate_residual_diagnostics[%d].blocked_reason", idx))
		checkForbidden(rd.Notes, fmt.Sprintf("candidate_residual_diagnostics[%d].notes", idx))
		for wIdx, dw := range rd.Metrics.DiagnosticWarnings {
			checkForbidden(dw, fmt.Sprintf("candidate_residual_diagnostics[%d].metrics.diagnostic_warnings[%d]", idx, wIdx))
		}
	}
	for idx, comp := range r.ResidualNullComparisons {
		checkForbidden(comp.Reason, fmt.Sprintf("residual_null_comparisons[%d].reason", idx))
	}

	// Double check raw JSON content for safety
	rawJSONLower := strings.ToLower(rawJSON)
	rawJSONClean := strings.ReplaceAll(rawJSONLower, "no_winner_claim_gate", "no_xxxxxx_claim_gate")
	rawJSONClean = strings.ReplaceAll(rawJSONClean, "no_bmc_beats_null_models_claim_gate", "no_bmc_beats_xxxxxx_claim_gate")
	rawJSONClean = strings.ReplaceAll(rawJSONClean, "confirms no bmc-beats-null-models claim is made.", "confirms no bmc-beats-xxxxxx claim is made.")
	rawJSONClean = strings.ReplaceAll(rawJSONClean, "confirms no winner claims are made.", "confirms no xxxxxx claims are made.")
	for _, phrase := range forbiddenPhrases {
		if strings.Contains(rawJSONClean, phrase) {
			fail("json_content", "forbidden phrase detected in JSON content")
		}
	}

	return errs
}
