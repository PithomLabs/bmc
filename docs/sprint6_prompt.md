# BMC Sprint 6 Planning: Friedmann-Residual Specification and Gate Design

You are planning **BMC Sprint 6** under **EBP 2.1** discipline.

Do **not** implement Friedmann residual recovery yet.

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
```

Sprint 5 established:

```text
friedmann_readiness: local_only_candidate
clock_choice_debt: active
full_bmc_toy_gate: blocked
local_only_candidate is not Friedmann recovery readiness
```

This means local φ-branches can be inspected, but this does **not** mean Friedmann residual recovery is allowed yet.

## Sprint 6 goal

Plan a narrow **Friedmann-residual specification and gate-design sprint**.

Sprint 6 should answer:

```text
What exactly would a Friedmann-style residual mean on a local relational branch, and what gates must be passed before any future residual computation is allowed?
```

Sprint 6 is not a recovery sprint.

Sprint 6 is a specification, mapping, and gating sprint.

## Core question

Before computing any Friedmann residual, Sprint 6 must define:

```text
1. What is the candidate classical target equation?
2. What variables in BMC-0A are being compared to that target?
3. Which clock is used, and only on which local branches?
4. What derivatives are required?
5. What units/sign conventions are assumed?
6. What counts as residual failure, contested result, or readiness-only result?
7. What null models must later be compared?
8. What would block the residual from being computed?
```

## Forbidden scope

Do **not** plan or implement:

```text
Actual Friedmann residual recovery
Claim that Friedmann equations are recovered
Claim that BMC validates classical cosmology
Claim that local branches are ready for Friedmann recovery
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

`clock_choice_debt` remains active.

`needNullModel` remains partial/deferred.

`needFaithfulnessReview` remains contested.

## Required Sprint 6 design

### 1. Name the sprint correctly

Use:

```text
BMC Sprint 6: Friedmann-Residual Specification and Gate Design
```

Do not use:

```text
Friedmann Recovery
Friedmann Validation
Friedmann Residual Pass
Classical Limit Recovered
```

### 2. Define candidate residual contracts, not results

Plan a new package such as:

```text
internal/bmc/friedmannspec/
```

Possible files:

```text
contracts.go
mapping.go
derivatives.go
branch_requirements.go
residual_spec.go
gates.go
report.go
validate.go
friedmannspec_test.go
```

The package should not compute a final pass/fail Friedmann recovery result.

It should define what would be required for a future residual computation.

### 3. Candidate classical target

For the massless scalar minisuperspace toy, identify the candidate classical target only as a **candidate target**, not a confirmed target.

Possible target family:

```text
flat FRW + massless scalar
```

Possible target relation, stated only as a candidate:

```text
H^2 ≈ C * rho_phi
```

where:

```text
H = dα/dt_candidate
rho_phi depends on scalar kinetic energy under chosen convention
C is convention-dependent
```

Do not hardcode this as true physics for BMC yet.

The plan must explicitly track:

```text
normalization_debt
unit_convention_debt
clock_choice_debt
sign_convention_debt
classical_target_debt
```

### 4. Candidate map from BMC variables to classical variables

Plan a mapping record:

