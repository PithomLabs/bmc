package guidance_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

func TestRK4StepperCorrectness(t *testing.T) {
	// Vector field: v(x, y) = (x, y)
	// Starting at (1.0, 1.0), with dt = 0.1
	// Analytic value: exp(0.1) ≈ 1.1051709180756477
	initial := model.MiniState{Alpha: 1.0, Phi: 1.0}
	dt := 0.1

	vel := func(state model.MiniState) (float64, float64) {
		return state.Alpha, state.Phi
	}

	euler := guidance.NewEulerStepper()
	rk4 := guidance.NewRK4Stepper()

	stateEuler := euler.Step(initial, dt, vel)
	stateRK4 := rk4.Step(initial, dt, vel)

	exactVal := math.Exp(0.1)

	eulerErr := math.Abs(stateEuler.Alpha - exactVal)
	rk4Err := math.Abs(stateRK4.Alpha - exactVal)

	if rk4Err >= eulerErr {
		t.Errorf("Expected RK4 error (%e) to be strictly less than Euler error (%e)", rk4Err, eulerErr)
	}

	// RK4 should match exact value to a very high degree of precision (approx 1e-7)
	if rk4Err > 1e-6 {
		t.Errorf("Expected RK4 error to be less than 1e-6, got %e", rk4Err)
	}
}
