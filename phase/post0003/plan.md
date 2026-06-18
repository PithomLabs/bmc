# Implementation Plan - BMC-POST-0003: Euler/RK4 and dt Convergence Audit

Measure whether trajectory outputs are numerically stable across Euler/RK4 steppers and time-step refinement on the same superposition profile. This is a numerical reliability remediation ticket, not a physics promotion ticket.

## User Review Required

> [!IMPORTANT]
> - **Test-Level Remediation Only**: To prevent audit-layer bloat, this convergence audit is implemented strictly as a test-level package check without introducing a new JSON report schema or command-line interface.
> - **Same Superposition Profile**: Uses the existing safe superposition profile parameters (`bmc0a-superposition-safe`) across all comparison runs.
> - **Reference Trajectory Convention**: Uses the finest RK4 run (dt/4) as the numerical comparison baseline (local numerical reference, not physical ground truth).

## Proposed Changes

### Component: `internal/bmc/convergence`

#### [NEW] [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence.go)
Create a new Go package `internal/bmc/convergence` defining the convergence structures and calculation logic.

- Define `ConvergenceRun` and `ConvergenceAudit` structs as requested:
  ```go
  type ConvergenceRun struct {
  	RunID                       string   `json:"run_id"`
  	Stepper                     string   `json:"stepper"`
  	DeltaLambda                 float64  `json:"delta_lambda"`
  	Steps                       int      `json:"steps"`
  	TrajectoryFinite            bool     `json:"trajectory_finite"`
  	NodeContactDetected         bool     `json:"node_contact_detected"`
  	FinalAlpha                  *float64 `json:"final_alpha"`
  	FinalPhi                    *float64 `json:"final_phi"`
  	EndpointDistanceToReference *float64 `json:"endpoint_distance_to_reference,omitempty"`
  	MaxPointwiseDistanceToReference *float64 `json:"max_pointwise_distance_to_reference,omitempty"`
  	Status                      string   `json:"status"`
  	Notes                       []string `json:"notes"`
  }

  type ConvergenceAudit struct {
  	ProfileID            string           `json:"profile_id"`
  	ToyAnalysisOnly      bool             `json:"toy_analysis_only"`
  	PhysicsClaim         string           `json:"physics_claim"`
  	ReferenceRunID       string           `json:"reference_run_id"`
  	Runs                 []ConvergenceRun `json:"runs"`
  	InterpretationStatus string           `json:"interpretation_status"`
  	Warnings             []string         `json:"warnings"`
  }
  ```
- Implement `RunAudit(params model.SuperpositionParams) (*ConvergenceAudit, error)` which:
  - Validates `delta_lambda` and `steps`. If `delta_lambda <= 0` or `steps <= 0`, blocks or returns error/invalid run.
  - Runs four trajectories using `guidance.Integrate` on the same wave function:
    1. Euler stepper with `dt`
    2. RK4 stepper with `dt`
    3. RK4 stepper with `dt/2`
    4. RK4 stepper with `dt/4` (Reference Run)
  - Detects node contact: Checks if at any point the wavefunction amplitude falls below the node threshold.
  - Performs finiteness check: Checks if any coordinate in the trajectory is NaN or Inf.
  - Computes `endpoint_distance_to_reference` using the formula `sqrt((alpha_run - alpha_ref)^2 + (phi_run - phi_ref)^2)`.
  - Computes `max_pointwise_distance_to_reference` by comparing aligned points: index $i$ in run corresponds to index $i \cdot \text{stepRatio}$ in reference.
  - Employs strict EBP warnings and non-claims in the audit.

#### [NEW] [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go)
Implement tests:
- `TestConvergenceAuditComputesEulerRK4EndpointDrift`: Verifies endpoint drift calculations and ensures the drift is computed from actual trajectories (e.g. by mutating or changing parameters and asserting change in drift).
- `TestConvergenceAuditUsesSameProfileForAllSteppers`: Asserts all runs share the same core profile parameters.
- `TestConvergenceAuditUsesFinestRK4AsReference`: Verifies RK4 with `dt/4` is the reference run and distance to itself is zero.
- `TestConvergenceAuditRejectsInvalidStepSize`: Verifies invalid step size is rejected/blocked.
- `TestConvergenceAuditRejectsNonfiniteTrajectory`: Verifies NaN/Inf coordinates are detected and block convergence metrics.
- `TestConvergenceAuditBlocksNodeContact`: Verifies node contact is detected and blocks metrics.
- `TestConvergenceAuditDeterministic`: Verifies that successive runs of the audit return identical results.

### Component: Documentation

#### [NEW] [bmc_post_0003_euler_rk4_dt_convergence.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0003_euler_rk4_dt_convergence.md)
Document that:
- BMC-POST-0003 compares numerical trajectories across Euler/RK4 and dt refinements.
- This is numerical reliability evidence, not physics validation.
- The finest RK4 run is a local numerical reference, not physical ground truth.
- This does not implement Friedmann recovery.
- This does not validate BMC physics.
- Full BMC remains blocked.

## Verification Plan

### Automated Tests
1. **Convergence Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/convergence -v -count=1
   ```
2. **Guidance Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/guidance -v -count=1
   ```
3. **All Go Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./... -count=1
   ```
4. **Lean build verification:**
   ```bash
   cd BMC && /home/chaschel/.elan/bin/lake build
   ```
