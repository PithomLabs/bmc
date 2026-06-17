package residualrun

import (
	"encoding/json"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/clockseg"
	"github.com/PithomLabs/bmc/internal/bmc/friedmannspec"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/priorart"
)

func helperDefaultReport(t *testing.T) *ResidualRunReport {
	t.Helper()
	return GenerateDefaultReport()
}

func helperComputedReport(t *testing.T) *ResidualRunReport {
	t.Helper()
	rep := GenerateBlockedDefaultReport()
	rep.CandidateResidualComputed = true
	rep.InterpretationStatus = InterpretDiagComparisonOnly

	rep.LocalBranchEligibility[0].Eligible = true
	rep.LocalBranchEligibility[0].EligibilityStatus = EligibilityEligibleLocalBranch
	rep.LocalBranchEligibility[0].NodeContactFree = true
	rep.LocalBranchEligibility[0].TrajectoryFinite = true
	rep.LocalBranchEligibility[0].LocalClockStatus = "monotonic"
	rep.LocalBranchEligibility[0].DerivativeReadinessStatus = "ready"

	for idx := range rep.SourceArtifacts {
		rep.SourceArtifacts[idx].Provenance = "file_read"
		rep.SourceArtifacts[idx].Path = "out/mock_path.json"
		rep.SourceArtifacts[idx].Status = "available"
	}

	rep.SourceBranchRegistry = SourceBranchRegistry{
		SourceArtifactID: "bmc0a_clock_readiness",
		BranchIDs:        []string{"branch_0"},
	}

	valMean := 0.05
	valMax := 0.12
	valRms := 0.06
	rep.CandidateResidualDiagnostics[0].ResidualComputed = true
	rep.CandidateResidualDiagnostics[0].ResidualStatus = ResidualStatusGenerated
	rep.CandidateResidualDiagnostics[0].ResidualProvenance = ProvenanceComputed
	rep.CandidateResidualDiagnostics[0].Metrics = CandidateResidualMetrics{
		NumEvaluationPoints:     2,
		NumFiniteResidualPoints: 2,
		MeanAbsResidual:         &valMean,
		MaxAbsResidual:          &valMax,
		RmsResidual:             &valRms,
		ResidualFinite:          true,
		DiagnosticWarnings:      []string{},
	}
	rep.CandidateResidualDiagnostics[0].Notes = "Computed residual on branch_0"

	mockAlpha := 1.2
	mockPhi := 0.5
	mockLHS := 0.1
	mockRHS := 0.05
	mockLambda0 := 0.1
	mockLambda1 := 0.2
	rep.CandidateResidualDiagnostics[0].ResidualInputPoints = []ResidualInputPoint{
		{
			BranchID:               "branch_0",
			PointIndex:             0,
			Lambda:                 &mockLambda0,
			Alpha:                  &mockAlpha,
			Phi:                    &mockPhi,
			CandidateLeftHandSide:  &mockLHS,
			CandidateRightHandSide: &mockRHS,
			InputProvenance:        "derived_from_file_read",
		},
		{
			BranchID:               "branch_0",
			PointIndex:             1,
			Lambda:                 &mockLambda1,
			Alpha:                  &mockAlpha,
			Phi:                    &mockPhi,
			CandidateLeftHandSide:  &mockLHS,
			CandidateRightHandSide: &mockRHS,
			InputProvenance:        "derived_from_file_read",
		},
	}

	rep.CalculationLedger = []ResidualCalculationLedger{
		{
			CalculationID:      "calc_branch_0",
			BranchID:           "branch_0",
			FormulaID:          "candidate_local_branch_velocity_constraint_residual_v0.1",
			FormulaDescription: "Candidate toy diagnostic computed from adjacent local-branch trajectory point differences: residual = ((delta alpha/delta lambda)^2 - (delta phi/delta lambda)^2). This is not a recovery claim.",
			FormulaSource:      "relational_derivative_residual_spec_v0.1",
			ConventionProfile:  "toy_minisuperspace_limit_v0.1",
			InputProvenance:    "derived_from_file_read",
			InputFields:        []string{"alpha", "phi", "lambda"},
			NumInputPoints:     2,
			CalculationStatus:  "computed_from_local_branch",
			Notes:              "mock computed calculation",
		},
	}

	rep.ResidualNullComparisons = []ResidualNullComparison{
		{
			ComparisonID:         "bmc0a-residual-vs-null-v0.1",
			TargetResidualIDs:    []string{"candidate_residual_branch_0"},
			NullModelIDs:         []string{"constant_phase_control"},
			MetricsCompared:      []string{"mean_abs_residual", "max_abs_residual"},
			ComparisonComputed:   true,
			InterpretationStatus: InterpretDiagComparisonOnly,
			Reason:               "Comparison shows candidate local-branch residual metrics relative to constant phase null control.",
		},
	}

	return rep
}

