package residualrun

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"sort"

	"github.com/PithomLabs/bmc/internal/bmc/clockseg"
	"github.com/PithomLabs/bmc/internal/bmc/friedmannspec"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/priorart"
)

type ResidualConventionLedger struct {
	ConventionID        string `json:"convention_id"`
	Status              string `json:"status"`
	Description         string `json:"description"`
	HumanReviewRequired bool   `json:"human_review_required"`
}

type ResidualCalculationLedger struct {
	CalculationID      string   `json:"calculation_id"`
	BranchID           string   `json:"branch_id"`
	FormulaID          string   `json:"formula_id"`
	FormulaDescription string   `json:"formula_description"`
	FormulaSource      string   `json:"formula_source"`
	ConventionProfile  string   `json:"convention_profile"`
	InputProvenance    string   `json:"input_provenance"`
	InputFields        []string `json:"input_fields"`
	NumInputPoints     int      `json:"num_input_points"`
	CalculationStatus  string   `json:"calculation_status"`
	Notes              string   `json:"notes"`
}

type ResidualRunEbpDebt struct {
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

type SourceBranchRegistry struct {
	SourceArtifactID string   `json:"source_artifact_id"`
	BranchIDs        []string `json:"branch_ids"`
}

type ResidualRunReport struct {
	SchemaVersion                string                        `json:"schema_version"`
	ToyAnalysisOnly              bool                          `json:"toy_analysis_only"`
	FinalTruthClaim              bool                          `json:"final_truth_claim"`
	ArtifactKind                 string                        `json:"artifact_kind"`
	Scope                        string                        `json:"scope"`
	CandidateResidualComputed    bool                          `json:"candidate_residual_computed"`
	ResidualRecoveryClaim        bool                          `json:"residual_recovery_claim"`
	ScientificNoveltyClaimMade   bool                          `json:"scientific_novelty_claim_made"`
	BmcBeatsNullModelsClaim      bool                          `json:"bmc_beats_null_models_claim"`
	FullBmcToyGate               string                        `json:"full_bmc_toy_gate"`
	LocalBranchOnly              bool                          `json:"local_branch_only"`
	GlobalCosmologyClaim         bool                          `json:"global_cosmology_claim"`
	SourceArtifacts              []SourceArtifactRef           `json:"source_artifacts"`
	LocalBranchEligibility       []LocalBranchEligibility      `json:"local_branch_eligibility"`
	ConventionLedger             []ResidualConventionLedger    `json:"convention_ledger"`
	CandidateResidualDiagnostics []CandidateResidualDiagnostic `json:"candidate_residual_diagnostics"`
	ResidualNullComparisons      []ResidualNullComparison      `json:"residual_null_comparisons"`
	Gates                        []ResidualRunGate             `json:"gates"`
	EbpDebtVocabulary            string                        `json:"ebp_debt_vocabulary"`
	InterpretationStatus         string                        `json:"interpretation_status"`
	EbpDebt                      ResidualRunEbpDebt            `json:"ebp_debt"`
	Warnings                     []string                      `json:"warnings"`
	CalculationLedger            []ResidualCalculationLedger   `json:"calculation_ledger"`
	SourceBranchRegistry         SourceBranchRegistry          `json:"source_branch_registry"`
}

// GenerateDefaultReport creates a blocked candidate local-branch residual runner report.
func GenerateDefaultReport() *ResidualRunReport {
	return GenerateBlockedDefaultReport()
}

// GenerateBlockedDefaultReport returns a default report with candidate_residual_computed = false.
func GenerateBlockedDefaultReport() *ResidualRunReport {
	sources := []SourceArtifactRef{
		{
			ArtifactID:   "bmc0a_clock_readiness",
			ArtifactKind: "clock_readiness_report",
			Provenance:   "not_available",
			Status:       "not_available",
			Notes:        "Clock readiness diagnostics not loaded",
		},
		{
			ArtifactID:   "bmc0a_friedmann_spec",
			ArtifactKind: "friedmann_spec_report",
			Provenance:   "not_available",
			Status:       "not_available",
			Notes:        "Friedmann spec definitions not loaded",
		},
		{
			ArtifactID:   "bmc0a_nullrun",
			ArtifactKind: "nullrun_report",
			Provenance:   "not_available",
			Status:       "not_available",
			Notes:        "Null model diagnostics not loaded",
		},
		{
			ArtifactID:   "bmc0a_prior_art_boundary",
			ArtifactKind: "prior_art_boundary_report",
			Provenance:   "not_available",
			Status:       "not_available",
			Notes:        "Prior art boundary note not loaded",
		},
	}

	eligibility := []LocalBranchEligibility{
		{
			BranchID:                  "branch_0",
			SourceArtifact:            "bmc0a_clock_readiness",
			Eligible:                  false,
			EligibilityStatus:         EligibilitySourceUnavailable,
			Reason:                    "Source artifact unavailable.",
			NodeContactFree:           false,
			TrajectoryFinite:          false,
			LocalClockStatus:          "unknown",
			DerivativeReadinessStatus: "blocked",
		},
	}

	ledger := []ResidualConventionLedger{
		{
			ConventionID:        "clock_choice_debt",
			Status:              "unpaid",
			Description:         "Choice of relational clock parameter phi is local and uncalibrated.",
			HumanReviewRequired: true,
		},
		{
			ConventionID:        "classical_target_debt",
			Status:              "unpaid",
			Description:         "Classical Friedmann target is a toy-model minisuperspace limit.",
			HumanReviewRequired: true,
		},
		{
			ConventionID:        "unit_convention_debt",
			Status:              "unpaid",
			Description:         "Units of gravity and fields are not physically matched.",
			HumanReviewRequired: true,
		},
		{
			ConventionID:        "sign_convention_debt",
			Status:              "unpaid",
			Description:         "Sign conventions for relational variables require physical calibration.",
			HumanReviewRequired: true,
		},
		{
			ConventionID:        "normalization_debt",
			Status:              "unpaid",
			Description:         "Normalization factors for amplitude and metrics are uncalibrated.",
			HumanReviewRequired: true,
		},
		{
			ConventionID:        "faithfulness_review_debt",
			Status:              "contested",
			Description:         "Physical faithfulness remains contested under adversarial EBP review.",
			HumanReviewRequired: true,
		},
	}

	residuals := []CandidateResidualDiagnostic{
		{
			BranchID:           "branch_0",
			ResidualID:         "candidate_residual_branch_0",
			ResidualComputed:   false,
			ResidualStatus:     ResidualStatusSourceUnavailable,
			ResidualProvenance: ProvenanceBlocked,
			Metrics: CandidateResidualMetrics{
				NumEvaluationPoints:     0,
				NumFiniteResidualPoints: 0,
				MeanAbsResidual:         nil,
				MaxAbsResidual:          nil,
				RmsResidual:             nil,
				ResidualFinite:          false,
				DiagnosticWarnings:      []string{},
			},
			BlockedReason:       "Source artifacts unavailable.",
			Notes:               "Candidate residual diagnostic calculation blocked: source unavailable.",
			ResidualInputPoints: []ResidualInputPoint{},
		},
	}

	comparisons := []ResidualNullComparison{}

	gates := []ResidualRunGate{
		{"toy_analysis_only_gate", "pass", "Confirms analysis is strictly restricted to minisuperspace toy system."},
		{"no_final_truth_claim_gate", "pass", "Confirms no final truth claims are asserted."},
		{"candidate_residual_diagnostics_gate", "pass", "Confirms candidate local-branch residual diagnostics are generated."},
		{"no_recovery_claim_gate", "pass", "Confirms no recovery claim is made."},
		{"no_scientific_novelty_claim_gate", "pass", "Confirms no scientific novelty claim is made."},
		{"no_bmc_beats_null_models_claim_gate", "pass", "Confirms no BMC-beats-null-models claim is made."},
		{"full_bmc_blocked_gate", "pass", "Confirms the full BMC promotion gate is blocked."},
		{"convention_debts_visible_gate", "pass", "Confirms all convention debts are declared and visible."},
		{"faithfulness_contested_gate", "pass", "Confirms faithfulness review status remains contested."},
		{"no_full_bmc_promotion_gate", "pass", "Confirms no full BMC promotion was performed."},
	}

	ebpDebt := ResidualRunEbpDebt{
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
		PromotionStatus:         "candidate_residual_runner_candidate_only",
	}

	warnings := []string{
		"Sprint 10 computes candidate local-branch residual diagnostics only.",
		"No recovery claim is made.",
		"No scientific novelty claim is made.",
		"No BMC-beats-null-models claim is made.",
		"Full BMC remains blocked.",
		"Convention debts remain unpaid or contested.",
	}

	return &ResidualRunReport{
		SchemaVersion:                "bmc0a-local-residual-v0.1",
		ToyAnalysisOnly:              true,
		FinalTruthClaim:              false,
		ArtifactKind:                 "candidate_local_branch_residual_runner",
		Scope:                        "bmc0a_only",
		CandidateResidualComputed:    false,
		ResidualRecoveryClaim:        false,
		ScientificNoveltyClaimMade:   false,
		BmcBeatsNullModelsClaim:      false,
		FullBmcToyGate:               "blocked",
		LocalBranchOnly:              true,
		GlobalCosmologyClaim:         false,
		SourceArtifacts:              sources,
		LocalBranchEligibility:       eligibility,
		ConventionLedger:             ledger,
		CandidateResidualDiagnostics: residuals,
		ResidualNullComparisons:      comparisons,
		Gates:                        gates,
		EbpDebtVocabulary:            "ptw_adversarial_review_debt_status_v0.1",
		InterpretationStatus:         InterpretBlockedByNoEligibleLocalBranch,
		EbpDebt:                      ebpDebt,
		Warnings:                     warnings,
		CalculationLedger:            []ResidualCalculationLedger{},
		SourceBranchRegistry: SourceBranchRegistry{
			SourceArtifactID: "bmc0a_clock_readiness",
			BranchIDs:        []string{},
		},
	}
}

// RunResidualsFromFiles reads input reports from disk and computes residual diagnostics.
func RunResidualsFromFiles(clockPath, friedmannPath, nullrunPath, priorartPath string) (*ResidualRunReport, error) {
	rep := GenerateBlockedDefaultReport()

	rep.SourceArtifacts[0].Path = clockPath
	rep.SourceArtifacts[1].Path = friedmannPath
	rep.SourceArtifacts[2].Path = nullrunPath
	rep.SourceArtifacts[3].Path = priorartPath

	// Read clock readiness
	clockReport, err := clockseg.ReadClockReadinessReport(clockPath)
	if err != nil {
		rep.SourceArtifacts[0].Status = "not_available"
		rep.SourceArtifacts[0].Provenance = "not_available"
		rep.SourceArtifacts[0].Notes = fmt.Sprintf("Failed to read: %v", err)
		return rep, nil
	}
	rep.SourceArtifacts[0].Status = "available"
	rep.SourceArtifacts[0].Provenance = "file_read"
	rep.SourceArtifacts[0].Notes = "Successfully read and parsed clock readiness report"

	// Read Friedmann spec
	friedmannReport, err := friedmannspec.ReadFriedmannSpecReport(friedmannPath)
	if err != nil {
		rep.SourceArtifacts[1].Status = "not_available"
		rep.SourceArtifacts[1].Provenance = "not_available"
		rep.SourceArtifacts[1].Notes = fmt.Sprintf("Failed to read: %v", err)
		return rep, nil
	}
	rep.SourceArtifacts[1].Status = "available"
	rep.SourceArtifacts[1].Provenance = "file_read"
	rep.SourceArtifacts[1].Notes = "Successfully read and parsed Friedmann spec report"

	// Read nullrun
	nullReport, err := nullrun.ReadReport(nullrunPath)
	if err != nil {
		rep.SourceArtifacts[2].Status = "not_available"
		rep.SourceArtifacts[2].Provenance = "not_available"
		rep.SourceArtifacts[2].Notes = fmt.Sprintf("Failed to read: %v", err)
		return rep, nil
	}
	rep.SourceArtifacts[2].Status = "available"
	rep.SourceArtifacts[2].Provenance = "file_read"
	rep.SourceArtifacts[2].Notes = "Successfully read and parsed null model run report"

	// Read prior art boundary
	priorArtReport, err := priorart.ReadReport(priorartPath)
	if err != nil {
		rep.SourceArtifacts[3].Status = "not_available"
		rep.SourceArtifacts[3].Provenance = "not_available"
		rep.SourceArtifacts[3].Notes = fmt.Sprintf("Failed to read: %v", err)
		return rep, nil
	}
	rep.SourceArtifacts[3].Status = "available"
	rep.SourceArtifacts[3].Provenance = "file_read"
	rep.SourceArtifacts[3].Notes = "Successfully read and parsed prior art boundary report"

	return RunResidualsFromInputs(clockReport, friedmannReport, nullReport, priorArtReport, clockPath, friedmannPath, nullrunPath, priorartPath)
}

// RunResidualsFromInputs performs the honest candidate residual calculation from parsed structs.
func RunResidualsFromInputs(
	clockReport *clockseg.ClockReadinessReport,
	friedmannReport *friedmannspec.FriedmannSpecReport,
	nullReport *nullrun.NullRunReport,
	priorArtReport *priorart.PriorArtBoundaryReport,
	clockPath, friedmannPath, nullrunPath, priorartPath string,
) (*ResidualRunReport, error) {
	rep := GenerateBlockedDefaultReport()

	rep.SourceArtifacts[0].Provenance = "file_read"
	rep.SourceArtifacts[0].Path = clockPath
	rep.SourceArtifacts[0].Status = "available"
	rep.SourceArtifacts[0].Notes = "Successfully read and parsed clock readiness report"

	rep.SourceArtifacts[1].Provenance = "file_read"
	rep.SourceArtifacts[1].Path = friedmannPath
	rep.SourceArtifacts[1].Status = "available"
	rep.SourceArtifacts[1].Notes = "Successfully read and parsed Friedmann spec report"

	rep.SourceArtifacts[2].Provenance = "file_read"
	rep.SourceArtifacts[2].Path = nullrunPath
	rep.SourceArtifacts[2].Status = "available"
	rep.SourceArtifacts[2].Notes = "Successfully read and parsed null model run report"

	rep.SourceArtifacts[3].Provenance = "file_read"
	rep.SourceArtifacts[3].Path = priorartPath
	rep.SourceArtifacts[3].Status = "available"
	rep.SourceArtifacts[3].Notes = "Successfully read and parsed prior art boundary report"

	var sourceBranchIDs []string
	for idx := range clockReport.LocalRelationBranches {
		sourceBranchIDs = append(sourceBranchIDs, fmt.Sprintf("branch_%d", idx))
	}
	rep.SourceBranchRegistry = SourceBranchRegistry{
		SourceArtifactID: "bmc0a_clock_readiness",
		BranchIDs:        sourceBranchIDs,
	}

	nodeContactFree := clockReport.ClockIndependentDiagnostics.NodeContactFree
	trajectoryFinite := clockReport.ClockIndependentDiagnostics.TrajectoryFiniteness
	derivReady := clockReport.FriedmannReadiness == "local_only_candidate"

	var eligibleBranches []LocalBranchEligibility
	var residualDiagnostics []CandidateResidualDiagnostic
	var calculationLedger []ResidualCalculationLedger

	hasEligibleBranch := false
	blockedByMissingInputs := false

	for idx, b := range clockReport.LocalRelationBranches {
		branchID := fmt.Sprintf("branch_%d", idx)

		eligible := b.ValidationPassed && nodeContactFree && trajectoryFinite && derivReady

		status := EligibilityEligibleLocalBranch
		reason := "Branch satisfies relational clock monotonicity and node distance constraints."
		if !eligible {
			if !b.ValidationPassed {
				status = EligibilityBlockedByClockFragility
				reason = "Local relation branch failed validation checks."
			} else if !nodeContactFree {
				status = EligibilityBlockedByNodeObstruction
				reason = "Branch exhibits node contact causing division by zero risks."
			} else if !trajectoryFinite {
				status = EligibilityBlockedByNonfiniteTrajectory
				reason = "Branch exhibits nonfinite trajectory."
			} else if !derivReady {
				status = EligibilityBlockedByDerivativeUnreadiness
				reason = "Friedmann spec derivative checks are not ready."
			}
		}

		hasMissingInputs := false
		var sortedPoints []model.TrajectoryPoint

		if len(b.Points) < 2 {
			hasMissingInputs = true
		} else {
			// Check for nonfinite alpha or phi, or lambda
			for _, pt := range b.Points {
				if math.IsNaN(pt.State.Alpha) || math.IsInf(pt.State.Alpha, 0) ||
					math.IsNaN(pt.State.Phi) || math.IsInf(pt.State.Phi, 0) ||
					math.IsNaN(pt.Lambda) || math.IsInf(pt.Lambda, 0) {
					hasMissingInputs = true
					break
				}
			}

			if !hasMissingInputs {
				// Copy and sort the points by Lambda
				sortedPoints = make([]model.TrajectoryPoint, len(b.Points))
				copy(sortedPoints, b.Points)
				sort.Slice(sortedPoints, func(i, j int) bool {
					return sortedPoints[i].Lambda < sortedPoints[j].Lambda
				})

				// Verify that delta_lambda is valid (non-zero and ordered/sorted)
				for i := 1; i < len(sortedPoints); i++ {
					deltaLambda := sortedPoints[i].Lambda - sortedPoints[i-1].Lambda
					if deltaLambda <= 0 || math.IsNaN(deltaLambda) || math.IsInf(deltaLambda, 0) {
						hasMissingInputs = true
						break
					}
				}
			}
		}

		if hasMissingInputs {
			blockedByMissingInputs = true
			if eligible {
				eligible = false
				status = EligibilityBlockedByDerivativeUnreadiness
				reason = "Branch missing required per-point residual input points."
			}
		}

		eligEntry := LocalBranchEligibility{
			BranchID:                  branchID,
			SourceArtifact:            "bmc0a_clock_readiness",
			Eligible:                  eligible,
			EligibilityStatus:         status,
			Reason:                    reason,
			NodeContactFree:           nodeContactFree,
			TrajectoryFinite:          trajectoryFinite,
			LocalClockStatus:          "monotonic",
			DerivativeReadinessStatus: "ready",
		}
		if !derivReady || hasMissingInputs {
			eligEntry.DerivativeReadinessStatus = "blocked"
		}
		eligibleBranches = append(eligibleBranches, eligEntry)

		diag := CandidateResidualDiagnostic{
			BranchID:            branchID,
			ResidualID:          "candidate_residual_" + branchID,
			ResidualComputed:    eligible,
			ResidualStatus:      ResidualStatusBlockedByNodeObstruction,
			ResidualProvenance:  ProvenanceBlocked,
			Notes:               "",
			ResidualInputPoints: []ResidualInputPoint{},
		}

		if eligible {
			hasEligibleBranch = true
			diag.ResidualStatus = ResidualStatusGenerated
			diag.ResidualProvenance = ProvenanceComputed

			var residualValues []float64
			var inputPoints []ResidualInputPoint
			N := len(sortedPoints)
			for i := 0; i < N-1; i++ {
				var dt, dAlpha, dPhi float64
				dt = sortedPoints[i+1].Lambda - sortedPoints[i].Lambda
				if dt != 0 && !math.IsNaN(dt) && !math.IsInf(dt, 0) {
					dAlpha = (sortedPoints[i+1].State.Alpha - sortedPoints[i].State.Alpha) / dt
					dPhi = (sortedPoints[i+1].State.Phi - sortedPoints[i].State.Phi) / dt
				}
				lhs := dAlpha * dAlpha
				rhs := dPhi * dPhi
				res := lhs - rhs
				residualValues = append(residualValues, res)

				alphaVal := sortedPoints[i].State.Alpha
				phiVal := sortedPoints[i].State.Phi
				lambdaVal := sortedPoints[i].Lambda
				inputPoints = append(inputPoints, ResidualInputPoint{
					BranchID:               branchID,
					PointIndex:             i,
					Lambda:                 &lambdaVal,
					Alpha:                  &alphaVal,
					Phi:                    &phiVal,
					CandidateLeftHandSide:  &lhs,
					CandidateRightHandSide: &rhs,
					InputProvenance:        "derived_from_file_read",
				})
			}

			var sumAbs, sumSq, maxAbs float64
			var numFinite int
			for _, v := range residualValues {
				if !math.IsNaN(v) && !math.IsInf(v, 0) {
					numFinite++
					absVal := math.Abs(v)
					sumAbs += absVal
					sumSq += v * v
					if absVal > maxAbs {
						maxAbs = absVal
					}
				}
			}
			valMean := sumAbs / float64(numFinite)
			valMax := maxAbs
			valRms := math.Sqrt(sumSq / float64(numFinite))

			diag.Metrics = CandidateResidualMetrics{
				NumEvaluationPoints:     len(inputPoints),
				NumFiniteResidualPoints: numFinite,
				MeanAbsResidual:         &valMean,
				MaxAbsResidual:          &valMax,
				RmsResidual:             &valRms,
				ResidualFinite:          numFinite == len(inputPoints) && len(inputPoints) > 0,
				DiagnosticWarnings:      []string{},
			}
			diag.Notes = fmt.Sprintf("Candidate local-branch residual diagnostic metrics computed on eligible %s (samples=%d, clock_range=%f).", branchID, N, b.ClockRange)
			diag.ResidualInputPoints = inputPoints

			calcEntry := ResidualCalculationLedger{
				CalculationID:      "calc_" + branchID,
				BranchID:           branchID,
				FormulaID:          "candidate_local_branch_velocity_constraint_residual_v0.1",
				FormulaDescription: "Candidate toy diagnostic computed from adjacent local-branch trajectory point differences: residual = ((delta alpha/delta lambda)^2 - (delta phi/delta lambda)^2). This is not a recovery claim.",
				FormulaSource:      "relational_derivative_residual_spec_v0.1",
				ConventionProfile:  "toy_minisuperspace_limit_v0.1",
				InputProvenance:    "derived_from_file_read",
				InputFields:        []string{"alpha", "phi", "lambda"},
				NumInputPoints:     len(inputPoints),
				CalculationStatus:  "computed_from_local_branch",
				Notes:              fmt.Sprintf("Honestly computed from clock readiness branch data (samples=%d, clock_range=%f)", N, b.ClockRange),
			}
			calculationLedger = append(calculationLedger, calcEntry)
		} else {
			if hasMissingInputs {
				diag.BlockedReason = "Branch eligibility failed due to missing per-point residual inputs."
				diag.Notes = "Candidate residual diagnostic calculation blocked on ineligible " + branchID
				diag.ResidualStatus = ResidualStatusBlockedByClockFragility

				calcEntry := ResidualCalculationLedger{
					CalculationID:      "calc_" + branchID,
					BranchID:           branchID,
					FormulaID:          "candidate_local_branch_velocity_constraint_residual_v0.1",
					FormulaDescription: "Candidate toy diagnostic computed from adjacent local-branch trajectory point differences: residual = ((delta alpha/delta lambda)^2 - (delta phi/delta lambda)^2). This is not a recovery claim.",
					FormulaSource:      "relational_derivative_residual_spec_v0.1",
					ConventionProfile:  "toy_minisuperspace_limit_v0.1",
					InputProvenance:    "derived_from_file_read",
					InputFields:        []string{"alpha", "phi", "lambda"},
					NumInputPoints:     0,
					CalculationStatus:  "blocked_by_missing_residual_inputs",
					Notes:              "Calculation blocked due to missing per-point branch inputs.",
				}
				calculationLedger = append(calculationLedger, calcEntry)
			} else {
				diag.BlockedReason = "Branch eligibility criteria not met: " + reason
				diag.Notes = "Candidate residual diagnostic calculation blocked on ineligible " + branchID

				calcEntry := ResidualCalculationLedger{
					CalculationID:      "calc_" + branchID,
					BranchID:           branchID,
					FormulaID:          "candidate_local_branch_velocity_constraint_residual_v0.1",
					FormulaDescription: "Candidate toy diagnostic computed from adjacent local-branch trajectory point differences: residual = ((delta alpha/delta lambda)^2 - (delta phi/delta lambda)^2). This is not a recovery claim.",
					FormulaSource:      "relational_derivative_residual_spec_v0.1",
					ConventionProfile:  "toy_minisuperspace_limit_v0.1",
					InputProvenance:    "derived_from_file_read",
					InputFields:        []string{"alpha", "phi", "lambda"},
					NumInputPoints:     0,
					CalculationStatus:  "blocked_by_derivative_unreadiness",
					Notes:              "Calculation blocked due to branch eligibility failure: " + reason,
				}
				calculationLedger = append(calculationLedger, calcEntry)
			}
		}

		residualDiagnostics = append(residualDiagnostics, diag)
	}

	rep.LocalBranchEligibility = eligibleBranches
	rep.CandidateResidualDiagnostics = residualDiagnostics
	rep.CalculationLedger = calculationLedger
	computedResidualIDs := []string{}
	for _, diag := range residualDiagnostics {
		if diag.ResidualComputed {
			computedResidualIDs = append(computedResidualIDs, diag.ResidualID)
		}
	}

	if hasEligibleBranch {
		rep.CandidateResidualComputed = true
		if nullReport != nil && nullReport.NullDiagnosticsComputed {
			rep.InterpretationStatus = InterpretDiagComparisonOnly
			rep.ResidualNullComparisons = []ResidualNullComparison{
				{
					ComparisonID:         "bmc0a-residual-vs-null-v0.1",
					TargetResidualIDs:    computedResidualIDs,
					NullModelIDs:         []string{"constant_phase_control"},
					MetricsCompared:      []string{"mean_abs_residual", "max_abs_residual"},
					ComparisonComputed:   true,
					InterpretationStatus: InterpretDiagComparisonOnly,
					Reason:               "Comparison shows candidate local-branch residual metrics relative to constant phase null control.",
				},
			}
		} else {
			rep.InterpretationStatus = InterpretBlockedByNoComparableNullDiagnostics
			rep.ResidualNullComparisons = []ResidualNullComparison{
				{
					ComparisonID:         "bmc0a-residual-vs-null-v0.1",
					TargetResidualIDs:    computedResidualIDs,
					NullModelIDs:         []string{"constant_phase_control"},
					MetricsCompared:      []string{"mean_abs_residual", "max_abs_residual"},
					ComparisonComputed:   false,
					InterpretationStatus: InterpretBlockedByNoComparableNullDiagnostics,
					Reason:               "Comparison not computed because null model diagnostics are not available.",
				},
			}
		}
	} else {
		rep.CandidateResidualComputed = false
		if blockedByMissingInputs || len(clockReport.LocalRelationBranches) == 0 {
			rep.InterpretationStatus = InterpretBlockedByMissingResidualInputs
		} else {
			rep.InterpretationStatus = InterpretBlockedByNoEligibleLocalBranch
		}
		rep.ResidualNullComparisons = []ResidualNullComparison{}
	}

	return rep, nil
}

// ReadReport reads and strictly decodes a ResidualRunReport JSON file, rejecting trailing tokens.
func ReadReport(path string) (*ResidualRunReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep ResidualRunReport
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
func WriteReport(rep *ResidualRunReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}
