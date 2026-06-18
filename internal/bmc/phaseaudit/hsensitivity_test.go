package phaseaudit_test

import (
	"encoding/json"
	"math"
	"math/cmplx"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/phaseaudit"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

// perturbedWaveFunction has a phase that oscillates rapidly, causing high finite difference drift.
type perturbedWaveFunction struct{}

func (pw perturbedWaveFunction) Psi(alpha, phi float64) complex128 {
	phase := alpha * math.Sin(100.0*alpha)
	return cmplx.Exp(complex(0, phase))
}

// nodeWaveFunction has a node at (0,0).
type nodeWaveFunction struct{}

func (nw nodeWaveFunction) Psi(alpha, phi float64) complex128 {
	return complex(alpha, phi)
}

func TestHSensitivityPlaneWaveControlStable(t *testing.T) {
	wf := wave.NewPlaneWave(2.0, 3.0)
	points := []phaseaudit.SamplePoint{
		{Label: "plane_wave_control_point_1", Alpha: 0.1, Phi: 0.2},
		{Label: "plane_wave_control_point_2", Alpha: 0.5, Phi: -0.5},
	}
	hLadder := []float64{1e-2, 5e-3, 2.5e-3, 1.25e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("pw_fixture", "plane_wave", wf, points, hLadder)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "stable_for_control_scope" {
		t.Errorf("Expected status 'stable_for_control_scope', got: '%s'", audit.AnalysisStatus)
	}

	for _, pt := range audit.Points {
		if !pt.Authoritative {
			t.Errorf("Expected point %s to be authoritative", pt.Label)
		}
		// Plane wave phase gradient should match exact values k=2, omega=3
		for _, g := range pt.GradientsByH {
			if math.Abs(g.DSdAlpha-2.0) > 1e-3 {
				t.Errorf("Expected dS/dalpha close to 2.0, got %f at h=%f", g.DSdAlpha, g.H)
			}
			if math.Abs(g.DSdPhi-3.0) > 1e-3 {
				t.Errorf("Expected dS/dphi close to 3.0, got %f at h=%f", g.DSdPhi, g.H)
			}
		}
	}
}

func TestHSensitivityDetectsHDriftOnPerturbedPhaseFunction(t *testing.T) {
	wf := perturbedWaveFunction{}
	points := []phaseaudit.SamplePoint{
		{Label: "perturbed_point", Alpha: 0.1, Phi: 0.1},
	}
	hLadder := []float64{1e-2, 5e-3, 2.5e-3, 1.25e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("perturbed_fixture", "superposition", wf, points, hLadder)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "sensitive_to_h" {
		t.Errorf("Expected status 'sensitive_to_h', got: '%s'", audit.AnalysisStatus)
	}

	if audit.Points[0].Authoritative {
		t.Error("Expected perturbed point to be marked non-authoritative due to drift")
	}
}

func TestHSensitivityBlocksNearNodeProbe(t *testing.T) {
	wf := nodeWaveFunction{}
	points := []phaseaudit.SamplePoint{
		{Label: "near_node_point", Alpha: 1e-9, Phi: 1e-9},
	}
	hLadder := []float64{1e-2, 5e-3, 2.5e-3, 1.25e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("node_fixture", "near_node", wf, points, hLadder)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if audit.AnalysisStatus != "blocked_by_node_contact" {
		t.Errorf("Expected status 'blocked_by_node_contact', got: '%s'", audit.AnalysisStatus)
	}

	ptAudit := audit.Points[0]
	if !ptAudit.NodeContactOrNearNodeDetected {
		t.Error("Expected NodeContactOrNearNodeDetected to be true")
	}
	if ptAudit.Authoritative {
		t.Error("Expected near node point to be marked non-authoritative")
	}
}

func TestHSensitivityRejectsInvalidHLadder(t *testing.T) {
	wf := wave.NewPlaneWave(1.0, 1.0)
	points := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}

	tests := []struct {
		name    string
		hLadder []float64
	}{
		{"Empty", []float64{}},
		{"Negative H", []float64{1e-2, -1e-3}},
		{"Non-descending H", []float64{1e-3, 1e-2}},
		{"Duplicate H", []float64{1e-2, 1e-2}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, tc.hLadder)
			if err == nil {
				t.Error("Expected error for invalid h ladder, but got nil")
			}
		})
	}
}

func TestHSensitivityRejectsNonfiniteH(t *testing.T) {
	wf := wave.NewPlaneWave(1.0, 1.0)
	points := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}

	_, err1 := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, []float64{math.NaN()})
	if err1 == nil {
		t.Error("Expected error for NaN h, but got nil")
	}

	_, err2 := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, []float64{math.Inf(1)})
	if err2 == nil {
		t.Error("Expected error for Inf h, but got nil")
	}
}

