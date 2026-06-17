# BMC Sprint 4: Clock-Monotonicity Fragility Investigation

Investigate why φ-clock monotonicity fails under four of nine safe-superposition parameter perturbations identified in Sprint 3, and classify whether each failure is a real trajectory obstruction, a clock-choice limitation, a numerical artifact, or expected Bohmian behavior in this toy model.

This is an **obstruction diagnosis sprint**, not a theory-promotion sprint.

## User Review Required

> [!IMPORTANT]
> **Promotion and Physics Boundaries**:
> The maximum allowed Sprint 4 promotion is `promoted_clock_fragility_diagnostic_artifact_after_repairs`.
> This sprint does NOT imply or prove:
> - Bohmian mechanics correctness
> - Quantum gravity
> - Problem of time resolution
> - Spacetime emergence
> - Friedmann recovery
> - That φ is or is not a valid physical clock in full quantum cosmology
>
> The Friedmann residual check remains deferred, the full BMC toy gate remains blocked, and `clock_choice_debt` is explicitly activated.

> [!WARNING]
> **A failed φ-clock is valuable diagnostic information.**
> A valid Bohmian trajectory with a nonmonotonic φ is fundamentally different from an invalid trajectory (node contact, non-finite velocity). Sprint 4's most important output is preserving this distinction in the report schema.

---

## Open Questions

> [!IMPORTANT]
> **CLI Subcommand Choice**: I propose a dedicated `diagnose-clock` subcommand rather than reusing `audit`. Rationale: the `audit` subcommand runs the Sprint 3 robustness sweeps and produces a `RobustnessReport`. The clock diagnostic has a different schema (`bmc0a-clock-fragility-v0.1`), different report structure, and different validation logic. A separate subcommand prevents ambiguity in profile routing and makes the CLI self-documenting. The `validate` and `summarize` subcommands will route to clock-fragility-specific logic based on schema version, following the existing pattern.
>
> Is `diagnose-clock` acceptable, or do you prefer extending `audit --profile bmc0a-clock-fragility`?

---

## Proposed Changes

### Clock Diagnostic Package (`internal/bmc/clockdiag/`)

New package. All files are `[NEW]`.

---

#### [NEW] [events.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/events.go)

Defines the `ClockEvent` struct and detection logic.

```go
type ClockEvent struct {
    Index                  int      `json:"index"`
    Lambda                 float64  `json:"lambda"`
    Alpha                  float64  `json:"alpha"`
    Phi                    float64  `json:"phi"`
    DPhiDLambda            float64  `json:"dphi_dlambda"`
    DAlphaDLambda          float64  `json:"dalpha_dlambda"`
    AmplitudeR             float64  `json:"amplitude_r"`
    QValue                 *float64 `json:"q_value"`
    QStatus                string   `json:"q_status,omitempty"`
    QReason                string   `json:"q_reason,omitempty"`
    PhaseGradientMagnitude *float64 `json:"phase_gradient_magnitude"`
    PhaseGradientStatus    string   `json:"phase_gradient_status,omitempty"`
    PhaseGradientReason    string   `json:"phase_gradient_reason,omitempty"`
    NearNode               bool     `json:"near_node"`
    EventKind              string   `json:"event_kind"`
    Severity               string   `json:"severity"`
}
```

`EventKind` values: `"sign_change"`, `"near_zero"`, `"direction_reversal"`, `"monotonicity_failure"`.

`Severity` values: `"info"`, `"warning"`, `"diagnostic"`.

Implements `DetectClockEvents(traj model.Trajectory, wf wave.WaveFunction, params model.SuperpositionParams) []ClockEvent`:
- Iterates consecutive trajectory point pairs.
- Computes `dφ/dλ` and `dα/dλ` as finite differences between consecutive points: `(φ[i] - φ[i-1]) / (λ[i] - λ[i-1])`.
- Emits events when:
  - `dφ/dλ` changes sign between consecutive steps (`sign_change`).
  - `|dφ/dλ| < 1e-10` (`near_zero`).
  - φ step direction reverses relative to the dominant direction established in the first 10 steps (`direction_reversal`).
  - The cumulative monotonicity invariant breaks (`monotonicity_failure`).
