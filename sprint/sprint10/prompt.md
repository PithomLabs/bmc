# BMC Sprint 10 Planning: Candidate Local-Branch Residual Runner

You are planning **BMC Sprint 10** under strict **EBP 2.1** discipline.

Sprint 10 follows accepted artifacts:

```text id="0imcyw"
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A superposition control + node-obstruction artifact
Sprint 3: BMC-0A robustness audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A candidate residual specification artifact
Sprint 7: BMC-0A null-model scaffold artifact
Sprint 8-Lite: BMC-0A prior-art boundary note
Sprint 9: BMC-0A null-model runner artifact
```

Current hard boundaries:

```text id="6lpfue"
full_bmc_toy_gate: blocked
recovery_claim: false
scientific_novelty_claim: false
bmc_beats_null_models_claim: false
needNullModel: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
```

## Sprint 10 goal

Plan a narrow **candidate local-branch residual runner**.

Sprint 10 may compute bounded candidate residual diagnostics on already-validated local branches.

Sprint 10 must not claim recovery.

Sprint 10 must not say BMC beats null models.

Sprint 10 must not unblock full BMC.

Sprint 10 should answer:

```text id="5m73es"
Given a local BMC-0A branch that passed clock-readiness and obstruction checks, what is the candidate residual behavior under explicitly declared conventions, and how does that diagnostic compare with Sprint 9 null diagnostics without promoting a physics claim?
```

## Important distinction

Sprint 10 may say:

```text id="kkgrur"
candidate residual diagnostics were computed
residual inputs were available or blocked
local-branch residual values were finite or nonfinite
target/null residual-style diagnostic comparison was generated
```

Sprint 10 must not say:

```text id="kxsrbo"
Friedmann behavior recovered
classical cosmology recovered
BMC validated
BMC beats null models
null models passed or failed
full BMC promoted
```

Allowed outcome labels:

```text id="aev73s"
candidate_residual_diagnostics_generated
residual_input_blocked
residual_nonfinite
local_branch_only
mixed_residual_diagnostics
insufficient_target_null_separation
target_null_residual_separation_candidate_unpromoted
blocked_by_clock_fragility
blocked_by_node_obstruction
blocked_by_convention_debt
```

Forbidden outcome labels:

```text id="d7467m"
recovered
validated
confirmed
proved
winner
outperformed
passed
failed
classical_limit_achieved
```

## Recommended package

Create:

```text id="ckffx0"
internal/bmc/residualrun/
```

Suggested files:

```text id="g9meqa"
contracts.go
inputs.go
branches.go
residual.go
compare.go
gates.go
report.go
validate.go
residualrun_test.go
```

Use Sprint 6 specification artifacts as the source of the candidate residual formula and convention ledger.

Use Sprint 5 local-branch readiness artifacts as the source of branch eligibility.

Use Sprint 9 null diagnostics as comparison inputs where applicable.

## CLI

Add:

```bash id="uvjqr5"
ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
ptw-bmc validate --report out/bmc0a_local_residual.json
ptw-bmc summarize --report out/bmc0a_local_residual.json
```

Schema version:

```text id="m3mgxk"
bmc0a-local-residual-v0.1
```

Unknown profile must fail safely.

## Required source artifacts

The report should name source artifacts:

```text id="nfg328"
bmc0a_clock_readiness
bmc0a_friedmann_spec
bmc0a_nullrun
bmc0a_prior_art_boundary
```

If the implementation does not actually read the prior JSON files yet, label the source as:

```text id="mrc5l4"
source_artifact_summary
```

Do not pretend file-backed provenance exists unless the runner actually reads those files.

## Local branch eligibility

Define:

```go id="5k805h"
type LocalBranchEligibility struct {
    BranchID string `json:"branch_id"`
    SourceArtifact string `json:"source_artifact"`
    Eligible bool `json:"eligible"`
    EligibilityStatus string `json:"eligibility_status"`
    Reason string `json:"reason"`

    NodeContactFree bool `json:"node_contact_free"`
    TrajectoryFinite bool `json:"trajectory_finite"`
    LocalClockStatus string `json:"local_clock_status"`
    DerivativeReadinessStatus string `json:"derivative_readiness_status"`
}
```

