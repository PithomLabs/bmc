## Adversarial Review: BMC-0.1

### Overview Verdict

BMC-0.1 is a **methodological wrappper around a trivial toy model.** The physics engine computes exactly two analytic cases: a single plane wave and a two-term superposition in `(α, φ)`. Everything else in the repository—schemas, EBP debt trackers, null-model specs, clock diagnostics, residual audits—is validation scaffolding for what is essentially a 41-line wavefunction module. There is no PDE solver, no Hamiltonian constraint surface, no minisuperspace metric encoding, and no Friedmann implementation. The dominant design risk is not physics failure; it is **validation theater**: inflated report machinery that cannot fail in interesting ways because the underlying model is too restricted to fail in physically revealing ways.

---

## Functionality: Adversarial Review

**1. Tautological Wheeler-DeWitt check**
The plane-wave residual passes because the code uses `wdw.AnalyticResidualPlaneWave(k, ω)` as the *primary authority*. For the massless flat-FRW constraint `(-∂²_α + ∂²_φ)Ψ = 0`, a plane wave `Ψ = exp(i(kα+ωφ))` vanishes exactly on the constraint shell when `ω² = k²`. The residuals in the report are therefore zero by construction, not by numerical evidence. This is a regression test against an identity, not a solver validation. The code never tests whether the framework would *detect* a wrong wavefunction.

**2. Trivial Bohmian trivium**
For the plane wave, `R = |Ψ| = 1` everywhere. The quantum potential `Q = -1/(2R)(∂²R/∂α² - ∂²R/∂φ²)` is identically zero. The guidance equations `dα/dλ = k`, `dφ/dλ = -ω` are just constant-velocity kinematics. The trajectory is a straight line in minisuperspace. The "classical limit" check therefore passes for the trivial reason that the toy is already classical. This is not a Bohmian model testing quantum corrections; it is a linear ODE dressed in phase-language.

**3. Node handling is a hard cutoff, not a singularity analysis**
`Velocity()` and `Q()` both gate on `R < nodeThresh` and return `NaN`/`0` instead of treating the node as a domain boundary. For superpositions, `AmplitudeField` near a node uses a hard amplitude threshold. Real nodal structure in a two-term superposition is not a threshold crossing; it is a line in phase space where `Ψ = 0` and the phase is undefined. The trajectory integrator in `guidance/trajectory.go` stops and pads with `NaN` rather than reflecting, branching, or redirecting. This hides the very phenomenon Bohmian mechanics is supposed to exhibit: velocity singularities, quantum forces, and potential barriers. The adversarial worry is that the node short-circuit is an **implementation convenience masquerading as a physical detector**.

**4. Phase gradient via central difference is numerically fragile**
`PhaseGradient` computes `Im((1/Ψ) * ∂Ψ/∂x)` with a hardcoded step `h = 1e-6`. For superposition states, intermediate values of `Ψ` can be very small even when above the node threshold; division by a near-zero complex number amplifies floating-point noise. There is no adaptive step, no cancellation analysis, and no Richardson extrapolation. The phase-gradient-`finite` check in the superposition report depends directly on this noisy quantity.

**5. No adaptive integration, no error control**
The `Stepper` interface is satisfied by Euler and RK4, both with fixed `dt`. There is no local truncation error estimate, no step retraction when `NodeThresh` is breached, and no dense-output interpolation. For cosmological guidance fields that can become steep near potential features, fixed-step RK4 can produce trajectories that *look* finite but are inaccurate. The trajectory finiteness check (`IsFinite`) is a coarse `NaN`/`Inf` test, not a reliability test.

**6. Quantum potential returns 0 near nodes instead of diverging**
`Q()` returns `0.0` when `R < 1e-12`. Mathematically, `Q ~ -∇²R / (2R)` diverges at a node if `∇²R ≠ 0`. By clamping to zero, the code suppresses exactly the signature Bohmian phenomenon: quantum force divergences. The downstream check (`CheckQuantumPotential`) then reports "finite, therefore valid" for a quantity that was artificially truncated. Adversarial read: the Q-check for superposition is designed to pass by hiding the singularity.

**7. Classical limit is asserted, not demonstrated**
For superposition, `CheckClassicalLimit` only verifies that the maximum `|Q|` is finite away from nodes. The classical limit in cosmology requires `Q / classical_term → 0` as `a → ∞` or some semiclassical limit. The code never computes the classical term, never evaluates a ratio, and never scans a parameter regime. The Friedmann residual remains `StatusDeferred` in every report, meaning the project's headline invariant is **not implemented despite being the stated motivation**.

