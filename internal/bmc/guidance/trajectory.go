package guidance

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// Integrate integrates the Bohmian guidance equations forward.
// It checks the node threshold to prevent integrations through nodes.
func Integrate(wf wave.WaveFunction, initial model.MiniState, stepper Stepper, dt float64, steps int, nodeThresh float64) model.Trajectory {
	points := make([]model.TrajectoryPoint, steps+1)
	points[0] = model.TrajectoryPoint{
		Lambda: 0.0,
		State:  initial,
	}

	velFunc := func(state model.MiniState) (float64, float64) {
		return Velocity(wf, state.Alpha, state.Phi, nodeThresh)
	}

	currentState := initial
	for i := 1; i <= steps; i++ {
		// If current state is already at a node, we stop integrating and fill remaining points with NaN
		rVal := wave.AmplitudeField(currentState.Alpha, currentState.Phi, wf)
		if rVal < nodeThresh {
			for j := i; j <= steps; j++ {
				points[j] = model.TrajectoryPoint{
					Lambda: float64(j) * dt,
					State: model.MiniState{
						Alpha: math.NaN(),
						Phi:   math.NaN(),
					},
				}
			}
			break
		}

		currentState = stepper.Step(currentState, dt, velFunc)
		points[i] = model.TrajectoryPoint{
			Lambda: float64(i) * dt,
			State:  currentState,
		}
	}

	return model.Trajectory{
		Points: points,
	}
}

// IsFinite checks if all coordinates in the trajectory are finite numbers.
func IsFinite(traj model.Trajectory) bool {
	for _, p := range traj.Points {
		if math.IsNaN(p.State.Alpha) || math.IsInf(p.State.Alpha, 0) ||
			math.IsNaN(p.State.Phi) || math.IsInf(p.State.Phi, 0) {
			return false
		}
	}
	return true
}

// CheckTrajectory validates if the trajectory is finite and integrated correctly.
func CheckTrajectory(traj model.Trajectory) model.CheckResult {
	finite := IsFinite(traj)
	count := len(traj.Points)

	var status model.CheckStatus
	var reason string
	if finite {
		status = model.StatusPass
		reason = "Bohmian trajectory is finite and successfully integrated."
	} else {
		status = model.StatusFail
		reason = "Bohmian trajectory contains non-finite values (NaN or Inf) due to node proximity or integration failure."
	}

	return model.CheckResult{
		Status:      status,
		Pass:        finite,
		Reason:      reason,
		Finite:      &finite,
		PointsCount: &count,
	}
}