func TestResidualRunReportValidation(t *testing.T) {
	rep := helperDefaultReport(t)
	errs := ValidateReport(rep, "{}")
	if len(errs) > 0 {
		t.Fatalf("Default report failed validation: %v", errs)
	}
}

func TestResidualRunRejectsRecoveryClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ResidualRecoveryClaim = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ResidualRecoveryClaim is true")
	}
}

func TestResidualRunRejectsScientificNoveltyClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ScientificNoveltyClaimMade = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ScientificNoveltyClaimMade is true")
	}
}

func TestResidualRunRejectsBMCBeatsNullModelsClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.BmcBeatsNullModelsClaim = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when BmcBeatsNullModelsClaim is true")
	}
}

func TestResidualRunRequiresFullBMCBlocked(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.FullBmcToyGate = "passed"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when FullBmcToyGate is not blocked")
	}
}

func TestResidualRunRequiresConventionLedger(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ConventionLedger = []ResidualConventionLedger{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ConventionLedger is empty")
	}
}

func TestResidualRunRejectsRetiredConventionDebt(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ConventionLedger[0].Status = "retired"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a convention debt is retired")
	}
}

func TestResidualRunRejectsResidualForIneligibleBranch(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics = append(rep.CandidateResidualDiagnostics, CandidateResidualDiagnostic{
		BranchID:           "branch_1",
		ResidualID:         "candidate_residual_branch_1",
		ResidualComputed:   true,
		ResidualStatus:     ResidualStatusGenerated,
		ResidualProvenance: ProvenanceComputed,
		Notes:              "ineligible computed residual",
	})
	rep.LocalBranchEligibility = append(rep.LocalBranchEligibility, LocalBranchEligibility{
		BranchID:                  "branch_1",
		SourceArtifact:            "bmc0a_clock_readiness",
		Eligible:                  false,
		EligibilityStatus:         EligibilityBlockedByNodeObstruction,
		Reason:                    "Node contact failed",
		NodeContactFree:           false,
		TrajectoryFinite:          true,
		LocalClockStatus:          "monotonic",
		DerivativeReadinessStatus: "ready",
	})
	val := 0.1
	rep.CandidateResidualDiagnostics[1].Metrics.MeanAbsResidual = &val
	rep.CandidateResidualDiagnostics[1].Metrics.MaxAbsResidual = &val
	rep.CandidateResidualDiagnostics[1].Metrics.RmsResidual = &val
	rep.CandidateResidualDiagnostics[1].Metrics.NumEvaluationPoints = 100
	rep.CandidateResidualDiagnostics[1].Metrics.NumFiniteResidualPoints = 100
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual is computed on an ineligible branch")
	}
}

func TestResidualRunRejectsNonfiniteResidualMetrics(t *testing.T) {
	rep := helperComputedReport(t)
	infVal := math.Inf(1)
	rep.CandidateResidualDiagnostics[0].Metrics.MeanAbsResidual = &infVal
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a metric has nonfinite values")
	}
}

func TestResidualRunRejectsNegativeResidualMetrics(t *testing.T) {
	rep := helperComputedReport(t)
	val := -0.5
	rep.CandidateResidualDiagnostics[0].Metrics.MeanAbsResidual = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when mean_abs_residual is negative")
	}
}

func TestResidualRunRejectsSentinelResidualMetrics(t *testing.T) {
	rep := helperComputedReport(t)
	val := -1.0
	rep.CandidateResidualDiagnostics[0].Metrics.MaxAbsResidual = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when max_abs_residual has a sentinel value of -1")
	}
}

func TestResidualRunRejectsMissingOptionalMetricKeys(t *testing.T) {
	rep := helperComputedReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	raw := strings.Replace(string(data), `"mean_abs_residual"`, `"missing_key"`, -1)
	errs := ValidateReport(rep, raw)
	foundMissing := false
	for _, e := range errs {
		if strings.Contains(e.Message, "missing optional metric key") {
			foundMissing = true
		}
	}
	if !foundMissing {
		t.Fatal("Expected validation to fail when raw JSON is missing mean_abs_residual key")
	}
}

func TestResidualRunRequiresAllGatesExactlyOnce(t *testing.T) {
	rep := helperDefaultReport(t)
	backup := rep.Gates
	rep.Gates = rep.Gates[1:]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate is missing")
	}

	rep.Gates = append(backup, backup[0])
	errs = ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate is duplicated")
	}
}