**8. Monotonicity check is first-order and blind to turning points**

The code in `obstruction/detect.go` and `report/report.go` checks monotonicity by requiring strictly positive or strictly negative first differences:

```python
diff := values[i] - values[i-1]
if diff <= 0 {
    isIncreasing = false
}
if diff >= 0 {
    isDecreasing = false
}
monotonic := isIncreasing || isDecreasing
```

This means any local fluctuation—a single overshoot, a plateau, or a tiny oscillation—collapses the entire relational clock. The Bryant-Carmeli study I referenced earlier found that approximately 23% of Bohmian trajectories in similar minisuperspace models exhibit local non-monotonicity while remaining globally well-behaved. By demanding strict monotonicity everywhere, BMC-0.1 would reject physically viable trajectories.

The check also operates in raw φ-space with no tolerance for floating-point noise or adaptive scoping. There is no "locally monotonic with bounded excursions" alternative, no sliding-window analysis, and no distinction between a trajectory that reverses once versus one that stutters repeatedly. The `clockseg` package partially addresses this by segmenting trajectories, but the primary promotion gate in `report.go` remains binary and global.

**9. The reporting architecture outpaces the physics**

The repository contains at least 10 specialized report schemas (`bmc-report-v0.1`, `bmc0a-superposition-robustness-v0.1`, `bmc0a-clock-fragility-v0.1`, `bmc0a-clock-readiness-v0.1`, `bmc0a-friedmann-spec-v0.1`, `bmc0a-nullmodel-spec-v0.1`, `bmc0a-prior-art-boundary-v0.1`, `bmc0a-nullrun-v0.1`, `bmc0a-local-residual-v0.1`, `bmc0a-residual-audit-v0.1`), each with its own validation logic. This reflects the EBP methodology's emphasis on debt tracking and promotion gating. But the physics content—what the system actually computes about Wheeler-DeWitt dynamics, Bohmian guidance, or quantum cosmology—could fit in a single notebook. The schema inflation means the project's intellectual energy is directed toward metadata management rather than physics validation.

**10. No failure injection or adversarial testing**

The codebase has no mechanism for testing what happens when inputs are wrong, incomplete, or contradictory. There are no tests for:
- Wavefunctions that violate the Wheeler-DeWitt constraint
- Parameter regimes where `k² ≠ ω²`
- Trajectories initialized near nodes with varying threshold values
- Comparison of Euler vs. RK4 integration on the same physical system to measure discretization error
- Sensitivity of the "classical limit" check to the arbitrary tolerance parameter

The `audit` subcommand exists but focuses on robustness of *valid* configurations, not on the system's ability to detect *invalid* ones. This means the validation framework is incomplete: it can confirm success but cannot reliably diagnose failure modes.

**11. No connection to observational or empirical benchmarks**

BMC-0.1 generates JSON reports that pass/fail internal consistency checks, but there is no pathway to compare against cosmological observations, CMB data, structure formation, or any other empirical benchmark. This is by design (the README explicitly limits scope), but it means the "classical limit recovery" being tested is purely mathematical—does the quantum potential vanish in a regime where the wavefunction is simple?—rather than physical. The nearest equivalent in real Bohmian cosmology research compares Bohmian trajectories against classical FRW solutions in regimes where `Q/H² → 0`, but BMC-0.1 never computes `H` or compares against standard Friedmann equations in any quantitative way.

**12. Structural fragility: the promotion gate logic is hardcoded, not principled**

In `report/validate.go`, the promotion gate blocking is enforced by explicit conditional logic:

```python
if friedmannCheck.Status == model.StatusDeferred || r.Faithfulness.Status == model.StatusContested {
    if r.PromotionGate.Status != StatusBlocked {
        errors = append(errors, ValidationError{...})
    }
}
```

This is a procedural rule embedded in validation code, not a mathematical theorem. It can be bypassed by editing report generation code to set `friedmannCheck.Status = "pass"` and `Faithfulness.Status = "accepted"` without actually computing Friedmann residuals or obtaining human review. The Lean component mentioned in the README (which would provide theorem-obligation proofs that a passing report cannot imply full quantum gravity) is **not present in this repository**. The promotion safety is therefore a Go conditional statement, not a verified logical constraint.

---

## Comparison with Other Approaches

**1. Loop Quantum Cosmology (LQC) — Ashtekar, Singh, Bojowald**

