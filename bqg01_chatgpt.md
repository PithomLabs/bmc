> **A Bohmian minimum viable theory for quantum cosmology, not full quantum gravity.**

That means: no Standard Model, no black holes, no full spacetime emergence, no claim that geometry is derived from pure non-geometry, no “problem of time solved.” Under EBP 2.1, the idea enters freely, but promotion is blocked until map, invariant, toy check, null model, obstruction, and faithfulness debts are paid. EBP explicitly says debt blocks promotion but does not kill the idea, and promotion is not truth. 

## What I would do differently

### 1. Stop trying to derive a 3D metric from one scalar field

The previous scalar-field-first route is too fragile. One scalar gradient cannot generically build a nondegenerate 3D metric. So I would not begin with “spacetime emerges from a scalar field.”

Instead, I would begin with **Bohmian geometrodynamics in minisuperspace**:

```text
Primitive actual configuration:
q = (a, φ)
```

where `a` is the scale factor and `φ` is a scalar matter clock. This is less radical, but much cleaner. It admits that geometry is still present in the toy model. The new claim becomes:

```text
Not: geometry emerges from nothing.
But: a classical cosmological history may emerge from a Bohmian trajectory guided by a timeless Wheeler-DeWitt wave function.
```

That is smaller, but testable.

### 2. Treat “time” as an internal-clock test, not a solved mystery

The Wheeler-DeWitt setting is naturally relevant because modern Bohmian quantum cosmology work still uses Wheeler-DeWitt equations plus deterministic guidance equations in minisuperspace; one recent paper constructs a Bohmian Hubble parameter from the pilot-wave phase. ([arXiv][1])

But I would not say “Bohmian mechanics solves the problem of time.” I would say:

```text
Candidate claim:
If φ is monotonic along Bohmian trajectories, then φ can function as an internal clock for a(t).
```

That directly competes with Page-Wootters relational time, where observed dynamics can be described through internal clock readings in a stationary global setting. ([APS Link][2])

### 3. Do not compete with string theory, LQG, or causal sets globally

I would use their gaps as **design constraints**, not as targets to defeat.

String theory has landscape/selection issues, so our MVT must have a tiny fixed parameter space, not a huge menu of vacua. The landscape is a real live issue in string cosmology and vacuum-energy discussions. ([arXiv][3])

LQG has long-standing dynamics/classical-limit/constraint issues; Rovelli’s review states that consistency and the correct classical limit are final tests for Hamiltonian-constraint proposals, and a recent LQG paper still describes the Hamiltonian constraint as elusive because graph-changing actions make calculations difficult. ([PMC][4]) So our MVT should not claim to solve full background-independent dynamics. It should only test one explicit minisuperspace Hamiltonian.

Causal set theory has a hard continuum-recovery problem: its core postulate replaces spacetime with locally finite partially ordered sets, but recovering smooth continuum physics remains a major burden. ([Springer Link][5]) So our MVT should not claim discrete-to-continuum recovery.

## The minimum viable theory I would draft

I would call it, cautiously:

```text
Bohmian Constraint Cosmology v0.1
```

Not “Bohmian Quantum Gravity.” That name overclaims too early.

# Bohmian Constraint Cosmology v0.1

## Status

This is a minimum viable theory candidate, not a complete theory of quantum gravity. It is a finite quantum-cosmology bridge model using Bohmian mechanics under a Wheeler-DeWitt constraint.

It does not claim to solve quantum gravity, derive spacetime from nothing, recover the Standard Model, solve black hole information loss, or prove Bohmian mechanics true.

## Core Claim

In a symmetry-reduced cosmological model with configuration variables (q = (a, \phi)), a Wheeler-DeWitt constraint equation plus a Bohmian guidance law may generate actual trajectories (a(\tau), \phi(\tau)). If (\phi) is monotonic along a trajectory, it may serve as an internal clock, allowing an effective cosmological history (a(\phi)). In a suitable semiclassical limit, this history should recover the classical Friedmann equation, while outside that limit it may exhibit quantum-potential corrections.

## Primitive Ontology

The primitive ontology is not full spacetime and not a field on a fully emergent manifold.

The primitive ontology is the actual minisuperspace configuration:

[
q = (a, \phi)
]

where:

* (a) is the cosmological scale factor.
* (\phi) is a scalar matter degree of freedom that may function as an internal clock.
* (\Psi(a,\phi)) is a constraint-satisfying wave function.
* The actual history is a Bohmian trajectory through configuration space.

## Constraint Equation

The wave function satisfies a Wheeler-DeWitt-type constraint:

[
\hat{H}\Psi(a,\phi) = 0
]

The exact Hamiltonian operator, factor ordering, potential (V(\phi)), and boundary conditions must be specified explicitly for each toy run.

## Bohmian Guidance Law

Write:

[
\Psi(a,\phi) = R(a,\phi)e^{iS(a,\phi)/\hbar}
]

The guidance law is defined by the phase gradient:

[
\dot{q}^A = G^{AB}\partial_B S
]

where (G^{AB}) is the minisuperspace metric or its chosen toy-model analogue.

