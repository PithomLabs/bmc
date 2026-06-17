# BMC Sprint 2: BMC-0A Two-Plane-Wave Superposition Control

Stress-test the existing BMC-0A pipeline when $\Psi$ has nonconstant amplitude, possible nodes, nonzero quantum potential $Q$, and nonconstant Bohmian velocity.

## User Review Required

> [!IMPORTANT]
> **Proposed Default Superposition Parameters**:
> - Wave 1: $c_1 = 1.0 + 0i$, $k_1 = 1.0$, $\omega_1 = 1.0$
> - Wave 2: $c_2 = 0.5 + 0i$, $k_2 = 2.0$, $\omega_2 = -2.0$ (opposite sign to create interference)
> - Initial state: $\alpha_0 = 0.0, \phi_0 = 0.0$
> - Integration: $\Delta\lambda = 0.05, N = 200$ steps
> - Node threshold: $10^{-5}$
>
> Please confirm if these default values are acceptable or if different default values are preferred.

> [!WARNING]
> **Split-gate preservation**:
> We will introduce a new technical gate `bmc0a_superposition_control_gate` while keeping `full_bmc_toy_gate` blocked. Under EBP 2.1, passing this superposition control gate requires that the simulation runs without triggering active obstructions (e.g. no node contact, clock remains monotonic, etc.).

---

## Proposed Changes

### Model Types and Configuration (`internal/bmc/model/`)

#### [MODIFY] [types.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/types.go)
- No structural changes required; reuse existing structures.

#### [MODIFY] [params.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/params.go)
- Add `SuperpositionParams` struct:
  ```go
  type SuperpositionParams struct {
      C1Real     float64 `json:"c1_real"`
      C1Imag     float64 `json:"c1_imag"`
      K1         float64 `json:"k1"`
      Omega1     float64 `json:"omega1"`
      C2Real     float64 `json:"c2_real"`
      C2Imag     float64 `json:"c2_imag"`
      K2         float64 `json:"k2"`
      Omega2     float64 `json:"omega2"`
      Alpha0     float64 `json:"alpha0"`
      Phi0       float64 `json:"phi0"`
      LambdaStep float64 `json:"lambda_step"`
      Steps      int     `json:"steps"`
      Tolerance  float64 `json:"tolerance"`
      NodeThresh float64 `json:"node_threshold"`
  }
  ```
- Add `DefaultSuperpositionParams()` returning the default values.
- Add `Validate()` method on `SuperpositionParams` checking:
  - Component constraints: $\omega_1^2 = k_1^2$ and $\omega_2^2 = k_2^2$ hold within tolerance.
  - Non-finite numbers in all fields.
  - `Steps > 0`, `LambdaStep > 0`, and `NodeThresh > 0`.

---

### Wave Functions (`internal/bmc/wave/`)

#### [NEW] [superposition.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/superposition.go)
- `SuperpositionWave` struct holding $c_1, c_2$ (as `complex128`), and $k_1, \omega_1, k_2, \omega_2$ (as `float64`).
- Implements `WaveFunction` interface:
  ```go
  func (s SuperpositionWave) Psi(alpha, phi float64) complex128 {
      psi1 := s.C1 * cmplx.Exp(complex(0, s.K1*alpha + s.Omega1*phi))
      psi2 := s.C2 * cmplx.Exp(complex(0, s.K2*alpha + s.Omega2*phi))
      return psi1 + psi2
  }
  ```
- Amplitude $R = |\Psi|$ is nonconstant, and phase $S = \text{arg}(\Psi)$ has nonconstant gradient.

---

### Guidance & Integration (`internal/bmc/guidance/`)

#### [MODIFY] [integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go)
- Add `RK4Stepper` struct implementing the `Stepper` interface.
- `Step(state model.MiniState, dt float64, vel func(model.MiniState) (float64, float64)) model.MiniState`
  - Standard Runge-Kutta 4th order update:
    $$k_1 = \vec{v}(\vec{q}_n)$$
    $$k_2 = \vec{v}(\vec{q}_n + 0.5 \cdot dt \cdot k_1)$$
    $$k_3 = \vec{v}(\vec{q}_n + 0.5 \cdot dt \cdot k_2)$$
    $$k_4 = \vec{v}(\vec{q}_n + dt \cdot k_3)$$
    $$\vec{q}_{n+1} = \vec{q}_n + \frac{dt}{6}(k_1 + 2k_2 + 2k_3 + k_4)$$

---

### Obstruction Detection (`internal/bmc/obstruction/`)

#### [MODIFY] [detect.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/detect.go)
- **Activate Node Obstruction**:
  - `DetectNodeObstruction(traj model.Trajectory, wf wave.WaveFunction, threshold float64) model.Obstruction`
    - Evaluates $R = |\Psi|$ at all trajectory points. If $R < \text{threshold}$, sets `applies: true` with severity `blocker` and details of node contact.
- **Activate Phase Unwrapping Obstruction**:
  - `DetectPhaseUnwrapObstruction(traj model.Trajectory, wf wave.WaveFunction) model.Obstruction`
    - Evaluates phase gradient along the trajectory. If the step-to-step phase difference exceeds $\pi/2$, phase unwrapping may fail; sets `applies: true` with severity `warning`.

---

### Report Generation & CLI (`internal/bmc/report/` & `cmd/ptw-bmc/`)

#### [MODIFY] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go)
- Change `Parameters` field to `interface{}` to allow serialization of either `PlaneWaveParams` or `SuperpositionParams`.
- Add `GenerateSuperposition(params model.SuperpositionParams, finalTruthClaim bool) (*Report, error)`:
  - Integrates using `RK4Stepper`.
  - Runs all checks, including node detection.
  - Sets technical gate to `bmc0a_superposition_control_gate`.

#### [MODIFY] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go)
- Update `Validate()` to handle the `bmc0a_superposition_control_gate` configuration.
- Ensure validation enforces that the technical gate status is consistent.

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Support `--profile bmc0a-superposition` in the `run` command.
- When this profile is selected, parses default superposition parameters, calls `report.GenerateSuperposition()`, and serializes the report.

---

### Lean Safety Contracts (`BMC/`)

#### [MODIFY] [ToyReport.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ToyReport.lean)
- Add `nodeObstruction : CheckStatus` to the `BMCReport` structure.

#### [MODIFY] [Promotion.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Promotion.lean)
- Add `reportPassesBMC0ASuperpositionGate(r : BMCReport) : Bool`:
  - Requires: `toyAnalysisOnly && !finalTruthClaim && wdwResidual == pass && trajectoryFinite == pass && clockMonotonic == pass && qFinite == pass && classicalLimit == pass && friedmannResidual == deferred && nodeObstruction == pass`.
- Add theorem obligations:
  - `superposition_gate_requires_no_node_obstruction`: `superposition gate = true → nodeObstruction = pass`
  - `superposition_witness_passes`: Decider showing a superposition report witness passes the gate.

---

### Verification and Test Plan

#### [NEW] [internal/bmc/wave/superposition_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/superposition_test.go)
- `TestSuperpositionSatisfiesWdW`: Verify that finite-difference WdW residual of the superposition of two plane waves is within tolerance when both component waves satisfy the constraint.

#### [NEW] [internal/bmc/guidance/integrate_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate_test.go)
- `TestRK4StepperCorrectness`: Verify that `RK4Stepper` integrates a test trajectory with higher accuracy than the Euler stepper.

#### [NEW] [internal/bmc/obstruction/node_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/node_test.go)
- `TestNodeObstructionDetection`: Construct a trajectory that passes near a node, verify `node_obstruction` applies as blocker.
