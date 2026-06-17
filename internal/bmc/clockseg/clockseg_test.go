package clockseg

import (
	"bytes"
	"encoding/json"
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// 1. TestClockReadinessReportValidation
func TestClockReadinessReportValidation(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	errs := ValidateClockReadinessReport(rep)
	if len(errs) > 0 {
		t.Errorf("expected generated report to validate successfully, got %d errors: %v", len(errs), errs)
	}
}

// 2. TestClockSegmentationDetectsLocalBranches
func TestClockSegmentationDetectsLocalBranches(t *testing.T) {
	// Let's create a synthetic non-monotonic trajectory
	// Phi goes up then down
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.1, Phi: 2.0}},
		{Lambda: 2.0, State: model.MiniState{Alpha: 1.2, Phi: 3.0}}, // Turning point
		{Lambda: 3.0, State: model.MiniState{Alpha: 1.3, Phi: 2.5}},
		{Lambda: 4.0, State: model.MiniState{Alpha: 1.4, Phi: 2.0}},
	}
	traj := model.Trajectory{Points: points}

	segs, tps := SegmentTrajectory(traj, 1e-10)

	if len(segs) != 2 {
		t.Fatalf("expected 2 segments, got %d", len(segs))
	}
	if len(tps) != 1 {
		t.Fatalf("expected 1 turning point, got %d", len(tps))
	}

	if segs[0].StartIndex != 0 || segs[0].EndIndex != 2 || segs[0].Direction != 1 {
		t.Errorf("segment 0 incorrect: %+v", segs[0])
	}
	if segs[1].StartIndex != 2 || segs[1].EndIndex != 4 || segs[1].Direction != -1 {
		t.Errorf("segment 1 incorrect: %+v", segs[1])
	}
	if tps[0].Index != 2 || tps[0].Phi != 3.0 {
		t.Errorf("turning point incorrect: %+v", tps[0])
	}
}

// 3. TestTurningPointsDetectedAtBranchBoundaries
func TestTurningPointsDetectedAtBranchBoundaries(t *testing.T) {
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.1, Phi: 2.0}},
		{Lambda: 2.0, State: model.MiniState{Alpha: 1.2, Phi: 3.0}}, // TP 1
		{Lambda: 3.0, State: model.MiniState{Alpha: 1.3, Phi: 2.0}},
		{Lambda: 4.0, State: model.MiniState{Alpha: 1.4, Phi: 1.0}}, // TP 2
		{Lambda: 5.0, State: model.MiniState{Alpha: 1.5, Phi: 1.5}},
	}
	traj := model.Trajectory{Points: points}

	segs, tps := SegmentTrajectory(traj, 1e-10)

	if len(segs) != 3 {
		t.Errorf("expected 3 segments, got %d", len(segs))
	}
	if len(tps) != 2 {
		t.Errorf("expected 2 turning points, got %d", len(tps))
	}

	for idx, tp := range tps {
		if segs[idx].EndIndex != tp.Index {
			t.Errorf("turning point %d index %d does not match segment %d end %d", idx, tp.Index, idx, segs[idx].EndIndex)
		}
		if segs[idx+1].StartIndex != tp.Index {
			t.Errorf("turning point %d index %d does not match segment %d start %d", idx, tp.Index, idx+1, segs[idx+1].StartIndex)
		}
	}
}

// 4. TestLocalRelationBranchesAreSingleValuedOnSegments
func TestLocalRelationBranchesAreSingleValuedOnSegments(t *testing.T) {
	// Case 1: single-valued
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.5, Phi: 2.0}},
		{Lambda: 2.0, State: model.MiniState{Alpha: 2.0, Phi: 3.0}},
	}
	singleValued, noise := IsAlphaPhiSingleValued(points, 1e-9)
	if !singleValued {
		t.Errorf("expected single-valued to be true, got false with noise %e", noise)
	}

	// Case 2: not single-valued (different Alpha for same/very close Phi)
	nonSingleValuedPoints := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0000000001}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 2.0, Phi: 1.0000000002}},
	}
	singleValued2, noise2 := IsAlphaPhiSingleValued(nonSingleValuedPoints, 1e-9)
	if singleValued2 {
		t.Errorf("expected single-valued to be false for identical/very close Phi with different Alphas, got true, noise %e", noise2)
	}
}

// 5. TestClockIndependentDiagnosticsDoNotRequirePhiClock
func TestClockIndependentDiagnosticsDoNotRequirePhiClock(t *testing.T) {
	// Ensure we can compute diagnostics even on a weird non-monotonic trajectory
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Verify all returned diagnostics are finite numbers
	d := rep.ClockIndependentDiagnostics
	if math.IsNaN(d.PathLengthInConfigurationSpace) || math.IsInf(d.PathLengthInConfigurationSpace, 0) {
		t.Error("path length must be finite")
	}
	if math.IsNaN(d.TotalLambdaInterval) || math.IsInf(d.TotalLambdaInterval, 0) {
		t.Error("lambda interval must be finite")
	}
	if d.NumValidTrajectoryPoints <= 0 {
		t.Error("must have valid trajectory points")
	}
	if d.NumClockSegments <= 0 {
		t.Error("must have segments")
	}
}

