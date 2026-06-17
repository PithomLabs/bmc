package residualaudit

import (
	"strings"

	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/residualrun"
)

func buildComparisonAudits(residualRep *residualrun.ResidualRunReport, nullRep *nullrun.NullRunReport) []ResidualComparisonAudit {
	if len(residualRep.ResidualNullComparisons) == 0 {
		return []ResidualComparisonAudit{{
			AuditID:              "audit_missing_residual_null_comparison",
			SourceComparisonID:   "",
			AuditComputed:        false,
			AuditStatus:          AuditStatusComparisonMissing,
			AuditProvenance:      ProvenanceBlocked,
			Findings:             []string{"No residual/null comparison record was present."},
			InterpretationStatus: InterpretComparisonIntegrityFailed,
			Notes:                "Comparison audit blocked because no comparison record exists.",
		}}
	}

	computedResiduals := map[string]bool{}
	for _, diag := range residualRep.CandidateResidualDiagnostics {
		if diag.ResidualComputed {
			computedResiduals[diag.ResidualID] = true
		}
	}
	computedNulls := map[string]bool{}
	for _, run := range nullRep.NullModelRuns {
		if run.RunStatus == nullrun.RunStatusDiagnosticsGenerated {
			computedNulls[run.NullModelID] = true
		}
	}

	audits := []ResidualComparisonAudit{}
	for _, comp := range residualRep.ResidualNullComparisons {
		audit := ResidualComparisonAudit{
			AuditID:              "audit_" + sanitizeID(comp.ComparisonID),
			SourceComparisonID:   comp.ComparisonID,
			AuditComputed:        true,
			AuditStatus:          AuditStatusComparisonAudited,
			AuditProvenance:      ProvenanceDerivedFromFileRead,
			TargetResidualIDs:    append([]string{}, comp.TargetResidualIDs...),
			NullModelIDs:         append([]string{}, comp.NullModelIDs...),
			MetricsAudited:       append([]string{}, comp.MetricsCompared...),
			Findings:             []string{},
			InterpretationStatus: InterpretComparisonStructurallyHonest,
			Notes:                "Comparison audit checks structure only and remains unpromoted.",
		}
		if len(comp.MetricsCompared) == 0 || len(comp.TargetResidualIDs) == 0 || len(comp.NullModelIDs) == 0 {
			audit.AuditComputed = false
			audit.AuditStatus = AuditStatusComparisonDecorative
			audit.AuditProvenance = ProvenanceBlocked
			audit.InterpretationStatus = InterpretComparisonIntegrityFailed
			audit.Findings = append(audit.Findings, "Comparison record is missing metrics, targets, or null references.")
		}
		for _, id := range comp.TargetResidualIDs {
			if !computedResiduals[id] {
				audit.AuditComputed = false
				audit.AuditStatus = AuditStatusComparisonDecorative
				audit.AuditProvenance = ProvenanceBlocked
				audit.InterpretationStatus = InterpretComparisonIntegrityFailed
				audit.Findings = append(audit.Findings, "Comparison target does not reference a computed residual diagnostic.")
			}
		}
		for _, id := range comp.NullModelIDs {
			if !computedNulls[id] {
				audit.AuditComputed = false
				audit.AuditStatus = AuditStatusComparisonBlocked
				audit.AuditProvenance = ProvenanceBlocked
				audit.InterpretationStatus = InterpretBlockedByMissingNullInputs
				audit.Findings = append(audit.Findings, "Comparison null reference does not reference a generated null diagnostic.")
			}
		}
		if audit.AuditComputed {
			audit.Findings = append(audit.Findings, "Comparison targets computed residual diagnostics and generated null diagnostics.")
		}
		audits = append(audits, audit)
	}
	return audits
}

func summarizeInterpretation(comps []ResidualComparisonAudit, stabs []ResidualStabilityAudit) string {
	if len(comps) == 0 {
		return InterpretComparisonIntegrityFailed
	}
	hasFailed := false
	hasMixed := false
	for _, comp := range comps {
		if !comp.AuditComputed || comp.InterpretationStatus == InterpretComparisonIntegrityFailed {
			hasFailed = true
		}
	}
	for _, stab := range stabs {
		if stab.StabilityStatus == StabilityUnstable || stab.StabilityStatus == StabilityNonfinite {
			return InterpretComparisonUnstable
		}
		if stab.StabilityStatus == StabilitySensitive || !stab.StabilityComputed {
			hasMixed = true
		}
	}
	if hasFailed {
		return InterpretComparisonIntegrityFailed
	}
	if hasMixed {
		return InterpretComparisonStabilityMixed
	}
	return InterpretComparisonStructurallyHonest
}

func sanitizeID(id string) string {
	id = strings.ReplaceAll(id, "-", "_")
	id = strings.ReplaceAll(id, ".", "_")
	if id == "" {
		return "comparison"
	}
	return id
}
