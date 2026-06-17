Sprint 11.1 is **implementation-complete based on the walkthrough**, and it appears to directly repair the rejected Sprint 11 blocker. Under our process, though, it should be marked **pending adversarial review**, not accepted yet.

```json id="s11_1_status"
{
  "sprint": "BMC Sprint 11.1",
  "artifact": "residual/null audit stability recompute repair",
  "implementation_status": "complete_reported",
  "acceptance_status": "pending_adversarial_review",
  "current_promotion_status": "residual_audit_candidate_only",
  "maximum_after_review": "promoted_residual_audit_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked",
  "physics_claim_promoted": false
}
```

## Initial assessment

This is the right repair shape. The key adversarial blocker was:

```text
decorative stability audit = perturbed = baseline * (1 + magnitude)
```

Your Sprint 11.1 report says that is now replaced by:

```text
copy residual_input_points
perturb copied interval-level proxy records
recompute metrics from copied records
compute baseline-vs-recomputed deltas
```

That is exactly what needed to change.

## Important boundary

Because you kept the fixed schema, the phrase **“interval-level proxy stability diagnostics”** is important. That keeps the claim honest.

It means Sprint 11.1 is not saying:

```text
we perturbed original trajectory geometry
we proved physical stability
we proved residual/null separation is meaningful
```

It is only saying:

```text
we tested whether the audit metric changes under deterministic perturbations of the serialized interval-level residual input records
```

That is acceptable under the current schema.

## What to adversarially verify

The next review should check only the repaired parts:

```text
1. The old baseline-scaling logic is truly gone.
2. alpha/phi/lambda perturbations modify copied residual_input_points or interval proxy records.
3. mean_abs_residual, max_abs_residual, and rms_residual are recomputed from modified records.
4. TestResidualAuditStabilityRecomputesAfterPerturbation would fail under baseline * (1 + magnitude).
5. validation rejects negative deltas and invalid computed/noncomputed stability states.
6. generated JSON labels the audit as interval-level proxy, not physical trajectory perturbation.
7. no comparison audit behavior regressed.
8. no recovery/null-failure/superiority wording appeared.
```

## EBP status

```text
needToyCheck: partial
needNullModel: partial
needFaithfulnessReview: contested
ResidualAuditIntegrity: candidate_after_repairs
ComparisonIntegrity: previously_accepted_pending_regression_check
StabilityAuditIntegrity: repaired_reported_pending_review
SourceProvenanceIntegrity: partial
LocalOnlyBoundary: reported_enforced
NoRecoveryClaimBoundary: reported_enforced
NoBMCBeatsNullsBoundary: reported_enforced

clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
full_bmc_toy_gate: blocked
```

## Bottom line

Sprint 11.1 looks like the correct surgical repair. Do **one focused adversarial review** now. If that review confirms the recomputation is real and not baseline scaling in disguise, then Sprint 11 can be accepted as:

```text
promoted_residual_audit_artifact_after_repairs
```

Still no physics promotion. This remains an audit artifact, not a BMC validation result.
