You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 9: Null-Model Runner for BMC-0A**.

Sprint 9 is allowed to compute **null-model diagnostics** and **target/null diagnostic comparison records**.

Sprint 9 is not allowed to compute a Friedmann residual, claim recovery, claim scientific novelty, claim BMC beats null models, or unblock full BMC.

## Current accepted context

Accepted prior artifacts:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A superposition control + node-obstruction artifact
Sprint 3: BMC-0A robustness audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A Friedmann residual specification artifact
Sprint 7: BMC-0A null-model scaffold artifact
Sprint 8-Lite: BMC-0A prior-art boundary note
```

Sprint 9 reportedly added:

```text
internal/bmc/nullrun/contracts.go
internal/bmc/nullrun/report.go
internal/bmc/nullrun/runner.go
internal/bmc/nullrun/validate.go
internal/bmc/nullrun/nullrun_test.go
cmd/ptw-bmc run-nullmodels routing
BMC/BMC/NullRun.lean
```

Reported schema:

```text
bmc0a-nullrun-v0.1
```

Reported CLI:

```bash
./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
./ptw-bmc validate --report out/bmc0a_nullrun.json
./ptw-bmc summarize --report out/bmc0a_nullrun.json
```

Reported summary:

```text
Residual Computed: false
Null Diagnostics Computed: true
Target/Null Comparison Computed: true
Recovery Claim: false
Scientific Novelty Claim Made: false
Full BMC: blocked
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
Interpretation Status: diagnostic_comparison_only
Promotion Status: null_model_runner_candidate_only
```

## Core review question

```text
Does Sprint 9 honestly compute bounded null-model diagnostics and diagnostic-only target/null comparison records, without fabricating results or implying BMC superiority, recovery, validation, or full promotion?
```

## Hard forbidden claims

The artifact must not claim or imply:

```text
BMC beats null models
BMC outperforms controls
null models passed
null models failed
winner
validated physics
confirmed recovery
Friedmann recovery
recovery of Friedmann
ready for recovery
full BMC validated
scientifically novel
breakthrough
problem of time solved
quantum gravity validated
full BMC promoted
```

Acceptable wording:

```text
diagnostic comparison only
deterministic fixture
computed from existing BMC diagnostics
source artifact summary
blocked
deferred
insufficient separation
mixed diagnostics
target/null separation candidate unpromoted
full BMC remains blocked
no residual was computed
no recovery claim is made
```

## Materials to review

Review actual source and generated artifacts, not only the walkthrough:

```text
internal/bmc/nullrun/**
cmd/ptw-bmc/main.go
BMC/BMC/NullRun.lean
BMC/BMC.lean
out/bmc0a_nullrun.json
all relevant Go tests
CLI output
Lean build output
README/walkthrough text if present
```

If expected files or generated artifacts are missing, report them as debt.

## 1. Scope review

Check that Sprint 9 did not add or imply:

```text
Friedmann residual runner
ready-for-recovery logic
classical-limit recovery logic
paper ingestion
claim extraction
general physics profile system
dashboard
full BMC promotion
scientific novelty claim
```

Allowed scope:

```text
null-model diagnostic generation
deterministic seeded null fixtures
diagnostic provenance records
target/null diagnostic comparison records
blocked/deferred null-model statuses
policy-only Lean safety contracts
```

Mark as blocker if Sprint 9 computes residuals or claims recovery/superiority.

## 2. Report identity and hard fields

Review `report.go`, `validate.go`, generated JSON, and CLI summary.

Expected identity:

```text
schema_version = bmc0a-nullrun-v0.1
artifact_kind = null_model_runner_report
scope = bmc0a_only
```

Expected hard fields:

```text
toy_analysis_only = true
final_truth_claim = false
residual_computed = false
null_diagnostics_computed = true
target_null_comparison_computed = true, only if at least one honest comparable diagnostic exists
recovery_claim = false
scientific_novelty_claim_made = false
full_bmc_toy_gate = blocked
ebp_debt_vocabulary = ptw_runtime_debt_status_v0.1
containsFinalTruthClaim = absent
needFaithfulnessReview = contested
```

Mark as blocker if any of these can be changed to an overclaiming state and still validate.

## 3. Null-model registry and run accounting

Expected seven null models, exactly once each:

```text
constant_phase_control
randomized_phase_control
matched_amplitude_randomized_phase_control
classical_frw_reference_trajectory
same_branch_segmentation_under_null_wavefunctions
node_neighborhood_stress_case
clock_choice_alternative_branch_diagnostic
```

Allowed `run_status` values:

```text
diagnostics_generated
blocked
deferred
```

Forbidden:

```text
passed
failed
validated
winner
outperformed
proved
recovered
confirmed
```

Important accounting check:

```text
The summary reports 7 registered null models, 4 with diagnostics, and 1 blocked. Verify where the remaining 2 are counted. If they are deferred, the summary should explicitly show Null Models Deferred: 2 or otherwise account for all 7.
```

Mark as repair-required if counts are opaque.

Mark as blocker if any null model is treated as passed/failed/winner.

## 4. Diagnostic provenance review

Each `NullModelRun` must include:

```text
diagnostic_provenance
```

Allowed values:

```text
computed_from_existing_bmc_diagnostics
deterministic_fixture
source_artifact_summary
blocked
deferred
```

Forbidden:

```text
assumed
fabricated
inferred_success
validated_physics
```

Review whether each diagnostic value is honestly tied to its provenance.

For any run marked `diagnostics_generated`, check:

```text
Is it computed from an existing diagnostic path?
Is it a deterministic fixture?
Is the seed recorded if stochastic/randomized?
Is the provenance clear?
Are unavailable metrics null rather than sentinel values?
```

Mark as blocker if diagnostic numbers appear fabricated, unexplained, or success-implying.

## 5. Randomness and determinism

For randomized nulls, verify:

```text
seed is recorded
rng_kind is recorded
generation is deterministic
repeated generation produces byte-identical JSON
```

Mark as repair-required if deterministic output is not tested.

Mark as blocker if random nulls are nondeterministic without disclosure.

## 6. Diagnostic metric validity

Review all numeric fields:

```text
min_amplitude_r
max_abs_q_away_from_nodes
max_phase_gradient
num_valid_trajectory_points
num_clock_segments
num_turning_points
```

Validation must reject:

```text
nonfinite numeric values
sentinel unavailable values such as -1
negative counts where impossible
missing diagnostic status
diagnostic_status outside allowed enum
```

Unavailable numeric values must be represented as JSON `null`, not sentinel numbers.

## 7. Classical reference trajectory handling

The `classical_frw_reference_trajectory` must be treated as a reference comparator only.

Required:

```text
notes include: reference comparator only; no residual or recovery interpretation
diagnostic_provenance = deterministic_fixture | source_artifact_summary | blocked
run_status = diagnostics_generated | blocked | deferred
```

Mark as blocker if this entry implies classical recovery.

## 8. Target/null diagnostic comparison review

Review `TargetNullDiagnosticComparison`.

Allowed interpretation statuses:

```text
diagnostic_comparison_only
mixed_diagnostics
insufficient_separation
target_null_separation_candidate_unpromoted
blocked_by_clock_fragility
blocked_by_node_obstruction
blocked_by_no_comparable_null_diagnostics
```

Forbidden:

```text
winner
BMC_beats_nulls
validated
confirmed
recovered
proved
outperformed
```

If `target_null_comparison_computed = true`, verify:

```text
at least one honest comparable null diagnostic exists
metrics_compared is nonempty
comparison does not assign winner/loser
reason does not imply superiority
interpretation_status remains unpromoted
```

If no comparable null diagnostics exist, verify:

```text
target_null_comparison_computed = false
interpretation_status = blocked_by_no_comparable_null_diagnostics
```

Mark as blocker if comparison language implies BMC superiority.

## 9. Gates

Expected gates, exactly once each:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
no_residual_computation_gate
null_diagnostics_computed_gate
target_null_comparison_computed_gate
no_winner_claim_gate
no_recovery_claim_gate
no_scientific_novelty_claim_gate
full_bmc_blocked_gate
faithfulness_contested_gate
```

All gates must have:

```text
status = pass
```

Validation must reject:

```text
missing gates
duplicated gates
unknown gates
empty gate names
non-pass gate statuses
```

Mark as blocker if gates are decorative rather than enforced.

## 10. Validator strictness

Review `ValidateReport` or equivalent.

It must reject:

```text
wrong schema_version
wrong artifact_kind
wrong scope
toy_analysis_only = false
final_truth_claim = true
residual_computed = true
null_diagnostics_computed = false
target_null_comparison_computed inconsistent with comparable diagnostics
recovery_claim = true
scientific_novelty_claim_made = true
full_bmc_toy_gate != blocked
missing required null models
duplicate null_model_id
unknown null_model_id
forbidden run_status
forbidden diagnostic_provenance
forbidden diagnostic_status
forbidden interpretation_status
numeric sentinel values
nonfinite numeric values
empty target_null_diagnostic_comparisons when comparison is computed
comparison_computed inconsistent with report-level target_null_comparison_computed
winner/outperform/validated/recovered/proved language
missing required gates
duplicated gates
unknown gates
any required gate status != pass
ebp_debt_vocabulary != ptw_runtime_debt_status_v0.1
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
warning missing “No residual was computed.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
```

Check strict JSON decoding:

```go
json.Decoder.DisallowUnknownFields()
```

Also check that trailing JSON tokens are rejected.

## 11. Forbidden phrase scan

Search source, generated JSON, summary output, tests, comments, warnings, and walkthrough for forbidden phrases or equivalent implications:

```text
BMC beats null models
BMC outperforms controls
null models passed
null models failed
winner
validated physics
confirmed recovery
Friedmann recovery
recovery of Friedmann
ready for recovery
full BMC validated
scientifically novel
breakthrough
problem of time solved
quantum gravity validated
assumed success
fabricated
inferred success
```

The phrase scanner should be case-insensitive and phrase-safe.

Error messages must not echo the forbidden phrase itself. They may say:

```text
forbidden phrase detected at report field warnings[2]
```

but not print the forbidden phrase.

## 12. CLI behavior

Run or inspect reported runs:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
./ptw-bmc validate --report out/bmc0a_nullrun.json
./ptw-bmc summarize --report out/bmc0a_nullrun.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Check:

```text
run-nullmodels subcommand exists
unknown profile fails safely
validate routes bmc0a-nullrun-v0.1 to nullrun validator
summarize routes bmc0a-nullrun-v0.1 to nullrun summarizer
existing schema routes are not weakened
errors are explicit but phrase-safe
```

## 13. Lean policy review

Review:

```text
BMC/BMC/NullRun.lean
BMC/BMC.lean
```

Acceptable Lean scope:

```text
tiny policy-only file
no physics proof
no null-model success proof
no BMC-beats-nulls proof
no residual proof
no recovery proof
no full BMC proof
```

Check:

```text
lake build succeeds
no sorry/admit exists
theorems only encode safety booleans
```

Mark as blocker if Lean claims physics validity.

## 14. Tests

Confirm tests or equivalents exist for:

```text
TestNullRunReportValidation
TestNullRunRejectsResidualComputed
TestNullRunRejectsMissingNullDiagnostics
TestNullRunRejectsMissingTargetNullComparison
TestNullRunRejectsRecoveryClaim
TestNullRunRejectsScientificNoveltyClaim
TestNullRunRequiresAllSevenNullModels
TestNullRunRejectsDuplicateNullModelID
TestNullRunRejectsUnknownNullModelID
TestNullRunRejectsForbiddenRunStatus
TestNullRunRejectsForbiddenDiagnosticProvenance
TestNullRunRejectsForbiddenInterpretationStatus
TestNullRunRejectsSentinelNumericMetrics
TestNullRunRejectsNonfiniteNumericMetrics
TestNullRunRequiresAllGatesExactlyOnce
TestNullRunRejectsNonPassGate
TestNullRunRejectsUnknownFields
TestNullRunRejectsTrailingJSONTokens
TestNullRunDeterministicJSON
TestNullRunForbiddenPhraseScan
TestNullRunForbiddenPhraseErrorsArePhraseSafe
TestNullRunCLIRoutingRunsGenerateValidateSummarize
TestNullRunUnknownProfileFailsAtCLI
TestNullRunClassicalReferenceIsReferenceOnly
TestNullRunAccountsForAllSevenNullModelsInSummary
```

Missing tests should be repair-required unless covered by stronger equivalents.

## EBP debt classification

Classify each item as one of:

```text
unpaid
partial
retired
contested
overclaimed
absent
```

Debt items:

```text
needLiteratureAudit
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
containsFinalTruthClaim
NullRunDiagnosticIntegrity
DiagnosticProvenanceIntegrity
NoWinnerClaimBoundary
NoResidualComputationBoundary
NoNullComparisonOverclaimBoundary
RecoveryClaimBoundary
LeanPolicyBoundary
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
    "needLiteratureAudit": "",
    "needMap": "",
    "needInvariant": "",
    "needToyCheck": "",
    "needNullModel": "",
    "needObstruction": "",
    "needFaithfulnessReview": "",
    "clock_choice_debt": "",
    "classical_target_debt": "",
    "unit_convention_debt": "",
    "sign_convention_debt": "",
    "normalization_debt": "",
    "containsFinalTruthClaim": "",
    "NullRunDiagnosticIntegrity": "",
    "DiagnosticProvenanceIntegrity": "",
    "NoWinnerClaimBoundary": "",
    "NoResidualComputationBoundary": "",
    "NoNullComparisonOverclaimBoundary": "",
    "RecoveryClaimBoundary": "",
    "LeanPolicyBoundary": ""
  },
  "scope_findings": [],
  "null_model_run_findings": [],
  "diagnostic_provenance_findings": [],
  "comparison_findings": [],
  "validator_findings": [],
  "cli_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|null_model_runner_candidate_only|promoted_null_model_runner_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 9 passes perfectly, the maximum allowed recommendation is:

```text
promoted_null_model_runner_artifact_after_repairs
```

Never recommend promotion as:

```text
null models passed
BMC beats null models
BMC outperforms controls
Friedmann recovery
ready for recovery
full BMC
scientific novelty
quantum gravity validated
problem of time solved
```

Remember:

```text
A null-model runner is not a proof engine.
A diagnostic comparison is not a victory claim.
A target/null difference is not a promoted physics claim.
A deterministic fixture is not a physical result.
Full BMC remains blocked.
```
