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
	"physics_success",
	"bmc_validated",
	"friedmann_recovered",
	"quantum_gravity_progress",
	"ready",
	"successful",
	"first ever",
	"scientifically novel",
	"scientifically original",
	"breakthrough",
	"unprecedented",
	"bmc is new physics",
	"bmc proves bohmian cosmology",
	"bmc validates quantum gravity",
	"friedmann recovery",
	"recovery of friedmann",
	"ready for recovery",
	"full bmc validated",
	"problem of time solved",
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
		OperatorFormStatus:      "explicit_unpaid",
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
			ResidualNorms:                []string{"L2", "L_infinity"},
			ToleranceStatus:              "unjustified",
			PassGateStatus:               "blocked",
			MustFailOnNonfinite:          true,
			MustFailOnBoundaryViolation:  true,
			MustFailOnUnreviewedOperator: true,
		},
		FailureModes: []FailureMode{
			{ID: "unreviewed_operator_form", Description: "The massive scalar Wheeler-DeWitt operator form has not undergone peer physical review.", BlocksPromotion: true},
			{ID: "ambiguous_factor_ordering", Description: "The choice of factor ordering of variables is ambiguous and lacks physical consensus.", BlocksPromotion: true},
			{ID: "ambiguous_units", Description: "The unit conventions and scaling factors are not explicitly unified.", BlocksPromotion: true},
			{ID: "missing_boundary_conditions", Description: "Boundary conditions for alpha and phi at grid edges are not defined.", BlocksPromotion: true},
			{ID: "residual_norm_not_defined", Description: "The norm under which the Wheeler-DeWitt residual is measured is not defined.", BlocksPromotion: true},
			{ID: "tolerance_not_justified", Description: "The convergence and pass/fail tolerance has not been numerically justified.", BlocksPromotion: true},
			{ID: "solver_convergence_failed", Description: "The numerical finite-difference solver fails to converge on the grid.", BlocksPromotion: true},
			{ID: "null_model_not_run", Description: "The WdW residual has not been evaluated against registered null models.", BlocksPromotion: true},
			{ID: "faithfulness_review_missing", Description: "The physical faithfulness review for the minisuperspace simplification is missing.", BlocksPromotion: true},
			{ID: "recovery_claim_forbidden", Description: "Any recovery claim of Friedmann classical limit is forbidden in this specification stage.", BlocksPromotion: true},
			{ID: "grid_unspecified", Description: "The numerical grid range and spacing are unspecified.", BlocksPromotion: true},
			{ID: "finite_difference_scheme_unspecified", Description: "The boundary stencils and solver scheme are not fully specified.", BlocksPromotion: true},
			{ID: "boundary_stencil_unstable", Description: "The numerical stencils at the grid boundaries are potentially unstable.", BlocksPromotion: true},
		},
		RequiredNullModels: []string{
			"constant_phase_control",
			"randomized_phase_control",
			"matched_amplitude_randomized_phase_control",
			"classical_frw_reference_trajectory",
			"same_branch_segmentation_under_null_wavefunctions",
			"node_neighborhood_stress_case",
		},
		RequiredFaithfulnessReviews: []string{
			"minisuperspace_truncation_faithfulness_review",
			"operator_ordering_physical_faithfulness_review",
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

// ValidateSpec performs EBP correctness checks on the specification.
func ValidateSpec(spec *MassiveScalarWdWSpec) []error {
	var errs []error

	// Helper to add formatting-safe errors (phrase-safe)
	fail := func(fieldName string) {
		errs = append(errs, fmt.Errorf("forbidden term or claim detected in field %s", fieldName))
	}

	// 1. Schema and basic fields
	if spec.SchemaVersion != "bmc0b-massive-scalar-spec-v0.1" {
		errs = append(errs, fmt.Errorf("schema_version: unsupported schema version: %s", spec.SchemaVersion))
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
	if spec.GridSpec.GridStatus != "required_before_solver" {
		// Enforce validation if not unspecified
		if spec.GridSpec.AlphaMin == nil || spec.GridSpec.AlphaMax == nil ||
			spec.GridSpec.PhiMin == nil || spec.GridSpec.PhiMax == nil ||
			spec.GridSpec.AlphaPoints == nil || spec.GridSpec.PhiPoints == nil {
			errs = append(errs, fmt.Errorf("grid_spec: partially empty grid fields cannot look ready"))
		} else {
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
	}

	// 4. Failure Mode validation
	requiredFailureModes := map[string]bool{
		"unreviewed_operator_form":            true,
		"ambiguous_factor_ordering":           true,
		"ambiguous_units":                     true,
		"missing_boundary_conditions":         true,
		"residual_norm_not_defined":           true,
		"tolerance_not_justified":             true,
		"solver_convergence_failed":           true,
		"null_model_not_run":                  true,
		"faithfulness_review_missing":         true,
		"recovery_claim_forbidden":            true,
		"grid_unspecified":                    true,
		"finite_difference_scheme_unspecified": true,
		"boundary_stencil_unstable":           true,
	}

	foundModes := make(map[string]bool)
	for _, fm := range spec.FailureModes {
		if requiredFailureModes[fm.ID] {
			foundModes[fm.ID] = true
			if !fm.BlocksPromotion {
				errs = append(errs, fmt.Errorf("failure_mode %s: blocks_promotion must be true for POST-0004", fm.ID))
			}
		}
	}

	for fmID := range requiredFailureModes {
		if !foundModes[fmID] {
			errs = append(errs, fmt.Errorf("failure_modes: missing required failure mode: %s", fmID))
		}
	}

	// 5. Null Model validation
	requiredNullModels := map[string]bool{
		"constant_phase_control":                             true,
		"randomized_phase_control":                           true,
		"matched_amplitude_randomized_phase_control":         true,
		"classical_frw_reference_trajectory":                 true,
		"same_branch_segmentation_under_null_wavefunctions":  true,
		"node_neighborhood_stress_case":                      true,
	}

	foundNullModels := make(map[string]bool)
	for _, nm := range spec.RequiredNullModels {
		if requiredNullModels[nm] {
			foundNullModels[nm] = true
		}
	}

	for nmID := range requiredNullModels {
		if !foundNullModels[nmID] {
			errs = append(errs, fmt.Errorf("required_null_models: missing required null model: %s", nmID))
		}
	}

	// 6. Case-insensitive forbidden language scan (phrase-safe errors)
	checkForbidden := func(val string, fieldName string) {
		valLower := strings.ToLower(val)
		for _, w := range forbiddenWords {
			if strings.Contains(valLower, w) {
				fail(fieldName)
				break
			}
		}
	}

	checkForbidden(spec.OperatorFormStatus, "operator_form_status")
	checkForbidden(spec.FactorOrderingStatus, "factor_ordering_status")
	checkForbidden(spec.UnitsConventionStatus, "units_convention_status")
	checkForbidden(spec.BoundaryConditionStatus, "boundary_condition_status")
	checkForbidden(spec.GridSpec.GridStatus, "grid_spec.grid_status")
	checkForbidden(spec.FiniteDifferenceSpec.StencilStatus, "finite_difference_spec.stencil_status")
	checkForbidden(spec.FiniteDifferenceSpec.BoundaryStencilStatus, "finite_difference_spec.boundary_stencil_status")
	checkForbidden(spec.FiniteDifferenceSpec.StabilityRequirement, "finite_difference_spec.stability_requirement")
	checkForbidden(spec.FiniteDifferenceSpec.ConvergenceRequirement, "finite_difference_spec.convergence_requirement")
	checkForbidden(spec.ResidualGateSpec.ToleranceStatus, "residual_gate_spec.tolerance_status")
	checkForbidden(spec.ResidualGateSpec.PassGateStatus, "residual_gate_spec.pass_gate_status")
	checkForbidden(spec.PhysicsClaim, "physics_claim")

	for i, fm := range spec.FailureModes {
		checkForbidden(fm.ID, fmt.Sprintf("failure_modes[%d].id", i))
		checkForbidden(fm.Description, fmt.Sprintf("failure_modes[%d].description", i))
	}

	for i, fr := range spec.RequiredFaithfulnessReviews {
		checkForbidden(fr, fmt.Sprintf("required_faithfulness_reviews[%d]", i))
	}

	for i, pb := range spec.PromotionBoundaries {
		checkForbidden(pb, fmt.Sprintf("promotion_boundaries[%d]", i))
	}

	return errs
}
