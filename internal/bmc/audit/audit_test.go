package audit_test

import (
	"encoding/json"
	"math"
	"os"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/audit"
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

func TestRobustnessReportValidation(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	rep, err := audit.GenerateAuditReport(safeParams, nodeProbeParams)
	if err != nil {
		t.Fatalf("Error generating audit report: %v", err)
	}

	// 1. Validation should pass normally
	errors := audit.ValidateRobustnessReport(rep)
	if len(errors) > 0 {
		t.Fatalf("Expected validation to pass, but got errors: %v", errors)
	}

	// 2. Reject final truth claim
	rep.FinalTruthClaim = true
	errors = audit.ValidateRobustnessReport(rep)
	if len(errors) == 0 {
		t.Error("Expected validation to fail for final_truth_claim = true, but no errors were returned")
	}
	rep.FinalTruthClaim = false // reset

	// 3. Reject toy analysis = false
	rep.ToyAnalysisOnly = false
	errors = audit.ValidateRobustnessReport(rep)
	if len(errors) == 0 {
		t.Error("Expected validation to fail for toy_analysis_only = false, but no errors were returned")
	}
	rep.ToyAnalysisOnly = true // reset

	// 4. Strict decoding test: reject unknown fields
	tmpFile := t.TempDir() + "/robustness_invalid.json"
	jsonData, err := json.Marshal(rep)
	if err != nil {
		t.Fatalf("Error marshaling report: %v", err)
	}

	// Add an unknown field manually to the JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal(jsonData, &parsed); err != nil {
		t.Fatalf("Error unmarshaling to map: %v", err)
	}
	parsed["unknown_field_xyz"] = "invalid_data"

	invalidData, err := json.Marshal(parsed)
	if err != nil {
		t.Fatalf("Error marshaling invalid data: %v", err)
	}

	if err := os.WriteFile(tmpFile, invalidData, 0644); err != nil {
		t.Fatalf("Error writing file: %v", err)
	}

	_, err = audit.ReadRobustnessReport(tmpFile)
	if err == nil {
		t.Error("Expected ReadRobustnessReport to fail due to strict decoding of unknown field, but it succeeded")
	}
}

func TestStepSizeConvergence(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	results, err := audit.RunStepSizeSweep(params)
	if err != nil {
		t.Fatalf("Error running step-size sweep: %v", err)
	}

	if len(results) != 4 {
		t.Fatalf("Expected 4 results, got %d", len(results))
	}

	// Verify step sizes and deterministic execution
	expectedSizes := []float64{0.1, 0.05, 0.025, 0.0125}
	expectedSteps := []int{100, 200, 400, 800}

	for i, r := range results {
		if r.StepSize != expectedSizes[i] {
			t.Errorf("Expected step size %f, got %f at index %d", expectedSizes[i], r.StepSize, i)
		}
		if r.Steps != expectedSteps[i] {
			t.Errorf("Expected step count %d, got %d at index %d", expectedSteps[i], r.Steps, i)
		}

		// Verify drift metrics are finite
		if r.EndpointDriftAlpha == nil || math.IsNaN(*r.EndpointDriftAlpha) || math.IsInf(*r.EndpointDriftAlpha, 0) {
			t.Errorf("Endpoint drift alpha is non-finite at index %d", i)
		}
		if r.EndpointDriftPhi == nil || math.IsNaN(*r.EndpointDriftPhi) || math.IsInf(*r.EndpointDriftPhi, 0) {
			t.Errorf("Endpoint drift phi is non-finite at index %d", i)
		}

		// The finest step size (0.0125) must have exactly zero drift relative to itself
		if r.StepSize == 0.0125 {
			if r.EndpointDriftAlpha == nil || *r.EndpointDriftAlpha != 0.0 || r.EndpointDriftPhi == nil || *r.EndpointDriftPhi != 0.0 {
				t.Errorf("Finest step size drift should be exactly 0, got alpha_drift=%v, phi_drift=%v", r.EndpointDriftAlpha, r.EndpointDriftPhi)
			}
		}

		// Safe profile results should be stable (technical gate status pass)
		if r.TechnicalGateStatus != model.StatusPass {
			t.Errorf("Expected technical gate status 'pass' at step size %f, got '%s'", r.StepSize, r.TechnicalGateStatus)
		}
	}
}

