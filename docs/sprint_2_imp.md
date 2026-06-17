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
- `TestRK4StepperCorrectness`: Verifies RK4 is higher accuracy than Euler on a linear field.
- `TestNodeObstructionDetection`: Constructs a trajectory containing a known node, verifying that node contact is correctly flagged.
- `TestSafeSuperpositionPassesChecks`: Generates the safe profile, verifying all checks pass.
- `TestNodeProbeShortCircuitAndValidationGate`: Runs the node-probe profile (starts at $\alpha_0=0, \phi_0=0$ which guarantees $\Psi = 0$). Verifies trajectory is short-circuited (empty), `node_contact_free` fails, and technical gate routes to `node_detection_validation_gate` with status `pass`.

All 14 Go tests passed successfully.

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


# BMC Sprint 2 Task List

- `[x]` 1. Implement Go Configuration & Wavefunction
  - `[x]` Update `internal/bmc/model/params.go` with `SuperpositionParams` & validation for `MaxPhaseGrad`
  - `[x]` Create `internal/bmc/wave/superposition.go` for `SuperpositionWave`
- `[x]` 2. Implement Stepper, Velocity, & WdW Updates
  - `[x]` Add `RK4Stepper` in `internal/bmc/guidance/integrate.go`
  - `[x]` Update `internal/bmc/guidance/velocity.go` to check `node_threshold` and refuse near-node computations
  - `[x]` Add `ComponentResidual` & `ComponentResidualsSuperposition` in `internal/bmc/wdw/residual.go`
- `[x]` 3. Implement Obstruction Detection
  - `[x]` Update `internal/bmc/obstruction/detect.go` with `node_detection`, `node_contact_free`, `q_finite_away_from_nodes`, and `phase_gradient_finite` checks
- `[x]` 4. Update Report, Validation, and CLI
  - `[x]` Define `ReportParameters` struct and modify `Report` in `internal/bmc/report/report.go`
  - `[x]` Implement initial node short-circuit logic in `report.go`
  - `[x]` Update `internal/bmc/report/validate.go` to support new checks and profiles (`bmc0a-superposition-safe` vs `bmc0a-superposition-node-probe`)
  - `[x]` Update `cmd/ptw-bmc/main.go` to support the two new profiles
- `[x]` 5. Update Lean Safety Contracts
  - `[x]` Update `BMC/BMC/ToyReport.lean` structure with the new fields
  - `[x]` Update `BMC/BMC/Promotion.lean` with safe gate, validation gate, and safety theorems
- `[x]` 6. Write Sprint 2 Go Tests
  - `[x]` Create `internal/bmc/wave/superposition_test.go`
  - `[x]` Create `internal/bmc/guidance/integrate_test.go` (testing RK4 stepper)
  - `[x]` Create `internal/bmc/obstruction/node_test.go` (testing safe trajectory vs node-probe short-circuit)
- `[x]` 7. Verification
  - `[x]` Run `go test ./...`
  - `[x]` Build CLI: `go build -buildvcs=false ./cmd/ptw-bmc`
  - `[x]` Run both safe and node-probe CLI profiles
  - `[x]` Validate both generated JSON reports
  - `[x]` Summarize both reports
  - `[x]` Run `cd BMC && lake build` to check Lean proofs
- `[x]` 8. Walkthrough
  - `[x]` Create the Sprint 2 walkthrough artifact


