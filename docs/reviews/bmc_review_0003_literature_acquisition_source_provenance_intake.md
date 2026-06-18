# BMC-REVIEW-0003: Literature Acquisition and Source-Provenance Intake

This document records the literature acquisition and source-provenance intake for the BMC-0B Wheeler-DeWitt (WdW) constraint operator under strict EBP 2.1 discipline. It inventories available and missing sources, defines intake templates, sets relevance criteria, and outlines requirements for notation mapping and human physics review.

---

## 1. Review Purpose and Non-Goals

### Purpose
The purpose of this review is to create a documentation-only source intake and provenance ledger for literature needed to review the BMC-0B minisuperspace massive-scalar Wheeler-DeWitt operator form, answering the following:
- What source materials are needed to review the BMC-0B operator form?
- Which sources, if any, are currently present in the repository?
- Which equations or claims are actually visible in available sources?
- Which sources are still missing?
- Which source passages must be reviewed by a human physicist before debt retirement?
- What exact follow-up review could map source equations to project conventions?

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

Maximum allowed EBP stage: `BMC-REVIEW-0003 accepted_for_source_intake_scope`

---

## 2. Current Accepted Boundary from REVIEW-0001 and REVIEW-0002

Any future review or design must respect the boundaries set by the first two review tickets:
- **REVIEW-0001**: Created the convention/provenance ledger, mapping project assumptions and marking the current fixture operator as project-only.
- **REVIEW-0002**: Confirmed that no external peer-reviewed literature equation is currently available in-repo to retire `OperatorFormDebt`.

This ticket (REVIEW-0003) focuses strictly on the acquisition and registration of source materials, keeping all physical debts unpaid.

---

## 3. Source-Intake Methodology

To prevent the ingestion of unverified equations or silent convention switching, the following methodology is applied to all sources:
1. **Classification**: Categorize the source into Case 1 (source text available in-repo), Case 2 (bibliographic reference exists but source text missing), or Case 3 (no source/reference available yet).
2. **Relevance Assessment**: Evaluate the source against explicit relevance criteria (Section 7). Only sources meeting these criteria can be used as evidence.
3. **Equation and Claim Registration**: Extract candidates using the templates defined in Sections 8 and 9.
4. **Notation and Convention Mapping**: Map all source variables to project conventions. Do not silently normalize signs or factors.
5. **Human Physics Review**: Submit the completed intake files for human review before any debt retirement is proposed.

---

## 4. Repository Source Inventory (Case 1: Source Text Available in Repo)

Only sources listed in this section have full text available in the repository and are capable of supporting future equation intake.