func TestThresholdSensitivity(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	results, err := audit.RunThresholdSweep(safeParams, nodeProbeParams)
	if err != nil {
		t.Fatalf("Error running threshold sweep: %v", err)
	}

	if len(results) != 6 {
		t.Fatalf("Expected 6 results, got %d", len(results))
	}

	for _, r := range results {
		if r.Profile == "safe" {
			// Safe profile remains node-contact-free under all thresholds
			if r.NodeContactFree != model.StatusPass {
				t.Errorf("Expected safe profile to be node-contact-free under threshold %e, got status '%s'", r.Threshold, r.NodeContactFree)
			}
			if r.TechnicalGateName != "bmc0a_superposition_safe_gate" || r.TechnicalGateStatus != model.StatusPass {
				t.Errorf("Expected safe profile to pass safe gate under threshold %e, got gate '%s' with status '%s'", r.Threshold, r.TechnicalGateName, r.TechnicalGateStatus)
			}
		} else if r.Profile == "node-probe" {
			// Node-probe remains correctly classified as node-contact/blocked
			if r.NodeContactFree != model.StatusFail {
				t.Errorf("Expected node-probe profile to fail node-contact-free check under threshold %e, got status '%s'", r.Threshold, r.NodeContactFree)
			}
			// But obstruction detector itself functions correctly, so validation gate passes
			if r.TechnicalGateName != "node_detection_validation_gate" || r.TechnicalGateStatus != model.StatusPass {
				t.Errorf("Expected node-probe to pass validation gate under threshold %e, got gate '%s' with status '%s'", r.Threshold, r.TechnicalGateName, r.TechnicalGateStatus)
			}
		}
	}
}

func TestPhaseGradientSensitivity(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	results, err := audit.RunPhaseBoundSweep(params)
	if err != nil {
		t.Fatalf("Error running phase bound sweep: %v", err)
	}

	if len(results) != 4 {
		t.Fatalf("Expected 4 results, got %d", len(results))
	}

	// Bounds: 25, 50, 100, 200
	expectedBounds := []float64{25, 50, 100, 200}
	for i, r := range results {
		if r.Bound != expectedBounds[i] {
			t.Errorf("Expected bound %f, got %f", expectedBounds[i], r.Bound)
		}
		if r.MaxObservedPhaseGrad == nil || math.IsNaN(*r.MaxObservedPhaseGrad) || math.IsInf(*r.MaxObservedPhaseGrad, 0) {
			t.Errorf("Max observed phase gradient is non-finite under bound %f", r.Bound)
		}

		// Assert status is consistent with binding boolean
		if r.IsBinding && r.PhaseGradientFinite == model.StatusPass {
			t.Errorf("Expected phase_gradient_finite check to fail when binding is true under bound %f", r.Bound)
		}
		if !r.IsBinding && r.PhaseGradientFinite != model.StatusPass {
			t.Errorf("Expected phase_gradient_finite check to pass when binding is false under bound %f", r.Bound)
		}
	}
}

func TestParameterPerturbations(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	// 1. Parameter perturbation sweep
	pResults, err := audit.RunPerturbationSweep(safeParams)
	if err != nil {
		t.Fatalf("Error running parameter sweep: %v", err)
	}

	if len(pResults) != 9 {
		t.Fatalf("Expected 9 parameter sweep results, got %d", len(pResults))
	}

	// Verify nested-loop deterministic order
	expectedC2 := []float64{0.45, 0.45, 0.45, 0.50, 0.50, 0.50, 0.55, 0.55, 0.55}
	expectedK2 := []float64{1.9, 2.0, 2.1, 1.9, 2.0, 2.1, 1.9, 2.0, 2.1}

	for i, r := range pResults {
		if math.Abs(r.C2Real-expectedC2[i]) > 1e-9 {
			t.Errorf("Expected c2_real %f, got %f at index %d", expectedC2[i], r.C2Real, i)
		}
		if math.Abs(r.K2-expectedK2[i]) > 1e-9 {
			t.Errorf("Expected k2 %f, got %f at index %d", expectedK2[i], r.K2, i)
		}
		if math.Abs(r.Omega2-(-expectedK2[i])) > 1e-9 {
			t.Errorf("Expected omega2 %f, got %f at index %d", -expectedK2[i], r.Omega2, i)
		}
	}

	// 2. Node offset sweep (NodeThresh = 1e-5)
	oResults, err := audit.RunNodeProbeOffsetSweep(nodeProbeParams)
	if err != nil {
		t.Fatalf("Error running node offset sweep: %v", err)
	}

	if len(oResults) != 4 {
		t.Fatalf("Expected 4 offset results, got %d", len(oResults))
	}

	// (0,0): must short-circuit
	if !oResults[0].ShortCircuitTriggered || oResults[0].Integrated {
		t.Error("Expected exact node (0, 0) to short-circuit and not integrate")
	}
	if oResults[0].NodeContactFree != model.StatusFail {
		t.Errorf("Expected exact node (0, 0) node_contact_free status to be 'fail', got '%s'", oResults[0].NodeContactFree)
	}

	// (1e-8, 0): should short-circuit (below 1e-5 threshold)
	if !oResults[1].ShortCircuitTriggered || oResults[1].Integrated {
		t.Error("Expected offset (1e-8, 0) to short-circuit and not integrate")
	}

	// (1e-6, 0): should short-circuit (below 1e-5 threshold)
	if !oResults[2].ShortCircuitTriggered || oResults[2].Integrated {
		t.Error("Expected offset (1e-6, 0) to short-circuit and not integrate")
	}

	// (1e-4, 0): is above threshold, integrates, and reports if phase/Q risks appear
	if oResults[3].ShortCircuitTriggered || !oResults[3].Integrated {
		t.Error("Expected offset (1e-4, 0) to integrate and not short-circuit")
	}
}

