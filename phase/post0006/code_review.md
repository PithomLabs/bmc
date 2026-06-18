{
  "summary": "POST-0006 adds a strict quantum-potential near-node domain-boundary audit that prevents clamped Q=0 from being treated as authoritative. `AuditPoint` checks the evaluation point plus all four central-difference stencil offsets for amplitude below `NearNodeAmplitudeFloor` (1e-8) and nonfinite values, blocks with explicit authority metadata and distinct statuses, and rejects invalid inputs early. Legacy `qpotential.Q` still clamps near-node values to 0.0, but the audit path exposes this as non-authoritative rather than authoritative evidence. Sample fixtures, forbidden-term phrase-safety, no-promotion EBP fields, and negative-case tests are all in place. Tests across `qpotential`, `phaseaudit`, `bmc0bspec`, `wdw`, `report`, `convergence`, and full workspace pass; Lean lake build passes. Remaining issue is diff hygiene: untracked `phase/post0006/` boundary files plus the new doc/target files remain untracked in the repo, failing the clean-handoff boundary.",
  "overall_verdict": "accept_with_repairs",
  "ebp_debt_review": {
    "BMCPost0006Status": "audit_only",
    "QPotentialNearNodeDomainBoundary": "accepted_for_qpotential_near_node_audit_scope",
    "needConstraintViolationTests": "clean",
    "needNumericalErrorAudit": "improved_for_qpotential_near_node_scope",
    "needNontrivialPhysicsCase": "clean",
    "needToyCheck": "clean",
    "needFaithfulnessReview": "unchanged",
    "containsFinalTruthClaim": "absent",
    "BMC0BStatus": "specified_only",
    "SolverStatus": "not_implemented",
    "NoRecoveryClaimBoundary": "clean",
    "NoSchemaCLIBloat": "clean",
    "full_bmc_toy_gate": "blocked"
  },
  "existing_q_behavior_findings": [
    "Legacy `qpotential.Q` still clamps `rVal < 1e-12` to `0.0` without authority metadata.",
    "The new `AuditPoint` wrapper is strict and independent; it never uses `qpotential.Q` for its authoritative path and surfaces non-authoritative results for the prior failure mode.",
    "Because the prior clamping is contained in `q` and wrapped safely by the audit path, near-node Q=0 is no longer authoritative in the POST-0006 surface."
  ],
  "authority_metadata_findings": [
    "`QPotentialAuditResult` includes `QPotential`, `Amplitude`, `Status`, `Authoritative`, and `Reason`.",
    "Seven explicit statuses are defined: authoritative, blocked-by-node-contact, blocked-by-near-node-amplitude, blocked-by-nonfinite-wave, blocked-by-nonfinite-derivative, blocked-by-domain-boundary, and audit-only-no-promotion.",
    "No CLI/schema/output bloat was added; authority metadata is purely in-memory audit output."
  ],
  "amplitude_floor_findings": [
    "Named constant `NearNodeAmplitudeFloor = 1e-8` is used.",
    "Exact node contact (`rVal == 0.0 || rVal < 1e-12`) is separately blocked before the floor check.",
    "Near-node blocking happens before Q is treated as usable evidence."
  ],
  "stencil_domain_findings": [
    "`AuditPoint` evaluates amplitude at the central point plus four stencil offsets (`alpha ± h`, `phi ± h`) using `QPotentialDerivativeStep = 1e-4`.",
    "Stencil amplitudes below the floor trigger `StatusQPotentialBlockedByDomainBoundary` with `Authoritative = false`.",
    "Stencil nonfinite amplitudes also block with a nonfinite-wave status."
  ],
  "nonfinite_blocking_findings": [
    "NaN/Inf alpha/phi coordinates are rejected at the `RunAudit` input layer.",
    "Nonfinite wavefunction amplitudes at the central or stencil points are blocked.",
    "Nonfinite second derivatives and nonfinite computed Q are blocked.",
    "Invalid/NaN/Inf derivative step `h` is rejected before evaluation."
  ],
  "zero_clamping_regression_findings": [
    "`TestQPotentialDoesNotClampNearNodeToAuthoritativeZero` confirms exact node contact returns `Authoritative=false`, `StatusQPotentialBlockedByNodeContact`, and emits `QPotential = 0.0` only as a non-authoritative flagged value.",
    "This directly covers the prior reported failure mode."
  ],
  "sample_fixture_findings": [
    "`plane_wave_control`: tested via `wave.NewPlaneWave`; finite Q only when safely away from nodes, and `Authoritative=true` is gated by amplitude/domain checks.",
    "`superposition_safe_or_audit_only`: `wave.NewSuperpositionWave` tested as `audit_only_no_promotion`; finite Q but explicitly non-authoritative.",
    "`near_node_probe`: node wavefunction at near-zero coordinates returns blocked/non-authoritative status."
  ],
  "no_promotion_field_findings": [
    "`QPotentialAudit` sets `toy_analysis_only=true`, `physics_claim='none'`.",
    "`QPotentialEBPStatus` sets `toy_analysis_only=true`, `physics_claim='none'`, `bmc0b_impact='none'`, `friedmann_recovery_impact='none'`, `promotion_recommendation='do_not_promote'`.",
    "Tests verify all no-promotion fields explicitly."
  ],
  "forbidden_inference_findings": [
    "`TestQPotentialForbiddenInferenceAudit` performs case-insensitive phrase checks on validation errors and accepted audit/EBP output fields.",
    "Forbidden terms appear only inside test arrays as rejected fixtures; no accepted status, doc claim, or EBP field contains forbidden language."
  ],
  "test_findings": [
    "Required tests present and passing: plane-wave control authoritative, superposition safe audit-only, near-node blocked, zero-clamp regression, nonfinite point rejection, invalid derivative step rejection, nonfinite derivative blocking, deterministic serialization, no-promotion fields, and forbidden inference audit.",
    "`go test ./internal/bmc/qpotential -v -count=1` and full `go test ./... -count=1` pass.",
    "`lake build` passes."
  ],
  "documentation_findings": [
    "`docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md` clearly states the audit-only scope: no BMC-0B, no WdW solve, no BMC validation, no Friedmann recovery, no classical-limit proof, no null models, no full BMC unblock.",
    "It explains the near-node instability, amplitude floor, derivative step, stencil checks, authority model, and unpaid residual debts."
  ],
  "schema_cli_output_findings": [
    "No new CLI route found in `cmd/ptw-bmc/main.go` or elsewhere.",
    "No new report schema or generated JSON output artifact in `out/` attributable to `qpotential`.",
    "Audit is in-memory; all existing `out/bmc0a_*` files predate this ticket."
  ],
  "diff_hygiene_findings": [
    "`git diff --stat` shows no modifications to already-tracked files.",
    "`git status --short` shows untracked target files plus untracked `phase/post0006/` boundary files (`code.md`, `prompt.md`).",
    "These boundary-process artifacts are outside the review target and replicate the prior diff-hygiene problem pattern, so final acceptance should require cleanup or explicit boundary justification."
  ],
  "test_and_build_findings": [
    "`go test ./internal/bmc/qpotential -v -count=1`: PASS",
    "`go test ./internal/bmc/phaseaudit -v -count=1`: PASS",
    "`go test ./internal/bmc/bmc0bspec -v -count=1`: PASS",
    "`go test ./internal/bmc/wdw -v -count=1`: PASS",
    "`go test ./internal/bmc/report -v -count=1`: PASS",
    "`go test ./internal/bmc/convergence -v -count=1`: PASS",
    "`go test ./... -count=1`: PASS",
    "`cd BMC && /home/chaschel/.elan/bin/lake build`: PASS"
  ],
  "missing_tests": [],
  "required_repairs_before_acceptance": [
    "Remove, move, or explicitly justify untracked boundary files in `phase/post0006/` so the final acceptance boundary matches the expected target files."
  ],
  "optional_repairs": [
    "Add a regression-style test where a stencil point falls just below the amplitude floor to confirm the domain-boundary blocked path is exercised directly.",
    "Add a test that exercises the top-level `RunAudit` with a nonfinite derivative-step path to confirm the input rejection is covered at the integration boundary."
  ],
  "promotion_recommendation": "accept_post_0006_for_qpotential_near_node_audit_scope",
  "next_smallest_useful_move": "Clean untracked `phase/post0006` boundary files and finalize POST-0006 acceptance. Do not promote, validate BMC, implement BMC-0B solver, produce numerical BMC-0B results, or claim Friedmann/classical-limit recovery."
}