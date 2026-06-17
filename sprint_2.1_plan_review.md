Approve the revised Sprint 2 plan **with one important correction**.

The structure is now much better:

```text id="trw0gc"
safe profile = may pass technical gate
node-probe profile = should trigger node obstruction, not safe gate
full BMC toy gate = remains blocked
Friedmann residual = remains deferred
```

## Required correction: node-probe defaults

Your node-probe profile may **not reliably hit a node** with:

```text id="0iuvxb"
α0 = 0.0, φ0 = 0.1
```

For:

```text id="5hg0n0"
Ψ = exp(i(α + φ)) - exp(i(2α - 2φ))
```

nodes occur when:

```text id="29gvgh"
α + φ = 2α - 2φ mod 2π
```

or:

```text id="8hm4f8"
α = 3φ - 2πn
```

So a guaranteed node is:

```text id="5l32du"
α0 = 0.0, φ0 = 0.0
```

because both phases are zero and `Ψ = 1 - 1 = 0`.

Use this for the node-probe profile:

```text id="d0jbr3"
bmc0a-superposition-node-probe
c1 = 1.0 + 0i
k1 = 1.0, ω1 = 1.0
c2 = -1.0 + 0i
k2 = 2.0, ω2 = -2.0
α0 = 0.0, φ0 = 0.0
Δλ = 0.05
steps = 200
node_threshold = 1e-5
expected: node detected immediately; safe technical gate blocked
```

Better still: for the **unit test**, do not depend on integration to discover a node. Construct a trajectory containing the known node point directly:

```text id="k4zy0k"
Trajectory point: α = 0, φ = 0
Expected: node_detection = pass, node_contact_free = fail/blocker
```

Then the profile can also run, but the detector test remains deterministic.

## Second correction: analytic residual formula

This proposed function is slightly too lossy:

```go id="gmqzg0"
func AnalyticResidualSuperposition(k1, omega1, k2, omega2 float64) float64
```

returning:

```text id="ui97kh"
k1² - omega1² + k2² - omega2²
```

That works only as a scalar constraint summary. The actual residual is component-wise:

```text id="vbiy4i"
c1(k1² - ω1²)e^{iθ1} + c2(k2² - ω2²)e^{iθ2}
```

For Sprint 2, component residuals are enough, but name it honestly:

```go id="m98019"
ComponentResidualsSuperposition(...) []ComponentResidual
```

Then the report can say:

```text id="61iajz"
All component residuals are zero; therefore by linearity the analytic superposition residual is zero.
```

Finite-difference residual remains a sanity check.

## Third correction: `MaxPhaseGrad` validation

Add validation:

```text id="6r1bzr"
MaxPhaseGrad > 0
MaxPhaseGrad finite
```

The plan currently validates `NodeThresh`, but not `MaxPhaseGrad`.

## Fourth correction: node-probe validation gate

This is good:

```text id="au2s6z"
node_detection = pass
node_contact_free = fail/blocker
```

But phrase the gate as:

```text id="kd8nx9"
node_detection_validation_gate
```

not a “technical pass” for the physics simulation. The node-probe profile passes only as an **obstruction-detector test**, not as a successful trajectory simulation.

## Approved coding-agent instruction

Use this concise clarification:

```text id="ah3xw5"
Sprint 2 plan approved with corrections:

1. Keep two profiles:
   - bmc0a-superposition-safe
   - bmc0a-superposition-node-probe

2. Change node-probe initial state to α0=0, φ0=0 so Ψ=0 is guaranteed for:
   Ψ = exp(i(α+φ)) - exp(i(2α-2φ)).

3. Add a deterministic unit test that constructs a trajectory containing the known node point α=0, φ=0. Do not rely only on integration to discover the node.

4. Treat the node-probe profile as passing only the node-detection validation gate, not the safe superposition technical gate.

5. Replace AnalyticResidualSuperposition scalar summary with component residual reporting:
   residual_j = k_j² - ω_j².
   Report that linearity implies the full superposition residual is zero when all component residuals are zero.

6. Validate MaxPhaseGrad > 0 and finite.

7. Preserve split gates:
   - safe profile may pass bmc0a_superposition_safe_gate
   - node-probe must block safe gate but pass node-detection validation
   - full_bmc_toy_gate remains blocked
   - Friedmann residual remains deferred
   - faithfulness remains contested

8. Do not implement massive scalar, Friedmann residual, LQC, Page-Wootters, black holes, fermions, gauge fields, perturbations, or full quantum gravity.
```

## Final EBP status

```text id="iqkbwg"
Sprint 2 plan: approved with required corrections
needToyCheck: active
needObstruction: now actively testing node/phase/Q risks
needInvariant: partial, Q finite away from nodes
needNullModel: still deferred
needFaithfulnessReview: contested
full BMC promotion: blocked
```

Once those corrections are applied, the coding agent can proceed with Sprint 2 implementation.
