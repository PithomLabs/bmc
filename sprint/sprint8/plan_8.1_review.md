Approve **Sprint 8.1 implementation plan**.

This is the right repair scope. It keeps Sprint 8-Lite small, fixes the adversarial blockers, and does **not** drift into the full literature-audit / paper-ingestion / claim-extraction system.

```json id="s8_1_plan_review"
{
  "sprint": "BMC Sprint 8.1",
  "plan_status": "approved",
  "scope": "prior_art_boundary_repair_only",
  "allowed_maximum_promotion_after_review": "promoted_prior_art_boundary_note_after_repairs",
  "forbidden_expansion": [
    "full_literature_review_engine",
    "paper_ingestion",
    "claim_extraction",
    "citation_database",
    "claim_to_profile_mapper",
    "null_model_runner",
    "friedmann_residual_runner"
  ]
}
```

## What looks correct

The plan directly addresses the adversarial findings:

```text id="vsvsn2"
1. Conservative source labels.
2. Softer boundary-use wording.
3. Claim-specific boundary-status validation.
4. Phrase-safe validation errors.
5. Case-insensitive forbidden phrase scan.
6. Debt-vocabulary clarification.
7. Stronger actual CLI routing tests.
8. README_GOALS clarification without expanding scope.
```

That is exactly what Sprint 8.1 should do.

## One minor revision before implementation

For the `EbpDebtVocabulary` addition, make sure the JSON field is snake_case:

```go id="bygs7k"
EbpDebtVocabulary string `json:"ebp_debt_vocabulary"`
```

and validate it explicitly:

```text id="fqs4rh"
ebp_debt_vocabulary == ptw_runtime_debt_status_v0.1
```

Also make sure the new warning text is neutral:

```text id="oqp29m"
Runtime EBP debt labels are not adversarial-review classifications.
```

That wording is good. Keep it.

## Acceptance boundary after implementation

After Sprint 8.1 is implemented and verified, the artifact can be considered for:

```text id="ixvpl4"
promoted_prior_art_boundary_note_after_repairs
```

but only if the walkthrough confirms:

```text id="sikfl7"
go test ./... passes
go build -buildvcs=false ./cmd/ptw-bmc passes
prior-art-boundary generation passes
validate passes
summarize passes
lake build passes
CLI unknown profile fails safely
forbidden phrase errors remain phrase-safe
source labels remain conservative
```

## EBP status

```text id="mm6ii5"
Claim: Sprint 8.1 repairs the prior-art boundary artifact.
Status: approved implementation plan.

Claim: Sprint 8.1 establishes scientific novelty.
Status: forbidden.

Claim: Sprint 8.1 completes the literature review.
Status: forbidden.

Claim: Sprint 8.1 may enable acceptance of the prior-art boundary note after review.
Status: allowed.

needLiteratureAudit: partial
needMap: partial
needInvariant: partial
needToyCheck: unpaid
needNullModel: unpaid
needObstruction: partial
needFaithfulnessReview: contested
full_bmc_toy_gate: blocked
promotion_status: prior_art_boundary_note_candidate_pending_repair
```

Final instruction:

```text id="fau6tm"
Go implement Sprint 8.1 exactly as planned, with the snake_case debt vocabulary field and no scope expansion. After verification, return a repair walkthrough with files changed, test count, CLI output, Lean build result, and remaining limitations.
```
