Use this for **Sprint 7 planning**. It keeps the next sprint focused on **null-model scaffolding**, not Friedmann residual computation.

# BMC Sprint 7 Planning: Null-Model Scaffold for Future Friedmann Residuals

You are planning **BMC Sprint 7** under **EBP 2.1** discipline.

Do **not** compute a Friedmann residual yet.

Produce a reviewable implementation plan only.

## Current accepted artifacts

The following artifacts are accepted under strict EBP limits:

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

Sprint 6.1 established:

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

Sprint 7 must respond to Sprint 6’s requirement that null-model debt be active and unpaid before any future Friedmann-residual promotion.

## Sprint 7 goal

Plan a narrow **null-model scaffold sprint**.

Sprint 7 should answer:

```text
What null-model reports, schemas, gates, and comparison contracts must exist before a future Friedmann-style residual can be interpreted as meaningful rather than accidental, generic, or artifact-driven?
```

Sprint 7 is not a residual computation sprint.

Sprint 7 is not a null-model result sprint.

Sprint 7 is a scaffold, registry, and validation sprint.

## Core question

Before any future Friedmann residual can be trusted, Sprint 7 must define:

```text
1. Which null models are required?
2. What each null model is meant to falsify or control for?
3. Which data each null model needs from prior artifacts?
4. Which metrics would be compared in a future sprint?
5. Which comparisons are forbidden in Sprint 7?
6. What gates block Friedmann residual interpretation until null-model results exist?
7. What report schema records the null-model requirements without pretending they are already run?
```

## Forbidden scope

Do **not** plan or implement:

```text
Actual Friedmann residual computation
Friedmann residual recovery
Numerical null-model comparison results
Claim that BMC beats null models
Claim that any null model failed
Claim that Friedmann recovery is supported
Ready-for-Friedmann claim
Full BMC promotion
Full quantum gravity
Proof of Bohmian mechanics
Solution to the problem of time
Spacetime emergence proof
Massive scalar model
LQC comparison
Page-Wootters comparison
Inhomogeneous perturbations
Black holes
Fermions
Gauge fields
Lorentz recovery
```

The full BMC toy gate remains blocked.

`needNullModel` remains active.

`clock_choice_debt` remains active.

`needFaithfulnessReview` remains contested.

## Required Sprint 7 design

### 1. Name the sprint correctly

Use:

```text
BMC Sprint 7: Null-Model Scaffold for Future Friedmann Residuals
```

Do not use:

```text
Null Models Passed
Friedmann Residual Validation
Friedmann Recovery Controls
Classical Limit Verified
```

### 2. New package

Plan a new package:

```text
internal/bmc/nullspec/
```

Possible files:

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

The package should not run null-model simulations or produce comparison winners.

It should define required null models, inputs, metrics, and future comparison gates.

### 3. Null-model registry

Plan a null-model record:

```go
type NullModelSpec struct {
    NullModelID string `json:"null_model_id"`
    Name string `json:"name"`
    Purpose string `json:"purpose"`

    ControlsFor []string `json:"controls_for"`
    RequiredInputs []string `json:"required_inputs"`
    RequiredMetrics []string `json:"required_metrics"`

    RequiredBeforeResidualPromotion bool `json:"required_before_residual_promotion"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Required null models:

```text
constant_phase_control
randomized_phase_control
matched_amplitude_randomized_phase_control
classical_frw_reference_trajectory
same_branch_segmentation_under_null_wavefunctions
node_neighborhood_stress_case
clock_choice_alternative_branch_diagnostic
```

Allowed `Status` values:

```text
planned
deferred
blocked
```

Do not allow:

```text
passed
failed
outperformed
validated
recovered
proved
```

Each null model must have:

```text
required_before_residual_promotion = true
```

### 4. Required inputs

Plan an input contract:

```go
type NullModelInputRequirement struct {
    InputID string `json:"input_id"`
    SourceArtifact string `json:"source_artifact"`
    RequiredFor []string `json:"required_for"`

    AvailabilityStatus string `json:"availability_status"`
    Reason string `json:"reason"`
}
```

Allowed `AvailabilityStatus` values:

```text
available
planned
deferred
blocked
```

Required source artifacts should include:

```text
bmc0a_superposition_safe
bmc0a_superposition_robustness
bmc0a_clock_fragility
bmc0a_clock_readiness
bmc0a_friedmann_spec
```

Sprint 7 may declare required inputs, but it must not treat missing inputs as results.

### 5. Future metric contracts

Plan future metric records:

