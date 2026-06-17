You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 5: Relational-Clock Readiness and Local Clock Segmentation**.

Sprint 5 follows the accepted Sprint 4 result:

```text
diagnostic_outcome: clock_fragile
trajectory_valid: pass
phi_clock_valid: fail
distinction_preserved: true
clock_choice_debt: active
```

Sprint 5 must remain a **relational-clock readiness and local segmentation sprint only**.

It must not implement or imply Friedmann recovery.

## Current accepted context

Accepted artifacts under strict EBP limits:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
```

Sprint 5 claims to add:

```text
internal/bmc/clockseg/
segment-clock CLI subcommand
bmc0a-clock-readiness-v0.1 report schema
local monotonic φ-branch segmentation
turning-point detection
local α(φ) single-valuedness checks
clock-independent diagnostics
ClockReadiness.lean policy/safety gates
```

## Forbidden claims

The artifact must **not** claim:

```text
Friedmann residual recovery
ready for Friedmann recovery
full BMC validation
quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence
valid φ-clock for full cosmology
invalid φ-clock for full cosmology
valid α-clock for full cosmology
LQC comparison completed
Page-Wootters comparison completed
```

The full BMC toy gate must remain blocked.

Friedmann residual must remain deferred.

`clock_choice_debt` must remain active.

Faithfulness for full BMC remains contested unless separate human review occurred.

## Materials to review

Review actual source and generated artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
internal/bmc/clockseg/**
BMC/**
BMC/BMC/ClockReadiness.lean
out/bmc0a_clock_readiness.json
out/bmc0a_clock_fragility.json
out/bmc0a_superposition_robustness.json
go.mod
all Go tests
Sprint 1 final report
Sprint 2 final report
Sprint 3 final report
Sprint 4 final report
Sprint 5 walkthrough/report
```

If any expected file or artifact is missing, report it as debt.

## Expected Sprint 5 design

Expected package:

```text
internal/bmc/clockseg/
```

Expected files:

```text
branches.go
segments.go
turning_points.go
local_relations.go
clock_independent.go
report.go
validate.go
clockseg_test.go
```

Expected CLI:

```bash
ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
ptw-bmc validate --report out/bmc0a_clock_readiness.json
ptw-bmc summarize --report out/bmc0a_clock_readiness.json
```

Expected schema:

```text
bmc0a-clock-readiness-v0.1
```

Expected Lean file:

```text
BMC/BMC/ClockReadiness.lean
```

## Central review principle

Sprint 5 may identify **local relational branches**, but:

```text
local branch ≠ global clock
local_only_candidate ≠ Friedmann recovery
clock-independent diagnostic ≠ physics proof
```

The key review question is:

```text
Does Sprint 5 safely prepare local-clock diagnostics without claiming readiness for Friedmann recovery?
```

Mark as blocker if the report or code treats `local_only_candidate` as “ready for Friedmann recovery.”

## Required review targets

### 1. `local_only_candidate` wording and validation

Review all report text, warnings, summaries, tests, comments, and CLI output.

Correct behavior:

```text
friedmann_readiness = local_only_candidate
```

may only mean:

```text
local φ branches exist and can be inspected safely
```

It must not mean:

```text
ready for Friedmann recovery
Friedmann residual can now be computed
Friedmann debt is retired
```

Check that the report includes a warning like:

```text
local_only_candidate is not Friedmann recovery readiness
```

Check that validation rejects missing warning if `friedmann_readiness = local_only_candidate`.

### 2. Local branch segmentation

Review `segments.go`.

Check whether:

```text
segments are deterministic
segments are non-overlapping
segments are ordered by index/lambda
segment start_index <= end_index
segment lambda ranges are finite
each segment has direction: increasing|decreasing|flat|mixed
each segment has monotonic_status: pass|fail|contested
invalid/too-short segments are marked honestly
```

