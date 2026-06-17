 A concrete gap surfaced: ResidualInputPoint currently has alpha/phi/LHS/RHS/provenance but no
  serialized lambda, while the runner computes with lambda internally. The plan should explicitly
  add that field, populate it from sorted branch points, and validate it only within each computed
  diagnostic.

• Explored
  └ Read main.go, validate.go, residualrun_test.go

• Proposed Plan


  # Sprint 10.3 Implementation Plan

  ## Summary

  Harden residualrun without expanding Sprint 10 scope: dynamically target actual computed
  residual diagnostics, serialize and validate lambda on residual input points, reject forbidden
  formula IDs phrase-safely, and add regression coverage for the exact adversarial cases. Keep
  recovery, superiority, full-BMC, and Sprint 11 behavior blocked.

  ## Key Changes

  - Add lambda to ResidualInputPoint as a JSON field and populate it from the sorted source branch
    point used for each finite-difference residual input.

  - Build ResidualNullComparison.TargetResidualIDs from CandidateResidualDiagnostics where
    ResidualComputed == true; never hard-code candidate_residual_branch_0.

  - Preserve comparison behavior:
      - if at least one computed residual and comparable null diagnostics exist, emit one computed
        comparison targeting all actual computed residual IDs;

      - if computed residuals exist but no comparable null diagnostics exist, emit a blocked
        comparison targeting the actual computed residual IDs;

      - if no computed residuals exist, emit no comparisons.

  - Strengthen validation:
      - every computed diagnostic must have nonempty residual_input_points;
      - each point must match the parent branch ID and sequential point index;
      - lambda, alpha, phi, candidate_left_hand_side, and candidate_right_hand_side must be non-
        nil finite numbers;

      - lambda must be strictly increasing only within that diagnostic’s own
        residual_input_points;

      - input_provenance must be exactly file_read or derived_from_file_read;
      - blocked diagnostics must keep residual_input_points empty.

  - Reject forbidden calculation formula IDs in ValidateReport with phrase-safe messages:
      - friedmann_residual
      - classical_residual
      - recovery_residual
      - cosmology_recovery_residual

  - Keep summary output computed/blocked/total-aware as already implemented; do not change
    promoted status beyond candidate_residual_runner_candidate_only.

  ## Test Plan

  - Add TestResidualRunComparisonTargetsActualComputedResidual against generated report output,
    not a helper-only unit:
      - construct inputs where branch_0 is blocked/uncomputed and branch_1 is eligible/computed;
      - run RunResidualsFromInputs;
      - assert branch_0 residual_computed = false;
      - assert branch_1 residual_computed = true;
      - assert residual_null_comparisons[0].target_residual_ids contains
        candidate_residual_branch_1;

      - assert it does not contain candidate_residual_branch_0.

  - Add validator tests for:
      - computed diagnostic with empty residual_input_points;
      - nonfinite lambda, alpha, phi, LHS, and RHS;
      - nil lambda or other required numeric point fields;
      - non-strict or duplicate lambda ordering within a single computed diagnostic;
      - invalid residual input provenance;
      - forbidden formula IDs using phrase-safe error assertions.

  - Add/confirm clockseg serialization coverage proving branch points are emitted in the clock-
    readiness report.

  - Run:
      - GOCACHE=/tmp/go-build-cache go test ./...
      - CLI generate/validate/summarize for run-residuals
      - lake build

  ## Assumptions

  - lambda should be serialized on each residual input point for standalone audit, but validation
    compares lambda order only inside each diagnostic, never across branches.

  - Dynamic comparison targets should include all computed residual diagnostics in deterministic
    report order.

  - Error messages must identify field/location but must not echo forbidden formula IDs or
    restricted phrases.

  - Sprint 10.3 remains a candidate local/toy residual-runner repair; it must not claim Friedmann
    recovery, classical-limit recovery, superiority over nulls, or full BMC promotion.
