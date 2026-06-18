Approve **BMC-POST-0007 plan**, with two technical tightenings before implementation.

```json id="post0007_plan_review"
{
  "ticket": "BMC-POST-0007",
  "title": "Branch-Cut and Stencil Boundary Regression Hardening",
  "plan_status": "approved_with_tightening",
  "scope": "regression_hardening_only",
  "solver_allowed": false,
  "trajectory_allowed": false,
  "bmc0b_solver_allowed": false,
  "physics_promotion_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

## Tightening 1: branch-cut test must prove the failure mode

Your fixture is good:

```text id="branch_fixture"
Ψ(alpha, phi) = exp(i * (π - 0.5 * alpha))
sample near alpha = 0
expected ∂S/∂alpha ≈ -0.5
```

Make sure the test would fail under naïve `arg(Ψ)` central differencing. Around `alpha = 0`, one side crosses the principal-argument branch, so a bad implementation would produce a huge false derivative instead of `-0.5`. The test should assert not only “stable,” but also **numerically close to -0.5** within a named tolerance.

Suggested assertion:

```text id="branch_assertion"
abs(observed_dS_dalpha - (-0.5)) <= BranchCutGradientTolerance
```

Also confirm `dS/dphi ≈ 0` for this fixture.

## Tightening 2: stencil-boundary test must hit the stencil, not the center

For `TestQPotentialBlocksStencilPointBelowAmplitudeFloor`, make the mock wavefunction satisfy:

```text id="stencil_fixture"
amplitude(center) > NearNodeAmplitudeFloor
amplitude(alpha + h, phi) < NearNodeAmplitudeFloor
amplitude(alpha - h, phi) > NearNodeAmplitudeFloor
amplitude(alpha, phi + h) > NearNodeAmplitudeFloor
amplitude(alpha, phi - h) > NearNodeAmplitudeFloor
```

Then assert:

```text id="stencil_assertion"
Authoritative == false
Status == q_potential_blocked_by_domain_boundary
QPotential is not treated as evidence
```

This prevents the test from accidentally passing because the center point itself was near-node.

## Keep the plan narrow

Expected changes should remain:

```text id="expected_post0007_files"
internal/bmc/phaseaudit/hsensitivity_test.go
internal/bmc/qpotential/domain_test.go
docs/postmortem/bmc_post_0007_branchcut_stencil_regression_hardening.md
```

Only modify production code if one of these regression tests exposes a real gap.

## EBP status

```text id="post0007_ebp"
BMCPost0007Status: plan_approved_with_tightening
PhaseGradientBranchCutRegression: targeted
QPotentialStencilBoundaryRegression: targeted
needNumericalErrorAudit: improved_after_implementation
needNontrivialPhysicsCase: unpaid
needFaithfulnessReview: unchanged
containsFinalTruthClaim: absent
BMC0BStatus: specified_only
SolverStatus: not_implemented
full_bmc_toy_gate: blocked
```

Proceed with POST-0007 as **test-first regression hardening**, not solver work.