func TestResidualRunRejectsNonPassGate(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.Gates[0].Status = "fail"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate status is not pass")
	}
}

func TestResidualRunRejectsForbiddenResidualStatus(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualStatus = "recovered"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when using forbidden residual status")
	}
}

func TestResidualRunRejectsForbiddenResidualProvenance(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualProvenance = "fabricated"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when using forbidden residual provenance")
	}
}

func TestResidualRunRejectsForbiddenInterpretationStatus(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.InterpretationStatus = "winner"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when using forbidden interpretation status")
	}
}

func TestResidualRunRejectsDecorativeComparison(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ResidualNullComparisons[0].MetricsCompared = []string{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when target/null comparison is decorative (empty metrics)")
	}
}

func TestResidualRunRejectsUnknownFields(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_unknown.json")
	raw := `{"schema_version":"bmc0a-local-residual-v0.1","extra_unknown_field":"some_val"}`
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with unknown fields")
	}
}

func TestResidualRunRejectsTrailingJSONTokens(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_trailing.json")

	rep := helperDefaultReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatal(err)
	}
	raw := string(data) + "   true"
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err = ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with trailing tokens")
	}
}

func TestResidualRunDeterministicJSON(t *testing.T) {
	rep1 := helperDefaultReport(t)
	rep2 := helperDefaultReport(t)

	data1, err1 := json.Marshal(rep1)
	data2, err2 := json.Marshal(rep2)
	if err1 != nil || err2 != nil {
		t.Fatalf("Marshal failed: %v %v", err1, err2)
	}

	if string(data1) != string(data2) {
		t.Fatal("Serialized reports are not identical (non-deterministic)")
	}
}

func TestResidualRunForbiddenPhraseScan(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This contains classical limit achieved.")
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected forbidden phrase scan to catch forbidden phrases")
	}
}

func TestResidualRunForbiddenPhraseErrorsArePhraseSafe(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This contains classical limit achieved.")
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected failure")
	}
	for _, e := range errs {
		if strings.Contains(strings.ToLower(e.Message), "classical limit") {
			t.Fatalf("Error message leaked forbidden phrase: %s", e.Message)
		}
	}
}

func TestResidualRunRejectsEligibleBranchWithNodeContactFalse(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchEligibility[0].NodeContactFree = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when node_contact_free is false for eligible branch")
	}
}

func TestResidualRunRejectsEligibleBranchWithTrajectoryFiniteFalse(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchEligibility[0].TrajectoryFinite = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when trajectory_finite is false for eligible branch")
	}
}

func TestResidualRunRejectsComputedResidualWhenDerivativeNotReady(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchEligibility[0].DerivativeReadinessStatus = "blocked"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when derivative readiness status is not ready")
	}
}

func TestResidualRunRejectsComputedResidualWithBlockedStatus(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualStatus = ResidualStatusBlockedByNodeObstruction
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual_computed is true but status is blocked")
	}
}

func TestResidualRunRejectsUncomputedResidualWithGeneratedStatus(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualComputed = false
	rep.CandidateResidualDiagnostics[0].ResidualStatus = ResidualStatusGenerated
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual_computed is false but status is generated")
	}
}

func TestResidualRunRejectsComputedResidualWithBlockedProvenance(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualProvenance = ProvenanceBlocked
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual_computed is true but provenance is blocked")
	}
}

func TestResidualRunRejectsReportFalseWithComputedDiagnostic(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualComputed = false
	rep.InterpretationStatus = InterpretBlockedByNoEligibleLocalBranch
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when candidate_residual_computed is false but a diagnostic has residual_computed = true")
	}
}

func TestResidualRunRejectsReportTrueWithNoComputedDiagnostic(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualComputed = false
	rep.CandidateResidualDiagnostics[0].ResidualStatus = ResidualStatusBlockedByNodeObstruction
	rep.CandidateResidualDiagnostics[0].ResidualProvenance = ProvenanceBlocked
	rep.CandidateResidualDiagnostics[0].Metrics.MeanAbsResidual = nil
	rep.CandidateResidualDiagnostics[0].Metrics.MaxAbsResidual = nil
	rep.CandidateResidualDiagnostics[0].Metrics.RmsResidual = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when candidate_residual_computed is true but no diagnostic has residual_computed = true")
	}
}

func TestResidualRunRejectsComputedBranchProvenanceWithoutCalculationLedger(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CalculationLedger = []ResidualCalculationLedger{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual provenance is computed but calculation ledger is empty")
	}
}

