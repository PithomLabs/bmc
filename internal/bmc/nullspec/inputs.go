package nullspec

// NullModelInputRequirement specifies the source and availability status of input data from prior artifacts.
type NullModelInputRequirement struct {
	InputID            string   `json:"input_id"`
	SourceArtifact     string   `json:"source_artifact"`
	RequiredFor        []string `json:"required_for"`
	AvailabilityStatus string   `json:"availability_status"`
	Reason             string   `json:"reason"`
}
