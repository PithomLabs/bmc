package nullspec

// NullModelMetricContract defines the future metrics to compare trajectories and control baselines.
type NullModelMetricContract struct {
	MetricID                        string   `json:"metric_id"`
	Description                     string   `json:"description"`
	AppliesTo                       []string `json:"applies_to"`
	RequiredBeforeResidualPromotion bool     `json:"required_before_residual_promotion"`
	Status                          string   `json:"status"`
	Reason                          string   `json:"reason"`
}
