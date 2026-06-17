# BMC Sprint 5: Relational-Clock Readiness and Local Clock Segmentation

This is the proposed implementation plan for BMC Sprint 5 under the EBP 2.1 discipline. The goal is a narrow readiness sprint to determine whether BMC-0A can define safe relational diagnostics when φ is not globally monotonic.

## User Review Required

> [!WARNING]
> This sprint **strictly forbids** implementing Friedmann residual recovery or attempting to solve the problem of time. Any attempt to claim full quantum gravity, validity of φ-clock for full cosmology, or ready for Friedmann recovery will be rejected. The focus is strictly on **local clock segmentation** and **clock-independent diagnostics**.

## Open Questions

> [!IMPORTANT]
> 1. Are there any additional clock-independent diagnostics that should be explicitly tracked in this sprint?
> 2. Should `α(φ)` single-valuedness check strictly fail if there is any numerical noise, or tolerate a tiny epsilon threshold?

## Implementation Plan

### JSON Plan Output

Below is the structured plan as requested:

```json
{
  "summary": "Sprint 5 is a relational-clock readiness sprint aimed at determining if BMC-0A can define safe relational diagnostics when phi is not globally monotonic. It will segment local monotonic branches, detect turning points, and evaluate clock-independent diagnostics without attempting Friedmann residual recovery or claiming full quantum gravity.",
  "proposed_actions": [
    "Create internal/bmc/clockseg package to handle local monotonic segment detection and branch extraction.",
    "Define structures for ClockSegment, ClockTurningPoint, LocalRelationBranch, and ClockIndependentDiagnostic.",
    "Implement turning-point detection using the parameterized near-zero dphi threshold (1e-10) and segment the trajectory based on phi monotonicity.",
    "Assess alpha(phi) single-valuedness on each local branch.",
    "Compute clock-independent diagnostics (trajectory_finiteness, node_contact_free, max_abs_q_away_from_nodes, min_amplitude_r, phase_gradient_finite).",
    "Add step-refinement branch stability checks for the four fragile configurations (dt=0.05, 0.025, 0.0125).",
    "Generate deterministic bmc0a-clock-readiness-v0.1 JSON reports.",
    "Implement strict validation for the readiness report following EBP guardrails.",
    "Add 'segment-clock' subcommand to ptw-bmc CLI.",
    "Add BMCClockReadinessReport and policy theorems to Lean."
  ],
  "files_to_add": [
    "internal/bmc/clockseg/segments.go",
    "internal/bmc/clockseg/branches.go",
    "internal/bmc/clockseg/turning_points.go",
    "internal/bmc/clockseg/local_relations.go",
    "internal/bmc/clockseg/clock_independent.go",
    "internal/bmc/clockseg/report.go",
    "internal/bmc/clockseg/validate.go",
    "internal/bmc/clockseg/clockseg_test.go",
    "BMC/BMC/ClockReadiness.lean"
  ],
  "files_to_modify": [
    "cmd/ptw-bmc/main.go",
    "BMC/BMC.lean"
  ],
  "test_plan": [
    "TestClockReadinessReportValidation",
    "TestClockSegmentationDetectsLocalBranches",
    "TestTurningPointsDetectedAtBranchBoundaries",
    "TestLocalRelationBranchesAreSingleValuedOnSegments",
    "TestClockIndependentDiagnosticsDoNotRequirePhiClock",
    "TestStepRefinementBranchAuditDeterministic",
    "TestClockReadinessRejectsFinalTruthClaim",
    "TestClockReadinessRejectsNonfiniteMetrics",
    "TestClockReadinessRequiresClockChoiceDebtActive",
    "TestClockReadinessDeterministicJSON"
  ],
  "cli_plan": [
    "Add subcommand 'segment-clock' (e.g., ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json)",
    "Update 'validate' to handle bmc0a-clock-readiness-v0.1 schema",
    "Update 'summarize' to handle bmc0a-clock-readiness-v0.1 schema"
  ],
  "lean_plan": [
    "Add BMC/BMC/ClockReadiness.lean defining BMCClockReadinessReport structure",
    "Add theorem clock_readiness_requires_toy_only",
    "Add theorem clock_readiness_blocks_final_truth",
    "Add theorem clock_readiness_requires_friedmann_deferred",
    "Add theorem clock_readiness_does_not_imply_full_bmc",
    "Add theorem clock_readiness_keeps_clock_choice_debt_active",
    "Add theorem clock_readiness_local_candidate_does_not_mean_friedmann_recovered",
    "Import BMC.ClockReadiness in BMC/BMC.lean"
  ],
  "assumptions": [
    "Sprint 5 strictly adheres to EBP 2.1 boundaries.",
    "Friedmann residual recovery remains completely blocked.",
    "Clock choice debt remains active."
  ],
  "proof_obligations": [
    "Demonstrate stable detection of local monotonic branches under step refinement.",
    "Demonstrate proper JSON schema validation rejecting invalid states, missing values, and non-finite metrics."
  ],
  "null_models": [
    "Classical FRW (placeholder)",
    "Standard WdW (placeholder)",
    "LQC (deferred)",
    "Page-Wootters (deferred)"
  ],
  "risks": [
    "Confusion between local clock readiness and full Friedmann recovery.",
    "Branch boundary detection sensitivity to step size refinement."
  ],
  "human_review_questions": [
    "Are there any additional clock-independent diagnostics that should be explicitly tracked in this sprint?",
    "Should alpha(phi) single-valuedness check strictly fail if there is any numerical noise, or tolerate a tiny epsilon threshold?"
  ],
  "ebp_debt_status": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "containsFinalTruthClaim": "absent",
    "LeanVerification": "planned"
  },
  "promotion_status": {
    "sprint5_clock_readiness": "promoted_clock_readiness_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked",
    "forbidden_promotions": [
      "full BMC",
      "full quantum gravity",
      "proof of Bohmian mechanics",
      "solution to the problem of time",
      "spacetime emergence proof",
      "Friedmann recovery",
      "ready for Friedmann recovery",
      "valid phi-clock for full cosmology",
      "valid alpha-clock for full cosmology"
    ]
  },
  "next_smallest_useful_move": "Approve this implementation plan so I can begin writing the internal/bmc/clockseg package and Lean safety contracts."
}
```

## Verification Plan

### Automated Tests
- Run `go test ./internal/bmc/clockseg` to verify branch detection, turning points, validation, and JSON generation.
- Verify Lean theorems compile and are verified via Lake build.

### Manual Verification
- Run `go run ./cmd/ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json`
- Verify the generated report schema using `go run ./cmd/ptw-bmc validate --report out/bmc0a_clock_readiness.json`
- Validate the summary view via `go run ./cmd/ptw-bmc summarize --report out/bmc0a_clock_readiness.json`
