package clockseg

import (
	"fmt"
	"math"
	"sort"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// ExtractLocalRelationBranch constructs a relational branch from a segment of a trajectory.
// It checks alpha(phi) single-valuedness using singleValuednessEpsilon and requires minBranchSamples.
func ExtractLocalRelationBranch(traj model.Trajectory, seg ClockSegment, epsilon float64, minSamples int) LocalRelationBranch {
	// 1. Slice points
	var points []model.TrajectoryPoint
	if seg.StartIndex >= 0 && seg.EndIndex < len(traj.Points) && seg.StartIndex <= seg.EndIndex {
		points = traj.Points[seg.StartIndex : seg.EndIndex+1]
	}

	samples := len(points)
	lambdaRange := seg.EndLambda - seg.StartLambda

	// Compute clock (phi) range
	var clockRange float64
	var minPhi, maxPhi float64
	if samples > 0 {
		minPhi = points[0].State.Phi
		maxPhi = points[0].State.Phi
		for _, p := range points {
			if p.State.Phi < minPhi {
				minPhi = p.State.Phi
			}
			if p.State.Phi > maxPhi {
				maxPhi = p.State.Phi
			}
		}
		clockRange = maxPhi - minPhi
	}

	// 2. Single-valuedness check
	singleValued, maxAbsNoise := IsAlphaPhiSingleValued(points, epsilon)

	// 3. Validation rules
	validationPassed := true
	var reason string

	if samples < minSamples {
		validationPassed = false
		reason = fmt.Sprintf("branch has too few samples (%d < %d)", samples, minSamples)
	} else if math.IsNaN(lambdaRange) || math.IsInf(lambdaRange, 0) || lambdaRange < 0 {
		validationPassed = false
		reason = "branch lambda range is non-finite or negative"
	} else if math.IsNaN(clockRange) || math.IsInf(clockRange, 0) || clockRange < 0 {
		validationPassed = false
		reason = "branch clock range is non-finite or negative"
	} else if clockRange < 1e-5 {
		validationPassed = false
		reason = fmt.Sprintf("branch clock range is too small (%e < 1e-5)", clockRange)
	} else if !singleValued {
		validationPassed = false
		reason = fmt.Sprintf("alpha(phi) is not single-valued (noise %e > epsilon %e)", maxAbsNoise, epsilon)
	} else {
		reason = "local relational branch satisfies all validation rules"
	}

	return LocalRelationBranch{
		Segment:          seg,
		Samples:          samples,
		LambdaRange:      lambdaRange,
		ClockRange:       clockRange,
		SingleValued:     singleValued,
		MaxAbsNoise:      maxAbsNoise,
		ValidationPassed: validationPassed,
		Reason:           reason,
	}
}

// IsAlphaPhiSingleValued checks if alpha(phi) is single-valued within tolerance epsilon.
func IsAlphaPhiSingleValued(points []model.TrajectoryPoint, epsilon float64) (bool, float64) {
	if len(points) == 0 {
		return true, 0.0
	}

	sorted := make([]model.TrajectoryPoint, len(points))
	copy(sorted, points)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].State.Phi < sorted[j].State.Phi
	})

	maxNoise := 0.0
	singleValued := true

	for i := 0; i < len(sorted); i++ {
		for j := i + 1; j < len(sorted); j++ {
			dPhi := math.Abs(sorted[i].State.Phi - sorted[j].State.Phi)
			if dPhi > epsilon {
				break
			}
			dAlpha := math.Abs(sorted[i].State.Alpha - sorted[j].State.Alpha)
			if dAlpha > maxNoise {
				maxNoise = dAlpha
			}
			if dAlpha > epsilon {
				singleValued = false
			}
		}
	}

	return singleValued, maxNoise
}
