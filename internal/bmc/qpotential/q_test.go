package qpotential_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/qpotential"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

func TestPlaneWaveQApproximatelyZero(t *testing.T) {
	pw := wave.NewPlaneWave(1.0, 1.0)
	tolerance := 1e-9
	h := 1e-4

	// Sample Q at a few points
	testPoints := []model.MiniState{
		{Alpha: 0.0, Phi: 0.0},
		{Alpha: 1.0, Phi: -1.0},
		{Alpha: -2.5, Phi: 3.2},
	}

	for _, pt := range testPoints {
		qVal := qpotential.Q(pw, pt.Alpha, pt.Phi, h)
		if math.Abs(qVal) > tolerance {
			t.Errorf("Expected Q ≈ 0 at (%f, %f), got %e", pt.Alpha, pt.Phi, qVal)
		}
		if qVal != 0.0 {
			t.Errorf("Expected Q to be exactly 0 (constant amplitude), got %e", qVal)
		}
	}
}
