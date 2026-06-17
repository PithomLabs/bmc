# Walkthrough: BMC Sprint 9.1 Nullrun Schema and Accounting Repair

This walkthrough details the repairs implemented in Sprint 9.1 to resolve the nullrun schema-integrity and summary accounting issues.

## Files Changed
- [internal/bmc/nullrun/report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/report.go) - Removed `omitempty` from optional float pointers and updated default report EBP debt statuses/vocabulary.
- [internal/bmc/nullrun/validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/validate.go) - Added raw JSON check for optional keys, updated float checks, implemented strict comparison validation, and enforced the adversarial-review EBP debt vocabulary.
- [internal/bmc/nullrun/nullrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/nullrun_test.go) - Appended 18 new unit tests covering Sprint 9.1 reforms.
- [cmd/ptw-bmc/main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go) - Updated CLI summary layout to account for deferred null runs and print exact totals.
- [implementation_plan.md](file:///home/chaschel/.gemini/antigravity-ide/brain/15752b17-b4e0-41bd-bcbb-1bcaaa65743f/implementation_plan.md) - Documented approved plan.
- [task.md](file:///home/chaschel/.gemini/antigravity-ide/brain/15752b17-b4e0-41bd-bcbb-1bcaaa65743f/task.md) - Verified task checklist.

---

## Validator Repairs
- Added check for raw JSON key presence under `diagnostics` to ensure `min_amplitude_r`, `max_abs_q_away_from_nodes`, and `max_phase_gradient` are always present (rejecting omitted keys).
- Updated float value validation to reject negative values and sentinel-like values such as `-1`.
- Added strict EBP debt status validation checking all fields against allowed values: `unpaid`, `partial`, `retired`, `contested`, `overclaimed`, `absent`, and `null_model_runner_candidate_only`.

---

## Report Schema Repairs
- Removed `omitempty` tag from pointer fields `MinAmplitudeR`, `MaxAbsQAwayFromNodes`, and `MaxPhaseGradient` in `NullDiagnostics`.
- Changed `EbpDebtVocabulary` to `"ptw_adversarial_review_debt_status_v0.1"`.
- Set default EBP debt fields to use review vocabulary (e.g. `unpaid`, `partial`, `absent`, `contested`) instead of `active`.

---

## Explicit Null Metric Output
Unavailable optional float metrics are serialized as explicit JSON `null` when nil.
Example snippet from `out/bmc0a_nullrun.json`:
```json
        "min_amplitude_r": null,
        "max_abs_q_away_from_nodes": null,
        "max_phase_gradient": null
```

---

## Comparison Validation Repairs
- Rejects decorative comparison records when `TargetNullComparisonComputed` is `true`.
- Requires `metrics_compared` and `null_model_ids` to be non-empty.
- Ensures all referenced `null_model_ids` exist and have status `diagnostics_generated` (rejecting blocked/deferred/unknown runs).
- Rejects `winner / outperformed / validated / recovered / proved` language in comparison reasons and statuses.
- Verifies that at least one computed comparison satisfies honest constraints.

---

## EBP Debt Vocabulary Repair
- Enforces `ebp_debt_vocabulary == ptw_adversarial_review_debt_status_v0.1`.
- Restricts all debt status values to `unpaid`, `partial`, `retired`, `contested`, `overclaimed`, `absent`.
- Restricts `promotion_status` to `null_model_runner_candidate_only`.

---

## Summary Accounting Repair
Summarizer now counts and displays deferred runs and totals:
```text
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
Null Models Deferred: 2
Null Models Accounted For: 7/7
```

---

## New/Updated Test Names
- `TestNullRunRejectsComparisonWithNoMetrics`
- `TestNullRunRejectsComparisonWithNoNullModelIDs`
- `TestNullRunRejectsComparisonReferencingBlockedNullRun`
- `TestNullRunRejectsComparisonReferencingDeferredNullRun`
- `TestNullRunRejectsComparisonReferencingUnknownNullRun`
- `TestNullRunRejectsNegativeMinAmplitude`
- `TestNullRunRejectsNegativeMaxAbsQ`
- `TestNullRunRejectsNegativeMaxPhaseGradient`
- `TestNullRunRejectsSentinelNegativeFloatMetrics`
- `TestNullRunUnavailableMetricsEmitExplicitNull`
- `TestNullRunClassicalReferenceUnavailableMetricsEmitNull`
- `TestNullRunRejectsMissingOptionalMetricKeys`
- `TestNullRunRejectsOutOfVocabularyDebtStatus`
- `TestNullRunAcceptsReviewDebtVocabulary`
- `TestNullRunRequiresCandidateOnlyPromotionStatus`
- `TestNullRunAccountsForAllSevenNullModelsInSummary`
- `TestNullRunSummaryIncludesDeferredCount`
- `TestNullRunSummaryIncludesAccountedForCount`

---

## Go Test Result
All package tests compile and pass successfully:
```bash
$ go test ./...
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/friedmannspec    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/nullrun  0.373s
ok      github.com/PithomLabs/bmc/internal/bmc/nullspec (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/priorart 0.371s
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```

---

## CLI Generation Output
```bash
$ ./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
Successfully ran null models profile 'bmc0a-nullrun' and generated report: out/bmc0a_nullrun.json
```

---

## CLI Validation Output
```bash
$ ./ptw-bmc validate --report out/bmc0a_nullrun.json
Null Model Run Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
```

---

## CLI Summary Output
```bash
$ ./ptw-bmc summarize --report out/bmc0a_nullrun.json
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
Null Models With Diagnostics: 4
Null Models Blocked: 1
Null Models Deferred: 2
Null Models Accounted For: 7/7
Interpretation Status: diagnostic_comparison_only
Promotion Status: null_model_runner_candidate_only
```

---

## Lean Build Result
Lean policy verification builds successfully:
```bash
$ cd BMC && /home/chaschel/.elan/bin/lake build
Build completed successfully (12 jobs).
```

---

## Remaining Limitations
- Sprint 9.1 is a diagnostic comparator only; it does not compute Friedmann residuals or support physics-limit recovery assertions.
- The comparison does not claim BMC beats null models or outperformed controls. Full BMC promotion remains blocked.
