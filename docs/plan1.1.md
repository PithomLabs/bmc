# BMC Sprint 1: BMC-0A Plane-Wave Control

Implement the first EBP 2.1-compliant toy check for Bohmian Minisuperspace Cosmology: the plane-wave control case in Go, plus Lean promotion-safety contracts.

## User Review Required

> [!IMPORTANT]
> **Split-gate architecture**: Per your clarification, we implement two separate gates:
> - `bmc0a_plane_control_gate` — passable in Sprint 1
> - `full_bmc_toy_gate` — blocked until Friedmann residual + faithfulness review
>
> This replaces the single `reportPassesToyGate` from the sprint plan.

> [!WARNING]
> **Friedmann residual** is explicitly deferred. The JSON report will show `"status": "deferred"` for this check. The Lean promotion gate will require `friedmannResidual = pass`, keeping full promotion blocked.

---

## Proposed Changes

### Go Module Initialization

Initialize the Go module at project root.

#### [NEW] [go.mod](file:///home/chaschel/Documents/go/bmc/go.mod)
- Module path: `github.com/PithomLabs/bmc`
- Go version: 1.22+ (latest stable)
- Zero external dependencies

---

### Model Types (`internal/bmc/model/`)

Core data structures shared across all packages.

#### [NEW] [types.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/types.go)
- `MiniState` struct: `Alpha float64`, `Phi float64`
- `TrajectoryPoint` struct: `Lambda float64`, `State MiniState`
- `Trajectory` struct: `Points []TrajectoryPoint`
- `Complex` type alias for `complex128`
- `CheckStatus` type (string enum): `"pass"`, `"fail"`, `"deferred"`, `"contested"`
- `CheckResult` struct: `Status CheckStatus`, `Pass bool`, `Reason string`, plus optional numeric fields
- `ObstructionSeverity` type: `"info"`, `"warning"`, `"blocker"`
- `Obstruction` struct: `Name string`, `Applies bool`, `Severity ObstructionSeverity`, `Evidence string`, `Consequence string`, `Status CheckStatus`

#### [NEW] [params.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/params.go)
- `PlaneWaveParams` struct — all numeric model parameters are `float64` except `Steps`:
  ```go
  type PlaneWaveParams struct {
      K          float64
      Omega      float64
      Alpha0     float64
      Phi0       float64
      LambdaStep float64
      Steps      int
      Tolerance  float64
  }
  ```
- `DefaultPlaneWaveParams()` function returning: `K=1, Omega=1, Alpha0=0, Phi0=0, LambdaStep=0.1, Steps=100, Tolerance=1e-9`
- `Validate(params PlaneWaveParams) error` — rejects:
  - Constraint violation: `|ω² - k²| > tolerance`
  - Non-finite parameters (NaN/Inf in any field)
  - Invalid step count (`Steps <= 0`)

---

### Wave Functions (`internal/bmc/wave/`)

Plane-wave wavefunction, amplitude, and phase extraction.

#### [NEW] [plane.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/plane.go)
- `PlaneWave` struct holding `K`, `Omega float64`
- `Psi(alpha, phi float64) complex128` — returns `exp(i(kα + ωφ))`
- `PsiGrid(alphaRange, phiRange []float64) [][]complex128` — evaluates on grid (for future use)
- Uses `math/cmplx` for complex arithmetic

#### [NEW] [amplitude.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/amplitude.go)
- `Amplitude(psi complex128) float64` — returns `cmplx.Abs(psi)` = `|Ψ|`
- `AmplitudeField(alpha, phi float64, wf WaveFunction) float64` — evaluates R at a point
- For plane wave, R = 1 everywhere (constant)

#### [NEW] [phase.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/phase.go)
- `Phase(psi complex128) float64` — returns `cmplx.Phase(psi)` = `arg(Ψ)`
- `PhaseGradient(alpha, phi float64, wf WaveFunction) (dSdAlpha, dSdPhi float64)` — returns `(k, ω)` for plane wave
- `WaveFunction` interface: `Psi(alpha, phi float64) complex128`
  - `PlaneWave` implements this interface

---

### WdW Residual (`internal/bmc/wdw/`)

Wheeler-DeWitt residual computation.

#### [NEW] [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go)

Two residual implementations — analytic is the **primary authority** for Sprint 1:

