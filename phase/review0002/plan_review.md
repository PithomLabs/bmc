Approve **BMC-REVIEW-0002 plan** as documentation-only operator-form literature faithfulness review.

```json id="bmc_review_0002_plan_review"
{
  "ticket": "BMC-REVIEW-0002",
  "title": "Operator-Form Literature Faithfulness Review",
  "plan_status": "approved",
  "scope": "documentation_only_literature_faithfulness_review",
  "solver_allowed": false,
  "numerical_results_allowed": false,
  "schema_or_cli_allowed": false,
  "physics_promotion_allowed": false,
  "operator_form_debt_retirement_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

The plan is correctly scoped. The important safeguard is that **missing external literature means OperatorFormDebt stays unpaid**, not partially retired. That matches the REVIEW-0001 provenance discipline, where the current fixture operator was recorded as a project diagnostic rather than a faithful continuous WdW operator. 

One tightening before implementation:

```text id="review0002_tightening"
In Section 5, if no usable literature source is present, label the literature candidates as acquisition requirements, not candidates.

Preferred wording:
operator_candidate_literature_required_001
status: external_literature_missing
debt_status: OperatorFormDebt: unpaid
carry_forward_status: source_acquisition_required

Avoid wording that makes missing candidates sound like real alternatives already identified.
```

Current EBP status:

```text id="review0002_ebp"
BMCReview0002Status: plan_approved
OperatorFormDebt: targeted_but_unpaid
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

Proceed with BMC-REVIEW-0002 as **source-availability and operator-form mapping only**.
