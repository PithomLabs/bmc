package nullspec

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// 1. TestNullModelSpecReportValidation
func TestNullModelSpecReportValidation(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, err := GenerateNullModelSpecReport(params)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	errs := ValidateNullModelSpecReport(r)
	if len(errs) > 0 {
		t.Errorf("expected validation to pass, got %d errors: %v", len(errs), errs)
	}
}

// 2. TestNullModelSpecRejectsResidualComputed
func TestNullModelSpecRejectsResidualComputed(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.ResidualComputed = true
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when residual_computed is true")
	}
}

// 3. TestNullModelSpecRejectsNullComparisonComputed
func TestNullModelSpecRejectsNullComparisonComputed(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.NullComparisonComputed = true
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when null_comparison_computed is true")
	}
}

// 4. TestNullModelSpecRejectsRecoveryClaim
func TestNullModelSpecRejectsRecoveryClaim(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.FriedmannRecoveryClaim = true
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when friedmann_recovery_claim is true")
	}
}

// 5. TestNullModelSpecRequiresAllRequiredNullModels
func TestNullModelSpecRequiresAllRequiredNullModels(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	// Remove one required null model
	r.NullModels = r.NullModels[1:]
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when a required null model is missing")
	}
}

// 6. TestNullModelSpecRejectsDuplicateNullModelIDs
func TestNullModelSpecRejectsDuplicateNullModelIDs(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	// Add a duplicate null model
	r.NullModels = append(r.NullModels, r.NullModels[0])
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when duplicate null models exist")
	}
}

// 7. TestNullModelSpecRequiresBeforeResidualPromotion
func TestNullModelSpecRequiresBeforeResidualPromotion(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.NullModels[0].RequiredBeforeResidualPromotion = false
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when required_before_residual_promotion is false")
	}
}

// 8. TestNullModelSpecRejectsPassedFailedNullModelStatus
func TestNullModelSpecRejectsPassedFailedNullModelStatus(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.NullModels[0].Status = "passed"
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when null model status is passed")
	}
}

// 9. TestNullModelSpecRequiresMetricContracts
func TestNullModelSpecRequiresMetricContracts(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.MetricContracts = nil
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when metric contracts are empty")
	}
}

// 10. TestNullModelSpecRejectsComputedFutureComparison
func TestNullModelSpecRejectsComputedFutureComparison(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.FutureComparisonContracts[0].ComparisonComputed = true
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when comparison_computed is true")
	}
}

// 11. TestNullModelSpecRequiresNoNullComparisonResultGate
func TestNullModelSpecRequiresNoNullComparisonResultGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	// Remove no_null_comparison_result_gate
	var newGates []NullSpecGate
	for _, g := range r.Gates {
		if g.Name != "no_null_comparison_result_gate" {
			newGates = append(newGates, g)
		}
	}
	r.Gates = newGates

	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when no_null_comparison_result_gate is missing")
	}
}

// 12. TestNullModelSpecRequiresFullBMCBlockedGate
func TestNullModelSpecRequiresFullBMCBlockedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	// Remove full_bmc_blocked_gate
	var newGates []NullSpecGate
	for _, g := range r.Gates {
		if g.Name != "full_bmc_blocked_gate" {
			newGates = append(newGates, g)
		}
	}
	r.Gates = newGates

	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when full_bmc_blocked_gate is missing")
	}
}

// 13. TestNullModelSpecRequiresNullModelDebtActive
func TestNullModelSpecRequiresNullModelDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.EbpDebt.NeedNullModel = "resolved"
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when needNullModel is resolved")
	}
}

// 14. TestNullModelSpecRejectsUnknownFields
func TestNullModelSpecRejectsUnknownFields(t *testing.T) {
	jsonData := `{"schema_version": "bmc0a-nullmodel-spec-v0.1", "extra_bad_field": true}`
	dec := json.NewDecoder(bytes.NewReader([]byte(jsonData)))
	dec.DisallowUnknownFields()

	var r NullModelSpecReport
	err := dec.Decode(&r)
	if err == nil {
		t.Error("expected strict unmarshal to reject unknown fields")
	}
}

