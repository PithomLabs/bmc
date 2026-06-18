```json
{
  "summary": "BMC-REVIEW-0003 adds `docs/reviews/bmc_review_0003_literature_acquisition_source_provenance_intake.md`. The document performs a disciplined source-intake and provenance review: it explicitly distinguishes Case 1 (source text available in repo), Case 2 (bibliographic reference exists but source text missing), and Case 3 (no source/reference available yet); inventories only conservative in-repo sources; lists cited-but-inaccessible references with `source_text_missing` status; defines a comprehensive missing-source acquisition ledger; provides equation-intake and claim-intake templates with strict no-acceptance-without-human-review rules; and keeps all physical debts unpaid because no external peer-reviewed source text is present. It does not implement a solver, produce numerical BMC-0B results, retire any debt by inventory alone, or authorize solver work. Forbidden positive claims are absent; any forbidden-word matches across the reviewed docs are confined to explicitly negated/rejected-example/no-promotion contexts. `go test ./... -count=1` passes and `lake build` passes. Untracked `phase/review0003/` boundary files remain and should be removed or moved to restore clean boundary hygiene.",
  "overall_verdict": "accept_with_repairs",
  "ebp_debt_review": {
    "BMCReview0003Status": "accepted_for_source_intake_scope",
    "SourceProvenanceIntake": "accepted_for_source_intake_scope",
    "OperatorFormDebt": "documented_unpaid",
    "MetricSignatureDebt": "dependent_unpaid",
    "ScalarPotentialNormalizationDebt": "dependent_unpaid",
    "FactorOrderingDebt": "dependent_unpaid",
    "UnitsConventionDebt": "dependent_unpaid",
    "BoundaryConditionDebt": "dependent_unpaid",
    "WavefunctionAdmissibilityDebt": "dependent_unpaid",
    "ResidualDefinitionDebt": "dependent_unpaid",
    "NullModelDesignDebt": "dependent_unpaid",
    "ClassicalRecoveryCriterionDebt": "dependent_unpaid",
    "FaithfulnessReviewDebt": "partial_documentation_only",
    "HumanPhysicsReviewDebt": "unpaid",
    "LeanProofDebt": "unpaid",
    "containsFinalTruthClaim": "absent",
    "BMC0BStatus": "specified_only",
    "SolverStatus": "not_implemented",
    "NoRecoveryClaimBoundary": "clean",
    "NoSchemaCLIBloat": "clean",
    "full_bmc_toy_gate": "blocked"
  },
  "case_distinction_findings": [
    "Section 3 defines the three-case methodology: Case 1 (source text available in repo), Case 2 (bibliographic reference exists but source text missing), Case 3 (no source/reference available yet).",
    "Only Case 1 is treated as capable of supporting future equation intake; Case 2 and Case 3 remain `source_acquisition_required`.",
    "The document explicitly states no case can retire debt without equation mapping and human physics review."
  ],
  "source_inventory_findings": [
    "Section 4 inventories Case 1 sources conservatively: Kimi chat notes, ChatGPT notes, and the residual code fixture.",
    "All are marked `available_for_review` and `Human review required? Yes`.",
    "No in-repo source is treated as peer-reviewed physics authority.",
    "Section 5 inventories Case 2 references (arXiv citations) with `source_text_missing` / `source_acquisition_required` status; no equations are invented or summarized as though read."
  ],
  "repository_source_findings": [
    "In-repo project documents identify current assumptions but are correctly classified as `project_doc` and `code_fixture`, not peer-reviewed literature.",
    "The residual code fixture is recorded as showing current implementation behavior only."
  ],
  "citation_reference_findings": [
    "Nine cited references are listed in Case 2 with ArXiv IDs and URLs but no source text present.",
    "Each is marked `source_text_missing` and `source_acquisition_required`.",
    "No summary, paraphrase, or invented equation content is attributed to these inaccessible sources."
  ],
  "missing_source_acquisition_findings": [
    "Section 6 provides a missing source acquisition ledger covering all required categories: minisuperspace WdW operator, massive scalar potential normalization, factor ordering, metric/signature, units/constants, boundary conditions, Bohmian guidance/Q-potential, classical recovery target, and null-model design.",
    "Each entry includes why needed, debt affected, minimum required content, acceptable source form, human reviewer needed, priority, and case classification.",
    "No missing category is treated as already satisfied."
  ],
  "source_relevance_criteria_findings": [
    "Section 7 defines eight explicit relevance criteria for BMC-0B operator-form sources.",
    "Sources not meeting these criteria are marked background-only and cannot be used for debt retirement.",
    "This prevents background sources from being misused as evidence."
  ],
  "equation_intake_template_findings": [
    "Section 8 provides a reusable equation-intake template with all required fields, including exact source location, equation text, variable mapping, metric/signature, factor ordering, potential term, units/constants, boundary assumptions, supported/unsupported claims, debt affected, human review question, and status.",
    "Intake rules explicitly prohibit acceptance without human physics review, inferring missing notation, silently normalizing signs, merging conventions without a mapping table, and treating equation similarity as debt retirement."
  ],
  "claim_intake_template_findings": [
    "Section 9 provides a reusable claim-intake template covering operator form, factor ordering, units convention, boundary condition, guidance law, quantum potential, classical recovery, null model, and interpretive claims.",
    "Each entry requires evidence requirements and human review questions."
  ],
  "candidate_source_category_findings": [
    "Section 10 defines three candidate source categories: primary peer-reviewed literature, secondary review articles, and project internal notes.",
    "Only primary peer-reviewed literature is permitted to support retirement of operator-form and physical debts.",
    "Project internal notes are restricted to diagnostic context."
  ],
  "operator_form_source_requirement_findings": [
    "Section 11 states the minimum source content required to revisit OperatorFormDebt: explicit WdW equation, configuration variables, metric/signature, kinetic term, potential term, factor ordering or ambiguity, units/constants, boundary/domain assumptions or explicit omission, and notation mapping.",
    "Operator-form review cannot proceed from partial or implicit source material."
  ],
  "massive_scalar_potential_source_requirement_findings": [
    "Section 12 requires source material to identify potential form, mass parameter convention, normalization constants, unit system, placement in Hamiltonian constraint, and placement in WdW operator.",
    "Scalar potential normalization cannot be inferred without source text."
  ],
  "factor_ordering_source_requirement_findings": [
    "Section 13 requires source material or human review for chosen ordering, alternative orderings, ordering ambiguity, and effects on kinetic operator, residuals, boundary behavior, and Bohmian guidance/Q-potential.",
    "Factor ordering cannot be silently assumed."
  ],
  "units_constants_source_requirement_findings": [
    "Section 14 preserves source requirements for hbar, G, c, Planck/dimensionless choices, mass parameter convention, potential normalization, metric/signature, and WdW sign convention.",
    "Units cannot be fixed without source provenance."
  ],
  "boundary_domain_source_requirement_findings": [
    "Section 15 requires sources or human review for domain of alpha, domain of phi, grid/domain bounds, boundary conditions, boundary residual treatment, wavefunction admissibility, node/near-node handling, nonfinite behavior, and stencil-boundary behavior.",
    "Boundary conditions are not treated as settled."
  ],
  "classical_recovery_source_requirement_findings": [
    "Section 16 requires classical recovery source material to define target classical equations, variable mapping, clock/time convention, observable or residual to compare, error metric, domain restrictions, failure conditions, and null-model comparison need.",
    "Classical recovery criteria are not treated as already defined."
  ],
  "notation_mapping_requirement_findings": [
    "Section 17 requires future equation intake to map source variables to alpha and phi, source scale factor to project alpha = ln(a), source scalar field to project phi, source time/lapse/clock to project lambda, source wavefunction to project Psi, source operator to project residual conventions, and source constants/units to the project convention table.",
    "Notation mapping cannot be skipped."
  ],
  "human_literature_review_findings": [
    "Section 18 includes a human review checklist asking which source is authoritative, which equation to map first, source variables and conventions, whether the source includes massive scalar potential normalization and factor ordering, metric/signature, boundary/domain assumptions, classical recovery justification, Bohmian guidance/Q-potential alignment, and what debts can be retired with what evidence.",
    "Human review is not optional or decorative."
  ],
  "debt_status_findings": [
    "Section 19 classifies all required debts. OperatorFormDebt is `documented_unpaid` due to missing peer-reviewed source texts in-repo.",
    "All dependent debts are `dependent_unpaid` and blocked by OperatorFormDebt.",
    "FaithfulnessReviewDebt is `partial_documentation_only`; HumanPhysicsReviewDebt and LeanProofDebt are `unpaid`.",
    "No debt is retired by inventory alone."
  ],
  "no_promotion_findings": [
    "Section 1 non-goals include all required no-promotion statements and the maximum stage.",
    "Section 20 repeats the no-promotion audit verbatim.",
    "No promotion-adjacent language appears in positive claims or statuses."
  ],
  "forbidden_phrase_findings": [
    "Forbidden or promotion-adjacent terms appear only in explicitly negated/rejected-example/no-promotion contexts across the reviewed documentation surface.",
    "No accepted status, checklist result, source status, or positive claim uses these terms."
  ],
  "documentation_only_findings": [
    "Review scope is documentation-only.",
    "No solver code, trajectory integration, BMC-0B numerical computation, CLI route, report schema, generated output artifact, or success dashboard was added.",
    "Untracked boundary directories `phase/review0003/` remain outside the review target and should be cleaned up."
  ],
  "test_and_build_findings": [
    "go test ./... -count=1: PASS",
    "lake build: PASS (14 jobs)"
  ],
  "missing_sections": [],
  "required_repairs_before_acceptance": [
    "Remove the untracked boundary files in `phase/review0003/`, or move them under documentation so the final acceptance boundary matches the documentation-only target."
  ],
  "optional_repairs": [],
  "promotion_recommendation": "accept_review_0003_for_source_intake_scope",
  "next_smallest_useful_move": "Acquire full text for the Case 2 bibliographic references, prioritizing the minisuperspace WdW operator source and metric/signature source, then run the equation-intake template through human physics review to begin retiring OperatorFormDebt and dependent debts."
}
```