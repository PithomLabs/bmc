package bmc0bspec

import (
	"fmt"
	"strings"
)

// GridSpec defines the grid and domain parameters for the minisuperspace.
type GridSpec struct {
	AlphaMin    *float64 `json:"alpha_min,omitempty"`
	AlphaMax    *float64 `json:"alpha_max,omitempty"`
	PhiMin      *float64 `json:"phi_min,omitempty"`
	PhiMax      *float64 `json:"phi_max,omitempty"`
	AlphaPoints *int     `json:"alpha_points,omitempty"`
	PhiPoints   *int     `json:"phi_points,omitempty"`
	GridStatus  string   `json:"grid_status"`
}

// FiniteDifferenceSpec defines the stencils, schemes, stability, and convergence requirements.
type FiniteDifferenceSpec struct {
	Scheme                 string `json:"scheme"`
	Order                  string `json:"order"`
	StencilStatus          string `json:"stencil_status"`
	BoundaryStencilStatus  string `json:"boundary_stencil_status"`
	StabilityRequirement   string `json:"stability_requirement"`
	ConvergenceRequirement string `json:"convergence_requirement"`
}

// ResidualGateSpec defines WdW residual norms, tolerances, and pass/fail gate statuses.
type ResidualGateSpec struct {
	ResidualNorms                []string `json:"residual_norms"`
	ToleranceStatus              string   `json:"tolerance_status"`
	PassGateStatus               string   `json:"pass_gate_status"`
	MustFailOnNonfinite          bool     `json:"must_fail_on_nonfinite"`
	MustFailOnBoundaryViolation  bool     `json:"must_fail_on_boundary_violation"`
	MustFailOnUnreviewedOperator bool     `json:"must_fail_on_unreviewed_operator"`
}

// FailureMode represents an explicit risk that blocks promotion.
type FailureMode struct {
	ID              string `json:"id"`
	Description     string `json:"description"`
	BlocksPromotion bool   `json:"blocks_promotion"`
}

// MassiveScalarWdWSpec holds the full BMC-0B specification.
type MassiveScalarWdWSpec struct {
	SchemaVersion               string               `json:"schema_version"`
	ProfileID                   string               `json:"profile_id"`
	ArtifactKind                string               `json:"artifact_kind"`
	ToyAnalysisOnly             bool                 `json:"toy_analysis_only"`
	PhysicsClaim                string               `json:"physics_claim"`
	Variables                   []string             `json:"variables"`
	OperatorFormStatus          string               `json:"operator_form_status"`
	FactorOrderingStatus        string               `json:"factor_ordering_status"`
	UnitsConventionStatus       string               `json:"units_convention_status"`
	BoundaryConditionStatus     string               `json:"boundary_condition_status"`
	GridSpec                    GridSpec             `json:"grid_spec"`
	FiniteDifferenceSpec        FiniteDifferenceSpec `json:"finite_difference_spec"`
	ResidualGateSpec            ResidualGateSpec     `json:"residual_gate_spec"`
	FailureModes                []FailureMode        `json:"failure_modes"`
	RequiredNullModels          []string             `json:"required_null_models"`
	RequiredFaithfulnessReviews []string             `json:"required_faithfulness_reviews"`
	PromotionBoundaries         []string             `json:"promotion_boundaries"`

	// Machine-checkable no-solver flags
	SolverImplemented        bool `json:"solver_implemented"`
	NumericalResultsComputed bool `json:"numerical_results_computed"`
	TrajectoriesIntegrated   bool `json:"trajectories_integrated"`
	RecoveryClaimMade        bool `json:"recovery_claim_made"`
}

var forbiddenWords = []string{
	"validated",
	"proved",
	"recovered",
	"ready",
	"successful",
	"physics_success",
	"bmc_validated",
	"friedmann_recovered",
	"quantum_gravity_progress",
	"bmc beats nulls",
	"full bmc unblocked",
}

// Field-specific whitelists
var operatorFormStatusWhitelist = map[string]bool{
	"pending_faithfulness_review": true,
	"explicit_unpaid":             true,
	"blocked_until_reviewed":      true,
}

var factorOrderingStatusWhitelist = map[string]bool{
	"pending_faithfulness_review": true,
	"explicit_unpaid":             true,
	"required_before_solver":      true,
	"blocked_until_reviewed":      true,
}

var unitsConventionStatusWhitelist = map[string]bool{
	"pending_faithfulness_review": true,
	"explicit_unpaid":             true,
	"required_before_solver":      true,
	"blocked_until_reviewed":      true,
}

