package qpotential_test

import (
	"encoding/json"
	"math"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// nodeWaveFunction has a node at (0,0).
type nodeWaveFunction struct{}

func (nw nodeWaveFunction) Psi(alpha, phi float64) complex128 {
	return complex(alpha, phi)
}

// nonfiniteWaveFunction returns NaN amplitude at (0.1, 0.1).
type nonfiniteWaveFunction struct{}

func (nfw nonfiniteWaveFunction) Psi(alpha, phi float64) complex128 {
	if alpha == 0.1 && phi == 0.1 {
		return complex(math.NaN(), math.NaN())
	}
	return complex(1.0, 0.0)
}

func TestQPotentialPlaneWaveControlAuthoritative(t *testing.T) {
	pw := wave.NewPlaneWave(2.0, 3.0)
	points := []qpotential.SamplePoint{
		{Label: "pw_pt_1", Alpha: 0.1, Phi: 0.2},
		{Label: "pw_pt_2", Alpha: -0.5, Phi: 0.5},
	}
	h := qpotential.QPotentialDerivativeStep

	audit, err := qpotential.RunAudit("pw_fixture", "plane_wave", pw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "stable_for_control_scope" {
		t.Errorf("Expected status 'stable_for_control_scope', got '%s'", audit.AnalysisStatus)
	}

	for _, pt := range audit.Points {
		if !pt.Authoritative {
			t.Errorf("Expected point %f, %f to be authoritative", pt.Alpha, pt.Phi)
		}
		if pt.Status != qpotential.StatusQPotentialAuthoritative {
			t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialAuthoritative, pt.Status)
		}
		if math.Abs(pt.QPotential) > 1e-7 {
			t.Errorf("Expected plane wave Q to be close to 0.0, got %e", pt.QPotential)
		}
	}
}

func TestQPotentialSuperpositionSafeAuditOnly(t *testing.T) {
	sw := wave.NewSuperpositionWave(1.0, 0.0, 1.0, 1.0, 1.0, 0.0, 2.0, 2.0)
	points := []qpotential.SamplePoint{
		{Label: "sw_pt", Alpha: 0.1, Phi: 0.2},
	}
	h := qpotential.QPotentialDerivativeStep

	audit, err := qpotential.RunAudit("sw_fixture", "superposition", sw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "audit_only_no_promotion" {
		t.Errorf("Expected status 'audit_only_no_promotion', got '%s'", audit.AnalysisStatus)
	}

	pt := audit.Points[0]
	if pt.Authoritative {
		t.Error("Expected superposition point to be non-authoritative")
	}
	if pt.Status != qpotential.StatusQPotentialAuditOnlyNoPromotion {
		t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialAuditOnlyNoPromotion, pt.Status)
	}
}

func TestQPotentialNearNodeBlocked(t *testing.T) {
	nw := nodeWaveFunction{}
	// At (1e-9, 1e-9), amplitude is sqrt(2)*1e-9 ≈ 1.4e-9 < 1e-8 (floor)
	points := []qpotential.SamplePoint{
		{Label: "near_node_pt", Alpha: 1e-9, Phi: 1e-9},
	}
	h := qpotential.QPotentialDerivativeStep

	audit, err := qpotential.RunAudit("node_fixture", "near_node", nw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "blocked_by_near_node_amplitude" {
		t.Errorf("Expected status 'blocked_by_near_node_amplitude', got '%s'", audit.AnalysisStatus)
	}

	pt := audit.Points[0]
	if pt.Authoritative {
		t.Error("Expected near-node point to be non-authoritative")
	}
	if pt.Status != qpotential.StatusQPotentialBlockedByNearNodeAmp {
		t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialBlockedByNearNodeAmp, pt.Status)
	}
}

func TestQPotentialDoesNotClampNearNodeToAuthoritativeZero(t *testing.T) {
	nw := nodeWaveFunction{}
	// Exactly at (0,0) amplitude is 0.0
	res := qpotential.AuditPoint(nw, 0.0, 0.0, qpotential.QPotentialDerivativeStep, "near_node")

	if res.Authoritative {
		t.Error("Expected exact node point to be non-authoritative")
	}
	if res.Status != qpotential.StatusQPotentialBlockedByNodeContact {
		t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialBlockedByNodeContact, res.Status)
	}
	if res.QPotential != 0.0 {
		t.Errorf("Expected QPotential output field to be 0.0 (but flagged non-authoritative), got %f", res.QPotential)
	}
}

