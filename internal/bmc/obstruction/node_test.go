package obstruction_test

import (
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/obstruction"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

func TestNodeObstructionDetection(t *testing.T) {
	// Construct superposition wave with c1 = 1.0, c2 = -1.0
	// This has a node at (0, 0)
	sw := wave.NewSuperpositionWave(1.0, 0.0, 1.0, 1.0, -1.0, 0.0, 2.0, -2.0)

	// Construct trajectory containing the known node point directly
	traj := model.Trajectory{
		Points: []model.TrajectoryPoint{
			{
				Lambda: 0.0,
				State:  model.MiniState{Alpha: 0.0, Phi: 0.0},
			},
		},
	}

	contact, _ := obstruction.DetectNodeContact(traj, sw, 1e-5)
	if !contact {
		t.Error("Expected DetectNodeContact to return true for trajectory containing the node (0, 0)")
	}
}

func TestSafeSuperpositionAmplitudeStaysAboveNodeThreshold(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	sw := wave.NewSuperpositionWave(
		params.C1Real, params.C1Imag, params.K1, params.Omega1,
		params.C2Real, params.C2Imag, params.K2, params.Omega2,
	)
	stepper := guidance.NewRK4Stepper()
	initialState := model.MiniState{Alpha: params.Alpha0, Phi: params.Phi0}
	traj := guidance.Integrate(sw, initialState, stepper, params.LambdaStep, params.Steps, params.NodeThresh)

	if len(traj.Points) == 0 {
		t.Fatal("Expected trajectory to have points, but got 0")
	}

	for i, p := range traj.Points {
		rVal := wave.AmplitudeField(p.State.Alpha, p.State.Phi, sw)
		if rVal < params.NodeThresh {
			t.Errorf("Point %d at (alpha=%f, phi=%f) has amplitude R=%e, which is below node threshold %e",
				i, p.State.Alpha, p.State.Phi, rVal, params.NodeThresh)
		}
	}
}

func TestSafeSuperpositionPassesChecks(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	rep, err := report.GenerateSuperposition(params, false)
	if err != nil {
		t.Fatalf("Error generating safe superposition report: %v", err)
	}

	// Technical gate must be bmc0a_superposition_safe_gate and it must pass
	if rep.TechnicalGate.Name != "bmc0a_superposition_safe_gate" {
		t.Errorf("Expected technical gate 'bmc0a_superposition_safe_gate', got '%s'", rep.TechnicalGate.Name)
	}
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical gate status 'pass', got '%s'", rep.TechnicalGate.Status)
	}

	// Promotion gate must be blocked
	if rep.PromotionGate.Status != report.StatusBlocked {
		t.Errorf("Expected promotion gate status 'blocked', got '%s'", rep.PromotionGate.Status)
	}

	// Individual checks should pass (except Friedmann which is deferred)
	checks := []string{
		"wdw_residual", "trajectory", "clock_monotonicity", "quantum_potential",
		"classical_limit", "node_detection", "node_contact_free",
		"q_finite_away_from_nodes", "phase_gradient_finite",
	}

	for _, name := range checks {
		check, exists := rep.Checks[name]
		if !exists {
			t.Errorf("Expected check '%s' to exist", name)
			continue
		}
		if check.Status != model.StatusPass {
			t.Errorf("Expected check '%s' to pass, got '%s' (reason: %s)", name, check.Status, check.Reason)
		}
	}

	fCheck, exists := rep.Checks["friedmann_residual"]
	if !exists {
		t.Error("Expected check 'friedmann_residual' to exist")
	} else if fCheck.Status != model.StatusDeferred {
		t.Errorf("Expected check 'friedmann_residual' to be deferred, got '%s'", fCheck.Status)
	}
}

func TestNodeProbeShortCircuitAndValidationGate(t *testing.T) {
	params := model.DefaultSuperpositionNodeProbeParams()
	rep, err := report.GenerateSuperposition(params, false)
	if err != nil {
		t.Fatalf("Error generating node-probe report: %v", err)
	}

	// Assertions required by Sprint 2:

	// 1. node_detection = pass
	ndCheck, exists := rep.Checks["node_detection"]
	if !exists {
		t.Fatal("Expected check 'node_detection' to exist")
	}
	if ndCheck.Status != model.StatusPass {
		t.Errorf("Expected node_detection check status to be 'pass', got '%s'", ndCheck.Status)
	}

	// 2. node_contact_free = fail
	ncCheck, exists := rep.Checks["node_contact_free"]
	if !exists {
		t.Fatal("Expected check 'node_contact_free' to exist")
	}
	if ncCheck.Status != model.StatusFail {
		t.Errorf("Expected node_contact_free check status to be 'fail', got '%s'", ncCheck.Status)
	}

	// 3. node obstruction severity = blocker
	foundNodeObs := false
	for _, obs := range rep.Obstructions {
		if obs.Name == "node_obstruction" {
			foundNodeObs = true
			if !obs.Applies {
				t.Error("Expected node_obstruction Applies to be true")
			}
			if obs.Severity != model.SeverityBlocker {
				t.Errorf("Expected node_obstruction Severity to be 'blocker', got '%s'", obs.Severity)
			}
		}
	}
	if !foundNodeObs {
		t.Error("Expected to find node_obstruction in report obstructions, but none was found")
	}

	// 4. trajectory is empty or explicitly short-circuited/blocked
	tCheck, exists := rep.Checks["trajectory"]
	if !exists {
		t.Fatal("Expected check 'trajectory' to exist")
	}
	if tCheck.Status != model.StatusFail {
		t.Errorf("Expected trajectory check status to be 'fail', got '%s'", tCheck.Status)
	}
	if tCheck.PointsCount == nil || *tCheck.PointsCount != 0 {
		t.Errorf("Expected trajectory points count to be 0 for short-circuit, got %v", tCheck.PointsCount)
	}

	// 5. safe superposition gate does not pass
	if rep.TechnicalGate.Name == "bmc0a_superposition_safe_gate" {
		t.Errorf("Expected technical gate to NOT be bmc0a_superposition_safe_gate, but it is")
	}

	// 6. technical gate = node_detection_validation_gate
	if rep.TechnicalGate.Name != "node_detection_validation_gate" {
		t.Errorf("Expected technical gate name 'node_detection_validation_gate', got '%s'", rep.TechnicalGate.Name)
	}

	// 7. technical gate status = pass
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical gate status 'pass', got '%s'", rep.TechnicalGate.Status)
	}

	// 8. full BMC toy gate remains blocked
	if rep.PromotionGate.Name != "full_bmc_toy_gate" {
		t.Errorf("Expected promotion gate name 'full_bmc_toy_gate', got '%s'", rep.PromotionGate.Name)
	}
	if rep.PromotionGate.Status != "blocked" {
		t.Errorf("Expected promotion gate status 'blocked', got '%s'", rep.PromotionGate.Status)
	}

	// 9. Validation of report should pass because node-probe setup successfully satisfies the node_detection_validation_gate
	errors := report.Validate(rep)
	if len(errors) > 0 {
		t.Errorf("Expected validation of node-probe report to pass, but got errors: %v", errors)
	}
}
