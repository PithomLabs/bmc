package report_test

import (
	"encoding/json"
	"math"
	"os"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

func TestReportDeterministicJSON(t *testing.T) {
	params := model.DefaultPlaneWaveParams()

	rep1, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report 1: %v", err)
	}

	rep2, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report 2: %v", err)
	}

	json1, err := json.Marshal(rep1)
	if err != nil {
		t.Fatalf("Error marshaling report 1: %v", err)
	}

	json2, err := json.Marshal(rep2)
	if err != nil {
		t.Fatalf("Error marshaling report 2: %v", err)
	}

	if string(json1) != string(json2) {
		t.Error("Generated JSON reports are not identical")
	}
}

func TestNoFinalTruthClaimAllowed(t *testing.T) {
	params := model.DefaultPlaneWaveParams()

	// Assert final truth claim (should fail validation)
	rep, err := report.Generate(params, true)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	errors := report.Validate(rep)
	if len(errors) == 0 {
		t.Fatal("Expected validation to fail for final_truth_claim = true, but no errors were returned")
	}

	foundFinalTruthError := false
	for _, valErr := range errors {
		if valErr.Field == "final_truth_claim" {
			foundFinalTruthError = true
			if valErr.Severity != report.ValidationFail {
				t.Errorf("Expected severity 'fail', got '%s'", valErr.Severity)
			}
		}
	}

	if !foundFinalTruthError {
		t.Error("Expected validation error on field 'final_truth_claim', but none was found")
	}
}

func TestToyAnalysisOnlyRejected(t *testing.T) {
	params := model.DefaultPlaneWaveParams()

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	// Manually violate the EBP toy_analysis_only constraint
	rep.ToyAnalysisOnly = false

	errors := report.Validate(rep)
	if len(errors) == 0 {
		t.Fatal("Expected validation to fail for toy_analysis_only = false, but no errors were returned")
	}

	foundToyError := false
	for _, valErr := range errors {
		if valErr.Field == "toy_analysis_only" {
			foundToyError = true
			if valErr.Severity != report.ValidationFail {
				t.Errorf("Expected severity 'fail', got '%s'", valErr.Severity)
			}
		}
	}

	if !foundToyError {
		t.Error("Expected validation error on field 'toy_analysis_only', but none was found")
	}
}

func TestWriteJSONDeterministic(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	tmpFile := t.TempDir() + "/report.json"
	err = report.WriteJSON(rep, tmpFile)
	if err != nil {
		t.Fatalf("Error writing JSON to temp file: %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("Error reading temp file: %v", err)
	}

	expectedData, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		t.Fatalf("Error marshaling: %v", err)
	}
	expectedData = append(expectedData, '\n')

	if string(data) != string(expectedData) {
		t.Error("On-disk JSON contents do not match expected marshaled bytes")
	}
}

func TestValidateKeepsToyOnlyStatus(t *testing.T) {
	params := model.DefaultPlaneWaveParams()

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	errors := report.Validate(rep)
	if len(errors) > 0 {
		t.Fatalf("Expected validation to pass, but got errors: %v", errors)
	}

	if !rep.ToyAnalysisOnly {
		t.Error("Expected ToyAnalysisOnly to be true")
	}
	if rep.FinalTruthClaim {
		t.Error("Expected FinalTruthClaim to be false")
	}
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical gate to pass, got %s", rep.TechnicalGate.Status)
	}
	if rep.PromotionGate.Status != report.StatusBlocked {
		t.Errorf("Expected promotion gate to be blocked, got %s", rep.PromotionGate.Status)
	}
}

func TestReportRejectsConstraintViolationStatus(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	// Manually set WdW residual check to fail while keeping technical gate as Pass
	wdwCheck := rep.Checks["wdw_residual"]
	wdwCheck.Status = model.StatusFail
	wdwCheck.Pass = false
	rep.Checks["wdw_residual"] = wdwCheck

	errors := report.Validate(rep)
	if len(errors) == 0 {
		t.Error("Expected report validation to fail due to inconsistent failed check and passing technical gate, but it passed")
	}
}

func TestPromotionGateBlocksConstraintViolation(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Error generating report: %v", err)
	}

	// Manually set WdW residual check to fail and set technical gate status to Fail
	wdwCheck := rep.Checks["wdw_residual"]
	wdwCheck.Status = model.StatusFail
	wdwCheck.Pass = false
	rep.Checks["wdw_residual"] = wdwCheck
	rep.TechnicalGate.Status = model.StatusFail

	// If the technical gate status has failed, promotion is blocked.
	// We assert that the validation behaves correctly (passes structural consistency checks, but remains blocked).
	errors := report.Validate(rep)
	for _, valErr := range errors {
		if valErr.Severity == report.ValidationFail {
			t.Errorf("Unexpected validation failure: %s", valErr.Message)
		}
	}

	if rep.TechnicalGate.Status != model.StatusFail {
		t.Errorf("Expected technical gate to be failed, got %s", rep.TechnicalGate.Status)
	}
	if rep.PromotionGate.Status != report.StatusBlocked {
		t.Errorf("Expected promotion gate to remain blocked, got %s", rep.PromotionGate.Status)
	}
}