func TestQPotentialRejectsNonfinitePoint(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0)
	h := qpotential.QPotentialDerivativeStep

	pointsNaN := []qpotential.SamplePoint{{Label: "pt", Alpha: math.NaN(), Phi: 0.1}}
	_, err1 := qpotential.RunAudit("fixture", "plane_wave", pw, pointsNaN, h)
	if err1 == nil {
		t.Error("Expected error for NaN point coordinates, got nil")
	}

	pointsInf := []qpotential.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: math.Inf(1)}}
	_, err2 := qpotential.RunAudit("fixture", "plane_wave", pw, pointsInf, h)
	if err2 == nil {
		t.Error("Expected error for Inf point coordinates, got nil")
	}
}

func TestQPotentialRunAuditRejectsInvalidDerivativeStep(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0)
	points := []qpotential.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}

	_, err1 := qpotential.RunAudit("fixture", "plane_wave", pw, points, 0.0)
	if err1 == nil {
		t.Error("Expected error for derivative step = 0.0, got nil")
	}

	_, err2 := qpotential.RunAudit("fixture", "plane_wave", pw, points, -1e-4)
	if err2 == nil {
		t.Error("Expected error for negative derivative step, got nil")
	}

	_, err3 := qpotential.RunAudit("fixture", "plane_wave", pw, points, math.NaN())
	if err3 == nil {
		t.Error("Expected error for NaN derivative step, got nil")
	}

	_, err4 := qpotential.RunAudit("fixture", "plane_wave", pw, points, math.Inf(1))
	if err4 == nil {
		t.Error("Expected error for Inf derivative step, got nil")
	}
}

func TestQPotentialBlocksNonfiniteDerivative(t *testing.T) {
	nfw := nonfiniteWaveFunction{}
	res := qpotential.AuditPoint(nfw, 0.1, 0.1, qpotential.QPotentialDerivativeStep, "plane_wave")

	if res.Authoritative {
		t.Error("Expected nonfinite wave point to be non-authoritative")
	}
	if res.Status != qpotential.StatusQPotentialBlockedByNonfiniteWave {
		t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialBlockedByNonfiniteWave, res.Status)
	}
}

func TestQPotentialDeterministic(t *testing.T) {
	pw := wave.NewPlaneWave(2.0, 3.0)
	points := []qpotential.SamplePoint{
		{Label: "pt", Alpha: 0.1, Phi: 0.2},
	}
	h := qpotential.QPotentialDerivativeStep

	audit1, err := qpotential.RunAudit("fixture", "plane_wave", pw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	audit2, err := qpotential.RunAudit("fixture", "plane_wave", pw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	j1, err := json.Marshal(audit1)
	if err != nil {
		t.Fatalf("Failed to marshal audit1: %v", err)
	}
	j2, err := json.Marshal(audit2)
	if err != nil {
		t.Fatalf("Failed to marshal audit2: %v", err)
	}

	if string(j1) != string(j2) {
		t.Error("QPotentialAudit serialization is not deterministic")
	}
}

func TestQPotentialNoPromotionFields(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0)
	points := []qpotential.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}
	h := qpotential.QPotentialDerivativeStep

	audit, err := qpotential.RunAudit("fixture", "plane_wave", pw, points, h)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if !audit.ToyAnalysisOnly {
		t.Error("Expected toy_analysis_only to be true")
	}
	if audit.PhysicsClaim != "none" {
		t.Errorf("Expected physics_claim to be 'none', got '%s'", audit.PhysicsClaim)
	}
	if !audit.EBP.ToyAnalysisOnly {
		t.Error("Expected EBP.toy_analysis_only to be true")
	}
	if audit.EBP.PhysicsClaim != "none" {
		t.Errorf("Expected EBP.physics_claim to be 'none', got '%s'", audit.EBP.PhysicsClaim)
	}
	if audit.EBP.BMC0BImpact != "none" {
		t.Errorf("Expected EBP.bmc0b_impact to be 'none', got '%s'", audit.EBP.BMC0BImpact)
	}
	if audit.EBP.FriedmannRecoveryImpact != "none" {
		t.Errorf("Expected EBP.friedmann_recovery_impact to be 'none', got '%s'", audit.EBP.FriedmannRecoveryImpact)
	}
	if audit.EBP.PromotionRecommendation != "do_not_promote" {
		t.Errorf("Expected EBP.promotion_recommendation to be 'do_not_promote', got '%s'", audit.EBP.PromotionRecommendation)
	}
}

