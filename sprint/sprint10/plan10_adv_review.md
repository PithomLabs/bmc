 {
  "summary": "Go tests, CLI generate/validate/summarize, and Lean build all run successfully, but
  Sprint 10 is not yet acceptable as an honest candidate residual runner. The generated report uses
  fixed residual constants while labeling them as computed from a local branch, the validator
  permits several invalid over-claiming or internally inconsistent states, the CLI summary miscounts
  eligible branches, and Sprint 10 walkthrough text still contains forbidden/restricted phrase
  material.",
  "overall_verdict": "reject_for_now",
  "ebp_debt_review": {
  "needLiteratureAudit": "partial",
  "needMap": "partial",
  "needInvariant": "partial",
  "needToyCheck": "partial",
  "needNullModel": "partial",
  "needObstruction": "partial",
  "needFaithfulnessReview": "contested",
  "clock_choice_debt": "unpaid",
  "classical_target_debt": "unpaid",
  "unit_convention_debt": "unpaid",
  "sign_convention_debt": "unpaid",
  "normalization_debt": "unpaid",
  "containsFinalTruthClaim": "absent",
  "ResidualRunDiagnosticIntegrity": "overclaimed",
  "ResidualNullComparisonIntegrity": "partial",
  "SourceProvenanceIntegrity": "partial",
  "LocalOnlyBoundary": "partial",
  "NoRecoveryClaimBoundary": "partial",
  "NoBMCBeatsNullsBoundary": "partial",
  "ConventionDebtVisibility": "partial",
  "LeanPolicyBoundary": "contested"
  },
  "scope_findings": [
  "No full cosmology engine, leaderboard, dashboard, paper ingestion system, or general physics
  profile system was found in the Sprint 10 code path.",
  "The implementation remains structurally scoped to candidate local-branch residual diagnostics,
  but the generated default report overstates computation provenance because the residual values are
  fixed constants rather than derived by an evident residual calculation."
  ],
  "branch_eligibility_findings": [
  "The default report has one eligible branch and one blocked branch, but the CLI summary prints the
  total branch count as 'Eligible Local Branches: 2'. This is materially misleading because branch_1
  is explicitly ineligible.",
  "Validation does not enforce the full eligibility predicate. A branch can be marked eligible while
  node_contact_free=false or trajectory_finite=false and still validate.",
  "Validation does not require derivative readiness before a computed residual is accepted."
  ],
  "source_provenance_findings": [
  "The generated report honestly uses source_artifact_summary rather than file_read for the required
  source artifacts.",
  "The validator does not require the expected source artifact IDs exactly once, does not reject
  duplicates, and does not require a path/status basis when provenance is file_read."
  ],
  "convention_ledger_findings": [
  "The default convention ledger contains all six required debts with expected unpaid/contested
  statuses.",
  "Validation does not reject empty convention debt descriptions.",
  "Validation does not reject human_review_required=false for unresolved convention debts.",
  "Validation permits ledger statuses such as partial for convention debts even where the Sprint 10
  default contract expects the five core convention debts to remain unpaid and faithfulness review
  to remain contested."
  ],
  "candidate_residual_findings": [
  "The default residual metrics are hard-coded constants in GenerateDefaultReport, yet
  residual_provenance is computed_from_bmc0a_local_branch. This should either perform an auditable
  residual calculation or mark the diagnostic as a deterministic fixture/source summary with
  corresponding interpretation limits.",
  "Validation allows residual_computed=true with a blocked residual_status on an eligible branch.",
  "Validation allows candidate_residual_computed=false while an individual diagnostic still has
  residual_computed=true, as long as numeric metric pointers are null."
  ],
  "metric_findings": [
  "Validation rejects nonfinite and negative pointer-valued residual magnitudes, including -1
  sentinel magnitudes.",
  "Validation does not reject num_finite_residual_points > num_evaluation_points.",
  "Validation does not reject negative num_evaluation_points or negative
  num_finite_residual_points.",
  "Validation does not reject residual_finite=true when the finite point count is zero.",
  "Validation does not require computed residual diagnostics to provide non-null mean/max/rms metric
  values."
  ],
  "comparison_findings": [
  "The default comparison is diagnostic-only and does not declare a superiority outcome.",
  "Validation only requires at least one referenced target residual to be computed; it does not
  require every target_residual_id to exist and be computed.",
  "Validation permits a computed comparison to include an unknown target residual ID if another
  target ID is valid.",
  "Null model IDs are not tied to an auditable Sprint 9 comparable diagnostic record, so the
  comparison remains source-summary/decorative unless strengthened."
  ],
  "validator_findings": [
  "Strict JSON decoding and trailing-token rejection are implemented in ReadReport.",
  "Hard identity fields and anti-overclaim booleans are mostly enforced.",
  "The validator is not strict enough for branch eligibility, residual status/provenance
  consistency, metric count invariants, convention ledger descriptions/review flags, source artifact
  completeness, and comparison target integrity.",
  "The forbidden phrase list omits at least one forbidden wording from the review prompt: classical
  cosmology recovered.",
  "The generated/walkthrough materials include restricted phrase text in sprint/sprint10
  documentation, including plan10.md and plan10_code.md."
  ],
  "cli_findings": [
  "run-residuals exists, unknown profile fails safely, validate routes the residual schema to the
  residualrun validator, and summarize routes to the residual summary.",
  "The summary line for eligible branches is incorrect because it prints
  len(local_branch_eligibility), not the count of branches with eligible=true and
  eligible_local_branch status."
  ],
  "lean_findings": [
  "lake build succeeds.",
  "BMC/BMC/ResidualRun.lean is policy-only and does not prove a physics result.",
  "No sorry/admit was found in the reviewed ResidualRun Lean file.",
  "The Lean gate hard-requires candidateResidualComputed=true, so the policy model does not
  represent the valid blocked/no-eligible-branch case."
  ],
  "overclaim_findings": [
  "The strongest overclaim is provenance-level: fixed constants are reported as computed local-
  branch residual metrics.",
  "The walkthrough says validation covers all rules, but several required adversarial cases are not
  covered and currently validate successfully.",
  "Sprint 10 documentation still contains restricted phrase material and should be phrase-cleaned or
  clearly isolated from generated/review artifacts."
  ],
  "missing_tests": [
  "TestResidualRunRequiresLocalOnlyBoundary",
  "TestResidualRunRejectsGlobalCosmologyClaim",
  "TestResidualRunRequiresAllSixConventionDebts",
  "TestResidualRunRejectsResidualWhenNoEligibleBranch",
  "TestResidualRunRejectsFiniteCountGreaterThanEvalCount",
  "TestResidualRunRejectsDeterministicFixtureSeparationCandidate",
  "TestResidualRunRejectsComparisonWithUnknownTargetResidual",
  "TestResidualRunRejectsComparisonWithUncomputedTargetResidual",
  "Tests for node_contact_free=false and trajectory_finite=false on an otherwise eligible branch",
  "Tests for derivative readiness not ready on a computed residual branch",
  "Tests for empty convention debt descriptions and human_review_required=false",
  "Tests for negative evaluation/finite counts and residual_finite=true with zero finite points",
  "Tests requiring every computed comparison target to exist and be computed"
  ],
  "required_repairs_before_acceptance": [
  "Replace hard-coded residual constants with an auditable bounded residual computation, or relabel
  them as deterministic fixtures/source summaries and keep interpretation strictly diagnostic-
  only.",
  "Fix CLI summary to count only truly eligible local branches.",
  "Enforce full branch eligibility: eligible=true, eligible_local_branch status,
  node_contact_free=true, trajectory_finite=true, and derivative readiness before any computed
  residual is accepted.",
  "Enforce residual_computed/status/provenance consistency.",
  "Enforce all metric count invariants and require non-null numeric metrics for computed
  diagnostics.",
  "Require every computed comparison target residual ID to exist and reference a computed
  diagnostic.",
  "Reject empty convention descriptions and human_review_required=false for unresolved convention
  debts.",
  "Strengthen source artifact validation for required IDs, duplicates, and file_read support.",
  "Add the missing forbidden phrase variant and clean Sprint 10 walkthrough/docs of restricted
  wording.",
  "Add regression tests for the validator gaps listed above."
  ],
  "optional_repairs": [
  "Represent the no-eligible-branch blocked case in the Lean policy model.",
  "Include explicit source paths only when files are actually read.",
  "Add a compact provenance note explaining whether each branch was derived from Sprint 5 artifacts
  or from a deterministic summary fixture."
  ],
  "faithfulness_verdict": {
  "status": "contested",
  "reason": "The artifact preserves many policy boundaries, but the residual diagnostics are not yet
  faithfully evidenced as computed from eligible local branches."
  },
  "promotion_recommendation": "do_not_promote",
  "next_smallest_useful_move": "Repair validator invariants and either implement an auditable
  residual calculation or downgrade the default residual diagnostics to deterministic fixtures with
  diagnostic-only interpretation."
  }
