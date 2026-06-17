# BMC Sprint 5 Planning: Relational-Clock Readiness and Local Clock Segmentation

You are planning **BMC Sprint 5** under **EBP 2.1** discipline.

Do **not** implement code yet. Produce a reviewable implementation plan only.

## Current accepted artifacts

The following artifacts are accepted under strict EBP limits:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
```

Sprint 4 found:

```text
diagnostic_outcome: clock_fragile
trajectory_valid: pass
phi_clock_valid: fail
distinction_preserved: true
clock_choice_debt: active
```

This means the toy trajectory can remain numerically well-formed while φ fails as a global monotonic relational clock.

## Sprint 5 goal

Plan a narrow readiness sprint:

```text
Determine whether BMC-0A can define safe relational diagnostics when φ is not globally monotonic.
```

Sprint 5 should answer:

```text
1. Can φ be used on local monotonic branches?
2. Where are the branch boundaries / turning points?
3. Can α(φ) be extracted safely only on valid local segments?
4. Are there clock-independent diagnostics we can use before any Friedmann-style residual?
5. Should Friedmann residual remain blocked until clock-readiness debt is partially paid?
```

Sprint 5 is a **relational-clock readiness sprint**, not a Friedmann recovery sprint.

## Forbidden scope

Do **not** plan or implement:

```text
Friedmann residual recovery
massive scalar field
LQC comparison
Page-Wootters comparison
inhomogeneous perturbations
black holes
fermions
gauge fields
Lorentz recovery
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
claim that φ is invalid in full quantum cosmology
claim that α is the correct replacement clock
```

The full BMC toy gate must remain blocked.

Friedmann residual remains deferred.

Faithfulness for full BMC remains contested.

`clock_choice_debt` remains active.

## Required diagnostic questions

Sprint 5 should answer:

```text
1. How many local monotonic φ-branches exist in each fragile trajectory?
2. What are the λ ranges of each branch?
3. Are branch boundaries associated with dφ/dλ ≈ 0, sign changes, high phase gradient, high Q, or low amplitude R?
4. Can α(φ) be constructed safely on each branch?
5. Are branch-level diagnostics stable under step refinement?
6. Does α provide a local candidate clock on regions where φ fails?
7. Can the report distinguish:
   - global φ-clock failure,
   - local φ-clock usability,
   - branch-level invalidity,
   - and fully invalid trajectories?
