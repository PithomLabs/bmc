Approve the **Sprint 11 implementation plan**, with one small wording/enum tightening before implementation. The plan is correctly scoped as a **residual/null comparison audit**, not a physics-promotion sprint. It also keeps source provenance, comparison integrity, stability checks, strict validation, CLI routing, and Lean policy boundaries visible. 

```json id="s11_plan_status"
{
  "sprint": "BMC Sprint 11",
  "plan_status": "approved_with_minor_tightening",
  "scope": "residual_null_comparison_audit_only",
  "promotion_status_before_review": "residual_audit_candidate_only",
  "full_bmc_toy_gate": "blocked",
  "forbidden_promotions": [
    "physics recovery",
    "classical-limit recovery",
    "model validation",
    "null-model failure conclusion",
    "full-BMC promotion",
    "scientific novelty"
  ]
}
```

## What is strong in the plan

The plan correctly makes Sprint 11 an **audit layer** over Sprint 10.3, not a new physics claim.

The strongest parts are:

```text id="s11_strengths"
1. It reads or blocks on the accepted source artifacts.
2. It audits whether residual/null comparisons are structurally meaningful.
3. It adds deterministic stability checks.
4. It keeps interpretation candidate-only.
5. It separates comparison integrity from physics success.
6. It keeps full BMC blocked.
7. It includes strict JSON validation and phrase-safe restricted-language handling.
8. It proposes Lean only as policy safety, not as a physics proof.
```

## One required tightening

I would rename this interpretation status:

```text id="old_status"
comparison_integrity_passed
```

to something less promotion-like:

```text id="better_status"
comparison_integrity_structurally_honest
```

Reason: the word `passed` can be misread as “the model passed a test,” even though the plan correctly says it only means structural honesty.

So use:

```text id="status_set_revision"
diagnostic_audit_only
comparison_integrity_structurally_honest
comparison_integrity_failed
comparison_stability_mixed
comparison_unstable
insufficient_target_null_separation
target_null_separation_candidate_unpromoted
blocked_by_missing_residual_inputs
blocked_by_missing_null_inputs
blocked_by_source_unavailable
```

Then update validator/tests/summary accordingly.

## EBP claim audit

```text id="s11_ebp_audit"
Claim: Sprint 11 audits comparison integrity.
Status: allowed.

Claim: Sprint 11 computes stability diagnostics.
Status: allowed, if deterministic and source-backed.

Claim: Sprint 11 says BMC beats null models.
Status: forbidden.

Claim: Sprint 11 says null models failed.
Status: forbidden.

Claim: Sprint 11 promotes recovery or full BMC.
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

## Final instruction

Proceed with Sprint 11 using the plan, but rename `comparison_integrity_passed` to `comparison_integrity_structurally_honest` or equivalent. Keep the sprint as an audit of comparison quality only. No victory, no recovery, no null failure, no full-BMC promotion.
