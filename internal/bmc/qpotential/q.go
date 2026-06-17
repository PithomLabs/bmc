package qpotential

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// Q computes the quantum potential Q at configuration point (alpha, phi).
// Q = -1/(2R) * (∂²R/∂α² - ∂²R/∂φ²)
// The derivatives are computed via central finite difference with step h.
func Q(wf wave.WaveFunction, alpha, phi, h float64) float64 {
	rVal := wave.AmplitudeField(alpha, phi, wf)
	if rVal < 1e-12 {
		// Guard against division by zero near wavefunction nodes.
		return 0.0
	}

	// Second derivative with respect to alpha
	rPlusAlpha := wave.AmplitudeField(alpha+h, phi, wf)
	rMinusAlpha := wave.AmplitudeField(alpha-h, phi, wf)
	d2RDAlpha2 := (rPlusAlpha - 2*rVal + rMinusAlpha) / (h * h)

	// Second derivative with respect to phi
	rPlusPhi := wave.AmplitudeField(alpha, phi+h, wf)
	rMinusPhi := wave.AmplitudeField(alpha, phi-h, wf)
	d2RDPhi2 := (rPlusPhi - 2*rVal + rMinusPhi) / (h * h)

	return -1.0 / (2.0 * rVal) * (d2RDAlpha2 - d2RDPhi2)
}

// QAlongTrajectory evaluates the quantum potential Q at every point on the trajectory.
func QAlongTrajectory(wf wave.WaveFunction, traj model.Trajectory, h float64) []float64 {
	qVals := make([]float64, len(traj.Points))
	for i, p := range traj.Points {
		qVals[i] = Q(wf, p.State.Alpha, p.State.Phi, h)
	}
	return qVals
}

// MaxAbsQ returns the maximum absolute value of Q in a slice of Q values.
func MaxAbsQ(qValues []float64) float64 {
	var maxVal float64
	for _, val := range qValues {
		absVal := math.Abs(val)
		if absVal > maxVal {
			maxVal = absVal
		}
	}
	return maxVal
}

// CheckQuantumPotential verifies if the maximum quantum potential stays within tolerance.
func CheckQuantumPotential(maxQ float64, tolerance float64) model.CheckResult {
	pass := maxQ <= tolerance

	var status model.CheckStatus
	var reason string
	if pass {
		status = model.StatusPass
		reason = "Quantum potential is approximately zero, matching plane-wave control requirements."
	} else {
		status = model.StatusFail
		reason = "Quantum potential exceeds plane-wave control tolerance."
	}

	return model.CheckResult{
		Status:  status,
		Pass:    pass,
		Reason:  reason,
		MaxAbsQ: &maxQ,
	}
}
