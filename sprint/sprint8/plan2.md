Go with the **slim Sprint 8**, not the big literature-audit subsystem.

The reason: prior art definitely exists for Bohmian quantum cosmology, minisuperspace Wheeler–DeWitt models, scalar-field cases, Bohmian trajectories, and classical-limit/singularity discussions, so we need a boundary note before future claims. But we do **not** need a full paper-ingestion/literature database yet. Pinto-Neto and Fabris review de Broglie–Bohm quantum cosmology in minisuperspace and beyond; Pinto-Neto and Struyve review Bohmian quantum gravity/cosmology; Colistete/Fabris/Pinto-Neto already studied scalar-field minisuperspace, WKB superpositions, Bohmian trajectories, singularities, and classical-limit behavior. ([arXiv][1])

Use this as the next implementation prompt.

# BMC Sprint 8-Lite Implementation Plan: BMC-0A Prior-Art Boundary Note

## Sprint name

```text id="07852g"
BMC Sprint 8-Lite: BMC-0A Prior-Art Boundary Note
```

## Purpose

Implement a small anti-bloat prior-art boundary artifact.

This sprint exists only to prevent accidental novelty inflation before we continue toward the BMC null-model runner and future local-branch residual experiments.

Do **not** build a full literature-audit subsystem.

Do **not** build paper ingestion.

Do **not** build claim extraction.

Do **not** build a general physics-branch audit engine.

Do **not** compute null-model comparisons.

Do **not** compute Friedmann residuals.

Do **not** claim scientific novelty.

## Core question

```text id="lk8we4"
Which parts of BMC-0A are clearly prior-art territory, which parts are implementation/workflow-specific, and which claims remain blocked from novelty or physics promotion?
```

## Scope

This sprint should create one small deterministic prior-art boundary report and validation layer.

Keep it BMC-focused.

Do not generalize to all physics yet.

## Files to add

Prefer this minimal package:

```text id="th66vy"
internal/bmc/priorart/contracts.go
internal/bmc/priorart/report.go
internal/bmc/priorart/validate.go
internal/bmc/priorart/priorart_test.go
```

Optional documentation artifact:

```text id="r3hox5"
docs/bmc0a_prior_art_boundary.md
```

Do not create a full `litaudit` framework yet.

## Files to modify

```text id="mjplke"
cmd/ptw-bmc/main.go
```

Optional, only if current validation/summarization routing requires it:

```text id="gi57q3"
internal/bmc/report/*
```

## CLI

Add one narrow command:

```bash id="rdfgjm"
ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
```

Route validation and summary by schema version:

```bash id="293h79"
ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
```

Schema version:

```text id="6nk4uf"
bmc0a-prior-art-boundary-v0.1
```

Unknown profile must fail safely.

## Report shape

Generate deterministic JSON:

```json id="bfa7uv"
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

## Prior-art source structure

Define:

```go id="6olq6u"
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

```text id="s05rmb"
review
paper
book
unknown
```

Allowed `ReviewStatus`:

```text id="q7uxk3"
seed_unreviewed
abstract_reviewed
skim_reviewed
human_review_required
```

Forbidden:

```text id="qceqgq"
proves_our_claim
confirms_novelty
settled
definitive
```

Seed at least these source classes:

```text id="up6cri"
bohmian_quantum_cosmology_review
bohmian_quantum_gravity_cosmology_review
scalar_field_minisuperspace_classical_limit_paper
gaussian_superposition_bohmian_trajectory_paper
modern_friedmann_scalar_field_bohmian_example
```

The report may use source IDs and metadata, but it must not pretend full literature completeness.

## Boundary claim structure

Define:

```go id="sih6ul"
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

```text id="ep4cma"
established_prior_art
likely_prior_art
implementation_variant
workflow_distinctive_candidate
unknown_requires_review
blocked
not_claimed
```

Forbidden:

```text id="y7gq78"
novel
first_ever
scientifically_original
breakthrough
proved_new
validated_physics
```

Required boundary claims:

```text id="z5svou"
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

Expected classifications:

```text id="m5h25i"
bohmian minisuperspace trajectories: established_prior_art or likely_prior_art
Wheeler-DeWitt toy equation: established_prior_art or likely_prior_art
scalar field clock checks: likely_prior_art / unknown_requires_review
quantum potential diagnostics: established_prior_art or likely_prior_art
node obstruction detection: implementation_variant or likely_prior_art
local branch segmentation: implementation_variant / unknown_requires_review
null-model scaffold: workflow_distinctive_candidate
EBP debt gates: workflow_distinctive_candidate
Lean policy contracts: workflow_distinctive_candidate
recovery claim: not_claimed
scientific novelty claim: blocked
```

## Gates

Define exactly these gates:

```text id="xv000k"
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

All gates must appear exactly once.

All gates must have status:

```text id="s3ibkt"
pass
```

## Validation requirements

Validation must reject:

```text id="x8j6jp"
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

```go id="9lz8zo"
json.Decoder.DisallowUnknownFields()
```

Also check EOF after decoding to reject trailing JSON tokens.

## Forbidden phrase scan

Scan generated JSON and summary output for:

```text id="04t1h6"
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

Allowed neutral phrases:

```text id="2nnzts"
prior-art boundary note
scientific novelty not established
workflow-distinctive candidate
source review required
human review required
candidate only
blocked
contested
```

## Deterministic output

Repeated generation must be byte-identical.

No map iteration should determine output ordering.

## Summary output

The summary should print:

```text id="95ho2s"
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

Do not print novelty or breakthrough language.

## Tests

Add tests:

```text id="r99m3t"
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

```bash id="y7q8vt"
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
./ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
./ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

No new Lean file is required for Sprint 8-Lite unless the implementation already has a clean place to add a tiny policy theorem.

Avoid Lean expansion unless it is very small.

## Acceptance criteria

Sprint 8-Lite may be accepted only if:

```text id="q935oa"
Go tests pass
CLI build succeeds
report generation works
validation passes for valid report
validation rejects novelty/recovery/residual/comparison overclaims
summary uses neutral wording
deterministic JSON holds
no scientific novelty claim appears
full BMC remains blocked
```

## Explicit non-goals

Do not implement:

```text id="azv7pi"
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

## Expected status after implementation

Before adversarial review:

```text id="t8x8y4"
prior_art_boundary_note_candidate_only
```

Maximum promotion after implementation, review, and repairs:

```text id="blyq41"
promoted_prior_art_boundary_note_after_repairs
```

After Sprint 8-Lite, return immediately to the BMC physics pipeline:

```text id="zz0j9v"
Sprint 9: BMC Null-Model Runner
Sprint 10: Candidate Local-Branch Residual Runner
Sprint 11: Paper-Claim Extraction Prototype
```

EBP status:

```text id="et4t5f"
Claim: Slim Sprint 8 is safer than full literature-audit subsystem.
Status: recommended.

Claim: BMC physics novelty is established.
Status: blocked / not claimed.

Claim: Workflow distinctiveness may be candidate-only.
Status: allowed but human review required.

needLiteratureAudit: partially addressed by Sprint 8-Lite
needNullModel: active
needFaithfulnessReview: contested
full_bmc_toy_gate: blocked
```

[1]: https://arxiv.org/abs/1306.0820?utm_source=chatgpt.com "Quantum cosmology from the de Broglie-Bohm perspective"
