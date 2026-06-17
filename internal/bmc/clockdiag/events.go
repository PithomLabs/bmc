package clockdiag

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

type ClockEvent struct {
	Index                  int      `json:"index"`
	Lambda                 float64  `json:"lambda"`
	Alpha                  float64  `json:"alpha"`
	Phi                    float64  `json:"phi"`
	DPhiDLambda            float64  `json:"dphi_dlambda"`
	DAlphaDLambda          float64  `json:"dalpha_dlambda"`
	AmplitudeR             float64  `json:"amplitude_r"`
	QValue                 *float64 `json:"q_value"`
	QStatus                string   `json:"q_status,omitempty"`
	QReason                string   `json:"q_reason,omitempty"`
	PhaseGradientMagnitude *float64 `json:"phase_gradient_magnitude"`
	PhaseGradientStatus    string   `json:"phase_gradient_status,omitempty"`
	PhaseGradientReason    string   `json:"phase_gradient_reason,omitempty"`
	NearNode               bool     `json:"near_node"`
	EventKind              string   `json:"event_kind"`
	Severity               string   `json:"severity"`
}

// DetectClockEvents iterates consecutive trajectory point pairs, computes velocities, and detects clock anomalies.
func DetectClockEvents(traj model.Trajectory, wf wave.WaveFunction, params model.SuperpositionParams, nearZeroDPhiThreshold float64) []ClockEvent {
	if len(traj.Points) < 2 {
		return nil
	}

	// 1. Calculate velocities for all consecutive intervals
	dphis := make([]float64, len(traj.Points))
	dalphas := make([]float64, len(traj.Points))
	hasVel := make([]bool, len(traj.Points))

	for i := 1; i < len(traj.Points); i++ {
		pPrev := traj.Points[i-1]
		pCurr := traj.Points[i]

		if math.IsNaN(pPrev.State.Alpha) || math.IsNaN(pPrev.State.Phi) ||
			math.IsNaN(pCurr.State.Alpha) || math.IsNaN(pCurr.State.Phi) {
			continue
		}

		dt := pCurr.Lambda - pPrev.Lambda
		if dt <= 0 {
			continue
		}

		dphis[i] = (pCurr.State.Phi - pPrev.State.Phi) / dt
		dalphas[i] = (pCurr.State.Alpha - pPrev.State.Alpha) / dt
		hasVel[i] = true
	}

	// 2. Establish dominant sign using first 10 valid steps
	dominantSign := 0.0
	sumDiff := 0.0
	count := 0
	for j := 1; j <= 10 && j < len(traj.Points); j++ {
		if hasVel[j] {
			sumDiff += dphis[j]
			count++
		}
	}
	if count > 0 {
		if sumDiff > 0 {
			dominantSign = 1.0
		} else if sumDiff < 0 {
			dominantSign = -1.0
		}
	}

	var events []ClockEvent

	// 3. Scan for events
	for i := 1; i < len(traj.Points); i++ {
		if !hasVel[i] {
			continue
		}

		pCurr := traj.Points[i]
		dphi := dphis[i]
		dalpha := dalphas[i]

		amplitudeR := wave.AmplitudeField(pCurr.State.Alpha, pCurr.State.Phi, wf)
		nearNode := amplitudeR < params.NodeThresh

		var qValPtr *float64
		var qStatus, qReason string
		var phaseGradPtr *float64
		var phaseGradStatus, phaseGradReason string

		if !nearNode {
			q := qpotential.Q(wf, pCurr.State.Alpha, pCurr.State.Phi, 1e-4)
			qValPtr = &q

			dSdAlpha, dSdPhi := wave.PhaseGradient(pCurr.State.Alpha, pCurr.State.Phi, wf)
			pg := math.Sqrt(dSdAlpha*dSdAlpha + dSdPhi*dSdPhi)
			phaseGradPtr = &pg
		} else {
			qStatus = "not_applicable"
			qReason = "point is near node; quantum potential undefined"
			phaseGradStatus = "not_applicable"
			phaseGradReason = "point is near node; phase gradient undefined"
		}

		// Helper to construct a base event
		createEvent := func(kind, severity string) ClockEvent {
			return ClockEvent{
				Index:                  i,
				Lambda:                 pCurr.Lambda,
				Alpha:                  pCurr.State.Alpha,
				Phi:                    pCurr.State.Phi,
				DPhiDLambda:            dphi,
				DAlphaDLambda:          dalpha,
				AmplitudeR:             amplitudeR,
				QValue:                 qValPtr,
				QStatus:                qStatus,
				QReason:                qReason,
				PhaseGradientMagnitude: phaseGradPtr,
				PhaseGradientStatus:    phaseGradStatus,
				PhaseGradientReason:    phaseGradReason,
				NearNode:               nearNode,
				EventKind:              kind,
				Severity:               severity,
			}
		}

		// Check: near_zero
		if math.Abs(dphi) < nearZeroDPhiThreshold {
			events = append(events, createEvent("near_zero", "info"))
		}

		// Check: sign_change (opposite sign compared to previous step)
		if i > 1 && hasVel[i-1] {
			dphiPrev := dphis[i-1]
			if dphi*dphiPrev < 0 {
				events = append(events, createEvent("sign_change", "warning"))
			}
		}

		// Check: direction_reversal (opposite sign to dominant direction)
		if dominantSign != 0 && dphi*dominantSign < 0 {
			events = append(events, createEvent("direction_reversal", "warning"))
		}

		// Check: monotonicity_failure (opposite or zero step compared to dominant sign/monotonicity rule)
		if dominantSign != 0 {
			isFailure := false
			if dominantSign > 0 && dphi <= 0 {
				isFailure = true
			} else if dominantSign < 0 && dphi >= 0 {
				isFailure = true
			}
			if isFailure {
				events = append(events, createEvent("monotonicity_failure", "diagnostic"))
			}
		}
	}

	return events
}
