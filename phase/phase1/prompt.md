# Phase 1 Postmortem Inventory: Freeze Sprint 1–11 and Classify the BMC Surface Area

You are working on the Go-based BMC/PTW codebase under strict EBP 2.1 discipline.

Current status:

```text
Sprint 1–11 audit stack is frozen.
Sprint 11.1 artifact hygiene repair is accepted.
out/bmc0a_residual_audit.json exists and validates.
go test ./... passed.
lake build passed.
```

Do **not** start Sprint 12.

Do **not** add a new physics model.

Do **not** add a new report schema unless the inventory itself requires a small postmortem schema.

Do **not** claim Friedmann recovery, classical-limit recovery, BMC validation, null-model failure, BMC superiority, scientific novelty, or full BMC promotion.

## Phase 1 goal

Create a **postmortem inventory** of the current BMC codebase after Sprints 1–11.

The purpose is to identify:

```text
1. what is real computation,
2. what is validation/audit scaffolding,
3. what is an active debt tracker,
4. what is a deferral stub,
5. what is potentially validation theater,
6. what failure-detection tests are missing,
7. what should be frozen before remediation.
```

This is an inventory phase, not a refactor phase.

## Required deliverable

Create a markdown postmortem document:

```text
docs/postmortem/bmc_sprint_1_11_inventory.md
```

If the repo prefers another docs folder, use:

```text
docs/sprint/bmc_sprint_1_11_inventory.md
```

but choose one location and document it clearly.

## Inventory sections

The document must contain these sections:

```text
1. Executive Summary
2. Current Accepted Artifact Stack
3. Report Schema Inventory
4. CLI Subcommand Inventory
5. Package Inventory
6. Test Inventory
7. Generated Artifact Inventory
8. Real Computation vs. Scaffolding Classification
9. Deferral Stub Inventory
10. Failure-Detection Gap Inventory
11. Numerical-Analysis Debt Inventory
12. Physics-Faithfulness Debt Inventory
13. Validation-Theater Risk Register
14. Freeze Recommendations
15. Next Remediation Tickets
16. EBP Status Summary
```

## 1. Executive Summary

State clearly:

```text
BMC-0.1 is currently a methodological toy benchmark and audit stack, not a physics result.
The Sprint 1–11 stack blocks overclaims but still lacks nontrivial physics failure detection.
Full BMC remains blocked.
```

Do not use promotional wording.

## 2. Current Accepted Artifact Stack

List Sprints 1–11:

```text
Sprint 1: plane-wave control
Sprint 2: superposition + node obstruction
Sprint 3: robustness audit
Sprint 4: clock fragility diagnostic
Sprint 5: local branch / clock readiness
Sprint 6: candidate Friedmann residual specification only
Sprint 7: null-model scaffold
Sprint 8-Lite: prior-art boundary note
Sprint 9 + 9.1: null-model runner
Sprint 10 + 10.3: candidate local-branch residual runner
Sprint 11 + 11.1: residual/null comparison audit artifact
```

For each sprint, classify:

```text
artifact_kind
accepted_status
real_computation_level
scaffolding_level
promotion_scope
forbidden_interpretations
remaining_debts
```

## 3. Report Schema Inventory

Inventory every schema found in the codebase and generated artifacts, including but not limited to:

```text
bmc-report-v0.1
bmc0a-superposition-robustness-v0.1
bmc0a-clock-fragility-v0.1
bmc0a-clock-readiness-v0.1
bmc0a-friedmann-spec-v0.1
bmc0a-nullmodel-spec-v0.1
bmc0a-prior-art-boundary-v0.1
bmc0a-nullrun-v0.1
bmc0a-local-residual-v0.1
bmc0a-residual-audit-v0.1
```

For each schema, record:

```text
schema_version
owning package
CLI generator if any
validator route
summarizer route
generated artifact path
active computation / audit scaffold / deferral stub / legacy
whether it can fail meaningfully
whether it blocks overclaims
whether it risks validation theater
```

## 4. CLI Subcommand Inventory

Inventory every `ptw-bmc` subcommand.

For each subcommand, record:

```text
name
profile(s)
source package
output schema
output path if standard
whether it computes, audits, validates, summarizes, or only specifies
current accepted status
unknown-profile behavior
```

Include at least:

```text
run
audit
diagnose-clock
segment-clock
spec-friedmann
spec-nullmodels
run-nullmodels
run-residuals
audit-residuals
validate
summarize
```

If names differ, use the actual repo names.

## 5. Package Inventory

Inventory `internal/bmc/...` packages.

For each package, classify as:

```text
core physics computation
trajectory/guidance computation
wavefunction model
constraint/residual computation
audit/report scaffold
validation/gating
specification only
prior-art/boundary note
Lean policy support
```

Flag packages that are mostly schema/report machinery.

## 6. Test Inventory

Inventory tests across the repo.

Classify each test into one of:

```text
positive/control test
negative/failure-detection test
integration test
CLI routing test
validation/gate test
determinism test
numerical-convergence test
sensitivity test
forbidden-phrase/anti-overclaim test
Lean policy test
```

Also mark gaps:

```text
wrong wavefunction tests missing
constraint violation tests missing
Euler vs RK4 convergence missing
step-size convergence missing
node-threshold sensitivity missing
phase-gradient h-sensitivity missing
Q near-node policy tests missing
dimensional/unit consistency missing
Friedmann target comparison missing
```

Do not invent tests. Inspect actual test files.

## 7. Generated Artifact Inventory

