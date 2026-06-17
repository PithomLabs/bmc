# Walkthrough: BMC Sprint 9 Null-Model Runner

This walkthrough details the design, implementation, and verification results for the narrow BMC-0A null-model runner under strict EBP 2.1 discipline.

## Files Changed/Added
- [internal/bmc/nullrun/contracts.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/contracts.go) - Added run status, diagnostic provenance, diagnostic status, and interpretation status constants.
- [internal/bmc/nullrun/report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/report.go) - Defined JSON report schema structures, registered 7 null models with seeds, and set default diagnostic comparators.
- [internal/bmc/nullrun/runner.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/runner.go) - Implements deterministic null diagnostic fixtures and real diagnostic calls where available.
- [internal/bmc/nullrun/validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/validate.go) - Enforces EBP 2.1 gates, schema parameters, target-null comparison constraints, and performs case-insensitive, phrase-safe forbidden phrase scanning.
- [internal/bmc/nullrun/nullrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/nullrun_test.go) - Includes unit tests validating schema correctness, validation failure states, phrase-safety, and subprocess-based CLI execution.
- [cmd/ptw-bmc/main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go) - Integrated CLI subcommand routing for `run-nullmodels` and added report format validations and summary output formatting.
- [BMC/BMC/NullRun.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/NullRun.lean) - Implemented optional Lean policy verification contracts.
- [BMC/BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean) - Imported `BMC.NullRun` to verify on lake build.

---

## Applied Revisions

### 1. Honest Diagnostics Wording
The implementation plan and runner logic were revised to reflect that we generate deterministic fixtures and real diagnostic calls where available, marking any unavailable runs as blocked or deferred rather than simulating or fabricating numbers.

### 2. Diagnostic Provenance
Each `NullModelRun` contains `diagnostic_provenance` (`DiagnosticProvenance string`) with permitted values:
- `computed_from_existing_bmc_diagnostics`
- `deterministic_fixture`
- `source_artifact_summary`
- `blocked`
- `deferred`

Any other values (including `assumed`, `fabricated`, `inferred_success`, and `validated_physics`) are strictly rejected. Additionally, these forbidden strings are globally blacklisted in the case-insensitive phrase scanner.

### 3. Comparable Diagnostics Existence Rule
Added validation constraint: if no null model runs have generated diagnostics (`RunStatus != RunStatusDiagnosticsGenerated`), the target-null comparison is blocked. Specifically:
- `target_null_comparison_computed` must be `false`.
- `interpretation_status` must be `blocked_by_no_comparable_null_diagnostics`.

### 4. Classical Reference Trajectory Handling
`classical_frw_reference_trajectory` is treated as a reference comparator rather than a null wavefunction:
- `run_status` must be `diagnostics_generated | blocked | deferred`.
- `diagnostic_provenance` must be `deterministic_fixture | source_artifact_summary | blocked`.
- `notes` must declare: `"reference comparator only; no residual or recovery interpretation"`.

---

## Verification Results

### 1. Go Unit Tests
All package tests compile and pass successfully:
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
ok      github.com/PithomLabs/bmc/internal/bmc/nullrun  0.315s
ok      github.com/PithomLabs/bmc/internal/bmc/nullspec (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/obstruction      (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/priorart (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/qpotential       (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/report   (cached)
ok      github.com/PithomLabs/bmc/internal/bmc/wave     (cached)
?       github.com/PithomLabs/bmc/internal/bmc/wdw      [no test files]
```

### 2. Lean Policy Safety Compilation
The Lean 4 theorem projections are corrected to account for right-associative conjunctions. The project builds cleanly with 0 warnings:
```bash
$ cd BMC && /home/chaschel/.elan/bin/lake build
Build completed successfully (12 jobs).
```

### 3. CLI Subcommands Demonstration

#### Report Generation:
```bash
$ ./ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json
Successfully ran null models profile 'bmc0a-nullrun' and generated report: out/bmc0a_nullrun.json
```

#### Report Validation:
```bash
$ ./ptw-bmc validate --report out/bmc0a_nullrun.json
Null Model Run Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.
```

#### Human-Readable Summary:
```bash
$ ./ptw-bmc summarize --report out/bmc0a_nullrun.json
BMC Sprint 9 Null-Model Runner Summary
Schema Version: bmc0a-nullrun-v0.1
Scope: bmc0a_only
Residual Computed: false
Null Diagnostics Computed: true
Target/Null Comparison Computed: true
Recovery Claim: false
Scientific Novelty Claim Made: false
Full BMC: blocked
Null Models Registered: 7
Null Models With Diagnostics: 4
Null Models Blocked: 1
Interpretation Status: diagnostic_comparison_only
Promotion Status: null_model_runner_candidate_only
```

---

## Remaining Limitations
- Sprint 9 is limited to a comparative diagnostic runner; it does not compute Friedmann residuals or support physics-limit recovery assertions.
- The comparison does not claim BMC beats null models or outperformed controls. Full BMC promotion remains blocked.
