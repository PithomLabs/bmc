Approve Sprint 7 planning **with revisions before implementation**. The scope is correct: this sprint should build a **null-model scaffold only**, not run comparisons, compute residuals, or claim BMC beats anything. The plan correctly keeps `needNullModel`, convention debts, clock-choice debt, and full BMC blockage visible. 

## Answers to open questions

For **Phase Resolution**: keep Sprint 7 strictly targeted to **BMC-0A single-scalar minisuperspace superposition**. Do not generalize to multi-field extensions yet.

For **Metric Thresholds**: register metrics only. Do **not** define pass/fail thresholds in Sprint 7. Thresholds should become a later sprint after the null-model scaffold itself is accepted.

## Required revisions before implementation

### 1. Fix promotion status

This line is too strong for a planning document:

```json
"sprint7_nullmodel_spec": "promoted_nullmodel_spec_artifact_after_repairs"
```

Use:

```json
"sprint7_nullmodel_spec": "planned_candidate_only"
```

Maximum future promotion can remain:

```text
promoted_nullmodel_spec_artifact_after_repairs
```

but only after implementation, tests, CLI verification, Lean build, adversarial review, and repairs.

### 2. Fix the warning text

This sentence is also too strong:

```text
The promotion status remains promoted_nullmodel_spec_artifact_after_repairs
```

Replace with:

```text
The promotion status remains planned_candidate_only. Maximum future promotion is promoted_nullmodel_spec_artifact_after_repairs after review and repairs.
```

### 3. Expand verification commands

The verification plan should not only run the package-level tests. Require the full pipeline:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
./ptw-bmc validate --report out/bmc0a_nullmodel_spec.json
./ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

### 4. Add debt fields to Lean scope

Lean should also track the same debt boundaries Sprint 6 protected:

```text
classicalTargetDebtActive
unitConventionDebtActive
signConventionDebtActive
normalizationDebtActive
```

Even though Sprint 7 focuses on null models, those debts still block Friedmann interpretation.

### 5. Add exact-gate cardinality

The validator should require all 10 gates **exactly once**, not merely present:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
no_residual_computation_gate
no_null_comparison_result_gate
null_model_registry_complete_gate
required_before_residual_promotion_gate
friedmann_recovery_claim_blocked_gate
full_bmc_blocked_gate
clock_choice_debt_active_gate
faithfulness_contested_gate
```

### 6. Add forbidden phrase scan

Add a test that scans generated JSON, summary text, and warnings for phrases like:

```text
null models passed
BMC beats null models
Friedmann recovery
ready for Friedmann recovery
classical limit verified
```

## Revised EBP status

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
Sprint 7 plan approved with revisions. Implement only null-model scaffolding, registries, input requirements, future metric contracts, future comparison contracts, safety gates, report validation, CLI routing, and Lean policy contracts. Do not run null models, compute comparisons, compute Friedmann residuals, claim BMC beats null models, or unblock full BMC.
```
