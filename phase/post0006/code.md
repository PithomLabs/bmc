# Walkthrough - BMC-POST-0006: Quantum Potential Near-Node Domain-Boundary Audit

This walkthrough documents the verified implementation of the domain-boundary audit of quantum-potential evaluations.

## Changes Made

### Files Added
- [domain.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/domain.go)
- [domain_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/domain_test.go)
- [bmc_post_0006_qpotential_near_node_domain_boundary.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md)

### Files Modified
- None.

### Detailed Status
- **Existing Q-potential behavior found**: The existing `Q` function in `q.go` silently clamped to `0.0` near nodes (`rVal < 1e-12`).
- **Whether near-node Q previously clamped or defaulted**: Yes, it clamped to `0.0` with no indication of authority loss.
- **Domain-boundary authority model added**: Wrapped evaluations in `domain.go` so that the audit returns metadata including `Authoritative` and `Status` fields.
- **Amplitude floor used**: Enforces `NearNodeAmplitudeFloor = 1e-8`.
- **Derivative step used**: `QPotentialDerivativeStep = 1e-4`.
- **Stencil/domain checks added**: Evaluates the wavefunction amplitude at the evaluation point and all offset stencil points. If any point falls below the floor, evaluations are flagged non-authoritative.
- **Node/nonfinite blocking behavior**: Any exact node contact ($R < 10^{-12}$) or stencil coordinate that crosses a node-sensitive boundary immediately blocks the calculation, marking the point as non-authoritative. Nonfinite wavefunctions or derivatives are similarly blocked.
- **Tests added**: Table-driven tests validating plane wave authority, superposition points audit-only constraints, node contact blocking, zero clamping rejection, invalid input validations, nonfinite derivative blocking, determinism, and lack of promotion fields.
- **Forbidden inference audit**: Verifies that status, schema, and error fields contain no forbidden words (`validated`, `proved`, `recovered`, `ready`, `successful`, etc.).
- **Documentation added**: Created a detailed postmortem document explaining node sensitivity, thresholds, and EBP promotion blocks.
- **Diff hygiene status**: Clean. Git diff shows no modifications to existing tracked files.

## Verification Results

- **`go test ./internal/bmc/qpotential`**: PASS
- **`go test ./internal/bmc/phaseaudit`**: PASS
- **`go test ./internal/bmc/bmc0bspec`**: PASS
- **`go test ./internal/bmc/wdw`**: PASS
- **`go test ./internal/bmc/report`**: PASS
- **`go test ./internal/bmc/convergence`**: PASS
- **`go test ./...`**: PASS
- **`lake build`**: PASS

## Summary and Next Steps

- **Remaining limitations**: Bounded domain-boundary audit only, no physics promotion or solver progress is claimed.
- **Whether POST-0006 is ready for review**: Yes, fully implemented and verified.
- **Next recommended ticket**: Conclude final review of the audit package.


# Task: BMC-POST-0006 Implementation

- [x] Create `internal/bmc/qpotential/domain.go` implementing near-node checks, constants, and structured audit outputs.
- [x] Enforce near-node blocking (amplitude floor check) making it explicit and non-authoritative.
- [x] Create `internal/bmc/qpotential/domain_test.go` with table-driven tests.
- [x] Implement explicit tests for near-node zero clamping rejection, invalid derivative steps, nonfinite derivative blocking, and phrase safety of errors.
- [x] Run verification tests and Lake build.
- [x] Scan for forbidden words.
- [x] Create `docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md` documentation.



