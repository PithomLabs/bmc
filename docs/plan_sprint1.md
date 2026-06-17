## Sprint 1

You are implementing the first EBP 2.1-compliant sprint for **Bohmian Minisuperspace Cosmology v0.1 / BMC-0.1**.

Read `README.md` first and treat it as the governing project plan. Do not expand scope beyond the first implementation step.

## Mission

Implement **BMC-0A plane-wave control** in **Go + Lean**.

The goal is not to prove quantum gravity. The goal is to build the smallest executable toy-check artifact:

```text
flat FRW minisuperspace + massless scalar control
variables: α = ln(a), φ
WdW toy equation: (-∂²/∂α² + ∂²/∂φ²) Ψ(α,φ) = 0
plane wave: Ψ(α,φ) = exp(i(kα + ωφ))
constraint: ω² = k²
expected: WdW residual ≈ 0, R constant, Q ≈ 0, finite Bohmian trajectory
```

## Hard Scope Limits

Do **not** implement yet:

```text
massive scalar numerical WdW solver
finite-difference PDE solver for general Ψ
packet/superposition states
node detection beyond placeholder fields
LQC comparison
Page-Wootters comparison
dashboard/UI/database/RAG/LLM integration
black holes
full quantum gravity
fermions
gauge fields
inhomogeneous perturbations
```

This sprint is only the deterministic plane-wave control plus Lean promotion-safety contracts.

## Required Go Structure

Create or update:

```text
cmd/ptw-bmc/
  main.go

internal/bmc/model/
  types.go
  params.go

internal/bmc/wave/
  plane.go
  amplitude.go
  phase.go

internal/bmc/wdw/
  residual.go

internal/bmc/guidance/
  velocity.go
  integrate.go
  trajectory.go

internal/bmc/qpotential/
  q.go

internal/bmc/invariant/
  classical_limit.go

internal/bmc/obstruction/
  obstruction.go
  detect.go

internal/bmc/report/
  report.go
  validate.go
  write_json.go
```

Keep implementation simple, deterministic, and testable.

## Required CLI

Implement:

```bash
go run -buildvcs=false ./cmd/ptw-bmc run --profile bmc0a-plane --out out/bmc0a_plane.json
go run -buildvcs=false ./cmd/ptw-bmc validate --report out/bmc0a_plane.json
go run -buildvcs=false ./cmd/ptw-bmc summarize --report out/bmc0a_plane.json
```

The `run` command should generate a deterministic JSON report.

The `validate` command should check report schema and pass/fail gates.

The `summarize` command should print a short human-readable summary without making physics overclaims.

## Required Physics Computations

For the plane wave:

```text
Ψ(α,φ) = exp(i(kα + ωφ))
R = |Ψ| = 1
S = kα + ωφ
∂S/∂α = k
∂S/∂φ = ω
dα/dλ = ∂S/∂α = k
dφ/dλ = -∂S/∂φ = -ω
Q = -1/(2R)(∂²R/∂α² - ∂²R/∂φ²) = 0
```

WdW residual:

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ = (k² - ω²) Ψ
```

Therefore the residual should be zero or near zero when `ω² = k²`.

Use configurable tolerances, but default to strict deterministic values where possible.

## Required JSON Report

Generate a report shaped like:

```json
{
  "schema_version": "bmc-report-v0.1",
  "model_id": "bmc0a_plane",
  "toy_analysis_only": true,
  "physics_claim": "minisuperspace_only",
  "final_truth_claim": false,
  "promotion_recommendation": "blocked_or_candidate",
  "parameters": {
    "k": 1,
    "omega": 1,
    "alpha0": 0,
    "phi0": 0,
    "lambda_step": 0.1,
    "steps": 100,
    "tolerance": 1e-9
  },
  "equations": {
    "wdw": "(-d_alpha_alpha + d_phi_phi) Psi = 0",
    "wavefunction": "Psi(alpha,phi)=exp(i(k alpha + omega phi))",
    "guidance": "dalpha/dlambda=k, dphi/dlambda=-omega",
    "quantum_potential": "Q=0 for constant amplitude"
  },
  "checks": {
    "wdw_residual": {
      "pass": true,
      "max_abs_residual": 0
    },
    "trajectory": {
      "pass": true,
      "finite": true,
      "points": 101
    },
    "clock_monotonicity": {
      "pass": true,
      "variable": "phi"
    },
    "quantum_potential": {
      "pass": true,
      "max_abs_q": 0
    },
    "classical_limit": {
      "pass": true,
      "reason": "Plane-wave control has Q approximately zero."
    },
    "friedmann_residual": {
      "pass": false,
      "reason": "Not implemented in BMC-0A plane-wave control; remains debt."
    }
  },
  "null_models": [
    {
      "name": "classical_frw",
      "status": "placeholder_obligation",
      "reason": "Classical-limit comparison begins with Q≈0 control."
    },
    {
      "name": "standard_wdw",
      "status": "placeholder_obligation",
      "reason": "Ensemble comparison deferred."
    },
    {
      "name": "lqc",
      "status": "deferred",
      "reason": "Not part of BMC-0A."
    },
    {
      "name": "page_wootters",
      "status": "deferred",
      "reason": "Formal relational-time null model deferred."
    }
  ],
  "obstructions": [],
  "faithfulness": {
    "status": "contested",
    "reason": "No human faithfulness review yet. This only tests a plane-wave minisuperspace control."
  },
  "ebp_debt": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "active"
  },
  "warnings": [
    "Toy analysis only.",
    "Does not test full quantum gravity.",
    "Does not test black holes, fermions, gauge fields, Lorentz recovery, or inhomogeneous perturbations.",
    "Passing this report cannot promote any final-truth claim."
  ]
}
```

Adjust numeric values if the implementation uses different defaults, but preserve the schema spirit and overclaim blockers.

## Required Go Tests

Add tests for:

```text
TestPlaneWaveSatisfiesWdWResidual
TestPlaneWaveQApproximatelyZero
TestPlaneWaveTrajectoryFinite
TestClockMonotonicityDetection
TestReportDeterministicJSON
TestValidateRejectsFinalTruthClaim
TestValidateKeepsToyOnlyStatus
```

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
```