func TestHSensitivityRejectsNonfinitePoint(t *testing.T) {
	wf := wave.NewPlaneWave(1.0, 1.0)
	hLadder := []float64{1e-2, 5e-3}

	ptsNaN := []phaseaudit.SamplePoint{{Label: "pt", Alpha: math.NaN(), Phi: 0.1}}
	_, err1 := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, ptsNaN, hLadder)
	if err1 == nil {
		t.Error("Expected error for NaN point coordinate, but got nil")
	}

	ptsInf := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: math.Inf(-1)}}
	_, err2 := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, ptsInf, hLadder)
	if err2 == nil {
		t.Error("Expected error for Inf point coordinate, but got nil")
	}
}

func TestHSensitivityDeterministic(t *testing.T) {
	wf := wave.NewPlaneWave(2.0, 3.0)
	points := []phaseaudit.SamplePoint{
		{Label: "pt1", Alpha: 0.1, Phi: 0.2},
	}
	hLadder := []float64{1e-2, 5e-3, 2.5e-3, 1.25e-3}

	audit1, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, hLadder)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	audit2, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, hLadder)
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
		t.Error("HSensitivityAudit serialization is not deterministic")
	}
}

func TestHSensitivityNoPromotionFields(t *testing.T) {
	wf := wave.NewPlaneWave(1.0, 1.0)
	points := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}
	hLadder := []float64{1e-2, 5e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, hLadder)
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

func TestHSensitivityDoesNotRequireCLIOrSchema(t *testing.T) {
	// The audit is run entirely in-memory. We test that it correctly executes
	// without any file output or external dependencies.
	wf := wave.NewPlaneWave(1.0, 1.0)
	points := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}
	hLadder := []float64{1e-2, 5e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, hLadder)
	if err != nil || audit == nil {
		t.Fatalf("Audit execution failed in-memory: %v", err)
	}
}

func TestHSensitivityValidationErrorsAndStatusesArePhraseSafe(t *testing.T) {
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

	wf := wave.NewPlaneWave(1.0, 1.0)
	points := []phaseaudit.SamplePoint{{Label: "pt", Alpha: 0.1, Phi: 0.1}}

	// 1. Check returned validation errors
	_, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, []float64{1e-3, 1e-2})
	if err != nil {
		assertPhraseSafe(err.Error())
	} else {
		t.Error("Expected validation error, got nil")
	}

	// 2. Check output audit status fields
	audit, err := phaseaudit.RunHSensitivityAudit("fixture", "plane_wave", wf, points, []float64{1e-2, 5e-3})
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

// branchCutWaveFunction returns Psi(alpha, phi) = exp(i * (Pi - 0.5 * alpha))
type branchCutWaveFunction struct{}

func (bc branchCutWaveFunction) Psi(alpha, phi float64) complex128 {
	phase := math.Pi - 0.5*alpha
	return cmplx.Exp(complex(0, phase))
}

func TestHSensitivityBranchCutSafeGradientDoesNotArgWrap(t *testing.T) {
	wf := branchCutWaveFunction{}
	// Near alpha = 0, phase is Pi. Shift +h will cross Pi.
	points := []phaseaudit.SamplePoint{
		{Label: "branch_cut_point", Alpha: 0.0, Phi: 0.0},
	}
	hLadder := []float64{1e-2, 5e-3, 2.5e-3, 1.25e-3}

	audit, err := phaseaudit.RunHSensitivityAudit("bc_fixture", "superposition", wf, points, hLadder)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	pt := audit.Points[0]
	if !pt.Authoritative {
		t.Error("Expected branch-cut point to be authoritative (no phase wrap jump)")
	}

	const BranchCutGradientTolerance = 1e-3
	for _, g := range pt.GradientsByH {
		// observed_dS_dalpha should be close to -0.5
		if math.Abs(g.DSdAlpha-(-0.5)) > BranchCutGradientTolerance {
			t.Errorf("Expected dS/dalpha close to -0.5, got %f at h=%f", g.DSdAlpha, g.H)
		}
		// observed_dS_dphi should be close to 0.0
		if math.Abs(g.DSdPhi-0.0) > BranchCutGradientTolerance {
			t.Errorf("Expected dS/dphi close to 0.0, got %f at h=%f", g.DSdPhi, g.H)
		}
	}
}

