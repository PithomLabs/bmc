# BMC Postmortem Remediation Walkthrough (POST-0001 & POST-0002)

This walkthrough documents the remediation work done on the Wheeler-DeWitt (WdW) constraint check logic.

## Sprint remediation logs

---

### BMC-POST-0001: Constraint Violation Detection

#### 1. Files Added & Modified
* [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) (Added `NumericalResidualAt` with strict input checks).
* [residual_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual_test.go) (Added 5 unit tests verifying valid control, violated constraint, and perturbed wavefunction fixtures).
* [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (Added report-level gate checks).
* [bmc_post_0001_constraint_violation_detection.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0001_constraint_violation_detection.md) (Remediation note).

#### 2. Key Accomplishments
* Proved wrong WdW constraint inputs (such as plane waves where $k^2 \neq \omega^2$) fail in isolated package checks.
* Defined physical tolerances explicitly in [residual_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual_test.go).

---

### BMC-POST-0002: Numerical WdW Residual Integration / Analytic Authority Displacement

#### 1. Files Added & Modified
* **Modified Data Model:** [types.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/types.go) (Exposed analytical and numerical residual magnitudes, statuses, and authority scopes on the `CheckResult` struct).
* **Modified WdW Package:** [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) (Defined status and authority constants, step size $10^{-4}$ and tolerance $10^{-6}$).
* **Modified Report Generation:** [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go) (Integrated `NumericalResidualAt` into the plane-wave report path, displacing analytic residual authority, and explicitly deferred superposition checks).
* **Modified Report Validation:** [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) (Added checks to ensure failed numerical checks propagate and block the technical gate).
* **Modified Report Tests:** [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (Added 5 required integration tests).
* **Added Documentation Note:** [bmc_post_0002_numerical_wdw_residual_integration.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0002_numerical_wdw_residual_integration.md).

#### 2. Key Accomplishments
* **Primary Authority Displacement:** The plane-wave report path now uses the numerical evaluator as the primary authority. The analytical residual acts only as an oracle/control.
* **Tolerances:** Step step = $10^{-4}$, tolerance = $10^{-6}$.
* **Superposition Deferral:** Superposition reports now explicitly mark numerical status as `numerical_residual_not_computed` with authority `not_authoritative`, ensuring it is treated as a clear deferral, not a silent success.
* **Disagreement Detection:** If the analytical and numerical magnitudes disagree beyond $10^{-4}$ tolerance, the status is flagged as `NumericalResidualError`.

---

### BMC-POST-0002.1: Artifact Freshness and Validation Consistency Repair

#### 1. Files Added & Modified
* **Modified Report Tests:** [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (Added `TestReportValidationRejectsNumericalResidualErrorClaimedAsPass` consistency check).
* **Modified Documentation:** [bmc_post_0002_numerical_wdw_residual_integration.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0002_numerical_wdw_residual_integration.md) (Clarified analytic/numerical disagreement semantics).
* **Tracked JSON Artifacts Regenerated:**
  - `out/bmc0a_plane.json`
  - `out/bmc0a_superposition_safe.json`
  - `out/bmc0a_superposition_node_probe.json`
  - `out/bmc0a_local_residual.json`
  - `out/bmc0a_residual_audit.json`

#### 2. Key Accomplishments
* **Consistency Check Added:** Added a test asserting that if `numerical_residual_status = numerical_residual_error` but the check is claimed as pass, validation correctly rejects the report.
* **Artifact Refreshment:** Fully regenerated all active JSON review fixtures using the updated CLI build.
  - `out/bmc0a_plane.json` correctly exhibits `numerical_residual_pass` and `diagnostic_authority`.
  - `out/bmc0a_superposition_safe.json` and `out/bmc0a_superposition_node_probe.json` show `numerical_residual_not_computed` and `not_authoritative` deferral metadata.
* **Clarified Semantics:** Update postmortem note to make clear that analytic/numerical disagreement is treated as an oracle-control audit error for the plane-wave path, and that this does not restore analytic residual as primary authority.

---

### BMC-POST-0003: Euler/RK4 and dt Convergence Audit

#### 1. Files Added & Modified
* **New Package Created**: [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence.go) (implements structures, RunAudit, node-contact detection, and lambda aligned pointwise verification).
* **New Tests Added**: [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go) (verifies endpoint drift, same profile checks, finest RK4 as reference, invalid steps, nonfinite trajectories, node contact, and determinism).
* **New Documentation**: [bmc_post_0003_euler_rk4_dt_convergence.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0003_euler_rk4_dt_convergence.md).
* **CLI/Schema**: Intentionally avoided to prevent audit-layer expansion and bloat.

#### 2. Key Accomplishments
* **Fixed Total Lambda Span**: Total span is held constant across all runs (`lambda_span = dt * steps = 10.0`) under the superposition safe profile.
* **Refinements and Steppers**:
  - Euler stepper with dt = 0.05 (steps = 200)
  - RK4 stepper with dt = 0.05 (steps = 200)
  - RK4 stepper with dt/2 = 0.025 (steps = 400)
  - RK4 stepper with dt/4 = 0.0125 (steps = 800) (Finest RK4 reference)
* **Lambda-Aligned Pointwise Comparison**: Compares coarse and refined runs at matching parameter values, asserting `abs(coarse_lambda - ref_lambda) <= 1e-9` at aligned indices.
* **Deterministic Live Trajectory Checks**: Verification tests prove endpoint distance and pointwise drift calculations are computed dynamically from actual trajectories by performing coordinate mutations and asserting expected output drift changes.
* **Node-Contact and Non-Finite Rejection**: Trajectories with node contact are blocked with status `blocked_by_node_contact` and leave distance metrics non-authoritative, but preserve partial coordinates. Non-finite values without node contact block metrics with status `blocked_by_nonfinite_trajectory`.

---

### BMC-POST-0003.2: Remediation Stack Validation and Diff Hygiene Repair

#### 1. Files Added & Modified
* **Modified convergence.go**: [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence.go) (rebuilds `RunAudit` to check and reject non-finite `LambdaStep` inputs (NaN/Inf) with explicit error).
* **Modified convergence_test.go**: [convergence_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/convergence/convergence_test.go) (added test verifying NaN/Inf `LambdaStep` rejection; strengthened the anti-hardcoding test by checking dynamic distance updates under altered trajectory parameters).
* **Modified report validation**: [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) (added checks requiring plane-wave passing reports to include `numerical_residual_status = numerical_residual_pass` and `numerical_residual_authority = diagnostic_authority`).
* **Modified report tests**: [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (added `TestReportValidationRejectsForgedPlaneWaveWdWPass` verifying that forged metadata triggers validation failures).
* **Workspace hygiene**: Reverted unrelated dirty deletes in `phase/`. The untracked directories `faq/`, `phase/post0001/`, `phase/post0002/`, `phase/post0003/` are ignored workspace artifacts and are not part of the PR boundaries.

#### 2. Key Accomplishments
* **Rejection of Nonfinite Step Sizes**: Calling `RunAudit` with a step size of `math.NaN()` or `math.Inf(±1)` now returns an invalid input error instead of drifting into trajectory integration failure.
* **Plane-Wave Validation Protection**: Added strict validation rules ensuring plane-wave WdW check passes require the numeric evaluator to be the active diagnostic authority, preventing forged pass results with missing or incorrect metadata.
* **Stronger Anti-Hardcoding Test**: The convergence package anti-hardcoding test now verifies that distance metrics update correctly by re-running the convergence audit against altered parameters, proving distance calculations are live.

---

## Verification Results

### Go Unit Tests
* **WdW Package Tests**: `go test ./internal/bmc/wdw -v -count=1` -> `PASS`.
* **Report Package Tests**: `go test ./internal/bmc/report -v -count=1` -> `PASS`.
* **Convergence Package Tests**: `go test ./internal/bmc/convergence -v -count=1` -> `PASS`.
* **All Go Tests**: `go test ./...` -> `PASS`.

### Lean Proof Verification
* `cd BMC && lake build` -> `Build completed successfully (14 jobs).` with zero warnings.

---

## EBP 2.1 Debt Status After Integration

```text
BMC-POST-0003 final acceptance: accepted_with_repairs
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

---

### BMC-POST-0004: BMC-0B Massive Scalar Numerical WdW Specification

#### 1. Files Added & Modified
* **New Package Created**: [spec.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/bmc0bspec/spec.go) (Defines spec structs, validator, default spec, required failure modes, required null models, no-solver machine checkable flags, and case-insensitive forbidden term scans).
* **New Tests Added**: [spec_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/bmc0bspec/spec_test.go) (Verifies that solver implementation, trajectory integration, and recovery claims are blocked; confirms required failure modes block promotion; validates grid bounds; tests case-insensitive forbidden phrase rejection with phrase-safe error formatting; checks deterministic serialization).
* **New Documentation**: [bmc_post_0004_bmc0b_massive_scalar_wdw_spec.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0004_bmc0b_massive_scalar_wdw_spec.md).
* **CLI/Schema**: Intentionally avoided (as `new_cli_required: false` and `schema_bump_required: false` are enforced).

#### 2. Key Accomplishments
* **No-Solver Enforcements**: Standardized machine-checkable flags that default to `false` and block EBP promotion if set to `true`.
* **Required Failure Modes**: Implemented all 13 required failure modes, each strictly blocking promotion.
* **Null Model Scaffolding**: Integrated 6 required null models as specifications to check numerical results against reference controls.
* **Phrase-Safe Rejection**: Implemented case-insensitive checks for forbidden terms (e.g. `"validated"`, `"proved"`, `"recovered"`, `"ready"`, `"successful"`) across status fields, claims, failure modes, faithfulness reviews, and boundaries. All error messages are phrase-safe and do not echo the blocked phrases.

---

## Verification Results

### Go Unit Tests
* **WdW Package Tests**: `go test ./internal/bmc/wdw -v -count=1` -> `PASS`.
* **Report Package Tests**: `go test ./internal/bmc/report -v -count=1` -> `PASS`.
* **Convergence Package Tests**: `go test ./internal/bmc/convergence -v -count=1` -> `PASS`.
* **BMC-0B Spec Package Tests**: `go test ./internal/bmc/bmc0bspec -v -count=1` -> `PASS`.
* **All Go Tests**: `go test ./...` -> `PASS`.

### Lean Proof Verification
* `cd BMC && lake build` -> `Build completed successfully (14 jobs).` with zero warnings.

---

## EBP 2.1 Debt Status After Integration

```text
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: partial
needNontrivialPhysicsCase: targeted_by_spec_only
needToyCheck: partial
needFaithfulnessReview: required_before_bmc0b_solver
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
BMC0BStatus: specified_only
OperatorFormDebt: explicit_unpaid
BoundaryConditionDebt: explicit_unpaid
SolverStatus: not_implemented
```

## Next Recommended Ticket
* **BMC-POST-0005: Phase Gradient h-Sensitivity Audit (Priority: P2)**
  - **Scope**: Perform sensitivity audit on phase gradient field calculations to ensure numerical stability and error control under changes in finite difference step size $h$.

---

## Strict EBP Reminders
* A specification is not a solver.
* A finite-difference plan is not a numerical result.
* A massive scalar BMC-0B profile is not Friedmann recovery.
* **Full BMC remains blocked.**




