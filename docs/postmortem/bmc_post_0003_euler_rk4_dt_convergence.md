# BMC-POST-0003: Euler/RK4 and dt Convergence Audit

## Summary
BMC-POST-0003 compares numerical trajectories across Euler/RK4 and dt refinements on the same superposition profile. This audit evaluates the self-consistency of trajectory integrations under refinement but does not imply physical correctness.

## Key Principles and Constraints
- **Numerical Self-Consistency Only**: This audit compares numerical self-consistency of toy trajectory integration under stepper and dt refinement; it does not test whether the trajectory is physically correct.
- **Reference Trajectory Convention**: The finest RK4 run (RK4 stepper with dt/4) is used as the numerical comparison baseline (local numerical reference), not physical ground truth.
- **No Physics Validation**: Observing numerical convergence does not validate Bohmian Minisuperspace Cosmology (BMC) physics.
- **No Friedmann Recovery**: Reaching a stable or convergent trajectory does not implement, prove, or suggest Friedmann cosmological recovery.
- **Full BMC remains blocked.**

## Runs and Metrics Schema
For each configuration, we evaluate:
- **`euler_dt`**: Euler stepper with dt = `LambdaStep`
- **`rk4_dt`**: RK4 stepper with dt = `LambdaStep`
- **`rk4_dt_2`**: RK4 stepper with dt = `LambdaStep / 2`
- **`rk4_dt_4`**: RK4 stepper with dt = `LambdaStep / 4` (Reference Run)

For each run, the audit reports:
- Stepper and delta lambda step sizes.
- Trajectory finiteness and node contact flags.
- Final coordinates (`final_alpha`, `final_phi`).
- Endpoint distance to the reference trajectory.
- Pointwise distance to reference aligned by lambda value.

## EBP status expected after implementation

```text
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: partial
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```
