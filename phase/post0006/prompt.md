# BMC-POST-0006 Implementation Prompt: Quantum Potential Near-Node Domain-Boundary Audit

You are working on the Go-based BMC/PTW repository under strict EBP 2.1 discipline.

## Ticket

```text
BMC-POST-0006: Quantum Potential Near-Node Domain-Boundary Audit
```

## Purpose

Repair the remaining near-node quantum-potential audit debt.

The current risk is:

```text
Near nodes, the quantum potential expression can become singular, unstable, or non-authoritative.
If the implementation silently clamps, defaults, or treats near-node quantum-potential values as valid, the audit can become validation theater.
```

This ticket must make near-node quantum-potential behavior explicit, blocked, and non-authoritative where appropriate.

## Strict EBP boundary

This is a **numerical/domain-boundary audit**, not physics promotion.

Do not claim:

```text
BMC validation
Friedmann recovery
classical-limit recovery
quantum gravity progress
BMC-0B solver progress
null-model failure
BMC superiority
scientific novelty
full BMC promotion
```

Do not implement:

```text
BMC-0B solver
massive scalar numerical solution
new trajectory solver
new CLI route
new report schema
new JSON output artifact
Friedmann residual computation
```

Allowed scope:

```text
BMC-0A quantum-potential near-node audit only
Go package/tests or targeted repair in existing quantum-potential package
small postmortem documentation note
no CLI/schema/output artifact by default
```

## Core review question

```text
Does the BMC code correctly treat quantum-potential values near nodes as domain-boundary/non-authoritative cases instead of silently producing authoritative Q values?
```

## Target files

Prefer a narrow package or targeted modification around existing Q-potential code.

Expected possible files:

```text
internal/bmc/qpotential/domain.go
internal/bmc/qpotential/domain_test.go
```

or, if the existing package name differs, use the existing package that computes quantum potential.

Add documentation:

```text
docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md
```

Avoid touching:

```text
cmd/ptw-bmc/main.go
internal/bmc/report/*
internal/bmc/model/*
out/*.json
BMC/Lean files
```

unless absolutely necessary and explicitly justified.

## Required behavior

Quantum potential must not silently return an authoritative value when the wave amplitude is too small or when derivatives are nonfinite.

Define named constants such as:

```go
const NearNodeAmplitudeFloor = 1e-8
const QPotentialDerivativeStep = 1e-4
const QPotentialMagnitudeWarning = 1e8
```

Use existing constants if already present and appropriate.

## Required statuses

Introduce or enforce statuses equivalent to:

```text
q_potential_authoritative
q_potential_blocked_by_node_contact
q_potential_blocked_by_near_node_amplitude
q_potential_blocked_by_nonfinite_wave
q_potential_blocked_by_nonfinite_derivative
q_potential_blocked_by_domain_boundary
q_potential_audit_only_no_promotion
```

Do not use promotion-style terms.

Avoid emitted status words such as:

```text
validated
proved
recovered
ready
successful
physics_success
bmc_validated
friedmann_recovered
quantum_gravity_progress
full bmc unblocked
bmc beats nulls
```

## Required data model

Use a small result struct or extend the existing local computation result with explicit authority metadata.

Example:

```go
type QPotentialAuditResult struct {
    Alpha float64 `json:"alpha"`
    Phi float64 `json:"phi"`
    Amplitude float64 `json:"amplitude"`
    QPotential float64 `json:"q_potential,omitempty"`
    Status string `json:"status"`
    Authoritative bool `json:"authoritative"`
    Reason string `json:"reason,omitempty"`
}
```

Required boundary fields if a top-level audit object is introduced:

```text
toy_analysis_only = true
physics_claim = "none"
bmc0b_impact = "none"
friedmann_recovery_impact = "none"
promotion_recommendation = "do_not_promote"
```

## Required computations

Audit Q-potential evaluation at deterministic sample points:

```text
plane_wave_control_points
superposition_safe_points
near_node_probe_points
```

Expected behavior:

```text
plane_wave_control_points:
  Q should be finite and authoritative if amplitude is safely nonzero.

superposition_safe_points:
  Q may be finite, but remains audit-only and not physics validation.

near_node_probe_points:
  Q must be blocked or non-authoritative if amplitude falls below NearNodeAmplitudeFloor, if derivatives are nonfinite, or if stencil points cross a node-sensitive domain boundary.
```

Do not silently clamp Q to zero near nodes.

