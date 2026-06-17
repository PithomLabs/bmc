package clockdiag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// 1. TestClockFragilityReportValidation
func TestClockFragilityReportValidation(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Should pass initially
	errs := ValidateClockFragilityReport(rep)
	if len(errs) > 0 {
		t.Errorf("expected generated report to validate successfully, got %d errors", len(errs))
	}

	// Test invalid schema version
	origSchema := rep.SchemaVersion
	rep.SchemaVersion = "bad-version"
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for invalid schema version")
	}
	rep.SchemaVersion = origSchema

	// Test invalid toy analysis only
	rep.ToyAnalysisOnly = false
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for toy_analysis_only = false")
	}
	rep.ToyAnalysisOnly = true

	// Test invalid diagnostic kind
	rep.DiagnosticKind = "wrong-kind"
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for invalid diagnostic_kind")
	}
	rep.DiagnosticKind = "clock_monotonicity_fragility"

	// Test unblocked promotion gate
	rep.PromotionGate.Status = model.StatusPass
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for unblocked promotion gate")
	}
	rep.PromotionGate.Status = report.StatusBlocked

	// Test missing technical gate status/reason
	origStatus := rep.TechnicalGate.Status
	rep.TechnicalGate.Status = ""
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for empty technical_gate.status")
	}
	rep.TechnicalGate.Status = origStatus

	origReason := rep.TechnicalGate.Reason
	rep.TechnicalGate.Reason = ""
	if len(ValidateClockFragilityReport(rep)) == 0 {
		t.Error("expected validation failure for empty technical_gate.reason")
	}
	rep.TechnicalGate.Reason = origReason
}

// 2. TestClockEventsDetectedInFailedPerturbations
func TestClockEventsDetectedInFailedPerturbations(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	if len(rep.ClockEvents) == 0 {
		t.Fatal("expected clock events to be detected in failed perturbations, got none")
	}

	// Verify that each kind is valid and event_kind ∈ {sign_change, direction_reversal, monotonicity_failure}
	foundMonotonicityFailure := false
	for _, event := range rep.ClockEvents {
		if event.EventKind == "monotonicity_failure" {
			foundMonotonicityFailure = true
		}
		if event.EventKind != "sign_change" && event.EventKind != "near_zero" &&
			event.EventKind != "direction_reversal" && event.EventKind != "monotonicity_failure" {
			t.Errorf("unexpected event kind: %s", event.EventKind)
		}
	}

	if !foundMonotonicityFailure {
		t.Error("expected to find monotonicity_failure events")
	}
}

// 3. TestFailedPerturbationsRecheckedUnderStepRefinement
func TestFailedPerturbationsRecheckedUnderStepRefinement(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Exactly 12 results (4 configs x 3 step sizes)
	expectedCount := 12
	if len(rep.FailedPerturbationRechecks) != expectedCount {
		t.Errorf("expected %d step refinement results, got %d", expectedCount, len(rep.FailedPerturbationRechecks))
	}

	// Verify step sizes and config ordering are correct and deterministic
	expectedStepSizes := []float64{0.05, 0.025, 0.0125}
	failedConfigs := DefaultFailedConfigs()

	for idx, s := range rep.FailedPerturbationRechecks {
		configIdx := idx / 3
		stepIdx := idx % 3

		expectedConfig := failedConfigs[configIdx]
		if s.C2Real != expectedConfig.C2Real || s.K2 != expectedConfig.K2 || s.Omega2 != expectedConfig.Omega2 {
			t.Errorf("result %d: expected config %+v, got c2=%.2f, k2=%.2f, omega2=%.2f",
				idx, expectedConfig, s.C2Real, s.K2, s.Omega2)
		}

		if s.StepSize != expectedStepSizes[stepIdx] {
			t.Errorf("result %d: expected step size %f, got %f", idx, expectedStepSizes[stepIdx], s.StepSize)
		}

		// Ensure trajectory validity was evaluated
		if !s.TrajectoryValid {
			// In our perturbations, the trajectory should be valid (finite) because they are safe perturbations
			t.Errorf("result %d (c2=%.2f, k2=%.1f, dt=%.4f) has invalid trajectory", idx, s.C2Real, s.K2, s.StepSize)
		}
	}
}

