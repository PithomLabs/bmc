# Walkthrough - BMC-POST-0007: Branch-Cut and Stencil Boundary Regression Hardening

This walkthrough documents the completed implementation of the regression hardening tests for phase-gradient branch-cuts and quantum-potential stencil-boundaries.

## Changes Made

### Files Added
- [bmc_post_0007_branchcut_stencil_regression_hardening.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0007_branchcut_stencil_regression_hardening.md)

### Files Modified
- [hsensitivity_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity_test.go)
- [domain_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/domain_test.go)

### Detailed Status
- **Phase-gradient branch-cut test**: Added `TestHSensitivityBranchCutSafeGradientDoesNotArgWrap` which uses the wavefunction $\Psi = \exp(i (\pi - 0.5 \alpha))$. Verifies that the numerical gradient near the phase-wrap boundary at $\alpha = 0$ remains authoritative and close to $-0.5$ (with $\frac{\partial S}{\partial \phi} \approx 0$).
- **Q-potential stencil-boundary test**: Added `TestQPotentialBlocksStencilPointBelowAmplitudeFloor`. Uses a mock wavefunction where the evaluation center is safe ($R = 1.0$) but exactly one stencil point ($\alpha + h$) is below `NearNodeAmplitudeFloor`. Verifies that the evaluation is correctly flagged `Authoritative = false` with status `q_potential_blocked_by_domain_boundary`.
- **Top-level invalid derivative step test**: Renamed test to `TestQPotentialRunAuditRejectsInvalidDerivativeStep` and verified it covers nonpositive $h$, NaN $h$, and Inf $h$.
- **Forbidden phrase scan**: Ran the scan to check that no status, schema, or error messages contain forbidden phrases (`validated`, `proved`, `recovered`, `ready`, `successful`, etc.).
- **Diff hygiene status**: Clean. Git diff shows no modifications to existing tracked code files, only to test files and the new postmortem markdown.

## Verification Results

- **`go test ./internal/bmc/phaseaudit`**: PASS
- **`go test ./internal/bmc/qpotential`**: PASS
- **`go test ./internal/bmc/bmc0bspec`**: PASS
- **`go test ./internal/bmc/wdw`**: PASS
- **`go test ./internal/bmc/report`**: PASS
- **`go test ./internal/bmc/convergence`**: PASS
- **`go test ./...`**: PASS
- **`lake build`**: PASS

## Summary and Next Steps

- **Remaining limitations**: The audit remains strictly test-first regression hardening. No solver, trajectory integration, or physics promotion is added.
- **Whether POST-0007 is ready for review**: Yes, all regression tests have been added and verified.
- **Next recommended ticket**: Conclude final review of the audit package.
