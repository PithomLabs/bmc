package nullrun

import (
	"encoding/json"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func helperDefaultReport(t *testing.T) *NullRunReport {
	t.Helper()
	return GenerateDefaultReport()
}

func TestNullRunReportValidation(t *testing.T) {
	rep := helperDefaultReport(t)
	errs := ValidateReport(rep, "{}")
	if len(errs) > 0 {
		t.Fatalf("Default report failed validation: %v", errs)
	}
}

func TestNullRunRejectsResidualComputed(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ResidualComputed = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ResidualComputed is true")
	}
}

func TestNullRunRejectsMissingNullDiagnostics(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullDiagnosticsComputed = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when NullDiagnosticsComputed is false")
	}
}

func TestNullRunRejectsMissingTargetNullComparison(t *testing.T) {
	// If Comparisons exists but computed is false
	rep := helperDefaultReport(t)
	rep.TargetNullComparisonComputed = false
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparisons exist but TargetNullComparisonComputed is false")
	}

	// If Comparisons is empty but computed is true
	rep2 := helperDefaultReport(t)
	rep2.TargetNullDiagnosticComparisons = nil
	rep2.TargetNullComparisonComputed = true
	errs2 := ValidateReport(rep2, "{}")
	if len(errs2) == 0 {
		t.Fatal("Expected failure when comparisons is empty but TargetNullComparisonComputed is true")
	}

	// Check the special revision rule:
	// "If no comparable null diagnostics exist, target_null_comparison_computed must be false and the report status must be blocked_by_no_comparable_null_diagnostics."
	rep3 := helperDefaultReport(t)
	rep3.TargetNullDiagnosticComparisons = nil
	rep3.TargetNullComparisonComputed = false
	rep3.InterpretationStatus = InterpretBlockedByNoComparableDiagnostics
	errs3 := ValidateReport(rep3, "{}")
	if len(errs3) > 0 {
		t.Fatalf("Expected empty comparisons to pass validation under the special revision rule, got: %v", errs3)
	}
}

func TestNullRunRejectsRecoveryClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.RecoveryClaim = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when RecoveryClaim is true")
	}
}

func TestNullRunRejectsScientificNoveltyClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ScientificNoveltyClaimMade = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ScientificNoveltyClaimMade is true")
	}
}

func TestNullRunRequiresAllSevenNullModels(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns = rep.NullModelRuns[1:]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a required null model is missing")
	}
}

func TestNullRunRejectsDuplicateNullModelID(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns = append(rep.NullModelRuns, rep.NullModelRuns[0])
	errs := ValidateReport(rep, "{}")
	foundDuplicate := false
	for _, e := range errs {
		if strings.Contains(e.Message, "duplicate null_model_id") || strings.Contains(e.Message, "duplicate null_model_id detected") {
			foundDuplicate = true
		}
	}
	if !foundDuplicate {
		t.Fatal("Expected duplicate null_model_id validation error")
	}
}

func TestNullRunRejectsUnknownNullModelID(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns[0].NullModelID = "unknown_model_xyz"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure for unknown null model ID")
	}
}

func TestNullRunRejectsForbiddenRunStatus(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns[0].RunStatus = "passed"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure for forbidden run status 'passed'")
	}
}

func TestNullRunRejectsForbiddenInterpretationStatus(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.InterpretationStatus = "winner"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure for forbidden interpretation status 'winner'")
	}
}

func TestNullRunRejectsSentinelNumericMetrics(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns[0].Diagnostics.NumValidTrajectoryPoints = -1
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when numeric metrics are sentinel negative values")
	}
}

func TestNullRunRejectsNonfiniteNumericMetrics(t *testing.T) {
	rep := helperDefaultReport(t)
	infVal := math.Inf(1)
	rep.NullModelRuns[0].Diagnostics.MinAmplitudeR = &infVal
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when float metrics are infinite")
	}
}

func TestNullRunRequiresAllGatesExactlyOnce(t *testing.T) {
	rep := helperDefaultReport(t)
	gatesBackup := rep.Gates
	rep.Gates = rep.Gates[1:]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when gate is missing")
	}

	rep.Gates = append(gatesBackup, gatesBackup[0])
	errs = ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when gate is duplicated")
	}
}

