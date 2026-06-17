I have successfully implemented the first remediation ticket: **BMC-POST-0001: Constraint Violation Detection**.

Here is the walkthrough of the files added/modified and the verification results.

### Files Added & Modified
* **Numerical Residual Evaluator Logic:** Exposed `NumericalResidualAt` in [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) to numerically evaluate $(-\partial^2\Psi/\partial\alpha^2 + \partial^2\Psi/\partial\phi^2)$ using central finite differences on the wavefunction.
* **Negative & Control Tests:** Created [residual_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual_test.go) with explicit tolerances to verify that physics-violating inputs fail validation.
* **Gate Propagation Tests:** Appended two tests in [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) to verify that failed checks block promotion.
* **Documentation Note:** Created the remediation note [bmc_post_0001_constraint_violation_detection.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0001_constraint_violation_detection.md).
* **Progress Tracking:** Updated the checklist in [task.md](file:///home/chaschel/.gemini/antigravity-ide/brain/37ad27f9-13fb-4063-8c6a-81bd1bae7b12/task.md) and recorded details in [walkthrough.md](file:///home/chaschel/.gemini/antigravity-ide/brain/37ad27f9-13fb-4063-8c6a-81bd1bae7b12/walkthrough.md).

---

### Verification Results

1. **WdW Package Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
   ```
   * **Result:** `PASS` (5 tests completed successfully).
     * `TestWDWNumericalResidualAcceptsConstraintShellPlaneWave`
     * `TestWDWNumericalResidualRejectsWrongPlaneWaveConstraint`
     * `TestWDWNumericalResidualRejectsInvalidWavefunctionFixture`
     * `TestNumericalResidualRejectsInvalidStep`
     * `TestNumericalResidualRejectsNonfiniteInputs`

2. **All Go Tests:**
   ```bash
   GOCACHE=/tmp/go-build-cache go test ./...
   ```
   * **Result:** `ok` (all tests passed).

3. **Lean Build:**
   ```bash
   cd BMC && /home/chaschel/.elan/bin/lake build
   ```
   * **Result:** `Build completed successfully (14 jobs).` with zero warnings or `sorry` statements.

---

### EBP Debt Status Coordinates

```text
needConstraintViolationTests: retired_for_plane_wave_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

* **Recommended Next Ticket:** `BMC-POST-0002: Independent Numerical WdW Residual Evaluator` to integrate the evaluator directly into the simulation pipeline, displacing the analytic plane-wave control oracle for all superposition wavefunctions.

> [!WARNING]
> **Strict EBP Reminder:** A wrong-input rejection test is not physics validation. The toy numerical residual evaluator is a finite-difference diagnostic tool for failure detection, not a complete WdW solver. Full BMC remains blocked.