func TestRobustnessOutcomeMixedWhenPerturbationFailuresOccur(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	rep, err := audit.GenerateAuditReport(safeParams, nodeProbeParams)
	if err != nil {
		t.Fatalf("Error generating audit report: %v", err)
	}

	// 1. Assert technical gate status is pass (audit ran successfully and is integral)
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical gate status 'pass' for audit integrity, got '%s'", rep.TechnicalGate.Status)
	}

	// 2. Assert that robustness_outcome is 'mixed' due to parameter perturbation failures
	hasFailures := false
	for _, p := range rep.ParameterPerturbationSweep {
		if p.SafeGateStatus != model.StatusPass || p.ClockMonotonic != model.StatusPass {
			hasFailures = true
			break
		}
	}

	if !hasFailures {
		t.Fatal("Expected some perturbation failures to occur in the default sweep parameters, but none occurred")
	}

	if rep.RobustnessOutcome != "mixed" {
		t.Errorf("Expected robustness_outcome to be 'mixed', got '%s'", rep.RobustnessOutcome)
	}
}

func TestDataIntegrityValidationRules(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	rep, err := audit.GenerateAuditReport(safeParams, nodeProbeParams)
	if err != nil {
		t.Fatalf("Error generating audit report: %v", err)
	}

	// Double check technical_gate status and robustness_outcome
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical_gate status 'pass' for audit integrity, got '%s'", rep.TechnicalGate.Status)
	}
	if rep.RobustnessOutcome != "mixed" {
		t.Errorf("Expected robustness_outcome to remain 'mixed', got '%s'", rep.RobustnessOutcome)
	}

	// Helper to clone a float64
	val := 1.5
	valNegOne := -1.0
	valNaN := math.NaN()

	// 1. Unavailable metric (nil/null) with status/reason: PASS
	rep.StepSizeSweep[0].MaxAbsQ = nil
	rep.StepSizeSweep[0].MaxAbsQStatus = "not_applicable"
	rep.StepSizeSweep[0].MaxAbsQReason = "no away-from-node points available"
	errs := audit.ValidateRobustnessReport(rep)
	if len(errs) > 0 {
		t.Errorf("Expected nil/null metric with status/reason to pass, got errors: %v", errs)
	}

	// 2. Unavailable metric (nil/null) without status/reason: FAIL
	rep.StepSizeSweep[0].MaxAbsQ = nil
	rep.StepSizeSweep[0].MaxAbsQStatus = ""
	rep.StepSizeSweep[0].MaxAbsQReason = ""
	errs = audit.ValidateRobustnessReport(rep)
	if len(errs) == 0 {
		t.Error("Expected nil/null metric without status/reason to fail validation, but it passed")
	}

	// Reset
	rep.StepSizeSweep[0].MaxAbsQ = &val
	rep.StepSizeSweep[0].MaxAbsQStatus = ""
	rep.StepSizeSweep[0].MaxAbsQReason = ""

	// 3. Unavailable metric (-1.0 sentinel) with status/reason: PASS
	rep.StepSizeSweep[0].MaxAbsQ = &valNegOne
	rep.StepSizeSweep[0].MaxAbsQStatus = "not_applicable"
	rep.StepSizeSweep[0].MaxAbsQReason = "sentinel value; no valid away-from-node samples"
	errs = audit.ValidateRobustnessReport(rep)
	if len(errs) > 0 {
		t.Errorf("Expected -1.0 sentinel with status/reason to pass, got errors: %v", errs)
	}

	// 4. Bare -1.0 sentinel without status/reason: FAIL
	rep.StepSizeSweep[0].MaxAbsQ = &valNegOne
	rep.StepSizeSweep[0].MaxAbsQStatus = ""
	rep.StepSizeSweep[0].MaxAbsQReason = ""
	errs = audit.ValidateRobustnessReport(rep)
	if len(errs) == 0 {
		t.Error("Expected bare -1.0 sentinel without status/reason to fail validation, but it passed")
	}

	// Reset
	rep.StepSizeSweep[0].MaxAbsQ = &val
	rep.StepSizeSweep[0].MaxAbsQStatus = ""
	rep.StepSizeSweep[0].MaxAbsQReason = ""

	// 5. NaN still fails: FAIL
	rep.StepSizeSweep[0].MaxAbsQ = &valNaN
	errs = audit.ValidateRobustnessReport(rep)
	if len(errs) == 0 {
		t.Error("Expected NaN metric to fail validation, but it passed")
	}
}