// 4. TestTrajectoryValidityDistinguishedFromClockValidity
func TestTrajectoryValidityDistinguishedFromClockValidity(t *testing.T) {
	// Let's check with some mock refinement results
	mockResults := []StepRefinementResult{
		{
			C2Real:          0.50,
			K2:              2.1,
			Omega2:          -2.1,
			StepSize:        0.05,
			PhiMonotonic:    false,
			TrajectoryValid: true,
		},
	}

	summary := ComputeTrajectoryValiditySummary(mockResults)
	if summary.TrajectoryValid != model.StatusPass {
		t.Errorf("expected trajectory_valid=pass, got %s", summary.TrajectoryValid)
	}
	if summary.PhiClockValid != model.StatusFail {
		t.Errorf("expected phi_clock_valid=fail, got %s", summary.PhiClockValid)
	}
	if !summary.DistinctionPreserved {
		t.Error("expected distinction_preserved=true when trajectory is valid but clock is invalid")
	}

	// Case: both fail
	mockResultsFailBoth := []StepRefinementResult{
		{
			C2Real:          0.50,
			K2:              2.1,
			Omega2:          -2.1,
			StepSize:        0.05,
			PhiMonotonic:    false,
			TrajectoryValid: false,
		},
	}
	summaryFailBoth := ComputeTrajectoryValiditySummary(mockResultsFailBoth)
	if summaryFailBoth.TrajectoryValid != model.StatusFail {
		t.Errorf("expected trajectory_valid=fail, got %s", summaryFailBoth.TrajectoryValid)
	}
	if summaryFailBoth.PhiClockValid != model.StatusFail {
		t.Errorf("expected phi_clock_valid=fail, got %s", summaryFailBoth.PhiClockValid)
	}
	if !summaryFailBoth.DistinctionPreserved {
		t.Error("expected distinction_preserved=true when both fail")
	}
}

// 5. TestAlternativeClockSummary
func TestAlternativeClockSummary(t *testing.T) {
	// Test alpha monotonicity aggregation
	mockResults := []StepRefinementResult{
		{PhiMonotonic: false, AlphaMonotonic: true},
		{PhiMonotonic: false, AlphaMonotonic: true},
	}

	summary := ComputeAlternativeClockSummary(mockResults)
	if summary.PhiMonotonic != model.StatusFail {
		t.Errorf("expected phi_monotonic=fail, got %s", summary.PhiMonotonic)
	}
	if summary.AlphaMonotonic != model.StatusPass {
		t.Errorf("expected alpha_monotonic=pass, got %s", summary.AlphaMonotonic)
	}
	if summary.ClockChoiceDebt != "active" {
		t.Errorf("expected clock_choice_debt=active, got %s", summary.ClockChoiceDebt)
	}
}

// 6. TestClockFragilityRejectsFinalTruthClaim
func TestClockFragilityRejectsFinalTruthClaim(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	rep.FinalTruthClaim = true
	errs := ValidateClockFragilityReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "final_truth_claim" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validation to reject final_truth_claim = true")
	}
}

// 7. TestClockFragilityRejectsNonfiniteMetrics
func TestClockFragilityRejectsNonfiniteMetrics(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	if len(rep.ClockEvents) == 0 {
		t.Fatal("expected at least one clock event")
	}

	// Inject NaN into lambda field of a clock event
	rep.ClockEvents[0].Lambda = math.NaN()
	errs := ValidateClockFragilityReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "clock_events[0].lambda" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validation to reject NaN lambda in clock event")
	}
}

// 8. TestClockFragilityDeterministicJSON
func TestClockFragilityDeterministicJSON(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep1, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 1: %v", err)
	}

	rep2, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 2: %v", err)
	}

	buf1, err := json.MarshalIndent(rep1, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal report 1: %v", err)
	}

	buf2, err := json.MarshalIndent(rep2, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal report 2: %v", err)
	}

	if !bytes.Equal(buf1, buf2) {
		t.Error("generated reports are not byte-identical")
	}
}

