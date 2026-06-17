# Phase 1 Postmortem Inventory: Sprints 1–11 and BMC Surface Area Classification

## 1. Executive Summary

- **Methodological Toy Benchmark:** Bohmian Minisuperspace Cosmology (BMC-0.1) is currently a methodological toy benchmark and audit stack, not a physics result. It operates entirely within a simplified minisuperspace model without inhomogeneous perturbations, gauge fields, or full quantum gravity recovery.
- **Lack of Nontrivial Physics Failure Detection:** The Sprint 1–11 stack succeeds in blocking premature promotions and overclaims via validation gates, but it still lacks nontrivial physics failure detection. Constraints can be violated without triggering explicit errors in the core numerical solvers because numerical evaluation does not verify physical consistency (e.g., matching the Friedmann target or auditing the central finite differences against wrong wavefunctions).
- **Full BMC Blocked:** Full Bohmian Minisuperspace Cosmology promotion remains blocked. All promotion gates are set to `blocked` or `contested`, and key physics debts (e.g., factor ordering, minisuperspace metric conventions, classical target equations, and empirical benchmarks) remain completely unpaid.

---

## 2. Current Accepted Artifact Stack

The Sprint 1–11 audit stack is frozen. The accepted historical status of each sprint is classified below.

### Sprint 1: Plane-Wave Control
* **artifact_kind:** Plane-wave control report
* **accepted_status:** Accepted (Sprint 1)
* **real_computation_level:** Toy (Analytical plane-wave WdW residual and Euler trajectory integration)
* **scaffolding_level:** Low (Basic report generator and schema validator)
* **promotion_scope:** Bounded plane-wave control artifact only
* **forbidden_interpretations:** Spacetime emergence, solving the problem of time, physical cosmology recovery, scientific novelty, or full BMC promotion
* **remaining_debts:** Needs superposition wavefunctions, node contact detection, numerical robustness audits, and comparison null models

### Sprint 2: Superposition + Node Obstruction
* **artifact_kind:** Two-term superposition control report
* **accepted_status:** Accepted after repairs (Sprint 2)
* **real_computation_level:** Toy (RK4 trajectory integration, node short-circuiting, phase gradient evaluation, and basic node contact checks)
* **scaffolding_level:** Medium (Obstruction registry and node detection checks)
* **promotion_scope:** Superposition control and node obstruction detection validation gate
* **forbidden_interpretations:** Classical limit recovery, WdW solver superiority, physical correctness of trajectory near nodes
* **remaining_debts:** Needs systematic step-size sweeps, threshold sensitivity audits, and clock diagnostics

### Sprint 3: Robustness Audit
* **artifact_kind:** Numerical robustness and convergence report
* **accepted_status:** Accepted (Sprint 3)
* **real_computation_level:** Toy (Step-size sweeps, node-threshold sweeps, phase gradient sweeps, parameter perturbations, and initial node probe offsets)
* **scaffolding_level:** High (Robustness outcome classifier and nested sweep loops)
* **promotion_scope:** Robustness sweep audit gate
* **forbidden_interpretations:** General numerical stability guarantees, physical stability of cosmology
* **remaining_debts:** Needs clock monotonicity diagnostics, local branch segmentation, and null model runner

### Sprint 4: Clock Fragility Diagnostic
* **artifact_kind:** Clock fragility report
* **accepted_status:** Accepted (Sprint 4)
* **real_computation_level:** Toy (Monotonicity correlations, turning point detection, and fragility diagnostics along trajectories)
* **scaffolding_level:** Medium (Fragility report generator)
* **promotion_scope:** Clock fragility diagnostic artifact
* **forbidden_interpretations:** Physical relational time recovery, general lapse problem solution
* **remaining_debts:** Needs clock-independent segmentation, Friedmann specification, and null model comparisons

### Sprint 5: Local Branch / Clock Readiness
* **artifact_kind:** Clock readiness report
* **accepted_status:** Accepted (Sprint 5)
* **real_computation_level:** Toy (Clock segmentation, local branch extraction, and clock-independent relational states)
* **scaffolding_level:** Medium (Readiness report generator)
* **promotion_scope:** Clock readiness & segmentation control artifact
* **forbidden_interpretations:** Global relational cosmology history recovery
* **remaining_debts:** Needs Friedmann residual spec, null model specifications, and runner

### Sprint 6: Candidate Friedmann Residual Specification Only
* **artifact_kind:** Friedmann spec report
* **accepted_status:** Accepted (Sprint 6)
* **real_computation_level:** Zero / Specification Only (No derivatives computed; placeholder specification of gates only)
* **scaffolding_level:** High (Prose mappings and gate conditions)
* **promotion_scope:** Candidate Friedmann spec gate
* **forbidden_interpretations:** Friedmann residual computation or recovery claims
* **remaining_debts:** Needs actual numerical Friedmann evaluation, null-model registry, and runner comparisons

### Sprint 7: Null-Model Scaffold
* **artifact_kind:** Null model spec report
* **accepted_status:** Accepted (Sprint 7)
* **real_computation_level:** Zero / Specification Only (registry definitions and planned null comparisons placeholder)
* **scaffolding_level:** High (Registry definitions and comparison contracts)
* **promotion_scope:** Null-model specification gate
* **forbidden_interpretations:** Running null models or claiming superiority over null models
* **remaining_debts:** Running null models, candidate local branch residuals, and comparison audits

### Sprint 8-Lite: Prior-Art Boundary Note
* **artifact_kind:** Prior-art boundary report
* **accepted_status:** Accepted after repairs (Sprint 8.1)
* **real_computation_level:** Zero / Specification Only (Documenting boundary notes and novelty constraints)
* **scaffolding_level:** High (Policy gates in Go/Lean enforcing novelty bans)
* **promotion_scope:** Prior-art boundary note
* **forbidden_interpretations:** Scientific novelty claim, literature review completion
* **remaining_debts:** Null-model runner, candidate local-branch residuals, and comparison audits