Allowed `EligibilityStatus` values:

```text id="tjtqqk"
eligible_local_branch
blocked_by_node_obstruction
blocked_by_clock_fragility
blocked_by_nonfinite_trajectory
blocked_by_derivative_unreadiness
source_unavailable
```

Only branches with:

```text id="bny7md"
eligible = true
eligibility_status = eligible_local_branch
node_contact_free = true
trajectory_finite = true
```

may receive candidate residual diagnostics.

## Convention ledger

Define:

```go id="t8zow6"
type ResidualConventionLedger struct {
    ConventionID string `json:"convention_id"`
    Status string `json:"status"`
    Description string `json:"description"`
    HumanReviewRequired bool `json:"human_review_required"`
}
```

Required convention debts:

```text id="zclg3x"
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
faithfulness_review_debt
```

Allowed `Status` values:

```text id="85p8jf"
unpaid
partial
contested
blocked
```

Do not mark these as retired in Sprint 10.

## Candidate residual record

Define:

```go id="rwnayu"
type CandidateResidualDiagnostic struct {
    BranchID string `json:"branch_id"`
    ResidualID string `json:"residual_id"`

    ResidualComputed bool `json:"residual_computed"`
    ResidualStatus string `json:"residual_status"`
    ResidualProvenance string `json:"residual_provenance"`

    Metrics CandidateResidualMetrics `json:"metrics"`
    BlockedReason string `json:"blocked_reason,omitempty"`
    Notes string `json:"notes"`
}
```

Allowed `ResidualStatus` values:

```text id="ewbyz4"
candidate_residual_diagnostics_generated
residual_input_blocked
residual_nonfinite
blocked_by_clock_fragility
blocked_by_node_obstruction
blocked_by_convention_debt
source_unavailable
```

Allowed `ResidualProvenance` values:

```text id="m1a2ci"
computed_from_bmc0a_local_branch
deterministic_fixture
source_artifact_summary
blocked
```

Forbidden `ResidualProvenance` values:

```text id="l0xwmv"
assumed
fabricated
validated_physics
inferred_success
```

## Candidate residual metrics

Define:

```go id="b8w10d"
type CandidateResidualMetrics struct {
    NumEvaluationPoints int `json:"num_evaluation_points"`
    NumFiniteResidualPoints int `json:"num_finite_residual_points"`

    MeanAbsResidual *float64 `json:"mean_abs_residual"`
    MaxAbsResidual *float64 `json:"max_abs_residual"`
    RmsResidual *float64 `json:"rms_residual"`

    ResidualFinite bool `json:"residual_finite"`
    DiagnosticWarnings []string `json:"diagnostic_warnings"`
}
```

Unavailable numeric metrics must be explicit JSON `null`, not omitted.

Validation must reject:

```text id="o3pwol"
NaN
+Inf
-Inf
negative residual magnitudes
sentinel values such as -1
missing optional metric keys
```

## Target/null residual-style comparison

Sprint 10 may compare target candidate residual diagnostics against Sprint 9 null diagnostics or null residual-style fixtures, but must not declare winners.

Define:

```go id="gqupyz"
type ResidualNullComparison struct {
    ComparisonID string `json:"comparison_id"`
    TargetResidualIDs []string `json:"target_residual_ids"`
    NullModelIDs []string `json:"null_model_ids"`
    MetricsCompared []string `json:"metrics_compared"`

    ComparisonComputed bool `json:"comparison_computed"`
    InterpretationStatus string `json:"interpretation_status"`
    Reason string `json:"reason"`
}
```

Allowed `InterpretationStatus` values:

