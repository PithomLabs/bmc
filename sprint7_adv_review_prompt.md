You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 7: Null-Model Scaffold for Future Friedmann Residuals**.

Sprint 7 follows the accepted Sprint 6 + 6.1 artifact:

```text
residual_computed: false
friedmann_recovery_claim: false
spec_scope: candidate_specification_only
full_bmc_toy_gate: blocked
needNullModel: active
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
needFaithfulnessReview: contested
```

Sprint 7 must remain a **null-model scaffold sprint only**.

It must not run null-model simulations, compute comparison results, compute Friedmann residuals, or imply that BMC outperforms any null model.

## Current accepted context

Accepted artifacts under strict EBP limits:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A Friedmann-residual specification/gate-design artifact
Sprint 6.1: BMC-0A Friedmann-spec data-integrity/gate repair artifact
```

Sprint 7 claims to add:

```text
internal/bmc/nullspec/
spec-nullmodels CLI subcommand
bmc0a-nullmodel-spec-v0.1 report schema
null-model registry
input requirement registry
future metric contracts
future comparison contracts
10 policy safety gates
NullModelSpec.lean policy/safety contracts
```

## Forbidden claims

The artifact must **not** claim:

```text
null models passed
null models failed
BMC beats null models
BMC outperforms controls
Friedmann residual recovery
Friedmann equations recovered
ready for Friedmann recovery
classical cosmology recovered
BMC validates FRW
full BMC validation
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

The full BMC toy gate must remain blocked.

The report must state:

```text
residual_computed: false
null_comparison_computed: false
friedmann_recovery_claim: false
spec_scope: null_model_specification_only
promotion_status: planned_candidate_only
```

The following debts must remain active or contested:

```text
needNullModel
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
needFaithfulnessReview
```

## Materials to review

