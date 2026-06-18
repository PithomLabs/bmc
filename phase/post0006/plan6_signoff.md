POST-0006 is **accepted in substance** for its bounded audit scope.

```json id="post0006_status"
{
  "ticket": "BMC-POST-0006",
  "review_verdict": "accept_with_repairs",
  "substantive_status": "accepted_for_qpotential_near_node_audit_scope",
  "QPotentialNearNodeDomainBoundary": "accepted_for_qpotential_near_node_audit_scope",
  "needNumericalErrorAudit": "improved_for_qpotential_near_node_scope",
  "BMC0BStatus": "specified_only",
  "SolverStatus": "not_implemented",
  "containsFinalTruthClaim": "absent",
  "full_bmc_toy_gate": "blocked"
}
```

What is now accepted:

```text
legacy Q near-node clamp identified
near-node Q no longer authoritative through audit path
central + stencil amplitude checks implemented
nonfinite wave/derivative/Q cases blocked
zero-clamping regression test present
no-promotion fields present
forbidden inference audit present
no CLI/schema/output artifact added
tests and Lake build pass
```

Next best technical move:

```text
BMC-POST-0007: Branch-Cut and Stencil Boundary Regression Hardening
```

Reason: POST-0005 and POST-0006 both introduced strong audit wrappers, but the optional review items point to the same next hardening theme: explicit regression tests for branch-cut behavior, stencil-boundary edge cases, and top-level invalid-input paths. This keeps us in numerical-audit hardening rather than jumping to solver work.

Current EBP boundary:

```text
No BMC validation.
No Friedmann recovery.
No BMC-0B solver.
No classical-limit claim.
Full BMC remains blocked.
```