// 15. TestNullModelSpecRejectsTrailingJSONTokens
func TestNullModelSpecRejectsTrailingJSONTokens(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)
	b, _ := json.Marshal(r)

	tempDir := t.TempDir()

	// Write valid JSON
	tmpFile := filepath.Join(tempDir, "tmp_test_null_valid.json")
	_ = os.WriteFile(tmpFile, b, 0644)

	_, err := ReadNullModelSpecReport(tmpFile)
	if err != nil {
		t.Fatalf("expected ReadNullModelSpecReport to succeed, got: %v", err)
	}

	// Trailing garbage JSON
	garbageFile := filepath.Join(tempDir, "tmp_test_null_garbage.json")
	garbageBytes := append(b, []byte(" }")...)
	_ = os.WriteFile(garbageFile, garbageBytes, 0644)

	_, err = ReadNullModelSpecReport(garbageFile)
	if err == nil {
		t.Error("expected ReadNullModelSpecReport to fail with trailing garbage")
	} else if !strings.Contains(err.Error(), "trailing garbage") {
		t.Errorf("expected error message to contain 'trailing garbage', got: %v", err)
	}
}

// 16. TestNullModelSpecDeterministicJSON
func TestNullModelSpecDeterministicJSON(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r1, _ := GenerateNullModelSpecReport(params)
	r2, _ := GenerateNullModelSpecReport(params)

	b1, _ := json.MarshalIndent(r1, "", "  ")
	b2, _ := json.MarshalIndent(r2, "", "  ")

	if !bytes.Equal(b1, b2) {
		t.Error("expected deterministic JSON outputs to be byte-identical")
	}
}

// 17. TestNullModelSpecCLIRouting
func TestNullModelSpecCLIRouting(t *testing.T) {
	tempDir := t.TempDir()
	binPath := filepath.Join(tempDir, "ptw-bmc")

	// 1. Build the CLI binary
	cmdBuild := exec.Command("go", "build", "-buildvcs=false", "-o", binPath, "../../../cmd/ptw-bmc/main.go")
	if out, err := cmdBuild.CombinedOutput(); err != nil {
		t.Fatalf("failed to build ptw-bmc CLI: %v\nOutput: %s", err, string(out))
	}

	// 2. Generate report
	reportPath := filepath.Join(tempDir, "bmc0a_nullmodel_spec.json")
	cmdGen := exec.Command(binPath, "spec-nullmodels", "--profile", "bmc0a-nullmodel-spec", "--out", reportPath)
	if out, err := cmdGen.CombinedOutput(); err != nil {
		t.Fatalf("failed to run spec-nullmodels: %v\nOutput: %s", err, string(out))
	}

	// Verify report file was written
	if _, err := os.Stat(reportPath); os.IsNotExist(err) {
		t.Fatal("expected report file to be written, but not found")
	}

	// 3. Test unknown profile failure
	cmdGenBad := exec.Command(binPath, "spec-nullmodels", "--profile", "unknown-profile", "--out", reportPath)
	if err := cmdGenBad.Run(); err == nil {
		t.Error("expected spec-nullmodels to fail for unknown profile, but it exited with 0")
	}

	// 4. Validate report (this verifies actual validate command routing compiles and dispatches correctly)
	cmdVal := exec.Command(binPath, "validate", "--report", reportPath)
	if out, err := cmdVal.CombinedOutput(); err != nil {
		t.Fatalf("failed to validate report: %v\nOutput: %s", err, string(out))
	} else if !strings.Contains(string(out), "PASSED") {
		t.Errorf("expected validate command output to contain 'PASSED', got: %s", string(out))
	}

	// 5. Summarize report (this verifies actual summarize command routing compiles and dispatches correctly)
	cmdSum := exec.Command(binPath, "summarize", "--report", reportPath)
	if out, err := cmdSum.CombinedOutput(); err != nil {
		t.Fatalf("failed to summarize report: %v\nOutput: %s", err, string(out))
	} else if !strings.Contains(string(out), "Null-Model Spec") {
		t.Errorf("expected summarize output to contain summary title, got: %s", string(out))
	}
}