If the existing code currently returns zero near nodes, this ticket must either:

```text
replace that behavior with blocked/non-authoritative status
```

or:

```text
wrap the existing computation so reports/tests cannot treat clamped near-node Q as authoritative
```

## Required validation

Reject invalid inputs:

```text
NaN/Inf alpha
NaN/Inf phi
nonpositive derivative step
NaN/Inf derivative step
missing fixture/profile ID if audit-level object is used
unknown fixture type
```

Block non-authoritative cases:

```text
amplitude below floor
nonfinite wavefunction
nonfinite derivative
nonfinite Q
stencil touches near-node region
division by near-zero amplitude
```

## Required tests

Add table-driven Go tests covering at least:

```text
TestQPotentialPlaneWaveControlAuthoritative
TestQPotentialSuperpositionSafeAuditOnly
TestQPotentialNearNodeBlocked
TestQPotentialDoesNotClampNearNodeToAuthoritativeZero
TestQPotentialRejectsNonfinitePoint
TestQPotentialRejectsInvalidDerivativeStep
TestQPotentialBlocksNonfiniteDerivative
TestQPotentialDeterministic
TestQPotentialNoPromotionFields
TestQPotentialForbiddenInferenceAudit
```

The most important test is:

```text
TestQPotentialDoesNotClampNearNodeToAuthoritativeZero
```

It must fail if near-node Q silently returns `0` with `Authoritative=true`.

## Forbidden inference audit

Add a forbidden-term scan for accepted status/report fields and validation errors.

Forbidden terms:

```text
validated
proved
recovered
ready
successful
physics_success
bmc_validated
friedmann_recovered
quantum_gravity_progress
full bmc unblocked
bmc beats nulls
```

Allowed wording:

```text
audit_only
toy_analysis_only
non_authoritative
blocked
domain_boundary
do_not_promote
```

Validation errors must be phrase-safe and must not echo forbidden phrases from contaminated input.

## Documentation

Create:

```text
docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md
```

The note must state:

```text
POST-0006 is a quantum-potential near-node domain-boundary audit.
It does not implement BMC-0B.
It does not solve the Wheeler-DeWitt equation.
It does not validate BMC.
It does not claim Friedmann recovery.
It does not prove classical-limit recovery.
It does not run null models.
It does not unblock full BMC.
```

It should explain:

```text
why Q-potential can become unstable near nodes
what amplitude floor was used
what stencil/domain checks were added
what counts as authoritative
what counts as non-authoritative
what remains unpaid
```

## Verification commands

Run:

```bash
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/qpotential -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/phaseaudit -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/bmc0bspec -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/report -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/convergence -v -count=1
GOCACHE=/tmp/go-build-cache go test ./... -count=1
cd BMC && /home/chaschel/.elan/bin/lake build
```

If the package is not named `qpotential`, replace that command with the actual package path.

Run forbidden scan:

```bash
grep -R "validated\|proved\|recovered\|ready\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" internal/bmc/qpotential docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md || true
```

Forbidden phrases may appear only in rejected test fixtures, not in accepted statuses, docs claims, or output fields.

## Required walkthrough

Return a walkthrough with:

```text
files added
files modified
existing Q-potential behavior found
whether near-node Q previously clamped or defaulted
domain-boundary authority model added
amplitude floor used
derivative step used
stencil/domain checks added
node/nonfinite blocking behavior
tests added
forbidden inference audit
documentation added
go test ./internal/bmc/qpotential result or actual package path
go test ./internal/bmc/phaseaudit result
go test ./internal/bmc/bmc0bspec result
go test ./internal/bmc/wdw result
go test ./internal/bmc/report result
go test ./internal/bmc/convergence result
go test ./... result
lake build result
remaining limitations
whether POST-0006 is ready for review
next recommended ticket
```

## Expected EBP status after implementation

```text
BMCPost0006Status: implemented_pending_review
QPotentialNearNodeDomainBoundary: audit_only
needNumericalErrorAudit: improved_for_qpotential_near_node_scope
needNontrivialPhysicsCase: unpaid
needFaithfulnessReview: unchanged
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
SolverStatus: not_implemented
BMC0BStatus: specified_only
promotion_recommendation: do_not_promote
```

## Strict EBP reminder

A Q-potential domain-boundary audit is not a solver.

Blocking near-node Q is not physics success.

A finite, stable Q value at a control point is not BMC validation.

This does not establish Friedmann recovery.

This does not unblock full BMC.
