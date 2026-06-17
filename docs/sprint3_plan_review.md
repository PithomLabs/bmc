Approve Sprint 3 plan **with required revisions before implementation**.

The direction is right: Sprint 3 should not add new physics. It should audit whether Sprint 2’s safe and node-probe results survive ordinary numerical choices.

## Main required corrections

### 1. Do not hard-code robustness gate to `pass`

This part is too optimistic:

```json
"technical_gate": { "name": "bmc0a_superposition_robustness_audit_gate", "status": "pass" }
```

Change it to computed status:

```json
"technical_gate": {
  "name": "bmc0a_superposition_robustness_audit_gate",
  "status": "pass|contested|fail",
  "reason": ""
}
```

The audit gate should pass only if all planned sweeps run, reports validate, no NaN/Inf is hidden, and gate flips are honestly recorded. A gate flip does not automatically fail the audit; hiding it does.

### 2. Keep total integration time fixed in step-size sweep

For step-size convergence, do not keep `N = 200` for every step size, because then each run ends at a different total λ.

Use a fixed total interval, preferably the Sprint 2 baseline:

```text
T = 10.0
LambdaStep values: 0.1, 0.05, 0.025, 0.0125
Steps: 100, 200, 400, 800
```

Then endpoint drift is meaningful.

### 3. Tests should verify honest reporting, not force perfect robustness

This line is a little too strong:

```text
TestStepSizeConvergence: Asserts endpoint drift stays small and smaller step sizes don't flip gate status.
```

Better:

```text
TestStepSizeConvergence:
- verifies all step-size runs complete deterministically,
- verifies endpoint drift is finite and reported,
- verifies gate status changes, if any, are explicitly recorded,
- optionally applies a conservative drift tolerance only after baseline values are known.
```

EBP rule: a robustness audit can reveal fragility. Fragility is not failure if reported honestly.

### 4. Add validation for robustness reports

Since this introduces a new schema, add:

```text
ValidateRobustnessReport()
SummarizeRobustnessReport()
```

or extend existing validation cleanly without weakening the original BMC report validator.

Do not let ordinary `validate` accept unknown schemas loosely.

### 5. Node-probe offset sweep needs exact expectations

For node-probe offsets, use threshold-aware expectations:

```text
(0, 0): must short-circuit.
(1e-8, 0): likely below 1e-5 threshold, should short-circuit.
(1e-6, 0): likely below 1e-5 threshold, should short-circuit.
(1e-4, 0): may be above 1e-5 threshold and may integrate, but must be flagged if phase/Q risk appears.
```

Do not require all offsets to behave identically. Require them to be classified honestly.

## Approved revised Sprint 3 scope

```text
Sprint 3: numerical robustness/convergence audit only
No massive scalar
No Friedmann residual
No LQC/Page-Wootters comparison
No perturbations beyond tiny parameter sweeps of the existing toy
No full BMC promotion
```

## EBP ledger

```text
needMap: partial, unchanged
needInvariant: partial, improved only for numerical stability of Q/phase metrics
needToyCheck: active, robustness subset now being audited
needNullModel: partial/deferred
needObstruction: active, threshold/near-node behavior audited
needFaithfulnessReview: contested
containsFinalTruthClaim: must remain absent
LeanVerification: planned policy layer only
promotion ceiling: promoted_robustness_audit_artifact_after_repairs
```

## Final instruction to coding agent

```text
Proceed with Sprint 3 planning after applying these revisions:

1. Robustness gate status must be computed, not hard-coded pass.
2. Step-size sweep must keep total integration interval fixed.
3. Tests should verify deterministic execution, finite reported metrics, honest gate-flip reporting, and only conservative tolerances.
4. Add strict validation/summarization support for the robustness report schema.
5. Node-probe offset sweep must classify each offset relative to NodeThresh rather than assuming identical outcomes.

Do not implement new physics. Full BMC toy gate remains blocked.
```

After those edits, the plan is implementation-ready.
