You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 10.2: Candidate Local-Branch Residual Runner — True Input Faithfulness Repair**.

Sprint 10.2 claims to repair the Sprint 10.1 faithfulness failure by removing hidden residual constants, deleting synthetic branch injection, adding per-point trajectory inputs, and computing candidate residual diagnostics from actual local-branch point data.

Your job is to determine whether that claim is actually true.

## Accepted prior context

Accepted prior artifacts:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A superposition control + node-obstruction artifact
Sprint 3: BMC-0A robustness audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A candidate residual specification artifact
Sprint 7: BMC-0A null-model scaffold artifact
Sprint 8-Lite: BMC-0A prior-art boundary note
Sprint 9 + 9.1: BMC-0A null-model runner artifact accepted after repairs
Sprint 10: rejected for faithfulness issues
Sprint 10.1: rejected for hidden constants and synthetic branch provenance
```

Sprint 10.1 was rejected because:

```text
1. residual magnitudes were derived from embedded constants scaled by clock_range
2. the runner fabricated an extra blocked branch not present in the file-backed clock-readiness artifact
3. calculation provenance overclaimed the actual computation
```

Sprint 10.2 reportedly repairs this by:

```text
1. computing residual values from per-point trajectory data
2. using finite differences over alpha, phi, and lambda
3. removing hidden residual magnitude constants
4. removing unconditional synthetic branch injection
5. adding source branch registry validation
6. blocking old artifacts without trajectory points
7. making CLI summary distinguish computed and blocked diagnostics
```

## Reported files changed

Review actual source and generated artifacts, not just the walkthrough:

```text
internal/bmc/residualrun/report.go
internal/bmc/residualrun/residual.go
internal/bmc/residualrun/validate.go
internal/bmc/residualrun/residualrun_test.go
cmd/ptw-bmc/main.go
BMC/BMC/ResidualRun.lean
BMC/BMC.lean
out/bmc0a_local_residual.json
out/bmc0a_clock_readiness.json
out/bmc0a_friedmann_spec.json
out/bmc0a_nullrun.json
out/bmc0a_prior_art_boundary.json
```

Also check whether these files were actually changed or already contained point data:

```text
internal/bmc/clockseg/branches.go
internal/bmc/clockseg/local_relations.go
```

The Sprint 10.2 plan said clockseg would be modified to emit trajectory points. The walkthrough did not list those clockseg files, so verify the true provenance of `points` in `out/bmc0a_clock_readiness.json`.

## Core review question

```text
Does Sprint 10.2 honestly compute candidate local-branch residual metrics from per-point file-backed trajectory data, without hidden constants, synthetic source-backed branches, recovery claims, superiority claims, or global-cosmology promotion?
```

## Hard forbidden claims

The artifact must not claim or imply:

```text
Friedmann recovery
recovery of Friedmann
classical cosmology recovered
classical limit achieved
BMC validates Friedmann behavior
BMC beats null models
BMC outperforms controls
null models passed
null models failed
winner
validated physics
confirmed recovery
ready for recovery
full BMC validated
full BMC promoted
scientifically novel
breakthrough
problem of time solved
quantum gravity validated
```

Acceptable wording:

```text
candidate residual diagnostic
candidate local-branch residual
candidate local branch velocity constraint residual
residual-style diagnostic
diagnostic comparison only
local branch only
full BMC remains blocked
no recovery claim is made
no BMC-beats-null-models claim is made
convention debts remain unpaid or contested
blocked by missing residual inputs
candidate residual runner candidate only
```

## 1. Scope review

Check that Sprint 10.2 did not add or imply:

```text
global cosmology runner
classical-limit recovery logic
Friedmann recovery engine
BMC superiority ranking
full BMC promotion
scientific novelty claim
paper ingestion
claim extraction
general physics profile framework
dashboard or leaderboard
```

Allowed scope:

```text
candidate local-branch residual diagnostics
file-backed source artifact reading
local branch eligibility checks
per-point residual input extraction
finite-difference residual calculation
source branch registry
calculation provenance ledger
convention debt ledger
residual/null diagnostic comparison records
policy-only Lean safety contracts
```

Mark as blocker if the sprint claims recovery, validation, superiority, or global cosmology.

## 2. Report identity and hard fields

Expected identity:

```text
schema_version = bmc0a-local-residual-v0.1
artifact_kind = candidate_local_branch_residual_runner
scope = bmc0a_only
```

Expected hard fields:

```text
toy_analysis_only = true
final_truth_claim = false
residual_recovery_claim = false
scientific_novelty_claim_made = false
bmc_beats_null_models_claim = false
full_bmc_toy_gate = blocked
local_branch_only = true
global_cosmology_claim = false
ebp_debt_vocabulary = ptw_adversarial_review_debt_status_v0.1
containsFinalTruthClaim = absent
needFaithfulnessReview = contested
promotion_status = candidate_residual_runner_candidate_only
```

Important conditional field:

```text
candidate_residual_computed must be true only when at least one residual diagnostic is genuinely computed from eligible source-backed local-branch point inputs.
```

Mark as blocker if any hard overclaim field can be changed and still validate.

## 3. Clockseg point provenance review

Verify how trajectory points enter `out/bmc0a_clock_readiness.json`.

Check:

```text
Does LocalRelationBranch actually include points?
Does ExtractLocalRelationBranch or equivalent populate those points from the original trajectory?
Are points serialized into bmc0a_clock_readiness.json?
Are points file-backed rather than generated inside residualrun?
```

Possible outcomes:

```text
A. clockseg was modified and points are genuinely emitted by Sprint 5 clock-readiness generation
B. points already existed before Sprint 10.2 and are legitimately file-backed
C. residualrun injects or reconstructs points without clear provenance
```

Mark as blocker if C is true and the report claims file-backed local-branch inputs.

## 4. Hidden-constant audit

This is the highest-priority review.

Inspect `internal/bmc/residualrun/report.go` and helper functions.

Look for any remaining constants or fallback numbers used to compute:

```text
mean_abs_residual
max_abs_residual
rms_residual
candidate_left_hand_side
candidate_right_hand_side
residual values
```

Previously bad constants included:

```text
0.208
0.502
0.251
```

But also search for any new disguised constants.

A computed residual metric must be derived from a residual series:

```text
for each adjacent point pair:
    delta_lambda = lambda[i+1] - lambda[i]
    d_alpha = (alpha[i+1] - alpha[i]) / delta_lambda
    d_phi = (phi[i+1] - phi[i]) / delta_lambda
    lhs = d_alpha^2
    rhs = d_phi^2
    residual = lhs - rhs
