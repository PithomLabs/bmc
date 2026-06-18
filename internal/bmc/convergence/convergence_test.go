package convergence

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

func TestConvergenceAuditComputesEulerRK4EndpointDrift(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	audit, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error running convergence audit: %v", err)
	}

	if len(audit.Runs) != 4 {
		t.Fatalf("Expected 4 runs, got %d", len(audit.Runs))
	}

	runsMap := make(map[string]ConvergenceRun)
	for _, r := range audit.Runs {
		runsMap[r.RunID] = r
	}

	// Verify all runs computed successfully
	for _, id := range []string{"euler_dt", "rk4_dt", "rk4_dt_2", "rk4_dt_4"} {
		r, ok := runsMap[id]
		if !ok {
			t.Fatalf("Missing run config %s", id)
		}
		if r.Status != StatusComputed {
			t.Errorf("Expected run %s status to be %s, got %s", id, StatusComputed, r.Status)
		}
		if !r.TrajectoryFinite {
			t.Errorf("Expected run %s trajectory to be finite", id)
		}
		if r.NodeContactDetected {
			t.Errorf("Expected run %s node contact to be false", id)
		}
	}

	// Live computation assertion: verify endpoint drift is calculated from actual endpoints
	euler := runsMap["euler_dt"]
	rk4_ref := runsMap["rk4_dt_4"]

	if euler.FinalAlpha == nil || euler.FinalPhi == nil || rk4_ref.FinalAlpha == nil || rk4_ref.FinalPhi == nil {
		t.Fatal("Final coordinates are nil for computed runs")
	}

	dx := *euler.FinalAlpha - *rk4_ref.FinalAlpha
	dy := *euler.FinalPhi - *rk4_ref.FinalPhi
	expectedDist := math.Sqrt(dx*dx + dy*dy)

	if euler.EndpointDistanceToReference == nil {
		t.Fatal("EndpointDistanceToReference is nil for euler_dt")
	}

	if math.Abs(*euler.EndpointDistanceToReference-expectedDist) > 1e-9 {
		t.Errorf("EndpointDistanceToReference mismatch: computed %f, expected %f", *euler.EndpointDistanceToReference, expectedDist)
	}

	// Verify that metrics are not hard-coded by recomputing under altered trajectory data
	params2 := params
	params2.Alpha0 = 1.0 // alter starting coordinates
	audit2, err := RunAudit(params2)
	if err != nil {
		t.Fatalf("Unexpected error running convergence audit on altered parameters: %v", err)
	}

	runsMap2 := make(map[string]ConvergenceRun)
	for _, r := range audit2.Runs {
		runsMap2[r.RunID] = r
	}

	euler2 := runsMap2["euler_dt"]
	if euler2.EndpointDistanceToReference == nil {
		t.Fatal("EndpointDistanceToReference is nil for altered euler_dt")
	}

	if math.Abs(*euler.EndpointDistanceToReference-*euler2.EndpointDistanceToReference) < 1e-9 {
		t.Error("Expected endpoint distance to change when computed from altered trajectory parameters, but it remained identical")
	}
}


func TestConvergenceAuditUsesSameProfileForAllSteppers(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	audit, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error running convergence audit: %v", err)
	}

	// Total lambda span check: lambda_span = dt * steps = 0.05 * 200 = 10.0
	expectedSpan := params.LambdaStep * float64(params.Steps)

	for _, r := range audit.Runs {
		span := r.DeltaLambda * float64(r.Steps)
		if math.Abs(span-expectedSpan) > 1e-9 {
			t.Errorf("Run %s has incorrect lambda span: got %f, expected %f", r.RunID, span, expectedSpan)
		}
	}
}

func TestConvergenceAuditUsesFinestRK4AsReference(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	audit, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if audit.ReferenceRunID != "rk4_dt_4" {
		t.Errorf("Expected ReferenceRunID to be 'rk4_dt_4', got '%s'", audit.ReferenceRunID)
	}

	for _, r := range audit.Runs {
		if r.RunID == "rk4_dt_4" {
			if r.EndpointDistanceToReference == nil || *r.EndpointDistanceToReference > 1e-15 {
				t.Errorf("Expected reference run to have 0 distance to itself, got %v", r.EndpointDistanceToReference)
			}
		}
	}
}

