Sprint 9 remains **accept_with_repairs**. Do **not** accept or promote it yet.

The review found real schema-integrity issues, not physics-overclaim issues. That is good news: the sprint stayed inside the safety boundary, but the report/validator still allow weak or decorative data. The right next move is **Sprint 9.1: Nullrun Schema and Accounting Repair**.

```json id="s9_review_status"
{
  "sprint": "BMC Sprint 9",
  "verdict": "accept_with_repairs",
  "accepted": false,
  "next_step": "Sprint 9.1 Nullrun Schema and Accounting Repair",
  "main_issue": "validator/report allow decorative comparisons, omitted null metrics, sentinel metric values, debt-vocabulary mismatch, and opaque run counts",
  "physics_boundary": {
    "residual_computed": false,
    "recovery_claim": false,
    "superiority_claim": false,
    "full_bmc_toy_gate": "blocked"
  }
}
```

Use this repair prompt.

Implement **BMC Sprint 9.1: Nullrun Schema and Accounting Repair**.

Current status: `accept_with_repairs`.

Do not start Sprint 10 yet.

Sprint 9 remains a BMC-0A null-model diagnostic runner only.

Do not compute a Friedmann residual.

Do not claim recovery.

Do not claim BMC beats null models.

Do not promote full BMC.

## Required repairs

### 1. Reject decorative comparison records

Patch:

```text
internal/bmc/nullrun/validate.go
```

Current issue:

```text
When target_null_comparison_computed = true, validation accepts comparison records with empty metrics_compared or empty/unusable null_model_ids.
```

Repair rule:

If:

```text
target_null_comparison_computed = true
```

then at least one `TargetNullDiagnosticComparison` must satisfy:

```text
comparison_computed = true
metrics_compared is nonempty
null_model_ids is nonempty
each null_model_id exists in null_model_runs
each referenced null_model_id has run_status = diagnostics_generated
interpretation_status is allowed and non-victory
```

Reject:

```text
empty metrics_compared
empty null_model_ids
null_model_ids that reference blocked/deferred runs
null_model_ids that do not exist
comparison_computed = false when report-level comparison is true
winner / outperformed / validated / recovered / proved language
```

Add tests:

```text
TestNullRunRejectsComparisonWithNoMetrics
TestNullRunRejectsComparisonWithNoNullModelIDs
TestNullRunRejectsComparisonReferencingBlockedNullRun
TestNullRunRejectsComparisonReferencingDeferredNullRun
TestNullRunRejectsComparisonReferencingUnknownNullRun
```

### 2. Reject sentinel float metrics

Patch:

```text
internal/bmc/nullrun/validate.go
```

Current issue:

```text
Validator rejects nonfinite floats but still accepts finite sentinel values such as -1.
```

For optional numeric metrics:

```text
min_amplitude_r
max_abs_q_away_from_nodes
max_phase_gradient
```

validation must reject:

```text
NaN
+Inf
-Inf
negative values
sentinel-like unavailable values such as -1
```

Because these are nonnegative diagnostic magnitudes or amplitudes.

Unavailable values must be represented as JSON null, not a sentinel.

Add tests:

```text
TestNullRunRejectsNegativeMinAmplitude
TestNullRunRejectsNegativeMaxAbsQ
TestNullRunRejectsNegativeMaxPhaseGradient
TestNullRunRejectsSentinelNegativeFloatMetrics
```

### 3. Preserve unavailable metrics as explicit JSON null

Patch:

```text
internal/bmc/nullrun/report.go
```

Current issue:

```go
MinAmplitudeR *float64 `json:"min_amplitude_r,omitempty"`
MaxAbsQAwayFromNodes *float64 `json:"max_abs_q_away_from_nodes,omitempty"`
MaxPhaseGradient *float64 `json:"max_phase_gradient,omitempty"`
```

Remove `omitempty`.

Use:

```go
MinAmplitudeR *float64 `json:"min_amplitude_r"`
MaxAbsQAwayFromNodes *float64 `json:"max_abs_q_away_from_nodes"`
MaxPhaseGradient *float64 `json:"max_phase_gradient"`
```

The generated JSON must emit unavailable metrics as:

```json
"min_amplitude_r": null,
"max_abs_q_away_from_nodes": null,
"max_phase_gradient": null
```

not omit them.

Add tests:

```text
TestNullRunUnavailableMetricsEmitExplicitNull
TestNullRunClassicalReferenceUnavailableMetricsEmitNull
```