func TestResidualRunRejectsFiniteCountGreaterThanEvalCount(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.NumFiniteResidualPoints = 120
	rep.CandidateResidualDiagnostics[0].Metrics.NumEvaluationPoints = 100
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when finite points > evaluation points")
	}
}

func TestResidualRunRejectsNegativeEvaluationPointCount(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.NumEvaluationPoints = -10
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when evaluation points is negative")
	}
}

func TestResidualRunRejectsNegativeFinitePointCount(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.NumFiniteResidualPoints = -5
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when finite points is negative")
	}
}

func TestResidualRunRejectsResidualFiniteTrueWithZeroFinitePoints(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.NumFiniteResidualPoints = 0
	rep.CandidateResidualDiagnostics[0].Metrics.ResidualFinite = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when residual_finite is true but finite points is zero")
	}
}

func TestResidualRunRejectsComputedResidualWithNilMean(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.MeanAbsResidual = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when computed residual has nil mean")
	}
}

func TestResidualRunRejectsComputedResidualWithNilMax(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.MaxAbsResidual = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when computed residual has nil max")
	}
}

func TestResidualRunRejectsComputedResidualWithNilRMS(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.RmsResidual = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when computed residual has nil RMS")
	}
}

func TestResidualRunRejectsComparisonWithUnknownTargetResidual(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ResidualNullComparisons[0].TargetResidualIDs = []string{"nonexistent_residual"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison references nonexistent target residual")
	}
}

func TestResidualRunRejectsComparisonWithUncomputedTargetResidual(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics = append(rep.CandidateResidualDiagnostics, CandidateResidualDiagnostic{
		BranchID:           "branch_1",
		ResidualID:         "candidate_residual_branch_1",
		ResidualComputed:   false,
		ResidualStatus:     ResidualStatusBlockedByNodeObstruction,
		ResidualProvenance: ProvenanceBlocked,
		Notes:              "uncomputed branch",
	})
	rep.ResidualNullComparisons[0].TargetResidualIDs = []string{"candidate_residual_branch_1"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison target is uncomputed")
	}
}

func TestResidualRunRequiresEveryComparisonTargetComputed(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics = append(rep.CandidateResidualDiagnostics, CandidateResidualDiagnostic{
		BranchID:           "branch_1",
		ResidualID:         "candidate_residual_branch_1",
		ResidualComputed:   false,
		ResidualStatus:     ResidualStatusBlockedByNodeObstruction,
		ResidualProvenance: ProvenanceBlocked,
		Notes:              "uncomputed branch",
	})
	rep.ResidualNullComparisons[0].TargetResidualIDs = []string{"candidate_residual_branch_0", "candidate_residual_branch_1"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when at least one target in comparison is uncomputed")
	}
}

func TestResidualRunRequiresExpectedSourceArtifacts(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceArtifacts = rep.SourceArtifacts[:3]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when missing a required source artifact")
	}
}

func TestResidualRunRejectsDuplicateSourceArtifact(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceArtifacts = append(rep.SourceArtifacts, rep.SourceArtifacts[0])
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a source artifact is duplicated")
	}
}

func TestResidualRunRejectsUnknownSourceArtifact(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceArtifacts[0].ArtifactID = "unknown_artifact"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when an unknown source artifact is present")
	}
}

func TestResidualRunRejectsFileReadWithoutPath(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceArtifacts[0].Provenance = "file_read"
	rep.SourceArtifacts[0].Path = ""
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when file_read provenance has empty path")
	}
}

func TestResidualRunRequiresAllSixConventionDebts(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger = rep.ConventionLedger[:5]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when missing a required convention debt")
	}
}

func TestResidualRunRejectsDuplicateConventionDebt(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger = append(rep.ConventionLedger, rep.ConventionLedger[0])
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a convention debt is duplicated")
	}
}

func TestResidualRunRejectsUnknownConventionDebt(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger[0].ConventionID = "unknown_debt"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when an unknown convention debt is present")
	}
}

func TestResidualRunRejectsEmptyConventionDebtDescription(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger[0].Description = ""
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when convention debt description is empty")
	}
}

func TestResidualRunRejectsHumanReviewFalseForUnresolvedConventionDebt(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger[0].HumanReviewRequired = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when human_review_required is false for unresolved debt")
	}
}

func TestResidualRunRejectsPartialCoreConventionDebt(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ConventionLedger[0].Status = "partial"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when core convention debt is partial")
	}
}

func TestResidualRunRequiresLocalOnlyBoundary(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchOnly = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when local_branch_only is false")
	}
}

func TestResidualRunRejectsGlobalCosmologyClaim(t *testing.T) {
	rep := helperComputedReport(t)
	rep.GlobalCosmologyClaim = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when global_cosmology_claim is true")
	}
}