then:
    mean_abs_residual = mean(abs(residual))
    max_abs_residual = max(abs(residual))
    rms_residual = sqrt(mean(residual^2))
```

The exact finite-difference scheme may differ, but the formula must be explicit and input-derived.

Mark as blocker if residual magnitudes are still constants, fixture numbers, or metadata-scaled constants.

## 5. Sensitivity test review

Confirm a real sensitivity test exists:

```text
TestResidualRunMetricsChangeWhenInputBranchDataChanges
```

The test must modify actual trajectory values such as:

```text
alpha
phi
lambda
candidate per-point input values, if used
```

It is not enough to modify:

```text
clock_range
branch count
metadata
labels
warnings
```

The test must prove at least one of these changes:

```text
mean_abs_residual
max_abs_residual
rms_residual
residual input point values
```

Mark as blocker if metric sensitivity is not tested or only metadata-sensitive.

## 6. Finite-difference input validation

Review whether the runner blocks or rejects:

```text
fewer than 2 trajectory points
duplicate lambda values
delta_lambda = 0
nonfinite lambda
nonfinite alpha
nonfinite phi
non-orderable point sequence
nonfinite residual values
```

Sorting by lambda is acceptable, but duplicate lambda after sorting must block.

Check whether `TestResidualRunRejectsNonfiniteOrDuplicateLambda` or equivalent covers this.

Mark as blocker if invalid point data can produce a computed residual.

## 7. Source branch registry review

Review `SourceBranchRegistry`.

Expected behavior:

```text
source_branch_registry is populated from parsed bmc0a_clock_readiness.json
every reported local_branch_eligibility branch ID is in the registry
every computed diagnostic branch ID is in the registry
every computed calculation ledger branch ID is in the registry
synthetic branch IDs are rejected unless explicitly fixture-only and not file-backed
```

Mark as blocker if a branch not present in the source artifact can be reported as file-backed.

Specifically verify the previous bad case is gone:

```text
unconditional synthetic branch_1 injection
```

## 8. Branch eligibility review

A branch may receive a computed residual only if:

```text
eligible = true
eligibility_status = eligible_local_branch
node_contact_free = true
trajectory_finite = true
derivative_readiness_status = ready
points exist and have length >= 2
```

Reported summary:

```text
Eligible Local Branches: 1
Computed Candidate Residual Diagnostics: 1
Blocked Candidate Residual Diagnostics: 0
Total Candidate Residual Diagnostics: 1
Residual/Null Comparisons: 1
```

Check that the one computed diagnostic references exactly the one eligible source-backed branch.

Mark as blocker if ineligible branches receive computed residuals.

## 9. Residual input point review

Review `ResidualInputPoint`.

Expected fields, or equivalents:

```text
branch_id
point_index
alpha
phi
candidate_left_hand_side
candidate_right_hand_side
input_provenance
```

For computed diagnostics:

```text
residual_input_points must be present and nonempty
each input point must be tied to the computed branch
each point must have finite values
candidate_left_hand_side and candidate_right_hand_side must be derived from finite differences, not constants
input_provenance must be file_read or derived_from_file_read
```

Mark as blocker if computed diagnostics have empty or decorative residual input points.

## 10. Candidate residual diagnostic consistency

Review each `CandidateResidualDiagnostic`.

Allowed residual statuses:

```text
candidate_residual_diagnostics_generated
residual_input_blocked
residual_nonfinite
blocked_by_clock_fragility
blocked_by_node_obstruction
blocked_by_convention_debt
source_unavailable
blocked_by_missing_residual_inputs
```

Allowed residual provenances:

```text
computed_from_bmc0a_local_branch
deterministic_fixture
source_artifact_summary
blocked
```

Forbidden provenances:

```text
assumed
fabricated
validated_physics
inferred_success
```

Validation must reject:

```text
residual_computed = true with blocked status
residual_computed = false with generated status
residual_computed = true with provenance = blocked
candidate_residual_computed = false while any diagnostic has residual_computed = true
candidate_residual_computed = true with no computed diagnostic
computed_from_bmc0a_local_branch without matching calculation ledger
computed diagnostic with empty residual_input_points
```

Mark as blocker if status, provenance, and metrics are internally inconsistent.

## 11. Calculation ledger review

Review `ResidualCalculationLedger`.

For computed entries, validation must enforce:

```text
calculation_id nonempty
branch_id nonempty
formula_source nonempty
formula_id nonempty
formula_description nonempty
convention_profile nonempty
input_fields nonempty
input_provenance = file_read or derived_from_file_read
num_input_points > 0
calculation_status = computed_from_local_branch
notes do not imply recovery, validation, or superiority
```

Formula ID should be guarded, for example:

```text
candidate_local_branch_velocity_constraint_residual_v0.1
```

Forbidden formula labels:

```text
friedmann_residual
classical_residual
recovery_residual
cosmology_recovery_residual
```

Mark as blocker if the ledger is decorative or hides numeric formula components.

## 12. Metric invariant review

Review `CandidateResidualMetrics`.

Required fields:

```text
num_evaluation_points
num_finite_residual_points
mean_abs_residual
max_abs_residual
rms_residual
residual_finite
diagnostic_warnings
```

Pointer-valued numeric metrics must be explicit JSON `null` when unavailable:

```text
mean_abs_residual
max_abs_residual
rms_residual
```

Validation must reject:

```text
NaN
+Inf
-Inf
negative residual magnitudes
sentinel values such as -1
missing optional metric keys
num_evaluation_points < 0
num_finite_residual_points < 0
num_finite_residual_points > num_evaluation_points
residual_finite = true when num_finite_residual_points = 0
computed residual diagnostics with nil mean_abs_residual
computed residual diagnostics with nil max_abs_residual
computed residual diagnostics with nil rms_residual
```

Mark as blocker if invalid metrics can validate.

## 13. Blocked report path review

Old artifacts without trajectory points must produce:

```text
candidate_residual_computed = false
interpretation_status = blocked_by_missing_residual_inputs
Computed Candidate Residual Diagnostics: 0
Residual/Null Comparisons: 0
no fake residual metrics
```

Confirm a test exists:

```text
TestResidualRunBlockedReportWhenInputsMissing
```

or a stronger equivalent.

Mark as blocker if missing point inputs still generate computed residuals.

## 14. Residual/null comparison review

Review `ResidualNullComparison`.

Allowed interpretation statuses:

```text
diagnostic_comparison_only
mixed_residual_diagnostics
insufficient_target_null_separation
target_null_residual_separation_candidate_unpromoted
blocked_by_no_comparable_null_diagnostics
blocked_by_no_eligible_local_branch
blocked_by_missing_residual_inputs
blocked_by_convention_debt
blocked_by_clock_fragility
blocked_by_node_obstruction
```

Forbidden:

```text
winner
BMC_beats_nulls
validated
confirmed
recovered
proved
outperformed
passed
failed
```

If `comparison_computed = true`, require:

```text
target_residual_ids nonempty
every target_residual_id exists
every referenced target residual has residual_computed = true
null_model_ids nonempty
metrics_compared nonempty
null inputs come from Sprint 9 file-backed diagnostics or are explicitly source summaries
reason does not imply superiority
interpretation_status remains unpromoted
```

If candidate residuals are blocked, comparisons must be absent or blocked.

Mark as blocker if comparison records are decorative or imply superiority.

## 15. Source artifact use review

Expected source artifact IDs exactly once:

```text
bmc0a_clock_readiness
bmc0a_friedmann_spec
bmc0a_nullrun
bmc0a_prior_art_boundary
```

If `candidate_residual_computed = true`, verify:

```text
bmc0a_clock_readiness is actually used for branch points
bmc0a_friedmann_spec is used at least as formula/convention source or explicitly labeled as context-only
bmc0a_nullrun is used for null comparison or explicitly labeled as comparison source
bmc0a_prior_art_boundary is used only as boundary/context source, not physics validation
```

Mark as repair-required if source roles are overstated.

Mark as blocker if file_read provenance is claimed for unread files.

## 16. Convention ledger review

Required convention debts exactly once:

```text
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
faithfulness_review_debt
```

Expected default statuses:

```text
clock_choice_debt = unpaid
classical_target_debt = unpaid
unit_convention_debt = unpaid
sign_convention_debt = unpaid
normalization_debt = unpaid
faithfulness_review_debt = contested
```

Validation must reject:

```text
missing convention debt
duplicate convention debt
unknown convention debt
empty description
human_review_required = false for unresolved debt
retired convention debt
partial for the five core convention debts
```

Mark as blocker if convention debts are hidden or treated as retired.

## 17. Local-only/global-claim boundary

The report must contain:

```text
local_branch_only = true
global_cosmology_claim = false
```

Validation must reject:

```text
local_branch_only = false
global_cosmology_claim = true
```

Also check summary/warnings for any global-cosmology implication.

Mark as blocker if local branch diagnostics are treated as full cosmology.

## 18. Gates

Expected gates exactly once:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
candidate_residual_diagnostics_gate
no_recovery_claim_gate
no_scientific_novelty_claim_gate
no_bmc_beats_null_models_claim_gate
full_bmc_blocked_gate
convention_debts_visible_gate
faithfulness_contested_gate
no_full_bmc_promotion_gate
```

