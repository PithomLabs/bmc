package priorart

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type PriorArtSource struct {
	SourceID      string   `json:"source_id"`
	Title         string   `json:"title"`
	Authors       []string `json:"authors"`
	Year          int      `json:"year"`
	SourceKind    string   `json:"source_kind"`
	ReviewStatus  string   `json:"review_status"`
	RelevanceTags []string `json:"relevance_tags"`
	BoundaryUse   string   `json:"boundary_use"`
}

type BoundaryClaim struct {
	ClaimID             string   `json:"claim_id"`
	ClaimText           string   `json:"claim_text"`
	BoundaryStatus      string   `json:"boundary_status"`
	RelatedSourceIDs    []string `json:"related_source_ids"`
	Reason              string   `json:"reason"`
	HumanReviewRequired bool     `json:"human_review_required"`
}

type PriorArtGate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type EbpDebt struct {
	NeedLiteratureAudit     string `json:"needLiteratureAudit"`
	NeedMap                 string `json:"needMap"`
	NeedInvariant           string `json:"needInvariant"`
	NeedToyCheck            string `json:"needToyCheck"`
	NeedNullModel           string `json:"needNullModel"`
	NeedObstruction         string `json:"needObstruction"`
	NeedFaithfulnessReview  string `json:"needFaithfulnessReview"`
	ClockChoiceDebt         string `json:"clock_choice_debt"`
	ClassicalTargetDebt     string `json:"classical_target_debt"`
	UnitConventionDebt      string `json:"unit_convention_debt"`
	SignConventionDebt      string `json:"sign_convention_debt"`
	NormalizationDebt       string `json:"normalization_debt"`
	ContainsFinalTruthClaim string `json:"containsFinalTruthClaim"`
	PromotionStatus         string `json:"promotion_status"`
}

type PriorArtBoundaryReport struct {
	SchemaVersion                 string           `json:"schema_version"`
	EbpDebtVocabulary             string           `json:"ebp_debt_vocabulary"`
	ToyAnalysisOnly               bool             `json:"toy_analysis_only"`
	FinalTruthClaim               bool             `json:"final_truth_claim"`
	ArtifactKind                  string           `json:"artifact_kind"`
	Scope                         string           `json:"scope"`
	ScientificNoveltyClaimMade    bool             `json:"scientific_novelty_claim_made"`
	ScientificNoveltyClaimAllowed bool             `json:"scientific_novelty_claim_allowed"`
	WorkflowDistinctivenessStatus string           `json:"workflow_distinctiveness_status"`
	ResidualComputed              bool             `json:"residual_computed"`
	NullComparisonComputed        bool             `json:"null_comparison_computed"`
	RecoveryClaim                 bool             `json:"recovery_claim"`
	FullBmcToyGate                string           `json:"full_bmc_toy_gate"`
	PriorArtSources               []PriorArtSource `json:"prior_art_sources"`
	BoundaryClaims                []BoundaryClaim  `json:"boundary_claims"`
	Gates                         []PriorArtGate   `json:"gates"`
	EbpDebt                       EbpDebt          `json:"ebp_debt"`
	Warnings                      []string         `json:"warnings"`
}

