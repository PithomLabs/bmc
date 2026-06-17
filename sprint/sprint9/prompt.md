# BMC Sprint 9 Planning: Null-Model Runner for BMC-0A

You are planning **BMC Sprint 9** under strict **EBP 2.1** discipline.

Sprint 9 follows accepted artifacts:

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

Current hard boundaries:

```text
residual_computed: false
recovery_claim: false
full_bmc_toy_gate: blocked
scientific_novelty_claim: false
needNullModel: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
```

## Sprint 9 goal

Plan a narrow **BMC null-model runner**.

Sprint 9 should answer:

```text
When we run the registered null models from Sprint 7 against the existing BMC-0A diagnostics, do the null controls reproduce the same kinds of branch, clock, node, and diagnostic behavior as the target BMC-0A configuration?
```

Sprint 9 may compute **null-model diagnostic outputs**.

Sprint 9 must not compute a Friedmann residual.

Sprint 9 must not claim recovery.

Sprint 9 must not claim BMC beats null models.

Sprint 9 must not unblock full BMC.

## Forbidden scope

Do **not** implement:

```text
Friedmann residual computation
Friedmann recovery claim
ready-for-recovery claim
full BMC promotion
scientific novelty claim
paper ingestion
claim extraction
general physics-branch workbench
massive scalar model
LQC comparison
Page-Wootters comparison
quantum gravity validation
problem-of-time solution
```

## Required distinction

Sprint 9 may say:

```text
null-model diagnostics were computed
null-model reports were generated
target and null diagnostic fields are comparable
```

Sprint 9 must not say:

```text
null models passed
null models failed
BMC beats null models
BMC outperforms controls
Friedmann behavior is recovered
classical cosmology is recovered
```

Allowed outcome labels:

```text
diagnostics_generated
comparison_ready
mixed_diagnostics
null_like_behavior_observed
target_null_separation_candidate_unpromoted
insufficient_separation
blocked_by_node_obstruction
blocked_by_clock_fragility
```

Forbidden outcome labels:

```text
passed
failed
winner
outperformed
validated
recovered
proved
confirmed
```

## Recommended package

Create:

```text
internal/bmc/nullrun/
```

Suggested files:

```text
contracts.go
profiles.go
fixtures.go
runner.go
diagnostics.go
compare.go
gates.go
report.go
validate.go
nullrun_test.go
```

Do not modify the accepted Sprint 7 nullspec package except for import/use if needed.

## CLI

Add:

```bash
ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
ptw-bmc validate --report out/bmc0a_nullrun.json
ptw-bmc summarize --report out/bmc0a_nullrun.json
```

Schema version:

```text
bmc0a-nullrun-v0.1
```

Unknown profile must fail safely.

## Required null models

Sprint 9 should run or instantiate diagnostic fixtures for the seven Sprint 7 null models:

```text
constant_phase_control
randomized_phase_control
matched_amplitude_randomized_phase_control
classical_frw_reference_trajectory
same_branch_segmentation_under_null_wavefunctions
node_neighborhood_stress_case
clock_choice_alternative_branch_diagnostic
```

If any cannot yet be run numerically, mark it:

```text
blocked
```

with a reason. Do not fake results.

## Determinism rule

Any randomized null model must be deterministic by recorded seed.

Required fields:

```go
type NullRunSeed struct {
    NullModelID string `json:"null_model_id"`
    Seed uint64 `json:"seed"`
    RngKind string `json:"rng_kind"`
}
```

Repeated generation with the same profile must produce byte-identical JSON.

## Null model run record

Define:

```go
type NullModelRun struct {
    NullModelID string `json:"null_model_id"`
    RunStatus string `json:"run_status"`
    Seed *NullRunSeed `json:"seed,omitempty"`

    DiagnosticStatus string `json:"diagnostic_status"`
    Diagnostics NullDiagnostics `json:"diagnostics"`

    BlockedReason string `json:"blocked_reason,omitempty"`
    Notes string `json:"notes"`
}
```

