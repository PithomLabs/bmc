# BMC-REVIEW-0004: Source Text Acquisition and Equation Intake

This document records the source text acquisition and equation-intake review for the BMC-0B Wheeler-DeWitt (WdW) constraint operator under strict EBP 2.1 discipline. It updates the source acquisition status, maps candidate equations from available repository files, provides notation and convention mappings, and compares candidate operator forms against current project fixtures.

---

## 1. Review Purpose and Non-Goals

### Purpose
The purpose of this review is to construct a documentation-only source-text acquisition and equation-intake review that:
1. Identifies which source texts are currently available for review.
2. Confirms whether any acquired source contains an explicit minisuperspace WdW operator equation.
3. Extracts source equation candidates into an intake ledger.
4. Maps source notation to current PTW/BMC notation only where the source text supports it.
5. Preserves all debts as unpaid unless source equation, notation mapping, convention mapping, and human physics review requirements are satisfied.

### Non-Goals
- This review is strictly documentation-only.
- It does not implement a solver.
- It does not produce numerical BMC-0B results.
- It does not retire any physical debts.
- It does not solve the Wheeler-DeWitt equation.
- It does not validate BMC.
- It does not recover Friedmann dynamics.
- It does not prove classical-limit recovery.
- It does not defeat null models.
- It does not unblock full BMC.

Maximum allowed EBP stage: `BMC-REVIEW-0004 accepted_for_equation_intake_scope`

---

## 2. Current Accepted Boundary from REVIEW-0001 through REVIEW-0003

- **REVIEW-0001**: Established the convention/provenance ledger, mapping project variables and marking the current fixture operator as project-only.
- **REVIEW-0002**: Confirmed that no external peer-reviewed literature equation is currently available in-repo to retire `OperatorFormDebt`.
- **REVIEW-0003**: Established a source-intake ledger, classifying sources into Case 1 (source text available in-repo), Case 2 (bibliographic reference exists but source text missing), and Case 3 (no source/reference available yet).

This review (REVIEW-0004) respects these boundaries by evaluating Case 1 sources for candidate equation intake while keeping all Case 2 and Case 3 sources blocked.

---

## 3. Source Acquisition Status Update

Case 1 project notes and code fixtures are currently available for review. External peer-reviewed publications (Case 2 and Case 3) remain missing.