- `AnalyticResidualPlaneWave(k, omega float64) float64`
  - Returns `k² - ω²` (exact, no floating-point discretization error)
  - **This is the primary pass/fail gate for Sprint 1**
- `FiniteDifferenceResidual(wf wave.WaveFunction, alpha, phi, h float64) complex128`
  - Computes `(-∂²Ψ/∂α² + ∂²Ψ/∂φ²)` using central finite differences with step `h`
  - Included as a **numerical sanity check only**, not the main authority
- `MaxAbsFiniteDiffResidual(wf wave.WaveFunction, points []model.MiniState, h float64) float64`
  - Evaluates finite-difference residual at all trajectory points, returns max |residual|
- `CheckResidual(analyticRes float64, tolerance float64) model.CheckResult`
  - Uses the analytic residual as the authority

---

### Bohmian Guidance (`internal/bmc/guidance/`)

Velocity field, integration, and trajectory construction.

#### [NEW] [velocity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/velocity.go)
- `Velocity(wf wave.WaveFunction, alpha, phi float64) (dAlpha, dPhi float64)`
  - `dα/dλ = ∂S/∂α = k`
  - `dφ/dλ = -∂S/∂φ = -ω`
  - Uses numerical phase gradient from `wave.PhaseGradient`

#### [NEW] [integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go)
- `Stepper` interface: `Step(state model.MiniState, dt float64, vel func(model.MiniState) (float64, float64)) model.MiniState`
- `EulerStepper` struct implementing `Stepper`
  - For plane wave with constant velocity, Euler is exact
- Deferred debt: RK4Stepper for packet/superposition states

#### [NEW] [trajectory.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/trajectory.go)
- `Integrate(wf wave.WaveFunction, initial model.MiniState, stepper Stepper, dt float64, steps int) model.Trajectory`
  - Produces `steps+1` points (including initial condition)
- `IsFinite(traj model.Trajectory) bool` — checks all points for NaN/Inf
- `CheckTrajectory(traj model.Trajectory) model.CheckResult`

---

### Quantum Potential (`internal/bmc/qpotential/`)

#### [NEW] [q.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/q.go)
- `Q(wf wave.WaveFunction, alpha, phi, h float64) float64`
  - `Q = -1/(2R) * (∂²R/∂α² - ∂²R/∂φ²)`
  - Uses central finite differences for second derivatives of R
  - For plane wave: R = 1, all derivatives of R = 0, so Q = 0
- `QAlongTrajectory(wf wave.WaveFunction, traj model.Trajectory, h float64) []float64`
- `MaxAbsQ(qValues []float64) float64`
- `CheckQuantumPotential(maxQ float64, tolerance float64) model.CheckResult`

---

### Classical Limit Invariant (`internal/bmc/invariant/`)

#### [NEW] [classical_limit.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/invariant/classical_limit.go)
- `CheckClassicalLimit(maxAbsQ float64, tolerance float64) model.CheckResult`
  - For plane wave: Q ≈ 0 means classical-like behavior
  - Reports pass with reason: "Plane-wave control has Q approximately zero."

---

### Obstruction Detection (`internal/bmc/obstruction/`)

4 implemented + 5 deferred placeholders.

#### [NEW] [obstruction.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/obstruction.go)
- Enum-like constants for all 9 obstruction names
- `NewDeferred(name, reason string) model.Obstruction` — creates deferred placeholder

#### [NEW] [detect.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/detect.go)
- **Implemented detection** (Sprint 1):
  - `DetectWdWResidual(maxRes, tolerance float64) model.Obstruction`
  - `DetectNonfiniteQ(qValues []float64) model.Obstruction`
  - `DetectClockNonmonotonicity(traj model.Trajectory, variable string) model.Obstruction`
  - `DetectOverclaimBlocker(finalTruthClaim bool) model.Obstruction`
    - Always **present** in the report
    - `applies: true` + severity `blocker` only when `finalTruthClaim = true` or forbidden language detected
    - `applies: false` when no overclaim detected (clean Sprint 1 report):
      ```json
      { "name": "full_qg_overclaim_blocker", "applies": false,
        "severity": "blocker", "status": "pass",
        "evidence": "No final-truth language detected.",
        "consequence": "Continue as toy-only artifact." }
      ```
