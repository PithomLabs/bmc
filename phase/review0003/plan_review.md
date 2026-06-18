Approve **BMC-REVIEW-0003 plan** as documentation-only source-intake/provenance work.

```json id="bmc_review_0003_plan_review"
{
  "ticket": "BMC-REVIEW-0003",
  "title": "Literature Acquisition and Source-Provenance Intake",
  "plan_status": "approved",
  "scope": "documentation_only_source_intake",
  "solver_allowed": false,
  "numerical_results_allowed": false,
  "schema_or_cli_allowed": false,
  "physics_promotion_allowed": false,
  "operator_form_debt_retirement_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

The plan is correctly scoped. REVIEW-0003 should **inventory source needs and provenance gaps**, not perform equation acceptance or solver preparation.

One tightening before implementation:

```text id="review0003_tightening"
In Sections 4–6, distinguish three cases clearly:

1. source text available in repo
2. bibliographic reference exists but source text missing
3. no source/reference available yet

Only case 1 can support future equation intake.
Cases 2 and 3 must remain source_acquisition_required.
```

Current EBP status:

```text id="review0003_ebp"
BMCReview0003Status: plan_approved
SourceProvenanceIntake: targeted
OperatorFormDebt: documented_unpaid
MetricSignatureDebt: dependent_unpaid
ScalarPotentialNormalizationDebt: dependent_unpaid
FactorOrderingDebt: dependent_unpaid
UnitsConventionDebt: dependent_unpaid
BoundaryConditionDebt: dependent_unpaid
ResidualDefinitionDebt: dependent_unpaid
FaithfulnessReviewDebt: partial_documentation_only
HumanPhysicsReviewDebt: unpaid
containsFinalTruthClaim: absent
BMC0BStatus: specified_only
SolverStatus: not_implemented
full_bmc_toy_gate: blocked
```

Proceed with REVIEW-0003 as **source inventory and acquisition ledger only**.
