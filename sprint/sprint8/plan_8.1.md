# Implementation Plan - BMC Sprint 8.1: Prior-Art Boundary Repair

This plan details the implementation of the Sprint 8.1 repairs to transition the Prior-Art Boundary Note from `accept_with_repairs` status.

## User Review Required

> [!IMPORTANT]
> The plan maintains the small, deterministic boundary-note design without introducing paper ingestion or complex bibliographic engines. All validation errors are revised to be completely phrase-safe.

## Proposed Changes

### Core Prior-Art Package

#### [MODIFY] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/report.go)
- Add the field `EbpDebtVocabulary string json:"ebp_debt_vocabulary"` to the `PriorArtBoundaryReport` struct.
- In `GenerateDefaultReport()`, set `EbpDebtVocabulary = "ptw_runtime_debt_status_v0.1"`.
- Set the `ReviewStatus` of all 5 seeded prior-art sources to `seed_unreviewed` or `human_review_required`.
- Soften `BoundaryUse` descriptions for the sources to avoid strong claims (e.g. use "Seed source for future review...").
- In warnings, add the note: `"Runtime EBP debt labels are not adversarial-review classifications."`

#### [MODIFY] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/validate.go)
- Validate that `ebp_debt_vocabulary` is exactly `"ptw_runtime_debt_status_v0.1"`.
- Implement claim-specific boundary-status validation mapping. Reject status assignments that deviate from the allowed sets:
  - `bmc_uses_bohmian_minisuperspace_trajectories`: `established_prior_art` | `likely_prior_art`
  - `bmc_uses_wheeler_dewitt_toy_equation`: `established_prior_art` | `likely_prior_art`
  - `bmc_uses_scalar_field_clock_checks`: `likely_prior_art` | `unknown_requires_review`
  - `bmc_uses_quantum_potential_diagnostics`: `established_prior_art` | `likely_prior_art`
  - `bmc_uses_node_obstruction_detection`: `implementation_variant` | `likely_prior_art` | `unknown_requires_review`
  - `bmc_uses_local_branch_segmentation`: `implementation_variant` | `unknown_requires_review`
  - `bmc_defines_null_model_scaffold`: `workflow_distinctive_candidate` | `unknown_requires_review`
  - `bmc_uses_ebp_debt_gates`: `workflow_distinctive_candidate` | `unknown_requires_review`
  - `bmc_uses_lean_policy_contracts`: `workflow_distinctive_candidate` | `unknown_requires_review`
  - `bmc_does_not_claim_recovery`: `not_claimed`
  - `bmc_does_not_claim_scientific_novelty`: `blocked`
- Ensure that validation errors do not echo invalid review statuses, boundary statuses, gate names, gate statuses, or scanned forbidden phrases.
- Update `strings.Contains` checks to use `strings.ToLower` for case-insensitive scanning.

#### [MODIFY] [priorart_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/priorart_test.go)
- Add 11 new tests:
  - `TestPriorArtBoundaryRejectsWrongStatusDirectionForPhysicsClaim`
  - `TestPriorArtBoundaryRejectsWrongStatusDirectionForWorkflowClaim`
  - `TestPriorArtBoundaryRejectsWrongStatusDirectionForNotClaimedRecovery`
  - `TestPriorArtBoundaryRejectsWrongStatusDirectionForNoveltyBoundary`
  - `TestPriorArtBoundaryForbiddenBoundaryStatusErrorIsPhraseSafe`
  - `TestPriorArtBoundaryForbiddenReviewStatusErrorIsPhraseSafe`
  - `TestPriorArtBoundaryForbiddenPhraseErrorIsPhraseSafe`
  - `TestPriorArtBoundaryDeclaresDebtVocabulary`
  - `TestPriorArtBoundaryCLIRoutingRunsGenerateValidateSummarize`
  - `TestPriorArtBoundaryUnknownProfileFailsAtCLI`
  - `TestPriorArtBoundaryForbiddenPhraseScanIsCaseInsensitive`

---

### Command Line Interface & Goals

#### [MODIFY] [README_GOALS.md](file:///home/chaschel/Documents/go/bmc/README_GOALS.md)
- Add a clear warning note to clarify that Sprint 8-Lite is limited strictly to a prior-art boundary note and does not implement claim extraction or citation databases.

---

## Verification Plan

### Automated Tests
- Run `go test ./...` in the root workspace.
- Run `cd BMC && /home/chaschel/.elan/bin/lake build` to verify Lean verification policy correctness.

### Manual Verification
- Compile `ptw-bmc` via `go build -buildvcs=false ./cmd/ptw-bmc`.
- Execute `./ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json` to verify generation.
- Run `./ptw-bmc validate --report out/bmc0a_prior_art_boundary.json`.
- Run `./ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json` and inspect output.
