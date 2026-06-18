# BMC-REVIEW-0001: Operator-Form, Units, and Convention Faithfulness Review

This document records the first faithfulness review under BMC-GATE-0001 for the BMC-0B Wheeler-DeWitt (WdW) problem statement. It maps current project assumptions, candidate operator forms, unit/convention choices, factor-ordering alternatives, boundary-condition assumptions, residual definitions, and literature provenance. It does not implement a solver, produce numerical BMC-0B results, validate BMC, recover Friedmann dynamics, or unblock full BMC.

## 1. Review Purpose and Non-Goals

### Purpose
- Draft a candidate BMC-0B mathematical problem statement for human review.
- Record the current project operator-form assumption and possible alternatives.
- Freeze a units/convention table to prevent silent switching.
- Identify the factor-ordering debt and evidence needed to retire it.
- Document boundary-condition, admissibility, residual, and classical-recovery assumptions.
- Produce a source/provenance ledger with explicit debt markers.

### Non-Goals
- BMC-REVIEW-0001 does not implement BMC-0B.
- BMC-REVIEW-0001 does not solve the Wheeler-DeWitt equation.
- BMC-REVIEW-0001 does not validate BMC.
- BMC-REVIEW-0001 does not recover Friedmann dynamics.
- BMC-REVIEW-0001 does not prove classical-limit recovery.
- BMC-REVIEW-0001 does not defeat null models.
- BMC-REVIEW-0001 does not unblock full BMC.

Maximum allowed EBP stage: `BMC-REVIEW-0001 accepted_for_faithfulness_review_scope`

## 2. Current Accepted Stack Boundary

This review must respect the accepted limited scopes of the closed remediation stack:
- **POST-0001**: Accepted only for plane-wave/toy WdW constraint violation detection scope.
- **POST-0002 / 0002.1**: Accepted only for plane-wave numerical residual report-path authority scope.
- **POST-0003 / 0003.2**: Accepted only for superposition trajectory numerical self-consistency/convergence audit scope.
- **POST-0004 / 0004.2**: Accepted only for BMC-0B massive scalar specification-only scope.
- **POST-0005**: Accepted only for phase-gradient h-sensitivity audit scope.
- **POST-0006**: Accepted only for quantum-potential near-node domain-boundary audit scope.
- **POST-0007**: Accepted only for branch-cut and stencil-boundary regression-hardening scope.
- **POST-0008**: Accepted only for stack closure ledger and freeze scope.
- **GATE-0001**: Accepted only for faithfulness-review and solver-design gate scope.

## 3. Candidate BMC-0B Mathematical Problem Statement (Unreviewed)

> Status: `candidate_problem_statement_unreviewed`

The candidate problem statement below collects current fixture and specification assumptions. Nothing in this section is final. Each item carries a debt marker indicating whether it is `imported_from_project_docs`, `assumed_by_current_fixture`, `requires_literature_confirmation`, or `requires_human_physics_review`.

### 3.1 Configuration Variables and Domain
- Variables: `alpha`, `phi` (POST-0004 spec; `imported_from_project_docs`)
- Coordinate convention: `alpha = ln(a)` (`imported_from_project_docs`)
- Scalar field variable: `phi` (`imported_from_project_docs`)
- Domain of `alpha` and `phi`: not yet fixed (`BoundaryConditionDebt: unpaid`)
- Grid bounds and resolution: specified-only contract in POST-0004; numerical values remain `required_before_solver`

### 3.2 Minisuperspace Metric/Signature
- Metric signature convention: unspecified in current project docs (`MetricSignatureDebt: unpaid`)
- The signature determines the WdW operator's hyperbolic/elliptic character and must be fixed before any operator form is declared faithful.

### 3.3 Hamiltonian Constraint
- Hamiltonian constraint form: not yet declared in project docs (`OperatorFormDebt: unpaid`)
- Candidate form requires explicit statement of kinetic term, potential term, and constraint surface.

