# BMC-POST-0004: BMC-0B Massive Scalar Numerical WdW Specification

This document records the design decisions and boundary definitions for the Wheeler-DeWitt (WdW) specification of the massive scalar minisuperspace profile (`BMC-0B`).

## 1. Specification-Only Scope

Under EBP 2.1 discipline, this ticket is strictly **specification-only**. It establishes the validation constraints, failure modes, null models, and faithfulness obligations **before** any numerical solver implementation or trajectory integration is attempted.

* BMC-POST-0004 is specification-only.
* No solver is implemented.
* No numerical BMC-0B result is produced.
* No trajectories are integrated.
* No Friedmann recovery is claimed.
* No BMC validation is claimed.
* Full BMC remains blocked.

## 2. Rationale and Prevention of Validation Theater

Jumping directly from simple plane-wave controls (`BMC-0A`) into a massive scalar solver introduces significant risks of validation theater. This specification forces the system to declare all operator conventions, grid requirements, boundary conditions, and null-model benchmarks beforehand.

### Required Failure Modes (13 Total)
All 13 required failure modes block EBP promotion (`BlocksPromotion = true`):
1. `unreviewed_operator_form`: The operator form is unreviewed.
2. `ambiguous_factor_ordering`: The factor ordering lacks physical consensus.
3. `ambiguous_units`: Unit conventions are not unified.
4. `missing_boundary_conditions`: Boundary conditions for alpha and phi at grid edges are not defined.
5. `grid_domain_too_small`: The grid range is too small, truncating the dynamics.
6. `boundary_artifact_contamination`: Contamination from boundary stencils corrupts the interior.
7. `nonfinite_solution_values`: The solver produces nonfinite values (NaN/Inf).
8. `residual_norm_not_defined`: Norm under which residual is measured is undefined.
9. `tolerance_not_justified`: Convergence tolerances are unjustified.
10. `solver_convergence_failed`: Solver convergence fails.
11. `null_model_not_run`: WdW residual has not been evaluated against null models.
12. `faithfulness_review_missing`: Minisuperspace faithfulness review is missing.
13. `recovery_claim_forbidden`: Cosmological recovery claims are forbidden.

### Required Null Models (6 Total)
The spec requires verification against 6 future null-model obligations:
1. `zero_potential_control`
2. `massless_scalar_limit_control`
3. `random_phase_same_amplitude_control`
4. `coarse_grid_boundary_artifact_control`
5. `wrong_potential_sign_control`
6. `random_boundary_condition_control`

## 3. Machine-Checkable Enforcements

The Go validator (`ValidateSpec`) enforces that:
- All no-solver flags (`SolverImplemented`, `NumericalResultsComputed`, `TrajectoriesIntegrated`, `RecoveryClaimMade`) must remain `false`.
- All 13 failure modes must block promotion.
- Forbidden terms (such as `"validated"`, `"proved"`, `"recovered"`, `"ready"`, `"successful"`, `"bmc beats nulls"`, `"full bmc unblocked"`) are rejected case-insensitively across all string fields, status fields, physics claims, failure descriptions, faithfulness reviews, and promotion boundaries.
- Grid validation requires explicit bounds and resolution check if grid status is not `"required_before_solver"`, `"blocked_until_specified"`, or `"explicit_unpaid"`.
