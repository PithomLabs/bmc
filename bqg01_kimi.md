---

## Starting Position: What We Actually Know vs. What We Don't

**Established facts (not up for debate):**
- Non-relativistic Bohmian mechanics works mathematically and reproduces all quantum predictions
- The measurement problem in quantum cosmology is genuine (no external observer)
- Jacobson 1995 derived Einstein's equations from thermodynamic assumptions on local Rindler horizons—but assumed pre-existing geometry
- The Wheeler-DeWitt equation ĤΨ = 0 is a constraint, not a dynamical equation
- Lorentz covariance is empirically exact; Lorentz invariance (no preferred frame) is an ontological claim experiment cannot settle

**What we do not know:**
- Whether any beable ontology extends to relativistic quantum field theory consistently
- Whether emergent geometry can be derived from pre-geometric variables
- Whether the problem of time admits any solution, Bohmian or otherwise
- Whether quantum gravity requires new degrees of freedom or just new organization of known ones

**Strategic implication:** Build upward from what works, not downward from what we wish were true.

---

## EBP 2.1 Capture

**Owner:** Draft
**Claim:** A Bohmian minisuperspace model with a single scalar field degree of freedom, where geometry emerges from the beable trajectory via an explicit map, can be constructed without contradiction and yields testable differences from standard Wheeler-DeWitt quantum cosmology.

**Source:** Original synthesis + adversarial repairs
**Born:** 2026-06-17

**Debt:** Full debt attached.

---

## Debt Retirement: Sequential, Ruthless

---

### `needMap` — First, Define the Smallest Possible Domain

**Decision:** Start with minisuperspace. Not because it's ambitious, but because it's the only place where the map can actually be written down.

**Domain:** Two-dimensional configuration space (a, φ), where a is the scale factor and φ is a homogeneous scalar field. Both are functions of a single parameter t (Bohmian time).

**Why a and not something "more fundamental"?** Because we need a codomain we can check against. The codomain is FRW geometry. If we start with variables that don't obviously map to FRW, we can't verify the map. Pre-geometric romanticism is deferred until the geometric case works.

**Codomain:** FRW spacetimes with metric ds² = -N(t)²dt² + a(t)²dΩ₃², where dΩ₃² is the metric on S³, ℝ³, or H³.

**Translation rule (explicit):**
1. The spatial metric is h_ij = a(t)² δ_ij (or appropriate curvature)
2. The lapse N(t) is derived from the Bohmian velocity: N = (da/dt)/|∂S/∂a|, where Ψ = Re^{iS}
3. The shift is zero by homogeneity
4. The 4-metric is assembled as above

**Check:** Is h_ij positive definite? Is N real? These are falsifiable conditions on the trajectory. If they fail, the map fails transparently.

**What this map does NOT claim:** It does not claim to derive a from something pre-geometric. It takes a as a beable and maps it to geometry. The "emergence" is of the 4D Lorentzian structure and Einstein dynamics, not of the spatial scale factor itself.

**Why this is honest:** We are not pretending to derive geometry from non-geometry. We are deriving spacetime geometry from spatial geometry + Bohmian dynamics. This is weaker but actually constructible.

---

### `needInvariant` — Find Something That Must Survive

**Candidate invariant:** The Bohmian quantum potential Q(a, φ) evaluated on the trajectory.

**Why this fails as an invariant:** Q is not preserved under the map. It is a feature of the domain (configuration space dynamics) that has no obvious geometric counterpart in the codomain.

**Repair:** The quantity that must survive is the **effective Friedmann equation** itself. In standard cosmology, this is:
(ȧ/a)² = (8πG/3)ρ - k/a² + Λ/3

In the Bohmian model, this becomes:
(ȧ/a)² = (8πG/3)ρ_eff - k/a² + Λ/3 + Q_Bohm(a, φ)

where Q_Bohm is the Bohmian quantum correction.

**Invariant statement:** The functional form of the Friedmann equation is preserved, with Q_Bohm as an additive correction that vanishes as ℏ → 0.

**Why this is checkable:** We can compare the Bohmian trajectory to the classical trajectory and to the WKB approximation. Q_Bohm is computable from Ψ. If it doesn't vanish correctly in the classical limit, the invariant is violated and the map fails.

**Sharper invariant (numerical):** The ratio Q_Bohm / (8πGρ/3) as a function of a. This must → 0 as a → ∞ (classical limit) and must be finite as a → 0 (quantum regime).

---

### `needToyCheck` — Actually Run It

**The model (specific equations):**

Wheeler-DeWitt equation for FRW with scalar field:
[-∂²/∂a² + (1/a²)∂²/∂φ² + a⁴V(φ) - ka²]Ψ(a,φ) = 0

(Units: 8πG = 1, ℏ = 1, c = 1. Mass parameter absorbed into field definition.)

**Boundary conditions:** 
- Ψ regular at a = 0 (no-boundary or tunneling condition)
- Ψ → 0 as a → ∞ (normalizable in some appropriate sense, or WKB outgoing)

**Guidance equations:**
da/dt = ∂S/∂a
dφ/dt = -(1/a²)∂S/∂φ

