# BMC-POST-0005 Implementation Plan: Phase Gradient h-Sensitivity Audit

Audit the finite-difference step size `h` sensitivity of phase gradient calculations to detect numerical instability and node-probe behaviors.

## User Review Required

> [!IMPORTANT]
> The audit is numerical-only. No physical promotion or solver progress is claimed. EBP status fields will strictly reject any promotion.

## Open Questions

None. The requirements, target files, and test specifications are clear.

## Proposed Changes

---

### Phase Audit Component

#### [NEW] [hsensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity.go)
Create the core phase-gradient sensitivity audit structure, constants, and execution logic.
- Enforce input validations (non-empty `hValues`, elements > 0, no NaN/Inf, descending order, non-empty sample points, valid coordinates, missing fixture ID, etc.).
- Compute `dS/dalpha` and `dS/dphi` across step sizes using finite differences on `wave.WaveFunction`.
- Compare successive step size drift (`gradient_l2_drift_between_successive_h`, `gradient_linf_drift_between_successive_h`, relative drift, sign flips, nonfinite values, and near-node detection).
- Assign status classifications: `stable_for_control_scope`, `sensitive_to_h`, `blocked_by_node_contact`, `blocked_by_nonfinite_gradient`, `invalid_input`, `audit_only_no_promotion`.
- Expose a clean, deterministic API: `RunAudit`.

#### [NEW] [hsensitivity_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity_test.go)
Implement table-driven test suites validating:
- `TestHSensitivityPlaneWaveControlStable`
- `TestHSensitivityDetectsHDriftOnPerturbedPhaseFunction`
- `TestHSensitivityBlocksNearNodeProbe`
- `TestHSensitivityRejectsInvalidHLadder`
- `TestHSensitivityRejectsNonfiniteH`
- `TestHSensitivityRejectsNonfinitePoint`
- `TestHSensitivityDeterministic`
- `TestHSensitivityNoPromotionFields`
- `TestHSensitivityDoesNotRequireCLIOrSchema`
- Scan error and status messages for forbidden terms (`validated`, `proved`, `recovered`, `successful`, `physics_success`, `bmc_validated`, `friedmann_recovered`, `quantum_gravity_progress`, `full bmc unblocked`, `bmc beats nulls`, `ready`).

---

### Documentation Component

#### [NEW] [bmc_post_0005_phase_gradient_h_sensitivity.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md)
Document the audit purpose, finite difference sensitivity context, thresholds, results on test wavefunctions, and the unchanged/blocked status of physical BMC-0B promotion.

---

## Verification Plan

### Automated Tests
- Run tests in the new package:
  ```bash
  GOCACHE=/tmp/go-build-cache go test ./internal/bmc/phaseaudit -v -count=1
  ```
- Run all workspace tests to ensure no regressions:
  ```bash
  GOCACHE=/tmp/go-build-cache go test ./... -count=1
  ```
- Run Lean/Lake build:
  ```bash
  cd BMC && /home/chaschel/.elan/bin/lake build
  ```

### Manual Verification
- Check git status for diff hygiene:
  ```bash
  git status --short
  git diff --stat
  ```
- Run the forbidden term scanning command:
  ```bash
  grep -R "validated\|proved\|recovered\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" internal/bmc/phaseaudit docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md || true
  ```
