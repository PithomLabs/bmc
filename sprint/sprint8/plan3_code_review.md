Sprint 8-Lite is **implementation-complete**, but under our process it should be marked:

```json id="s8lite_status"
{
  "sprint": "BMC Sprint 8-Lite",
  "artifact": "BMC-0A prior-art boundary note",
  "implementation_status": "complete_reported",
  "acceptance_status": "pending_adversarial_review",
  "current_promotion_status": "prior_art_boundary_note_candidate_only",
  "maximum_future_promotion": "promoted_prior_art_boundary_note_after_repairs"
}
```

## Initial review

This looks well-scoped. It did **not** bloat into the full literature-audit subsystem. The walkthrough shows the right constraints:

```text id="zqilps"
No full literature database.
No paper ingestion.
No automated claim extraction.
No general physics-branch workbench.
No null-model comparison.
No residual computation.
No scientific novelty claim.
Full BMC remains blocked.
```

The optional Lean file is acceptable because it stayed tiny: about 64 lines, policy-only, no physics proof, no prior-art completeness claim.

## Good signs

```text id="n5xol3"
1. Schema is narrow: bmc0a-prior-art-boundary-v0.1.
2. Scope is bmc0a_only.
3. Five prior-art source seeds are recorded without pretending completeness.
4. Eleven boundary claims are declared.
5. All ten gates pass.
6. Summary uses “Prior-Art Sources Seeded,” not “Literature Reviewed.”
7. Novelty claim is explicitly false.
8. Recovery claim is false.
9. Residual and null comparison flags are false.
10. Full BMC remains blocked.
```

## One thing to check carefully in adversarial review

The review should inspect whether the actual seeded source titles/authors are **real and accurate**, or whether the artifact only uses placeholder IDs.

Placeholders are allowed for Sprint 8-Lite, but the report must label them honestly as:

```text id="0iu4bs"
seed_unreviewed
human_review_required
```

It must not imply the seed list is a real completed bibliography.

## Current EBP status

```text id="i6z8m6"
needLiteratureAudit: partial
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
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
promotion_status: prior_art_boundary_note_candidate_only
```

## Next step

Run **Sprint 8-Lite adversarial review**.

The review should focus on:

```text id="wbhal6"
1. Does the artifact avoid claiming scientific novelty?
2. Does it avoid claiming literature completeness?
3. Are seed sources labeled conservatively?
4. Are boundary claims classified without overreach?
5. Does forbidden phrase scanning really catch novelty/recovery language?
6. Does validation reject residual, null-comparison, recovery, and novelty claims?
7. Does the CLI route generation, validation, and summary correctly?
8. Does Lean remain policy-only?
```

Do not proceed to Sprint 9 until Sprint 8-Lite is accepted or repaired.
