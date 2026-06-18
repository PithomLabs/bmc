Approve **BMC-POST-0004.2 repair plan**, with one important tightening: your plan is directionally right, but it should explicitly require **field-specific exact tests**, not only “comprehensive assertions.” The prior re-review blockers were exact-set duplicate/cardinality gaps, broad status acceptance, phrase leakage, and dirty untracked `phase/post0004` files. 

```json id="post0004_2_plan_review"
{
  "ticket": "BMC-POST-0004.2",
  "title": "Exact-Set, Phrase-Safe, and Diff Hygiene Repair",
  "plan_status": "approved_with_tightening",
  "scope": "specification_only",
  "solver_allowed": false,
  "trajectory_allowed": false,
  "numerical_result_allowed": false,
  "new_cli_allowed": false,
  "full_bmc_toy_gate": "blocked"
}
```

## Required tightening before implementation

Make sure `spec_test.go` includes explicit tests for all four exact-set classes:

```text id="exact_set_tests"
duplicate FailureMode rejected
duplicate RequiredNullModel rejected
duplicate RequiredFaithfulnessReview rejected
duplicate ResidualNorm rejected

wrong cardinality FailureModes rejected
wrong cardinality RequiredNullModels rejected
wrong cardinality RequiredFaithfulnessReviews rejected
wrong cardinality ResidualNorms rejected
```

Also make sure the phrase-safe test checks **actual returned errors**, not just accepted documentation. The known bad phrase was `"solver-ready"`, so no validation error should contain:

```text id="forbidden_error_leaks"
ready
validated
proved
recovered
successful
```

## Diff hygiene tightening

Your plan mentions `git status --short`, but the repair must explicitly handle the previously reported untracked files:

```text id="dirty_files"
phase/post0004/adv_prompt.md
phase/post0004/walkthrough.md
```

They should be moved out of the repo, ignored intentionally, or explicitly excluded from the POST-0004.2 review boundary.

## Acceptance boundary

POST-0004.2 is acceptable only if:

```text id="post0004_2_acceptance"
1. exact-set validation rejects duplicates and wrong cardinality;
2. field-specific status validation rejects vague non-debt statuses;
3. phrase-safe errors do not leak forbidden terms;
4. forbidden-term tests cover representative stored fields;
5. only target files are modified, or extra files are explicitly justified;
6. go test ./internal/bmc/bmc0bspec -v -count=1 passes;
7. go test ./... -count=1 passes;
8. lake build passes;
9. no solver, trajectory, numerical BMC-0B result, recovery claim, validation claim, or full-BMC claim is introduced.
```

## EBP status

```text id="post0004_2_ebp"
BMC0BStatus: specified_only_pending_repair
OperatorFormDebt: explicit_unpaid_pending_field_specific_validation
ResidualGateDebt: explicit_unpaid_pending_exact_set_validation
NullModelDebt: explicit_unpaid_pending_exact_set_validation
SolverStatus: not_implemented
containsFinalTruthClaim: absent
full_bmc_toy_gate: blocked
```

Proceed with the repair. After implementation, it should still go through a focused re-review before POST-0004 is cleanly accepted.