- For each event, records the amplitude R, Q value (via `qpotential` if away from node, else `nil` with status/reason), phase gradient magnitude (if away from node), and `near_node` flag (`R < NodeThresh`).

---

#### [NEW] [correlation.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/correlation.go)

Descriptive correlation/comparison between clock failures and trajectory metrics.

```go
type CorrelationSummary struct {
    ParameterSet          string   `json:"parameter_set"`
    C2Real                float64  `json:"c2_real"`
    K2                    float64  `json:"k2"`
    Omega2                float64  `json:"omega2"`
    ClockMonotonic        bool     `json:"clock_monotonic"`
    NumClockEvents        int      `json:"num_clock_events"`
    MinAmplitudeR         *float64 `json:"min_amplitude_r"`
    MinAmplitudeRStatus   string   `json:"min_amplitude_r_status,omitempty"`
    MinAmplitudeRReason   string   `json:"min_amplitude_r_reason,omitempty"`
    MaxAbsQ               *float64 `json:"max_abs_q"`
    MaxAbsQStatus         string   `json:"max_abs_q_status,omitempty"`
    MaxAbsQReason         string   `json:"max_abs_q_reason,omitempty"`
    MaxPhaseGradMagnitude *float64 `json:"max_phase_gradient_magnitude"`
    MaxPhaseGradStatus    string   `json:"max_phase_gradient_status,omitempty"`
    MaxPhaseGradReason    string   `json:"max_phase_gradient_reason,omitempty"`
    MinDistToNodeThresh   *float64 `json:"min_dist_to_node_thresh"`
    MinDistToNodeStatus   string   `json:"min_dist_to_node_status,omitempty"`
    MinDistToNodeReason   string   `json:"min_dist_to_node_reason,omitempty"`
}
```

Implements `ComputeCorrelations(safeParams model.SuperpositionParams, failedConfigs []FailedPerturbationConfig) ([]CorrelationSummary, error)`:
- Runs all nine Sprint 3 perturbation parameter sets (not just the four failures).
- For each, integrates the trajectory, detects clock events, computes min R, max |Q|, max phase gradient, min distance to `NodeThresh`.
- Returns deterministic comparison table. Correlation is descriptive only—no statistical tests.

---

#### [NEW] [diagnose.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/diagnose.go)

Core diagnostic logic: re-runs the four failed perturbations under step refinement and checks alternative clock candidates.

```go
type FailedPerturbationConfig struct {
    C2Real float64
    K2     float64
    Omega2 float64
}

type StepRefinementResult struct {
    C2Real          float64            `json:"c2_real"`
    K2              float64            `json:"k2"`
    Omega2          float64            `json:"omega2"`
    StepSize        float64            `json:"step_size"`
    Steps           int                `json:"steps"`
    PhiMonotonic    bool               `json:"phi_monotonic"`
    AlphaMonotonic  bool               `json:"alpha_monotonic"`
    NumClockEvents  int                `json:"num_clock_events"`
    TrajectoryValid bool               `json:"trajectory_valid"`
    ClockEvents     []ClockEvent       `json:"clock_events"`
}

type AlternativeClockSummary struct {
    PhiMonotonic     model.CheckStatus `json:"phi_monotonic"`
    AlphaMonotonic   model.CheckStatus `json:"alpha_monotonic"`
    BothMonotonic    bool              `json:"both_monotonic"`
    NeitherMonotonic bool              `json:"neither_monotonic"`
    ClockChoiceDebt  string            `json:"clock_choice_debt"`
}

type TrajectoryValiditySummary struct {
    TrajectoryValid       model.CheckStatus `json:"trajectory_valid"`
    PhiClockValid         model.CheckStatus `json:"phi_clock_valid"`
    DistinctionPreserved  bool              `json:"distinction_preserved"`
    Reason                string            `json:"reason"`
}
```

