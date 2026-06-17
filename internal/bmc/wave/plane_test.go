package wave_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/wave"
	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

func TestPlaneWaveSatisfiesWdWResidual(t *testing.T) {
	k := 1.0
	omega := 1.0
	tolerance := 1e-9

	// Analytic residual is k^2 - omega^2
	residual := wdw.AnalyticResidualPlaneWave(k, omega)
	if math.Abs(residual) > tolerance {
		t.Fatalf("Expected analytic residual of plane wave to satisfy tolerance, got %e", residual)
	}

	if residual != 0.0 {
		t.Fatalf("Expected analytic residual of plane wave to be exactly 0, got %e", residual)
	}

	// Double-check the wavefunction implementation
	pw := wave.NewPlaneWave(k, omega)
	psiVal := pw.Psi(0.0, 0.0)
	if psiVal != 1.0 {
		t.Errorf("Expected Psi(0,0) = 1, got %v", psiVal)
	}
}
