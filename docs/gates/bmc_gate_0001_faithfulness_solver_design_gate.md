# BMC-GATE-0001: Faithfulness Review and Solver-Design Gate

This document defines the strict checkpoints, review templates, and validation criteria that must be satisfied before any future BMC-0B solver, numerical cosmology run, or recovery claim can be proposed.

## 1. Gate Purpose & Non-Goals

The purpose of this gate is to establish a rigorous framework for reviewing mathematical formulation, physical correctness, and numerical reliability **prior** to any future solver implementation. 

### Non-Goals
* This gate does not implement a solver.
* This gate does not validate BMC.
* This gate does not solve the Wheeler-DeWitt (WdW) equation.
* This gate does not recover Friedmann dynamics.
* This gate does not prove classical-limit recovery.
* This gate does not run or defeat null models.
* This gate does not unblock full BMC.

## 2. Current Accepted Remediation Stack Scope

Any future proposal must build upon and respect the accepted limited-scopes of the closed remediation stack:
* **POST-0001**: Accepted only for plane-wave/toy WdW constraint violation detection scope.
* **POST-0002 / 0002.1**: Accepted only for plane-wave numerical residual report-path authority scope.
* **POST-0003 / 0003.2**: Accepted only for superposition trajectory numerical self-consistency/convergence audit scope.
* **POST-0004 / 0004.2**: Accepted only for BMC-0B massive scalar specification-only scope.
* **POST-0005**: Accepted only for phase-gradient h-sensitivity audit scope.
* **POST-0006**: Accepted only for quantum-potential near-node domain-boundary audit scope.
* **POST-0007**: Accepted only for branch-cut and stencil-boundary regression-hardening scope.
* **POST-0008**: Accepted only for stack closure ledger and freeze scope.

## 3. Required Mathematical Problem Statement

Future solver proposals must explicitly define the complete mathematical system:
* Configuration variables ($\alpha, \phi$) and their physical meaning.
* Minisuperspace metric and signature.
* Hamiltonian constraint.
* Wheeler-DeWitt operator form.
* Potential term for the massive scalar model.
* Domain of $\alpha$ and $\phi$.
* Wavefunction boundary conditions and admissibility criteria.
* Definition of numerical residuals, residual norm, and tolerance policy.
* Authoritative pass/fail criteria.

## 4. Operator-Form Review

Future implementation is blocked until a human-reviewed operator form is specified, addressing:
* What is the exact WdW equation being solved?
* Which signs represent convention choices vs. physical/mathematical requirements?
* Where does the scalar potential enter the equation?
* What is the kinetic operator form and its mathematical properties?
* What is imported from literature vs. assumed by PTW?
* What remains unproved?

## 5. Factor-Ordering Review

The choice of factor ordering introduces physical debt that must be made explicit:
* What factor ordering is assumed?
* Which alternative factor orderings were considered, and why were they rejected?
* Does the ordering affect computed residuals, boundary conditions, or Bohmian guidance signs?
* What is the EBP debt status of this choice?

## 6. Units and Convention Review

A fixed convention table must be declared to prevent silent convention switching:
* Coordinate mapping: $\alpha = \ln(a)$ and scalar field $\phi$.
* Trajectory parameter convention (e.g. $\lambda$ or time coordinate).
* Units selection (e.g. Planck units).
* Mass parameter convention.
* Potential normalization.
* Metric signature convention.
* WdW sign convention.
* Phase-gradient convention.
* Guidance-sign convention (e.g. separating phase gradient $\partial S / \partial \phi$ from Bohmian velocity $\dot{\phi}$).
* Quantum-potential convention.

## 7. Minisuperspace Metric/Signature Review

The metric signature on the configuration space grid determines the hyperbolic vs. elliptic nature of the partial differential equation. Future solver design must:
* Define the signature conventions.
* Assess the physical validity of the chosen signature.
* Address grid alignment and coordinate singularities.

## 8. Boundary-Condition Review

Future work must explicitly specify:
* Domain bounds and grid boundaries.
* Admissible boundary conditions.
* Boundary residual treatment.
* Node, near-node, and stencil-boundary handling.
* Rejection of vague boundary labels (e.g. "default", "reasonable", "standard", "ready", "validated").

