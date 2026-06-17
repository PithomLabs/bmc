## prompt

the coding agent is still in the implementation plan stage and just asking clarification question



Tell the coding agent this:

Clarification: split the gates. Do not make Sprint 1 pass the full BMC toy gate.

For Sprint 1, implement only:

```text
BMC-0A Plane-Wave Control Gate
```

This gate may pass if:

```text
toy_analysis_only = true
final_truth_claim = false
WdW residual passes
trajectory is finite
φ is monotonic for the selected parameters
Q is finite
Q ≈ 0
classical control behavior passes
friedmann_residual is explicitly deferred
faithfulness is explicitly contested/deferred
```

Do **not** require `friedmannResidualPass = true` for Sprint 1.

Instead of using plain booleans for all checks, use a status enum/string such as:

```text
pass
fail
deferred
contested
```

In the JSON report, Friedmann should be:

```json
"friedmann_residual": {
  "status": "deferred",
  "implemented": false,
  "reason": "Not part of BMC-0A plane-wave control; remains debt for the full BMC toy gate."
}
```

Add two separate gate sections:

```json
"technical_gate": {
  "name": "bmc0a_plane_control_gate",
  "status": "pass"
},
"promotion_gate": {
  "name": "full_bmc_toy_gate",
  "status": "blocked",
  "reason": "Friedmann residual and faithfulness review remain unpaid debt."
}
```

In Lean, define two gates:

```lean
reportPassesBMC0AControlGate
reportPassesFullBMCToyGate
```

`reportPassesBMC0AControlGate` should require `friedmannResidual = deferred`.

`reportPassesFullBMCToyGate` should require `friedmannResidual = pass`.

Sprint 1 may pass the control gate but must remain blocked from the full toy promotion gate.

This is an EBP 2.1 promotion-boundary rule: passing the plane-wave control retires only a narrow part of `needToyCheck`. It does not retire the Friedmann invariant, full toy-check debt, or faithfulness debt.

Proceed with implementation using this split-gate design. Do not implement the massive scalar solver, packet states, LQC comparison, Page-Wootters comparison, or full Friedmann residual yet.

That is the right clarification. The coding agent should not wait for another design choice here; the contradiction is resolved by **separating the control gate from the promotion gate**.
