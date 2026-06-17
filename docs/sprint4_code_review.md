You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 4: Clock-Monotonicity Fragility Investigation**.

Sprint 4 investigates why the φ-clock monotonicity check failed in four of nine safe-superposition parameter perturbations identified in Sprint 3.

This is an **obstruction diagnosis sprint**, not a theory-promotion sprint.

## Current accepted context

The project already has these accepted artifacts under strict EBP limits:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
```

Sprint 3 found:

```text
technical_gate: pass
robustness_outcome: mixed
reason: four of nine parameter perturbation runs failed due to clock-monotonicity checks
```

Sprint 4 should diagnose that fragility.

## Forbidden claims

The artifact must **not** claim:

```text
Bohmian mechanics is proven
quantum gravity is solved
the problem of time is solved
spacetime emergence is proven
Friedmann recovery is achieved
full BMC is validated
φ is proven valid or invalid as a physical clock in full quantum cosmology
α is proven to be the correct replacement clock
black holes are handled
fermions are handled
gauge fields are handled
Lorentz recovery is achieved
LQC comparison is completed
Page-Wootters comparison is completed
```

The full BMC toy gate must remain blocked.

Friedmann residual must remain deferred.

`clock_choice_debt` must remain active.

Faithfulness for full BMC remains contested unless a separate human faithfulness review occurred.

## Materials to review

Review actual source and generated artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
internal/bmc/clockdiag/**
BMC/**
BMC/BMC/ClockFragility.lean
out/bmc0a_clock_fragility.json
out/bmc0a_superposition_robustness.json
go.mod
all Go tests
Sprint 1 final report
Sprint 2 final report
Sprint 3 final report
Sprint 4 walkthrough/report
```

If any expected file or artifact is missing, report it as debt.

## Expected Sprint 4 design

Sprint 4 should add a new package:

```text
internal/bmc/clockdiag/
```

Expected files:

```text
events.go
correlation.go
diagnose.go
report.go
validate.go
clockdiag_test.go
```

Expected CLI command:

```bash
ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json
```

Validation and summary should route by schema version:

```text
bmc0a-clock-fragility-v0.1
```

Expected Lean file:

```text
BMC/BMC/ClockFragility.lean
```

## Central review principle

The most important Sprint 4 distinction is:

```text
valid trajectory with bad/nonmonotonic φ-clock
```

versus:

```text
invalid trajectory due to node contact, nonfinite velocity, NaN/Inf, or failed integration
```

A nonmonotonic φ-clock does **not** automatically invalidate the Bohmian trajectory. It may mean only that φ is a poor relational clock in that parameter region.

Mark as a blocker if the implementation collapses these two cases into one undifferentiated failure.

## Required review targets

### 1. Clock event detection

Review `DetectClockEvents`.

It should detect events such as:

```text
dφ/dλ sign change
|dφ/dλ| near zero
φ direction reversal
monotonicity failure
```

Check whether:

```text
near_zero_dphi_threshold is parameterized and reported
default value is 1e-10
finite differences are computed consistently
events include α, φ, λ/index, dφ/dλ, dα/dλ
events include amplitude R
events include Q status/reason
events include phase-gradient status/reason
near-node status is recorded
```

If the threshold is hard-coded without being reported, mark as repair.

### 2. Rechecks of failed Sprint 3 perturbations

Sprint 4 should re-run the four failed Sprint 3 perturbations:

```text
c2 = 0.50, k2 = 2.1, omega2 = -2.1
c2 = 0.55, k2 = 1.9, omega2 = -1.9
c2 = 0.55, k2 = 2.0, omega2 = -2.0
c2 = 0.55, k2 = 2.1, omega2 = -2.1
```

Each should be run at fixed total interval:

```text
T = 10.0
dt = 0.05, 0.025, 0.0125
steps = 200, 400, 800
```

Expected number of results:

```text
4 configs × 3 step sizes = 12 results
```

Review whether:

```text
all 12 results are present
ordering is deterministic
φ monotonicity is reported
α monotonicity is reported
trajectory validity is reported
clock events are recorded
no NaN/Inf is hidden
```

### 3. Diagnostic outcome

Expected possible outcomes:

```text
clock_stable
clock_fragile
mixed
contested
```

Review whether `diagnostic_outcome` is computed from the recheck results, not hard-coded.

If the report says `clock_fragile`, check that the evidence supports persistence under refinement.

If the report says `clock_stable`, check that this means the prior failures disappear under refinement and are reported as likely numerical artifacts.

If the report says `mixed`, check that some configs stabilize while others remain fragile.

### 4. Alternative clock summary

Review whether Sprint 4 checks α only as an **alternative clock candidate**, not as a replacement or solution.

Correct behavior:

```text
φ_monotonic: pass|fail|contested
α_monotonic: pass|fail|contested
both_monotonic: true|false
neither_monotonic: true|false
clock_choice_debt: active
```

Incorrect behavior:

```text
α is declared the correct clock for cosmology
φ is declared invalid in full quantum cosmology
clock_choice_debt is retired
```

Mark any such overclaim as a blocker.

### 5. Trajectory validity versus clock validity

