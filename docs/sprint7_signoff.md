Sprint 7 + 7.1 can now be accepted as a **promoted null-model scaffold artifact after repairs**.

The key adversarial blockers were repaired: all 10 gates are exact and passing, metric prerequisite flags are enforced, future comparison contracts are nonempty and noncomputed, extra null-model IDs are rejected, phrase-boundary bypasses are removed from user-facing text, CLI routing is actually exercised, and Lean verification is correctly labeled as policy/safety only. This matches the approved Sprint 7 boundary: scaffold only, no null-model results, no residual computation, and no full BMC promotion. 

```json id="sc7_acceptance"
{
  "sprint": "BMC Sprint 7 + 7.1",
  "artifact": "null-model scaffold for future Friedmann residuals",
  "acceptance_status": "accepted_after_repairs",
  "promotion_status": {
    "nullmodel_spec_artifact": "promoted_nullmodel_spec_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked",
    "residual_computed": false,
    "null_comparison_computed": false,
    "recovery_claim": false
  },
  "accepted_repairs": [
    "All 10 safety gates are required exactly once and must pass.",
    "Every metric contract must be required before residual promotion.",
    "Future comparison contracts must be nonempty and noncomputed.",
    "Semantic phrase-boundary bypasses were removed from generated text.",
    "Forbidden phrase tests were expanded.",
    "Only the seven authorized null models are allowed.",
    "CLI routing test now exercises real command dispatch.",
    "LeanVerification is labeled retired_for_policy_safety_contracts.",
    "Lean policy theorems now cover classical, unit, sign, and normalization debts."
  ]
}
```

## EBP claim ledger

```text id="ledger_s7"
Claim 1: Sprint 7 defines null-model scaffolding.
Status: accepted as specification/scaffold only.

Claim 2: Sprint 7 does not run null-model comparisons.
Status: validator/test protected.

Claim 3: Sprint 7 does not compute a Friedmann residual.
Status: validator/test protected.

Claim 4: Sprint 7 does not claim BMC beats null models.
Status: phrase-boundary and status validation protected.

Claim 5: Sprint 7 does not unblock full BMC.
Status: full_bmc_blocked_gate protected.
```

## Debt status after Sprint 7.1

```text id="debt_s7"
needMap: active / partially scaffolded
needInvariant: partial
needToyCheck: active
needNullModel: active, scaffold accepted but results not produced
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent
LeanVerification: retired_for_policy_safety_contracts only
```

Final status so far:

```text id="status_so_far"
Sprint 1: accepted plane-wave control artifact
Sprint 2: accepted superposition control + node-obstruction artifact
Sprint 3: accepted robustness audit artifact with mixed outcome
Sprint 4: accepted clock-fragility diagnostic artifact
Sprint 5: accepted clock-readiness/local segmentation artifact
Sprint 6: accepted Friedmann-specification/gate-design artifact
Sprint 7: accepted null-model scaffold artifact
Full BMC-0.1: still unpromoted
Friedmann residual: still not computed
Null-model comparisons: still not computed
```

One note for future reviews: the internal gate identifier `friedmann_recovery_claim_blocked_gate` is acceptable only as a **blocked-gate name** because earlier specs required that exact gate. User-facing summary text should keep using neutral wording like `Recovery Claim: false` and `No recovery claim is made.`