8. What diagnostics remain meaningful without choosing φ as a global clock?
```

## Required Sprint 5 design

### 1. New package

Plan a new package:

```text
internal/bmc/clockseg/
```

Possible files:

```text
segments.go
branches.go
turning_points.go
local_relations.go
clock_independent.go
report.go
validate.go
```

### 2. Monotonic segment detection

Define branch/segment records:

```go
type ClockSegment struct {
    SegmentID       string  `json:"segment_id"`
    StartIndex      int     `json:"start_index"`
    EndIndex        int     `json:"end_index"`
    StartLambda     float64 `json:"start_lambda"`
    EndLambda       float64 `json:"end_lambda"`
    StartAlpha      float64 `json:"start_alpha"`
    EndAlpha        float64 `json:"end_alpha"`
    StartPhi        float64 `json:"start_phi"`
    EndPhi          float64 `json:"end_phi"`

    ClockVariable   string `json:"clock_variable"`
    MonotonicStatus string `json:"monotonic_status"`
    Direction       string `json:"direction"`

    MinAmplitudeR   *float64 `json:"min_amplitude_r"`
    MinAmplitudeRStatus string `json:"min_amplitude_r_status,omitempty"`
    MinAmplitudeRReason string `json:"min_amplitude_r_reason,omitempty"`

    MaxAbsQ         *float64 `json:"max_abs_q"`
    MaxAbsQStatus   string `json:"max_abs_q_status,omitempty"`
    MaxAbsQReason   string `json:"max_abs_q_reason,omitempty"`

    MaxPhaseGradient *float64 `json:"max_phase_gradient"`
    MaxPhaseGradientStatus string `json:"max_phase_gradient_status,omitempty"`
    MaxPhaseGradientReason string `json:"max_phase_gradient_reason,omitempty"`

    NearNodeContact bool `json:"near_node_contact"`
    SegmentValidity string `json:"segment_validity"`
    Reason          string `json:"reason"`
}
```

Allowed `monotonic_status` values:

```text
pass
fail
contested
```

Allowed `direction` values:

```text
increasing
decreasing
flat
mixed
```

Allowed `segment_validity` values:

```text
valid
invalid
contested
```

### 3. Turning-point detection

Plan turning-point records:

```go
type ClockTurningPoint struct {
    Index        int     `json:"index"`
    Lambda      float64 `json:"lambda"`
    Alpha       float64 `json:"alpha"`
    Phi         float64 `json:"phi"`
    DPhiDLambda float64 `json:"dphi_dlambda"`
    EventKind   string  `json:"event_kind"`
    AssociatedWithLowR bool `json:"associated_with_low_r"`
    AssociatedWithHighQ bool `json:"associated_with_high_q"`
    AssociatedWithHighPhaseGradient bool `json:"associated_with_high_phase_gradient"`
    Reason      string  `json:"reason"`
}
```

Allowed `event_kind` values:

```text
near_zero_dphi
sign_change
direction_reversal
branch_boundary
```

Use the same parameterized near-zero threshold from Sprint 4, not a hidden constant.

### 4. Local α(φ) branch extraction

Plan local relational branch diagnostics:

```go
type LocalRelationBranch struct {
    SegmentID       string `json:"segment_id"`
    Relation        string `json:"relation"`
    ClockVariable   string `json:"clock_variable"`
    DependentVariable string `json:"dependent_variable"`

    NumSamples      int     `json:"num_samples"`
    ClockRange      float64 `json:"clock_range"`
    DependentRange  float64 `json:"dependent_range"`

    IsSingleValued  bool   `json:"is_single_valued"`
    BranchStatus    string `json:"branch_status"`
    Reason          string `json:"reason"`
}
```

For Sprint 5, this should only verify whether a local relation like `α(φ)` is single-valued on monotonic branches.

Do **not** fit Friedmann equations.

Do **not** compute Friedmann residual.

### 5. Clock-independent diagnostics

Plan diagnostics that do not require φ as a global clock:

```go
type ClockIndependentDiagnostic struct {
    Name   string `json:"name"`
    Status string `json:"status"`
    Value  *float64 `json:"value,omitempty"`
    ValueStatus string `json:"value_status,omitempty"`
    ValueReason string `json:"value_reason,omitempty"`
    Reason string `json:"reason"`
}
```

Examples:

```text
trajectory_finiteness
node_contact_free
max_abs_q_away_from_nodes
min_amplitude_r
phase_gradient_finite
path_length_in_configuration_space
total_lambda_interval
```

These diagnostics can help decide whether a trajectory is usable even when φ is not a global clock.

### 6. Step-refinement branch stability

Use the same four fragile configurations from Sprint 4:

```text
c2 = 0.50, k2 = 2.1, omega2 = -2.1
c2 = 0.55, k2 = 1.9, omega2 = -1.9
c2 = 0.55, k2 = 2.0, omega2 = -2.0
c2 = 0.55, k2 = 2.1, omega2 = -2.1
```

Run:

```text
T = 10.0
dt = 0.05, 0.025, 0.0125
steps = 200, 400, 800
```

Report whether the number and location of local monotonic branches are stable under refinement.

Do not require perfect equality. Require honest reporting.

### 7. Readiness classification

Define:

```text
friedmann_readiness: blocked|local_only_candidate|contested
```

Allowed meanings:

```text
blocked:
  no safe global or local relational clock handling is available.

local_only_candidate:
  φ fails globally but local monotonic branches exist and are segmentable.
  Friedmann residual is still not implemented, but local branch readiness exists.

contested:
  branch behavior is unstable or unclear under refinement.
