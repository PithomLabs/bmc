package friedmannspec

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// 1. TestFriedmannSpecReportValidation
func TestFriedmannSpecReportValidation(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, err := GenerateFriedmannSpecReport(params)
	if err != nil {
		t.Fatalf("failed to generate report: %v", err)
	}

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) > 0 {
		t.Errorf("expected validation to pass, got %d errors: %v", len(errs), errs)
	}
}

// 2. TestFriedmannSpecRejectsResidualComputed
func TestFriedmannSpecRejectsResidualComputed(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.ResidualComputed = true
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when residual_computed is true")
	}
}

// 3. TestFriedmannSpecRejectsRecoveryClaim
func TestFriedmannSpecRejectsRecoveryClaim(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.FriedmannRecoveryClaim = true
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when friedmann_recovery_claim is true")
	}
}

// 4. TestFriedmannSpecRequiresFullBMCBlocked
func TestFriedmannSpecRequiresFullBMCBlocked(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.PromotionGate.Status = model.StatusPass
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when promotion gate is not blocked")
	}
}

// 5. TestFriedmannSpecRequiresClockChoiceDebtActive
func TestFriedmannSpecRequiresClockChoiceDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.EbpDebt.ClockChoiceDebt = "resolved"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when clock_choice_debt is not active")
	}
}

// 6. TestFriedmannSpecRequiresClassicalTargetDebtActive
func TestFriedmannSpecRequiresClassicalTargetDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.EbpDebt.ClassicalTargetDebt = "resolved"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when classical_target_debt is not active")
	}
}

// 7. TestFriedmannSpecRequiresUnitConventionDebtActive
func TestFriedmannSpecRequiresUnitConventionDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.EbpDebt.UnitConventionDebt = "resolved"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when unit_convention_debt is not active")
	}
}

// 8. TestFriedmannSpecRequiresNormalizationDebtActive
func TestFriedmannSpecRequiresNormalizationDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.EbpDebt.NormalizationDebt = "resolved"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when normalization_debt is not active")
	}
}

// 9. TestFriedmannSpecRequiresNullModelDebtActive
func TestFriedmannSpecRequiresNullModelDebtActive(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.EbpDebt.NeedNullModel = "resolved"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when needNullModel is not active")
	}
}

// 10. TestFriedmannSpecRejectsValidatedCandidateMap
func TestFriedmannSpecRejectsValidatedCandidateMap(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.CandidateMaps[0].Status = "validated"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when candidate map is validated")
	}
}

// 11. TestFriedmannSpecRejectsReadyBranchResidualReadiness
func TestFriedmannSpecRejectsReadyBranchResidualReadiness(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.BranchRequirements[0].BranchResidualReadiness = "ready"
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when branch residual readiness is ready")
	}
}

// 12. TestFriedmannSpecRejectsEmptyNullModelRequirements
func TestFriedmannSpecRejectsEmptyNullModelRequirements(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	r.NullModelRequirements = nil
	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when null model requirements are empty")
	}
}

// 13. TestFriedmannSpecRejectsUnknownFields
func TestFriedmannSpecRejectsUnknownFields(t *testing.T) {
	jsonData := `{"schema_version": "bmc0a-friedmann-spec-v0.1", "extra_bad_field": true}`
	dec := json.NewDecoder(bytes.NewReader([]byte(jsonData)))
	dec.DisallowUnknownFields()

	var r FriedmannSpecReport
	err := dec.Decode(&r)
	if err == nil {
		t.Error("expected strict unmarshal to reject unknown fields")
	}
}

// 14. TestFriedmannSpecDeterministicJSON
func TestFriedmannSpecDeterministicJSON(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r1, _ := GenerateFriedmannSpecReport(params)
	r2, _ := GenerateFriedmannSpecReport(params)

	b1, _ := json.MarshalIndent(r1, "", "  ")
	b2, _ := json.MarshalIndent(r2, "", "  ")

	if !bytes.Equal(b1, b2) {
		t.Error("expected deterministic JSON outputs to be byte-identical")
	}
}