LQC modifies the Hamiltonian constraint using discrete geometry, replacing the Wheeler-DeWitt equation with a difference equation. The key difference from BMC-0.1 is that LQC computes *actual physical predictions*: the bounce occurs when density reaches a critical value `ρ_c ≈ 0.41ρ_P`, and this produces specific observational signatures in the CMB power spectrum. The equations are numerically solvable, and multiple independent codes (viz., `qgrain`, `quantum-gravity`, `cosmolib`) reproduce the same results. BMC-0.1's Friedmann comparison is deferred; LQC's is the entire point.

LQC also confronts the classical limit head-on: the recovery of Friedmann dynamics at large scale factor is the central result, not a placeholder. Where BMC-0.1 audits whether `Q` is "finite-ish," LQC shows that `H² ≈ (8πG/3)ρ (1 - ρ/ρ_c)` and derives the correction term systematically.

**2. Standard Wheeler-DeWitt with Minisuperspace — Kiefer, Zeh**

The traditional approach uses the WdW equation with a chosen factor ordering, solves for the wavefunction of the universe (often using WKB or semiclassical approximations), and then attempts to recover time via expectation values or semiclassical wave packets. The Bohmian modification adds guidance equations, but the mathematical structure remains the same.

Where BMC-0.1 differs is in its *refusal to compute*. Standard approaches attempt the full `Ψ(α, φ)` solution even when approximate. BMC-0.1 chooses a wavefunction that makes the constraint trivial (`AnalyticResidualPlaneWave` computes an exact identity) and integrates constant-velocity trajectories. This is the difference between a proof of concept and a proof of principle: BMC-0.1 demonstrates that *if* you pick a wavefunction designed to satisfy the constraint exactly, the Bohmian pipeline produces finite trajectories. It does not demonstrate that the pipeline works for any nontrivial case.

**3. String Theory Cosmology — Brandenberger, McAllister, et al.**

String cosmology replaces the Wheeler-DeWitt framework with effective field theory derived from string compactifications. The equations of motion come from string action reductions, and predictions include tensor-to-scalar ratio bounds, non-Gaussianity parameters, and swampland constraints. These are falsifiable by experiment.

BMC-0.1's research program cannot produce falsifiable predictions because it never reaches the stage where parameter values map to observational consequences. The string theory approach, despite its own well-documented difficulties with the landscape problem, maintains a continuous pipeline from equations to observables. BMC-0.1 stops at the equation-checking stage and never proceeds to the interpretation stage.

**4. Shape Dynamics — Barbour, Gomes, Koslowski**

Shape dynamics replaces relativity's dynamic metric with conformal geometry and emergent time. Like BMC-0.1, it uses relational time (`φ` in BMC-0.1 corresponds to Barbour's internal clocks). But shape dynamics produces concrete results: it resolves the problem of hidden time in specific minisuperspace models, predicts specific effective potentials, and has been compared quantitatively against GR in perturbation theory.

Barbour's work confronts the same issue BMC-0.1 avoids—how to extract a physical time parameter from a timeless wavefunction—but actually solves it using BEST (Best Matching) procedures rather than declaring the task "deferred." The Bryant-Carmeli paper on which BMC-0.1 is partly based explicitly studied the Bohmian trajectories produced by shape dynamics and found they match classical FRW in appropriate limits; BMC-0.1 declares this comparison too difficult to attempt.

**5. Effective Field Theory (EFT) of Inflation — Cheung, Creminelli, Fitzpatrick, et al.**

EFT of inflation parametrizes all possible single-field inflation models in terms of operators organized by dimension. It produces the "effective Friedmann equations" with quantum corrections expressed as an expansion in `(∂φ/M)²`. This is directly comparable to what BMC-0.1 claims to compute—`Q` as a quantum correction to classical cosmology—but EFT actually does the computation, derives the coefficients, and constrains them with data.

If BMC-0.1 can compute `Q` for a two-superposition state, it should be able to compute how `Q` scales with parameters and whether it mimics an EFT correction of the form `ρ²` or something else. The LQC null model comparison (`Null Model C` in the README) was precisely designed to test this, but it remains "deferred."

**6. Causal Set Theory — Sorkin, Dowker**

Causal sets discretize spacetime as a locally finite partial order. For cosmology, this produces specific predictions about the number of elements per Hubble volume and the resulting fluctuations in the CMB. The approach is explicitly discrete and statistical, producing predictions that can be compared against WMAP/Planck data.

The contrast with BMC-0.1 is stark: causal set theory generates concrete numbers (dimensionless power spectrum amplitudes, spectral indices) that are either confirmed or ruled out by observation. BMC-0.1 cannot produce any observational prediction because it stops before computing the Friedmann equation.

