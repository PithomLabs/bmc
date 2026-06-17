Based on your reported checklist, **Sprint 8.1 repair is complete**.

I would now mark Sprint 8-Lite + 8.1 as:

```json id="s8_1_status"
{
  "sprint": "BMC Sprint 8-Lite + 8.1",
  "artifact": "BMC-0A prior-art boundary note",
  "implementation_status": "repair_complete_reported",
  "acceptance_status": "accepted_based_on_reported_verification",
  "promotion_status": "promoted_prior_art_boundary_note_after_repairs",
  "full_bmc_toy_gate": "blocked",
  "scientific_novelty_claim": false,
  "residual_computed": false,
  "null_comparison_computed": false,
  "recovery_claim": false
}
```

## Accepted repairs

```text id="e8p8og"
1. Source labels are now conservative.
2. Boundary-use wording is softened.
3. Claim-specific boundary-status direction is enforced.
4. Validation errors are phrase-safe.
5. Forbidden phrase scanning is case-insensitive.
6. EBP debt vocabulary is declared explicitly.
7. CLI routing and unknown-profile behavior are tested.
8. README_GOALS clarifies Sprint 8-Lite scope.
9. Go tests pass.
10. Lean policy build passes.
11. CLI generation, validation, and summary are manually verified.
```

## EBP status after Sprint 8.1

```text id="dt1hd9"
needLiteratureAudit: partial
needMap: partial
needInvariant: partial
needToyCheck: active
needNullModel: active
needObstruction: partial
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent
PriorArtBoundaryIntegrity: retired for Sprint 8-Lite scope
NoNoveltyClaimBoundary: retired for Sprint 8-Lite scope
NoCompletenessClaimBoundary: retired for Sprint 8-Lite scope
NoResidualComputationBoundary: retired
NoNullComparisonBoundary: retired
RecoveryClaimBoundary: retired
LeanPolicyBoundary: retired for policy-only scope
```

## What this does and does not mean

This means:

```text id="ln31sq"
Sprint 8-Lite successfully prevents novelty inflation for BMC-0A.
```

It does **not** mean:

```text id="hdt71q"
the literature review is complete
BMC physics is novel
BMC physics is validated
null models have been run
a residual has been computed
full BMC is promoted
```

## Next step

Now we return to the physics pipeline:

```text id="gpnxqk"
Sprint 9: BMC Null-Model Runner
```

Sprint 9 should be the first sprint that actually runs the null-model comparisons scaffolded in Sprint 7, while still avoiding any “BMC beats null models” overclaim until the reports and adversarial review say so.
