# Implementation Plan - BMC Sprint 9: Null-Model Runner for BMC-0A

This plan outlines the design and implementation of a narrow null-model runner for BMC-0A, mapping and comparing target diagnostics with null controls without asserting any recovery or superiority.

## User Review Required

> [!IMPORTANT]
> The sprint is strictly constrained to a diagnostic comparison. No Friedmann residuals will be computed, and no physics-limit recovery or victory claims will be made.

## Proposed Changes

### Core Null-Model Runner Package

Create the package `internal/bmc/nullrun/` with the following files:

#### [NEW] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/contracts.go)
Defines enums and constants:
- **RunStatus**: `diagnostics_generated`, `blocked`, `deferred`
- **DiagnosticStatus**: `finite`, `nonfinite`, `node_blocked`, `clock_fragile`, `local_only`, `not_available`
- **InterpretationStatus**: `diagnostic_comparison_only`, `mixed_diagnostics`, `insufficient_separation`, `target_null_separation_candidate_unpromoted`, `blocked_by_clock_fragility`, `blocked_by_node_obstruction`
- **Forbidden statuses**: `passed`, `failed`, `validated`, `winner`, `outperformed`, `proved`, `recovered`, `confirmed`

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/report.go)
Defines structs for the JSON report schema version `bmc0a-nullrun-v0.1`:
- `NullRunSeed`, `NullDiagnostics`, `NullModelRun`, `TargetNullDiagnosticComparison`, `NullRunGate`, `NullRunEbpDebt`, and `NullRunReport`.
- Implements `ReadReport()` (with `json.Decoder.DisallowUnknownFields()` and trailing garbage check) and `WriteReport()`.
- Implements `GenerateReport()` which instantiates runs for the 7 registered null models and target diagnostic comparison records.

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/validate.go)
Implements all validation requirements:
- Rejects invalid schema version, scope, toy analysis options, and non-blocked toy gates.
- Rejects any presence of recovery or scientific novelty claims.
- Validates gate presence, enums, EBP debt status, and ensures validation errors do not echo invalid enums or forbidden phrase scan hits (fully phrase-safe).
- Performs case-insensitive forbidden phrase scanning.

#### [NEW] [runner.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/runner.go)
Implements numeric and status simulation runs for the null models. Any stochastic elements use a recorded seed (deterministic RNG).

#### [NEW] [nullrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/nullrun_test.go)
Contains 22 unit tests checking all schema invariants, invalid enums, non-finite values, duplicate IDs, gate counts, phrase-safety, case-insensitivity, and CLI routing execution.

---

### CLI Subcommands

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Add subcommand `run-nullmodels` with flags `--profile` and `--out`.
- Integrate `validate` routing for schema `bmc0a-nullrun-v0.1`.
- Integrate `summarize` routing for schema `bmc0a-nullrun-v0.1`.

---

### Lean Policy

#### [NEW] [NullRun.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/NullRun.lean)
Defines structure `BMCNullRunReport` and policy safety theorems ensuring null-model runner blocks recovery claims and residual computations.

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
Import `BMC.NullRun` to build it during `lake build`.

---

## Verification Plan

### Automated Tests
- Run `go test ./...` in the root workspace.
- Run `cd BMC && /home/chaschel/.elan/bin/lake build` to verify Lean verification policy correctness.

### Manual Verification
- Compile `ptw-bmc` via `go build -buildvcs=false ./cmd/ptw-bmc`.
- Execute `./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json` to verify generation.
- Run `./ptw-bmc validate --report out/bmc0a_nullrun.json`.
- Run `./ptw-bmc summarize --report out/bmc0a_nullrun.json` and inspect output.
