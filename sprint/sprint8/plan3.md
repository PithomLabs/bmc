# Implementation Plan - BMC Sprint 8-Lite: BMC-0A Prior-Art Boundary Note

This plan details the implementation of a small, deterministic prior-art boundary artifact that prevents accidental novelty inflation before the project continues to the BMC null-model runner.

## User Review Required

> [!IMPORTANT]
> This sprint is strictly constrained to a lean, deterministic prior-art boundary note. It does not implement any full literature audit, paper ingestion, claim extraction, general workbench, or residual/null-model computations.

## Proposed Changes

### Core Prior-Art Package

Create the package `internal/bmc/priorart/` with the following files:

#### [NEW] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/contracts.go)
Defines constants for allowed `SourceKind`, `ReviewStatus`, `BoundaryStatus`, and the forbidden sets of statuses for reviews and boundaries.
- **SourceKind**: `review`, `paper`, `book`, `unknown`
- **ReviewStatus**: `seed_unreviewed`, `abstract_reviewed`, `skim_reviewed`, `human_review_required`
- **Forbidden ReviewStatus**: `proves_our_claim`, `confirms_novelty`, `settled`, `definitive`
- **BoundaryStatus**: `established_prior_art`, `likely_prior_art`, `implementation_variant`, `workflow_distinctive_candidate`, `unknown_requires_review`, `blocked`, `not_claimed`
- **Forbidden BoundaryStatus**: `novel`, `first_ever`, `scientifically_original`, `breakthrough`, `proved_new`, `validated_physics`

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/report.go)
Defines the `PriorArtSource`, `BoundaryClaim`, `PriorArtBoundaryReport`, and `EbpDebt` structs matching the requested JSON structure. Includes:
- `GenerateReport()` function to construct the deterministic report with seeded sources and boundary claims.
- `ReadReport(path string)` using `json.Decoder` with `DisallowUnknownFields()` and checks for trailing tokens/garbage.
- `WriteReport(r *PriorArtBoundaryReport, path string)` to write pretty-printed JSON.

The seeded sources:
1. `bohmian_quantum_cosmology_review`
2. `bohmian_quantum_gravity_cosmology_review`
3. `scalar_field_minisuperspace_classical_limit_paper`
4. `gaussian_superposition_bohmian_trajectory_paper`
5. `modern_friedmann_scalar_field_bohmian_example`

The seeded boundary claims (and their statuses):
1. `bmc_uses_bohmian_minisuperspace_trajectories` (`established_prior_art`)
2. `bmc_uses_wheeler_dewitt_toy_equation` (`established_prior_art`)
3. `bmc_uses_scalar_field_clock_checks` (`likely_prior_art`)
4. `bmc_uses_quantum_potential_diagnostics` (`established_prior_art`)
5. `bmc_uses_node_obstruction_detection` (`implementation_variant`)
6. `bmc_uses_local_branch_segmentation` (`implementation_variant`)
7. `bmc_defines_null_model_scaffold` (`workflow_distinctive_candidate`)
8. `bmc_uses_ebp_debt_gates` (`workflow_distinctive_candidate`)
9. `bmc_uses_lean_policy_contracts` (`workflow_distinctive_candidate`)
10. `bmc_does_not_claim_recovery` (`not_claimed`)
11. `bmc_does_not_claim_scientific_novelty` (`blocked`)

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/validate.go)
Implements validation constraints on the parsed report structure, including:
- Checking schema version, toy analysis flags, novelty/recovery claim prohibitions, residual/null comparison flags, gates, and EBP debts.
- Scanning the JSON content and summary outputs for forbidden phrases without printing those phrases in the error logs or summary.
- Ensuring required sources and boundary claims exist, with no duplicates or forbidden statuses.

#### [NEW] [priorart_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/priorart_test.go)
Contains all 20 required unit tests validating correctness, error boundaries, determinism, duplicate IDs, forbidden phrase scanning, and gate constraints.

---

### Command Line Interface

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Add subcommand `prior-art-boundary` that parses `--profile` and `--out`.
- Under `validateCmd`, detect schema version `bmc0a-prior-art-boundary-v0.1` and route validation.
- Under `summarizeCmd`, detect schema version `bmc0a-prior-art-boundary-v0.1` and route summarization.
- Ensure unknown profile fails safely.

---

### Lean Verification Policy

#### [NEW] [PriorArtBoundary.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/PriorArtBoundary.lean)
Defines the `BMCPriorArtBoundaryReport` structure, gate policy rules, and policy safety proofs in Lean corresponding to Sprint 8-Lite logic.

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
Import `BMC.PriorArtBoundary` to build it during `lake build`.

---

## Verification Plan

### Automated Tests
- Run tests via `go test ./...` in the root workspace.
- Run `cd BMC && /home/chaschel/.elan/bin/lake build` to verify Lean verification policy correctness.

### Manual Verification
- Compile `ptw-bmc` via `go build -buildvcs=false ./cmd/ptw-bmc`.
- Execute `ptw-bmc prior-art-boundary` to generate `out/bmc0a_prior_art_boundary.json`.
- Run `ptw-bmc validate` to verify it passes.
- Run `ptw-bmc summarize` to output the summary.
