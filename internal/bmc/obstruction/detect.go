package obstruction

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// DetectWdWResidual checks if the Wheeler-DeWitt residual exceeds tolerance.
func DetectWdWResidual(maxRes, tolerance float64) model.Obstruction {
	absRes := math.Abs(maxRes)
	applies := absRes > tolerance

	var status model.CheckStatus
	var evidence string
	var consequence string

	if applies {
		status = model.StatusFail
		evidence = fmt.Sprintf("Wheeler-DeWitt residual %e exceeds tolerance %e", absRes, tolerance)
		consequence = "Downgrade model promotion status; WdW equation is not satisfied."
	} else {
		status = model.StatusPass
		evidence = fmt.Sprintf("Wheeler-DeWitt residual %e satisfies tolerance %e", absRes, tolerance)
		consequence = "Residual is clean."
	}

	return model.Obstruction{
		Name:        WdWResidualObstruction,
		Applies:     applies,
		Severity:    model.SeverityBlocker,
		Evidence:    evidence,
		Consequence: consequence,
		Status:      status,
	}
}

// DetectNonfiniteQ checks if any value of Q along the trajectory is NaN or Inf.
func DetectNonfiniteQ(qValues []float64) model.Obstruction {
	applies := false
	evidence := "All quantum potential values along the trajectory are finite."
	status := model.StatusPass

	for i, q := range qValues {
		if math.IsNaN(q) || math.IsInf(q, 0) {
			applies = true
			evidence = fmt.Sprintf("Non-finite quantum potential value %f found at index %d", q, i)
			status = model.StatusFail
			break
		}
	}

	consequence := "Continue trajectory analysis."
	if applies {
		consequence = "Block promotion; trajectory enters undefined quantum potential region."
	}

	return model.Obstruction{
		Name:        NonfiniteQObstruction,
		Applies:     applies,
		Severity:    model.SeverityBlocker,
		Evidence:    evidence,
		Consequence: consequence,
		Status:      status,
	}
}

// DetectClockNonmonotonicity checks if the selected relational clock variable (usually phi) is strictly monotonic.
func DetectClockNonmonotonicity(traj model.Trajectory, variable string) model.Obstruction {
	// Filter out NaNs (short-circuited trajectories)
	var values []float64
	for _, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
			continue
		}
		if variable == "phi" {
			values = append(values, p.State.Phi)
		} else {
			values = append(values, p.State.Alpha)
		}
	}

	if len(values) <= 1 {
		return model.Obstruction{
			Name:        ClockNonmonotonicityObstruction,
			Applies:     true,
			Severity:    model.SeverityBlocker,
			Evidence:    "Trajectory has insufficient finite points to determine monotonicity.",
			Consequence: "Block promotion; relational clock is undefined.",
			Status:      model.StatusFail,
		}
	}

	// Check monotonicity
	isIncreasing := true
	isDecreasing := true

	for i := 1; i < len(values); i++ {
		diff := values[i] - values[i-1]
		if diff <= 0 {
			isIncreasing = false
		}
		if diff >= 0 {
			isDecreasing = false
		}
	}

	monotonic := isIncreasing || isDecreasing
	applies := !monotonic

	var status model.CheckStatus
	var evidence string
	var consequence string

	if applies {
		status = model.StatusFail
		evidence = fmt.Sprintf("Relational clock variable '%s' is not strictly monotonic.", variable)
		consequence = "Block promotion; relational history a(phi) cannot be uniquely defined."
	} else {
		status = model.StatusPass
		evidence = fmt.Sprintf("Relational clock variable '%s' is strictly monotonic.", variable)
		consequence = "Relational clock is valid."
	}

	return model.Obstruction{
		Name:        ClockNonmonotonicityObstruction,
		Applies:     applies,
		Severity:    model.SeverityBlocker,
		Evidence:    evidence,
		Consequence: consequence,
		Status:      status,
	}
}