Check that segmentation is not silently applied only to the easiest safe baseline if Sprint 5 claims readiness over fragile configurations.

### 3. Fragile configuration coverage

Sprint 5 should account for the four fragile Sprint 4 configurations:

```text
c2 = 0.50, k2 = 2.1, omega2 = -2.1
c2 = 0.55, k2 = 1.9, omega2 = -1.9
c2 = 0.55, k2 = 2.0, omega2 = -2.0
c2 = 0.55, k2 = 2.1, omega2 = -2.1
```

With:

```text
T = 10.0
dt = 0.05, 0.025, 0.0125
steps = 200, 400, 800
```

Expected:

```text
4 configs × 3 step sizes = 12 branch-audit runs
```

Review whether:

```text
all 12 are present
branch counts are reported per run
turning points are reported per run
local relation branches are reported per run or clearly aggregated
branch stability under refinement is reported honestly
```

If the summary only reports one safe baseline segment, verify that the full JSON still contains the 12 fragile branch audits. If not, mark as required repair.

### 4. Turning-point detection

Review `turning_points.go`.

Check that turning points use the parameterized threshold:

```text
near_zero_dphi_threshold = 1e-10
```

or a reported/configured value, not a hidden constant.

Check whether turning points are:

```text
sorted by index/lambda
associated with sign change / near-zero dφ / direction reversal / branch boundary
linked to low R, high Q, or high phase gradient only descriptively
```

No causal proof should be claimed.

### 5. Local α(φ) branch extraction

Review `local_relations.go`.

Expected parameters:

```text
single_valuedness_epsilon = 1e-9
min_branch_samples = 3
```

Check whether:

```text
epsilon is reported
min_branch_samples is reported
too-short branches fail or are contested, not passed
single-valuedness is evaluated only on local monotonic branches
α(φ) is not evaluated across global nonmonotonic φ loops
branch_status values are valid and honest
```

Mark as blocker if a global `α(φ)` relation is claimed when φ is nonmonotonic globally.

### 6. Clock-independent diagnostics

Review `clock_independent.go`.

Expected diagnostics include:

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

Check whether these diagnostics genuinely do not require a global φ-clock.

Check optional numeric handling:

```text
unavailable metric = null
unavailable metric has explicit status/reason
NaN/Inf rejected
bare sentinels not used
```

### 7. Readiness classification

Expected values:

```text
readiness_outcome = local_only_candidate|blocked|contested
friedmann_readiness = blocked|local_only_candidate|contested
```

Review whether:

```text
local_only_candidate requires valid local relation branches
blocked is used if no safe local relational handling exists
contested is used if branch behavior is unstable
full_bmc_toy_gate remains blocked
```

Mark as blocker if report uses:

```text
ready_for_friedmann_recovery
```

or equivalent language.

### 8. Strict validation

Review `ValidateClockReadinessReport`.

It should reject:

```text
final_truth_claim = true
toy_analysis_only = false
missing/wrong schema_version
wrong readiness_kind
missing near_zero_dphi_threshold
near_zero_dphi_threshold <= 0
missing single_valuedness_epsilon
single_valuedness_epsilon <= 0
missing min_branch_samples
min_branch_samples < 3
promotion_gate.status != blocked
missing technical_gate status/reason
invalid readiness_outcome values
invalid friedmann_readiness values
clock_choice_debt != active
nonfinite numeric metrics
unavailable numeric values without status/reason
empty step_refinement_branch_audit
empty clock_segments when readiness_outcome = local_only_candidate
friedmann_readiness = local_only_candidate with no valid local relation branches
unknown JSON fields
```

Check for strict decoding, for example:

```go
json.Decoder.DisallowUnknownFields()
```

### 9. Deterministic JSON

Check whether repeated generation produces byte-identical JSON.

Review deterministic ordering for:

```text
source_artifacts
step_refinement_branch_audit
clock_segments
turning_points
local_relation_branches
clock_independent_diagnostics
warnings
```