```text id="fsbuyq"
diagnostic_comparison_only
mixed_residual_diagnostics
insufficient_target_null_separation
target_null_residual_separation_candidate_unpromoted
blocked_by_no_comparable_null_diagnostics
blocked_by_convention_debt
blocked_by_clock_fragility
blocked_by_node_obstruction
```

Forbidden:

```text id="4nsd55"
winner
BMC_beats_nulls
validated
confirmed
recovered
proved
outperformed
```

If comparison is computed, require:

```text id="zlkjy5"
target_residual_ids nonempty
null_model_ids nonempty
metrics_compared nonempty
comparison_computed = true
no blocked/deferred nulls used as comparable diagnostics
no winner/superiority language
```

## Report shape

Generate deterministic JSON:

```json id="cid0ab"
{
  "schema_version": "bmc0a-local-residual-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "artifact_kind": "candidate_local_branch_residual_runner",
  "scope": "bmc0a_only",
  "candidate_residual_computed": true,
  "residual_recovery_claim": false,
  "scientific_novelty_claim_made": false,
  "bmc_beats_null_models_claim": false,
  "full_bmc_toy_gate": "blocked",
  "source_artifacts": [],
  "local_branch_eligibility": [],
  "convention_ledger": [],
  "candidate_residual_diagnostics": [],
  "residual_null_comparisons": [],
  "gates": [],
  "ebp_debt_vocabulary": "ptw_adversarial_review_debt_status_v0.1",
  "ebp_debt": {
    "needLiteratureAudit": "partial",
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "partial",
    "needNullModel": "partial",
    "needObstruction": "partial",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "unpaid",
    "classical_target_debt": "unpaid",
    "unit_convention_debt": "unpaid",
    "sign_convention_debt": "unpaid",
    "normalization_debt": "unpaid",
    "containsFinalTruthClaim": "absent",
    "promotion_status": "candidate_residual_runner_candidate_only"
  },
  "warnings": [
    "Sprint 10 computes candidate local-branch residual diagnostics only.",
    "No recovery claim is made.",
    "No scientific novelty claim is made.",
    "No BMC-beats-null-models claim is made.",
    "Full BMC remains blocked.",
    "Convention debts remain unpaid or contested."
  ]
}
```

## Required gates

Exactly once each:

```text id="mdknn8"
toy_analysis_only_gate
no_final_truth_claim_gate
candidate_residual_diagnostics_gate
no_recovery_claim_gate
no_scientific_novelty_claim_gate
no_bmc_beats_null_models_claim_gate
full_bmc_blocked_gate
convention_debts_visible_gate
faithfulness_contested_gate
no_full_bmc_promotion_gate
```

All gates must have:

```text id="bmd8iw"
status = pass
```

## Validation requirements

Validation must reject:

```text id="whfh9u"
wrong schema_version
wrong artifact_kind
wrong scope
toy_analysis_only = false
final_truth_claim = true
candidate_residual_computed = false
residual_recovery_claim = true
scientific_novelty_claim_made = true
bmc_beats_null_models_claim = true
full_bmc_toy_gate != blocked
empty local_branch_eligibility
empty convention_ledger
missing required convention debts
convention debt marked retired
candidate residual computed for ineligible branch
residual_computed = true when branch is ineligible
candidate_residual_diagnostics empty when candidate_residual_computed = true
nonfinite numeric residual metrics
negative residual magnitudes
sentinel residual values
missing optional residual metric keys
forbidden residual_status
forbidden residual_provenance
forbidden interpretation_status
comparison computed with empty target_residual_ids
comparison computed with empty null_model_ids
comparison computed with empty metrics_compared
winner/outperform/validated/recovered/proved language
missing required gates
duplicated gates
unknown gates
any gate status != pass
ebp_debt_vocabulary != ptw_adversarial_review_debt_status_v0.1
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
clock_choice_debt != unpaid
classical_target_debt != unpaid
unit_convention_debt != unpaid
sign_convention_debt != unpaid
normalization_debt != unpaid
warning missing “No recovery claim is made.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
```

