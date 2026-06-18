Approve **BMC-GATE-0001 plan** as documentation-first gate-design work.

```json id="bmc_gate_0001_plan_review"
{
  "ticket": "BMC-GATE-0001",
  "title": "Faithfulness Review and Solver-Design Gate",
  "plan_status": "approved_with_minor_tightening",
  "scope": "gate_design_only",
  "documentation_first": true,
  "solver_allowed": false,
  "trajectory_allowed": false,
  "bmc0b_solver_allowed": false,
  "physics_promotion_allowed": false,
  "new_cli_allowed": false,
  "new_schema_allowed": false,
  "new_output_artifact_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

The plan is correctly scoped. It defines the **conditions for future solver discussion**, not permission to implement the solver.

Two minor tightenings before implementation:

```text id="gate_tightening_1"
Rename or clarify “Future Work Requires a New Gate” as:
“Solver Implementation Requires a Separate Gate/Ticket”

Reason:
BMC-GATE-0001 is already a gate. The document should make clear that passing this gate only defines requirements; it does not authorize implementation.
```

```text id="gate_tightening_2"
Expand the forbidden phrase scan to include:
ready

Reason:
“ready” has repeatedly been treated as a dangerous promotion-adjacent word in this stack.
```

Updated scan suggestion:

```bash id="gate_scan"
grep -R "validated\|proved\|recovered\|ready\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" docs/gates docs/postmortem || true
```

Current EBP status:

```text id="bmc_gate_0001_ebp"
BMCGate0001Status: plan_approved_with_minor_tightening
FaithfulnessReviewGate: targeted
SolverDesignGate: targeted
OperatorFormDebt: unpaid
FactorOrderingDebt: unpaid
UnitsConventionDebt: unpaid
BoundaryConditionDebt: unpaid
NullModelDesignDebt: unpaid
ClassicalRecoveryCriterionDebt: unpaid
containsFinalTruthClaim: absent
BMC0BStatus: specified_only
SolverStatus: not_implemented
full_bmc_toy_gate: blocked
```

Proceed with BMC-GATE-0001 as a **gate document only**.
