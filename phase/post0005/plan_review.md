Approve **BMC-POST-0005 plan**, with required tightening. It is correctly scoped as a **BMC-0A numerical h-sensitivity audit**, not BMC-0B solver work.

```json id="post0005_plan_review"
{
  "ticket": "BMC-POST-0005",
  "title": "Phase Gradient h-Sensitivity Audit",
  "plan_status": "approved_with_tightening",
  "scope": "numerical_audit_only",
  "bmc0b_solver_allowed": false,
  "trajectory_solver_allowed": false,
  "new_cli_allowed": false,
  "new_schema_allowed": false,
  "new_output_artifacts_allowed": false,
  "promotion_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

One precondition remains: POST-0004.2 had only a diff-hygiene repair left, specifically untracked `phase/post0004/*` files. Confirm that cleanup before POST-0005 implementation. The final review allowed POST-0004 acceptance only after that cleanup. 

## Required tightenings

Add these to the implementation plan:

```text id="post0005_tightening"
1. Separate phase-gradient convention from guidance-velocity convention.
2. Guard against phase branch-cut artifacts.
3. Make near-node blocking explicit and non-authoritative.
4. Add explicit h-ladder ordering tests.
5. Add manual grep for “ready” too.
6. Run package regression tests for bmc0bspec, wdw, report, and convergence, not only phaseaudit and ./...
```

## Phase convention warning

For plane wave control, the audit should compare the **phase gradient** only:

```text id="phase_gradient_convention"
∂S/∂alpha = k
∂S/∂phi = omega
```

Do **not** mix this with Bohmian guidance velocity signs such as:

```text id="guidance_convention"
dphi/dlambda = -∂S/∂phi
```

That guidance-sign convention belongs to trajectory dynamics, not this h-sensitivity audit unless explicitly reviewed.

## Branch-cut warning

If `S = arg(Ψ)` is finite-differenced directly, phase wrapping can create fake jumps near `π ↔ -π`.

Safer options:

```text id="phase_wrap_options"
1. Use the existing project phase-gradient function if already reviewed.
2. If computing directly, document branch-cut behavior and add a test that detects wrap sensitivity.
3. Consider an internal derivative form like Im((∂ψ)/ψ) only if it matches existing conventions and is clearly documented as numerical audit machinery, not new physics authority.
```

## Verification command tightening

Use the full set:

```bash id="post0005_verification"
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/phaseaudit -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/bmc0bspec -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/report -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/convergence -v -count=1
GOCACHE=/tmp/go-build-cache go test ./... -count=1
cd BMC && /home/chaschel/.elan/bin/lake build
```

And include `ready` in the forbidden scan:

```bash id="post0005_forbidden_scan"
grep -R "validated\|proved\|recovered\|ready\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" internal/bmc/phaseaudit docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md || true
```

## EBP status

```text id="post0005_ebp"
BMCPost0005Status: plan_approved_with_tightening
PhaseGradientHSensitivity: audit_only
needNumericalErrorAudit: targeted_for_phase_gradient_scope
needNontrivialPhysicsCase: unpaid
needFaithfulnessReview: unchanged
containsFinalTruthClaim: absent
BMC0BStatus: specified_only
SolverStatus: not_implemented
full_bmc_toy_gate: blocked
```

Proceed only after POST-0004.2 diff hygiene is clean.
