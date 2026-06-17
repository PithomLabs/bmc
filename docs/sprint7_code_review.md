Sprint 7 is **implementation-complete and ready for adversarial review**, but **not accepted yet**.

Based on the walkthrough, it appears aligned with the approved Sprint 7 scope: a **null-model scaffold only**, with no residual computation, no null-model comparison results, no claim that BMC beats null models, and full BMC still blocked. That matches the planning boundary: Sprint 7 should define null-model specs, inputs, metric contracts, comparison contracts, gates, validation, CLI routing, and Lean policy contracts only. 

## Current EBP status

```json id="30tkhc"
{
  "sprint": "BMC Sprint 7",
  "artifact": "null-model scaffold for future Friedmann residuals",
  "implementation_status": "complete_reported",
  "acceptance_status": "pending_adversarial_review",
  "maximum_future_promotion": "promoted_nullmodel_spec_artifact_after_repairs",
  "residual_computed": false,
  "null_comparison_computed": false,
  "friedmann_recovery_claim": false,
  "full_bmc_toy_gate": "blocked",
  "needNullModel": "active",
  "clock_choice_debt": "active",
  "promotion_status": "planned_candidate_only"
}
```

## Good signs

```text id="vjcj4h"
1. All 7 required null models are registered as planned.
2. comparison_computed = false is enforced.
3. residual_computed = false is enforced.
4. full_bmc_blocked_gate is present and passing.
5. no_null_comparison_result_gate is present and passing.
6. needNullModel remains active.
7. Lean builds with policy-only safety theorems.
```

## Two things to flag for adversarial review

First, this line deserves scrutiny:

```text id="sfpg64"
Rephrases "Friedmann recovery" to "recovery of Friedmann" to avoid forbidden phrases.
```

That could be harmless if it only appears in a negative field like:

```text id="v3ewap"
Recovery of Friedmann Claim: false
```

But it could also become a phrase-scan loophole. The reviewer should check whether the code is truly blocking the concept, not merely avoiding exact-string detection.

Second, the walkthrough says both:

```text id="f55lts"
18 comprehensive tests
```

and later:

```text id="a0ien6"
17 comprehensive unit tests
```

That is minor, but the adversarial review should verify the actual test count and names.

## Debt ledger

```text id="ccm4zr"
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: active, scaffold implemented but no results yet
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent_reported
LeanVerification: reported_retired for policy/safety contracts only
```

## Next step

Run **Sprint 7 adversarial review**. Do not proceed to Sprint 8 until Sprint 7 is accepted or repaired.

The review should focus on whether Sprint 7 really stays a scaffold:

```text id="x4s6ig"
1. No null comparison results are computed.
2. No Friedmann residual is computed.
3. No null model is marked passed/failed.
4. No “BMC beats null models” implication appears.
5. All 7 required null models are present exactly once.
6. All 10 gates are present exactly once.
7. forbidden-phrase scanning is semantic enough and not bypassed by wording like “recovery of Friedmann.”
8. full BMC remains blocked.
```