### 3.4 Wheeler-DeWitt Operator Form
- Current fixture convention: `FiniteDifferenceResidual` evaluates `-d2Psi/dAlpha2 + d2Psi/dPhi2` (POST-0001/0002 code)
- This numerical operator is a diagnostic, not a faithful statement of the continuous WdW operator (`imported_from_project_docs`)
- Factor ordering is not encoded in the fixture (`FactorOrderingDebt: unpaid`)
- Sign conventions, potential placement, and metric-dependent kinetic terms remain unreviewed (`OperatorFormDebt: unpaid`)

### 3.5 Massive Scalar Potential Term
- Potential normalization: not declared in project docs (`ScalarPotentialNormalizationDebt: unpaid`)
- Mass parameter convention: not declared (`UnitsConventionDebt: unpaid`)

### 3.6 Numerical Residual Definition
- Current numerical residual: `-d2Psi/dAlpha2 + d2Psi/dPhi2` (`imported_from_project_docs`)
- Residual norm obligations in POST-0004: `l2_residual_norm`, `linf_residual_norm`, `boundary_residual_norm` (`imported_from_project_docs`)
- Tolerance justification: `explicit_unpaid` (`ResidualDefinitionDebt: unpaid`)
- Authoritative pass/fail criteria: not yet defined for BMC-0B (`ResidualDefinitionDebt: unpaid`)

### 3.7 Wavefunction Admissibility
- Admissible class: not declared (`WavefunctionAdmissibilityDebt: unpaid`)
- Node/near-node policy: amplitude floor `1e-8` in audit code (`imported_from_project_docs`)
- Nonfinite-value policy: blocked in audit code (`imported_from_project_docs`)
- Superposition authority: `audit_only_no_promotion` in phase and Q audits (`imported_from_project_docs`)

### 3.8 Classical Recovery Criterion
- No recovery claim is currently demonstrated (`ClassicalRecoveryCriterionDebt: unpaid`)
- Target classical equations, mapping, clock conventions, and error metrics are undeclared.

## 4. Configuration Variables and Domain Assumptions

| Item | Current Project Assumption | Required Human Review |
|---|---|---|
| `alpha` | `ln(a)` coordinate | Confirm minisuperspace coordinate choice and domain bounds |
| `phi` | Scalar field variable | Confirm field variable and physical range |
| Domain bounds | Not fixed | Specify `alpha`/`phi` bounds and grid resolution policy |
| Boundary conditions | `explicit_unpaid` | Declare admissible boundary conditions before solver design |

## 5. Minisuperspace Metric/Signature Convention

| Item | Current Project Assumption | Possible Alternative | Debt Status |
|---|---|---|---|
| Metric signature | Not declared | `(-,+)` or `(+,-)` on `(alpha, phi)` | `unpaid` |
| Grid alignment | Not declared | Aligned vs. staggered stencils | `unpaid` |
| Coordinate singularity handling | Not declared | Regular grid vs. excision | `unpaid` |

The metric signature must be fixed before the operator form is declared faithful, because it controls the sign of kinetic terms and the PDE character.

## 6. Hamiltonian Constraint Form

- Current status: not explicitly declared in project docs (`OperatorFormDebt: unpaid`)
- Required inputs before solver design:
  - Kinetic operator form
  - Scalar potential placement
  - Constraint surface definition
  - Source equation or literature reference

## 7. Wheeler-DeWitt Operator-Form Candidates

| Candidate ID | Operator Form | Sign Convention | Factor Ordering | Scalar Potential Placement | Source/Provenance | Impact on Residual | Impact on Guidance/Q | Debt Status |
|---|---|---|---|---|---|---|---|---|
| `operator_candidate_project_fixture_only` | `-d2/dAlpha2 + d2/dPhi2` (finite-difference diagnostic) | Signs match toy residual evaluator in POST-0001/0002 | Not encoded | Not included | `imported_from_project_docs` | Defines current numerical residual diagnostic only | Phase-gradient and Q conventions are separate | `requires_literature_confirmation` |
| Alternative orderings | Not reviewed | Unknown until literature review is performed | Unknown | Unknown | `requires_literature_review` | Unknown | Unknown | `unpaid` |