Review `TrajectoryValiditySummary`.

It should distinguish:

```text
trajectory_valid: pass|fail|contested
phi_clock_valid: pass|fail|contested
distinction_preserved: true|false
reason: explanatory text
```

The key valid Sprint 4 result may be:

```text
trajectory_valid = pass
phi_clock_valid = fail
distinction_preserved = true
```

This means the path is numerically well-formed, but φ is not a good monotonic clock over that path.

Mark as blocker if the report treats clock failure as automatic trajectory invalidity.

### 6. Correlation summary

Review whether the diagnostic compares clock failures with:

```text
min amplitude R
max |Q| away from nodes
max phase-gradient magnitude
min distance to node threshold
parameter values c2, k2, omega2
clock event count
```

This must be descriptive only.

No statistical or causal claim should be made unless explicitly supported.

Correct wording:

```text
correlated with
associated with
observed alongside
suggests possible relation
```

Incorrect wording:

```text
proves Q causes clock failure
proves near-node behavior causes clock failure
```

### 7. Optional metrics and data integrity

Review all optional numeric metrics.

Correct behavior:

```text
unavailable metric uses null
unavailable metric has explicit status/reason
NaN/Inf rejected
bare sentinel values rejected unless explicitly marked and justified
```

If metrics use undocumented `-1`, `NaN`, `Inf`, or omitted reasons, mark as required repair.

### 8. Strict validation

Review `ValidateClockFragilityReport`.

It should reject:

```text
final_truth_claim = true
toy_analysis_only = false
wrong or missing schema_version
wrong diagnostic_kind
promotion_gate.status != blocked
missing technical_gate status/reason
nonfinite numeric metrics
unavailable numeric values without status/reason
empty failed_perturbation_rechecks
empty clock_events when diagnostic_outcome = clock_fragile
invalid diagnostic_outcome values
clock_choice_debt != active
unknown JSON fields
```

Check that strict decoding uses something like:

```go
json.Decoder.DisallowUnknownFields()
```

### 9. Deterministic JSON

Review whether the generated report is byte-stable across repeated runs.

Check deterministic ordering for:

```text
failed perturbation rechecks
clock events
correlation summary
warnings
source artifacts
```

### 10. CLI behavior

Review:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json
./ptw-bmc validate --report out/bmc0a_clock_fragility.json
./ptw-bmc summarize --report out/bmc0a_clock_fragility.json
cd BMC && lake build
```

Check whether:

```text
diagnose-clock subcommand exists
unknown diagnose-clock profiles fail safely
validate routes clock-fragility schema to clockdiag validator
summarize routes clock-fragility schema to clockdiag summarizer
ordinary BMC and robustness validators are not weakened
errors are explicit
```

### 11. Lean safety contracts

Review:

```text
BMC/BMC/ClockFragility.lean
BMC/BMC.lean
```

Check that:

```text
lake build succeeds
no sorry/admit exists
BMCClockFragilityReport structure exists
reportPassesBMC0AClockFragilityDiagnosticGate exists
reportPassesFullBMCForClockFragility exists
clock_fragility_requires_toy_only exists
clock_fragility_blocks_final_truth exists
clock_fragility_requires_friedmann_deferred exists
clock_fragility_does_not_imply_full_bmc exists
clock_fragility_keeps_clock_choice_debt_active exists
witness tests exist
Lean remains policy/safety only
no theorem claims clock physics or full quantum cosmology
```

### 12. Overclaim scan

Search source, tests, comments, CLI output, report text, README, and walkthrough for forbidden implications:

```text
solves quantum gravity
proves Bohmian mechanics
solves the problem of time
derives spacetime
recovers Friedmann
validates full BMC
validates quantum gravity
proves φ is invalid as clock
proves α is valid clock
permanent limitation
```

The phrase “persistent under tested refinements” is acceptable.

The phrase “permanent limitation” is too strong unless explicitly qualified.

## Code review checks

Review these implementation details:

```text
clockdiag package structure
parameterized near_zero_dphi_threshold
strict schema validation
computed diagnostic_outcome
computed technical_gate
finite metric validation
NaN/Inf rejection
optional metric status/reason handling
trajectory validity versus clock validity separation
step refinement rechecks over exactly 12 runs
all 9 perturbations included in correlation summary
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
ClockFragilityDiagnosticIntegrity
TrajectoryVsClockDistinction
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
    "ClockFragilityDiagnosticIntegrity": "",
    "TrajectoryVsClockDistinction": ""
  },
  "diagnostic_findings": [],
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
  "promotion_recommendation": "do_not_promote|clock_fragility_candidate_only|promoted_clock_fragility_diagnostic_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 4 passes perfectly, the maximum allowed recommendation is:

```text
promoted_clock_fragility_diagnostic_artifact_after_repairs
```

Never recommend promotion as:

```text
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
Friedmann recovery
valid φ-clock for full cosmology
invalid φ-clock for full cosmology
valid α-clock for full cosmology
```

Remember:

```text
A failed φ-clock can be valuable information.
A valid trajectory with a bad clock is different from an invalid trajectory.
Do not hide clock-choice debt.
Do not turn a diagnostic into a proof.
```
