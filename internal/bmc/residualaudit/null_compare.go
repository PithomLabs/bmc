package residualaudit

type ResidualComparisonAudit struct {
	AuditID              string   `json:"audit_id"`
	SourceComparisonID   string   `json:"source_comparison_id"`
	AuditComputed        bool     `json:"audit_computed"`
	AuditStatus          string   `json:"audit_status"`
	AuditProvenance      string   `json:"audit_provenance"`
	TargetResidualIDs    []string `json:"target_residual_ids"`
	NullModelIDs         []string `json:"null_model_ids"`
	MetricsAudited       []string `json:"metrics_audited"`
	Findings             []string `json:"findings"`
	InterpretationStatus string   `json:"interpretation_status"`
	Notes                string   `json:"notes"`
}