// 6. TestStepRefinementBranchAuditDeterministic
func TestStepRefinementBranchAuditDeterministic(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep1, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 1: %v", err)
	}
	rep2, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 2: %v", err)
	}

	if len(rep1.StepRefinementBranchAudit) != len(rep2.StepRefinementBranchAudit) {
		t.Fatal("step refinement audit list sizes mismatch")
	}

	for i := range rep1.StepRefinementBranchAudit {
		a1 := rep1.StepRefinementBranchAudit[i]
		a2 := rep2.StepRefinementBranchAudit[i]
		if a1.C2Real != a2.C2Real || a1.K2 != a2.K2 || a1.StepSize != a2.StepSize {
			t.Errorf("deterministic mismatch at index %d", i)
		}
	}
}

// 7. TestClockReadinessRejectsFinalTruthClaim
func TestClockReadinessRejectsFinalTruthClaim(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	rep.FinalTruthClaim = true
	errs := ValidateClockReadinessReport(rep)
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

// 8. TestClockReadinessRejectsNonfiniteMetrics
func TestClockReadinessRejectsNonfiniteMetrics(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	rep.ClockIndependentDiagnostics.PathLengthInConfigurationSpace = math.NaN()
	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "clock_independent_diagnostics.path_length" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validation to reject NaN path length")
	}
}

// 9. TestClockReadinessRequiresClockChoiceDebtActive
func TestClockReadinessRequiresClockChoiceDebtActive(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	rep.EbpDebt.ClockChoiceDebt = "resolved"
	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "ebp_debt.clock_choice_debt" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validation to reject resolved clock choice debt")
	}
}

// 10. TestClockReadinessDeterministicJSON
func TestClockReadinessDeterministicJSON(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep1, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 1: %v", err)
	}

	rep2, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report 2: %v", err)
	}

	buf1, err := json.MarshalIndent(rep1, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal 1: %v", err)
	}

	buf2, err := json.MarshalIndent(rep2, "", "  ")
	if err != nil {
		t.Fatalf("failed to marshal 2: %v", err)
	}

	if !bytes.Equal(buf1, buf2) {
		t.Error("generated JSON reports are not byte-identical")
	}
}

// 11. TestClockReadinessRejectsReadyForFriedmannLanguage
func TestClockReadinessRejectsReadyForFriedmannLanguage(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Case A: friedmann_readiness set to local_only_candidate, but we clear all local branches
	rep.LocalRelationBranches = nil
	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "friedmann_readiness" && valErr.Message == "reject friedmann_readiness = local_only_candidate if no valid local relation branches exist" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validator to reject local_only_candidate readiness when there are no valid local branches")
	}

	// Regenerate to restore valid branches
	rep, _ = GenerateClockReadinessReport(safeParams)
	// Case B: friedmann_readiness set to a forbidden option
	rep.FriedmannReadiness = "ready_for_friedmann"
	errs2 := ValidateClockReadinessReport(rep)
	found2 := false
	for _, valErr := range errs2 {
		if valErr.Field == "friedmann_readiness" {
			found2 = true
			break
		}
	}
	if !found2 {
		t.Error("expected validator to reject ready_for_friedmann as an option")
	}
}

// 12. TestClockSegmentsAreNonOverlappingAndOrdered
func TestClockSegmentsAreNonOverlappingAndOrdered(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Inject overlapping segment
	rep.ClockSegments = append(rep.ClockSegments, ClockSegment{
		StartIndex:  0, // Overlaps!
		EndIndex:    5,
		StartLambda: 0.0,
		EndLambda:   1.0,
		Direction:   1,
	})

	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "clock_segments[1].start_index" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validator to reject overlapping segment start index")
	}
}

// 13. TestLocalRelationRejectsTooFewSamples
func TestLocalRelationRejectsTooFewSamples(t *testing.T) {
	// Use Segment with only 2 points
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.1, Phi: 2.0}},
	}
	traj := model.Trajectory{Points: points}
	seg := ClockSegment{StartIndex: 0, EndIndex: 1, StartLambda: 0.0, EndLambda: 1.0, Direction: 1}

	// Should fail since minSamples is 3
	branch := ExtractLocalRelationBranch(traj, seg, 1e-9, 3)
	if branch.ValidationPassed {
		t.Error("expected branch to fail validation due to too few samples")
	}
}