Use:

```go id="o0us8y"
json.Decoder.DisallowUnknownFields()
```

Also reject trailing JSON tokens.

## Forbidden phrase scan

Use case-insensitive scanning.

Reject generated JSON, summary output, comments, and warnings containing:

```text id="lpuy9i"
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
classical limit achieved
```

Errors must be phrase-safe. Do not echo forbidden strings.

## Summary output

`ptw-bmc summarize --report out/bmc0a_local_residual.json` should show:

```text id="7pxzvr"
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
Interpretation Status: diagnostic_comparison_only|mixed_residual_diagnostics|insufficient_target_null_separation|target_null_residual_separation_candidate_unpromoted|blocked_by_no_comparable_null_diagnostics|blocked_by_convention_debt|blocked_by_clock_fragility|blocked_by_node_obstruction
Promotion Status: candidate_residual_runner_candidate_only
```

Do not print recovery/superiority language.

## Tests

Plan tests:

```text id="psad0u"
TestResidualRunReportValidation
TestResidualRunRejectsRecoveryClaim
TestResidualRunRejectsScientificNoveltyClaim
TestResidualRunRejectsBMCBeatsNullModelsClaim
TestResidualRunRequiresFullBMCBlocked
TestResidualRunRequiresConventionLedger
TestResidualRunRejectsRetiredConventionDebt
TestResidualRunRejectsResidualForIneligibleBranch
TestResidualRunRejectsNonfiniteResidualMetrics
TestResidualRunRejectsNegativeResidualMetrics
TestResidualRunRejectsSentinelResidualMetrics
TestResidualRunRejectsMissingOptionalMetricKeys
TestResidualRunRequiresAllGatesExactlyOnce
TestResidualRunRejectsNonPassGate
TestResidualRunRejectsForbiddenResidualStatus
TestResidualRunRejectsForbiddenResidualProvenance
TestResidualRunRejectsForbiddenInterpretationStatus
TestResidualRunRejectsDecorativeComparison
TestResidualRunRejectsUnknownFields
TestResidualRunRejectsTrailingJSONTokens
TestResidualRunDeterministicJSON
TestResidualRunForbiddenPhraseScan
TestResidualRunForbiddenPhraseErrorsArePhraseSafe
TestResidualRunCLIRoutingRunsGenerateValidateSummarize
TestResidualRunUnknownProfileFailsAtCLI
```

## Lean policy plan

Add a tiny policy-only Lean file only if it remains small:

```text id="26hcsl"
BMC/BMC/ResidualRun.lean
```

It may encode:

```text id="mxorft"
toyAnalysisOnly
finalTruthClaim
candidateResidualComputed
recoveryClaim
scientificNoveltyClaim
bmcBeatsNullModelsClaim
fullBMCBlocked
faithfulnessContested
```

Allowed theorem purpose:

```text id="befl2z"
residualrun_forbids_recovery_claim
residualrun_forbids_bmc_beats_nulls_claim
residualrun_requires_full_bmc_blocked
residualrun_does_not_imply_classical_limit
```

Do not prove physics validity.

Do not prove recovery.

Do not prove BMC superiority.

## Verification commands

Run:

```bash id="l2awhk"
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc validate --report out/bmc0a_local_residual.json
./ptw-bmc summarize --report out/bmc0a_local_residual.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

## Required output

Return an implementation plan in this JSON shape:

```json id="yrf4x2"
{
  "summary": "",
  "proposed_actions": [],
  "files_to_add": [],
  "files_to_modify": [],
  "test_plan": [],
  "cli_plan": [],
  "lean_plan": [],
  "assumptions": [],
  "branch_eligibility_plan": [],
  "residual_diagnostic_plan": [],
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
    "sprint10_residualrun": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP reminder

A candidate residual diagnostic is not recovery.

A residual/null comparison is not a victory claim.

A local branch is not full cosmology.

Explicit conventions are not retired convention debt.

Full BMC remains blocked.