Current project fixtures (`wave.PlaneWave`, `wdw.FiniteDifferenceResidual`) do not encode a faithful continuous operator form, factor ordering, or potential term. Any future solver proposal must present operator-form candidates from the literature with explicit source provenance.

## 8. Massive Scalar Potential Normalization

| Item | Current Project Assumption | Possible Alternative | Debt Status |
|---|---|---|---|
| Potential form | Not declared | `V(phi) = (1/2) m^2 phi^2` or alternative | `unpaid` |
| Mass parameter convention | Not declared | Planck units, reduced mass, etc. | `unpaid` |
| Potential normalization | Not declared | Dimensionless vs. dimensional normalization | `unpaid` |

Until the potential normalization is fixed, residuals cannot be compared across different unit systems.

## 9. Units and Constants Convention Table

| Convention | Current Project Assumption | Possible Alternative | Source/Provenance | Debt Status | Human Review Required |
|---|---|---|---|---|---|
| `alpha = ln(a)` | Used in fixtures and docs | Alternative scale factor parametrization | `imported_from_project_docs` | `requires_literature_confirmation` | Yes |
| `phi` scalar field variable | Used in fixtures and docs | Alternative field variable or normalization | `imported_from_project_docs` | `requires_literature_confirmation` | Yes |
| `lambda` trajectory parameter | Referenced in docs; not fixed in BMC-0B scope | Time coordinate vs. proper time | `imported_from_project_docs` | `requires_human_physics_review` | Yes |
| Unit system | Not declared | Planck units vs. alternative | `external_literature_missing` | `unpaid` | Yes |
| `hbar` convention | Not declared | `hbar = 1` vs. explicit `hbar` | `external_literature_missing` | `unpaid` | Yes |
| `G` convention | Not declared | `G = 1` vs. explicit `G` | `external_literature_missing` | `unpaid` | Yes |
| `c` convention | Not declared | `c = 1` vs. explicit `c` | `external_literature_missing` | `unpaid` | Yes |
| Mass parameter convention | Not declared | Dimensionless mass vs. physical mass | `external_literature_missing` | `unpaid` | Yes |
| Potential normalization | Not declared | `V = (1/2) m^2 phi^2` vs. alternative | `external_literature_missing` | `unpaid` | Yes |
| Minisuperspace metric signature | Not declared | `(-,+)` or `(+,-)` | `external_literature_missing` | `unpaid` | Yes |
| WdW sign convention | Toy residual uses `-d2/dAlpha2 + d2/dPhi2` | Literature-dependent sign choices | `imported_from_project_docs` | `requires_literature_confirmation` | Yes |
| Phase-gradient convention | `Im((1/Psi) * dPsi/dx)` (POST-0005/0007 code) | Equivalent analytic identity confirmed | `imported_from_project_docs` | `requires_literature_confirmation` | Yes |
| Guidance-sign convention | Phase gradient documented as separate from Bohmian velocity (POST-0005 doc) | Confirm trajectory velocity sign convention | `imported_from_project_docs` | `requires_human_physics_review` | Yes |
| Q-potential convention | `Q = -1/(2R) * (d2R/dAlpha2 - d2R/dPhi2)` (POST-0006 doc and code) | Literature-dependent factor ordering | `imported_from_project_docs` | `requires_literature_confirmation` | Yes |
| Inner product or diagnostic measure | Not declared | L2 inner product vs. alternative | `external_literature_missing` | `unpaid` | Yes |
| Residual norm convention | POST-0004 spec lists `l2`, `linf`, `boundary` | Justify tolerance thresholds | `imported_from_project_docs` | `requires_human_physics_review` | Yes |

## 10. Factor-Ordering Alternatives