func TestReportNumericalWDWResidualAcceptsConstraintShellPlaneWave(t *testing.T) {
	// Valid plane wave: k² = ω² = 4.0
	params := model.DefaultPlaneWaveParams()
	params.K = 2.0
	params.Omega = 2.0

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Unexpected error generating report: %v", err)
	}

	wdwCheck := rep.Checks["wdw_residual"]
	if wdwCheck.Status != model.StatusPass || !wdwCheck.Pass {
		t.Errorf("Expected WdW residual check to pass, got Status=%s, Pass=%t", wdwCheck.Status, wdwCheck.Pass)
	}

	if wdwCheck.NumericalResidualStatus == nil || *wdwCheck.NumericalResidualStatus != wdw.NumericalResidualPass {
		t.Errorf("Expected numerical residual status to be '%s', got %v", wdw.NumericalResidualPass, wdwCheck.NumericalResidualStatus)
	}

	if wdwCheck.NumericalResidualAuthority == nil || *wdwCheck.NumericalResidualAuthority != wdw.NumericalAuthorityDiagnostic {
		t.Errorf("Expected authority to be '%s', got %v", wdw.NumericalAuthorityDiagnostic, wdwCheck.NumericalResidualAuthority)
	}

	if wdwCheck.NumericalResidualMagnitude == nil || *wdwCheck.NumericalResidualMagnitude > wdw.WDWNumericalResidualTolerance {
		t.Errorf("Expected numerical residual magnitude to satisfy tolerance, got %v", wdwCheck.NumericalResidualMagnitude)
	}

	errors := report.Validate(rep)
	for _, valErr := range errors {
		if valErr.Severity == report.ValidationFail {
			t.Errorf("Unexpected validation error: %s", valErr.Message)
		}
	}
}

func TestReportNumericalWDWResidualBlocksWrongPlaneWaveConstraint(t *testing.T) {
	// Violated plane wave: k² = 4.0, ω² = 9.0 (k² != ω²)
	params := model.DefaultPlaneWaveParams()
	params.K = 2.0
	params.Omega = 3.0
	params.Tolerance = 10.0 // Allow parameter validation to pass

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Unexpected error generating report: %v", err)
	}

	wdwCheck := rep.Checks["wdw_residual"]
	if wdwCheck.Status != model.StatusFail || wdwCheck.Pass {
		t.Errorf("Expected WdW residual check to fail, got Status=%s, Pass=%t", wdwCheck.Status, wdwCheck.Pass)
	}

	if wdwCheck.NumericalResidualStatus == nil || *wdwCheck.NumericalResidualStatus != wdw.NumericalResidualViolationDetected {
		t.Errorf("Expected numerical residual status to be '%s', got %v", wdw.NumericalResidualViolationDetected, wdwCheck.NumericalResidualStatus)
	}

	if wdwCheck.NumericalResidualMagnitude == nil || *wdwCheck.NumericalResidualMagnitude <= wdw.WDWNumericalResidualTolerance {
		t.Errorf("Expected numerical residual magnitude to exceed tolerance, got %v", wdwCheck.NumericalResidualMagnitude)
	}

	if rep.TechnicalGate.Status != model.StatusFail {
		t.Errorf("Expected Technical Gate status to be 'fail', got '%s'", rep.TechnicalGate.Status)
	}

	// Structural validation should pass since report.Generate consistently sets TechnicalGate.Status to fail
	errors := report.Validate(rep)
	for _, valErr := range errors {
		if valErr.Severity == report.ValidationFail {
			t.Errorf("Unexpected validation failure: %s", valErr.Message)
		}
	}
}

func TestReportValidationRejectsNumericalResidualViolationClaimedAsPass(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Unexpected error generating report: %v", err)
	}

	// Mutate report: claim pass while numerical residual violation status is present
	wdwCheck := rep.Checks["wdw_residual"]
	wdwCheck.Status = model.StatusPass
	wdwCheck.Pass = true
	numStatus := wdw.NumericalResidualViolationDetected
	wdwCheck.NumericalResidualStatus = &numStatus
	rep.Checks["wdw_residual"] = wdwCheck

	errors := report.Validate(rep)
	foundMismatch := false
	for _, valErr := range errors {
		if valErr.Field == "checks.wdw_residual" {
			foundMismatch = true
		}
	}

	if !foundMismatch {
		t.Error("Expected report validation to reject report when pass is claimed with a numerical residual violation detected")
	}
}

func TestPromotionGateBlocksNumericalWDWViolation(t *testing.T) {
	// Wrong plane wave parameters
	params := model.DefaultPlaneWaveParams()
	params.K = 2.0
	params.Omega = 5.0
	params.Tolerance = 30.0 // Allow parameter validation to pass

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Unexpected error generating report: %v", err)
	}

	if rep.Checks["wdw_residual"].Status != model.StatusFail {
		t.Errorf("Expected failed status, got %s", rep.Checks["wdw_residual"].Status)
	}

	if rep.PromotionGate.Status != report.StatusBlocked {
		t.Errorf("Expected promotion gate to remain blocked, got %s", rep.PromotionGate.Status)
	}
}

func TestAnalyticResidualRemainsOracleControlOnly(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	params.K = 2.0
	params.Omega = 3.0 // k² - ω² = -5.0
	params.Tolerance = 10.0 // Allow parameter validation to pass

	rep, err := report.Generate(params, false)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	wdwCheck := rep.Checks["wdw_residual"]
	if wdwCheck.AnalyticResidualMagnitude == nil || math.Abs(*wdwCheck.AnalyticResidualMagnitude-5.0) > 1e-9 {
		t.Errorf("Expected analytic residual magnitude to be oracle control of 5.0, got %v", wdwCheck.AnalyticResidualMagnitude)
	}

	// Ensure numerical residual acts as diagnostic authority
	if wdwCheck.NumericalResidualAuthority == nil || *wdwCheck.NumericalResidualAuthority != wdw.NumericalAuthorityDiagnostic {
		t.Errorf("Expected numerical authority to be diagnostic authority, got %v", wdwCheck.NumericalResidualAuthority)
	}
}



