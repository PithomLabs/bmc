Implement **BMC Sprint 8.1: Prior-Art Boundary Repair**.

Do not start Sprint 9 yet.

Current status: `accept_with_repairs`.

Sprint 8-Lite remains a small BMC-0A prior-art boundary note. Keep it small. Do not expand into a full literature-audit system.

## Required repairs

### 1. Make prior-art source labels honest and conservative

Review:

```text id="7m2zkr"
internal/bmc/priorart/report.go
```

The generated prior-art sources currently use placeholder-like titles/authors while some entries are marked `skim_reviewed` or `abstract_reviewed`.

Repair rule:

```text id="ghhn6l"
If source metadata is placeholder/minimal, mark review_status as seed_unreviewed or human_review_required.
```

Do not mark any source as:

```text id="b2dpn4"
skim_reviewed
abstract_reviewed
```

unless the report contains defensible real metadata and the implementation actually justifies that review status.

So the safest Sprint 8.1 repair is:

```text id="2q7swl"
Set all seeded prior-art source entries to review_status = seed_unreviewed or human_review_required.
Soften boundary_use wording.
```

Avoid strong phrases like:

```text id="ar4tf0"
establishes baseline
settles prior art
proves overlap
confirms known status
```

Use conservative wording like:

```text id="fmc861"
Seed source for future human review of Bohmian minisuperspace prior art.
Seed source for checking overlap with scalar-field minisuperspace claims.
Seed source for future review of Bohmian quantum-cosmology background.
```

### 2. Enforce claim-specific boundary-status direction

Review:

```text id="6vx7ss"
internal/bmc/priorart/validate.go
```

The validator must not merely check that `boundary_status` is one of the allowed values.

It must enforce the expected status direction per required claim.

Required claim-specific allowed status sets:

```text id="19y4g6"
bmc_uses_bohmian_minisuperspace_trajectories:
  established_prior_art | likely_prior_art

bmc_uses_wheeler_dewitt_toy_equation:
  established_prior_art | likely_prior_art

bmc_uses_scalar_field_clock_checks:
  likely_prior_art | unknown_requires_review

bmc_uses_quantum_potential_diagnostics:
  established_prior_art | likely_prior_art

bmc_uses_node_obstruction_detection:
  implementation_variant | likely_prior_art | unknown_requires_review

bmc_uses_local_branch_segmentation:
  implementation_variant | unknown_requires_review

bmc_defines_null_model_scaffold:
  workflow_distinctive_candidate | unknown_requires_review

bmc_uses_ebp_debt_gates:
  workflow_distinctive_candidate | unknown_requires_review

bmc_uses_lean_policy_contracts:
  workflow_distinctive_candidate | unknown_requires_review

bmc_does_not_claim_recovery:
  not_claimed

bmc_does_not_claim_scientific_novelty:
  blocked
```

Validation must reject a known physics-side ingredient being reclassified as `workflow_distinctive_candidate` if that claim-specific set does not allow it.

Add tests:

```text id="jkl021"
TestPriorArtBoundaryRejectsWrongStatusDirectionForPhysicsClaim
TestPriorArtBoundaryRejectsWrongStatusDirectionForWorkflowClaim
TestPriorArtBoundaryRejectsWrongStatusDirectionForNotClaimedRecovery
TestPriorArtBoundaryRejectsWrongStatusDirectionForNoveltyBoundary
```

### 3. Make validation errors phrase-safe

Review all validation paths for:

```text id="62tlh8"
review_status
source_kind
boundary_status
gate name
gate status
forbidden phrase scan
```

Validation errors must not echo forbidden or invalid enum strings back to logs.

Bad:

```text id="8tl072"
forbidden boundary status "scientifically_original"
```

Good:

```text id="m5cwx9"
forbidden boundary status detected at boundary_claims[3]
```

Bad:

```text id="kd5lip"
invalid review_status "confirms_novelty"
```

Good:

```text id="r16fb9"
invalid review_status detected at prior_art_sources[2]
```

Add tests:

```text id="f5v7fs"
TestPriorArtBoundaryForbiddenBoundaryStatusErrorIsPhraseSafe
TestPriorArtBoundaryForbiddenReviewStatusErrorIsPhraseSafe
TestPriorArtBoundaryForbiddenPhraseErrorIsPhraseSafe
```

The tests should verify that returned error messages do not contain the forbidden strings.

### 4. Clarify EBP debt vocabulary

The generated report uses runtime EBP labels such as:

```text id="5ghfkk"
active
partial
contested
absent
```

The adversarial review classification used a separate vocabulary:

```text id="u2n4rw"
unpaid
partial
retired
contested
overclaimed
absent
```

Do not confuse these.

Repair with one of these two approaches.

Preferred minimal approach:

```text id="xcf1co"
Add a report field:
"ebp_debt_vocabulary": "ptw_runtime_debt_status_v0.1"
```

Then validate that `ebp_debt` uses the PTW runtime labels intentionally.

Also add a warning or note:

```text id="tzuju8"
Runtime EBP debt labels are not adversarial-review classifications.
```

Alternative approach:

```text id="28hmfl"
Rename ebp_debt to runtime_ebp_debt.
```

Prefer the minimal field addition to avoid unnecessary schema churn.

Add test:

```text id="r2o3ei"
TestPriorArtBoundaryDeclaresDebtVocabulary
```

### 5. Strengthen CLI tests

The adversarial review says:

```text id="uwmx5k"
TestPriorArtBoundaryCLIRouting exists by name but does not actually exercise CLI routing.
TestPriorArtBoundaryUnknownProfileFails exists by name but tests schema validation rather than CLI unknown-profile failure.
```

Repair tests so they actually run the command path.

Acceptable methods:

```text id="t56cuw"
Compile test binary and run subprocess commands.
```

or:

```text id="2xjl74"
Factor CLI routing into testable functions and test those directly.
```

Required tests:

```text id="c8hjgj"
TestPriorArtBoundaryCLIRoutingRunsGenerateValidateSummarize
TestPriorArtBoundaryUnknownProfileFailsAtCLI
```

These must verify:

```text id="k32tid"
prior-art-boundary generation succeeds for bmc0a-prior-art-boundary
validate routes schema bmc0a-prior-art-boundary-v0.1 to priorart validator
summarize routes schema bmc0a-prior-art-boundary-v0.1 to priorart summarizer
unknown profile fails safely
```

### 6. Optional: case-insensitive forbidden-language scanning

Add case-insensitive scanning if simple.

It should catch variants of forbidden phrases without echoing them in error output.

Test:

```text id="wqpcas"
TestPriorArtBoundaryForbiddenPhraseScanIsCaseInsensitive
```

### 7. Optional: clarify README_GOALS

If `README_GOALS` could confuse future full-audit roadmap with Sprint 8-Lite scope, add a clear note:

```text id="x8ivxp"
Sprint 8-Lite is not the paper-claim extraction or full literature-audit system. It is only a BMC-0A prior-art boundary note.
```

Do not expand the roadmap text.

## Required verification commands

Run:

```bash id="7ljhi1"
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json
./ptw-bmc validate --report out/bmc0a_prior_art_boundary.json
./ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also verify deterministic generation, either by test or command.

## Return walkthrough

Return:

```text id="ux5299"
files changed
source metadata/status repairs
claim-specific boundary status validator changes
phrase-safe error changes
EBP debt vocabulary clarification
CLI test improvements
new/updated test names
Go test result
CLI generation/validation/summary output
Lean build result
remaining limitations
```

## Forbidden scope

Do not implement:

```text id="u5adeq"
full literature-review engine
paper ingestion
PDF parsing
LLM claim extraction
claim-to-profile mapper
citation database
general physics-branch audit registry
null-model runner
Friedmann residual runner
scientific novelty claim
full BMC promotion
```

## Expected status after this repair

Before final acceptance:

```text id="zaoo3q"
prior_art_boundary_note_candidate_after_repairs
```

Maximum after review confirms repairs:

```text id="d9cqtp"
promoted_prior_art_boundary_note_after_repairs
```
