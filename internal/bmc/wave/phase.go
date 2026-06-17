package wave

import (
	"math/cmplx"
)

// Phase returns the wrapped phase S = arg(Ψ) in range [-pi, pi].
func Phase(psi complex128) float64 {
	return cmplx.Phase(psi)
}

// DefaultPhaseFDStepSize is the default step size used for numerical finite difference phase gradients.
const DefaultPhaseFDStepSize = 1e-6

// PhaseGradient computes (∂S/∂α, ∂S/∂φ) at (alpha, phi).
// It uses the analytic identity ∂S/∂x = Im((1/Ψ) * ∂Ψ/∂x) to avoid phase-wrapping/branch-cut issues.
// Derivatives of Ψ are approximated using central finite differences with step h.
func PhaseGradient(alpha, phi float64, wf WaveFunction) (dSdAlpha, dSdPhi float64) {
	psi := wf.Psi(alpha, phi)
	if cmplx.Abs(psi) < 1e-12 {
		// Avoid division by zero at nodes
		return 0.0, 0.0
	}

	h := DefaultPhaseFDStepSize

	// Partial derivative with respect to alpha
	psiPlusAlpha := wf.Psi(alpha+h, phi)
	psiMinusAlpha := wf.Psi(alpha-h, phi)
	dPsiDAlpha := (psiPlusAlpha - psiMinusAlpha) / complex(2*h, 0)
	dSdAlpha = imag(dPsiDAlpha / psi)

	// Partial derivative with respect to phi
	psiPlusPhi := wf.Psi(alpha, phi+h)
	psiMinusPhi := wf.Psi(alpha, phi-h)
	dPsiDPhi := (psiPlusPhi - psiMinusPhi) / complex(2*h, 0)
	dSdPhi = imag(dPsiDPhi / psi)

	return dSdAlpha, dSdPhi
}
