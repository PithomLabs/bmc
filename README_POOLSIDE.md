Based on my analysis of the codebase at `/home/chaschel/Documents/go/bmc`, here's what this project is all about:

## Project Overview

This is the **Bohmian Minisuperspace Cosmology (BMC)** project - a research codebase implementing a "toy model" for quantum cosmological trajectory analysis using Bohmian mechanics under an EBP (Elephant Protocol) governance framework.

## Core Purpose

The project builds a **diagnostic benchmark** (not a proof machine) for testing whether a small Bohmian minisuperspace model can support safe relational-dynamics diagnostics. It deliberately avoids overclaiming quantum gravity, Friedmann recovery, or solutions to the "problem of time."

## Technical Architecture

### Go Implementation (`internal/bmc/`)
The codebase implements:

1. **Wavefunction Models**:
   - Plane wave control (`Ψ = exp(i(kα + ωφ))`)
   - Two-plane-wave superposition with node detection

2. **Physics Computations**:
   - Wheeler-DeWitt equation residual checking
   - Bohmian guidance trajectory integration (Euler and RK4 steppers)
   - Quantum potential calculation along trajectories
   - Clock monotonicity detection (φ as relational time)

3. **CLI Tool** (`cmd/ptw-bmc`):
   - `run` - Execute model profiles
   - `validate` - Schema and EBP constraint validation
   - `summarize` - Human-readable report summaries
   - `audit` - Numerical robustness analysis
   - `diagnose-clock`, `segment-clock` - Clock diagnostics
   - `prior-art-boundary`, `run-nullmodels`, `run-residuals` - Sprint-specific workflows

4. **Multi-sprint Structure**: Each sprint adds a specific analysis layer:
   - Sprint 1-3: Plane-wave and superposition core + robustness
   - Sprint 6: Friedmann-residual specification (deferred)
   - Sprint 7: Null-model scaffold
   - Sprint 8: Prior-art boundary (literature audit)
   - Sprint 9: Null-model runner
   - Sprint 10: Candidate local-branch residual runner
   - Sprint 11: Residual/null comparison audit

### Lean Verification (`BMC/BMC/`)
Formal policy contracts preventing overclaiming:
- `ToyReport.lean` - Basic check status types
- `NullModelSpec.lean`, `NullRun.lean`, `ResidualAudit.lean`, `PriorArtBoundary.lean`, `FriedmannSpec.lean` - Gate conditions that **forbid** recovery claims, final-truth language, and scientific novelty until debts are paid

## Key Design Principle: EBP 2.1

Ideas enter freely, but **promotion costs debt**. Required debts include:
- `needMap` - Explicit mapping from beables to physics claims
- `needInvariant` - Invariant quantities that must survive the map
- `needToyCheck` - Actual numerical execution
- `needNullModel` - Comparison against competing approaches
- `needObstruction` - File known blockers honestly
- `needFaithfulnessReview` - Formal model matches intended claims

## Model Variables

- `α = ln(a)` - Log scale factor
- `φ` - Homogeneous scalar field (candidate clock)
- `Ψ = R e^{iS}` - Wheeler-DeWitt wavefunction
- Guidance: `dα/dλ = ∂S/∂α`, `dφ/dλ = -∂S/∂φ`

## Safety Enforcement

The codebase uses **multiple layers** to prevent overclaiming:
1. Runtime validation in Go rejects forbidden phrases
2. Lean theorems prove gate conditions cannot be satisfied with recovery claims
3. Reports include mandatory warnings disclaiming toy-only scope
4. All promotion gates are set to "blocked" until debts retire