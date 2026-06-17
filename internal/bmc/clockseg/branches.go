package clockseg

import "github.com/PithomLabs/bmc/internal/bmc/model"

// ClockTurningPoint represents a point in the trajectory where phi monotonicity
// changes or near-zero dphi/dlambda velocity is encountered.
type ClockTurningPoint struct {
	Index       int     `json:"index"`
	Lambda      float64 `json:"lambda"`
	Alpha       float64 `json:"alpha"`
	Phi         float64 `json:"phi"`
	DPhiDLambda float64 `json:"dphi_dlambda"`
}

// ClockSegment represents a continuous interval of the trajectory where the
// phi clock is strictly monotonic (increasing or decreasing).
type ClockSegment struct {
	StartIndex  int     `json:"start_index"`
	EndIndex    int     `json:"end_index"`
	StartLambda float64 `json:"start_lambda"`
	EndLambda   float64 `json:"end_lambda"`
	Direction   int     `json:"direction"` // +1 for increasing phi, -1 for decreasing phi
}

// LocalRelationBranch represents the alpha(phi) relational branch analysis
// extracted from a single ClockSegment.
type LocalRelationBranch struct {
	Segment          ClockSegment            `json:"segment"`
	Samples          int                     `json:"samples"`
	LambdaRange      float64                 `json:"lambda_range"`
	ClockRange       float64                 `json:"clock_range"`
	SingleValued     bool                    `json:"single_valued"`
	MaxAbsNoise      float64                 `json:"max_abs_noise"`
	ValidationPassed bool                    `json:"validation_passed"`
	Reason           string                  `json:"reason"`
	Points           []model.TrajectoryPoint `json:"points,omitempty"`
}

// ClockIndependentDiagnostic contains physical metrics along the trajectory
// that do not depend on the global monotonicity or choice of relational clock phi.
type ClockIndependentDiagnostic struct {
	PathLengthInConfigurationSpace float64 `json:"path_length_in_configuration_space"`
	TotalLambdaInterval            float64 `json:"total_lambda_interval"`
	NumValidTrajectoryPoints       int     `json:"num_valid_trajectory_points"`
	NumClockSegments               int     `json:"num_clock_segments"`
	NumTurningPoints               int     `json:"num_turning_points"`
	MinAmplitudeR                  float64 `json:"min_amplitude_r"`
	MaxAbsQAwayFromNodes           float64 `json:"max_abs_q_away_from_nodes"`
	MaxPhaseGradient               float64 `json:"max_phase_gradient"`
	NodeContactFree                bool    `json:"node_contact_free"`
	TrajectoryFiniteness           bool    `json:"trajectory_finiteness"`
}
