You are implementing **BMC Sprint 8-Lite: BMC-0A Prior-Art Boundary Note** under strict EBP 2.1 discipline.

This sprint must stay small.

Do **not** build a full literature-audit system.

Do **not** build paper ingestion.

Do **not** build claim extraction.

Do **not** build a general physics-branch workbench.

Do **not** compute null-model comparisons.

Do **not** compute Friedmann residuals.

Do **not** claim scientific novelty.

## Goal

Implement a small deterministic prior-art boundary artifact that prevents accidental novelty inflation before the project continues to the BMC null-model runner.

The report should answer:

```text
Which BMC-0A ingredients are prior-art territory?
Which are implementation/workflow-specific?
Which novelty or recovery claims remain blocked?
```

## Package to add

Create:

```text
internal/bmc/priorart/
```

Files:

```text
contracts.go
report.go
validate.go
priorart_test.go
```

Avoid a large `litaudit` framework.

## CLI

Add this command:

```bash
ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
```

Route validation and summary by schema version:

```text
bmc0a-prior-art-boundary-v0.1
```

Required commands must work:

```bash
ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
```

Unknown profile must fail safely.

## Report schema

Generate deterministic JSON with this identity:

```json
{
  "schema_version": "bmc0a-prior-art-boundary-v0.1",
  "toy_analysis_only": true,
  "final_truth_claim": false,
  "artifact_kind": "prior_art_boundary_note",
  "scope": "bmc0a_only",
  "scientific_novelty_claim_made": false,
  "scientific_novelty_claim_allowed": false,
  "workflow_distinctiveness_status": "candidate_only",
  "residual_computed": false,
  "null_comparison_computed": false,
  "recovery_claim": false,
  "full_bmc_toy_gate": "blocked",
  "prior_art_sources": [],
  "boundary_claims": [],
  "gates": [],
  "ebp_debt": {
    "needLiteratureAudit": "partial",
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
    "promotion_status": "prior_art_boundary_note_candidate_only"
  },
  "warnings": [
    "This is a BMC-0A prior-art boundary note only.",
    "No scientific novelty claim is made.",
    "No recovery claim is made.",
    "No residual was computed.",
    "No null-model comparison was computed.",
    "Full BMC remains blocked."
  ]
}
```

## Types

Define:

```go
type PriorArtSource struct {
    SourceID string `json:"source_id"`
    Title string `json:"title"`
    Authors []string `json:"authors"`
    Year int `json:"year"`
    SourceKind string `json:"source_kind"`
    ReviewStatus string `json:"review_status"`
    RelevanceTags []string `json:"relevance_tags"`
    BoundaryUse string `json:"boundary_use"`
}
```

Allowed `SourceKind`:

```text
review
paper
book
unknown
```

Allowed `ReviewStatus`:

```text
seed_unreviewed
abstract_reviewed
skim_reviewed
human_review_required
```

Forbidden `ReviewStatus`:

```text
proves_our_claim
confirms_novelty
settled
definitive
```

Define:

```go
type BoundaryClaim struct {
    ClaimID string `json:"claim_id"`
    ClaimText string `json:"claim_text"`
    BoundaryStatus string `json:"boundary_status"`
    RelatedSourceIDs []string `json:"related_source_ids"`
    Reason string `json:"reason"`
    HumanReviewRequired bool `json:"human_review_required"`
}
```

Allowed `BoundaryStatus`:

```text
established_prior_art
likely_prior_art
implementation_variant
workflow_distinctive_candidate
unknown_requires_review
blocked
not_claimed
```

Forbidden `BoundaryStatus`:

```text
novel
first_ever
scientifically_original
breakthrough
proved_new
validated_physics
```

## Required seed source classes

Seed at least these source IDs or equivalent:

```text
bohmian_quantum_cosmology_review
bohmian_quantum_gravity_cosmology_review
scalar_field_minisuperspace_classical_limit_paper
gaussian_superposition_bohmian_trajectory_paper
modern_friedmann_scalar_field_bohmian_example
```

Do not claim literature completeness.

## Required boundary claims

Register exactly these boundary claims or a clearly equivalent set:

```text
bmc_uses_bohmian_minisuperspace_trajectories
bmc_uses_wheeler_dewitt_toy_equation
bmc_uses_scalar_field_clock_checks
bmc_uses_quantum_potential_diagnostics
bmc_uses_node_obstruction_detection
bmc_uses_local_branch_segmentation
bmc_defines_null_model_scaffold
bmc_uses_ebp_debt_gates
bmc_uses_lean_policy_contracts
bmc_does_not_claim_recovery
bmc_does_not_claim_scientific_novelty
```

Expected status direction:

```text
Bohmian minisuperspace trajectories: established_prior_art or likely_prior_art
Wheeler-DeWitt toy equation: established_prior_art or likely_prior_art
scalar field clock checks: likely_prior_art or unknown_requires_review
quantum potential diagnostics: established_prior_art or likely_prior_art
node obstruction detection: implementation_variant or likely_prior_art
local branch segmentation: implementation_variant or unknown_requires_review
null-model scaffold: workflow_distinctive_candidate
EBP debt gates: workflow_distinctive_candidate
Lean policy contracts: workflow_distinctive_candidate
recovery claim: not_claimed
scientific novelty claim: blocked
```

## Gates

Required gates, exactly once each:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
no_scientific_novelty_claim_gate
prior_art_sources_seeded_gate
boundary_claims_declared_gate
no_residual_computation_gate
no_null_comparison_result_gate
no_recovery_claim_gate
full_bmc_blocked_gate
faithfulness_contested_gate
```

All gates must have:

```text
status = pass
```

## Validation

Implement strict validation.

Reject:

```text
wrong schema_version
wrong artifact_kind
wrong scope
toy_analysis_only = false
final_truth_claim = true
scientific_novelty_claim_made = true
scientific_novelty_claim_allowed = true
workflow_distinctiveness_status != candidate_only
residual_computed = true
null_comparison_computed = true
recovery_claim = true
full_bmc_toy_gate != blocked
empty prior_art_sources
empty boundary_claims
missing required boundary claims
duplicate source_id
duplicate claim_id
forbidden source review status
forbidden boundary status
boundary_status = novel
boundary_status = first_ever
boundary_status = scientifically_original
boundary_status = breakthrough
missing required gates
duplicated gates
unknown gates
any gate status != pass
containsFinalTruthClaim != absent
needFaithfulnessReview != contested
warning missing “No scientific novelty claim is made.”
warning missing “Full BMC remains blocked.”
unknown JSON fields
trailing JSON tokens
```

Use:

```go
json.Decoder.DisallowUnknownFields()
```

Also reject trailing JSON tokens after decode.

## Forbidden phrase scan

Generated JSON and summary output must not contain:

```text
first ever
scientifically novel
scientifically original
breakthrough
unprecedented
BMC is new physics
BMC proves Bohmian cosmology
BMC validates quantum gravity
Friedmann recovery
recovery of Friedmann
ready for recovery
full BMC validated
problem of time solved
```

Allowed neutral wording:

```text
prior-art boundary note
scientific novelty not established
workflow-distinctive candidate
source review required
human review required
candidate only
blocked
contested
```

## Summary output

`ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json` should print:

```text
BMC Sprint 8-Lite Prior-Art Boundary Summary
Schema Version: bmc0a-prior-art-boundary-v0.1
Scope: bmc0a_only
Scientific Novelty Claim Made: false
Scientific Novelty Claim Allowed: false
Workflow Distinctiveness: candidate_only
Residual Computed: false
Null Comparison Computed: false
Recovery Claim: false
Prior-Art Sources Seeded: N
Boundary Claims Declared: N
Full BMC: blocked
Promotion Status: prior_art_boundary_note_candidate_only
```

## Tests

Add tests:

```text
TestPriorArtBoundaryReportValidation
TestPriorArtBoundaryRejectsNoveltyClaimMade
TestPriorArtBoundaryRejectsScientificNoveltyAllowed
TestPriorArtBoundaryRejectsResidualComputed
TestPriorArtBoundaryRejectsNullComparisonComputed
TestPriorArtBoundaryRejectsRecoveryClaim
TestPriorArtBoundaryRequiresSeedSources
TestPriorArtBoundaryRequiresBoundaryClaims
TestPriorArtBoundaryRejectsMissingRequiredBoundaryClaims
TestPriorArtBoundaryRejectsDuplicateSourceIDs
TestPriorArtBoundaryRejectsDuplicateClaimIDs
TestPriorArtBoundaryRejectsForbiddenBoundaryStatuses
TestPriorArtBoundaryRequiresAllGatesExactlyOnce
TestPriorArtBoundaryRejectsNonPassGate
TestPriorArtBoundaryRejectsUnknownFields
TestPriorArtBoundaryRejectsTrailingJSONTokens
TestPriorArtBoundaryDeterministicJSON
TestPriorArtBoundaryForbiddenPhraseScan
TestPriorArtBoundaryCLIRouting
TestPriorArtBoundaryUnknownProfileFails
```

## Verification commands

Run:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
./ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
./ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

No new Lean file is required unless the change is tiny and policy-only.

## Return walkthrough

Return:

```text
files added
files modified
schema version
sample summary output
test count
verification command results
prior-art source count
boundary claim count
gate list and statuses
EBP debt status
remaining limitations
```

## Expected status

Before adversarial review:

```text
prior_art_boundary_note_candidate_only
```

Maximum after adversarial review and repairs:

```text
promoted_prior_art_boundary_note_after_repairs
```

## Explicit non-goals

Do not implement:

```text
full literature-review engine
paper PDF ingestion
LLM claim extraction
claim-to-profile mapper
general physics-branch audit registry
large citation database
null-model runner
Friedmann residual runner
full BMC promotion
scientific novelty claim
```
