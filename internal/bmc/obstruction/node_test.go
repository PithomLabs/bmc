package obstruction_test

import (
	"testing"

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

	// 1. Technical gate name must be node_detection_validation_gate and status must be pass
	if rep.TechnicalGate.Name != "node_detection_validation_gate" {
		t.Errorf("Expected technical gate name 'node_detection_validation_gate', got '%s'", rep.TechnicalGate.Name)
	}
	if rep.TechnicalGate.Status != model.StatusPass {
		t.Errorf("Expected technical gate status 'pass', got '%s'", rep.TechnicalGate.Status)
	}

	// 2. Trajectory should be empty (short-circuited)
	// Wait! Let's check checks map to see if checks are set as expected
	tCheck, exists := rep.Checks["trajectory"]
	if !exists {
		t.Fatal("Expected check 'trajectory' to exist")
	}
	if tCheck.Status != model.StatusFail {
		t.Errorf("Expected trajectory check status to be 'fail', got '%s'", tCheck.Status)
	}
	if tCheck.Pass {
		t.Error("Expected trajectory check Pass to be false")
	}

	// 3. Node checks
	ncCheck, exists := rep.Checks["node_contact_free"]
	if !exists {
		t.Fatal("Expected check 'node_contact_free' to exist")
	}
	if ncCheck.Status != model.StatusFail {
		t.Errorf("Expected node_contact_free check status to be 'fail', got '%s'", ncCheck.Status)
	}

	ndCheck, exists := rep.Checks["node_detection"]
	if !exists {
		t.Fatal("Expected check 'node_detection' to exist")
	}
	if ndCheck.Status != model.StatusPass {
		t.Errorf("Expected node_detection check status to be 'pass', got '%s'", ndCheck.Status)
	}

	// 4. Phase and Q checks must be contested
	pgCheck, exists := rep.Checks["phase_gradient_finite"]
	if !exists {
		t.Fatal("Expected check 'phase_gradient_finite' to exist")
	}
	if pgCheck.Status != model.StatusContested {
		t.Errorf("Expected phase_gradient_finite status to be 'contested', got '%s'", pgCheck.Status)
	}

	qCheck, exists := rep.Checks["q_finite_away_from_nodes"]
	if !exists {
		t.Fatal("Expected check 'q_finite_away_from_nodes' to exist")
	}
	if qCheck.Status != model.StatusContested {
		t.Errorf("Expected q_finite_away_from_nodes status to be 'contested', got '%s'", qCheck.Status)
	}

	// 5. Validation of report should pass because node-probe setup successfully satisfies the node_detection_validation_gate
	errors := report.Validate(rep)
	if len(errors) > 0 {
		t.Errorf("Expected validation of node-probe report to pass, but got errors: %v", errors)
	}
}
