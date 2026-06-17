package guidance_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/obstruction"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

func TestPlaneWaveTrajectoryFinite(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0)
	stepper := guidance.NewEulerStepper()
	initial := model.MiniState{Alpha: 0.0, Phi: 0.0}
	dt := 0.1
	steps := 100

	traj := guidance.Integrate(pw, initial, stepper, dt, steps, 1e-5)
	if len(traj.Points) != steps+1 {
		t.Fatalf("Expected %d points in trajectory, got %d", steps+1, len(traj.Points))
	}

	if !guidance.IsFinite(traj) {
		t.Error("Expected trajectory to be completely finite, but non-finite values were found")
	}

	// For dalpha/dlambda = k = 1, alpha should be k * steps * dt = 1.0 * 100 * 0.1 = 10.0
	finalPoint := traj.Points[steps]
	expectedAlpha := 10.0
	if math.Abs(finalPoint.State.Alpha-expectedAlpha) > 1e-9 {
		t.Errorf("Expected final Alpha to be close to %f, got %f", expectedAlpha, finalPoint.State.Alpha)
	}

	// For dphi/dlambda = -omega = -1, phi should be -omega * steps * dt = -1.0 * 100 * 0.1 = -10.0
	expectedPhi := -10.0
	if math.Abs(finalPoint.State.Phi-expectedPhi) > 1e-9 {
		t.Errorf("Expected final Phi to be close to %f, got %f", expectedPhi, finalPoint.State.Phi)
	}
}

func TestClockMonotonicityDetection(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0) // omega = 1 != 0
	stepper := guidance.NewEulerStepper()
	initial := model.MiniState{Alpha: 0.0, Phi: 0.0}
	dt := 0.1
	steps := 10

	traj := guidance.Integrate(pw, initial, stepper, dt, steps, 1e-5)
	obs := obstruction.DetectClockNonmonotonicity(traj, "phi")

	if obs.Applies {
		t.Errorf("Expected clock to be monotonic (applies=false), but applies=true: %s", obs.Evidence)
	}
	if obs.Status != model.StatusPass {
		t.Errorf("Expected pass status, got %s", obs.Status)
	}
}

func TestClockMonotonicityFailsWhenOmegaZero(t *testing.T) {
	// If omega = 0, dphi/dlambda = -omega = 0.
	// This means phi will be constant throughout the trajectory, making it non-monotonic as a clock.
	pw := wave.NewPlaneWave(1.0, 0.0)
	stepper := guidance.NewEulerStepper()
	initial := model.MiniState{Alpha: 0.0, Phi: 0.0}
	dt := 0.1
	steps := 10

	traj := guidance.Integrate(pw, initial, stepper, dt, steps, 1e-5)
	obs := obstruction.DetectClockNonmonotonicity(traj, "phi")

	if !obs.Applies {
		t.Error("Expected clock to be non-monotonic (applies=true), but applies=false")
	}
	if obs.Status != model.StatusFail {
		t.Errorf("Expected fail status, got %s", obs.Status)
	}
}
