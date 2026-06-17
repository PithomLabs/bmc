package nullspec

// NullModelSpec defines the specification for a required null model control.
type NullModelSpec struct {
	NullModelID                     string   `json:"null_model_id"`
	Name                            string   `json:"name"`
	Purpose                         string   `json:"purpose"`
	ControlsFor                     []string `json:"controls_for"`
	RequiredInputs                  []string `json:"required_inputs"`
	RequiredMetrics                 []string `json:"required_metrics"`
	RequiredBeforeResidualPromotion bool     `json:"required_before_residual_promotion"`
	Status                          string   `json:"status"`
	Reason                          string   `json:"reason"`
}
