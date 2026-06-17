# Implementation Plan - BMC Sprint 10.3: Residualrun Comparison and Validator Hardening

This plan outlines repairs and validation hardening for the candidate local-branch residual runner under Sprint 10.3/EBP 2.1 discipline.

## User Review Required

> [!IMPORTANT]
> - **Dynamic Comparison Targets**: Rather than hard-coding comparison target residual IDs to `candidate_residual_branch_0`, we dynamically aggregate all actual computed diagnostics (`residual_computed = true`) to form the target list.
> - **Stand-alone Lambda Coordinate Audit**: We add a `Lambda` field to `ResidualInputPoint` to support standalone verification of finite-difference ordering and finiteness directly from the output ledger.
> - **Audited Input Points**: We enforce strict check constraints on every `ResidualInputPoint` to reject nonfinite parameters, mismatched branch IDs, negative indexes, or improper provenances.
> - **Formula ID Blacklisting**: Any formula ID in the calculation ledger containing recovery or Friedmann references (e.g. `cosmology_recovery_residual`) is strictly rejected during validation.

## Open Questions

- No open questions. The sprint instructions specify precise validation rules, test naming conventions, and behavior patterns.

## Proposed Changes

### 1. Core Residual-Runner Package

#### [MODIFY] [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residual.go)
- Add `Lambda *float64 json:"lambda"` to `ResidualInputPoint` structure without `omitempty`.

#### [MODIFY] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/report.go)
- In `RunResidualsFromInputs`, dynamically construct `computedTargetResidualIDs` from diagnostics where `ResidualComputed == true`.
- Supply `TargetResidualIDs: computedTargetResidualIDs` in the generated `ResidualNullComparison` blocks (both computed and blocked comparison paths).
- Populate the new `Lambda` field of `ResidualInputPoint` during finite difference execution.
- Change `diag.ResidualStatus` to `ResidualStatusInputBlocked` when per-point inputs are missing.

#### [MODIFY] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/validate.go)
- For every `ResidualInputPoint` in a computed diagnostic, validate that:
  - `Alpha`, `Phi`, `Lambda`, `CandidateLeftHandSide`, and `CandidateRightHandSide` are non-nil and finite.
  - `PointIndex` is $\ge 0$.
  - `BranchID` is non-empty and matches the diagnostic's `BranchID`.
  - `InputProvenance` is exactly `"file_read"` or `"derived_from_file_read"`.
  - `Lambda` points are strictly monotonically increasing ($\Delta\lambda > 0$ and finite) relative to their predecessors.
- Scan calculation ledger formula IDs to reject those containing `friedmann_residual`, `classical_residual`, `recovery_residual`, or `cosmology_recovery_residual` (case-insensitive). Keep errors phrase-safe.

#### [MODIFY] [residualrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residualrun_test.go)
- Update `helperComputedReport` to include mock `Lambda` values.
- Implement the requested regression tests:
  - `TestResidualRunComparisonTargetsActualComputedResidual` (mocking branch 0 as blocked and branch 1 as eligible, verifying targets).
  - `TestResidualRunRejectsNonfiniteResidualInputPointAlpha`
  - `TestResidualRunRejectsNonfiniteResidualInputPointPhi`
  - `TestResidualRunRejectsNonfiniteResidualInputPointLHS`
  - `TestResidualRunRejectsNonfiniteResidualInputPointRHS`
  - `TestResidualRunRejectsResidualInputPointBranchMismatch`
  - `TestResidualRunRejectsNegativeResidualInputPointIndex`
  - `TestResidualRunRejectsForbiddenResidualInputPointProvenance`
  - `TestResidualRunRejectsForbiddenFormulaID`
  - `TestResidualRunForbiddenFormulaIDErrorIsPhraseSafe`
  - `TestResidualRunRejectsComputedDiagnosticWithEmptyResidualInputPoints`
  - `TestResidualRunRejectsComputedResidualWithoutSourceBackedBranchPoints`
- Ensure all other existing tests compile and pass.

---

## Verification Plan

### Automated Tests
- Run `go test ./...` in the root workspace.
- Run `cd BMC && /home/chaschel/.elan/bin/lake build` to verify Lean correctness.

### Manual Verification
- Re-run CLI workflow:
  ```bash
  go build -buildvcs=false ./cmd/ptw-bmc
  ./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
  ./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
  ./ptw-bmc validate --report out/bmc0a_local_residual.json
  ./ptw-bmc summarize --report out/bmc0a_local_residual.json
  ```
