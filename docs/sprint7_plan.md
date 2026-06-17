# BMC Sprint 7: Null-Model Scaffold for Future Friedmann Residuals

This sprint implements a strict null-model specification, input registry, metric contract list, comparison contract registry, and EBP safety gates for future Friedmann-style cosmological comparisons. Under strict EBP 2.1 rules, no actual numerical null-model results or Friedmann residuals will be computed. This sprint serves purely as a specification and data-integrity scaffold to ensure future results can be compared against meaningful baselines rather than generic noise.

## User Review Required

> [!WARNING]
> This sprint **strictly forbids** numerical null-model comparisons, claiming that BMC outperforms any null models, computing actual residuals, or unblocking the full BMC gate. All statuses for null models, metric contracts, and comparison contracts must remain strictly planned/deferred/blocked. The promotion status remains `promoted_nullmodel_spec_artifact_after_repairs`, and full BMC remains blocked.

## Open Questions

> [!NOTE]
> 1. **Phase Resolution**: Are the proposed null models sufficient for multi-field extensions, or are they strictly targeted at the BMC-0A single-scalar minisuperspace superposition? We assume they are strictly configured for BMC-0A.
> 2. **Metric Thresholds**: Should the metric contracts define concrete mathematical pass/fail thresholds in this sprint, or should they only register the metrics for future threshold definitions? We assume registry only.

## Proposed Changes

---

### Go Code Package: `internal/bmc/nullspec`

#### [NEW] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/contracts.go)
Defines constants for allowed status values (`planned`, `deferred`, `blocked`), allowed spec scopes (`null_model_specification_only`), and input availability statuses.

#### [NEW] [registry.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/registry.go)
Defines the `NullModelSpec` structure containing the ID, name, purpose, and control/input/metric associations. Registers the 7 required null models:
1. `constant_phase_control`
2. `randomized_phase_control`
3. `matched_amplitude_randomized_phase_control`
4. `classical_frw_reference_trajectory`
5. `same_branch_segmentation_under_null_wavefunctions`
6. `node_neighborhood_stress_case`
7. `clock_choice_alternative_branch_diagnostic`

#### [NEW] [inputs.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/inputs.go)
Defines `NullModelInputRequirement` tracking input sources from prior sprints (such as `bmc0a_superposition_safe`, `bmc0a_superposition_robustness`, `bmc0a_clock_fragility`, `bmc0a_clock_readiness`, and `bmc0a_friedmann_spec`).

#### [NEW] [metrics.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/metrics.go)
Defines `NullModelMetricContract` to register candidate metrics (e.g., `branch_count_stability`, `turning_point_stability`, `local_relation_single_valuedness_rate`, etc.) without evaluation.

#### [NEW] [comparison_contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/comparison_contracts.go)
Defines `FutureNullComparisonContract` requiring `comparison_computed = false` as a hard invariant.

#### [NEW] [gates.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/gates.go)
Defines `NullSpecGate` evaluating the 10 required safety gates.

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/report.go)
Assembles the deterministic `NullModelSpecReport` structure matching the `bmc0a-nullmodel-spec-v0.1` schema. Provides `GenerateNullModelSpecReport`, `ReadNullModelSpecReport` (with trailing token EOF protection), `WriteJSON`, and `SummarizeNullModelSpecReport`.

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/validate.go)
Enforces strict schema validation rejecting any recovery claims, computed residuals, computed null comparisons, unblocked gates, or unknown JSON fields.

#### [NEW] [nullspec_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec/nullspec_test.go)
Implements all 17 required unit tests validating EBP safety borders.

---

### Command Line Interface

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Adds subcommand `spec-nullmodels`.
- Routes validation and summarization of schema version `bmc0a-nullmodel-spec-v0.1`.

---

### Lean Formal Verification

#### [NEW] [NullModelSpec.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/NullModelSpec.lean)
Defines the `BMCNullModelSpecReport` structure and formally proves the 11 safety theorems.

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
Imports `BMC.NullModelSpec`.

---

## Verification Plan

### Automated Tests
- Run `go test ./internal/bmc/nullspec` to verify all validator invariants.
- Run `cd BMC && lake build` to verify Lean safety proofs.

### Manual Verification
- Generate specification: `ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json`
- Validate specification: `ptw-bmc validate --report out/bmc0a_nullmodel_spec.json`
- Summarize specification: `ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json`

---

## EBP Plan Specification (Required Output)

