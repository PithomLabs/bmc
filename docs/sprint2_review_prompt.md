You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 2: BMC-0A Two-Plane-Wave Superposition Control**.

The artifact extends the already accepted BMC-0A plane-wave control by testing a two-plane-wave superposition where:

```text
Ψ(α, φ) = c1 exp(i(k1 α + ω1 φ)) + c2 exp(i(k2 α + ω2 φ))
```

The purpose is narrow:

```text
Stress-test the BMC-0A pipeline when Ψ has nonconstant amplitude, possible nodes, nonzero quantum potential Q, and nonconstant Bohmian velocity.
```

The artifact must **not** claim:

```text
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann residual recovery
black holes
fermions
gauge fields
Lorentz recovery
inhomogeneous perturbations
LQC comparison completed
Page-Wootters comparison completed
```

## Materials to review

Review the actual source and artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
BMC/**
out/bmc0a_superposition_safe.json
out/bmc0a_superposition_node_probe.json
go.mod
all Go tests
Sprint 1 final report
Sprint 2 walkthrough/report
```

If any file or artifact is missing, report it as implementation debt.

## Expected Sprint 2 design

There should be two profiles:

```text
1. bmc0a-superposition-safe
   Purpose: non-node superposition control.
   Expected: no node contact, Q finite away from nodes, phase gradient finite, technical gate may pass.

2. bmc0a-superposition-node-probe
   Purpose: obstruction-detection test.
   Expected: starts at a known node, node detection passes, node_contact_free fails/blocker, integration short-circuits, safe technical gate blocked.
```

The node-probe profile should not be treated as a successful trajectory simulation. It only passes as a node-detection validation artifact.

The full BMC toy gate must remain blocked.

Friedmann residual must remain deferred.

Faithfulness must remain contested unless a separate human faithfulness review has occurred.

## Required physics checks

Attack these points strictly.

### 1. Superposition WdW residual

For each component:

```text
residual_j = k_j² - ω_j²
```

Check that each component satisfies:

```text
ω_j² = k_j²
```

Because the WdW equation is linear, if all component residuals are zero, the analytic superposition residual is zero.

Review whether:

```text
component residuals are reported
finite-difference residual is only a sanity check
finite-difference residual is not falsely treated as stronger than the analytic component check
```

### 2. Safe profile node avoidance

For the safe profile:

```text
c1 = 1 + 0i
k1 = 1, ω1 = 1
c2 = 0.5 + 0i
k2 = 2, ω2 = -2
α0 = 0, φ0 = 0
node_threshold = 1e-5
```

Review whether the trajectory truly avoids near-node regions along all integrated points.

Check whether `node_contact_free = pass` is supported by actual amplitude samples, not just assumed from unequal amplitudes.

### 3. Node-probe short-circuit

For the node-probe profile:

```text
c1 = 1 + 0i
k1 = 1, ω1 = 1
c2 = -1 + 0i
k2 = 2, ω2 = -2
α0 = 0, φ0 = 0
```

At the initial point:

```text
Ψ(0,0) = 1 - 1 = 0
```

Review whether the implementation detects the node before attempting:

```text
velocity calculation
RK4 stepping
phase-gradient calculation
Q calculation at Ψ = 0
```

The correct behavior is:

```text
node_detection = pass
node_contact_free = fail/blocker
trajectory = fail or contested / short-circuited
phase_gradient_finite = contested or fail, but not pass based on invalid computation
q_finite_away_from_nodes = contested, pass-vacuous, or not-applicable with clear reason
safe_superposition_gate = blocked
node_detection_validation_gate = pass
full_bmc_toy_gate = blocked
```

If the code computes through Ψ=0 and merely catches NaN afterward, mark this as a blocker.

### 4. Quantum potential handling

For superpositions, Q may be nonzero and unstable near nodes.

Review whether:

```text
Q is computed only away from nodes
Q finite checks skip or block near-node points
node-probe does not claim Q was physically computed at Ψ=0
q_finite_away_from_nodes is honestly described
```

If the node-probe has no away-from-node points because it short-circuits immediately, check that the report does not misleadingly call Q globally finite.

### 5. Phase gradient handling

The phase-gradient formula should be based on:

```text
∂x S = Im(Ψ⁻¹ ∂x Ψ)
```

Review whether the implementation:

```text
refuses near-node phase-gradient evaluation
checks nonfinite gradients
uses max_phase_gradient
does not rely only on raw arg jumps
```

### 6. RK4 test quality

Review `TestRK4StepperCorrectness`.

Reject it as weak if it only tests a constant-velocity field where Euler is already exact.

A good RK4 test should use a non-constant velocity field with known behavior or compare RK4 vs Euler on a curve where RK4 is demonstrably more accurate.

### 7. Clock monotonicity

For both profiles, check whether φ monotonicity is measured from the generated trajectory.

For node-probe, if trajectory is empty/short-circuited, clock monotonicity should not be falsely passed. It should be fail, contested, or not-applicable with clear reason.

### 8. Report and validation gates

Check that the reports distinguish:

```text
bmc0a_superposition_safe_gate
node_detection_validation_gate
full_bmc_toy_gate
```

The safe profile may pass only the safe gate.

The node-probe profile may pass only node-detection validation, not the safe gate.

The full BMC toy gate must remain blocked in both.

### 9. Lean safety contracts

Review:

```text
BMC/BMC/ToyReport.lean
BMC/BMC/Promotion.lean
```

Check that:

```text
lake build succeeds
no sorry/admit exists
safe superposition gate is defined
node-detection gate is defined
full toy gate remains blocked by Friedmann/faithfulness
witnesses demonstrate safe profile can pass safe gate
witnesses demonstrate node-probe passes node-detection gate but fails safe gate
no Lean theorem claims physics or full QG
```

### 10. Overclaim scan

Search code, report text, README, comments, tests, CLI output, and summaries for forbidden implication:

```text
solves quantum gravity
proves Bohmian mechanics
solves problem of time
derives spacetime
recovers Friedmann
validates full BMC
validates quantum gravity
```

Any such wording is a promotion blocker.

## Code review checks

Review:

```text
parameter validation
status/pass consistency
typed ReportParameters
deterministic JSON
CLI profile routing
validation for safe profile
validation for node-probe profile
NaN/Inf handling
MaxPhaseGrad validation
NodeThresh validation
RK4 behavior
near-node short-circuit behavior
```

## EBP debt classification

Classify each debt item as one of:

```text
unpaid
partial
retired
contested
overclaimed
```

Debt items:

```text
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
containsFinalTruthClaim
LeanVerification
SafeSuperpositionControl
NodeDetectionValidation
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
    "containsFinalTruthClaim": "",
    "LeanVerification": "",
    "SafeSuperpositionControl": "",
    "NodeDetectionValidation": ""
  },
  "physics_findings": [],
  "code_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|superposition_control_candidate_only|promoted_superposition_control_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 2 passes perfectly, the maximum allowed recommendation is:

```text
promoted_superposition_control_artifact_after_repairs
```

Never recommend promotion as:

```text
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann recovery
```