```go
type FriedmannCandidateMap struct {
    MapID string `json:"map_id"`

    AlphaMeaning string `json:"alpha_meaning"`
    PhiMeaning string `json:"phi_meaning"`
    ClockVariable string `json:"clock_variable"`
    ClockScope string `json:"clock_scope"`

    CandidateScaleFactor string `json:"candidate_scale_factor"`
    CandidateHubbleDefinition string `json:"candidate_hubble_definition"`
    CandidateEnergyDensityDefinition string `json:"candidate_energy_density_definition"`

    UnitConventionStatus string `json:"unit_convention_status"`
    SignConventionStatus string `json:"sign_convention_status"`
    NormalizationStatus string `json:"normalization_status"`
    ClockChoiceDebt string `json:"clock_choice_debt"`

    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Allowed `Status` values:

```text
candidate_only
contested
blocked
```

Do not allow:

```text
validated
recovered
proved
```

### 5. Local branch requirements

Use Sprint 5 local branches as input.

Plan a branch-requirement record:

```go
type FriedmannBranchRequirement struct {
    BranchID string `json:"branch_id"`
    SourceConfigID string `json:"source_config_id"`

    PhiLocalMonotonic bool `json:"phi_local_monotonic"`
    AlphaPhiSingleValued bool `json:"alpha_phi_single_valued"`
    MinBranchSamples int `json:"min_branch_samples"`
    ActualSamples int `json:"actual_samples"`

    ClockRange float64 `json:"clock_range"`
    LambdaRange float64 `json:"lambda_range"`

    DerivativeReady bool `json:"derivative_ready"`
    DerivativeMethod string `json:"derivative_method"`
    DerivativeDebt string `json:"derivative_debt"`

    NodeContactFree bool `json:"node_contact_free"`
    QFiniteAwayFromNodes bool `json:"q_finite_away_from_nodes"`

    BranchResidualReadiness string `json:"branch_residual_readiness"`
    Reason string `json:"reason"`
}
```

Allowed `BranchResidualReadiness` values:

```text
blocked
candidate_only
contested
```

Do not allow:

```text
ready
pass
recovered
```

### 6. Derivative-readiness specification

A future Friedmann residual would need derivatives on local branches.

Plan derivative-readiness checks, but do not compute a final Friedmann residual.

Track:

```text
dα/dφ stability
d²α/dφ² availability if needed
finite difference sensitivity
branch endpoint exclusion
turning-point exclusion
node-neighborhood exclusion
minimum samples after exclusions
```

Plan a derivative record:

```go
type DerivativeReadinessCheck struct {
    BranchID string `json:"branch_id"`
    Method string `json:"method"`

    ExcludesEndpoints bool `json:"excludes_endpoints"`
    ExcludesTurningPointNeighborhoods bool `json:"excludes_turning_point_neighborhoods"`
    ExcludesNearNodePoints bool `json:"excludes_near_node_points"`

    MinSamplesRequired int `json:"min_samples_required"`
    SamplesAvailable int `json:"samples_available"`

    StepSensitivityStatus string `json:"step_sensitivity_status"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Allowed `Status` values:

```text
candidate_only
blocked
contested
```

### 7. Residual formula candidate registry

Plan a registry of candidate residual formulas.

Example:

```go
type ResidualFormulaCandidate struct {
    FormulaID string `json:"formula_id"`
    Description string `json:"description"`

    ClassicalTarget string `json:"classical_target"`
    RequiredVariables []string `json:"required_variables"`
    RequiredDerivatives []string `json:"required_derivatives"`

    ConventionDebt []string `json:"convention_debt"`
    NullModelsRequired []string `json:"null_models_required"`

    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Allowed `Status` values:

```text
candidate_only
blocked
contested
```

Do not evaluate the formula numerically in Sprint 6.

### 8. Null-model planning

Sprint 6 should not perform null-model comparison yet, but must specify the null models required before any future Friedmann residual can be promoted.

Required null-model ledger:

```text
constant-phase control
randomized phase control
matched amplitude / randomized phase control
classical FRW reference trajectory
same branch segmentation under null wavefunctions
node-neighborhood stress case
clock-choice alternative branch diagnostic
```

Plan a null-model requirement record:

```go
type FriedmannNullModelRequirement struct {
    NullModelID string `json:"null_model_id"`
    Purpose string `json:"purpose"`
    RequiredBeforeResidualPromotion bool `json:"required_before_residual_promotion"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Allowed `Status` values:

```text
planned
deferred
blocked
```

### 9. Gate design

Plan a gate record:

```go
type FriedmannSpecGate struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Required gates:

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

Allowed `Status` values:

```text
pass
blocked
contested
```

Important: `no_residual_computation_gate` should pass only if Sprint 6 does **not** compute a residual.

### 10. Proposed CLI

Plan:

```bash
ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
```

`validate` and `summarize` should route by schema version:

```text
bmc0a-friedmann-spec-v0.1
```

### 11. Report shape

Plan deterministic JSON:

```json
{
  "schema_version": "bmc0a-friedmann-spec-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "spec_kind": "friedmann_residual_specification",
  "source_artifacts": [
    "Sprint 1: BMC-0A plane-wave control artifact",
    "Sprint 2: BMC-0A two-plane-wave superposition control artifact",
    "Sprint 2: BMC-0A node-obstruction detection artifact",
    "Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
    "Sprint 4: BMC-0A clock-fragility diagnostic artifact",
    "Sprint 5: BMC-0A clock-readiness/local segmentation artifact"
  ],
  "residual_computed": false,
  "friedmann_recovery_claim": false,
  "candidate_maps": [],
  "branch_requirements": [],
  "derivative_readiness_checks": [],
  "residual_formula_candidates": [],
  "null_model_requirements": [],
  "gates": [],
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked",
    "reason": "Sprint 6 defines residual specifications only. No residual recovery is computed."
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
    "normalization_debt": "active",
    "containsFinalTruthClaim": "absent",
    "LeanVerification": "planned",
    "promotion_status": "planned_candidate_only"
  },
  "warnings": [
    "Sprint 6 is a specification sprint only.",
    "No Friedmann residual was computed.",
    "No Friedmann recovery is claimed.",
    "Local branch readiness is not Friedmann recovery readiness.",
    "Full BMC remains blocked."
  ]
}
```

### 12. Validation requirements

Plan strict validation rejecting:

```text
final_truth_claim = true
toy_analysis_only = false
schema_version missing or wrong
spec_kind wrong
residual_computed = true
friedmann_recovery_claim = true
promotion_gate.status != blocked
clock_choice_debt != active
classical_target_debt != active
unit_convention_debt != active
normalization_debt != active
needNullModel not active
missing no_residual_computation_gate
no_residual_computation_gate.status != pass
invalid candidate map status
candidate map status = validated
candidate map status = recovered
invalid branch residual readiness
branch residual readiness = ready/pass/recovered
invalid residual formula status
empty null_model_requirements
warnings missing “No Friedmann residual was computed”
warnings missing “No Friedmann recovery is claimed”
unknown JSON fields
nonfinite numeric values
```

Use strict decoding:

```go
json.Decoder.DisallowUnknownFields()
```

### 13. Lean planning

Do not formalize Friedmann physics yet.

Add only policy/safety contracts:

```text
BMC/BMC/FriedmannSpec.lean
```

Possible structure:

```lean
structure BMCFriedmannSpecReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  residualComputed : Bool
  friedmannRecoveryClaim : Bool
  fullBMCBlocked : Bool
  clockChoiceDebtActive : Bool
  classicalTargetDebtActive : Bool
  unitConventionDebtActive : Bool
  normalizationDebtActive : Bool
  nullModelDebtActive : Bool
  faithfulnessContested : Bool
```

Possible theorems:

```text
friedmann_spec_requires_toy_only
friedmann_spec_blocks_final_truth
friedmann_spec_forbids_residual_computation
friedmann_spec_forbids_recovery_claim
friedmann_spec_requires_full_bmc_blocked
friedmann_spec_requires_clock_choice_debt_active
friedmann_spec_requires_classical_target_debt_active
friedmann_spec_requires_unit_convention_debt_active
friedmann_spec_requires_normalization_debt_active
friedmann_spec_requires_null_model_debt_active
friedmann_spec_does_not_imply_friedmann_recovery
friedmann_spec_does_not_imply_full_bmc
```

Lean remains policy/safety only.

No theorem may claim:

```text
Friedmann equations recovered
classical cosmology recovered
BMC validates FRW
Bohmian cosmology proven
```

### 14. Test plan

Plan tests for:

```text
TestFriedmannSpecReportValidation
TestFriedmannSpecRejectsResidualComputed
TestFriedmannSpecRejectsRecoveryClaim
TestFriedmannSpecRequiresFullBMCBlocked
TestFriedmannSpecRequiresClockChoiceDebtActive
TestFriedmannSpecRequiresClassicalTargetDebtActive
TestFriedmannSpecRequiresUnitConventionDebtActive
TestFriedmannSpecRequiresNormalizationDebtActive
TestFriedmannSpecRequiresNullModelDebtActive
TestFriedmannSpecRejectsValidatedCandidateMap
TestFriedmannSpecRejectsReadyBranchResidualReadiness
TestFriedmannSpecRejectsEmptyNullModelRequirements
TestFriedmannSpecRejectsUnknownFields
TestFriedmannSpecDeterministicJSON
TestFriedmannSpecCLIRouting
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
    "normalization_debt": "",
    "containsFinalTruthClaim": "",
    "LeanVerification": ""
  },
  "promotion_status": {
    "sprint6_friedmann_spec": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP guardrails

The maximum allowed Sprint 6 promotion is:

```text
promoted_friedmann_spec_artifact_after_repairs
```

Do not recommend promotion as:

```text
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
A specification is not a computation.
A residual formula candidate is not a residual result.
A local branch is not a global clock.
A local-only candidate is not Friedmann recovery readiness.
Null-model debt remains active.
Clock-choice debt remains active.
Full BMC remains blocked.
```
