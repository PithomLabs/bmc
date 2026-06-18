Accepted. **BMC-POST-0002 is now cleanly closed after POST-0002.1 repairs.**

```json id="post0002_final_acceptance"
{
  "ticket": "BMC-POST-0002",
  "repair_ticket": "BMC-POST-0002.1",
  "final_status": "accepted_for_plane_wave_report_path_scope",
  "artifact_freshness": "repaired",
  "validation_consistency": "repaired",
  "go_wdw_tests": "passed",
  "go_report_tests": "passed",
  "all_go_tests": "passed",
  "lean_build": "passed",
  "physics_claim_promoted": false,
  "full_bmc_toy_gate": "blocked"
}
```

What this earns:

```text id="0hcpm6"
The live plane-wave report path now uses numerical WdW residual diagnostics as authority, stale generated artifacts have been refreshed, and validation now rejects both violation/error numerical residual states if falsely marked as pass.
```

What it still does **not** earn:

```text id="uxrf8b"
full WdW solver
superposition numerical WdW authority
Friedmann recovery
classical-limit recovery
BMC validation
null-model failure
BMC superiority
scientific novelty
full BMC promotion
```

Updated EBP status:

```text id="qckmzw"
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: partial
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

Now we can proceed to:

```text id="ephsiy"
BMC-POST-0003: Euler/RK4 and dt Convergence Audit
```

This is the right next repair because the next unresolved risk is numerical integration reliability on identical superposition profiles.
