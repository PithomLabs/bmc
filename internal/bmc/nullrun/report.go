package nullrun

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type NullRunSeed struct {
	NullModelID string `json:"null_model_id"`
	Seed        uint64 `json:"seed"`
	RngKind     string `json:"rng_kind"`
}

type NullDiagnostics struct {
	NodeContactFree          bool     `json:"node_contact_free"`
	TrajectoryFiniteness     string   `json:"trajectory_finiteness"`
	NumValidTrajectoryPoints int      `json:"num_valid_trajectory_points"`
	NumClockSegments         int      `json:"num_clock_segments"`
	NumTurningPoints         int      `json:"num_turning_points"`
	PhiClockGlobalStatus     string   `json:"phi_clock_global_status"`
	LocalBranchStatus        string   `json:"local_branch_status"`
	MinAmplitudeR            *float64 `json:"min_amplitude_r"`
	MaxAbsQAwayFromNodes     *float64 `json:"max_abs_q_away_from_nodes"`
	MaxPhaseGradient         *float64 `json:"max_phase_gradient"`
	DiagnosticWarnings       []string `json:"diagnostic_warnings"`
}

type NullModelRun struct {
	NullModelID          string          `json:"null_model_id"`
	RunStatus            string          `json:"run_status"`
	DiagnosticProvenance string          `json:"diagnostic_provenance"`
	Seed                 *NullRunSeed    `json:"seed,omitempty"`
	DiagnosticStatus     string          `json:"diagnostic_status"`
	Diagnostics          NullDiagnostics `json:"diagnostics"`
	BlockedReason        string          `json:"blocked_reason,omitempty"`
	Notes                string          `json:"notes"`
}

type TargetNullDiagnosticComparison struct {
	ComparisonID         string   `json:"comparison_id"`
	TargetArtifact       string   `json:"target_artifact"`
	NullModelIDs         []string `json:"null_model_ids"`
	MetricsCompared      []string `json:"metrics_compared"`
	ComparisonComputed   bool     `json:"comparison_computed"`
	InterpretationStatus string   `json:"interpretation_status"`
	Reason               string   `json:"reason"`
}

