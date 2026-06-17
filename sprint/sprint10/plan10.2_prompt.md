# BMC Sprint 10.2: Residualrun True Input Faithfulness Repair

Implement **BMC Sprint 10.2**.

Current status:

```text
Sprint 10.1 adversarial verdict: reject_for_now
Promotion recommendation: do_not_promote
Main blocker: residual magnitudes still come from embedded constants scaled by clock_range, not an auditable per-point local-branch residual calculation.
Secondary blocker: runner fabricates an extra blocked branch not present in the file-backed clock-readiness artifact.
```

Do **not** start Sprint 11.

Do **not** claim recovery, validation, superiority over null models, classical-limit recovery, global cosmology, scientific novelty, or full BMC promotion.

## Sprint 10.2 goal

Repair the residual runner so it follows this rule:

```text
No explicit per-point residual inputs, no computed residual report.
```

A computed candidate residual diagnostic is allowed only if the runner can compute a residual series from actual file-backed local-branch input data.

If that data is not present, the output must be a blocked report.

## Required repairs

### 1. Delete embedded residual magnitude constants

Patch:

```text
internal/bmc/residualrun/report.go
```

Remove any computed-path use of constants like:

```text
0.208
0.502
0.251
```

or any equivalent hidden constants used to produce:

```text
mean_abs_residual
max_abs_residual
rms_residual
```

These metrics must instead be derived from a residual series:

```go
residualValues := []float64{...}

meanAbsResidual = mean(abs(residualValues))
maxAbsResidual = max(abs(residualValues))
rmsResidual = sqrt(mean(residualValues^2))
numEvaluationPoints = len(residualValues)
numFiniteResidualPoints = countFinite(residualValues)
residualFinite = numFiniteResidualPoints == numEvaluationPoints && numEvaluationPoints > 0
```

If `residualValues` cannot be built from real input fields, block the report.

### 2. Add explicit per-point residual input requirements

Define a structure like:

```go
type ResidualInputPoint struct {
    BranchID string `json:"branch_id"`
    PointIndex int `json:"point_index"`

    Alpha *float64 `json:"alpha"`
    Phi *float64 `json:"phi"`

    CandidateLeftHandSide *float64 `json:"candidate_left_hand_side"`
    CandidateRightHandSide *float64 `json:"candidate_right_hand_side"`

    InputProvenance string `json:"input_provenance"`
}
```

The exact field names may differ if the existing Sprint 5/Sprint 6 artifacts already provide better names, but the principle is mandatory:

```text
There must be per-point residual inputs.
The residual value must be computed from those inputs.
The report must disclose which input fields were used.
```

Candidate residual value may be:

```go
residual := CandidateLeftHandSide - CandidateRightHandSide
```

or another explicitly declared candidate formula from Sprint 6, but the formula must be visible in the calculation ledger.

Do not invent a physical Friedmann residual if the source artifacts do not provide the required terms. Use a candidate residual diagnostic only.

### 3. Block when per-point inputs are absent

If file-backed artifacts do not contain enough per-point residual inputs:

```text
candidate_residual_computed = false
candidate_residual_diagnostics = blocked or empty
residual_null_comparisons = empty or comparison_computed=false
interpretation_status = blocked_by_missing_residual_inputs
```

Add this allowed status if missing:

```text
blocked_by_missing_residual_inputs
```

The CLI should still succeed by generating a valid blocked report.

### 4. Delete unconditional synthetic branch injection

Remove any logic that unconditionally appends:

```text
branch_1
```

or any branch not present in the file-backed clock-readiness artifact.

Rules:

```text
Only report branches present in the source clock-readiness artifact.
If a synthetic branch is needed for tests, mark it deterministic_fixture and keep it out of default file-backed reports.
No synthetic branch may claim source_artifact = bmc0a_clock_readiness with file_read provenance.
```

Validation should reject reported file-backed branch IDs that are not present in the clock-readiness source branch set.

### 5. Strengthen source-to-branch validation

When source provenance is `file_read`, validation must require:

```text
reported local_branch_eligibility branch IDs are present in the parsed clock-readiness source artifact
computed residual diagnostics reference only those branch IDs
calculation ledger branch IDs match source-backed branch IDs
```

If the validator cannot inspect the source file directly, then the generated report must include a `source_branch_ids` or `source_branch_registry` field derived from the file, and validation must check against that.

Suggested structure:

```go
type SourceBranchRegistry struct {
    SourceArtifactID string `json:"source_artifact_id"`
    BranchIDs []string `json:"branch_ids"`
}
```

### 6. Make calculation ledger fully auditable

For computed entries, `ResidualCalculationLedger` must include and validation must enforce:

