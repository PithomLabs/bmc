# BMC-POST-0005: Phase Gradient h-Sensitivity Audit

This document records the results and design decisions for the numerical h-sensitivity audit of phase-gradient calculations under BMC-POST-0005.

## 1. Scope and EBP Limitations

Under strict EBP 2.1 discipline, this work is a **numerical sensitivity audit only**.

* POST-0005 is a numerical h-sensitivity audit.
* It does not implement BMC-0B.
* It does not solve the Wheeler-DeWitt equation.
* It does not validate BMC.
* It does not claim Friedmann recovery.
* It does not prove classical-limit recovery.
* It does not run null models.
* It does not unblock full BMC.

## 2. Rationale: Why Phase Gradients Need h-Sensitivity Checks

Phase-gradient calculations in BMC-0A are evaluated numerically using central finite differences on configuration space grids with a default step size `h`. Near wavefunction nodes or in region of highly-oscillatory phase behavior, the numerical derivatives can be highly sensitive to the step size `h` or suffer from branch-cut artifacts. 

Directly finite-differencing the wrapped phase `S = arg(Ψ)` can introduce spurious jumps of size `2π`. To prevent this, we formulate derivatives using the analytic identity:
$$\frac{\partial S}{\partial x} = \text{Im}\left(\frac{1}{\Psi} \frac{\partial \Psi}{\partial x}\right)$$
This derivative form is treated strictly as internal numerical audit machinery, not new physics authority. 

Evaluating the stability of the phase gradient across an `h` ladder allows the system to flag unstable or non-authoritative gradients instead of silently relying on unverified grid resolutions.

## 3. Audit Design

### h-Ladder
The default audit uses a step size refinement ladder:
* $h = 1.0 \times 10^{-2}$
* $h = 5.0 \times 10^{-3}$
* $h = 2.5 \times 10^{-3}$
* $h = 1.25 \times 10^{-3}$

### Sample Fixtures
1. **`plane_wave`**: Evaluated on analytic plane wave control wavefunctions where exact gradients are known:
   $$\frac{\partial S}{\partial \alpha} = k, \quad \frac{\partial S}{\partial \phi} = \omega$$
   These are compared without Bohmian guidance sign conventions (e.g. $\dot{\phi} = -\frac{\partial S}{\partial \phi}$ belongs to trajectory dynamics and is separate).
2. **`superposition`**: Evaluated on multi-wave superpositions to track numerical drift across `h`.
3. **`near_node`**: Points placed near wavefunction zeroes to verify that near-node behavior is correctly blocked.

### Authoritativeness and Status Criteria
* **`stable_for_control_scope`**: Achieved only when all evaluated points are stable (low drift) and authoritative.
* **`sensitive_to_h`**: Occurs if any point gradient drifts beyond `AbsoluteDriftTolerance` ($10^{-6}$) or `RelativeDriftTolerance` ($10^{-4}$ when denominator $> 10^{-5}$).
* **`blocked_by_node_contact`**: Occurs if any point coordinate or finite-difference stencil point has an amplitude below the `NearNodeAmplitudeFloor` ($10^{-8}$). Such points are flagged non-authoritative.
* **`blocked_by_nonfinite_gradient`**: Occurs if `NaN` or `Inf` values are encountered.

## 4. EBP Status & Remaining Debt

The status remains:
* `PhaseGradientHSensitivity: audit_only`
* `SolverStatus: not_implemented`
* `BMC0BStatus: specified_only`
* `promotion_recommendation: do_not_promote`

Physical promotion to any solver status remains unpaid.
