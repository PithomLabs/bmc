 {
  "summary": "Implement BMC Sprint 11 as a narrow residual/null comparison audit package and CLI
  route. The artifact will read the accepted Sprint 5/9/10.3 inputs when available, audit whether
  residual/null comparisons are structurally meaningful and stable, emit deterministic candidate-
  only JSON, and keep all recovery, superiority, novelty, and full-BMC promotion claims blocked.",
  "proposed_actions": [
  "Create internal/bmc/residualaudit as a sibling package to residualrun and nullrun, following
  the existing report/read/write/validate/test pattern.",
  "Define a deterministic ResidualAuditReport with schema_version bmc0a-residual-audit-v0.1,
  required EBP debt fields, source_artifacts, comparison_audits, stability_audits, gates,
  warnings, and top-level interpretation_status for summary routing.",
  "Implement RunAuditFromFiles and RunAuditFromInputs. File mode reads clock readiness, nullrun,
  local residual, Friedmann spec, and prior-art boundary reports; each successfully read input
  gets file_read provenance, otherwise the generated audit remains valid but blocked.",
  "Generate comparison audit records from residualrun.ResidualNullComparisons, verifying target
  residual IDs map to computed CandidateResidualDiagnostics, null IDs map to comparable Sprint 9
  null diagnostics or are honestly marked unavailable/source-summary, and metrics_audited is
  nonempty.",
  "Generate stability audit records from residualrun residual_input_points for each computed
  diagnostic, using small deterministic perturbations and preserving candidate-only
  interpretation.",
  "Add ptw-bmc audit-residuals --profile bmc0a-residual-audit --out ... plus validate and
  summarize routing for the new schema.",
  "Add strict JSON decoding with DisallowUnknownFields and trailing-token rejection.",
  "Add phrase-safe, case-insensitive validation for restricted language, with explicit handling
  for required negated warnings so the warnings do not cause false positives."
  ],
  "files_to_add": [
  "internal/bmc/residualaudit/contracts.go",
  "internal/bmc/residualaudit/inputs.go",
  "internal/bmc/residualaudit/audit.go",
  "internal/bmc/residualaudit/stability.go",
  "internal/bmc/residualaudit/null_compare.go",
  "internal/bmc/residualaudit/gates.go",
  "internal/bmc/residualaudit/report.go",
  "internal/bmc/residualaudit/validate.go",
  "internal/bmc/residualaudit/residualaudit_test.go",
  "BMC/BMC/ResidualAudit.lean"
  ],
  "files_to_modify": [
  "cmd/ptw-bmc/main.go: add audit-residuals subcommand, validate routing, summarize routing, and
  usage text.",
  "BMC/BMC.lean: import BMC.ResidualAudit if the Lean file is added."
  ],
  "test_plan": [
  "Add TestResidualAuditReportValidation against the deterministic default/generated report.",
  "Add boundary tests rejecting recovery_claim, scientific_novelty_claim_made,
  bmc_beats_null_models_claim, full_bmc_toy_gate != blocked, local_branch_only=false, and
  global_cosmology_claim=true.",
  "Add source artifact tests for required IDs, duplicate IDs, unknown IDs, file_read without path,
  and file_read without available/read_success status.",
  "Add comparison audit tests rejecting decorative comparison audits, empty metrics_audited, empty
  target_residual_ids for computed audits, empty null_model_ids for computed audits, unknown
  target residual IDs, uncomputed target residual IDs, and unavailable/null references not marked
  blocked or summary-only.",
  "Add enum tests rejecting forbidden audit_status, audit_provenance, interpretation_status,
  perturbation_kind, and stability_status.",
  "Add stability tests rejecting nonfinite numeric values, negative perturbation magnitudes except
  perturbation_kind=none, and missing baseline/perturbed/delta values when
  stability_computed=true.",
  "Add gate tests requiring all Sprint 11 gates exactly once, rejecting non-pass gate status and
  unknown gates.",
  "Add strict JSON tests for unknown fields and trailing JSON tokens.",
  "Add deterministic JSON test by generating the same report twice and comparing bytes.",
  "Add phrase-scan and phrase-safe error tests; validation errors must identify field/location
  without echoing restricted strings.",
  "Add CLI tests covering audit-residuals generate, validate, summarize, and unknown-profile
  failure."
  ],
  "cli_plan": [
  "audit-residuals requires --profile and --out.",
  "Only bmc0a-residual-audit is accepted; any other profile exits nonzero without writing a
  successful report.",
  "Default input paths are out/bmc0a_clock_readiness.json, out/bmc0a_nullrun.json, out/
  bmc0a_local_residual.json, out/bmc0a_friedmann_spec.json, and out/bmc0a_prior_art_boundary.json
  resolved from the workspace root.",
  "validate routes schema bmc0a-residual-audit-v0.1 to residualaudit.ReadReport and
  residualaudit.ValidateReport.",
  "summarize prints the requested Sprint 11 summary fields: schema, scope,
  residual_audit_computed, blocked claim booleans, full_bmc_toy_gate, comparison audit count,
  stability audit count, interpretation_status, and promotion_status."
  ],
  "lean_plan": [
  "Add BMC/BMC/ResidualAudit.lean only as a small policy model.",
  "Encode booleans for toyAnalysisOnly, finalTruthClaim, recoveryClaim, scientificNoveltyClaim,
  bmcBeatsNullModelsClaim, fullBMCBlocked, and faithfulnessContested.",
  "Add theorems only for policy boundaries: forbids recovery claim, forbids null-superiority
  claim, requires full BMC blocked, and audit does not imply residual success.",
  "Do not encode or prove physics validity, residual success, null failure, recovery, or
  promotion.",
  "Import the file from BMC/BMC.lean and verify with lake build."
  ],
  "assumptions": [
  "Sprint 11 is additive and does not change Sprint 10.3 residualrun semantics.",
  "A computed audit requires the local residual report to be file-read and
  residual_audit_computed=true; missing source files produce a valid blocked report with no audit
  records.",
  "The top-level report should include interpretation_status so the CLI summary can print a single
  status value.",
  "Status available and read_success are both accepted as successful file_read source statuses,
  matching the prompt and existing package style.",
  "Stability perturbations are deterministic and local to residual_input_points; no stochastic
  sampling is introduced.",
  "Comparison integrity means structural honesty only, never a physics pass/fail conclusion."
  ],
  "comparison_audit_plan": [
  "For each residualrun residual_null_comparison, emit one ResidualComparisonAudit with audit_id
  audit_<source_comparison_id>.",
  "Set audit_computed=true and audit_status=comparison_audited only when every target_residual_id
  exists, every target diagnostic has residual_computed=true, every requested metric exists and is
  finite, metrics_audited is nonempty, and every null_model_id maps to a nullrun record with
  generated diagnostics.",
  "Set audit_status=comparison_decorative when a comparison record exists but has empty metrics,
  empty targets, empty null IDs, metadata-only separation, or cannot be tied to actual computed
  diagnostics.",
  "Set audit_status=comparison_missing when the residual report is computed but has no residual/
  null comparison record.",
  "Set audit_status=source_unavailable or comparison_blocked when required source reports or
  comparable null diagnostics are unavailable.",
  "Use interpretation_status comparison_integrity_passed only for structurally honest comparisons;
  use comparison_integrity_failed, insufficient_target_null_separation,
  blocked_by_missing_null_inputs, or blocked_by_source_unavailable for failed/blocked cases.",
  "Never infer or print null success/failure or model superiority."
  ],
  "stability_audit_plan": [
  "For each computed CandidateResidualDiagnostic, recompute baseline mean_abs_residual,
  max_abs_residual, and rms_residual from residual_input_points before perturbing.",
  "Generate deterministic stability records for baseline_metric=mean_abs_residual using
  alpha_point_perturbation, phi_point_perturbation, and lambda_spacing_perturbation when at least
  two audited input points are present.",
  "Use perturbation_magnitude=1e-6 for alpha/phi additive perturbations and 1e-6 relative spacing
  perturbation for lambda, applied to the first non-endpoint interval when possible, otherwise the
  first interval.",
  "Also emit branch_subset_resampling when at least four residual input points exist by
  recomputing on every other interval; otherwise emit not_computed with
  blocked_by_missing_inputs.",
  "Compute absolute_delta and relative_delta from baseline and perturbed metric values;
  relative_delta is null only when baseline is zero.",
  "Classify stable_under_small_perturbation when relative_delta <= 1e-6 or absolute_delta <= 1e-
  12, sensitive_to_small_perturbation when finite but larger, unstable_or_ill_conditioned for
  nonfinite recomputation, and blocked statuses for missing inputs."
  ],
  "proof_obligations": [
  "All comparison targets in computed audits must point to computed residual diagnostics.",
  "All computed audit null references must point to actual generated Sprint 9 null diagnostics or
  be blocked/source-summary rather than presented as computed.",
  "All computed comparison audits must contain nonempty metrics_audited.",
  "All computed stability audits must have finite baseline, perturbed, absolute_delta, and valid
  relative_delta semantics.",
  "Blocked source-input paths must remain valid JSON reports with residual_audit_computed=false
  and no audit records.",
  "All EBP gates remain pass-only policy gates and do not promote Sprint 11 beyond candidate audit
  status."
  ],
  "risks": [
  "The existing nullrun report has simpler source_artifacts than residualrun; the audit package
  should record its own source_artifact provenance rather than reusing nullrun source fields as
  file-backed proof.",
  "Required negated warnings contain restricted concepts; the validator should use targeted
  allowlist/sanitization for required negated warning text while still rejecting positive or
  unqualified restricted language.",
  "Stability diagnostics could be overinterpreted; report wording must use weak/mixed/unstable/
  candidate-useful language only.",
  "If residual_input_points are sparse or absent, stability audit should block rather than
  synthesize inputs.",
  "The repository currently has uncommitted Sprint 10.x/generated artifacts; implementation should
  avoid reverting unrelated work."
  ],
  "human_review_questions": [],
  "ebp_debt_status": {
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
  "containsFinalTruthClaim": "absent"
  },
  "promotion_status": {
  "sprint11_residual_audit": "residual_audit_candidate_only",
  "full_bmc_toy_gate": "blocked",
  "forbidden_promotions": [
  "physics recovery",
  "classical-limit recovery",
  "model validation",
  "null-model failure conclusion",
  "full-BMC promotion",
  "scientific novelty"
  ]
  },
  "next_smallest_useful_move": "Scaffold internal/bmc/residualaudit with report types, constants,
  deterministic blocked/default report generation, strict ReadReport/ValidateReport, and the first
  validation test before wiring CLI generation."
  }