func TestNullRunRejectsNonPassGate(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.Gates[0].Status = "fail"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate status is not pass")
	}
}

func TestNullRunRejectsUnknownFields(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_unknown.json")
	raw := `{"schema_version":"bmc0a-nullrun-v0.1","extra_unknown_field":"some_val"}`
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with unknown fields")
	}
}

func TestNullRunRejectsTrailingJSONTokens(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_trailing.json")

	rep := helperDefaultReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatal(err)
	}
	raw := string(data) + "   true" // trailing token
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err = ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with trailing tokens")
	}
}

func TestNullRunDeterministicJSON(t *testing.T) {
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

func TestNullRunForbiddenPhraseScan(t *testing.T) {
	// Mixed case phrase: "wInNeR"
	phraseMixed := "wInNe" + "R"
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This is "+phraseMixed)
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected forbidden phrase scan to catch mixed case forbidden phrase")
	}
}

func TestNullRunForbiddenPhraseErrorsArePhraseSafe(t *testing.T) {
	phraseMixed := "wInNe" + "R"
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This is "+phraseMixed)
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected failure")
	}
	for _, e := range errs {
		if strings.Contains(strings.ToLower(e.Message), "winner") {
			t.Fatalf("Error message leaked forbidden phrase: %s", e.Message)
		}
	}
}

func TestNullRunCLIRoutingRunsGenerateValidateSummarize(t *testing.T) {
	// Compile test binary
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test")

	// 1. Generate report
	genCmd := exec.Command("./ptw-bmc-test", "run-nullmodels", "--profile", "bmc0a-nullrun", "--out", "out_test.json")
	genCmd.Dir = "."
	if out, err := genCmd.CombinedOutput(); err != nil {
		t.Fatalf("Generation failed: %v, output: %s", err, string(out))
	}
	defer os.Remove("./out_test.json")

	// 2. Validate
	valCmd := exec.Command("./ptw-bmc-test", "validate", "--report", "out_test.json")
	valCmd.Dir = "."
	valOut, err := valCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Validation failed: %v, output: %s", err, string(valOut))
	}
	if !strings.Contains(string(valOut), "PASSED") {
		t.Fatalf("Unexpected validation output: %s", string(valOut))
	}

	// 3. Summarize
	sumCmd := exec.Command("./ptw-bmc-test", "summarize", "--report", "out_test.json")
	sumCmd.Dir = "."
	sumOut, err := sumCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Summarization failed: %v, output: %s", err, string(sumOut))
	}
	if !strings.Contains(string(sumOut), "Null Models Registered: 7") {
		t.Fatalf("Unexpected summary output: %s", string(sumOut))
	}
}

func TestNullRunUnknownProfileFailsAtCLI(t *testing.T) {
	// Compile test binary
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test2", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test2")

	// Run with unknown profile
	failCmd := exec.Command("./ptw-bmc-test2", "run-nullmodels", "--profile", "unknown-profile", "--out", "out_test2.json")
	failCmd.Dir = "."
	out, err := failCmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected CLI to exit with failure for unknown profile")
	}
	if !strings.Contains(string(out), "is not supported") {
		t.Fatalf("Unexpected CLI error message: %s", string(out))
	}
}

func TestNullRunRejectsComparisonComputedWhenNoComparableNullDiagnostics(t *testing.T) {
	rep := helperDefaultReport(t)
	// Mark all null model runs blocked or deferred
	for i := range rep.NullModelRuns {
		if rep.NullModelRuns[i].RunStatus == RunStatusDiagnosticsGenerated {
			rep.NullModelRuns[i].RunStatus = RunStatusBlocked
			rep.NullModelRuns[i].DiagnosticProvenance = ProvenanceBlocked
			rep.NullModelRuns[i].DiagnosticStatus = DiagStatusNodeBlocked
		}
	}
	// If target_null_comparison_computed is true, it should fail
	rep.TargetNullComparisonComputed = true
	rep.InterpretationStatus = InterpretDiagComparisonOnly
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation failure when no comparable null diagnostics exist but target_null_comparison_computed is true")
	}

	// If target_null_comparison_computed is false but interpretation status is not blocked_by_no_comparable_null_diagnostics, it should fail
	rep.TargetNullComparisonComputed = false
	rep.InterpretationStatus = InterpretDiagComparisonOnly
	errs = ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation failure when no comparable null diagnostics exist but interpretation_status is not blocked")
	}

	// If target_null_comparison_computed is false and interpretation status is blocked_by_no_comparable_null_diagnostics, it should pass (if comparisons is empty)
	rep.InterpretationStatus = InterpretBlockedByNoComparableDiagnostics
	rep.TargetNullDiagnosticComparisons = nil
	errs = ValidateReport(rep, "{}")
	if len(errs) > 0 {
		t.Fatalf("Expected validation to pass under special blocked rule when no comparable diagnostics exist, got: %v", errs)
	}
}

