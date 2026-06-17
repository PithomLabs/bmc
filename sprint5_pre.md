The next step is **Sprint 5 planning**, but **not Friedmann recovery yet**.

Sprint 4 found the key obstruction: in the tested BMC-0A region, the trajectory can remain numerically valid while the φ-clock fails as a monotonic relational clock. The adversarial review confirmed this distinction: `trajectory_valid=pass`, `phi_clock_valid=fail`, and `distinction_preserved=true`. 

So Sprint 5 should be:

```text id="ggvdq9"
BMC Sprint 5: Relational-Clock Handling / Clock-Independent Diagnostic Readiness
```

## Sprint 5 goal

Before using any Friedmann-style residual, we need to answer:

```text id="sjf0t7"
Can BMC-0A define meaningful relational diagnostics when φ is not globally monotonic?
```

Not:

```text id="z4l221"
Can we recover Friedmann yet?
```

## Why this is the right next step

A Friedmann residual usually compares evolution against a time parameter or relational clock. If φ is nonmonotonic in important toy regions, then any Friedmann-style check using φ as the clock may be misleading.

Sprint 5 should therefore build a **clock-readiness layer**:

```text id="xvyn7f"
1. Detect valid monotonic clock intervals.
2. Segment trajectories into locally monotonic φ regions.
3. Test whether α(φ) is meaningful only on those segments.
4. Add clock-independent diagnostics where possible.
5. Keep clock_choice_debt active.
```

## What Sprint 5 should implement

Narrowly:

```text id="nxh82q"
internal/bmc/clockseg/
```

with diagnostics like:

```text id="wjxm0j"
monotonic interval segmentation
turning-point detection
local α(φ) branch extraction
clock-branch validity report
clock-independent trajectory summaries
```

Example report fields:

```json id="pisxlk"
{
  "schema_version": "bmc0a-clock-readiness-v0.1",
  "diagnostic_kind": "relational_clock_readiness",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "global_phi_clock_valid": "fail",
  "local_phi_clock_segments": [],
  "alpha_as_clock_candidate": "contested",
  "clock_independent_diagnostics": [],
  "clock_choice_debt": "active",
  "friedmann_readiness": "blocked",
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked"
  }
}
```

## Success criteria

Sprint 5 succeeds if it can honestly say one of these:

```text id="guj7su"
A. φ fails globally but has usable local monotonic branches.
B. φ fails so badly that φ-based Friedmann checks must remain blocked.
C. α or another variable is a candidate clock, but only as contested diagnostic debt.
D. clock-independent diagnostics are safer for the next phase.
```

Any of these is useful.

## Forbidden in Sprint 5

Do **not** implement yet:

```text id="hinqtp"
Friedmann residual recovery
massive scalar model
LQC comparison
Page-Wootters comparison
full BMC promotion
claim that φ is invalid in full cosmology
claim that α is the correct clock
```

## EBP status going into Sprint 5

```text id="bcy5ca"
needMap: partial
needInvariant: partial
needToyCheck: active
needNullModel: partial/deferred
needObstruction: active, clock obstruction discovered
needFaithfulnessReview: contested
clock_choice_debt: active
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

The next smallest useful move is:

```text id="fpfkn8"
Write the Sprint 5 planning prompt for relational-clock segmentation and clock-readiness diagnostics.
```
