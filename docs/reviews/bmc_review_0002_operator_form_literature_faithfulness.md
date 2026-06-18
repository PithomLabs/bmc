# BMC-REVIEW-0002: Operator-Form Literature Faithfulness Review

This document records the literature faithfulness review of minisuperspace Wheeler-DeWitt (WdW) operator forms under strict EBP 2.1 discipline. It maps candidate operator forms, metric signatures, factor orderings, potential normalizations, notation conventions, and boundary implications against the current project fixture. 

---

## 1. Review Purpose and Non-Goals

### Purpose
The primary purpose of this review is to map candidate minisuperspace Wheeler-DeWitt operator forms against the current project fixture, answering the following questions:
- What WdW operator forms for homogeneous scalar-field cosmology appear in the available sources?
- How do their signs, metric/signature choices, factor orderings, potential terms, and unit conventions differ?
- Which parts of the current PTW/BMC fixture match a literature-backed form?
- Which parts remain project-only assumptions?
- What operator-form candidates should be carried forward for human physics review?

### Non-Goals
- This review is strictly documentation-only.
- It does not implement a solver.
- It does not produce numerical BMC-0B results.
- It does not retire `OperatorFormDebt` or any related physical debts.
- It does not solve the Wheeler-DeWitt equation.
- It does not validate BMC.
- It does not recover Friedmann dynamics.
- It does not prove classical-limit recovery.
- It does not defeat null models.
- It does not unblock full BMC.

Maximum allowed EBP stage: `BMC-REVIEW-0002 accepted_for_operator_form_review_scope`

---

## 2. Current Accepted Boundary from REVIEW-0001

BMC-REVIEW-0001 accepted only a documentation/provenance scope. It explicitly left the following debts unpaid:
- `OperatorFormDebt`
- `FactorOrderingDebt`
- `UnitsConventionDebt`
- `BoundaryConditionDebt`
- `MetricSignatureDebt`
- `ScalarPotentialNormalizationDebt`
- `WavefunctionAdmissibilityDebt`
- `ResidualDefinitionDebt`
- `NullModelDesignDebt`
- `ClassicalRecoveryCriterionDebt`
- `FaithfulnessReviewDebt`
- `HumanPhysicsReviewDebt`
- `LeanProofDebt`

This review (BMC-REVIEW-0002) targets the most upstream debt, `OperatorFormDebt`, but because external literature is not available in the repository, this review cannot retire `OperatorFormDebt` and leaves it, along with all dependent debts, unpaid.

---

## 3. Source Availability Audit

No external literature equation is available in-repo for retiring OperatorFormDebt. This review can only map current project assumptions and define literature-acquisition requirements. OperatorFormDebt remains unpaid.

