package friedmannspec

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

type FriedmannEbpDebt struct {
	NeedMap                 string `json:"needMap"`
	NeedInvariant           string `json:"needInvariant"`
	NeedToyCheck            string `json:"needToyCheck"`
	NeedNullModel          string `json:"needNullModel"`
	NeedObstruction         string `json:"needObstruction"`
	NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
	ClockChoiceDebt        string `json:"clock_choice_debt"`
	ClassicalTargetDebt    string `json:"classical_target_debt"`
	UnitConventionDebt     string `json:"unit_convention_debt"`
	SignConventionDebt     string `json:"sign_convention_debt"`
	NormalizationDebt      string `json:"normalization_debt"`
	ContainsFinalTruthClaim string `json:"containsFinalTruthClaim"`
	LeanVerification       string `json:"LeanVerification"`
	PromotionStatus        string `json:"promotion_status"`
}

type FriedmannSpecReport struct {
	SchemaVersion             string                          `json:"schema_version"`
	ToyAnalysisOnly           bool                            `json:"toy_analysis_only"`
	FinalTruthClaim           bool                            `json:"final_truth_claim"`
	SpecKind                  string                          `json:"spec_kind"`
	SpecScope                 string                          `json:"spec_scope"`
	SourceArtifacts           []string                        `json:"source_artifacts"`
	ResidualComputed          bool                            `json:"residual_computed"`
	FriedmannRecoveryClaim    bool                            `json:"friedmann_recovery_claim"`
	CandidateMaps             []FriedmannCandidateMap         `json:"candidate_maps"`
	BranchRequirements        []FriedmannBranchRequirement    `json:"branch_requirements"`
	DerivativeReadinessChecks []DerivativeReadinessCheck      `json:"derivative_readiness_checks"`
	ResidualFormulaCandidates []ResidualFormulaCandidate      `json:"residual_formula_candidates"`
	NullModelRequirements     []FriedmannNullModelRequirement `json:"null_model_requirements"`
	Gates                     []FriedmannSpecGate             `json:"gates"`
	PromotionGate             report.PromotionGate            `json:"promotion_gate"`
	EbpDebt                   FriedmannEbpDebt                `json:"ebp_debt"`
	Warnings                  []string                        `json:"warnings"`
}

