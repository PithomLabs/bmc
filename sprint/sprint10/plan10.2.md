# Implementation Plan - BMC Sprint 10.2: Residualrun True Input Faithfulness Repair

This plan outlines repairs to the candidate local-branch residual runner for BMC-0A to enforce true per-point input calculations, eliminate synthetic branches and magnitude constants, and strengthen source-backed registry validation under EBP 2.1.

## User Review Required

> [!IMPORTANT]
> - **Per-Point Local-Branch Trajectory Inputs**: We add trajectory points to `LocalRelationBranch` in `clockseg` so they are saved to `bmc0a_clock_readiness.json`.
> - **True Residual Calculation**: The residual runner reads these trajectory points and computes a residual series via finite differences:
>   - $LHS = (d\ln(a)/d\lambda)^2$
>   - $RHS = (d\phi/d\lambda)^2$
>   - $residual = LHS - RHS$
>   - Statistical metrics (mean, max, RMS) are computed directly from the series. No hidden or embedded constants are used.
> - **Source Branch Registry**: We introduce `SourceBranchRegistry` inside the report to validate that all reported branches match the branches present in the parsed clock-readiness source file.
> - **Blocked Status**: If input files or per-point trajectory inputs are missing, the report is generated as blocked with `interpretation_status = blocked_by_missing_residual_inputs` and `candidate_residual_computed = false`.

## Proposed Changes

### 1. Clock Segmentation Package

#### [MODIFY] [branches.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/branches.go)
- Add `Points []model.TrajectoryPoint json:"points,omitempty"` to `LocalRelationBranch`.

#### [MODIFY] [local_relations.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/local_relations.go)
- Assign `Points: points` when initializing `LocalRelationBranch` in `ExtractLocalRelationBranch`.

---

### 2. Core Residual-Runner Package

#### [MODIFY] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/contracts.go)
- Add constant `InterpretBlockedByMissingResidualInputs = "blocked_by_missing_residual_inputs"`.

#### [MODIFY] [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residual.go)
- Define `ResidualInputPoint` struct containing alpha, phi, candidate LHS, RHS, and provenance.
- Add `ResidualInputPoints []ResidualInputPoint json:"residual_input_points,omitempty"` to `CandidateResidualDiagnostic`.

#### [MODIFY] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/report.go)
- Define `SourceBranchRegistry` struct.
- Add `SourceBranchRegistry SourceBranchRegistry json:"source_branch_registry"` to `ResidualRunReport`.
- In `GenerateBlockedDefaultReport()`, initialize `SourceBranchRegistry` with empty branch list, and `ResidualNullComparisons` as empty list.
- In `RunResidualsFromInputs()`:
  - Remove unconditional synthetic `branch_1` injection.
  - Populate `SourceBranchRegistry` with all branch IDs from the parsed `clockReport`.
  - For each eligible branch:
    - Verify that `b.Points` is non-empty and has length $\ge 2$. If not, treat as blocked.
    - Compute `dAlpha` and `dPhi` at each point using finite differences relative to `Lambda`.
    - Set `CandidateLeftHandSide = dAlpha * dAlpha` and `CandidateRightHandSide = dPhi * dPhi`.
    - Compute `residual = LHS - RHS` for each point.
    - Calculate `MeanAbsResidual`, `MaxAbsResidual`, and `RmsResidual` directly from the computed residual values.
    - Populate `ResidualInputPoints` list.
    - Populate `ResidualCalculationLedger` with `input_provenance = "derived_from_file_read"`, non-zero `NumInputPoints`, and descriptive formulas.
  - If inputs are missing or points are insufficient, return a blocked report with `candidate_residual_computed = false`, `InterpretationStatus = "blocked_by_missing_residual_inputs"`, and empty diagnostic/comparison slices.

#### [MODIFY] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/validate.go)
- Include `InterpretBlockedByMissingResidualInputs` in the list of valid interpretation statuses.
- If `r.CandidateResidualComputed` is false, verify that `r.InterpretationStatus` is a valid blocked status.
- When `bmc0a_clock_readiness` source has `Provenance == "file_read"`, verify:
  - Every `b` in `r.LocalBranchEligibility` has `b.BranchID` present in `r.SourceBranchRegistry.BranchIDs`.
  - Every computed diagnostic `rd` has `rd.BranchID` in the registry.
  - Every computed calculation ledger `cl` has `cl.BranchID` in the registry.
- Enforce strict validation of `CalculationLedger` for computed entries:
  - `InputProvenance` must be `"file_read"` or `"derived_from_file_read"`.
  - `NumInputPoints` must be $> 0$.
  - `InputFields` must not be empty.
  - Verify all formula ID, description, and convention profile fields are present.

#### [MODIFY] [residualrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residualrun_test.go)
- Update `helperComputedReport` to include mock `SourceBranchRegistry` and `ResidualInputPoints` to satisfy validator constraints.
- Implement the requested test cases:
  - `TestResidualRunBlockedReportWhenInputsMissing`
  - `TestResidualRunMetricsChangeWhenInputBranchDataChanges`
  - `TestResidualRunRejectsSyntheticFileBackedBranch`
  - `TestResidualRunRejectsReportedBranchNotInSourceRegistry`
  - `TestResidualRunCalculationLedgerRequiresFormulaTransparency`
  - `TestResidualRunRejectsComputedLedgerWithZeroInputPoints`
  - `TestResidualRunRejectsComputedLedgerWithEmptyInputFields`
  - `TestResidualRunSummaryDistinguishesComputedAndBlockedDiagnostics`
  - `TestResidualRunRejectsComputedBranchProvenanceWithoutCalculationLedger`
  - `TestResidualRunRejectsReportFalseWithComputedDiagnostic`
  - `TestResidualRunRejectsReportTrueWithNoComputedDiagnostic`

---

### 3. CLI Subcommands

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Update `summarizeCmd()` for local-residual profiles to print separate counts for computed and blocked residual diagnostics, along with the total count:
  ```text
  Computed Candidate Residual Diagnostics: N
  Blocked Candidate Residual Diagnostics: N
  Total Candidate Residual Diagnostics: N
  ```

---

## Verification Plan

### Automated Tests
- `go test ./...` in the root workspace.
- `cd BMC && lake build` to verify Lean policy correctness.

### Manual Verification
- Re-run CLI segmentation to generate `out/bmc0a_clock_readiness.json` with branch trajectory points:
  `./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json`
- Re-generate residuals:
  `./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json`
- Validate and summarize the generated report:
  `./ptw-bmc validate --report out/bmc0a_local_residual.json`
  `./ptw-bmc summarize --report out/bmc0a_local_residual.json`