// DetectOverclaimBlocker checks if any final-truth claim flag is raised.
func DetectOverclaimBlocker(finalTruthClaim bool) model.Obstruction {
	var status model.CheckStatus
	var evidence string
	var consequence string

	if finalTruthClaim {
		status = model.StatusFail
		evidence = "Final-truth claim was explicitly asserted in configuration."
		consequence = "Block promotion; EBP 2.1 requires all claims to remain toy-only."
	} else {
		status = model.StatusPass
		evidence = "No final-truth language detected."
		consequence = "Continue as toy-only artifact."
	}

	return model.Obstruction{
		Name:        FullQgOverclaimBlocker,
		Applies:     finalTruthClaim,
		Severity:    model.SeverityBlocker,
		Evidence:    evidence,
		Consequence: consequence,
		Status:      status,
	}
}

// DetectNodeContact checks if the trajectory came into contact with a node.
func DetectNodeContact(traj model.Trajectory, wf wave.WaveFunction, nodeThresh float64) (contact bool, evidence string) {
	if len(traj.Points) == 0 {
		return true, "Trajectory is empty; initial state was below node threshold."
	}
	for i, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
			return true, fmt.Sprintf("Trajectory entered node region (NaN found at step %d).", i)
		}
		rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, wf)
		if rVal < nodeThresh {
			return true, fmt.Sprintf("Wavefunction amplitude %e is below node threshold %e at step %d.", rVal, nodeThresh, i)
		}
	}
	return false, "No trajectory points came within node threshold."
}

// DetectPhaseGradientFinite checks if the phase gradient is finite and within bounds.
func DetectPhaseGradientFinite(traj model.Trajectory, wf wave.WaveFunction, nodeThresh, maxPhaseGrad float64) (finite bool, evidence string) {
	hasFinitePoints := false
	for i, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
			continue
		}
		rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, wf)
		if rVal < nodeThresh {
			continue
		}
		hasFinitePoints = true
		dSa, dSp := wave.PhaseGradient(p.State.Alpha, p.State.Phi, wf)
		if math.IsNaN(dSa) || math.IsInf(dSa, 0) || math.IsNaN(dSp) || math.IsInf(dSp, 0) {
			return false, fmt.Sprintf("Non-finite phase gradient found at step %d.", i)
		}
		magnitude := math.Sqrt(dSa*dSa + dSp*dSp)
		if magnitude > maxPhaseGrad {
			return false, fmt.Sprintf("Phase gradient magnitude %f exceeded max bound %f at step %d.", magnitude, maxPhaseGrad, i)
		}
	}
	if !hasFinitePoints {
		return false, "No valid non-node points to evaluate phase gradient."
	}
	return true, "Phase gradient remains finite and bounded along the trajectory."
}

// DetectQFiniteAwayFromNodes checks if Q is finite for all points outside node regions.
func DetectQFiniteAwayFromNodes(traj model.Trajectory, qValues []float64, wf wave.WaveFunction, nodeThresh float64) (finite bool, evidence string) {
	hasFinitePoints := false
	for i, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsNaN(p.State.Phi) {
			continue
		}
		rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, wf)
		if rVal < nodeThresh {
			continue
		}
		if i >= len(qValues) {
			break
		}
		hasFinitePoints = true
		q := qValues[i]
		if math.IsNaN(q) || math.IsInf(q, 0) {
			return false, fmt.Sprintf("Non-finite quantum potential value %f found at step %d.", q, i)
		}
	}
	if !hasFinitePoints {
		return false, "No valid non-node points to evaluate quantum potential."
	}
	return true, "Quantum potential is finite for all points away from nodes."
}

