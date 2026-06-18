# BMC-POST-0002 Code Review Prompt: Numerical WdW Residual Integration / Analytic Authority Displacement

You are reviewing **BMC-POST-0002: Numerical WdW Residual Integration / Analytic Authority Displacement** under strict EBP 2.1 discipline.

POST-0002 reportedly changed the BMC plane-wave report path so that:

```text
NumericalResidualAt is now the diagnostic authority for the Wheeler-DeWitt check.
The analytic residual remains only as an oracle/control.
Wrong numerical residual status cannot be reported as pass.
Superposition numerical WdW residual checks are explicitly deferred/not authoritative.
```

Your task is to verify whether this is actually true in code, tests, generated report behavior, validation behavior, and documentation.

## Context

Accepted prior remediation:

```text
BMC-POST-0001 added NumericalResidualAt and negative tests proving wrong WdW inputs can fail in isolated WdW tests.
```

POST-0002 should integrate that into the report path.

It must not claim:

```text
Friedmann recovery
classical-limit recovery
BMC validation
null-model failure
BMC superiority
scientific novelty
full BMC promotion
```

Full BMC remains blocked.

## Files to inspect

Review actual code, not only the walkthrough:

```text
internal/bmc/model/types.go
internal/bmc/wdw/residual.go
internal/bmc/wdw/residual_test.go
internal/bmc/report/report.go
internal/bmc/report/validate.go
internal/bmc/report/report_test.go
docs/postmortem/bmc_post_0002_numerical_wdw_residual_integration.md
cmd/ptw-bmc/main.go
```

If generated output exists, inspect:

```text
/tmp/bmc_post_0002_plane_wave_report.json
out/*.json if relevant
```

## Core review question

```text
Does the plane-wave report path use the numerical WdW residual as the diagnostic authority, while keeping the analytic residual only as oracle/control?
```

## Review targets

### 1. Numerical authority path

Verify that in the plane-wave report generation path:

```text
NumericalResidualAt is actually called.
NumericalResidualMagnitude is populated.
NumericalResidualTolerance is populated.
NumericalResidualStatus is populated.
NumericalResidualAuthority is diagnostic_authority.
WdW pass/fail status follows the numerical residual, not the analytic residual alone.
```

Mark as blocker if analytic residual can still make the WdW check pass when the numerical residual reports violation/error.

### 2. Analytic residual oracle/control

Verify that analytic residual remains only as:

```text
oracle/control metadata
not the primary authority
not sufficient for pass
```

Expected authority wording should not imply analytic residual is authoritative.

Mark as blocker if analytic residual is still the decisive gate.

### 3. Wrong-constraint report-path failure

Confirm a report-level or report-generation test exists for:

```text
k² != ω²
```

Expected:

```text
numerical_residual_violation_detected
check.Status is not pass
check.Pass is false
promotion gate remains blocked
validation rejects inconsistent pass claims
```

Required or equivalent tests:

```text
TestReportNumericalWDWResidualBlocksWrongPlaneWaveConstraint
TestReportValidationRejectsNumericalResidualViolationClaimedAsPass
TestPromotionGateBlocksNumericalWDWViolation
```

Mark as blocker if wrong-constraint failure exists only in isolated WdW tests but not report-path tests.

### 4. Valid report-path control

Confirm a valid plane-wave report-path test exists:

```text
k² = ω²
```

Expected:

```text
numerical_residual_pass
numerical residual magnitude <= tolerance
analytic residual near zero
report validates
full BMC remains blocked due to unresolved debts
```

Required or equivalent test:

```text
TestReportNumericalWDWResidualAcceptsConstraintShellPlaneWave
```

Mark as blocker if valid control was not tested through the report path.

### 5. Inconsistent state validation

Validation must reject a report where:

```text
numerical_residual_status = numerical_residual_violation_detected
check.Status = pass
check.Pass = true
```

Validation must also reject:

```text
numerical_residual_status = numerical_residual_error
check.Status = pass
check.Pass = true
```

Mark as blocker if these inconsistent states validate.

### 6. Status/authority constants

Check that status and authority values are constants, not loose ad hoc strings.

Expected values:

```text
numerical_residual_pass
numerical_residual_violation_detected
numerical_residual_not_computed
numerical_residual_error

diagnostic_authority
oracle_control_only
not_authoritative
```

Validation should reject unknown numerical residual statuses or authorities when fields are present.

Mark as repair-required if constants exist but validation does not reject unknown values.

