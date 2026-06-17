Approved. This revised Sprint 1 plan is now implementation-ready.

It correctly incorporates the key fixes:

```text
split-gate architecture
BMC-0A control gate passable
full BMC toy gate blocked
Friedmann residual deferred
CheckStatus used instead of plain booleans
analytic plane-wave residual as primary authority
zero-dependency Go CLI
minimal Lean without Mathlib
Sprint 1 witness theorem instead of a false universal theorem
```

This is the right EBP 2.1 boundary: Sprint 1 may pass the **plane-wave control gate**, but it must remain unpromoted as a full BMC toy artifact because Friedmann residual and faithfulness review remain unpaid debt. 

Two tiny cleanup notes before the coding agent proceeds:

1. In the JSON reference, the obstruction example still shows:

```json
{ "name": "full_qg_overclaim_blocker", "applies": true, ... }
```

For the clean Sprint 1 report, that should be:

```json
{ "name": "full_qg_overclaim_blocker", "applies": false, "severity": "blocker", "status": "pass" }
```

`applies: true` only when final-truth/overclaim language is detected.

2. The file summary says “7 test functions,” but the revised test list now has more than 7. This is harmless, but update the summary to say:

```text
Tests: plane-wave residual, Q≈0, finite trajectory, clock monotonicity, report determinism, final-truth rejection, toy-only validation, invalid params, constraint violation, invalid steps, nonfinite params.
```

With those minor wording fixes, tell the coding agent:

```text
Proceed with Sprint 1 implementation. No adversarial LLM review is needed before coding. The adversarial review happens after runnable Go code, Lean files, generated JSON, test results, and the implementation report exist.
```
