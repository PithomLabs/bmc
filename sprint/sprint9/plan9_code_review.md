Sprint 9 is **implementation-complete based on the walkthrough**, but it should be marked:

```json id="s9_status"
{
  "sprint": "BMC Sprint 9",
  "artifact": "BMC-0A null-model runner",
  "implementation_status": "complete_reported",
  "acceptance_status": "pending_adversarial_review",
  "current_promotion_status": "null_model_runner_candidate_only",
  "maximum_future_promotion": "promoted_null_model_runner_artifact_after_repairs",
  "residual_computed": false,
  "recovery_claim": false,
  "scientific_novelty_claim": false,
  "full_bmc_toy_gate": "blocked"
}
```

## Initial review

This looks aligned with the approved Sprint 9 boundary. The important revisions were reportedly implemented:

```text id="eh9xvl"
1. No fake numeric results.
2. Diagnostic provenance is explicit.
3. Blocked/deferred paths are allowed when diagnostics are unavailable.
4. Classical reference trajectory is treated as a reference comparator only.
5. Target/null comparison is diagnostic-only.
6. No residual computation.
7. No recovery claim.
8. No “BMC beats null models” language.
9. Full BMC remains blocked.
```

## One count issue to check

The summary says:

```text id="y5yq69"
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
```

That accounts for 5 of 7. The adversarial review should check where the remaining 2 are counted. They may be `deferred`, which is allowed, but the summary should ideally show:

```text id="yg20ak"
Null Models Deferred: 2
```

or otherwise make the count transparent.

This is not automatically a blocker, but it is a clear review target.

## Review targets before acceptance

Sprint 9 adversarial review should verify:

```text id="y7l3c0"
1. All 7 null models are present exactly once.
2. Every run has honest diagnostic_provenance.
3. No unavailable metric uses sentinel values like -1.
4. The 4 diagnostics-generated runs are real deterministic fixtures or actual existing diagnostics.
5. The 1 blocked run has a clear reason.
6. Any deferred runs are explicitly counted and explained.
7. target_null_comparison_computed = true is justified by at least one honest comparable diagnostic.
8. No interpretation status implies BMC superiority.
9. Classical reference trajectory remains reference-only.
10. Forbidden phrase scanning catches winner/recovery/validation language.
11. Lean remains policy-only.
```

## EBP debt status after implementation, before review

```text id="s9_debt"
needLiteratureAudit: partial
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: partial_candidate, not retired
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent_reported
LeanVerification: policy_only_reported
full_bmc_toy_gate: blocked
```

## Status meaning

Sprint 9 may eventually promote only this:

```text id="j3izje"
promoted_null_model_runner_artifact_after_repairs
```

It must not promote:

```text id="firxzh"
null models passed
BMC beats null models
BMC outperforms controls
Friedmann recovery
ready for recovery
full BMC
```

Next step:

```text id="hjbrs0"
Run Sprint 9 adversarial review before proceeding to Sprint 10.
```