### Sprint 9 + 9.1: Null-Model Runner
* **artifact_kind:** Null model run report
* **accepted_status:** Accepted after repairs (Sprint 9.1)
* **real_computation_level:** Toy (Runs diagnostic calculations for 4 of 7 registered null models, blocks 1, defers 2)
* **scaffolding_level:** High (Null runner validation gates and strict envelope enforcement)
* **promotion_scope:** Null-model runner diagnostic artifact
* **forbidden_interpretations:** BMC beats null models, null models passed/failed, recovery claims
* **remaining_debts:** Candidate local-branch residuals, comparison audits, and target/null separation

### Sprint 10 + 10.3: Candidate Local-Branch Residual Runner
* **artifact_kind:** Candidate local-branch residual report
* **accepted_status:** Accepted after repairs (Sprint 10.3)
* **real_computation_level:** Toy (Per-point candidate local-branch residual computation, local lambda monotonicity validation, and dynamic target collection)
* **scaffolding_level:** High (Residual run validation gates and provenance checks)
* **promotion_scope:** Candidate residual runner artifact
* **forbidden_interpretations:** Friedmann recovery, classical-limit recovery, global cosmology claim, BMC validation
* **remaining_debts:** Comparison audits, stability audits under real perturbations, and convention debt retirement

### Sprint 11 + 11.1: Residual/Null Comparison Audit Artifact
* **artifact_kind:** Residual audit report
* **accepted_status:** Accepted after repairs (Sprint 11.1)
* **real_computation_level:** Toy (Interval-level proxy stability diagnostics, comparison audits checking structural honesty)
* **scaffolding_level:** High (Comparison audit validation gates, EBP warning checks)
* **promotion_scope:** Residual/null comparison audit artifact
* **forbidden_interpretations:** Friedmann recovery, classical-limit recovery, BMC superiority, scientific novelty, full BMC promotion
* **remaining_debts:** Real trajectory-level perturbation stability (instead of interval proxies), retired convention debts, and nontrivial physics cases

---

## 3. Report Schema Inventory

All schemas currently defined in the repository are cataloged below.