// 15. TestFriedmannSpecForbiddenWords
func TestFriedmannSpecForbiddenWords(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	forbiddenPhrases := []string{
		"ready for Friedmann recovery",
		"recovers Friedmann",
		"Friedmann residual passes",
		"classical cosmology recovered",
		"full BMC validated",
	}

	// 1. Scan full JSON serialization of the report to catch any forbidden phrases in strings.
	b, err := json.Marshal(r)
	if err != nil {
		t.Fatalf("failed to marshal report: %v", err)
	}
	reportStr := strings.ToLower(string(b))
	for _, p := range forbiddenPhrases {
		if strings.Contains(reportStr, strings.ToLower(p)) {
			t.Errorf("found forbidden phrase in generated report JSON: %q", p)
		}
	}

	// 2. Scan the printed summary output to ensure no forbidden phrases are in stdout.
	oldStdout := os.Stdout
	rPipe, wPipe, err := os.Pipe()
	if err != nil {
		t.Fatalf("failed to create pipe: %v", err)
	}
	os.Stdout = wPipe

	SummarizeFriedmannSpecReport(r)

	wPipe.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = io.Copy(&buf, rPipe)
	summaryStr := strings.ToLower(buf.String())
	for _, p := range forbiddenPhrases {
		if strings.Contains(summaryStr, strings.ToLower(p)) {
			t.Errorf("found forbidden phrase in summary output: %q", p)
		}
	}
}

// 16. TestFriedmannSpecRequiresFullBMCBlockedGate
func TestFriedmannSpecRequiresFullBMCBlockedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) > 0 {
		t.Fatalf("expected default report to pass, got: %v", errs)
	}

	foundGate := false
	for _, g := range r.Gates {
		if g.Name == "full_bmc_blocked_gate" {
			foundGate = true
			if g.Status != "pass" {
				t.Error("expected full_bmc_blocked_gate to be pass")
			}
		}
	}
	if !foundGate {
		t.Error("expected full_bmc_blocked_gate to exist")
	}
}

// 17. TestFriedmannSpecRejectsMissingFullBMCBlockedGate
func TestFriedmannSpecRejectsMissingFullBMCBlockedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	// Remove full_bmc_blocked_gate
	var newGates []FriedmannSpecGate
	for _, g := range r.Gates {
		if g.Name != "full_bmc_blocked_gate" {
			newGates = append(newGates, g)
		}
	}
	r.Gates = newGates

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when full_bmc_blocked_gate is missing")
	}
}

// 18. TestFriedmannSpecRejectsNonPassFullBMCBlockedGate
func TestFriedmannSpecRejectsNonPassFullBMCBlockedGate(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	// Change status of full_bmc_blocked_gate to blocked
	for idx, g := range r.Gates {
		if g.Name == "full_bmc_blocked_gate" {
			r.Gates[idx].Status = "blocked"
		}
	}

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when full_bmc_blocked_gate is not pass")
	}
}

// 19. TestFriedmannSpecRequiresNullModelsBeforeResidualPromotion
func TestFriedmannSpecRequiresNullModelsBeforeResidualPromotion(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	// Set required_before_residual_promotion to false on one of them
	r.NullModelRequirements[0].RequiredBeforeResidualPromotion = false

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when required_before_residual_promotion is false")
	}
}

// 20. TestFriedmannSpecCandidateMapRequiresClassicalTargetDebt
func TestFriedmannSpecCandidateMapRequiresClassicalTargetDebt(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	// Set ClassicalTargetDebt to empty
	r.CandidateMaps[0].ClassicalTargetDebt = ""

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when classical_target_debt is empty")
	}
}

// 21. TestFriedmannSpecRejectsRetiredClassicalTargetDebt
func TestFriedmannSpecRejectsRetiredClassicalTargetDebt(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)

	// Set ClassicalTargetDebt to retired
	r.CandidateMaps[0].ClassicalTargetDebt = "retired"

	errs := ValidateFriedmannSpecReport(r)
	if len(errs) == 0 {
		t.Error("expected validation to fail when classical_target_debt is retired")
	}
}

// 22. TestFriedmannSpecRejectsTrailingJSONTokens
func TestFriedmannSpecRejectsTrailingJSONTokens(t *testing.T) {
	params := model.DefaultSuperpositionSafeParams()
	r, _ := GenerateFriedmannSpecReport(params)
	b, _ := json.Marshal(r)

	tempDir := t.TempDir()

	// Write valid JSON to temp file
	tmpFile := filepath.Join(tempDir, "tmp_test_spec_valid.json")
	_ = os.WriteFile(tmpFile, b, 0644)

	_, err := ReadFriedmannSpecReport(tmpFile)
	if err != nil {
		t.Fatalf("expected ReadFriedmannSpecReport to succeed, got: %v", err)
	}

	// Trailing garbage JSON
	garbageFile := filepath.Join(tempDir, "tmp_test_spec_garbage.json")
	garbageBytes := append(b, []byte(" }")...)
	_ = os.WriteFile(garbageFile, garbageBytes, 0644)

	_, err = ReadFriedmannSpecReport(garbageFile)
	if err == nil {
		t.Error("expected ReadFriedmannSpecReport to fail with trailing garbage")
	} else if !strings.Contains(err.Error(), "trailing garbage") {
		t.Errorf("expected error message to contain 'trailing garbage', got: %v", err)
	}
}

