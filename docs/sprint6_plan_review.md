Approve Sprint 6 planning **with revisions before implementation**.

The scope is correct: Sprint 6 should be a **Friedmann-residual specification and gate-design sprint**, not a residual computation sprint. This is the right next move after Sprint 5, because Sprint 5 only established `friedmann_readiness = local_only_candidate`, kept `clock_choice_debt = active`, and kept full BMC blocked. 

## Required revisions before implementation

### 1. Fix the report type name

This line is likely a copy/paste error:

```text
Assembles the deterministic ClockReadinessReport matching the bmc0a-friedmann-spec-v0.1 schema.
```

Use:

```text
Assembles the deterministic FriedmannSpecReport matching the bmc0a-friedmann-spec-v0.1 schema.
```

Do not reuse `ClockReadinessReport` for Sprint 6. Sprint 6 needs a separate artifact identity.

### 2. Add full verification commands

The verification plan should include:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
./ptw-bmc validate --report out/bmc0a_friedmann_spec.json
./ptw-bmc summarize --report out/bmc0a_friedmann_spec.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Running only `go test ./internal/bmc/friedmannspec` is too narrow because CLI routing and existing validators could regress.

### 3. Make `residual_computed = false` a hard invariant

The validator must reject any report where:

```json
"residual_computed": true
```

The same applies to:

```json
"friedmann_recovery_claim": true
```

This is the central Sprint 6 safety gate.

### 4. Add a report-level field for candidate-only scope

Add something like:

```json
"spec_scope": "candidate_specification_only"
```

Allowed values should be narrow:

```text
candidate_specification_only
blocked
contested
```

Do not allow:

```text
ready
validated
recovered
```

### 5. Keep branch requirements as readiness contracts only

`FriedmannBranchRequirement` should not imply that a branch is ready for actual residual computation. Use:

```text
branch_residual_readiness: candidate_only|blocked|contested
```

Avoid:

```text
ready
pass
validated
```

### 6. Add explicit convention-debt fields

Make sure the report and validator require these to remain active:

```text
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
null_model_debt
faithfulness_review_debt
```

Your plan lists most of these, but `sign_convention_debt` should be elevated to the EBP debt ledger too, not only inside candidate maps.

### 7. Add tests for forbidden words in report output

Add a test that rejects or scans for dangerous phrases in generated report, summary, and warnings:

```text
ready for Friedmann recovery
recovers Friedmann
Friedmann residual passes
classical cosmology recovered
full BMC validated
```

This catches accidental drift.

## Approved decisions

```json
{
  "new_package": "internal/bmc/friedmannspec approved",
  "new_cli": "spec-friedmann approved",
  "schema": "bmc0a-friedmann-spec-v0.1 approved",
  "candidate_maps": "approved as candidate_only/contested/blocked",
  "branch_requirements": "approved as readiness contracts only",
  "derivative_readiness": "approved as specification only",
  "residual_formula_registry": "approved only if no numerical residual is evaluated",
  "null_model_requirements": "approved as required future debt",
  "lean_scope": "approved policy/safety only",
  "friedmann_residual_computation": "forbidden in Sprint 6",
  "full_bmc_gate": "must remain blocked"
}
```

## EBP status entering Sprint 6

```text
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: active
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
containsFinalTruthClaim: absent
LeanVerification: planned
promotion_status: planned_candidate_only
```

## Final instruction

```text
Sprint 6 plan approved with revisions. Implement only Friedmann-residual specification, candidate variable maps, derivative-readiness contracts, null-model requirements, and safety gates. Do not compute any Friedmann residual, do not claim Friedmann recovery, and do not unblock full BMC.
```