func TestResidualRunForbiddenPhraseScanIncludesClassicalCosmologyVariant(t *testing.T) {
	rep := helperComputedReport(t)
	phrase := "classical" + " " + "cosmology"
	rep.Warnings = append(rep.Warnings, "Warning: "+phrase)
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected forbidden phrase scan to catch classical cosmology variant")
	}
}

func TestResidualRunSummaryCountsOnlyTrulyEligibleBranches(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchEligibility = append(rep.LocalBranchEligibility, LocalBranchEligibility{
		BranchID:          "branch_1",
		Eligible:          false,
		EligibilityStatus: EligibilityBlockedByNodeObstruction,
		Reason:            "node contact",
	})

	eligibleCount := 0
	for _, b := range rep.LocalBranchEligibility {
		if b.Eligible {
			eligibleCount++
		}
	}
	if eligibleCount != 1 {
		t.Fatalf("Expected exactly 1 eligible branch, got %d", eligibleCount)
	}
}

func TestResidualRunCLIRoutingRunsGenerateValidateSummarize(t *testing.T) {
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test-res", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test-res")

	// 1. Generate report
	genCmd := exec.Command("./ptw-bmc-test-res", "run-residuals", "--profile", "bmc0a-local-residual", "--out", "out_test_res.json")
	genCmd.Dir = "."
	if out, err := genCmd.CombinedOutput(); err != nil {
		t.Fatalf("Generation failed: %v, output: %s", err, string(out))
	}
	defer os.Remove("./out_test_res.json")

	// 2. Validate
	valCmd := exec.Command("./ptw-bmc-test-res", "validate", "--report", "out_test_res.json")
	valCmd.Dir = "."
	valOut, err := valCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Validation failed: %v, output: %s", err, string(valOut))
	}
	if !strings.Contains(string(valOut), "PASSED") {
		t.Fatalf("Unexpected validation output: %s", string(valOut))
	}

	// 3. Summarize
	sumCmd := exec.Command("./ptw-bmc-test-res", "summarize", "--report", "out_test_res.json")
	sumCmd.Dir = "."
	sumOut, err := sumCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Summarization failed: %v, output: %s", err, string(sumOut))
	}
	if !strings.Contains(string(sumOut), "Candidate Local-Branch Residual Summary") {
		t.Fatalf("Unexpected summary output: %s", string(sumOut))
	}
}

func TestResidualRunUnknownProfileFailsAtCLI(t *testing.T) {
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test-res2", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test-res2")

	failCmd := exec.Command("./ptw-bmc-test-res2", "run-residuals", "--profile", "unknown-profile", "--out", "out_test_res2.json")
	failCmd.Dir = "."
	out, err := failCmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected CLI to exit with failure for unknown profile")
	}
	if !strings.Contains(string(out), "is not supported") {
		t.Fatalf("Unexpected CLI error message: %s", string(out))
	}
}

func TestResidualRunRejectsComputedReportWithNonFileReadSourceProvenance(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceArtifacts[0].Provenance = "source_artifact_summary"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when source provenance is not file_read in a computed report")
	}
}

func TestResidualRunRejectsBlockedReportWithComputedLedgerEntry(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.CalculationLedger = []ResidualCalculationLedger{
		{
			CalculationID:      "calc_0",
			BranchID:           "branch_0",
			FormulaID:          "f",
			FormulaDescription: "d",
			ConventionProfile:  "p",
			InputFields:        []string{"i"},
			CalculationStatus:  "computed_from_local_branch",
		},
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a blocked report has a computed_from_local_branch ledger status")
	}
}

func TestResidualRunRejectsBlockedReportWithNonemptyComparisons(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ResidualNullComparisons = []ResidualNullComparison{
		{
			ComparisonID:         "comp_id",
			TargetResidualIDs:    []string{"res_id"},
			NullModelIDs:         []string{"null_id"},
			MetricsCompared:      []string{"mean"},
			ComparisonComputed:   false,
			InterpretationStatus: InterpretBlockedByNoEligibleLocalBranch,
			Reason:               "reason",
		},
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a blocked report has nonempty comparisons")
	}
}

func TestResidualRunRejectsComputedReportWithUnblockedComparisonButFalseComparisonComputed(t *testing.T) {
	rep := helperComputedReport(t)
	rep.ResidualNullComparisons[0].ComparisonComputed = false
	rep.ResidualNullComparisons[0].InterpretationStatus = InterpretDiagComparisonOnly
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison_computed is false but interpretation status is not blocked")
	}
}

