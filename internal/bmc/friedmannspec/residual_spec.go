package friedmannspec

// ResidualFormulaCandidate represents a registered candidate formula for future residual computation.
type ResidualFormulaCandidate struct {
	FormulaID string `json:"formula_id"`
	Description string `json:"description"`

	ClassicalTarget     string   `json:"classical_target"`
	RequiredVariables   []string `json:"required_variables"`
	RequiredDerivatives []string   `json:"required_derivatives"`

	ConventionDebt     []string `json:"convention_debt"`
	NullModelsRequired []string `json:"null_models_required"`

	Status string `json:"status"`
	Reason string `json:"reason"`
}

// FriedmannNullModelRequirement specifies a null model required to be verified
// before any future residual computation can be promoted.
type FriedmannNullModelRequirement struct {
	NullModelID                     string `json:"null_model_id"`
	Purpose                         string `json:"purpose"`
	RequiredBeforeResidualPromotion bool   `json:"required_before_residual_promotion"`
	Status                          string `json:"status"`
	Reason                          string `json:"reason"`
}
