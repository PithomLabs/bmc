# Walkthrough - BMC-POST-0005: Phase Gradient h-Sensitivity Audit

This walkthrough documents the verified implementation of the finite-difference h-sensitivity audit for phase gradients.

## Changes Made

### Files Added
- [hsensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity.go)
- [hsensitivity_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity_test.go)
- [bmc_post_0005_phase_gradient_h_sensitivity.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md)

### Files Modified
- None.

### Detailed Status
- **Audit package created**: Created `internal/bmc/phaseaudit`.
- **h ladder used**: $h = 10^{-2}, 5 \times 10^{-3}, 2.5 \times 10^{-3}, 1.25 \times 10^{-3}$ (strictly validated for descending order, non-empty, and finite values).
- **Sample fixtures used**: `plane_wave`, `superposition`, and `near_node`.
- **Phase-gradient computation source**: Utilizes `Im((1/Ψ) * ∂Ψ/∂x)` formulation internally to protect against branch-cut phase wrapping artifacts.
- **Threshold constants**: Defined named constants:
  - `DefaultHCoarse = 1e-2`
  - `DefaultHMin = 1.25e-3`
  - `AbsoluteDriftTolerance = 1e-6`
  - `RelativeDriftTolerance = 1e-4`
  - `NearNodeAmplitudeFloor = 1e-8`
  - `SignFlipTolerance = 1e-7`
- **Stability metrics**: Evaluates successive L2 and L-infinity drifts, max component drift, and relative drift where the denominator is safe ($>10^{-5}$). Detects sign flips across all $h$ evaluations.
- **Node/nonfinite blocking behavior**: Any amplitude below `NearNodeAmplitudeFloor` ($10^{-8}$) or nonfinite complex evaluations immediately blocks the point, making it non-authoritative.
- **Forbidden inference audit**: Rejects forbidden terms in status, schema, and EBP fields (`ready`, `validated`, `proved`, `recovered`, `successful`, etc.) case-insensitively.
- **Tests added**: Enforces all requested test cases including plane wave stability, perturbed function h-drift detection, near-node blocking, invalid h-ladders (ordering and validation), NaN/Inf coordinate rejection, and phrase safety of error messages.
- **Documentation added**: Detailed postmortem guide outlining context, conventions, and EBP boundaries.
- **Diff hygiene status**: Clean. Git diff shows no modifications to existing tracked files.

## Verification Results

- **`go test ./internal/bmc/phaseaudit`**: PASS
- **`go test ./internal/bmc/bmc0bspec`**: PASS
- **`go test ./internal/bmc/wdw`**: PASS
- **`go test ./internal/bmc/report`**: PASS
- **`go test ./internal/bmc/convergence`**: PASS
- **`go test ./...`**: PASS
- **`lake build`**: PASS

## Summary and Next Steps

- **Remaining limitations**: Numerical audit only, no physics promotion or solver work.
- **Whether POST-0005 is ready for review**: Yes, fully implemented and verified.
- **Next recommended ticket**: Complete the final peer review of BMC-POST-0005.

# Task: BMC-POST-0005 Implementation

- [x] Confirm POST-0004.2 diff hygiene is clean (no untracked files).
- [x] Create `internal/bmc/phaseaudit/hsensitivity.go` with core logic, constants, and structured audit outputs.
- [x] Ensure phase-gradient convention uses `∂S/∂alpha = k` and `∂S/∂phi = omega` for plane waves, distinct from trajectory velocity signs.
- [x] Implement branch-cut guards using `Im((∂Ψ/Ψ))` formulation to avoid phase wrapping, matching `wave.PhaseGradient` logic.
- [x] Enforce near-node blocking (amplitude floor check) making it explicit and non-authoritative.
- [x] Create `internal/bmc/phaseaudit/hsensitivity_test.go` with table-driven tests.
- [x] Implement explicit tests for `h`-ladder ordering (invalid ladder detection, nonfinite `h` rejection).
- [x] Run verification tests and Lake build.
- [x] Scan for forbidden words (including `ready`).
- [x] Create `docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md` documentation.



