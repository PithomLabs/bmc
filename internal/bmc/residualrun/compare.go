package residualrun

// ResidualNullComparison holds target/null diagnostic comparison data.
type ResidualNullComparison struct {
	ComparisonID         string   `json:"comparison_id"`
	TargetResidualIDs    []string `json:"target_residual_ids"`
	NullModelIDs         []string `json:"null_model_ids"`
	MetricsCompared      []string `json:"metrics_compared"`
	ComparisonComputed   bool     `json:"comparison_computed"`
	InterpretationStatus string   `json:"interpretation_status"`
	Reason               string   `json:"reason"`
}
