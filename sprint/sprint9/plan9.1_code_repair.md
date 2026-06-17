Approve **Sprint 9.1 implementation plan with two small revisions**.

The plan directly repairs the Sprint 9 adversarial findings without expanding scope. It stays focused on schema integrity, comparison honesty, explicit null metrics, debt vocabulary, and run accounting.

```json id="s9_1_plan_review"
{
  "sprint": "BMC Sprint 9.1",
  "plan_status": "approved_with_minor_revisions",
  "scope": "nullrun_schema_and_accounting_repair_only",
  "do_not_start": "Sprint 10",
  "maximum_future_promotion_after_review": "promoted_null_model_runner_artifact_after_repairs",
  "forbidden_expansion": [
    "friedmann_residual_runner",
    "recovery_logic",
    "winner_or_superiority_claims",
    "paper_ingestion",
    "claim_extraction",
    "general_physics_profile_system",
    "full_bmc_promotion"
  ]
}
```

## Required minor revisions

### 1. Rename the debt vocabulary label

Because Sprint 9.1 is switching from runtime labels like `active` to the review vocabulary:

```text id="ig6nnd"
unpaid
partial
retired
contested
overclaimed
absent
```

do **not** keep:

```text id="f9tzb8"
ebp_debt_vocabulary = ptw_runtime_debt_status_v0.1
```

Use a clearer label, for example:

```text id="kdp8sg"
ebp_debt_vocabulary = ptw_adversarial_review_debt_status_v0.1
```

or:

```text id="hjxi11"
ebp_debt_vocabulary = ptw_review_debt_status_v0.1
```

Then validate against that exact value.

### 2. Ensure `blocked_by_no_comparable_null_diagnostics` is in contracts

The earlier repair logic requires this status when no honest comparable diagnostics exist:

```text id="qpj682"
blocked_by_no_comparable_null_diagnostics
```

Make sure it is explicitly included in the allowed `InterpretationStatus` constants and validator, even if the default report currently has computed comparisons.

## Approved repairs

```text id="fve5n2"
1. Reject decorative comparison records.
2. Require nonempty metrics_compared.
3. Require comparison null_model_ids to reference diagnostics_generated runs.
4. Reject comparisons pointing to blocked/deferred/unknown null runs.
5. Reject negative/sentinel float metrics.
6. Emit unavailable optional metrics as explicit JSON null.
7. Validate raw JSON key presence for optional metric fields.
8. Replace active debt labels with review vocabulary.
9. Show deferred null-model count in summary.
10. Show accounted-for count as 7/7.
```

## EBP status

```text id="f9ff0j"
Claim: Sprint 9.1 repairs nullrun schema/accounting integrity.
Status: approved implementation plan.

Claim: Sprint 9.1 computes Friedmann residuals.
Status: forbidden.

Claim: Sprint 9.1 proves BMC beats nulls.
Status: forbidden.

Claim: Sprint 9.1 may make the null-model runner acceptable after review.
Status: allowed.

needNullModel: partial after honest runner, not retired
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
full_bmc_toy_gate: blocked
```

Final instruction:

```text id="zfskma"
Go implement Sprint 9.1 with the two revisions above. After verification, return a repair walkthrough with files changed, updated tests, CLI summary output including deferred/accounted-for counts, debt vocabulary value, explicit null metric evidence, Go test result, and Lean build result.
```
