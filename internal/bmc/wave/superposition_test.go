package wave_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

func TestSuperpositionSatisfiesWdW(t *testing.T) {
	// Component 1: k1 = 1.0, omega1 = 1.0 -> residual = 0
	// Component 2: k2 = 2.0, omega2 = -2.0 -> residual = 0
	k1, omega1 := 1.0, 1.0
	k2, omega2 := 2.0, -2.0

	compRes := wdw.ComponentResidualsSuperposition(k1, omega1, k2, omega2)
	if len(compRes) != 2 {
		t.Fatalf("Expected 2 component residuals, got %d", len(compRes))
	}

	for _, c := range compRes {
		if math.Abs(c.Residual) > 1e-9 {
			t.Errorf("Component %d WdW constraint violated: residual is %e", c.Component, c.Residual)
		}
	}
}
