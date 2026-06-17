Implement **BMC Sprint 10.1: Residualrun Faithfulness and Validator Repair**.

Current Sprint 10 adversarial verdict:

```text id="sp4sgo"
reject_for_now
```

Do not start Sprint 11.

Do not promote Sprint 10.

Do not claim recovery, validation, superiority over null diagnostics, global cosmology, scientific novelty, or full BMC.

## Main repair goal

Repair Sprint 10 so the candidate local-branch residual runner is honest about what it computes.

The core blocker:

```text id="su0j5m"
Generated residual metrics are fixed constants, but residual_provenance says computed_from_bmc0a_local_branch.
```

This must be fixed before acceptance.

Choose one honest path.

## Path A: preferred if feasible

Implement an auditable bounded candidate residual calculation.

Requirements:

```text id="vq9lzt"
1. Residual metrics must be derived from branch data, not fixed constants.
2. The report must show enough provenance to audit what inputs were used.
3. residual_provenance may be computed_from_bmc0a_local_branch only if the calculation actually uses local-branch data.
4. Candidate residual diagnostics must remain local-only and unpromoted.
```

Add a compact calculation ledger:

```go id="v917ee"
type ResidualCalculationLedger struct {
    CalculationID string `json:"calculation_id"`
    BranchID string `json:"branch_id"`
    FormulaSource string `json:"formula_source"`
    InputProvenance string `json:"input_provenance"`
    NumInputPoints int `json:"num_input_points"`
    CalculationStatus string `json:"calculation_status"`
    Notes string `json:"notes"`
}
```

Allowed `CalculationStatus` values:

```text id="ujbw58"
computed_from_local_branch
blocked_by_missing_inputs
blocked_by_derivative_unreadiness
blocked_by_convention_debt
source_summary_only
```

## Path B: safer fallback

If an auditable calculation is not yet possible, downgrade the generated residual diagnostics.

Requirements:

```text id="m5nbzz"
1. Set residual_provenance to deterministic_fixture or source_artifact_summary.
2. Do not use computed_from_bmc0a_local_branch.
3. Keep interpretation_status = diagnostic_comparison_only or blocked_by_convention_debt.
4. Do not allow target_null_residual_separation_candidate_unpromoted from fixture-only diagnostics.
5. Explain in notes that the values are fixture-level plumbing diagnostics, not local-branch computations.
```

## Required validator repairs

Patch:

```text id="hqg1ab"
internal/bmc/residualrun/validate.go
```

### 1. Enforce full branch eligibility

A residual diagnostic with `residual_computed = true` must reference a branch where:

```text id="fvgrke"
eligible = true
eligibility_status = eligible_local_branch
node_contact_free = true
trajectory_finite = true
derivative_readiness_status = ready
```

Reject computed diagnostics if any condition fails.

Add tests:

```text id="z3tjcu"
TestResidualRunRejectsEligibleBranchWithNodeContactFalse
TestResidualRunRejectsEligibleBranchWithTrajectoryFiniteFalse
TestResidualRunRejectsComputedResidualWhenDerivativeNotReady
```

### 2. Fix CLI eligible branch count

Patch:

```text id="kzrjj0"
cmd/ptw-bmc/main.go
```

The summary must count only branches satisfying the full eligibility predicate.

Do not print total branch count as eligible count.

Add test:

```text id="y1zl02"
TestResidualRunSummaryCountsOnlyTrulyEligibleBranches
```

### 3. Enforce residual status/provenance consistency

Validation must reject:

```text id="gv2gk7"
residual_computed = true with blocked status
residual_computed = false with generated status
residual_computed = true with provenance blocked
candidate_residual_computed = false while any diagnostic has residual_computed = true
candidate_residual_computed = true with no computed diagnostic
computed_from_bmc0a_local_branch without auditable calculation provenance
```

Add tests:

```text id="ycnqjh"
TestResidualRunRejectsComputedResidualWithBlockedStatus
TestResidualRunRejectsUncomputedResidualWithGeneratedStatus
TestResidualRunRejectsComputedResidualWithBlockedProvenance
TestResidualRunRejectsReportFalseWithComputedDiagnostic
TestResidualRunRejectsReportTrueWithNoComputedDiagnostic
TestResidualRunRejectsComputedBranchProvenanceWithoutCalculationLedger
```

### 4. Enforce metric invariants

Validation must reject:

```text id="ghppd7"
num_evaluation_points < 0
num_finite_residual_points < 0
num_finite_residual_points > num_evaluation_points
residual_finite = true when num_finite_residual_points = 0
computed residual diagnostics with nil mean_abs_residual
computed residual diagnostics with nil max_abs_residual
computed residual diagnostics with nil rms_residual
negative residual magnitudes
sentinel residual magnitudes
nonfinite residual magnitudes
missing optional metric keys
```

Add tests:

```text id="lfdzde"
TestResidualRunRejectsFiniteCountGreaterThanEvalCount
TestResidualRunRejectsNegativeEvaluationPointCount
TestResidualRunRejectsNegativeFinitePointCount
TestResidualRunRejectsResidualFiniteTrueWithZeroFinitePoints
TestResidualRunRejectsComputedResidualWithNilMean
TestResidualRunRejectsComputedResidualWithNilMax
TestResidualRunRejectsComputedResidualWithNilRMS
```

