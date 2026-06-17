Our next step should be:

```text id="hy6u58"
Sprint 8: Literature / Prior-Art Audit Profile
```

Do **not** compute the Friedmann residual yet. Do **not** run null-model comparisons yet. Sprint 7 gave us the null-model scaffold, but before we interpret any future result, we should first check what parts of our BMC toy model are already known in the Bohmian quantum cosmology literature and what parts, if any, are distinctive as PTW/EBP workflow. The Sprint 7 scaffold was explicitly designed to keep null-model and recovery claims blocked until later validation exists. 

## Why Sprint 8 should be literature audit

Because we already clarified that much of the physics terrain has prior art:

```text id="jpbz9s"
Bohmian minisuperspace cosmology
Wheeler-DeWitt scalar-field models
Bohmian trajectories
quantum potential analysis
internal clock choices
singularity/bounce studies
classical-limit discussions
```

So before we build more physics machinery, we need to map:

```text id="3p6kk5"
What is standard?
What is copied/anticipated by prior work?
What is only our implementation variant?
What is genuinely distinctive?
Which claims must be downgraded?
Which papers become required source artifacts?
```

This protects us from accidentally treating old ideas as new.

## Sprint 8 goal

```text id="t4ce9p"
Build a prior-art audit layer that maps BMC-0A claims against existing Bohmian quantum cosmology literature.
```

This is not a physics-result sprint. It is a **claim-overlap and provenance sprint**.

## Concrete Sprint 8 outputs

I would build these artifacts:

```text id="jw1vvo"
1. paper registry
2. prior-art claim registry
3. BMC claim-to-literature overlap map
4. novelty/debt classifier
5. required-source ledger
6. citation/provenance report
7. forbidden-novelty-claim gate
```

The generated report should answer:

```text id="z1tx5h"
Claim C001: BMC uses Bohmian trajectories in minisuperspace.
Prior-art status: established.

Claim C002: BMC checks scalar-field clock monotonicity.
Prior-art status: likely related; requires source comparison.

Claim C003: BMC uses EBP gates, null-model scaffolds, and Lean policy contracts.
Prior-art status: unknown / possibly distinctive workflow; requires audit.

Claim C004: BMC recovers Friedmann dynamics.
Prior-art status: not claimed; blocked.
```

## Proposed package

```text id="2zz6cn"
internal/bmc/litaudit/
```

Possible files:

```text id="18hw8o"
contracts.go
papers.go
claims.go
overlap.go
novelty.go
gates.go
report.go
validate.go
litaudit_test.go
```

CLI:

```bash id="ph3zi7"
ptw-bmc audit-literature --profile bmc0a-prior-art --out out/bmc0a_prior_art_audit.json
ptw-bmc validate --report out/bmc0a_prior_art_audit.json
ptw-bmc summarize --report out/bmc0a_prior_art_audit.json
```

## Required gates

```text id="lm8hkh"
toy_analysis_only_gate
no_novelty_claim_gate
no_friedmann_recovery_claim_gate
prior_art_registry_nonempty_gate
claim_overlap_map_complete_gate
required_sources_declared_gate
full_bmc_blocked_gate
faithfulness_contested_gate
```

## Important statuses

The literature audit should classify claims like this:

```text id="otxmy4"
established_prior_art
likely_prior_art
partial_overlap
implementation_variant
workflow_distinctive_candidate
unknown_requires_review
not_claimed
blocked
```

It must **not** use:

```text id="9ggzi3"
novel
proven_new
first_ever
scientifically original
breakthrough
```

unless a proper human-reviewed literature review supports it later.

## What comes after Sprint 8

After Sprint 8, we can choose between two next paths:

```text id="8d0a3f"
Sprint 9A: Paper-claim extraction pipeline
Sprint 9B: BMC null-model runner
```

My recommendation:

```text id="wdjkc9"
Sprint 8: Literature/prior-art audit
Sprint 9: Paper-claim extraction pipeline
Sprint 10: BMC null-model runner
Sprint 11: Candidate local-branch residual runner
```

That order fits the roadmap we locked in: first strengthen the workbench’s ability to read and classify claims, then run more physics diagnostics.

## EBP status

```text id="5rzpev"
Claim: Next step should be Sprint 8 literature/prior-art audit.
Status: recommended process step.

Claim: We are ready to compute Friedmann residuals.
Status: blocked.

Claim: We are ready to claim novelty.
Status: blocked.

Claim: Sprint 7 scaffold is accepted.
Status: accepted as null-model scaffold only.

needLiteratureAudit: active
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: active, scaffold accepted but no results
needObstruction: active
needFaithfulnessReview: contested
full_bmc_toy_gate: blocked
promotion_status: prior-art audit candidate only
```