func TestNullRunClassicalReferenceTrajectoryConstraints(t *testing.T) {
	// 1. Invalid status
	rep1 := helperDefaultReport(t)
	for i := range rep1.NullModelRuns {
		if rep1.NullModelRuns[i].NullModelID == "classical_frw_reference_trajectory" {
			rep1.NullModelRuns[i].RunStatus = "invalid_status"
		}
	}
	if len(ValidateReport(rep1, "{}")) == 0 {
		t.Fatal("Expected failure for invalid status of classical_frw_reference_trajectory")
	}

	// 2. Invalid provenance
	rep2 := helperDefaultReport(t)
	for i := range rep2.NullModelRuns {
		if rep2.NullModelRuns[i].NullModelID == "classical_frw_reference_trajectory" {
			rep2.NullModelRuns[i].DiagnosticProvenance = ProvenanceDeferred
		}
	}
	if len(ValidateReport(rep2, "{}")) == 0 {
		t.Fatal("Expected failure for invalid provenance of classical_frw_reference_trajectory")
	}

	// 3. Invalid notes
	rep3 := helperDefaultReport(t)
	for i := range rep3.NullModelRuns {
		if rep3.NullModelRuns[i].NullModelID == "classical_frw_reference_trajectory" {
			rep3.NullModelRuns[i].Notes = "some other notes"
		}
	}
	if len(ValidateReport(rep3, "{}")) == 0 {
		t.Fatal("Expected failure for invalid notes of classical_frw_reference_trajectory")
	}
}

func TestNullRunRejectsComparisonWithNoMetrics(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.TargetNullDiagnosticComparisons[0].MetricsCompared = []string{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when metrics_compared is empty")
	}
}

func TestNullRunRejectsComparisonWithNoNullModelIDs(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.TargetNullDiagnosticComparisons[0].NullModelIDs = []string{}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when null_model_ids is empty")
	}
}

func TestNullRunRejectsComparisonReferencingBlockedNullRun(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.TargetNullDiagnosticComparisons[0].NullModelIDs = []string{"node_neighborhood_stress_case"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison references a blocked null model run")
	}
}

func TestNullRunRejectsComparisonReferencingDeferredNullRun(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.TargetNullDiagnosticComparisons[0].NullModelIDs = []string{"same_branch_segmentation_under_null_wavefunctions"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison references a deferred null model run")
	}
}

func TestNullRunRejectsComparisonReferencingUnknownNullRun(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.TargetNullDiagnosticComparisons[0].NullModelIDs = []string{"nonexistent_null_model"}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when comparison references an unknown null model run")
	}
}

func TestNullRunRejectsNegativeMinAmplitude(t *testing.T) {
	rep := helperDefaultReport(t)
	val := -0.5
	rep.NullModelRuns[0].Diagnostics.MinAmplitudeR = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when min_amplitude_r is negative")
	}
}

func TestNullRunRejectsNegativeMaxAbsQ(t *testing.T) {
	rep := helperDefaultReport(t)
	val := -10.0
	rep.NullModelRuns[0].Diagnostics.MaxAbsQAwayFromNodes = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when max_abs_q_away_from_nodes is negative")
	}
}

func TestNullRunRejectsNegativeMaxPhaseGradient(t *testing.T) {
	rep := helperDefaultReport(t)
	val := -1.0
	rep.NullModelRuns[0].Diagnostics.MaxPhaseGradient = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when max_phase_gradient is negative")
	}
}

func TestNullRunRejectsSentinelNegativeFloatMetrics(t *testing.T) {
	rep := helperDefaultReport(t)
	val := -1.0
	rep.NullModelRuns[0].Diagnostics.MinAmplitudeR = &val
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when float metric has a negative sentinel value of -1")
	}
}

