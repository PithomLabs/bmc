package guidance

import (
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// Stepper defines the integration step contract.
type Stepper interface {
	Step(state model.MiniState, dt float64, vel func(model.MiniState) (float64, float64)) model.MiniState
}

// EulerStepper implements a first-order Euler integration step.
type EulerStepper struct{}

// NewEulerStepper constructs an EulerStepper.
func NewEulerStepper() EulerStepper {
	return EulerStepper{}
}

// Step performs a single forward Euler integration step: state + dt * velocity.
func (e EulerStepper) Step(state model.MiniState, dt float64, vel func(model.MiniState) (float64, float64)) model.MiniState {
	dAlpha, dPhi := vel(state)
	return model.MiniState{
		Alpha: state.Alpha + dt*dAlpha,
		Phi:   state.Phi + dt*dPhi,
	}
}

// RK4Stepper implements a fourth-order Runge-Kutta integration step.
type RK4Stepper struct{}

// NewRK4Stepper constructs an RK4Stepper.
func NewRK4Stepper() RK4Stepper {
	return RK4Stepper{}
}

// Step performs a single Runge-Kutta 4th order integration step.
func (rk RK4Stepper) Step(state model.MiniState, dt float64, vel func(model.MiniState) (float64, float64)) model.MiniState {
	// k1
	k1Alpha, k1Phi := vel(state)

	// k2
	s2 := model.MiniState{
		Alpha: state.Alpha + 0.5*dt*k1Alpha,
		Phi:   state.Phi + 0.5*dt*k1Phi,
	}
	k2Alpha, k2Phi := vel(s2)

	// k3
	s3 := model.MiniState{
		Alpha: state.Alpha + 0.5*dt*k2Alpha,
		Phi:   state.Phi + 0.5*dt*k2Phi,
	}
	k3Alpha, k3Phi := vel(s3)

	// k4
	s4 := model.MiniState{
		Alpha: state.Alpha + dt*k3Alpha,
		Phi:   state.Phi + dt*k3Phi,
	}
	k4Alpha, k4Phi := vel(s4)

	return model.MiniState{
		Alpha: state.Alpha + (dt/6.0)*(k1Alpha+2.0*k2Alpha+2.0*k3Alpha+k4Alpha),
		Phi:   state.Phi + (dt/6.0)*(k1Phi+2.0*k2Phi+2.0*k3Phi+k4Phi),
	}
}

