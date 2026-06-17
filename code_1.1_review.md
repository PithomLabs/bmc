You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review the completed **BMC Sprint 1: BMC-0A Plane-Wave Control** implementation.

The artifact claims only this narrow result:

```text
In flat FRW minisuperspace with variables α = ln(a), φ, the plane-wave Wheeler-DeWitt control Ψ(α,φ)=exp(i(kα+ωφ)) with ω²=k² passes a narrow control gate: WdW residual ≈ 0, Q ≈ 0, finite Bohmian trajectory, monotonic φ clock, deterministic JSON, and full promotion blocked.
```

It must not claim:

```text
full quantum gravity
Bohmian quantum gravity proven
spacetime derived
problem of time solved
black holes handled
fermions handled
gauge fields handled
Lorentz recovery
inhomogeneous perturbations
```

## Materials to review

Review the actual source code and generated artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
BMC/**
out/bmc0a_plane.json
go.mod
all Go tests
coding-agent final walkthrough/report
```

## Important context

Sprint 1 intentionally uses split gates:

```text
bmc0a_plane_control_gate — passable in Sprint 1
full_bmc_toy_gate — must remain blocked
```

Friedmann residual must be:

```text
status = deferred
implemented = false
```

Faithfulness must remain contested or deferred unless a real human/reviewer faithfulness review has occurred.

Lean/Lake may not have been run in the coding environment. If `lake build` was not run successfully, mark Lean verification as unpaid or contested.

## Review tasks

Check the following strictly.

### Physics checks

1. Is the WdW residual correct?

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ = (k² - ω²)Ψ
```

2. Is the plane-wave constraint enforced?

```text
ω² = k²
```

3. Is Q correctly zero for constant amplitude?

```text
R = 1
Q = -1/(2R)(∂²R/∂α² - ∂²R/∂φ²) = 0
```

4. Are the guidance equations implemented consistently?

```text
dα/dλ = k
dφ/dλ = -ω
```

5. Does φ monotonicity fail correctly when ω = 0?

6. Does the report avoid pretending that Friedmann residual was checked?

### Code checks

1. Is JSON deterministic byte-for-byte?
2. Are NaN/Inf parameters rejected?
3. Are invalid steps rejected?
4. Is constraint violation rejected?
5. Does validation reject final_truth_claim=true?
6. Does validation reject toy_analysis_only=false?
7. Does validation enforce status/pass consistency?
8. Does the promotion gate remain blocked when Friedmann or faithfulness is not passed?
9. Is the CLI zero-dependency and simple?
10. Are there any hidden overclaims in comments, summary output, report text, or test names?

### Lean checks

1. Does the Lean project build with `lake build`?
2. Are there any `sorry`s?
3. Do the split gates exist?

```text
reportPassesBMC0AControlGate
reportPassesFullBMCToyGate
```

4. Does the control gate require Friedmann deferred?
5. Does the full toy gate require Friedmann pass and faithfulness pass?
6. Does the Sprint 1 witness pass the control gate but fail the full gate?
7. Are the Lean theorems only promotion-safety contracts, not physics proofs?

## EBP debt classification

Classify each as:

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
LeanVerification
```

## Required output JSON

Return exactly:

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
    "containsFinalTruthClaim": "",
    "LeanVerification": ""
  },
  "physics_findings": [],
  "code_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|control_gate_candidate_only|promoted_control_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict rule

Even if Sprint 1 passes perfectly, the maximum allowed recommendation is:

```text
promoted_control_artifact_after_repairs
```

Never recommend promotion as full BMC, full quantum gravity, a proof of Bohmian mechanics, or a solution to the problem of time.
