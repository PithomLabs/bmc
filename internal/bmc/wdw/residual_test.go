package wdw_test

import (
	"math"
	"math/cmplx"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

// Explicit tolerances for the numerical WdW residual checks.
// - residualStep: chosen as 1e-4 because it balances finite-difference truncation error and floating-point roundoff.
// - validResidualTol: set to 1e-6, which is small enough to ensure the wavefunction satisfies the constraint shell.
// - invalidResidualMin: set to 1e-3, ensuring that physics violations are clearly detectable and well above numerical noise.
const (
	residualStep       = 1e-4
	validResidualTol   = 1e-6
	invalidResidualMin = 1e-3
)

// planeWave returns a plane-wave wavefunction Psi(alpha, phi) = exp(i*(k*alpha + omega*phi))
func planeWave(k, omega float64) func(float64, float64) complex128 {
	return func(alpha, phi float64) complex128 {
		phase := k*alpha + omega*phi
		return cmplx.Exp(complex(0, phase))
	}
}

// TestWDWNumericalResidualAcceptsConstraintShellPlaneWave verifies that a valid plane wave
// satisfying k² = ω² yields a numerical residual within the valid tolerance.
func TestWDWNumericalResidualAcceptsConstraintShellPlaneWave(t *testing.T) {
	// k² = ω² = 4.0
	psi := planeWave(2.0, -2.0)

	res, err := wdw.NumericalResidualAt(psi, 1.0, 2.0, residualStep)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	mag := cmplx.Abs(res)
	if mag > validResidualTol {
		t.Errorf("Expected numerical residual magnitude to satisfy tolerance %e, but got %e", validResidualTol, mag)
	}
}

// TestWDWNumericalResidualRejectsWrongPlaneWaveConstraint verifies that a plane wave violating
// the constraint shell (k² != ω²) is successfully rejected with a residual above the minimum threshold.
func TestWDWNumericalResidualRejectsWrongPlaneWaveConstraint(t *testing.T) {
	// k² = 4.0, ω² = 9.0 -> WdW constraint violated (residual should be close to -5.0)
	psi := planeWave(2.0, 3.0)

	res, err := wdw.NumericalResidualAt(psi, 0.5, 1.0, residualStep)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	mag := cmplx.Abs(res)
	if mag < invalidResidualMin {
		t.Errorf("Expected numerical residual magnitude to exceed %e for invalid plane wave, but got %e", invalidResidualMin, mag)
	}
}

// TestWDWNumericalResidualRejectsInvalidWavefunctionFixture verifies that a wavefunction
// which deviates from a plane wave (by an extra factor) is rejected.
func TestWDWNumericalResidualRejectsInvalidWavefunctionFixture(t *testing.T) {
	// Psi(alpha, phi) = exp(i*(2*alpha + 2*phi)) * (1 + 0.5*alpha²)
	// Even though k² = ω², the alpha² perturbation violates the WdW differential equation.
	psi := func(alpha, phi float64) complex128 {
		base := planeWave(2.0, 2.0)(alpha, phi)
		return base * complex(1.0+0.5*alpha*alpha, 0.0)
	}

	res, err := wdw.NumericalResidualAt(psi, 1.0, 1.0, residualStep)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	mag := cmplx.Abs(res)
	if mag < invalidResidualMin {
		t.Errorf("Expected numerical residual magnitude to exceed %e for invalid wavefunction fixture, but got %e", invalidResidualMin, mag)
	}
}

// TestNumericalResidualRejectsInvalidStep verifies that step h <= 0 is rejected.
func TestNumericalResidualRejectsInvalidStep(t *testing.T) {
	psi := planeWave(1.0, 1.0)

	_, err := wdw.NumericalResidualAt(psi, 0.0, 0.0, 0.0)
	if err == nil {
		t.Error("Expected error for h = 0, got nil")
	}

	_, err = wdw.NumericalResidualAt(psi, 0.0, 0.0, -1e-4)
	if err == nil {
		t.Error("Expected error for negative h, got nil")
	}
}

// TestNumericalResidualRejectsNonfiniteInputs verifies that NaN or Inf values in inputs
// or wavefunction outputs are caught.
func TestNumericalResidualRejectsNonfiniteInputs(t *testing.T) {
	psi := planeWave(1.0, 1.0)

	// Nonfinite alpha
	_, err := wdw.NumericalResidualAt(psi, math.NaN(), 0.0, residualStep)
	if err == nil {
		t.Error("Expected error for NaN alpha, got nil")
	}

	// Nonfinite phi
	_, err = wdw.NumericalResidualAt(psi, 0.0, math.Inf(1), residualStep)
	if err == nil {
		t.Error("Expected error for Inf phi, got nil")
	}

	// Nonfinite h
	_, err = wdw.NumericalResidualAt(psi, 0.0, 0.0, math.NaN())
	if err == nil {
		t.Error("Expected error for NaN h, got nil")
	}

	// Nonfinite wavefunction value
	nanPsi := func(alpha, phi float64) complex128 {
		return complex(math.NaN(), 0.0)
	}
	_, err = wdw.NumericalResidualAt(nanPsi, 0.0, 0.0, residualStep)
	if err == nil {
		t.Error("Expected error for non-finite wavefunction evaluation, got nil")
	}
}
