package wdw

import (
	"math"
	"math/cmplx"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
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