```go
type NullModelMetricContract struct {
    MetricID string `json:"metric_id"`
    Description string `json:"description"`

    AppliesTo []string `json:"applies_to"`
    RequiredBeforeResidualPromotion bool `json:"required_before_residual_promotion"`

    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Candidate future metrics:

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

Allowed `Status` values:

```text
planned
deferred
blocked
```

Do not compute metric values in Sprint 7.

### 6. Future comparison contracts

Plan comparison contracts:

```go
type FutureNullComparisonContract struct {
    ComparisonID string `json:"comparison_id"`
    BaselineArtifact string `json:"baseline_artifact"`
    NullModelIDs []string `json:"null_model_ids"`
    Metrics []string `json:"metrics"`

    ComparisonComputed bool `json:"comparison_computed"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Required invariant:

```text
comparison_computed = false
```

Allowed `Status` values:

```text
planned
deferred
blocked
```

Forbidden:

```text
pass
fail
winner
outperformed
validated
recovered
```

### 7. Gate design

Plan a gate record:

```go
type NullSpecGate struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Required gates:

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

Allowed `Status` values:

```text
pass
blocked
contested
```

Important:

```text
no_null_comparison_result_gate passes only if no null-model comparison result is computed.
required_before_residual_promotion_gate passes only if all required null models are marked required_before_residual_promotion = true.
full_bmc_blocked_gate passes only if full BMC remains blocked.
```

### 8. Proposed CLI

Plan:

```bash
ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
```

`validate` and `summarize` should route by schema version:

```text
bmc0a-nullmodel-spec-v0.1
```

### 9. Report shape

Plan deterministic JSON:

```json
{
  "schema_version": "bmc0a-nullmodel-spec-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "spec_kind": "null_model_scaffold",
  "spec_scope": "null_model_specification_only",
  "residual_computed": false,
  "null_comparison_computed": false,
  "friedmann_recovery_claim": false,
  "source_artifacts": [
    "Sprint 1: BMC-0A plane-wave control artifact",
    "Sprint 2: BMC-0A two-plane-wave superposition control artifact",
    "Sprint 2: BMC-0A node-obstruction detection artifact",
    "Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
    "Sprint 4: BMC-0A clock-fragility diagnostic artifact",
    "Sprint 5: BMC-0A clock-readiness/local segmentation artifact",
    "Sprint 6: BMC-0A Friedmann-residual specification/gate-design artifact"
  ],
  "null_models": [],
  "input_requirements": [],
  "metric_contracts": [],
  "future_comparison_contracts": [],
  "gates": [],
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked",
    "reason": "Sprint 7 defines null-model scaffolding only. No null-model comparison results or Friedmann residuals are computed."
  },
  "ebp_debt": {
    "needMap": "active",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "active",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "classical_target_debt": "active",
    "unit_convention_debt": "active",
    "sign_convention_debt": "active",
    "normalization_debt": "active",
    "containsFinalTruthClaim": "absent",
    "LeanVerification": "planned",
    "promotion_status": "planned_candidate_only"
  },
  "warnings": [
    "Sprint 7 is a null-model scaffold sprint only.",
    "No Friedmann residual was computed.",
    "No null-model comparison result was computed.",
    "No Friedmann recovery is claimed.",
    "Null-model debt remains active.",
    "Full BMC remains blocked."
  ]
}
```

### 10. Validation requirements

Plan strict validation rejecting:

```text
final_truth_claim = true
toy_analysis_only = false
schema_version missing or wrong
spec_kind wrong
spec_scope wrong
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
null model status = passed|failed|validated|recovered|proved
null_model.required_before_residual_promotion != true
empty metric_contracts
metric status outside planned|deferred|blocked
future_comparison_contract.comparison_computed = true
missing no_null_comparison_result_gate
no_null_comparison_result_gate.status != pass
missing full_bmc_blocked_gate
full_bmc_blocked_gate.status != pass
missing required_before_residual_promotion_gate
required_before_residual_promotion_gate.status != pass
warnings missing “No null-model comparison result was computed”
warnings missing “No Friedmann residual was computed”
unknown JSON fields
nonfinite numeric values
```

Use strict decoding:

```go
json.Decoder.DisallowUnknownFields()
```

Also verify EOF after decode to reject trailing JSON tokens.

### 11. Lean planning

Do not formalize null-model physics yet.

Add only policy/safety contracts:

```text
BMC/BMC/NullModelSpec.lean
```

Possible structure:

```lean
structure BMCNullModelSpecReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  residualComputed : Bool
  nullComparisonComputed : Bool
  friedmannRecoveryClaim : Bool
  fullBMCBlocked : Bool
  nullModelDebtActive : Bool
  clockChoiceDebtActive : Bool
  faithfulnessContested : Bool
```

Possible theorems:

```text
null_model_spec_requires_toy_only
null_model_spec_blocks_final_truth
null_model_spec_forbids_residual_computation
null_model_spec_forbids_null_comparison_result
null_model_spec_forbids_friedmann_recovery_claim
null_model_spec_requires_full_bmc_blocked
null_model_spec_requires_null_model_debt_active
null_model_spec_requires_clock_choice_debt_active
null_model_spec_requires_faithfulness_contested
null_model_spec_does_not_imply_friedmann_recovery
null_model_spec_does_not_imply_full_bmc
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

### 12. Test plan

Plan tests for:

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

## Required output

Return a plan in this JSON shape:

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
  "proof_obligations": [],
  "null_models": [],
  "risks": [],
  "human_review_questions": [],
  "ebp_debt_status": {
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
    "LeanVerification": ""
  },
  "promotion_status": {
    "sprint7_nullmodel_spec": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP guardrails

The maximum allowed Sprint 7 promotion is:

```text
promoted_nullmodel_spec_artifact_after_repairs
```

Do not recommend promotion as:

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
A specification is not a computation.
A residual formula candidate is not a residual result.
Null-model debt remains active.
Clock-choice debt remains active.
Full BMC remains blocked.
```
