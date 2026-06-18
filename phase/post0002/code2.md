I have successfully implemented the second remediation ticket: **BMC-POST-0002: Numerical WdW Residual Integration / Analytic Authority Displacement**.

Here is the walkthrough of the files added/modified and the verification results.

### Files Added & Modified
* **Modified Data Model:** Added optional analytic and numerical residual metadata fields to `CheckResult` in [types.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/types.go).
* **Modified WdW Package:** Defined step, tolerance, status, and authority constants in [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go).
* **Modified Report Generation:** Integrated the numerical evaluator into the plane-wave report path and deferred superposition in [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go).
* **Modified Report Validation:** Hardened the validator in [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) to reject inconsistent pass claims.
* **Modified Report Tests:** Added 5 required report-path integration tests in [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go).
* **Added Documentation Note:** Created the remediation note [bmc_post_0002_numerical_wdw_residual_integration.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0002_numerical_wdw_residual_integration.md).
* **Progress Tracking:** Updated the checklist in [task.md](file:///home/chaschel/.gemini/antigravity-ide/brain/37ad27f9-13fb-4063-8c6a-81bd1bae7b12/task.md) and recorded details in [walkthrough.md](file:///home/chaschel/.gemini/antigravity-ide/brain/37ad27f9-13fb-4063-8c6a-81bd1bae7b12/walkthrough.md).

---

### Verification Results

1. **WdW Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
   ```
   * **Result:** `PASS`.

2. **Report Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/report -v -count=1
   ```
   * **Result:** `PASS` (all 12 tests pass, including the new integration tests).

3. **All Go Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./...
   ```
   * **Result:** `ok` (all tests passed).

4. **Lean Build:**
   ```bash
   cd BMC && /home/chaschel/.elan/bin/lake build
   ```
   * **Result:** `Build completed successfully (14 jobs).` with zero warnings.

5. **CLI Verification (Profile: `bmc0a-plane`):**
   * Report generation: `Successfully ran profile 'bmc0a-plane' and generated report.`
   * Report validation: `Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.`
   * Report summary: Outputs numerical residual values and authority (`Numerical Wheeler-DeWitt residual 0.000000e+00 satisfies tolerance 1.000000e-06`).

---

### EBP Debt Status Coordinates

```text
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

* **Recommended Next Ticket:** `BMC-POST-0003: Euler/RK4 and dt Convergence Audit` to compare numerical drift between Euler and RK4 steppers on identical superposition profiles.

> [!WARNING]
> **Strict EBP Reminder:** A numerical residual diagnostic is not a full WdW solver. A valid plane-wave numerical check is not Friedmann recovery. Blocking a wrong constraint is failure detection, not physics validation. Full BMC remains blocked.