var boundaryConditionStatusWhitelist = map[string]bool{
	"explicit_unpaid":         true,
	"required_before_solver":  true,
	"blocked_until_reviewed":  true,
	"blocked_until_specified": true,
}

var gridStatusWhitelist = map[string]bool{
	"required_before_solver":  true,
	"blocked_until_specified": true,
	"explicit_unpaid":         true,
}

var toleranceStatusWhitelist = map[string]bool{
	"explicit_unpaid":           true,
	"required_before_solver":    true,
	"required_before_promotion": true,
	"blocked_until_reviewed":    true,
}

var passGateStatusWhitelist = map[string]bool{
	"explicit_unpaid":           true,
	"required_before_solver":    true,
	"required_before_promotion": true,
	"blocked_until_reviewed":    true,
}

var stencilStatusWhitelist = map[string]bool{
	"explicit_unpaid":        true,
	"required_before_solver": true,
	"blocked_until_reviewed": true,
}

var boundaryStencilStatusWhitelist = map[string]bool{
	"explicit_unpaid":        true,
	"required_before_solver": true,
	"blocked_until_reviewed": true,
}

var stabilityRequirementWhitelist = map[string]bool{
	"courant_friedrichs_lewy_pending": true,
	"explicit_unpaid":                 true,
	"required_before_solver":          true,
	"blocked_until_reviewed":          true,
}

var convergenceRequirementWhitelist = map[string]bool{
	"richardson_extrapolation_pending": true,
	"explicit_unpaid":                  true,
	"required_before_solver":           true,
	"blocked_until_reviewed":           true,
}

// DefaultSpec creates a default, physically humble specification for BMC-0B.
func DefaultSpec() *MassiveScalarWdWSpec {
	return &MassiveScalarWdWSpec{
		SchemaVersion:           "bmc0b-massive-scalar-spec-v0.1",
		ProfileID:               "bmc0b_massive_scalar",
		ArtifactKind:            "massive_scalar_wdw_specification",
		ToyAnalysisOnly:         true,
		PhysicsClaim:            "minisuperspace_only",
		Variables:               []string{"alpha", "phi"},
		OperatorFormStatus:      "pending_faithfulness_review",
		FactorOrderingStatus:    "explicit_unpaid",
		UnitsConventionStatus:   "explicit_unpaid",
		BoundaryConditionStatus: "explicit_unpaid",
		GridSpec: GridSpec{
			GridStatus: "required_before_solver",
		},
		FiniteDifferenceSpec: FiniteDifferenceSpec{
			Scheme:                 "central_finite_difference",
			Order:                  "second_order",
			StencilStatus:          "explicit_unpaid",
			BoundaryStencilStatus:  "explicit_unpaid",
			StabilityRequirement:   "courant_friedrichs_lewy_pending",
			ConvergenceRequirement: "richardson_extrapolation_pending",
		},
		ResidualGateSpec: ResidualGateSpec{
			ResidualNorms:                []string{"l2_residual_norm", "linf_residual_norm", "boundary_residual_norm"},
			ToleranceStatus:              "explicit_unpaid",
			PassGateStatus:               "blocked_until_reviewed",
			MustFailOnNonfinite:          true,
			MustFailOnBoundaryViolation:  true,
			MustFailOnUnreviewedOperator: true,
		},
		FailureModes: []FailureMode{
			{ID: "unreviewed_operator_form", Description: "The massive scalar Wheeler-DeWitt operator form has not undergone peer physical review.", BlocksPromotion: true},
			{ID: "ambiguous_factor_ordering", Description: "The choice of factor ordering of variables is ambiguous and lacks physical consensus.", BlocksPromotion: true},
			{ID: "ambiguous_units", Description: "The unit conventions and scaling factors are not explicitly unified.", BlocksPromotion: true},
			{ID: "missing_boundary_conditions", Description: "Boundary conditions for alpha and phi at grid edges are not defined.", BlocksPromotion: true},
			{ID: "grid_domain_too_small", Description: "The chosen numerical grid range is too small, truncating the wavefunction dynamics.", BlocksPromotion: true},
			{ID: "boundary_artifact_contamination", Description: "Reflection or contamination from boundary stencils corrupts the interior solution.", BlocksPromotion: true},
			{ID: "nonfinite_solution_values", Description: "The numerical solver produces nonfinite values (NaN/Inf) on the grid.", BlocksPromotion: true},
			{ID: "residual_norm_not_defined", Description: "The norm under which the Wheeler-DeWitt residual is measured is not defined.", BlocksPromotion: true},
			{ID: "tolerance_not_justified", Description: "The convergence and pass/fail tolerance has not been numerically justified.", BlocksPromotion: true},
			{ID: "solver_convergence_failed", Description: "The numerical finite-difference solver fails to converge on the grid.", BlocksPromotion: true},
			{ID: "null_model_not_run", Description: "The WdW residual has not been evaluated against registered null models.", BlocksPromotion: true},
			{ID: "faithfulness_review_missing", Description: "The physical faithfulness review for the minisuperspace simplification is missing.", BlocksPromotion: true},
			{ID: "recovery_claim_forbidden", Description: "Any recovery claim of Friedmann classical limit is forbidden in this specification stage.", BlocksPromotion: true},
		},
		RequiredNullModels: []string{
			"zero_potential_control",
			"massless_scalar_limit_control",
			"random_phase_same_amplitude_control",
			"coarse_grid_boundary_artifact_control",
			"wrong_potential_sign_control",
			"random_boundary_condition_control",
		},
		RequiredFaithfulnessReviews: []string{
			"operator_form_review",
			"factor_ordering_review",
			"units_convention_review",
			"boundary_condition_review",
			"minisuperspace_metric_signature_review",
			"residual_norm_tolerance_review",
			"classical_target_recovery_criterion_review",
		},
		PromotionBoundaries: []string{
			"physics_claim_restricted_to_toy_spec_only",
			"no_solver_implemented_for_bmc0b_stage",
		},
		SolverImplemented:        false,
		NumericalResultsComputed: false,
		TrajectoriesIntegrated:   false,
		RecoveryClaimMade:        false,
	}
}

