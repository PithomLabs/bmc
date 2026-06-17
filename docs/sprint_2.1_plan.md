# BMC Sprint 2: BMC-0A Two-Plane-Wave Superposition Control

Stress-test the existing BMC-0A pipeline when $\Psi$ has nonconstant amplitude, possible nodes, nonzero quantum potential $Q$, and nonconstant Bohmian velocity.

## User Review Required

> [!IMPORTANT]
> **Proposed Superposition Profiles**:
>
> 1. `bmc0a-superposition-safe` (No node contact, Q is finite and computable, technical gate may pass):
>    - $c_1 = 1.0 + 0i, k_1 = 1.0, \omega_1 = 1.0$
>    - $c_2 = 0.5 + 0i, k_2 = 2.0, \omega_2 = -2.0$
>    - Initial state: $\alpha_0 = 0.0, \phi_0 = 0.0$
>    - Integration: $\Delta\lambda = 0.05, N = 200$ steps
>    - Node threshold: $10^{-5}$
>
> 2. `bmc0a-superposition-node-probe` (Intentionally crosses/approaches a node to test obstruction detection; blocks safe technical gate but passes node-detection validation):
>    - $c_1 = 1.0 + 0i, k_1 = 1.0, \omega_1 = 1.0$
>    - $c_2 = -1.0 + 0i, k_2 = 2.0, \omega_2 = -2.0$ (equal amplitudes with opposite sign)
>    - Initial state: $\alpha_0 = 0.0, \phi_0 = 0.1$
>    - Integration: $\Delta\lambda = 0.05, N = 200$ steps
>    - Node threshold: $10^{-5}$

---

## Proposed Changes

