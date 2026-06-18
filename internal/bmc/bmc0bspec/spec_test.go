package bmc0bspec_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/bmc0bspec"
)

func TestMassiveScalarWdWSpecIsSpecificationOnly(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// 1. Verify default validation passes
	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) > 0 {
		t.Fatalf("Default spec should pass validation, but got errors: %v", errs)
	}

	// 2. Verify solver implemented blocks validation
	spec.SolverImplemented = true
	if len(bmc0bspec.ValidateSpec(spec)) == 0 {
		t.Error("Expected validation to fail when solver_implemented is true")
	}

	// 3. Verify numerical results computed blocks validation
	spec = bmc0bspec.DefaultSpec()
	spec.NumericalResultsComputed = true
	if len(bmc0bspec.ValidateSpec(spec)) == 0 {
		t.Error("Expected validation to fail when numerical_results_computed is true")
	}

	// 4. Verify trajectories integrated blocks validation
	spec = bmc0bspec.DefaultSpec()
	spec.TrajectoriesIntegrated = true
	if len(bmc0bspec.ValidateSpec(spec)) == 0 {
		t.Error("Expected validation to fail when trajectories_integrated is true")
	}
}

func TestMassiveScalarWdWSpecBlocksUnreviewedOperatorPromotion(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// Find the unreviewed_operator_form failure mode and change BlocksPromotion to false
	found := false
	for i, fm := range spec.FailureModes {
		if fm.ID == "unreviewed_operator_form" {
			spec.FailureModes[i].BlocksPromotion = false
			found = true
			break
		}
	}

	if !found {
		t.Fatal("Required failure mode 'unreviewed_operator_form' not found in default spec")
	}

	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when 'unreviewed_operator_form' does not block promotion")
	}

	// Remove unreviewed_operator_form failure mode entirely
	spec = bmc0bspec.DefaultSpec()
	var remaining []bmc0bspec.FailureMode
	for _, fm := range spec.FailureModes {
		if fm.ID != "unreviewed_operator_form" {
			remaining = append(remaining, fm)
		}
	}
	spec.FailureModes = remaining

	errs = bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when 'unreviewed_operator_form' is missing")
	}
}

func TestMassiveScalarWdWSpecRequiresBoundaryConditions(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// Verify missing_boundary_conditions failure mode is present and blocks promotion
	found := false
	for _, fm := range spec.FailureModes {
		if fm.ID == "missing_boundary_conditions" {
			found = true
			if !fm.BlocksPromotion {
				t.Error("Expected 'missing_boundary_conditions' to block promotion")
			}
		}
	}
	if !found {
		t.Error("Expected 'missing_boundary_conditions' failure mode to be present in default spec")
	}

	// Set BoundaryConditionStatus to a forbidden term like "validated" or "ready"
	spec.BoundaryConditionStatus = "validated"
	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when BoundaryConditionStatus contains forbidden word")
	}
}

func TestMassiveScalarWdWSpecRequiresResidualGate(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// Verify residual_norm_not_defined is present and blocks promotion
	found := false
	for _, fm := range spec.FailureModes {
		if fm.ID == "residual_norm_not_defined" {
			found = true
			if !fm.BlocksPromotion {
				t.Error("Expected 'residual_norm_not_defined' to block promotion")
			}
		}
	}
	if !found {
		t.Error("Expected 'residual_norm_not_defined' failure mode to be present in default spec")
	}

	// Set ToleranceStatus to a forbidden term like "validated"
	spec.ResidualGateSpec.ToleranceStatus = "validated"
	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when ToleranceStatus contains forbidden word")
	}
}

func TestMassiveScalarWdWSpecRequiresNullModels(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// Remove one required null model and verify validation fails
	if len(spec.RequiredNullModels) == 0 {
		t.Fatal("RequiredNullModels is empty in default spec")
	}
	spec.RequiredNullModels = spec.RequiredNullModels[1:]

	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when one or more required null models are missing")
	}
}

func TestMassiveScalarWdWSpecBlocksRecoveryClaim(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()

	// Verify recovery claim made blocks promotion
	spec.RecoveryClaimMade = true
	errs := bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when recovery_claim_made is true")
	}

	// Verify description containing recovery claim triggers forbidden language check
	spec = bmc0bspec.DefaultSpec()
	spec.OperatorFormStatus = "recovered"
	errs = bmc0bspec.ValidateSpec(spec)
	if len(errs) == 0 {
		t.Error("Expected validation to fail when OperatorFormStatus contains recovered")
	}
}

func TestMassiveScalarWdWSpecRejectsForbiddenStatuses(t *testing.T) {
	forbiddenTerms := []string{
		"validated", "proved", "recovered", "physics_success",
		"bmc_validated", "friedmann_recovered", "quantum_gravity_progress",
		"ready", "successful",
	}

	for _, term := range forbiddenTerms {
		t.Run(term, func(t *testing.T) {
			spec := bmc0bspec.DefaultSpec()
			spec.OperatorFormStatus = "operator_" + term

			errs := bmc0bspec.ValidateSpec(spec)
			if len(errs) == 0 {
				t.Fatalf("Expected validation to fail for term '%s'", term)
			}

			// Verify that error messages are phrase-safe and do NOT contain the forbidden term
			for _, err := range errs {
				if strings.Contains(strings.ToLower(err.Error()), term) {
					t.Errorf("Validation error echoed forbidden term: %v", err)
				}
			}
		})
	}
}

func TestMassiveScalarWdWSpecDeterministic(t *testing.T) {
	spec1 := bmc0bspec.DefaultSpec()
	spec2 := bmc0bspec.DefaultSpec()

	json1, err := json.Marshal(spec1)
	if err != nil {
		t.Fatalf("Failed to marshal spec1: %v", err)
	}

	json2, err := json.Marshal(spec2)
	if err != nil {
		t.Fatalf("Failed to marshal spec2: %v", err)
	}

	if string(json1) != string(json2) {
		t.Error("Default spec JSON serialization is not deterministic")
	}
}
