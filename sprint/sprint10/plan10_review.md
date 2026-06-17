Approve **Sprint 10 implementation plan with revisions**.

The plan is pointed in the right direction: **candidate local-branch residual diagnostics only**, not recovery, not classical-limit validation, not full BMC. But Sprint 10 is the first sprint that computes something close to the dangerous word “residual,” so we need tighter guardrails than earlier sprints.

```json id="s10_plan_review"
{
  "sprint": "BMC Sprint 10",
  "plan_status": "approved_with_revisions",
  "scope": "candidate_local_branch_residual_diagnostics_only",
  "forbidden": [
    "Friedmann recovery claim",
    "classical limit claim",
    "BMC beats null models claim",
    "full BMC promotion",
    "scientific novelty claim"
  ],
  "maximum_status_before_review": "candidate_residual_runner_candidate_only"
}
```

## Required revisions before implementation

### 1. Do not hard-code `candidate_residual_computed = true`

This is the most important fix.

The report shape currently assumes:

```json id="d7kegr"
"candidate_residual_computed": true
```

That is only valid if at least one eligible local branch actually receives an honest candidate residual diagnostic.

Add a blocked path:

```text id="xdgqst"
If no eligible local branch exists, candidate_residual_computed must be false.
```

Add allowed report-level outcome:

```text id="27dj0g"
blocked_by_no_eligible_local_branch
```

So Sprint 10 should support both:

```text id="57n4ie"
candidate_residual_computed = true
```

and:

```text id="trw8h5"
candidate_residual_computed = false
```

depending on branch eligibility.

Validation rule:

```text id="s5ok4q"
candidate_residual_computed = true requires at least one candidate_residual_diagnostic with residual_computed = true and an eligible branch.
candidate_residual_computed = false requires a blocked interpretation status and no fake residual metrics.
```

### 2. Tighten `deterministic_fixture` use

`deterministic_fixture` is acceptable for plumbing, but dangerous for the default artifact if it looks like a physics result.

Preferred default provenance:

```text id="3wv2n4"
computed_from_bmc0a_local_branch
```

Allowed fallback:

```text id="p9n4sh"
source_artifact_summary
blocked
```

Use `deterministic_fixture` only if clearly marked as fixture-only and not used as evidence for target/null separation.

Add validation:

```text id="7eum04"
If ResidualProvenance = deterministic_fixture, InterpretationStatus must be diagnostic_comparison_only or blocked_by_convention_debt, not target_null_residual_separation_candidate_unpromoted.
```

### 3. Add source/provenance honesty rule

For each source artifact, record whether it was actually read or only summarized.

Add:

```go id="my2uie"
type SourceArtifactRef struct {
    ArtifactID string `json:"artifact_id"`
    ArtifactKind string `json:"artifact_kind"`
    Provenance string `json:"provenance"`
    Path string `json:"path,omitempty"`
    Status string `json:"status"`
    Notes string `json:"notes"`
}
```

Allowed `Provenance`:

```text id="n4cv2j"
file_read
source_artifact_summary
not_available
```

Forbidden:

```text id="xnh6vu"
assumed
fabricated
validated_physics
```

If the runner does not read `out/bmc0a_clock_readiness.json`, `out/bmc0a_friedmann_spec.json`, and `out/bmc0a_nullrun.json`, it must not pretend it did.

### 4. Add “candidate” wording everywhere

Use:

```text id="ae6taw"
candidate residual diagnostic
candidate local-branch residual
residual-style diagnostic
```

Avoid bare phrases like:

```text id="7yviqe"
Friedmann residual computed
classical residual
classical-limit residual
```

Even in comments and tests, keep the wording guarded.

### 5. Require all convention debts visible and unpaid/contested

The convention ledger must include all six:

```text id="dipry0"
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
faithfulness_review_debt
```

Validation must reject `retired`.

For Sprint 10 defaults, use:

```text id="2pbdyo"
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
faithfulness_review_debt: contested
```

Do not allow the residual runner to retire convention debts.

### 6. Target/null comparison must be conditional

If no honest residual diagnostic exists, no residual/null comparison should be computed.

Validation rule:

```text id="jh1o6i"
residual_null_comparisons[].comparison_computed = true requires at least one target residual with residual_computed = true and nonempty metrics_compared.
```

Also require null IDs to reference Sprint 9 diagnostics or explicitly label them as source summaries.

### 7. Add explicit local-only boundary

The report should contain:

```json id="zsiyk6"
"local_branch_only": true,
"global_cosmology_claim": false
```

Validation must reject:

```text id="5elxs1"
local_branch_only = false
global_cosmology_claim = true
```

This is important because one eligible branch is not full cosmology.

## Approved implementation boundary

Allowed:

```text id="byx0sw"
internal/bmc/residualrun/*
cmd/ptw-bmc run-residuals routing
optional tiny BMC/BMC/ResidualRun.lean
```

Forbidden:

```text id="mc8lgg"
full residual/recovery engine
large dashboard
paper claim extraction
general profile framework
BMC superiority ranking
classical-limit promotion
full BMC promotion
```

## EBP status

```text id="x89mui"
Claim: Sprint 10 may compute candidate local-branch residual diagnostics.
Status: allowed only for eligible local branches.

Claim: Sprint 10 may compute recovery.
Status: forbidden.

Claim: Sprint 10 may compare residual-style diagnostics against nulls.
Status: allowed only as diagnostic comparison, not superiority.

Claim: Sprint 10 retires convention debts.
Status: forbidden.

needNullModel: partial
needToyCheck: partial if honest diagnostics run
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
full_bmc_toy_gate: blocked
```

Final instruction:

```text id="s10_go"
Go implement Sprint 10 with the revisions above. Keep candidate_residual_computed conditional, preserve source provenance honesty, keep all convention debts unpaid/contested, add local_branch_only/global_cosmology_claim guards, and avoid all recovery/superiority language.
```
