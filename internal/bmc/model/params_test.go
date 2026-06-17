package model_test

import (
	"math"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/model"
)

func TestValidateDefaultParams(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	if err := params.Validate(); err != nil {
		t.Fatalf("Default params should be valid, but got error: %v", err)
	}
}

func TestValidateRejectsConstraintViolation(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	// Constraint omega^2 = k^2 must hold.
	// With k = 1.0, omega = 2.0, the difference is 3.0, which exceeds tolerance.
	params.Omega = 2.0
	if err := params.Validate(); err == nil {
		t.Error("Expected error for Wheeler-DeWitt constraint violation, but got nil")
	}
}

func TestValidateRejectsInvalidStepCount(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	params.Steps = 0
	if err := params.Validate(); err == nil {
		t.Error("Expected error for Steps = 0, but got nil")
	}

	params.Steps = -10
	if err := params.Validate(); err == nil {
		t.Error("Expected error for Steps < 0, but got nil")
	}
}

func TestValidateRejectsNonfiniteParameters(t *testing.T) {
	testCases := []struct {
		name      string
		modifyFn  func(*model.PlaneWaveParams)
		expectErr string
	}{
		{
			name: "K is NaN",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.K = math.NaN()
			},
			expectErr: "K must be finite",
		},
		{
			name: "Omega is Inf",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.Omega = math.Inf(1)
			},
			expectErr: "Omega must be finite",
		},
		{
			name: "Alpha0 is NaN",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.Alpha0 = math.NaN()
			},
			expectErr: "Alpha0 must be finite",
		},
		{
			name: "Phi0 is Inf",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.Phi0 = math.Inf(-1)
			},
			expectErr: "Phi0 must be finite",
		},
		{
			name: "LambdaStep is negative",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.LambdaStep = -0.1
			},
			expectErr: "LambdaStep must be finite and positive",
		},
		{
			name: "Tolerance is NaN",
			modifyFn: func(p *model.PlaneWaveParams) {
				p.Tolerance = math.NaN()
			},
			expectErr: "Tolerance must be finite and positive",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := model.DefaultPlaneWaveParams()
			tc.modifyFn(&p)
			err := p.Validate()
			if err == nil {
				t.Fatal("Expected error for non-finite parameter, got nil")
			}
		})
	}
}

func TestValidateRejectsNegativeLambdaStep(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	params.LambdaStep = -0.5
	if err := params.Validate(); err == nil {
		t.Error("Expected error for negative LambdaStep, but got nil")
	}
}

func TestValidateRejectsZeroLambdaStep(t *testing.T) {
	params := model.DefaultPlaneWaveParams()
	params.LambdaStep = 0.0
	if err := params.Validate(); err == nil {
		t.Error("Expected error for zero LambdaStep, but got nil")
	}
}