// validateExactSet helper enforces cardinality, uniqueness, presence of all required, and absence of extra items.
func validateExactSet(items []string, required []string, errPrefix string) []error {
	var errs []error

	// 1. Check cardinality
	if len(items) != len(required) {
		errs = append(errs, fmt.Errorf("%s: incorrect cardinality", errPrefix))
	}

	// 2. Check duplicates
	seen := make(map[string]bool)
	for _, item := range items {
		if seen[item] {
			errs = append(errs, fmt.Errorf("%s: duplicate item detected", errPrefix))
		}
		seen[item] = true
	}

	// 3. Check presence of required and absence of extra/unknown items
	reqMap := make(map[string]bool)
	for _, r := range required {
		reqMap[r] = true
		if !seen[r] {
			errs = append(errs, fmt.Errorf("%s: missing required item", errPrefix))
		}
	}

	for _, item := range items {
		if !reqMap[item] {
			errs = append(errs, fmt.Errorf("%s: unknown/extra item detected", errPrefix))
		}
	}

	return errs
}

// ValidateSpec performs EBP correctness checks on the specification.
func ValidateSpec(spec *MassiveScalarWdWSpec) []error {
	var errs []error

	// Helper to add formatting-safe errors (phrase-safe)
	fail := func(fieldName string) {
		errs = append(errs, fmt.Errorf("forbidden term or claim detected in field %s", fieldName))
	}

	// Helper to check whitelisted statuses
	checkStatus := func(val string, whitelist map[string]bool, fieldName string) {
		if !whitelist[val] {
			errs = append(errs, fmt.Errorf("invalid status value in field %s", fieldName))
		}
	}

	// 1. Schema and basic fields (strictly phrase-safe, no echoing of spec.SchemaVersion)
	if spec.SchemaVersion != "bmc0b-massive-scalar-spec-v0.1" {
		errs = append(errs, fmt.Errorf("schema_version: unsupported schema version"))
	}
	if !spec.ToyAnalysisOnly {
		errs = append(errs, fmt.Errorf("toy_analysis_only: must be true for POST-0004"))
	}
	if spec.PhysicsClaim != "minisuperspace_only" {
		errs = append(errs, fmt.Errorf("physics_claim: must be exactly 'minisuperspace_only'"))
	}

	// 2. Machine-checkable no-solver flags
	if spec.SolverImplemented {
		errs = append(errs, fmt.Errorf("solver_implemented: must be false for specification-only stage"))
	}
	if spec.NumericalResultsComputed {
		errs = append(errs, fmt.Errorf("numerical_results_computed: must be false for specification-only stage"))
	}
	if spec.TrajectoriesIntegrated {
		errs = append(errs, fmt.Errorf("trajectories_integrated: must be false for specification-only stage"))
	}
	if spec.RecoveryClaimMade {
		errs = append(errs, fmt.Errorf("recovery_claim_made: must be false for specification-only stage"))
	}

	// 3. Grid validation (obligation vs fake readiness)
	if spec.GridSpec.AlphaMin == nil || spec.GridSpec.AlphaMax == nil ||
		spec.GridSpec.PhiMin == nil || spec.GridSpec.PhiMax == nil ||
		spec.GridSpec.AlphaPoints == nil || spec.GridSpec.PhiPoints == nil {
		// Grid status must clearly state the grid is not yet specified (phrase-safe)
		if spec.GridSpec.GridStatus != "required_before_solver" &&
			spec.GridSpec.GridStatus != "blocked_until_specified" &&
			spec.GridSpec.GridStatus != "explicit_unpaid" {
			errs = append(errs, fmt.Errorf("grid_spec: incomplete grid status must remain blocked before solver work"))
		}
	} else {
		// All grid values are present: validate bounds and points count
		if *spec.GridSpec.AlphaMin >= *spec.GridSpec.AlphaMax {
			errs = append(errs, fmt.Errorf("grid_spec: alpha_min must be less than alpha_max"))
		}
		if *spec.GridSpec.PhiMin >= *spec.GridSpec.PhiMax {
			errs = append(errs, fmt.Errorf("grid_spec: phi_min must be less than phi_max"))
		}
		if *spec.GridSpec.AlphaPoints < 3 {
			errs = append(errs, fmt.Errorf("grid_spec: alpha_points must be at least 3"))
		}
		if *spec.GridSpec.PhiPoints < 3 {
			errs = append(errs, fmt.Errorf("grid_spec: phi_points must be at least 3"))
		}
	}

	// 4. Failure Mode validation (exactly 13 required failure modes)
	requiredFailureModes := []string{
		"unreviewed_operator_form",
		"ambiguous_factor_ordering",
		"ambiguous_units",
		"missing_boundary_conditions",
		"grid_domain_too_small",
		"boundary_artifact_contamination",
		"nonfinite_solution_values",
		"residual_norm_not_defined",
		"tolerance_not_justified",
		"solver_convergence_failed",
		"null_model_not_run",
		"faithfulness_review_missing",
		"recovery_claim_forbidden",
	}

	fmIDs := make([]string, len(spec.FailureModes))
	for i, fm := range spec.FailureModes {
		fmIDs[i] = fm.ID
		if !fm.BlocksPromotion {
			errs = append(errs, fmt.Errorf("failure_modes: blocks_promotion must be true for required failure modes"))
		}
	}

	errs = append(errs, validateExactSet(fmIDs, requiredFailureModes, "failure_modes")...)

	// 5. Null Model validation (exactly 6 required null models)
	requiredNullModels := []string{
		"zero_potential_control",
		"massless_scalar_limit_control",
		"random_phase_same_amplitude_control",
		"coarse_grid_boundary_artifact_control",
		"wrong_potential_sign_control",
		"random_boundary_condition_control",
	}
	errs = append(errs, validateExactSet(spec.RequiredNullModels, requiredNullModels, "required_null_models")...)

	// 6. Faithfulness Review validation (exactly 7 required reviews)
	requiredFaithfulnessReviews := []string{
		"operator_form_review",
		"factor_ordering_review",
		"units_convention_review",
		"boundary_condition_review",
		"minisuperspace_metric_signature_review",
		"residual_norm_tolerance_review",
		"classical_target_recovery_criterion_review",
	}
	errs = append(errs, validateExactSet(spec.RequiredFaithfulnessReviews, requiredFaithfulnessReviews, "required_faithfulness_reviews")...)

	// 7. Residual Gate validation
	if !spec.ResidualGateSpec.MustFailOnNonfinite {
		errs = append(errs, fmt.Errorf("residual_gate_spec.must_fail_on_nonfinite: must be true"))
	}
	if !spec.ResidualGateSpec.MustFailOnBoundaryViolation {
		errs = append(errs, fmt.Errorf("residual_gate_spec.must_fail_on_boundary_violation: must be true"))
	}
	if !spec.ResidualGateSpec.MustFailOnUnreviewedOperator {
		errs = append(errs, fmt.Errorf("residual_gate_spec.must_fail_on_unreviewed_operator: must be true"))
	}

	// Required norm obligations (exactly these three)
	requiredNorms := []string{
		"l2_residual_norm",
		"linf_residual_norm",
		"boundary_residual_norm",
	}
	errs = append(errs, validateExactSet(spec.ResidualGateSpec.ResidualNorms, requiredNorms, "residual_norms")...)

	// 8. Finite difference scheme & order validation
	if spec.FiniteDifferenceSpec.Scheme == "" {
		errs = append(errs, fmt.Errorf("finite_difference_spec.scheme: cannot be empty"))
	}
	if spec.FiniteDifferenceSpec.Order == "" {
		errs = append(errs, fmt.Errorf("finite_difference_spec.order: cannot be empty"))
	}

	// 9. Whitelist Status checks
	checkStatus(spec.OperatorFormStatus, operatorFormStatusWhitelist, "operator_form_status")
	checkStatus(spec.FactorOrderingStatus, factorOrderingStatusWhitelist, "factor_ordering_status")
	checkStatus(spec.UnitsConventionStatus, unitsConventionStatusWhitelist, "units_convention_status")
	checkStatus(spec.BoundaryConditionStatus, boundaryConditionStatusWhitelist, "boundary_condition_status")
	checkStatus(spec.GridSpec.GridStatus, gridStatusWhitelist, "grid_spec.grid_status")
	checkStatus(spec.FiniteDifferenceSpec.StencilStatus, stencilStatusWhitelist, "finite_difference_spec.stencil_status")
	checkStatus(spec.FiniteDifferenceSpec.BoundaryStencilStatus, boundaryStencilStatusWhitelist, "finite_difference_spec.boundary_stencil_status")
	checkStatus(spec.FiniteDifferenceSpec.StabilityRequirement, stabilityRequirementWhitelist, "finite_difference_spec.stability_requirement")
	checkStatus(spec.FiniteDifferenceSpec.ConvergenceRequirement, convergenceRequirementWhitelist, "finite_difference_spec.convergence_requirement")
	checkStatus(spec.ResidualGateSpec.ToleranceStatus, toleranceStatusWhitelist, "residual_gate_spec.tolerance_status")
	checkStatus(spec.ResidualGateSpec.PassGateStatus, passGateStatusWhitelist, "residual_gate_spec.pass_gate_status")

	// 10. Case-insensitive forbidden language scan (phrase-safe errors)
	checkForbidden := func(val string, fieldName string) {
		valLower := strings.ToLower(val)
		for _, w := range forbiddenWords {
			if strings.Contains(valLower, w) {
				fail(fieldName)
				break
			}
		}
	}

	checkForbidden(spec.SchemaVersion, "schema_version")
	checkForbidden(spec.ProfileID, "profile_id")
	checkForbidden(spec.ArtifactKind, "artifact_kind")
	checkForbidden(spec.PhysicsClaim, "physics_claim")
	for i, v := range spec.Variables {
		checkForbidden(v, fmt.Sprintf("variables[%d]", i))
	}
	checkForbidden(spec.OperatorFormStatus, "operator_form_status")
	checkForbidden(spec.FactorOrderingStatus, "factor_ordering_status")
	checkForbidden(spec.UnitsConventionStatus, "units_convention_status")
	checkForbidden(spec.BoundaryConditionStatus, "boundary_condition_status")
	checkForbidden(spec.GridSpec.GridStatus, "grid_spec.grid_status")
	checkForbidden(spec.FiniteDifferenceSpec.Scheme, "finite_difference_spec.scheme")
	checkForbidden(spec.FiniteDifferenceSpec.Order, "finite_difference_spec.order")
	checkForbidden(spec.FiniteDifferenceSpec.StencilStatus, "finite_difference_spec.stencil_status")
	checkForbidden(spec.FiniteDifferenceSpec.BoundaryStencilStatus, "finite_difference_spec.boundary_stencil_status")
	checkForbidden(spec.FiniteDifferenceSpec.StabilityRequirement, "finite_difference_spec.stability_requirement")
	checkForbidden(spec.FiniteDifferenceSpec.ConvergenceRequirement, "finite_difference_spec.convergence_requirement")
	for i, rn := range spec.ResidualGateSpec.ResidualNorms {
		checkForbidden(rn, fmt.Sprintf("residual_norms[%d]", i))
	}
	checkForbidden(spec.ResidualGateSpec.ToleranceStatus, "residual_gate_spec.tolerance_status")
	checkForbidden(spec.ResidualGateSpec.PassGateStatus, "residual_gate_spec.pass_gate_status")

	for i, fm := range spec.FailureModes {
		checkForbidden(fm.ID, fmt.Sprintf("failure_modes[%d].id", i))
		checkForbidden(fm.Description, fmt.Sprintf("failure_modes[%d].description", i))
	}

	for i, nm := range spec.RequiredNullModels {
		checkForbidden(nm, fmt.Sprintf("required_null_models[%d]", i))
	}

	for i, fr := range spec.RequiredFaithfulnessReviews {
		checkForbidden(fr, fmt.Sprintf("required_faithfulness_reviews[%d]", i))
	}

	for i, pb := range spec.PromotionBoundaries {
		checkForbidden(pb, fmt.Sprintf("promotion_boundaries[%d]", i))
	}

	return errs
}