## Map to Cosmology

The map is not a derivation of spacetime from non-geometry.

It is a map from Bohmian configuration trajectories to effective FRW histories:

[
(a(\tau),\phi(\tau)) \mapsto a(\phi)
]

when (\phi) is monotonic.

The effective spacetime metric is then reconstructed only at the symmetry-reduced level:

[
ds^2 = -N(\tau)^2 d\tau^2 + a(\tau)^2 d\Sigma_k^2
]

## Main Invariant

The first invariant is not entropy.

The first invariant is the classical-limit recovery condition:

[
Q / V_{\text{classical}} \to 0
]

where (Q) is the Bohmian quantum potential. In that limit, the Bohmian trajectory must recover the classical Hamilton-Jacobi/Friedmann behavior.

## First Toy Check

Choose one explicit potential (V(\phi)), one factor ordering, and one boundary condition.

Then compute:

1. A solution or approximate solution (\Psi(a,\phi)).
2. The phase (S(a,\phi)).
3. The guidance equations.
4. The trajectory (a(\tau),\phi(\tau)).
5. The relational history (a(\phi)).
6. The quantum-potential correction.
7. Whether the classical Friedmann equation is recovered when (Q) becomes negligible.

## Null Models

The candidate must be compared against:

1. Standard semiclassical Wheeler-DeWitt without Bohmian trajectories.
2. Page-Wootters relational time.
3. Loop quantum cosmology bounce models.
4. A simple WKB clock treatment.

If the Bohmian trajectory adds no operational distinction, the ontology may be explanatory rather than predictive.

## Failure Conditions

The model fails as a minimum viable theory if:

1. No conserved current or usable guidance law exists.
2. (\phi) cannot function as a monotonic internal clock in the chosen regime.
3. No classical Friedmann limit is recovered.
4. Results depend entirely on arbitrary factor ordering or boundary choices.
5. The Bohmian version gives no distinction from simpler null models.
6. The formal model does not faithfully represent the intended physical claim.

## Promotion Status

Unpromoted.

Current status: alive with debt.

The next task is to execute one finite toy check.

## EBP 2.1 debt ledger for this restart

| Debt                     |                    Status | How to retire it                                                                                                                    |
| ------------------------ | ------------------------: | ----------------------------------------------------------------------------------------------------------------------------------- |
| `needMap`                | **Partially payable now** | Map only from `(a, φ)` trajectories to FRW histories. Do not claim full spacetime emergence.                                        |
| `needInvariant`          |              **Narrowed** | Use classical-limit recovery first, not entropy. Check whether `Q/V_classical → 0` gives Friedmann behavior.                        |
| `needToyCheck`           |                **Unpaid** | Choose one potential, solve or approximate `Ψ`, compute trajectories.                                                               |
| `needNullModel`          |                **Unpaid** | Compare against Page-Wootters, WKB Wheeler-DeWitt, and loop quantum cosmology.                                                      |
| `needObstruction`        |                  **Open** | Main obstructions: factor ordering, inner product/measure, clock monotonicity, boundary conditions, empirical indistinguishability. |
| `needFaithfulnessReview` |              **Required** | Check whether the formal toy really supports the claim, or only a weaker mathematical analogy.                                      |

## The key strategic shift

The old version tried to say:

```text
Bohmian field ontology may make spacetime emerge and dissolve the problem of time.
```

The new ruthless version should say:

```text
A Bohmian trajectory law may turn a timeless Wheeler-DeWitt minisuperspace state into a concrete relational cosmological history, provided an internal clock works and the classical Friedmann limit is recovered.
```

That is the minimum viable theory. It is humble enough to test, but still philosophically powerful.

## Next smallest useful move

Start with the simplest possible Phase 0:

```text
Model: flat FRW + massless scalar field
Variables: α = ln(a), φ
Constraint: WdW equation in α, φ
Candidate solution: plane-wave or Gaussian packet in minisuperspace
Guidance law: derive dα/dτ and dφ/dτ from S
Clock test: is φ monotonic?
Classical-limit test: does α(φ) reproduce the expected classical relation?
Null model: compare the same setup with Page-Wootters relational time
```

That is where I would begin. No black holes yet. No full QFT. No Standard Model. No “emergent all of spacetime.” Just one room of the castle, built strongly enough that we can attack it.

[1]: https://arxiv.org/html/2512.18818v2?utm_source=chatgpt.com "Bohmian Quantum Cosmology from the Wheeler-DeWitt ..."
[2]: https://link.aps.org/doi/10.1103/PhysRevD.27.2885?utm_source=chatgpt.com "Dynamics described by stationary observables | Phys. Rev. D"
[3]: https://arxiv.org/abs/hep-th/0511037?utm_source=chatgpt.com "Living in the Multiverse"
[4]: https://pmc.ncbi.nlm.nih.gov/articles/PMC5567241/?utm_source=chatgpt.com "Loop Quantum Gravity - PMC"
[5]: https://link.springer.com/article/10.1007/s41114-019-0023-1?utm_source=chatgpt.com "The causal set approach to quantum gravity - Springer Nature"

