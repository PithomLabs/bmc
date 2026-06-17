package residualaudit

const (
	SchemaVersion = "bmc0a-residual-audit-v0.1"

	AuditStatusComparisonAudited    = "comparison_audited"
	AuditStatusComparisonBlocked    = "comparison_blocked"
	AuditStatusComparisonMissing    = "comparison_missing"
	AuditStatusComparisonDecorative = "comparison_decorative"
	AuditStatusComparisonUnstable   = "comparison_unstable"
	AuditStatusComparisonMixed      = "comparison_mixed"
	AuditStatusSourceUnavailable    = "source_unavailable"

	ProvenanceFileRead            = "file_read"
	ProvenanceDerivedFromFileRead = "derived_from_file_read"
	ProvenanceSourceSummary       = "source_artifact_summary"
	ProvenanceBlocked             = "blocked"
	ProvenanceNotAvailable        = "not_available"

	InterpretDiagnosticAuditOnly              = "diagnostic_audit_only"
	InterpretComparisonStructurallyHonest     = "comparison_integrity_structurally_honest"
	InterpretComparisonIntegrityFailed        = "comparison_integrity_failed"
	InterpretComparisonStabilityMixed         = "comparison_stability_mixed"
	InterpretComparisonUnstable               = "comparison_unstable"
	InterpretInsufficientTargetNullSeparation = "insufficient_target_null_separation"
	InterpretTargetNullSeparationCandidate    = "target_null_separation_candidate_unpromoted"
	InterpretBlockedByMissingResidualInputs   = "blocked_by_missing_residual_inputs"
	InterpretBlockedByMissingNullInputs       = "blocked_by_missing_null_inputs"
	InterpretBlockedBySourceUnavailable       = "blocked_by_source_unavailable"

	PerturbAlphaPoint    = "alpha_point_perturbation"
	PerturbPhiPoint      = "phi_point_perturbation"
	PerturbLambdaSpacing = "lambda_spacing_perturbation"
	PerturbBranchSubset  = "branch_subset_resampling"
	PerturbNone          = "none"

	StabilityStable       = "stable_under_small_perturbation"
	StabilitySensitive    = "sensitive_to_small_perturbation"
	StabilityUnstable     = "unstable_or_ill_conditioned"
	StabilityMissingInput = "blocked_by_missing_inputs"
	StabilityNonfinite    = "blocked_by_nonfinite_values"
	StabilityNotComputed  = "not_computed"
)
