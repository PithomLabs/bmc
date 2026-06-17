# BMC Sprint 11 Planning: Residual/Null Comparison Audit and Interpretation Hardening

You are planning **BMC Sprint 11** under strict **EBP 2.1** discipline.

Sprint 11 follows accepted artifacts:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A superposition + node-obstruction artifact
Sprint 3: BMC-0A robustness audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A candidate residual specification artifact
Sprint 7: BMC-0A null-model scaffold artifact
Sprint 8-Lite: BMC-0A prior-art boundary note
Sprint 9 + 9.1: BMC-0A null-model runner artifact
Sprint 10 + 10.3: candidate local-branch residual runner artifact accepted after repairs
```

Current hard boundary:

```text
BMC-0A can compute candidate local-branch residual diagnostics from file-backed per-point branch data.

But this does not imply:
- Friedmann recovery
- classical-limit recovery
- BMC validation
- BMC beats null models
- full BMC promotion
- scientific novelty
```

## Sprint 11 goal

Plan a narrow **residual/null comparison audit**.

Sprint 11 should answer:

```text
Is the residual/null comparison layer meaningful, stable, and non-decorative across local branch perturbations, null diagnostics, and report-validation checks?
```

Sprint 11 may audit comparison stability.

Sprint 11 may compute comparison-quality diagnostics.

Sprint 11 may detect whether the comparison is weak, unstable, decorative, or candidate-useful.

Sprint 11 must not promote the residual result into physics recovery.

## Allowed claims

Sprint 11 may say:

```text
residual/null comparison audit was generated
comparison target integrity was checked
comparison stability diagnostics were computed
comparison was blocked, weak, mixed, unstable, or candidate-useful
comparison remains unpromoted
```

Sprint 11 must not say:

```text
BMC beats null models
null models failed
null models passed
Friedmann behavior recovered
classical cosmology recovered
BMC validated
full BMC promoted
```

## Recommended package

Create:

```text
internal/bmc/residualaudit/
```

Suggested files:

```text
contracts.go
inputs.go
audit.go
stability.go
null_compare.go
gates.go
report.go
validate.go
residualaudit_test.go
```

## CLI

Add:

```bash
ptw-bmc audit-residuals --profile bmc0a-residual-audit --out out/bmc0a_residual_audit.json
ptw-bmc validate --report out/bmc0a_residual_audit.json
ptw-bmc summarize --report out/bmc0a_residual_audit.json
```

Schema version:

```text
bmc0a-residual-audit-v0.1
```

Unknown profile must fail safely.

## Source artifacts

Sprint 11 should read or explicitly summarize:

```text
bmc0a_clock_readiness
bmc0a_nullrun
bmc0a_local_residual
bmc0a_friedmann_spec
bmc0a_prior_art_boundary
```

If files are read, provenance must be `file_read`.

If not read, provenance must be `source_artifact_summary` or `not_available`.

Do not claim file-backed provenance unless the file is actually read.

## Audit focus

Sprint 11 should audit these questions:

```text
1. Are all comparison targets actual computed candidate residual diagnostics?
2. Are null references actual Sprint 9 null diagnostics or clearly source summaries?
3. Does the comparison use nonempty metrics?
4. Are computed metrics stable under small local-branch perturbations?
5. Does the comparison disappear or change if the local residual input changes?
6. Is any separation candidate merely caused by branch count, metadata, or fixture structure?
7. Are blocked/null/unavailable cases accounted for honestly?
8. Does interpretation remain unpromoted?
```

## Comparison audit records

Define:

```go
type ResidualComparisonAudit struct {
    AuditID string `json:"audit_id"`
    SourceComparisonID string `json:"source_comparison_id"`

    AuditComputed bool `json:"audit_computed"`
    AuditStatus string `json:"audit_status"`
    AuditProvenance string `json:"audit_provenance"`

    TargetResidualIDs []string `json:"target_residual_ids"`
    NullModelIDs []string `json:"null_model_ids"`
    MetricsAudited []string `json:"metrics_audited"`

    Findings []string `json:"findings"`
    InterpretationStatus string `json:"interpretation_status"`
    Notes string `json:"notes"`
}
```

Allowed `AuditStatus` values:

```text
comparison_audited
comparison_blocked
comparison_missing
comparison_decorative
comparison_unstable
comparison_mixed
source_unavailable
```

Allowed `AuditProvenance` values:

```text
file_read
derived_from_file_read
source_artifact_summary
blocked
```

Allowed `InterpretationStatus` values:

```text
diagnostic_audit_only
comparison_integrity_passed
comparison_integrity_failed
comparison_stability_mixed
comparison_unstable
insufficient_target_null_separation
target_null_separation_candidate_unpromoted
blocked_by_missing_residual_inputs
blocked_by_missing_null_inputs
blocked_by_source_unavailable
```

Forbidden interpretation values:

```text
winner
passed
failed
validated
confirmed
recovered
proved
outperformed
BMC_beats_nulls
```

Important: `comparison_integrity_passed` means the comparison record is structurally honest. It does **not** mean BMC passed a physics test.

## Stability audit records

Define:

```go
type ResidualStabilityAudit struct {
    StabilityID string `json:"stability_id"`
    BranchID string `json:"branch_id"`

    StabilityComputed bool `json:"stability_computed"`
    PerturbationKind string `json:"perturbation_kind"`
    PerturbationMagnitude float64 `json:"perturbation_magnitude"`

    BaselineMetric string `json:"baseline_metric"`
    BaselineValue *float64 `json:"baseline_value"`
    PerturbedValue *float64 `json:"perturbed_value"`
    AbsoluteDelta *float64 `json:"absolute_delta"`
    RelativeDelta *float64 `json:"relative_delta"`

    StabilityStatus string `json:"stability_status"`
    Notes string `json:"notes"`
}
```

Allowed `PerturbationKind` values:

```text
alpha_point_perturbation
phi_point_perturbation
lambda_spacing_perturbation
branch_subset_resampling
none
```

Allowed `StabilityStatus` values:

```text
stable_under_small_perturbation
sensitive_to_small_perturbation
unstable_or_ill_conditioned
blocked_by_missing_inputs
blocked_by_nonfinite_values
not_computed
```

Validation must reject nonfinite/negative perturbation magnitudes unless `PerturbationKind = none`.

## Report shape

Generate deterministic JSON:

```json
{
  "schema_version": "bmc0a-residual-audit-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "artifact_kind": "residual_null_comparison_audit",
  "scope": "bmc0a_only",
  "residual_audit_computed": true,
  "recovery_claim": false,
  "scientific_novelty_claim_made": false,
  "bmc_beats_null_models_claim": false,
  "full_bmc_toy_gate": "blocked",
  "local_branch_only": true,
  "global_cosmology_claim": false,
  "source_artifacts": [],
  "comparison_audits": [],
  "stability_audits": [],
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
    "promotion_status": "residual_audit_candidate_only"
  },
  "warnings": [
    "Sprint 11 audits residual/null comparison integrity only.",
    "No recovery claim is made.",
    "No BMC-beats-null-models claim is made.",
    "Full BMC remains blocked.",
    "Convention debts remain unpaid or contested."
  ]
}
```

If source inputs are missing, the report should remain valid but blocked:

```text
residual_audit_computed = false
comparison_audits = []
stability_audits = []
interpretation/status fields must indicate blocked_by_source_unavailable or equivalent
```

## Required gates

Exactly once each:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
residual_audit_scope_gate
comparison_integrity_gate
stability_audit_gate
no_recovery_claim_gate
no_scientific_novelty_claim_gate
no_bmc_beats_null_models_claim_gate
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
residual_audit_computed inconsistent with audit records
recovery_claim = true
scientific_novelty_claim_made = true
bmc_beats_null_models_claim = true
full_bmc_toy_gate != blocked
local_branch_only = false
global_cosmology_claim = true
missing required source artifacts
duplicate source artifacts
unknown source artifacts
file_read without path
file_read without available/read_success status
empty comparison_audits when residual_audit_computed = true
comparison audit with empty metrics_audited
comparison audit with empty target_residual_ids when audit_computed = true
comparison audit with empty null_model_ids when audit_computed = true
forbidden audit_status
forbidden audit_provenance
forbidden interpretation_status
stability audit with nonfinite numeric values
stability audit with negative perturbation magnitude
stability audit with missing baseline/perturbed values when stability_computed = true
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
classical limit achieved
classical cosmology recovered
```