All gates must have:

```text
status = pass
```

Validation must reject:

```text
missing gates
duplicated gates
unknown gates
empty gate names
non-pass gate statuses
```

Mark as blocker if gates are decorative rather than enforced.

## 19. Validator strictness

Review `ValidateReport` or equivalent.

It must reject:

```text
wrong schema_version
wrong artifact_kind
wrong scope
toy_analysis_only = false
final_truth_claim = true
candidate_residual_computed inconsistent with branch/residual records
residual_recovery_claim = true
scientific_novelty_claim_made = true
bmc_beats_null_models_claim = true
full_bmc_toy_gate != blocked
local_branch_only = false
global_cosmology_claim = true
empty local_branch_eligibility
missing required source artifacts
duplicate source artifacts
unknown source artifacts
file_read without path
file_read without available/read_success status
empty convention_ledger
missing required convention debts
duplicate convention debts
unknown convention debts
convention debt marked retired
human_review_required = false for unresolved convention debt
candidate residual computed for ineligible branch
candidate residual computed without source-backed branch points
candidate residual computed without calculation ledger
calculation ledger with empty formula fields
computed ledger with zero input points
residual_computed = true when branch is ineligible
candidate_residual_diagnostics empty when candidate_residual_computed = true
computed diagnostic with empty residual_input_points
nonfinite numeric residual metrics
negative residual magnitudes
sentinel residual values
missing optional residual metric keys
num_finite_residual_points > num_evaluation_points
residual_finite = true with zero finite points
forbidden residual_status
forbidden residual_provenance
forbidden interpretation_status
comparison computed with empty target_residual_ids
comparison computed with unknown target_residual_ids
comparison computed using uncomputed target residuals
comparison computed with empty null_model_ids
comparison computed with empty metrics_compared
winner/outperform/validated/recovered/proved language
missing required gates
duplicated gates
unknown gates
any gate status != pass
ebp_debt_vocabulary != ptw_adversarial_review_debt_status_v0.1
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
clock_choice_debt != unpaid
classical_target_debt != unpaid
unit_convention_debt != unpaid
sign_convention_debt != unpaid
normalization_debt != unpaid
warning missing “No recovery claim is made.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
```

