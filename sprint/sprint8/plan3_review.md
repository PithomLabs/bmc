Approve **Sprint 8-Lite implementation plan with minor revisions**.

This stays correctly scoped: a **small BMC-0A prior-art boundary note**, not a full literature-audit subsystem, not paper ingestion, not claim extraction, and not a general physics workbench. That is the right anti-bloat move.

## Key approval

```json id="s8lite_review"
{
  "sprint": "BMC Sprint 8-Lite",
  "plan_status": "approved_with_minor_revisions",
  "scope": "bmc0a_prior_art_boundary_note_only",
  "allowed_maximum_promotion_after_review": "promoted_prior_art_boundary_note_after_repairs",
  "forbidden": [
    "scientific_novelty_claim",
    "full_literature_audit_claim",
    "paper_ingestion",
    "claim_extraction_pipeline",
    "null_model_runner",
    "friedmann_residual_runner",
    "full_bmc_promotion"
  ]
}
```

## Minor revisions before implementation

### 1. Lean should remain optional, not mandatory

Your plan adds:

```text id="58d71y"
BMC/BMC/PriorArtBoundary.lean
```

That is acceptable **only if it stays tiny and policy-only**. But Sprint 8-Lite should not grow because of Lean.

Use this rule:

```text id="tg9v8l"
Add PriorArtBoundary.lean only if it is less than ~80 lines, contains no physics claims, and only proves policy booleans such as no novelty claim, no residual computation, and full BMC blocked.
```

Do **not** make Lean required for prior-art completeness. Lean cannot prove literature completeness.

### 2. Avoid storing “modern_friedmann_scalar_field_bohmian_example” as if fully verified

That seed source is fine, but its `ReviewStatus` should be conservative:

```text id="ze4kmr"
seed_unreviewed
```

or:

```text id="v0zlr6"
human_review_required
```

until someone actually checks it.

### 3. Keep source metadata minimal

Do not spend time perfecting citation metadata in Sprint 8-Lite. The goal is boundary protection, not bibliography management.

Acceptable source fields:

```text id="8v4gfo"
source_id
title
authors
year
source_kind
review_status
relevance_tags
boundary_use
```

No DOI completeness requirement yet.

### 4. Watch the forbidden phrase scanner

The phrase scanner must not print the forbidden phrases in its own error output. Your plan already says this. Keep it that way.

A good error message is:

```text id="8t5kzc"
forbidden phrase detected at warning[2]
```

not:

```text id="uec0ly"
forbidden phrase “...” detected
```

### 5. Summary output must not imply “literature checked”

Use:

```text id="hw6f4h"
Prior-Art Sources Seeded
```

not:

```text id="qsl94q"
Literature Reviewed
```

That distinction matters.

## Approved implementation boundary

Implement only:

```text id="s8lite_allowed"
internal/bmc/priorart/contracts.go
internal/bmc/priorart/report.go
internal/bmc/priorart/validate.go
internal/bmc/priorart/priorart_test.go
cmd/ptw-bmc prior-art-boundary routing
optional tiny BMC/BMC/PriorArtBoundary.lean policy file
```

Do not add:

```text id="s8lite_forbidden"
internal/bmc/litaudit
paper registry framework
paper ingestion
PDF parsing
LLM extraction
claim-to-profile mapper
citation database
general physics branch registry
```

## EBP status

```text id="s8lite_ebp"
Claim: Sprint 8-Lite prevents novelty inflation.
Status: approved candidate purpose.

Claim: Sprint 8-Lite establishes scientific novelty.
Status: forbidden.

Claim: Sprint 8-Lite completes literature review.
Status: forbidden.

Claim: Sprint 8-Lite may seed prior-art boundary sources.
Status: allowed.

needLiteratureAudit: partial_after_implementation
needNullModel: active
needFaithfulnessReview: contested
full_bmc_toy_gate: blocked
promotion_status: prior_art_boundary_note_candidate_only
```

Final instruction:

```text id="go_s8lite"
Go implement Sprint 8-Lite with the minor revisions above. Keep it small, deterministic, and BMC-0A scoped. After verification, return a walkthrough with files changed, test count, CLI output, gate statuses, and remaining limitations.
```
