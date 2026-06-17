**Go** = executable toy checks, numerical/finite tests, JSON reports.
**Lean** = formal contracts, theorem obligations, promotion blockers, and faithfulness gates.

This follows EBP 2.1: ideas enter freely, but promotion requires debt retirement; Lean is not the entry gate, but it is a strong debt-retirement instrument for definitions, theorem stubs, counterexamples, and executable-check contracts.  The BMC draft already gives the target toy model: FRW minisuperspace with homogeneous scalar field, Wheeler-DeWitt equation, Bohmian guidance, quantum potential, effective Friedmann check, and explicit non-claims. 

Here is the synthesized implementation plan.

# Implementation Plan: Bohmian Minisuperspace Cosmology v0.1

## 0. Status

Project name:

```text
Bohmian Minisuperspace Cosmology v0.1
```

Short name:

```text
BMC-0.1
```

Implementation languages:

```text
Go: executable toy checks, numerical integration, JSON reports
Lean: formal contracts, theorem obligations, promotion blockers
```

Current EBP status:

```text
alive
unpromoted
workbench-ready
toy-check debt active
```

No final-truth claim is allowed. BMC-0.1 is not a theory of everything, not full quantum gravity, not a derivation of geometry from non-geometry, and not a solution to the problem of time.

Its only allowed claim is:

```text
In a symmetry-reduced FRW minisuperspace model with homogeneous scalar field, a Wheeler-DeWitt constraint plus Bohmian guidance may generate an actual trajectory whose relational cosmological history can be checked against the classical Friedmann limit and rival toy models.
```

---

# 1. Core Design Choice

Do not begin with the full massive-scalar numerical grid.

Begin with a smaller analytic control model first.

The implementation should have two toy stages:

```text
BMC-0A: flat FRW + massless scalar field, analytic or semi-analytic wavefunctions
BMC-0B: FRW + massive scalar field, finite-difference numerical Wheeler-DeWitt grid
```

Why this matters:

```text
BMC-0A tests the entire Bohmian pipeline without hiding bugs inside a PDE solver.
BMC-0B tests the more ambitious numerical model after the pipeline is already trusted.
```

This prevents a common failure mode: confusing numerical artifacts with physics.

---

# 2. Minimum Viable Physics Scope

## 2.1 Variables

Use logarithmic scale factor:

```text
α = ln(a)
```

Configuration:

```text
q = (α, φ)
```

where:

```text
α = log scale factor
φ = homogeneous scalar field
```

Relational cosmological history:

```text
a(φ) = exp(α(φ))
```

This is safer than focusing first on absolute Bohmian time `t`, because the relational curve `a(φ)` is the physically meaningful output in a timeless Wheeler-DeWitt setting.

---

## 2.2 Phase 0A: Massless Scalar Control Model

Use a simplified Wheeler-DeWitt form:

```text
(-∂²/∂α² + ∂²/∂φ²) Ψ(α, φ) = 0
```

The sign convention should be configurable and recorded in the report.

Start with two wavefunction families:

```text
1. Plane-wave control state
2. Gaussian wavepacket or superposition state
```

Plane wave:

```text
Ψ(α, φ) = exp(i(kα + ωφ))
```

Constraint condition:

```text
ω² = k²
```

Expected behavior:

```text
R = constant
Q = 0
Bohmian trajectory should reproduce classical-like relational behavior
```

This is the baseline regression test.

Then use a nontrivial packet/superposition:

```text
Ψ = Σ_n c_n exp(i(k_n α + ω_n φ))
```

Expected behavior:

```text
R is nonconstant
Q may be nonzero
nodes may appear
guidance may fail near nodes
```

This tests the real obstruction behavior.

---

## 2.3 Phase 0B: Massive Scalar Numerical Model

After Phase 0A passes, implement the numerical Wheeler-DeWitt toy:

```text
[-∂²/∂a² + (1/a²)∂²/∂φ² + a⁴V(φ) - ka²]Ψ(a,φ) = 0
```

with:

```text
V(φ) = 1/2 m²φ²
```

