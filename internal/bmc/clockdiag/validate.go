package clockdiag

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// ValidateClockFragilityReport performs structural and semantic checks on the ClockFragilityReport.
func ValidateClockFragilityReport(r *ClockFragilityReport) []report.ValidationError {
	var errs []report.ValidationError

	// Helper to add failure validation error
	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// 1. Schema version
	if r.SchemaVersion != "bmc0a-clock-fragility-v0.1" {
		fail("schema_version", fmt.Sprintf("unsupported schema version: %s (expected 'bmc0a-clock-fragility-v0.1')", r.SchemaVersion))
	}

	// 2. Toy analysis assertion
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true for clock fragility report")
	}

	// 3. Final Truth Claim
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}

	// 4. Diagnostic Kind
	if r.DiagnosticKind != "clock_monotonicity_fragility" {
		fail("diagnostic_kind", fmt.Sprintf("unsupported diagnostic kind: %s (expected 'clock_monotonicity_fragility')", r.DiagnosticKind))
	}

	// 5. NearZeroDPhiThreshold
	if math.IsNaN(r.NearZeroDPhiThreshold) || math.IsInf(r.NearZeroDPhiThreshold, 0) || r.NearZeroDPhiThreshold <= 0 {
		fail("near_zero_dphi_threshold", "near_zero_dphi_threshold must be finite and positive")
	}

	// 6. Promotion Gate
	if r.PromotionGate.Status != report.StatusBlocked {
		fail("promotion_gate.status", "promotion_gate status must remain blocked")
	}

	// 7. Technical Gate
	if r.TechnicalGate.Status == "" {
		fail("technical_gate.status", "technical_gate status cannot be empty")
	}
	if r.TechnicalGate.Reason == "" {
		fail("technical_gate.reason", "technical_gate reason cannot be empty")
	}

	// 8. EBP Debt checks
	if r.EbpDebt.ClockChoiceDebt != "active" {
		fail("ebp_debt.clock_choice_debt", "clock_choice_debt must remain active")
	}

	// 9. Diagnostic Outcome
	validOutcomes := map[string]bool{
		"clock_stable":  true,
		"clock_fragile": true,
		"mixed":         true,
		"contested":     true,
	}
	if !validOutcomes[r.DiagnosticOutcome] {
		fail("diagnostic_outcome", fmt.Sprintf("invalid diagnostic outcome: %s", r.DiagnosticOutcome))
	}

	// 10. Failed Perturbation Rechecks (Must not be empty)
	if len(r.FailedPerturbationRechecks) == 0 {
		fail("failed_perturbation_rechecks", "failed_perturbation_rechecks cannot be empty")
	}

	// 11. Empty clock events check if fragile
	if r.DiagnosticOutcome == "clock_fragile" && len(r.ClockEvents) == 0 {
		fail("clock_events", "diagnostic outcome is clock_fragile, but no clock events were reported as evidence")
	}

	// 12. Validate StepRefinementResult sweeps
	for idx, s := range r.FailedPerturbationRechecks {
		prefix := fmt.Sprintf("failed_perturbation_rechecks[%d]", idx)
		if math.IsNaN(s.C2Real) || math.IsInf(s.C2Real, 0) {
			fail(prefix+".c2_real", "c2_real must be finite")
		}
		if math.IsNaN(s.K2) || math.IsInf(s.K2, 0) {
			fail(prefix+".k2", "k2 must be finite")
		}
		if math.IsNaN(s.Omega2) || math.IsInf(s.Omega2, 0) {
			fail(prefix+".omega2", "omega2 must be finite")
		}
		if math.IsNaN(s.StepSize) || math.IsInf(s.StepSize, 0) || s.StepSize <= 0 {
			fail(prefix+".step_size", "step_size must be positive and finite")
		}
		if s.Steps <= 0 {
			fail(prefix+".steps", "steps must be greater than zero")
		}
		if s.NumClockEvents < 0 {
			fail(prefix+".num_clock_events", "num_clock_events cannot be negative")
		}

		// Validate nested clock events in refinement results
		for eIdx, e := range s.ClockEvents {
			ePrefix := fmt.Sprintf("%s.clock_events[%d]", prefix, eIdx)
			errs = append(errs, validateClockEvent(e, ePrefix)...)
		}
	}

	// 13. Validate aggregated ClockEvents
	for idx, e := range r.ClockEvents {
		ePrefix := fmt.Sprintf("clock_events[%d]", idx)
		errs = append(errs, validateClockEvent(e, ePrefix)...)
	}

	// 14. Validate CorrelationSummary entries
	for idx, c := range r.CorrelationSummary {
		prefix := fmt.Sprintf("correlation_summary[%d]", idx)
		if c.ParameterSet == "" {
			fail(prefix+".parameter_set", "parameter_set name cannot be empty")
		}
		if math.IsNaN(c.C2Real) || math.IsInf(c.C2Real, 0) {
			fail(prefix+".c2_real", "c2_real must be finite")
		}
		if math.IsNaN(c.K2) || math.IsInf(c.K2, 0) {
			fail(prefix+".k2", "k2 must be finite")
		}
		if math.IsNaN(c.Omega2) || math.IsInf(c.Omega2, 0) {
			fail(prefix+".omega2", "omega2 must be finite")
		}
		if c.NumClockEvents < 0 {
			fail(prefix+".num_clock_events", "num_clock_events cannot be negative")
		}

		errs = append(errs, validateOptionalMetric(c.MinAmplitudeR, c.MinAmplitudeRStatus, c.MinAmplitudeRReason, prefix+".min_amplitude_r")...)
		errs = append(errs, validateOptionalMetric(c.MaxAbsQ, c.MaxAbsQStatus, c.MaxAbsQReason, prefix+".max_abs_q")...)
		errs = append(errs, validateOptionalMetric(c.MaxPhaseGradMagnitude, c.MaxPhaseGradStatus, c.MaxPhaseGradReason, prefix+".max_phase_gradient_magnitude")...)
		errs = append(errs, validateOptionalMetric(c.MinDistToNodeThresh, c.MinDistToNodeStatus, c.MinDistToNodeReason, prefix+".min_dist_to_node_thresh")...)
	}

	// 15. Validate Alternative Clock Summary
	if r.AlternativeClockSummary.PhiMonotonic != model.StatusPass &&
		r.AlternativeClockSummary.PhiMonotonic != model.StatusFail &&
		r.AlternativeClockSummary.PhiMonotonic != model.StatusContested {
		fail("alternative_clock_summary.phi_monotonic", "invalid phi_monotonic status")
	}
	if r.AlternativeClockSummary.AlphaMonotonic != model.StatusPass &&
		r.AlternativeClockSummary.AlphaMonotonic != model.StatusFail &&
		r.AlternativeClockSummary.AlphaMonotonic != model.StatusContested {
		fail("alternative_clock_summary.alpha_monotonic", "invalid alpha_monotonic status")
	}
	if r.AlternativeClockSummary.ClockChoiceDebt != "active" {
		fail("alternative_clock_summary.clock_choice_debt", "clock_choice_debt must remain active")
	}

	// 16. Validate Trajectory Validity Summary
	if r.TrajectoryValiditySummary.TrajectoryValid != model.StatusPass &&
		r.TrajectoryValiditySummary.TrajectoryValid != model.StatusFail &&
		r.TrajectoryValiditySummary.TrajectoryValid != model.StatusContested {
		fail("trajectory_validity_summary.trajectory_valid", "invalid trajectory_valid status")
	}
	if r.TrajectoryValiditySummary.PhiClockValid != model.StatusPass &&
		r.TrajectoryValiditySummary.PhiClockValid != model.StatusFail &&
		r.TrajectoryValiditySummary.PhiClockValid != model.StatusContested {
		fail("trajectory_validity_summary.phi_clock_valid", "invalid phi_clock_valid status")
	}
	if r.TrajectoryValiditySummary.Reason == "" {
		fail("trajectory_validity_summary.reason", "reason cannot be empty")
	}

	return errs
}

