Approve Sprint 5 planning **with revisions before implementation**. The direction is correct: local clock segmentation and clock-independent diagnostics are the right next step after Sprint 4’s `clock_fragile` result.

## Answers to open questions

Yes, add a few clock-independent diagnostics explicitly:

```text
path_length_in_configuration_space
total_lambda_interval
num_valid_trajectory_points
num_clock_segments
num_turning_points
min_amplitude_r
max_abs_q_away_from_nodes
max_phase_gradient
node_contact_free
trajectory_finiteness
```

For `α(φ)` single-valuedness, **do not fail on tiny numerical noise**. Use a parameterized tolerance, for example:

```text
single_valuedness_epsilon = 1e-9
```

and include it in the report and validator. Also require enough samples per branch, such as:

```text
min_branch_samples = 3
```

A two-point branch may be monotonic but too weak for meaningful relational diagnostics.

## Required corrections

### 1. Fix promotion status

This is too strong for a planning document:

```json
"sprint5_clock_readiness": "promoted_clock_readiness_artifact_after_repairs"
```

Use:

```json
"sprint5_clock_readiness": "planned_candidate_only"
```

Maximum future promotion remains:

```text
promoted_clock_readiness_artifact_after_repairs
```

but only after implementation, tests, CLI verification, Lean build, walkthrough, adversarial review, and repairs.

### 2. Update the null-model field

Sprint 5 does not add null models. Prefer:

```json
"null_models": [
  "No new null models planned. Existing null-model debt remains partial/deferred."
]
```

Listing Classical FRW, WdW, LQC, and Page-Wootters may imply more work than this sprint actually performs.

### 3. Add branch-quality validation

Add validation rules for:

```text
segments are deterministic and non-overlapping
segment start_index <= end_index
segment has at least min_branch_samples for local relation extraction
segment lambda range is finite
clock_range is finite and nonnegative
local relation branch cannot pass if clock_range is too small
turning points are sorted by index/lambda
```

### 4. Strengthen readiness wording

Use:

```text
local_only_candidate
```

only to mean:

```text
local φ branches exist and can be inspected safely
```

It must **not** mean:

```text
ready for Friedmann recovery
```

Good validation rule:

```text
reject friedmann_readiness = local_only_candidate if no valid local relation branches exist
```

but also include a warning:

```text
local_only_candidate is not Friedmann recovery readiness
```

### 5. Expand test plan slightly

Add these tests:

```text
TestClockReadinessRejectsReadyForFriedmannLanguage
TestClockSegmentsAreNonOverlappingAndOrdered
TestLocalRelationRejectsTooFewSamples
TestLocalOnlyCandidateDoesNotUnblockFullBMC
```

## Approved decisions

```json
{
  "segment_clock_subcommand": "approved",
  "clockseg_package": "approved",
  "local_phi_branch_segmentation": "approved",
  "alpha_phi_single_valuedness_check": "approved_with_parameterized_epsilon",
  "clock_independent_diagnostics": "approved",
  "friedmann_readiness_field": "approved_only_with_blocked_or_local_only_candidate_or_contested",
  "Lean_scope": "approved_policy_only",
  "full_bmc_gate": "must_remain_blocked"
}
```

## Revised EBP status

```text
needMap: partial
needInvariant: partial
needToyCheck: active
needNullModel: partial/deferred
needObstruction: active
needFaithfulnessReview: contested
clock_choice_debt: active
containsFinalTruthClaim: absent
LeanVerification: planned
promotion_status: planned_candidate_only
```

Final instruction:

```text
Sprint 5 plan approved with revisions. Implement only local clock segmentation, local α(φ) branch checks, and clock-independent diagnostics. Do not implement Friedmann residual recovery, massive scalar, LQC/Page-Wootters comparison, full BMC promotion, or any full quantum-gravity claim.
```
