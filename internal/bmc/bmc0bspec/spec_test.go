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

func TestMassiveScalarWdWSpecRejectsDuplicateFailureMode(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	// Duplicate the first failure mode
	spec.FailureModes = append(spec.FailureModes, spec.FailureModes[0])
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when failure modes contain duplicates")
	}
}

func TestMassiveScalarWdWSpecRejectsWrongFailureModeCardinality(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	// Remove one failure mode
	spec.FailureModes = spec.FailureModes[1:]
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when failure modes cardinality is incorrect")
	}
}

func TestMassiveScalarWdWSpecRejectsDuplicateNullModel(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.RequiredNullModels = append(spec.RequiredNullModels, spec.RequiredNullModels[0])
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when null models contain duplicates")
	}
}

func TestMassiveScalarWdWSpecRejectsWrongNullModelCardinality(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.RequiredNullModels = spec.RequiredNullModels[1:]
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when null models cardinality is incorrect")
	}
}

func TestMassiveScalarWdWSpecRejectsDuplicateFaithfulnessReview(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.RequiredFaithfulnessReviews = append(spec.RequiredFaithfulnessReviews, spec.RequiredFaithfulnessReviews[0])
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when reviews contain duplicates")
	}
}

func TestMassiveScalarWdWSpecRejectsWrongFaithfulnessReviewCardinality(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.RequiredFaithfulnessReviews = spec.RequiredFaithfulnessReviews[1:]
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when reviews cardinality is incorrect")
	}
}

func TestMassiveScalarWdWSpecRejectsDuplicateResidualNorm(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.ResidualGateSpec.ResidualNorms = append(spec.ResidualGateSpec.ResidualNorms, spec.ResidualGateSpec.ResidualNorms[0])
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when residual norms contain duplicates")
	}
}

func TestMassiveScalarWdWSpecRejectsWrongResidualNormCardinality(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.ResidualGateSpec.ResidualNorms = spec.ResidualGateSpec.ResidualNorms[1:]
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when residual norms cardinality is incorrect")
	}
}

func TestMassiveScalarWdWSpecRejectsNonDebtOperatorStatus(t *testing.T) {
	invalidStatuses := []string{"validated", "specified_only", "not_computed", "pending", "not_started"}
	for _, status := range invalidStatuses {
		t.Run(status, func(t *testing.T) {
			spec := bmc0bspec.DefaultSpec()
			spec.OperatorFormStatus = status
			if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
				t.Errorf("Expected validation to fail when operator form status is non-debt/%s", status)
			}
		})
	}
}

func TestMassiveScalarWdWSpecRejectsNonDebtFactorOrderingStatus(t *testing.T) {
	invalidStatuses := []string{"ready", "specified_only", "not_computed", "pending", "not_started"}
	for _, status := range invalidStatuses {
		t.Run(status, func(t *testing.T) {
			spec := bmc0bspec.DefaultSpec()
			spec.FactorOrderingStatus = status
			if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
				t.Errorf("Expected validation to fail when factor ordering status is non-debt/%s", status)
			}
		})
	}
}

func TestMassiveScalarWdWSpecRejectsNonDebtUnitsStatus(t *testing.T) {
	invalidStatuses := []string{"successful", "specified_only", "not_computed", "pending", "not_started"}
	for _, status := range invalidStatuses {
		t.Run(status, func(t *testing.T) {
			spec := bmc0bspec.DefaultSpec()
			spec.UnitsConventionStatus = status
			if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
				t.Errorf("Expected validation to fail when units status is non-debt/%s", status)
			}
		})
	}
}

func TestMassiveScalarWdWSpecRejectsNonDebtBoundaryStatus(t *testing.T) {
	invalidStatuses := []string{"proved", "specified_only", "not_computed", "pending", "not_started"}
	for _, status := range invalidStatuses {
		t.Run(status, func(t *testing.T) {
			spec := bmc0bspec.DefaultSpec()
			spec.BoundaryConditionStatus = status
			if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
				t.Errorf("Expected validation to fail when boundary condition status is non-debt/%s", status)
			}
		})
	}
}


