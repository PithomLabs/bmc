package clockdiag

import (
	"fmt"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

type CorrelationSummary struct {
	ParameterSet          string   `json:"parameter_set"`
	C2Real                float64  `json:"c2_real"`
	K2                    float64  `json:"k2"`
	Omega2                float64  `json:"omega2"`
	ClockMonotonic        bool     `json:"clock_monotonic"`
	NumClockEvents        int      `json:"num_clock_events"`
	MinAmplitudeR         *float64 `json:"min_amplitude_r"`
	MinAmplitudeRStatus   string   `json:"min_amplitude_r_status,omitempty"`
	MinAmplitudeRReason   string   `json:"min_amplitude_r_reason,omitempty"`
	MaxAbsQ               *float64 `json:"max_abs_q"`
	MaxAbsQStatus         string   `json:"max_abs_q_status,omitempty"`
	MaxAbsQReason         string   `json:"max_abs_q_reason,omitempty"`
	MaxPhaseGradMagnitude *float64 `json:"max_phase_gradient_magnitude"`
	MaxPhaseGradStatus    string   `json:"max_phase_gradient_status,omitempty"`
	MaxPhaseGradReason    string   `json:"max_phase_gradient_reason,omitempty"`
	MinDistToNodeThresh   *float64 `json:"min_dist_to_node_thresh"`
	MinDistToNodeStatus   string   `json:"min_dist_to_node_status,omitempty"`
	MinDistToNodeReason   string   `json:"min_dist_to_node_reason,omitempty"`
}

// ComputeCorrelations runs all 9 superposition configs and computes correlation metrics.
func ComputeCorrelations(safeParams model.SuperpositionParams, nearZeroDPhiThreshold float64) ([]CorrelationSummary, error) {
	c2Vals := []float64{0.45, 0.50, 0.55}
	k2Vals := []float64{1.9, 2.0, 2.1}

	summaries := make([]CorrelationSummary, 0, len(c2Vals)*len(k2Vals))

	for _, c2Real := range c2Vals {
		for _, k2 := range k2Vals {
			p := safeParams
			p.C2Real = c2Real
			p.K2 = k2
			p.Omega2 = -k2

			// 1. Initialize wavefunction
			sw := wave.NewSuperpositionWave(
				p.C1Real, p.C1Imag, p.K1, p.Omega1,
				p.C2Real, p.C2Imag, p.K2, p.Omega2,
			)

			// 2. Integrate trajectory
			stepper := guidance.NewRK4Stepper()
			initialState := model.MiniState{Alpha: p.Alpha0, Phi: p.Phi0}
			traj := guidance.Integrate(sw, initialState, stepper, p.LambdaStep, p.Steps, p.NodeThresh)

			// 3. Monotonicity check
			phiMonotonic := true
			if len(traj.Points) <= 1 {
				phiMonotonic = false
			} else {
				phiInc := true
				phiDec := true
				for i := 1; i < len(traj.Points); i++ {
					pPrev := traj.Points[i-1]
					pCurr := traj.Points[i]
					if math.IsNaN(pPrev.State.Phi) || math.IsNaN(pCurr.State.Phi) {
						continue
					}
					diff := pCurr.State.Phi - pPrev.State.Phi
					if diff <= 0 {
						phiInc = false
					}
					if diff >= 0 {
						phiDec = false
					}
				}
				phiMonotonic = phiInc || phiDec
			}

			// 4. Detect clock events
			events := DetectClockEvents(traj, sw, p, nearZeroDPhiThreshold)

			// 5. Compute metrics
			var minR *float64
			var minRStatus, minRReason string
			var minDist *float64
			var minDistStatus, minDistReason string
			var maxQ *float64
			var maxQStatus, maxQReason string
			var maxPG *float64
			var maxPGStatus, maxPGReason string

			var rList []float64
			var distList []float64
			var qList []float64
			var pgList []float64

			for _, pt := range traj.Points {
				if math.IsNaN(pt.State.Alpha) || math.IsNaN(pt.State.Phi) {
					continue
				}

				rVal := wave.AmplitudeField(pt.State.Alpha, pt.State.Phi, sw)
				rList = append(rList, rVal)
				distList = append(distList, rVal-p.NodeThresh)

				if rVal >= p.NodeThresh {
					qVal := qpotential.Q(sw, pt.State.Alpha, pt.State.Phi, 1e-4)
					if !math.IsNaN(qVal) && !math.IsInf(qVal, 0) {
						qList = append(qList, qVal)
					}

					dSdAlpha, dSdPhi := wave.PhaseGradient(pt.State.Alpha, pt.State.Phi, sw)
					pgVal := math.Sqrt(dSdAlpha*dSdAlpha + dSdPhi*dSdPhi)
					if !math.IsNaN(pgVal) && !math.IsInf(pgVal, 0) {
						pgList = append(pgList, pgVal)
					}
				}
			}

			if len(rList) > 0 {
				minVal := rList[0]
				for _, r := range rList {
					if r < minVal {
						minVal = r
					}
				}
				minR = &minVal

				minD := distList[0]
				for _, d := range distList {
					if d < minD {
						minD = d
					}
				}
				minDist = &minD
			} else {
				minRStatus = "not_applicable"
				minRReason = "trajectory contains no valid points"
				minDistStatus = "not_applicable"
				minDistReason = "trajectory contains no valid points"
			}

			if len(qList) > 0 {
				maxVal := math.Abs(qList[0])
				for _, q := range qList {
					if math.Abs(q) > maxVal {
						maxVal = math.Abs(q)
					}
				}
				maxQ = &maxVal
			} else {
				maxQStatus = "not_applicable"
				maxQReason = "no points away from node to evaluate quantum potential"
			}

			if len(pgList) > 0 {
				maxVal := pgList[0]
				for _, pg := range pgList {
					if pg > maxVal {
						maxVal = pg
					}
				}
				maxPG = &maxVal
			} else {
				maxPGStatus = "not_applicable"
				maxPGReason = "no points away from node to evaluate phase gradient"
			}

			summaries = append(summaries, CorrelationSummary{
				ParameterSet:          fmt.Sprintf("c2=%.2f, k2=%.1f, omega2=%.1f", c2Real, k2, -k2),
				C2Real:                c2Real,
				K2:                    k2,
				Omega2:                -k2,
				ClockMonotonic:        phiMonotonic,
				NumClockEvents:        len(events),
				MinAmplitudeR:         minR,
				MinAmplitudeRStatus:   minRStatus,
				MinAmplitudeRReason:   minRReason,
				MaxAbsQ:               maxQ,
				MaxAbsQStatus:         maxQStatus,
				MaxAbsQReason:         maxQReason,
				MaxPhaseGradMagnitude: maxPG,
				MaxPhaseGradStatus:    maxPGStatus,
				MaxPhaseGradReason:    maxPGReason,
				MinDistToNodeThresh:   minDist,
				MinDistToNodeStatus:   minDistStatus,
				MinDistToNodeReason:   minDistReason,
			})
		}
	}

	return summaries, nil
}
