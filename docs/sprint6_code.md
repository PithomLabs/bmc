# BMC Sprint 6 Walkthrough: Friedmann-Residual Specification and Gate Design

Sprint 6 has been successfully implemented, verified, and validated under strict **EBP 2.1** discipline guidelines. All revisions requested by the planning review have been addressed.

## Implemented Components

### 1. Go Package: `internal/bmc/friedmannspec`

- **[contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/contracts.go)**: Establishes Sprint 6 constants for status values and narrow allowed specification scopes (`candidate_specification_only`, `blocked`, `contested`).
- **[mapping.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/mapping.go)**: Logs candidate FRW target maps, validating that convention, normalization, and sign debts remain contested/active.
- **[branch_requirements.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/branch_requirements.go)**: Models local branch relational readiness contracts without implying actual computation readiness.
- **[derivatives.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/derivatives.go)**: Defines methods and exclusion zones (endpoints, turning points, near-node points) for numerical derivatives.
- **[residual_spec.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/residual_spec.go)**: Registers candidate classical relations and null model requirements without computing physics residuals.
- **[gates.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/gates.go)**: Encodes the 10 required policy safety gates.
- **[report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/report.go)**: Generates the deterministic `FriedmannSpecReport` (separated artifact identity matching the `bmc0a-friedmann-spec-v0.1` schema).
- **[validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/validate.go)**: Implements strict schema validation, including:
  - Hard invariants rejecting `residual_computed = true` or `friedmann_recovery_claim = true`.
  - Rejection of unknown JSON fields.
  - Ensuring the EBP ledger requires active status for clock choice, classical target, unit convention, sign convention, normalization, null model, and faithfulness review debts.
- **[friedmannspec_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/friedmannspec_test.go)**: Contains 15 unit tests verifying all spec rules, schema validations, deterministic serialization, and forbidden word scans.

### 2. Command Line Interface: `cmd/ptw-bmc`

- Implemented `spec-friedmann` subcommand to run the profile and output the `FriedmannSpecReport`.
- Updated `validate` and `summarize` subcommands to correctly route and validate the `bmc0a-friedmann-spec-v0.1` schema.

### 3. Lean Formal Verification: `BMC/BMC/FriedmannSpec.lean`

- Implemented `BMCFriedmannSpecReport` structure matching the Go report.
- Proved 12 formal policy safety theorems confirming that passing the Sprint 6 gate strictly forbids residual computation, blocks final truth claims, requires full BMC blocked, requires all EBP debts to remain active, and does not imply Friedmann recovery.

---

## Verification Outcomes

### 1. Go Unit Tests
Run: `go test ./...`
```text
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/friedmannspec    0.003s
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```

### 2. Lean Policy Safety Proofs
Run: `cd BMC && lake build`
```text
Build completed successfully (9 jobs).
```
All Lean formal proofs verified cleanly.

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

The generated report file is located at: [bmc0a_friedmann_spec.json](file:///home/chaschel/Documents/go/bmc/out/bmc0a_friedmann_spec.json).


# BMC Sprint 6: Friedmann-Residual Specification and Gate Design Task List

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
- [x] Generate `walkthrough.md` with results


