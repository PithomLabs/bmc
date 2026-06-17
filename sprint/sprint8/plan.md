I drafted Sprint 8 as a **prior-art / claim-overlap audit sprint**, not a physics-result sprint. The reason is strong: Bohmian quantum cosmology, minisuperspace Wheeler–DeWitt models, scalar-field treatments, Bohmian trajectories, singularity/classical-limit discussions, and Bohmian quantum-gravity/cosmology reviews already exist in the literature, so Sprint 8 should prevent accidental novelty overclaims before we compute anything further. ([arXiv][1])

# BMC Sprint 8 Implementation Plan: Literature / Prior-Art Audit Profile

## Sprint name

```text
BMC Sprint 8: Literature / Prior-Art Audit Profile
```

## Sprint purpose

Sprint 8 builds a strict prior-art audit layer for BMC-0A and the surrounding PTW/EBP workflow.

The sprint does not compute physics results.

The sprint does not run null-model comparisons.

The sprint does not compute a Friedmann residual.

The sprint does not claim novelty.

The sprint does not promote full BMC.

Its purpose is to answer:

```text
Which parts of BMC-0A are already established prior art?
Which parts are likely prior art but require source review?
Which parts are implementation variants?
Which parts are workflow-distinctive candidates?
Which claims must be downgraded, blocked, or tied to required sources before future promotion?
```

## Current accepted context