### 1. `bmc-report-v0.1`
* **owning package:** [report](file:///home/chaschel/Documents/go/bmc/internal/bmc/report)
* **CLI generator:** `ptw-bmc run`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** 
  - `out/bmc0a_plane.json`
  - `out/bmc0a_superposition_safe.json`
  - `out/bmc0a_superposition_node_probe.json`
* **classification:** Active computation / Audit scaffold
* **whether it can fail meaningfully:** Yes, if WdW residual or trajectory checks fail.
* **whether it blocks overclaims:** Yes, checks `final_truth_claim` and enforces `toy_analysis_only`.
* **whether it risks validation theater:** Low, but has basic threshold checks.

### 2. `bmc0a-superposition-robustness-v0.1`
* **owning package:** [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit)
* **CLI generator:** `ptw-bmc audit`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_superposition_robustness.json`
* **classification:** Audit scaffold
* **whether it can fail meaningfully:** Yes, if the step-size sweep or parameter perturbations fail validation checks.
* **whether it blocks overclaims:** Yes, enforces `promotion_gate` status to be blocked and checks `final_truth_claim`.
* **whether it risks validation theater:** Yes, runs fixed-step sweeps and narrow parameter bounds.

### 3. `bmc0a-clock-fragility-v0.1`
* **owning package:** [clockdiag](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag)
* **CLI generator:** `ptw-bmc diagnose-clock`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_clock_fragility.json`
* **classification:** Audit scaffold
* **whether it can fail meaningfully:** Yes, if clock monotonicity fails.
* **whether it blocks overclaims:** Yes, blocks promotion if fragility is too high.
* **whether it risks validation theater:** Medium, correlation and event counts are heuristic.

### 4. `bmc0a-clock-readiness-v0.1`
* **owning package:** [clockseg](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg)
* **CLI generator:** `ptw-bmc segment-clock`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** 
  - `out/bmc0a_clock_readiness.json`
  - `out/bmc0a_clock_readiness_no_points.json`
* **classification:** Audit scaffold
* **whether it can fail meaningfully:** Yes, if segments cannot be extracted.
* **whether it blocks overclaims:** Yes, blocks full promotion.
* **whether it risks validation theater:** Yes, checks segment count and monotonic sub-intervals without physical time metric.

### 5. `bmc0a-friedmann-spec-v0.1`
* **owning package:** [friedmannspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec)
* **CLI generator:** `ptw-bmc spec-friedmann`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_friedmann_spec.json`
* **classification:** Specification only / Deferral stub
* **whether it can fail meaningfully:** No, as it represents only the specification of the gate.
* **whether it blocks overclaims:** Yes, keeps the gate blocked.
* **whether it risks validation theater:** High, because it specifies gate conditions that are deferred.

### 6. `bmc0a-nullmodel-spec-v0.1`
* **owning package:** [nullspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec)
* **CLI generator:** `ptw-bmc spec-nullmodels`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_nullmodel_spec.json`
* **classification:** Specification only
* **whether it can fail meaningfully:** No, specification only.
* **whether it blocks overclaims:** Yes.
* **whether it risks validation theater:** High, plans models that are not computed.

### 7. `bmc0a-prior-art-boundary-v0.1`
* **owning package:** [priorart](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart)
* **CLI generator:** `ptw-bmc prior-art-boundary`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_prior_art_boundary.json`
* **classification:** Specification only
* **whether it can fail meaningfully:** Yes, if forbidden claims are found in text checks.
* **whether it blocks overclaims:** Yes, blocks if claims are violated.
* **whether it risks validation theater:** Medium, checks string/text matching.

### 8. `bmc0a-nullrun-v0.1`
* **owning package:** [nullrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun)
* **CLI generator:** `ptw-bmc run-nullmodels`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_nullrun.json`
* **classification:** Audit scaffold / Active computation
* **whether it can fail meaningfully:** Yes, if registered null models are missing or corrupt.
* **whether it blocks overclaims:** Yes, prevents beats-null claims.
* **whether it risks validation theater:** Yes, comparison metrics are calculated on placeholder/toy runs.

### 9. `bmc0a-local-residual-v0.1`
* **owning package:** [residualrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun)
* **CLI generator:** `ptw-bmc run-residuals`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_local_residual.json`
* **classification:** Active computation / Audit scaffold
* **whether it can fail meaningfully:** Yes, if point values are non-finite or monotonicity is violated.
* **whether it blocks overclaims:** Yes.
* **whether it risks validation theater:** Yes, candidate residual uses simplified analytic wave function.

### 10. `bmc0a-residual-audit-v0.1`
* **owning package:** [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit)
* **CLI generator:** `ptw-bmc audit-residuals`
* **validator route:** `ptw-bmc validate --report <file>`
* **summarizer route:** `ptw-bmc summarize --report <file>`
* **generated artifact path:** `out/bmc0a_residual_audit.json`
* **classification:** Audit scaffold
* **whether it can fail meaningfully:** Yes, if stability limits are violated.
* **whether it blocks overclaims:** Yes.
* **whether it risks validation theater:** High, stability check uses interval-level proxies instead of re-integrating trajectories.

---

## 4. CLI Subcommand Inventory

The subcommands of `ptw-bmc` implemented in [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go) are listed below.

### 1. `run`
* **profiles:** `bmc0a-plane`, `bmc0a-superposition-safe`, `bmc0a-superposition-node-probe`
* **source package:** [report](file:///home/chaschel/Documents/go/bmc/internal/bmc/report)
* **output schema:** `bmc-report-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_plane.json`, etc.)
* **action class:** Computes guidance trajectory, checks WdW residual, clock, and quantum potential, audits obstructions.
* **current accepted status:** Accepted plane-wave and superposition control artifact.
* **unknown-profile behavior:** Prints error and usage to stderr, then exits with code 1.

### 2. `validate`
* **profiles:** Dynamically matches `schema_version` inside the provided JSON file.
* **source package:** Switch maps to [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit), [clockdiag](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag), [clockseg](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg), [friedmannspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec), [nullspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec), [priorart](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart), [nullrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun), [residualrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun), [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit), or [report](file:///home/chaschel/Documents/go/bmc/internal/bmc/report).
* **output schema:** None (prints validation success/failure messages).
* **output path if standard:** Stdout / Stderr.
* **action class:** Validates (runs strict JSON parsing, enforces EBP 2.1 promotion constraints).
* **current accepted status:** Validates all Sprint 1-11 reports.
* **unknown-profile behavior:** Parses as a general report and validates via `report.Validate(rep)`.

### 3. `summarize`
* **profiles:** Dynamically matches `schema_version` inside the provided JSON file.
* **source package:** Same as `validate`.
* **output schema:** None (prints human-readable summary).
* **output path if standard:** Stdout.
* **action class:** Summarizes.
* **current accepted status:** Summarizes all Sprint 1-11 reports.
* **unknown-profile behavior:** Parses as a general report and summarizes via general layout logic.

### 4. `audit`
* **profiles:** `bmc0a-superposition-robustness`
* **source package:** [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit)
* **output schema:** `bmc0a-superposition-robustness-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_superposition_robustness.json`).
* **action class:** Audits.
* **current accepted status:** Accepted superposition robustness sweep.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 5. `diagnose-clock`
* **profiles:** `bmc0a-clock-fragility`
* **source package:** [clockdiag](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag)
* **output schema:** `bmc0a-clock-fragility-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_clock_fragility.json`).
* **action class:** Audits / Diagnoses.
* **current accepted status:** Accepted.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 6. `segment-clock`
* **profiles:** `bmc0a-clock-readiness`
* **source package:** [clockseg](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg)
* **output schema:** `bmc0a-clock-readiness-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_clock_readiness.json`).
* **action class:** Audits / Segments.
* **current accepted status:** Accepted.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 7. `spec-friedmann`
* **profiles:** `bmc0a-friedmann-spec`
* **source package:** [friedmannspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec)
* **output schema:** `bmc0a-friedmann-spec-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_friedmann_spec.json`).
* **action class:** Specifies only.
* **current accepted status:** Accepted.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 8. `spec-nullmodels`
* **profiles:** `bmc0a-nullmodel-spec`
* **source package:** [nullspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec)
* **output schema:** `bmc0a-nullmodel-spec-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_nullmodel_spec.json`).
* **action class:** Specifies only.
* **current accepted status:** Accepted.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 9. `prior-art-boundary`
* **profiles:** `bmc0a-prior-art-boundary`
* **source package:** [priorart](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart)
* **output schema:** `bmc0a-prior-art-boundary-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_prior_art_boundary.json`).
* **action class:** Specifies only.
* **current accepted status:** Accepted after repairs.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 10. `run-nullmodels`
* **profiles:** `bmc0a-nullrun`
* **source package:** [nullrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun)
* **output schema:** `bmc0a-nullrun-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_nullrun.json`).
* **action class:** Audits / Computes.
* **current accepted status:** Accepted after repairs.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 11. `run-residuals`
* **profiles:** `bmc0a-local-residual`
* **source package:** [residualrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun)
* **output schema:** `bmc0a-local-residual-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_local_residual.json`).
* **action class:** Audits / Computes.
* **current accepted status:** Accepted after repairs.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

### 12. `audit-residuals`
* **profiles:** `bmc0a-residual-audit`
* **source package:** [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit)
* **output schema:** `bmc0a-residual-audit-v0.1`
* **output path if standard:** User-specified via `--out` (usually `out/bmc0a_residual_audit.json`).
* **action class:** Audits.
* **current accepted status:** Accepted after repairs.
* **unknown-profile behavior:** Prints error to stderr, exits with code 1.

---

## 5. Package Inventory

Classification of all packages under `internal/bmc/...` and `BMC/BMC/...`.

| Package Name | Scope Classification | Description / Machinism Flag |
| :--- | :--- | :--- |
| [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit) | Audit/Report Scaffold | Run numerical robustness and convergence sweeps. **[Machinery]** |
| [clockdiag](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag) | Audit/Report Scaffold | Monotonicity fragility investigation. **[Machinery]** |
| [clockseg](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg) | Audit/Report Scaffold | Clock segmentation and sub-branch extraction. **[Machinery]** |
| [friedmannspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec) | Specification only | Metadata/schema design for the Friedmann check. **[Machinery]** |
| [guidance](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance) | Trajectory/Guidance computation | Implements the Euler and RK4 numerical steppers. |
| [invariant](file:///home/chaschel/Documents/go/bmc/internal/bmc/invariant) | Validation/Gating | Classical limit comparative checkers. |
| [model](file:///home/chaschel/Documents/go/bmc/internal/bmc/model) | Core physics data models | Defines parameters and state types. |
| [nullrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun) | Audit/Report Scaffold | Runs toy null models. **[Machinery]** |
| [nullspec](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullspec) | Specification only | Structure schemas for null model registry. **[Machinery]** |
| [obstruction](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction) | Validation/Gating | Detects singularities, node contacts, phase bounds. |
| [priorart](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart) | Prior-art/boundary note | Enforces EBP boundaries. **[Machinery]** |
| [qpotential](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential) | Constraint/residual computation | Finite difference quantum potential evaluator. |
| [report](file:///home/chaschel/Documents/go/bmc/internal/bmc/report) | Audit/Report Scaffold | Superposition and plane wave pipeline reports. **[Machinery]** |
| [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit) | Audit/Report Scaffold | Audits structural honesty and stability of comparisons. **[Machinery]** |
| [residualrun](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun) | Audit/Report Scaffold | Computes candidate local residuals. **[Machinery]** |
| [wave](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave) | Wavefunction model | Defines plane and superposition wavefunctions. |
| [wdw](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw) | Constraint/residual computation | Wheeler-DeWitt residual calculation. |
| [BMC/BMC](file:///home/chaschel/Documents/go/bmc/BMC/BMC) | Lean Policy Support | Models promotion rules as Lean theorem obligations. |

---

## 6. Test Inventory

Inventory of existing Go tests across the codebase, mapped to testing categories.

### Test Categories and Placements

1. **positive/control test:**
   - [wave/plane_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/plane_test.go) (verifies exact analytical wavefunction amplitude/phase values).
   - [guidance/trajectory_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/trajectory_test.go) (verifies basic trajectory path integration).
   - [model/params_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/model/params_test.go) (validates default parameters).

2. **negative/failure-detection test:**
   - [obstruction/node_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/node_test.go) (verifies that node contact correctly fails validation).
   - [audit/audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go) (verifies that final truth claims and `toy_analysis_only = false` fail validation).

3. **integration test:**
   - [report/report_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report_test.go) (runs the entire plane-wave and superposition pipeline end-to-end).

4. **CLI routing test:**
   - [priorart/priorart_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/priorart_test.go) (checks that CLI flag parsing rejects invalid command flags).

5. **validation/gate test:**
   - [nullrun/nullrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/nullrun_test.go) (verifies that corrupt null model run entries fail parsing).
   - [residualrun/residualrun_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualrun/residualrun_test.go) (checks schema validation and bounds).
   - [residualaudit/residualaudit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit/residualaudit_test.go) (verifies structural comparison validation).

6. **determinism test:**
   - [audit/audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go) (asserts that the convergence step-size sweep runs deterministically).

7. **numerical-convergence test:**
   - [audit/audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go) (measures endpoint drift of trajectories at decreasing step sizes).

8. **sensitivity test:**
   - [audit/audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go) (asserts thresholds sensitivity sweep under node-probe parameters).

9. **forbidden-phrase/anti-overclaim test:**
   - [priorart/priorart_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/priorart/priorart_test.go) (scans for case-insensitive final truth claims and forbidden phrases).

10. **Lean policy test:**
    - [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean) (lake build parses and validates safety policies in Lean).

---

### Failure-Detection Test Gaps

The following critical validation and physical failure-detection tests are **completely missing**:

* **Wrong wavefunction tests missing:** No test in the repository ensures that a physically incorrect wavefunction (e.g., an arbitrary, non-solution function) is rejected by the residual solver or constraint checks.
* **Constraint violation tests missing:** There are no tests verifying that a plane-wave configuration violating the coordinate relationship $k^2 = \omega^2$ is caught and fails the WdW check.
* **Euler vs RK4 convergence missing:** The code contains both `EulerStepper` and `RK4Stepper`, but no test compares their drift on the same wavefunction to audit numerical errors.
* **Step-size convergence error rate missing:** The step-size sweep checks endpoints but does **not** test that the error decreases with the expected mathematical order of RK4, $O(\Delta\lambda^4)$.
* **Node-threshold sensitivity missing:** The sweep sweeps thresholds, but there is no test verifying that tiny perturbations of a trajectory passing very close to a node correctly flip the safety gate status.
* **Phase-gradient h-sensitivity missing:** The finite difference step $h$ for wavefunction phase derivatives is hardcoded without auditing the numerical stability of derivatives under step size sweep.
* **Q near-node policy tests missing:** No test verifies that starting/passing near a node produces a domain-boundary error rather than a fake valid $Q=0$ (caused by the amplitude guard returning $0.0$).
* **Dimensional/unit consistency missing:** No checks or tests exist for unit coherence (dimensions of $a$, $\phi$, and coordinates).
* **Friedmann target comparison missing:** Because Friedmann checks are deferred, no tests check the trajectory against the actual classical Friedmann cosmological equations.

---

## 7. Generated Artifact Inventory

The JSON artifacts stored under [out/](file:///home/chaschel/Documents/go/bmc/out/) are inventoried below.

| Artifact Path | schema_version | Generator Command | Validation Command | Summary Command | Accepted Sprint | Track in Repo | Reproducible | Stale or Current |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| `out/bmc0a_plane.json` | `bmc-report-v0.1` | `ptw-bmc run --profile bmc0a-plane --out out/bmc0a_plane.json` | `ptw-bmc validate --report out/bmc0a_plane.json` | `ptw-bmc summarize --report out/bmc0a_plane.json` | Sprint 1 | Yes | Yes | Current |
| `out/bmc0a_superposition_safe.json` | `bmc-report-v0.1` | `ptw-bmc run --profile bmc0a-superposition-safe --out out/bmc0a_superposition_safe.json` | `ptw-bmc validate --report out/bmc0a_superposition_safe.json` | `ptw-bmc summarize --report out/bmc0a_superposition_safe.json` | Sprint 2 | Yes | Yes | Current |
| `out/bmc0a_superposition_node_probe.json` | `bmc-report-v0.1` | `ptw-bmc run --profile bmc0a-superposition-node-probe --out out/bmc0a_superposition_node_probe.json` | `ptw-bmc validate --report out/bmc0a_superposition_node_probe.json` | `ptw-bmc summarize --report out/bmc0a_superposition_node_probe.json` | Sprint 2 | Yes | Yes | Current |
| `out/bmc0a_superposition_robustness.json` | `bmc0a-superposition-robustness-v0.1` | `ptw-bmc audit --profile bmc0a-superposition-robustness --out out/bmc0a_superposition_robustness.json` | `ptw-bmc validate --report out/bmc0a_superposition_robustness.json` | `ptw-bmc summarize --report out/bmc0a_superposition_robustness.json` | Sprint 3 | Yes | Yes | Current |
| `out/bmc0a_clock_fragility.json` | `bmc0a-clock-fragility-v0.1` | `ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json` | `ptw-bmc validate --report out/bmc0a_clock_fragility.json` | `ptw-bmc summarize --report out/bmc0a_clock_fragility.json` | Sprint 4 | Yes | Yes | Current |
| `out/bmc0a_clock_readiness.json` | `bmc0a-clock-readiness-v0.1` | `ptw-bmc segment-clock --profile bmc0a-clock-readiness --out out/bmc0a_clock_readiness.json` | `ptw-bmc validate --report out/bmc0a_clock_readiness.json` | `ptw-bmc summarize --report out/bmc0a_clock_readiness.json` | Sprint 5 | Yes | Yes | Current |
| `out/bmc0a_clock_readiness_no_points.json` | `bmc0a-clock-readiness-v0.1` | Manual/Custom flag | `ptw-bmc validate --report out/bmc0a_clock_readiness_no_points.json` | `ptw-bmc summarize --report out/bmc0a_clock_readiness_no_points.json` | Sprint 5 | Yes | Yes | Current |
| `out/bmc0a_friedmann_spec.json` | `bmc0a-friedmann-spec-v0.1` | `ptw-bmc spec-friedmann --profile bmc0a-friedmann-spec --out out/bmc0a_friedmann_spec.json` | `ptw-bmc validate --report out/bmc0a_friedmann_spec.json` | `ptw-bmc summarize --report out/bmc0a_friedmann_spec.json` | Sprint 6 | Yes | Yes | Current |
| `out/bmc0a_nullmodel_spec.json` | `bmc0a-nullmodel-spec-v0.1` | `ptw-bmc spec-nullmodels --profile bmc0a-nullmodel-spec --out out/bmc0a_nullmodel_spec.json` | `ptw-bmc validate --report out/bmc0a_nullmodel_spec.json` | `ptw-bmc summarize --report out/bmc0a_nullmodel_spec.json` | Sprint 7 | Yes | Yes | Current |
| `out/bmc0a_prior_art_boundary.json` | `bmc0a-prior-art-boundary-v0.1` | `ptw-bmc prior-art-boundary --profile bmc0a-prior-art-boundary --out out/bmc0a_prior_art_boundary.json` | `ptw-bmc validate --report out/bmc0a_prior_art_boundary.json` | `ptw-bmc summarize --report out/bmc0a_prior_art_boundary.json` | Sprint 8.1 | Yes | Yes | Current |
| `out/bmc0a_nullrun.json` | `bmc0a-nullrun-v0.1` | `ptw-bmc run-nullmodels --profile bmc0a-nullrun --out out/bmc0a_nullrun.json` | `ptw-bmc validate --report out/bmc0a_nullrun.json` | `ptw-bmc summarize --report out/bmc0a_nullrun.json` | Sprint 9.1 | Yes | Yes | Current |
| `out/bmc0a_local_residual.json` | `bmc0a-local-residual-v0.1` | `ptw-bmc run-residuals --profile bmc0a-local-residual --out out/bmc0a_local_residual.json` | `ptw-bmc validate --report out/bmc0a_local_residual.json` | `ptw-bmc summarize --report out/bmc0a_local_residual.json` | Sprint 10.3 | Yes | Yes | Current |
| `out/bmc0a_residual_audit.json` | `bmc0a-residual-audit-v0.1` | `ptw-bmc audit-residuals --profile bmc0a-residual-audit --out out/bmc0a_residual_audit.json` | `ptw-bmc validate --report out/bmc0a_residual_audit.json` | `ptw-bmc summarize --report out/bmc0a_residual_audit.json` | Sprint 11.1 | Yes | Yes | Current |

---

## 8. Real Computation vs. Scaffolding Classification

This table classifies each feature block based on the level of real physical computation vs. administrative scaffolding.

| Feature Block | Classification | Technical Rationale / Blocker |
| :--- | :--- | :--- |
| **Plane-wave WdW residual** | `toy_computation` | Trivial mathematical check ($k^2 - \omega^2 = 0$) on analytically constant wave amplitudes. |
| **Plane-wave guidance trajectory** | `toy_computation` | Linear integrations with no physical curvature. |
| **Two-term superposition wavefunction** | `toy_computation` | Standard analytical quantum superposition, evaluated locally. |
| **Node obstruction detection** | `audit_scaffold` | Stops integrations when amplitude falls below a hard threshold. |
| **Clock fragility diagnostic** | `audit_scaffold` | Measures monotonicity statistics on discrete points. |
| **Clock segmentation/branch extraction**| `audit_scaffold` | Identifies local segments where clock is monotonic. |
| **Null model runner** | `audit_scaffold` | Executes placeholder null comparisons without actual physical separation. |
| **Candidate local residual runner** | `audit_scaffold` | Evaluates residuals point-by-point under explicit assumptions. |
| **Residual/null audit** | `audit_scaffold` | Confirms comparing structures are matching and honest. |
| **Prior-art boundary note** | `policy_scaffold` | Validates text strings to prevent overclaims. |
| **Friedmann residual spec** | `specification_only` | Gate schema and derivatives definition only. Deferred. |
| **Null model spec** | `specification_only` | Registry layout and list definitions. |
| **Lean policy files** | `policy_scaffold` | Translates Booleans in JSON reports to Lean theorems. |

- **Risk of Validation Theater:** High across the entire comparison audit stack. The complex JSON validation pipelines (e.g. `audit-residuals` checking `local-residual`, `nullrun`, etc.) check structural formatting, while the underlying physics remains a highly simplified, non-physical toy benchmark.

---

## 9. Deferral Stub Inventory

All deferred items, stubs, and unpaid debts in the repository are recorded below.

### 1. Friedmann Residual Evaluation
* **File:** [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go#L134-L139), [detect.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction/detect.go#L245)
* **Symbol/Function/Field:** `checks["friedmann_residual"]`, `ClassicalLimitFailure`
* **Deferred Item:** Comparison of numerical trajectory values against the classical Friedmann cosmology target equations.
* **Why Deferred:** Minisuperspace dynamics do not naturally recover classical cosmology without fine-tuning, and implementing the numerical solver comparison was deferred.
* **Retirement Condition:** An independent solver that compares the Bohmian trajectory against the classical Friedmann equations and calculates the residual.
* **Dependencies:** None. Current reports run successfully with this check flagged as `deferred`.

### 2. Classical Target Comparison
* **File:** [classical_limit.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/invariant/classical_limit.go)
* **Symbol/Function/Field:** `CheckClassicalLimit`
* **Deferred Item:** Verification of classical Einstein-Hilbert limits from the Bohmian guidance equations.
* **Why Deferred:** Relies on $Q \approx 0$ which only holds trivially for plane waves and fails in superpositions.
* **Retirement Condition:** Explicit numerical comparison against an actual classical gravity target simulation.
* **Dependencies:** None. Currently bypassed using a dummy check comparing `maxAbsQ` to a tolerance.

### 3. Null Model Comparisons Beyond Diagnostic Summaries
* **File:** [nullrun/runner.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/nullrun/runner.go)
* **Symbol/Function/Field:** `RunNullModels`
* **Deferred Item:** Full mathematical and physical comparisons of residuals against a comprehensive suite of null models.
* **Why Deferred:** Running complex models (e.g. Page-Wootters, LQC) requires implementing separate simulators.
* **Retirement Condition:** Integration of full solvers for the 7 registered null models.
* **Dependencies:** Current audits depend only on the registered status flags of the null models.

### 4. Observational/Empirical Benchmarks
* **File:** [ebp_debt](file:///home/chaschel/Documents/go/bmc/out/bmc0a_residual_audit.json#L303-L318)
* **Symbol/Function/Field:** `NeedLiteratureAudit`, `normalization_debt`, etc.
* **Deferred Item:** Auditing coordinates and parameters against actual cosmological observations (e.g. Planck satellite data).
* **Why Deferred:** Code is a pure toy benchmark.
* **Retirement Condition:** A database of observational parameters used to seed initial states.
* **Dependencies:** Bypassed entirely.

### 5. Lean Theorem Obligations Beyond Policy Booleans
* **File:** [Promotion.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Promotion.lean)
* **Symbol/Function/Field:** `reportPassesFullBMCToyGate`
* **Deferred Item:** Proving physical properties of the cosmology solver (e.g. stability, uniqueness, or WdW conservation) in Lean.
* **Why Deferred:** Current Lean theorems only assert boolean implications of JSON report structures.
* **Retirement Condition:** Nontrivial Lean theorems modeling the wave equation and proving bounds on trajectory integration.
* **Dependencies:** Bypassed; Lean build only checks the boolean promotion gate policy.

---

## 10. Failure-Detection Gap Inventory

The prioritizing of adversarial tests to close validation gaps is outlined below.

### 1. Wrong-Constraint Wavefunction Rejection
* **Priority:** P1 (Sprint P2)
* **Target Package:** [wdw](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw)
* **Minimal Test Fixture:** Pass a wavefunction $\Psi(\alpha, \phi) = \exp(i k \alpha^2)$ (which is not a solution to the WdW equation) into the central finite difference evaluator `FiniteDifferenceResidual` and assert that the residual is non-zero and fails.
* **Expected Behavior:** `FiniteDifferenceResidual` must fail the constraint validation check with a non-zero residual.

### 2. Coordinate Violated Control Check
* **Priority:** P1 (Sprint P2)
* **Target Package:** [wdw](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw)
* **Minimal Test Fixture:** Run `CheckResidual` with a plane wave where $k^2 \neq \omega^2$.
* **Expected Behavior:** Must fail validation and output an explicit blocker check.

### 3. Invalid Superposition Node Contact
* **Priority:** P1 (Sprint P2)
* **Target Package:** [obstruction](file:///home/chaschel/Documents/go/bmc/internal/bmc/obstruction)
* **Minimal Test Fixture:** Initialize a trajectory starting exactly inside a node region (e.g., $R < 1e-12$).
* **Expected Behavior:** Obstruction detection must trigger an explicit node contact obstruction failure immediately, short-circuiting integration.

### 4. Step-Size Stepper Drift Audit
* **Priority:** P2 (Sprint P2)
* **Target Package:** [guidance](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance)
* **Minimal Test Fixture:** Integrate the same safe superposition wavefunction using both `EulerStepper` and `RK4Stepper` with the same step size and initial configuration.
* **Expected Behavior:** Measures the numerical drift between the two steppers and asserts it matches the theoretical truncation error.

### 5. RK4 Convergence Sweep Verification
* **Priority:** P2 (Sprint P3)
* **Target Package:** [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit)
* **Minimal Test Fixture:** Evaluate trajectory drift under $\Delta\lambda \in \{0.1, 0.05, 0.025\}$ and verify that the endpoint scaling matches $O(\Delta\lambda^4)$.
* **Expected Behavior:** Fails validation if the scaling order is inconsistent with Runge-Kutta 4th order.

### 6. Perturbed Node Contact Sweeper
* **Priority:** P2 (Sprint P3)
* **Target Package:** [audit](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit)
* **Minimal Test Fixture:** Run trajectory integrations near nodes with tiny initial parameter offsets ($10^{-9}$).
* **Expected Behavior:** Identifies parameter sensitivity and marks unstable trajectories as failed.

### 7. Finite Difference Step-Size Sensitivity Sweep
* **Priority:** P3 (Sprint P3)
* **Target Package:** [qpotential](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential)
* **Minimal Test Fixture:** Compute Q along a trajectory using finite-difference steps $h \in \{1e-3, 1e-4, 1e-5\}$.
* **Expected Behavior:** Checks if the calculated Q values diverge or remain stable, validating the hardcoded step $h$.

### 8. Monotonicity Independent Clock Readiness check
* **Priority:** P3 (Sprint P3)
* **Target Package:** [clockseg](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg)
* **Minimal Test Fixture:** Pass a trajectory with clock variable $\phi$ that stops or reverses but has a well-defined physical time.
* **Expected Behavior:** Ensures the clock readiness check evaluates clock quality based on physical indicators rather than a boolean monotonicity flag.

---

## 11. Numerical-Analysis Debt Inventory

The repository contains several key numerical weaknesses:

* **Fixed Step Integrations:** Both Euler and RK4 steppers in [integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go) use fixed step sizes $\Delta\lambda$.
* **No Adaptive Step Sizing:** If a trajectory approaches a wavefunction node (where the phase gradient diverges), a fixed-step solver will jump across the singularity or produce non-physical trajectories without adjusting step size.
* **No Local Truncation Error Estimates:** The solvers do not evaluate local errors per step, meaning numerical drift can accumulate silently.
* **No Convergence Table:** While a step sweep is implemented, no formal convergence table (evaluating asymptotic error rates) is exported.
* **Hardcoded Finite Difference step:** The step size $h = 10^{-4}$ in [q.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/q.go#L20) and [residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go#L36) is hardcoded, introducing truncation and roundoff errors that are not audited.
* **Near-Node Division Clamping:** In [q.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/q.go#L15), the division by $R$ is clamped to zero when amplitude is small:
  ```go
  if rVal < 1e-12 {
      return 0.0
  }
  ```
  This returns $Q = 0$ near nodes, indicating perfect classical recovery instead of flagging a division-by-zero domain boundary failure.

---

## 12. Physics-Faithfulness Debt Inventory

The physical equations solved by the code have several severe simplifications and debts:

* **Factor Ordering Unmodeled:** The WdW operator $- \frac{\partial^2}{\partial\alpha^2} + \frac{\partial^2}{\partial\phi^2}$ assumes a specific factor ordering without auditing alternate options.
* **Minisuperspace Metric Conventions:** The kinetic terms assume flat $(+,-)$ signature without physical justification.
* **No Classical Target Equations:** The classical Friedmann equations are completely absent, making classical limit verification impossible.
* **No Unit / Dimensional Checks:** Coordinate values are dimensionless, and no scale factor conversion is audited.
* **No $Q$/Classical Ratio:** The ratio of quantum potential energy to classical kinetic energy is not computed.
* **No Nontrivial Potential-Bearing Model:** Only constant-potential (plane wave) or zero-potential superpositions are supported.
* **No Observational Benchmarks:** The code cannot accept or audit parameters against standard astronomical benchmarks.

---

## 13. Validation-Theater Risk Register

The risk registry catalogs locations where administrative validation protocols have replaced real physical verification.

| Risk ID | Risk Description | Code Evidence | Severity | Affected Packages | Recommended Remediation |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **VT-0001** | **Near-Node Q potential Clamp** | `q.go:L15` clamps $Q=0$ when amplitude is small. | **HIGH** | [qpotential](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential) | Raise a domain-boundary error or return NaN instead of $0.0$. |
| **VT-0002** | **Analytic Tautology Authority** | Plane-wave WdW check evaluates a simple formula ($k^2 - \omega^2$) instead of solving the wave equation. | **MEDIUM** | [wdw](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw) | Enforce numerical double-integration of the wave equation over states. |
| **VT-0003** | **Decorative Comparison Audits** | The residual comparisons pass validation as long as keys are present, regardless of physics. | **MEDIUM** | [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit) | Validate that the numerical difference is smaller than a physics-backed bound. |
| **VT-0004** | **Conditional promotion gates** | Promotion checks are written as hardcoded Go conditionals rather than theorem-backed mathematical guarantees. | **MEDIUM** | [cmd/ptw-bmc](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc) | Generate execution traces that can be verified directly in Lean. |
| **VT-0005** | **Decorative stability sweep** | Stability audits use proxy rescalings on the interval level instead of re-integrating trajectories under perturbation. | **HIGH** | [residualaudit](file:///home/chaschel/Documents/go/bmc/internal/bmc/residualaudit) | Integrate perturbed trajectories directly inside the stability audit tool. |

---

## 14. Freeze Recommendations

To protect the integrity of the codebase, the following freeze rules must be enforced:

1. **Freeze New Report Schemas:** Do not define any new JSON schemas or report versions. The existing `v0.1` schemas are sufficient for toy auditing.
2. **Freeze New Audit Layers:** Do not add additional audit packages or CLI reporting subcommands.
3. **Freeze Promotion Language:** Do not change the status of any promotion gates from `blocked` or `contested`.
4. **Freeze Sprint 1-11 Artifacts:** Mark the existing `out/*.json` files as historical controls.
5. **Restrict Work to Remediation:** Only allow changes that implement failure-detection tests or resolve numerical-analysis debts.

---

## 15. Next Remediation Tickets

The prioritized engineering tickets for the next remediation phase are detailed below.

### BMC-POST-0001: Constraint Violation Detection
* **Priority:** P1
* **Scope:** Add negative tests to [wdw](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw) and [report](file:///home/chaschel/Documents/go/bmc/internal/bmc/report).
* **Likely Files:** [wdw/residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go), [audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go).
* **Success Criteria:** Validation must fail when running reports with a wavefunction that does not solve the WdW equation, or a plane wave where $k^2 \neq \omega^2$.
* **Forbidden Claims:** Do not claim full WdW solver validation.
* **EBP Debts Affected:** `needToyCheck`, `needObstruction`.

### BMC-POST-0002: Independent Numerical WdW Residual Evaluator
* **Priority:** P1
* **Scope:** Implement a solver that evaluates the second-order wave equation derivatives numerically across all points in a trajectory.
* **Likely Files:** [wdw/residual.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wdw/residual.go).
* **Success Criteria:** Disables the analytic fallback ($k^2 - \omega^2 = 0$) and calculates finite-difference residuals on the grid.
* **Forbidden Claims:** Do not claim analytical completeness.
* **EBP Debts Affected:** `needInvariant`.

### BMC-POST-0003: Euler/RK4 and dt Convergence Audit
* **Priority:** P2
* **Scope:** Compare numerical drift between Euler and RK4 steppers on identical superposition profiles.
* **Likely Files:** [guidance/integrate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/integrate.go), [audit/convergence.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/convergence.go).
* **Success Criteria:** Outputs error ratios and verifies that RK4 error scales as $O(\Delta\lambda^4)$ asymptotically.
* **Forbidden Claims:** Do not claim global numerical stability.
* **EBP Debts Affected:** `needToyCheck`.

### BMC-POST-0004: Node Domain Boundary Policy for Q and Velocity
* **Priority:** P1
* **Scope:** Remove the division-by-zero clamping in Q potential evaluation.
* **Likely Files:** [qpotential/q.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/qpotential/q.go), [guidance/velocity.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/guidance/velocity.go).
* **Success Criteria:** Evaluating Q or velocity at points where amplitude $R < \text{NodeThresh}$ must return NaN or trigger a validation error, rather than returning a clamped $0.0$.
* **Forbidden Claims:** Do not claim node-passing capability.
* **EBP Debts Affected:** `needObstruction`, `needInvariant`.

### BMC-POST-0005: Phase Gradient h-Sensitivity Audit
* **Priority:** P3
* **Scope:** Implement a sweep of the finite difference parameter $h$ to audit derivative sensitivity.
* **Likely Files:** [wave/phase.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/wave/phase.go), [audit/audit_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/audit/audit_test.go).
* **Success Criteria:** The sweep detects if the derivative values diverge or become unstable.
* **Forbidden Claims:** Do not claim derivative convergence.
* **EBP Debts Affected:** `needInvariant`.

### BMC-POST-0006: Local Clock-Readiness Gate Integration
* **Priority:** P2
* **Scope:** Integrate the local readiness check directly into the promotion validation command.
* **Likely Files:** [clockseg/validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockseg/validate.go), [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go).
* **Success Criteria:** Bypassing clock readiness validation fails the main `ptw-bmc validate` command.
* **Forbidden Claims:** Do not claim physical lapse solution.
* **EBP Debts Affected:** `needMap`, `needInvariant`.

### BMC-POST-0007: Friedmann Residual Specification Split
* **Priority:** P2
* **Scope:** Separate the Friedmann specification from numerical simulation.
* **Likely Files:** [friedmannspec/residual_spec.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/friedmannspec/residual_spec.go).
* **Success Criteria:** Cleans up unused stubs and isolates planned gates.
* **Forbidden Claims:** Do not claim Friedmann recovery.
* **EBP Debts Affected:** `needMap`.

### BMC-POST-0008: Schema Inventory and Shared Envelope Plan
* **Priority:** P3
* **Scope:** Design a unified envelope for all ten report schemas.
* **Likely Files:** [report/report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/report/report.go).
* **Success Criteria:** Reduces duplicate validator logic across CLI subcommands.
* **Forbidden Claims:** None.
* **EBP Debts Affected:** `needToyCheck`.

---

## 16. EBP Status Summary

The current EBP 2.1 status coordinates for the BMC-0.1 workspace are summarized below:

```text
promotion_status: postmortem_inventory_only
full_bmc_toy_gate: blocked
containsFinalTruthClaim: absent
needFaithfulnessReview: contested
needToyCheck: partial
needNullModel: partial
needNumericalErrorAudit: unpaid
needConstraintViolationTests: unpaid
needNontrivialPhysicsCase: unpaid
```