Implements:
- `RunStepRefinementRechecks(safeParams model.SuperpositionParams, failedConfigs []FailedPerturbationConfig) ([]StepRefinementResult, error)`:
  - For each of the four failed perturbations, runs with `T=10.0` and `dt ∈ {0.05, 0.025, 0.0125}` (steps: 200, 400, 800).
  - Returns 12 results (4 configs × 3 step sizes), deterministic order.
  - Each result includes φ-monotonicity, α-monotonicity, clock events, and trajectory validity (no NaN/Inf in valid trajectory points).

- `ComputeAlternativeClockSummary(refinementResults []StepRefinementResult) AlternativeClockSummary`:
  - Aggregates across all step refinement results.
  - `phi_monotonic`: `pass` if all runs have monotonic φ, `fail` if any fail, `contested` if results are mixed across step sizes.
  - `alpha_monotonic`: same logic for α.
  - `clock_choice_debt`: always `"active"`.

- `ComputeTrajectoryValiditySummary(refinementResults []StepRefinementResult) TrajectoryValiditySummary`:
  - `trajectory_valid`: `pass` if all trajectories are finite and well-formed (no NaN/Inf).
  - `phi_clock_valid`: `pass` if φ is monotonic across all runs, `fail` otherwise.
  - `distinction_preserved`: `true` iff the report can distinguish the two states. Specifically: if `trajectory_valid == pass` and `phi_clock_valid == fail`, then the distinction is preserved. If `trajectory_valid == fail`, then the distinction is also preserved (trajectory itself is invalid). Only set to `false` if the code cannot tell them apart.

---

#### [NEW] [report.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/report.go)

Defines `ClockFragilityReport`, read/write/summarize functions.

```go
type ClockFragilityReport struct {
    SchemaVersion              string                    `json:"schema_version"`
    ToyAnalysisOnly            bool                      `json:"toy_analysis_only"`
    FinalTruthClaim            bool                      `json:"final_truth_claim"`
    SourceArtifacts            []string                  `json:"source_artifacts"`
    DiagnosticKind             string                    `json:"diagnostic_kind"`
    FailedPerturbationRechecks []StepRefinementResult    `json:"failed_perturbation_rechecks"`
    ClockEvents                []ClockEvent              `json:"clock_events"`
    CorrelationSummary         []CorrelationSummary      `json:"correlation_summary"`
    AlternativeClockSummary    AlternativeClockSummary   `json:"alternative_clock_summary"`
    TrajectoryValiditySummary  TrajectoryValiditySummary `json:"trajectory_validity_summary"`
    TechnicalGate              TechnicalGate             `json:"technical_gate"`
    DiagnosticOutcome          string                    `json:"diagnostic_outcome"`
    PromotionGate              report.PromotionGate      `json:"promotion_gate"`
    EbpDebt                    ClockEbpDebt              `json:"ebp_debt"`
    Warnings                   []string                  `json:"warnings"`
}

type TechnicalGate struct {
    Name   string            `json:"name"`
    Status model.CheckStatus `json:"status"`
    Reason string            `json:"reason"`
}

type ClockEbpDebt struct {
    NeedMap                string `json:"needMap"`
    NeedInvariant          string `json:"needInvariant"`
    NeedToyCheck           string `json:"needToyCheck"`
    NeedNullModel          string `json:"needNullModel"`
    NeedObstruction        string `json:"needObstruction"`
    NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
    ClockChoiceDebt        string `json:"clock_choice_debt"`
}
```