Accepted prior BMC artifacts:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
Sprint 6: BMC-0A Friedmann-residual specification/gate-design artifact
Sprint 7: BMC-0A null-model scaffold artifact
```

Current hard boundaries:

```text
residual_computed: false
null_comparison_computed: false
recovery_claim: false
full_bmc_toy_gate: blocked
needNullModel: active
needLiteratureAudit: active
needFaithfulnessReview: contested
clock_choice_debt: active
classical_target_debt: active
unit_convention_debt: active
sign_convention_debt: active
normalization_debt: active
```

## Maximum allowed promotion

The maximum allowed Sprint 8 promotion is:

```text
promoted_prior_art_audit_artifact_after_repairs
```

Forbidden promotion labels:

```text
novel_physics_claim
first_ever_claim
scientifically_original_claim
friedmann_recovery
ready_for_friedmann_recovery
full_bmc
full_quantum_gravity
bohmian_cosmology_proven
problem_of_time_solved
```

## Core design principle

Sprint 8 must separate:

```text
prior_art_status
implementation_status
workflow_distinctiveness
physics_claim_status
promotion_status
```

A claim may be executable in PTW while still being old physics.

A claim may be workflow-distinctive while not being scientifically novel.

A claim may be cited in prior art while still needing faithful reimplementation review.

A claim may be interesting while still blocked.

## Proposed Go package

Create:

```text
internal/bmc/litaudit/
```

Suggested files:

```text
contracts.go
papers.go
claims.go
overlap.go
novelty.go
sources.go
gates.go
report.go
validate.go
litaudit_test.go
```

## CLI plan

Add:

```bash
ptw-bmc audit-literature --profile bmc0a-prior-art --out out/bmc0a_prior_art_audit.json
ptw-bmc validate --report out/bmc0a_prior_art_audit.json
ptw-bmc summarize --report out/bmc0a_prior_art_audit.json
```

Schema version:

```text
bmc0a-prior-art-audit-v0.1
```

Unknown profiles must fail safely.

Validation and summarization must route by schema version.

## Data model: source papers

Define:

```go
type PriorArtSource struct {
    SourceID string `json:"source_id"`
    Title string `json:"title"`
    Authors []string `json:"authors"`
    Year int `json:"year"`
    SourceKind string `json:"source_kind"`
    VenueOrArchive string `json:"venue_or_archive"`
    DOIOrArxiv string `json:"doi_or_arxiv"`
    ReviewStatus string `json:"review_status"`
    RelevanceTags []string `json:"relevance_tags"`
    Notes string `json:"notes"`
}
```

Allowed `SourceKind` values:

```text
review
paper
book
thesis
lecture_notes
unknown
```

Allowed `ReviewStatus` values:

```text
seed_unreviewed
abstract_reviewed
skim_reviewed
full_text_reviewed
human_review_required
excluded
```

Forbidden `ReviewStatus` values:

```text
proves_our_claim
confirms_novelty
settled
definitive
```

Sprint 8 may use `seed_unreviewed`, `abstract_reviewed`, or `human_review_required`.

Do not pretend a paper is fully reviewed unless the implementation actually records evidence for that status.

## Data model: BMC claims

Define:

```go
type BMCClaimForAudit struct {
    ClaimID string `json:"claim_id"`
    ClaimText string `json:"claim_text"`
    ClaimKind string `json:"claim_kind"`
    SourceArtifact string `json:"source_artifact"`
    BMCStatus string `json:"bmc_status"`
    RequiresLiteratureAudit bool `json:"requires_literature_audit"`
    RequiresHumanReview bool `json:"requires_human_review"`
    Notes string `json:"notes"`
}
```

Required BMC claims to register:

```text
bmc_uses_wdw_minisuperspace_equation
bmc_uses_bohmian_guidance_from_phase
bmc_uses_quantum_potential_diagnostic
bmc_detects_nodes_as_obstructions
bmc_tests_scalar_field_clock_monotonicity
bmc_segments_local_relational_branches
bmc_defines_candidate_friedmann_residual_spec
bmc_defines_null_model_scaffold
bmc_uses_ebp_gates_and_debt_ledger
bmc_uses_lean_policy_safety_contracts
```

Allowed `ClaimKind` values:

```text
equation_structure_claim
guidance_claim
diagnostic_claim
clock_choice_claim
branch_structure_claim
residual_spec_claim
null_model_scaffold_claim
workflow_governance_claim
formal_policy_claim
interpretive_claim
```

Allowed `BMCStatus` values:

```text
accepted_scaffold_only
accepted_control_only
accepted_diagnostic_only
blocked
contested
not_claimed
```

Forbidden `BMCStatus` values:

```text
proven_true
novel
first_ever
physics_validated
friedmann_recovered
full_bmc_promoted
```

## Data model: claim-overlap records

Define:

```go
type ClaimOverlapRecord struct {
    ClaimID string `json:"claim_id"`
    RelatedSourceIDs []string `json:"related_source_ids"`
    OverlapStatus string `json:"overlap_status"`
    OverlapReason string `json:"overlap_reason"`
    RequiredCitationBeforePromotion bool `json:"required_citation_before_promotion"`
    HumanReviewRequired bool `json:"human_review_required"`
    Notes string `json:"notes"`
}
```

Allowed `OverlapStatus` values:

```text
established_prior_art
likely_prior_art
partial_overlap
implementation_variant
workflow_distinctive_candidate
unknown_requires_review
not_claimed
blocked
```

Forbidden `OverlapStatus` values:

```text
novel
proven_new
first_ever
scientifically_original
breakthrough
unprecedented
settled
```

Interpretation rules:

```text
established_prior_art:
  Similar claim or method appears clearly in prior literature.

likely_prior_art:
  Abstract/summary suggests overlap, but full source review is still required.

partial_overlap:
  Some ingredients overlap, but exact BMC/PTW implementation differs.

implementation_variant:
  Physics idea is known; BMC implementation may be a software variant.

workflow_distinctive_candidate:
  Possible distinction lies in EBP/PTW workflow, not necessarily physics.

unknown_requires_review:
  No safe classification yet.

not_claimed:
  BMC does not make this claim.

blocked:
  Claim must not be promoted.