- **Deferred placeholders** (filed as `status: deferred`):
  - `node_obstruction`
  - `phase_unwrap_obstruction`
  - `classical_limit_failure`
  - `lapse_or_time_interpretation_debt`
  - `measure_problem_deferred`
- `DetectAll(...)` — runs all detectors, returns slice of obstructions

---

### Report Generation (`internal/bmc/report/`)

JSON report, validation, and human-readable summary.

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go)
- `Report` struct matching the JSON schema from the sprint plan, with these modifications:
  - Each check uses `CheckResult` with `Status` field (not just boolean `Pass`)
  - **Consistency rule**: `Status` is the single source of truth. If `Pass bool` is included in JSON for readability, validation enforces: `status = "pass" ↔ pass = true`, `status ∈ {"fail","deferred","contested"} ↔ pass = false`
  - `TechnicalGate` struct: `Name string`, `Status CheckStatus`
  - `PromotionGate` struct: `Name string`, `Status CheckStatus`, `Reason string`
  - Friedmann residual: `Status: "deferred"`, `Implemented: false`
- `Generate(params model.PlaneWaveParams) (*Report, error)` — runs entire pipeline:
  1. Construct plane wave from params
  2. Compute WdW residual
  3. Integrate trajectory
  4. Compute Q along trajectory
  5. Check clock monotonicity
  6. Check classical limit
  7. Detect obstructions
  8. Assemble report with both gates
- `SchemaVersion` = `"bmc-report-v0.1"`

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go)
- `Validate(report *Report) []ValidationError`
  - Schema version present
  - `toy_analysis_only = true`
  - `final_truth_claim = false`
  - All required fields present
  - **Status/Pass consistency**: for every check, `status = "pass" ↔ pass = true`
  - Technical gate status is consistent with individual checks
  - Promotion gate is blocked if Friedmann or faithfulness are not passed
  - Warnings array is non-empty
- `ValidationError` struct with severity levels

#### [NEW] [write_json.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/write_json.go)
- `WriteJSON(report *Report, path string) error` — deterministic JSON output
  - Uses `json.MarshalIndent` with sorted keys where possible
  - Creates parent directories if needed

---

### CLI (`cmd/ptw-bmc/`)

#### [NEW] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Manual subcommand dispatch using `os.Args` + `flag.NewFlagSet`
- Three subcommands:

**`run`**:
- `--profile bmc0a-plane` (only supported profile for Sprint 1)
- `--out path/to/output.json`
- Calls `report.Generate()` + `report.WriteJSON()`

**`validate`**:
- `--report path/to/report.json`
- Reads JSON, calls `report.Validate()`
- Prints pass/fail with details
- Exit code 0 if validation passes, 1 if fails

**`summarize`**:
- `--report path/to/report.json`
- Prints human-readable summary
- Includes EBP non-claims warnings
- Does NOT make physics overclaims

---

### Lean Project (`BMC/`)

Minimal Lean 4 Lake project, no Mathlib dependency.

#### [NEW] [lakefile.lean](file:///home/chaschel/Documents/go/bmc/BMC/lakefile.lean)
- Lake project named `BMC`
- No external dependencies

#### [NEW] [lean-toolchain](file:///home/chaschel/Documents/go/bmc/BMC/lean-toolchain)
- Pinned to current stable Lean 4 version

#### [NEW] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
- Root file importing `BMC.ToyReport` and `BMC.Promotion`

#### [NEW] [BMC/ToyReport.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ToyReport.lean)
- `CheckStatus` inductive: `pass | fail | deferred | contested` with `DecidableEq, Repr`
- Helper predicates: `checkPassed`, `checkDeferred`
- `BMCReport` structure with `CheckStatus` fields instead of `Bool`:
  ```lean
  structure BMCReport where
    toyAnalysisOnly     : Bool
    finalTruthClaim     : Bool
    wdwResidual         : CheckStatus
    trajectoryFinite    : CheckStatus
    clockMonotonic      : CheckStatus
    qFinite             : CheckStatus
    classicalLimit      : CheckStatus
    friedmannResidual   : CheckStatus
    faithfulness        : CheckStatus
  ```

#### [NEW] [BMC/Promotion.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Promotion.lean)
- **Control gate** (Sprint 1):
  ```lean
  def reportPassesBMC0AControlGate (r : BMCReport) : Bool :=
    r.toyAnalysisOnly &&
    !r.finalTruthClaim &&
    checkPassed r.wdwResidual &&
    checkPassed r.trajectoryFinite &&
    checkPassed r.clockMonotonic &&
    checkPassed r.qFinite &&
    checkPassed r.classicalLimit &&
    checkDeferred r.friedmannResidual
  ```