| Source ID | Source title or description | Location in repo | Source type | Text available? | Relevant equation visible? | Relevant to which debt? | Status | Human review required? |
|---|---|---|---|---|---|---|---|---|
| `source_project_kimi_notes` | Kimi chat notes / transcripts on Bohmian cosmology | [bqg_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg_kimi.md), [bqg01_kimi.md](file:///home/chaschel/Documents/go/bmc/bqg01_kimi.md) | `project_doc` | Yes | Yes (toy minisuperspace equations) | `OperatorFormDebt` | `available_for_review` | Yes |
| `source_project_chatgpt_notes` | ChatGPT chat notes on Bohmian pilot-wave cosmology | [bqg01_chatgpt.md](file:///home/chaschel/Documents/go/bmc/bqg01_chatgpt.md) | `project_doc` | Yes | Yes (toy minisuperspace equations) | `OperatorFormDebt` | `available_for_review` | Yes |
| `source_project_residual_code` | Residual code fixture evaluator | [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) | `code_fixture` | Yes | Yes (toy finite-difference residual) | `OperatorFormDebt` | `available_for_review` | Yes |

---

## 5. Existing Citation/Reference Inventory (Case 2: Bibliographic Reference Exists but Source Text Missing)

The following papers are cited in project documentation, but their full text is not present in the repository. They are currently classified as missing and cannot support equation intake until acquired.

| Source ID | Citation details / Title | URL or Arxiv ID | Source type | Text available? | Status | Debt affected |
|---|---|---|---|---|---|---|
| `source_pinto_neto_fabris_2013` | "Quantum cosmology from the de Broglie-Bohm perspective" | [arXiv:1306.0820](https://arxiv.org/abs/1306.0820) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `OperatorFormDebt`, `FactorOrderingDebt` |
| `source_pinto_neto_struyve_2018` | "Bohmian quantum gravity and cosmology" | [arXiv:1801.03353](https://arxiv.org/abs/1801.03353) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `OperatorFormDebt`, `UnitsConventionDebt` |
| `source_colistete_fabris_pinto_neto_1997` | "Singularities and Classical Limit in Quantum Cosmology" | [arXiv:gr-qc/9711047](https://arxiv.org/abs/gr-qc/9711047) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `OperatorFormDebt`, `ClassicalRecoveryCriterionDebt` |
| `source_analytical_solutions_2023` | "Just some simple (but nontrivial) analytical solutions for de Broglie-Bohm quantum cosmology" | [arXiv:2301.06088](https://arxiv.org/abs/2301.06088) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `OperatorFormDebt`, `BoundaryConditionDebt` |
| `source_bohmian_cosmology_2025` | "Bohmian Quantum Cosmology from the Wheeler-DeWitt Equation" | [arXiv:2512.18818](https://arxiv.org/abs/2512.18818) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `OperatorFormDebt`, `MetricSignatureDebt` |
| `source_jacobson_1995` | "Thermodynamics of Spacetime: The Einstein Equation of State" | [arXiv:gr-qc/9504004](https://arxiv.org/abs/gr-qc/9504004) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `ClassicalRecoveryCriterionDebt` |
| `source_lqg_continuum_2014` | "The continuum limit of loop quantum gravity" | [arXiv:1409.1450](https://arxiv.org/abs/1409.1450) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `NullModelDesignDebt` |
| `source_sft_review_2024` | "String Field Theory: A Review" | [arXiv:2405.19421](https://arxiv.org/abs/2405.19421) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `NullModelDesignDebt` |
| `source_asymptotic_safety_2025` | "Asymptotic Safety and Canonical Quantum Gravity" | [arXiv:2507.14296](https://arxiv.org/abs/2507.14296) | `bibliographic_reference_only` | No | `source_text_missing` / `source_acquisition_required` | `NullModelDesignDebt` |

---

## 6. Missing Source Acquisition Ledger (Case 3: No Source/Reference Available Yet)

The following table lists critical categories of missing sources where no specific bibliographic reference or source text is currently available in the repository, or where references exist but are missing full text.

| Needed source category | Why needed | Debt affected | Minimum required content | Acceptable source form | Human reviewer needed | Priority | Case |
|---|---|---|---|---|---|---|---|
| minisuperspace WdW operator with scalar field | Define the continuous kinetic terms and metric factors | `OperatorFormDebt` | Explicit differential equation for scale factor and scalar | Peer-reviewed publication text | Yes | Critical | Case 2 (`source_pinto_neto_fabris_2013`) |
| massive scalar potential normalization | Standardize how the potential $V(\phi)$ enters the constraint | `ScalarPotentialNormalizationDebt` | Formula showing coupling constant factors | Peer-reviewed publication text | Yes | High | Case 2 (`source_bohmian_cosmology_2025`) |
| factor-ordering discussion | Address mathematical ambiguities in derivative orderings | `FactorOrderingDebt` | Equations comparing alternative orderings ($p$-parameter) | Peer-reviewed publication text | Yes | High | Case 3 (no reference in docs yet) |
| minisuperspace metric/signature convention | Fix the signature of configuration space metrics | `MetricSignatureDebt` | explicit metric tensor definition for $(a, \phi)$ | Peer-reviewed publication text | Yes | Critical | Case 2 (`source_bohmian_cosmology_2025`) |
| units/constants convention | Normalize physical constants ($\hbar$, $G$, $c$) and units | `UnitsConventionDebt` | List of coordinate units and constants | Peer-reviewed publication text | Yes | High | Case 2 (`source_pinto_neto_struyve_2018`) |
| boundary-condition discussion | Specify mathematically admissible boundary conditions | `BoundaryConditionDebt` | Equations for boundary behavior at $a \to 0$ | Peer-reviewed publication text | Yes | Medium | Case 3 (no reference in docs yet) |
| Bohmian quantum cosmology guidance/Q-potential formulation | Derive trajectory equations from phase gradients | `FaithfulnessReviewDebt` | Formula for pilot-wave velocity and $Q$ potential | Peer-reviewed publication text | Yes | High | Case 2 (`source_analytical_solutions_2023`) |
| classical recovery target equations | Map wave trajectories to classical Friedmann solutions | `ClassicalRecoveryCriterionDebt` | Friedmann equation with corrections | Peer-reviewed publication text | Yes | High | Case 2 (`source_colistete_fabris_pinto_neto_1997`) |
| null-model or control-model design reference | Benchmark against rival quantum gravity frameworks | `NullModelDesignDebt` | Equations for alternative relational clocks | Peer-reviewed publication text | Yes | Medium | Case 2 (`source_jacobson_1995`) |

---

## 7. Source Relevance Criteria

To prevent irrelevant or background material from being used to retire physical debts, a source is relevant to the BMC-0B operator form only if it satisfies at least one of the following criteria:
1. Contains an explicit minisuperspace Wheeler-DeWitt equation.
2. Uses scalar-field cosmology or contains a massive scalar potential term.
3. Explicitly states the configuration space metric/signature or kinetic term convention.
4. States factor ordering or discusses the ordering ambiguity in quantum cosmology.
5. Explicitly states units/constants conventions.
6. Explicitly states boundary/domain assumptions for the wavefunction.
7. Connects the wavefunction to Bohmian quantum cosmology guidance equations or Q-potential formulations.
8. States classical target equations or recovery criteria.

Sources that do not meet these criteria are marked as **background only** and cannot be used as evidence for debt retirement.

---

## 8. Equation-Intake Template

The template below must be used for any future extraction of equations from acquired sources.

```text
Equation ID: [e.g., EQ-SOURCE-0001]
Source ID: [e.g., source_pinto_neto_fabris_2013]
Exact source location: [e.g., Section II, Equation 15, Page 4]
Equation text or short excerpt: [LaTex formula]
Variables used by source: [list variables and source definitions]
Project notation mapping: [mapping to project variables alpha, phi, etc.]
Metric/signature convention: [e.g., (-, +)]
Factor ordering: [specify ordering parameters used in the source]
Potential term: [definition and normalization of V(phi)]
Units/constants: [e.g., hbar=1, G=1, c=1, or explicit dimensions]
Boundary assumptions: [e.g., vanishing at singularity]
What this equation supports: [details]
What this equation does not support: [details]
Debt affected: [e.g., OperatorFormDebt]
Human review question: [question for the human reviewer]
Status: source_equation_candidate_unreviewed
```

### Intake Rules:
- Do not mark any equation as accepted without human physics review.
- Do not infer missing notation.
- Do not normalize signs silently.
- Do not merge multiple source conventions without a mapping table.
- Do not treat equation similarity as debt retirement.

---

## 9. Claim-Intake Template

The template below must be used for any future extraction of qualitative claims or interpretations.

```text
Claim ID: [e.g., CLAIM-SOURCE-0001]
Source ID: [e.g., source_pinto_neto_struyve_2018]
Claim text or summary: [description of claim]
Claim type: [operator_form | factor_ordering | units_convention | boundary_condition | guidance_law | quantum_potential | classical_recovery | null_model | interpretive_claim]
Mathematical dependency: [dependent variables or equations]
Physical dependency: [physical assumptions]
Relation to BMC-0B: [how it maps to BMC-0B]
Relation to current project fixture: [comparison to residual.go]
Evidence required: [what is needed to verify the claim]
Debt affected: [debt name]
Status: unreviewed
Human review question: [question for the human reviewer]
```

---

## 10. Candidate Source Categories

1. **Primary Peer-Reviewed Literature**: Journal publications containing original derivations. Only this category can support the retirement of operator-form and physical debts.
2. **Secondary Review Articles**: Book chapters, review papers, or lecture notes summarizing derivations. Used to verify physical consensus.
3. **Project Internal Notes**: Chat transcripts, developer notes, or text files in the repository. These provide diagnostic context only and cannot retire physical debts.

---

## 11. Minisuperspace WdW Operator Source Requirements

Any acquired source for the WdW operator kinetic term must explicitly state:
- The coordinates utilized (e.g. scale factor $a$ vs. log scale factor $\alpha = \ln(a)$).
- The metric tensor on the minisuperspace configuration space.
- The kinetic operator derivation (e.g., Laplace-Beltrami operator form).

---

## 12. Massive Scalar Potential Source Requirements

Any acquired source for potential normalization must state:
- The coupling factor of the potential $V(\phi)$ in the Hamiltonian constraint.
- The normalization of the mass parameter ($m$) relative to Planck scales.
- The scaling factor of the potential term as a function of the scale factor $a$ (e.g., whether it is scaled by $a^6$ or $e^{6\alpha}$).

---

## 13. Factor-Ordering Source Requirements

Any acquired source for factor ordering must specify:
- The mathematical form of the ordering (e.g., covariant ordering vs. Laplacian ordering).
- The ordering parameter (e.g. $p$ parameter in the Laplace-Beltrami generalization).
- The physical justification for the chosen ordering (e.g., covariance under coordinate transformations).

---

## 14. Units/Constants/Convention Source Requirements

Any acquired source for unit conventions must state:
- The units of the coordinates $\alpha$ and $\phi$.
- The exact scaling of physical constants ($\hbar$, $G$, $c$).
- The normalization of the wavefunction $\Psi$ (e.g. square-integrability measure).

---

## 15. Boundary-Condition Source Requirements

Any acquired source for domain boundaries must state:
- The coordinate range of the scale factor $a$ (whether $a > 0$ or $a \ge 0$).
- The behavior of the wavefunction $\Psi$ at the singularity ($a \to 0$ or $\alpha \to -\infty$).
- The boundary conditions applied at spatial infinity.

---

## 16. Classical-Recovery Target Source Requirements

Any acquired source for classical recovery must state:
- The target classical equations (e.g. Friedmann equations with scalar field).
- The clock parameter used to derive trajectories (proper time vs. conformal time).
- The mathematical mapping between Pilot-wave trajectories and classical spacetime metrics.

---

## 17. Source-to-Project Notation Mapping Requirements

Any future review utilizing an acquired source must construct a translation table between the source variables and the project variables (e.g., mapping $\chi$ in literature to $\phi$ in the code). The table must assign a mapping confidence score.

---

## 18. Human Literature-Review Checklist

The human physics review panel must address the following questions:
1. Which source is authoritative for the BMC-0B operator form?
2. Which equation should be mapped first?
3. What are the source's variables and conventions?
4. Does the source include massive scalar potential normalization?
5. Does the source state factor ordering?
6. Does the source state minisuperspace metric/signature?
7. Does the source define boundary/domain assumptions?
8. Does the source justify any classical recovery target?
9. Does the source align with Bohmian guidance/Q-potential usage?
10. What debts can be retired, and what evidence is required?

---

## 19. Debt Status After Source Intake

Because no external peer-reviewed literature source text is available in-repo, all physical debts remain unpaid.

| Debt | Status | Detail |
|---|---|---|
| `OperatorFormDebt` | `documented_unpaid` | Missing peer-reviewed source texts in-repo |
| `MetricSignatureDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `ScalarPotentialNormalizationDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `FactorOrderingDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `UnitsConventionDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `BoundaryConditionDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `WavefunctionAdmissibilityDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `ResidualDefinitionDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `NullModelDesignDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `ClassicalRecoveryCriterionDebt` | `dependent_unpaid` | Blocked by `OperatorFormDebt` |
| `FaithfulnessReviewDebt` | `partial_documentation_only` | Inventory completed but no equations ingested |
| `HumanPhysicsReviewDebt` | `unpaid` | Human review checklist drafted but not executed |
| `LeanProofDebt` | `unpaid` | Mathematical formalization not started |

---

## 20. No-Promotion Audit

- BMC-REVIEW-0003 does not implement BMC-0B.
- BMC-REVIEW-0003 does not solve the Wheeler-DeWitt equation.
- BMC-REVIEW-0003 does not validate BMC.
- BMC-REVIEW-0003 does not recover Friedmann dynamics.
- BMC-REVIEW-0003 does not prove classical-limit recovery.
- BMC-REVIEW-0003 does not defeat null models.
- BMC-REVIEW-0003 does not unblock full BMC.

The maximum allowed EBP stage is:
`BMC-REVIEW-0003 accepted_for_source_intake_scope`

---

## 21. Next Allowed Ticket Recommendation

Recommended next actions are strictly limited to:
- Literature acquisition: Acquiring peer-reviewed publications and text sources to resolve Case 2 and Case 3 source gaps.
- Human review checklist execution: Reviewing Case 1 project notes and control models with human physics experts.
- Regression testing: Hardening Go constraint tests or Lean build components.

Disallowed next steps:
- Implementing any solver code or trajectory solver for BMC-0B.
- Formulating or running numerical residuals for the massive scalar field model.
- Physics promotion or claims of Friedmann dynamics recovery.
