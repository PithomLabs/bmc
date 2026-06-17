# BMC Sprint 6: Friedmann-Residual Specification and Gate Design

This sprint plans the specification, variable mapping, and safety gating required for a future Friedmann-style residual calculation, under strict EBP 2.1 guidelines. No actual physics residual will be computed or evaluated numerically in this sprint.

## User Review Required

> [!WARNING]
> This sprint **strictly forbids** implementing Friedmann residual recovery or attempting to solve the problem of time. Any attempt to claim full quantum gravity, validity of φ-clock for full cosmology, or ready for Friedmann recovery will be rejected. The focus is strictly on **specification** and **gate design**.

## Proposed Changes

### Go Code Package: `internal/bmc/friedmannspec`

---

#### [NEW] [contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/contracts.go)
Defines EBP spec configurations and lists convention, target, and choice debts.

#### [NEW] [mapping.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/mapping.go)
Defines the `FriedmannCandidateMap` structure to log candidate FRW targets and convention/sign/normalization debt status. Allowed statuses are `candidate_only`, `contested`, and `blocked`.

#### [NEW] [branch_requirements.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/branch_requirements.go)
Defines `FriedmannBranchRequirement` which takes Sprint 5 segments and validates index ranges, minimum samples, clock range, and derivative preparation. Allowed readiness values are `blocked`, `candidate_only`, and `contested`.

#### [NEW] [derivatives.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/derivatives.go)
Defines `DerivativeReadinessCheck` to specify numerical derivative methods, turning point exclusions, near-node exclusions, and finite difference step sensitivity.

#### [NEW] [residual_spec.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/residual_spec.go)
Defines `ResidualFormulaCandidate` registry to specify candidate classical target relations, required variables, and convention debts without executing numerical evaluation.

#### [NEW] [gates.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/gates.go)
Defines the `FriedmannSpecGate` structure and evaluates the 10 required gate criteria (such as `no_residual_computation_gate`, `full_bmc_blocked_gate`, `toy_analysis_only_gate`, etc.).

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/report.go)
Assembles the deterministic `ClockReadinessReport` matching the `bmc0a-friedmann-spec-v0.1` schema.

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/validate.go)
Enforces strict schema validation including rejecting `residual_computed = true`, `friedmann_recovery_claim = true`, unblocked promotion status, and checking for unknown JSON fields.

#### [NEW] [friedmannspec_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/friedmannspec_test.go)
Includes the 15 required unit tests validating all spec boundaries.

---

### Command Line Interface

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)
- Adds subcommand `spec-friedmann`.
- Routes validation and summarization of schema version `bmc0a-friedmann-spec-v0.1`.

---

### Lean Formal Verification

#### [NEW] [FriedmannSpec.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/FriedmannSpec.lean)
Defines `BMCFriedmannSpecReport` structure and proves the 12 policy-only safety theorems.

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)
Imports `BMC.FriedmannSpec`.

---

## Verification Plan

### Automated Tests
- Run `go test ./internal/bmc/friedmannspec` to verify all spec and validator rules.
- Run `lake build` to verify Lean safety theorems compile and compile cleanly.

### Manual Verification
- Generate specification: `ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json`
- Validate specification: `ptw-bmc validate --report out/bmc0a_friedmann_spec.json`
- Summarize specification: `ptw-bmc summarize --report out/bmc0a_friedmann_spec.json`