func TestMassiveScalarWdWSpecValidationErrorsArePhraseSafe(t *testing.T) {
	forbiddenWordsInErrors := []string{"ready", "solver-ready", "successful", "validated", "proved", "recovered"}

	assertPhraseSafe := func(errs []error) {
		for _, err := range errs {
			errStr := strings.ToLower(err.Error())
			for _, word := range forbiddenWordsInErrors {
				if strings.Contains(errStr, word) {
					t.Errorf("Validation error echoed forbidden word '%s': %v", word, err)
				}
			}
		}
	}

	// 1. Test incomplete grid status branch
	spec1 := bmc0bspec.DefaultSpec()
	spec1.GridSpec.GridStatus = "ready" // Forbidden status
	errs1 := bmc0bspec.ValidateSpec(spec1)
	if len(errs1) == 0 {
		t.Error("Expected error on bad grid status")
	}
	assertPhraseSafe(errs1)

	// 2. Test forbidden term scan branch
	spec2 := bmc0bspec.DefaultSpec()
	spec2.SchemaVersion = "bmc0b-spec-validated"
	errs2 := bmc0bspec.ValidateSpec(spec2)
	if len(errs2) == 0 {
		t.Error("Expected error on forbidden term in SchemaVersion")
	}
	assertPhraseSafe(errs2)

	// 3. Test exact-set duplicate branch
	spec3 := bmc0bspec.DefaultSpec()
	spec3.RequiredNullModels = append(spec3.RequiredNullModels, spec3.RequiredNullModels[0])
	errs3 := bmc0bspec.ValidateSpec(spec3)
	if len(errs3) == 0 {
		t.Error("Expected error on duplicate null model")
	}
	assertPhraseSafe(errs3)

	// 4. Test field-specific bad status branch
	spec4 := bmc0bspec.DefaultSpec()
	spec4.OperatorFormStatus = "invalid_status_value"
	errs4 := bmc0bspec.ValidateSpec(spec4)
	if len(errs4) == 0 {
		t.Error("Expected error on non-whitelisted operator form status")
	}
	assertPhraseSafe(errs4)
}

func TestMassiveScalarWdWSpecRejectsForbiddenTermsAcrossStoredStrings(t *testing.T) {
	fieldsToTest := []struct {
		name   string
		mutate func(spec *bmc0bspec.MassiveScalarWdWSpec, term string)
	}{
		{"SchemaVersion", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.SchemaVersion = "bmc0b-" + term }},
		{"ProfileID", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.ProfileID = term }},
		{"ArtifactKind", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.ArtifactKind = term }},
		{"PhysicsClaim", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.PhysicsClaim = term }},
		{"Variables", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.Variables = []string{"alpha", term} }},
		{"GridSpec.GridStatus", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.GridSpec.GridStatus = term }},
		{"FiniteDifferenceSpec.Scheme", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.Scheme = term }},
		{"FiniteDifferenceSpec.Order", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.Order = term }},
		{"FiniteDifferenceSpec.StencilStatus", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.StencilStatus = term }},
		{"FiniteDifferenceSpec.BoundaryStencilStatus", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.BoundaryStencilStatus = term }},
		{"FiniteDifferenceSpec.StabilityRequirement", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.StabilityRequirement = term }},
		{"FiniteDifferenceSpec.ConvergenceRequirement", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FiniteDifferenceSpec.ConvergenceRequirement = term }},
		{"ResidualGateSpec.ResidualNorms", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) {
			spec.ResidualGateSpec.ResidualNorms = []string{"l2_residual_norm", "linf_residual_norm", term}
		}},
		{"ResidualGateSpec.ToleranceStatus", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.ResidualGateSpec.ToleranceStatus = term }},
		{"ResidualGateSpec.PassGateStatus", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.ResidualGateSpec.PassGateStatus = term }},
		{"FailureMode.Description", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.FailureModes[0].Description = term }},
		{"RequiredNullModels", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) {
			spec.RequiredNullModels = []string{"zero_potential_control", "massless_scalar_limit_control", "random_phase_same_amplitude_control", "coarse_grid_boundary_artifact_control", "wrong_potential_sign_control", term}
		}},
		{"RequiredFaithfulnessReviews", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) {
			spec.RequiredFaithfulnessReviews = []string{"operator_form_review", "factor_ordering_review", "units_convention_review", "boundary_condition_review", "minisuperspace_metric_signature_review", "residual_norm_tolerance_review", term}
		}},
		{"PromotionBoundaries", func(spec *bmc0bspec.MassiveScalarWdWSpec, term string) { spec.PromotionBoundaries = []string{"physics_claim_restricted_to_toy_spec_only", term} }},
	}

	forbiddenTerms := []string{
		"validated", "proved", "recovered", "ready", "successful",
		"physics_success", "bmc_validated", "friedmann_recovered",
		"quantum_gravity_progress", "bmc beats nulls", "full bmc unblocked",
	}

	for _, f := range fieldsToTest {
		for _, term := range forbiddenTerms {
			t.Run(f.name+"_"+term, func(t *testing.T) {
				spec := bmc0bspec.DefaultSpec()
				f.mutate(spec, term)

				errs := bmc0bspec.ValidateSpec(spec)
				if len(errs) == 0 {
					t.Fatalf("Expected validation error when field %s contains forbidden term '%s'", f.name, term)
				}

				// Ensure no error messages echo the forbidden term (phrase-safe)
				for _, err := range errs {
					if strings.Contains(strings.ToLower(err.Error()), term) {
						t.Errorf("Validation error for field %s echoed forbidden term '%s': %v", f.name, term, err)
					}
				}
			})
		}
	}
}

