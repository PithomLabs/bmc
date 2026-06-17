Accepted.

```json id="s11_1_final_acceptance"
{
  "sprint": "BMC Sprint 11.1",
  "artifact": "residual/null comparison audit artifact",
  "status": "accepted_after_repairs",
  "promotion_status": "promoted_residual_audit_artifact_after_repairs",
  "artifact_path": "out/bmc0a_residual_audit.json",
  "validation": "passed",
  "go_tests": "passed",
  "lean_build": "passed",
  "physics_claim_promoted": false,
  "full_bmc_toy_gate": "blocked"
}
```

This closes the Sprint 1–11 audit stack.

What was earned:

```text id="earned_s11"
BMC now has a candidate-only residual/null comparison audit artifact that checks comparison integrity and interval-level proxy stability, with strict EBP boundaries.
```

What was **not** earned:

```text id="not_earned_s11"
Friedmann recovery
classical-limit recovery
BMC validation
null-model failure
BMC superiority
scientific novelty
full BMC promotion
```

## Updated EBP status

```text id="ebp_after_s11"
ResidualAuditIntegrity: retired for Sprint 11 audit-artifact scope
StabilityAuditIntegrity: retired for Sprint 11.1 interval-proxy audit scope
ComparisonIntegrity: partial
SourceProvenanceIntegrity: partial
NoRecoveryClaimBoundary: retired
NoBMCBeatsNullsBoundary: retired
LeanPolicyBoundary: retired

needLiteratureAudit: partial
needMap: partial
needInvariant: partial
needToyCheck: partial
needNullModel: partial
needObstruction: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
full_bmc_toy_gate: blocked
```

## Process decision

Freeze the audit/report expansion here.

Next phase:

```text id="next_phase"
Phase 1 — Postmortem Inventory
```

Goal:

```text id="phase1_goal"
Inventory what exists, classify what is real computation vs. scaffolding/debt, and identify the smallest failure-detection repairs before adding any more physics claims or schemas.
```

No Sprint 12 yet. The next prompt should be for **Phase 1 Postmortem Inventory**.