Allowed `RunStatus` values:

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
```

## Diagnostic fields

Use diagnostic fields already aligned with Sprints 3–7:

```go
type NullDiagnostics struct {
    NodeContactFree bool `json:"node_contact_free"`
    TrajectoryFiniteness string `json:"trajectory_finiteness"`

    NumValidTrajectoryPoints int `json:"num_valid_trajectory_points"`
    NumClockSegments int `json:"num_clock_segments"`
    NumTurningPoints int `json:"num_turning_points"`

    PhiClockGlobalStatus string `json:"phi_clock_global_status"`
    LocalBranchStatus string `json:"local_branch_status"`

    MinAmplitudeR *float64 `json:"min_amplitude_r,omitempty"`
    MaxAbsQAwayFromNodes *float64 `json:"max_abs_q_away_from_nodes,omitempty"`
    MaxPhaseGradient *float64 `json:"max_phase_gradient,omitempty"`

    DiagnosticWarnings []string `json:"diagnostic_warnings"`
}
```

Allowed diagnostic statuses:

```text
finite
nonfinite
node_blocked
clock_fragile
local_only
not_available
```

Use `null` for unavailable numeric metrics. Do not use sentinel numbers such as `-1`.

## Target comparison record

Sprint 9 may compare diagnostic summaries, but not declare winners.

Define:

```go
type TargetNullDiagnosticComparison struct {
    ComparisonID string `json:"comparison_id"`
    TargetArtifact string `json:"target_artifact"`
    NullModelIDs []string `json:"null_model_ids"`

    MetricsCompared []string `json:"metrics_compared"`
    ComparisonComputed bool `json:"comparison_computed"`

    InterpretationStatus string `json:"interpretation_status"`
    Reason string `json:"reason"`
}
```

Required:

```text
comparison_computed = true
```

because Sprint 9 is the null-model runner.

Allowed `InterpretationStatus` values:

```text
diagnostic_comparison_only
mixed_diagnostics
insufficient_separation
target_null_separation_candidate_unpromoted
blocked_by_clock_fragility
blocked_by_node_obstruction
```

Forbidden:

```text
winner
BMC_beats_nulls
validated
confirmed
recovered
proved
```

Important: computed comparison does **not** mean promoted physics claim.

## Report shape

Generate deterministic JSON:

```json
{
  "schema_version": "bmc0a-nullrun-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "artifact_kind": "null_model_runner_report",
  "scope": "bmc0a_only",
  "residual_computed": false,
  "null_diagnostics_computed": true,
  "target_null_comparison_computed": true,
  "recovery_claim": false,
  "scientific_novelty_claim_made": false,
  "full_bmc_toy_gate": "blocked",
  "source_artifacts": [],
  "null_model_runs": [],
  "target_null_diagnostic_comparisons": [],
  "gates": [],
  "ebp_debt_vocabulary": "ptw_runtime_debt_status_v0.1",
  "ebp_debt": {
    "needLiteratureAudit": "partial",
    "needMap": "active",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "classical_target_debt": "active",
    "unit_convention_debt": "active",
    "sign_convention_debt": "active",
    "normalization_debt": "active",
    "containsFinalTruthClaim": "absent",
    "promotion_status": "null_model_runner_candidate_only"
  },
  "warnings": [
    "Sprint 9 computes null-model diagnostics only.",
    "No recovery claim is made.",
    "No residual was computed.",
    "No scientific novelty claim is made.",
    "Full BMC remains blocked.",
    "Runtime EBP debt labels are not adversarial-review classifications."
  ]
}
```

## Required gates

Exactly once each:

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

## Validation requirements

Validation must reject:

```text
wrong schema_version
wrong artifact_kind
wrong scope
toy_analysis_only = false
final_truth_claim = true
residual_computed = true
null_diagnostics_computed = false
target_null_comparison_computed = false
recovery_claim = true
scientific_novelty_claim_made = true
full_bmc_toy_gate != blocked
missing required null models
duplicate null_model_id
unknown null_model_id
forbidden run_status
forbidden diagnostic_status
numeric sentinel values such as -1 for unavailable metrics
nonfinite numeric values
empty target_null_diagnostic_comparisons
comparison_computed != true
forbidden interpretation_status
winner/outperform/validated/recovered/proved language
missing required gates
duplicated gates
unknown gates
any gate status != pass
ebp_debt_vocabulary != ptw_runtime_debt_status_v0.1
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
warning missing “No residual was computed.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
```

Use:

```go
json.Decoder.DisallowUnknownFields()
```

Also reject trailing JSON tokens.

## Forbidden phrase scan

Use case-insensitive scanning.

Reject generated JSON, summary output, comments, and warnings containing:

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
```

Errors must be phrase-safe. Do not echo forbidden strings.

## Summary output

`ptw-bmc summarize --report out/bmc0a_nullrun.json` should show:

```text
BMC Sprint 9 Null-Model Runner Summary
Schema Version: bmc0a-nullrun-v0.1
Scope: bmc0a_only
Residual Computed: false
Null Diagnostics Computed: true
Target/Null Comparison Computed: true
Recovery Claim: false
Scientific Novelty Claim Made: false
Full BMC: blocked
Null Models Registered: 7
Null Models With Diagnostics: N
Null Models Blocked: N
Interpretation Status: diagnostic_comparison_only|mixed_diagnostics|insufficient_separation|target_null_separation_candidate_unpromoted|blocked_by_clock_fragility|blocked_by_node_obstruction
Promotion Status: null_model_runner_candidate_only
```

Do not print winner or recovery language.

## Tests

Plan tests:

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
```

## Lean policy plan

Add a tiny policy-only Lean file only if it remains small:

```text
BMC/BMC/NullRun.lean
```

It may encode:

```text
toyAnalysisOnly
finalTruthClaim
residualComputed
nullDiagnosticsComputed
targetNullComparisonComputed
recoveryClaim
scientificNoveltyClaim
fullBMCBlocked
faithfulnessContested
```

Allowed theorem purpose:

```text
nullrun_forbids_residual_computation
nullrun_forbids_recovery_claim
nullrun_requires_full_bmc_blocked
nullrun_does_not_imply_bmc_beats_nulls
```

Do not prove physics validity.

Do not prove null models pass or fail.

Do not prove recovery.

## Verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
./ptw-bmc validate --report out/bmc0a_nullrun.json
./ptw-bmc summarize --report out/bmc0a_nullrun.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

## Required output

Return an implementation plan in this JSON shape:

```json
{
  "summary": "",
  "proposed_actions": [],
  "files_to_add": [],
  "files_to_modify": [],
  "test_plan": [],
  "cli_plan": [],
  "lean_plan": [],
  "assumptions": [],
  "null_model_plan": [],
  "diagnostic_plan": [],
  "comparison_plan": [],
  "proof_obligations": [],
  "risks": [],
  "human_review_questions": [],
  "ebp_debt_status": {
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
    "containsFinalTruthClaim": ""
  },
  "promotion_status": {
    "sprint9_nullrun": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP reminder

A null-model runner is not a proof engine.

A diagnostic comparison is not a victory claim.

A null model is not “passed” or “failed” in Sprint 9.

A target/null difference is only a candidate signal until reviewed.

Full BMC remains blocked.