where Ψ = |Ψ|e^{iS}.

**Quantum potential:**
Q = -(1/2|Ψ|)(∂²|Ψ|/∂a² - (1/a²)∂²|Ψ|/∂φ²)

**Effective Friedmann equation from guidance:**
(1/a da/dt)² = (1/3)[(1/2)(dφ/dt)² + V(φ) + Q]

**Toy check protocol:**
1. Choose V(φ) = (1/2)m²φ² (massive scalar, simplest case)
2. Solve WdW numerically on a grid (a, φ) ∈ [0, a_max] × [-φ_max, φ_max]
3. Extract S(a, φ) from Ψ (phase unwrapping required)
4. Integrate guidance equations from initial (a₀, φ₀) with da/dt > 0 (expanding universe)
5. Compute Q along trajectory
6. Verify: (i) Q → 0 as a → ∞, (ii) effective equation matches classical Friedmann at late times, (iii) quantum corrections are significant at early times

**Success criterion:** The Bohmian trajectory exists for all t, does not hit singularities (a = 0 with da/dt = 0), and the effective Friedmann equation holds with computable Q.

**Failure modes (explicit):**
- S is not globally defined (nodes in Ψ where |Ψ| = 0)
- Trajectory hits a = 0 in finite t (Bohmian bounce fails)
- Q diverges, making effective equation meaningless
- Late-time behavior does not match classical cosmology

**Why this is a real check:** These are not philosophical objections. They are numerical outcomes. The model either computes or it doesn't.

---

### `needNullModel` — Engage the Actual Rivals

**Rival 1: Standard Wheeler-DeWitt quantum cosmology (no Bohmian beables)**

What it does: Solves WdW, computes wave function, extracts probabilities via Born rule (in some interpretation).

What the Bohmian model claims to do differently: Provides a single trajectory, not a probability distribution. The "universe" is one actual (a(t), φ(t)), not a superposition.

**Engagement question:** Does the Bohmian trajectory sample the WdW probability distribution |Ψ|²?

**Test:** Compute the ensemble of Bohmian trajectories from different initial conditions (a₀, φ₀) weighted by |Ψ(a₀, φ₀)|². Check if the distribution of (a, φ) at late times matches |Ψ(a, φ)|².

**Expected outcome (from non-relativistic BM):** Yes, by the quantum equilibrium hypothesis. But this is a theorem for finite-dimensional systems, not proven for WdW. The toy check tests whether equilibrium holds in this minisuperspace.

**If equilibrium fails:** The Bohmian model makes different predictions from standard WdW. This is a feature, not a bug—it means the models are distinguishable.

**If equilibrium holds:** The Bohmian model is empirically equivalent to standard WdW in this regime. The beable is explanatory overhead unless it resolves the problem of time or measurement problem in a way standard WdW cannot.

**Rival 2: Loop quantum cosmology (LQC)**

What it does: Replaces a with a discrete variable (eigenvalues of the area operator), modifies the Friedmann equation with ρ²/ρ_crit correction.

**Engagement question:** Does the Bohmian quantum correction Q_Bohm mimic the LQC correction, or are they different?

**Test:** Compare functional form of Q_Bohm(a) vs. LQC correction ρ(a)²/ρ_crit.

**Expected outcome:** Different functional forms. LQC correction is algebraic in ρ; Bohmian correction depends on second derivatives of |Ψ|.

**Distinguishability:** In principle, yes, if we had observational access to the quantum correction. In practice, both are currently unobservable.

**Honest null-model conclusion:** The Bohmian model is not necessary (standard WdW may suffice), not unique (LQC provides alternative quantum corrections), and not currently superior in predictive power. It is ontologically distinct (single trajectory vs. superposition) and may resolve the measurement problem in cosmology. These are real but non-empirical distinctions.

**Disclaimers (explicit):**
- We do not claim this model is the only way to quantize cosmology
- We do not claim it is necessary for resolving the problem of time
- We claim only that it is a consistent, calculable alternative with a clear ontology

---

### `needObstruction` — File Honestly, Resolve or Downgrade

**Obstruction 1: Node problem**

In Bohmian mechanics, the velocity field v = ∇S/m is undefined where Ψ = 0 (nodes). In the WdW context, nodes are generic in excited states.

**Status:** Applies to any Bohmian WdW model.
**Consequence:** The model must be restricted to node-free wave functionals, or the guidance equation must be modified near nodes, or nodes must be accepted as singularities in the velocity field.

**Resolution for minisuperspace:** Use the ground state or a carefully chosen superposition that avoids nodes in the relevant region. State explicitly: "Results hold for node-free Ψ; extension to generic states requires node regularization, which is not addressed here."

**Obstruction 2: Measure problem (infinite-dimensional extension)**

As filed previously: no natural measure on field configuration space.

**Status:** Does not apply to minisuperspace (finite-dimensional). Applies to full theory.
**Consequence:** The model is explicitly restricted to minisuperspace. Claims about full field theory are conjectural and marked as such.

**Downgrade:** "This is a minisuperspace model. Extension to inhomogeneous perturbations and full field theory requires solving the measure problem, which remains open."