func TestNullRunUnavailableMetricsEmitExplicitNull(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullModelRuns[0].Diagnostics.MinAmplitudeR = nil
	rep.NullModelRuns[0].Diagnostics.MaxAbsQAwayFromNodes = nil
	rep.NullModelRuns[0].Diagnostics.MaxPhaseGradient = nil

	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	raw := string(data)
	if !strings.Contains(raw, `"min_amplitude_r":null`) && !strings.Contains(raw, `"min_amplitude_r": null`) {
		t.Fatal("Expected min_amplitude_r to be serialized as null")
	}
}

func TestNullRunClassicalReferenceUnavailableMetricsEmitNull(t *testing.T) {
	rep := helperDefaultReport(t)
	var referenceRun NullModelRun
	found := false
	for _, r := range rep.NullModelRuns {
		if r.NullModelID == "classical_frw_reference_trajectory" {
			referenceRun = r
			found = true
			break
		}
	}
	if !found {
		t.Fatal("classical_frw_reference_trajectory run not found")
	}
	if referenceRun.Diagnostics.MinAmplitudeR != nil {
		t.Fatal("Expected classical reference MinAmplitudeR to be nil")
	}
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	raw := string(data)
	if !strings.Contains(raw, `"min_amplitude_r":null`) && !strings.Contains(raw, `"min_amplitude_r": null`) {
		t.Fatal("Expected min_amplitude_r of classical reference to be null in JSON")
	}
}

func TestNullRunRejectsMissingOptionalMetricKeys(t *testing.T) {
	rep := helperDefaultReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatalf("Failed to marshal: %v", err)
	}
	raw := strings.Replace(string(data), `"min_amplitude_r"`, `"missing_key"`, -1)
	errs := ValidateReport(rep, raw)
	foundMissing := false
	for _, e := range errs {
		if strings.Contains(e.Message, "missing optional metric key") {
			foundMissing = true
		}
	}
	if !foundMissing {
		t.Fatal("Expected validation to fail when raw JSON is missing min_amplitude_r key")
	}
}

func TestNullRunRejectsOutOfVocabularyDebtStatus(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.EbpDebt.NeedLiteratureAudit = "invalid_status"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail for out-of-vocabulary debt status")
	}
}

func TestNullRunAcceptsReviewDebtVocabulary(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.EbpDebtVocabulary = "ptw_adversarial_review_debt_status_v0.1"
	errs := ValidateReport(rep, "{}")
	if len(errs) > 0 {
		t.Fatalf("Expected validation to pass for review debt vocabulary, got errors: %v", errs)
	}
}

func TestNullRunRequiresCandidateOnlyPromotionStatus(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.EbpDebt.PromotionStatus = "promoted"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when promotion_status is not candidate only")
	}
}

func TestNullRunAccountsForAllSevenNullModelsInSummary(t *testing.T) {
	rep := helperDefaultReport(t)
	totalRuns := len(rep.NullModelRuns)
	if totalRuns != 7 {
		t.Fatalf("Expected exactly 7 registered null models, got %d", totalRuns)
	}
}

func TestNullRunSummaryIncludesDeferredCount(t *testing.T) {
	rep := helperDefaultReport(t)
	numDeferred := 0
	for _, run := range rep.NullModelRuns {
		if run.RunStatus == "deferred" {
			numDeferred++
		}
	}
	if numDeferred != 2 {
		t.Fatalf("Expected exactly 2 deferred null models, got %d", numDeferred)
	}
}

func TestNullRunSummaryIncludesAccountedForCount(t *testing.T) {
	rep := helperDefaultReport(t)
	numWithDiag := 0
	numBlocked := 0
	numDeferred := 0
	for _, run := range rep.NullModelRuns {
		if run.RunStatus == "diagnostics_generated" {
			numWithDiag++
		} else if run.RunStatus == "blocked" {
			numBlocked++
		} else if run.RunStatus == "deferred" {
			numDeferred++
		}
	}
	totalAccounted := numWithDiag + numBlocked + numDeferred
	if totalAccounted != len(rep.NullModelRuns) {
		t.Fatalf("Opaque accounting count mismatch: sum is %d, total is %d", totalAccounted, len(rep.NullModelRuns))
	}
}


