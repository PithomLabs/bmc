BMC-REVIEW-0004 is **accepted in substance** for equation-intake scope. 

```json id="bmc_review_0004_final_status"
{
  "ticket": "BMC-REVIEW-0004",
  "review_verdict": "accept_with_repairs",
  "substantive_status": "accepted_for_equation_intake_scope",
  "SourceTextAcquisition": "accepted_for_equation_intake_scope",
  "EquationIntake": "accepted_for_equation_intake_scope",
  "OperatorFormDebt": "documented_unpaid",
  "MetricSignatureDebt": "dependent_unpaid",
  "ScalarPotentialNormalizationDebt": "dependent_unpaid",
  "FactorOrderingDebt": "dependent_unpaid",
  "UnitsConventionDebt": "dependent_unpaid",
  "BoundaryConditionDebt": "dependent_unpaid",
  "ResidualDefinitionDebt": "dependent_unpaid",
  "FaithfulnessReviewDebt": "partial_documentation_only",
  "HumanPhysicsReviewDebt": "unpaid",
  "LeanProofDebt": "unpaid",
  "BMC0BStatus": "specified_only",
  "SolverStatus": "not_implemented",
  "containsFinalTruthClaim": "absent",
  "full_bmc_toy_gate": "blocked"
}
```

What is accepted:

```text id="review0004_accepted"
REVIEW-0004 correctly intakes Case 1 project-note/code-fixture equations only as unreviewed candidates.
It separates project-note, code-fixture, and peer-reviewed-source equation categories.
No peer-reviewed source equation is available.
Case 2/3 sources remain acquisition-required.
Notation and convention mappings remain human-review-dependent.
No equation is treated as accepted, faithful, or debt-retiring.
No solver, numerical result, schema, CLI, output artifact, recovery claim, or physics promotion was added.
```

Current stack position:

```text id="bmc_stack_position_review0004"
POST-0001 through POST-0008: closed in substance
GATE-0001: accepted for gate-design scope
REVIEW-0001: accepted for convention/provenance review scope
REVIEW-0002: accepted for operator-form review scope
REVIEW-0003: accepted for source-intake scope
REVIEW-0004: accepted for equation-intake scope
BMC-0B solver: blocked
OperatorFormDebt: documented_unpaid
full_bmc_toy_gate: blocked
```

Next best move:

```text id="next_ticket"
BMC-REVIEW-0005: Human Physics Review Checklist for Case 1 Equation Candidates
```

Purpose:

```text id="review0005_purpose"
Run a structured human-review checklist on EQ-KIMI-0001 through EQ-KIMI-0005 and EQ-CODE-0001, not to retire debt, but to classify which project-note/code-fixture equations are worth comparing against future peer-reviewed sources.
```

Parallel but separate task:

```text id="parallel_source_task"
Acquire full peer-reviewed Case 2 source text for the minisuperspace WdW operator, metric/signature convention, scalar potential normalization, and factor-ordering discussion.
```

Do **not** proceed to solver design or numerical BMC-0B work yet.
