Sprint 6 + 6.1 can now be accepted as a **promoted Friedmann-specification artifact after repairs**.

The key repaired blockers were the right ones: strict gate cardinality, required `full_bmc_blocked_gate`, null-model promotion prerequisites, explicit `classical_target_debt`, Lean sign-convention debt, faithfulness boundary validation, and trailing JSON-token rejection.

```json id="fbr6jz"
{
  "sprint": "BMC Sprint 6 + 6.1",
  "artifact": "friedmann-residual specification and gate design",
  "acceptance_status": "accepted_after_repairs",
  "promotion_status": {
    "friedmann_spec_artifact": "promoted_friedmann_spec_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked",
    "residual_computed": false,
    "friedmann_recovery_claim": false,
    "spec_scope": "candidate_specification_only"
  },
  "accepted_repairs": [
    "All 10 EBP safety gates are required exactly once.",
    "full_bmc_blocked_gate must exist and pass.",
    "no_residual_computation_gate must exist and pass.",
    "faithfulness_contested_gate must exist and pass.",
    "All null-model requirements must be required before residual promotion.",
    "FriedmannCandidateMap now explicitly tracks classical_target_debt.",
    "Lean now tracks signConventionDebtActive.",
    "Lean proves friedmann_spec_requires_sign_convention_debt_active.",
    "Validator rejects trailing JSON tokens.",
    "Go tests, CLI validation, summary, and Lean build reportedly pass."
  ]
}
```

## EBP claim ledger

```text id="9bl5mx"
Claim 1: Sprint 6 specifies candidate Friedmann-residual requirements.
Status: accepted as specification-only artifact.

Claim 2: Sprint 6 does not compute a Friedmann residual.
Status: validator/Lean protected.

Claim 3: Sprint 6 does not claim Friedmann recovery.
Status: validator/Lean protected.

Claim 4: Sprint 6 does not unblock full BMC.
Status: full_bmc_blocked_gate protected.

Claim 5: Sprint 6 prepares future null-model and convention-debt requirements.
Status: accepted as debt specification, not debt retirement.
```

## Debt status after Sprint 6.1

```text id="us7ygz"
needMap: active / partially scaffolded by candidate maps
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
containsFinalTruthClaim: absent
LeanVerification: retired for current policy/safety contracts only
```

Final status so far:

```text id="uiay60"
Sprint 1: accepted plane-wave control artifact
Sprint 2: accepted superposition control + node-obstruction artifact
Sprint 3: accepted robustness audit artifact with mixed outcome
Sprint 4: accepted clock-fragility diagnostic artifact
Sprint 5: accepted clock-readiness/local segmentation artifact
Sprint 6: accepted Friedmann-specification/gate-design artifact
Full BMC-0.1: still unpromoted
Friedmann recovery: not claimed
```

The next safest step is **Sprint 7 planning: Null-Model Scaffold for Future Friedmann Residuals**.

Do not compute the Friedmann residual yet. Sprint 6 made `needNullModel` active and required null models before any residual promotion, so Sprint 7 should build the null-model comparison scaffolding first.
