# Implementation Plan - BMC-POST-0004: BMC-0B Massive Scalar Numerical WdW Specification

Define the finite-difference schema, grid/domain assumptions, boundary conditions, residual gates, failure modes, null-model expectations, and EBP promotion boundaries for a future BMC-0B massive scalar numerical Wheeler-DeWitt profile. This ticket is **specification-only**; no solver or trajectories are implemented.

## User Review Required

> [!IMPORTANT]
> - **Specification-Only**: This ticket does not implement a massive scalar solver, trajectories, or cosmological recovery. It serves strictly as a workbench preparation.
> - **Humble Status Obligations**: All statuses (operator form, boundary conditions, grid, etc.) remain in pending/not-computed states to prevent premature claims.
> - **Forbidden Term Checks**: The validation checker strictly rejects words indicating validation, success, recovery, or proof.

## Proposed Changes

### Component: `internal/bmc/bmc0bspec`

#### [NEW] [spec.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/bmc0bspec/spec.go)
Create a new Go package `internal/bmc/bmc0bspec` defining spec structs, validation logic, and default specifications.

- Define structs:
  ```go
  type GridSpec struct {
  	AlphaMin    *float64 `json:"alpha_min,omitempty"`
  	AlphaMax    *float64 `json:"alpha_max,omitempty"`
  	PhiMin      *float64 `json:"phi_min,omitempty"`
  	PhiMax      *float64 `json:"phi_max,omitempty"`
  	AlphaPoints *int     `json:"alpha_points,omitempty"`
  	PhiPoints   *int     `json:"phi_points,omitempty"`
  	GridStatus  string   `json:"grid_status"`
  }

  type FiniteDifferenceSpec struct {
  	Scheme                 string `json:"scheme"`
  	Order                  string `json:"order"`
  	StencilStatus          string `json:"stencil_status"`
  	BoundaryStencilStatus  string `json:"boundary_stencil_status"`
  	StabilityRequirement   string `json:"stability_requirement"`
  	ConvergenceRequirement string `json:"convergence_requirement"`
  }

  type ResidualGateSpec struct {
  	ResidualNorms                []string `json:"residual_norms"`
  	ToleranceStatus              string   `json:"tolerance_status"`
  	PassGateStatus               string   `json:"pass_gate_status"`
  	MustFailOnNonfinite          bool     `json:"must_fail_on_nonfinite"`
  	MustFailOnBoundaryViolation  bool     `json:"must_fail_on_boundary_violation"`
  	MustFailOnUnreviewedOperator bool     `json:"must_fail_on_unreviewed_operator"`
  }

  type FailureMode struct {
  	ID              string `json:"id"`
  	Description     string `json:"description"`
  	BlocksPromotion bool   `json:"blocks_promotion"`
  }

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
  }
  ```
- Implement `DefaultSpec() *MassiveScalarWdWSpec` containing the default specification under EBP requirements.
- Implement `ValidateSpec(spec *MassiveScalarWdWSpec) []error` which:
  - Enforces `ToyAnalysisOnly = true` and `PhysicsClaim = "minisuperspace_only"`.
  - Verifies presence of all 13 required failure modes.
  - Verifies presence of all 6 required null models.
  - Rejects any forbidden status words (`validated`, `proved`, `recovered`, `physics_success`, `bmc_validated`, `friedmann_recovered`, `quantum_gravity_progress`) case-insensitively across status fields.
  - Enforces operator form, boundary conditions, and stencils remain in unreviewed/pending states (e.g. `pending_faithfulness_review`, `specified_only`, `blocked_until_reviewed`).

#### [NEW] [spec_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/bmc0bspec/spec_test.go)
Implement tests:
- `TestMassiveScalarWdWSpecIsSpecificationOnly`
- `TestMassiveScalarWdWSpecRequiresReviewedOperatorForm`
- `TestMassiveScalarWdWSpecRequiresBoundaryConditions`
- `TestMassiveScalarWdWSpecRequiresResidualGate`
- `TestMassiveScalarWdWSpecRequiresNullModels`
- `TestMassiveScalarWdWSpecBlocksRecoveryClaim`
- `TestMassiveScalarWdWSpecRejectsForbiddenStatuses`
- `TestMassiveScalarWdWSpecDeterministic`

### Component: Documentation

#### [NEW] [bmc_post_0004_bmc0b_massive_scalar_wdw_spec.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0004_bmc0b_massive_scalar_wdw_spec.md)
Document that:
- BMC-POST-0004 is specification-only.
- No solver is implemented, no numerical result is produced, no trajectories are integrated, and no Friedmann recovery or physics validation is claimed.
- The purpose is to prevent jumping from BMC-0A toy controls into BMC-0B numerical claims without first fixing operator conventions, boundary conditions, grid assumptions, residual gates, null models, and faithfulness-review obligations.

## Verification Plan

### Automated Tests
1. **Spec Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/bmc0bspec -v -count=1
   ```
2. **All Go Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./... -count=1
   ```
3. **Lean build verification:**
   ```bash
   cd BMC && /home/chaschel/.elan/bin/lake build
   ```