## Required Lean Structure

Add minimal Lean files if a Lean/Lake project already exists. If none exists, create the smallest reasonable Lake project without disrupting the Go module.

Recommended files:

```text
BMC/
  ToyReport.lean
  Promotion.lean
```

Lean should not prove physics or numerics. Lean should encode promotion-safety contracts only.

Minimum Lean structures:

```lean
structure BMCReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  wdwResidualPass : Bool
  trajectoryFinitePass : Bool
  clockMonotonicPass : Bool
  qFinitePass : Bool
  classicalLimitPass : Bool
  friedmannResidualPass : Bool
  faithfulnessAccepted : Bool
```

Minimum promotion gate:

```lean
def reportPassesToyGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.wdwResidualPass &&
  r.trajectoryFinitePass &&
  r.clockMonotonicPass &&
  r.qFinitePass &&
  r.classicalLimitPass &&
  r.friedmannResidualPass &&
  r.faithfulnessAccepted
```

Required theorem obligations:

```lean
theorem final_truth_blocks_toy_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesToyGate r = false := by
  simp [reportPassesToyGate, h]

theorem toy_gate_requires_toy_only
  (r : BMCReport)
  (h : reportPassesToyGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesToyGate] at h
  exact h.1

theorem faithfulness_required
  (r : BMCReport)
  (h : reportPassesToyGate r = true) :
  r.faithfulnessAccepted = true := by
  simp [reportPassesToyGate] at h
  exact h.2.2.2.2.2.2.2.2
```

If the exact Lean proof shape differs, repair it cleanly. Do not use `sorry`.

Run:

```bash
lake build
```

If Lean is not installed or the repo is not configured for Lean, still add the Lean files and document the unrun build honestly in the final report.

## EBP 2.1 Requirements

Every generated report must preserve these claims:

```text
BMC-0A is toy-only.
BMC-0A is not full quantum gravity.
BMC-0A is not a solution to the problem of time.
BMC-0A is not a proof of Bohmian mechanics.
BMC-0A only tests whether the plane-wave Wheeler-DeWitt/Bohmian guidance pipeline is internally coherent.
```

Do not write comments, docs, CLI output, or report text that says or implies:

```text
solves quantum gravity
proves Bohmian quantum gravity
derives spacetime
solves the problem of time
settles quantum cosmology
```

## Final Response Format

When finished, report back in this exact JSON shape:

```json
{
  "summary": "",
  "implemented_files": [],
  "commands_run": [],
  "test_results": [],
  "assumptions": [],
  "proof_obligations": [],
  "null_models": [],
  "risks": [],
  "human_review_questions": [],
  "promotion_status": "alive_unpromoted_toy_check_in_progress"
}
```

If something fails, do not hide it. Report the failure and the next smallest repair.

One correction I made from the README plan: for Sprint 1, `friedmann_residual` should remain **not implemented / debt**, because the plane-wave control is mainly validating the WdW residual, phase, guidance, trajectory, and `Q≈0` pipeline. The coding agent should not fake a Friedmann check before the model earns it.
