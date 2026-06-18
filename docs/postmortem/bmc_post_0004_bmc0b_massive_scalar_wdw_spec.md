# BMC-POST-0004: BMC-0B Massive Scalar Numerical WdW Specification

This document records the design decisions and boundary definitions for the Wheeler-DeWitt (WdW) specification of the massive scalar minisuperspace profile (`BMC-0B`).

## 1. Specification-Only Scope

Under EBP 2.1 discipline, this ticket is strictly **specification-only**. It establishes the validation constraints, failure modes, null models, and faithfulness obligations **before** any numerical solver implementation or trajectory integration is attempted.

- **No Solver Implemented**: The code does not contain a numerical finite-difference solver.
- **No Cosmological/Trajectory Computation**: No wavefunctions are solved, no quantum trajectories are integrated, and no numerical residuals are computed.
- **No Physical Claims**: No classical-limit recovery, Friedmann emergence, or quantum cosmology validation is claimed.

## 2. Rational and Prevention of Validation Theater

Jumping directly from simple plane-wave controls (`BMC-0A`) into a massive scalar solver introduces significant risks of validation theater. This specification forces the system to declare all operator conventions, grid requirements, boundary conditions, and null-model benchmarks beforehand.

### Required Failure Modes (13 Total)
All 13 required failure modes block EBP promotion (`BlocksPromotion = true`):
1. `unreviewed_operator_form`: The operator form is unreviewed.
2. `ambiguous_factor_ordering`: The factor ordering lacks physical consensus.
3. `ambiguous_units`: Unit conventions are not unified.
4. `missing_boundary_conditions`: Alpha/Phi boundary conditions are undefined.
5. `residual_norm_not_defined`: Norm under which residual is measured is undefined.
6. `tolerance_not_justified`: Convergence tolerances are unjustified.
7. `solver_convergence_failed`: Solver convergence fails.
8. `null_model_not_run`: WdW residual has not been evaluated against null models.
9. `faithfulness_review_missing`: Minisuperspace faithfulness review is missing.
10. `recovery_claim_forbidden`: Cosmological recovery claims are forbidden.
11. `grid_unspecified`: Grid resolution or range is unspecified.
12. `finite_difference_scheme_unspecified`: Discretization stencil/scheme is unspecified.
13. `boundary_stencil_unstable`: Grid edge stencils are potentially unstable.

### Required Null Models (6 Total)
The spec requires verification against 6 comparison null models before residual promotion:
1. `constant_phase_control`
2. `randomized_phase_control`
3. `matched_amplitude_randomized_phase_control`
4. `classical_frw_reference_trajectory`
5. `same_branch_segmentation_under_null_wavefunctions`
6. `node_neighborhood_stress_case`

## 3. Machine-Checkable Enforcements

The Go validator (`ValidateSpec`) enforces that:
- All no-solver flags (`SolverImplemented`, `NumericalResultsComputed`, `TrajectoriesIntegrated`, `RecoveryClaimMade`) must remain `false`.
- All 13 failure modes must block promotion.
- Forbidden terms (such as `"validated"`, `"proved"`, `"recovered"`, `"ready"`, `"successful"`) are rejected case-insensitively across all string fields, status fields, physics claims, failure descriptions, faithfulness reviews, and promotion boundaries.
- Grid validation requires explicit bounds and resolution check if grid status is not `required_before_solver`.
