# Walkthrough: BMC Sprint 10 Candidate Local-Branch Residual Runner

This walkthrough details the implementation, validation, and verification results for the candidate local-branch residual runner under EBP 2.1 discipline.

## Files Added/Changed
- [internal/bmc/residualrun/contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/contracts.go) - Declared enums and constants for eligibility, residual status, provenances, and interpretation statuses.
- [internal/bmc/residualrun/inputs.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/inputs.go) - Defined the `SourceArtifactRef` struct.
- [internal/bmc/residualrun/branches.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/branches.go) - Defined the `LocalBranchEligibility` struct.
- [internal/bmc/residualrun/residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residual.go) - Defined `CandidateResidualMetrics` and `CandidateResidualDiagnostic` structures.
- [internal/bmc/residualrun/compare.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/compare.go) - Defined `ResidualNullComparison` struct.
- [internal/bmc/residualrun/gates.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/gates.go) - Defined the `ResidualRunGate` struct.
- [internal/bmc/residualrun/report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/report.go) - Defined `ResidualRunReport` structure and default report generator `GenerateDefaultReport()`.
- [internal/bmc/residualrun/validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/validate.go) - Implemented validation rules (raw JSON optional metrics presence, EBP debt, comparisons, gates, local-only bounds, phrase-safe scanner).
- [internal/bmc/residualrun/residualrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residualrun_test.go) - Wrote 25 unit tests covering all rules and CLI routings.
- [cmd/ptw-bmc/main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go) - Integrated CLI subcommand routing for `run-residuals`, `validate`, and `summarize`.
- [BMC/BMC/ResidualRun.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ResidualRun.lean) - Implemented Lean policy safety theorems.
- [BMC/BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean) - Added ResidualRun import.

---

## Applied Revisions

### 1. Conditional Candidate Residual Setting
- `candidate_residual_computed` is conditional on the existence of at least one eligible branch with computed diagnostics.
- If no eligible branch exists, `candidate_residual_computed` must be `false` and the interpretation status must be blocked (`InterpretBlockedByNoEligibleLocalBranch`).

### 2. Tightened Deterministic Fixture Use
- Validation requires that if `ResidualProvenance = deterministic_fixture`, the interpretation status is restricted to `diagnostic_comparison_only` or `blocked_by_convention_debt`, preventing it from being promoted as a separation candidate.

### 3. Source/Provenance Honesty Struct
- Added `SourceArtifactRef` to record whether source artifacts were read (`file_read`) or summarized (`source_artifact_summary`). In default configurations, we use `source_artifact_summary` without pretending files were read.

### 4. "Candidate" Wording
- Transitioned all terminology in report fields, gates, comments, and logs to guarded phrases: `candidate residual diagnostic`, `candidate local-branch residual`, `residual-style diagnostic`, avoiding bare recovery or confirmation terms.

### 5. Convention Debts Enforcement
- Enforced that all six required debts must be declared in the ledger and marked `unpaid` or `contested` (rejecting `retired`).

### 6. Conditional Comparison Validation
- Target/null comparisons require that at least one target residual is computed (`residual_computed = true`) and that compared metrics are non-empty.

### 7. Explicit Local-Only Boundary
- The report strictly carries `"local_branch_only": true` and `"global_cosmology_claim": false` flags. Any other values are rejected by the validator.

---

## Verification Results

### 1. Go Unit Tests
All 25 Go unit tests pass cleanly:
```bash
$ go test ./...
?       github.com/PithomLabs/bmc/cmd/ptw-bmc   [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/audit    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockdiag        (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/clockseg (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/friedmannspec    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/guidance (cached)
?       github.com/PithomLabs/bmc/internal/bmc/invariant        [no test files]
ok      github.com/PithomLabs/bmc/internal/bmc/model    (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/nullrun  (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/nullspec (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/priorart (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/residualrun      0.349s
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```

### 2. Lean Policy Safety Compilation
All safety theorems in `ResidualRun.lean` build successfully:
```bash
$ cd BMC && /home/chaschel/.elan/bin/lake build
Build completed successfully (13 jobs).
```

### 3. CLI Subcommands Demonstration

#### Report Generation:
```bash
$ ./ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json
Successfully ran residual profile 'bmc0a-local-residual' and generated report: out/bmc0a_local_residual.json
```

#### Report Validation:
```bash
$ ./ptw-bmc validate --report out/bmc0a_local_residual.json
Candidate Local-Branch Residual Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
```

#### Human-Readable Summary:
```bash
$ ./ptw-bmc summarize --report out/bmc0a_local_residual.json
BMC Sprint 10 Candidate Local-Branch Residual Summary
Schema Version: bmc0a-local-residual-v0.1
Scope: bmc0a_only
Candidate Residual Computed: true
Recovery Claim: false
Scientific Novelty Claim Made: false
BMC Beats Null Models Claim: false
Full BMC: blocked
Eligible Local Branches: 2
Candidate Residual Diagnostics: 2
Residual/Null Comparisons: 1
Interpretation Status: diagnostic_comparison_only
Promotion Status: candidate_residual_runner_candidate_only
```

---

## Remaining Limitations
- Sprint 10 is limited to candidate local-branch residual diagnostics; it does not claim global Friedmann recovery.
- All convention debts remain unpaid or contested. Full BMC promotion remains blocked.