- Current assumed ordering: not fixed (`FactorOrderingDebt: unpaid`)
- Alternatives not yet reviewed: supersymmetric orderings, curvature-coupled orderings, etc. (`requires_literature_review`)
- Whether ordering changes residuals: unknown until operator form is fixed (`unpaid`)
- Whether ordering changes boundary behavior: unknown (`unpaid`)
- Whether ordering changes Bohmian phase-gradient or Q-potential quantities: unknown (`unpaid`)
- Evidence needed to retire debt:
  - Explicit operator-form candidate from literature with factor ordering stated
  - Numerical comparison of residuals under alternative orderings
  - Human physics review confirming chosen ordering

## 11. Boundary-Condition Assumptions

| Item | Current Project Assumption | Required Declaration |
|---|---|---|
| Domain bounds | Not fixed | Declare `alpha`/`phi` bounds before solver design |
| Grid bounds | Specified-only contract in POST-0004 | Numerical values remain `required_before_solver` |
| Boundary conditions | `explicit_unpaid` | Specify admissible boundary conditions |
| Boundary residual treatment | Not declared | Block or clamp policy for boundary stencils |
| Node/near-node handling | Amplitude floor `1e-8` in audit code | Confirm floor choice and authority policy |
| Nonfinite behavior | Blocked in audit code | Confirm blocking policy in solver design |
| Stencil-boundary behavior | Blocked in POST-0006/0007 code | Confirm domain-boundary contamination policy |

Vague labels such as `default`, `reasonable`, `standard`, `ready`, or `validated` are explicitly rejected for boundary-condition declarations.

## 12. Wavefunction/Domain Admissibility Assumptions

| Item | Current Project Assumption | Debt Status |
|---|---|---|
| Admissible class | Not declared | `WavefunctionAdmissibilityDebt: unpaid` |
| Smoothness/differentiability | Not declared | `unpaid` |
| Node handling | Amplitude floor `1e-8` | `imported_from_project_docs` |
| Near-node threshold | `1e-8` in audit code | `imported_from_project_docs` |
| Nonfinite-value policy | Blocked | `imported_from_project_docs` |
| Domain-boundary policy | Blocked in Q audit | `imported_from_project_docs` |
| Superposition authority | `audit_only_no_promotion` | `imported_from_project_docs` |

Square-integrability, behavior at infinite limits, and boundary regularity must be declared before solver design.

## 13. Residual Definition Implications

- Authoritative residual definition: not yet declared (`ResidualDefinitionDebt: unpaid`)
- Residual norm: POST-0004 lists `l2`, `linf`, `boundary` (`imported_from_project_docs`)
- Tolerance threshold: `explicit_unpaid` (`ResidualDefinitionDebt: unpaid`)
- Grid/refinement policy: not declared (`ResidualDefinitionDebt: unpaid`)
- Pass/fail criteria: not declared (`ResidualDefinitionDebt: unpaid`)
- Diagnostic authority rules: numerical residual is diagnostic authority for plane-wave report path (POST-0002); superposition remains `not_authoritative` (`imported_from_project_docs`)
- Oracle/control comparison rules: analytic vs. numerical disagreement triggers `NumericalResidualError` (POST-0002) (`imported_from_project_docs`)

## 14. Current Toy-Fixture Convention Map

| Fixture | Convention Reflected | Source | Status |
|---|---|---|---|
| `wave.PlaneWave` | `Psi(alpha,phi) = exp(i(k*alpha + omega*phi))` | `imported_from_project_docs` | Control fixture only |
| `wave.SuperpositionWave` | Two-plane-wave superposition with complex coefficients | `imported_from_project_docs` | Audit-only, not physics validation |
| `wdw.FiniteDifferenceResidual` | `-d2Psi/dAlpha2 + d2Psi/dPhi2` central finite difference | `imported_from_project_docs` | Toy diagnostic, not continuous operator |
| `wave.PhaseGradient` | `Im((1/Psi) * dPsi/dx)` identity | `imported_from_project_docs` | Audit machinery |
| `qpotential.Q` | `-1/(2R)*(d2R/dAlpha2 - d2R/dPhi2)` | `imported_from_project_docs` | Audit machinery with near-node clamping |

