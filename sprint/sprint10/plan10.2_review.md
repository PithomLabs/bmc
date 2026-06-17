Approve **Sprint 10.2 implementation plan with revisions**.

This is the right repair direction. It finally attacks the real blocker: the runner must compute from **per-point branch data**, or it must honestly block.

```json id="s10_2_plan_status"
{
  "sprint": "BMC Sprint 10.2",
  "plan_status": "approved_with_revisions",
  "core_rule": "No per-point residual inputs, no computed residual report.",
  "chosen_path": "true per-point candidate residual calculation",
  "allowed_scope": "candidate local-branch residual diagnostics only",
  "forbidden_scope": [
    "Friedmann recovery claim",
    "classical-limit recovery claim",
    "BMC beats null models claim",
    "global cosmology claim",
    "full BMC promotion"
  ]
}
```

## Required revisions before implementation

### 1. Rename the formula carefully

Your proposed formula is fine as a **candidate toy residual diagnostic**, but do not frame it as a Friedmann residual.

Use a guarded formula ID like:

```text id="formula_guarded"
candidate_guidance_velocity_balance_residual_v0.1
```

or:

```text id="formula_guarded_2"
candidate_local_branch_velocity_constraint_residual_v0.1
```

Avoid names like:

```text id="formula_avoid"
friedmann_residual
classical_residual
cosmology_recovery_residual
```

The calculation:

```text id="formula_text"
LHS = (dα/dλ)^2
RHS = (dφ/dλ)^2
residual = LHS - RHS
```

is acceptable only as a **candidate local diagnostic under explicit conventions**, not as a classical cosmology test.

### 2. Validate finite-difference inputs

When computing `dAlpha` and `dPhi`, reject or block if:

```text id="fd_checks"
delta_lambda = 0
delta_lambda is nonfinite
alpha or phi is nonfinite
there are fewer than 2 trajectory points
trajectory points are not ordered or cannot be ordered by lambda/index
```

If `model.TrajectoryPoint` does not actually expose `Lambda`, use an explicit available field, or add a clear index/time parameter. Do not silently assume a clock coordinate exists.

### 3. Make blocked behavior the default for old artifacts

Older `bmc0a_clock_readiness.json` files will not have `points`.

That must produce:

```text id="blocked_old_artifacts"
candidate_residual_computed = false
interpretation_status = blocked_by_missing_residual_inputs
Computed Candidate Residual Diagnostics: 0
```

Do not backfill fake points.

### 4. Do not use `omitempty` for essential audit fields

For computed diagnostics, `residual_input_points` should be present and nonempty.

For blocked diagnostics, it can be empty, but validation must prevent:

```text id="bad_state"
residual_computed = true
residual_input_points = []
```

### 5. Sensitivity test must change residual inputs

This test is critical:

```text id="sensitivity_test"
TestResidualRunMetricsChangeWhenInputBranchDataChanges
```

It must modify actual point values such as `alpha`, `phi`, or the finite-difference-relevant values.

Changing only `clock_range`, branch count, or metadata is not enough.

### 6. Summary should separate computed vs blocked

Your summary repair is required. Use:

```text id="summary_required"
Eligible Local Branches: N
Computed Candidate Residual Diagnostics: N
Blocked Candidate Residual Diagnostics: N
Total Candidate Residual Diagnostics: N
Residual/Null Comparisons: N
```

This prevents another misleading “2 diagnostics” situation.

## Approved core changes

```text id="approved_changes"
1. Add trajectory points to clockseg branch output.
2. Remove hidden residual magnitude constants.
3. Compute residual series from finite differences over real branch points.
4. Add residual input points to the report.
5. Add source branch registry.
6. Reject reported branches not found in source registry.
7. Block when point inputs are absent.
8. Remove synthetic branch injection.
9. Add sensitivity test.
10. Keep all anti-overclaim boundaries.
```

## EBP status

```text id="ebp_s10_2"
Claim: Sprint 10.2 may compute candidate local-branch residual diagnostics from per-point inputs.
Status: allowed if truly input-derived.

Claim: Sprint 10.2 recovers Friedmann behavior.
Status: forbidden.

Claim: Sprint 10.2 validates BMC.
Status: forbidden.

Claim: Sprint 10.2 retires convention debts.
Status: forbidden.

needToyCheck: partial
needNullModel: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
full_bmc_toy_gate: blocked
```

Final instruction:

```text id="go_s10_2"
Go implement Sprint 10.2 with the revisions above. The computed path must derive metrics from per-point residual values. If per-point inputs are missing or invalid, generate a blocked report. No hidden constants, no synthetic source-backed branches, no recovery language.
```