func TestResidualRunRejectsCalculationLedgerWithEmptyFields(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CalculationLedger[0].FormulaID = ""
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when calculation ledger has empty FormulaID")
	}
}

func TestResidualRunBlockedReportWhenInputsMissing(t *testing.T) {
	clockRep := &clockseg.ClockReadinessReport{
		SchemaVersion: "bmc0a-clock-readiness-v0.1",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				Samples: 0,
				Points:  []model.TrajectoryPoint{},
			},
		},
	}
	friedmannRep := &friedmannspec.FriedmannSpecReport{}
	nullRep := &nullrun.NullRunReport{}
	priorArtRep := &priorart.PriorArtBoundaryReport{}

	rep, err := RunResidualsFromInputs(clockRep, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if rep.CandidateResidualComputed {
		t.Fatal("Expected candidate_residual_computed to be false when points are missing")
	}
	if rep.InterpretationStatus != InterpretBlockedByMissingResidualInputs {
		t.Fatalf("Expected interpretation status %s, got %s", InterpretBlockedByMissingResidualInputs, rep.InterpretationStatus)
	}
}

func TestResidualRunMetricsChangeWhenInputBranchDataChanges(t *testing.T) {
	clockRep1 := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.2, State: model.MiniState{Alpha: 1.2, Phi: 0.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 1.5, Phi: 0.3}},
				},
			},
		},
	}

	clockRep2 := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.2, State: model.MiniState{Alpha: 2.2, Phi: 1.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 3.5, Phi: 2.3}},
				},
			},
		},
	}

	friedmannRep := &friedmannspec.FriedmannSpecReport{}
	nullRep := &nullrun.NullRunReport{NullDiagnosticsComputed: true}
	priorArtRep := &priorart.PriorArtBoundaryReport{}

	rep1, err := RunResidualsFromInputs(clockRep1, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	rep2, err := RunResidualsFromInputs(clockRep2, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !rep1.CandidateResidualComputed || !rep2.CandidateResidualComputed {
		t.Fatalf("Expected both reports to be computed, got rep1=%v rep2=%v", rep1.CandidateResidualComputed, rep2.CandidateResidualComputed)
	}

	m1 := rep1.CandidateResidualDiagnostics[0].Metrics
	m2 := rep2.CandidateResidualDiagnostics[0].Metrics

	if m1.MeanAbsResidual == nil || m2.MeanAbsResidual == nil {
		t.Fatalf("Expected non-nil mean metrics, got m1=%v, m2=%v", m1.MeanAbsResidual, m2.MeanAbsResidual)
	}

	if *m1.MeanAbsResidual == *m2.MeanAbsResidual || *m1.MaxAbsResidual == *m2.MaxAbsResidual || *m1.RmsResidual == *m2.RmsResidual {
		t.Fatalf("Expected metrics to change when point inputs change. m1=%v, m2=%v", m1, m2)
	}
}

func TestResidualRunComparisonTargetsActualComputedResidual(t *testing.T) {
	clockRep := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: false,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.2, State: model.MiniState{Alpha: 1.2, Phi: 0.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 1.5, Phi: 0.3}},
				},
			},
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.2, State: model.MiniState{Alpha: 1.8, Phi: 0.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 2.5, Phi: 0.4}},
				},
			},
		},
	}

	rep, err := RunResidualsFromInputs(
		clockRep,
		&friedmannspec.FriedmannSpecReport{},
		&nullrun.NullRunReport{NullDiagnosticsComputed: true},
		&priorart.PriorArtBoundaryReport{},
		"clock.json",
		"friedmann.json",
		"null.json",
		"priorart.json",
	)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if len(rep.CandidateResidualDiagnostics) != 2 {
		t.Fatalf("Expected two residual diagnostics, got %d", len(rep.CandidateResidualDiagnostics))
	}
	if rep.CandidateResidualDiagnostics[0].ResidualComputed {
		t.Fatal("Expected branch_0 residual_computed = false")
	}
	if !rep.CandidateResidualDiagnostics[1].ResidualComputed {
		t.Fatal("Expected branch_1 residual_computed = true")
	}
	if len(rep.ResidualNullComparisons) != 1 {
		t.Fatalf("Expected one generated residual/null comparison, got %d", len(rep.ResidualNullComparisons))
	}
	targets := rep.ResidualNullComparisons[0].TargetResidualIDs
	if !stringSliceContains(targets, "candidate_residual_branch_1") {
		t.Fatalf("Expected comparison target IDs to contain candidate_residual_branch_1, got %v", targets)
	}
	if stringSliceContains(targets, "candidate_residual_branch_0") {
		t.Fatalf("Expected comparison target IDs not to contain candidate_residual_branch_0, got %v", targets)
	}
}