Check strict JSON decoding:

```go
json.Decoder.DisallowUnknownFields()
```

Also check that trailing JSON tokens are rejected.

## 20. Forbidden phrase scan

Search source, generated JSON, summary output, tests, comments, warnings, and walkthrough for forbidden phrases or equivalent implications:

```text
BMC beats null models
BMC outperforms controls
null models passed
null models failed
winner
validated physics
confirmed recovery
Friedmann recovery
recovery of Friedmann
ready for recovery
full BMC validated
scientifically novel
breakthrough
problem of time solved
quantum gravity validated
assumed success
fabricated
inferred success
classical limit achieved
classical cosmology recovered
```

The phrase scanner should be case-insensitive and phrase-safe.

Error messages must not echo the forbidden phrase itself.

## 21. CLI behavior

Run or inspect:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json
./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
./ptw-bmc validate --report out/bmc0a_local_residual.json
./ptw-bmc summarize --report out/bmc0a_local_residual.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Check:

```text
segment-clock emits branch points if residual computation depends on them
run-residuals subcommand exists
unknown profile fails safely
validate routes bmc0a-local-residual-v0.1 to residualrun validator
summarize routes bmc0a-local-residual-v0.1 to residualrun summarizer
summary counts only truly eligible branches
summary distinguishes computed vs blocked diagnostics
existing schema routes are not weakened
errors are explicit but phrase-safe
```