Mark as blocker if unknown success-like values can validate.

### 7. Tolerance and step audit

Verify that numerical step and tolerance are defined in one place, for example:

```text
WDWNumericalResidualStep = 1e-4
WDWNumericalResidualTolerance = 1e-6
```

Confirm tests use explicit tolerances and do not rely on vague nonzero checks.

Mark as repair-required if tolerances are scattered or undocumented.

### 8. Superposition deferral

Verify that superposition report path sets:

```text
numerical_residual_status = numerical_residual_not_computed
numerical_residual_authority = not_authoritative
```

It must not imply numerical WdW check passed for superposition.

Mark as blocker if superposition numerical residual is reported as pass without computation.

### 9. Documentation review

Check documentation note states:

```text
Numerical residual is a toy finite-difference diagnostic.
Analytic residual remains oracle/control.
This is not a full WdW solver.
This does not implement Friedmann recovery.
This does not validate BMC physics.
Full BMC remains blocked.
```

Mark as repair-required if documentation overclaims.

### 10. CLI review

Run or inspect:

```bash
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/wdw -v -count=1
GOCACHE=/tmp/go-build-cache go test ./internal/bmc/report -v -count=1
GOCACHE=/tmp/go-build-cache go test ./... -count=1
cd BMC && /home/chaschel/.elan/bin/lake build
```

Also run, using the actual accepted profile:

```bash
./ptw-bmc run --profile bmc0a-plane --out /tmp/bmc_post_0002_plane_wave_report.json
./ptw-bmc validate --report /tmp/bmc_post_0002_plane_wave_report.json
./ptw-bmc summarize --report /tmp/bmc_post_0002_plane_wave_report.json
```

Check summary output does not claim recovery or validation.

## EBP debt classification

Classify:

```text
needConstraintViolationTests
needNumericalErrorAudit
needNontrivialPhysicsCase
needToyCheck
needFaithfulnessReview
containsFinalTruthClaim
full_bmc_toy_gate
AnalyticAuthorityDisplacement
ReportPathFailureDetection
SuperpositionNumericalResidualScope
NoRecoveryClaimBoundary
LeanPolicyBoundary
```

Allowed values:

```text
unpaid
partial
retired_for_plane_wave_scope
contested
overclaimed
absent
blocked
```

## Required output JSON

Return exactly this JSON shape:

```json
{
  "summary": "",
  "overall_verdict": "accept|accept_with_repairs|reject_for_now",
  "ebp_debt_review": {
    "needConstraintViolationTests": "",
    "needNumericalErrorAudit": "",
    "needNontrivialPhysicsCase": "",
    "needToyCheck": "",
    "needFaithfulnessReview": "",
    "containsFinalTruthClaim": "",
    "full_bmc_toy_gate": "",
    "AnalyticAuthorityDisplacement": "",
    "ReportPathFailureDetection": "",
    "SuperpositionNumericalResidualScope": "",
    "NoRecoveryClaimBoundary": "",
    "LeanPolicyBoundary": ""
  },
  "numerical_authority_findings": [],
  "analytic_oracle_findings": [],
  "wrong_constraint_report_path_findings": [],
  "valid_control_report_path_findings": [],
  "validation_consistency_findings": [],
  "status_authority_constant_findings": [],
  "tolerance_step_findings": [],
  "superposition_deferral_findings": [],
  "documentation_findings": [],
  "cli_findings": [],
  "missing_tests": [],
  "required_repairs_before_acceptance": [],
  "optional_repairs": [],
  "authority_displacement_verdict": {
    "status": "accepted|contested|rejected",
    "reason": ""
  },
  "promotion_recommendation": "do_not_promote|accept_post_0002_for_plane_wave_report_path_scope",
  "next_smallest_useful_move": ""
}
```

## Strict recommendation limit

Even if POST-0002 passes perfectly, the maximum allowed recommendation is:

```text
accept_post_0002_for_plane_wave_report_path_scope
```

Never recommend:

```text
Friedmann recovery
classical-limit recovery
BMC validation
null-model failure
BMC superiority
full BMC
scientific novelty
```

Remember:

```text
A numerical residual diagnostic is not a full WdW solver.
A valid plane-wave numerical check is not Friedmann recovery.
Blocking a wrong constraint is failure detection, not physics validation.
Superposition numerical WdW authority remains deferred unless explicitly implemented.
Full BMC remains blocked.
```
