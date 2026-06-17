package friedmannspec

// FriedmannCandidateMap defines the candidate mapping from BMC variables to classical FRW variables,
// tracking unit, sign, normalization, and clock choice debts.
type FriedmannCandidateMap struct {
	MapID string `json:"map_id"`

	AlphaMeaning  string `json:"alpha_meaning"`
	PhiMeaning    string `json:"phi_meaning"`
	ClockVariable string `json:"clock_variable"`
	ClockScope    string `json:"clock_scope"`

	CandidateScaleFactor             string `json:"candidate_scale_factor"`
	CandidateHubbleDefinition         string `json:"candidate_hubble_definition"`
	CandidateEnergyDensityDefinition string `json:"candidate_energy_density_definition"`

	UnitConventionStatus string `json:"unit_convention_status"`
	SignConventionStatus string `json:"sign_convention_status"`
	NormalizationStatus  string `json:"normalization_status"`
	ClockChoiceDebt      string `json:"clock_choice_debt"`
	ClassicalTargetDebt  string `json:"classical_target_debt"`

	Status string `json:"status"`
	Reason string `json:"reason"`
}