## 22. Lean policy review

Review:

```text
BMC/BMC/ResidualRun.lean
BMC/BMC.lean
```

Acceptable Lean scope:

```text
tiny policy-only file
no physics proof
no residual correctness proof
no recovery proof
no BMC-beats-nulls proof
no classical-limit proof
no full BMC proof
supports both computed and blocked residual-run configurations
```

Check:

```text
lake build succeeds
no sorry/admit exists
theorems only encode safety booleans
```

Mark as blocker if Lean claims physics validity.

## 23. Tests

Confirm tests or equivalents exist for:

```text
TestResidualRunBlockedReportWhenInputsMissing
TestResidualRunMetricsChangeWhenInputBranchDataChanges
TestResidualRunRejectsSyntheticFileBackedBranch
TestResidualRunRejectsReportedBranchNotInSourceRegistry
TestResidualRunCalculationLedgerRequiresFormulaTransparency
TestResidualRunRejectsComputedLedgerWithZeroInputPoints
TestResidualRunRejectsComputedLedgerWithEmptyInputFields
TestResidualRunSummaryDistinguishesComputedAndBlockedDiagnostics
TestResidualRunRejectsComputedBranchProvenanceWithoutCalculationLedger
TestResidualRunRejectsReportFalseWithComputedDiagnostic
TestResidualRunRejectsReportTrueWithNoComputedDiagnostic
TestResidualRunRejectsNonfiniteOrDuplicateLambda
TestResidualRunRejectsComputedDiagnosticWithEmptyResidualInputPoints
TestResidualRunRejectsComputedResidualWithoutSourceBackedBranchPoints
TestResidualRunSegmentClockEmitsBranchPoints
```

