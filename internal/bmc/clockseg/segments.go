package clockseg

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// SegmentTrajectory segments the trajectory into monotonic phi intervals and returns turning points.
func SegmentTrajectory(traj model.Trajectory, nearZeroThreshold float64) ([]ClockSegment, []ClockTurningPoint) {
	if len(traj.Points) < 2 {
		return nil, nil
	}

	var segments []ClockSegment
	var turningPoints []ClockTurningPoint

	n := len(traj.Points)
	dphis := make([]float64, n)
	hasDPhi := make([]bool, n)

	for i := 1; i < n; i++ {
		pPrev := traj.Points[i-1]
		pCurr := traj.Points[i]

		if math.IsNaN(pPrev.State.Phi) || math.IsNaN(pCurr.State.Phi) ||
			math.IsNaN(pPrev.State.Alpha) || math.IsNaN(pCurr.State.Alpha) {
			continue
		}

		dt := pCurr.Lambda - pPrev.Lambda
		if dt <= 0 {
			continue
		}
		dphis[i] = (pCurr.State.Phi - pPrev.State.Phi) / dt
		hasDPhi[i] = true
	}

	startIdx := 0
	currentDir := 0 // 0: undetermined, 1: increasing, -1: decreasing

	for i := 1; i < n; i++ {
		if !hasDPhi[i] {
			continue
		}

		diff := traj.Points[i].State.Phi - traj.Points[i-1].State.Phi

		if currentDir == 0 {
			if diff > 1e-12 {
				currentDir = 1
			} else if diff < -1e-12 {
				currentDir = -1
			}
			continue
		}

		isReversal := false
		if currentDir == 1 && diff < 0 {
			isReversal = true
		} else if currentDir == -1 && diff > 0 {
			isReversal = true
		}

		if isReversal {
			tp := ClockTurningPoint{
				Index:       i - 1,
				Lambda:      traj.Points[i-1].Lambda,
				Alpha:       traj.Points[i-1].State.Alpha,
				Phi:         traj.Points[i-1].State.Phi,
				DPhiDLambda: dphis[i-1],
			}
			turningPoints = append(turningPoints, tp)

			segments = append(segments, ClockSegment{
				StartIndex:  startIdx,
				EndIndex:    i - 1,
				StartLambda: traj.Points[startIdx].Lambda,
				EndLambda:   traj.Points[i-1].Lambda,
				Direction:   currentDir,
			})

			startIdx = i - 1
			currentDir = -currentDir
		}
	}

	if startIdx < n-1 {
		segments = append(segments, ClockSegment{
			StartIndex:  startIdx,
			EndIndex:    n - 1,
			StartLambda: traj.Points[startIdx].Lambda,
			EndLambda:   traj.Points[n-1].Lambda,
			Direction:   currentDir,
		})
	}

	return segments, turningPoints
}