This is not the first implementation step. It is the second-stage implementation after the analytic control pipeline works.

---

# 3. Bohmian Guidance Layer

Represent the wavefunction as:

```text
Ψ = R exp(iS)
```

Compute:

```text
R = |Ψ|
S = arg(Ψ)
```

Guidance equations for the BMC toy:

```text
dα/dλ = ∂S/∂α
dφ/dλ = -∂S/∂φ
```

or, in the `a` variable:

```text
da/dλ = ∂S/∂a
dφ/dλ = -(1/a²)∂S/∂φ
```

The exact convention must be recorded in the report.

For Phase 0A, prefer `α` because it avoids singular behavior at `a = 0`.

---

# 4. Quantum Potential Layer

For Phase 0A in `(α, φ)` coordinates:

```text
Q = -1/(2R) (∂²R/∂α² - ∂²R/∂φ²)
```

For Phase 0B in `(a, φ)` coordinates:

```text
Q = -1/(2R) (∂²R/∂a² - (1/a²)∂²R/∂φ²)
```

The implementation must report:

```text
Q along each trajectory
max |Q|
mean |Q|
Q/classical-term ratio
finite/nonfinite Q count
node proximity
```

---

# 5. Map Debt

The map is intentionally weak and honest.

Domain:

```text
Bohmian trajectory in minisuperspace:
γ(λ) = (α(λ), φ(λ))
```

Codomain:

```text
FRW relational history:
a(φ) = exp(α(φ))
```

Optional reconstructed metric:

```text
ds² = -N(λ)²dλ² + a(λ)²dΣ_k²
```

But the first promoted map should be:

```text
γ(λ) -> a(φ)
```

not:

```text
γ(λ) -> full spacetime geometry
```

Reason:

```text
The relational history avoids overclaiming about lapse, proper time, and full spacetime emergence.
```

---

# 6. Invariant Debt

The first invariant is not entropy.

The first invariant is:

```text
classical-limit recovery
```

Operational form:

```text
Q_ratio = |Q| / |classical_term|
```

Success condition:

```text
Q_ratio -> small in the classical/large-a region
Friedmann residual -> small in the same region
```

Report both:

```text
QRatioLate
FriedmannResidualLate
```

BMC-0.1 succeeds only if:

```text
QRatioLate <= threshold
FriedmannResidualLate <= threshold
trajectory remains finite
φ is monotonic enough to define a(φ)
```

---

# 7. Toy Check Debt

## 7.1 BMC-0A toy check

Inputs:

```text
wavefunction family
k values
grid range for α and φ
trajectory initial conditions
integration step size
node threshold
classical-limit region
```

Outputs:

```text
WdW residual
phase-gradient field
Bohmian trajectories
relational histories a(φ)
Q field
Q along trajectories
classical-limit metrics
node obstruction report
JSON report
```

Success criteria:

```text
WdW residual is below tolerance
phase is usable away from nodes
trajectories integrate without nonfinite values
φ is monotonic on selected trajectories
a(φ) is extractable
Q is finite away from nodes
plane-wave control gives Q ≈ 0
classical-like trajectory is recovered in the control case
```

Failure criteria:

```text
WdW residual too large
phase unwrapping fails
trajectory enters node region
Q diverges
φ is not monotonic
no relational history can be extracted
classical control case fails
```

---

## 7.2 BMC-0B toy check

Only begin after BMC-0A passes.

Inputs:

```text
V(φ)=1/2m²φ²
grid in a and φ
boundary condition
finite-difference scheme
initial trajectory points
integration settings
```

Outputs:

```text
numerical Ψ
WdW residual grid
phase S
amplitude R
guidance field
trajectory ensemble
Q field
effective Friedmann residual
comparison with LQC correction form
JSON report
```

Success criteria:

```text
numerical solution is stable enough for guidance
Q finite on selected trajectory regions
late-time classical recovery appears
quantum correction is computable
failure modes are explicitly reported
```

---

# 8. Null Model Debt

BMC-0.1 must compare against at least three null models.

## 8.1 Null Model A: Classical FRW

Question:

