Sprint 7 is **not accepted yet**. The adversarial review is a clear `accept_with_repairs`.

The artifact stayed inside the big boundary — no null-model results, no Friedmann residual, no BMC-beats-nulls claim — but the validator is still too permissive. The biggest repair is to make every Sprint 7 safety gate and future-comparison prerequisite enforceable, not just documented. This matches the original Sprint 7 requirement that the scaffold define gates and validation, not merely list them. 

```json id="s7_status"
{
  "sprint": "BMC Sprint 7",
  "current_status": "accept_with_repairs",
  "accepted": false,
  "next_step": "Sprint 7.1 nullspec validator and phrase-boundary repair",
  "maximum_future_promotion": "promoted_nullmodel_spec_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked",
  "residual_computed": false,
  "null_comparison_computed": false,
  "friedmann_recovery_claim": false,
  "needNullModel": "active"
}
```

Use this repair prompt.

Sprint 7 adversarial review received. Do not start Sprint 8.

Current status: `accept_with_repairs`.

Implement **Sprint 7.1: Nullspec Validator and Phrase-Boundary Repair** only.

## Required repairs before Sprint 7 acceptance

### 1. Require every Sprint 7 safety gate to exist exactly once and pass

Update:

```text
internal/bmc/nullspec/validate.go
```

The validator must require all 10 Sprint 7 gates exactly once:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
no_residual_computation_gate
no_null_comparison_result_gate
null_model_registry_complete_gate
required_before_residual_promotion_gate
friedmann_recovery_claim_blocked_gate
full_bmc_blocked_gate
clock_choice_debt_active_gate
faithfulness_contested_gate
```

Each required gate should have:

```text
status = pass
```

Reports must fail validation if any required gate is:

```text
missing
duplicated
blocked
contested
unknown
empty
```

Add tests:

```text
TestNullModelSpecRequiresAllGatesExactlyOnce
TestNullModelSpecRejectsNonPassNoResidualComputationGate
TestNullModelSpecRejectsNonPassToyAnalysisOnlyGate
TestNullModelSpecRejectsNonPassNoFinalTruthClaimGate
TestNullModelSpecRejectsNonPassRecoveryClaimBlockedGate
TestNullModelSpecRejectsNonPassClockChoiceDebtGate
TestNullModelSpecRejectsNonPassFaithfulnessContestedGate
```

### 2. Validate metric prerequisite flags

Every metric contract must have:

```text
required_before_residual_promotion = true
```

Validation must reject any metric contract where this is false.

Add test:

```text
TestNullModelSpecRequiresMetricsBeforeResidualPromotion
```

Do not compute metric values or thresholds.

### 3. Require future comparison contracts to be nonempty

The report must include at least one future comparison contract.

Validation must reject:

```text
future_comparison_contracts = []
null future_comparison_contracts
```

Every future comparison contract must still have:

```text
comparison_computed = false
status = planned|deferred|blocked
```

Add tests:

```text
TestNullModelSpecRejectsEmptyFutureComparisonContracts
TestNullModelSpecRejectsComputedFutureComparison
```

### 4. Remove the phrase-scan bypass “recovery of Friedmann”

Update generated JSON, summary output, comments, warnings, and walkthrough text.

Remove wording like:

```text
Recovery of Friedmann Claim
No recovery of Friedmann is claimed
Confirms that no recovery of Friedmann is claimed
```

Use neutral wording instead:

```text
Recovery Claim: false
No recovery claim is made.
No recovery-style interpretation is allowed.
```

Also expand forbidden phrase tests to catch:

```text
recovery of Friedmann
Friedmann support
Friedmann-compatible
control victory
nulls defeated
BMC beats null models
```

Add test:

```text
TestNullModelSpecRejectsSemanticBypassPhrases
```

### 5. Consider exact seven-null-model registry

Prefer rejecting any extra unknown `null_model_id`.

Required null models, exactly once:

```text
constant_phase_control
randomized_phase_control
matched_amplitude_randomized_phase_control
classical_frw_reference_trajectory
same_branch_segmentation_under_null_wavefunctions
node_neighborhood_stress_case
clock_choice_alternative_branch_diagnostic
```

Validation should reject unknown extra IDs unless there is a documented extension mechanism. For Sprint 7, keep the registry exact.

Add test:

```text
TestNullModelSpecRejectsUnknownExtraNullModelID
```

### 6. Improve CLI routing test

Update `TestNullModelSpecCLIRouting` so it does more than inspect schema version.

It should verify that the `ptw-bmc` validation/summarization routing actually dispatches the `bmc0a-nullmodel-spec-v0.1` schema to the nullspec validator/summarizer. If full subprocess testing is too heavy, factor routing into a testable function.

Also test unknown profile failure for:

```bash
ptw-bmc spec-nullmodels --profile unknown-profile
```

### 7. Update LeanVerification debt label

If Lean policy contracts build cleanly, update report wording from:

```text
LeanVerification: planned
```

to something more precise:

```text
LeanVerification: retired_for_policy_safety_contracts
```

Do not imply Lean proves null-model physics or Friedmann recovery.

### 8. Optional Lean strengthening

Optional but useful: add explicit Lean theorem names for:

```text
classicalTargetDebtActive
unitConventionDebtActive
signConventionDebtActive
normalizationDebtActive
```

Keep these policy-only. Do not prove null-model success, Friedmann recovery, or physical validity.

## Required verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json
./ptw-bmc validate --report out/bmc0a_nullmodel_spec.json
./ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Return a concise repair report with:

```text
files changed
validator repairs completed
new/updated tests
exact gate requirements enforced
metric prerequisite validation added
future comparison nonempty validation added
phrase-boundary repairs completed
CLI routing test improved
command results
updated Sprint 7 status
```

## Forbidden scope

Do not implement:

```text
null-model simulations
null-model comparison results
BMC beats null models claim
Friedmann residual computation
Friedmann recovery
ready-for-Friedmann claim
full BMC promotion
massive scalar model
LQC comparison
Page-Wootters comparison
full quantum gravity claim
proof of Bohmian mechanics
solution to the problem of time
```

Sprint 7 remains a null-model scaffold artifact only.

After these repairs pass, Sprint 7 can be considered for:

```text id="x8a2e5"
promoted_nullmodel_spec_artifact_after_repairs
```

Until then, it remains:

```text id="j4s9qt"
nullmodel_spec_candidate_only_pending_repairs
```
