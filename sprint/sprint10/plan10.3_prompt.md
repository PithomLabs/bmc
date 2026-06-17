# BMC Sprint 10.3: Residualrun Comparison and Validator Hardening

Implement **BMC Sprint 10.3**.

Current Sprint 10.2 adversarial verdict:

```text
overall_verdict: accept_with_repairs
faithfulness_verdict: accepted
promotion_recommendation: promoted_candidate_residual_runner_artifact_after_repairs, after repairs
```

Do not start Sprint 11.

Do not claim recovery, validation, superiority over null models, classical-limit recovery, global cosmology, scientific novelty, or full BMC promotion.

## Sprint 10.3 goal

Repair the remaining Sprint 10.2 issues:

```text
1. Stop hard-coding residual/null comparison target_residual_ids to candidate_residual_branch_0.
2. Strengthen residual_input_points validation.
3. Reject forbidden calculation formula IDs.
4. Add missing regression tests.
5. Regenerate and validate the residual report.
```

The generated artifact already appears faithful at the computation level. Sprint 10.3 is a hardening pass before final acceptance.

## Required repair 1: build comparison targets from actual computed diagnostics

Patch:

```text
internal/bmc/residualrun/report.go
```

Current issue:

```text
Residual/null comparison generation hard-codes target_residual_ids to candidate_residual_branch_0.
```

Repair rule:

```text
target_residual_ids must be built from actual CandidateResidualDiagnostic records where residual_computed = true.
```

Do not hard-code:

```text
candidate_residual_branch_0
branch_0
```

Implementation rule:

```go
var computedTargetResidualIDs []string
for _, d := range report.CandidateResidualDiagnostics {
    if d.ResidualComputed {
        computedTargetResidualIDs = append(computedTargetResidualIDs, d.ResidualID)
    }
}
```

If `computedTargetResidualIDs` is empty:

```text
do not create a computed comparison
set comparison_computed = false if a blocked comparison record is emitted
or leave residual_null_comparisons empty
```

Add regression test:

```text
TestResidualRunComparisonTargetsActualComputedResidual
```

Test case:

```text
branch_0 is blocked
branch_1 is eligible/computed
comparison must target candidate_residual_branch_1, not candidate_residual_branch_0
```

## Required repair 2: strengthen residual_input_points numeric validation

Patch:

```text
internal/bmc/residualrun/validate.go
```

For every `ResidualInputPoint` in a computed diagnostic, validation must reject:

```text
nonfinite alpha
nonfinite phi
nonfinite candidate_left_hand_side
nonfinite candidate_right_hand_side
negative point_index
empty branch_id
branch_id mismatch with diagnostic branch_id
```

If `Lambda` is included in `ResidualInputPoint`, also reject:

```text
nonfinite lambda
duplicate or zero delta lambda if sufficient adjacent context is available
```

If `Lambda` is not included, keep finite-difference validation in the computation path and consider adding `lambda` as an optional audit field later.

Add tests:

```text
TestResidualRunRejectsNonfiniteResidualInputPointAlpha
TestResidualRunRejectsNonfiniteResidualInputPointPhi
TestResidualRunRejectsNonfiniteResidualInputPointLHS
TestResidualRunRejectsNonfiniteResidualInputPointRHS
TestResidualRunRejectsResidualInputPointBranchMismatch
TestResidualRunRejectsNegativeResidualInputPointIndex
```

## Required repair 3: restrict residual_input_points provenance

Patch:

```text
internal/bmc/residualrun/validate.go
```

Allowed `ResidualInputPoint.InputProvenance` values:

```text
file_read
derived_from_file_read
```

Reject:

```text
source_artifact_summary
deterministic_fixture
blocked
assumed
fabricated
validated_physics
inferred_success
empty string
```

Add test:

```text
TestResidualRunRejectsForbiddenResidualInputPointProvenance
```

## Required repair 4: reject forbidden calculation formula IDs

Patch:

```text
internal/bmc/residualrun/validate.go
```

Reject calculation ledger formula IDs such as:

```text
friedmann_residual
classical_residual
recovery_residual
cosmology_recovery_residual
```

Use case-insensitive matching.

Allowed guarded formula ID:

```text
candidate_local_branch_velocity_constraint_residual_v0.1
```

Do not echo forbidden formula IDs in validation errors. Keep errors phrase-safe.

Add tests:

```text
TestResidualRunRejectsForbiddenFormulaID
TestResidualRunForbiddenFormulaIDErrorIsPhraseSafe
```

## Required repair 5: add missing regression tests

Add or confirm stronger equivalents for:

```text
TestResidualRunRejectsComputedDiagnosticWithEmptyResidualInputPoints
TestResidualRunRejectsComputedResidualWithoutSourceBackedBranchPoints
TestResidualRunSegmentClockEmitsBranchPoints
TestResidualRunComparisonTargetsActualComputedResidual
TestResidualRunRejectsNonfiniteResidualInputPointAlpha
TestResidualRunRejectsNonfiniteResidualInputPointPhi
TestResidualRunRejectsNonfiniteResidualInputPointLHS
TestResidualRunRejectsNonfiniteResidualInputPointRHS
TestResidualRunRejectsForbiddenResidualInputPointProvenance
TestResidualRunRejectsForbiddenFormulaID
```

Keep all previous Sprint 10, 10.1, and 10.2 tests passing.

## Optional repair: improve blocked diagnostic status

Current blocked missing-input path reportedly uses:

```text
blocked_by_clock_fragility
```

Prefer one of:

```text
residual_input_blocked
blocked_by_missing_residual_inputs
```

This is optional unless it affects validation or summary clarity.

## Optional repair: add lambda to residual_input_points

Consider adding:

```go
Lambda *float64 `json:"lambda"`
```

to `ResidualInputPoint` for standalone audit of finite-difference provenance.

This is useful but not required if the calculation ledger and source branch points are sufficient.

## Required verification

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc validate --report out/bmc0a_local_residual.json
./ptw-bmc summarize --report out/bmc0a_local_residual.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also verify deterministic output.

## Required walkthrough after repair

Return:

```text
files changed
comparison target repair
residual_input_points validation repair
formula ID validation repair
new/updated tests
Go test result
CLI segment-clock output
CLI residual generation output
CLI validation output
CLI summary output
Lean build result
remaining limitations
```

## Expected status after Sprint 10.3

Before review:

```text
candidate_residual_runner_candidate_after_repairs
```

Maximum after review confirms repairs:

```text
promoted_candidate_residual_runner_artifact_after_repairs
```

## Strict EBP reminder

A candidate residual diagnostic is not recovery.

A residual/null comparison is not a victory claim.

A local branch is not full cosmology.

A file-backed calculation is still a toy diagnostic.

Full BMC remains blocked.