Validation errors must be phrase-safe and must not echo forbidden strings.

## Summary output

`ptw-bmc summarize --report out/bmc0a_residual_audit.json` should show:

```text
BMC Sprint 11 Residual/Null Comparison Audit Summary
Schema Version: bmc0a-residual-audit-v0.1
Scope: bmc0a_only
Residual Audit Computed: true
Recovery Claim: false
Scientific Novelty Claim Made: false
BMC Beats Null Models Claim: false
Full BMC: blocked
Comparison Audits: N
Stability Audits: N
Interpretation Status: diagnostic_audit_only|comparison_integrity_passed|comparison_integrity_failed|comparison_stability_mixed|comparison_unstable|insufficient_target_null_separation|target_null_separation_candidate_unpromoted|blocked_by_missing_residual_inputs|blocked_by_missing_null_inputs|blocked_by_source_unavailable
Promotion Status: residual_audit_candidate_only
```

Do not print recovery/superiority language.

## Tests

Plan tests:

```text
TestResidualAuditReportValidation
TestResidualAuditRejectsRecoveryClaim
TestResidualAuditRejectsScientificNoveltyClaim
TestResidualAuditRejectsBMCBeatsNullModelsClaim
TestResidualAuditRequiresFullBMCBlocked
TestResidualAuditRequiresLocalOnlyBoundary
TestResidualAuditRejectsGlobalCosmologyClaim
TestResidualAuditRequiresSourceArtifacts
TestResidualAuditRejectsDuplicateSourceArtifact
TestResidualAuditRejectsUnknownSourceArtifact
TestResidualAuditRejectsFileReadWithoutPath
TestResidualAuditRejectsDecorativeComparisonAudit
TestResidualAuditRejectsForbiddenAuditStatus
TestResidualAuditRejectsForbiddenAuditProvenance
TestResidualAuditRejectsForbiddenInterpretationStatus
TestResidualAuditRejectsNonfiniteStabilityMetrics
TestResidualAuditRejectsNegativePerturbationMagnitude
TestResidualAuditRequiresAllGatesExactlyOnce
TestResidualAuditRejectsNonPassGate
TestResidualAuditRejectsUnknownFields
TestResidualAuditRejectsTrailingJSONTokens
TestResidualAuditDeterministicJSON
TestResidualAuditForbiddenPhraseScan
TestResidualAuditForbiddenPhraseErrorsArePhraseSafe
TestResidualAuditCLIRoutingRunsGenerateValidateSummarize
TestResidualAuditUnknownProfileFailsAtCLI
```

## Lean policy plan

Add a tiny policy-only Lean file only if it stays small:

```text
BMC/BMC/ResidualAudit.lean
```

It may encode:

```text
toyAnalysisOnly
finalTruthClaim
recoveryClaim
scientificNoveltyClaim
bmcBeatsNullModelsClaim
fullBMCBlocked
faithfulnessContested
```

Allowed theorem purpose:

```text
residualaudit_forbids_recovery_claim
residualaudit_forbids_bmc_beats_nulls_claim
residualaudit_requires_full_bmc_blocked
residualaudit_does_not_imply_residual_success
```

Do not prove physics validity.

Do not prove residual success.

Do not prove null failure.

Do not prove recovery.

## Verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc audit-residuals --profile bmc0a-residual-audit --out out/bmc0a_residual_audit.json
./ptw-bmc validate --report out/bmc0a_residual_audit.json
./ptw-bmc summarize --report out/bmc0a_residual_audit.json
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
  "comparison_audit_plan": [],
  "stability_audit_plan": [],
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
    "sprint11_residual_audit": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP reminder

A comparison audit is not a victory claim.

A stability audit is not recovery.

A local branch diagnostic is not full cosmology.

A null diagnostic is not a failed null model.

Full BMC remains blocked.
