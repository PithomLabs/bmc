{
  "summary": "Sprint 7 is mostly a null-model scaffold: generated identity fields are correct,
  residual_computed/null_comparison_computed/friedmann_recovery_claim are false, no null
  simulations or comparison values were found, deterministic JSON holds, Go tests pass, CLI
  routing works, and lake build succeeds. Acceptance still needs repairs because validation is
  permissive for key safety gates and metric prerequisites, future comparison contracts may be
  omitted, and generated/report text uses the phrase-scan bypass "recovery of Friedmann."",
  "overall_verdict": "accept_with_repairs",
  "ebp_debt_review": {
  "needMap": "partial",
  "needInvariant": "partial",
  "needToyCheck": "unpaid",
  "needNullModel": "partial",
  "needObstruction": "unpaid",
  "needFaithfulnessReview": "contested",
  "clock_choice_debt": "unpaid",
  "classical_target_debt": "unpaid",
  "unit_convention_debt": "unpaid",
  "sign_convention_debt": "unpaid",
  "normalization_debt": "unpaid",
  "containsFinalTruthClaim": "absent",
  "LeanVerification": "partial",
  "NullModelSpecDiagnosticIntegrity": "partial",
  "NullComparisonBoundary": "partial",
  "NoResidualComputationBoundary": "partial",
  "FriedmannRecoveryBoundary": "partial"
  },
  "nullspec_findings": [
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/validate.go",
  "line": 109,
  "issue": "Metric contracts do not validate required_before_residual_promotion.",
  "evidence": "ValidateNullModelSpecReport checks metric status only; a metric_contract with
  required_before_residual_promotion=false would still validate."
  },
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/validate.go",
  "line": 116,
  "issue": "Future comparison contracts are not required to be present.",
  "evidence": "The validator iterates future_comparison_contracts but does not reject an empty
  slice, despite Sprint 7 requiring future comparison contracts."
  },
  {
  "severity": "minor_repair",
  "file": "internal/bmc/nullspec/validate.go",
  "line": 82,
  "issue": "Null-model registry validation allows extra unknown null models.",
  "evidence": "The validator ensures the seven required IDs occur once, but it does not reject
  additional null_model_id entries outside the required set."
  }
  ],
  "physics_boundary_findings": [],
  "code_findings": [
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/validate.go",
  "line": 141,
  "issue": "Required safety gates are not all required to pass.",
  "evidence": "Only no_null_comparison_result_gate, full_bmc_blocked_gate, and
  required_before_residual_promotion_gate have explicit pass checks.
  no_residual_computation_gate, toy_analysis_only_gate, no_final_truth_claim_gate,
  friedmann_recovery_claim_blocked_gate, null_model_registry_complete_gate,
  clock_choice_debt_active_gate, and faithfulness_contested_gate can be blocked or contested and
  still validate."
  },
  {
  "severity": "minor_repair",
  "file": "internal/bmc/nullspec/report.go",
  "line": 289,
  "issue": "Generated EBP debt says LeanVerification is planned even though NullModelSpec.lean
  builds.",
  "evidence": "lake build succeeds and BMC/BMC/NullModelSpec.lean exists, but the generated debt
  ledger still records LeanVerification as planned."
  }
  ],
  "cli_findings": [
  {
  "severity": "minor_repair",
  "file": "internal/bmc/nullspec/nullspec_test.go",
  "line": 249,
  "issue": "TestNullModelSpecCLIRouting does not exercise CLI routing.",
  "evidence": "The test only checks the generated schema version; it does not invoke validate/
  summarize routing or unknown profile behavior."
  }
  ],
  "lean_findings": [],
  "overclaim_findings": [
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/report.go",
  "line": 269,
  "issue": "Generated gate text uses the phrase-scan bypass "recovery of Friedmann."",
  "evidence": "The gate reason says "Confirms that no recovery of Friedmann is claimed." The
  prompt explicitly calls out "recovery of Friedmann" as a phrase-scan bypass to inspect."
  },
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/report.go",
  "line": 304,
  "issue": "Generated warnings use the phrase-scan bypass "recovery of Friedmann."",
  "evidence": "out/bmc0a_nullmodel_spec.json contains "No recovery of Friedmann is claimed."
  Prefer the accepted neutral wording "No recovery claim is made.""
  },
  {
  "severity": "repair_required",
  "file": "internal/bmc/nullspec/report.go",
  "line": 385,
  "issue": "CLI summary prints "Recovery of Friedmann Claim."",
  "evidence": "The summarize output exposes the same bypass phrase even though it is negated."
  },
  {
  "severity": "minor_repair",
  "file": "internal/bmc/nullspec/nullspec_test.go",
  "line": 266,
  "issue": "Forbidden phrase test omits the known bypass phrase.",
  "evidence": "The test scans for "Friedmann recovery" but not "recovery of Friedmann", so the
  generated report passes while containing the bypass phrase."
  }
  ],
  "missing_tests": [
  "Test that no_residual_computation_gate.status != pass is rejected.",
  "Test that toy_analysis_only_gate.status != pass is rejected.",
  "Test that no_final_truth_claim_gate.status != pass is rejected.",
  "Test that friedmann_recovery_claim_blocked_gate.status != pass is rejected.",
  "Test that metric_contracts[*].required_before_residual_promotion must be true.",
  "Test that future_comparison_contracts cannot be empty.",
  "Test that unknown extra null_model_id entries are rejected if the registry is intended to be
  exactly the seven-model set.",
  "Test that CLI routing actually validates and summarizes the null-model schema through the ptw-
  bmc command.",
  "Test forbidden semantic bypass phrases including "recovery of Friedmann"."
  ],
  "required_repairs_before_acceptance": [
  "Require every Sprint 7 safety gate to exist exactly once and have the intended pass status
  where the gate represents a satisfied safety condition.",
  "Validate metric_contracts[].required_before_residual_promotion == true.",
  "Reject empty future_comparison_contracts.",
  "Remove "recovery of Friedmann" from generated JSON, summary output, and walkthrough text; use
  neutral wording such as "No recovery claim is made."",
  "Expand tests to cover gate status failures, metric prerequisite failures, empty comparison
  contracts, real CLI routing, and bypass phrase scans."
  ],
  "optional_repairs": [
  "Reject extra unknown null_model_id entries to make the registry exactly the seven expected
  null models.",
  "Update the generated LeanVerification debt label from planned to partial if the intended
  meaning is policy/safety Lean coverage rather than future physics verification.",
  "Add explicit Lean theorem names for classicalTargetDebtActive, unitConventionDebtActive,
  signConventionDebtActive, and normalizationDebtActive, even though the booleans are already
  included in the gate."
  ],
  "faithfulness_verdict": {
  "status": "contested",
  "reason": "The artifact remains a scaffold and does not compute null-model results or Friedmann
  residuals, but diagnostic integrity is only partial until validator gates and phrase-boundary
  repairs are tightened."
  },
  "promotion_recommendation": "promoted_nullmodel_spec_artifact_after_repairs",
  "next_smallest_useful_move": "Patch ValidateNullModelSpecReport to enforce all Sprint 7 gate
  statuses, metric prerequisite flags, and nonempty future comparison contracts, then replace the
  "recovery of Friedmann" wording and add regression tests."
  }
