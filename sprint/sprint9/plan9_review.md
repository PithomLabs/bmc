Approve **Sprint 9 implementation plan with important revisions**.

The scope is right: a **BMC-0A null-model runner**, not a residual runner, not a recovery sprint, not a full BMC promotion. But I would tighten a few points before implementation so we do not accidentally generate fake “comparison” results.

```json id="s9_plan_review"
{
  "sprint": "BMC Sprint 9",
  "plan_status": "approved_with_revisions",
  "scope": "null_model_diagnostic_runner_only",
  "allowed": [
    "null-model diagnostic generation",
    "target/null diagnostic comparison records",
    "deterministic seeded null fixtures",
    "blocked/deferred status when a null cannot be honestly run"
  ],
  "forbidden": [
    "Friedmann residual computation",
    "recovery claim",
    "BMC beats null models claim",
    "winner language",
    "full BMC promotion",
    "scientific novelty claim"
  ]
}
```

## Required revisions before implementation

### 1. Replace “numeric and status simulation” wording

This line is risky:

```text id="x2zvl9"
Implements numeric and status simulation runs for the null models.
```

Use this instead:

```text id="m5y4bw"
Implements deterministic null diagnostic fixtures and real diagnostic calls where available. If a null model cannot be honestly evaluated with existing BMC diagnostics, mark it blocked or deferred with a reason.
```

Do **not** fabricate diagnostic numbers just to populate the report.

### 2. Add provenance for every diagnostic

Each `NullModelRun` should say whether diagnostics came from a real computation, deterministic fixture, inherited source artifact, or blocked path.

Add:

```go id="9pzj2l"
DiagnosticProvenance string `json:"diagnostic_provenance"`
```

Allowed values:

```text id="ore8u7"
computed_from_existing_bmc_diagnostics
deterministic_fixture
source_artifact_summary
blocked
deferred
```

Forbidden values:

```text id="hwhkcz"
assumed
fabricated
inferred_success
validated_physics
```

### 3. Do not force comparison if no comparable diagnostics exist

The original plan says:

```text id="eqqwj0"
target_null_comparison_computed = true
```

That is okay only if at least one target/null diagnostic comparison is honestly computed.

Add a safety rule:

```text id="el8t69"
If no comparable null diagnostics exist, target_null_comparison_computed must be false and the report status must be blocked_by_no_comparable_null_diagnostics.
```

But if the current implementation can generate at least one honest deterministic fixture comparison, then `target_null_comparison_computed = true` is acceptable.

### 4. Treat `classical_frw_reference_trajectory` carefully

This is not a normal null wavefunction. It is more like a reference comparator.

Allow it, but label it conservatively:

```text id="5a8k50"
run_status: diagnostics_generated | blocked | deferred
diagnostic_provenance: deterministic_fixture | source_artifact_summary | blocked
notes: reference comparator only; no residual or recovery interpretation
```

Do not use it to imply classical recovery.

### 5. Keep Lean optional and tiny

`NullRun.lean` is acceptable only if it remains policy-only.

It may prove:

```text id="hr7hrk"
no residual computation
no recovery claim
full BMC blocked
null-run does not imply BMC beats nulls
```

It must not prove:

```text id="nq6kui"
null models pass
target wins
BMC is physically validated
classical limit is recovered
```

## Approved package boundary

Allowed:

```text id="9vz7du"
internal/bmc/nullrun/contracts.go
internal/bmc/nullrun/report.go
internal/bmc/nullrun/validate.go
internal/bmc/nullrun/runner.go
internal/bmc/nullrun/nullrun_test.go
cmd/ptw-bmc run-nullmodels routing
optional tiny BMC/BMC/NullRun.lean
```

Do not add:

```text id="rsay5t"
residual runner
large comparison engine
dashboard
paper claim extraction
general physics profile system
```

## Status after implementation

Before adversarial review:

```text id="umw2wu"
null_model_runner_candidate_only
```

Maximum after implementation, review, and repairs:

```text id="xd7tl4"
promoted_null_model_runner_artifact_after_repairs
```

## EBP status

```text id="b7epoc"
Claim: Sprint 9 may compute null-model diagnostics.
Status: allowed.

Claim: Sprint 9 may compare target/null diagnostics.
Status: allowed only as diagnostic comparison, not winner claim.

Claim: Sprint 9 may compute Friedmann residuals.
Status: forbidden.

Claim: Sprint 9 may say BMC beats null models.
Status: forbidden.

needNullModel: active → may become partial after honest null diagnostics
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
full_bmc_toy_gate: blocked
```

Final instruction:

```text id="8hxr24"
Go implement Sprint 9 with the revisions above: no fake numeric results, explicit diagnostic provenance, blocked/deferred status where needed, no residual computation, no recovery claim, no winner language, and full BMC remains blocked.
```
