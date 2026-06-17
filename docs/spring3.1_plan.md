# BMC Sprint 3: Numerical Robustness and Convergence Audit (Revised)

Numerical robustness and convergence audit of the accepted `bmc0a-superposition-safe` and `bmc0a-superposition-node-probe` models. We verify that these results are stable under numerical variations (step size, node threshold, phase-gradient bound) and small parameter perturbations, and verify that our tools report behavior honestly.

## User Review Required

> [!IMPORTANT]
> **Promotion and Physics Boundaries**:
> The maximum allowed promotion in Sprint 3 is `promoted_robustness_audit_artifact_after_repairs`.
> This is a purely numerical audit of the minisuperspace toy model. It does NOT imply or prove:
> - Bohmian mechanics correctness
> - Quantum gravity
> - Problem of time resolution
> - Spacetime emergence
> - Friedmann recovery
>
> The Friedmann residual check remains deferred, and the full BMC toy gate remains blocked.

---

## Proposed Changes

### Audit Logic & Report Generation (`internal/bmc/audit/`)

We will introduce a new Go package `internal/bmc/audit` to perform the numerical sweeps and generate the robustness JSON report.

#### [NEW] [convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/convergence.go)
- Implements step-size convergence sweep.
- **Fixed Integration Interval**: Keeps total integration interval fixed at $T = 10.0$.
  - Runs for `LambdaStep = [0.1, 0.05, 0.025, 0.0125]` with corresponding `Steps = [100, 200, 400, 800]`.
- Computes trajectory endpoint drift (difference in final coordinates relative to the finest step size `0.0125`).
- Evaluates `max |Q| away from nodes`, `min amplitude R`, `max phase-gradient magnitude`, clock monotonicity, and individual check statuses.

#### [NEW] [threshold_sensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/threshold_sensitivity.go)
- Runs both `safe` and `node-probe` profiles with `NodeThresh = [1e-4, 1e-5, 1e-6]`.
- Reports whether the safe profile stays node-contact-free and whether the node-probe remains blocked as node contact.
- Audits the statuses of `q_finite_away_from_nodes` and `phase_gradient_finite` under these thresholds.

#### [NEW] [phase_bound_sensitivity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/phase_bound_sensitivity.go)
- Runs the safe profile with `MaxPhaseGrad = [25, 50, 100, 200]`.
- Reports the max observed phase gradient, whether the configured bound is binding, and whether it alters the check status.

#### [NEW] [perturbation.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/perturbation.go)
- Audits behavior under small perturbations while preserving each component's WdW constraint:
  - $c_2 \in \{0.45, 0.50, 0.55\}$
  - $k_2 \in \{1.9, 2.0, 2.1\}$ (with $\omega_2 = -k_2$)
- Audits node contact, Q finiteness, clock monotonicity, and gate status changes.
- **Node-probe offset sweep**: Runs offsets from the exact node $(\alpha_0, \phi_0) = (0, 0)$, specifically $(1\times 10^{-8}, 0)$, $(1\times 10^{-6}, 0)$, and $(1\times 10^{-4}, 0)$. Classifies offsets relative to `NodeThresh` ($10^{-5}$):
  - $(0, 0)$, $(1\times 10^{-8}, 0)$, $(1\times 10^{-6}, 0)$ must short-circuit.
  - $(1\times 10^{-4}, 0)$ is above threshold, integrates, and reports if phase/Q risks appear.

#### [NEW] [robustness_report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/robustness_report.go)
- Defines the `RobustnessReport` JSON struct:
  - `schema_version`: `"bmc0a-superposition-robustness-v0.1"`
  - `toy_analysis_only`: `true`
  - `final_truth_claim`: `false`
  - `source_artifacts`: `["bmc0a_superposition_safe", "bmc0a_superposition_node_probe"]`
  - `audit_kind`: `"numerical_robustness_convergence"`
  - **Computed Robustness Gate**: `technical_gate` will have computed `status` ("pass", "fail", or "contested") and `reason`.
    - Passes if all planned sweeps run, reports validate successfully, no NaN/Inf values are hidden, and gate flips are honestly recorded.
  - `promotion_gate`: `{ "name": "full_bmc_toy_gate", "status": "blocked" }`
  - Detailed sweep structures for step sizes, thresholds, phase-gradient bounds, perturbations, and node probe offsets.
- Implements `GenerateAuditReport()` and `WriteJSON()` for serialization.
- Implements **ValidateRobustnessReport()** and **SummarizeRobustnessReport()** to strictly validate the robustness report schema and print a structured summary without modifying or weakening the original plane/superposition report validator.

---

### Command CLI Integration (`cmd/ptw-bmc/`)

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Integrate a new `audit` subcommand:
  `ptw-bmc audit --profile bmc0a-superposition-robustness --out out/bmc0a_superposition_robustness.json`
- Update `validate` and `summarize` subcommands to automatically route to `ValidateRobustnessReport` / `SummarizeRobustnessReport` based on the JSON schema version, avoiding loose/lax schema acceptance.

---

### Lean Policy/Safety Gates (`BMC/`)

#### [NEW] [Robustness.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Robustness.lean)
- Formalize the structure of `BMCRobustnessReport`.
- Define `reportPassesBMC0ARobustnessAuditGate` check logic.
- Implement safety theorems verifying that a passing robustness report does not imply full quantum gravity/Bohmian mechanics recovery and keeps the final truth claim blocked:
  - `robustness_audit_does_not_imply_full_bmc`
  - `robustness_audit_requires_toy_only`
  - `robustness_audit_blocks_final_truth`
  - `robustness_audit_requires_friedmann_deferred`
- Compile this package as part of `lake build`.

---

## Verification Plan

### Automated Tests
- **`TestRobustnessReportValidation`**: Validates schema verification, ensuring that the validator rejects unknown fields or final truth claims in the robustness report.
- **`TestStepSizeConvergence`**:
  - Verifies all step-size runs complete deterministically.
  - Verifies endpoint drift is finite and reported.
  - Verifies gate status changes, if any, are explicitly and honestly recorded.
- **`TestThresholdSensitivity`**: Verifies safe profiles pass and node-probe profiles fail under all thresholds.
- **`TestPhaseGradientSensitivity`**: Verifies bound behaves consistently.
- **`TestParameterPerturbations`**: Asserts consistency under small parameter offsets.

### Manual Verification
Execute:
```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc audit --profile bmc0a-superposition-robustness --out out/bmc0a_superposition_robustness.json
./ptw-bmc validate --report out/bmc0a_superposition_robustness.json
./ptw-bmc summarize --report out/bmc0a_superposition_robustness.json
cd BMC && /home/chaschel/.elan/bin/lake build
```
