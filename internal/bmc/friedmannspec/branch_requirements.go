package friedmannspec

// FriedmannBranchRequirement evaluates the readiness contract of a segmented local branch
// for a potential future residual computation.
type FriedmannBranchRequirement struct {
	BranchID       string `json:"branch_id"`
	SourceConfigID string `json:"source_config_id"`

	PhiLocalMonotonic    bool `json:"phi_local_monotonic"`
	AlphaPhiSingleValued bool `json:"alpha_phi_single_valued"`
	MinBranchSamples     int  `json:"min_branch_samples"`
	ActualSamples        int  `json:"actual_samples"`

	ClockRange  float64 `json:"clock_range"`
	LambdaRange float64 `json:"lambda_range"`

	DerivativeReady  bool   `json:"derivative_ready"`
	DerivativeMethod string `json:"derivative_method"`
	DerivativeDebt   string `json:"derivative_debt"`

	NodeContactFree      bool `json:"node_contact_free"`
	QFiniteAwayFromNodes bool `json:"q_finite_away_from_nodes"`

	BranchResidualReadiness string `json:"branch_residual_readiness"`
	Reason                  string `json:"reason"`
}