```text
Does the Bohmian trajectory reduce to the classical FRW trajectory when Q is negligible?
```

Metric:

```text
late-time Friedmann residual
```

## 8.2 Null Model B: Standard Wheeler-DeWitt without beables

Question:

```text
Does the Bohmian ensemble merely reproduce the same |Ψ|² distribution, adding ontology but no operational difference?
```

Metric:

```text
trajectory ensemble density vs |Ψ|²
```

This can be postponed until after single-trajectory tests pass.

## 8.3 Null Model C: Loop Quantum Cosmology-style correction

Question:

```text
Does Q_Bohm mimic an algebraic bounce correction, or is it functionally distinct?
```

Metric:

```text
functional comparison between Q_Bohm and a simple ρ²-style correction
```

## 8.4 Null Model D: Page-Wootters relational time

Question:

```text
Does BMC produce a relational history a(φ) that differs from or merely restates an internal-clock relational-time model?
```

For BMC-0.1 this can be recorded as a formal null-model obligation, not fully executed.

---

# 9. Obstruction Debt

The implementation must automatically file obstructions.

Required obstruction flags:

```text
node_obstruction
phase_unwrap_obstruction
nonfinite_q_obstruction
clock_nonmonotonicity_obstruction
wdw_residual_obstruction
classical_limit_failure
lapse_or_time_interpretation_debt
measure_problem_deferred
full_qg_overclaim_blocker
```

Each obstruction should have:

```text
applies: true/false
severity: info/warning/blocker
evidence: short string
consequence: downgrade, block, or continue
```

---

# 10. Faithfulness Debt

Faithfulness question:

```text
Does the executable Go model actually test the BMC-0.1 claim?
```

The answer is allowed to be:

```text
yes
no
contested
```

For BMC-0.1, faithfulness should remain:

```text
contested until BMC-0A report exists
```

The first Go report must include this warning:

```text
This report tests only a symmetry-reduced minisuperspace toy model. It does not test full quantum gravity, full Lorentz recovery, black hole information, fermions, gauge fields, or inhomogeneous perturbations.
```

---

# 11. Go Implementation Architecture

Recommended repository structure:

```text
cmd/ptw-bmc/
  main.go

internal/bmc/
  model/
    types.go
    params.go
  wave/
    plane.go
    packet.go
    phase.go
    amplitude.go
  wdw/
    residual.go
    finite_difference.go
  guidance/
    velocity.go
    integrate.go
    trajectory.go
  qpotential/
    q.go
    derivatives.go
  invariant/
    classical_limit.go
    friedmann.go
  nullmodel/
    classical_frw.go
    lqc_compare.go
    wdw_ensemble.go
  obstruction/
    obstruction.go
    detect.go
  report/
    report.go
    validate.go
    write_json.go

testdata/bmc/
  bmc0a_plane_expected.json
  bmc0a_packet_expected.json

out/
  bmc0a_report.json
```

Minimum CLI commands:

```text
ptw-bmc run --profile bmc0a-plane --out out/bmc0a_plane.json
ptw-bmc run --profile bmc0a-packet --out out/bmc0a_packet.json
ptw-bmc validate --report out/bmc0a_plane.json
ptw-bmc summarize --report out/bmc0a_packet.json
```

Do not start with dashboards.

Do not start with a database.

Do not start with LLM/RAG integration.

Do not start with PTW UI.

First produce deterministic JSON reports.

---

# 12. Go Report Schema

Minimum JSON report fields:

```json
{
  "schema_version": "bmc-report-v0.1",
  "model_id": "bmc0a_plane",
  "toy_analysis_only": true,
  "physics_claim": "minisuperspace_only",
  "final_truth_claim": false,
  "promotion_recommendation": "blocked_or_candidate",
  "parameters": {},
  "equations": {},
  "checks": {
    "wdw_residual": {},
    "trajectory": {},
    "clock_monotonicity": {},
    "quantum_potential": {},
    "classical_limit": {},
    "friedmann_residual": {}
  },
  "null_models": [],
  "obstructions": [],
  "faithfulness": {
    "status": "contested",
    "reason": "No human review yet."
  },
  "ebp_debt": {
    "needMap": "partial",
    "needInvariant": "partial",
    "needToyCheck": "active",
    "needNullModel": "partial",
    "needObstruction": "active",
    "needFaithfulnessReview": "active"
  }
}
```

