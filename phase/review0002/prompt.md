# BMC-REVIEW-0002 Prompt: Operator-Form Literature Faithfulness Review

You are implementing the next documentation-only review ticket for the Go-based BMC/PTW repository under strict EBP 2.1 discipline.

## Ticket

```text
BMC-REVIEW-0002: Operator-Form Literature Faithfulness Review
```

## Context

BMC-REVIEW-0001 created the first operator, units, convention, and provenance review:

```text
docs/reviews/bmc_review_0001_operator_units_conventions.md
```

That review accepted only a documentation/provenance scope. It explicitly left the following debts unpaid:

```text
OperatorFormDebt
FactorOrderingDebt
UnitsConventionDebt
BoundaryConditionDebt
MetricSignatureDebt
ScalarPotentialNormalizationDebt
WavefunctionAdmissibilityDebt
ResidualDefinitionDebt
NullModelDesignDebt
ClassicalRecoveryCriterionDebt
FaithfulnessReviewDebt
HumanPhysicsReviewDebt
LeanProofDebt
```

BMC-REVIEW-0002 should target the most upstream debt first:

```text
OperatorFormDebt
```

Most later work depends on the exact Wheeler-DeWitt operator form being reviewed against literature.

## Purpose

Create a documentation-only literature faithfulness review that maps candidate minisuperspace Wheeler-DeWitt operator forms against the current project fixture.

The review should answer:

```text
What WdW operator forms for homogeneous scalar-field cosmology appear in the available sources?
How do their signs, metric/signature choices, factor orderings, potential terms, and unit conventions differ?
Which parts of the current PTW/BMC fixture match a literature-backed form?
Which parts remain project-only assumptions?
What operator-form candidates should be carried forward for human physics review?
```

This ticket must not implement a solver.

It must not produce numerical BMC-0B results.

It must not retire OperatorFormDebt unless explicit source provenance and human-review criteria are satisfied.

## Strict EBP boundary

This is a **literature faithfulness review only**.

Do not implement:

```text
BMC-0B solver
massive scalar numerical solution
trajectory solver
Friedmann residual computation
classical-limit recovery
new benchmark result artifact
new report schema
new CLI route
generated JSON output artifact
public-facing success dashboard
physics promotion
null-model victory
BMC superiority
scientific novelty
full BMC promotion
```

Allowed scope:

```text
documentation
operator-form comparison table
source/provenance ledger
equation-mapping table
notation/convention mapping
debt ledger
human-review questions
no code unless absolutely necessary
no CLI
no schema
no output artifact
```

Prefer documentation-only.

## Expected primary artifact

Create:

```text
docs/reviews/bmc_review_0002_operator_form_literature_faithfulness.md
```

Do not add code by default.

## Source and provenance rules

Use only:

```text
1. Sources already present in the repository.
2. Equations or references already cited in project documentation.
3. User-provided source text, if available.
```

If external literature is not available in the repository, do **not** invent citations, equations, or bibliographic details.

Instead, create a source-debt section:

```text
external_literature_missing: true
source_acquisition_required: true
human_literature_review_required: true
```

Allowed provenance statuses:

```text
project_doc
code_fixture
repo_source
external_literature_missing
human_review_required
```

Allowed review statuses:

```text
imported_from_project_docs
assumed_by_fixture
mapped_to_available_source
requires_literature_confirmation
requires_human_physics_review
unreviewed
contested
```

Forbidden statuses:

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

These forbidden terms may appear only in explicitly negated/no-promotion/rejected-example sections.

## Files to inspect

Inspect at minimum:

```text
docs/reviews/bmc_review_0001_operator_units_conventions.md
docs/gates/bmc_gate_0001_faithfulness_solver_design_gate.md
docs/postmortem/bmc_post_0008_remediation_stack_closure_ledger.md
docs/postmortem/bmc_post_0004_bmc0b_massive_scalar_wdw_spec.md
docs/postmortem/bmc_post_0001_constraint_violation_detection.md
docs/postmortem/bmc_post_0002_numerical_wdw_residual_integration.md
docs/postmortem/bmc_post_0003_euler_rk4_dt_convergence.md
docs/postmortem/bmc_post_0005_phase_gradient_h_sensitivity.md
docs/postmortem/bmc_post_0006_qpotential_near_node_domain_boundary.md
docs/postmortem/bmc_post_0007_branchcut_stencil_regression_hardening.md
```

Inspect code only to document current fixture behavior, not to infer physics truth:

```text
internal/bmc/wdw
internal/bmc/wave
internal/bmc/report
internal/bmc/bmc0bspec
internal/bmc/qpotential
internal/bmc/phaseaudit
```

## Required review sections

The review document must contain:

```text
1. Review purpose and non-goals
2. Current accepted boundary from REVIEW-0001
3. Source availability audit
4. Current project fixture operator map
5. Literature operator-form candidates
6. Notation and convention mapping
7. Metric/signature mapping
8. Kinetic-term comparison
9. Scalar-potential placement comparison
10. Factor-ordering comparison
11. Units/constants implications
12. Boundary/domain implications
13. Residual-definition implications
14. Impact on Bohmian guidance and Q-potential quantities
15. Operator-form candidate carry-forward list
16. OperatorFormDebt status
17. Remaining dependent debts
18. Human physics-review questions
19. No-promotion audit
20. Next allowed ticket recommendation
```

## Required source availability audit

The review must state whether usable literature sources are present in the repository.

Create a table:

```text
Source or citation
Location in repo
Source type
Usable equation present?
Relevant to BMC-0B operator form?
Status
Debt
```

If no usable external literature is present, state:

```text
No external literature equation is available in-repo for retiring OperatorFormDebt.
This review can only map current project assumptions and define literature-acquisition requirements.
OperatorFormDebt remains unpaid.
```

## Required current fixture operator map

Document the current project fixture operator without promoting it.

At minimum include:

```text
Current fixture residual:
-d2Psi/dAlpha2 + d2Psi/dPhi2

Status:
toy diagnostic / project fixture only

Not yet faithful because:
continuous operator form not sourced
metric/signature not fixed
factor ordering not fixed
massive scalar potential not included
units/constants not fixed
boundary conditions not fixed
```

Do not call the fixture a faithful WdW operator.

## Required literature operator-form candidate table

Create a table with columns:

```text
Candidate ID
Source/provenance
Equation form
Variables
Metric/signature convention
Kinetic term
Potential term
Factor ordering
Units/constants
Boundary assumptions
Relation to current fixture
Debt status
Human review required
```

If real source equations are unavailable, include placeholder rows:

```text
operator_candidate_project_fixture_only
operator_candidate_literature_required_001
operator_candidate_literature_required_002
```

and mark them as:

```text
requires_literature_confirmation
external_literature_missing
human_review_required
```

## Required notation/convention mapping

Create a mapping table:

```text
Project notation
Source notation
Meaning
Mapping confidence
Debt status
Human review question
```

Include at least:

```text
alpha
a
phi
mass parameter
potential term
lapse/time or clock parameter
WdW wavefunction Psi
residual operator
```

If source notation is missing, mark:

```text
source_not_available
requires_literature_confirmation
```

## Required metric/signature mapping

The review must identify whether current docs fix the minisuperspace metric/signature.

If not fixed, state:

```text
MetricSignatureDebt remains unpaid.
OperatorFormDebt cannot be retired until metric/signature is fixed.
```

Include a table:

```text
Candidate signature
Effect on kinetic signs
Effect on WdW operator
Relation to current fixture
Debt status
```

## Required scalar-potential placement comparison

Review whether the massive scalar potential term is specified.

If not specified, state:

```text
ScalarPotentialNormalizationDebt remains unpaid.
BMC-0B operator form cannot be finalized without potential normalization.
```

Include table:

