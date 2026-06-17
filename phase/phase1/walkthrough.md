# Phase 1 Postmortem Inventory Walkthrough

This walkthrough outlines the work completed for the Phase 1 Postmortem Inventory under EBP 2.1 discipline.

## 1. Accomplishments & Summary Statistics
We completed a comprehensive audit and classification of the Bohmian Minisuperspace Cosmology (BMC-0.1) repository. The resulting artifact has been created and verified.

* **Inventory Document Path:** [bmc_sprint_1_11_inventory.md](file:///home/chaschel/Documents/go/bmc/docs/postmortem/bmc_sprint_1_11_inventory.md)
* **Summary of Schema Count:** 10 schemas found (`bmc-report-v0.1`, `bmc0a-superposition-robustness-v0.1`, `bmc0a-clock-fragility-v0.1`, `bmc0a-clock-readiness-v0.1`, `bmc0a-friedmann-spec-v0.1`, `bmc0a-nullmodel-spec-v0.1`, `bmc0a-prior-art-boundary-v0.1`, `bmc0a-nullrun-v0.1`, `bmc0a-local-residual-v0.1`, `bmc0a-residual-audit-v0.1`)
* **Summary of CLI Count:** 12 subcommands mapped (`run`, `validate`, `summarize`, `audit`, `diagnose-clock`, `segment-clock`, `spec-friedmann`, `spec-nullmodels`, `prior-art-boundary`, `run-nullmodels`, `run-residuals`, `audit-residuals`)
* **Summary of Package Count:** 18 packages/subfolders classified (17 under `internal/bmc/`, plus the Lean policy support package under `BMC/BMC/`)
* **Summary of Test Categories:** 10 testing categories identified, spanning positive/control, negative/failure-detection, integration, CLI routing, validation, determinism, numerical-convergence, sensitivity, anti-overclaim, and Lean policy tests.

---

## 2. Key Findings & Registers

### Top Validation-Theater Risks
1. **Near-Node Q-Potential Clamp:** `qpotential.Q` clamps the evaluation of $Q$ to $0.0$ near wavefunction nodes (amplitude $< 10^{-12}$) to prevent division by zero, creating a fake indication of classical Einstein limit validity instead of failing.
2. **Analytic Tautology Authority:** Wheeler-DeWitt residual check relies on evaluating a simple analytical subtraction ($k^2 - \omega^2$) instead of validating the wave equation numerically across the generated trajectory points.
3. **Decorative Stability Audits:** Stability sweeps in Sprint 11 use interval proxy rescalings rather than re-integrating trajectories under real coordinate/parameter perturbations.

### Top Missing Failure-Detection Tests
1. **Wrong Wavefunction Rejection:** No check in the codebase ensures that a wavefunction which does not solve the Wheeler-DeWitt equation is rejected.
2. **Coordinate Violated Control Check:** No test verifies that parameter pairs where $k^2 \neq \omega^2$ fail WdW checks.
3. **Node Contact Short-Circuit Failure:** The node-probe short-circuit must explicitly block promotion gates instead of returning finite/contested mock validity.

---

## 3. Recommended Next Ticket
* **BMC-POST-0001: Constraint Violation Detection (Priority: P1)**
  - **Scope:** Introduce negative test fixtures validating that incorrect wavefunctions or violated coordinate relationships ($k^2 \neq \omega^2$) explicitly trigger validation failures in the WdW check.

---

## 4. Verification Results
* **Go Tests:** `go test ./...` passed successfully (cached/run checked).
* **Lean Build:** `lake build` completed successfully with 14 jobs built and zero `sorry`/`admit` obligations.

---

## 5. EBP 2.1 Reminders
* This postmortem inventory is an administrative and structural registry, not remediation.
* Creating a schema inventory does not constitute physics progress.
* Highlighting failure-detection gaps is not the implementation of failure detection.
* Bounding the residual comparison audit is not a classical cosmology recovery.
* **Full BMC remains blocked.**