func TestConvergenceAuditRejectsInvalidStepSize(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	params.LambdaStep = -0.01

	_, err := RunAudit(params)
	if err == nil {
		t.Error("Expected error for negative step size, got nil")
	}

	params.LambdaStep = math.NaN()
	_, err = RunAudit(params)
	if err == nil {
		t.Error("Expected error for NaN step size, got nil")
	}

	params.LambdaStep = math.Inf(1)
	_, err = RunAudit(params)
	if err == nil {
		t.Error("Expected error for +Inf step size, got nil")
	}

	params.LambdaStep = math.Inf(-1)
	_, err = RunAudit(params)
	if err == nil {
		t.Error("Expected error for -Inf step size, got nil")
	}

	params.LambdaStep = 0.05
	params.Steps = 0
	_, err = RunAudit(params)
	if err == nil {
		t.Error("Expected error for zero steps, got nil")
	}
}


func TestConvergenceAuditRejectsNonfiniteTrajectory(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	// Using infinite Alpha0 will produce non-finite values in the trajectory
	params.Alpha0 = math.Inf(1)

	audit, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	for _, r := range audit.Runs {
		if r.Status != StatusBlockedNonfinite {
			t.Errorf("Expected run %s status to be %s, got %s", r.RunID, StatusBlockedNonfinite, r.Status)
		}
		if r.TrajectoryFinite {
			t.Errorf("Expected trajectory to be marked as not finite for run %s", r.RunID)
		}
	}

	if audit.InterpretationStatus != InterpretationBlockedNonfinite {
		t.Errorf("Expected interpretation status to be %s, got %s", InterpretationBlockedNonfinite, audit.InterpretationStatus)
	}
}

func TestConvergenceAuditBlocksNodeContact(t *testing.T) {
	params := model.DefaultSuperpositionNodeProbeParams()
	audit, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	for _, r := range audit.Runs {
		if r.Status != StatusBlockedNodeContact {
			t.Errorf("Expected run %s status to be %s, got %s", r.RunID, StatusBlockedNodeContact, r.Status)
		}
		if !r.NodeContactDetected {
			t.Errorf("Expected run %s node contact to be detected", r.RunID)
		}
		if r.EndpointDistanceToReference != nil {
			t.Errorf("Expected endpoint distance to be nil on node contact for run %s", r.RunID)
		}
		if r.MaxPointwiseDistanceToReference != nil {
			t.Errorf("Expected pointwise distance to be nil on node contact for run %s", r.RunID)
		}
	}

	if audit.InterpretationStatus != InterpretationBlockedNodeContact {
		t.Errorf("Expected interpretation status to be %s, got %s", InterpretationBlockedNodeContact, audit.InterpretationStatus)
	}
}

func TestConvergenceAuditDeterministic(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()

	audit1, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error 1: %v", err)
	}

	audit2, err := RunAudit(params)
	if err != nil {
		t.Fatalf("Unexpected error 2: %v", err)
	}

	if audit1.InterpretationStatus != audit2.InterpretationStatus {
		t.Errorf("Nondeterministic interpretation status: %s vs %s", audit1.InterpretationStatus, audit2.InterpretationStatus)
	}

	for i := range audit1.Runs {
		r1 := audit1.Runs[i]
		r2 := audit2.Runs[i]
		if r1.RunID != r2.RunID {
			t.Fatalf("Mismatch run IDs at index %d", i)
		}
		if r1.Status != r2.Status {
			t.Errorf("Run %s status mismatch: %s vs %s", r1.RunID, r1.Status, r2.Status)
		}
		if r1.EndpointDistanceToReference != nil && r2.EndpointDistanceToReference != nil {
			if *r1.EndpointDistanceToReference != *r2.EndpointDistanceToReference {
				t.Errorf("Run %s drift mismatch: %f vs %f", r1.RunID, *r1.EndpointDistanceToReference, *r2.EndpointDistanceToReference)
			}
		}
	}
}