| Source ID | Prior REVIEW-0003 case | Current source-text status | Location in repo or provided context | Text available? | Equation search performed? | Relevant equation located? | Relevant debts | Current status | Human review required? |
|---|---|---|---|---|---|---|---|---|---|
| `source_project_kimi_notes` | Case 1 | `source_text_available` | [bqg_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg_kimi.md), [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) | Yes | Yes | Yes | `OperatorFormDebt`, `MetricSignatureDebt`, `FactorOrderingDebt` | `source_equation_candidate_unreviewed` | Yes |
| `source_project_chatgpt_notes` | Case 1 | `source_text_available` | [bqg01_chatgpt.md](file:///home/chaschel/Documents/go/bmc/bqg01_chatgpt.md) | Yes | Yes | Yes | `OperatorFormDebt`, `UnitsConventionDebt` | `source_equation_candidate_unreviewed` | Yes |
| `source_project_residual_code` | Case 1 | `source_text_available` | [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) | Yes | Yes | Yes | `OperatorFormDebt` | `source_equation_candidate_unreviewed` | Yes |
| `source_pinto_neto_fabris_2013` | Case 2 | `source_text_missing` | None | No | No | No | `OperatorFormDebt`, `FactorOrderingDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_pinto_neto_struyve_2018` | Case 2 | `source_text_missing` | None | No | No | No | `OperatorFormDebt`, `UnitsConventionDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_colistete_fabris_pinto_neto_1997` | Case 2 | `source_text_missing` | None | No | No | No | `OperatorFormDebt`, `ClassicalRecoveryCriterionDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_analytical_solutions_2023` | Case 2 | `source_text_missing` | None | No | No | No | `OperatorFormDebt`, `BoundaryConditionDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_bohmian_cosmology_2025` | Case 2 | `source_text_missing` | None | No | No | No | `OperatorFormDebt`, `MetricSignatureDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_jacobson_1995` | Case 2 | `source_text_missing` | None | No | No | No | `ClassicalRecoveryCriterionDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_lqg_continuum_2014` | Case 2 | `source_text_missing` | None | No | No | No | `NullModelDesignDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_sft_review_2024` | Case 2 | `source_text_missing` | None | No | No | No | `NullModelDesignDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |
| `source_asymptotic_safety_2025` | Case 2 | `source_text_missing` | None | No | No | No | `NullModelDesignDebt` | `source_acquisition_required` / `equation_intake_blocked` | Yes |

---

## 4. Case 2 and Case 3 Resolution Status

No Case 2 or Case 3 source text has been acquired or provided in the repository. As a result, no Case 2 or Case 3 items can be resolved. Their status remains `source_text_missing` and `source_acquisition_required`. They cannot support equation intake or contribute to debt retirement at this stage.

---

## 5. Newly Available Source-Text Inventory

No new peer-reviewed external publications are available. The only source texts available for review are the previously cataloged Case 1 repository files:
- **Project Notes**: [bqg_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg_kimi.md), [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md), [bqg01_chatgpt.md](file:///home/chaschel/Documents/go/bmc/bqg01_chatgpt.md).
- **Code Fixtures**: [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go).

---

## 6. Sources Still Missing

All Case 2 and Case 3 sources listed in Section 3 remain missing. This includes the full texts for `source_pinto_neto_fabris_2013`, `source_pinto_neto_struyve_2018`, `source_colistete_fabris_pinto_neto_1997`, `source_analytical_solutions_2023`, `source_bohmian_cosmology_2025`, and others.

---

## 7. Equation-Search Methodology

1. **Sources Inspected**: [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md), [bqg01_chatgpt.md](file:///home/chaschel/Documents/go/bmc/bqg01_chatgpt.md), [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go).
2. **Keywords Searched**: `Wheeler-DeWitt`, `WdW`, `Hamiltonian`, `guidance`, `quantum potential`, `residual`, `-d2Psi`.
3. **Sections/Pages**: Evaluated all lines containing LaTex or text mathematical formatting.
4. **Equations Located**: Located the toy WdW equation, guidance equations, quantum potential formula, and emergent metric in project notes. Located the toy finite-difference residual in the Go code.
5. **Context Sufficiency**: Context in project notes is sufficient to extract toy candidates for review but is not sufficient to establish a peer-reviewed continuous physical standard.
6. **Limits of Search**: Restricted to internal workspace text files. No external web-indexing search was performed.

---

## 8. Equation-Intake Ledger

The extracted equations are categorized as `project_note_equation_candidate_unreviewed` or `code_fixture_equation_candidate_unreviewed`. None are peer-reviewed.

| Equation ID | Source ID | Exact source location | Equation text or short excerpt | Variables used by source | Project notation mapping | Metric/signature convention | Factor ordering | Potential term | Units/constants | Boundary assumptions | What this equation supports | What this equation does not support | Debt affected | Human review question | Status |
|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|---|
| `EQ-KIMI-0001` | `source_project_kimi_notes` | `bqg01_kimi.md` line 88 | `[-∂²/∂a² + (1/a²)∂²/∂φ² + a⁴V(φ) - ka²]Ψ(a,φ) = 0` | `a`, `φ`, `V(φ)`, `k`, `Ψ` | `a` -> scale factor, `φ` -> field, `Ψ` -> Psi | `(-, +)` on `(a, φ)` | Laplacian ordering (implicit) | `a⁴V(φ)` | `8πG=1`, `hbar=1`, `c=1` | `Ψ` regular at `a=0` | Toy massive-scalar cosmology | Coordinates `α = ln(a)` | `OperatorFormDebt` | Is $a$ preferred over $\alpha = \ln(a)$ for grid solvers? | `project_note_equation_candidate_unreviewed` |
| `EQ-KIMI-0002` | `source_project_kimi_notes` | `bqg01_kimi.md` lines 98-99 | `da/dt = ∂S/∂a`, `dφ/dt = -(1/a²)∂S/∂φ` | `a`, `φ`, `S`, `t` | `a` -> scale factor, `φ` -> field, `t` -> Bohmian time | Compatible with `EQ-KIMI-0001` | None | None | None | None | Pilot-wave guidance law | Proper time clocks | `FaithfulnessReviewDebt` | Does the coordinate factor $1/a^2$ in the $\phi$ derivative cause singularities at $a \to 0$? | `project_note_equation_candidate_unreviewed` |
| `EQ-KIMI-0003` | `source_project_kimi_notes` | `bqg01_kimi.md` line 103 | `Q = -(1/2|Ψ|)(∂²|Ψ|/∂a² - (1/a²)∂²|Ψ|/∂φ²)` | `a`, `φ`, `|Ψ|` | `a` -> scale factor, `φ` -> field, `|Ψ|` -> amplitude | Compatible with `EQ-KIMI-0001` | Laplacian-like | None | None | None | Quantum potential audit | `α` coordinate stencils | `FaithfulnessReviewDebt` | Does $Q$ remain bounded at domain boundaries? | `project_note_equation_candidate_unreviewed` |
| `EQ-KIMI-0004` | `source_project_kimi_notes` | `bqg01_kimi.md` line 107 | `(1/a da/dt)² = (1/3)[(1/2)(dφ/dt)² + V(φ) + Q]` | `a`, `φ`, `t`, `V(φ)`, `Q` | `a` -> scale factor, `φ` -> field, `t` -> time, `Q` -> Q | Compatible with `EQ-KIMI-0001` | None | `V(φ)` | None | None | Effective Friedmann recovery check | Numerical solver constraints | `ClassicalRecoveryCriterionDebt` | Does this recovery limit require $\hbar \to 0$? | `project_note_equation_candidate_unreviewed` |
| `EQ-KIMI-0005` | `source_project_kimi_notes` | `bqg01_kimi.md` line 228 | `ds² = -[(da/dt)/(∂S/∂a)]²dt² + a(t)²dΩ₃²` | `a`, `S`, `t`, `ds²`, `dΩ₃²` | `a` -> scale factor, `t` -> time | Compatible with `EQ-KIMI-0001` | None | None | None | None | Spacetime metric reconstruction | Inhomogeneous perturbations | `MetricSignatureDebt` | How does this map handle non-monotonic clocks? | `project_note_equation_candidate_unreviewed` |
| `EQ-CODE-0001` | `source_project_residual_code` | `residual.go` line 34 | `(-d2Psi/dAlpha2 + d2Psi/dPhi2)` | `alpha`, `phi`, `h`, `wf` | `alpha` -> log scale factor, `phi` -> field | `(-, +)` on `(alpha, phi)` | Constant coefficients | None | `hbar=1` | None | Diagnostic verification checks | Physical massive-scalar potential | `OperatorFormDebt` | Does this central difference match any literature form? | `code_fixture_equation_candidate_unreviewed` |

Only the third category (`peer_reviewed_source_equation_candidate_unreviewed`) can later become eligible for debt-retirement review, and only after notation mapping, convention mapping, and human physics review. Because no peer-reviewed external publications are available in-repo, no peer-reviewed equations are present in this ledger.

---

## 9. Notation-Mapping Ledger

The following notation mappings are extracted from available project notes. Mappings are not finalized.

| Source symbol | Source meaning | Project symbol | Project meaning | Mapping confidence | Source evidence | Debt affected | Human review question | Status |
|---|---|---|---|---|---|---|---|---|
| `a` | Scale factor | `a` | Scale factor | `mapping_uncertain` | [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) line 42 | `UnitsConventionDebt` | Does the project code utilize `a` or only `alpha = ln(a)`? | `requires_human_physics_review` |
| `φ` | Scalar field | `phi` | Scalar field | `mapping_uncertain` | [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) line 42 | `UnitsConventionDebt` | Are the normalization factors identical? | `requires_human_physics_review` |
| `Ψ(a, φ)` | Wavefunction | `Psi` | Wavefunction | `mapping_uncertain` | [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) line 91 | `WavefunctionAdmissibilityDebt` | Does `Psi` denote the amplitude or the full complex state? | `requires_human_physics_review` |
| `t` | Bohmian time | `lambda` | Clock parameter | `mapping_uncertain` | [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) line 42 | `UnitsConventionDebt` | How does `t` scale relative to the trajectory parameter? | `requires_human_physics_review` |

---

## 10. Convention-Mapping Ledger

Agreement and conflict statuses between candidate project note equations and the current code fixtures.

| Convention category | Source convention | Project convention | Agreement status | Conflict or ambiguity | Debt affected | Required repair | Human review question |
|---|---|---|---|---|---|---|---|
| metric/signature | `(-, +)` on `(a, φ)` | `(-, +)` on `(alpha, phi)` | Conflict | Source uses `a` directly; project uses `alpha = ln(a)` | `MetricSignatureDebt` | Coordinate transformation required | Should we transform to log scale factor? |
| WdW sign convention | `[-∂²/∂a² + (1/a²)∂²/∂φ²...]` | `[-d2/dAlpha2 + d2/dPhi2]` | Conflict | Kinetic coefficients and signs differ | `OperatorFormDebt` | Re-derive kinetic signs under log scale | Do the kinetic signs match the metric? |
| factor ordering | Laplacian ordering | Constant coefficients | Conflict | Project code has no first-derivative terms | `FactorOrderingDebt` | Add first-derivative terms to stencil | Which factor ordering is physical? |
| units/constants | `8πG=1`, `hbar=1`, `c=1` | `hbar=1` | Ambiguity | Constant scaling factors are omitted | `UnitsConventionDebt` | Declare unified unit system | What units are used for residuals? |
| mass parameter convention | Absorbed in field definition | Implicit | Ambiguity | Mass parameter normalization is unspecified | `UnitsConventionDebt` | Specify mass scaling explicitly | How is field mass scaled? |
| scalar potential normalization | `a⁴V(φ)` | None (massless) | Conflict | Massive potential is absent in project fixture | `ScalarPotentialNormalizationDebt` | Implement potential multiplication in code | What is the normalization factor? |
| boundary/domain assumptions | `Ψ` regular at `a=0` | Stencil-boundary blocking | Conflict | Project uses artificial coordinate clamping | `BoundaryConditionDebt` | Define physical boundaries at $a \to 0$ | How is the singularity handled? |
| wavefunction admissibility | Outgoing/WKB limits | Amplitude floor `1e-8` | Conflict | Project code clamps near-node values | `WavefunctionAdmissibilityDebt` | Define mathematical function space | Is the floor physical? |
| inner product or diagnostic measure | Unspecified | L2 norm | Ambiguity | Volume element not defined in project docs | `ResidualDefinitionDebt` | Define configuration space volume | What measure is used for norms? |
| residual norm | Unspecified | L2, Linf, Boundary | Ambiguity | Tolerance criteria are unjustified | `ResidualDefinitionDebt` | Justify tolerance values | What is the pass/fail threshold? |

---

## 11. Operator-Form Candidate Intake

Because no peer-reviewed external publications are available in-repo, the only candidates are project note toy models. They are recorded below.

### Candidate 1: `operator_candidate_source_equation_unreviewed` (Project Note Toy)
- **Source ID**: `source_project_kimi_notes`
- **Equation ID**: `EQ-KIMI-0001`
- **Source equation text**: `[-∂²/∂a² + (1/a²)∂²/∂φ² + a⁴V(φ) - ka²]Ψ = 0`
- **Source variables**: `a`, `φ`, `V(φ)`, `k`, `Ψ`
- **Project mapping**: `a` -> scale factor, `φ` -> field, `Ψ` -> Psi
- **Metric/signature**: `(-, +)` on `(a, φ)`
- **Kinetic term**: `-∂²/∂a² + (1/a²)∂²/∂φ²`
- **Potential term**: `a⁴V(φ)`
- **Factor ordering**: Laplacian ordering (implicit)
- **Units/constants**: `8πG=1`, `hbar=1`, `c=1`
- **Boundary assumptions**: `Ψ` regular at `a=0`
- **Relation to current fixture**: Diverges due to variables ($a$ vs. $\alpha = \ln(a)$) and massive potential term.
- **Debt status**: `documented_unpaid`
- **Human review required**: Yes.

---

## 12. Metric/Signature Intake
Project note `EQ-KIMI-0001` assumes a metric signature of `(-, +)` on coordinates `(a, φ)`. Under this metric, the scale factor kinetic term has a negative sign, reflecting the Lorentzian character of minisuperspace. This convention is unreviewed.

---

## 13. Massive Scalar Potential Intake
Project note `EQ-KIMI-0001` places the massive scalar potential as a multiplicative term `a⁴V(φ)Ψ` where $V(\phi) = \frac{1}{2}m^2\phi^2$. The factor of $a^4$ arises from the Hamiltonian constraint density. This potential term is absent in the current code fixture.

---

## 14. Factor-Ordering Intake
Project note `EQ-KIMI-0001` uses an implicit Laplacian factor ordering, which does not introduce first-derivative terms because the kinetic operator is written simply in terms of second derivatives. Alternative orderings (e.g. including covariant derivatives with respect to the minisuperspace metric) would introduce first-derivative terms like $\frac{1}{a}\frac{\partial}{\partial a}$, altering the discrete residual.

---

## 15. Units/Constants Intake
Project notes assume Planck units where $8\pi G = \hbar = c = 1$. Under this system, all variables and parameters become dimensionless, making numerical verification sensitive to coordinate scaling.

---

## 16. Boundary/Domain Intake
Project notes assume that the wavefunction $\Psi$ is regular at the singularity $a = 0$ (such as the Vilenkin tunneling or Hartle-Hawking no-boundary wavefunctions). In contrast, the current project stencils use artificial boundary clamping at the grid edges.

---

## 17. Bohmian Guidance and Q-Potential Intake
Project note `EQ-KIMI-0002` defines the guidance equations as:
- $\dot{a} = \frac{\partial S}{\partial a}$
- $\dot{\phi} = -\frac{1}{a^2}\frac{\partial S}{\partial \phi}$

The factor of $-\frac{1}{a^2}$ in the field velocity arises from the inverse metric coefficient on minisuperspace. This factor is absent in the project's plane-wave guidance stencils.

---

## 18. Classical-Recovery Target Intake
Project note `EQ-KIMI-0004` relates the Bohmian trajectory to the effective Friedmann equation containing the quantum potential $Q$. The recovery target requires that $Q \to 0$ at late times ($a \to \infty$) to recover classical cosmology.

---

## 19. Source-to-Project Fixture Comparison

A comparison between `EQ-KIMI-0001` and the current project fixture:

- **Current fixture residual**: `-d2Psi/dAlpha2 + d2Psi/dPhi2`

### Comparison Dimensions:
1. **Variables**: Different. Fixture uses `alpha = ln(a)` and `phi`. Kimi notes use `a` and `φ`.
2. **Metric/signature**: Same signature `(-, +)` but different coordinate metrics. Kimi notes metric depends on $a$ (introducing $1/a^2$ factor).
3. **Kinetic signs**: Same. Both feature a negative sign on scale factor derivatives.
4. **Potential term**: Kimi notes include `a⁴V(φ)`. Fixture has no potential.
5. **Factor ordering**: Both assume constant factor coefficients without first-derivative terms.
6. **Units/constants**: Both assume $\hbar = 1$, but Kimi notes explicitly fix $8\pi G = c = 1$.
7. **Boundary assumptions**: Kimi notes assume regularity at $a = 0$. Fixture blocks stencils at boundary grid coordinates.
8. **Fixture as limit**: The current fixture cannot be interpreted as a limiting case of `EQ-KIMI-0001` because the coordinate transformation $\alpha = \ln(a)$ introduces first-derivative terms under Laplacian ordering which are absent in the fixture.
9. **Missing evidence**: Verification requires deriving the WdW operator in $\alpha = \ln(a)$ coordinates under various factor orderings from peer-reviewed literature.

---

## 20. Debt Status After Equation Intake

Because no peer-reviewed external publications are available in the repository, the candidate equations are project-note candidates only. All physical debts remain unpaid.

`candidate_equation_intake_complete`
`debt_retirement_blocked_pending_human_review`

- **OperatorFormDebt**: `documented_unpaid`
- **MetricSignatureDebt**: `dependent_unpaid`
- **ScalarPotentialNormalizationDebt**: `dependent_unpaid`
- **FactorOrderingDebt**: `dependent_unpaid`
- **UnitsConventionDebt**: `dependent_unpaid`
- **BoundaryConditionDebt**: `dependent_unpaid`
- **WavefunctionAdmissibilityDebt**: `dependent_unpaid`
- **ResidualDefinitionDebt**: `dependent_unpaid`
- **NullModelDesignDebt**: `dependent_unpaid`
- **ClassicalRecoveryCriterionDebt**: `dependent_unpaid`
- **FaithfulnessReviewDebt**: `partial_documentation_only`
- **HumanPhysicsReviewDebt**: `unpaid`
- **LeanProofDebt**: `unpaid`

---

## 21. Human Physics-Review Checklist

The human physics review panel must verify:
1. Is the source equation `EQ-KIMI-0001` the correct BMC-0B operator-form source?
2. Are the source variables correctly mapped to alpha and phi?
3. Is the metric/signature explicit or inferred?
4. Is factor ordering specified or ambiguous?
5. Is the massive scalar potential included and normalized?
6. Are units/constants fixed?
7. Are boundary/domain assumptions stated?
8. Can the current project fixture be interpreted as a limiting diagnostic case?
9. What equation should be carried forward?
10. Which debts remain blocked?
11. What evidence is required before any debt retirement?

---

## 22. No-Promotion Audit

- BMC-REVIEW-0004 does not implement BMC-0B.
- BMC-REVIEW-0004 does not solve the Wheeler-DeWitt equation.
- BMC-REVIEW-0004 does not validate BMC.
- BMC-REVIEW-0004 does not recover Friedmann dynamics.
- BMC-REVIEW-0004 does not prove classical-limit recovery.
- BMC-REVIEW-0004 does not defeat null models.
- BMC-REVIEW-0004 does not unblock full BMC.

The maximum allowed EBP stage is:
`BMC-REVIEW-0004 accepted_for_equation_intake_scope`

---

## 23. Next Allowed Ticket Recommendation

Recommended next actions are strictly limited to:
- Literature acquisition: Acquiring peer-reviewed publications and text sources to resolve Case 2 and Case 3 source gaps.
- Human review preparation: Preparing specific checklists for reviewing the Case 1 project-note equation candidates.
- Regression testing: Hardening Go constraint tests or Lean build components.

Disallowed next steps:
- Implementing any solver code or trajectory solver for BMC-0B.
- Formulating or running numerical residuals for the massive scalar field model.
- Physics promotion or claims of Friedmann dynamics recovery.
