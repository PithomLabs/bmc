# BMC Sprint 3 Planning: Numerical Robustness and Convergence Audit

You are planning **BMC Sprint 3** under **EBP 2.1** discipline.

Do **not** implement code yet. Produce a reviewable implementation plan only.

## Current accepted artifacts

Sprint 1 is accepted as:

```text
BMC-0A plane-wave control artifact
```

Sprint 2 is accepted as:

```text
BMC-0A two-plane-wave superposition control artifact
BMC-0A node-obstruction detection artifact
```

These are **control artifacts only**. They do not prove Bohmian mechanics, quantum gravity, the problem of time, spacetime emergence, Friedmann recovery, or full BMC.

## Sprint 3 goal

Plan a narrow numerical robustness/convergence sprint for the already accepted BMC-0A superposition artifact.

The purpose is:

```text
Check whether the accepted safe-superposition and node-probe results are stable under numerical choices such as step size, node threshold, phase-gradient bound, and small parameter perturbations.
```

Sprint 3 should answer:

```text
Are Sprint 2 results robust, or are they artifacts of one hand-picked step size, threshold, or parameter choice?
```

## Forbidden scope

Do **not** plan or implement:

```text
massive scalar field
Friedmann residual recovery
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
```

Friedmann residual remains deferred.

Faithfulness for full BMC remains contested.

Full BMC toy gate remains blocked.

## Required Sprint 3 planning targets

Plan the smallest useful Sprint 3 around these checks:

### 1. Step-size convergence

For the safe superposition profile, run the same profile with several `LambdaStep` values, for example:

```text
0.1
0.05
0.025
0.0125
```

Keep total integration interval approximately fixed.

Report:

```text
trajectory endpoint drift
max |Q| away from nodes
min amplitude R along trajectory
max phase-gradient magnitude
clock monotonicity status
technical gate status
```

Expected result:

```text
Smaller step sizes should not radically change pass/fail status.
```

Do not require perfect convergence yet. This is a toy numerical audit.

### 2. Threshold sensitivity

Run the safe and node-probe profiles with several node thresholds:

```text
1e-4
1e-5
1e-6
```

Report whether:

```text
safe profile remains node-contact-free
node-probe remains detected as node contact
q_finite_away_from_nodes status remains honest
phase_gradient_finite status remains honest
```

Expected result:

```text
Safe profile should remain safe under reasonable thresholds.
Node-probe should remain blocked under all thresholds.
```

### 3. Phase-gradient bound sensitivity

Run the safe profile with several `MaxPhaseGrad` values:

```text
25
50
100
200
```

Report:

```text
max observed phase-gradient
whether the configured bound is actually binding
whether phase_gradient_finite status changes
```

Expected result:

```text
The bound should not be silently irrelevant. If it never binds, report that as descriptive, not as proof.
```

### 4. Small parameter perturbation audit

Perturb the safe profile slightly while preserving each component’s WdW constraint.

Examples:

```text
c2 = 0.45, 0.50, 0.55
k2 = 1.9, 2.0, 2.1
omega2 = -k2
```

Report:

```text
whether node contact appears
whether Q remains finite away from nodes
whether clock monotonicity remains stable
whether safe gate remains pass or becomes contested/fail
```

Expected result:

```text
Small perturbations should not produce unexplained gate flips.
If gate flips occur, they must be documented as obstruction evidence, not hidden.
```

### 5. Node-probe robustness

For the node-probe profile, test small offsets around the exact node:

```text
(α0, φ0) = (0, 0)
(1e-8, 0)
(1e-6, 0)
(1e-4, 0)
```

Report:

```text
initial amplitude R
whether short-circuit triggers
whether integration is allowed
whether node_contact_free fails or passes
whether phase/Q are contested near node
```

Expected result:

```text
Exact node should short-circuit.
Near-node cases should either short-circuit or be flagged as high-risk depending on NodeThresh.
```

## Go design proposal

Plan additions, but do not implement yet.

Suggested package additions:

```text
internal/bmc/audit/
```

Possible files:

```text
convergence.go
threshold_sensitivity.go
phase_bound_sensitivity.go
perturbation.go
robustness_report.go
```

Suggested CLI profile or command:

```bash
ptw-bmc audit --profile bmc0a-superposition-robustness --out out/bmc0a_superposition_robustness.json
```

Keep CLI stdlib-only.

Do not add external dependencies.

## Report shape

Plan a deterministic JSON report with fields like:

```json
{
  "schema_version": "bmc0a-superposition-robustness-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "source_artifacts": [
    "bmc0a_superposition_safe",
    "bmc0a_superposition_node_probe"
  ],
  "audit_kind": "numerical_robustness_convergence",
  "step_size_sweep": [],
  "node_threshold_sweep": [],
  "phase_gradient_bound_sweep": [],
  "parameter_perturbation_sweep": [],
  "node_probe_offset_sweep": [],
  "technical_gate": {
    "name": "bmc0a_superposition_robustness_audit_gate",
    "status": "pass|fail|contested"
  },
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked"
  },
  "ebp_debt": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested"
  },
  "warnings": []
}
```

## Lean planning

Do not formalize numerical convergence mathematics yet.

Plan only promotion-safety contracts.

Possible Lean additions:

```lean
reportPassesBMC0ARobustnessAuditGate
robustness_audit_does_not_imply_full_bmc
robustness_audit_requires_toy_only
robustness_audit_blocks_final_truth
robustness_audit_requires_friedmann_deferred
```

Lean must remain a policy/safety layer for Sprint 3, not a proof of numerical convergence.

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
    "containsFinalTruthClaim": "",
    "LeanVerification": ""
  },
  "promotion_status": {
    "sprint3_robustness_audit": "",
    "full_bmc_toy_gate": "",
    "forbidden_promotions": []
  },
  "next_smallest_useful_move": ""
}
```

## Strict EBP guardrails

The maximum allowed Sprint 3 promotion is:

```text
promoted_robustness_audit_artifact_after_repairs
```

Do not recommend promotion as:

```text
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann recovery
```

Remember:

```text
Ideas enter free. Promotion costs debt.
A robustness audit can retire numerical-dependence debt.
It cannot retire full physics debt.
```