Review actual source and generated artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
internal/bmc/nullspec/**
BMC/**
BMC/BMC/NullModelSpec.lean
out/bmc0a_nullmodel_spec.json
out/bmc0a_friedmann_spec.json
out/bmc0a_clock_readiness.json
out/bmc0a_clock_fragility.json
out/bmc0a_superposition_robustness.json
go.mod
all Go tests
Sprint 1 final report
Sprint 2 final report
Sprint 3 final report
Sprint 4 final report
Sprint 5 final report
Sprint 6 final report
Sprint 7 walkthrough/report
```

If any expected file or artifact is missing, report it as debt.

## Expected Sprint 7 design

Expected package:

```text
internal/bmc/nullspec/
```

Expected files:

```text
contracts.go
registry.go
inputs.go
metrics.go
comparison_contracts.go
gates.go
report.go
validate.go
nullspec_test.go
```

Expected CLI:

```bash
ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
ptw-bmc validate --report out/bmc0a_nullmodel_spec.json
ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json
```

Expected schema:

```text
bmc0a-nullmodel-spec-v0.1
```

Expected Lean file:

```text
BMC/BMC/NullModelSpec.lean
```

## Central review principle

Sprint 7 may define **future null-model requirements**, but:

```text
a null-model scaffold is not a null-model result
a future comparison contract is not a comparison result
a metric contract is not a measured metric value
a specification is not a computation
a residual formula candidate is not a residual result
null-model debt remains active
clock-choice debt remains active
full BMC remains blocked
```

The key review question is:

```text
Does Sprint 7 safely define the null-model scaffold required for future Friedmann-residual interpretation, without computing results or implying BMC beats any null model?
```

Mark as blocker if any code path computes null-model comparison results, marks a null model passed/failed, computes a Friedmann residual, or unblocks full BMC.

## Required review targets

### 1. Report identity and schema

Review `report.go`, `validate.go`, CLI output, and generated JSON.

Correct behavior:

```text
schema_version = bmc0a-nullmodel-spec-v0.1
spec_kind = null_model_scaffold
spec_scope = null_model_specification_only
residual_computed = false
null_comparison_computed = false
friedmann_recovery_claim = false
```

Mark as blocker if:

```text
ClockReadinessReport or FriedmannSpecReport is reused as the Sprint 7 artifact identity
spec_scope allows ready/validated/recovered
residual_computed can become true
null_comparison_computed can become true
friedmann_recovery_claim can become true
```

### 2. Null-model registry

Review `registry.go`.

Expected null models, exactly once each:

```text
constant_phase_control
randomized_phase_control
matched_amplitude_randomized_phase_control
classical_frw_reference_trajectory
same_branch_segmentation_under_null_wavefunctions
node_neighborhood_stress_case
clock_choice_alternative_branch_diagnostic
```

Each null model must include:

```text
null_model_id
name
purpose
controls_for
required_inputs
required_metrics
required_before_residual_promotion = true
status = planned|deferred|blocked
reason
```

Forbidden null-model statuses:

```text
passed
failed
winner
outperformed
validated
recovered
proved
ready
pass
```

Mark as blocker if any null model is treated as already run.

### 3. Input requirements

Review `inputs.go`.

Expected source artifacts include:

```text
bmc0a_superposition_safe
bmc0a_superposition_robustness
bmc0a_clock_fragility
bmc0a_clock_readiness
bmc0a_friedmann_spec
```

Allowed availability statuses:

```text
available
planned
deferred
blocked
```

Forbidden:

```text
validated
proved
result_passed
comparison_complete
```

Check that missing inputs are treated as availability/debt, not as null-model results.

### 4. Metric contracts

Review `metrics.go`.

Expected candidate metric contracts include:

```text
branch_count_stability
turning_point_stability
local_relation_single_valuedness_rate
derivative_readiness_rate
node_contact_rate
min_amplitude_distribution
max_abs_q_distribution
phase_gradient_distribution
residual_formula_input_availability
```

Each metric contract must include:

```text
metric_id
description
applies_to
required_before_residual_promotion = true
status = planned|deferred|blocked
reason
```

Metric contracts must not include measured values, pass/fail thresholds, winners, or interpretations.

Mark as blocker if Sprint 7 computes metric values.

### 5. Future comparison contracts

Review `comparison_contracts.go`.

Expected:

```text
FutureNullComparisonContract
```

Every future comparison contract must have:

```text
comparison_computed = false
status = planned|deferred|blocked
```

Forbidden status values:

```text
pass
fail
winner
outperformed
validated
recovered
proved
ready
```

Mark as blocker if any comparison result is computed or implied.

### 6. Gate design

Review `gates.go`.

Expected gates, exactly once each:

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

Allowed gate statuses:

```text
pass
blocked
contested
```

Important:

```text
no_residual_computation_gate passes only if no Friedmann residual is computed
no_null_comparison_result_gate passes only if no null-model comparison result is computed
required_before_residual_promotion_gate passes only if every required null model is marked required_before_residual_promotion = true
full_bmc_blocked_gate passes only if full BMC remains blocked
```

Mark as blocker if any required gate is missing, duplicated, or permissive.

### 7. Strict validation

Review `ValidateNullModelSpecReport`.

It should reject:

```text
final_truth_claim = true
toy_analysis_only = false
missing or wrong schema_version
wrong spec_kind
wrong spec_scope
residual_computed = true
null_comparison_computed = true
friedmann_recovery_claim = true
promotion_gate.status != blocked
needNullModel != active
clock_choice_debt != active
classical_target_debt != active
unit_convention_debt != active
sign_convention_debt != active
normalization_debt != active
needFaithfulnessReview != contested
missing required null models
duplicate null_model_id
null model status outside planned|deferred|blocked
null model status = passed|failed|validated|recovered|proved|winner|outperformed
null_model.required_before_residual_promotion != true
empty metric_contracts
metric status outside planned|deferred|blocked
metric contract with computed numeric result
future_comparison_contract.comparison_computed = true
future comparison status = pass|fail|winner|outperformed|validated|recovered|proved|ready
missing no_null_comparison_result_gate
no_null_comparison_result_gate.status != pass
missing no_residual_computation_gate
no_residual_computation_gate.status != pass
missing full_bmc_blocked_gate
full_bmc_blocked_gate.status != pass
missing required_before_residual_promotion_gate
required_before_residual_promotion_gate.status != pass
warnings missing “No null-model comparison result was computed”
warnings missing “No Friedmann residual was computed”
unknown JSON fields
trailing JSON tokens
nonfinite numeric values
```

Check for strict decoding:

```go
json.Decoder.DisallowUnknownFields()
```

Also check EOF after decode to reject trailing JSON tokens.

### 8. Deterministic JSON

Check whether repeated generation produces byte-identical JSON.

Review deterministic ordering for:

```text
source_artifacts
null_models
input_requirements
metric_contracts
future_comparison_contracts
gates
warnings
```

### 9. CLI behavior

Run or review reported runs:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
./ptw-bmc validate --report out/bmc0a_nullmodel_spec.json
./ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json
cd BMC && lake build
```

Check whether:

```text
spec-nullmodels subcommand exists
unknown spec-nullmodels profiles fail safely
validate routes null-model spec schema to nullspec validator
summarize routes null-model spec schema to nullspec summarizer
ordinary BMC, robustness, clock-fragility, clock-readiness, and Friedmann-spec validators are not weakened
errors are explicit
```

### 10. Lean safety contracts

Review:

```text
BMC/BMC/NullModelSpec.lean
BMC/BMC.lean
```

Check that:

```text
lake build succeeds
no sorry/admit exists
BMCNullModelSpecReport exists
null_model_spec_requires_toy_only exists
null_model_spec_blocks_final_truth exists
null_model_spec_forbids_residual_computation exists
null_model_spec_forbids_null_comparison_result exists
null_model_spec_forbids_friedmann_recovery_claim exists
null_model_spec_requires_full_bmc_blocked exists
null_model_spec_requires_null_model_debt_active exists
null_model_spec_requires_clock_choice_debt_active exists
null_model_spec_requires_faithfulness_contested exists
null_model_spec_does_not_imply_friedmann_recovery exists
null_model_spec_does_not_imply_full_bmc exists
```

Also check whether Lean tracks these debts, either in this Sprint 7 file or inherited from Sprint 6 boundary logic:

```text
classicalTargetDebtActive
unitConventionDebtActive
signConventionDebtActive
normalizationDebtActive
```

Lean remains policy/safety only.

No theorem may claim:

```text
null model passed
BMC beats null models
Friedmann equations recovered
classical cosmology recovered
BMC validates FRW
Bohmian cosmology proven
```

### 11. Forbidden phrase and semantic scan

Search source, tests, comments, report text, CLI output, summaries, README, and walkthrough for dangerous phrases:

```text
null models passed
null models failed
BMC beats null models
BMC outperforms controls
Friedmann recovery
ready for Friedmann recovery
classical limit verified
classical cosmology recovered
FRW recovered
full BMC validated
validates quantum gravity
solves the problem of time
proves Bohmian mechanics
derives spacetime
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Pay special attention to phrase-scan bypasses, such as:

```text
recovery of Friedmann
Friedmann support
Friedmann-compatible
control victory
nulls defeated
```

Acceptable phrasing:

```text
null-model scaffold
null-model specification only
future null-model requirement
future metric contract
future comparison contract
no null-model comparison result was computed
no Friedmann residual was computed
no recovery claim is made
full BMC remains blocked
null-model debt remains active
clock-choice debt remains active
```

Mark as blocker if wording dodges the phrase scan while implying a forbidden claim.

### 12. Test coverage

Confirm tests include at least:

```text
TestNullModelSpecReportValidation
TestNullModelSpecRejectsResidualComputed
TestNullModelSpecRejectsNullComparisonComputed
TestNullModelSpecRejectsRecoveryClaim
TestNullModelSpecRequiresAllRequiredNullModels
TestNullModelSpecRejectsDuplicateNullModelIDs
TestNullModelSpecRequiresBeforeResidualPromotion
TestNullModelSpecRejectsPassedFailedNullModelStatus
TestNullModelSpecRequiresMetricContracts
TestNullModelSpecRejectsComputedFutureComparison
TestNullModelSpecRequiresNoNullComparisonResultGate
TestNullModelSpecRequiresFullBMCBlockedGate
TestNullModelSpecRequiresNullModelDebtActive
TestNullModelSpecRejectsUnknownFields
TestNullModelSpecRejectsTrailingJSONTokens
TestNullModelSpecDeterministicJSON
TestNullModelSpecCLIRouting
```

If the walkthrough claims 18 tests but task list claims 17, verify actual test count and names.

## Code review checks

Review these details:

```text
nullspec package structure
spec_scope allowed values
residual_computed hard false invariant
null_comparison_computed hard false invariant
friedmann_recovery_claim hard false invariant
all 7 required null models present exactly once
all null models required_before_residual_promotion = true
null model statuses planned/deferred/blocked only
metric contracts present and not evaluated
future comparison contracts present and not computed
10 safety gates present exactly once
strict schema validation
unknown-field rejection
trailing-token rejection
forbidden semantic phrase scan
deterministic output ordering
CLI routing by schema version
full_bmc_toy_gate blocked
needNullModel active
clock_choice_debt active
classical_target_debt active
unit_convention_debt active
sign_convention_debt active
normalization_debt active
needFaithfulnessReview contested
no external dependencies added
```

## EBP debt classification

Classify each item as one of:

```text
unpaid
partial
retired
contested
overclaimed
absent
```

Debt items:

```text
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
containsFinalTruthClaim
LeanVerification
NullModelSpecDiagnosticIntegrity
NullComparisonBoundary
NoResidualComputationBoundary
FriedmannRecoveryBoundary
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
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
    "containsFinalTruthClaim": "",
    "LeanVerification": "",
    "NullModelSpecDiagnosticIntegrity": "",
    "NullComparisonBoundary": "",
    "NoResidualComputationBoundary": "",
    "FriedmannRecoveryBoundary": ""
  },
  "nullspec_findings": [],
  "physics_boundary_findings": [],
  "code_findings": [],
  "cli_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|nullmodel_spec_candidate_only|promoted_nullmodel_spec_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 7 passes perfectly, the maximum allowed recommendation is:

```text
promoted_nullmodel_spec_artifact_after_repairs
```

Never recommend promotion as:

```text
null models passed
BMC beats null models
Friedmann recovery
ready for Friedmann recovery
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Remember:

```text
A null-model scaffold is not a null-model result.
A future comparison contract is not a comparison result.
A metric contract is not a measured metric value.
A specification is not a computation.
Null-model debt remains active.
Clock-choice debt remains active.
Full BMC remains blocked.
```