```

Do not use:

```text
ready_for_friedmann_recovery
```

That is too strong for Sprint 5.

## Proposed CLI

Plan:

```bash
ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
```

Rationale:

```text
segment-clock is clearer than diagnose-clock because Sprint 5 specifically segments local clock branches.
```

`validate` and `summarize` should route by schema version:

```text
bmc0a-clock-readiness-v0.1
```

## Report shape

Plan a deterministic JSON report:

```json
{
  "schema_version": "bmc0a-clock-readiness-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "source_artifacts": [
    "bmc0a_superposition_safe",
    "bmc0a_superposition_robustness",
    "bmc0a_clock_fragility"
  ],
  "readiness_kind": "relational_clock_readiness",
  "near_zero_dphi_threshold": 1e-10,
  "technical_gate": {
    "name": "bmc0a_clock_readiness_gate",
    "status": "pass|contested|fail",
    "reason": ""
  },
  "readiness_outcome": "local_only_candidate|blocked|contested",
  "friedmann_readiness": "blocked|local_only_candidate|contested",
  "global_clock_summary": {
    "phi_global_clock": "fail",
    "alpha_global_clock": "pass|fail|contested",
    "clock_choice_debt": "active"
  },
  "step_refinement_branch_audit": [],
  "clock_segments": [],
  "turning_points": [],
  "local_relation_branches": [],
  "clock_independent_diagnostics": [],
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked",
    "reason": "Friedmann residual, null-model comparison, and full faithfulness review remain unpaid."
  },
  "ebp_debt": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active"
  },
  "warnings": []
}
```

## Validation requirements

Plan strict validation:

```text
reject final_truth_claim = true
reject toy_analysis_only = false
reject missing/wrong schema_version
reject wrong readiness_kind
reject missing near_zero_dphi_threshold
reject near_zero_dphi_threshold <= 0
reject promotion_gate.status != blocked
reject missing technical_gate status/reason
reject invalid readiness_outcome values
reject invalid friedmann_readiness values
reject clock_choice_debt != active
reject nonfinite numeric metrics
reject unavailable numeric values unless paired with explicit status/reason
reject empty step_refinement_branch_audit
reject empty clock_segments if readiness_outcome = local_only_candidate
reject `friedmann_readiness = local_only_candidate` if no valid local relation branches exist
reject unknown JSON fields using DisallowUnknownFields
```

## Lean planning

Do not formalize relational-clock physics yet.

Add only policy/safety contracts:

```lean
BMC/BMC/ClockReadiness.lean
```

Possible definitions:

```lean
structure BMCClockReadinessReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  technicalGatePassed : Bool
  friedmannDeferred : Bool
  fullBMCBlocked : Bool
  clockChoiceDebtActive : Bool
  readinessIsLocalOnlyCandidate : Bool
```

Possible theorems:

```lean
clock_readiness_requires_toy_only
clock_readiness_blocks_final_truth
clock_readiness_requires_friedmann_deferred
clock_readiness_does_not_imply_full_bmc
clock_readiness_keeps_clock_choice_debt_active
clock_readiness_local_candidate_does_not_mean_friedmann_recovered
```

Lean remains policy/safety only.

## Test plan

Plan tests for:

```text
TestClockReadinessReportValidation
TestClockSegmentationDetectsLocalBranches
TestTurningPointsDetectedAtBranchBoundaries
TestLocalRelationBranchesAreSingleValuedOnSegments
TestClockIndependentDiagnosticsDoNotRequirePhiClock
TestStepRefinementBranchAuditDeterministic
TestClockReadinessRejectsFinalTruthClaim
TestClockReadinessRejectsNonfiniteMetrics
TestClockReadinessRequiresClockChoiceDebtActive
TestClockReadinessDeterministicJSON
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
    "containsFinalTruthClaim": "",
    "LeanVerification": ""
  },
  "promotion_status": {
    "sprint5_clock_readiness": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP guardrails

The maximum allowed Sprint 5 promotion is:

```text
promoted_clock_readiness_artifact_after_repairs
```

Do not recommend promotion as:

```text
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann recovery
ready for Friedmann recovery
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Remember:

```text
A local relational branch is not a global clock.
A local-only candidate is not Friedmann recovery.
Clock-independent diagnostics are readiness scaffolding, not physics proof.
Clock-choice debt remains active.
```