Missing tests should be repair-required unless covered by stronger equivalents.

## EBP debt classification

Classify each item as one of:

```text
unpaid
partial
retired
contested
overclaimed
absent
```

Debt items:

```text
needLiteratureAudit
needMap
needInvariant
needToyCheck
needNullModel
needObstruction
needFaithfulnessReview
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
containsFinalTruthClaim
ResidualRunDiagnosticIntegrity
ResidualNullComparisonIntegrity
SourceProvenanceIntegrity
LocalOnlyBoundary
NoRecoveryClaimBoundary
NoBMCBeatsNullsBoundary
ConventionDebtVisibility
LeanPolicyBoundary
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
    "needLiteratureAudit": "",
    "needMap": "",
    "needInvariant": "",
    "needToyCheck": "",
    "needNullModel": "",
    "needObstruction": "",
    "needFaithfulnessReview": "",
    "clock_choice_debt": "",
    "classical_target_debt": "",
    "unit_convention_debt": "",
    "sign_convention_debt": "",
    "normalization_debt": "",
    "containsFinalTruthClaim": "",
    "ResidualRunDiagnosticIntegrity": "",
    "ResidualNullComparisonIntegrity": "",
    "SourceProvenanceIntegrity": "",
    "LocalOnlyBoundary": "",
    "NoRecoveryClaimBoundary": "",
    "NoBMCBeatsNullsBoundary": "",
    "ConventionDebtVisibility": "",
    "LeanPolicyBoundary": ""
  },
  "scope_findings": [],
  "clockseg_point_provenance_findings": [],
  "hidden_constant_findings": [],
  "sensitivity_test_findings": [],
  "finite_difference_findings": [],
  "source_branch_registry_findings": [],
  "branch_eligibility_findings": [],
  "residual_input_point_findings": [],
  "candidate_residual_findings": [],
  "calculation_ledger_findings": [],
  "metric_findings": [],
  "blocked_path_findings": [],
  "comparison_findings": [],
  "source_artifact_use_findings": [],
  "convention_ledger_findings": [],
  "local_boundary_findings": [],
  "gate_findings": [],
  "validator_findings": [],
  "cli_findings": [],
  "lean_findings": [],
  "overclaim_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "faithfulness_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|candidate_residual_runner_candidate_only|promoted_candidate_residual_runner_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 10.2 passes perfectly, the maximum allowed recommendation is:

```text
promoted_candidate_residual_runner_artifact_after_repairs
```

Never recommend promotion as:

```text
Friedmann recovery
classical-limit recovery
BMC beats null models
BMC outperforms controls
full BMC
scientific novelty
quantum gravity validated
problem of time solved
```

Remember:

```text
A candidate residual diagnostic is not recovery.
A residual/null comparison is not a victory claim.
A local branch is not full cosmology.
A file-backed calculation is still a toy diagnostic.
No per-point residual inputs means no computed residual report.
Full BMC remains blocked.
```