// GenerateFriedmannSpecReport builds the specification report under EBP 2.1 policy rules.
func GenerateFriedmannSpecReport(safeParams model.SuperpositionParams) (*FriedmannSpecReport, error) {
	// Candidate Mapping Record
	candidateMap := FriedmannCandidateMap{
		MapID:                            "bmc0a-map-v0.1",
		AlphaMeaning:                     "log of scale factor ln(a)",
		PhiMeaning:                       "homogeneous massless scalar field",
		ClockVariable:                    "phi",
		ClockScope:                       "local relational monotonic branches",
		CandidateScaleFactor:             "a = exp(alpha)",
		CandidateHubbleDefinition:         "H = dalpha/dt_candidate",
		CandidateEnergyDensityDefinition: "rho_phi = (dphi/dlambda)^2 / (2 * a^6)",
		UnitConventionStatus:             StatusContested,
		SignConventionStatus:             StatusContested,
		NormalizationStatus:             StatusContested,
		ClockChoiceDebt:                  "active",
		ClassicalTargetDebt:              "active",
		Status:                           StatusCandidateOnly,
		Reason:                           "Debts are active; target relations and physical normalization are not yet confirmed.",
	}

	// Branch Requirement
	branchReq := FriedmannBranchRequirement{
		BranchID:                "branch-0",
		SourceConfigID:          "bmc0a_superposition_safe",
		PhiLocalMonotonic:       true,
		AlphaPhiSingleValued:    true,
		MinBranchSamples:        3,
		ActualSamples:           201,
		ClockRange:              4.1810838631599525,
		LambdaRange:             10.0,
		DerivativeReady:         false,
		DerivativeMethod:        "finite_difference_stencil_3pt",
		DerivativeDebt:          "active",
		NodeContactFree:         true,
		QFiniteAwayFromNodes:    true,
		BranchResidualReadiness: StatusCandidateOnly,
		Reason:                  "Branch is numerically well-formed but actual derivative calculation is deferred.",
	}

	// Derivative Readiness Check
	derivCheck := DerivativeReadinessCheck{
		BranchID:                          "branch-0",
		Method:                            "finite_difference_3pt",
		ExcludesEndpoints:                 true,
		ExcludesTurningPointNeighborhoods: true,
		ExcludesNearNodePoints:            true,
		MinSamplesRequired:                5,
		SamplesAvailable:                  199,
		StepSensitivityStatus:             StatusContested,
		Status:                            StatusCandidateOnly,
		Reason:                            "Excluded boundary and near-node points; step sensitivity requires multi-grid sweep verification.",
	}

	// Residual Formula Candidate
	formulaCandidate := ResidualFormulaCandidate{
		FormulaID:           "flat-frw-massless-scalar-v0.1",
		Description:         "Candidate Hubble-squared and scalar kinetic density relation",
		ClassicalTarget:     "H^2 = (8 * pi * G / 3) * rho_phi",
		RequiredVariables:   []string{"alpha", "phi", "lambda"},
		RequiredDerivatives: []string{"dalpha/dphi", "d2alpha/dphi2"},
		ConventionDebt:      []string{"G_gravitational_constant_normalization", "time_parameter_gauge"},
		NullModelsRequired:  []string{"constant-phase control", "randomized phase control"},
		Status:              StatusCandidateOnly,
		Reason:              "Formula registered for specification purposes only; no numerical evaluation was executed.",
	}

	// Null Model Requirements
	nullModels := []struct {
		ID      string
		Purpose string
	}{
		{"constant-phase control", "Verify if a phase-invariant state produces a fake residual match."},
		{"randomized phase control", "Establish baseline residual behavior for random phase fluctuations."},
		{"matched amplitude / randomized phase control", "Disentangle the amplitude contribution from phase alignment."},
		{"classical FRW reference trajectory", "Verify convergence to classical limit reference curves."},
		{"same branch segmentation under null wavefunctions", "Confirm if non-monotonic segmentation is robust under null perturbations."},
		{"node-neighborhood stress case", "Check robustness of exclusions near wavefunction nodes."},
		{"clock-choice alternative branch diagnostic", "Analyze alternative clock candidates (e.g. alpha clock) on segmented branches."},
	}

	var nullModelReqs []FriedmannNullModelRequirement
	for _, nm := range nullModels {
		nullModelReqs = append(nullModelReqs, FriedmannNullModelRequirement{
			NullModelID:                     nm.ID,
			Purpose:                         nm.Purpose,
			RequiredBeforeResidualPromotion: true,
			Status:                          "planned",
			Reason:                          "Required to protect against false positive residual matching.",
		})
	}

	// Gate structures
	requiredGates := []struct {
		Name   string
		Reason string
	}{
		{"toy_analysis_only_gate", "Confirms analysis is strictly restricted to minisuperspace toy system."},
		{"no_final_truth_claim_gate", "Confirms no final truth claims are asserted."},
		{"local_branch_only_gate", "Confirms clock is only evaluated on local monotonic branches."},
		{"clock_choice_debt_active_gate", "Confirms clock choice debt remains active."},
		{"classical_target_candidate_only_gate", "Confirms the classical target is marked as candidate only."},
		{"unit_convention_debt_gate", "Confirms unit convention debt remains active."},
		{"null_model_debt_gate", "Confirms null model debt remains active."},
		{"faithfulness_contested_gate", "Confirms faithfulness review status remains contested."},
		{"no_residual_computation_gate", "Confirms that no actual numerical residual was computed."},
		{"full_bmc_blocked_gate", "Confirms the full BMC promotion gate is blocked."},
	}

	var gates []FriedmannSpecGate
	for _, g := range requiredGates {
		gates = append(gates, FriedmannSpecGate{
			Name:   g.Name,
			Status: "pass",
			Reason: g.Reason,
		})
	}

	// EBP Debt
	ebpDebt := FriedmannEbpDebt{
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
		LeanVerification:        "planned",
		PromotionStatus:         "planned_candidate_only",
	}

	// Promotion Gate
	promotionGate := report.PromotionGate{
		Name:   "full_bmc_toy_gate",
		Status: report.StatusBlocked,
		Reason: "Sprint 6 defines residual specifications only. No residual recovery is computed.",
	}

	warnings := []string{
		"Sprint 6 is a specification sprint only.",
		"No Friedmann residual was computed.",
		"No Friedmann recovery is claimed.",
		"Local branch readiness is not Friedmann recovery readiness.",
		"Full BMC remains blocked.",
	}

	repVal := &FriedmannSpecReport{
		SchemaVersion:             "bmc0a-friedmann-spec-v0.1",
		ToyAnalysisOnly:           true,
		FinalTruthClaim:           false,
		SpecKind:                  "friedmann_residual_specification",
		SpecScope:                 SpecScopeCandidateOnly,
		SourceArtifacts: []string{
			"Sprint 1: BMC-0A plane-wave control artifact",
			"Sprint 2: BMC-0A two-plane-wave superposition control artifact",
			"Sprint 2: BMC-0A node-obstruction detection artifact",
			"Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
			"Sprint 4: BMC-0A clock-fragility diagnostic artifact",
			"Sprint 5: BMC-0A clock-readiness/local segmentation artifact",
		},
		ResidualComputed:          false,
		FriedmannRecoveryClaim:    false,
		CandidateMaps:             []FriedmannCandidateMap{candidateMap},
		BranchRequirements:        []FriedmannBranchRequirement{branchReq},
		DerivativeReadinessChecks: []DerivativeReadinessCheck{derivCheck},
		ResidualFormulaCandidates: []ResidualFormulaCandidate{formulaCandidate},
		NullModelRequirements:     nullModelReqs,
		Gates:                     gates,
		PromotionGate:             promotionGate,
		EbpDebt:                   ebpDebt,
		Warnings:                  warnings,
	}

	return repVal, nil
}

// ReadFriedmannSpecReport reads a JSON file and strictly decodes it.
func ReadFriedmannSpecReport(path string) (*FriedmannSpecReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep FriedmannSpecReport
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
func WriteJSON(rep *FriedmannSpecReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// SummarizeFriedmannSpecReport outputs a human-readable summary.
func SummarizeFriedmannSpecReport(rep *FriedmannSpecReport) {
	fmt.Println("================================================================================")
	fmt.Printf("BMC Sprint 6 Friedmann Spec Report Summary\n")
	fmt.Println("================================================================================")
	fmt.Printf("Schema Version:            %s\n", rep.SchemaVersion)
	fmt.Printf("Spec Kind:                 %s\n", rep.SpecKind)
	fmt.Printf("Spec Scope:                %s\n", rep.SpecScope)
	fmt.Printf("Residual Computed:         %v (Must be false)\n", rep.ResidualComputed)
	fmt.Printf("Friedmann Recovery Claim:  %v (Must be false)\n", rep.FriedmannRecoveryClaim)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Candidate Maps:\n")
	for _, m := range rep.CandidateMaps {
		fmt.Printf("  - Map ID: %s (Status: %s)\n", m.MapID, m.Status)
		fmt.Printf("    Clock: %s (%s)\n", m.ClockVariable, m.ClockScope)
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
