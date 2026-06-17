Approve Sprint 2 plan **with corrections**. The direction is right: two-plane-wave superposition is the correct next stress test after the plane-wave control. But the proposed defaults and gate wording need tightening before implementation.

## Default parameters: revise them

Your proposed defaults are:

```text
c1 = 1.0 + 0i, k1 = 1.0, ω1 = 1.0
c2 = 0.5 + 0i, k2 = 2.0, ω2 = -2.0
α0 = 0.0, φ0 = 0.0
Δλ = 0.05, steps = 200
node_threshold = 1e-5
```

These are **acceptable as a non-node default**, but they may not reliably stress node behavior because `|c1| ≠ |c2|`. With amplitudes `1.0` and `0.5`, exact cancellation is impossible, so you get interference but not true nodes.

I recommend two profiles instead:

```text
bmc0a-superposition-safe
c1 = 1.0 + 0i
k1 = 1.0, ω1 = 1.0
c2 = 0.5 + 0i
k2 = 2.0, ω2 = -2.0
α0 = 0.0, φ0 = 0.0
Δλ = 0.05
steps = 200
node_threshold = 1e-5
expected: no node contact, Q computable, technical gate may pass
```

and:

```text
bmc0a-superposition-node-probe
c1 = 1.0 + 0i
k1 = 1.0, ω1 = 1.0
c2 = -1.0 + 0i
k2 = 1.0, ω2 = -1.0
or another equal-amplitude setup chosen to create near-cancellation
expected: node or near-node obstruction should trigger
```

This is better EBP design: one profile tests **successful nontrivial superposition**, the other tests **obstruction detection**.

## Required correction 1: do not say the gate requires “no active obstructions”

This sentence is too broad:

```text
passing this superposition control gate requires that the simulation runs without triggering active obstructions
```

For Sprint 2, this should be narrower:

```text
The safe superposition profile may pass only if no blocker obstruction is triggered.
The node-probe profile is expected to trigger node obstruction and should not pass the technical gate.
A successful obstruction-triggering test still retires obstruction-detection debt.
```

In EBP terms, a failing/node-triggering profile can still be useful. It should not pass the technical gate, but it can prove the workbench detects the failure honestly.

## Required correction 2: add `qFiniteAwayFromNodes`, not just `qFinite`

For superpositions, `Q` can become unstable near nodes. The gate should not require global `qFinite = pass` unless the trajectory never enters near-node regions.

Use:

```text
qFiniteAwayFromNodes = pass
nodeObstruction = pass
```

where `nodeObstruction = pass` means:

```text
No node contact for safe profile
OR
node contact correctly detected for node-probe profile
```

Better yet, separate these:

```text
nodeDetection = pass
nodeContact = pass | fail | deferred
qFiniteAwayFromNodes = pass
```

For the **safe profile**, `nodeContact` should be `pass` meaning “no contact detected.”

For the **node-probe profile**, `nodeContact` should be `fail/blocker` but `nodeDetection` should be `pass`.

## Required correction 3: avoid `interface{}` for `Parameters` if possible

Changing `Parameters` to `interface{}` is convenient but weakens validation. Prefer a typed wrapper:

```go
type ReportParameters struct {
    PlaneWave     *model.PlaneWaveParams     `json:"plane_wave,omitempty"`
    Superposition *model.SuperpositionParams `json:"superposition,omitempty"`
}
```

This keeps JSON deterministic and validation inspectable.

Use `interface{}` only if the coding agent strongly prefers it and validation remains strict.

## Required correction 4: WdW residual for superposition should have analytic and finite-diff paths

Because the equation is linear, if each component satisfies the WdW equation, the superposition does too.

Add:

```go
AnalyticResidualSuperposition(s SuperpositionWave) float64
```

or component residual reporting:

```json
"component_residuals": [
  { "component": 1, "k": 1, "omega": 1, "residual": 0 },
  { "component": 2, "k": 2, "omega": -2, "residual": 0 }
]
```

Finite-difference residual is useful, but do not make it the sole authority.

## Required correction 5: phase unwrap obstruction should use `Ψ⁻¹ ∂Ψ`, not raw phase jumps only

Raw `arg(Ψ)` jumps can be misleading because phase is periodic. Since Sprint 1 already uses:

```text
∂x S = Im(Ψ⁻¹ ∂x Ψ)
```

Sprint 2 should treat phase obstruction as:

```text
If |Ψ| < node_threshold, phase-gradient is contested.
If phase-gradient magnitude becomes nonfinite or exceeds configured bound, phase-gradient obstruction applies.
```

The `π/2` phase-jump heuristic can be reported as a warning, but the blocker should be near-node/nonfinite-gradient behavior.

## Required Lean correction

Do not make the superposition gate require simply:

```lean
nodeObstruction == pass
```

That is ambiguous. Add separate fields:

```lean
nodeDetection : CheckStatus
nodeContactFree : CheckStatus
qFiniteAwayFromNodes : CheckStatus
phaseGradientFinite : CheckStatus
```

Then:

```lean
def reportPassesBMC0ASuperpositionSafeGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.nodeDetection &&
  checkPassed r.nodeContactFree &&
  checkPassed r.qFiniteAwayFromNodes &&
  checkPassed r.phaseGradientFinite &&
  checkPassed r.classicalLimit &&
  checkDeferred r.friedmannResidual
```

For the node-probe profile, define a different gate or witness:

```lean
reportPassesBMC0ANodeDetectionGate
```

This should pass if node contact is detected correctly, while the trajectory/promotion gate remains blocked.

## EBP debt ledger for Sprint 2 plan

```text
needMap: partial, unchanged
needInvariant: partial, now testing Q computability away from nodes
needToyCheck: active, extending beyond plane wave
needNullModel: still partial/deferred
needObstruction: active, node and phase-gradient obstruction become real checks
needFaithfulnessReview: active/contested
containsFinalTruthClaim: must remain false
promotion: control-level only, full BMC blocked
```

## Final approval status

```text
Sprint 2 plan: approved with required revisions
Default safe profile: acceptable
Need additional node-probe profile: yes
Do not start massive scalar/Friedmann/LQC/Page-Wootters yet
Adversarial review: after implementation, not before
```

The corrected next instruction to the coding agent is:

```text
Proceed with Sprint 2 planning using two profiles:
1. bmc0a-superposition-safe — non-node superposition, technical gate may pass.
2. bmc0a-superposition-node-probe — intentionally tests node obstruction; should not pass the safe technical gate, but should pass node-detection validation.

Preserve split gates. Full BMC toy gate remains blocked.
```
