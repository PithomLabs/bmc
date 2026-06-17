# BMC-POST-0001: Constraint Violation Detection Remediation Note

## Summary
BMC-POST-0001 implements failure-detection capabilities to verify that the Wheeler-DeWitt (WdW) constraint check rejects deliberately invalid wavefunctions and incorrect plane-wave constraint parameters ($k^2 \neq \omega^2$).

## What Was Added
- **Toy Numerical WdW Residual Evaluator:** Implemented `NumericalResidualAt` in [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go) to approximate $(-\partial^2\Psi/\partial\alpha^2 + \partial^2\Psi/\partial\phi^2)$ using central finite differences with step $h$, including strict validation bounds.
- **Negative and Control Tests:** Created [residual_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual_test.go) containing tests with explicit physical tolerances to verify valid control shell wavefunctions, violated constraint plane waves, perturbed wavefunction fixtures, invalid steps, and nonfinite values.
- **Gate Integration Tests:** Added tests to [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) asserting that a failed constraint correctly blocks promotion gates.

---

## Strict EBP Boundaries and Disclaimers

> [!WARNING]
> - **No Physics Validation:** A wrong-input rejection test is not physics validation.
> - **No Complete Solver:** The toy numerical WdW residual evaluator is a finite-difference diagnostic tool for failure detection, not a complete WdW equation solver.
> - **No Friedmann Recovery:** Passing a valid plane-wave control check does not implement or prove Friedmann cosmological recovery.
> - **Full BMC Blocked:** The promotion status remains locked at `blocked` and the faithfulness review remains contested.

---

## EBP Debt Status Coordinates

```text
needConstraintViolationTests: retired_for_plane_wave_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```