---

# 13. Lean Implementation Architecture

Recommended Lean structure:

```text
BMC/
  EBP.lean
  Minisuperspace.lean
  Map.lean
  Invariant.lean
  ToyReport.lean
  Obstruction.lean
  Faithfulness.lean
  Promotion.lean
```

Lean should not try to prove the numerics.

Lean should formalize:

```text
what the claim is
what the map is
what the invariant means
what counts as a report
what blocks promotion
what the toy model does not imply
```

---

# 14. Lean Contract Sketch

Lean concepts to define:

```lean
inductive DebtItem
| needMap
| needInvariant
| needToyCheck
| needNullModel
| needObstruction
| needFaithfulnessReview

inductive PromotionStatus
| alive
| blocked
| toyCandidate
| promotedToyArtifact
```

Core structures:

```lean
structure MiniState where
  alpha : Float
  phi : Float

structure Trajectory where
  points : List MiniState

structure FRWHistory where
  relationalPoints : List (Float × Float)

structure BMCReport where
  toyAnalysisOnly : Bool
  finalTruthClaim : Bool
  wdwResidualPass : Bool
  trajectoryFinitePass : Bool
  clockMonotonicPass : Bool
  qFinitePass : Bool
  classicalLimitPass : Bool
  friedmannResidualPass : Bool
  faithfulnessAccepted : Bool
```

Promotion rule:

```lean
def reportPassesToyGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.wdwResidualPass &&
  r.trajectoryFinitePass &&
  r.clockMonotonicPass &&
  r.qFinitePass &&
  r.classicalLimitPass &&
  r.friedmannResidualPass &&
  r.faithfulnessAccepted
```

Theorem obligation:

```lean
theorem no_full_qg_claim_from_toy_report
  (r : BMCReport) :
  reportPassesToyGate r = true ->
  r.toyAnalysisOnly = true := by
  intro h
  -- proof follows from definition
  simp [reportPassesToyGate] at h
  exact h.left
```

This theorem is deliberately humble. It proves only that a passing toy report remains toy-only.

---

# 15. Lean Theorem Obligations

Minimum theorem obligations:

```text
T1: A passing toy report cannot imply full quantum gravity.
T2: A report with finalTruthClaim=true cannot be promoted.
T3: A report with faithfulnessAccepted=false cannot be promoted.
T4: A report with toyAnalysisOnly=true must remain classified as toy-level.
T5: If φ is not monotonic, relational history a(φ) is blocked.
T6: If Q is nonfinite, classical-limit recovery is blocked.
T7: If WdW residual fails tolerance, toy-check debt remains active.
```

These are promotion-safety theorems, not physics proofs.

---

# 16. Test Plan

## Go tests

Required tests:

```text
TestPlaneWaveSatisfiesWdWResidual
TestPlaneWaveQApproximatelyZero
TestPlaneWaveTrajectoryFinite
TestPacketReportsNodeObstruction
TestClockMonotonicityDetection
TestNonfiniteQBlocksPromotion
TestReportDeterministicJSON
TestNoFinalTruthClaimAllowed
```

## Lean tests/build

Required Lean checks:

```text
lake build
```

Required Lean theorem files:

```text
Promotion.lean
Faithfulness.lean
Obstruction.lean
```

---

# 17. EBP Gates

## Gate 0: Capture

Output:

```text
BMC-0.1 captured as alive, unpromoted.
```

## Gate 1: Map specified

Pass condition:

```text
Domain, codomain, and translation rule are recorded.
```

Status:

```text
needMap partially retired
```

## Gate 2: Invariant specified

Pass condition:

```text
Classical-limit recovery and Q-ratio are defined.
```

Status:

```text
needInvariant partially retired
```

## Gate 3: Toy check executed

Pass condition:

```text
BMC-0A plane-wave and packet reports generated.
```