### 10. CLI behavior

Run or review reported runs:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
./ptw-bmc validate --report out/bmc0a_clock_readiness.json
./ptw-bmc summarize --report out/bmc0a_clock_readiness.json
cd BMC && lake build
```

Check whether:

```text
segment-clock subcommand exists
unknown segment-clock profiles fail safely
validate routes readiness schema to clockseg validator
summarize routes readiness schema to clockseg summarizer
ordinary BMC, robustness, and clock-fragility validators are not weakened
errors are explicit
```

### 11. Lean safety contracts

Review:

```text
BMC/BMC/ClockReadiness.lean
BMC/BMC.lean
```

Check that:

```text
lake build succeeds
no sorry/admit exists
BMCClockReadinessReport exists
clock_readiness_requires_toy_only exists
clock_readiness_blocks_final_truth exists
clock_readiness_requires_friedmann_deferred exists
clock_readiness_does_not_imply_full_bmc exists
clock_readiness_keeps_clock_choice_debt_active exists
clock_readiness_local_candidate_does_not_mean_friedmann_recovered exists
witness tests exist
Lean remains policy/safety only
no theorem claims relational-clock physics or Friedmann recovery
```

### 12. Overclaim scan

Search source, tests, comments, report text, CLI output, summary, README, and walkthrough for forbidden phrases or implications:

```text
ready for Friedmann recovery
recovers Friedmann
Friedmann residual passes
validates full BMC
validates quantum gravity
solves the problem of time
proves Bohmian mechanics
derives spacetime
valid φ-clock for full cosmology
valid α-clock for full cosmology
global clock solved
```

Acceptable phrasing:

```text
local_only_candidate
local branch candidate
clock-readiness diagnostic
local-only relational diagnostic
not Friedmann recovery readiness
clock_choice_debt remains active
```

## Code review checks

Review these details:

```text
clockseg package structure
parameterized near_zero_dphi_threshold
parameterized single_valuedness_epsilon
parameterized min_branch_samples
strict schema validation
computed readiness_outcome
computed technical_gate
finite metric validation
optional metric status/reason handling
no bare sentinel values
local branch segmentation correctness
branch non-overlap and ordering
turning point ordering
local α(φ) branch single-valuedness
clock-independent diagnostics independence from global φ
12 fragile step-refinement branch audits
deterministic output ordering
CLI routing by schema version
full_bmc_toy_gate blocked
Friedmann residual deferred
clock_choice_debt active
no external dependencies added
```

## EBP debt classification

Classify each item as one of:

```text
unpaid
partial
retired
contested
overclaimed
absent
```

Debt items:

```text
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
clock_choice_debt
containsFinalTruthClaim
LeanVerification
ClockReadinessDiagnosticIntegrity
LocalBranchValidity
FriedmannReadinessBoundary
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
    "needMap": "",
    "needInvariant": "",
    "needToyCheck": "",
    "needNullModel": "",
    "needObstruction": "",
    "needFaithfulnessReview": "",
    "clock_choice_debt": "",
    "containsFinalTruthClaim": "",
    "LeanVerification": "",
    "ClockReadinessDiagnosticIntegrity": "",
    "LocalBranchValidity": "",
    "FriedmannReadinessBoundary": ""
  },
  "readiness_findings": [],
  "physics_boundary_findings": [],
  "code_findings": [],
  "cli_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|clock_readiness_candidate_only|promoted_clock_readiness_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 5 passes perfectly, the maximum allowed recommendation is:

```text
promoted_clock_readiness_artifact_after_repairs
```

Never recommend promotion as:

```text
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann recovery
ready for Friedmann recovery
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Remember:

```text
A local relational branch is not a global clock.
A local-only candidate is not Friedmann recovery.
Clock-independent diagnostics are readiness scaffolding, not physics proof.
Clock-choice debt remains active.
```
