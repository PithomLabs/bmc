# BMC Sprint 2 Walkthrough: Two-Plane-Wave Superposition Control

We have successfully implemented and verified the two-plane-wave superposition control model (**BMC-0A**) in **Go + Lean**, obeying all EBP 2.1 protocol requirements and including the initial node short-circuit guard.

---

## 1. What was Implemented

### Go Configuration & Wavefunction (`internal/bmc/`)
- [params.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/params.go): Added `SuperpositionParams` supporting two-component plane-wave superpositions with parameters ($c_1, c_2, k_1, \omega_1, k_2, \omega_2$), initial state $(\alpha_0, \phi_0)$, $\Delta\lambda$, integration steps, `NodeThresh`, and `MaxPhaseGrad`. Implemented rigorous range validation.
- [superposition.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/superposition.go): Formulated the superposition wave function $\Psi(\alpha, \phi) = c_1 \exp(i(k_1\alpha + \omega_1\phi)) + c_2 \exp(i(k_2\alpha + \omega_2\phi))$.
- [integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go): Implemented `RK4Stepper` supporting fourth-order Runge-Kutta updates.
- [velocity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/velocity.go) & [trajectory.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/trajectory.go): Refuse velocity evaluation and short-circuit integration if wavefunction amplitude $R < \text{node\_threshold}$, replacing undefined gradient steps with NaN values safely.
- [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go): Implemented component-wise WdW residual checks ($k_j^2 - \omega_j^2$).
- [detect.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/detect.go): Implemented active detectors for `node_detection`, `node_contact_free`, `q_finite_away_from_nodes`, and `phase_gradient_finite`.
- [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go): Swapped interface parameters for a structured, typed `ReportParameters` wrapper. Implemented the initial node short-circuit guard.

### Lean Formalization (`BMC/`)
- [ToyReport.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ToyReport.lean) & [Promotion.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Promotion.lean): Updated checks structure and safety gates.
  - Defined `reportPassesBMC0ASuperpositionSafeGate` requiring all checks (excluding Friedmann) to pass.
  - Defined `reportPassesBMC0ANodeDetectionGate` requiring only node detection and WdW residual to pass.
  - Verified promotion safety theorems for both gates and tested them against safe and node-probe witnesses.

---

## 2. What was Tested & Verified

### Automated Go Tests
- `TestSuperpositionSatisfiesWdW`: Checks component-wise residuals.
- `TestRK4StepperCorrectness`: Verifies RK4 has strictly lower error than Euler on a non-constant ODE ($d\alpha/d\lambda=\alpha, d\phi/d\lambda=-\phi$).
- `TestNodeObstructionDetection`: Constructs a trajectory containing a known node, verifying that node contact is correctly flagged.
- `TestSafeSuperpositionAmplitudeStaysAboveNodeThreshold`: Samples all points along the safe profile trajectory and independently asserts that the amplitude stays above the node threshold.
- `TestSafeSuperpositionPassesChecks`: Generates the safe profile, verifying all checks pass.
- `TestNodeProbeShortCircuitAndValidationGate`: Runs the node-probe profile (starts at $\alpha_0=0, \phi_0=0$ which guarantees $\Psi = 0$). Verifies that the trajectory is short-circuited (empty with 0 points count), `node_detection` passes, `node_contact_free` fails, a `node_obstruction` is filed with blocker severity, the safe superposition gate does not pass, the technical gate matches `node_detection_validation_gate` with status `pass`, and the full BMC toy gate remains `blocked`.

All Go tests passed successfully.

### CLI Profile Runs & validation
1. **`bmc0a-superposition-safe`**:
   - Technical gate: `bmc0a_superposition_safe_gate` (status: `pass`)
   - Promotion gate: `blocked`
   - Individual checks pass.
2. **`bmc0a-superposition-node-probe`**:
   - Technical gate: `node_detection_validation_gate` (status: `pass`)
   - Trajectory: fail / short-circuited.
   - Obstructions: `node_obstruction`, `phase_unwrap_obstruction`, `nonfinite_q_obstruction`, `clock_nonmonotonicity_obstruction` all actively trigger and blocker severity applies.
   - Promotion gate: `blocked`.

Both generated reports were successfully validated and summarized.

### Lean compilation
Ran `/home/chaschel/.elan/bin/lake build` to compile the safety contracts and witness deciders:
```text
Build completed successfully (5 jobs).
```
Zero errors, warnings, sorries, or admits.

---

## 3. Results Summary

- **Safe Superposition profile**: Successfully integrates without node contact.
- **Node-Probe Superposition profile**: Correctly short-circuits at the initial step.
- **Lean Promotion Gates**: Deciders show the safe witness passes the safe gate, and the node-probe witness passes the detection validation gate but fails the safe gate. Both are blocked from promotion to full QG.

