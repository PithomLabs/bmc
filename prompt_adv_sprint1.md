## prompt

also before we start, I want to ensure that we employ another LLM for adversarial review

Yes — that should be part of the workflow before any implementation is accepted.

Under EBP 2.1, adversarial review is not decoration; it is how we file obstruction debt, null-model debt, and faithfulness debt. EBP explicitly welcomes objections, counterexamples, null models, faithfulness failures, formalization failures, no-go results, and better rival explanations because each creates clearer debt.  It also says Lean does not decide physical faithfulness and does not replace human judgment, so the other LLM should attack both the **physics** and the **formalization boundary**, not merely the code. 

Here is the extra prompt to run **after the coding agent produces Sprint 1 output**, but **before we accept or build on it**.

You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review the implementation of **Bohmian Minisuperspace Cosmology v0.1 / BMC-0A plane-wave control**.

The project goal is deliberately narrow:

```text
flat FRW minisuperspace + massless scalar control
variables: α = ln(a), φ
toy Wheeler-DeWitt equation: (-∂²/∂α² + ∂²/∂φ²) Ψ(α,φ) = 0
plane wave: Ψ(α,φ) = exp(i(kα + ωφ))
constraint: ω² = k²
expected: WdW residual ≈ 0, R constant, Q ≈ 0, finite Bohmian trajectory
```

This artifact must **not** claim to solve quantum gravity, prove Bohmian mechanics, derive spacetime, solve the problem of time, recover Lorentz symmetry, explain black holes, handle fermions, handle gauge fields, or validate full Wheeler-DeWitt quantum gravity.

Your job is to attack the artifact.

## Materials to Review

Review:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
BMC/**/*.lean
out/bmc0a_plane.json
all Go tests
all Lean files
coding agent final report
```

If any of these files are missing, report that as implementation debt.

## Review Goals

Evaluate whether the implementation actually satisfies the first sprint:

```text
BMC-0A plane-wave control in Go
deterministic JSON report
Go tests for residual, Q≈0, finite trajectory, monotonic clock, report determinism, final-truth blocker
Lean promotion-safety contracts
no overclaiming
```

## EBP 2.1 Review Ledger

For each debt item, classify as one of:

```text
unpaid
partial
retired
contested
overclaimed
```

Debt items:

```text
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
containsFinalTruthClaim
```

Be strict: writing equations is not the same as passing a toy check. Passing a toy check is not proof of nature.

## Physics Review Questions

Attack these points:

1. Does the implementation use the correct toy WdW residual?

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ = (k² - ω²) Ψ
```

for:

```text
Ψ = exp(i(kα + ωφ))
```

2. Does the implementation correctly enforce or report the constraint?

```text
ω² = k²
```

3. Does it compute:

```text
R = 1
S = kα + ωφ
Q = -1/(2R)(∂²R/∂α² - ∂²R/∂φ²) = 0
```

without fake numerical complexity?

4. Are the guidance equations consistent with the selected sign convention?

```text
dα/dλ = ∂S/∂α = k
dφ/dλ = -∂S/∂φ = -ω
```

5. Is `φ` monotonic for the selected parameters? If `ω = 0`, does the clock test fail correctly?

6. Does the model avoid pretending that the plane-wave control has already checked the Friedmann residual?

7. Does the report clearly state that the Friedmann check remains debt in BMC-0A?

8. Does any report, CLI output, comment, README section, or test name imply that BMC-0A validates quantum gravity?

## Code Review Questions

Attack these points:

1. Is the JSON deterministic byte-for-byte across repeated runs?

2. Are tolerances explicit and recorded?

3. Are invalid inputs rejected?

Examples:

```text
k or omega NaN
k or omega Inf
steps <= 0
lambda_step <= 0
constraint violation |k² - omega²| above tolerance
output path missing or unwritable
unknown profile
```

4. Does validation reject:

```text
final_truth_claim=true
toy_analysis_only=false
missing warnings
missing EBP debt fields
NaN/Inf numeric values
```

5. Are all package boundaries reasonable and simple?

6. Is there any unnecessary architecture, dashboard, database, RAG, LLM integration, or premature generality?

7. Are tests meaningful or just asserting implementation details?

8. Are error messages useful?

## Lean Review Questions

Attack these points:

1. Do Lean files build without `sorry`?

2. Do the theorem names match the actual safety obligations?

3. Does Lean only prove promotion-safety claims rather than pretending to prove physics?

4. Does `reportPassesToyGate` block promotion when:

```text
finalTruthClaim = true
toyAnalysisOnly = false
faithfulnessAccepted = false
friedmannResidualPass = false
```

5. Does the Lean layer clearly preserve the boundary:

```text
passing toy report != full quantum gravity
```

6. Are there any weak definitions that make the theorems trivial in a misleading way?

Trivial promotion-safety theorems are acceptable only if the report honestly says they are policy guards, not physics proofs.

## Required Output

Return your review in this exact JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
    "needMap": "",
    "needInvariant": "",
    "needToyCheck": "",
    "needNullModel": "",
    "needObstruction": "",
    "needFaithfulnessReview": "",
    "containsFinalTruthClaim": ""
  },
  "physics_findings": [],
  "code_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "null_model_notes": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|toy_candidate_only|promoted_toy_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict Review Rule

If the artifact passes the plane-wave control, your maximum allowed promotion recommendation is:

```text
toy_candidate_only
```

or, after all required repairs:

```text
promoted_toy_artifact_after_repairs
```

You may not recommend promotion as full quantum gravity, a solution to the problem of time, or a proof of Bohmian mechanics.

My recommended workflow is:

```text
1. Coding agent implements Sprint 1.
2. Coding agent returns JSON report.
3. We send the implementation + report to another LLM using the adversarial prompt above.
4. You paste the adversarial review back here.
5. I synthesize: accept, repair, or reject under EBP 2.1.
6. Only then do we proceed to packet/superposition BMC-0A or BMC-0B.
```

For Sprint 1, I would make the adversarial review a **hard gate**. No second sprint until the reviewer signs off or we repair the blocker findings.
