package residualaudit

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/priorart"
	"github.com/PithomLabs/bmc/internal/bmc/residualrun"
)

type ResidualAuditEbpDebt struct {
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

type ResidualAuditReport struct {
	SchemaVersion              string                    `json:"schema_version"`
	ToyAnalysisOnly            bool                      `json:"toy_analysis_only"`
	FinalTruthClaim            bool                      `json:"final_truth_claim"`
	ArtifactKind               string                    `json:"artifact_kind"`
	Scope                      string                    `json:"scope"`
	ResidualAuditComputed      bool                      `json:"residual_audit_computed"`
	RecoveryClaim              bool                      `json:"recovery_claim"`
	ScientificNoveltyClaimMade bool                      `json:"scientific_novelty_claim_made"`
	BmcBeatsNullModelsClaim    bool                      `json:"bmc_beats_null_models_claim"`
	FullBmcToyGate             string                    `json:"full_bmc_toy_gate"`
	LocalBranchOnly            bool                      `json:"local_branch_only"`
	GlobalCosmologyClaim       bool                      `json:"global_cosmology_claim"`
	SourceArtifacts            []SourceArtifactRef       `json:"source_artifacts"`
	ComparisonAudits           []ResidualComparisonAudit `json:"comparison_audits"`
	StabilityAudits            []ResidualStabilityAudit  `json:"stability_audits"`
	Gates                      []ResidualAuditGate       `json:"gates"`
	EbpDebtVocabulary          string                    `json:"ebp_debt_vocabulary"`
	EbpDebt                    ResidualAuditEbpDebt      `json:"ebp_debt"`
	InterpretationStatus       string                    `json:"interpretation_status"`
	Warnings                   []string                  `json:"warnings"`
}

func GenerateBlockedDefaultReport() *ResidualAuditReport {
	return &ResidualAuditReport{
		SchemaVersion:              SchemaVersion,
		ToyAnalysisOnly:            true,
		FinalTruthClaim:            false,
		ArtifactKind:               "residual_null_comparison_audit",
		Scope:                      "bmc0a_only",
		ResidualAuditComputed:      false,
		RecoveryClaim:              false,
		ScientificNoveltyClaimMade: false,
		BmcBeatsNullModelsClaim:    false,
		FullBmcToyGate:             "blocked",
		LocalBranchOnly:            true,
		GlobalCosmologyClaim:       false,
		SourceArtifacts:            defaultSourceArtifacts(),
		ComparisonAudits:           []ResidualComparisonAudit{},
		StabilityAudits:            []ResidualStabilityAudit{},
		Gates:                      defaultGates(),
		EbpDebtVocabulary:          "ptw_adversarial_review_debt_status_v0.1",
		EbpDebt: ResidualAuditEbpDebt{
			NeedLiteratureAudit:     "partial",
			NeedMap:                 "partial",
			NeedInvariant:           "partial",
			NeedToyCheck:            "partial",
			NeedNullModel:           "partial",
			NeedObstruction:         "partial",
			NeedFaithfulnessReview:  "contested",
			ClockChoiceDebt:         "unpaid",
			ClassicalTargetDebt:     "unpaid",
			UnitConventionDebt:      "unpaid",
			SignConventionDebt:      "unpaid",
			NormalizationDebt:       "unpaid",
			ContainsFinalTruthClaim: "absent",
			PromotionStatus:         "residual_audit_candidate_only",
		},
		InterpretationStatus: InterpretBlockedBySourceUnavailable,
		Warnings: []string{
			"Sprint 11 audits residual/null comparison integrity only.",
			"No recovery claim is made.",
			"No BMC-beats-null-models claim is made.",
			"Full BMC remains blocked.",
			"Convention debts remain unpaid or contested.",
		},
	}
}

func GenerateDefaultReport() *ResidualAuditReport {
	return GenerateBlockedDefaultReport()
}

func RunAuditFromFiles(clockPath, nullrunPath, residualPath, friedmannPath, priorartPath string) (*ResidualAuditReport, error) {
	rep := GenerateBlockedDefaultReport()

	if _, err := os.ReadFile(clockPath); err != nil {
		markSourceUnavailable(rep.SourceArtifacts, "bmc0a_clock_readiness", clockPath, err)
		return rep, nil
	}
	markSourceRead(rep.SourceArtifacts, "bmc0a_clock_readiness", clockPath)

	nullRep, err := nullrun.ReadReport(nullrunPath)
	if err != nil {
		markSourceUnavailable(rep.SourceArtifacts, "bmc0a_nullrun", nullrunPath, err)
		return rep, nil
	}
	markSourceRead(rep.SourceArtifacts, "bmc0a_nullrun", nullrunPath)

	residualRep, err := residualrun.ReadReport(residualPath)
	if err != nil {
		markSourceUnavailable(rep.SourceArtifacts, "bmc0a_local_residual", residualPath, err)
		return rep, nil
	}
	markSourceRead(rep.SourceArtifacts, "bmc0a_local_residual", residualPath)

	if _, err := os.ReadFile(friedmannPath); err != nil {
		markSourceUnavailable(rep.SourceArtifacts, "bmc0a_friedmann_spec", friedmannPath, err)
		return rep, nil
	}
	markSourceRead(rep.SourceArtifacts, "bmc0a_friedmann_spec", friedmannPath)

	if _, err := priorart.ReadReport(priorartPath); err != nil {
		markSourceUnavailable(rep.SourceArtifacts, "bmc0a_prior_art_boundary", priorartPath, err)
		return rep, nil
	}
	markSourceRead(rep.SourceArtifacts, "bmc0a_prior_art_boundary", priorartPath)

	return RunAuditFromInputs(residualRep, nullRep, rep.SourceArtifacts), nil
}

func RunAuditFromInputs(residualRep *residualrun.ResidualRunReport, nullRep *nullrun.NullRunReport, sources []SourceArtifactRef) *ResidualAuditReport {
	rep := GenerateBlockedDefaultReport()
	if len(sources) > 0 {
		rep.SourceArtifacts = sources
	}
	if residualRep == nil || nullRep == nil || !residualRep.CandidateResidualComputed {
		rep.InterpretationStatus = InterpretBlockedByMissingResidualInputs
		return rep
	}
	if !nullRep.NullDiagnosticsComputed {
		rep.InterpretationStatus = InterpretBlockedByMissingNullInputs
		return rep
	}

	rep.ComparisonAudits = buildComparisonAudits(residualRep, nullRep)
	rep.StabilityAudits = buildStabilityAudits(residualRep.CandidateResidualDiagnostics)
	rep.ResidualAuditComputed = len(rep.ComparisonAudits) > 0 && len(rep.StabilityAudits) > 0
	rep.InterpretationStatus = summarizeInterpretation(rep.ComparisonAudits, rep.StabilityAudits)
	if !rep.ResidualAuditComputed {
		rep.InterpretationStatus = InterpretBlockedBySourceUnavailable
	}
	return rep
}

func ReadReport(path string) (*ResidualAuditReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()
	var rep ResidualAuditReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}
	var dummy json.RawMessage
	if err := dec.Decode(&dummy); err != io.EOF {
		return nil, fmt.Errorf("trailing garbage in JSON: %v", err)
	}
	return &rep, nil
}

func WriteReport(rep *ResidualAuditReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}
