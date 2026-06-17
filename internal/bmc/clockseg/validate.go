package clockseg

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/report"
)

// ValidateClockReadinessReport performs all strict EBP and schema validation checks on the report.
func ValidateClockReadinessReport(r *ClockReadinessReport) []report.ValidationError {
	var errs []report.ValidationError

	fail := func(field, msg string) {
		errs = append(errs, report.ValidationError{
			Field:    field,
			Message:  msg,
			Severity: report.ValidationFail,
		})
	}

	// 1. Schema version
	if r.SchemaVersion != "bmc0a-clock-readiness-v0.1" {
		fail("schema_version", fmt.Sprintf("unsupported schema version: %s (expected 'bmc0a-clock-readiness-v0.1')", r.SchemaVersion))
	}

	// 2. Toy analysis assertion
	if !r.ToyAnalysisOnly {
		fail("toy_analysis_only", "toy_analysis_only must be true")
	}

	// 3. Final Truth Claim
	if r.FinalTruthClaim {
		fail("final_truth_claim", "final_truth_claim must be false")
	}

	// 4. Epsilon and branch samples parameters
	if math.IsNaN(r.SingleValuednessEpsilon) || math.IsInf(r.SingleValuednessEpsilon, 0) || r.SingleValuednessEpsilon <= 0 {
		fail("single_valuedness_epsilon", "single_valuedness_epsilon must be finite and positive")
	}
	if r.MinBranchSamples < 3 {
		fail("min_branch_samples", fmt.Sprintf("min_branch_samples must be at least 3 (got %d)", r.MinBranchSamples))
	}

	// 5. Null models check
	if len(r.NullModels) != 1 || r.NullModels[0] != "No new null models planned. Existing null-model debt remains partial/deferred." {
		fail("null_models", "null_models must exactly state: 'No new null models planned. Existing null-model debt remains partial/deferred.'")
	}

	// 5.5 Readiness scope check
	if r.ReadinessScope != "All 12 step-refinement branch-audit runs are present, covering 4 fragile configurations across 3 step sizes." {
		fail("readiness_scope", "readiness_scope must exactly state that all 12 step-refinement branch-audit runs are present")
	}

	// 6. Friedmann readiness options
	validReadiness := map[string]bool{
		"blocked":              true,
		"local_only_candidate": true,
		"contested":            true,
	}
	if !validReadiness[r.FriedmannReadiness] {
		fail("friedmann_readiness", fmt.Sprintf("invalid friedmann_readiness option: %s", r.FriedmannReadiness))
	}

	// 7. Local relation branches validation
	validBranchCount := 0
	for idx, b := range r.LocalRelationBranches {
		prefix := fmt.Sprintf("local_relation_branches[%d]", idx)
		if b.Segment.StartIndex < 0 || b.Segment.EndIndex < 0 {
			fail(prefix+".segment", "segment indices cannot be negative")
		}
		if b.Segment.StartIndex > b.Segment.EndIndex {
			fail(prefix+".segment", "segment start_index must be <= end_index")
		}
		if math.IsNaN(b.LambdaRange) || math.IsInf(b.LambdaRange, 0) || b.LambdaRange < 0 {
			fail(prefix+".lambda_range", "lambda range must be finite and nonnegative")
		}
		if math.IsNaN(b.ClockRange) || math.IsInf(b.ClockRange, 0) || b.ClockRange < 0 {
			fail(prefix+".clock_range", "clock range must be finite and nonnegative")
		}

		if b.ValidationPassed {
			validBranchCount++
			if !b.SingleValued {
				fail(prefix+".validation_passed", "branch cannot pass validation if alpha(phi) is not single-valued")
			}
			if b.ClockRange < 1e-5 {
				fail(prefix+".clock_range", fmt.Sprintf("local relation branch cannot pass if clock_range is too small (%e < 1e-5)", b.ClockRange))
			}
			if b.Samples < r.MinBranchSamples {
				fail(prefix+".samples", fmt.Sprintf("local relation branch cannot pass with fewer than min_branch_samples (%d < %d)", b.Samples, r.MinBranchSamples))
			}
		} else {
			if b.Reason == "" {
				fail(prefix+".reason", "failed branch must specify validation failure reason")
			}
		}
	}

	// 8. Reject friedmann_readiness = local_only_candidate if no valid local relation branches exist
	if r.FriedmannReadiness == "local_only_candidate" && validBranchCount == 0 {
		fail("friedmann_readiness", "reject friedmann_readiness = local_only_candidate if no valid local relation branches exist")
	}

	// 9. Warnings and readiness checks
	hasReadinessWarning := false
	for _, w := range r.Warnings {
		if w == "local_only_candidate is not Friedmann recovery readiness" {
			hasReadinessWarning = true
			break
		}
	}
	if r.FriedmannReadiness == "local_only_candidate" && !hasReadinessWarning {
		fail("warnings", "must include warning: 'local_only_candidate is not Friedmann recovery readiness'")
	}

	// 10. Segments checks
	for idx, s := range r.ClockSegments {
		prefix := fmt.Sprintf("clock_segments[%d]", idx)
		if s.StartIndex < 0 || s.EndIndex < 0 {
			fail(prefix+".start_index/end_index", "segment indices cannot be negative")
		}
		if s.StartIndex > s.EndIndex {
			fail(prefix+".start_index/end_index", fmt.Sprintf("segment start_index (%d) must be <= end_index (%d)", s.StartIndex, s.EndIndex))
		}
		if math.IsNaN(s.StartLambda) || math.IsInf(s.StartLambda, 0) ||
			math.IsNaN(s.EndLambda) || math.IsInf(s.EndLambda, 0) {
			fail(prefix+".lambdas", "segment lambdas must be finite")
		}
		if s.StartLambda > s.EndLambda {
			fail(prefix+".lambdas", "segment start_lambda must be <= end_lambda")
		}

		if idx > 0 {
			prev := r.ClockSegments[idx-1]
			if s.StartIndex < prev.EndIndex {
				fail(prefix+".start_index", fmt.Sprintf("segments are overlapping: index %d starts at %d which is before previous end %d", idx, s.StartIndex, prev.EndIndex))
			}
		}
	}

	// 11. Turning points checks
	for idx, tp := range r.ClockTurningPoints {
		prefix := fmt.Sprintf("clock_turning_points[%d]", idx)
		if tp.Index < 0 {
			fail(prefix+".index", "turning point index cannot be negative")
		}
		if math.IsNaN(tp.Lambda) || math.IsInf(tp.Lambda, 0) {
			fail(prefix+".lambda", "turning point lambda must be finite")
		}
		if math.IsNaN(tp.Alpha) || math.IsInf(tp.Alpha, 0) ||
			math.IsNaN(tp.Phi) || math.IsInf(tp.Phi, 0) {
			fail(prefix+".coordinates", "turning point coordinates must be finite")
		}
		if math.IsNaN(tp.DPhiDLambda) || math.IsInf(tp.DPhiDLambda, 0) {
			fail(prefix+".dphi_dlambda", "turning point velocity must be finite")
		}

		if idx > 0 {
			prev := r.ClockTurningPoints[idx-1]
			if tp.Index < prev.Index {
				fail(prefix+".index", fmt.Sprintf("turning points are not sorted by index: TP %d has index %d, previous has %d", idx, tp.Index, prev.Index))
			}
			if tp.Index == prev.Index && tp.Lambda <= prev.Lambda {
				fail(prefix+".lambda", fmt.Sprintf("turning points at same index are not sorted by lambda: TP %d has lambda %f, previous has %f", idx, tp.Lambda, prev.Lambda))
			}
		}
	}

	// 12. Clock-Independent Diagnostics checks
	d := r.ClockIndependentDiagnostics
	if math.IsNaN(d.PathLengthInConfigurationSpace) || math.IsInf(d.PathLengthInConfigurationSpace, 0) || d.PathLengthInConfigurationSpace < 0 {
		fail("clock_independent_diagnostics.path_length", "path length must be finite and nonnegative")
	}
	if math.IsNaN(d.TotalLambdaInterval) || math.IsInf(d.TotalLambdaInterval, 0) || d.TotalLambdaInterval < 0 {
		fail("clock_independent_diagnostics.total_lambda_interval", "lambda interval must be finite and nonnegative")
	}
	if d.NumValidTrajectoryPoints < 0 {
		fail("clock_independent_diagnostics.num_valid_trajectory_points", "num valid trajectory points cannot be negative")
	}
	if d.NumClockSegments < 0 {
		fail("clock_independent_diagnostics.num_clock_segments", "num clock segments cannot be negative")
	}
	if d.NumTurningPoints < 0 {
		fail("clock_independent_diagnostics.num_turning_points", "num turning points cannot be negative")
	}
	if math.IsNaN(d.MinAmplitudeR) || math.IsInf(d.MinAmplitudeR, 0) || d.MinAmplitudeR < 0 {
		fail("clock_independent_diagnostics.min_amplitude_r", "min amplitude R must be finite and nonnegative")
	}
	if math.IsNaN(d.MaxAbsQAwayFromNodes) || math.IsInf(d.MaxAbsQAwayFromNodes, 0) || d.MaxAbsQAwayFromNodes < 0 {
		fail("clock_independent_diagnostics.max_abs_q", "max absolute Q must be finite and nonnegative")
	}
	if math.IsNaN(d.MaxPhaseGradient) || math.IsInf(d.MaxPhaseGradient, 0) || d.MaxPhaseGradient < 0 {
		fail("clock_independent_diagnostics.max_phase_gradient", "max phase gradient must be finite and nonnegative")
	}

	// 13. Promotion gate must remain blocked
	if r.PromotionGate.Status != report.StatusBlocked {
		fail("promotion_gate.status", "promotion gate status must remain blocked")
	}

	// 14. EBP debt ledger check
	if r.EbpDebt.ClockChoiceDebt != "active" {
		fail("ebp_debt.clock_choice_debt", "clock_choice_debt must remain active")
	}
	if r.EbpDebt.PromotionStatus != "planned_candidate_only" {
		fail("ebp_debt.promotion_status", "promotion_status must be planned_candidate_only")
	}
	if r.EbpDebt.Sprint5ClockReadiness != "planned_candidate_only" {
		fail("ebp_debt.sprint5_clock_readiness", "sprint5_clock_readiness must be planned_candidate_only")
	}

	return errs
}