// 18. TestNullModelSpecForbiddenWords
func TestNullModelSpecForbiddenWords(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	forbiddenPhrases := []string{
		"null models passed",
		"BMC beats null models",
		"Friedmann recovery",
		"ready for Friedmann recovery",
		"classical limit verified",
	}

	// Scan full JSON serialization of the report
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("failed to marshal report: %v", err)
	}
	reportStr := strings.ToLower(string(b))
	for _, p := range forbiddenPhrases {
		if strings.Contains(reportStr, strings.ToLower(p)) {
			t.Errorf("found forbidden phrase in generated report JSON: %q", p)
		}
	}

	// Scan the printed summary output
	oldStdout := os.Stdout
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = wPipe

	SummarizeNullModelSpecReport(r)

	wPipe.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, rPipe)
	summaryStr := strings.ToLower(buf.String())
	for _, p := range forbiddenPhrases {
		if strings.Contains(summaryStr, strings.ToLower(p)) {
			t.Errorf("found forbidden phrase in summary output: %q", p)
		}
	}
}

// 19. TestNullModelSpecRequiresAllGatesExactlyOnce
func TestNullModelSpecRequiresAllGatesExactlyOnce(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	// Duplicate a gate
	r.Gates = append(r.Gates, r.Gates[0])
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when duplicate gates exist")
	}
}

// 20. TestNullModelSpecRejectsNonPassNoResidualComputationGate
func TestNullModelSpecRejectsNonPassNoResidualComputationGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "no_residual_computation_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when no_residual_computation_gate is not pass")
	}
}

// 21. TestNullModelSpecRejectsNonPassToyAnalysisOnlyGate
func TestNullModelSpecRejectsNonPassToyAnalysisOnlyGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "toy_analysis_only_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when toy_analysis_only_gate is not pass")
	}
}

// 22. TestNullModelSpecRejectsNonPassNoFinalTruthClaimGate
func TestNullModelSpecRejectsNonPassNoFinalTruthClaimGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "no_final_truth_claim_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when no_final_truth_claim_gate is not pass")
	}
}

// 23. TestNullModelSpecRejectsNonPassRecoveryClaimBlockedGate
func TestNullModelSpecRejectsNonPassRecoveryClaimBlockedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "friedmann_recovery_claim_blocked_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when friedmann_recovery_claim_blocked_gate is not pass")
	}
}

// 24. TestNullModelSpecRejectsNonPassClockChoiceDebtGate
func TestNullModelSpecRejectsNonPassClockChoiceDebtGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "clock_choice_debt_active_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when clock_choice_debt_active_gate is not pass")
	}
}

// 25. TestNullModelSpecRejectsNonPassFaithfulnessContestedGate
func TestNullModelSpecRejectsNonPassFaithfulnessContestedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	for idx, g := range r.Gates {
		if g.Name == "faithfulness_contested_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when faithfulness_contested_gate is not pass")
	}
}

// 26. TestNullModelSpecRequiresMetricsBeforeResidualPromotion
func TestNullModelSpecRequiresMetricsBeforeResidualPromotion(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.MetricContracts[0].RequiredBeforeResidualPromotion = false
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when required_before_residual_promotion is false on a metric contract")
	}
}

// 27. TestNullModelSpecRejectsEmptyFutureComparisonContracts
func TestNullModelSpecRejectsEmptyFutureComparisonContracts(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.FutureComparisonContracts = nil
	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when future comparison contracts are empty")
	}
}

// 28. TestNullModelSpecRejectsSemanticBypassPhrases
func TestNullModelSpecRejectsSemanticBypassPhrases(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	forbiddenPhrases := []string{
		"recovery of Friedmann",
		"Friedmann support",
		"Friedmann-compatible",
		"control victory",
		"nulls defeated",
		"BMC beats null models",
	}

	b, _ := json.Marshal(r)
	reportStr := strings.ToLower(string(b))

	for _, p := range forbiddenPhrases {
		if strings.Contains(reportStr, strings.ToLower(p)) {
			t.Errorf("found forbidden bypass phrase in generated report JSON: %q", p)
		}
	}
}

// 29. TestNullModelSpecRejectsUnknownExtraNullModelID
func TestNullModelSpecRejectsUnknownExtraNullModelID(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateNullModelSpecReport(params)

	r.NullModels = append(r.NullModels, NullModelSpec{
		NullModelID:                     "unknown_null_model",
		Name:                            "Unknown Spec ID",
		RequiredBeforeResidualPromotion: true,
		Status:                          StatusPlanned,
	})

	errs := ValidateNullModelSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when an unknown extra null model ID is registered")
	}
}