- `SchemaVersion`: `"bmc0a-clock-fragility-v0.1"`
- `DiagnosticKind`: `"clock_monotonicity_fragility"`
- `DiagnosticOutcome`: computed as:
  - `"clock_stable"`: all failed perturbations become monotonic under step refinement (failures were numerical artifacts).
  - `"clock_fragile"`: failures persist under refinement for at least one config.
  - `"mixed"`: some configs stabilize, others persist.
  - `"contested"`: inconclusive (shouldn't happen with deterministic sweeps, but preserved for completeness).
- `TechnicalGate`: computed status:
  - `"pass"`: all planned diagnostics ran, reports validated, no NaN/Inf hidden, distinction preserved.
  - `"fail"`: diagnostic execution failed or distinction was not preserved.
  - `"contested"`: edge case.

Implements:
- `ReadClockFragilityReport(path string) (*ClockFragilityReport, error)`: strict JSON decoding with `DisallowUnknownFields()`.
- `WriteJSON(rep *ClockFragilityReport, path string) error`: pretty-printed JSON.
- `SummarizeClockFragilityReport(rep *ClockFragilityReport)`: human-readable summary.
- `GenerateClockFragilityReport(safeParams model.SuperpositionParams) (*ClockFragilityReport, error)`: orchestrates all diagnostics:
  1. Runs `DetectClockEvents` on each failed perturbation at default step size.
  2. Runs `RunStepRefinementRechecks` for the four failed configs.
  3. Runs `ComputeCorrelations` across all nine perturbation configs.
  4. Computes `AlternativeClockSummary` and `TrajectoryValiditySummary`.
  5. Computes `DiagnosticOutcome` and `TechnicalGate`.
  6. Assembles full report with EBP debt and warnings.

---

#### [NEW] [validate.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/validate.go)

Strict validation of `ClockFragilityReport`.

Rules:
- Reject `final_truth_claim = true`.
- Reject `toy_analysis_only = false`.
- Reject `schema_version ≠ "bmc0a-clock-fragility-v0.1"`.
- Reject `diagnostic_kind ≠ "clock_monotonicity_fragility"`.
- Reject `promotion_gate.status ≠ "blocked"`.
- Reject missing/empty `technical_gate.status` or `technical_gate.reason`.
- Reject non-finite numeric values in all sweep results and clock events.
- Reject unavailable numeric values (`nil` or `-1.0`) unless paired with explicit `status`/`reason`.
- Reject empty `failed_perturbation_rechecks`.
- Reject empty `clock_events` if `diagnostic_outcome == "clock_fragile"` (fragile outcome requires evidence).
- Reject invalid `diagnostic_outcome` values (must be one of: `clock_stable`, `clock_fragile`, `mixed`, `contested`).
- Reject `ebp_debt.clock_choice_debt ≠ "active"` (clock choice debt must remain active).
- Uses `validateOptionalMetric` pattern from Sprint 3 data-integrity repair for all `*float64` fields.

---

### CLI Integration (`cmd/ptw-bmc/`)

#### [MODIFY] [main.go](file:///home/chaschel/Documents/go/bmc/cmd/ptw-bmc/main.go)

- Add `diagnose-clock` subcommand:
  ```bash
  ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json
  ```
- Update `validate` to route on schema version `"bmc0a-clock-fragility-v0.1"` → `clockdiag.ValidateClockFragilityReport`.
- Update `summarize` to route on schema version `"bmc0a-clock-fragility-v0.1"` → `clockdiag.SummarizeClockFragilityReport`.
- Update `printUsage()` to list the new subcommand.

---

### Lean Policy/Safety Gates (`BMC/`)

#### [NEW] [ClockFragility.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/ClockFragility.lean)

Policy-only. No clock physics formalization.

```lean
import BMC.ToyReport

structure BMCClockFragilityReport where
  toyAnalysisOnly      : Bool
  finalTruthClaim      : Bool
  technicalGatePassed  : Bool
  friedmannDeferred    : Bool
  fullBMCBlocked       : Bool
  clockChoiceDebtActive : Bool
  deriving Repr

def reportPassesBMC0AClockFragilityDiagnosticGate (r : BMCClockFragilityReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.technicalGatePassed &&
  r.friedmannDeferred &&
  r.fullBMCBlocked &&
  r.clockChoiceDebtActive

def reportPassesFullBMCForClockFragility (r : BMCClockFragilityReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.technicalGatePassed &&
  !r.friedmannDeferred &&
  !r.fullBMCBlocked

-- Policy safety theorems

theorem clock_fragility_requires_toy_only ...
theorem clock_fragility_blocks_final_truth ...
theorem clock_fragility_requires_friedmann_deferred ...
theorem clock_fragility_does_not_imply_full_bmc ...
theorem clock_fragility_keeps_clock_choice_debt_active ...
```

Following the exact pattern established in [Robustness.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC/Robustness.lean): structure, gate decider, full-BMC decider, five safety theorems, witness, and witness-passes/fails-gate checks.

#### [MODIFY] [BMC.lean](file:///home/chaschel/Documents/go/bmc/BMC/BMC.lean)

Add `import BMC.ClockFragility`.

---

## Verification Plan

### Automated Tests

All tests in [NEW] [clockdiag_test.go](file:///home/chaschel/Documents/go/bmc/internal/bmc/clockdiag/clockdiag_test.go):

| Test | Verifies |
|------|----------|
| `TestClockFragilityReportValidation` | Generated report passes validation; rejects `final_truth_claim=true`, `toy_analysis_only=false`, bad schema version, bad diagnostic kind, unblocked promotion gate, missing technical gate reason |
| `TestClockEventsDetectedInFailedPerturbations` | For each of the 4 failed perturbation configs at default step size, at least one clock event is detected with `event_kind ∈ {sign_change, direction_reversal, monotonicity_failure}` |
| `TestFailedPerturbationsRecheckedUnderStepRefinement` | Exactly 12 results (4 configs × 3 step sizes); all step sizes and config values are deterministic and correctly ordered; trajectory validity is reported for each |
| `TestTrajectoryValidityDistinguishedFromClockValidity` | When `trajectory_valid = pass` and `phi_clock_valid = fail`, `distinction_preserved = true`. When both fail, `distinction_preserved = true` (trajectory invalidity is its own diagnosis) |
| `TestAlternativeClockSummary` | α-monotonicity status is computed; `clock_choice_debt = "active"` always |
| `TestClockFragilityRejectsFinalTruthClaim` | Setting `final_truth_claim=true` causes validation failure |
| `TestClockFragilityRejectsNonfiniteMetrics` | Injecting `NaN` into clock event fields causes validation failure |
| `TestClockFragilityDeterministicJSON` | Generating the report twice produces byte-identical JSON output |

### Manual Verification

```bash
go test ./...
go build -buildvcs=false ./cmd/ptw-bmc
./ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json
./ptw-bmc validate --report out/bmc0a_clock_fragility.json
./ptw-bmc summarize --report out/bmc0a_clock_fragility.json
cd BMC && /home/chaschel/.elan/bin/lake build
```

---

## Sprint 4 Plan (Structured)

```json
{
  "summary": "Investigate why φ-clock monotonicity fails under four safe-superposition parameter perturbations. Classify each failure as real obstruction, clock-choice limitation, numerical artifact, or expected Bohmian behavior. Create a new clockdiag package with event detection, step-refinement rechecks, correlation audit, alternative clock analysis, and trajectory-vs-clock validity distinction.",
  "proposed_actions": [
    "Create internal/bmc/clockdiag/ package with events.go, correlation.go, diagnose.go, report.go, validate.go",
    "Add diagnose-clock CLI subcommand",
    "Route validate/summarize for clock-fragility schema version",
    "Create ClockFragility.lean with policy/safety contracts",
    "Import ClockFragility from BMC.lean",
    "Add comprehensive unit tests in clockdiag_test.go"
  ],
  "files_to_add": [
    "internal/bmc/clockdiag/events.go",
    "internal/bmc/clockdiag/correlation.go",
    "internal/bmc/clockdiag/diagnose.go",
    "internal/bmc/clockdiag/report.go",
    "internal/bmc/clockdiag/validate.go",
    "internal/bmc/clockdiag/clockdiag_test.go",
    "BMC/BMC/ClockFragility.lean"
  ],
  "files_to_modify": [
    "cmd/ptw-bmc/main.go",
    "BMC/BMC.lean"
  ],
  "test_plan": [
    "TestClockFragilityReportValidation",
    "TestClockEventsDetectedInFailedPerturbations",
    "TestFailedPerturbationsRecheckedUnderStepRefinement",
    "TestTrajectoryValidityDistinguishedFromClockValidity",
    "TestAlternativeClockSummary",
    "TestClockFragilityRejectsFinalTruthClaim",
    "TestClockFragilityRejectsNonfiniteMetrics",
    "TestClockFragilityDeterministicJSON"
  ],
  "cli_plan": [
    "ptw-bmc diagnose-clock --profile bmc0a-clock-fragility --out out/bmc0a_clock_fragility.json",
    "ptw-bmc validate --report out/bmc0a_clock_fragility.json (routed via schema_version)",
    "ptw-bmc summarize --report out/bmc0a_clock_fragility.json (routed via schema_version)"
  ],
  "lean_plan": [
    "BMC/BMC/ClockFragility.lean: structure, diagnostic gate decider, full-BMC decider, 5 safety theorems, witness, witness validation",
    "BMC/BMC.lean: add import BMC.ClockFragility"
  ],
  "assumptions": [
    "The four failed perturbation configs from Sprint 3 are stable and reproducible",
    "Clock events can be detected via consecutive-point finite differences without requiring sub-step interpolation",
    "Describing monotonicity of α as an alternative clock candidate is sufficient diagnostic; no clock replacement is performed",
    "The WdW constraint omega2^2 = k2^2 is preserved for all perturbation configs (omega2 = -k2 ensures this)"
  ],
  "proof_obligations": [
    "clock_fragility_requires_toy_only",
    "clock_fragility_blocks_final_truth",
    "clock_fragility_requires_friedmann_deferred",
    "clock_fragility_does_not_imply_full_bmc",
    "clock_fragility_keeps_clock_choice_debt_active"
  ],
  "null_models": [
    "No new null models planned. Sprint 4 is a diagnostic of existing behavior."
  ],
  "risks": [
    "All four perturbation failures may resolve under step refinement, making the diagnostic outcome 'clock_stable' rather than 'clock_fragile'. This is still a valid and informative outcome.",
    "If α is also nonmonotonic where φ fails, neither variable is a good internal clock. This creates stronger clock_choice_debt but does not invalidate the trajectory.",
    "Clock event detection via finite differences may miss sub-step sign changes. This is acknowledged as a limitation of the discrete sampling; sub-step interpolation is out of scope."
  ],
  "human_review_questions": [
    "Is 'diagnose-clock' the preferred CLI subcommand name, or should clock fragility be a profile under the existing 'audit' subcommand?",
    "Should the correlation summary include all 9 perturbation configs (including 5 that pass), or only the 4 that fail? The plan includes all 9 for descriptive comparison.",
    "Is the near-zero dφ/dλ threshold of 1e-10 appropriate, or should it be parameterized?"
  ],
  "ebp_debt_status": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "contested",
    "clock_choice_debt": "active",
    "containsFinalTruthClaim": "false",
    "LeanVerification": "active (policy/safety only, no clock physics formalization)"
  },
  "promotion_status": {
    "sprint4_clock_fragility_diagnostic": "promoted_clock_fragility_diagnostic_artifact_after_repairs",
    "full_bmc_toy_gate": "blocked",
    "forbidden_promotions": [
      "full BMC",
      "full quantum gravity",
      "proof of Bohmian mechanics",
      "solution to the problem of time",
      "spacetime emergence proof",
      "Friedmann recovery",
      "valid φ-clock for full cosmology"
    ]
  },
  "next_smallest_useful_move": "If Sprint 4 finds clock_fragile: investigate whether a different relational clock variable (e.g. a monotonic function of α and φ) resolves the fragility, or formally document clock-choice as a permanent obstruction for this parameter region. If Sprint 4 finds clock_stable: proceed to Friedmann residual check (Sprint 5 candidate)."
}
```
