Verdict: **approve the plan with small corrections before implementation.** It correctly applies the split-gate architecture: `bmc0a_plane_control_gate` can pass in Sprint 1, while `full_bmc_toy_gate` remains blocked until Friedmann residual and faithfulness review are paid. That resolves the earlier contradiction. 

## What is strong

The plan is EBP-compatible because it keeps Sprint 1 narrow: plane-wave control only, zero external Go dependencies, minimal Lean without Mathlib, deterministic JSON, and explicit non-claims. It also correctly defers Friedmann residual instead of faking it, and uses `CheckStatus = pass | fail | deferred | contested` rather than ambiguous booleans. 

The Go/Lean split is also right:

```text
Go: execute the finite toy check.
Lean: enforce promotion-safety contracts.
```

That is exactly the boundary we want.

## Required corrections before coding

### 1. Fix `PlaneWaveParams` type

The plan says:

```go
PlaneWaveParams struct: K, Omega, Alpha0, Phi0, LambdaStep, Steps int, Tolerance float64
```

That is wrong because `K`, `Omega`, `Alpha0`, `Phi0`, and `LambdaStep` must be `float64`; only `Steps` should be `int`.

Use:

```go
type PlaneWaveParams struct {
	K          float64
	Omega     float64
	Alpha0    float64
	Phi0      float64
	LambdaStep float64
	Steps     int
	Tolerance float64
}
```

### 2. Do not rely on finite differences for the plane-wave residual test only

The plan says `Residual` computes finite differences, while also noting the analytic result is `(k² - ω²)Ψ`. For Sprint 1, implement both:

```text
AnalyticResidualPlaneWave
FiniteDifferenceResidual
```

The **primary pass/fail gate** should use the analytic residual for the plane-wave control. Finite difference can be reported as a numerical sanity check, but do not let floating finite-difference error become the main authority over an exactly solvable control case.

### 3. Avoid `Pass bool` becoming a second truth source

The plan uses both:

```json
"status": "pass",
"pass": true
```

That is acceptable for user readability, but validation must enforce consistency:

```text
status = "pass" iff pass = true
status in {"fail","deferred","contested"} iff pass = false
```

Otherwise we reintroduce ambiguity.

### 4. Fix obstruction wording: overclaim blocker should not “always apply” as a blocker

The plan says:

```text
DetectOverclaimBlocker(finalTruthClaim bool) — always applies as blocker
```

That is too strong. It should always be **present**, but only **applies as blocker** when `finalTruthClaim = true` or when forbidden language is detected.

For a clean Sprint 1 report:

```json
{
  "name": "full_qg_overclaim_blocker",
  "applies": false,
  "severity": "blocker",
  "status": "pass",
  "evidence": "No final-truth language detected.",
  "consequence": "Continue as toy-only artifact."
}
```

### 5. Add invalid-parameter tests

The test list is good, but add:

```text
TestValidateRejectsConstraintViolation
TestValidateRejectsInvalidStepCount
TestValidateRejectsNonfiniteParameters
TestClockMonotonicityFailsWhenOmegaZero
```

Especially `omega = 0` matters because `φ` will not be a usable clock.

### 6. Lean theorem `control_gate_does_not_imply_full_gate` should be a concrete witness

Do not try to prove this universally, because it is not true universally. A control-gate-passing report *could* also pass the full gate in a later sprint if Friedmann and faithfulness are passed, but in Sprint 1 we want a witness showing separation.

Use a theorem like:

```lean
def sprint1Witness : BMCReport := {
  toyAnalysisOnly := true,
  finalTruthClaim := false,
  wdwResidual := CheckStatus.pass,
  trajectoryFinite := CheckStatus.pass,
  clockMonotonic := CheckStatus.pass,
  qFinite := CheckStatus.pass,
  classicalLimit := CheckStatus.pass,
  friedmannResidual := CheckStatus.deferred,
  faithfulness := CheckStatus.contested
}

theorem sprint1_witness_passes_control :
  reportPassesBMC0AControlGate sprint1Witness = true := by
  decide

theorem sprint1_witness_fails_full_gate :
  reportPassesFullBMCToyGate sprint1Witness = false := by
  decide
```

That is cleaner than a misleading universal statement.

## EBP 2.1 ledger for the plan

```text
needMap: partial
needInvariant: partial
needToyCheck: active, narrowed to plane-wave control
needNullModel: partial placeholders only
needObstruction: active, with four Sprint-1 detectors and deferred placeholders
needFaithfulnessReview: active/contested
containsFinalTruthClaim: blocked by validation
promotion_status: alive_unpromoted_control_gate_only
```

## Approval note to send coding agent

You can send this:

```text
Plan approved with corrections:

1. Fix PlaneWaveParams: all numeric model parameters except Steps are float64.
2. Implement analytic plane-wave residual as the primary Sprint 1 check; finite-difference residual may be included only as a sanity check.
3. Keep CheckStatus as the authority; if `pass` bool is included in JSON, validation must enforce consistency with status.
4. Overclaim blocker should always be present but applies=false when no final-truth claim is detected; applies=true only when overclaim/final-truth language exists.
5. Add invalid-parameter tests: constraint violation, nonfinite params, invalid step count, and omega=0 clock failure.
6. In Lean, make `control_gate_does_not_imply_full_gate` a Sprint 1 witness theorem, not a universal theorem.

Proceed with Sprint 1 only. Do not implement packet states, RK4, massive scalar, Friedmann residual, LQC, Page-Wootters, dashboards, or full QG claims.
```

With those fixes, the plan is ready for implementation.