func validateClockEvent(e ClockEvent, prefix string) []report.ValidationError {
	var errs []report.ValidationError
	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	if e.Index < 0 {
		fail(prefix+".index", "index cannot be negative")
	}
	if math.IsNaN(e.Lambda) || math.IsInf(e.Lambda, 0) {
		fail(prefix+".lambda", "lambda must be finite")
	}
	if math.IsNaN(e.Alpha) || math.IsInf(e.Alpha, 0) {
		fail(prefix+".alpha", "alpha must be finite")
	}
	if math.IsNaN(e.Phi) || math.IsInf(e.Phi, 0) {
		fail(prefix+".phi", "phi must be finite")
	}
	if math.IsNaN(e.DPhiDLambda) || math.IsInf(e.DPhiDLambda, 0) {
		fail(prefix+".dphi_dlambda", "dphi_dlambda must be finite")
	}
	if math.IsNaN(e.DAlphaDLambda) || math.IsInf(e.DAlphaDLambda, 0) {
		fail(prefix+".dalpha_dlambda", "dalpha_dlambda must be finite")
	}
	if math.IsNaN(e.AmplitudeR) || math.IsInf(e.AmplitudeR, 0) || e.AmplitudeR < 0 {
		fail(prefix+".amplitude_r", "amplitude_r must be positive and finite")
	}
	if e.EventKind != "sign_change" && e.EventKind != "near_zero" && e.EventKind != "direction_reversal" && e.EventKind != "monotonicity_failure" {
		fail(prefix+".event_kind", fmt.Sprintf("invalid event_kind: %s", e.EventKind))
	}
	if e.Severity != "info" && e.Severity != "warning" && e.Severity != "diagnostic" {
		fail(prefix+".severity", fmt.Sprintf("invalid severity: %s", e.Severity))
	}

	errs = append(errs, validateOptionalMetric(e.QValue, e.QStatus, e.QReason, prefix+".q_value")...)
	errs = append(errs, validateOptionalMetric(e.PhaseGradientMagnitude, e.PhaseGradientStatus, e.PhaseGradientReason, prefix+".phase_gradient_magnitude")...)

	return errs
}

func validateOptionalMetric(val *float64, status, reason, name string) []report.ValidationError {
	var errs []report.ValidationError

	if val == nil {
		if status == "" || reason == "" {
			errs = append(errs, report.ValidationError{
				Field:    name,
				Message:  fmt.Sprintf("unavailable metric %s must be paired with explicit status and reason", name),
				Severity: report.ValidationFail,
			})
		}
		return errs
	}

	v := *val
	if math.IsNaN(v) || math.IsInf(v, 0) {
		errs = append(errs, report.ValidationError{
			Field:    name,
			Message:  fmt.Sprintf("metric %s contains non-finite value", name),
			Severity: report.ValidationFail,
		})
	}
	return errs
}
