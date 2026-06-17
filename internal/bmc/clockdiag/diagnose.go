package clockdiag

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

type FailedPerturbationConfig struct {
	C2Real float64 `json:"c2_real"`
	K2     float64 `json:"k2"`
	Omega2 float64 `json:"omega2"`
}

type StepRefinementResult struct {
	C2Real          float64      `json:"c2_real"`
	K2              float64      `json:"k2"`
	Omega2          float64      `json:"omega2"`
	StepSize        float64      `json:"step_size"`
	Steps           int          `json:"steps"`
	PhiMonotonic    bool         `json:"phi_monotonic"`
	AlphaMonotonic  bool         `json:"alpha_monotonic"`
	NumClockEvents  int          `json:"num_clock_events"`
	TrajectoryValid bool         `json:"trajectory_valid"`
	ClockEvents     []ClockEvent `json:"clock_events"`
}

type AlternativeClockSummary struct {
	PhiMonotonic     model.CheckStatus `json:"phi_monotonic"`
	AlphaMonotonic   model.CheckStatus `json:"alpha_monotonic"`
	BothMonotonic    bool              `json:"both_monotonic"`
	NeitherMonotonic bool              `json:"neither_monotonic"`
	ClockChoiceDebt  string            `json:"clock_choice_debt"`
}

type TrajectoryValiditySummary struct {
	TrajectoryValid      model.CheckStatus `json:"trajectory_valid"`
	PhiClockValid        model.CheckStatus `json:"phi_clock_valid"`
	DistinctionPreserved bool              `json:"distinction_preserved"`
	Reason               string            `json:"reason"`
}

// DefaultFailedConfigs returns the four failed configurations identified in Sprint 3.
func DefaultFailedConfigs() []FailedPerturbationConfig {
	return []FailedPerturbationConfig{
		{C2Real: 0.50, K2: 2.1, Omega2: -2.1},
		{C2Real: 0.55, K2: 1.9, Omega2: -1.9},
		{C2Real: 0.55, K2: 2.0, Omega2: -2.0},
		{C2Real: 0.55, K2: 2.1, Omega2: -2.1},
	}
}

// RunStepRefinementRechecks runs step refinement on the four failed configurations.
func RunStepRefinementRechecks(safeParams model.SuperpositionParams, failedConfigs []FailedPerturbationConfig, nearZeroDPhiThreshold float64) ([]StepRefinementResult, error) {
	stepSizes := []float64{0.05, 0.025, 0.0125}
	results := make([]StepRefinementResult, 0, len(failedConfigs)*len(stepSizes))

	for _, config := range failedConfigs {
		// Create the superposition wave function for this configuration
		sw := wave.NewSuperpositionWave(
			safeParams.C1Real, safeParams.C1Imag, safeParams.K1, safeParams.Omega1,
			config.C2Real, safeParams.C2Imag, config.K2, config.Omega2,
		)

		for _, dt := range stepSizes {
			steps := int(math.Round(10.0 / dt))

			stepper := guidance.NewRK4Stepper()
			initialState := model.MiniState{Alpha: safeParams.Alpha0, Phi: safeParams.Phi0}
			traj := guidance.Integrate(sw, initialState, stepper, dt, steps, safeParams.NodeThresh)

			// 1. Trajectory Validity check
			trajValid := guidance.IsFinite(traj) && len(traj.Points) > 1

			// 2. Monotonicity checks
			phiMonotonic := true
			alphaMonotonic := true

			if len(traj.Points) <= 1 {
				phiMonotonic = false
				alphaMonotonic = false
			} else {
				phiInc := true
				phiDec := true
				alphaInc := true
				alphaDec := true

				for i := 1; i < len(traj.Points); i++ {
					pPrev := traj.Points[i-1]
					pCurr := traj.Points[i]

					if math.IsNaN(pPrev.State.Phi) || math.IsNaN(pCurr.State.Phi) {
						continue
					}

					phiDiff := pCurr.State.Phi - pPrev.State.Phi
					if phiDiff <= 0 {
						phiInc = false
					}
					if phiDiff >= 0 {
						phiDec = false
					}

					alphaDiff := pCurr.State.Alpha - pPrev.State.Alpha
					if alphaDiff <= 0 {
						alphaInc = false
					}
					if alphaDiff >= 0 {
						alphaDec = false
					}
				}

				phiMonotonic = phiInc || phiDec
				alphaMonotonic = alphaInc || alphaDec
			}

			// 3. Detect clock events
			events := DetectClockEvents(traj, sw, safeParams, nearZeroDPhiThreshold)

			results = append(results, StepRefinementResult{
				C2Real:          config.C2Real,
				K2:              config.K2,
				Omega2:          config.Omega2,
				StepSize:        dt,
				Steps:           steps,
				PhiMonotonic:    phiMonotonic,
				AlphaMonotonic:  alphaMonotonic,
				NumClockEvents:  len(events),
				TrajectoryValid: trajValid,
				ClockEvents:     events,
			})
		}
	}

	return results, nil
}

