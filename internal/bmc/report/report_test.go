package report_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
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

