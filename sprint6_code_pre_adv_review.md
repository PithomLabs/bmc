Sprint 6 is **implementation-complete and ready for adversarial review**, but **not accepted yet**.

Based on your walkthrough, the implementation appears aligned with the approved Sprint 6 boundary: it defines Friedmann-residual **specification**, candidate maps, derivative-readiness contracts, null-model requirements, gates, CLI routing, validation, and Lean policy theorems — while keeping:

```text id="w8e86j"
residual_computed: false
friedmann_recovery_claim: false
spec_scope: candidate_specification_only
full_bmc_gate: blocked
promotion_status: planned_candidate_only
```

That is exactly the right safety posture.

## Current EBP status

```json id="h7qxgg"
{
  "sprint": "BMC Sprint 6",
  "artifact": "friedmann-residual specification and gate design",
  "implementation_status": "complete_reported",
  "acceptance_status": "pending_adversarial_review",
  "maximum_future_promotion": "promoted_friedmann_spec_artifact_after_repairs",
  "residual_computed": false,
  "friedmann_recovery_claim": false,
  "full_bmc_toy_gate": "blocked",
  "clock_choice_debt": "active",
  "classical_target_debt": "active",
  "unit_convention_debt": "active",
  "sign_convention_debt": "active",
  "normalization_debt": "active",
  "null_model_debt": "active"
}
```

## Key claims allowed

```text id="h2mrxd"
1. Sprint 6 specifies what a future Friedmann-style residual would require.
2. Sprint 6 defines safety gates that prevent premature Friedmann recovery claims.
3. Sprint 6 does not compute a residual.
4. Sprint 6 does not recover Friedmann equations.
5. Sprint 6 does not unblock full BMC.
```

## Key claims still forbidden

```text id="s0qvtc"
Friedmann recovery
ready for Friedmann recovery
classical cosmology recovered
full BMC validation
Bohmian cosmology proof
quantum gravity claim
solution to the problem of time
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

## What adversarial review should focus on

The review should verify:

```text id="m50cya"
1. No code path computes a Friedmann residual.
2. residual_computed is always false and validation rejects true.
3. friedmann_recovery_claim is always false and validation rejects true.
4. Candidate maps cannot be marked validated/recovered.
5. Branch requirements cannot be marked ready/pass/recovered.
6. Null-model debt remains active.
7. Convention debts remain active.
8. CLI validate/summarize routes the new schema correctly.
9. Lean theorems are policy-only and prove no physics recovery.
10. Forbidden phrase scan catches dangerous wording.
```

## Debt ledger after reported implementation

```text id="pb6tyd"
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: active
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent_reported
LeanVerification: reported_retired
```

Next step: run **Sprint 6 adversarial review**. Do not proceed to Sprint 7 until Sprint 6 is accepted or repaired.
