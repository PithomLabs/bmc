package priorart

import (
	"fmt"
	"strings"

	"github.com/PithomLabs/bmc/internal/bmc/report"
)

var forbiddenPhrases = []string{
	"first ever",
	"scientifically novel",
	"scientifically original",
	"breakthrough",
	"unprecedented",
	"bmc is new physics",
	"bmc proves bohmian cosmology",
	"bmc validates quantum gravity",
	"friedmann recovery",
	"recovery of friedmann",
	"ready for recovery",
	"full bmc validated",
	"problem of time solved",
}

// ValidateReport performs all strict Sprint 8-Lite validation checks.
func ValidateReport(r *PriorArtBoundaryReport, rawJSON string) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// 1. Schema, kind, scope, and ebp debt vocabulary
	if r.SchemaVersion != "bmc0a-prior-art-boundary-v0.1" {
		fail("schema_version", "unsupported schema version")
	}
	if r.EbpDebtVocabulary != "ptw_runtime_debt_status_v0.1" {
		fail("ebp_debt_vocabulary", "invalid ebp_debt_vocabulary detected")
	}
	if r.ArtifactKind != "prior_art_boundary_note" {
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

	// 3. Safety and novelty flags
	if r.ScientificNoveltyClaimMade {
		fail("scientific_novelty_claim_made", "scientific_novelty_claim_made must be false")
	}
	if r.ScientificNoveltyClaimAllowed {
		fail("scientific_novelty_claim_allowed", "scientific_novelty_claim_allowed must be false")
	}
	if r.WorkflowDistinctivenessStatus != "candidate_only" {
		fail("workflow_distinctiveness_status", "workflow_distinctiveness_status must be candidate_only")
	}
	if r.ResidualComputed {
		fail("residual_computed", "residual_computed must be false")
	}
	if r.NullComparisonComputed {
		fail("null_comparison_computed", "null_comparison_computed must be false")
	}
	if r.RecoveryClaim {
		fail("recovery_claim", "recovery_claim must be false")
	}
	if r.FullBmcToyGate != "blocked" {
		fail("full_bmc_toy_gate", "full_bmc_toy_gate must be blocked")
	}

	// 4. Sources validation
	if len(r.PriorArtSources) == 0 {
		fail("prior_art_sources", "prior_art_sources cannot be empty")
	}
	seenSourceIDs := make(map[string]bool)
	requiredSeedSources := map[string]bool{
		"bohmian_quantum_cosmology_review":                 true,
		"bohmian_quantum_gravity_cosmology_review":          true,
		"scalar_field_minisuperspace_classical_limit_paper": true,
		"gaussian_superposition_bohmian_trajectory_paper":   true,
		"modern_friedmann_scalar_field_bohmian_example":     true,
	}
	for idx, src := range r.PriorArtSources {
		if src.SourceID == "" {
			fail("prior_art_sources", "source_id cannot be empty")
		}
		if seenSourceIDs[src.SourceID] {
			fail("prior_art_sources", "duplicate source_id detected")
		}
		seenSourceIDs[src.SourceID] = true

		// Source kind validation (phrase-safe)
		if src.SourceKind != SourceKindReview && src.SourceKind != SourceKindPaper &&
			src.SourceKind != SourceKindBook && src.SourceKind != SourceKindUnknown {
			fail("prior_art_sources", fmt.Sprintf("invalid source_kind detected at prior_art_sources[%d]", idx))
		}

		// Review status validation (phrase-safe)
		if src.ReviewStatus != ReviewStatusSeedUnreviewed &&
			src.ReviewStatus != ReviewStatusAbstractReviewed &&
			src.ReviewStatus != ReviewStatusSkimReviewed &&
			src.ReviewStatus != ReviewStatusHumanReviewRequired {
			fail("prior_art_sources", fmt.Sprintf("invalid review_status detected at prior_art_sources[%d]", idx))
		}
		// Forbidden checks explicitly (phrase-safe)
		if src.ReviewStatus == "proves_our_claim" || src.ReviewStatus == "confirms_novelty" ||
			src.ReviewStatus == "settled" || src.ReviewStatus == "definitive" {
			fail("prior_art_sources", fmt.Sprintf("forbidden review_status detected at prior_art_sources[%d]", idx))
		}
	}

	for id := range requiredSeedSources {
		if !seenSourceIDs[id] {
			fail("prior_art_sources", "missing required seed source ID")
		}
	}

	// 5. Boundary claims validation
	if len(r.BoundaryClaims) == 0 {
		fail("boundary_claims", "boundary_claims cannot be empty")
	}
	seenClaimIDs := make(map[string]bool)
	requiredBoundaryClaims := map[string]bool{
		"bmc_uses_bohmian_minisuperspace_trajectories": true,
		"bmc_uses_wheeler_dewitt_toy_equation":         true,
		"bmc_uses_scalar_field_clock_checks":           true,
		"bmc_uses_quantum_potential_diagnostics":       true,
		"bmc_uses_node_obstruction_detection":          true,
		"bmc_uses_local_branch_segmentation":           true,
		"bmc_defines_null_model_scaffold":              true,
		"bmc_uses_ebp_debt_gates":                      true,
		"bmc_uses_lean_policy_contracts":               true,
		"bmc_does_not_claim_recovery":                  true,
		"bmc_does_not_claim_scientific_novelty":        true,
	}

	// Claim-specific boundary status mapping
	allowedBoundaryStatusMap := map[string]map[string]bool{
		"bmc_uses_bohmian_minisuperspace_trajectories": {
			BoundaryStatusEstablishedPriorArt: true,
			BoundaryStatusLikelyPriorArt:      true,
		},
		"bmc_uses_wheeler_dewitt_toy_equation": {
			BoundaryStatusEstablishedPriorArt: true,
			BoundaryStatusLikelyPriorArt:      true,
		},
		"bmc_uses_scalar_field_clock_checks": {
			BoundaryStatusLikelyPriorArt:        true,
			BoundaryStatusUnknownRequiresReview: true,
		},
		"bmc_uses_quantum_potential_diagnostics": {
			BoundaryStatusEstablishedPriorArt: true,
			BoundaryStatusLikelyPriorArt:      true,
		},
		"bmc_uses_node_obstruction_detection": {
			BoundaryStatusImplementationVariant: true,
			BoundaryStatusLikelyPriorArt:        true,
			BoundaryStatusUnknownRequiresReview: true,
		},
		"bmc_uses_local_branch_segmentation": {
			BoundaryStatusImplementationVariant: true,
			BoundaryStatusUnknownRequiresReview: true,
		},
		"bmc_defines_null_model_scaffold": {
			BoundaryStatusWorkflowDistinctiveCandidate: true,
			BoundaryStatusUnknownRequiresReview:        true,
		},
		"bmc_uses_ebp_debt_gates": {
			BoundaryStatusWorkflowDistinctiveCandidate: true,
			BoundaryStatusUnknownRequiresReview:        true,
		},
		"bmc_uses_lean_policy_contracts": {
			BoundaryStatusWorkflowDistinctiveCandidate: true,
			BoundaryStatusUnknownRequiresReview:        true,
		},
		"bmc_does_not_claim_recovery": {
			BoundaryStatusNotClaimed: true,
		},
		"bmc_does_not_claim_scientific_novelty": {
			BoundaryStatusBlocked: true,
		},
	}

	for idx, clm := range r.BoundaryClaims {
		if clm.ClaimID == "" {
			fail("boundary_claims", "claim_id cannot be empty")
		}
		if seenClaimIDs[clm.ClaimID] {
			fail("boundary_claims", "duplicate claim_id detected")
		}
		seenClaimIDs[clm.ClaimID] = true

		// Boundary status validation (phrase-safe)
		if clm.BoundaryStatus != BoundaryStatusEstablishedPriorArt &&
			clm.BoundaryStatus != BoundaryStatusLikelyPriorArt &&
			clm.BoundaryStatus != BoundaryStatusImplementationVariant &&
			clm.BoundaryStatus != BoundaryStatusWorkflowDistinctiveCandidate &&
			clm.BoundaryStatus != BoundaryStatusUnknownRequiresReview &&
			clm.BoundaryStatus != BoundaryStatusBlocked &&
			clm.BoundaryStatus != BoundaryStatusNotClaimed {
			fail("boundary_claims", fmt.Sprintf("invalid boundary_status detected at boundary_claims[%d]", idx))
		}

		// Explicit forbidden checks (phrase-safe)
		bs := clm.BoundaryStatus
		if bs == "novel" || bs == "first_ever" || bs == "scientifically_original" ||
			bs == "breakthrough" || bs == "proved_new" || bs == "validated_physics" {
			fail("boundary_claims", fmt.Sprintf("forbidden boundary_status detected at boundary_claims[%d]", idx))
		}

		// Enforce claim-specific boundary-status direction
		allowedMap, exists := allowedBoundaryStatusMap[clm.ClaimID]
		if exists && !allowedMap[clm.BoundaryStatus] {
			fail("boundary_claims", fmt.Sprintf("invalid boundary_status direction detected at boundary_claims[%d]", idx))
		}
	}

	for id := range requiredBoundaryClaims {
		if !seenClaimIDs[id] {
			fail("boundary_claims", "missing required boundary claim ID")
		}
	}

	// 6. Gates validation
	requiredGates := map[string]bool{
		"toy_analysis_only_gate":           true,
		"no_final_truth_claim_gate":        true,
		"no_scientific_novelty_claim_gate": true,
		"prior_art_sources_seeded_gate":    true,
		"boundary_claims_declared_gate":    true,
		"no_residual_computation_gate":     true,
		"no_null_comparison_result_gate":   true,
		"no_recovery_claim_gate":           true,
		"full_bmc_blocked_gate":            true,
		"faithfulness_contested_gate":      true,
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

	// 7. EBP Debt checks
	if r.EbpDebt.ContainsFinalTruthClaim != "absent" {
		fail("ebp_debt.containsFinalTruthClaim", "containsFinalTruthClaim must be absent")
	}
	if r.EbpDebt.NeedFaithfulnessReview != "contested" {
		fail("ebp_debt.needFaithfulnessReview", "needFaithfulnessReview must be contested")
	}

	// 8. Warnings validation
	hasNoveltyWarn := false
	hasBlockedWarn := false
	for _, w := range r.Warnings {
		if strings.Contains(w, "No scientific novelty claim is made.") {
			hasNoveltyWarn = true
		}
		if strings.Contains(w, "Full BMC remains blocked.") {
			hasBlockedWarn = true
		}
	}
	if !hasNoveltyWarn {
		fail("warnings", "missing warning: “No scientific novelty claim is made.”")
	}
	if !hasBlockedWarn {
		fail("warnings", "missing warning: “Full BMC remains blocked.”")
	}

	// 9. Case-insensitive forbidden phrase scanner (phrase-safe error messaging)
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
	for idx, src := range r.PriorArtSources {
		checkForbidden(src.Title, fmt.Sprintf("prior_art_sources[%d].title", idx))
		checkForbidden(src.BoundaryUse, fmt.Sprintf("prior_art_sources[%d].boundary_use", idx))
		for aIdx, author := range src.Authors {
			checkForbidden(author, fmt.Sprintf("prior_art_sources[%d].authors[%d]", idx, aIdx))
		}
		for tIdx, tag := range src.RelevanceTags {
			checkForbidden(tag, fmt.Sprintf("prior_art_sources[%d].relevance_tags[%d]", idx, tIdx))
		}
	}
	for idx, clm := range r.BoundaryClaims {
		checkForbidden(clm.ClaimText, fmt.Sprintf("boundary_claims[%d].claim_text", idx))
		checkForbidden(clm.Reason, fmt.Sprintf("boundary_claims[%d].reason", idx))
	}

	// Double check raw JSON content for safety
	rawJSONLower := strings.ToLower(rawJSON)
	for _, phrase := range forbiddenPhrases {
		if strings.Contains(rawJSONLower, phrase) {
			fail("json_content", "forbidden phrase detected in JSON content")
		}
	}

	return errs
}
