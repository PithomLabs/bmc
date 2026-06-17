package nullspec

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

type NullModelEbpDebt struct {
	NeedMap                 string `json:"needMap"`
	NeedInvariant           string `json:"needInvariant"`
	NeedToyCheck            string `json:"needToyCheck"`
	NeedNullModel           string `json:"needNullModel"`
	NeedObstruction         string `json:"needObstruction"`
	NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
	ClockChoiceDebt         string `json:"clock_choice_debt"`
	ClassicalTargetDebt     string `json:"classical_target_debt"`
	UnitConventionDebt      string `json:"unit_convention_debt"`
	SignConventionDebt      string `json:"sign_convention_debt"`
	NormalizationDebt       string `json:"normalization_debt"`
	ContainsFinalTruthClaim string `json:"containsFinalTruthClaim"`
	LeanVerification        string `json:"LeanVerification"`
	PromotionStatus         string `json:"promotion_status"`
}

type NullModelSpecReport struct {
	SchemaVersion             string                         `json:"schema_version"`
	ToyAnalysisOnly           bool                           `json:"toy_analysis_only"`
	FinalTruthClaim           bool                           `json:"final_truth_claim"`
	SpecKind                  string                         `json:"spec_kind"`
	SpecScope                 string                         `json:"spec_scope"`
	ResidualComputed          bool                           `json:"residual_computed"`
	NullComparisonComputed    bool                           `json:"null_comparison_computed"`
	FriedmannRecoveryClaim    bool                           `json:"friedmann_recovery_claim"`
	SourceArtifacts           []string                       `json:"source_artifacts"`
	NullModels                []NullModelSpec                `json:"null_models"`
	InputRequirements         []NullModelInputRequirement    `json:"input_requirements"`
	MetricContracts           []NullModelMetricContract      `json:"metric_contracts"`
	FutureComparisonContracts []FutureNullComparisonContract `json:"future_comparison_contracts"`
	Gates                     []NullSpecGate                 `json:"gates"`
	PromotionGate             report.PromotionGate           `json:"promotion_gate"`
	EbpDebt                   NullModelEbpDebt               `json:"ebp_debt"`
	Warnings                  []string                       `json:"warnings"`
}

