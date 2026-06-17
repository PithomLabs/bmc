package guidance

import (
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// Velocity computes the Bohmian velocity field (dα/dλ, dφ/dλ) at a configuration point.
// If the wavefunction amplitude R = |Ψ| is below the nodeThresh, it refuses the calculation
// and returns (NaN, NaN) to prevent integration or derivative calculations near singularities.
// Guidance equations:
//
//	dα/dλ =  ∂S/∂α
//	dφ/dλ = -∂S/∂φ
func Velocity(wf wave.WaveFunction, alpha, phi, nodeThresh float64) (dAlpha, dPhi float64) {
	rVal := wave.AmplitudeField(alpha, phi, wf)
	if rVal < nodeThresh {
		return math.NaN(), math.NaN()
	}

	dSdAlpha, dSdPhi := wave.PhaseGradient(alpha, phi, wf)
	return dSdAlpha, -dSdPhi
}
