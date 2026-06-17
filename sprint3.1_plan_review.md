Approve the revised Sprint 3 plan. It is now implementation-ready, with a few final guardrails to prevent false “robustness” claims.

## Final required guardrails

### 1. Split “audit integrity” from “robustness outcome”

Do not make `technical_gate = pass` mean “the model is robust in every sense.”

Make it mean:

```text
The audit ran correctly, reported all sweep outcomes honestly, hid no NaN/Inf, preserved EBP boundaries, and kept full BMC blocked.
```

Then include a separate field:

```json
"robustness_outcome": "stable|mixed|fragile|contested"
```

This matters because a good audit may discover fragility. Finding fragility is not a failed audit.

### 2. Use strict JSON decoding

For `ValidateRobustnessReport()`, make sure Go uses strict decoding, for example with `json.Decoder.DisallowUnknownFields()`. Otherwise “rejects unknown fields” will not actually happen by default.

### 3. Clarify threshold-test language

This line:

```text
TestThresholdSensitivity: safe profiles pass and node-probe profiles fail under all thresholds.
```

Should become:

```text
safe profile remains node-contact-free under all thresholds;
node-probe remains correctly classified as node-contact/blocked or near-node-risk depending on threshold;
node_detection_validation_gate passes when obstruction behavior is correctly reported.
```

Node-probe “fails” only as a trajectory/safe-gate artifact. It should pass as an obstruction-detection validation artifact.

### 4. Add deterministic ordering

All sweep arrays should be emitted in fixed order:

```text
step sizes: descending or ascending, but fixed
thresholds: fixed order
phase bounds: fixed order
perturbation grid: deterministic nested-loop order
node offsets: fixed order
```

This preserves byte-stable JSON and makes adversarial review easier.

## Approved status

```json
{
  "sprint": "BMC Sprint 3",
  "plan_status": "approved_for_implementation",
  "scope": "numerical_robustness_convergence_audit_only",
  "maximum_promotion": "promoted_robustness_audit_artifact_after_repairs",
  "full_bmc_toy_gate": "blocked",
  "friedmann_residual": "deferred",
  "faithfulness": "contested"
}
```

## EBP debt status before implementation

```text
needMap: partial
needInvariant: partial
needToyCheck: active
needNullModel: partial/deferred
needObstruction: active
needFaithfulnessReview: contested
containsFinalTruthClaim: must remain absent
LeanVerification: planned
```

## Final instruction to coding agent

Proceed with Sprint 3 implementation using the revised plan, plus these guards:

```text
1. Add robustness_outcome separate from technical_gate.
2. Make technical_gate report audit integrity, not final robustness truth.
3. Use strict JSON validation with unknown-field rejection.
4. Preserve deterministic sweep ordering.
5. Keep node-probe classified as obstruction validation, not successful trajectory simulation.
6. Keep full_bmc_toy_gate blocked.
7. Do not implement new physics.
```

After implementation, run the full command list and then send Sprint 3 to adversarial review before accepting it.
