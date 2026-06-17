package model

// MiniState represents a point in the symmetry-reduced flat FRW minisuperspace configuration space.
type MiniState struct {
	Alpha float64 `json:"alpha"` // ln(a)
	Phi   float64 `json:"phi"`   // homogeneous scalar field
}

// TrajectoryPoint represents a state along a Bohmian parameter/time trajectory.
type TrajectoryPoint struct {
	Lambda float64   `json:"lambda"` // trajectory parameter (analogous to time)
	State  MiniState `json:"state"`
}

// Trajectory holds a sequence of points representing a relational cosmological history.
type Trajectory struct {
	Points []TrajectoryPoint `json:"points"`
}

// Complex is a type alias for Go's native complex128 type.
type Complex = complex128

// CheckStatus defines the status of a check, complying with EBP 2.1 split gates.
type CheckStatus string

const (
	StatusPass      CheckStatus = "pass"
	StatusFail      CheckStatus = "fail"
	StatusDeferred  CheckStatus = "deferred"
	StatusContested CheckStatus = "contested"
)

// CheckResult represents the outcome of a particular check or validation.
type CheckResult struct {
	Status                     CheckStatus `json:"status"`
	Pass                       bool        `json:"pass"`
	Reason                     string      `json:"reason,omitempty"`
	MaxAbsResidual             *float64    `json:"max_abs_residual,omitempty"`
	MaxAbsQ                    *float64    `json:"max_abs_q,omitempty"`
	Finite                     *bool       `json:"finite,omitempty"`
	PointsCount                *int        `json:"points,omitempty"`
	Variable                   *string     `json:"variable,omitempty"`
	AnalyticResidualMagnitude  *float64    `json:"analytic_residual_magnitude,omitempty"`
	NumericalResidualMagnitude *float64    `json:"numerical_residual_magnitude,omitempty"`
	NumericalResidualTolerance *float64    `json:"numerical_residual_tolerance,omitempty"`
	NumericalResidualStatus    *string     `json:"numerical_residual_status,omitempty"`
	NumericalResidualAuthority *string     `json:"numerical_residual_authority,omitempty"`
}

// ObstructionSeverity defines how critical an obstruction is.
type ObstructionSeverity string

const (
	SeverityInfo    ObstructionSeverity = "info"
	SeverityWarning ObstructionSeverity = "warning"
	SeverityBlocker ObstructionSeverity = "blocker"
)

// Obstruction represents an active or potential blocker identified under EBP 2.1.
type Obstruction struct {
	Name        string              `json:"name"`
	Applies     bool                `json:"applies"`
	Severity    ObstructionSeverity `json:"severity"`
	Evidence    string              `json:"evidence"`
	Consequence string              `json:"consequence"`
	Status      CheckStatus         `json:"status"`
}