Inventory generated artifacts under `out/`.

For each artifact, record:

```text
path
schema_version
generator command
validation command
summary command
accepted sprint
whether it should be tracked in repo
whether it is reproducible
whether it is stale or current
```

At minimum check:

```text
out/bmc0a_clock_readiness.json
out/bmc0a_nullrun.json
out/bmc0a_local_residual.json
out/bmc0a_residual_audit.json
```

and any other existing `out/*.json`.

## 8. Real Computation vs. Scaffolding Classification

Create a table with rows like:

```text
Plane-wave WdW residual
Plane-wave guidance trajectory
Two-term superposition wavefunction
Node obstruction detection
Clock fragility diagnostic
Clock segmentation/local branch extraction
Null model runner
Candidate local residual runner
Residual/null audit
Prior-art boundary note
Friedmann residual spec
Null model spec
Lean policy files
```

Classify each as:

```text
real_computation
toy_computation
audit_scaffold
policy_scaffold
specification_only
deferred
risk_of_validation_theater
```

## 9. Deferral Stub Inventory

Find all deferred statuses, TODOs, placeholder specs, blocked gates, and future-work markers.

Record:

```text
file
symbol/function/field
deferred item
why deferred
what would retire the debt
whether current reports depend on it
```

Special attention:

```text
Friedmann residual
classical target comparison
null model comparisons beyond diagnostic summaries
observational/empirical benchmarks
Lean theorem obligations beyond policy booleans
```

## 10. Failure-Detection Gap Inventory

Create a prioritized list of missing adversarial tests.

Required entries:

```text
wrong-constraint wavefunction should fail WdW residual
k² != ω² plane wave should fail constraint check
invalid superposition near node should produce explicit obstruction, not fake finite validity
Euler vs RK4 drift should be measured on same profile
dt convergence should be measured
node-threshold sweep should show sensitivity or stability
phase-gradient h-sensitivity should be measured
Q near-node should produce domain-boundary status, not Q=0 validity
global monotonicity should not be the only clock-readiness gate
```

For each, specify:

```text
target package
minimal test fixture
expected failure/pass behavior
whether it should be Sprint P2 or P3
```

## 11. Numerical-Analysis Debt Inventory

Document current numerical weaknesses:

```text
fixed step Euler/RK4
no adaptive integration
no local truncation error estimate
no convergence table
hardcoded phase-gradient step h if present
near-node division sensitivity
NaN/Inf finiteness checks not equal to numerical reliability
```

Do not fix them yet. Inventory only.

## 12. Physics-Faithfulness Debt Inventory

Document unpaid physics debts:

```text
factor ordering not modeled
minisuperspace metric convention not fully audited
classical target/Friedmann equation not implemented as physics comparison
units/dimensions not checked
Q/classical-term ratio not computed
no nontrivial potential-bearing model
no observational benchmark path
full BMC blocked
```

## 13. Validation-Theater Risk Register

Create a risk table with:

```text
risk_id
risk_description
evidence
severity
affected packages/schemas
recommended remediation
```

Include at least:

```text
schema proliferation beyond physics
analytic plane-wave tautology as primary authority
Q near-node clamp risk
classical-limit check independence from real Friedmann target
null comparison becoming decorative
audit stability becoming decorative, now repaired in Sprint 11.1 but still note its history
promotion gates as Go conditionals rather than theorem-backed guarantees
```

## 14. Freeze Recommendations

Recommend what should be frozen.

Expected recommendations:

```text
Freeze new report schemas.
Freeze new audit layers.
Freeze promotion language.
Freeze existing Sprint 1–11 artifacts as historical audit stack.
Allow only remediation tickets that improve failure detection or nontrivial physics.
```

## 15. Next Remediation Tickets

Create concrete tickets:

```text
BMC-POST-0001: Constraint Violation Detection
BMC-POST-0002: Independent Numerical WdW Residual Evaluator
BMC-POST-0003: Euler/RK4 and dt Convergence Audit
BMC-POST-0004: Node Domain Boundary Policy for Q and Velocity
BMC-POST-0005: Phase Gradient h-Sensitivity Audit
BMC-POST-0006: Local Clock-Readiness Gate Integration
BMC-POST-0007: Friedmann Residual Specification Split
BMC-POST-0008: Schema Inventory and Shared Envelope Plan
```

For each ticket include:

```text
title
priority
scope
files likely touched
success criteria
forbidden claims
EBP debts affected
```

## 16. EBP Status Summary

End with:

```text
promotion_status: postmortem_inventory_only
full_bmc_toy_gate: blocked
containsFinalTruthClaim: absent
needFaithfulnessReview: contested
needToyCheck: partial
needNullModel: partial
needNumericalErrorAudit: unpaid
needConstraintViolationTests: unpaid
needNontrivialPhysicsCase: unpaid
```

## Verification

Run:

```bash
GOCACHE=/tmp/go-build-cache go test ./...
cd BMC && /home/chaschel/.elan/bin/lake build
```

Do not require new code tests unless the inventory doc generation itself includes tests.

## Required output

Return a walkthrough with:

```text
files added
files modified
inventory document path
summary of schema count
summary of CLI count
summary of package count
summary of test categories
top validation-theater risks
top missing failure-detection tests
recommended next ticket
go test result
lake build result
remaining limitations
```

## Strict EBP reminder

A postmortem inventory is not remediation.

A schema inventory is not physics progress.

A failure-gap list is not failure detection.

A residual audit is not recovery.

Full BMC remains blocked.
