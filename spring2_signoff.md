Good. Based on this repair report, Sprint 2 can now be accepted as a **promoted superposition control artifact after repairs** — with the same strict boundary as before.

```json
{
  "sprint": "BMC Sprint 2",
  "artifact": "BMC-0A two-plane-wave superposition control",
  "acceptance_status": "accepted_after_repairs",
  "promotion_status": {
    "bmc0a_superposition_safe_gate": "promoted_control_artifact_after_repairs",
    "node_detection_validation_gate": "promoted_obstruction_detection_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked"
  },
  "accepted_repairs": [
    "Added explicit safe-profile amplitude sampling test.",
    "Enhanced exact-name node-probe short-circuit and validation-gate test.",
    "Replaced weak RK4 test with non-constant ODE accuracy test.",
    "Preserved node short-circuit guard.",
    "Preserved Lean promotion-safety contracts."
  ],
  "verification_reported": [
    "go test ./... passed",
    "go build -buildvcs=false ./cmd/ptw-bmc passed",
    "safe superposition profile run/validate/summarize passed",
    "node-probe profile run/validate/summarize passed",
    "lake build passed with no sorry/admit"
  ],
  "remaining_blockers": [
    "Friedmann residual remains deferred.",
    "Null-model comparison remains unpaid.",
    "Faithfulness review for full BMC remains contested.",
    "Full BMC, full quantum gravity, Bohmian proof, problem-of-time solution, and spacetime-emergence claims remain forbidden."
  ]
}
```

## EBP debt status

```text
needMap: partial
needInvariant: partial, improved by Q-finite-away-from-nodes checks
needToyCheck: partially retired for BMC-0A superposition control only
needNullModel: partial/deferred
needObstruction: partially retired for node/phase/Q short-circuit behavior
needFaithfulnessReview: contested for full BMC
containsFinalTruthClaim: absent based on reported scan
LeanVerification: retired for current safety contracts
```

One wording caution: for the node-probe profile, keep saying **“Q is undefined/contested at the node and correctly blocked or short-circuited”**, not “Q was computed and found nonfinite.” That distinction matters because the safe behavior is refusing invalid computation, not calculating through `Ψ = 0`.

Final status:

```text
Sprint 1: accepted plane-wave control artifact
Sprint 2: accepted superposition control + node-obstruction artifact
Full BMC-0.1: still unpromoted
```

The next step is now **Sprint 3 planning**, but keep it narrow. I would not jump to full Friedmann recovery yet. The safest next target is a **numerical robustness/convergence sprint** for the accepted superposition artifact: step-size convergence, threshold sensitivity, and trajectory stability under small parameter perturbations.
