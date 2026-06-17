# Implementation Plan - BMC Sprint 10: Candidate Local-Branch Residual Runner

This plan outlines the design and implementation of a narrow candidate local-branch residual runner for BMC-0A under strict EBP 2.1 discipline.

## User Review Required

> [!IMPORTANT]
> The sprint is strictly constrained to a candidate local-branch residual diagnostic runner. No Friedmann recovery is claimed, classical limit limit is not recovered, and full BMC remains blocked. All convention debts remain unpaid or contested.
> Optional float pointer fields in `CandidateResidualMetrics` will not carry `omitempty` in order to serialize as explicit JSON `null` when unavailable. The validator will check the raw JSON to ensure these keys are explicitly present and reject missing keys.

## Proposed Changes

### Core Residual-Runner Package

Create the package `internal/bmc/residualrun/` with the following files:

#### [NEW] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/contracts.go)
Defines enums and constants:
- **EligibilityStatus**: `eligible_local_branch`, `blocked_by_node_obstruction`, `blocked_by_clock_fragility`, `blocked_by_nonfinite_trajectory`, `blocked_by_derivative_unreadiness`, `source_unavailable`
- **ResidualStatus**: `candidate_residual_diagnostics_generated`, `residual_input_blocked`, `residual_nonfinite`, `blocked_by_clock_fragility`, `blocked_by_node_obstruction`, `blocked_by_convention_debt`, `source_unavailable`
- **ResidualProvenance**: `computed_from_bmc0a_local_branch`, `deterministic_fixture`, `source_artifact_summary`, `blocked`
- **InterpretationStatus**: `diagnostic_comparison_only`, `mixed_residual_diagnostics`, `insufficient_target_null_separation`, `target_null_residual_separation_candidate_unpromoted`, `blocked_by_no_comparable_null_diagnostics`, `blocked_by_convention_debt`, `blocked_by_clock_fragility`, `blocked_by_node_obstruction`
- **Forbidden outcome labels**: `recovered`, `validated`, `confirmed`, `proved`, `winner`, `outperformed`, `passed`, `failed`, `classical_limit_achieved`

#### [NEW] [inputs.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/inputs.go)
Defines structure for reading or summarizing required source artifacts (`bmc0a_clock_readiness`, `bmc0a_friedmann_spec`, `bmc0a_nullrun`, `bmc0a_prior_art_boundary`) with fallback to `source_artifact_summary` when not backed by files.

#### [NEW] [branches.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/branches.go)
Defines `LocalBranchEligibility` struct. Implements checks ensuring candidate residual diagnostics are only computed for branches with `eligible = true`, `node_contact_free = true`, and `trajectory_finite = true`.

#### [NEW] [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residual.go)
Defines `CandidateResidualMetrics` and `CandidateResidualDiagnostic` structures.
Ensure `mean_abs_residual`, `max_abs_residual`, and `rms_residual` pointers omit the `omitempty` tag.

#### [NEW] [compare.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/compare.go)
Defines `ResidualNullComparison` struct. Enforces that target/null comparisons check for non-empty target residuals, null models, and compared metrics, rejecting victory language.

#### [NEW] [gates.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/gates.go)
Verifies the presence and status of the 10 required gates exactly once. All must pass.

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/report.go)
Defines `ResidualRunReport` schema version `bmc0a-local-residual-v0.1`.
Implements `GenerateDefaultReport()`, `ReadReport()`, and `WriteReport()`. Uses EBP debt status values `unpaid`, `partial`, and `contested` with vocabulary `ptw_adversarial_review_debt_status_v0.1`.

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/validate.go)
Implements all validation requirements:
- Validates schema parameters, eligibility directions, and convention debts.
- Checks raw JSON key presence for optional metric fields under `diagnostics`.
- Rejects nonfinite, negative, and sentinel values for residual metrics.
- Enforces strict target-null comparison rules (no blocked/deferred null references).
- Rejects forbidden outcomes and victory language.
- Performs case-insensitive forbidden phrase scanning with phrase-safe error logging.

#### [NEW] [residualrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residualrun_test.go)
Contains 25 test cases verifying all validation rules, phrase safety, and CLI routing execution.

---

### CLI Subcommands

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Add subcommand `run-residuals` with flags `--profile` and `--out`.
- Integrate `validate` routing for schema `bmc0a-local-residual-v0.1`.
- Integrate `summarize` routing for schema `bmc0a-local-residual-v0.1` displaying:
  ```text
  BMC Sprint 10 Candidate Local-Branch Residual Summary
  Schema Version: bmc0a-local-residual-v0.1
  Scope: bmc0a_only
  Candidate Residual Computed: true
  Recovery Claim: false
  Scientific Novelty Claim Made: false
  BMC Beats Null Models Claim: false
  Full BMC: blocked
  Eligible Local Branches: N
  Candidate Residual Diagnostics: N
  Residual/Null Comparisons: N
  Interpretation Status: ...
  Promotion Status: candidate_residual_runner_candidate_only
  ```

---

### Lean Policy

#### [NEW] [ResidualRun.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ResidualRun.lean)
Defines structure `BMCResidualRunReport` and policy safety theorems:
- `residualrun_forbids_recovery_claim`
- `residualrun_forbids_bmc_beats_nulls_claim`
- `residualrun_requires_full_bmc_blocked`
- `residualrun_does_not_imply_classical_limit`

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
Import `BMC.ResidualRun` to build it during `lake build`.

---

## Verification Plan

### Automated Tests
- Run `go test ./...` in the root workspace.
- Run `cd BMC && /home/chaschel/.elan/bin/lake build` to verify Lean verification policy correctness.

### Manual Verification
- Compile `ptw-bmc` via `go build -buildvcs=false ./cmd/ptw-bmc`.
- Execute `./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json` to verify generation.
- Run `./ptw-bmc validate --report out/bmc0a_local_residual.json`.
- Run `./ptw-bmc summarize --report out/bmc0a_local_residual.json` and inspect output.
