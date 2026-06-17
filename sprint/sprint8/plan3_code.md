# Walkthrough: BMC Sprint 8-Lite Prior-Art Boundary Note

This walkthrough summarizes the execution of Sprint 8-Lite under strict EBP 2.1 policy rules.

## Files Added
- [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/contracts.go) - Defines allowed classifications (e.g. source kind, review status, boundary status) and their forbidden/safety-contract boundaries.
- [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/report.go) - Declares report structs, defaults for seeded sources and boundary claims, and IO helper methods.
- [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/validate.go) - Implements schema validator, EBP boundary validations, and phrase-safe scanner.
- [priorart_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/priorart_test.go) - Test suite with 20 targeted unit tests.
- [PriorArtBoundary.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/PriorArtBoundary.lean) - Tiny Lean model (~64 lines) representing the safety policy bounds for the sprint.

## Files Modified
- [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go) - Integrates the `prior-art-boundary` subcommand, and routes validation and summarization by schema version.
- [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean) - Imports the new `PriorArtBoundary` module into the Lean verification build.

## Schema Version
`bmc0a-prior-art-boundary-v0.1`

## Test Count
- Go Package Tests: **20 unit tests** (all passing).
- Lean Verification Theorems: **6 policy safety proofs** + **2 witness checks** (all successfully verified).

## CLI Verification Output

### 1. Generating Report
```bash
$ ./ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
Successfully ran prior-art boundary profile 'bmc0a-prior-art-boundary' and generated report: out/bmc0a_prior_art_boundary.json
```

### 2. Validating Report
```bash
$ ./ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
Prior-Art Boundary Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
```

### 3. Summarizing Report
```bash
$ ./ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
BMC Sprint 8-Lite Prior-Art Boundary Summary
Schema Version: bmc0a-prior-art-boundary-v0.1
Scope: bmc0a_only
Scientific Novelty Claim Made: false
Scientific Novelty Claim Allowed: false
Workflow Distinctiveness: candidate_only
Residual Computed: false
Null Comparison Computed: false
Recovery Claim: false
Prior-Art Sources Seeded: 5
Boundary Claims Declared: 11
Full BMC: blocked
Promotion Status: prior_art_boundary_note_candidate_only
```

## Gate Statuses
All 10 safety gates evaluate to **pass** in the verified artifact:
- `toy_analysis_only_gate`: **pass**
- `no_final_truth_claim_gate`: **pass**
- `no_scientific_novelty_claim_gate`: **pass**
- `prior_art_sources_seeded_gate`: **pass**
- `boundary_claims_declared_gate`: **pass**
- `no_residual_computation_gate`: **pass**
- `no_null_comparison_result_gate`: **pass**
- `no_recovery_claim_gate`: **pass**
- `full_bmc_blocked_gate`: **pass**
- `faithfulness_contested_gate`: **pass**

## EBP Debt Status
- `needLiteratureAudit`: `partial`
- `needMap`: `active`
- `needInvariant`: `partial`
- `needToyCheck`: `active`
- `needNullModel`: `active`
- `needObstruction`: `active`
- `needFaithfulnessReview`: `contested`
- `clock_choice_debt`: `active`
- `classical_target_debt`: `active`
- `unit_convention_debt`: `active`
- `sign_convention_debt`: `active`
- `normalization_debt`: `active`
- `containsFinalTruthClaim`: `absent`
- `promotion_status`: `prior_art_boundary_note_candidate_only`

## Remaining Limitations
- Sprint 8-Lite does not implement a full literature database, paper indexing, ingestion pipelines, or automated claim-to-profile mapping.
- All literature references are minimal seed placeholders to demonstrate verification structures rather than complete bibliography lists.
- Full BMC promotion remains blocked until the null-model runner and future residual comparison contracts are defined.
