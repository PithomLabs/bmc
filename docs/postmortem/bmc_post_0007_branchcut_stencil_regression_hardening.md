# BMC-POST-0007: Branch-Cut and Stencil Boundary Regression Hardening

This document records the design decisions and regression test validations implemented under BMC-POST-0007.

## 1. Scope and EBP Limitations

Under strict EBP 2.1 discipline, this work is **regression hardening only**.

* POST-0007 is a regression hardening ticket.
* It does not implement BMC-0B.
* It does not solve the Wheeler-DeWitt equation.
* It does not validate BMC.
* It does not claim Friedmann recovery.
* It does not prove classical-limit recovery.
* It does not run null models.
* It does not unblock full BMC.

## 2. Hardening Details

This ticket adds targeted regression tests to protect the numerical audit stack from regression risks around phase-gradients and quantum-potentials.

### Phase-Gradient Branch-Cut wrapping
The finite-difference phase-gradient implementation utilizes the identity $\text{Im}\left(\frac{1}{\Psi} \frac{\partial \Psi}{\partial x}\right)$ to avoid wrapping artifacts near the principal-argument branch cut. 
We introduced a regression test `TestHSensitivityBranchCutSafeGradientDoesNotArgWrap` with a mock wavefunction:
$$\Psi(\alpha, \phi) = \exp(i (\pi - 0.5 \alpha))$$
Evaluated near $\alpha = 0$, a naïve implementation finite-differencing `arg(Ψ)` directly would cross the principal argument boundary and calculate a huge false derivative (on the order of $-313.66$). The regression test verifies that the calculated gradient remains stable and close to $-0.5$ (with $\frac{\partial S}{\partial \phi} \approx 0$).

### Quantum-Potential Stencil Boundary Crossings
The quantum-potential audit requires that evaluations are flagged non-authoritative if any of the central or finite-difference stencil points used for second derivatives fall below the `NearNodeAmplitudeFloor` ($10^{-8}$). 
We introduced the regression test `TestQPotentialBlocksStencilPointBelowAmplitudeFloor` using a mock wavefunction where the evaluation center point is safe ($R = 1.0$) but exactly one stencil point ($\alpha + h$) is zero. The test verifies that the audit detects this boundary crossing, flags `Authoritative = false`, and returns status `q_potential_blocked_by_domain_boundary`.

### Top-Level Input Validation
The regression test `TestQPotentialRunAuditRejectsInvalidDerivativeStep` ensures that the public audit entry point correctly rejects invalid derivative steps (nonpositive values $h \le 0$, `NaN`, or `Inf`) prior to performing any calculations.

## 3. EBP Status & Remaining Debt

The status remains:
* `PhaseGradientBranchCutRegression: targeted`
* `QPotentialStencilBoundaryRegression: targeted`
* `SolverStatus: not_implemented`
* `BMC0BStatus: specified_only`
* `promotion_recommendation: do_not_promote`

Physical promotion to any solver status remains unpaid.
