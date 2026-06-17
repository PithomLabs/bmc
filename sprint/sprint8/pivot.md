## prompt

one moment, does sprint 8 risk making our project bloated? first it started as a toy model, then I try to make it as workbench implementation based on EBP 2.1 and also expand it to include testing other branches of physics. what do you think


Yes, Sprint 8 **does risk bloating the project** if we implement it as a big literature-review subsystem too early.

Your instinct is correct. The project started as:

```text
a tiny BMC toy model
```

Then it became:

```text
a disciplined EBP-governed workbench
```

Then the roadmap expanded toward:

```text
a phased physics-claim audit platform
```

That expansion is legitimate, but it creates a real danger: **architecture inflation before the core toy model earns its next physics result**.

## My honest view

Sprint 8 is conceptually right, but the implementation plan I gave is probably **too heavy for now**.

It tries to add:

```text
paper registry
claim registry
overlap registry
source ledger
novelty boundary
Lean policy file
CLI routing
validation gates
forbidden phrase scanner
full prior-art report schema
```

That is coherent, but it may be premature.

The toy model is still in the middle of proving its own usefulness. Before we build a general literature-audit machine, we should avoid turning the project into a framework that keeps adding scaffolds but never reaches the next experimental step.

## The danger pattern

The danger is this:

```text
toy model → governance layer → audit layer → literature layer → paper extraction layer → profile mapper → dashboard → no actual physics test
```

That would be bad.

Not because the pieces are wrong, but because the center of gravity drifts away from the toy model.

The workbench should grow **from pressure created by the toy model**, not from imagined future needs.

## Better framing

Sprint 8 should not be a full “literature/prior-art audit profile” yet.

It should be a **small prior-art note attached to BMC-0A**.

A slimmer Sprint 8 would answer only:

```text
Are we accidentally treating known Bohmian minisuperspace ideas as new?
Which 5–10 sources must be cited before future claims?
Which BMC claims are standard prior art vs PTW workflow-specific?
```

That is enough.

No big framework yet.

## Revised Sprint 8 recommendation

I would shrink Sprint 8 to:

```text
Sprint 8: BMC-0A Prior-Art Boundary Note
```

Not:

```text
Sprint 8: Literature / Prior-Art Audit Profile
```

The output should be one small JSON/Markdown artifact, not a full subsystem.

### Minimal outputs

```text
1. prior_art_sources.md
2. bmc0a_prior_art_boundary.json
3. one validator test that blocks novelty claims
4. one CLI summary or maybe no CLI at all yet
```

Maybe package:

```text
internal/bmc/priorart/
```

But even that may be optional. A static artifact plus validation may suffice.

## What to avoid for now

Do not build yet:

```text
general paper ingestion
full claim extraction
claim-to-profile mapper
large literature registry
general physics branch audit engine
dashboard
citation database
multi-domain profiles
```

Those belong later, after the BMC pipeline produces more meaningful results.

## What we should keep from Sprint 8

Keep only the anti-overclaim function.

The useful Sprint 8 purpose is:

```text
Prevent novelty inflation.
```

So the artifact should say:

```text
Bohmian minisuperspace trajectories are prior art.
WdW scalar-field models are prior art.
Quantum potential analysis is prior art.
Clock-choice issues are prior art or likely prior art.
PTW/EBP gated audit workflow may be distinctive, but this is only candidate workflow distinctiveness.
No scientific novelty claim is allowed.
```

That is enough.

## What I would do next instead

I would revise the near-term sequence:

```text
Sprint 8: Small BMC-0A Prior-Art Boundary Note
Sprint 9: BMC Null-Model Runner
Sprint 10: Candidate Local-Branch Residual Runner
Sprint 11: Paper-Claim Extraction Prototype
```

This keeps us close to the toy model.

The paper-claim workbench still matters, but not yet. It should arrive after we have a stronger example of what a profile actually tests.

## Why this is healthier

The current BMC track has momentum:

```text
controls → superposition → robustness → clock fragility → local branches → residual spec → null-model scaffold
```

The natural next physics/software step is:

```text
actually run null models
```

But before that, a small prior-art boundary note is wise.

So the balanced choice is:

```text
Do enough literature audit to avoid overclaiming,
but not so much that we stop progressing the toy model.
```

## Revised decision

I would **not** implement the full Sprint 8 plan I just wrote.

I would downgrade it.

Use this instead:

```text
Sprint 8: BMC-0A Prior-Art Boundary Note
Scope: small
Goal: anti-novelty guardrail
No big registry
No paper-ingestion system
No general physics-branch audit system
```

Then proceed to the null-model runner.

## EBP status

```text
Claim: Full Sprint 8 prior-art audit profile risks bloat.
Status: accepted concern.

Claim: Some prior-art boundary check is still needed.
Status: recommended.

Claim: We should pause the whole toy model to build a general literature workbench.
Status: rejected for now.

Claim: PTW can later expand into paper-claim extraction and other physics branches.
Status: roadmap-level, not immediate implementation.

Recommended next step: shrink Sprint 8, then return to BMC null-model execution.
Promotion status: keep BMC-focused, avoid framework inflation.
```