### 5. Strengthen comparison target integrity

For every computed comparison:

```text id="tu6k8a"
every target_residual_id must exist
every referenced target residual must have residual_computed = true
target_residual_ids must not be empty
null_model_ids must not be empty
metrics_compared must not be empty
comparison reason must remain diagnostic-only
```

Reject reports where one valid target hides another unknown or uncomputed target.

Add tests:

```text id="znhyzm"
TestResidualRunRejectsComparisonWithUnknownTargetResidual
TestResidualRunRejectsComparisonWithUncomputedTargetResidual
TestResidualRunRequiresEveryComparisonTargetComputed
```

### 6. Strengthen source artifact validation

Expected source artifact IDs exactly once:

```text id="k26gd8"
bmc0a_clock_readiness
bmc0a_friedmann_spec
bmc0a_nullrun
bmc0a_prior_art_boundary
```

Validation must reject:

```text id="w58se9"
missing source artifact
duplicate source artifact
unknown source artifact
provenance outside allowed vocabulary
file_read with empty path
file_read without supportive status
```

Add tests:

```text id="klgp4b"
TestResidualRunRequiresExpectedSourceArtifacts
TestResidualRunRejectsDuplicateSourceArtifact
TestResidualRunRejectsUnknownSourceArtifact
TestResidualRunRejectsFileReadWithoutPath
```

### 7. Strengthen convention ledger validation

Required convention debts exactly once:

```text id="f7yodw"
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
faithfulness_review_debt
```

Sprint 10.1 default expectations:

```text id="k1mo66"
clock_choice_debt = unpaid
classical_target_debt = unpaid
unit_convention_debt = unpaid
sign_convention_debt = unpaid
normalization_debt = unpaid
faithfulness_review_debt = contested
```

Validation must reject:

```text id="l8hbb0"
missing convention debt
duplicate convention debt
unknown convention debt
empty description
human_review_required = false for unresolved debt
retired convention debt
partial for the five core convention debts in the default report
```

Add tests:

```text id="e1ad2a"
TestResidualRunRequiresAllSixConventionDebts
TestResidualRunRejectsDuplicateConventionDebt
TestResidualRunRejectsUnknownConventionDebt
TestResidualRunRejectsEmptyConventionDebtDescription
TestResidualRunRejectsHumanReviewFalseForUnresolvedConventionDebt
TestResidualRunRejectsPartialCoreConventionDebt
```

### 8. Strengthen local-only boundary tests

Validation must reject:

```text id="r80cwc"
local_branch_only = false
global_cosmology_claim = true
```

Add tests:

```text id="c6ctoy"
TestResidualRunRequiresLocalOnlyBoundary
TestResidualRunRejectsGlobalCosmologyClaim
```

### 9. Phrase-clean generated artifacts and docs

The forbidden phrase scanner must include the missing restricted wording variant identified by review.

Do not place restricted phrases in generated reports, summaries, warnings, walkthroughs, comments, or docs.

For tests, construct restricted phrases programmatically if necessary so static docs stay clean.

Validation errors must remain phrase-safe and must not echo restricted phrases.

Add or update tests:

```text id="q6y57t"
TestResidualRunForbiddenPhraseScanIncludesClassicalCosmologyVariant
TestResidualRunForbiddenPhraseErrorsArePhraseSafe
```

### 10. Lean policy repair

Patch:

```text id="hyn9g0"
BMC/BMC/ResidualRun.lean
```

The policy model must allow both cases:

```text id="4y0xbr"
candidateResidualComputed = true
candidateResidualComputed = false when blocked by no eligible branch
```

Lean must remain policy-only.

Do not prove physics validity.

Do not prove residual correctness.

Do not prove recovery or superiority.

Add policy theorem or structure field for blocked/no-eligible case if needed.

## Required verification

Run:

```bash id="hndxfq"
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc validate --report out/bmc0a_local_residual.json
./ptw-bmc summarize --report out/bmc0a_local_residual.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also verify deterministic output.

## Required walkthrough after repair

Return:

```text id="ykj4v0"
chosen path: A auditable calculation or B fixture downgrade
files changed
branch eligibility repairs
source artifact validation repairs
convention ledger repairs
residual status/provenance consistency repairs
metric invariant repairs
comparison integrity repairs
CLI summary count repair
phrase-cleaning repair
Lean policy repair
new/updated test names
Go test result
CLI generation output
CLI validation output
CLI summary output
Lean build result
remaining limitations
```

## Expected status after Sprint 10.1

Before review:

```text id="w274zg"
candidate_residual_runner_candidate_after_repairs
```

Maximum after adversarial review confirms repairs:

```text id="pfitpt"
promoted_candidate_residual_runner_artifact_after_repairs
```

## Strict EBP reminder

A candidate residual diagnostic is not recovery.

A residual/null comparison is not a victory claim.

A local branch is not full cosmology.

A fixture is not a physical result.

A source summary is not file-backed provenance.

Explicit conventions are not retired convention debt.

Full BMC remains blocked.
