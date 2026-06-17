# BMC Sprint 6 & 6.1 Walkthrough: Friedmann-Residual Specification, Gate Design, and Data-Integrity Repair

Sprint 6 and its repair sprint (Sprint 6.1) have been successfully completed under strict **EBP 2.1** guidelines. No actual physics residual is computed, and full BMC remains blocked.

## Completed Repairs in Sprint 6.1

### 1. Enforced Gate Cardinality and Status Requirements
- **Cardinality Enforced**: The validator now strictly counts and requires all 10 EBP safety gates to appear exactly once in the report's `gates` array. Any missing or duplicate gates trigger a validation failure.
- **Specific Gate Invariants**:
  - `no_residual_computation_gate` must have status `"pass"`.
  - `full_bmc_blocked_gate` must exist and have status `"pass"`.
  - `faithfulness_contested_gate` must exist and have status `"pass"`.
- **Allowed Statuses**: Only `"pass"`, `"blocked"`, and `"contested"` are accepted for gates, preventing any accidental promotion or recovery claims.

### 2. Enforced Null-Model Promotion Prerequisites
- All null model requirements in `null_model_requirements` are strictly checked. The validator rejects the report if any requirement has `required_before_residual_promotion = false`.
- The EBP ledger debt status for `needNullModel` remains `"active"`.

### 3. Added Candidate Map Classical-Target Debt Status
- The `FriedmannCandidateMap` structure in [mapping.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/mapping.go) now includes the explicit `classical_target_debt` status field.
- The validator enforces that this field only has one of the narrow values: `"active"`, `"contested"`, or `"blocked"`. Any empty, missing, retired, validated, recovered, or proved status is strictly rejected.
- The generated BMC-0A map has `classical_target_debt = "active"`.

### 4. Added Lean Sign-Convention Debt Contract
- Updated [FriedmannSpec.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/FriedmannSpec.lean) by adding the `signConventionDebtActive : Bool` field to `BMCFriedmannSpecReport`.
- Updated `reportPassesBMC0AFriedmannSpecGate` to require `signConventionDebtActive = true`.
- Proved the policy-only safety theorem `friedmann_spec_requires_sign_convention_debt_active`. All theorems compile and verify cleanly with `lake build`.

### 5. Added Faithfulness Boundary Validation
- Added validator checks requiring `needFaithfulnessReview = "contested"` and that `faithfulness_contested_gate` exists with status `"pass"`.

### 6. Trailing JSON Token Protection
- Updated the report reader `ReadFriedmannSpecReport` in [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/report.go) to verify EOF (by calling `dec.Decode(&dummy)`) to ensure that any trailing garbage tokens or characters after the primary JSON document are strictly rejected.

---

## Verification Outcomes

### 1. Go Unit Tests
Run: `go test ./...`
```text
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/friedmannspec    0.004s
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```
All **22 tests** in `friedmannspec_test.go` passed successfully, including:
- `TestFriedmannSpecRequiresFullBMCBlockedGate`
- `TestFriedmannSpecRejectsMissingFullBMCBlockedGate`
- `TestFriedmannSpecRejectsNonPassFullBMCBlockedGate`
- `TestFriedmannSpecRequiresNullModelsBeforeResidualPromotion`
- `TestFriedmannSpecCandidateMapRequiresClassicalTargetDebt`
- `TestFriedmannSpecRejectsRetiredClassicalTargetDebt`
- `TestFriedmannSpecRejectsTrailingJSONTokens`

### 2. Lean Policy Safety Proofs
Run: `cd BMC && lake build`
```text
Build completed successfully (9 jobs).
```
All safety proofs verified with zero warnings or errors.

### 3. CLI Demonstration
Run:
```bash
./ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
./ptw-bmc validate --report out/bmc0a_friedmann_spec.json
./ptw-bmc summarize --report out/bmc0a_friedmann_spec.json
```
Output:
```text
Successfully ran Friedmann spec profile 'bmc0a-friedmann-spec' and generated report: out/bmc0a_friedmann_spec.json
Friedmann Spec Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
================================================================================
BMC Sprint 6 Friedmann Spec Report Summary
================================================================================
Schema Version:            bmc0a-friedmann-spec-v0.1
Spec Kind:                 friedmann_residual_specification
Spec Scope:                candidate_specification_only
Residual Computed:         false (Must be false)
Friedmann Recovery Claim:  false (Must be false)
--------------------------------------------------------------------------------
Candidate Maps:
  - Map ID: bmc0a-map-v0.1 (Status: candidate_only)
    Clock: phi (local relational monotonic branches)
--------------------------------------------------------------------------------
Gates:
  - toy_analysis_only_gate                   [pass   ]
  - no_final_truth_claim_gate                [pass   ]
  - local_branch_only_gate                   [pass   ]
  - clock_choice_debt_active_gate            [pass   ]
  - classical_target_candidate_only_gate     [pass   ]
  - unit_convention_debt_gate                [pass   ]
  - null_model_debt_gate                     [pass   ]
  - faithfulness_contested_gate              [pass   ]
  - no_residual_computation_gate             [pass   ]
  - full_bmc_blocked_gate                    [pass   ]
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

The generated report is saved at: [bmc0a_friedmann_spec.json](file:///home/chaschel/Documents/go/bmc/out/bmc0a_friedmann_spec.json).



# BMC Sprint 6 & 6.1 Task List

## Sprint 6: Friedmann-Residual Specification and Gate Design
- [x] Create `internal/bmc/friedmannspec` package
  - [x] Implement EBP contracts and spec constants (`contracts.go`)
  - [x] Implement candidate maps (`mapping.go`)
  - [x] Implement derivative-readiness specs (`derivatives.go`)
  - [x] Implement local branch requirements (`branch_requirements.go`)
  - [x] Implement residual formula registry (`residual_spec.go`)
  - [x] Implement 10 safety gates (`gates.go`)
  - [x] Implement deterministic report structures and generation (`report.go`)
  - [x] Implement strict validator enforcing all EBP rules (`validate.go`)
  - [x] Implement comprehensive test suite (`friedmannspec_test.go`)
- [x] Add CLI subcommand `spec-friedmann` and update routing (`cmd/ptw-bmc/main.go`)
- [x] Add Lean safety contracts (`BMC/BMC/FriedmannSpec.lean`) and import them (`BMC/BMC.lean`)
- [x] Run Go tests and Lake build to verify correctness

## Sprint 6.1: Friedmann Spec Gate/Data-Integrity Repair
- [x] Update validator to enforce exact EBP gate cardinality (each of 10 gates must appear exactly once)
- [x] Enforce that `full_bmc_blocked_gate` status must be `"pass"` in validator
- [x] Enforce `required_before_residual_promotion = true` for all null model requirements in validator
- [x] Add `classical_target_debt` status field to `FriedmannCandidateMap` and generated report
- [x] Enforce that candidate map `classical_target_debt` status must be `"active"`, `"contested"`, or `"blocked"` in validator
- [x] Add `signConventionDebtActive` contract field and theorem `friedmann_spec_requires_sign_convention_debt_active` to Lean safety verification
- [x] Add faithfulness review boundary validation to report validator
- [x] Implement trailing JSON token protection in JSON report readers
- [x] Add 7 new unit tests covering all repairs
- [x] Re-run full Go test and Lean build validation pipelines
- [x] Document repairs in `walkthrough.md`


