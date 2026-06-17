 {
  "summary": "Sprint 10.1 is not yet faithful. It reads the expected files and produces a computed
  report, but the residual magnitudes are still derived from embedded constants scaled by
  clock_range rather than an auditable residual calculation from local-branch residual inputs, and
  the runner fabricates an extra blocked branch not present in the source clock-readiness
  artifact.",
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
  "SourceProvenanceIntegrity": "overclaimed",
  "LocalOnlyBoundary": "partial",
  "NoRecoveryClaimBoundary": "partial",
  "NoBMCBeatsNullsBoundary": "partial",
  "ConventionDebtVisibility": "partial",
  "LeanPolicyBoundary": "partial"
  },
  "scope_findings": [],
  "source_file_read_findings": [
  "All four expected input artifacts exist and the generated report records them as file_read.",
  "However, the residual calculation path uses clockReport.LocalRelationBranches and
  nullReport.NullDiagnosticsComputed only; friedmannReport and priorArtReport are read but not used
  in the residual metric computation, so the all-source file-backed provenance is stronger than the
  actual computation path supports."
  ],
  "calculation_ledger_findings": [
  "The ledger is present, but it repeats the same formula text for computed and blocked entries and
  does not make the constants 0.208, 0.502, and 0.251 auditable.",
  "ValidateReport only checks a subset of transparency fields; it does not enforce nonempty
  formula_source, nonempty input_provenance, allowed calculation_status values, or num_input_points
  > 0 for computed ledger entries."
  ],
  "hidden_constant_findings": [
  "Blocker: internal/bmc/residualrun/report.go lines 417-419 still compute mean_abs_residual,
  max_abs_residual, and rms_residual from embedded constants divided by clock_range. Changing branch
  samples changes only the point counts, not the residual magnitudes, and no residual series or
  derivative target from the source artifacts is evaluated.",
  "The formula description says the residual is computed from local branch data, but the only
  numeric dependency for the magnitudes is clock_range; samples and validation_passed are listed in
  input_fields without affecting the residual magnitudes."
  ],
  "branch_eligibility_findings": [
  "Blocker: internal/bmc/residualrun/report.go lines 475-507 unconditionally appends branch_1 when
  the source artifact has only one local branch. In the generated input out/
  bmc0a_clock_readiness.json there is only branch_0, so branch_1 is a synthetic blocked branch while
  its source_artifact and ledger input_provenance claim file-backed provenance."
  ],
  "candidate_residual_findings": [
  "The computed diagnostic for branch_0 is internally consistent with eligibility flags.",
  "The blocked diagnostic for branch_1 is not source-faithful because the branch does not exist in
  the file-backed clock-readiness input."
  ],
  "metric_findings": [
  "Metric validation covers NaN/Inf, negative values, missing optional metric keys, finite-count
  bounds, and nil metrics for computed diagnostics.",
  "The metric computation itself remains too fixture-like because the residual magnitudes are hard-
  coded scale constants rather than computed residuals over evaluation points."
  ],
  "blocked_path_findings": [
  "The runner returns a blocked report when required input files cannot be read.",
  "A named regression test specifically exercising missing input files is absent or not present
  under the requested name/equivalent coverage."
  ],
  "comparison_findings": [
  "The comparison record remains diagnostic-only and does not appear to assert superiority.",
  "Because the target residual metric is not faithfully computed, the computed residual/null
  comparison is downstream of the residual integrity failure."
  ],
  "convention_ledger_findings": [],
  "local_boundary_findings": [
  "The generated report keeps local_branch_only=true and global_cosmology_claim=false.",
  "The synthetic branch_1 weakens the local-only provenance boundary because not every reported
  branch corresponds to a local branch in the source artifact."
  ],
  "validator_findings": [
  "ValidateReport does not reject calculation ledger entries with empty formula_source or
  input_provenance.",
  "ValidateReport does not enforce the allowed calculation_status set.",
  "ValidateReport does not require num_input_points > 0 for computed_from_local_branch ledger
  entries.",
  "ValidateReport does not prove that reported local_branch_eligibility records correspond to actual
  branches in the clock-readiness source artifact."
  ],
  "cli_findings": [
  "go test ./... passed.",
  "lake build passed.",
  "The summary reports total Candidate Residual Diagnostics: 2, which obscures that only one
  diagnostic is computed and the other is blocked; this is misleading for the reported CLI summary."
  ],
  "lean_findings": [
  "Lean scope is policy-only, supports computed and blocked witnesses, builds successfully, and
  contains no sorry/admit.",
  "Lean does not repair the Go-side faithfulness issue because it only encodes safety booleans."
  ],
  "overclaim_findings": [
  "The phrase-level anti-overclaim gates are mostly preserved.",
  "The overclaim is provenance/faithfulness rather than physics promotion: the report labels the
  residual as computed_from_bmc0a_local_branch even though the magnitudes are still based on
  embedded constants."
  ],
  "missing_tests": [
  "TestResidualRunBlockedReportWhenInputsMissing",
  "TestResidualRunCalculationLedgerRequiresFormulaTransparency with checks for formula_source,
  input_provenance, allowed status, and computed num_input_points",
  "TestResidualRunMetricsChangeWhenInputBranchDataChanges",
  "A test rejecting reported branches that are not present in the file-backed clock-readiness
  artifact"
  ],
  "required_repairs_before_acceptance": [
  "Replace the 0.208/0.502/0.251 constant-derived residual magnitudes with an auditable calculation
  over actual local-branch residual inputs or mark the report blocked until those inputs exist.",
  "Remove the unconditional synthetic branch_1, or clearly mark it as non-file-backed/source-
  unavailable without claiming it is sourced from bmc0a_clock_readiness.",
  "Make the calculation ledger disclose every numeric formula component and validate formula_source,
  input_provenance, allowed calculation_status, and positive input point counts for computed
  entries.",
  "Add a sensitivity test proving residual magnitudes change when the actual residual input data
  changes, not merely when clock_range changes.",
  "Make the CLI summary distinguish computed diagnostics from blocked diagnostics."
  ],
  "optional_repairs": [
  "Use the friedmann spec input directly in the residual formula or downgrade its role in source
  provenance to source summary/context only.",
  "Add validation tying reported branch IDs back to file-backed source branch records."
  ],
  "faithfulness_verdict": {
  "status": "rejected",
  "reason": "The sprint does not honestly repair the Sprint 10 faithfulness failure because the
  residual magnitudes are still hidden constants scaled by one branch field, and one reported branch
  is synthetic rather than file-backed."
  },
  "promotion_recommendation": "do_not_promote",
  "next_smallest_useful_move": "Implement a real file-backed residual calculation from explicit per-
  point local-branch inputs, delete the fabricated branch fixture, and add a failing sensitivity
  test before regenerating the report."
  }