- **Full toy gate** (future sprints):
  ```lean
  def reportPassesFullBMCToyGate (r : BMCReport) : Bool :=
    r.toyAnalysisOnly &&
    !r.finalTruthClaim &&
    checkPassed r.wdwResidual &&
    checkPassed r.trajectoryFinite &&
    checkPassed r.clockMonotonic &&
    checkPassed r.qFinite &&
    checkPassed r.classicalLimit &&
    checkPassed r.friedmannResidual &&
    checkPassed r.faithfulness
  ```
- **Theorem obligations**:
  - `final_truth_blocks_control_gate`: `finalTruthClaim = true → control gate = false`
  - `final_truth_blocks_toy_gate`: `finalTruthClaim = true → toy gate = false`
  - `control_gate_requires_toy_only`: `control gate = true → toyAnalysisOnly = true`
  - `toy_gate_requires_toy_only`: `toy gate = true → toyAnalysisOnly = true`
  - `faithfulness_required_for_full_gate`: `toy gate = true → faithfulness = pass`
  - `friedmann_deferred_in_control_gate`: `control gate = true → friedmannResidual = deferred`
- **Sprint 1 witness** (concrete example, not a universal theorem):
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
    reportPassesBMC0AControlGate sprint1Witness = true := by decide

  theorem sprint1_witness_fails_full_gate :
    reportPassesFullBMCToyGate sprint1Witness = false := by decide
  ```

All proofs by `simp` / `decide` — no `sorry`.

---

### Go Tests

#### [NEW] [internal/bmc/wave/plane_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/plane_test.go)
- `TestPlaneWaveSatisfiesWdWResidual`: construct plane wave with k=1,ω=1, verify **analytic** residual = 0

#### [NEW] [internal/bmc/qpotential/q_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/q_test.go)
- `TestPlaneWaveQApproximatelyZero`: verify max|Q| < 1e-9 along trajectory

#### [NEW] [internal/bmc/guidance/trajectory_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/trajectory_test.go)
- `TestPlaneWaveTrajectoryFinite`: 100-step trajectory, all points finite
- `TestClockMonotonicityDetection`: verify φ is monotonic for ω ≠ 0
- `TestClockMonotonicityFailsWhenOmegaZero`: verify φ is flagged non-monotonic when ω = 0

#### [NEW] [internal/bmc/report/report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go)
- `TestReportDeterministicJSON`: generate twice, compare byte-for-byte
- `TestValidateRejectsFinalTruthClaim`: set `finalTruthClaim=true`, validate should fail
- `TestValidateKeepsToyOnlyStatus`: normal report passes validation, remains toy-only

#### [NEW] [internal/bmc/model/params_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/params_test.go)
- `TestValidateRejectsConstraintViolation`: k=1, ω=2 → `|ω²-k²| > tolerance` → error
- `TestValidateRejectsInvalidStepCount`: Steps=0 or Steps<0 → error
- `TestValidateRejectsNonfiniteParameters`: NaN or Inf in any numeric field → error

---

### Output Artifacts

#### [NEW] out/bmc0a_plane.json
- Generated by `go run ./cmd/ptw-bmc run --profile bmc0a-plane --out out/bmc0a_plane.json`
- Not checked in; generated during verification

---

## JSON Report Shape (Reference)

The generated report will follow this structure:

```json
{
  "schema_version": "bmc-report-v0.1",
  "model_id": "bmc0a_plane",
  "toy_analysis_only": true,
  "physics_claim": "minisuperspace_only",
  "final_truth_claim": false,
  "parameters": {
    "k": 1, "omega": 1, "alpha0": 0, "phi0": 0,
    "lambda_step": 0.1, "steps": 100, "tolerance": 1e-9
  },
  "equations": {
    "wdw": "(-d²/dα² + d²/dφ²) Ψ = 0",
    "wavefunction": "Ψ(α,φ) = exp(i(kα + ωφ))",
    "guidance": "dα/dλ = k, dφ/dλ = -ω",
    "quantum_potential": "Q = -1/(2R)(d²R/dα² - d²R/dφ²)"
  },
  "checks": {
    "wdw_residual": { "status": "pass", "pass": true, "max_abs_residual": 0 },
    "trajectory": { "status": "pass", "pass": true, "finite": true, "points": 101 },
    "clock_monotonicity": { "status": "pass", "pass": true, "variable": "phi" },
    "quantum_potential": { "status": "pass", "pass": true, "max_abs_q": 0 },
    "classical_limit": { "status": "pass", "pass": true, "reason": "Plane-wave control has Q approximately zero." },
    "friedmann_residual": {
      "status": "deferred",
      "implemented": false,
      "reason": "Not part of BMC-0A plane-wave control; remains debt for the full BMC toy gate."
    }
  },
  "technical_gate": {
    "name": "bmc0a_plane_control_gate",
    "status": "pass"
  },
  "promotion_gate": {
    "name": "full_bmc_toy_gate",
    "status": "blocked",
    "reason": "Friedmann residual and faithfulness review remain unpaid debt."
  },
  "null_models": [
    { "name": "classical_frw", "status": "placeholder_obligation", "reason": "..." },
    { "name": "standard_wdw", "status": "placeholder_obligation", "reason": "..." },
    { "name": "lqc", "status": "deferred", "reason": "Not part of BMC-0A." },
    { "name": "page_wootters", "status": "deferred", "reason": "..." }
  ],
  "obstructions": [
    { "name": "full_qg_overclaim_blocker", "applies": true, "severity": "blocker", "status": "pass", "..." : "..." }
  ],
  "faithfulness": { "status": "contested", "reason": "No human faithfulness review yet. ..." },
  "ebp_debt": {
    "needMap": "partial", "needInvariant": "partial",
    "needToyCheck": "active", "needNullModel": "partial",
    "needObstruction": "active", "needFaithfulnessReview": "active"
  },
  "warnings": [
    "Toy analysis only.",
    "Does not test full quantum gravity.",
    "Does not test black holes, fermions, gauge fields, Lorentz recovery, or inhomogeneous perturbations.",
    "Passing this report cannot promote any final-truth claim."
  ]
}
```

---

## File Summary

| Component | Files | Purpose |
|-----------|-------|---------|
| Go module | `go.mod` | Module `github.com/PithomLabs/bmc` |
| Model types | `internal/bmc/model/{types,params}.go` | Shared data structures |
| Wave | `internal/bmc/wave/{plane,amplitude,phase}.go` | Plane-wave Ψ, R, S |
| WdW | `internal/bmc/wdw/residual.go` | Wheeler-DeWitt residual |
| Guidance | `internal/bmc/guidance/{velocity,integrate,trajectory}.go` | Bohmian guidance + Euler stepper |
| Q potential | `internal/bmc/qpotential/q.go` | Quantum potential Q |
| Invariant | `internal/bmc/invariant/classical_limit.go` | Classical-limit check |
| Obstruction | `internal/bmc/obstruction/{obstruction,detect}.go` | 4 active + 5 deferred |
| Report | `internal/bmc/report/{report,validate,write_json}.go` | JSON report pipeline |
| CLI | `cmd/ptw-bmc/main.go` | run / validate / summarize |
| Lean | `BMC/{lakefile,BMC,BMC/ToyReport,BMC/Promotion}.lean` | Promotion-safety contracts |
| Tests | 4 test files across packages | 7 test functions |

**Total: ~20 new files, 0 modified files, 0 external dependencies**

---

## Verification Plan

### Automated Tests

```bash
# Go tests
go test ./...

# Go build
go build -buildvcs=false ./cmd/ptw-bmc

# Generate report
go run -buildvcs=false ./cmd/ptw-bmc run --profile bmc0a-plane --out out/bmc0a_plane.json

# Validate report
go run -buildvcs=false ./cmd/ptw-bmc validate --report out/bmc0a_plane.json

# Summarize report
go run -buildvcs=false ./cmd/ptw-bmc summarize --report out/bmc0a_plane.json

# Lean build (if toolchain available)
cd BMC && lake build
```

### Expected Outcomes
- All 7 Go tests pass
- Report JSON is deterministic
- Technical gate: `pass`
- Promotion gate: `blocked`
- WdW residual: 0 (or < 1e-9)
- Q: 0 (or < 1e-9)
- Trajectory: 101 finite points
- Lean: `lake build` succeeds with no `sorry`
