# BMC-POST-0007 Planning Prompt: Branch-Cut and Stencil Boundary Regression Hardening

You are planning the next bounded remediation ticket for the Go-based BMC/PTW repository under strict EBP 2.1 discipline.

## Ticket

```text
BMC-POST-0007: Branch-Cut and Stencil Boundary Regression Hardening
```

## Context

Accepted remediation stack so far:

```text
BMC-POST-0001: Constraint Violation Detection
BMC-POST-0002 / 0002.1: Numerical WdW Residual Authority Repair
BMC-POST-0003 / 0003.2: Euler/RK4 and dt Convergence Audit
BMC-POST-0004 / 0004.2: BMC-0B Massive Scalar WdW Specification Contract
BMC-POST-0005: Phase Gradient h-Sensitivity Audit
BMC-POST-0006: Quantum Potential Near-Node Domain-Boundary Audit
```

POST-0005 and POST-0006 were accepted in substance as bounded numerical-audit artifacts. They did not validate BMC, did not implement BMC-0B, did not recover Friedmann dynamics, and did not unblock full BMC.

## Purpose

Plan a small hardening ticket that adds targeted regression tests and, only if necessary, small validation helpers for two remaining numerical audit risks:

```text
1. Phase-gradient branch-cut regression risk.
2. Quantum-potential stencil boundary regression risk.
```

This ticket should make sure the audit stack fails safely when phase wrapping, near-node stencil contamination, or invalid derivative-step inputs appear.

## Strict EBP boundary

This is a **regression hardening plan**, not implementation yet.

Do not propose:

```text
BMC-0B solver
massive scalar numerical solution
trajectory solver
Friedmann residual computation
classical-limit recovery
new report schema
new CLI route
generated JSON output artifacts
physics validation
null-model victory
BMC superiority
scientific novelty
full BMC promotion
```

Allowed scope:

```text
small Go tests
small internal helper repairs only if tests expose a real gap
small postmortem note
no schema/CLI/output artifact by default
```

## Target areas to inspect

Likely relevant packages:

```text
internal/bmc/phaseaudit
internal/bmc/qpotential
internal/bmc/wave
```

Potential documentation:

```text
docs/postmortem/bmc_post_0007_branchcut_stencil_regression_hardening.md
```

## Planning focus 1: explicit branch-cut regression test

POST-0005 uses the branch-cut-safe identity:

```text
Im((1/Ψ) * ∂Ψ/∂x)
```

Plan a regression test that would fail if the implementation accidentally switched back to naïvely finite-differencing `arg(Ψ)` across the `π ↔ -π` branch cut.

The plan should specify:

```text
test fixture or mock wavefunction
sample point near phase wrapping
expected behavior for branch-cut-safe derivative
what failure would look like under naïve arg-difference
whether this belongs in phaseaudit tests or wave tests
```

Preferred test name:

```text
TestHSensitivityBranchCutSafeGradientDoesNotArgWrap
```

or a clearer equivalent.

## Planning focus 2: qpotential stencil boundary edge test

POST-0006 checks central and stencil amplitudes against `NearNodeAmplitudeFloor`.

Plan a targeted regression test where:

```text
central point amplitude is safe
one stencil point falls below the amplitude floor
audit must return Authoritative=false
status must indicate domain-boundary or near-node stencil contamination
Q must not be treated as authoritative evidence
```

Preferred test name:

```text
TestQPotentialBlocksStencilPointBelowAmplitudeFloor
```

or a clearer equivalent.

## Planning focus 3: top-level invalid derivative-step path

POST-0006 review suggested adding a top-level integration-boundary test for invalid derivative step.

Plan a test that verifies the public audit entry point rejects:

```text
h <= 0
NaN h
Inf h
```

before any quantum-potential computation is treated as meaningful.

Preferred test name:

```text
TestQPotentialRunAuditRejectsInvalidDerivativeStep
```

or equivalent.

## Planning focus 4: preserve no-promotion fields

The plan must preserve these EBP boundaries:

```text
toy_analysis_only = true
physics_claim = "none"
bmc0b_impact = "none"
friedmann_recovery_impact = "none"
promotion_recommendation = "do_not_promote"
```

No test should use words like:

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

except as rejected fixtures in forbidden-term tests.

## Expected output

Return a concise implementation plan with this structure:

```json
{
  "ticket": "BMC-POST-0007",
  "title": "Branch-Cut and Stencil Boundary Regression Hardening",
  "plan_status": "proposed",
  "scope": "regression_hardening_only",
  "summary": "",
  "files_to_inspect": [],
  "files_expected_to_add": [],
  "files_expected_to_modify": [],
  "tests_to_add": [],
  "implementation_steps": [],
  "validation_rules": [],
  "forbidden_claims": [],
  "ebp_debt_status": {
    "BMCPost0007Status": "planned",
    "PhaseGradientBranchCutRegression": "targeted",
    "QPotentialStencilBoundaryRegression": "targeted",
    "needNumericalErrorAudit": "improved_after_implementation",
    "needNontrivialPhysicsCase": "unpaid",
    "needFaithfulnessReview": "unchanged",
    "containsFinalTruthClaim": "absent",
    "BMC0BStatus": "specified_only",
    "SolverStatus": "not_implemented",
    "full_bmc_toy_gate": "blocked"
  },
  "verification_plan": [],
  "risks": [],
  "human_review_questions": [],
  "next_recommended_action": ""
}
```

## Verification commands to include in the plan

The plan should include these commands after implementation:

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

## Strict reminder

A branch-cut regression test is not physics validation.

A stencil-boundary regression test is not Friedmann recovery.

A stronger near-node audit is not BMC success.

This ticket should harden failure detection only.

Full BMC remains blocked.