// GenerateDefaultReport creates a deterministic prior-art boundary report.
func GenerateDefaultReport() *PriorArtBoundaryReport {
	sources := []PriorArtSource{
		{
			SourceID:      "bohmian_quantum_cosmology_review",
			Title:         "Bohmian Quantum Cosmology: A Review",
			Authors:       []string{"Author A"},
			Year:          2020,
			SourceKind:    SourceKindReview,
			ReviewStatus:  ReviewStatusSeedUnreviewed,
			RelevanceTags: []string{"bohmian", "cosmology"},
			BoundaryUse:   "Seed source for future review of Bohmian quantum-cosmology background.",
		},
		{
			SourceID:      "bohmian_quantum_gravity_cosmology_review",
			Title:         "Bohmian Quantum Gravity and Cosmology",
			Authors:       []string{"Author B"},
			Year:          2018,
			SourceKind:    SourceKindReview,
			ReviewStatus:  ReviewStatusSeedUnreviewed,
			RelevanceTags: []string{"quantum-gravity", "cosmology"},
			BoundaryUse:   "Seed source for future human review of Bohmian minisuperspace prior art.",
		},
		{
			SourceID:      "scalar_field_minisuperspace_classical_limit_paper",
			Title:         "On the Classical Limit of Scalar Field Minisuperspace",
			Authors:       []string{"Author C"},
			Year:          2015,
			SourceKind:    SourceKindPaper,
			ReviewStatus:  ReviewStatusSeedUnreviewed,
			RelevanceTags: []string{"scalar-field", "classical-limit"},
			BoundaryUse:   "Seed source for checking overlap with scalar-field minisuperspace claims.",
		},
		{
			SourceID:      "gaussian_superposition_bohmian_trajectory_paper",
			Title:         "Bohmian Trajectories for Gaussian Superpositions in Minisuperspace",
			Authors:       []string{"Author D"},
			Year:          2022,
			SourceKind:    SourceKindPaper,
			ReviewStatus:  ReviewStatusSeedUnreviewed,
			RelevanceTags: []string{"gaussian", "superposition", "trajectory"},
			BoundaryUse:   "Seed source for future human review of Bohmian minisuperspace prior art.",
		},
		{
			SourceID:      "modern_friedmann_scalar_field_bohmian_example",
			Title:         "Bohmian Trajectories in a Modern Friedmann Scalar Field Model",
			Authors:       []string{"Author E"},
			Year:          2024,
			SourceKind:    SourceKindPaper,
			ReviewStatus:  ReviewStatusSeedUnreviewed,
			RelevanceTags: []string{"friedmann", "bohmian", "scalar-field"},
			BoundaryUse:   "Seed source for checking overlap with scalar-field minisuperspace claims.",
		},
	}

	claims := []BoundaryClaim{
		{
			ClaimID:             "bmc_uses_bohmian_minisuperspace_trajectories",
			ClaimText:           "BMC uses Bohmian minisuperspace trajectories for quantum cosmological evolution.",
			BoundaryStatus:      BoundaryStatusEstablishedPriorArt,
			RelatedSourceIDs:    []string{"bohmian_quantum_cosmology_review", "bohmian_quantum_gravity_cosmology_review"},
			Reason:              "Standard minisuperspace trajectory formulation is well-documented in literature.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_wheeler_dewitt_toy_equation",
			ClaimText:           "BMC uses the Wheeler-DeWitt toy equation to model wavefunction evolution.",
			BoundaryStatus:      BoundaryStatusEstablishedPriorArt,
			RelatedSourceIDs:    []string{"bohmian_quantum_cosmology_review"},
			Reason:              "Wheeler-DeWitt equation is the standard prior art foundation for minisuperspace.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_scalar_field_clock_checks",
			ClaimText:           "BMC uses scalar field values as relational clocks to segment trajectories.",
			BoundaryStatus:      BoundaryStatusLikelyPriorArt,
			RelatedSourceIDs:    []string{"scalar_field_minisuperspace_classical_limit_paper"},
			Reason:              "Relational clock checking using scalar fields is widely discussed, but requires human review for implementation specifics.",
			HumanReviewRequired: true,
		},
		{
			ClaimID:             "bmc_uses_quantum_potential_diagnostics",
			ClaimText:           "BMC computes quantum potential values along trajectories.",
			BoundaryStatus:      BoundaryStatusEstablishedPriorArt,
			RelatedSourceIDs:    []string{"bohmian_quantum_gravity_cosmology_review"},
			Reason:              "Quantum potential is a core element of Bohmian mechanics prior art.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_node_obstruction_detection",
			ClaimText:           "BMC implements detection of node proximity to flag potential mathematical failures.",
			BoundaryStatus:      BoundaryStatusImplementationVariant,
			RelatedSourceIDs:    []string{"gaussian_superposition_bohmian_trajectory_paper"},
			Reason:              "Node detection is specific to our numerical implementation of trajectory integration.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_local_branch_segmentation",
			ClaimText:           "BMC segments trajectories locally to handle non-monotonic clock behavior.",
			BoundaryStatus:      BoundaryStatusImplementationVariant,
			RelatedSourceIDs:    []string{"gaussian_superposition_bohmian_trajectory_paper"},
			Reason:              "Branch segmentation is workflow/implementation specific, needs confirmation.",
			HumanReviewRequired: true,
		},
		{
			ClaimID:             "bmc_defines_null_model_scaffold",
			ClaimText:           "BMC implements a null-model validation scaffold to rule out trivial matching.",
			BoundaryStatus:      BoundaryStatusWorkflowDistinctiveCandidate,
			RelatedSourceIDs:    []string{},
			Reason:              "Scaffolding structure is specific to the BMC workflow verification process.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_ebp_debt_gates",
			ClaimText:           "BMC uses EBP debt gates to track unresolved physical and mathematical assumptions.",
			BoundaryStatus:      BoundaryStatusWorkflowDistinctiveCandidate,
			RelatedSourceIDs:    []string{},
			Reason:              "EBP debt tracking is unique to the workflow governance model.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_uses_lean_policy_contracts",
			ClaimText:           "BMC uses Lean policy contracts to formally prove safety constraints.",
			BoundaryStatus:      BoundaryStatusWorkflowDistinctiveCandidate,
			RelatedSourceIDs:    []string{},
			Reason:              "Lean policy proving is unique to the BMC workflow verification pipeline.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_does_not_claim_recovery",
			ClaimText:           "BMC does not claim to recover the classical Friedmann equations from quantum trajectories.",
			BoundaryStatus:      BoundaryStatusNotClaimed,
			RelatedSourceIDs:    []string{"modern_friedmann_scalar_field_bohmian_example"},
			Reason:              "Classical dynamics recovery is explicitly deferred; no recovery claim is made.",
			HumanReviewRequired: false,
		},
		{
			ClaimID:             "bmc_does_not_claim_scientific_novelty",
			ClaimText:           "BMC does not claim any scientific novelty or new physics.",
			BoundaryStatus:      BoundaryStatusBlocked,
			RelatedSourceIDs:    []string{},
			Reason:              "Scientific novelty claims are strictly blocked under current EBP status.",
			HumanReviewRequired: false,
		},
	}

	gates := []PriorArtGate{
		{"toy_analysis_only_gate", "pass", "Confirms analysis is strictly restricted to minisuperspace toy system."},
		{"no_final_truth_claim_gate", "pass", "Confirms no final truth claims are asserted."},
		{"no_scientific_novelty_claim_gate", "pass", "Confirms no scientific novelty claim is made."},
		{"prior_art_sources_seeded_gate", "pass", "Confirms prior-art sources are seeded."},
		{"boundary_claims_declared_gate", "pass", "Confirms boundary claims are declared."},
		{"no_residual_computation_gate", "pass", "Confirms that no actual numerical residual was computed."},
		{"no_null_comparison_result_gate", "pass", "Confirms that no null model comparison results are computed."},
		{"no_recovery_claim_gate", "pass", "Confirms no recovery claim is made."},
		{"full_bmc_blocked_gate", "pass", "Confirms the full BMC promotion gate is blocked."},
		{"faithfulness_contested_gate", "pass", "Confirms faithfulness review status remains contested."},
	}

	ebpDebt := EbpDebt{
		NeedLiteratureAudit:     "partial",
		NeedMap:                 "active",
		NeedInvariant:           "partial",
		NeedToyCheck:            "active",
		NeedNullModel:           "active",
		NeedObstruction:         "active",
		NeedFaithfulnessReview:  "contested",
		ClockChoiceDebt:         "active",
		ClassicalTargetDebt:     "active",
		UnitConventionDebt:      "active",
		SignConventionDebt:      "active",
		NormalizationDebt:       "active",
		ContainsFinalTruthClaim: "absent",
		PromotionStatus:         "prior_art_boundary_note_candidate_only",
	}

	warnings := []string{
		"This is a BMC-0A prior-art boundary note only.",
		"No scientific novelty claim is made.",
		"No recovery claim is made.",
		"No residual was computed.",
		"No null-model comparison was computed.",
		"Full BMC remains blocked.",
		"Runtime EBP debt labels are not adversarial-review classifications.",
	}

	return &PriorArtBoundaryReport{
		SchemaVersion:                 "bmc0a-prior-art-boundary-v0.1",
		EbpDebtVocabulary:             "ptw_runtime_debt_status_v0.1",
		ToyAnalysisOnly:               true,
		FinalTruthClaim:               false,
		ArtifactKind:                  "prior_art_boundary_note",
		Scope:                         "bmc0a_only",
		ScientificNoveltyClaimMade:    false,
		ScientificNoveltyClaimAllowed: false,
		WorkflowDistinctivenessStatus: "candidate_only",
		ResidualComputed:              false,
		NullComparisonComputed:        false,
		RecoveryClaim:                 false,
		FullBmcToyGate:                "blocked",
		PriorArtSources:               sources,
		BoundaryClaims:                claims,
		Gates:                         gates,
		EbpDebt:                       ebpDebt,
		Warnings:                      warnings,
	}
}

// ReadReport reads and strictly decodes a PriorArtBoundaryReport JSON file, rejecting trailing tokens.
func ReadReport(path string) (*PriorArtBoundaryReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep PriorArtBoundaryReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}

	var dummy json.RawMessage
	if err := dec.Decode(&dummy); err != io.EOF {
		return nil, fmt.Errorf("trailing garbage in JSON: %v", err)
	}

	return &rep, nil
}

// WriteReport writes the pretty-printed JSON report to a file.
func WriteReport(rep *PriorArtBoundaryReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}