```text
Potential candidate
Placement in Hamiltonian constraint
Placement in WdW operator
Units/constants required
Source/provenance
Debt status
```

Do not invent the final potential normalization.

## Required factor-ordering comparison

Include:

```text
current fixture ordering status
available candidate orderings
missing literature orderings
impact on residual
impact on boundary behavior
impact on Bohmian guidance/Q-potential
debt status
```

If no literature ordering is available, state:

```text
FactorOrderingDebt remains unpaid.
No solver design may assume this debt retired.
```

## Required impact assessment

Explain how different operator forms may affect:

```text
numerical residual
boundary residual
near-node behavior
phase-gradient interpretation
Q-potential form
null-model design
classical recovery criterion
```

But do not run tests or produce numerical results.

## Required carry-forward candidate list

End with a carry-forward list:

```text
Candidate carried forward
Why carried forward
Debt status
Required source review
Required human review
Blocked implementation reason
```

At minimum, carry forward:

```text
operator_candidate_project_fixture_only
```

but keep it marked:

```text
project_fixture_only
not faithful yet
requires literature review
blocks solver implementation
```

## Required no-promotion audit

The review must state:

```text
BMC-REVIEW-0002 does not implement BMC-0B.
BMC-REVIEW-0002 does not solve the Wheeler-DeWitt equation.
BMC-REVIEW-0002 does not validate BMC.
BMC-REVIEW-0002 does not recover Friedmann dynamics.
BMC-REVIEW-0002 does not prove classical-limit recovery.
BMC-REVIEW-0002 does not defeat null models.
BMC-REVIEW-0002 does not unblock full BMC.
```

Maximum allowed EBP stage:

```text
BMC-REVIEW-0002 accepted_for_operator_form_review_scope
```

## Expected implementation walkthrough

After implementation, provide:

```json
{
  "ticket": "BMC-REVIEW-0002",
  "implementation_status": "complete",
  "scope": "operator_form_literature_faithfulness_review_only",
  "files_added": [],
  "files_modified": [],
  "summary": "",
  "source_availability_status": "",
  "operator_form_status": "",
  "metric_signature_status": "",
  "scalar_potential_status": "",
  "factor_ordering_status": "",
  "carried_forward_candidates": [],
  "remaining_debts": [],
  "forbidden_inference_scan": "",
  "verification_results": [],
  "ebp_debt_status": {
    "BMCReview0002Status": "implemented_pending_review",
    "OperatorFormDebt": "documented_unpaid",
    "MetricSignatureDebt": "documented_unpaid",
    "ScalarPotentialNormalizationDebt": "documented_unpaid",
    "FactorOrderingDebt": "documented_unpaid",
    "UnitsConventionDebt": "documented_unpaid",
    "BoundaryConditionDebt": "documented_unpaid",
    "ResidualDefinitionDebt": "documented_unpaid",
    "FaithfulnessReviewDebt": "partial_documentation_only",
    "HumanPhysicsReviewDebt": "unpaid",
    "containsFinalTruthClaim": "absent",
    "BMC0BStatus": "specified_only",
    "SolverStatus": "not_implemented",
    "full_bmc_toy_gate": "blocked"
  },
  "next_recommended_action": ""
}
```

## Verification plan

Run:

```bash
GOCACHE=/tmp/go-build-cache go test ./... -count=1
cd BMC && /home/chaschel/.elan/bin/lake build
```

Run forbidden phrase scan:

```bash
grep -R "validated\|proved\|recovered\|ready\|successful\|physics_success\|bmc_validated\|friedmann_recovered\|quantum_gravity_progress\|full bmc unblocked\|bmc beats nulls" docs/reviews docs/gates docs/postmortem || true
```

Matches are acceptable only in explicitly negated/no-promotion/rejected-example contexts.

## Strict reminder

A literature review is not a solver.

A sourced operator candidate is not a physics result.

A candidate operator form is not accepted until provenance, convention mapping, and human physics review retire the debt.

Full BMC remains blocked.
