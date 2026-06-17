Sprint 6 is **not accepted yet**. The adversarial review is a clean `accept_with_repairs`: the no-computation boundary mostly held, but the validator and Lean contract are not strict enough yet.

The important distinction:

```text id="ovpg2j"
promotion_recommendation = promoted_friedmann_spec_artifact_after_repairs
current_status = candidate_only_pending_repairs
```

So the next move is **Sprint 6.1 data-integrity and policy-gate repair**, not Sprint 7.

## Key findings to accept

The reviewer’s concrete blockers are valid:

```text id="pwry02"
1. Validator must require full_bmc_blocked_gate by name and require status = pass.
2. Validator must require every null_model_requirement.required_before_residual_promotion = true.
3. FriedmannCandidateMap must explicitly track classical_target_debt.
4. Lean must include signConventionDebtActive and theorem friedmann_spec_requires_sign_convention_debt_active.
5. Tests and walkthrough must be updated to match actual validator behavior.
```

Optional but worth doing:

```text id="cy0gm2"
- Require all 10 gate names exactly once.
- Require faithfulness_contested_gate by name/status.
- Add EOF checking after JSON decode to reject trailing tokens.
```

## Current EBP status

```json id="f2h9as"
{
  "sprint": "BMC Sprint 6",
  "overall_verdict": "accept_with_repairs",
  "current_acceptance_status": "not_accepted_yet",
  "promotion_status": "friedmann_spec_candidate_only_pending_repairs",
  "maximum_future_promotion": "promoted_friedmann_spec_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked_required_but_validator_needs_repair",
  "residual_computed": "false_reported",
  "friedmann_recovery_claim": "false_reported",
  "needNullModel": "active_unpaid",
  "clock_choice_debt": "active_unpaid",
  "classical_target_debt": "active_unpaid_but_map_needs_explicit_field",
  "sign_convention_debt": "active_unpaid_but_lean_needs_contract",
  "faithfulness": "contested"
}
```

Use this repair prompt.

Sprint 6 adversarial review received. Do not start Sprint 7.

Current status: `accept_with_repairs`.

Implement **Sprint 6.1: Friedmann Spec Gate/Data-Integrity Repair** only.

## Required repairs before Sprint 6 acceptance

### 1. Enforce `full_bmc_blocked_gate`

Update `internal/bmc/friedmannspec/validate.go`.

Validation must require:

```text
gate name: full_bmc_blocked_gate
gate status: pass
```

Reports must fail validation if:

```text
full_bmc_blocked_gate is missing
full_bmc_blocked_gate.status != pass
full_bmc_blocked_gate appears more than once, if exact gate cardinality is enforced
```

Add tests:

```text
TestFriedmannSpecRequiresFullBMCBlockedGate
TestFriedmannSpecRejectsMissingFullBMCBlockedGate
TestFriedmannSpecRejectsNonPassFullBMCBlockedGate
```

### 2. Enforce null-model promotion prerequisites

Every `null_model_requirements[]` item must have:

```text
required_before_residual_promotion = true
```

Validation must reject any null model requirement where this is false.

Add test:

```text
TestFriedmannSpecRequiresNullModelsBeforeResidualPromotion
```

Also keep:

```text
needNullModel = active
null_model_debt = active
```

Do not mark null-model debt retired.

### 3. Add explicit `classical_target_debt` to candidate maps

Update `internal/bmc/friedmannspec/mapping.go`.

`FriedmannCandidateMap` must explicitly include a field such as:

```go
ClassicalTargetDebt string `json:"classical_target_debt"`
```

Allowed values should be narrow, preferably:

```text
active
contested
blocked
```

For the generated BMC-0A map, set:

```text
classical_target_debt = active
```

Validation must reject a candidate map with missing, empty, retired, validated, recovered, or proved classical-target status.

Add tests:

```text
TestFriedmannSpecCandidateMapRequiresClassicalTargetDebt
TestFriedmannSpecRejectsRetiredClassicalTargetDebt
```

### 4. Add missing Lean sign-convention debt contract

Update:

```text
BMC/BMC/FriedmannSpec.lean
```

Add field to `BMCFriedmannSpecReport`:

```lean
signConventionDebtActive : Bool
```

Add theorem:

```lean
friedmann_spec_requires_sign_convention_debt_active
```

The theorem should remain policy-only. It must not prove Friedmann recovery, FRW validity, or any physics result.

Run Lake build after repair.

### 5. Tighten gate cardinality if easy

Prefer enforcing all 10 required gates exactly once:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
local_branch_only_gate
clock_choice_debt_active_gate
classical_target_candidate_only_gate
unit_convention_debt_gate
null_model_debt_gate
faithfulness_contested_gate
no_residual_computation_gate
full_bmc_blocked_gate
```

Each required gate must be present exactly once.

Status requirements:

```text
no_residual_computation_gate = pass
full_bmc_blocked_gate = pass
faithfulness_contested_gate = pass
```

No gate should imply recovery, validation, or full BMC promotion.

### 6. Validate faithfulness boundary

Prefer adding validation for:

```text
needFaithfulnessReview = contested
faithfulness_contested_gate exists
faithfulness_contested_gate.status = pass
```

This keeps faithfulness review debt visible.

### 7. Add JSON trailing-token protection

In report readers, after decoding with `DisallowUnknownFields`, verify EOF so trailing JSON tokens are rejected.

Add test:

```text
TestFriedmannSpecRejectsTrailingJSONTokens
```

### 8. Update walkthrough claims

Update `sprint6_code.md` or walkthrough text so it does not overstate validator coverage.

Only claim strict validation covers a gate/debt after the validator actually enforces it.

## Required verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
./ptw-bmc validate --report out/bmc0a_friedmann_spec.json
./ptw-bmc summarize --report out/bmc0a_friedmann_spec.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Return a concise repair report with:

```text
files changed
validator repairs completed
new/updated tests
exact gate requirements enforced
classical_target_debt map field added
Lean sign-convention theorem added
command results
updated Sprint 6 status
```

## Forbidden scope

Do not implement:

```text
Friedmann residual computation
Friedmann recovery
ready-for-Friedmann claim
full BMC promotion
massive scalar model
LQC comparison
Page-Wootters comparison
full quantum gravity claim
proof of Bohmian mechanics
solution to the problem of time
```

Sprint 6 remains a specification artifact only.

After these repairs pass, Sprint 6 can be accepted as:

```text id="h88q6o"
promoted_friedmann_spec_artifact_after_repairs
```

Until then, it remains:

```text id="8qaj37"
friedmann_spec_candidate_only_pending_repairs
```
