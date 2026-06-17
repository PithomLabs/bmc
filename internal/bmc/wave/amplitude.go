package wave

import (
	"math/cmplx"
)

// Amplitude returns the modulus R = |Ψ|.
func Amplitude(psi complex128) float64 {
	return cmplx.Abs(psi)
}

// AmplitudeField evaluates the modulus R = |Ψ| at a given configuration point.
func AmplitudeField(alpha, phi float64, wf WaveFunction) float64 {
	return Amplitude(wf.Psi(alpha, phi))
}
