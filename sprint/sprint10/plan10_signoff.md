Sprint 10.3 can now be accepted **based on the reported implementation and verification**.

```json id="s10_3_acceptance"
{
  "sprint": "BMC Sprint 10.3",
  "artifact": "candidate local-branch residual runner",
  "acceptance_status": "accepted_after_repairs_reported",
  "promotion_status": "promoted_candidate_residual_runner_artifact_after_repairs",
  "physics_claim_promoted": false,
  "full_bmc_toy_gate": "blocked",
  "recovery_claim": false,
  "bmc_beats_null_models_claim": false,
  "scientific_novelty_claim": false
}
```

## Why this passes

The remaining Sprint 10.2 repair targets were addressed:

```text id="s10_3_pass"
1. residual_input_points now serialize lambda explicitly.
2. num_evaluation_points == len(residual_input_points) is enforced.
3. comparison target IDs are dynamically collected from actual computed residual diagnostics.
4. hard-coded candidate_residual_branch_0 targeting is removed.
5. residual input points now validate lambda/alpha/phi/LHS/RHS finiteness.
6. lambda monotonicity is checked locally inside each computed diagnostic.
7. input provenance is restricted to file_read or derived_from_file_read.
8. forbidden formula IDs are rejected phrase-safely.
9. regression tests were added.
10. Go tests, CLI validation/summarization, and Lean build passed.
```

This closes the main Sprint 10 chain:

```text id="s10_chain_status"
Sprint 10: rejected for fixed residual constants
Sprint 10.1: rejected for hidden constants + synthetic branch
Sprint 10.2: accepted with repairs; real per-point calculation mostly established
Sprint 10.3: accepted after comparison/validator hardening
```

## What Sprint 10.3 now earns

It earns this:

```text id="earned"
BMC-0A now has a candidate local-branch residual runner whose computed diagnostics are reported as derived from file-backed per-point branch data, with validator checks for point-level auditability and dynamic comparison targeting.
```

It does **not** earn this:

```text id="not_earned"
Friedmann recovery
classical-limit recovery
BMC validation
BMC beats null models
global cosmology claim
full BMC promotion
```

## Worktree note

The uncommitted/untracked artifacts are not an EBP blocker, but they are an engineering hygiene issue before the next sprint. Before Sprint 11, I would snapshot the state:

```bash id="git_snapshot"
git status --short
git add internal/bmc/residualrun cmd/ptw-bmc BMC/BMC/ResidualRun.lean BMC/BMC.lean internal/bmc/clockseg
git commit -m "Repair BMC residual runner faithfulness and validation"
```

Only include generated `out/` artifacts if your repo policy tracks them.

## Updated EBP debt status

```text id="s10_3_debt"
needLiteratureAudit: partial
needMap: partial
needInvariant: partial
needToyCheck: partial
needNullModel: partial
needObstruction: partial
needFaithfulnessReview: contested
clock_choice_debt: unpaid
classical_target_debt: unpaid
unit_convention_debt: unpaid
sign_convention_debt: unpaid
normalization_debt: unpaid
containsFinalTruthClaim: absent

ResidualRunDiagnosticIntegrity: retired for Sprint 10 local-runner scope
ResidualNullComparisonIntegrity: retired for Sprint 10 local-runner scope
SourceProvenanceIntegrity: retired for Sprint 10 local-runner scope
LocalOnlyBoundary: retired
NoRecoveryClaimBoundary: retired
NoBMCBeatsNullsBoundary: retired
ConventionDebtVisibility: retired for visibility only, not debt retirement
LeanPolicyBoundary: retired for policy-only scope
full_bmc_toy_gate: blocked
```

## Roadmap position now

We are still in:

```text id="tier_position"
Tier 1: BMC Toy Benchmark
```

But we have now crossed a meaningful milestone:

```text id="milestone"
The BMC toy benchmark can produce a guarded candidate residual diagnostic from local branch trajectory data and compare it diagnostically against null-layer outputs without claiming recovery.
```

The next safe phase is **not** “claim success.” It is:

```text id="next_phase"
Sprint 11: Residual/null comparison audit and interpretation hardening
```

The goal of Sprint 11 should be to audit whether the residual/null comparison is meaningful, stable, and non-decorative across perturbations—not to promote physics.
