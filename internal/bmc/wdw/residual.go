package wdw

import (
	"errors"
	"math"
	"math/cmplx"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

const (
	WDWNumericalResidualStep      = 1e-4
	WDWNumericalResidualTolerance = 1e-6

	NumericalResidualPass              = "numerical_residual_pass"
	NumericalResidualViolationDetected = "numerical_residual_violation_detected"
	NumericalResidualNotComputed       = "numerical_residual_not_computed"
	NumericalResidualError             = "numerical_residual_error"

	NumericalAuthorityDiagnostic = "diagnostic_authority"
	NumericalAuthorityOracleOnly = "oracle_control_only"
	NumericalAuthorityNone       = "not_authoritative"
)

// AnalyticResidualPlaneWave returns the exact analytic WdW residual for a plane wave: k² - ω².
// For Sprint 1, this is the primary pass/fail authority.
func AnalyticResidualPlaneWave(k, omega float64) float64 {
	return k*k - omega*omega
}

// FiniteDifferenceResidual computes the WdW residual (-∂²Ψ/∂α² + ∂²Ψ/∂φ²) using central finite differences.
// This is included as a numerical sanity check only.
func FiniteDifferenceResidual(wf wave.WaveFunction, alpha, phi, h float64) complex128 {
	psiVal := wf.Psi(alpha, phi)

	// Second derivative with respect to alpha
	psiPlusAlpha := wf.Psi(alpha+h, phi)
	psiMinusAlpha := wf.Psi(alpha-h, phi)
	d2PsiDAlpha2 := (psiPlusAlpha - complex(2.0, 0.0)*psiVal + psiMinusAlpha) / complex(h*h, 0.0)

	// Second derivative with respect to phi
	psiPlusPhi := wf.Psi(alpha, phi+h)
	psiMinusPhi := wf.Psi(alpha, phi-h)
	d2PsiDPhi2 := (psiPlusPhi - complex(2.0, 0.0)*psiVal + psiMinusPhi) / complex(h*h, 0.0)

	return -d2PsiDAlpha2 + d2PsiDPhi2
}

// MaxAbsFiniteDiffResidual evaluates the finite-difference residual across a slice of states.
func MaxAbsFiniteDiffResidual(wf wave.WaveFunction, points []model.MiniState, h float64) float64 {
	var maxVal float64
	for _, p := range points {
		res := FiniteDifferenceResidual(wf, p.Alpha, p.Phi, h)
		val := cmplx.Abs(res)
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// CheckResidual constructs a CheckResult based on the analytic residual.
func CheckResidual(analyticRes float64, tolerance float64) model.CheckResult {
	absRes := math.Abs(analyticRes)
	pass := absRes <= tolerance

	var status model.CheckStatus
	var reason string
	if pass {
		status = model.StatusPass
		reason = "Analytic Wheeler-DeWitt residual satisfies tolerance."
	} else {
		status = model.StatusFail
		reason = "Wheeler-DeWitt constraint violated: analytic residual exceeds tolerance."
	}

	return model.CheckResult{
		Status:         status,
		Pass:           pass,
		Reason:         reason,
		MaxAbsResidual: &absRes,
	}
}

// ComponentResidual represents the Wheeler-DeWitt constraint residual of a single wave component.
type ComponentResidual struct {
	Component int     `json:"component"`
	K         float64 `json:"k"`
	Omega     float64 `json:"omega"`
	Residual  float64 `json:"residual"`
}

// ComponentResidualsSuperposition computes individual component residuals k_j^2 - omega_j^2.
func ComponentResidualsSuperposition(k1, omega1, k2, omega2 float64) []ComponentResidual {
	return []ComponentResidual{
		{
			Component: 1,
			K:         k1,
			Omega:     omega1,
			Residual:  k1*k1 - omega1*omega1,
		},
		{
			Component: 2,
			K:         k2,
			Omega:     omega2,
			Residual:  k2*k2 - omega2*omega2,
		},
	}
}

// NumericalResidualAt evaluates the WdW residual (-∂²Ψ/∂α² + ∂²Ψ/∂φ²) numerically
// using central finite differences at a specific (alpha, phi) configuration.
// It acts as a toy numerical WdW residual evaluator for failure detection.
func NumericalResidualAt(
	psi func(alpha, phi float64) complex128,
	alpha float64,
	phi float64,
	h float64,
) (complex128, error) {
	if h <= 0 {
		return 0, errors.New("finite difference step h must be strictly positive")
	}

	if math.IsNaN(alpha) || math.IsInf(alpha, 0) ||
		math.IsNaN(phi) || math.IsInf(phi, 0) ||
		math.IsNaN(h) || math.IsInf(h, 0) {
		return 0, errors.New("numerical inputs (alpha, phi, h) must be finite")
	}

	psiVal := psi(alpha, phi)
	psiPlusAlpha := psi(alpha+h, phi)
	psiMinusAlpha := psi(alpha-h, phi)
	psiPlusPhi := psi(alpha, phi+h)
	psiMinusPhi := psi(alpha, phi-h)

	isFinite := func(c complex128) bool {
		r := real(c)
		i := imag(c)
		return !math.IsNaN(r) && !math.IsInf(r, 0) && !math.IsNaN(i) && !math.IsInf(i, 0)
	}

	if !isFinite(psiVal) || !isFinite(psiPlusAlpha) || !isFinite(psiMinusAlpha) || !isFinite(psiPlusPhi) || !isFinite(psiMinusPhi) {
		return 0, errors.New("wavefunction evaluated to non-finite value")
	}

	d2PsiDAlpha2 := (psiPlusAlpha - 2*psiVal + psiMinusAlpha) / complex(h*h, 0.0)
	d2PsiDPhi2 := (psiPlusPhi - 2*psiVal + psiMinusPhi) / complex(h*h, 0.0)

	res := -d2PsiDAlpha2 + d2PsiDPhi2

	if !isFinite(res) {
		return 0, errors.New("numerical residual calculation produced non-finite values")
	}

	return res, nil
}


