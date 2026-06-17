The next step is **Sprint 2: BMC-0A Packet/Superposition Control**.

Do **not** jump to the massive scalar numerical Wheeler-DeWitt model yet. Sprint 1 proved that the pipeline works in the easiest exact case: plane wave, constant amplitude, `Q = 0`, no nodes, no real phase trouble. The next EBP step is to attack the first nontrivial obstruction:

```text
What happens when Ψ is a superposition, R is nonconstant, Q may be nonzero, and nodes/phase issues can appear?
```

## Sprint 2 target

```text
BMC-0A.2 packet/superposition control
```

Implement a simple superposition of two WdW-valid plane waves:

```text
Ψ(α, φ) =
c1 exp(i(k1 α + ω1 φ)) +
c2 exp(i(k2 α + ω2 φ))

with:
ω1² = k1²
ω2² = k2²
```

This keeps the same flat massless minisuperspace equation:

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ = 0
```

but now the amplitude is generally not constant, so:

```text
R = |Ψ| may vary
Q may be nonzero
nodes may appear
phase gradients may become unstable near nodes
```

That is exactly the next stress test.

## Why this is the right next step

Sprint 1 retired only the **plane-wave control subset** of `needToyCheck`.

Sprint 2 begins paying the next debt:

| EBP debt                 | Sprint 2 purpose                                                           |
| ------------------------ | -------------------------------------------------------------------------- |
| `needToyCheck`           | Test nontrivial wavefunction behavior beyond the exact plane wave.         |
| `needObstruction`        | Activate node, phase-unwrapping, nonfinite-Q, and clock-failure detectors. |
| `needInvariant`          | Check whether `Q` is computable and finite away from nodes.                |
| `needFaithfulnessReview` | Confirm the Go toy still tests only BMC-0A, not full BMC.                  |
| `needNullModel`          | Still mostly deferred; no need to compare LQC/Page-Wootters yet.           |

## Required Sprint 2 additions

### Go additions

Add a new wave type:

```text
internal/bmc/wave/superposition.go
```

with:

```go
type PlaneComponent struct {
    Coeff complex128
    K     float64
    Omega float64
}

type SuperpositionWave struct {
    Components []PlaneComponent
}
```

Validation must require every component to satisfy:

```text
omega² = k²
```

Add node detection:

```text
node_obstruction
```

using:

```text
R = |Ψ| < node_threshold
```

Add phase-gradient risk detection:

```text
phase_unwrap_obstruction
```

For Sprint 2, this can be simple:

```text
if trajectory enters near-node region, phase gradient is contested
```

Add variable guidance:

```text
dα/dλ = ∂S/∂α
dφ/dλ = -∂S/∂φ
```

computed using the branch-safe identity already used in Sprint 1:

```text
∂x S = Im(Ψ⁻¹ ∂x Ψ)
```

But now it must fail safely if `|Ψ|` is too small.

### Integrator

Add RK4 **only now**, because the velocity field is no longer constant.

Keep Euler as the Sprint 1 control integrator.

Add:

```text
RK4Stepper
```

but keep the report explicit:

```text
integrator: rk4
integrator_debt: numerical accuracy not formally proved
```

### Report changes

Add a new profile:

```bash
ptw-bmc run --profile bmc0a-superposition --out out/bmc0a_superposition.json
```

New report fields should include:

```json
{
  "model_id": "bmc0a_superposition",
  "wave_family": "two_plane_wave_superposition",
  "node_threshold": 1e-6,
  "checks": {
    "wdw_residual": {},
    "trajectory": {},
    "clock_monotonicity": {},
    "quantum_potential": {},
    "node_detection": {},
    "phase_gradient": {},
    "friedmann_residual": {
      "status": "deferred"
    }
  },
  "technical_gate": {
    "name": "bmc0a_superposition_control_gate"
  },
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked"
  }
}
```

## Success criteria

Sprint 2 succeeds only if:

```text
1. Each component satisfies the WdW constraint.
2. The superposed Ψ satisfies the WdW residual check numerically.
3. Q is computable away from nodes.
4. Node regions are detected and not hidden.
5. Trajectories either integrate safely or fail with explicit obstruction reports.
6. φ monotonicity is measured, not assumed.
7. Friedmann residual remains deferred.
8. Full BMC promotion remains blocked.
```

## Failure criteria

Sprint 2 fails, or remains blocked, if:

```text
1. The code hides node singularities.
2. Q becomes NaN/Inf without obstruction reporting.
3. Phase gradients are computed through Ψ≈0 as if nothing happened.
4. The report implies packet success proves full BMC.
5. The full promotion gate accidentally passes.
```

## Lean changes

Do not formalize the physics of superpositions yet.

Only add promotion-safety contracts:

```lean
reportPassesBMC0ASuperpositionControlGate
```

This gate should require:

```text
toyAnalysisOnly = true
finalTruthClaim = false
wdwResidual = pass
trajectoryFinite = pass OR trajectoryObstructionFiled = pass
qFiniteAwayFromNodes = pass
nodeDetection = pass
friedmannResidual = deferred
```

The exact Lean structure can remain policy-level. The key theorem should be:

```text
superposition control gate does not imply full BMC toy gate
```

using a concrete witness, like Sprint 1.

## Next coding-agent prompt should say

```text
Implement Sprint 2 only: BMC-0A two-plane-wave superposition control.

Do not implement massive scalar, Friedmann residual, LQC comparison, Page-Wootters comparison, black holes, fermions, gauge fields, or full quantum gravity.

Goal:
Stress-test the existing BMC-0A pipeline when Ψ has nonconstant amplitude, possible nodes, nonzero Q, and nonconstant Bohmian velocity.

Required:
- SuperpositionWave
- component constraint validation
- numerical WdW residual
- RK4Stepper
- node detection
- phase-gradient obstruction near nodes
- Q computation away from nodes
- deterministic JSON report
- split gates preserved
- full BMC toy gate remains blocked
- Lean promotion-safety gate for superposition control
```

## EBP status after Sprint 1

```text
BMC-0A plane-wave control: promoted control artifact after repairs
BMC-0A superposition control: next active work item
Full BMC-0.1: unpromoted
Friedmann residual: unpaid
Null-model comparison: unpaid
Faithfulness review for full BMC: contested
```

So the immediate next move is:

```text
Start Sprint 2 planning for a two-plane-wave superposition control.
```

This is the right “next smallest useful move” because it attacks the first real Bohmian obstruction: **nodes, phase gradients, and nonzero quantum potential**.