**Obstruction 3: Problem of time relocation**

The adversarial review noted: the problem of time may relocate into explaining why the Bohmian foliation looks like Lorentzian time.

**Status:** Applies.
**Consequence:** The model does not "dissolve" the problem of time. It reframes it: the Bohmian time t is the fundamental parameter, and the question becomes why the emergent FRW time coordinate τ (proper time of comoving observers) approximates t in the coarse-grained limit.

**Resolution:** In minisuperspace, there is only one time parameter. The identification t ≈ τ is explicit by construction (the lapse N(t) relates them). The problem of time relocation is deferred to inhomogeneous models, where multiple clocks could disagree.

**Honest statement:** "The problem of time is reframed, not solved. In minisuperspace, the reframing is trivial. In full gravity, it is non-trivial and unresolved."

---

### `needFaithfulnessReview` — Match Formalism to Claim

**Claim:** The model provides a Bohmian ontology for quantum cosmology: a single actual trajectory guided by the WdW wave function.

**Formalism check:**
- Is there a single trajectory? Yes: (a(t), φ(t)) from guidance equations.
- Is it guided by Ψ? Yes: S and Q are derived from Ψ.
- Does Ψ satisfy WdW? Yes: by construction.
- Is the ontology clear? Yes: a and φ are beables; Ψ is the pilot wave; t is a parameter.

**Potential mismatch:** The WdW equation is a constraint (ĤΨ = 0), not a Schrödinger equation. In standard Bohmian mechanics, Ψ evolves via Schrödinger equation. Here, Ψ is "static" with respect to WdW time (which doesn't exist), but the trajectory evolves in t.

**Resolution:** This is a known feature of Bohmian quantum gravity, not a bug. The "static" Ψ is analogous to the energy eigenstate in non-relativistic QM: the wave function is time-independent, but the particle moves (in a stationary state, the particle is at rest; in a superposition, it moves). In the WdW context, all states are "energy eigenstates" (E = 0), and the "motion" is in configuration space, not physical space.

**Faithfulness verdict:** The formalism matches the claim, with the caveat that "static Ψ" means "WdW-constrained," not "time-independent in the Schrödinger sense."

---

## The Minimum Viable Theory: Final Form

**Name:** Bohmian Minisuperspace Cosmology (BMC)

**Scope:** FRW universe with homogeneous scalar field. Single trajectory ontology.

**Equations:**
1. WdW: [-∂²/∂a² + (1/a²)∂²/∂φ² + a⁴V(φ) - ka²]Ψ = 0
2. Guidance: da/dt = ∂S/∂a, dφ/dt = -(1/a²)∂S/∂φ
3. Emergent metric: ds² = -[(da/dt)/(∂S/∂a)]²dt² + a(t)²dΩ₃²
4. Effective Friedmann: (ȧ/a)² = (1/3)[(1/2)(dφ/dt)² + V(φ) + Q]

**Map:** (a(t), φ(t)) → FRW spacetime via (3)

**Invariant:** Functional form of Friedmann equation preserved, Q_Bohm as additive correction

**Toy check:** Numerical integration of (1)-(4) for V(φ) = (1/2)m²φ²

**Null model:** Standard WdW (ensemble prediction), LQC (algebraic correction)

**Obstructions:** Node problem (restricted to node-free Ψ), measure problem (minisuperspace only), time relocation (deferred)

**Faithfulness:** Formalism matches claim; "static Ψ" = WdW-constrained

**Final-truth language:** None. All claims are conditional: "In this model," "if the WdW solution exists," "the toy check tests whether."

---

## What Is Explicitly NOT Claimed

- Geometry is not derived from non-geometric beables (a is geometric)
- Lorentz symmetry is not emergent (it is assumed in the FRW ansatz)
- The problem of time is not dissolved (it is reframed and deferred)
- The model is not a theory of everything (no fermions, no gauge fields, no inhomogeneity)
- The model is not unique or necessary (standard WdW and LQC are alternatives)

---

## Next Smallest Move

Execute the toy check: solve the WdW equation numerically for V(φ) = (1/2)m²φ², integrate the guidance equations, compute Q, verify the effective Friedmann equation.

This is a finite calculation. It either works or it fails. If it fails, the model is modified or abandoned. If it works, the next move is inhomogeneous perturbations (tensor and scalar modes), which is a genuine extension, not a triviality.

The EBP 2.1 ledger for this draft:

| Debt | Status |
|------|--------|
| `needMap` | Retired (explicit, checkable) |
| `needInvariant` | Retired (Friedmann form + Q_Bohm) |
| `needToyCheck` | Active (equations written, execution pending) |
| `needNullModel` | Retired (engaged WdW and LQC, disclaimed uniqueness) |
| `needObstruction` | Retired (filed, resolved by restriction, or downgraded) |
| `needFaithfulnessReview` | Retired (formalism matches claim) |
| `containsFinalTruthClaim` | False (no final-truth language) |

**Status:** Workbench-ready. Promotion pending toy check execution.

This is the difference between the original document and this draft: the original claimed the throne without paying debts. This draft sits on the workbench with one debt remaining, and that debt is executable.
