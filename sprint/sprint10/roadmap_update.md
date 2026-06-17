We are currently in **Tier 1: BMC Toy Benchmark**, still early, but no longer at the “just scaffolding” stage.

Plain English:

```text
We have built the safety rails, the toy wave model, the branch/clock diagnostics, the null-model comparison layer, and the prior-art boundary.
Now we are preparing the first dangerous-but-important step: computing candidate residual diagnostics on local branches — without calling it recovery.
```

## Where we are on the roadmap

### Tier 0 — PTW / EBP Core

Status: **partially established**

This is the workbench discipline: schemas, gates, validator strictness, no-overclaim checks, Lean policy stubs, CLI report generation, and adversarial review loop.

We have enough Tier 0 machinery to safely run BMC toy sprints, but PTW is not yet a full general physics workbench.

### Tier 1 — BMC Toy Benchmark

Status: **active; this is where we are now**

This is the current main track.

We have completed:

```text
Sprint 1: plane-wave control
Sprint 2: superposition + node obstruction
Sprint 3: robustness audit
Sprint 4: clock-fragility diagnostic
Sprint 5: local branch / clock-readiness segmentation
Sprint 6: candidate residual specification
Sprint 7: null-model scaffold
Sprint 8-Lite: prior-art boundary note
Sprint 9 + 9.1: null-model runner accepted after repairs
```

We are now at:

```text
Sprint 10: candidate local-branch residual runner
```

Sprint 10 is planned and approved **with guardrail revisions**, but implementation has not yet been accepted.

## What Sprint 10 means in simple terms

Imagine we built a small test track.

First, we checked whether the toy car can move without crashing into obvious walls: nodes, nonfinite values, bad clocks, broken branches.

Then we built fake/control tracks — null models — to make sure we are not fooling ourselves.

Now Sprint 10 asks:

```text
On a local branch that is clean enough to inspect, what does the candidate residual look like?
```

But we are **not** allowed to say:

```text
This recovers Friedmann cosmology.
This validates Bohmian cosmology.
This beats null models.
This proves anything about quantum gravity.
```

We are only saying:

```text
We can compute a bounded local diagnostic and compare it against controls.
```

## Current truth-status

The honest status is:

```json
{
  "current_tier": "Tier 1: BMC toy benchmark",
  "current_sprint": "Sprint 10 planning/implementation boundary",
  "last_accepted_artifact": "Sprint 9.1 null-model runner",
  "next_artifact": "candidate local-branch residual runner",
  "physics_claim_promoted": false,
  "friedmann_recovery_claim": false,
  "full_bmc_promoted": false,
  "null_model_layer_status": "partial",
  "faithfulness_review": "contested"
}
```

## What we have actually earned

We have earned this:

```text
BMC-0A is now a disciplined toy benchmark with:
- node obstruction checks
- clock fragility diagnostics
- local branch segmentation
- prior-art humility
- null-model scaffolding
- null-model runner with honest accounting
- strict anti-overclaim validation
```

We have **not** earned this:

```text
BMC recovers classical cosmology.
BMC solves the problem of time.
BMC beats null models.
Bohmian quantum cosmology is validated.
The framework has physical confirmation.
```

## The big picture

We are still in the **toy-model refinery** stage.

That is good. This is exactly the phase where PTW should be boring, strict, and suspicious.

The next few steps are likely:

```text
Sprint 10: candidate local-branch residual runner
Sprint 10.1: adversarial repair, almost certainly needed
Sprint 11: residual/null comparison hardening
Sprint 12: convention-debt review and faithfulness audit
Sprint 13: decide whether BMC-0A deserves a stronger toy-status promotion
```

## EBP debt snapshot

```text
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
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

So the shortest update is:

```text
We are past scaffolding and now entering the first real diagnostic-computation phase, but still inside a toy local-branch sandbox. The next milestone is Sprint 10: compute candidate residual diagnostics without interpreting them as recovery.
```