```text
calculation_id nonempty
branch_id nonempty
formula_source nonempty
formula_id nonempty
formula_description nonempty
convention_profile nonempty
input_fields nonempty
input_provenance = file_read or derived_from_file_read
num_input_points > 0
calculation_status = computed_from_local_branch
notes must not imply recovery, validation, or superiority
```

Add allowed `InputProvenance` value if needed:

```text
derived_from_file_read
```

The ledger must disclose every numeric formula component. Hidden constants are forbidden in the computed path unless explicitly declared as convention constants with source and human-review debt.

### 7. Add sensitivity test

Add a test proving that residual metrics change when the actual residual input data changes.

Required test:

```text
TestResidualRunMetricsChangeWhenInputBranchDataChanges
```

Test idea:

```text
1. Build a temporary file-backed input fixture with residual input points.
2. Run residual calculation and capture mean/max/rms.
3. Modify one or more residual input point values.
4. Re-run residual calculation.
5. Assert mean/max/rms changed.
```

Changing only `clock_range` or point count is not enough. The residual magnitudes must depend on the per-point residual inputs.

### 8. Add blocked-input test

Add:

```text
TestResidualRunBlockedReportWhenInputsMissing
```

It must verify:

```text
missing required files or missing per-point residual inputs
=> candidate_residual_computed=false
=> no computed residual diagnostics
=> no computed comparison
=> valid blocked report
```

### 9. Make CLI summary distinguish computed vs blocked diagnostics

Patch:

```text
cmd/ptw-bmc/main.go
```

Summary should print:

```text
Eligible Local Branches: N
Computed Candidate Residual Diagnostics: N
Blocked Candidate Residual Diagnostics: N
Total Candidate Residual Diagnostics: N
Residual/Null Comparisons: N
```

Do not print a single ambiguous line like:

```text
Candidate Residual Diagnostics: 2
```

when only one is computed.

Add test:

```text
TestResidualRunSummaryDistinguishesComputedAndBlockedDiagnostics
```

### 10. Strengthen validator against hidden fixture computation

Validation must reject:

```text
computed_from_bmc0a_local_branch without calculation ledger
computed_from_bmc0a_local_branch with zero input points
computed_from_bmc0a_local_branch with empty input_fields
computed_from_bmc0a_local_branch with formula fields missing
computed residual diagnostics with nil metrics
computed residual diagnostics when candidate_residual_computed=false
candidate_residual_computed=true with no computed residual diagnostic
reported source-backed branches not present in source_branch_registry
```

### 11. Preserve all anti-overclaim boundaries

The report must keep:

```text
toy_analysis_only = true
final_truth_claim = false
residual_recovery_claim = false
scientific_novelty_claim_made = false
bmc_beats_null_models_claim = false
full_bmc_toy_gate = blocked
local_branch_only = true
global_cosmology_claim = false
promotion_status = candidate_residual_runner_candidate_only
```

Do not change Lean into a physics proof. Lean must remain policy-only.

## Required tests

Add or update tests:

```text
TestResidualRunBlockedReportWhenInputsMissing
TestResidualRunMetricsChangeWhenInputBranchDataChanges
TestResidualRunRejectsSyntheticFileBackedBranch
TestResidualRunRejectsReportedBranchNotInSourceRegistry
TestResidualRunCalculationLedgerRequiresFormulaTransparency
TestResidualRunRejectsComputedLedgerWithZeroInputPoints
TestResidualRunRejectsComputedLedgerWithEmptyInputFields
TestResidualRunSummaryDistinguishesComputedAndBlockedDiagnostics
TestResidualRunRejectsComputedBranchProvenanceWithoutCalculationLedger
TestResidualRunRejectsReportFalseWithComputedDiagnostic
TestResidualRunRejectsReportTrueWithNoComputedDiagnostic
```

Keep all existing Sprint 10 and 10.1 tests passing.

## Required verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc validate --report out/bmc0a_local_residual.json
./ptw-bmc summarize --report out/bmc0a_local_residual.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also verify deterministic output generation.

## Required walkthrough after repair

Return:

```text
chosen path: true per-point calculation or blocked-until-inputs-exist
files changed
embedded constants removed
synthetic branch injection removed
per-point residual input structure added
source branch registry or equivalent added
calculation ledger repairs
blocked report behavior
sensitivity test result
CLI summary repair
new/updated tests
Go test result
CLI generation output
CLI validation output
CLI summary output
Lean build result
remaining limitations
```

## Expected status after Sprint 10.2

Before review:

```text
candidate_residual_runner_candidate_after_repairs
```

Maximum after adversarial review confirms repairs:

```text
promoted_candidate_residual_runner_artifact_after_repairs
```

## Strict EBP reminder

A candidate residual diagnostic is not recovery.

A residual/null comparison is not a victory claim.

A local branch is not full cosmology.

A file-backed calculation is still a toy diagnostic.

No per-point residual inputs means no computed residual report.

Full BMC remains blocked.