Status:

```text
needToyCheck partially retired if reports pass
```

## Gate 4: Null models filed

Pass condition:

```text
Classical FRW, standard WdW, LQC, and Page-Wootters obligations recorded.
```

Status:

```text
needNullModel partially retired
```

## Gate 5: Obstructions filed

Pass condition:

```text
Node, Q divergence, clock failure, WdW residual, measure problem, and overclaim blockers are represented in the report.
```

Status:

```text
needObstruction partially retired
```

## Gate 6: Faithfulness reviewed

Pass condition:

```text
Human or reviewer explicitly accepts that the Go toy report tests the stated BMC-0.1 claim and nothing stronger.
```

Status:

```text
needFaithfulnessReview retired only after review
```

---

# 18. Promotion Rule

BMC-0.1 may be promoted only as:

```text
promoted toy artifact
```

Never as:

```text
quantum gravity solution
problem of time solution
full Bohmian QG theory
emergent spacetime proof
```

Promotion wording if all gates pass:

```text
BMC-0.1 is a promoted toy-level Bohmian minisuperspace artifact under current debt conditions. It demonstrates that the specified Wheeler-DeWitt/Bohmian guidance pipeline can produce finite relational trajectories and pass the selected classical-limit checks in the tested model. It does not establish full quantum gravity.
```

---

# 19. First Implementation Sprint

## Sprint goal

Build BMC-0A plane-wave control.

## Deliverables

```text
Go package internal/bmc/model
Go package internal/bmc/wave
Go package internal/bmc/wdw
Go package internal/bmc/guidance
Go package internal/bmc/qpotential
Go package internal/bmc/report
CLI cmd/ptw-bmc
Lean BMC/Promotion.lean
Lean BMC/ToyReport.lean
out/bmc0a_plane.json
```

## Sprint success

```text
Plane-wave WdW residual passes.
Q ≈ 0.
Trajectory finite.
JSON deterministic.
Lean promotion blocker builds.
Report remains toy-only.
```

## Sprint failure

```text
If plane-wave control fails, do not proceed to packet states or massive scalar numerics.
Fix equations, signs, derivative code, or report contracts first.
```

---

# 20. Second Implementation Sprint

## Sprint goal

Build BMC-0A packet/superposition test.

## Deliverables

```text
packet wavefunction
phase extraction
node detection
Q computation
trajectory integration
obstruction report
```

## Sprint success

```text
Nodes are detected.
Q is computed away from nodes.
Trajectories either pass or fail transparently.
Clock monotonicity is reported.
No overclaiming occurs.
```

---

# 21. Third Implementation Sprint

## Sprint goal

Begin BMC-0B massive scalar numerical WdW model.

## Deliverables

```text
finite-difference grid
boundary condition module
numerical residual checker
massive scalar potential
first numerical Ψ experiment
```

This sprint should begin only after BMC-0A is stable.

---

# 22. What Not To Build Yet

Do not build yet:

```text
black hole information module
fermion ontology
gauge fields
inhomogeneous perturbations
full field-theoretic measure
Lorentz recovery
quantum gravity dashboard
PTW UI integration
automatic paper claims
large RAG corpus
```

Those are later phases.

The next smallest useful move is only:

```text
Implement BMC-0A plane-wave control in Go and the toy-only promotion blocker in Lean.
```

The key correction to Kimi’s draft is that we should **not** mark most debts retired yet. Under EBP, stating a map or invariant makes it inspectable, but toy success and faithfulness review still matter; EBP says toy checks are wind tunnels and faithfulness asks whether the formalization matches the intended claim.  Kimi’s BMC draft is valuable because it defines the equations, failure modes, null models, and explicit non-claims, but its ledger is too optimistic; it says several debts are retired while also saying toy execution is still pending. 

The immediate next step is therefore:

```text
BMC-0A:
Implement the massless flat-FRW plane-wave control in Go.
Generate deterministic JSON.
Add Lean theorem obligations proving that even a passing toy report cannot promote a full quantum-gravity claim.
```

That is the smallest EBP-compliant move that actually retires debt.