Important: because Go pointer fields decode both missing and null as nil, add a raw JSON presence check if needed. The validator should reject generated or user-supplied JSON where these fields are missing from a diagnostics object.

Possible approach:

```text
Before or during validation, inspect the raw JSON object for each diagnostics record and require the keys:
min_amplitude_r
max_abs_q_away_from_nodes
max_phase_gradient
```

Add test:

```text
TestNullRunRejectsMissingOptionalMetricKeys
```

### 4. Use accepted EBP debt status values or declare runtime vocabulary correctly

Current issue:

```text
Generated ebp_debt fields use active, but the adversarial review vocabulary uses unpaid, partial, retired, contested, overclaimed, absent.
```

Choose one of two repairs.

Preferred for Sprint 9.1:

```text
Use the adversarial-review debt vocabulary in the Sprint 9 report.
```

Suggested generated debt values:

```json
{
  "needLiteratureAudit": "partial",
  "needMap": "partial",
  "needInvariant": "partial",
  "needToyCheck": "unpaid",
  "needNullModel": "partial",
  "needObstruction": "partial",
  "needFaithfulnessReview": "contested",
  "clock_choice_debt": "unpaid",
  "classical_target_debt": "unpaid",
  "unit_convention_debt": "unpaid",
  "sign_convention_debt": "unpaid",
  "normalization_debt": "unpaid",
  "containsFinalTruthClaim": "absent",
  "promotion_status": "null_model_runner_candidate_only"
}
```

Validation must reject debt values outside:

```text
unpaid
partial
retired
contested
overclaimed
absent
null_model_runner_candidate_only
```

For `promotion_status`, allow only:

```text
null_model_runner_candidate_only
```

Add tests:

```text
TestNullRunRejectsOutOfVocabularyDebtStatus
TestNullRunAcceptsReviewDebtVocabulary
TestNullRunRequiresCandidateOnlyPromotionStatus
```

Alternative repair, only if you intentionally retain runtime labels:

```text
Keep ebp_debt_vocabulary = ptw_runtime_debt_status_v0.1 and document that active is a runtime label, not adversarial-review classification.
```

But since the review explicitly objected, prefer the first approach for Sprint 9.

### 5. Account for deferred null models in the summary

Patch:

```text
cmd/ptw-bmc/main.go
```

Current issue:

```text
Summary shows 7 registered, 4 with diagnostics, 1 blocked, leaving 2 unaccounted.
```

Summary must explicitly show:

```text
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
Null Models Deferred: 2
Null Models Accounted For: 7/7
```

Validation or summary test should verify that:

```text
diagnostics_generated_count + blocked_count + deferred_count = total_null_model_runs
```

Add tests:

```text
TestNullRunAccountsForAllSevenNullModelsInSummary
TestNullRunSummaryIncludesDeferredCount
TestNullRunSummaryIncludesAccountedForCount
```

### 6. Re-run forbidden phrase and phrase-safe error protections

Ensure the new repairs do not introduce forbidden wording.

Forbidden phrases remain:

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

Validation errors must remain phrase-safe and must not echo forbidden strings.

Add or update tests if needed:

```text
TestNullRunForbiddenPhraseScan
TestNullRunForbiddenPhraseErrorsArePhraseSafe
```

## Required verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
./ptw-bmc validate --report out/bmc0a_nullrun.json
./ptw-bmc summarize --report out/bmc0a_nullrun.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also verify deterministic output generation.

## Return walkthrough

Return:

```text
files changed
validator repairs
report schema repairs
explicit null metric output
comparison validation repairs
EBP debt vocabulary repair
summary accounting repair
new/updated test names
Go test result
CLI generation output
CLI validation output
CLI summary output
Lean build result
remaining limitations
```

## Expected status after Sprint 9.1

Before final review:

```text
null_model_runner_candidate_after_repairs
```

Maximum after review confirms repairs:

```text
promoted_null_model_runner_artifact_after_repairs
```

## Forbidden scope

Do not implement:

```text
Friedmann residual runner
recovery logic
BMC superiority claims
winner/loser ranking
paper ingestion
claim extraction
general physics profile system
dashboard
full BMC promotion
```

Remember:

```text
A diagnostic comparison is not a victory claim.
Explicit null metrics are not physical success.
A deterministic fixture is not a promoted physics result.
Full BMC remains blocked.
```

After this repair passes, Sprint 9 can be accepted as a **null-model runner artifact**, but `needNullModel` should still be only **partial**, not retired. The runner exists; it does not yet justify residual interpretation or full BMC promotion.
