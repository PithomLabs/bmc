```json
{
  "summary": "BMC-REVIEW-0004 adds `docs/reviews/bmc_review_0004_source_text_acquisition_equation_intake.md`. The document performs a disciplined Case 1 equation-intake review: it updates source acquisition status, clearly separates `project_note_equation_candidate_unreviewed`, `code_fixture_equation_candidate_unreviewed`, and `peer_reviewed_source_equation_candidate_unreviewed` categories; inventories project-note equations from Kimi/ChatGPT sources and the residual code fixture with intact debt markers; preserves Case 2/3 as `source_acquisition_required`/`equation_intake_blocked`; provides notation and convention mapping ledgers; compares candidate equations against the current fixture without claiming equivalence; and keeps all physical debts unpaid unless source text, notation mapping, convention mapping, and human physics review justify retirement. No solver code, numerical BMC-0B computation, CLI/schema/output artifact, or physics promotion was introduced. Forbidden positive claims are absent; any forbidden-word matches across the docs remain confined to explicitly negated/rejected-example/no-promotion contexts. `go test ./... -count=1` and `lake build` pass. Untracked `phase/review0004/` boundary files remain and should be removed or moved to restore clean boundary hygiene.",
  "overall_verdict": "accept_with_repairs",
  "ebp_debt_review": {
    "BMCReview0004Status": "accepted_for_equation_intake_scope",
    "SourceTextAcquisition": "accepted_for_equation_intake_scope",
    "EquationIntake": "accepted_for_equation_intake_scope",
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
  "source_acquisition_status_findings": [
    "Section 3 updates source acquisition status for all prior Case 1/2/3 entries, restoring them from REVIEW-0003.",
    "Case 1 project notes and code fixtures are marked `source_text_available` with `source_equation_candidate_unreviewed` status and `Human review required? Yes`.",
    "Case 2 bibliographic references remain `source_text_missing`, `source_acquisition_required`, and `equation_intake_blocked`.",
    "No status implies proof, validation, recovery, solver permission, debt retirement, or physics success."
  ],
  "case_resolution_findings": [
    "Section 4 explicitly states no Case 2 or Case 3 source text has been acquired and these items cannot be resolved yet.",
    "Case 1 intake is permitted only as candidate intake; Case 2/3 remain blocked from equation intake and debt retirement."
  ],
  "source_text_inventory_findings": [
    "Section 5 inventories only Kimi/ChatGPT project notes and the residual fixture as newly available Case 1 source texts.",
    "No new peer-reviewed external publication is claimed as available."
  ],
  "missing_source_findings": [
    "Section 6 lists all Case 2/3 sources as still missing: minisuperspace WdW operator, metric/signature, scalar potential normalization, factor ordering, units/constants, boundary conditions, Bohmian guidance/Q-potential, classical recovery target, null-model/control-model sources.",
    "No missing category is treated as satisfied."
  ],
  "equation_search_methodology_findings": [
    "Section 7 records sources inspected, keywords searched, evaluated sections/pages, equations located, sufficiency judgment, and search limits.",
    "Context is judged sufficient only to extract toy candidates, not to establish a peer-reviewed physical standard.",
    "No claim of completeness beyond the inspected internal files is made."
  ],
  "equation_intake_ledger_findings": [
    "Section 8 provides the equation-intake ledger with all required fields for EQ-KIMI-0001 through EQ-KIMI-0005 and EQ-CODE-0001.",
    "All entries are marked `project_note_equation_candidate_unreviewed` or `code_fixture_equation_candidate_unreviewed`.",
    "No peer-reviewed equation status is assigned.",
    "Intake rules prohibit acceptance without human physics review, inferring missing notation, silently normalizing signs, merging conventions without a mapping table, and treating equation similarity as debt retirement."
  ],
  "equation_category_distinction_findings": [
    "The ledger correctly separates `project_note_equation_candidate_unreviewed`, `code_fixture_equation_candidate_unreviewed`, and `peer_reviewed_source_equation_candidate_unreviewed`.",
    "The document explicitly states only the peer-reviewed category can later become eligible for debt-retirement review, and only after notation mapping, convention mapping, and human physics review.",
    "No peer-reviewed equation status is assigned without peer-reviewed source text."
  ],
  "notation_mapping_findings": [
    "Section 9 maps source symbols (`a`, `φ`, `Ψ`, `t`) to project symbols with `mapping_uncertain` confidence and `requires_human_physics_review` status.",
    "Maps are not treated as settled."
  ],
  "convention_mapping_findings": [
    "Section 10 compares conventions across metric/signature, WdW sign convention, factor ordering, units/constants, mass parameter convention, scalar potential normalization, boundary/domain assumptions, wavefunction admissibility, inner product/diagnostic measure, and residual norm.",
    "Conflicts and ambiguities are explicitly identified and tied to specific debts.",
    "Convention conflicts are not smoothed over or silently normalized."
  ],
  "operator_form_candidate_findings": [
    "Section 11 records the candidate operator-form intake from project notes and code fixture only.",
    "The candidate is not treated as a faithful/source-backed operator form; debt remains `documented_unpaid` and `Human review required: Yes`."
  ],
  "metric_signature_intake_findings": [
    "Section 12 records the metric signature from project notes as unreviewed and notes the project's log-coordinate convention differs from the source's direct `a` coordinate.",
    "Status is implicitly dependent-unpaid."
  ],
  "scalar_potential_intake_findings": [
    "Section 13 records the potential form and placement from project notes, noting absence in the current code fixture.",
    "`ScalarPotentialNormalizationDebt` remains unpaid and is not accepted from project notes alone."
  ],
  "factor_ordering_intake_findings": [
    "Section 14 records the project-note factor ordering and notes impacts on residuals and boundary behavior.",
    "`FactorOrderingDebt` remains unpaid and solver design is not permitted to assume it retired."
  ],
  "units_constants_intake_findings": [
    "Section 15 records unit/constant assumptions from project notes and notes differences from the current fixture.",
    "`UnitsConventionDebt` remains unpaid."
  ],
  "boundary_domain_intake_findings": [
    "Section 16 records boundary/domain assumptions from project notes and conflicts with current stencil-boundary blocking.",
    "`BoundaryConditionDebt` remains unpaid."
  ],
  "bohmian_guidance_qpotential_intake_findings": [
    "Section 17 records guidance equations and Q-potential formula from project notes and notes differences from current audit machinery.",
    "Project audit formulas are not treated as physically faithful without source review."
  ],
  "classical_recovery_intake_findings": [
    "Section 18 records classical recovery target equations from project notes and notes remaining questions.",
    "Classical recovery criteria are not treated as already defined."
  ],
  "source_to_project_fixture_comparison_findings": [
    "Section 19 compares project-note equations against `-d2Psi/dAlpha2 + d2Psi/dPhi2` across variables, metric/signature, kinetic signs, potential term, factor ordering, units/constants, boundary assumptions, and limiting-case interpretability.",
    "Equivalence is not claimed; missing evidence is explicitly listed."
  ],
  "debt_status_findings": [
    "Section 20 classifies all required debts as `documented_unpaid` or `dependent_unpaid`.",
    "FaithfulnessReviewDebt is `partial_documentation_only`; HumanPhysicsReviewDebt and LeanProofDebt are `unpaid`.",
    "The section additionally records `candidate_equation_intake_complete` and `debt_retirement_blocked_pending_human_review` to indicate intake progress without debt retirement.",
    "No debt is retired by project-note or code-fixture equation intake."
  ],
  "human_physics_review_findings": [
    "Section 21 includes the required human physics-review checklist questions.",
    "Human review is not optional or decorative."
  ],
  "no_promotion_findings": [
    "Section 1 non-goals include the six required no-promotion statements and the maximum stage.",
    "Section 22 repeats the no-promotion audit verbatim.",
    "No promotion-adjacent language appears in positive claims or statuses."
  ],
  "forbidden_phrase_findings": [
    "Forbidden or promotion-adjacent terms appear only in explicitly negated/rejected-example/no-promotion contexts across the reviewed documentation surface.",
    "No accepted status, checklist result, source status, equation status, or positive claim uses these terms."
  ],
  "documentation_only_findings": [
    "Review scope is documentation-only.",
    "No solver code, trajectory integration, BMC-0B numerical computation, CLI route, report schema, generated output artifact, or success dashboard was added.",
    "Untracked boundary directory `phase/review0004/` remains outside the review target and should be cleaned up."
  ],
  "test_and_build_findings": [
    "go test ./... -count=1: PASS",
    "lake build: PASS (14 jobs)"
  ],
  "missing_sections": [],
  "required_repairs_before_acceptance": [
    "Remove the untracked boundary files in `phase/review0004/`, or move them under documentation so the final acceptance boundary matches the documentation-only target."
  ],
  "optional_repairs": [],
  "promotion_recommendation": "accept_review_0004_for_equation_intake_scope",
  "next_smallest_useful_move": "Acquire full text for the Case 2 bibliographic references and execute the human physics-review checklist on the Case 1 project-note equation candidates (`EQ-KIMI-0001`-`EQ-KIMI-0005`, `EQ-CODE-0001`) before attempting any further operator-form debt classification."
}
```