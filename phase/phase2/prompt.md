# BMC-POST-0001 Implementation Plan: Constraint Violation Detection

You are working on the Go-based BMC/PTW codebase under strict EBP 2.1 discipline.

Current status:

```text
Phase 1 Postmortem Inventory is accepted.
docs/postmortem/bmc_sprint_1_11_inventory.md exists.
Sprint 1–11 audit stack is frozen.
Full BMC remains blocked.
```

Do **not** start Sprint 12.

Do **not** add a new report schema.

Do **not** add a new audit layer.

Do **not** claim Friedmann recovery, classical-limit recovery, BMC validation, null-model failure, BMC superiority, scientific novelty, or full BMC promotion.

## Goal

Implement the first remediation ticket:

```text
BMC-POST-0001: Constraint Violation Detection
```

Purpose:

```text
Make the BMC toy pipeline fail correctly on wrong physics inputs before adding new physics ambitions.
```

The current risk is that the plane-wave Wheeler-DeWitt check can pass by analytic tautology. BMC-POST-0001 should prove the system can reject deliberately invalid wavefunctions or violated constraint parameters.

## Scope

This ticket should add **negative/failure-detection tests and minimal supporting code only**.

Allowed:

```text
wrong-constraint plane-wave tests
k² != ω² failure tests
wrong wavefunction fixture tests
numerical residual evaluator if needed for failure detection
report/gate propagation checks for failed constraint status
clear failure status vocabulary
small documentation note in docs/postmortem or docs/remediation
```

Forbidden:

```text
new physics model
new BMC sprint schema
new audit layer
Friedmann residual implementation
classical-limit recovery logic
model ranking
null-model pass/fail claims
full BMC promotion
```

## Required implementation strategy

### 1. Keep analytic residual as oracle/control, not primary authority

Do **not** delete the analytic plane-wave residual.

Instead, add or expose an independent check that can detect violations.

Preferred approach:

```text
Analytic residual remains a known-good oracle.
A numerical WdW residual evaluator becomes the failure-detection path.
```

The numerical evaluator can be minimal and toy-level.

Suggested package target:

```text
internal/bmc/wdw
```

or a small sibling file in the existing WdW package.

Suggested function:

```go
func NumericalResidualAt(
    psi func(alpha, phi float64) complex128,
    alpha float64,
    phi float64,
    h float64,
) (complex128, error)
```

It should approximate:

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ
```

using central finite differences.

It must reject:

```text
h <= 0
nonfinite alpha
nonfinite phi
nonfinite h
nonfinite residual components
```

### 2. Add wrong-constraint plane-wave test

Add test:

```text
TestWDWNumericalResidualRejectsWrongPlaneWaveConstraint
```

Use a plane wave:

```text
Ψ = exp(i(kα + ωφ))
```

with:

```text
k² != ω²
```

Expected:

```text
numerical residual magnitude > tolerance
constraint status = fail or violation_detected
```

The test must not pass merely because analytic code says so.

### 3. Add valid-control test

Add test:

```text
TestWDWNumericalResidualAcceptsConstraintShellPlaneWave
```

Use:

```text
k² = ω²
```

Expected:

```text
numerical residual magnitude <= tolerance
```

This test should show the numerical evaluator can distinguish valid from invalid plane waves.

### 4. Add wrong wavefunction fixture

Add at least one non-plane-wave or deliberately invalid fixture.

Example:

```text
Ψ(α, φ) = exp(i(kα + ωφ)) * (1 + εα²)
```

or:

```text
Ψ(α, φ) = exp(i(kα)) with missing φ dependence when the target expects ω² = k²
```

Add test:

```text
TestWDWNumericalResidualRejectsInvalidWavefunctionFixture
```

Expected:

```text
residual magnitude > tolerance
violation detected
```

### 5. Add report/gate propagation test if existing report path supports it

If the main BMC report path has a WdW or constraint check status, add a test that verifies a failed constraint cannot be promoted.

Possible test name:

```text
TestReportRejectsConstraintViolationStatus
```

or:

```text
TestPromotionGateBlocksConstraintViolation
```

Expected:

```text
constraint violation -> report validation fails or promotion gate is blocked
```

Do not force a large refactor if the report path cannot accept custom wavefunctions yet. If custom injection is not currently possible, document that as a follow-up ticket.

### 6. Add status vocabulary only if needed

If existing status values are insufficient, add minimal statuses such as:

```text
constraint_violation_detected
wdw_residual_failed
```

Do not add a new schema.

Do not add a new large report type.

### 7. Documentation

Add a short remediation note, preferably:

```text
docs/postmortem/bmc_post_0001_constraint_violation_detection.md
```

or append a short section to:

```text
docs/postmortem/bmc_sprint_1_11_inventory.md
```

The note should state:

```text
BMC-POST-0001 adds negative tests proving that wrong WdW inputs can be detected.
This does not implement Friedmann recovery.
This does not validate BMC.
Full BMC remains blocked.
```

## Required tests

Add or confirm:

```text
TestWDWNumericalResidualAcceptsConstraintShellPlaneWave
TestWDWNumericalResidualRejectsWrongPlaneWaveConstraint
TestWDWNumericalResidualRejectsInvalidWavefunctionFixture
TestNumericalResidualRejectsInvalidStep
TestNumericalResidualRejectsNonfiniteInputs
```

If report integration is feasible:

```text
TestReportRejectsConstraintViolationStatus
TestPromotionGateBlocksConstraintViolation
```

## Acceptance criteria

This ticket is accepted only if:

```text
1. A deliberately wrong WdW input fails.
2. A valid plane-wave control passes.
3. Invalid numerical inputs are rejected.
4. Failure status cannot be misreported as success.
5. go test ./... passes.
6. lake build passes.
7. No recovery, superiority, null-failure, novelty, or full-BMC claims are introduced.
```

## Verification commands

Run:

```bash
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -count=1
GOCACHE=/tmp/go-build-cache go test ./...
cd BMC && /home/chaschel/.elan/bin/lake build
```

If report-path tests are added, include their package-specific test command too.

## Required walkthrough

Return:

```text
files added
files modified
numerical residual function added or not
negative tests added
valid control tests added
report/gate propagation tests added or deferred
documentation note path
go test ./internal/bmc/wdw result
go test ./... result
lake build result
remaining limitations
next recommended ticket
```

## EBP debt status expected after implementation

```text
needConstraintViolationTests: partial or retired_for_plane_wave_scope
needNumericalErrorAudit: unpaid
needNontrivialPhysicsCase: unpaid
needToyCheck: partial
needFaithfulnessReview: contested
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

## Strict EBP reminder

A wrong-input rejection test is not physics validation.

A numerical residual evaluator is not a full WdW solver.

A passed plane-wave control is not Friedmann recovery.

Full BMC remains blocked.
