# BMC-POST-0007 Implementation Plan: Branch-Cut and Stencil Boundary Regression Hardening

targeted regression hardening to protect phase-gradient calculation and quantum-potential domain boundaries from numerical regressions.

## User Review Required

> [!IMPORTANT]
> This ticket is strictly regression hardening for the numerical audit stack. It does not validate physical models, unblock solver implementation, or promote the BMC stage.

## Open Questions

None. The target test cases are well defined.

## Proposed Changes

---

### Phase Gradient Audit Component

#### [MODIFY] [hsensitivity_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/phaseaudit/hsensitivity_test.go)
- Add `TestHSensitivityBranchCutSafeGradientDoesNotArgWrap` regression test.
- Use a mock wavefunction `Psi(alpha, phi) = exp(i * (Pi - 0.5 * alpha))`.
- Evaluate near `alpha = 0.0` (where phase wrapping at $\pi \leftrightarrow -\pi$ occurs under naïve finite difference of `arg(Ψ)`).
- Verify the audit's output gradient remains stable and close to `-0.5`, avoiding spurious jumps.

---

### Quantum Potential Component

#### [MODIFY] [domain_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/domain_test.go)
- Add `TestQPotentialBlocksStencilPointBelowAmplitudeFloor` regression test.
- Construct a mock wavefunction where the center point is safe, but one stencil point ($\alpha + h$) is below `NearNodeAmplitudeFloor`.
- Verify the audit blocks evaluation with `Authoritative = false` and status `q_potential_blocked_by_domain_boundary`.
- Add `TestQPotentialRunAuditRejectsInvalidDerivativeStep` verifying that the public audit entry point rejects $h \le 0$, `NaN`, and `Inf`.

---

### Documentation Component

#### [NEW] [bmc_post_0007_branchcut_stencil_regression_hardening.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_post_0007_branchcut_stencil_regression_hardening.md)
Document the postmortem design details of the regression tests, explaining how they protect the audit stack from phase-wrap and stencil contamination errors.

---

## Verification Plan

### Automated Tests
Run package regression tests:
```bash
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/phaseaudit -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/qpotential -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/bmc0bspec -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/report -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/convergence -v -count=1
GOCACHE=/tmp/go-build-cache go test ./... -count=1
cd BMC && /home/chaschel/.elan/bin/lake build
```
Verify forbidden phrases are absent.