```

## Data model: novelty boundary

Define:

```go
type NoveltyBoundary struct {
    NoveltyClaimMade bool `json:"novelty_claim_made"`
    ScientificNoveltyClaimAllowed bool `json:"scientific_novelty_claim_allowed"`
    WorkflowDistinctivenessClaimAllowed bool `json:"workflow_distinctiveness_claim_allowed"`
    Reason string `json:"reason"`
}
```

Required values for Sprint 8:

```text
novelty_claim_made = false
scientific_novelty_claim_allowed = false
workflow_distinctiveness_claim_allowed = true only as candidate
```

Acceptable reason:

```text
Sprint 8 maps prior-art overlap only. It does not establish scientific novelty. Workflow distinctiveness remains a candidate requiring human review.
```

## Data model: required source ledger

Define:

```go
type RequiredSourceLedgerEntry struct {
    ClaimID string `json:"claim_id"`
    RequiredSourceIDs []string `json:"required_source_ids"`
    MissingSourceDebt bool `json:"missing_source_debt"`
    HumanReviewRequired bool `json:"human_review_required"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Allowed statuses:

```text
source_declared
source_missing
source_review_required
blocked_until_source_review
```

Forbidden statuses:

```text
source_proves_claim
source_confirms_novelty
source_settles_claim
```

## Report shape

Generate deterministic JSON:

```json
{
  "schema_version": "bmc0a-prior-art-audit-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "spec_kind": "prior_art_audit",
  "spec_scope": "literature_overlap_mapping_only",
  "residual_computed": false,
  "null_comparison_computed": false,
  "recovery_claim": false,
  "novelty_claim_made": false,
  "scientific_novelty_claim_allowed": false,
  "workflow_distinctiveness_claim_allowed": "candidate_only",
  "source_registry": [],
  "bmc_claims": [],
  "claim_overlap_records": [],
  "required_source_ledger": [],
  "novelty_boundary": {},
  "gates": [],
  "promotion_gate": {
    "name": "prior_art_audit_gate",
    "status": "candidate_only",
    "reason": "Sprint 8 maps literature overlap only. It does not establish scientific novelty."
  },
  "ebp_debt": {
    "needLiteratureAudit": "active",
    "needMap": "active",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "active",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "classical_target_debt": "active",
    "unit_convention_debt": "active",
    "sign_convention_debt": "active",
    "normalization_debt": "active",
    "containsFinalTruthClaim": "absent",
    "LeanVerification": "planned_for_policy_safety_contracts",
    "promotion_status": "prior_art_audit_candidate_only"
  },
  "warnings": [
    "Sprint 8 is a prior-art overlap audit only.",
    "No scientific novelty claim is made.",
    "No recovery claim is made.",
    "No residual was computed.",
    "No null-model comparison was computed.",
    "Full BMC remains blocked."
  ]
}
```

## Required gates

Define:

```go
type PriorArtAuditGate struct {
    Name string `json:"name"`
    Status string `json:"status"`
    Reason string `json:"reason"`
}
```

Required gates, exactly once each:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
no_scientific_novelty_claim_gate
prior_art_registry_nonempty_gate
bmc_claim_registry_nonempty_gate
claim_overlap_records_nonempty_gate
required_source_ledger_gate
no_residual_computation_gate
no_null_comparison_result_gate
no_recovery_claim_gate
full_bmc_blocked_gate
faithfulness_contested_gate
```

Allowed gate statuses:

```text
pass
blocked
contested
```

For Sprint 8, all required gates must be `pass`, except gates that explicitly represent contested faithfulness may use `pass` only if the report records `needFaithfulnessReview = contested`.

Validation must reject missing, duplicated, unknown, empty, blocked, or contested gates.

## Validation requirements

`ValidatePriorArtAuditReport` must reject:

```text
wrong schema_version
wrong spec_kind
wrong spec_scope
toy_analysis_only = false
final_truth_claim = true
residual_computed = true
null_comparison_computed = true
recovery_claim = true
novelty_claim_made = true
scientific_novelty_claim_allowed = true
workflow_distinctiveness_claim_allowed not equal candidate_only
empty source_registry
empty bmc_claims
empty claim_overlap_records
empty required_source_ledger
missing required BMC claims
duplicate claim_id
duplicate source_id
duplicate overlap record for same claim_id and source set
overlap status outside allowed list
overlap status = novel|first_ever|breakthrough|scientifically_original
claim status = proven_true|novel|first_ever|physics_validated|friedmann_recovered|full_bmc_promoted
source review status outside allowed list
source review status = proves_our_claim|confirms_novelty|settled|definitive
novelty_boundary.novelty_claim_made = true
novelty_boundary.scientific_novelty_claim_allowed = true
missing required gates
duplicated gates
unknown gates
any required gate status != pass
promotion_gate.status not equal candidate_only
needLiteratureAudit missing
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
full BMC not blocked
warning missing “No scientific novelty claim is made.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
nonfinite numeric values
```

Use:

```go
json.Decoder.DisallowUnknownFields()
```

Also check EOF after decode to reject trailing tokens.

## Forbidden phrase scanning

Scan generated JSON, summary output, report text, warnings, comments, and tests for:

```text
first ever
scientifically novel
scientifically original
breakthrough
unprecedented
proves our model
confirms our model
BMC is new physics
BMC proves Bohmian cosmology
BMC validates quantum gravity
Friedmann recovery
recovery of Friedmann
ready for recovery
full BMC validated
problem of time solved
```

Allowed neutral phrases:

```text
prior-art audit
overlap mapping
workflow-distinctive candidate
scientific novelty not established
source review required
human review required
candidate only
blocked
contested
```

## Deterministic output requirements

Repeated report generation must be byte-identical.

Ensure deterministic ordering for:

```text
source_registry
bmc_claims
claim_overlap_records
required_source_ledger
gates
warnings
```

No map iteration should determine JSON output order.

## Summary output

`ptw-bmc summarize --report out/bmc0a_prior_art_audit.json` should display:

```text
BMC Sprint 8 Prior-Art Audit Summary
Schema Version: bmc0a-prior-art-audit-v0.1
Spec Kind: prior_art_audit
Spec Scope: literature_overlap_mapping_only
Scientific Novelty Claim Made: false
Scientific Novelty Claim Allowed: false
Workflow Distinctiveness: candidate_only
Residual Computed: false
Null Comparison Computed: false
Recovery Claim: false
Sources Registered: N
BMC Claims Registered: N
Claim Overlap Records: N
Full BMC: blocked
Promotion Status: prior_art_audit_candidate_only
```

Do not print “novel result,” “scientific novelty established,” or similar wording.

## Lean policy plan

Add:

```text
BMC/BMC/PriorArtAudit.lean
```

Import it in:

```text
BMC/BMC.lean
```

Define a policy-only structure:

```lean
structure BMCPriorArtAuditReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  residualComputed : Bool
  nullComparisonComputed : Bool
  recoveryClaim : Bool
  noveltyClaimMade : Bool
  scientificNoveltyAllowed : Bool
  fullBMCBlocked : Bool
  literatureAuditActive : Bool
  faithfulnessContested : Bool
```

Policy theorems:

```text
prior_art_audit_requires_toy_only
prior_art_audit_blocks_final_truth
prior_art_audit_forbids_residual_computation
prior_art_audit_forbids_null_comparison_result
prior_art_audit_forbids_recovery_claim
prior_art_audit_forbids_novelty_claim
prior_art_audit_forbids_scientific_novelty_allowed
prior_art_audit_requires_full_bmc_blocked
prior_art_audit_requires_literature_audit_active
prior_art_audit_requires_faithfulness_contested
prior_art_audit_does_not_imply_original_physics
prior_art_audit_does_not_imply_full_bmc
```

Lean must remain policy/safety only.

Do not prove scientific novelty.

Do not prove prior-art completeness.

Do not prove BMC physics.

## Test plan

Add tests:

```text
TestPriorArtAuditReportValidation
TestPriorArtAuditRejectsNoveltyClaimMade
TestPriorArtAuditRejectsScientificNoveltyAllowed
TestPriorArtAuditRejectsResidualComputed
TestPriorArtAuditRejectsNullComparisonComputed
TestPriorArtAuditRejectsRecoveryClaim
TestPriorArtAuditRequiresSourceRegistry
TestPriorArtAuditRequiresBMCClaimRegistry
TestPriorArtAuditRequiresClaimOverlapRecords
TestPriorArtAuditRequiresRequiredSourceLedger
TestPriorArtAuditRejectsMissingRequiredBMCClaims
TestPriorArtAuditRejectsDuplicateClaimIDs
TestPriorArtAuditRejectsDuplicateSourceIDs
TestPriorArtAuditRejectsForbiddenOverlapStatuses
TestPriorArtAuditRejectsForbiddenClaimStatuses
TestPriorArtAuditRejectsForbiddenSourceReviewStatuses
TestPriorArtAuditRequiresAllGatesExactlyOnce
TestPriorArtAuditRejectsNonPassGate
TestPriorArtAuditRejectsUnknownFields
TestPriorArtAuditRejectsTrailingJSONTokens
TestPriorArtAuditDeterministicJSON
TestPriorArtAuditForbiddenPhraseScan
TestPriorArtAuditCLIRouting
TestPriorArtAuditUnknownProfileFails
```

## Verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc audit-literature --profile bmc0a-prior-art --out out/bmc0a_prior_art_audit.json
./ptw-bmc validate --report out/bmc0a_prior_art_audit.json
./ptw-bmc summarize --report out/bmc0a_prior_art_audit.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

## Acceptance criteria

Sprint 8 may be accepted only if:

```text
Go tests pass
CLI build succeeds
audit-literature generates deterministic JSON
validate passes for valid report
validate rejects forbidden reports
summarize uses neutral wording
Lean policy build succeeds
no novelty claim appears
no recovery claim appears
no residual is computed
no null-model comparison is computed
full BMC remains blocked
needLiteratureAudit remains active or partially scaffolded
```

## Required output after implementation

Return a walkthrough with:

```text
files added
files modified
schema version
sample summary output
test count
verification command results
source registry count
BMC claim count
claim-overlap record count
required source ledger count
gate list and statuses
EBP debt status
remaining limitations
```

## Expected post-sprint status

If implementation, tests, CLI verification, Lean build, adversarial review, and repairs pass:

```text
promoted_prior_art_audit_artifact_after_repairs
```

Until then:

```text
prior_art_audit_candidate_only
```

## Explicit non-goals

Do not implement:

```text
full literature review completeness proof
automated paper PDF ingestion
LLM-based paper extraction
claim-to-profile mapper
null-model runner
Friedmann residual runner
physics novelty claim
Friedmann recovery claim
full BMC promotion
quantum gravity validation
problem-of-time solution
```

Those belong to later sprints.

Initial seed literature for Sprint 8 should include review and representative papers on Bohmian quantum cosmology rather than pretending the topic starts with us: Pinto-Neto/Fabris review de Broglie–Bohm quantum cosmology, Pinto-Neto/Struyve review Bohmian quantum gravity and cosmology, Colistete/Fabris/Pinto-Neto cover singularities and classical limits with scalar fields, and later papers provide analytical or FRW-scalar-field Bohmian minisuperspace examples. ([arXiv][1])

EBP status:

```text
Claim: Sprint 8 should be prior-art overlap mapping only.
Status: implementation plan candidate.

Claim: BMC physics is novel.
Status: not established / forbidden.

Claim: PTW/EBP workflow may be distinctive.
Status: candidate only; requires audit and human review.

needLiteratureAudit: active
needMap: active
needInvariant: partial
needToyCheck: active
needNullModel: active
needFaithfulnessReview: contested
full_bmc_toy_gate: blocked
```

[1]: https://arxiv.org/abs/1306.0820?utm_source=chatgpt.com "Quantum cosmology from the de Broglie-Bohm perspective"