// DetectAllPlaneWave runs all active plane-wave obstruction detectors.
func DetectAllPlaneWave(params model.PlaneWaveParams, traj model.Trajectory, maxRes float64, qValues []float64, finalTruthClaim bool) []model.Obstruction {
	obstructions := make([]model.Obstruction, 0, 9)

	obstructions = append(obstructions, DetectWdWResidual(maxRes, params.Tolerance))
	obstructions = append(obstructions, DetectNonfiniteQ(qValues))
	obstructions = append(obstructions, DetectClockNonmonotonicity(traj, "phi"))
	obstructions = append(obstructions, DetectOverclaimBlocker(finalTruthClaim))

	// Deferred placeholders
	obstructions = append(obstructions, NewDeferred(NodeObstruction, "Node detection deferred to Sprint 2."))
	obstructions = append(obstructions, NewDeferred(PhaseUnwrapObstruction, "Phase unwrapping evaluation deferred to Sprint 2."))
	obstructions = append(obstructions, NewDeferred(ClassicalLimitFailure, "Classical limit comparison requires massive scalar comparison (Sprint 3)."))
	obstructions = append(obstructions, NewDeferred(LapseOrTimeInterpretationDebt, "Relational clock time interpretation remains active debt."))
	obstructions = append(obstructions, NewDeferred(MeasureProblemDeferred, "Ensemble measure problem is deferred."))

	return obstructions
}

// DetectAllSuperposition runs active superposition obstruction detectors.
func DetectAllSuperposition(params model.SuperpositionParams, traj model.Trajectory, maxRes float64, qValues []float64, finalTruthClaim bool, sw wave.SuperpositionWave) []model.Obstruction {
	obstructions := make([]model.Obstruction, 0, 9)

	obstructions = append(obstructions, DetectWdWResidual(maxRes, params.Tolerance))
	obstructions = append(obstructions, DetectClockNonmonotonicity(traj, "phi"))
	obstructions = append(obstructions, DetectOverclaimBlocker(finalTruthClaim))

	// Node contact blocker detection
	contact, nodeEvidence := DetectNodeContact(traj, sw, params.NodeThresh)
	var nodeStatus model.CheckStatus = model.StatusPass
	if contact {
		nodeStatus = model.StatusFail
	}
	obstructions = append(obstructions, model.Obstruction{
		Name:        NodeObstruction,
		Applies:     contact,
		Severity:    model.SeverityBlocker,
		Evidence:    nodeEvidence,
		Consequence: "Block safe promotion gate; trajectory came into contact with a node.",
		Status:      nodeStatus,
	})

	// Phase unwrap / gradient finite blocker detection
	gradientFinite, gradEvidence := DetectPhaseGradientFinite(traj, sw, params.NodeThresh, params.MaxPhaseGrad)
	var gradStatus model.CheckStatus = model.StatusPass
	if !gradientFinite {
		gradStatus = model.StatusFail
	}
	obstructions = append(obstructions, model.Obstruction{
		Name:        PhaseUnwrapObstruction,
		Applies:     !gradientFinite,
		Severity:    model.SeverityBlocker,
		Evidence:    gradEvidence,
		Consequence: "Block safe promotion gate; phase gradient is non-finite or exceeded bounds.",
		Status:      gradStatus,
	})

	// Q potential blocker detection
	qFinite, qEvidence := DetectQFiniteAwayFromNodes(traj, qValues, sw, params.NodeThresh)
	var qStatus model.CheckStatus = model.StatusPass
	if !qFinite {
		qStatus = model.StatusFail
	}
	obstructions = append(obstructions, model.Obstruction{
		Name:        NonfiniteQObstruction,
		Applies:     !qFinite,
		Severity:    model.SeverityBlocker,
		Evidence:    qEvidence,
		Consequence: "Block safe promotion gate; quantum potential is non-finite away from nodes.",
		Status:      qStatus,
	})

	// Deferred placeholders
	obstructions = append(obstructions, NewDeferred(ClassicalLimitFailure, "Classical limit comparison requires massive scalar comparison (Sprint 3)."))
	obstructions = append(obstructions, NewDeferred(LapseOrTimeInterpretationDebt, "Relational clock time interpretation remains active debt."))
	obstructions = append(obstructions, NewDeferred(MeasureProblemDeferred, "Ensemble measure problem is deferred."))

	return obstructions
}