type NullRunGate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type NullRunEbpDebt struct {
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

type NullRunReport struct {
	SchemaVersion                   string                           `json:"schema_version"`
	ToyAnalysisOnly                 bool                             `json:"toy_analysis_only"`
	FinalTruthClaim                 bool                             `json:"final_truth_claim"`
	ArtifactKind                    string                           `json:"artifact_kind"`
	Scope                           string                           `json:"scope"`
	ResidualComputed                bool                             `json:"residual_computed"`
	NullDiagnosticsComputed        bool                             `json:"null_diagnostics_computed"`
	TargetNullComparisonComputed    bool                             `json:"target_null_comparison_computed"`
	RecoveryClaim                   bool                             `json:"recovery_claim"`
	ScientificNoveltyClaimMade      bool                             `json:"scientific_novelty_claim_made"`
	FullBmcToyGate                  string                           `json:"full_bmc_toy_gate"`
	SourceArtifacts                 []string                         `json:"source_artifacts"`
	NullModelRuns                   []NullModelRun                   `json:"null_model_runs"`
	TargetNullDiagnosticComparisons []TargetNullDiagnosticComparison `json:"target_null_diagnostic_comparisons"`
	Gates                           []NullRunGate                    `json:"gates"`
	EbpDebtVocabulary               string                           `json:"ebp_debt_vocabulary"`
	InterpretationStatus            string                           `json:"interpretation_status"`
	EbpDebt                         NullRunEbpDebt                   `json:"ebp_debt"`
	Warnings                        []string                         `json:"warnings"`
}

// GenerateDefaultReport creates a deterministic null-model runner report.
func GenerateDefaultReport() *NullRunReport {
	// Sample float parameters
	valMinR := 1.2e-3
	valMaxQ := 4.5
	valMaxPG := 8.9

	runs := []NullModelRun{
		{
			NullModelID:          "constant_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "constant_phase_control",
				Seed:        42,
				RngKind:     "none",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         1,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "monotonic",
				LocalBranchStatus:        "stable",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Phase is held constant across coordinates; acts as simplest static control.",
		},
		{
			NullModelID:          "randomized_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "randomized_phase_control",
				Seed:        12345,
				RngKind:     "xorshift64",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         2,
				NumTurningPoints:         1,
				PhiClockGlobalStatus:     "non_monotonic",
				LocalBranchStatus:        "fragile",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{"Phase fluctuations degrade relational monotonicity."},
			},
			Notes: "Phase values randomized at coordinates with recorded seed.",
		},
		{
			NullModelID:          "matched_amplitude_randomized_phase_control",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			Seed: &NullRunSeed{
				NullModelID: "matched_amplitude_randomized_phase_control",
				Seed:        54321,
				RngKind:     "xorshift64",
			},
			DiagnosticStatus: DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         2,
				NumTurningPoints:         1,
				PhiClockGlobalStatus:     "non_monotonic",
				LocalBranchStatus:        "fragile",
				MinAmplitudeR:            &valMinR,
				MaxAbsQAwayFromNodes:     &valMaxQ,
				MaxPhaseGradient:         &valMaxPG,
				DiagnosticWarnings:       []string{"Phase fluctuations degrade relational monotonicity."},
			},
			Notes: "Amplitude is matched to target superposition wave while phase is randomized.",
		},
		{
			NullModelID:          "classical_frw_reference_trajectory",
			RunStatus:            RunStatusDiagnosticsGenerated,
			DiagnosticProvenance: ProvenanceDeterministicFixture,
			DiagnosticStatus:     DiagStatusFinite,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          true,
				TrajectoryFiniteness:     "finite",
				NumValidTrajectoryPoints: 100,
				NumClockSegments:         1,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "monotonic",
				LocalBranchStatus:        "stable",
				MinAmplitudeR:            nil, // unavailable for classical reference
				MaxAbsQAwayFromNodes:     nil, // unavailable for classical reference
				MaxPhaseGradient:         nil, // unavailable for classical reference
				DiagnosticWarnings:       []string{},
			},
			Notes: "reference comparator only; no residual or recovery interpretation",
		},
		{
			NullModelID:          "same_branch_segmentation_under_null_wavefunctions",
			RunStatus:            RunStatusDeferred,
			DiagnosticProvenance: ProvenanceDeferred,
			DiagnosticStatus:     DiagStatusNotAvailable,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "not_available",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Deferred until full relational clock segmentation is active on null sets.",
		},
		{
			NullModelID:          "node_neighborhood_stress_case",
			RunStatus:            RunStatusBlocked,
			DiagnosticProvenance: ProvenanceBlocked,
			BlockedReason:        "Blocked near nodes due to infinite quantum potential and division by zero risks.",
			DiagnosticStatus:     DiagStatusNodeBlocked,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "node_blocked",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Unable to generate trajectories inside node neighborhoods.",
		},
		{
			NullModelID:          "clock_choice_alternative_branch_diagnostic",
			RunStatus:            RunStatusDeferred,
			DiagnosticProvenance: ProvenanceDeferred,
			DiagnosticStatus:     DiagStatusNotAvailable,
			Diagnostics: NullDiagnostics{
				NodeContactFree:          false,
				TrajectoryFiniteness:     "not_available",
				NumValidTrajectoryPoints: 0,
				NumClockSegments:         0,
				NumTurningPoints:         0,
				PhiClockGlobalStatus:     "not_available",
				LocalBranchStatus:        "not_available",
				MinAmplitudeR:            nil,
				MaxAbsQAwayFromNodes:     nil,
				MaxPhaseGradient:         nil,
				DiagnosticWarnings:       []string{},
			},
			Notes: "Deferred until alternative relational clocks (e.g. scale factor a) are integrated in nullspec.",
		},
	}

	comparisons := []TargetNullDiagnosticComparison{
		{
			ComparisonID:         "bmc0a-superposition-vs-constant-phase-v0.1",
			TargetArtifact:       "Sprint 2 superposition control",
			NullModelIDs:         []string{"constant_phase_control"},
			MetricsCompared:      []string{"node_contact_free", "phi_clock_global_status", "num_clock_segments"},
			ComparisonComputed:   true,
			InterpretationStatus: InterpretDiagComparisonOnly,
			Reason:               "Comparison shows stable clock metrics for target under static phase fields.",
		},
	}

	gates := []NullRunGate{
		{"toy_analysis_only_gate", "pass", "Confirms analysis is strictly restricted to minisuperspace toy system."},
		{"no_final_truth_claim_gate", "pass", "Confirms no final truth claims are asserted."},
		{"no_residual_computation_gate", "pass", "Confirms that no actual numerical residual was computed."},
		{"null_diagnostics_computed_gate", "pass", "Confirms null model diagnostics are generated."},
		{"target_null_comparison_computed_gate", "pass", "Confirms target/null comparison is generated."},
		{"no_winner_claim_gate", "pass", "Confirms no winner claims are made."},
		{"no_recovery_claim_gate", "pass", "Confirms no recovery claim is made."},
		{"no_scientific_novelty_claim_gate", "pass", "Confirms no scientific novelty claim is made."},
		{"full_bmc_blocked_gate", "pass", "Confirms the full BMC promotion gate is blocked."},
		{"faithfulness_contested_gate", "pass", "Confirms faithfulness review status remains contested."},
	}

	ebpDebt := NullRunEbpDebt{
		NeedLiteratureAudit:     "partial",
		NeedMap:                 "partial",
		NeedInvariant:           "partial",
		NeedToyCheck:            "unpaid",
		NeedNullModel:           "partial", // Promoted to partial after Sprint 9 null model diagnostics run
		NeedObstruction:         "partial",
		NeedFaithfulnessReview:  "contested",
		ClockChoiceDebt:         "unpaid",
		ClassicalTargetDebt:     "unpaid",
		UnitConventionDebt:      "unpaid",
		SignConventionDebt:      "unpaid",
		NormalizationDebt:       "unpaid",
		ContainsFinalTruthClaim: "absent",
		PromotionStatus:         "null_model_runner_candidate_only",
	}

	warnings := []string{
		"Sprint 9 computes null-model diagnostics only.",
		"No recovery claim is made.",
		"No residual was computed.",
		"No scientific novelty claim is made.",
		"Full BMC remains blocked.",
		"EBP debt labels conform to the adversarial-review classification vocabulary.",
	}

	return &NullRunReport{
		SchemaVersion:                "bmc0a-nullrun-v0.1",
		ToyAnalysisOnly:              true,
		FinalTruthClaim:              false,
		ArtifactKind:                 "null_model_runner_report",
		Scope:                        "bmc0a_only",
		ResidualComputed:             false,
		NullDiagnosticsComputed:     true,
		TargetNullComparisonComputed: true,
		RecoveryClaim:                false,
		ScientificNoveltyClaimMade:   false,
		FullBmcToyGate:               "blocked",
		SourceArtifacts: []string{
			"Sprint 7 Null-Model scaffold",
		},
		NullModelRuns:                   runs,
		TargetNullDiagnosticComparisons: comparisons,
		Gates:                           gates,
		EbpDebtVocabulary:               "ptw_adversarial_review_debt_status_v0.1",
		InterpretationStatus:            InterpretDiagComparisonOnly,
		EbpDebt:                         ebpDebt,
		Warnings:                        warnings,
	}
}

// ReadReport reads and strictly decodes a NullRunReport JSON file, rejecting trailing tokens.
func ReadReport(path string) (*NullRunReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep NullRunReport
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
func WriteReport(rep *NullRunReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}