// TestOptionalMetricValidation checks that nil/null optional metrics with status/reason pass,
// that nil/null optional metrics without status/reason fail, and that -1.0 is treated as a normal number.
func TestOptionalMetricValidation(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockFragilityReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// 1. A nil optional metric with status/reason passes
	rep.CorrelationSummary[0].MaxAbsQ = nil
	rep.CorrelationSummary[0].MaxAbsQStatus = "not_applicable"
	rep.CorrelationSummary[0].MaxAbsQReason = "test reason"

	errs := ValidateClockFragilityReport(rep)
	if len(errs) > 0 {
		t.Errorf("expected validation to pass for nil optional metric with status/reason, got errors: %v", errs)
	}

	// 2. A nil optional metric without status/reason fails
	rep.CorrelationSummary[0].MaxAbsQ = nil
	rep.CorrelationSummary[0].MaxAbsQStatus = ""
	rep.CorrelationSummary[0].MaxAbsQReason = "test reason" // missing status

	errs = ValidateClockFragilityReport(rep)
	if len(errs) == 0 {
		t.Error("expected validation to fail when optional metric is nil but status is empty")
	}

	rep.CorrelationSummary[0].MaxAbsQStatus = "not_applicable"
	rep.CorrelationSummary[0].MaxAbsQReason = "" // missing reason

	errs = ValidateClockFragilityReport(rep)
	if len(errs) == 0 {
		t.Error("expected validation to fail when optional metric is nil but reason is empty")
	}

	// 3. Verify that -1.0 is treated as a normal number (not a sentinel)
	// That is, if value is -1.0, it does not require status and reason.
	// So if status and reason are empty, it should pass.
	val := -1.0
	rep.CorrelationSummary[0].MaxAbsQ = &val
	rep.CorrelationSummary[0].MaxAbsQStatus = ""
	rep.CorrelationSummary[0].MaxAbsQReason = ""

	errs = ValidateClockFragilityReport(rep)
	if len(errs) > 0 {
		t.Errorf("expected validation to pass for -1.0 without status/reason (no longer treated as sentinel), got errors: %v", errs)
	}
}

// TestDiagnosticOutcomeBranches verifies classification of outcome under step refinement rechecks.
func TestDiagnosticOutcomeBranches(t *testing.T) {
	// Case 1: clock_fragile when all finest-step configs remain nonmonotonic
	mockFragileResults := []StepRefinementResult{
		{C2Real: 0.50, K2: 2.1, StepSize: 0.0125, PhiMonotonic: false},
		{C2Real: 0.55, K2: 1.9, StepSize: 0.0125, PhiMonotonic: false},
	}

	// Case 2: mixed when only some finest-step configs remain nonmonotonic
	mockMixedResults := []StepRefinementResult{
		{C2Real: 0.50, K2: 2.1, StepSize: 0.0125, PhiMonotonic: true},
		{C2Real: 0.55, K2: 1.9, StepSize: 0.0125, PhiMonotonic: false},
	}

	// Case 3: clock_stable when all finest-step configs become monotonic
	mockStableResults := []StepRefinementResult{
		{C2Real: 0.50, K2: 2.1, StepSize: 0.0125, PhiMonotonic: true},
		{C2Real: 0.55, K2: 1.9, StepSize: 0.0125, PhiMonotonic: true},
	}

	computeOutcome := func(refinementResults []StepRefinementResult) string {
		configFinestStable := make(map[string]bool)
		configFinestTotal := make(map[string]int)

		for _, r := range refinementResults {
			if r.StepSize == 0.0125 {
				key := fmt.Sprintf("c2=%.2f,k2=%.1f", r.C2Real, r.K2)
				configFinestTotal[key]++
				if r.PhiMonotonic {
					configFinestStable[key] = true
				}
			}
		}

		stableCount := 0
		for _, stable := range configFinestStable {
			if stable {
				stableCount++
			}
		}

		if stableCount == len(configFinestTotal) && len(configFinestTotal) > 0 {
			return "clock_stable"
		} else if stableCount == 0 && len(configFinestTotal) > 0 {
			return "clock_fragile"
		} else {
			return "mixed"
		}
	}

	if outcome := computeOutcome(mockFragileResults); outcome != "clock_fragile" {
		t.Errorf("expected clock_fragile, got %s", outcome)
	}

	if outcome := computeOutcome(mockMixedResults); outcome != "mixed" {
		t.Errorf("expected mixed, got %s", outcome)
	}

	if outcome := computeOutcome(mockStableResults); outcome != "clock_stable" {
		t.Errorf("expected clock_stable, got %s", outcome)
	}
}
