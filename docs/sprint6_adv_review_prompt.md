You are an adversarial reviewer for an EBP 2.1-governed research/software artifact.

Review **BMC Sprint 6: Friedmann-Residual Specification and Gate Design**.

Sprint 6 follows the accepted Sprint 5 artifact:

```text
friedmann_readiness: local_only_candidate
clock_choice_debt: active
full_bmc_toy_gate: blocked
local_only_candidate is not Friedmann recovery readiness
```

Sprint 6 must remain a **specification and gate-design sprint only**.

It must not compute or imply Friedmann recovery.

## Current accepted context

Accepted artifacts under strict EBP limits:

```text
Sprint 1: BMC-0A plane-wave control artifact
Sprint 2: BMC-0A two-plane-wave superposition control artifact
Sprint 2: BMC-0A node-obstruction detection artifact
Sprint 3: BMC-0A numerical robustness/convergence audit artifact
Sprint 4: BMC-0A clock-fragility diagnostic artifact
Sprint 5: BMC-0A clock-readiness/local segmentation artifact
```

Sprint 6 claims to add:

```text
internal/bmc/friedmannspec/
spec-friedmann CLI subcommand
bmc0a-friedmann-spec-v0.1 report schema
candidate Friedmann/FRW variable maps
branch residual-readiness contracts
derivative-readiness contracts
residual formula candidate registry
future null-model requirements
10 policy safety gates
FriedmannSpec.lean policy/safety contracts
```

## Forbidden claims

The artifact must **not** claim:

```text
Friedmann residual recovery
Friedmann equations recovered
ready for Friedmann recovery
classical cosmology recovered
BMC validates FRW
full BMC validation
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

The full BMC toy gate must remain blocked.

The report must state:

```text
residual_computed: false
friedmann_recovery_claim: false
spec_scope: candidate_specification_only
promotion_status: planned_candidate_only
```

The following debts must remain active or contested:

```text
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
null_model_debt
needFaithfulnessReview
```

## Materials to review

Review actual source and generated artifacts, not only the walkthrough:

```text
README.md
cmd/ptw-bmc/**
internal/bmc/**
internal/bmc/friedmannspec/**
BMC/**
BMC/BMC/FriedmannSpec.lean
out/bmc0a_friedmann_spec.json
out/bmc0a_clock_readiness.json
out/bmc0a_clock_fragility.json
out/bmc0a_superposition_robustness.json
go.mod
all Go tests
Sprint 1 final report
Sprint 2 final report
Sprint 3 final report
Sprint 4 final report
Sprint 5 final report
Sprint 6 walkthrough/report
```

If any expected file or artifact is missing, report it as debt.

## Expected Sprint 6 design

Expected package:

```text
internal/bmc/friedmannspec/
```

Expected files:

```text
contracts.go
mapping.go
branch_requirements.go
derivatives.go
residual_spec.go
gates.go
report.go
validate.go
friedmannspec_test.go
```

Expected CLI:

```bash
ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
ptw-bmc validate --report out/bmc0a_friedmann_spec.json
ptw-bmc summarize --report out/bmc0a_friedmann_spec.json
```

Expected schema:

```text
bmc0a-friedmann-spec-v0.1
```

Expected Lean file:

```text
BMC/BMC/FriedmannSpec.lean
```

## Central review principle

Sprint 6 may define **candidate residual contracts**, but:

```text
a specification is not a computation
a residual formula candidate is not a residual result
a local branch is not a global clock
local_only_candidate is not Friedmann recovery readiness
null-model debt remains active
clock-choice debt remains active
full BMC remains blocked
```

The key review question is:

```text
Does Sprint 6 safely define what a future Friedmann-style residual would require, without computing the residual or implying recovery?
```

Mark as blocker if any code path computes a Friedmann residual, evaluates recovery, or unblocks full BMC.

## Required review targets

### 1. Report identity and schema

Review `report.go`, `validate.go`, CLI output, and generated JSON.

Correct behavior:

```text
schema_version = bmc0a-friedmann-spec-v0.1
spec_kind = friedmann_residual_specification
spec_scope = candidate_specification_only
residual_computed = false
friedmann_recovery_claim = false
```

Mark as blocker if:

```text
ClockReadinessReport is reused as the Sprint 6 artifact identity
spec_scope allows ready/validated/recovered
residual_computed can become true
friedmann_recovery_claim can become true
```

### 2. No residual computation

Search all Sprint 6 code.

There must be no numerical Friedmann residual computation.

Allowed:

```text
candidate formula description
required variables registry
required derivative registry
candidate map declarations
gate definitions
future null-model requirements
```

Forbidden:

```text
actual residual value
residual pass/fail
recovery metric
thresholded Friedmann error
H² - Cρ numerical calculation
claim that candidate equation matches trajectory
```

Mark as blocker if any such computation exists.

### 3. Candidate maps

Review `mapping.go`.

Expected structure:

```text
FriedmannCandidateMap
```

Allowed map status values:

```text
candidate_only
contested
blocked
```

Forbidden status values:

```text
validated
recovered
proved
ready
pass
```

Check that candidate maps explicitly track:

```text
clock_choice_debt
classical_target_debt
unit_convention_debt
sign_convention_debt
normalization_debt
```

The candidate classical target may be described as:

```text
flat FRW + massless scalar candidate target
```

but must not be asserted as the confirmed physics target.

### 4. Branch residual-readiness contracts

Review `branch_requirements.go`.

Expected:

```text
FriedmannBranchRequirement
```

Allowed readiness values:

```text
blocked
candidate_only
contested
```

Forbidden:

```text
ready
pass
recovered
validated
```

Check that branch contracts use Sprint 5 local branches only as **inputs to future-readiness contracts**, not as evidence that residual computation is now allowed.

Check that branch contracts account for:

```text
local φ monotonicity
α(φ) single-valuedness
minimum samples
clock range
lambda range
node contact
Q finiteness away from nodes
derivative readiness
```

### 5. Derivative-readiness contracts

Review `derivatives.go`.

Expected:

```text
DerivativeReadinessCheck
```

It may specify derivative requirements, but must not compute a final Friedmann residual.

Check that it tracks:

```text
dα/dφ stability
d²α/dφ² availability if needed
finite-difference sensitivity
branch endpoint exclusion
turning-point exclusion
near-node exclusion
minimum samples after exclusions
```

Allowed status values:

```text
candidate_only
blocked
contested
```

Mark as blocker if derivative readiness is treated as residual success.

### 6. Residual formula candidate registry

Review `residual_spec.go`.

Expected:

```text
ResidualFormulaCandidate
```

It may describe formula candidates, required variables, required derivatives, convention debts, and required null models.

It must not evaluate formulas numerically.

Allowed status values:

```text
candidate_only
blocked
contested
```

Forbidden:

```text
validated
recovered
proved
ready
pass
```

### 7. Null-model requirements

Sprint 6 should not perform null-model comparison, but must specify future null-model requirements.

Expected requirements include:

```text
constant-phase control
randomized phase control
matched amplitude / randomized phase control
classical FRW reference trajectory
same branch segmentation under null wavefunctions
node-neighborhood stress case
clock-choice alternative branch diagnostic
```

Review whether each requirement has:

```text
purpose
required_before_residual_promotion = true
status = planned|deferred|blocked
reason
```

Mark as blocker if null-model debt is marked retired.

### 8. Gate design

Review `gates.go`.

Expected gates:

```text
toy_analysis_only_gate
no_final_truth_claim_gate
local_branch_only_gate
clock_choice_debt_active_gate
classical_target_candidate_only_gate
unit_convention_debt_gate
null_model_debt_gate
faithfulness_contested_gate
no_residual_computation_gate
full_bmc_blocked_gate
```

Allowed gate statuses:

```text
pass
blocked
contested
```

Important:

```text
no_residual_computation_gate passes only if no residual is computed
full_bmc_blocked_gate passes only if full BMC remains blocked
null_model_debt_gate passes only if null-model debt remains active/planned, not retired
```

### 9. Strict validation

Review `ValidateFriedmannSpecReport`.

It should reject:

```text
final_truth_claim = true
toy_analysis_only = false
missing or wrong schema_version
wrong spec_kind
invalid spec_scope
spec_scope = ready|validated|recovered
residual_computed = true
friedmann_recovery_claim = true
promotion_gate.status != blocked
clock_choice_debt != active
classical_target_debt != active
unit_convention_debt != active
sign_convention_debt != active
normalization_debt != active
null_model_debt != active
needNullModel not active
missing no_residual_computation_gate
no_residual_computation_gate.status != pass
missing full_bmc_blocked_gate
full_bmc_blocked_gate.status != pass
candidate map status = validated|recovered|proved|ready|pass
branch residual readiness = ready|pass|recovered|validated
residual formula status = validated|recovered|proved|ready|pass
empty null_model_requirements
warnings missing “No Friedmann residual was computed”
warnings missing “No Friedmann recovery is claimed”
unknown JSON fields
nonfinite numeric values
```

Check for strict decoding:

```go
json.Decoder.DisallowUnknownFields()
```

### 10. Deterministic JSON

Check whether repeated generation produces byte-identical JSON.

Review deterministic ordering for:

```text
source_artifacts
candidate_maps
branch_requirements
derivative_readiness_checks
residual_formula_candidates
null_model_requirements
gates
warnings
```

### 11. CLI behavior

Run or review reported runs:

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json
./ptw-bmc validate --report out/bmc0a_friedmann_spec.json
./ptw-bmc summarize --report out/bmc0a_friedmann_spec.json
cd BMC && lake build
```

Check whether:

```text
spec-friedmann subcommand exists
unknown spec-friedmann profiles fail safely
validate routes friedmann spec schema to friedmannspec validator
summarize routes friedmann spec schema to friedmannspec summarizer
ordinary BMC, robustness, clock-fragility, and clock-readiness validators are not weakened
errors are explicit
```

### 12. Lean safety contracts

Review:

```text
BMC/BMC/FriedmannSpec.lean
BMC/BMC.lean
```

Check that:

```text
lake build succeeds
no sorry/admit exists
BMCFriedmannSpecReport exists
friedmann_spec_requires_toy_only exists
friedmann_spec_blocks_final_truth exists
friedmann_spec_forbids_residual_computation exists
friedmann_spec_forbids_recovery_claim exists
friedmann_spec_requires_full_bmc_blocked exists
friedmann_spec_requires_clock_choice_debt_active exists
friedmann_spec_requires_classical_target_debt_active exists
friedmann_spec_requires_unit_convention_debt_active exists
friedmann_spec_requires_sign_convention_debt_active exists
friedmann_spec_requires_normalization_debt_active exists
friedmann_spec_requires_null_model_debt_active exists
friedmann_spec_does_not_imply_friedmann_recovery exists
friedmann_spec_does_not_imply_full_bmc exists
Lean remains policy/safety only
no theorem claims Friedmann recovery, FRW validation, or classical cosmology recovery
```

### 13. Forbidden phrase scan

Search source, tests, comments, report text, CLI output, summaries, README, and walkthrough for dangerous phrases:

```text
ready for Friedmann recovery
recovers Friedmann
Friedmann residual passes
classical cosmology recovered
FRW recovered
full BMC validated
validates quantum gravity
solves the problem of time
proves Bohmian mechanics
derives spacetime
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Acceptable phrases:

```text
candidate specification only
candidate map
candidate residual formula
future residual requirement
no residual computed
no Friedmann recovery claimed
full BMC remains blocked
null-model debt remains active
clock-choice debt remains active
```

## Code review checks

Review these details:

```text
friedmannspec package structure
spec_scope allowed values
residual_computed hard false invariant
friedmann_recovery_claim hard false invariant
candidate maps candidate_only/contested/blocked only
branch readiness candidate_only/contested/blocked only
derivative readiness candidate_only/contested/blocked only
residual formula candidates not numerically evaluated
null-model requirements present and not retired
10 safety gates present
strict schema validation
unknown-field rejection
finite metric validation if any numeric fields exist
deterministic output ordering
CLI routing by schema version
full_bmc_toy_gate blocked
clock_choice_debt active
classical_target_debt active
unit_convention_debt active
sign_convention_debt active
normalization_debt active
needNullModel active
faithfulness contested
no external dependencies added
```

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
LeanVerification
FriedmannSpecDiagnosticIntegrity
FriedmannReadinessBoundary
NoResidualComputationBoundary
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
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
    "LeanVerification": "",
    "FriedmannSpecDiagnosticIntegrity": "",
    "FriedmannReadinessBoundary": "",
    "NoResidualComputationBoundary": ""
  },
  "spec_findings": [],
  "physics_boundary_findings": [],
  "code_findings": [],
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
  "promotion_recommendation": "do_not_promote|friedmann_spec_candidate_only|promoted_friedmann_spec_artifact_after_repairs",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if Sprint 6 passes perfectly, the maximum allowed recommendation is:

```text
promoted_friedmann_spec_artifact_after_repairs
```

Never recommend promotion as:

```text
Friedmann recovery
ready for Friedmann recovery
full BMC
full quantum gravity
proof of Bohmian mechanics
solution to the problem of time
spacetime emergence proof
valid φ-clock for full cosmology
valid α-clock for full cosmology
```

Remember:

```text
A specification is not a computation.
A residual formula candidate is not a residual result.
A local branch is not a global clock.
A local-only candidate is not Friedmann recovery readiness.
Null-model debt remains active.
Clock-choice debt remains active.
Full BMC remains blocked.
```
