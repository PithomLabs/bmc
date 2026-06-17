# BMC-POST-0002: Numerical WdW Residual Integration / Analytic Authority Displacement

## Summary
BMC-POST-0002 moves the toy numerical WdW residual evaluator into the actual report/simulation path. The analytic plane-wave residual now serves only as an oracle/control, and the numerical residual becomes the primary diagnostic authority for WdW constraint failure detection in plane waves.

## What Was Modified & Integrated
- **Exposed CheckResult Metadata:** Extended the `CheckResult` struct in [types.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/types.go) with optional fields carrying numerical residual magnitudes, tolerances, statuses, and authority scopes.
- **Integrated Evaluator:** Updated `Generate` in [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go) to compute both analytic and numerical residuals. If they disagree beyond numerical tolerance ($1e-4$), the check is flagged as `NumericalResidualError`. Otherwise, the numerical residual is evaluated against `WDWNumericalResidualTolerance` ($1e-6$).
- **Report Validation & Gate Checks:** Updated `Validate` in [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/validate.go) to verify that if `NumericalResidualStatus` equals `numerical_residual_violation_detected` or `numerical_residual_error`, the WdW check must fail, ensuring that failed constraints propagate and block the technical control gate.
- **Superposition Deferral:** Superposition reports now explicitly mark numerical status as `numerical_residual_not_computed` with authority `not_authoritative`, documenting the deferral of numerical evaluation for superpositions.
- **Integration Tests:** Appended five required integration tests to [report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go).

---

## EBP status expected after implementation

```text
needConstraintViolationTests: retired_for_plane_wave_report_path_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

---

## Strict EBP Boundaries and Disclaimers

> [!WARNING]
> - **No Complete WdW Solver:** The integrated toy numerical residual diagnostic is not a complete Wheeler-DeWitt solver.
> - **No Friedmann Recovery:** Reaching a valid plane-wave numerical check does not implement, prove, or suggest Friedmann cosmological recovery.
> - **No Physics Validation:** Rejects wrong constraints (failure detection) but does not validate Bohmian Minisuperspace Cosmology physics.
> - **Full BMC remains blocked.**
