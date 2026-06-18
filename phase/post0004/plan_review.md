Approve **BMC-POST-0004 implementation plan** with required tightening.

```json id="post0004_plan_review"
{
  "ticket": "BMC-POST-0004",
  "title": "BMC-0B Massive Scalar Numerical WdW Specification",
  "plan_status": "approved_with_required_tightening",
  "scope": "specification_only",
  "solver_allowed": false,
  "trajectory_allowed": false,
  "new_cli_required": false,
  "physics_claim_promoted": false,
  "full_bmc_toy_gate": "blocked"
}
```

## What is good

The plan correctly avoids jumping from BMC-0A toy controls into a nontrivial BMC-0B solver. This is the right next step because it forces the workbench to define operator conventions, grid assumptions, boundary conditions, residual gates, nulls, and faithfulness obligations **before** numerical implementation.

The package name `internal/bmc/bmc0bspec` is good. It makes the boundary clear: this is a spec package, not a solver package.

## Required tightening 1: rename one test

This test name is misleading:

```text id="bad_test_name"
TestMassiveScalarWdWSpecRequiresReviewedOperatorForm
```

The spec should **not** require the operator form to already be reviewed. It should require that the operator-form review debt is explicit and blocks solver/promotion.

Better name:

```text id="better_test_name"
TestMassiveScalarWdWSpecRequiresOperatorFormReviewDebt
```

or:

```text id="better_test_name_2"
TestMassiveScalarWdWSpecBlocksUnreviewedOperatorPromotion
```

## Required tightening 2: validate forbidden terms beyond status fields

Do not scan only status fields. Also scan:

```text id="scan_fields"
PhysicsClaim
FailureMode.ID
FailureMode.Description
RequiredFaithfulnessReviews
PromotionBoundaries
documentation-facing strings if stored in the spec
```

Otherwise a forbidden claim could slip through a description or promotion boundary.

## Required tightening 3: make “no solver” machine-checkable

Add explicit fields or validation rules such as:

```go id="no_solver_fields"
SolverImplemented bool `json:"solver_implemented"`
NumericalResultsComputed bool `json:"numerical_results_computed"`
TrajectoriesIntegrated bool `json:"trajectories_integrated"`
RecoveryClaimMade bool `json:"recovery_claim_made"`
```

Default must be:

```text id="no_solver_defaults"
solver_implemented = false
numerical_results_computed = false
trajectories_integrated = false
recovery_claim_made = false
```

Validation must reject any `true` value for POST-0004.

## Required tightening 4: grid fields must be obligations, not fake readiness

If numeric grid values are included, validate them:

```text id="grid_numeric_validation"
alpha_min < alpha_max
phi_min < phi_max
alpha_points >= 3
phi_points >= 3
```

But if the grid is intentionally unspecified, then `GridStatus` must be something like:

```text id="grid_status"
required_before_solver
```

Do not let a partially empty grid look ready.

## Required tightening 5: required failure modes should mostly block promotion

For the 13 required failure modes, validation should require `BlocksPromotion = true` for all of these especially:

```text id="must_block"
unreviewed_operator_form
ambiguous_factor_ordering
ambiguous_units
missing_boundary_conditions
residual_norm_not_defined
tolerance_not_justified
solver_convergence_failed
null_model_not_run
faithfulness_review_missing
recovery_claim_forbidden
```

If any are non-blocking, require an explicit reason. Simpler: make all 13 block promotion for POST-0004.

## Required tightening 6: avoid “validated” in EBP status

When implementation returns its expected status, use:

```text id="status_terms"
BMC0BStatus: specified_only
OperatorFormDebt: explicit_unpaid
BoundaryConditionDebt: explicit_unpaid
SolverStatus: not_implemented
```

Avoid:

```text id="avoid_terms"
validated
ready
successful
recovered
```

## EBP claim audit

```text id="post0004_ebp_review"
Claim: POST-0004 defines BMC-0B massive scalar WdW obligations.
Status: approved.

Claim: POST-0004 implements a solver.
Status: forbidden.

Claim: POST-0004 computes trajectories or residuals.
Status: forbidden.

Claim: POST-0004 advances Friedmann recovery.
Status: forbidden.

needNontrivialPhysicsCase: targeted_by_spec_only
needFaithfulnessReview: required_before_bmc0b_solver
needNumericalErrorAudit: partial
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

## Final instruction

Proceed with POST-0004 as **specification-only**. Add `bmc0bspec`, spec validation, required failure modes/null models/faithfulness obligations, forbidden-term rejection, no-solver flags, docs, and tests. After implementation, run a focused code review before accepting it cleanly.