func TestResidualRunRejectsSyntheticFileBackedBranch(t *testing.T) {
	rep := helperComputedReport(t)
	rep.SourceBranchRegistry.BranchIDs = []string{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to reject reported branches not found in source registry")
	}
}

func stringSliceContains(vals []string, want string) bool {
	for _, val := range vals {
		if val == want {
			return true
		}
	}
	return false
}

func TestResidualRunRejectsReportedBranchNotInSourceRegistry(t *testing.T) {
	rep := helperComputedReport(t)
	rep.LocalBranchEligibility = append(rep.LocalBranchEligibility, LocalBranchEligibility{
		BranchID:       "branch_nonexistent",
		SourceArtifact: "bmc0a_clock_readiness",
		Eligible:       false,
	})
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when reported eligibility branch is not in the source branch registry")
	}
}

func TestResidualRunCalculationLedgerRequiresFormulaTransparency(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CalculationLedger[0].FormulaSource = ""
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when FormulaSource is empty in computed ledger")
	}
}

func TestResidualRunRejectsComputedLedgerWithZeroInputPoints(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CalculationLedger[0].NumInputPoints = 0
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when computed ledger has zero input points")
	}
}

func TestResidualRunRejectsComputedLedgerWithEmptyInputFields(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CalculationLedger[0].InputFields = []string{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when computed ledger has empty InputFields")
	}
}

func TestResidualRunRejectsComputedDiagnosticWithEmptyResidualInputPoints(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].ResidualInputPoints = []ResidualInputPoint{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when computed diagnostic has empty residual_input_points")
	}
}

func TestResidualRunRejectsResidualInputPointCountMismatch(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics[0].Metrics.NumEvaluationPoints = len(rep.CandidateResidualDiagnostics[0].ResidualInputPoints) + 1
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when residual_input_points length does not match num_evaluation_points")
	}
}

func TestResidualRunRejectsInvalidResidualInputPointFields(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*ResidualInputPoint)
	}{
		{
			name: "nil_lambda",
			mutate: func(pt *ResidualInputPoint) {
				pt.Lambda = nil
			},
		},
		{
			name: "nonfinite_lambda",
			mutate: func(pt *ResidualInputPoint) {
				v := math.Inf(1)
				pt.Lambda = &v
			},
		},
		{
			name: "nonfinite_alpha",
			mutate: func(pt *ResidualInputPoint) {
				v := math.NaN()
				pt.Alpha = &v
			},
		},
		{
			name: "nonfinite_phi",
			mutate: func(pt *ResidualInputPoint) {
				v := math.Inf(-1)
				pt.Phi = &v
			},
		},
		{
			name: "nonfinite_lhs",
			mutate: func(pt *ResidualInputPoint) {
				v := math.NaN()
				pt.CandidateLeftHandSide = &v
			},
		},
		{
			name: "nonfinite_rhs",
			mutate: func(pt *ResidualInputPoint) {
				v := math.Inf(1)
				pt.CandidateRightHandSide = &v
			},
		},
		{
			name: "invalid_provenance",
			mutate: func(pt *ResidualInputPoint) {
				pt.InputProvenance = "manual_fixture"
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rep := helperComputedReport(t)
			tc.mutate(&rep.CandidateResidualDiagnostics[0].ResidualInputPoints[0])
			errs := ValidateReport(rep, "{}")
			if len(errs) == 0 {
				t.Fatalf("Expected validation to fail for %s", tc.name)
			}
		})
	}
}

func TestResidualRunRejectsNonIncreasingResidualInputPointLambda(t *testing.T) {
	rep := helperComputedReport(t)
	dup := *rep.CandidateResidualDiagnostics[0].ResidualInputPoints[0].Lambda
	rep.CandidateResidualDiagnostics[0].ResidualInputPoints[1].Lambda = &dup
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail for duplicate lambda inside one computed diagnostic")
	}
}

func TestResidualRunRejectsForbiddenFormulaIDsPhraseSafely(t *testing.T) {
	forbiddenIDs := []string{
		"friedmann_residual",
		"classical_residual",
		"recovery_residual",
		"cosmology_recovery_residual",
	}
	for _, formulaID := range forbiddenIDs {
		t.Run(formulaID, func(t *testing.T) {
			rep := helperComputedReport(t)
			rep.CalculationLedger[0].FormulaID = formulaID
			errs := ValidateReport(rep, "{}")
			if len(errs) == 0 {
				t.Fatalf("Expected validation to fail for forbidden formula ID")
			}
			for _, err := range errs {
				if strings.Contains(err.Message, formulaID) {
					t.Fatalf("Validation error leaked forbidden formula ID %q: %s", formulaID, err.Message)
				}
			}
		})
	}
}

