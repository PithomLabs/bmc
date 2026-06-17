package nullrun

// Allowed RunStatus values
const (
	RunStatusDiagnosticsGenerated = "diagnostics_generated"
	RunStatusBlocked              = "blocked"
	RunStatusDeferred             = "deferred"
)

// Allowed DiagnosticProvenance values
const (
	ProvenanceComputed            = "computed_from_existing_bmc_diagnostics"
	ProvenanceDeterministicFixture = "deterministic_fixture"
	ProvenanceSourceArtifact      = "source_artifact_summary"
	ProvenanceBlocked             = "blocked"
	ProvenanceDeferred            = "deferred"
)

// Allowed DiagnosticStatus values
const (
	DiagStatusFinite       = "finite"
	DiagStatusNonfinite    = "nonfinite"
	DiagStatusNodeBlocked  = "node_blocked"
	DiagStatusClockFragile = "clock_fragile"
	DiagStatusLocalOnly    = "local_only"
	DiagStatusNotAvailable = "not_available"
)

// Allowed InterpretationStatus values
const (
	InterpretDiagComparisonOnly               = "diagnostic_comparison_only"
	InterpretMixedDiagnostics                 = "mixed_diagnostics"
	InterpretInsufficientSeparation           = "insufficient_separation"
	InterpretTargetNullSeparationCandidate    = "target_null_separation_candidate_unpromoted"
	InterpretBlockedByClockFragility          = "blocked_by_clock_fragility"
	InterpretBlockedByNodeObstruction         = "blocked_by_node_obstruction"
	InterpretBlockedByNoComparableDiagnostics = "blocked_by_no_comparable_null_diagnostics"
)
