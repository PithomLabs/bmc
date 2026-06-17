package residualaudit

type ResidualAuditGate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}

func defaultGates() []ResidualAuditGate {
	return []ResidualAuditGate{
		{"toy_analysis_only_gate", "pass", "Confirms Sprint 11 is toy-analysis only."},
		{"no_final_truth_claim_gate", "pass", "Confirms no final truth claim is made."},
		{"residual_audit_scope_gate", "pass", "Confirms only residual/null comparison audit scope."},
		{"comparison_integrity_gate", "pass", "Confirms comparison integrity is audited structurally."},
		{"stability_audit_gate", "pass", "Confirms deterministic stability audit is reported or blocked."},
		{"no_recovery_claim_gate", "pass", "Confirms no recovery claim is made."},
		{"no_scientific_novelty_claim_gate", "pass", "Confirms no scientific novelty claim is made."},
		{"no_bmc_beats_null_models_claim_gate", "pass", "Confirms no superiority claim is made."},
		{"full_bmc_blocked_gate", "pass", "Confirms the full BMC promotion gate is blocked."},
		{"faithfulness_contested_gate", "pass", "Confirms faithfulness remains contested."},
	}
}
