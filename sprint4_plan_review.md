Approve Sprint 4 planning **with revisions before implementation**. The plan is well aligned with the Sprint 3 result: it investigates the **clock-monotonicity fragility** instead of jumping to Friedmann recovery or new physics. The proposed `diagnose-clock` subcommand is acceptable and cleaner than overloading `audit`, because this is a different schema and diagnostic mode. 

## Required corrections

### 1. Fix `promotion_status`

This is too strong for a planning document:

```json
"sprint4_clock_fragility_diagnostic": "promoted_clock_fragility_diagnostic_artifact_after_repairs"
```

It should be:

```json
"sprint4_clock_fragility_diagnostic": "planned_candidate_only"
```

Maximum future promotion remains:

```text
promoted_clock_fragility_diagnostic_artifact_after_repairs
```

but only **after implementation, tests, Lean build, CLI verification, walkthrough, and adversarial review**.

### 2. Use `containsFinalTruthClaim: "absent"` instead of `"false"`

For consistency with prior EBP ledgers:

```json
"containsFinalTruthClaim": "absent"
```

not:

```json
"containsFinalTruthClaim": "false"
```

### 3. Do not let Sprint 4 classify “expected Bohmian behavior” too confidently

It can say:

```text
consistent_with_toy_bohmian_trajectory_behavior
```

but not:

```text
expected Bohmian behavior
```

unless there is an analytic derivation. Keep that category softer:

```text
clock_choice_limitation_or_toy_trajectory_feature
```

### 4. Parameterize the near-zero threshold

The plan currently uses:

```text
|dφ/dλ| < 1e-10
```

That is fine as a default, but add it to report parameters:

```json
"near_zero_dphi_threshold": 1e-10
```

and include it in validation. This avoids hiding a threshold choice.

### 5. Do not make Sprint 5 automatically Friedmann

This sentence is too fast:

```text
If Sprint 4 finds clock_stable: proceed to Friedmann residual check.
```

Better:

```text
If Sprint 4 finds clock_stable, consider a small Sprint 5 plan for relational-clock handling or Friedmann-residual readiness review.
```

A stable φ-clock would be useful, but Friedmann recovery still needs its own careful map/invariant definition.

## Approved decisions

```json
{
  "diagnose_clock_subcommand": "approved",
  "clockdiag_package": "approved",
  "event_detection": "approved_with_parameterized_threshold",
  "step_refinement_rechecks": "approved",
  "alternative_clock_summary": "approved_as_diagnostic_only",
  "trajectory_validity_vs_clock_validity_distinction": "strongly_approved",
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
Sprint 4 plan approved with the corrections above. Implement only the clock-fragility diagnostic. Do not implement Friedmann recovery, massive scalar, LQC/Page-Wootters comparison, full BMC promotion, or any full quantum-gravity claim.
```
