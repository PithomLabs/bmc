# BMC Sprint 7 & 7.1 Walkthrough: Null-Model Scaffold and Phrase-Boundary Repair

Sprint 7 and its repair sprint (Sprint 7.1) have been successfully completed under strict **EBP 2.1** guidelines. No actual physics comparisons or residual computations are performed, and full BMC remains blocked.

## Completed Repairs in Sprint 7.1

### 1. Gate Cardialities and Status Validation
- Enforces that all 10 EBP safety gates appear exactly once in the report's `gates` array.
- Enforces that every gate has status `"pass"`. If any gate is missing, duplicated, blocked, contested, unknown, or empty, the report strictly fails validation.

### 2. Metric Prerequisite Validation
- Every registered metric in `metric_contracts` must have `required_before_residual_promotion = true`. The validator rejects the report if any metric has this flag set to `false`.

### 3. Future Comparison Contracts Validation
- Enforces that `future_comparison_contracts` is nonempty. Reports with empty or null future comparison contracts fail validation.
- All contracts must have `comparison_computed = false` and status strictly in `"planned"`, `"deferred"`, or `"blocked"`.

### 4. Phrase-Boundary and Semantic Repairs
- Removed all bypass phrasings (e.g. "recovery of Friedmann", "Friedmann recovery claim") from source code, JSON structure, warnings, summary outputs, and comments.
- Replaced them with neutral wording such as `"No recovery claim is made"` and `"No recovery-style interpretation is allowed."`
- Expanded the forbidden phrase scanner to catch:
  - `recovery of Friedmann`
  - `Friedmann support`
  - `Friedmann-compatible`
  - `control victory`
  - `nulls defeated`
  - `BMC beats null models`

### 5. Exact Seven-Null-Model Registry
- Enforces that only the 7 authorized null models are registered. Any unknown or extra null model IDs in the report trigger validation failures.

### 6. Subprocess CLI Routing Test
- Rewrote `TestNullModelSpecCLIRouting` to compile the `ptw-bmc` CLI binary, execute report generation, test unknown profile failures, and verify that actual validation and summarization dispatch correctly route the `bmc0a-nullmodel-spec-v0.1` schema.

### 7. Lean Verification Debt Label
- Updated report wording from `LeanVerification: planned` to `LeanVerification: retired_for_policy_safety_contracts` since policy safety proofs verify cleanly in Lean.

---

## Verification Outcomes

### 1. Go Unit Tests
Run: `go test ./...`
```text
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/friedmannspec    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/nullspec 0.211s
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```
All **29 tests** in `nullspec_test.go` passed successfully.

### 2. Lean Policy Safety Proofs
Run: `cd BMC && lake build`
```text
Build completed successfully (10 jobs).
```
All safety proofs verified with zero warnings or errors.

### 3. CLI Demonstration
Run:
```bash
./ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
./ptw-bmc validate --report out/bmc0a_nullmodel_spec.json
./ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json
```
Output:
```text
Successfully ran Null Model spec profile 'bmc0a-nullmodel-spec' and generated report: out/bmc0a_nullmodel_spec.json
Null Model Spec Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
================================================================================
BMC Sprint 7 Null-Model Spec Report Summary
================================================================================
Schema Version:            bmc0a-nullmodel-spec-v0.1
Spec Kind:                 null_model_scaffold
Spec Scope:                null_model_specification_only
Residual Computed:         false (Must be false)
Null Comparison Computed:  false (Must be false)
Recovery Claim:            false (Must be false)
--------------------------------------------------------------------------------
Null Model specs registered:
  - ID: constant_phase_control                       Status: [planned]
  - ID: randomized_phase_control                     Status: [planned]
  - ID: matched_amplitude_randomized_phase_control   Status: [planned]
  - ID: classical_frw_reference_trajectory           Status: [planned]
  - ID: same_branch_segmentation_under_null_wavefunctions Status: [planned]
  - ID: node_neighborhood_stress_case                Status: [planned]
  - ID: clock_choice_alternative_branch_diagnostic   Status: [planned]
--------------------------------------------------------------------------------
Gates:
  - toy_analysis_only_gate                   [pass   ]
  - no_final_truth_claim_gate                [pass   ]
  - no_residual_computation_gate             [pass   ]
  - no_null_comparison_result_gate           [pass   ]
  - null_model_registry_complete_gate        [pass   ]
  - required_before_residual_promotion_gate  [pass   ]
  - friedmann_recovery_claim_blocked_gate    [pass   ]
  - full_bmc_blocked_gate                    [pass   ]
  - clock_choice_debt_active_gate            [pass   ]
  - faithfulness_contested_gate              [pass   ]
--------------------------------------------------------------------------------
EBP Debt Ledger:
  needMap:                 active
  needNullModel:           active
  clock_choice_debt:       active
  classical_target_debt:   active
  unit_convention_debt:    active
  sign_convention_debt:    active
  normalization_debt:      active
  promotion_status:        planned_candidate_only
================================================================================
```

The generated report is saved at: [bmc0a_nullmodel_spec.json](file:///home/chaschel/Documents/go/bmc/out/bmc0a_nullmodel_spec.json).

# BMC Sprint 7 & 7.1 Task List

## Sprint 7: Null-Model Scaffold for Future Friedmann Residuals
- [x] Create `internal/bmc/nullspec` package
  - [x] Implement EBP contracts and spec constants (`contracts.go`)
  - [x] Implement null-model specification registry (`registry.go`)
  - [x] Implement input requirements contract (`inputs.go`)
  - [x] Implement future metric contracts (`metrics.go`)
  - [x] Implement future comparison contracts (`comparison_contracts.go`)
  - [x] Implement 10 EBP safety gates (`gates.go`)
  - [x] Implement deterministic report structures and generation (`report.go`)
  - [x] Implement strict validator enforcing all EBP rules (`validate.go`)
  - [x] Implement 17 comprehensive unit tests (`nullspec_test.go`)
- [x] Add CLI subcommand `spec-nullmodels` and update validation/summarization routing (`cmd/ptw-bmc/main.go`)
- [x] Add Lean safety contracts (`BMC/BMC/NullModelSpec.lean`) and import them (`BMC/BMC.lean`)
- [x] Run Go tests and Lake build to verify correctness

## Sprint 7.1: Nullspec Validator and Phrase-Boundary Repair
- [x] Update validator to require all 10 safety gates exactly once and status must be `"pass"`
- [x] Enforce that every metric contract must have `required_before_residual_promotion = true` in validator
- [x] Enforce that future comparison contracts are nonempty in validator
- [x] Remove any phrase-scan bypass `"recovery of Friedmann"` from JSON structures, summary outputs, warnings, and comments
- [x] Expand forbidden phrase tests to catch `"recovery of Friedmann"`, `"Friedmann support"`, `"Friedmann-compatible"`, `"control victory"`, `"nulls defeated"`, and `"BMC beats null models"`
- [x] Enforce exact seven-null-model registry (reject unknown extra null model IDs)
- [x] Enhance CLI routing test `TestNullModelSpecCLIRouting` to compile the binary and test actual command dispatching and unknown profiles
- [x] Update report EBP debt LeanVerification label to `retired_for_policy_safety_contracts`
- [x] Strengthen Lean verification by adding explicit safety theorems for `classicalTargetDebtActive`, `unitConventionDebtActive`, `signConventionDebtActive`, and `normalizationDebtActive`
- [x] Re-run full Go test and Lean build validation pipelines
- [x] Document repairs in `walkthrough.md`