// 14. TestLocalOnlyCandidateDoesNotUnblockFullBMC
func TestLocalOnlyCandidateDoesNotUnblockFullBMC(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Try to unblock the promotion status in debt ledger
	rep.EbpDebt.PromotionStatus = "promoted_clock_readiness_artifact_after_repairs"
	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "ebp_debt.promotion_status" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validator to reject unblocking promotion_status even with local_only_candidate readiness")
	}
}

// 15. TestAll12FragileBranchAuditRuns
func TestAll12FragileBranchAuditRuns(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	if len(rep.StepRefinementBranchAudit) != 12 {
		t.Errorf("expected 12 branch-audit runs, got %d", len(rep.StepRefinementBranchAudit))
	}

	expectedStepSizes := []float64{0.05, 0.025, 0.0125}
	failedConfigs := DefaultFailedConfigs()

	for idx, r := range rep.StepRefinementBranchAudit {
		configIdx := idx / 3
		stepIdx := idx % 3

		expectedConfig := failedConfigs[configIdx]
		if r.C2Real != expectedConfig.C2Real || r.K2 != expectedConfig.K2 || r.Omega2 != expectedConfig.Omega2 {
			t.Errorf("run %d config mismatch: expected %+v, got c2=%.2f, k2=%.2f, omega2=%.2f",
				idx, expectedConfig, r.C2Real, r.K2, r.Omega2)
		}
		if r.StepSize != expectedStepSizes[stepIdx] {
			t.Errorf("run %d step size mismatch: expected %f, got %f", idx, expectedStepSizes[stepIdx], r.StepSize)
		}
	}
}

// 16. TestFlatSegmentEdgeCase
func TestFlatSegmentEdgeCase(t *testing.T) {
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.1, Phi: 1.0}},
		{Lambda: 2.0, State: model.MiniState{Alpha: 1.2, Phi: 1.0}},
	}
	traj := model.Trajectory{Points: points}
	segs, tps := SegmentTrajectory(traj, 1e-10)
	if len(segs) != 1 {
		t.Fatalf("expected 1 segment for flat trajectory, got %d", len(segs))
	}
	if segs[0].Direction != 0 {
		t.Errorf("expected direction 0 for flat segment, got %d", segs[0].Direction)
	}
	if len(tps) != 0 {
		t.Errorf("expected 0 turning points for flat trajectory, got %d", len(tps))
	}
}

// 17. TestSinglePointFinalSegmentEdgeCase
func TestSinglePointFinalSegmentEdgeCase(t *testing.T) {
	points := []model.TrajectoryPoint{
		{Lambda: 0.0, State: model.MiniState{Alpha: 1.0, Phi: 1.0}},
		{Lambda: 1.0, State: model.MiniState{Alpha: 1.1, Phi: 2.0}},
		{Lambda: 2.0, State: model.MiniState{Alpha: 1.2, Phi: 1.9}},
	}
	traj := model.Trajectory{Points: points}
	segs, _ := SegmentTrajectory(traj, 1e-10)
	if len(segs) != 2 {
		t.Fatalf("expected 2 segments, got %d", len(segs))
	}
	if segs[1].StartIndex != 1 || segs[1].EndIndex != 2 {
		t.Errorf("expected final segment from 1 to 2, got %d to %d", segs[1].StartIndex, segs[1].EndIndex)
	}
}

// 18. TestUnsortedTurningPointRejection
func TestUnsortedTurningPointRejection(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	// Inject unsorted turning points
	rep.ClockTurningPoints = []ClockTurningPoint{
		{Index: 5, Lambda: 5.0, Alpha: 1.0, Phi: 1.0, DPhiDLambda: 1.0},
		{Index: 2, Lambda: 2.0, Alpha: 1.0, Phi: 1.0, DPhiDLambda: 1.0},
	}

	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "clock_turning_points[1].index" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validator to reject unsorted turning points")
	}
}

// 19. TestFiniteClockIndependentDiagnostics
func TestFiniteClockIndependentDiagnostics(t *testing.T) {
	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := GenerateClockReadinessReport(safeParams)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	rep.ClockIndependentDiagnostics.MaxAbsQAwayFromNodes = math.NaN()
	errs := ValidateClockReadinessReport(rep)
	found := false
	for _, valErr := range errs {
		if valErr.Field == "clock_independent_diagnostics.max_abs_q" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected validator to reject NaN Q diagnostic")
	}
}

// 20. TestDisallowUnknownFieldsRejection
func TestDisallowUnknownFieldsRejection(t *testing.T) {
	jsonData := `{"schema_version": "bmc0a-clock-readiness-v0.1", "unexpected_key": "val"}`
	dec := json.NewDecoder(bytes.NewReader([]byte(jsonData)))
	dec.DisallowUnknownFields()
	var r ClockReadinessReport
	err := dec.Decode(&r)
	if err == nil {
		t.Error("expected decoder to reject unexpected keys, but it succeeded")
	}
}