func TestQPotentialForbiddenInferenceAudit(t *testing.T) {
	forbiddenWords := []string{
		"validated", "proved", "recovered", "ready", "successful",
		"physics_success", "bmc_validated", "friedmann_recovered",
		"quantum_gravity_progress", "full bmc unblocked", "bmc beats nulls",
	}

	assertPhraseSafe := func(str string) {
		strLower := strings.ToLower(str)
		for _, w := range forbiddenWords {
			if strings.Contains(strLower, w) {
				t.Errorf("Forbidden word '%s' detected in string: '%s'", w, str)
			}
		}
	}

	pw := wave.NewPlaneWave(1.0, 1.0)
	points := []qpotential.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}

	// Check validation error phrase-safety
	_, err := qpotential.RunAudit("fixture", "plane_wave", pw, points, -1e-4)
	if err != nil {
		assertPhraseSafe(err.Error())
	} else {
		t.Error("Expected error, got nil")
	}

	// Check audit fields phrase-safety
	audit, err := qpotential.RunAudit("fixture", "plane_wave", pw, points, qpotential.QPotentialDerivativeStep)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	assertPhraseSafe(audit.SchemaVersion)
	assertPhraseSafe(audit.PhysicsClaim)
	assertPhraseSafe(audit.AnalysisStatus)
	assertPhraseSafe(audit.EBP.PhysicsClaim)
	assertPhraseSafe(audit.EBP.BMC0BImpact)
	assertPhraseSafe(audit.EBP.FriedmannRecoveryImpact)
	assertPhraseSafe(audit.EBP.PromotionRecommendation)
}

// stencilBlockedWaveFunction returns amplitude > NearNodeAmplitudeFloor except at (0.5 + H, 0.5) where it is below it.
type stencilBlockedWaveFunction struct {
	H float64
}

func (sb stencilBlockedWaveFunction) Psi(alpha, phi float64) complex128 {
	// Center is at (0.5, 0.5). Shift alpha + h is (0.5 + h, 0.5)
	if math.Abs(alpha-(0.5+sb.H)) < 1e-9 && math.Abs(phi-0.5) < 1e-9 {
		return complex(0.0, 0.0) // Amplitude 0 < NearNodeAmplitudeFloor
	}
	return complex(1.0, 0.0) // Amplitude 1 > NearNodeAmplitudeFloor
}

func TestQPotentialBlocksStencilPointBelowAmplitudeFloor(t *testing.T) {
	h := qpotential.QPotentialDerivativeStep
	wf := stencilBlockedWaveFunction{H: h}

	// Stencil point check
	res := qpotential.AuditPoint(wf, 0.5, 0.5, h, "plane_wave")

	if res.Authoritative {
		t.Error("Expected stencil point below floor to mark result non-authoritative")
	}
	if res.Status != qpotential.StatusQPotentialBlockedByDomainBoundary {
		t.Errorf("Expected status '%s', got '%s'", qpotential.StatusQPotentialBlockedByDomainBoundary, res.Status)
	}
}

