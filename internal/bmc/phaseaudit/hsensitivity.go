package phaseaudit

import (
	"errors"
	"math"
	"math/cmplx"

	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// Named constants for audit thresholds. These are numerical audit benchmarks,
// not physical validation thresholds.
const (
	DefaultHCoarse             = 1e-2
	DefaultHMin                = 1.25e-3
	AbsoluteDriftTolerance     = 1e-6
	RelativeDriftTolerance     = 1e-4
	NearNodeAmplitudeFloor     = 1e-8
	SignFlipTolerance          = 1e-7
	DenominatorSafeFloor       = 1e-5
)

// GradVal holds the computed gradient components for a specific h.
type GradVal struct {
	H        float64 `json:"h"`
	DSdAlpha float64 `json:"ds_dalpha"`
	DSdPhi   float64 `json:"ds_dphi"`
}

// PointAudit represents the h-sensitivity check results for a single sample point.
type PointAudit struct {
	Label                         string    `json:"label"`
	Alpha                         float64   `json:"alpha"`
	Phi                           float64   `json:"phi"`
	GradientsByH                  []GradVal `json:"gradients_by_h"`
	MaxComponentDrift             float64   `json:"max_component_drift"`
	L2DriftBetweenSuccessiveH     []float64 `json:"l2_drift_between_successive_h"`
	LinfDriftBetweenSuccessiveH   []float64 `json:"linf_drift_between_successive_h"`
	SignFlipDetected              bool      `json:"sign_flip_detected"`
	NonfiniteDetected             bool      `json:"non_finite_detected"`
	NodeContactOrNearNodeDetected bool      `json:"node_contact_or_near_node_detected"`
	Authoritative                 bool      `json:"authoritative"`
}

// HSensitivitySummary aggregates the audit results.
type HSensitivitySummary struct {
	TotalPoints            int `json:"total_points"`
	StablePoints           int `json:"stable_points"`
	SensitivePoints        int `json:"sensitive_points"`
	NodeBlockedPoints      int `json:"node_blocked_points"`
	NonfiniteBlockedPoints int `json:"nonfinite_blocked_points"`
}

// HSensitivityEBPStatus represents the EBP status for this audit.
type HSensitivityEBPStatus struct {
	ToyAnalysisOnly         bool   `json:"toy_analysis_only"`
	PhysicsClaim            string `json:"physics_claim"`
	BMC0BImpact             string `json:"bmc0b_impact"`
	FriedmannRecoveryImpact string `json:"friedmann_recovery_impact"`
	PromotionRecommendation string `json:"promotion_recommendation"`
}

// HSensitivityAudit holds the full result of the h-sensitivity audit.
type HSensitivityAudit struct {
	SchemaVersion   string                `json:"schema_version"`
	FixtureID       string                `json:"fixture_id"`
	ToyAnalysisOnly bool                  `json:"toy_analysis_only"`
	PhysicsClaim    string                `json:"physics_claim"`
	AnalysisStatus  string                `json:"analysis_status"`
	HValues         []float64             `json:"h_values"`
	Points          []PointAudit          `json:"points"`
	Summary         HSensitivitySummary   `json:"summary"`
	EBP             HSensitivityEBPStatus `json:"ebp"`
}

// SamplePoint defines a coordinate point to audit.
type SamplePoint struct {
	Label string  `json:"label"`
	Alpha float64 `json:"alpha"`
	Phi   float64 `json:"phi"`
}

// isComplexFinite checks if real and imaginary parts of a complex number are finite.
func isComplexFinite(c complex128) bool {
	r := real(c)
	i := imag(c)
	return !math.IsNaN(r) && !math.IsInf(r, 0) && !math.IsNaN(i) && !math.IsInf(i, 0)
}

// ComputePhaseGradient computes (∂S/∂α, ∂S/∂φ) at (alpha, phi) using central finite difference step size h.
// It uses the Im((1/Ψ) * ∂Ψ/∂x) identity to avoid branch-cut/phase wrapping artifacts.
// This is documented as internal numerical audit machinery, not new physics authority.
func ComputePhaseGradient(alpha, phi float64, wf wave.WaveFunction, h float64) (dSdAlpha, dSdPhi float64, nonfinite bool) {
	psi := wf.Psi(alpha, phi)
	if !isComplexFinite(psi) {
		return 0.0, 0.0, true
	}
	if cmplx.Abs(psi) < NearNodeAmplitudeFloor {
		return 0.0, 0.0, false
	}

	// ∂Ψ/∂α using central finite difference
	psiPlusAlpha := wf.Psi(alpha+h, phi)
	psiMinusAlpha := wf.Psi(alpha-h, phi)
	if !isComplexFinite(psiPlusAlpha) || !isComplexFinite(psiMinusAlpha) {
		return 0.0, 0.0, true
	}
	dPsiDAlpha := (psiPlusAlpha - psiMinusAlpha) / complex(2*h, 0)
	dSdAlpha = imag(dPsiDAlpha / psi)

	// ∂Ψ/∂φ using central finite difference
	psiPlusPhi := wf.Psi(alpha, phi+h)
	psiMinusPhi := wf.Psi(alpha, phi-h)
	if !isComplexFinite(psiPlusPhi) || !isComplexFinite(psiMinusPhi) {
		return 0.0, 0.0, true
	}
	dPsiDPhi := (psiPlusPhi - psiMinusPhi) / complex(2*h, 0)
	dSdPhi = imag(dPsiDPhi / psi)

	nonfinite = math.IsNaN(dSdAlpha) || math.IsInf(dSdAlpha, 0) ||
		math.IsNaN(dSdPhi) || math.IsInf(dSdPhi, 0)

	return dSdAlpha, dSdPhi, nonfinite
}

// RunHSensitivityAudit performs the phase-gradient h-sensitivity audit.
func RunHSensitivityAudit(fixtureID string, fixtureType string, wf wave.WaveFunction, points []SamplePoint, hLadder []float64) (*HSensitivityAudit, error) {
	// 1. Validate fixtureID and fixtureType
	if fixtureID == "" {
		return nil, errors.New("invalid input: missing fixture/profile identifier")
	}
	if fixtureType != "plane_wave" && fixtureType != "superposition" && fixtureType != "near_node" {
		return nil, errors.New("invalid input: unknown fixture type")
	}

	// 2. Validate hLadder
	if len(hLadder) == 0 {
		return nil, errors.New("invalid input: empty h ladder")
	}
	var lastH float64
	for i, h := range hLadder {
		if h <= 0 {
			return nil, errors.New("invalid input: h <= 0")
		}
		if math.IsNaN(h) || math.IsInf(h, 0) {
			return nil, errors.New("invalid input: NaN/Inf h")
		}
		if i > 0 {
			if h >= lastH {
				return nil, errors.New("invalid input: non-descending or duplicate h ladder")
			}
		}
		lastH = h
	}

	// 3. Validate points
	if len(points) == 0 {
		return nil, errors.New("invalid input: empty sample points")
	}
	for _, pt := range points {
		if math.IsNaN(pt.Alpha) || math.IsInf(pt.Alpha, 0) ||
			math.IsNaN(pt.Phi) || math.IsInf(pt.Phi, 0) {
			return nil, errors.New("invalid input: NaN/Inf point coordinates")
		}
	}

	var pointAudits []PointAudit
	var totalStable, totalSensitive, totalNodeBlocked, totalNonfiniteBlocked int

	for _, pt := range points {
		var gradientsByH []GradVal
		var nonfiniteDetected bool
		var nodeContactOrNearNodeDetected bool

		// Check node contact at the evaluation point itself
		psi := wf.Psi(pt.Alpha, pt.Phi)
		if !isComplexFinite(psi) {
			nonfiniteDetected = true
		}
		if cmplx.Abs(psi) < NearNodeAmplitudeFloor {
			nodeContactOrNearNodeDetected = true
		}

		for _, h := range hLadder {
			// Check stencil points for non-finiteness or near-node contact
			for _, shift := range []float64{-h, h} {
				psiAlphaShift := wf.Psi(pt.Alpha+shift, pt.Phi)
				psiPhiShift := wf.Psi(pt.Alpha, pt.Phi+shift)
				if !isComplexFinite(psiAlphaShift) || !isComplexFinite(psiPhiShift) {
					nonfiniteDetected = true
				}
				if cmplx.Abs(psiAlphaShift) < NearNodeAmplitudeFloor || cmplx.Abs(psiPhiShift) < NearNodeAmplitudeFloor {
					nodeContactOrNearNodeDetected = true
				}
			}

			// Compute gradient
			dSdAlpha, dSdPhi, gradNonfinite := ComputePhaseGradient(pt.Alpha, pt.Phi, wf, h)
			if gradNonfinite {
				nonfiniteDetected = true
			}

			gradientsByH = append(gradientsByH, GradVal{
				H:        h,
				DSdAlpha: dSdAlpha,
				DSdPhi:   dSdPhi,
			})
		}

		// Calculate drifts between successive h values
		var l2Drifts []float64
		var linfDrifts []float64
		var maxDrift float64

		for i := 0; i < len(hLadder)-1; i++ {
			g1 := gradientsByH[i]
			g2 := gradientsByH[i+1]

			diffAlpha := g2.DSdAlpha - g1.DSdAlpha
			diffPhi := g2.DSdPhi - g1.DSdPhi

			l2 := math.Sqrt(diffAlpha*diffAlpha + diffPhi*diffPhi)
			linf := math.Max(math.Abs(diffAlpha), math.Abs(diffPhi))

			l2Drifts = append(l2Drifts, l2)
			linfDrifts = append(linfDrifts, linf)

			if l2 > maxDrift {
				maxDrift = l2
			}
		}

		// Check for sign flips across all h values
		var signFlipDetected bool
		for i := 0; i < len(gradientsByH); i++ {
			for j := i + 1; j < len(gradientsByH); j++ {
				g1 := gradientsByH[i]
				g2 := gradientsByH[j]

				if (g1.DSdAlpha > SignFlipTolerance && g2.DSdAlpha < -SignFlipTolerance) ||
					(g1.DSdAlpha < -SignFlipTolerance && g2.DSdAlpha > SignFlipTolerance) {
					signFlipDetected = true
				}
				if (g1.DSdPhi > SignFlipTolerance && g2.DSdPhi < -SignFlipTolerance) ||
					(g1.DSdPhi < -SignFlipTolerance && g2.DSdPhi > SignFlipTolerance) {
					signFlipDetected = true
				}
			}
		}

		driftWithinTolerance := true
		if len(l2Drifts) > 0 {
			for _, l2 := range l2Drifts {
				if l2 > AbsoluteDriftTolerance {
					gMag1 := math.Sqrt(gradientsByH[0].DSdAlpha*gradientsByH[0].DSdAlpha + gradientsByH[0].DSdPhi*gradientsByH[0].DSdPhi)
					if gMag1 > DenominatorSafeFloor {
						relDrift := l2 / gMag1
						if relDrift > RelativeDriftTolerance {
							driftWithinTolerance = false
						}
					} else {
						driftWithinTolerance = false
					}
				}
			}
		}

		authoritative := !nonfiniteDetected && !nodeContactOrNearNodeDetected && !signFlipDetected && driftWithinTolerance

		// Classify points
		if nodeContactOrNearNodeDetected {
			totalNodeBlocked++
		} else if nonfiniteDetected {
			totalNonfiniteBlocked++
		} else if !authoritative {
			totalSensitive++
		} else {
			totalStable++
		}

		pointAudits = append(pointAudits, PointAudit{
			Label:                         pt.Label,
			Alpha:                         pt.Alpha,
			Phi:                           pt.Phi,
			GradientsByH:                  gradientsByH,
			MaxComponentDrift:             maxDrift,
			L2DriftBetweenSuccessiveH:     l2Drifts,
			LinfDriftBetweenSuccessiveH:   linfDrifts,
			SignFlipDetected:              signFlipDetected,
			NonfiniteDetected:             nonfiniteDetected,
			NodeContactOrNearNodeDetected: nodeContactOrNearNodeDetected,
			Authoritative:                 authoritative,
		})
	}

	var analysisStatus string
	if totalNodeBlocked > 0 {
		analysisStatus = "blocked_by_node_contact"
	} else if totalNonfiniteBlocked > 0 {
		analysisStatus = "blocked_by_nonfinite_gradient"
	} else if totalSensitive > 0 {
		analysisStatus = "sensitive_to_h"
	} else {
		analysisStatus = "stable_for_control_scope"
	}

	audit := &HSensitivityAudit{
		SchemaVersion:   "phase-gradient-h-sensitivity-audit-v0.1",
		FixtureID:       fixtureID,
		ToyAnalysisOnly: true,
		PhysicsClaim:    "none",
		AnalysisStatus:  analysisStatus,
		HValues:         hLadder,
		Points:          pointAudits,
		Summary: HSensitivitySummary{
			TotalPoints:            len(points),
			StablePoints:           totalStable,
			SensitivePoints:        totalSensitive,
			NodeBlockedPoints:      totalNodeBlocked,
			NonfiniteBlockedPoints: totalNonfiniteBlocked,
		},
		EBP: HSensitivityEBPStatus{
			ToyAnalysisOnly:         true,
			PhysicsClaim:            "none",
			BMC0BImpact:             "none",
			FriedmannRecoveryImpact: "none",
			PromotionRecommendation: "do_not_promote",
		},
	}

	return audit, nil
}
