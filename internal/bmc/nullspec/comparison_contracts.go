package nullspec

// FutureNullComparisonContract defines the comparison plans to evaluate BMC trajectories against null baselines.
type FutureNullComparisonContract struct {
	ComparisonID       string   `json:"comparison_id"`
	BaselineArtifact   string   `json:"baseline_artifact"`
	NullModelIDs       []string `json:"null_model_ids"`
	Metrics            []string `json:"metrics"`
	ComparisonComputed bool     `json:"comparison_computed"`
	Status             string   `json:"status"`
	Reason             string   `json:"reason"`
}