---

## 4. BMC Sprint 3: Numerical Robustness & Convergence Audit

We have successfully implemented and verified the numerical robustness and convergence audit for Sprint 3.

### What was Implemented

1. **Audit Package (`internal/bmc/audit/`)**:
   - [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/convergence.go): Implemented step-size convergence sweep using a fixed integration interval ($T = 10.0$) with step sizes `[0.1, 0.05, 0.025, 0.0125]`. Drift is computed relative to the finest run.
   - [threshold_sensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/threshold_sensitivity.go): Sweeps thresholds `[1e-4, 1e-5, 1e-6]` on both profiles.
   - [phase_bound_sensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/phase_bound_sensitivity.go): Sweeps phase gradient bounds `[25, 50, 100, 200]`.
   - [perturbation.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/perturbation.go): Performs parameter grid perturbations (nested loop: `c2`, `k2`) and node-probe offsets.
   - [robustness_report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/robustness_report.go): Defines `RobustnessReport` schema, strict validation with unknown-field rejection, and custom validation/summarization logic. Emits results in deterministic sorted orders.
2. **CLI Subcommand Integration**:
   - Added `ptw-bmc audit` to run and output the JSON report.
   - Routed `validate` and `summarize` subcommands to robustness-specific logic based on schema version.
3. **Lean Policy Gates**:
   - [Robustness.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Robustness.lean): Formalized report policies, the robustness technical gate (`reportPassesBMC0ARobustnessAuditGate`), and theorems proving that audit success does not imply full quantum gravity/Bohmian mechanics recovery, keeping the final truth claim blocked.

### What was Tested & Verified

- **Go Unit Tests**:
  - `TestRobustnessReportValidation`: Verifies strict decoding (rejecting unknown fields) and rejects final truth claims.
  - `TestStepSizeConvergence`: Confirms deterministic execution and finite drift calculations.
  - `TestThresholdSensitivity`: Verifies safe and node-probe threshold-aware gate behaviors.
  - `TestPhaseGradientSensitivity`: Verifies status consistency under different bounds.
  - `TestParameterPerturbations`: Confirms grid perturbations and offset sweeps.
- **Lean Compilation**:
  - `lake build` completed successfully with zero errors, warnings, sorries, or admits.
- **CLI Commands**:
  - Report generated via `ptw-bmc audit`.
  - Report validated via `ptw-bmc validate` (passed).
  - Report summarized via `ptw-bmc summarize` (outcome: `mixed` due to clock monotonicity failure under specific perturbations, technical gate: `pass`).



# BMC Sprint 3 Task List

- `[x]` 1. Create Go Audit Package Files
  - `[x]` Implement `internal/bmc/audit/robustness_report.go` (Report structures, ValidateRobustnessReport, and SummarizeRobustnessReport with strict JSON decoding)
  - `[x]` Implement `internal/bmc/audit/convergence.go` (Step-size convergence sweep, T=10.0, step sizes 0.1, 0.05, 0.025, 0.0125, steps 100, 200, 400, 800)
  - `[x]` Implement `internal/bmc/audit/threshold_sensitivity.go` (Threshold sweeps 1e-4, 1e-5, 1e-6)
  - `[x]` Implement `internal/bmc/audit/phase_bound_sensitivity.go` (Phase bounds 25, 50, 100, 200)
  - `[x]` Implement `internal/bmc/audit/perturbation.go` (c2, k2 perturbations, and threshold-aware node offsets)
- `[x]` 2. Integrate with CLI
  - `[x]` Add `audit` subcommand in `cmd/ptw-bmc/main.go`
  - `[x]` Route `validate` and `summarize` subcommands to robustness-specific logic based on schema version
- `[x]` 3. Implement Lean 4 Safety Gates
  - `[x]` Create `BMC/BMC/Robustness.lean` with report structure, audit gate logic, and policy-safety theorems
- `[x]` 4. Implement Go Audit Unit Tests
  - `[x]` Create `internal/bmc/audit/audit_test.go` verifying step-size convergence, threshold sensitivity, phase sensitivity, perturbations, and strict validation
- `[x]` 5. Verification Run
  - `[x]` Run `go test ./...`
  - `[x]` Rebuild CLI and run `ptw-bmc audit`
  - `[x]` Validate and summarize the robustness audit report
  - `[x]` Run `cd BMC && lake build` to verify Lean policies compile cleanly
- `[x]` 6. Walkthrough
  - `[x]` Create the Sprint 3 walkthrough artifact
