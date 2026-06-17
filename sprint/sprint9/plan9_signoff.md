Sprint 9 + 9.1 can now be accepted as:

```json id="s9_1_acceptance"
{
  "sprint": "BMC Sprint 9 + 9.1",
  "artifact": "BMC-0A null-model runner",
  "acceptance_status": "accepted_after_repairs",
  "promotion_status": "promoted_null_model_runner_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked",
  "residual_computed": false,
  "recovery_claim": false,
  "scientific_novelty_claim": false,
  "bmc_beats_null_models_claim": false
}
```

## Why it passes now

The adversarial blockers were repaired:

```text id="eir3ls"
1. Decorative comparison records are rejected.
2. Comparisons must include metrics and real diagnostics-generated null IDs.
3. Blocked/deferred/unknown null IDs cannot be used as comparable diagnostics.
4. Negative/sentinel float metrics are rejected.
5. Unavailable optional metrics now serialize as explicit JSON null.
6. Missing optional metric keys are rejected.
7. EBP debt vocabulary is now ptw_adversarial_review_debt_status_v0.1.
8. Debt labels no longer use active.
9. Deferred null models are counted in summary.
10. All seven null models are accounted for.
11. Go tests pass.
12. CLI generation, validation, and summary pass.
13. Lean policy build passes.
```

The repaired summary is exactly the kind of accounting we needed:

```text id="e8lkr5"
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
Null Models Deferred: 2
Null Models Accounted For: 7/7
```

That fixes the opacity problem from Sprint 9.

## What Sprint 9 now means

Sprint 9 establishes:

```text id="dlofxj"
A validated null-model diagnostic runner exists for BMC-0A.
```

It does **not** establish:

```text id="43qyyi"
BMC beats null models.
Null models passed or failed.
Friedmann behavior is recovered.
Classical cosmology is recovered.
Full BMC is promoted.
```

The correct interpretation is:

```text id="6ce1v2"
BMC-0A now has a bounded null-diagnostic comparison layer. It can generate and validate diagnostic comparisons, but any target/null separation remains unpromoted.
```

## Updated EBP debt status

```text id="9lxgkx"
needLiteratureAudit: partial
needMap: partial
needInvariant: partial
needToyCheck: unpaid
needNullModel: partial
needObstruction: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
containsFinalTruthClaim: absent
NullRunDiagnosticIntegrity: retired for Sprint 9 scope
DiagnosticProvenanceIntegrity: retired for Sprint 9 scope
NoWinnerClaimBoundary: retired
NoResidualComputationBoundary: retired
NoNullComparisonOverclaimBoundary: retired
RecoveryClaimBoundary: retired
LeanPolicyBoundary: retired for policy-only scope
```

## Next step

Now we can move to:

```text id="9cxm62"
Sprint 10: Candidate Local-Branch Residual Runner
```

But Sprint 10 must remain narrow. It should compute **candidate residual diagnostics only on already-validated local branches**, with all convention debts still visible.

Sprint 10 must still forbid:

```text id="hp506h"
Friedmann recovery claim
ready-for-recovery claim
BMC beats null models claim
full BMC promotion
scientific novelty claim
```

The safe Sprint 10 goal is:

```text id="dsiyir"
Compute candidate local-branch residual diagnostics under explicit conventions, compare them against Sprint 9 null diagnostics, and keep all interpretation unpromoted.
```
