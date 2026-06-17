Approve **Sprint 10.3 implementation plan** with one minor tightening.

This plan correctly targets the remaining Sprint 10.2 repair items without expanding scope.

```json id="s10_3_plan_status"
{
  "sprint": "BMC Sprint 10.3",
  "plan_status": "approved_with_minor_tightening",
  "scope": "residualrun comparison and validator hardening",
  "do_not_start": "Sprint 11",
  "current_status": "candidate_residual_runner_candidate_only",
  "maximum_after_review": "promoted_candidate_residual_runner_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked"
}
```

## What is approved

Your plan directly addresses the four remaining issues:

```text id="s10_3_approved"
1. Dynamic comparison targets instead of hard-coded candidate_residual_branch_0.
2. Strong residual_input_points validation.
3. Lambda included for standalone finite-difference audit.
4. Forbidden formula ID rejection.
5. Missing regression tests added.
6. Summary remains computed/blocked/total-aware.
```

This is the correct final hardening pass before accepting Sprint 10.x as a candidate residual-runner artifact.

## One minor tightening

For `TestResidualRunComparisonTargetsActualComputedResidual`, make sure the test does **not** merely inspect a helper function. It should verify the generated comparison record itself.

The test should assert:

```text id="s10_3_test_requirement"
branch_0 residual_computed = false
branch_1 residual_computed = true
residual_null_comparisons[0].target_residual_ids contains candidate_residual_branch_1
residual_null_comparisons[0].target_residual_ids does not contain candidate_residual_branch_0
```

This is the exact regression from the adversarial review.

## Lambda validation note

Your Lambda validation is good. Make sure the validator only requires strictly increasing Lambda **within each computed diagnostic’s residual_input_points**, sorted or listed in point order. It should not compare Lambda across unrelated branches.

## EBP status

```text id="s10_3_ebp"
Claim: Sprint 10.3 repairs comparison target selection.
Status: approved plan.

Claim: Sprint 10.3 strengthens residual input-point validation.
Status: approved plan.

Claim: Sprint 10.3 promotes Friedmann recovery or classical-limit recovery.
Status: forbidden.

Claim: Sprint 10.3 promotes full BMC.
Status: forbidden.

needToyCheck: partial
needNullModel: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
containsFinalTruthClaim: absent
```

Final instruction:

```text id="s10_3_go"
Go implement Sprint 10.3. Keep the target residual IDs dynamically derived from actual computed diagnostics, validate every residual input point including Lambda/provenance/branch alignment, reject forbidden formula IDs phrase-safely, and keep all recovery/superiority/full-BMC claims blocked.
```