func TestMassiveScalarWdWSpecRejectsBadGridBounds(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.GridSpec.GridStatus = "explicit_unpaid"

	alphaMin := 0.0
	alphaMax := 10.0
	phiMin := 0.0
	phiMax := 10.0
	alphaPoints := 10
	phiPoints := 10

	spec.GridSpec.AlphaMin = &alphaMin
	spec.GridSpec.AlphaMax = &alphaMax
	spec.GridSpec.PhiMin = &phiMin
	spec.GridSpec.PhiMax = &phiMax
	spec.GridSpec.AlphaPoints = &alphaPoints
	spec.GridSpec.PhiPoints = &phiPoints

	// 1. Verify it passes with valid grid
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) > 0 {
		t.Fatalf("Expected valid spec with specified grid to pass, got: %v", errs)
	}

	// 2. Alpha bounds bad
	alphaMinBad := 12.0
	spec.GridSpec.AlphaMin = &alphaMinBad
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation error when alpha_min >= alpha_max")
	}

	// 3. Phi bounds bad
	spec.GridSpec.AlphaMin = &alphaMin
	phiMinBad := 12.0
	spec.GridSpec.PhiMin = &phiMinBad
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation error when phi_min >= phi_max")
	}
}

func TestMassiveScalarWdWSpecRejectsTooFewGridPoints(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.GridSpec.GridStatus = "explicit_unpaid"

	alphaMin := 0.0
	alphaMax := 10.0
	phiMin := 0.0
	phiMax := 10.0
	alphaPoints := 2 // Too few
	phiPoints := 10

	spec.GridSpec.AlphaMin = &alphaMin
	spec.GridSpec.AlphaMax = &alphaMax
	spec.GridSpec.PhiMin = &phiMin
	spec.GridSpec.PhiMax = &phiMax
	spec.GridSpec.AlphaPoints = &alphaPoints
	spec.GridSpec.PhiPoints = &phiPoints

	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation error when alpha_points < 3")
	}
}

func TestMassiveScalarWdWSpecBlocksRecoveryClaim(t *testing.T) {
	spec := bmc0bspec.DefaultSpec()
	spec.RecoveryClaimMade = true
	if errs := bmc0bspec.ValidateSpec(spec); len(errs) == 0 {
		t.Error("Expected validation to fail when recovery_claim_made is true")
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