No fixture encodes a faithful continuous WdW operator form with factor ordering, potential term, or metric signature.

## 15. Literature/Source Provenance Ledger

| Claim or Equation | Current Source | Source Type | Status | Debt | Human Review Question |
|---|---|---|---|---|---|
| Plane-wave control `Psi = exp(i(k*alpha + omega*phi))` | Project docs and fixtures | `code_fixture` | `imported_from_project_docs` | `requires_literature_confirmation` | Which paper introduces these coordinates and conventions? |
| Numerical residual `-d2/dAlpha2 + d2/dPhi2` | POST-0001/0002 residual.go | `project_doc` | `imported_from_project_docs` | `requires_literature_confirmation` | Which continuous operator does this discretize? |
| Phase-gradient identity `Im((1/Psi)*dPsi/dx)` | POST-0005/0007 docs | `project_doc` | `imported_from_project_docs` | `requires_literature_confirmation` | Confirm equivalence to standard definition |
| Q-potential form `Q = -1/(2R)*(d2R/dAlpha2 - d2R/dPhi2)` | POST-0006 doc and q.go | `project_doc` | `imported_from_project_docs` | `requires_literature_confirmation` | Which Bohmian Q potential formulation is used? |
| WdW operator form for massive scalar | Not declared | `external_literature_missing` | `unreviewed` | `OperatorFormDebt: unpaid` | What is the exact WdW equation with massive scalar potential? |
| Factor ordering | Not declared | `external_literature_missing` | `unreviewed` | `FactorOrderingDebt: unpaid` | Which ordering is physically motivated? |
| Metric/signature | Not declared | `external_literature_missing` | `unreviewed` | `MetricSignatureDebt: unpaid` | What signature is used in the minisuperspace model? |
| Units/conventions (hbar, G, c, mass) | Not declared | `external_literature_missing` | `unreviewed` | `UnitsConventionDebt: unpaid` | What unit system and constant conventions are assumed? |
| Boundary conditions | Not declared | `external_literature_missing` | `unreviewed` | `BoundaryConditionDebt: unpaid` | What boundary conditions are physically admissible? |
| Classical recovery target | Not declared | `external_literature_missing` | `unreviewed` | `ClassicalRecoveryCriterionDebt: unpaid` | What equations must be recovered and how? |

## 16. Project-Assumption Ledger

All items below are labeled with their current provenance and debt status. None are final.

| Assumption | Provenance | Debt Status |
|---|---|---|
| `alpha = ln(a)` | `imported_from_project_docs` | `requires_literature_confirmation` |
| `phi` scalar field variable | `imported_from_project_docs` | `requires_literature_confirmation` |
| Plane-wave phase-gradient identity | `imported_from_project_docs` | `requires_literature_confirmation` |
| Q-potential formula | `imported_from_project_docs` | `requires_literature_confirmation` |
| Numerical residual form `-d2/dAlpha2 + d2/dPhi2` | `imported_from_project_docs` | `requires_literature_confirmation` |
| Near-node amplitude floor `1e-8` | `imported_from_project_docs` | `requires_human_physics_review` |
| Superposition authority: `audit_only_no_promotion` | `imported_from_project_docs` | `requires_human_physics_review` |
| No solver implemented | `imported_from_project_docs` | `SolverStatus: not_implemented` |
| BMC-0B status: `specified_only` | `imported_from_project_docs` | `BMC0BStatus: specified_only` |
| Operator form: unreviewed | `assumed_by_current_fixture` | `OperatorFormDebt: unpaid` |
| Factor ordering: unreviewed | `assumed_by_current_fixture` | `FactorOrderingDebt: unpaid` |
| Units/conventions: unreviewed | `assumed_by_current_fixture` | `UnitsConventionDebt: unpaid` |
| Boundary conditions: unreviewed | `assumed_by_current_fixture` | `BoundaryConditionDebt: unpaid` |
| Metric/signature: unreviewed | `assumed_by_current_fixture` | `MetricSignatureDebt: unpaid` |
| Potential normalization: unreviewed | `assumed_by_current_fixture` | `ScalarPotentialNormalizationDebt: unpaid` |
| Wavefunction admissibility: unreviewed | `assumed_by_current_fixture` | `WavefunctionAdmissibilityDebt: unpaid` |
| Residual definition: unreviewed | `assumed_by_current_fixture` | `ResidualDefinitionDebt: unpaid` |
| Null-model design: unreviewed | `assumed_by_current_fixture` | `NullModelDesignDebt: unpaid` |
| Classical recovery criterion: unreviewed | `assumed_by_current_fixture` | `ClassicalRecoveryCriterionDebt: unpaid` |

