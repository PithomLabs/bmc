# BMC Sprint 5 Walkthrough: Relational-Clock Readiness and Local Clock Segmentation

Sprint 5 has been successfully implemented, verified, and proven under EBP 2.1 discipline. 

## Summary of Changes

### 1. Go Implementation (`internal/bmc/clockseg`)
- **[NEW] [branches.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/branches.go)**: Defines the core data models for `ClockSegment`, `ClockTurningPoint`, `LocalRelationBranch`, and `ClockIndependentDiagnostic`.
- **[NEW] [segments.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/segments.go)**: Implements dynamic segmentation of trajectory points into strictly monotonic $\phi$-clock intervals.
- **[NEW] [turning_points.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/turning_points.go)**: Provides sorting and ordering verification helpers for turning points.
- **[NEW] [local_relations.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/local_relations.go)**: Extracts relational branch properties and runs noise-tolerant $\alpha(\phi)$ single-valuedness checks using `single_valuedness_epsilon = 1e-9` and `min_branch_samples = 3`.
- **[NEW] [clock_independent.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/clock_independent.go)**: Computes the 10 explicit clock-independent diagnostics:
  1. `path_length_in_configuration_space`
  2. `total_lambda_interval`
  3. `num_valid_trajectory_points`
  4. `num_clock_segments`
  5. `num_turning_points`
  6. `min_amplitude_r`
  7. `max_abs_q_away_from_nodes`
  8. `max_phase_gradient`
  9. `node_contact_free`
  10. `trajectory_finiteness`
- **[NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/report.go)**: Generates the deterministic JSON report for the `bmc0a-clock-readiness-v0.1` schema, including step refinement sweeps at `0.05`, `0.025`, and `0.0125` for the four fragile configurations.
- **[NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/validate.go)**: Enforces strict EBP boundaries:
  - Rejects `friedmann_readiness = local_only_candidate` if no valid branches exist.
  - Ensures warnings include `"local_only_candidate is not Friedmann recovery readiness"`.
  - Ensures segments are deterministic, ordered, non-overlapping, and have finite ranges.
  - Ensures turning points are strictly ordered by index/lambda.
  - Blocks promotion gates and requires `sprint5_clock_readiness` to be `planned_candidate_only`.
- **[NEW] [clockseg_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/clockseg_test.go)**: Implements 14 unit tests checking all requirements, including:
  - `TestClockReadinessRejectsReadyForFriedmannLanguage`
  - `TestClockSegmentsAreNonOverlappingAndOrdered`
  - `TestLocalRelationRejectsTooFewSamples`
  - `TestLocalOnlyCandidateDoesNotUnblockFullBMC`

### 2. Command Line Interface (`cmd/ptw-bmc`)
- **[MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)**:
  - Added `segment-clock` subcommand.
  - Wired `validate` to verify files against the `bmc0a-clock-readiness-v0.1` schema.
  - Wired `summarize` to display a clean, human-readable summary of the readiness report.

### 3. Formal Policy Verification (`BMC`)
- **[NEW] [ClockReadiness.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ClockReadiness.lean)**: Formally defines `BMCClockReadinessReport` and verifies the EBP safety theorems:
  - `clock_readiness_requires_toy_only`
  - `clock_readiness_blocks_final_truth`
  - `clock_readiness_requires_friedmann_deferred`
  - `clock_readiness_does_not_imply_full_bmc`
  - `clock_readiness_keeps_clock_choice_debt_active`
  - `clock_readiness_local_candidate_does_not_mean_friedmann_recovered`
- **[MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)**: Imports `BMC.ClockReadiness`.

---

## Verification and Build Outcomes

### 1. Go Unit Tests
Run: `go test ./...`
```text
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg 0.120s
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```

### 2. Lean Policy Safety Proofs
Run: `lake build` in `BMC/`
```text
Build completed successfully (8 jobs).
```
All safety proofs verify with zero warnings.

### 3. CLI Demonstration
Run:
```bash
./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
./ptw-bmc validate --report out/bmc0a_clock_readiness.json
./ptw-bmc summarize --report out/bmc0a_clock_readiness.json
```
Output:
```text
Successfully ran clock readiness profile 'bmc0a-clock-readiness' and generated report: out/bmc0a_clock_readiness.json
Clock Readiness Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
================================================================================
BMC Sprint 5 Clock Readiness Report Summary
================================================================================
Schema Version:            bmc0a-clock-readiness-v0.1
Friedmann Readiness:       local_only_candidate
Toy Analysis Only:         true
Final Truth Claim:         false
Single Valuedness Epsilon: 1.000000e-09
Min Branch Samples:        3
--------------------------------------------------------------------------------
Technical Gate:            bmc0a_clock_readiness_gate (Status: pass)
Reason:                    Local clock segmentation and clock-independent diagnostics completed successfully.
--------------------------------------------------------------------------------
Promotion Gate:            full_bmc_gate (Status: blocked)
Reason:                    Sprint 5 clock-readiness promotion status is planned_candidate_only. Full BMC promotion remains blocked.
--------------------------------------------------------------------------------
Null Models:
  - No new null models planned. Existing null-model debt remains partial/deferred.
--------------------------------------------------------------------------------
Clock-Independent Diagnostics:
  Path Length in Conf Space: 14.246828
  Total Lambda Interval:     10.000000
  Valid Trajectory Points:   201
  Clock Segments Count:      1
  Turning Points Count:      0
  Min Amplitude R:           0.501533
  Max Abs Q (away nodes):    7.890430
  Max Phase Gradient:        3.972534
  Node Contact Free:         true
  Trajectory Finiteness:     true
--------------------------------------------------------------------------------
EBP Debt Ledger:
  needMap:                 partial
  needInvariant:           partial
  needToyCheck:            active
  needNullModel:           partial/deferred
  needObstruction:         active
  needFaithfulnessReview:  contested
  clock_choice_debt:       active
  containsFinalTruthClaim: absent
  LeanVerification:        planned
  promotion_status:        planned_candidate_only
  sprint5_clock_readiness:  planned_candidate_only
--------------------------------------------------------------------------------
Sweeps & Detections:
  Total clock turning points:   0
  Total clock segments:         1
  Total local relation branches:1
  Step Refinement Audit runs:   12
================================================================================
```
The generated report file is successfully located at: [bmc0a_clock_readiness.json](file:///home/chaschel/Documents/go/bmc/out/bmc0a_clock_readiness.json).
