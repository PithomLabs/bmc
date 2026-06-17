package model

import (
	"errors"
	"fmt"
	"math"
)

// PlaneWaveParams configures the plane-wave control model parameters.
type PlaneWaveParams struct {
	K          float64 `json:"k"`
	Omega      float64 `json:"omega"`
	Alpha0     float64 `json:"alpha0"`
	Phi0       float64 `json:"phi0"`
	LambdaStep float64 `json:"lambda_step"`
	Steps      int     `json:"steps"`
	Tolerance  float64 `json:"tolerance"`
}

// DefaultPlaneWaveParams returns a standard, deterministic set of parameters for plane-wave control.
func DefaultPlaneWaveParams() PlaneWaveParams {
	return PlaneWaveParams{
		K:          1.0,
		Omega:      1.0,
		Alpha0:     0.0,
		Phi0:       0.0,
		LambdaStep: 0.1,
		Steps:      100,
		Tolerance:  1e-9,
	}
}

// Validate checks that the plane-wave parameters are valid, finite, and satisfy the WdW constraint.
func (p PlaneWaveParams) Validate() error {
	// 1. Check for non-finite values
	if math.IsNaN(p.K) || math.IsInf(p.K, 0) {
		return errors.New("parameter K must be finite")
	}
	if math.IsNaN(p.Omega) || math.IsInf(p.Omega, 0) {
		return errors.New("parameter Omega must be finite")
	}
	if math.IsNaN(p.Alpha0) || math.IsInf(p.Alpha0, 0) {
		return errors.New("parameter Alpha0 must be finite")
	}
	if math.IsNaN(p.Phi0) || math.IsInf(p.Phi0, 0) {
		return errors.New("parameter Phi0 must be finite")
	}
	if math.IsNaN(p.LambdaStep) || math.IsInf(p.LambdaStep, 0) || p.LambdaStep <= 0 {
		return errors.New("parameter LambdaStep must be finite and positive")
	}
	if math.IsNaN(p.Tolerance) || math.IsInf(p.Tolerance, 0) || p.Tolerance <= 0 {
		return errors.New("parameter Tolerance must be finite and positive")
	}

	// 2. Check steps count
	if p.Steps <= 0 {
		return fmt.Errorf("steps must be greater than zero, got %d", p.Steps)
	}

	// 3. WdW constraint check: ω² = k²
	lhs := p.Omega * p.Omega
	rhs := p.K * p.K
	diff := math.Abs(lhs - rhs)
	if diff > p.Tolerance {
		return fmt.Errorf("Wheeler-DeWitt constraint violation: omega^2 = k^2 must hold within tolerance %e (omega^2=%f, k^2=%f, diff=%e)", p.Tolerance, lhs, rhs, diff)
	}

	return nil
}

// SuperpositionParams configures the two-plane-wave superposition model parameters.
type SuperpositionParams struct {
	C1Real       float64 `json:"c1_real"`
	C1Imag       float64 `json:"c1_imag"`
	K1           float64 `json:"k1"`
	Omega1       float64 `json:"omega1"`
	C2Real       float64 `json:"c2_real"`
	C2Imag       float64 `json:"c2_imag"`
	K2           float64 `json:"k2"`
	Omega2       float64 `json:"omega2"`
	Alpha0       float64 `json:"alpha0"`
	Phi0         float64 `json:"phi0"`
	LambdaStep   float64 `json:"lambda_step"`
	Steps        int     `json:"steps"`
	Tolerance    float64 `json:"tolerance"`
	NodeThresh   float64 `json:"node_threshold"`
	MaxPhaseGrad float64 `json:"max_phase_gradient"`
}

// DefaultSuperpositionSafeParams returns the standard default parameters for a safe, non-node superposition.
func DefaultSuperpositionSafeParams() SuperpositionParams {
	return SuperpositionParams{
		C1Real:       1.0,
		C1Imag:       0.0,
		K1:           1.0,
		Omega1:       1.0,
		C2Real:       0.5,
		C2Imag:       0.0,
		K2:           2.0,
		Omega2:       -2.0,
		Alpha0:       0.0,
		Phi0:         0.0,
		LambdaStep:   0.05,
		Steps:        200,
		Tolerance:    1e-9,
		NodeThresh:   1e-5,
		MaxPhaseGrad: 100.0,
	}
}

// DefaultSuperpositionNodeProbeParams returns parameters set up to intentionally start on/near a node.
func DefaultSuperpositionNodeProbeParams() SuperpositionParams {
	return SuperpositionParams{
		C1Real:       1.0,
		C1Imag:       0.0,
		K1:           1.0,
		Omega1:       1.0,
		C2Real:       -1.0,
		C2Imag:       0.0,
		K2:           2.0,
		Omega2:       -2.0,
		Alpha0:       0.0,
		Phi0:         0.0,
		LambdaStep:   0.05,
		Steps:        200,
		Tolerance:    1e-9,
		NodeThresh:   1e-5,
		MaxPhaseGrad: 100.0,
	}
}

// Validate checks that the superposition parameters are valid, finite, and satisfy component WdW constraints.
func (p SuperpositionParams) Validate() error {
	// 1. Check for non-finite values in all fields
	fields := []struct {
		val  float64
		name string
	}{
		{p.C1Real, "c1_real"}, {p.C1Imag, "c1_imag"}, {p.K1, "k1"}, {p.Omega1, "omega1"},
		{p.C2Real, "c2_real"}, {p.C2Imag, "c2_imag"}, {p.K2, "k2"}, {p.Omega2, "omega2"},
		{p.Alpha0, "alpha0"}, {p.Phi0, "phi0"}, {p.LambdaStep, "lambda_step"},
		{p.Tolerance, "tolerance"}, {p.NodeThresh, "node_threshold"}, {p.MaxPhaseGrad, "max_phase_gradient"},
	}

	for _, f := range fields {
		if math.IsNaN(f.val) || math.IsInf(f.val, 0) {
			return fmt.Errorf("parameter %s must be finite", f.name)
		}
	}

	// 2. Positivity checks
	if p.LambdaStep <= 0 {
		return errors.New("parameter LambdaStep must be positive")
	}
	if p.Tolerance <= 0 {
		return errors.New("parameter Tolerance must be positive")
	}
	if p.NodeThresh <= 0 {
		return errors.New("parameter NodeThresh must be positive")
	}
	if p.MaxPhaseGrad <= 0 {
		return errors.New("parameter MaxPhaseGrad must be positive")
	}

	// 3. Steps count
	if p.Steps <= 0 {
		return fmt.Errorf("steps must be greater than zero, got %d", p.Steps)
	}

	// 4. Component 1 WdW constraint: omega1^2 = k1^2
	diff1 := math.Abs(p.Omega1*p.Omega1 - p.K1*p.K1)
	if diff1 > p.Tolerance {
		return fmt.Errorf("component 1 Wheeler-DeWitt constraint violation: omega1^2 = k1^2 must hold within tolerance %e (omega1^2=%f, k1^2=%f, diff=%e)", p.Tolerance, p.Omega1*p.Omega1, p.K1*p.K1, diff1)
	}

	// 5. Component 2 WdW constraint: omega2^2 = k2^2
	diff2 := math.Abs(p.Omega2*p.Omega2 - p.K2*p.K2)
	if diff2 > p.Tolerance {
		return fmt.Errorf("component 2 Wheeler-DeWitt constraint violation: omega2^2 = k2^2 must hold within tolerance %e (omega2^2=%f, k2^2=%f, diff=%e)", p.Tolerance, p.Omega2*p.Omega2, p.K2*p.K2, diff2)
	}

	return nil
}