## 17. Unpaid Debt Ledger

The following debts remain unpaid unless explicitly retired by cited source review:

| Debt | Status |
|---|---|
| `OperatorFormDebt` | `unpaid` |
| `FactorOrderingDebt` | `unpaid` |
| `UnitsConventionDebt` | `unpaid` |
| `BoundaryConditionDebt` | `unpaid` |
| `MetricSignatureDebt` | `unpaid` |
| `ScalarPotentialNormalizationDebt` | `unpaid` |
| `WavefunctionAdmissibilityDebt` | `unpaid` |
| `ResidualDefinitionDebt` | `unpaid` |
| `NullModelDesignDebt` | `unpaid` |
| `ClassicalRecoveryCriterionDebt` | `unpaid` |
| `FaithfulnessReviewDebt` | `partial_documentation_only` |
| `HumanPhysicsReviewDebt` | `unpaid` |
| `LeanProofDebt` | `unpaid` |

## 18. Human Physics-Review Questions

A human review panel must answer, at minimum:
- What is the exact WdW equation with massive scalar potential, and from which source is it taken?
- Which signs are convention choices versus mathematically forced by the chosen minisuperspace model?
- What factor ordering is adopted, and what is the physical motivation?
- What metric signature is used, and does it match the chosen coordinate system?
- What are the admissible boundary conditions, and why?
- What units and constants are fixed, and what is the source?
- What classical equations are targeted for recovery, and what is the mapping from `(alpha, phi)` to classical variables?
- What null models must be defeated before any numerical result is trusted?
- What does it mean for a numerical residual to be authoritative, and when is it not?
- What human-review evidence would retire each unpaid debt?

## 19. No-Promotion Audit

BMC-REVIEW-0001 does not implement BMC-0B.
BMC-REVIEW-0001 does not solve the Wheeler-DeWitt equation.
BMC-REVIEW-0001 does not validate BMC.
BMC-REVIEW-0001 does not recover Friedmann dynamics.
BMC-REVIEW-0001 does not prove classical-limit recovery.
BMC-REVIEW-0001 does not defeat null models.
BMC-REVIEW-0001 does not unblock full BMC.

The maximum allowed EBP stage is:
`BMC-REVIEW-0001 accepted_for_faithfulness_review_scope`

## 20. Next Allowed Ticket Recommendation

Allowed next branches:
- Additional faithfulness review tickets for specific unresolved debts (operator form, factor ordering, units/conventions, boundary conditions, classical recovery criterion).
- Human physics-review checklist ticket.
- Null-model design document ticket.
- BMC-0B solver design document ticket (separate from implementation).
- Additional negative regression tests for the audit stack.

Disallowed next branches:
- Immediate BMC-0B solver implementation.
- Claiming Friedmann recovery.
- Promoting BMC.
- Adding benchmark result artifacts.
- Claiming null-model failure.
- Claiming scientific novelty.
- Unblocking full BMC.

---