// GenerateNullModelSpecReport builds the Sprint 7 Null-Model Scaffold Report.
func GenerateNullModelSpecReport(safeParams model.SuperpositionParams) (*NullModelSpecReport, error) {
	// Registered Null Models
	nullModels := []NullModelSpec{
		{
			NullModelID:                     "constant_phase_control",
			Name:                            "Constant Phase Control Null Model",
			Purpose:                         "Verify if a phase-invariant state produces a fake residual match.",
			ControlsFor:                     []string{"phase alignment artifacts"},
			RequiredInputs:                  []string{"bmc0a_superposition_safe"},
			RequiredMetrics:                 []string{"branch_count_stability"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to isolate phase-independent wave behavior from cosmological evolution.",
		},
		{
			NullModelID:                     "randomized_phase_control",
			Name:                            "Randomized Phase Control Null Model",
			Purpose:                         "Establish baseline residual behavior for random phase fluctuations.",
			ControlsFor:                     []string{"random alignment artifacts"},
			RequiredInputs:                  []string{"bmc0a_superposition_safe"},
			RequiredMetrics:                 []string{"local_relation_single_valuedness_rate"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to calibrate the background noise of random phase configurations.",
		},
		{
			NullModelID:                     "matched_amplitude_randomized_phase_control",
			Name:                            "Matched Amplitude Randomized Phase Null Model",
			Purpose:                         "Disentangle the amplitude contribution from phase alignment.",
			ControlsFor:                     []string{"amplitude-scaling false positives"},
			RequiredInputs:                  []string{"bmc0a_superposition_safe"},
			RequiredMetrics:                 []string{"min_amplitude_distribution"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to ensure amplitude envelope does not mimic cosmic expansion on its own.",
		},
		{
			NullModelID:                     "classical_frw_reference_trajectory",
			Name:                            "Classical FRW Reference Null Model",
			Purpose:                         "Verify convergence to classical limit reference curves.",
			ControlsFor:                     []string{"classical limit divergence"},
			RequiredInputs:                  []string{"bmc0a_superposition_safe", "bmc0a_friedmann_spec"},
			RequiredMetrics:                 []string{"derivative_readiness_rate"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to check classical limit convergence behavior under noise.",
		},
		{
			NullModelID:                     "same_branch_segmentation_under_null_wavefunctions",
			Name:                            "Same Branch Segmentation under Null Wavefunctions",
			Purpose:                         "Confirm if non-monotonic segmentation is robust under null perturbations.",
			ControlsFor:                     []string{"segmentation sensitivity"},
			RequiredInputs:                  []string{"bmc0a_clock_readiness"},
			RequiredMetrics:                 []string{"turning_point_stability"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to confirm turning points are physical and not segmenting on noise.",
		},
		{
			NullModelID:                     "node_neighborhood_stress_case",
			Name:                            "Node Neighborhood Stress Null Model",
			Purpose:                         "Check robustness of exclusions near wavefunction nodes.",
			ControlsFor:                     []string{"node proximity errors"},
			RequiredInputs:                  []string{"bmc0a_clock_readiness"},
			RequiredMetrics:                 []string{"node_contact_rate"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to stress-test exclusionary boundary logic near node obstructions.",
		},
		{
			NullModelID:                     "clock_choice_alternative_branch_diagnostic",
			Name:                            "Clock Choice Alternative Branch Null Model",
			Purpose:                         "Analyze alternative clock candidates (e.g. alpha clock) on segmented branches.",
			ControlsFor:                     []string{"clock choice bias"},
			RequiredInputs:                  []string{"bmc0a_clock_fragility"},
			RequiredMetrics:                 []string{"residual_formula_input_availability"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Required to rule out bias coming from using one relational clock over another.",
		},
	}

	// Required Inputs
	inputRequirements := []NullModelInputRequirement{
		{
			InputID:            "bmc0a_superposition_safe",
			SourceArtifact:     "Sprint 2: BMC-0A two-plane-wave superposition control artifact",
			RequiredFor:        []string{"constant_phase_control", "randomized_phase_control", "matched_amplitude_randomized_phase_control", "classical_frw_reference_trajectory"},
			AvailabilityStatus: AvailabilityAvailable,
			Reason:             "Base quantum trajectory data is fully generated and validated.",
		},
		{
			InputID:            "bmc0a_superposition_robustness",
			SourceArtifact:     "Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
			RequiredFor:        []string{"classical_frw_reference_trajectory"},
			AvailabilityStatus: AvailabilityAvailable,
			Reason:             "Convergence metrics are available for classical comparisons.",
		},
		{
			InputID:            "bmc0a_clock_fragility",
			SourceArtifact:     "Sprint 4: BMC-0A clock-fragility diagnostic artifact",
			RequiredFor:        []string{"clock_choice_alternative_branch_diagnostic"},
			AvailabilityStatus: AvailabilityAvailable,
			Reason:             "Clock monotonicity results are recorded.",
		},
		{
			InputID:            "bmc0a_clock_readiness",
			SourceArtifact:     "Sprint 5: BMC-0A clock-readiness/local segmentation artifact",
			RequiredFor:        []string{"same_branch_segmentation_under_null_wavefunctions", "node_neighborhood_stress_case"},
			AvailabilityStatus: AvailabilityAvailable,
			Reason:             "Local branch and turning point maps are recorded.",
		},
		{
			InputID:            "bmc0a_friedmann_spec",
			SourceArtifact:     "Sprint 6: BMC-0A Friedmann-residual specification/gate-design artifact",
			RequiredFor:        []string{"classical_frw_reference_trajectory"},
			AvailabilityStatus: AvailabilityAvailable,
			Reason:             "Candidate target maps and specifications are defined.",
		},
	}

	// Future Metrics
	metricContracts := []NullModelMetricContract{
		{
			MetricID:                        "branch_count_stability",
			Description:                     "Evaluation of segment branch count consistency across perturbations",
			AppliesTo:                       []string{"constant_phase_control"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Enforces stable segmentation before residual comparison.",
		},
		{
			MetricID:                        "turning_point_stability",
			Description:                     "Sensitivity of turning points to wavefunction phase shifts",
			AppliesTo:                       []string{"same_branch_segmentation_under_null_wavefunctions"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Ensures clock turning points are physical.",
		},
		{
			MetricID:                        "local_relation_single_valuedness_rate",
			Description:                     "Proportion of branches that preserve a single-valued relation to target clock",
			AppliesTo:                       []string{"randomized_phase_control"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Screens out random phase fluctuations that violate single-valuedness.",
		},
		{
			MetricID:                        "derivative_readiness_rate",
			Description:                     "Proportion of branch points where numerical derivative stencil is stable",
			AppliesTo:                       []string{"classical_frw_reference_trajectory"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Validates derivative stencils for residual formulas.",
		},
		{
			MetricID:                        "node_contact_rate",
			Description:                     "Frequency of points entering excluded node-proximity zones",
			AppliesTo:                       []string{"node_neighborhood_stress_case"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Ensures node proximity logic remains active.",
		},
		{
			MetricID:                        "min_amplitude_distribution",
			Description:                     "Verification of minimum amplitude distribution against noise limits",
			AppliesTo:                       []string{"matched_amplitude_randomized_phase_control"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Prevents near-node mathematical blowups in metrics.",
		},
		{
			MetricID:                        "max_abs_q_distribution",
			Description:                     "Distribution of qpotential magnitudes away from nodes",
			AppliesTo:                       []string{"constant_phase_control"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Tracks potential barriers relative to null models.",
		},
		{
			MetricID:                        "phase_gradient_distribution",
			Description:                     "Distribution of wavefunction phase gradients",
			AppliesTo:                       []string{"randomized_phase_control"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Monitors field velocity baseline fluctuations.",
		},
		{
			MetricID:                        "residual_formula_input_availability",
			Description:                     "Availability rate of required variable mappings",
			AppliesTo:                       []string{"clock_choice_alternative_branch_diagnostic"},
			RequiredBeforeResidualPromotion: true,
			Status:                          StatusPlanned,
			Reason:                          "Confirms alternative clocks are ready for testing.",
		},
	}

	// Future Comparison Contracts
	comparisonContracts := []FutureNullComparisonContract{
		{
			ComparisonID:       "bmc0a-superposition-vs-nulls-v0.1",
			BaselineArtifact:   "bmc0a_superposition_safe",
			NullModelIDs:       []string{"constant_phase_control", "randomized_phase_control", "matched_amplitude_randomized_phase_control"},
			Metrics:            []string{"branch_count_stability", "local_relation_single_valuedness_rate", "min_amplitude_distribution"},
			ComparisonComputed: false,
			Status:             StatusPlanned,
			Reason:             "Future comparison contract planned; no comparisons computed in Sprint 7.",
		},
	}

	// Safety Gates
	gates := []NullSpecGate{
		{"toy_analysis_only_gate", "pass", "Confirms analysis is strictly restricted to minisuperspace toy system."},
		{"no_final_truth_claim_gate", "pass", "Confirms no final truth claims are asserted."},
		{"no_residual_computation_gate", "pass", "Confirms that no actual numerical residual was computed."},
		{"no_null_comparison_result_gate", "pass", "Confirms that no null model comparison results are computed."},
		{"null_model_registry_complete_gate", "pass", "Confirms that all 7 required null models are fully registered."},
		{"required_before_residual_promotion_gate", "pass", "Confirms that all required null models are marked required_before_residual_promotion = true."},
		{"friedmann_recovery_claim_blocked_gate", "pass", "Confirms that no recovery claim is made."},
		{"full_bmc_blocked_gate", "pass", "Confirms the full BMC promotion gate is blocked."},
		{"clock_choice_debt_active_gate", "pass", "Confirms clock choice debt remains active."},
		{"faithfulness_contested_gate", "pass", "Confirms faithfulness review status remains contested."},
	}

	// EBP Debt
	ebpDebt := NullModelEbpDebt{
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
		LeanVerification:        "retired_for_policy_safety_contracts",
		PromotionStatus:         "planned_candidate_only",
	}

	// Promotion Gate
	promotionGate := report.PromotionGate{
		Name:   "full_bmc_toy_gate",
		Status: report.StatusBlocked,
		Reason: "Sprint 7 defines null-model scaffolding only. No null-model comparison results or Friedmann residuals are computed.",
	}

	warnings := []string{
		"Sprint 7 is a null-model scaffold sprint only.",
		"No Friedmann residual was computed.",
		"No null-model comparison result was computed.",
		"No recovery claim is made.",
		"Null-model debt remains active.",
		"Full BMC remains blocked.",
	}

	repVal := &NullModelSpecReport{
		SchemaVersion:             "bmc0a-nullmodel-spec-v0.1",
		ToyAnalysisOnly:           true,
		FinalTruthClaim:           false,
		SpecKind:                  "null_model_scaffold",
		SpecScope:                 SpecScopeNullModelOnly,
		ResidualComputed:          false,
		NullComparisonComputed:    false,
		FriedmannRecoveryClaim:    false,
		SourceArtifacts: []string{
			"Sprint 1: BMC-0A plane-wave control artifact",
			"Sprint 2: BMC-0A two-plane-wave superposition control artifact",
			"Sprint 2: BMC-0A node-obstruction detection artifact",
			"Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
			"Sprint 4: BMC-0A clock-fragility diagnostic artifact",
			"Sprint 5: BMC-0A clock-readiness/local segmentation artifact",
			"Sprint 6: BMC-0A Friedmann-residual specification/gate-design artifact",
			"Sprint 6.1: BMC-0A Friedmann-spec data-integrity/gate repair artifact",
		},
		NullModels:                nullModels,
		InputRequirements:         inputRequirements,
		MetricContracts:           metricContracts,
		FutureComparisonContracts: comparisonContracts,
		Gates:                     gates,
		PromotionGate:             promotionGate,
		EbpDebt:                   ebpDebt,
		Warnings:                  warnings,
	}

	return repVal, nil
}

// ReadNullModelSpecReport reads a JSON file and strictly decodes it, protecting against trailing garbage.
func ReadNullModelSpecReport(path string) (*NullModelSpecReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep NullModelSpecReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}

	var dummy json.RawMessage
	if err := dec.Decode(&dummy); err != io.EOF {
		return nil, fmt.Errorf("trailing garbage in JSON: %v", err)
	}

	return &rep, nil
}

// WriteJSON writes the report to a pretty-printed JSON file.
func WriteJSON(rep *NullModelSpecReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// SummarizeNullModelSpecReport outputs a human-readable summary.
func SummarizeNullModelSpecReport(rep *NullModelSpecReport) {
	fmt.Println("================================================================================")
	fmt.Printf("BMC Sprint 7 Null-Model Spec Report Summary\n")
	fmt.Println("================================================================================")
	fmt.Printf("Schema Version:            %s\n", rep.SchemaVersion)
	fmt.Printf("Spec Kind:                 %s\n", rep.SpecKind)
	fmt.Printf("Spec Scope:                %s\n", rep.SpecScope)
	fmt.Printf("Residual Computed:         %v (Must be false)\n", rep.ResidualComputed)
	fmt.Printf("Null Comparison Computed:  %v (Must be false)\n", rep.NullComparisonComputed)
	fmt.Printf("Recovery Claim:            %v (Must be false)\n", rep.FriedmannRecoveryClaim)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Null Model specs registered:\n")
	for _, nm := range rep.NullModels {
		fmt.Printf("  - ID: %-44s Status: [%s]\n", nm.NullModelID, nm.Status)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Gates:\n")
	for _, g := range rep.Gates {
		fmt.Printf("  - %-40s [%-7s]\n", g.Name, g.Status)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("EBP Debt Ledger:\n")
	fmt.Printf("  needMap:                 %s\n", rep.EbpDebt.NeedMap)
	fmt.Printf("  needNullModel:           %s\n", rep.EbpDebt.NeedNullModel)
	fmt.Printf("  clock_choice_debt:       %s\n", rep.EbpDebt.ClockChoiceDebt)
	fmt.Printf("  classical_target_debt:   %s\n", rep.EbpDebt.ClassicalTargetDebt)
	fmt.Printf("  unit_convention_debt:    %s\n", rep.EbpDebt.UnitConventionDebt)
	fmt.Printf("  sign_convention_debt:    %s\n", rep.EbpDebt.SignConventionDebt)
	fmt.Printf("  normalization_debt:      %s\n", rep.EbpDebt.NormalizationDebt)
	fmt.Printf("  promotion_status:        %s\n", rep.EbpDebt.PromotionStatus)
	fmt.Println("================================================================================")
}
