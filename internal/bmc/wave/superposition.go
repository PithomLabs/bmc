package wave

import (
	"math/cmplx"
)

// SuperpositionWave represents the wavefunction formed by the superposition of two plane waves:
// Ψ(α, φ) = c1 * exp(i(k1*α + ω1*φ)) + c2 * exp(i(k2*α + ω2*φ))
type SuperpositionWave struct {
	C1     complex128
	K1     float64
	Omega1 float64
	C2     complex128
	K2     float64
	Omega2 float64
}

// NewSuperpositionWave constructs a SuperpositionWave from real/imaginary parts of coefficients and parameters.
func NewSuperpositionWave(c1Real, c1Imag, k1, omega1, c2Real, c2Imag, k2, omega2 float64) SuperpositionWave {
	return SuperpositionWave{
		C1:     complex(c1Real, c1Imag),
		K1:     k1,
		Omega1: omega1,
		C2:     complex(c2Real, c2Imag),
		K2:     k2,
		Omega2: omega2,
	}
}

// Psi evaluates the superposition wavefunction at (alpha, phi).
func (sw SuperpositionWave) Psi(alpha, phi float64) complex128 {
	phase1 := sw.K1*alpha + sw.Omega1*phi
	phase2 := sw.K2*alpha + sw.Omega2*phi

	psi1 := sw.C1 * cmplx.Exp(complex(0, phase1))
	psi2 := sw.C2 * cmplx.Exp(complex(0, phase2))

	return psi1 + psi2
}