## 9. Wavefunction/Domain Admissibility Review

Define the criteria for wavefunction admissibility:
* Square-integrability or normalizability of $\Psi(\alpha, \phi)$ on the chosen domain.
* Behavior at infinite limits ($\alpha \to \pm\infty, \phi \to \pm\infty$).
* Regularity at configuration space boundaries.

## 10. Residual Norm and Tolerance Review

Numerical residuals must be measured under authoritative norms:
* Which norms ($L_2$, $L_\infty$, boundary norms) are computed?
* What are the justification rules for tolerance thresholds?
* Are failing stencils blocked or silently clamped?

## 11. Solver-Design Review

A solver-design document must be submitted prior to code changes:
* Solver type (e.g. finite difference, spectral).
* Grid strategy and stability constraints (e.g. CFL conditions).
* Convergence criteria under grid refinement.
* Residual calculation and error estimation.
* Test fixtures and expected failure cases.
* Null-model comparison plan.

## 12. Null-Model Design Review

Null models are future obligations to detect false positives:
* Analytic plane-wave controls.
* Wrong-sign or wrong-operator controls.
* Wrong-potential controls.
* Randomized phase/amplitude controls.
* Coarse-grid and boundary-artifact false-positive controls.

## 13. Faithfulness-to-Literature Review

Verify provenance and assumptions:
* List of source papers and books.
* Mapping of notations and assumptions.
* Unimplemented or untested claims.
* Open human faithfulness questions.

## 14. Classical Target/Recovery Criterion Review

Define classical recovery targets:
* Target classical equations (e.g. Friedmann, Klein-Gordon).
* Variable mapping and clock/time conventions.
* Acceptable error metrics.
* Null model comparisons.
* Explicit statement that no recovery is currently claimed or demonstrated.

## 15. Failure-Mode Registry

The following failure modes block EBP promotion and solver work:
* `operator_form_ambiguous`: WdW equation is not physically or mathematically defined.
* `factor_ordering_unpaid`: Factor ordering lacks physical or mathematical consensus.
* `units_convention_ambiguous`: Coordinate units or constants are not unified.
* `boundary_condition_ambiguous`: Domain boundary conditions are not declared.
* `solver_residual_not_authoritative`: Stencils or near-node evaluations do not report authority loss.
* `null_model_not_defined`: No false-positive null models are established.
* `faithfulness_review_missing`: No literature provenance tracking is provided.
* `classical_target_undefined`: Target classical equations and clock mappings are missing.

## 16. Human Physics-Review Checklist

A human review panel must verify:
* Physical correctness of the Hamiltonian constraint.
* Mathematical self-consistency of boundary conditions.
* Physical significance of the classical target mapping.

## 17. No-Promotion Audit

The gate enforces that:
* Passing BMC-GATE-0001 does not implement a solver.
* Passing BMC-GATE-0001 does not validate BMC.
* Passing BMC-GATE-0001 does not recover Friedmann dynamics.
* Passing BMC-GATE-0001 does not prove classical-limit recovery.
* Passing BMC-GATE-0001 does not defeat null models.
* Passing BMC-GATE-0001 does not unblock full BMC.

The maximum allowed EBP stage is:
`BMC-GATE-0001 accepted_for_gate_design_scope`

## 18. Solver Implementation Requires a Separate Gate/Ticket

Passing BMC-GATE-0001 only defines requirements; it does not authorize implementation. Any future solver, trajectory integration, or physics promotion branch must start from a separate, explicit ticket. Implementation remains blocked until a new gate achieves a completed:
* operator-form review
* factor-ordering review
* units-convention review
* boundary-condition review
* faithfulness review
* null-model design
* solver design
* human physics review
* no-promotion audit

## 19. Explicit Blocked Status After This Gate

* **Full BMC Gate Status**: Blocked.
* **Solver Status**: Not implemented.
* **BMC-0B Status**: Specified only.
* **Promotion Recommendation**: Do not promote.