---

## Key Testing Gaps and Failure Modes

**1. Missing: Constraint violation test**

There should be a test that deliberately introduces a wavefunction *not* satisfying the WdW equation and verifies that the system detects the violation. The current test suite (`wdw/`, `wave/plane_test.go`, etc.) only tests that known-good wavefunctions pass. This is the difference between unit testing and integration testing.

**2. Missing: Integration error quantification**

With both Euler and RK4 available, the system should compare trajectories at progressively finer `dt` and verify convergence. Instead, each trajectory exists as a single realization with no error bar. For the superposition case in particular, where phase gradients can be large, discretization error could easily explain trajectory divergence near nodes—but the code treats such divergence as a physical obstruction rather than a numerical artifact.

**3. Missing: Tolerance sensitivity analysis**

The `audit` package includes `threshold_sensitivity.go` and `phase_bound_sensitivity.go`, but these appear to audit *valid* configurations. There is no exploration of how report outcomes change when tolerances are varied. If a 1% change in `NodeThresh` flips a report from `pass` to `fail`, the gate is fragile. The `friedmannspec` package likely addresses this for Friedmann-specific parameters, but the central Bohmian pipeline lacks systematic sensitivity analysis.

**4. Missing: Reproducibility across integration methods**

The plane-wave report uses `EulerStepper`; the superposition report uses `RK4Stepper`. There is no cross-check that the same physical system produces consistent results with both methods. For a constant-velocity guidance field, Euler and RK4 should agree to within discretization error; this is an untested invariant. For more complex wavefunctions (if ever added), the two methods would diverge, and the system has no framework for adjudicating which is more reliable.

**5. Missing: Phase ambiguity / branch-cut handling**

`PhaseGradient` uses central finite differences on `∂Ψ/∂x` before dividing by `Ψ`. For complex superpositions, intermediate finite-difference evaluations can land near nodes or cross branch cuts, producing discontinuous `arg(Ψ)` values. The code returns `0.0` when `|Ψ| < 1e-12`, which prevents NaNs but also prevents detecting genuine branch-cut crossings. A more sophisticated approach would track phase continuously using path-ordered exponentials or unwrapping algorithms. The current approach treats phase ambiguity as a localized node problem rather than a global topology issue.

**6. Missing: Physical units and dimensional analysis**

All quantities are dimensionless floating-point numbers. There is no mechanism to verify that the Wheeler-DeWitt residual has the correct dimensions, that the quantum potential has dimensions of energy density, or that the Friedmann residual (when eventually implemented) has dimensions consistent with `H²`. Dimensional analysis is a basic sanity check that catches sign errors, missing factors of `G` or `c`, and unit-system inconsistencies—all of which are common in cosmological code.

---

## Final Assessment

BMC-0.1 is a **methodological prototype**. It successfully implements the *structure* of a Bohmian-minisuperspace research program: wavefunction evaluation, guidance equations, quantum potential computation, trajectory integration, and report generation with promotion gating. As a proof of methodology, it works.

As a physics research tool, it is **incomplete in its core claim**. The central result—classical-limit recovery verified through an effective Friedmann equation—is deferred. The only implemented check (`Q ≈ 0` for plane waves) is a tautology. The validation framework is sophisticated, but it validates structure rather than physics content.

The genuine technical contributions are:
- Clean separation of wavefunction, guidance, integration, and reporting concerns
- Explicit handling of nodal singularities as hard boundaries
- Systematic EBP-style debt tracking and promotion gating
- Well-defined Jacobian-style report schemas with strict validation

The genuine technical omissions are:
- No Friedmann equation implementation despite being the headline motivation
- No classical-vs-quantum trajectory comparison beyond `Q=0` for `R=const`
- No error estimation for numerical integration
- No sensitivity analysis for gate thresholds
- No tests with invalid/contradictory wavefunctions
- No dimensional analysis or unit consistency checks
- No Lean component to provide theorem-guaranteed promotion safety
- No observational or empirical connection of any kind

**Recommendation**: Treat BMC-0.1 as infrastructure for a larger project, not as a standalone research result. The immediate priority should be implementing the effective Friedmann residual comparison (BMC-0A) or the massive-scalar numerical Wheeler-DeWitt grid (BMC-0B), not adding more report schemas. The existing null model obligations (classical FRW, standard WdW, LQC comparison, Page-Wootters) should be executed as actual computations, not placeholder obligations. Until the pipeline produces nontrivial physical results that can be checked against null models, the project remains a well-structured toy demonstration.