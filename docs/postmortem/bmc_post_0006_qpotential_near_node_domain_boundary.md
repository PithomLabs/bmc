# BMC-POST-0006: Quantum Potential Near-Node Domain-Boundary Audit

This document records the results and design decisions for the domain-boundary audit of quantum-potential evaluations under BMC-POST-0006.

## 1. Scope and EBP Limitations

Under strict EBP 2.1 discipline, this work is a **numerical and domain-boundary audit only**.

* POST-0006 is a quantum-potential near-node domain-boundary audit.
* It does not implement BMC-0B.
* It does not solve the Wheeler-DeWitt equation.
* It does not validate BMC.
* It does not claim Friedmann recovery.
* It does not prove classical-limit recovery.
* It does not run null models.
* It does not unblock full BMC.

## 2. Rationale: Instability of Quantum Potential Near Wavefunction Nodes

The quantum potential is defined as:
$$Q = -\frac{1}{2R} \left( \frac{\partial^2 R}{\partial \alpha^2} - \frac{\partial^2 R}{\partial \phi^2} \right)$$
where $R = |\Psi|$ is the wavefunction amplitude. 

Near wavefunction nodes where $R \to 0$, the $1/R$ term makes the expression singular. Any floating-point noise in the numerical second-derivative stencil points is amplified by the division of $R$ (and by the division of $h^2$ in finite differences), causing unstable and non-authoritative values. 

If the implementation silently clamps these values to zero or treats them as valid, it creates validation theater. This audit makes all near-node and stencil domain crossings explicit, blocked, and non-authoritative.

## 3. Audit Design & Enforcements

### Amplitude Floor & derivative step
* `NearNodeAmplitudeFloor = 1e-8`
* `QPotentialDerivativeStep = 1e-4`
* `QPotentialMagnitudeWarning = 1e8`

### Stencil and Domain Boundary Checks
Before evaluating $Q$, the audit checks the wavefunction amplitude $R$:
1. **Node Contact**: If $R == 0.0$ or $R < 10^{-12}$, the point status is set to `q_potential_blocked_by_node_contact` with `Authoritative = false`.
2. **Near Node**: If $R < NearNodeAmplitudeFloor$, the point status is set to `q_potential_blocked_by_near_node_amplitude` with `Authoritative = false`.
3. **Stencil Crossing**: If the amplitude at any stencil offset point used for the second-derivative finite difference ($\alpha \pm h$, $\phi \pm h$) falls below `NearNodeAmplitudeFloor`, the point is flagged with `q_potential_blocked_by_domain_boundary` and `Authoritative = false`.
4. **Nonfinite wave / derivative**: If NaN/Inf values are encountered during wavefunction amplitude evaluations or second-derivative computations, the point is flagged as non-authoritative with status `q_potential_blocked_by_nonfinite_wave` or `q_potential_blocked_by_nonfinite_derivative`.

### Status and Authority Model
* `q_potential_authoritative` is set for plane wave controls where the amplitude is safely nonzero.
* `q_potential_audit_only_no_promotion` is set for multi-wave superposition points where Q is finite, highlighting that it is strictly an audit, not physics validation.

## 4. EBP Status & Remaining Debt

The status remains:
* `QPotentialNearNodeDomainBoundary: audit_only`
* `SolverStatus: not_implemented`
* `BMC0BStatus: specified_only`
* `promotion_recommendation: do_not_promote`

Physical promotion to any solver status remains unpaid.
