package wave

import (
	"math/cmplx"
)

// WaveFunction defines the contract for minipspace wavefunction evaluations.
type WaveFunction interface {
	Psi(alpha, phi float64) complex128
}

// PlaneWave represents the analytic control wavefunction Ψ(α,φ) = exp(i(kα + ωφ)).
type PlaneWave struct {
	K     float64
	Omega float64
}

// NewPlaneWave constructs a PlaneWave with parameters k and omega.
func NewPlaneWave(k, omega float64) PlaneWave {
	return PlaneWave{K: k, Omega: omega}
}

// Psi evaluates the wavefunction at (alpha, phi).
func (pw PlaneWave) Psi(alpha, phi float64) complex128 {
	phase := pw.K*alpha + pw.Omega*phi
	return cmplx.Exp(complex(0, phase))
}

// PsiGrid evaluates the wavefunction on a configuration space grid.
func (pw PlaneWave) PsiGrid(alphaRange, phiRange []float64) [][]complex128 {
	grid := make([][]complex128, len(alphaRange))
	for i, alpha := range alphaRange {
		grid[i] = make([]complex128, len(phiRange))
		for j, phi := range phiRange {
			grid[i][j] = pw.Psi(alpha, phi)
		}
	}
	return grid
}
