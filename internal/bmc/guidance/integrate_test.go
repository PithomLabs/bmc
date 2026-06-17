package guidance_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

func TestRK4StepperCorrectness(t *testing.T) {
	// Simple non-constant ODE:
	// dα/dλ = α
	// dφ/dλ = -φ
	// Starting at (1.0, 1.0), with dt = 0.1
	// Exact solution at λ=0.1:
	// α(0.1) = exp(0.1) ≈ 1.1051709180756477
	// φ(0.1) = exp(-0.1) ≈ 0.9048374180359595
	initial := model.MiniState{Alpha: 1.0, Phi: 1.0}
	dt := 0.1

	vel := func(state model.MiniState) (float64, float64) {
		return state.Alpha, -state.Phi
	}

	euler := guidance.NewEulerStepper()
	rk4 := guidance.NewRK4Stepper()

	stateEuler := euler.Step(initial, dt, vel)
	stateRK4 := rk4.Step(initial, dt, vel)

	exactAlpha := math.Exp(0.1)
	exactPhi := math.Exp(-0.1)

	eulerErrAlpha := math.Abs(stateEuler.Alpha - exactAlpha)
	eulerErrPhi := math.Abs(stateEuler.Phi - exactPhi)
	eulerErr := eulerErrAlpha + eulerErrPhi

	rk4ErrAlpha := math.Abs(stateRK4.Alpha - exactAlpha)
	rk4ErrPhi := math.Abs(stateRK4.Phi - exactPhi)
	rk4Err := rk4ErrAlpha + rk4ErrPhi

	if rk4Err >= eulerErr {
		t.Errorf("Expected RK4 error (%e) to be strictly less than Euler error (%e)", rk4Err, eulerErr)
	}

	// RK4 should match exact values to a high degree of precision (e.g. within 1e-5)
	if rk4ErrAlpha > 1e-5 || rk4ErrPhi > 1e-5 {
		t.Errorf("Expected RK4 errors to be less than 1e-5, got alpha_err=%e, phi_err=%e", rk4ErrAlpha, rk4ErrPhi)
	}
}