### Model Types and Configuration (`internal/bmc/model/`)

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
      MaxPhaseGrad float64 `json:"max_phase_gradient"` // Configured bound for phase obstruction
  }
  ```
- Add `DefaultSuperpositionSafeParams()` and `DefaultSuperpositionNodeProbeParams()`.
- Add `Validate()` method on `SuperpositionParams` checking:
  - Both component constraints: $\omega_1^2 = k_1^2$ and $\omega_2^2 = k_2^2$ hold within tolerance.
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

---

### Guidance & Integration (`internal/bmc/guidance/`)

#### [MODIFY] [integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go)
- Add `RK4Stepper` struct implementing the `Stepper` interface:
  - Updates the configuration using standard Runge-Kutta 4th order.

---

### Wheeler-DeWitt Residual (`internal/bmc/wdw/`)

#### [MODIFY] [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go)
- Add analytic residual check for the superposition:
  ```go
  func AnalyticResidualSuperposition(k1, omega1, k2, omega2 float64) float64
  ```
  Returns `k1² - omega1² + k2² - omega2²`.
- The report will include:
  ```json
  "component_residuals": [
    { "component": 1, "k": 1, "omega": 1, "residual": 0 },
    { "component": 2, "k": 2, "omega": -2, "residual": 0 }
  ]
  ```

---

### Obstruction Detection (`internal/bmc/obstruction/`)

#### [MODIFY] [detect.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/detect.go)
- **Node Detection**:
  - `node_detection` check status is `pass` if the detection logic executes successfully.
  - `node_contact_free` check status is `pass` if $R = |\Psi| \ge \text{node\_threshold}$ along the entire trajectory. If $R < \text{node\_threshold}$ at any point, `node_contact_free` is `fail/blocker`.
- **Phase Gradient Obstruction**:
  - If $R < \text{node\_threshold}$, the phase gradient is marked `contested`.
  - If the phase gradient magnitude $\sqrt{(\partial_\alpha S)^2 + (\partial_\phi S)^2}$ becomes non-finite or exceeds `max_phase_gradient`, the `phase_gradient_finite` check is `fail/blocker`.
- **Quantum Potential Obstruction**:
  - `q_finite_away_from_nodes` check status is `pass` if $Q$ remains finite for all trajectory points where $R \ge \text{node\_threshold}$.

---

### Report Generation & Validation (`internal/bmc/report/`)

#### [MODIFY] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go)
- Introduce a typed wrapper struct instead of `interface{}` to keep parameters checkable:
  ```go
  type ReportParameters struct {
      PlaneWave     *model.PlaneWaveParams     `json:"plane_wave,omitempty"`
      Superposition *model.SuperpositionParams `json:"superposition,omitempty"`
  }
  ```
- Change `Parameters` field in `Report` to `ReportParameters`.
- Add checks map fields matching the new Lean structure: `node_detection`, `node_contact_free`, `q_finite_away_from_nodes`, `phase_gradient_finite`.

#### [MODIFY] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go)
- Validate status-pass consistency for the new checks.
- For `bmc0a-superposition-safe` profile: require all checks (except Friedmann) to be `pass`.
- For `bmc0a-superposition-node-probe` profile: require `node_detection` to be `pass` and `node_contact_free` to be `fail/blocker`. Technical gate status must be consistent.

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Add support for both `--profile bmc0a-superposition-safe` and `--profile bmc0a-superposition-node-probe`.

---

### Lean Safety Contracts (`BMC/`)

#### [MODIFY] [BMC/ToyReport.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ToyReport.lean)
- Update `BMCReport` structure:
  ```lean
  structure BMCReport where
    toyAnalysisOnly      : Bool
    finalTruthClaim      : Bool
    wdwResidual          : CheckStatus
    trajectoryFinite     : CheckStatus
    clockMonotonic       : CheckStatus
    nodeDetection        : CheckStatus
    nodeContactFree      : CheckStatus
    qFiniteAwayFromNodes : CheckStatus
    phaseGradientFinite  : CheckStatus
    classicalLimit       : CheckStatus
    friedmannResidual    : CheckStatus
    faithfulness         : CheckStatus
  ```

#### [MODIFY] [BMC/Promotion.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Promotion.lean)
- Define the safe superposition gate:
  ```lean
  def reportPassesBMC0ASuperpositionSafeGate (r : BMCReport) : Bool :=
    r.toyAnalysisOnly &&
    !r.finalTruthClaim &&
    checkPassed r.wdwResidual &&
    checkPassed r.trajectoryFinite &&
    checkPassed r.clockMonotonic &&
    checkPassed r.nodeDetection &&
    checkPassed r.nodeContactFree &&
    checkPassed r.qFiniteAwayFromNodes &&
    checkPassed r.phaseGradientFinite &&
    checkPassed r.classicalLimit &&
    checkDeferred r.friedmannResidual
  ```
- Define the node-detection-only validation gate:
  ```lean
  def reportPassesBMC0ANodeDetectionGate (r : BMCReport) : Bool :=
    r.toyAnalysisOnly &&
    !r.finalTruthClaim &&
    checkPassed r.nodeDetection &&
    checkPassed r.wdwResidual &&
    checkDeferred r.friedmannResidual
  ```
- Add safety theorems verifying that even with valid node detection, both gates block promotion to a full quantum gravity claim.

---

### Verification and Test Plan

#### [NEW] [internal/bmc/wave/superposition_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/superposition_test.go)
- `TestSuperpositionSatisfiesWdW`: Verify that component residuals are 0 and analytic superposition residual is 0.

#### [NEW] [internal/bmc/guidance/integrate_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate_test.go)
- `TestRK4StepperCorrectness`: Verify that `RK4Stepper` integrates a test trajectory with higher accuracy than the Euler stepper.

#### [NEW] [internal/bmc/obstruction/node_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/node_test.go)
- `TestNodeObstructionDetection`: Run the `node-probe` profile, verify that `node_contact_free` is flagged as `fail/blocker` but `node_detection` is `pass`.
- `TestSafeSuperpositionPassesChecks`: Run the `safe` profile, verify all technical checks pass.
