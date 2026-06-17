## prompt

Friedmann residual in the toy gate: The sprint plan says friedmann_residual should remain not implemented / debt for Sprint 1. But the Lean reportPassesToyGate requires friedmannResidualPass = true to pass. This is a contradiction — a report that honestly marks Friedmann as unimplemented will always fail the Lean toy gate. How should we handle this?


Yes. That is a real contradiction. We should **not** make Sprint 1 pass a gate that includes `friedmannResidualPass = true` while also saying Friedmann is explicitly deferred.

The EBP-correct fix is to split the gates:

```text
BMC-0A Plane-Wave Control Gate
≠
Full BMC-0.1 Toy Promotion Gate
```

EBP says debt does not kill an idea; unpaid debt blocks promotion, and each session should retire or file one useful piece of debt rather than pretending all debt is paid.  It also says a toy check may retire debt only in the scope actually tested, and faithfulness review can keep an idea unpromoted if the formalization does not yet match the larger claim. 

## Correct handling

For Sprint 1, the Lean gate should **not** be named `reportPassesToyGate` if it requires Friedmann. That name is too broad.

Use two or three gates instead:

```text
reportPassesBMC0AControlGate
```

This checks only:

```text
toy_analysis_only = true
final_truth_claim = false
WdW residual passes
trajectory finite
φ monotonicity passes
Q ≈ 0 passes
Friedmann residual is deferred, not falsely passed
```

Then later:

```text
reportPassesFullBMCToyGate
```

This checks:

```text
all BMC-0A controls
plus Friedmann residual pass
plus faithfulness accepted
plus null-model/obstruction requirements
```

The Kimi BMC draft places the effective Friedmann equation as part of the fuller toy-check protocol, especially the massive/numerical stage where `Q_Bohm` modifies the Friedmann equation.  Sprint 1 is only the analytic plane-wave control, so it should keep Friedmann as explicit debt.

## Better type design

Do **not** use plain booleans for every check. A boolean forces this bad ambiguity:

```text
false = failed?
false = not implemented?
false = deferred?
```

Use a status enum:

```lean
inductive CheckStatus
| pass
| fail
| deferred
| contested
deriving DecidableEq, Repr
```

Then the Lean contract can say:

```lean
def checkPassed : CheckStatus -> Bool
| CheckStatus.pass => true
| _ => false

def checkDeferred : CheckStatus -> Bool
| CheckStatus.deferred => true
| _ => false
```

## Correct Lean shape

Replace the earlier Lean sketch with this structure:

```lean
inductive CheckStatus
| pass
| fail
| deferred
| contested
deriving DecidableEq, Repr

structure BMCReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  wdwResidual : CheckStatus
  trajectoryFinite : CheckStatus
  clockMonotonic : CheckStatus
  qFinite : CheckStatus
  qApproximatelyZero : CheckStatus
  classicalControl : CheckStatus
  friedmannResidual : CheckStatus
  faithfulness : CheckStatus
deriving Repr

def checkPassed : CheckStatus -> Bool
| CheckStatus.pass => true
| _ => false

def checkDeferred : CheckStatus -> Bool
| CheckStatus.deferred => true
| _ => false

def reportPassesBMC0AControlGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.qFinite &&
  checkPassed r.qApproximatelyZero &&
  checkPassed r.classicalControl &&
  checkDeferred r.friedmannResidual

def reportPassesFullBMCToyGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.qFinite &&
  checkPassed r.classicalControl &&
  checkPassed r.friedmannResidual &&
  checkPassed r.faithfulness
```

Then add safety theorems:

```lean
theorem bmc0a_control_requires_friedmann_deferred
  (r : BMCReport)
  (h : reportPassesBMC0AControlGate r = true) :
  r.friedmannResidual = CheckStatus.deferred := by
  simp [reportPassesBMC0AControlGate, checkPassed, checkDeferred] at h
  exact h.2.2.2.2.2.2.2.2

theorem full_toy_gate_requires_friedmann_pass
  (r : BMCReport)
  (h : reportPassesFullBMCToyGate r = true) :
  r.friedmannResidual = CheckStatus.pass := by
  simp [reportPassesFullBMCToyGate, checkPassed] at h
  exact h.2.2.2.2.2.2.2.1

theorem final_truth_blocks_bmc0a_control
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AControlGate r = false := by
  simp [reportPassesBMC0AControlGate, h]

theorem final_truth_blocks_full_toy_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesFullBMCToyGate r = false := by
  simp [reportPassesFullBMCToyGate, h]
```

The exact `.2.2...` proof paths may need adjustment in Lean, but the logic is right.

## Correct JSON change

Change this:

```json
"friedmann_residual": {
  "pass": false,
  "reason": "Not implemented in BMC-0A plane-wave control; remains debt."
}
```

to this:

```json
"friedmann_residual": {
  "status": "deferred",
  "implemented": false,
  "reason": "Not part of BMC-0A plane-wave control; remains debt for the full BMC toy gate."
}
```

And add two separate recommendations:

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

## What the coding agent should now implement

Tell the coding agent:

```text
Revise the Lean and Go report model to distinguish:
1. BMC-0A plane-wave control gate
2. Full BMC toy promotion gate

Sprint 1 may pass only the BMC-0A control gate.

Sprint 1 must not pass the full BMC toy gate because Friedmann residual is deferred and faithfulness review is not accepted yet.
```

## EBP status after this correction

```text
Claim:
BMC-0A can pass a narrow plane-wave control gate while Friedmann remains deferred.
Status: correct.

Debt:
needToyCheck: partially retired only for plane-wave control.
needInvariant: still partial because Friedmann invariant is not checked.
needFaithfulnessReview: active/contested.
Promotion: blocked for full BMC-0.1.

Next move:
Patch the coding-agent prompt and adversarial-review prompt to require separate gates.
```

This is exactly the kind of thing EBP is supposed to catch: not a physics failure, but a **promotion-boundary bug**.
