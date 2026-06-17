package residualrun

// Allowed EligibilityStatus values
const (
	EligibilityEligibleLocalBranch           = "eligible_local_branch"
	EligibilityBlockedByNodeObstruction      = "blocked_by_node_obstruction"
	EligibilityBlockedByClockFragility       = "blocked_by_clock_fragility"
	EligibilityBlockedByNonfiniteTrajectory  = "blocked_by_nonfinite_trajectory"
	EligibilityBlockedByDerivativeUnreadiness = "blocked_by_derivative_unreadiness"
	EligibilitySourceUnavailable             = "source_unavailable"
)

// Allowed ResidualStatus values
const (
	ResidualStatusGenerated                 = "candidate_residual_diagnostics_generated"
	ResidualStatusInputBlocked              = "residual_input_blocked"
	ResidualStatusNonfinite                 = "residual_nonfinite"
	ResidualStatusBlockedByClockFragility   = "blocked_by_clock_fragility"
	ResidualStatusBlockedByNodeObstruction  = "blocked_by_node_obstruction"
	ResidualStatusBlockedByConventionDebt   = "blocked_by_convention_debt"
	ResidualStatusSourceUnavailable         = "source_unavailable"
)

// Allowed ResidualProvenance values
const (
	ProvenanceComputed            = "computed_from_bmc0a_local_branch"
	ProvenanceDeterministicFixture = "deterministic_fixture"
	ProvenanceSourceArtifact      = "source_artifact_summary"
	ProvenanceBlocked             = "blocked"
)

// Allowed InterpretationStatus values
const (
	InterpretDiagComparisonOnly                 = "diagnostic_comparison_only"
	InterpretMixedResidualDiagnostics           = "mixed_residual_diagnostics"
	InterpretInsufficientTargetNullSeparation   = "insufficient_target_null_separation"
	InterpretTargetNullResidualSeparationCandidate = "target_null_residual_separation_candidate_unpromoted"
	InterpretBlockedByNoComparableNullDiagnostics = "blocked_by_no_comparable_null_diagnostics"
	InterpretBlockedByConventionDebt             = "blocked_by_convention_debt"
	InterpretBlockedByClockFragility             = "blocked_by_clock_fragility"
	InterpretBlockedByNodeObstruction            = "blocked_by_node_obstruction"
	InterpretBlockedByNoEligibleLocalBranch      = "blocked_by_no_eligible_local_branch"
	InterpretBlockedByMissingResidualInputs      = "blocked_by_missing_residual_inputs"
)
