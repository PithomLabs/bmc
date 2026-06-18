package qpotential

import (
	"errors"
	"math"

	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// Named constants for quantum-potential domain checks.
const (
	NearNodeAmplitudeFloor     = 1e-8
	QPotentialDerivativeStep   = 1e-4
	QPotentialMagnitudeWarning = 1e8
)

// Status strings.
const (
	StatusQPotentialAuthoritative           = "q_potential_authoritative"
	StatusQPotentialBlockedByNodeContact    = "q_potential_blocked_by_node_contact"
	StatusQPotentialBlockedByNearNodeAmp    = "q_potential_blocked_by_near_node_amplitude"
	StatusQPotentialBlockedByNonfiniteWave  = "q_potential_blocked_by_nonfinite_wave"
	StatusQPotentialBlockedByNonfiniteDeriv = "q_potential_blocked_by_nonfinite_derivative"
	StatusQPotentialBlockedByDomainBoundary = "q_potential_blocked_by_domain_boundary"
	StatusQPotentialAuditOnlyNoPromotion    = "q_potential_audit_only_no_promotion"
)

// QPotentialAuditResult represents the audit outcome for a single point.
type QPotentialAuditResult struct {
	Alpha         float64 `json:"alpha"`
	Phi           float64 `json:"phi"`
	Amplitude     float64 `json:"amplitude"`
	QPotential    float64 `json:"q_potential,omitempty"`
	Status        string  `json:"status"`
	Authoritative bool    `json:"authoritative"`
	Reason        string  `json:"reason,omitempty"`
}

// QPotentialEBPStatus represents EBP metadata for POST-0006.
type QPotentialEBPStatus struct {
	ToyAnalysisOnly         bool   `json:"toy_analysis_only"`
	PhysicsClaim            string `json:"physics_claim"`
	BMC0BImpact             string `json:"bmc0b_impact"`
	FriedmannRecoveryImpact string `json:"friedmann_recovery_impact"`
	PromotionRecommendation string `json:"promotion_recommendation"`
}

// QPotentialAudit holds top-level metadata and audited points.
type QPotentialAudit struct {
	SchemaVersion   string                 `json:"schema_version"`
	FixtureID       string                 `json:"fixture_id"`
	ToyAnalysisOnly bool                   `json:"toy_analysis_only"`
	PhysicsClaim    string                 `json:"physics_claim"`
	AnalysisStatus  string                 `json:"analysis_status"`
	Points          []QPotentialAuditResult `json:"points"`
	EBP             QPotentialEBPStatus    `json:"ebp"`
}

// SamplePoint defines coordinates for a point.
type SamplePoint struct {
	Label string  `json:"label"`
	Alpha float64 `json:"alpha"`
	Phi   float64 `json:"phi"`
}

// AuditPoint evaluates quantum potential at a single point with strict boundary audits.
func AuditPoint(wf wave.WaveFunction, alpha, phi, h float64, fixtureType string) QPotentialAuditResult {
	if math.IsNaN(alpha) || math.IsInf(alpha, 0) || math.IsNaN(phi) || math.IsInf(phi, 0) {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Status:        StatusQPotentialBlockedByDomainBoundary,
			Authoritative: false,
			Reason:        "invalid coordinates",
		}
	}

	rVal := wave.AmplitudeField(alpha, phi, wf)
	if math.IsNaN(rVal) || math.IsInf(rVal, 0) {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Amplitude:     rVal,
			Status:        StatusQPotentialBlockedByNonfiniteWave,
			Authoritative: false,
			Reason:        "nonfinite wavefunction amplitude",
		}
	}

	if rVal == 0.0 || rVal < 1e-12 {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Amplitude:     rVal,
			QPotential:    0.0, // Emitted zero must be flagged non-authoritative
			Status:        StatusQPotentialBlockedByNodeContact,
			Authoritative: false,
			Reason:        "amplitude is zero (node contact)",
		}
	}

	if rVal < NearNodeAmplitudeFloor {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Amplitude:     rVal,
			QPotential:    0.0, // Emitted zero must be flagged non-authoritative
			Status:        StatusQPotentialBlockedByNearNodeAmp,
			Authoritative: false,
			Reason:        "amplitude is below floor",
		}
	}

	// Stencil point checks
	stencilOffsets := []struct{ a, p float64 }{
		{alpha + h, phi},
		{alpha - h, phi},
		{alpha, phi + h},
		{alpha, phi - h},
	}

	for _, st := range stencilOffsets {
		rSt := wave.AmplitudeField(st.a, st.p, wf)
		if math.IsNaN(rSt) || math.IsInf(rSt, 0) {
			return QPotentialAuditResult{
				Alpha:         alpha,
				Phi:           phi,
				Amplitude:     rVal,
				Status:        StatusQPotentialBlockedByNonfiniteWave,
				Authoritative: false,
				Reason:        "nonfinite wavefunction amplitude at stencil point",
			}
		}
		if rSt < NearNodeAmplitudeFloor {
			return QPotentialAuditResult{
				Alpha:         alpha,
				Phi:           phi,
				Amplitude:     rVal,
				Status:        StatusQPotentialBlockedByDomainBoundary,
				Authoritative: false,
				Reason:        "stencil point touches near-node region",
			}
		}
	}

	rPlusAlpha := wave.AmplitudeField(alpha+h, phi, wf)
	rMinusAlpha := wave.AmplitudeField(alpha-h, phi, wf)
	d2RDAlpha2 := (rPlusAlpha - 2*rVal + rMinusAlpha) / (h * h)

	rPlusPhi := wave.AmplitudeField(alpha, phi+h, wf)
	rMinusPhi := wave.AmplitudeField(alpha, phi-h, wf)
	d2RDPhi2 := (rPlusPhi - 2*rVal + rMinusPhi) / (h * h)

	if math.IsNaN(d2RDAlpha2) || math.IsInf(d2RDAlpha2, 0) ||
		math.IsNaN(d2RDPhi2) || math.IsInf(d2RDPhi2, 0) {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Amplitude:     rVal,
			Status:        StatusQPotentialBlockedByNonfiniteDeriv,
			Authoritative: false,
			Reason:        "nonfinite derivative",
		}
	}

	qVal := -1.0 / (2.0 * rVal) * (d2RDAlpha2 - d2RDPhi2)
	if math.IsNaN(qVal) || math.IsInf(qVal, 0) {
		return QPotentialAuditResult{
			Alpha:         alpha,
			Phi:           phi,
			Amplitude:     rVal,
			Status:        StatusQPotentialBlockedByNonfiniteDeriv,
			Authoritative: false,
			Reason:        "nonfinite quantum potential value",
		}
	}

	status := StatusQPotentialAuthoritative
	authoritative := true
	var reason string

	if fixtureType == "superposition" {
		status = StatusQPotentialAuditOnlyNoPromotion
		authoritative = false
		reason = "superposition point is audit only"
	}

	return QPotentialAuditResult{
		Alpha:         alpha,
		Phi:           phi,
		Amplitude:     rVal,
		QPotential:    qVal,
		Status:        status,
		Authoritative: authoritative,
		Reason:        reason,
	}
}

