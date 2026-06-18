I have completed the repairs under **BMC-POST-0003.2: Remediation Stack Validation and Diff Hygiene Repair**.

For a full logs overview, please refer to the updated walkthrough artifact:
[walkthrough.md](file:///home/chaschel/.gemini/antigravity-ide/brain/37ad27f9-13fb-4063-8c6a-81bd1bae7b12/walkthrough.md)

---

### Implementation Walkthrough (BMC-POST-0003.2)

#### 1. Files Added & Modified
* **New Package Created**: Yes, `internal/bmc/convergence`.
* **Files Added**:
  * [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence.go)
  * [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go)
  * [bmc_post_0003_euler_rk4_dt_convergence.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0003_euler_rk4_dt_convergence.md)
* **Files Modified**:
  * [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) (Modified to require explicit WdW pass status and diagnostic authority for plane waves).
  * [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (Modified to add tests verifying forged plane-wave WdW check passes fail validation).
* **CLI/Schema**: Intentionally avoided to prevent audit-layer expansion and bloat.

#### 2. Parameters & Reference Trajectory
* **Profile Used**: `bmc0a-superposition-safe` (standard superposition safe params from `DefaultSuperpositionSafeParams`).
* **Step Sizes and Counts** (Enforcing fixed total lambda span = `10.0`):
  * **Euler dt**: $dt = 0.05$, $steps = 200$.
  * **RK4 dt**: $dt = 0.05$, $steps = 200$.
  * **RK4 dt/2**: $dt = 0.025$, $steps = 400$.
  * **RK4 dt/4**: $dt = 0.0125$, $steps = 800$.
* **Reference Trajectory Convention**: Finest RK4 run (RK4 with $dt/4$) is used as the numerical comparison baseline (local numerical reference, not physical ground truth).

#### 3. Rejection of Non-finite Input Step Sizes
* **Step Size (dt) Rejection**: Rebuilt `RunAudit` to check if `params.LambdaStep` is `math.IsNaN(dt)` or `math.IsInf(dt, 0)`. If triggered, it returns an invalid input error instead of allowing a non-finite step size to run.
* **Non-finite Step Size Tests**: Added NaN, +Inf, and -Inf `LambdaStep` rejection verification cases to `TestConvergenceAuditRejectsInvalidStepSize` in [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go).

#### 4. WdW Authority Validation Checks
* **Required Plane-Wave WdW pass fields**: Extended `Validate` in [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) to require a passing plane-wave WdW check (`StatusPass`) to contain:
  * `numerical_residual_status == "numerical_residual_pass"`
  * `numerical_residual_authority == "diagnostic_authority"`
* **Adversarial Tests**: Added `TestReportValidationRejectsForgedPlaneWaveWdWPass` in [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) to verify that forged pass claims with missing or incorrect metadata (e.g. `numerical_residual_not_computed` or `not_authoritative`) are rejected.

#### 5. Strengthening Anti-Hardcoding Test
* The package anti-hardcoding test in [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go#L63-L81) has been strengthened to verify that distance metrics are computed dynamically from actual trajectories by executing the audit against altered parameters and asserting that the distance changes.

#### 6. Diff & Workspace Hygiene
* Unrelated deleted files in `phase/` have been restored. Other untracked files and directories (`faq/`, `phase/post0001/`, etc.) are ignored workspace artifacts.

#### 7. Verification Results
* **`go test ./internal/bmc/convergence -v -count=1`**: `PASS`
* **`go test ./...`**: `PASS`
* **`lake build` (Lean)**: `PASS` (Build completed successfully with 14 jobs and zero warnings).

---

### Expected EBP Status After Implementation
```text
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: targeted
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
PlaneWaveWDWAuthority: validated
ConvergenceAuditScope: validated
DiffBoundaryHygiene: clean
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

### Remaining Limitations & Status
* Convergence is numerical self-consistency evidence, not physics correctness or Friedmann recovery validation.
* **Is POST-0003 acceptable now?** Yes, both POST-0003 and the POST-0003.2 repairs are complete and verified.
* **Next recommended remediation ticket**: `BMC-POST-0004: Massive Scalar Numerical Wheeler-DeWitt Equation Specification (Priority: P2)` (defining the finite difference schema, grid dimensions, boundary conditions, and gate requirements for the massive scalar model (BMC-0B) without implementing the solver yet).