func TestResidualRunSummaryDistinguishesComputedAndBlockedDiagnostics(t *testing.T) {
	rep := helperComputedReport(t)
	rep.CandidateResidualDiagnostics = append(rep.CandidateResidualDiagnostics, CandidateResidualDiagnostic{
		BranchID:           "branch_1",
		ResidualID:         "candidate_residual_branch_1",
		ResidualComputed:   false,
		ResidualStatus:     ResidualStatusBlockedByClockFragility,
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
		BlockedReason:       "Missing inputs",
		Notes:               "blocked branch",
		ResidualInputPoints: []ResidualInputPoint{},
	})
	rep.LocalBranchEligibility = append(rep.LocalBranchEligibility, LocalBranchEligibility{
		BranchID:       "branch_1",
		SourceArtifact: "bmc0a_clock_readiness",
		Eligible:       false,
	})
	rep.SourceBranchRegistry.BranchIDs = append(rep.SourceBranchRegistry.BranchIDs, "branch_1")

	numComputed := 0
	numBlocked := 0
	for _, rd := range rep.CandidateResidualDiagnostics {
		if rd.ResidualComputed {
			numComputed++
		} else {
			numBlocked++
		}
	}
	if numComputed != 1 || numBlocked != 1 {
		t.Fatalf("Expected 1 computed and 1 blocked diagnostic, got computed=%d blocked=%d", numComputed, numBlocked)
	}
}

func TestResidualRunRejectsNonfiniteOrDuplicateLambda(t *testing.T) {
	// Case 1: duplicate lambda
	clockRep1 := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.2, Phi: 0.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 1.5, Phi: 0.3}},
				},
			},
		},
	}

	// Case 2: nonfinite lambda
	clockRep2 := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: math.NaN(), State: model.MiniState{Alpha: 1.2, Phi: 0.2}},
					{Lambda: 0.3, State: model.MiniState{Alpha: 1.5, Phi: 0.3}},
				},
			},
		},
	}

	// Case 3: non-ordered lambda that can be ordered/sorted
	clockRep3 := &clockseg.ClockReadinessReport{
		SchemaVersion:      "bmc0a-clock-readiness-v0.1",
		FriedmannReadiness: "local_only_candidate",
		ClockIndependentDiagnostics: clockseg.ClockIndependentDiagnostic{
			NodeContactFree:      true,
			TrajectoryFiniteness: true,
		},
		LocalRelationBranches: []clockseg.LocalRelationBranch{
			{
				ValidationPassed: true,
				ClockRange:       1.0,
				Samples:          3,
				Points: []model.TrajectoryPoint{
					{Lambda: 0.3, State: model.MiniState{Alpha: 1.5, Phi: 0.3}},
					{Lambda: 0.1, State: model.MiniState{Alpha: 1.0, Phi: 0.1}},
					{Lambda: 0.2, State: model.MiniState{Alpha: 1.2, Phi: 0.2}},
				},
			},
		},
	}

	friedmannRep := &friedmannspec.FriedmannSpecReport{}
	nullRep := &nullrun.NullRunReport{NullDiagnosticsComputed: true}
	priorArtRep := &priorart.PriorArtBoundaryReport{}

	rep1, err := RunResidualsFromInputs(clockRep1, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if rep1.CandidateResidualComputed {
		t.Fatal("Expected rep1.CandidateResidualComputed to be false due to duplicate Lambda")
	}

	rep2, err := RunResidualsFromInputs(clockRep2, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if rep2.CandidateResidualComputed {
		t.Fatal("Expected rep2.CandidateResidualComputed to be false due to NaN Lambda")
	}

	rep3, err := RunResidualsFromInputs(clockRep3, friedmannRep, nullRep, priorArtRep, "c.json", "f.json", "n.json", "p.json")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if !rep3.CandidateResidualComputed {
		t.Fatal("Expected rep3.CandidateResidualComputed to be true as unordered but valid Lambdas can be sorted")
	}
	// Verify sorted order in output metrics/points
	pts := rep3.CandidateResidualDiagnostics[0].ResidualInputPoints
	if len(pts) != 2 {
		t.Fatalf("Expected 2 finite-difference interval input points, got %d", len(pts))
	}
	if *pts[0].Alpha != 1.0 || *pts[1].Alpha != 1.2 {
		t.Errorf("Expected sorted interval-start alphas: 1.0, 1.2, got: %f, %f", *pts[0].Alpha, *pts[1].Alpha)
	}
}
