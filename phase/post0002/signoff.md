Accepted.

```json id="post0002_acceptance"
{
  "ticket": "BMC-POST-0002",
  "title": "Numerical WdW Residual Integration / Analytic Authority Displacement",
  "status": "accepted",
  "scope": "plane-wave report-path numerical WdW diagnostic authority",
  "wdw_tests": "passed",
  "report_tests": "passed",
  "all_go_tests": "passed",
  "lean_build": "passed",
  "cli_profile_verified": "bmc0a-plane",
  "physics_claim_promoted": false,
  "full_bmc_toy_gate": "blocked"
}
```

## What was earned

This is a real remediation improvement. The important shift is now complete:

```text id="post0002_earned"
The plane-wave report path no longer relies on the analytic residual alone as the WdW authority. The numerical residual diagnostic now controls the WdW pass/fail status, while the analytic residual remains an oracle/control.
```

That means the report pipeline can now carry failure information, not just isolated tests.

## What was not earned

```text id="post0002_not_earned"
full WdW solver
superposition numerical WdW authority
Friedmann recovery
classical-limit recovery
BMC physics validation
null-model failure
BMC superiority
scientific novelty
full BMC promotion
```

The superposition deferral is correct. It is better to explicitly mark it `not_authoritative` than to fake coverage.

## Updated EBP debt status

```text id="post0002_ebp"
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

## Next ticket

Proceed to:

```text id="next_ticket"
BMC-POST-0003: Euler/RK4 and dt Convergence Audit
```

Purpose:

```text id="post0003_goal"
Measure whether trajectory outputs are numerically stable across steppers and time-step refinement on the same superposition profile.
```

This is the right next remediation because POST-0001/0002 handled **constraint violation detection**, and the next major adversarial gap is **integration reliability**.