// RunAudit runs a quantum-potential domain audit on sample points.
func RunAudit(fixtureID string, fixtureType string, wf wave.WaveFunction, points []SamplePoint, h float64) (*QPotentialAudit, error) {
	if fixtureID == "" {
		return nil, errors.New("invalid input: missing fixture/profile ID")
	}
	if fixtureType != "plane_wave" && fixtureType != "superposition" && fixtureType != "near_node" {
		return nil, errors.New("invalid input: unknown fixture type")
	}
	if h <= 0 {
		return nil, errors.New("invalid input: nonpositive derivative step")
	}
	if math.IsNaN(h) || math.IsInf(h, 0) {
		return nil, errors.New("invalid input: NaN/Inf derivative step")
	}

	for _, pt := range points {
		if math.IsNaN(pt.Alpha) || math.IsInf(pt.Alpha, 0) ||
			math.IsNaN(pt.Phi) || math.IsInf(pt.Phi, 0) {
			return nil, errors.New("invalid input: NaN/Inf coordinates")
		}
	}

	var results []QPotentialAuditResult
	var totalNodeBlocked, totalNearNodeBlocked, totalNonfiniteBlocked, totalStable, totalAuditOnly int

	for _, pt := range points {
		res := AuditPoint(wf, pt.Alpha, pt.Phi, h, fixtureType)
		results = append(results, res)

		switch res.Status {
		case StatusQPotentialBlockedByNodeContact:
			totalNodeBlocked++
		case StatusQPotentialBlockedByNearNodeAmp, StatusQPotentialBlockedByDomainBoundary:
			totalNearNodeBlocked++
		case StatusQPotentialBlockedByNonfiniteWave, StatusQPotentialBlockedByNonfiniteDeriv:
			totalNonfiniteBlocked++
		case StatusQPotentialAuditOnlyNoPromotion:
			totalAuditOnly++
		case StatusQPotentialAuthoritative:
			totalStable++
		}
	}

	analysisStatus := "stable_for_control_scope"
	if totalNodeBlocked > 0 {
		analysisStatus = "blocked_by_node_contact"
	} else if totalNearNodeBlocked > 0 {
		analysisStatus = "blocked_by_near_node_amplitude"
	} else if totalNonfiniteBlocked > 0 {
		analysisStatus = "blocked_by_nonfinite_wave"
	} else if totalAuditOnly > 0 {
		analysisStatus = "audit_only_no_promotion"
	}

	return &QPotentialAudit{
		SchemaVersion:   "qpotential-near-node-domain-boundary-audit-v0.1",
		FixtureID:       fixtureID,
		ToyAnalysisOnly: true,
		PhysicsClaim:    "none",
		AnalysisStatus:  analysisStatus,
		Points:          results,
		EBP: QPotentialEBPStatus{
			ToyAnalysisOnly:         true,
			PhysicsClaim:            "none",
			BMC0BImpact:             "none",
			FriedmannRecoveryImpact: "none",
			PromotionRecommendation: "do_not_promote",
		},
	}, nil
}
