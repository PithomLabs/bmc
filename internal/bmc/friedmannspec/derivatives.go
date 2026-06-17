package friedmannspec

// DerivativeReadinessCheck specifies numerical derivative checks on a local branch,
// defining necessary exclusions and step-size sensitivity.
type DerivativeReadinessCheck struct {
	BranchID string `json:"branch_id"`
	Method   string `json:"method"`

	ExcludesEndpoints                 bool `json:"excludes_endpoints"`
	ExcludesTurningPointNeighborhoods bool `json:"excludes_turning_point_neighborhoods"`
	ExcludesNearNodePoints            bool `json:"excludes_near_node_points"`

	MinSamplesRequired int `json:"min_samples_required"`
	SamplesAvailable   int `json:"samples_available"`

	StepSensitivityStatus string `json:"step_sensitivity_status"`
	Status                string `json:"status"`
	Reason                string `json:"reason"`
}