```json
{
  "summary": "Plan a strict null-model specification, input registry, metric contract list, comparison contract registry, and EBP safety gates for future Friedmann cosmological comparisons. Under strict EBP 2.1 rules, no actual numerical null-model results or Friedmann residuals will be computed.",
  "proposed_actions": [
    "Create internal/bmc/nullspec package containing contracts, structures, registry, inputs, metrics, comparison contracts, safety gates, report serialization, validation, and unit tests.",
    "Integrate spec-nullmodels subcommand into ptw-bmc CLI and route schema validation and summarization.",
    "Add Lean formal verification safety theorems to BMC/BMC/NullModelSpec.lean and import them."
  ],
  "files_to_add": [
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/contracts.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/registry.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/inputs.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/metrics.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/comparison_contracts.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/gates.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/report.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/validate.go",
    "/home/chaschel/Documents/go/bmc/internal/bmc/nullspec/nullspec_test.go",
    "/home/chaschel/Documents/go/bmc/BMC/BMC/NullModelSpec.lean"
  ],
  "files_to_modify": [
    "/home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go",
    "/home/chaschel/Documents/go/bmc/BMC/BMC.lean"
  ],
  "test_plan": [
    "TestNullModelSpecReportValidation",
    "TestNullModelSpecRejectsResidualComputed",
    "TestNullModelSpecRejectsNullComparisonComputed",
    "TestNullModelSpecRejectsRecoveryClaim",
    "TestNullModelSpecRequiresAllRequiredNullModels",
    "TestNullModelSpecRejectsDuplicateNullModelIDs",
    "TestNullModelSpecRequiresBeforeResidualPromotion",
    "TestNullModelSpecRejectsPassedFailedNullModelStatus",
    "TestNullModelSpecRequiresMetricContracts",
    "TestNullModelSpecRejectsComputedFutureComparison",
    "TestNullModelSpecRequiresNoNullComparisonResultGate",
    "TestNullModelSpecRequiresFullBMCBlockedGate",
    "TestNullModelSpecRequiresNullModelDebtActive",
    "TestNullModelSpecRejectsUnknownFields",
    "TestNullModelSpecRejectsTrailingJSONTokens",
    "TestNullModelSpecDeterministicJSON",
    "TestNullModelSpecCLIRouting"
  ],
  "cli_plan": [
    "Add spec-nullmodels subcommand mapping to generated NullModelSpecReport under profile bmc0a-nullmodel-spec",
    "Route bmc0a-nullmodel-spec-v0.1 schema version in validate and summarize subcommands"
  ],
  "lean_plan": [
    "Declare BMCNullModelSpecReport structure representing the spec report",
    "Define reportPassesBMC0AFriedmannNullSpecGate predicate",
    "Prove 11 policy safety theorems including null_model_spec_requires_toy_only, null_model_spec_forbids_residual_computation, null_model_spec_forbids_null_comparison_result, and null_model_spec_requires_full_bmc_blocked"
  ],
  "assumptions": [
    "Null models are configured strictly for BMC-0A single-scalar minisuperspace superposition.",
    "Metric contracts are registered for spec purposes only without quantitative thresholds defined in this sprint."
  ],
  "proof_obligations": [
    "Prove null_model_spec_requires_toy_only",
    "Prove null_model_spec_blocks_final_truth",
    "Prove null_model_spec_forbids_residual_computation",
    "Prove null_model_spec_forbids_null_comparison_result",
    "Prove null_model_spec_forbids_friedmann_recovery_claim",
    "Prove null_model_spec_requires_full_bmc_blocked",
    "Prove null_model_spec_requires_null_model_debt_active",
    "Prove null_model_spec_requires_clock_choice_debt_active",
    "Prove null_model_spec_requires_faithfulness_contested",
    "Prove null_model_spec_does_not_imply_friedmann_recovery",
    "Prove null_model_spec_does_not_imply_full_bmc"
  ],
  "null_models": [
    "constant_phase_control",
    "randomized_phase_control",
    "matched_amplitude_randomized_phase_control",
    "classical_frw_reference_trajectory",
    "same_branch_segmentation_under_null_wavefunctions",
    "node_neighborhood_stress_case",
    "clock_choice_alternative_branch_diagnostic"
  ],
  "risks": [
    "Potential copy-paste schema pollution from prior readiness reports: avoided by establishing distinct NullModelSpecReport structures and a separate bmc0a-nullmodel-spec-v0.1 schema namespace.",
    "Scope creep into numerical model simulations: strictly mitigated by hard validations and compilation blockers."
  ],
  "human_review_questions": [
    "Should metric contracts define concrete mathematical pass/fail thresholds in this sprint, or should they only register the metrics for future threshold definitions?"
  ],
  "ebp_debt_status": {
    "needMap": "active",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "active",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "classical_target_debt": "active",
    "unit_convention_debt": "active",
    "sign_convention_debt": "active",
    "normalization_debt": "active",
    "containsFinalTruthClaim": "absent",
    "LeanVerification": "planned"
  },
  "promotion_status": {
    "sprint7_nullmodel_spec": "promoted_nullmodel_spec_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked",
    "forbidden_promotions": [
      "null models passed",
      "BMC beats null models",
      "Friedmann recovery",
      "ready for Friedmann recovery",
      "full BMC",
      "full quantum gravity",
      "proof of Bohmian mechanics",
      "solution to the problem of time"
    ]
  },
  "next_smallest_useful_move": "Prepare design plans and wait for user review and approval before creating Go/Lean files."
}
```
