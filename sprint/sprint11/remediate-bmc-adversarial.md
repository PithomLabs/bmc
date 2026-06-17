# BMC Codebase Remediation Plan

Current status: in_progress.

Approach: three phases — Postmortem, Constrain, Refine.

## Phase 1 — Postmortem
Purpose: freeze the current surface area and prevent further drift before refactoring.

P1.1 Inventory all report schemas and CLI subcommands; map each to an actual implementation path or deferral stub. Document which are active debt vs. pure theater.
P1.2 Inventory all tests. Mark each as single-profile validation, negative-case missing, integration missing, or speed regression.
P1.3 Capture the smallest reproducible trace of the adversarial failures: node-threshold sensitivity, Euler vs RK4 drift, ClassicalLimit check independence from inner-potential computation, and no Friedmann residual despite promotion logic referencing it.
Deliverable: a checkpoint inventory saved under `docs/sprint` or `docs/postmortem` naming.

## Phase 2 — Constrain
Purpose: reduce complexity until the pipeline is auditable end-to-end.

P2.1 Freeze the plane-wave and safe-superposition profiles as the only `run` profiles until BMC-0A plane wave and superposition both produce a non-trivial trajectory and a real comparison target.
P2.2 Remove plane-wave WdW analytic-residual tautology from the primary authority path; replace with a solver that can also score an intentionally wrong wavefunction so the residual test can detect violations.
P2.3 Enforce RK4 as the default stepper with reproducible fixed dt; capture Euler drift and error as an audit subreport, not a silent alternative.
P2.4 Replace Q near-node clamp with an explicit node-domain rejection that records “divergent in limit” instead of silently returning 0.0.
Deliverable: updated defaults in `model/params.go` and `guidance/integrate.go`, plus one failing adversarial test.

## Phase 3 — Refine
Purpose: restore the missing physics and tighten validation logic.

P3.1 Implement an effective Friedmann residual in `(α, φ)`: compare guidance-derived ȧ, ä against the classical Friedmann right-hand side; expose tolerance and classical-limit region as parameters.
P3.2 Replace global strict monotonicity with a clock-readiness gate based on local monotonicity windows, mirroring the existing `clockseg` package but promoted into the main report path.
P3.3 Add adversarial test cases: wrong-constraint wavefunction, large-phase-gradient superposition, and node-threshold sweeps; require detectable false/fail paths.
P3.4 Reduce schema proliferation by collapsing specialized schemas into the core report plus structured audit sections; keep `bmc-report-v0.1` as the single truth schema unless physics changes require a version bump.
Deliverable: improved `internal/bmc/friedmannspec`, `internal/bmc/clockseg`, new tests under `internal/bmc/wave`, `internal/bmc/report`.

## Risk / dependency notes
P2 and P3.1 may change report contents; keep validation and summarizer behavior symmetric via the existing `schema_version` branch in `cmd/ptw-bmc/main.go`.
Keep EBP-style non-claims and promotion-blocking logic unchanged until report contents change; the schema contract is stronger than the La