| Source or citation | Location in repo | Source type | Usable equation present? | Relevant to BMC-0B operator form? | Status | Debt |
|---|---|---|---|---|---|---|
| ChatGPT/Kimi notes / chat transcripts | [bqg01_chatgpt.md](file:///home/chaschel/Documents/go/bmc/bqg01_chatgpt.md), [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md), [bqg_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg_kimi.md) | `project_doc` | Yes (toy models discussed in text) | Only as pre-conceptual discussion | `imported_from_project_docs` | `OperatorFormDebt: unpaid` |
| Sprint 8 plans | [plan.md](file:///home/chaschel/Documents/go/bmc/sprint/sprint8/plan.md), [plan2.md](file:///home/chaschel/Documents/go/bmc/sprint/sprint8/plan2.md), [plan3.md](file:///home/chaschel/Documents/go/bmc/sprint/sprint8/plan3.md) | `project_doc` | No (mentions external arXiv URLs, e.g., 2512.18818v2) | Only as bibliographical references | `imported_from_project_docs` | `OperatorFormDebt: unpaid` |
| Numerical code | [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) | `code_fixture` | Yes (toy finite-difference equation) | Only as diagnostic check | `assumed_by_fixture` | `OperatorFormDebt: unpaid` |
| External peer-reviewed literature | None present | `external_literature_missing` | No | Yes (standard massive scalar WdW equations) | `external_literature_missing` | `OperatorFormDebt: unpaid` |

---

## 4. Current Project Fixture Operator Map

The project currently uses a hardcoded finite-difference residual for verification:

- **Current fixture residual**: `-d2Psi/dAlpha2 + d2Psi/dPhi2`
- **Status**: `toy diagnostic / project fixture only`
- **Not yet faithful because**:
  - Continuous operator form is not sourced from peer-reviewed literature.
  - Minisuperspace metric/signature is not fixed.
  - Wavefunction factor ordering is not encoded or resolved.
  - Massive scalar field potential term is not included in the operator.
  - Coordinate unit and physical constant conventions are not fixed.
  - Domain boundary conditions for $\alpha$ and $\phi$ are not fixed.

The current code fixture is a toy evaluator for testing constraint violation detection machinery. It is not a representation of a faithful Wheeler-DeWitt constraint operator.

---

## 5. Literature Operator-Form Candidates

Due to the absence of peer-reviewed publications within the repository, candidate operator forms from the literature cannot be verified or mapped. They are classified below as acquisition requirements rather than identified alternatives.

| Candidate ID | Source/provenance | Equation form | Variables | Metric/signature convention | Kinetic term | Potential term | Factor ordering | Units/constants | Boundary assumptions | Relation to current fixture | Debt status | Human review required |
|---|---|---|---|---|---|---|---|---|---|---|---|---|
| `operator_candidate_project_fixture_only` | [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) | `(-d2/dAlpha2 + d2/dPhi2)Psi = 0` | `alpha`, `phi` | `(-,+)` | `-d2/dAlpha2 + d2/dPhi2` | None | None | None | None | Identical to current code fixture | `assumed_by_fixture` | Yes |
| `operator_candidate_literature_required_001` | External literature acquisition | Not available | Not available | Not available | Not available | Not available | Not available | Not available | Not available | Not available | `OperatorFormDebt: unpaid` | Yes |
| `operator_candidate_literature_required_002` | External literature acquisition | Not available | Not available | Not available | Not available | Not available | Not available | Not available | Not available | Not available | `OperatorFormDebt: unpaid` | Yes |

### Status Details:
- `operator_candidate_literature_required_001`:
  - status: `external_literature_missing`
  - debt_status: `OperatorFormDebt: unpaid`
  - carry_forward_status: `source_acquisition_required`
- `operator_candidate_literature_required_002`:
  - status: `external_literature_missing`
  - debt_status: `OperatorFormDebt: unpaid`
  - carry_forward_status: `source_acquisition_required`

---

## 6. Notation and Convention Mapping

Because no peer-reviewed literature is present in the repository, all literature notations are marked as unavailable.

| Project notation | Source notation | Meaning | Mapping confidence | Debt status | Human review question |
|---|---|---|---|---|---|
| `alpha` | `source_not_available` | Scale factor logarithm: `ln(a)` | `requires_literature_confirmation` | `UnitsConventionDebt: unpaid` | Is `alpha` defined as `ln(a)` or does it carry additional scaling constants? |
| `a` | `source_not_available` | FRW scale factor | `requires_literature_confirmation` | `UnitsConventionDebt: unpaid` | Does the coordinate range allow $a \to 0$, or is it bounded away from the singularity? |
| `phi` | `source_not_available` | Homogeneous scalar field variable | `requires_literature_confirmation` | `UnitsConventionDebt: unpaid` | Is the field normalized with respect to the reduced Planck mass? |
| mass parameter | `source_not_available` | Mass of the scalar field ($m$) | `requires_literature_confirmation` | `UnitsConventionDebt: unpaid` | How does the mass parameter scale under units conversion? |
| potential term | `source_not_available` | Massive scalar field potential: $V(\phi) = \frac{1}{2}m^2\phi^2$ | `requires_literature_confirmation` | `ScalarPotentialNormalizationDebt: unpaid` | What normalization factor is applied to $V(\phi)$ in the Hamiltonian? |
| lapse/time or clock parameter | `source_not_available` | Trajectory parametrization clock parameter | `requires_literature_confirmation` | `UnitsConventionDebt: unpaid` | Is the trajectory time variable equivalent to proper time? |
| WdW wavefunction Psi | `source_not_available` | Minisuperspace wavefunction $\Psi(\alpha, \phi)$ | `requires_literature_confirmation` | `WavefunctionAdmissibilityDebt: unpaid` | What boundary conditions are physically and mathematically admissible for $\Psi$? |
| residual operator | `source_not_available` | Wheeler-DeWitt constraint operator | `requires_literature_confirmation` | `OperatorFormDebt: unpaid` | Does the continuous operator match the discrete central finite difference? |

---

## 7. Metric/Signature Mapping

The minisuperspace metric signature on the $(\alpha, \phi)$ configuration space determines the hyperbolic or elliptic character of the Wheeler-DeWitt partial differential equation. The current project documentation does not fix this signature.

MetricSignatureDebt remains unpaid. OperatorFormDebt cannot be retired until metric/signature is fixed.

| Candidate signature | Effect on kinetic signs | Effect on WdW operator | Relation to current fixture | Debt status |
|---|---|---|---|---|
| `(-,+)` on `(alpha, phi)` | Negates `alpha` derivative relative to `phi` derivative | Yields `-d2/dAlpha2 + d2/dPhi2` | Matches current toy residual sign structure | `MetricSignatureDebt: dependent_unpaid` |
| `(+,-)` on `(alpha, phi)` | Negates `phi` derivative relative to `alpha` derivative | Yields `+d2/dAlpha2 - d2/dPhi2` | Opposite signs to current toy residual | `MetricSignatureDebt: dependent_unpaid` |
| General metric $f(a)$ | Introduces metric coefficients to derivatives | Non-constant coefficients on derivatives | Diverges from current toy residual | `MetricSignatureDebt: dependent_unpaid` |

---

## 8. Kinetic-Term Comparison

Continuous operator forms in flat FRW minisuperspace require applying the Laplace-Beltrami operator with respect to the minisuperspace metric. Depending on coordinates (e.g., scale factor $a$ vs. $\alpha = \ln(a)$), the kinetic operator term has different forms (e.g. including first-derivative terms $a \frac{\partial}{\partial a}$ or $\frac{\partial}{\partial a} a \frac{\partial}{\partial a}$).

Without literature sources, the kinetic signs, coefficients, and coordinate choices remain project-only assumptions.

---

## 9. Scalar-Potential Placement Comparison

Review whether the massive scalar potential term is specified.

ScalarPotentialNormalizationDebt remains unpaid. BMC-0B operator form cannot be finalized without potential normalization.

| Potential candidate | Placement in Hamiltonian constraint | Placement in WdW operator | Units/constants required | Source/provenance | Debt status |
|---|---|---|---|---|---|
| $V(\phi) = \frac{1}{2}m^2\phi^2$ | Added to scalar kinetic energy | Appears as multiplicative operator term $e^{6\alpha}V(\phi)\Psi$ | Planck mass, $\hbar$, $G$, $c$ | `external_literature_missing` | `ScalarPotentialNormalizationDebt: dependent_unpaid` |
| $V(\phi) = 0$ | Absent | Absent | None | `code_fixture` | `assumed_by_fixture` |

---

## 10. Factor-Ordering Comparison

FactorOrderingDebt remains unpaid. No solver design may assume this debt retired.

- **Current fixture ordering status**: Not encoded. The toy operator `-d2/dAlpha2 + d2/dPhi2` contains no first-derivative or cross-derivative terms.
- **Available candidate orderings**: Laplacian ordering, covariant ordering, Vilkovisky-DeWitt ordering.
- **Missing literature orderings**: Detailed factor-ordering equations for minisuperspace models with scalar fields are missing in-repo.
- **Impact on residual**: Factor-ordering terms introduce first-derivative terms which change the discrete residual evaluation.
- **Impact on boundary behavior**: Boundary terms and singularities at $a \to 0$ are highly sensitive to first-derivative factor-ordering coefficients.
- **Impact on Bohmian guidance/Q-potential**: Alterations to factor ordering modify the real and imaginary parts of the wavefunction constraint, affecting phase gradient and the quantum potential $Q$.

---

## 11. Units/Constants Implications

The choice of unit systems (e.g., Planck units, reduced Planck units, or dimensionless conventions) alters the physical scaling of the coordinates and the mass parameter. If constants like $\hbar$, $G$, and $c$ are set to unity, their dimensional scaling is lost.

Without fixing unit conversions, residual tolerances cannot be normalized or compared. `UnitsConventionDebt` remains unpaid.

---

## 12. Boundary/Domain Implications

Domain boundaries ($a \to 0$, $a \to \infty$) must be mapped to grid boundaries in $\alpha$. If $a \to 0$ is the singularity, then $\alpha \to -\infty$. A grid boundary in numerical stencils must truncate this domain, introducing boundary artifacts.

Without physical literature guiding boundary conditions (Dirichlet, Neumann, or DeWitt vanishing wavefunction), boundary-condition stencils remain arbitrary. `BoundaryConditionDebt` remains unpaid.

---

## 13. Residual-Definition Implications

Mathematical residuals depend on the volume element of the minisuperspace configuration space. L2 norms require a choice of measure. Without a literature-backed operator form and metric, the residual norm cannot be defined. `ResidualDefinitionDebt` remains unpaid.

---

## 14. Impact on Bohmian Guidance and Q-Potential Quantities

The guidance equation relates trajectory velocities to spatial derivatives of the wavefunction phase. Changing the kinetic terms or metric coefficients in the WdW operator directly alters the guidance equations. The quantum potential $Q$ is computed from the spatial curvature of the wavefunction amplitude $R$, which depends on the chosen configuration space metric and factor ordering.

---

## 15. Operator-Form Candidate Carry-Forward List

All candidates listed below are carried forward as requirements or control references for future human physics review. None are authorized for solver implementation.

| Candidate carried forward | Why carried forward | Debt status | Required source review | Required human review | Blocked implementation reason |
|---|---|---|---|---|---|
| `operator_candidate_project_fixture_only` | Control fixture diagnostic in existing test suite | `assumed_by_fixture` | `requires_literature_confirmation` | `requires_human_physics_review` | `project_fixture_only`, `not faithful yet`, `blocks solver implementation` |
| `operator_candidate_literature_required_001` | To represent standard minisuperspace WdW | `OperatorFormDebt: unpaid` | `source_acquisition_required` | `requires_human_physics_review` | `external_literature_missing`, `blocks solver implementation` |
| `operator_candidate_literature_required_002` | To represent alternative factor orderings | `OperatorFormDebt: unpaid` | `source_acquisition_required` | `requires_human_physics_review` | `external_literature_missing`, `blocks solver implementation` |

---

## 16. OperatorFormDebt Status

`OperatorFormDebt: unpaid`

Because no usable peer-reviewed literature is present in the repository, `OperatorFormDebt` remains unpaid.

---

## 17. Remaining Dependent Debts

The following dependent debts remain unpaid and block solver design and implementation:
- `MetricSignatureDebt: dependent_unpaid`
- `ScalarPotentialNormalizationDebt: dependent_unpaid`
- `FactorOrderingDebt: dependent_unpaid`
- `UnitsConventionDebt: dependent_unpaid`
- `BoundaryConditionDebt: dependent_unpaid`
- `WavefunctionAdmissibilityDebt: dependent_unpaid`
- `ResidualDefinitionDebt: dependent_unpaid`
- `NullModelDesignDebt: dependent_unpaid`
- `ClassicalRecoveryCriterionDebt: dependent_unpaid`
- `FaithfulnessReviewDebt: partial_documentation_only`
- `HumanPhysicsReviewDebt: unpaid`
- `LeanProofDebt: unpaid`

---

## 18. Human Physics-Review Questions

A human review panel must answer, at minimum:
1. What is the exact continuous WdW equation with massive scalar potential, and from which source is it taken?
2. Which signs are convention choices versus mathematically forced by the chosen minisuperspace model?
3. What factor ordering is adopted, and what is the physical motivation?
4. What metric signature is used, and does it match the chosen coordinate system?
5. What are the admissible boundary conditions, and why?
6. What units and constants are fixed, and what is the source?
7. What classical equations are targeted for recovery, and what is the mapping from `(alpha, phi)` to classical variables?
8. What null models must be defeated before any numerical result is trusted?
9. What does it mean for a numerical residual to be authoritative, and when is it not?
10. What human-review evidence would retire each unpaid debt?

---

## 19. No-Promotion Audit

- BMC-REVIEW-0002 does not implement BMC-0B.
- BMC-REVIEW-0002 does not solve the Wheeler-DeWitt equation.
- BMC-REVIEW-0002 does not validate BMC.
- BMC-REVIEW-0002 does not recover Friedmann dynamics.
- BMC-REVIEW-0002 does not prove classical-limit recovery.
- BMC-REVIEW-0002 does not defeat null models.
- BMC-REVIEW-0002 does not unblock full BMC.

The maximum allowed EBP stage is:
`BMC-REVIEW-0002 accepted_for_operator_form_review_scope`

---

## 20. Next Allowed Ticket Recommendation

Recommended next actions are restricted to:
- Literature acquisition: Acquiring peer-reviewed publications and text sources to resolve `OperatorFormDebt`.
- Human physics-review preparation: Creating specific checklists and questionnaire templates to prepare for the human review gate.
- Regression testing: Adding more robust assertion coverage in the Go test suite.

Disallowed next steps:
- Immediate implementation of any BMC-0B numerical solver.
- Production of massive scalar numerical solutions.
- Friedmann recovery claims.
- Unblocking full BMC.