// ComputeAlternativeClockSummary aggregates the step refinement results.
func ComputeAlternativeClockSummary(refinementResults []StepRefinementResult) AlternativeClockSummary {
	var phiPassCount, phiFailCount int
	var alphaPassCount, alphaFailCount int

	for _, r := range refinementResults {
		if r.PhiMonotonic {
			phiPassCount++
		} else {
			phiFailCount++
		}

		if r.AlphaMonotonic {
			alphaPassCount++
		} else {
			alphaFailCount++
		}
	}

	var phiStatus model.CheckStatus
	if phiPassCount == len(refinementResults) {
		phiStatus = model.StatusPass
	} else if phiFailCount == len(refinementResults) {
		phiStatus = model.StatusFail
	} else {
		phiStatus = model.StatusContested
	}

	var alphaStatus model.CheckStatus
	if alphaPassCount == len(refinementResults) {
		alphaStatus = model.StatusPass
	} else if alphaFailCount == len(refinementResults) {
		alphaStatus = model.StatusFail
	} else {
		alphaStatus = model.StatusContested
	}

	return AlternativeClockSummary{
		PhiMonotonic:     phiStatus,
		AlphaMonotonic:   alphaStatus,
		BothMonotonic:    phiStatus == model.StatusPass && alphaStatus == model.StatusPass,
		NeitherMonotonic: phiStatus == model.StatusFail && alphaStatus == model.StatusFail,
		ClockChoiceDebt:  "active",
	}
}

// ComputeTrajectoryValiditySummary assesses distinction between trajectory and clock validity.
func ComputeTrajectoryValiditySummary(refinementResults []StepRefinementResult) TrajectoryValiditySummary {
	var trajPassCount, trajFailCount int
	var phiPassCount, phiFailCount int

	for _, r := range refinementResults {
		if r.TrajectoryValid {
			trajPassCount++
		} else {
			trajFailCount++
		}

		if r.PhiMonotonic {
			phiPassCount++
		} else {
			phiFailCount++
		}
	}

	var trajStatus model.CheckStatus
	if trajPassCount == len(refinementResults) {
		trajStatus = model.StatusPass
	} else if trajFailCount == len(refinementResults) {
		trajStatus = model.StatusFail
	} else {
		trajStatus = model.StatusContested
	}

	var phiStatus model.CheckStatus
	if phiPassCount == len(refinementResults) {
		phiStatus = model.StatusPass
	} else if phiFailCount == len(refinementResults) {
		phiStatus = model.StatusFail
	} else {
		phiStatus = model.StatusContested
	}

	distinctionPreserved := false
	var reason string

	if trajStatus == model.StatusPass && phiStatus == model.StatusFail {
		distinctionPreserved = true
		reason = "Trajectory is numerically well-formed (valid) despite relational phi-clock nonmonotonicity (invalid)."
	} else if trajStatus == model.StatusFail {
		distinctionPreserved = true
		reason = "Trajectory itself is invalid (contains NaN or node contact), preventing clock evaluation."
	} else {
		distinctionPreserved = true
		reason = fmt.Sprintf("Trajectory validity is %s; phi-clock validity is %s.", trajStatus, phiStatus)
	}

	return TrajectoryValiditySummary{
		TrajectoryValid:      trajStatus,
		PhiClockValid:        phiStatus,
		DistinctionPreserved: distinctionPreserved,
		Reason:               reason,
	}
}
