package audit

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
)

type StepSizeResult struct {
	StepSize                 float64           `json:"step_size"`
	Steps                    int               `json:"steps"`
	EndpointDriftAlpha       *float64          `json:"endpoint_drift_alpha"`
	EndpointDriftAlphaStatus string            `json:"endpoint_drift_alpha_status,omitempty"`
	EndpointDriftAlphaReason string            `json:"endpoint_drift_alpha_reason,omitempty"`
	EndpointDriftPhi         *float64          `json:"endpoint_drift_phi"`
	EndpointDriftPhiStatus   string            `json:"endpoint_drift_phi_status,omitempty"`
	EndpointDriftPhiReason   string            `json:"endpoint_drift_phi_reason,omitempty"`
	MaxAbsQ                  *float64          `json:"max_abs_q"`
	MaxAbsQStatus            string            `json:"max_abs_q_status,omitempty"`
	MaxAbsQReason            string            `json:"max_abs_q_reason,omitempty"`
	MinAmplitudeR            *float64          `json:"min_amplitude_r"`
	MinAmplitudeRStatus      string            `json:"min_amplitude_r_status,omitempty"`
	MinAmplitudeRReason      string            `json:"min_amplitude_r_reason,omitempty"`
	MaxPhaseGrad             *float64          `json:"max_phase_gradient"`
	MaxPhaseGradStatus       string            `json:"max_phase_gradient_status,omitempty"`
	MaxPhaseGradReason       string            `json:"max_phase_gradient_reason,omitempty"`
	ClockMonotonic           bool              `json:"clock_monotonic"`
	TechnicalGateStatus      model.CheckStatus `json:"technical_gate_status"`
}

type ThresholdSensitivityResult struct {
	Threshold           float64           `json:"threshold"`
	Profile             string            `json:"profile"`
	NodeContactFree     model.CheckStatus `json:"node_contact_free"`
	QFiniteAwayFromNodes model.CheckStatus `json:"q_finite_away_from_nodes"`
	PhaseGradientFinite model.CheckStatus `json:"phase_gradient_finite"`
	TechnicalGateName   string            `json:"technical_gate_name"`
	TechnicalGateStatus model.CheckStatus `json:"technical_gate_status"`
}

type PhaseGradientBoundResult struct {
	Bound                      float64           `json:"bound"`
	MaxObservedPhaseGrad       *float64          `json:"max_observed_phase_gradient"`
	MaxObservedPhaseGradStatus string            `json:"max_observed_phase_gradient_status,omitempty"`
	MaxObservedPhaseGradReason string            `json:"max_observed_phase_gradient_reason,omitempty"`
	IsBinding                  bool              `json:"is_binding"`
	PhaseGradientFinite        model.CheckStatus `json:"phase_gradient_finite"`
}

type ParameterPerturbationResult struct {
	C2Real               float64           `json:"c2_real"`
	K2                   float64           `json:"k2"`
	Omega2               float64           `json:"omega2"`
	NodeContact          bool              `json:"node_contact"`
	QFiniteAwayFromNodes model.CheckStatus `json:"q_finite_away_from_nodes"`
	ClockMonotonic       model.CheckStatus `json:"clock_monotonic"`
	SafeGateStatus       model.CheckStatus `json:"safe_gate_status"`
}

type NodeProbeOffsetResult struct {
	OffsetAlpha            float64           `json:"offset_alpha"`
	OffsetPhi              float64           `json:"offset_phi"`
	InitialAmplitude       *float64          `json:"initial_amplitude"`
	InitialAmplitudeStatus string            `json:"initial_amplitude_status,omitempty"`
	InitialAmplitudeReason string            `json:"initial_amplitude_reason,omitempty"`
	ShortCircuitTriggered  bool              `json:"short_circuit_triggered"`
	Integrated             bool              `json:"integrated"`
	NodeContactFree        model.CheckStatus `json:"node_contact_free"`
	PhaseGradientFinite    model.CheckStatus `json:"phase_gradient_finite"`
	QFiniteAwayFromNodes   model.CheckStatus `json:"q_finite_away_from_nodes"`
}

type TechnicalGate struct {
	Name   string            `json:"name"`
	Status model.CheckStatus `json:"status"`
	Reason string            `json:"reason"`
}

type RobustnessReport struct {
	SchemaVersion             string                       `json:"schema_version"`
	ToyAnalysisOnly           bool                         `json:"toy_analysis_only"`
	FinalTruthClaim           bool                         `json:"final_truth_claim"`
	SourceArtifacts           []string                     `json:"source_artifacts"`
	AuditKind                 string                       `json:"audit_kind"`
	StepSizeSweep             []StepSizeResult             `json:"step_size_sweep"`
	NodeThresholdSweep        []ThresholdSensitivityResult `json:"node_threshold_sweep"`
	PhaseGradientBoundSweep   []PhaseGradientBoundResult   `json:"phase_gradient_bound_sweep"`
	ParameterPerturbationSweep []ParameterPerturbationResult `json:"parameter_perturbation_sweep"`
	NodeProbeOffsetSweep      []NodeProbeOffsetResult      `json:"node_probe_offset_sweep"`
	TechnicalGate             TechnicalGate                `json:"technical_gate"`
	RobustnessOutcome         string                       `json:"robustness_outcome"` // stable|mixed|fragile|contested
	PromotionGate             report.PromotionGate         `json:"promotion_gate"`
	EbpDebt                   report.EbpDebt               `json:"ebp_debt"`
	Warnings                  []string                     `json:"warnings"`
}

// ReadRobustnessReport reads and decodes a report strictly rejecting unknown fields.
func ReadRobustnessReport(path string) (*RobustnessReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep RobustnessReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}
	return &rep, nil
}

// WriteJSON serializes the robustness report with pretty-printing.
func WriteJSON(rep *RobustnessReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// validateOptionalMetric checks that an optional numeric metric is finite if present,
// and rejects null/nil or sentinel -1.0 values unless they are paired with a status/reason.
func validateOptionalMetric(val *float64, status, reason, name string) []report.ValidationError {
	var errs []report.ValidationError

	if val == nil || *val == -1.0 {
		if status == "" || reason == "" {
			errs = append(errs, report.ValidationError{
				Field:    name,
				Message:  fmt.Sprintf("unavailable metric %s must be paired with explicit status and reason", name),
				Severity: report.ValidationFail,
			})
		}
		return errs
	}

	v := *val
	if math.IsNaN(v) || math.IsInf(v, 0) {
		errs = append(errs, report.ValidationError{
			Field:    name,
			Message:  fmt.Sprintf("metric %s contains non-finite value", name),
			Severity: report.ValidationFail,
		})
	}

	return errs
}

func float64Ptr(v float64) *float64 {
	return &v
}

func formatFloatPtr(p *float64, fmtStr string) string {
	if p == nil {
		return "null"
	}
	return fmt.Sprintf(fmtStr, *p)
}

// ValidateRobustnessReport performs schema and EBP verification on the robustness report.
func ValidateRobustnessReport(r *RobustnessReport) []report.ValidationError {
	var errors []report.ValidationError

	if r.SchemaVersion != "bmc0a-superposition-robustness-v0.1" {
		errors = append(errors, report.ValidationError{
			Field:    "schema_version",
			Message:  fmt.Sprintf("unsupported schema version: %s", r.SchemaVersion),
			Severity: report.ValidationFail,
		})
	}

	if !r.ToyAnalysisOnly {
		errors = append(errors, report.ValidationError{
			Field:    "toy_analysis_only",
			Message:  "EBP 2.1 requires toy_analysis_only to be true for BMC-0A audits",
			Severity: report.ValidationFail,
		})
	}

	if r.FinalTruthClaim {
		errors = append(errors, report.ValidationError{
			Field:    "final_truth_claim",
			Message:  "EBP 2.1 promotion blocker: final_truth_claim must be false",
			Severity: report.ValidationFail,
		})
	}

	if r.AuditKind != "numerical_robustness_convergence" {
		errors = append(errors, report.ValidationError{
			Field:    "audit_kind",
			Message:  "incorrect audit kind",
			Severity: report.ValidationFail,
		})
	}

	if len(r.StepSizeSweep) == 0 {
		errors = append(errors, report.ValidationError{
			Field:    "step_size_sweep",
			Message:  "empty step size sweep",
			Severity: report.ValidationFail,
		})
	}

	if len(r.NodeThresholdSweep) == 0 {
		errors = append(errors, report.ValidationError{
			Field:    "node_threshold_sweep",
			Message:  "empty node threshold sweep",
			Severity: report.ValidationFail,
		})
	}

	if len(r.PhaseGradientBoundSweep) == 0 {
		errors = append(errors, report.ValidationError{
			Field:    "phase_gradient_bound_sweep",
			Message:  "empty phase bound sweep",
			Severity: report.ValidationFail,
		})
	}

	if len(r.ParameterPerturbationSweep) == 0 {
		errors = append(errors, report.ValidationError{
			Field:    "parameter_perturbation_sweep",
			Message:  "empty parameter perturbation sweep",
			Severity: report.ValidationFail,
		})
	}

	if len(r.NodeProbeOffsetSweep) == 0 {
		errors = append(errors, report.ValidationError{
			Field:    "node_probe_offset_sweep",
			Message:  "empty node probe offset sweep",
			Severity: report.ValidationFail,
		})
	}

	// 5. Reject NaN/Inf in reported numeric fields and ensure unavailable fields are paired with status/reason
	for i, res := range r.StepSizeSweep {
		if math.IsNaN(res.StepSize) || math.IsInf(res.StepSize, 0) {
			errors = append(errors, report.ValidationError{
				Field:    fmt.Sprintf("step_size_sweep[%d].step_size", i),
				Message:  "StepSizeResult contains non-finite step_size",
				Severity: report.ValidationFail,
			})
		}
		errors = append(errors, validateOptionalMetric(res.EndpointDriftAlpha, res.EndpointDriftAlphaStatus, res.EndpointDriftAlphaReason, fmt.Sprintf("step_size_sweep[%d].endpoint_drift_alpha", i))...)
		errors = append(errors, validateOptionalMetric(res.EndpointDriftPhi, res.EndpointDriftPhiStatus, res.EndpointDriftPhiReason, fmt.Sprintf("step_size_sweep[%d].endpoint_drift_phi", i))...)
		errors = append(errors, validateOptionalMetric(res.MaxAbsQ, res.MaxAbsQStatus, res.MaxAbsQReason, fmt.Sprintf("step_size_sweep[%d].max_abs_q", i))...)
		errors = append(errors, validateOptionalMetric(res.MinAmplitudeR, res.MinAmplitudeRStatus, res.MinAmplitudeRReason, fmt.Sprintf("step_size_sweep[%d].min_amplitude_r", i))...)
		errors = append(errors, validateOptionalMetric(res.MaxPhaseGrad, res.MaxPhaseGradStatus, res.MaxPhaseGradReason, fmt.Sprintf("step_size_sweep[%d].max_phase_gradient", i))...)
	}

	for i, res := range r.NodeThresholdSweep {
		if math.IsNaN(res.Threshold) || math.IsInf(res.Threshold, 0) {
			errors = append(errors, report.ValidationError{
				Field:    fmt.Sprintf("node_threshold_sweep[%d]", i),
				Message:  "ThresholdSensitivityResult contains non-finite threshold",
				Severity: report.ValidationFail,
			})
		}
	}

	for i, res := range r.PhaseGradientBoundSweep {
		if math.IsNaN(res.Bound) || math.IsInf(res.Bound, 0) {
			errors = append(errors, report.ValidationError{
				Field:    fmt.Sprintf("phase_gradient_bound_sweep[%d].bound", i),
				Message:  "PhaseGradientBoundResult contains non-finite bound",
				Severity: report.ValidationFail,
			})
		}
		errors = append(errors, validateOptionalMetric(res.MaxObservedPhaseGrad, res.MaxObservedPhaseGradStatus, res.MaxObservedPhaseGradReason, fmt.Sprintf("phase_gradient_bound_sweep[%d].max_observed_phase_gradient", i))...)
	}

	for i, res := range r.ParameterPerturbationSweep {
		if math.IsNaN(res.C2Real) || math.IsInf(res.C2Real, 0) ||
			math.IsNaN(res.K2) || math.IsInf(res.K2, 0) ||
			math.IsNaN(res.Omega2) || math.IsInf(res.Omega2, 0) {
			errors = append(errors, report.ValidationError{
				Field:    fmt.Sprintf("parameter_perturbation_sweep[%d]", i),
				Message:  "ParameterPerturbationResult contains non-finite values (NaN or Inf)",
				Severity: report.ValidationFail,
			})
		}
	}

	for i, res := range r.NodeProbeOffsetSweep {
		if math.IsNaN(res.OffsetAlpha) || math.IsInf(res.OffsetAlpha, 0) ||
			math.IsNaN(res.OffsetPhi) || math.IsInf(res.OffsetPhi, 0) {
			errors = append(errors, report.ValidationError{
				Field:    fmt.Sprintf("node_probe_offset_sweep[%d]", i),
				Message:  "NodeProbeOffsetResult contains non-finite offset values",
				Severity: report.ValidationFail,
			})
		}
		errors = append(errors, validateOptionalMetric(res.InitialAmplitude, res.InitialAmplitudeStatus, res.InitialAmplitudeReason, fmt.Sprintf("node_probe_offset_sweep[%d].initial_amplitude", i))...)
	}

	// Technical Gate audit integrity check
	if r.TechnicalGate.Name != "bmc0a_superposition_robustness_audit_gate" {
		errors = append(errors, report.ValidationError{
			Field:    "technical_gate.name",
			Message:  fmt.Sprintf("incorrect gate name: %s", r.TechnicalGate.Name),
			Severity: report.ValidationFail,
		})
	}

	// Check robustness outcome value
	validOutcome := false
	for _, o := range []string{"stable", "mixed", "fragile", "contested"} {
		if r.RobustnessOutcome == o {
			validOutcome = true
			break
		}
	}
	if !validOutcome {
		errors = append(errors, report.ValidationError{
			Field:    "robustness_outcome",
			Message:  fmt.Sprintf("invalid outcome: %s", r.RobustnessOutcome),
			Severity: report.ValidationFail,
		})
	}

	if r.PromotionGate.Status != report.StatusBlocked {
		errors = append(errors, report.ValidationError{
			Field:    "promotion_gate.status",
			Message:  "promotion gate must be blocked for Sprint 3",
			Severity: report.ValidationFail,
		})
	}

	return errors
}

// SummarizeRobustnessReport prints a human-readable summary of the numerical audit.
func SummarizeRobustnessReport(r *RobustnessReport) {
	fmt.Println("================================================================================")
	fmt.Printf("BMC Superposition Model - Numerical Robustness & Convergence Audit\n")
	fmt.Printf("Schema Version:      %s\n", r.SchemaVersion)
	fmt.Printf("Toy-Only:            %t\n", r.ToyAnalysisOnly)
	fmt.Printf("Final-Truth:         %t\n", r.FinalTruthClaim)
	fmt.Printf("Robustness Outcome:  %s\n", r.RobustnessOutcome)
	fmt.Println("================================================================================")
	fmt.Printf("Technical Gate (Audit Integrity): %s (%s)\n", r.TechnicalGate.Status, r.TechnicalGate.Name)
	if r.TechnicalGate.Reason != "" {
		fmt.Printf("  Reason: %s\n", r.TechnicalGate.Reason)
	}
	fmt.Printf("Promotion Gate:                  %s (%s)\n", r.PromotionGate.Status, r.PromotionGate.Name)
	if r.PromotionGate.Reason != "" {
		fmt.Printf("  Reason: %s\n", r.PromotionGate.Reason)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Step Size Sweep (Total T = 10.0):")
	for _, res := range r.StepSizeSweep {
		fmt.Printf("  - dt=%-6.4f steps=%-3d drift=(α:%s, φ:%s) max|Q|=%s min|Psi|=%s max|p_grad|=%s clock_mon:%-5t gate:%s\n",
			res.StepSize, res.Steps,
			formatFloatPtr(res.EndpointDriftAlpha, "%e"),
			formatFloatPtr(res.EndpointDriftPhi, "%e"),
			formatFloatPtr(res.MaxAbsQ, "%-6.3f"),
			formatFloatPtr(res.MinAmplitudeR, "%-8.6f"),
			formatFloatPtr(res.MaxPhaseGrad, "%-6.1f"),
			res.ClockMonotonic, res.TechnicalGateStatus)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Threshold Sensitivity Sweep:")
	for _, res := range r.NodeThresholdSweep {
		fmt.Printf("  - thresh=%-5e profile=%-10s node_free:%-9s q_finite:%-9s p_grad:%-9s gate:%s (%s)\n",
			res.Threshold, res.Profile, res.NodeContactFree, res.QFiniteAwayFromNodes, res.PhaseGradientFinite, res.TechnicalGateStatus, res.TechnicalGateName)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Phase Gradient Bound Sweep:")
	for _, res := range r.PhaseGradientBoundSweep {
		fmt.Printf("  - bound=%-3.0f max_observed=%s binding:%-5t status:%s\n",
			res.Bound, formatFloatPtr(res.MaxObservedPhaseGrad, "%-6.2f"), res.IsBinding, res.PhaseGradientFinite)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Parameter Perturbation Sweep:")
	for _, res := range r.ParameterPerturbationSweep {
		fmt.Printf("  - c2_real=%-4.2f k2=%-3.1f omega2=%-4.1f node_contact:%-5t q_finite:%-9s clock_mon:%-9s gate:%s\n",
			res.C2Real, res.K2, res.Omega2, res.NodeContact, res.QFiniteAwayFromNodes, res.ClockMonotonic, res.SafeGateStatus)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Node Probe Offset Sweep (NodeThresh = 1e-5):")
	for _, res := range r.NodeProbeOffsetSweep {
		fmt.Printf("  - offset=(%e, %e) init_amp=%s short_circuit:%-5t integrated:%-5t node_free:%-9s p_grad:%-9s q_finite:%-9s\n",
			res.OffsetAlpha, res.OffsetPhi, formatFloatPtr(res.InitialAmplitude, "%e"), res.ShortCircuitTriggered, res.Integrated, res.NodeContactFree, res.PhaseGradientFinite, res.QFiniteAwayFromNodes)
	}
	fmt.Println("================================================================================")
}

// GenerateAuditReport runs all sweeps and constructs the RobustnessReport.
func GenerateAuditReport(safeParams, nodeProbeParams model.SuperpositionParams) (*RobustnessReport, error) {
	stepSizeSweep, err := RunStepSizeSweep(safeParams)
	if err != nil {
		return nil, err
	}

	thresholdSweep, err := RunThresholdSweep(safeParams, nodeProbeParams)
	if err != nil {
		return nil, err
	}

	phaseBoundSweep, err := RunPhaseBoundSweep(safeParams)
	if err != nil {
		return nil, err
	}

	perturbationSweep, err := RunPerturbationSweep(safeParams)
	if err != nil {
		return nil, err
	}

	offsetSweep, err := RunNodeProbeOffsetSweep(nodeProbeParams)
	if err != nil {
		return nil, err
	}

	// Compute robustness outcome
	// stable: all safe profile audits pass without gate flips
	// mixed: some safe profile audits fail or get contested
	outcome := "stable"
	hasFails := false
	for _, r := range stepSizeSweep {
		if r.TechnicalGateStatus != model.StatusPass {
			hasFails = true
		}
	}
	for _, r := range thresholdSweep {
		if r.Profile == "safe" && r.TechnicalGateStatus != model.StatusPass {
			hasFails = true
		}
	}
	for _, r := range perturbationSweep {
		if r.SafeGateStatus != model.StatusPass {
			hasFails = true
		}
	}
	if hasFails {
		outcome = "mixed"
	}

	// Technical Gate (integrity)
	techStatus := model.StatusPass
	techReason := "All planning sweeps executed successfully and EBP boundaries are respected."

	warnings := []string{
		"Toy audit analysis only.",
		"Does not test full quantum gravity.",
		"Does not prove Bohmian mechanics, problem of time, or spacetime emergence.",
		"Passing this audit cannot promote any final-truth claim.",
	}

	r := &RobustnessReport{
		SchemaVersion:             "bmc0a-superposition-robustness-v0.1",
		ToyAnalysisOnly:           true,
		FinalTruthClaim:           false,
		SourceArtifacts:           []string{"bmc0a_superposition_safe", "bmc0a_superposition_node_probe"},
		AuditKind:                 "numerical_robustness_convergence",
		StepSizeSweep:             stepSizeSweep,
		NodeThresholdSweep:        thresholdSweep,
		PhaseGradientBoundSweep:   phaseBoundSweep,
		ParameterPerturbationSweep: perturbationSweep,
		NodeProbeOffsetSweep:      offsetSweep,
		TechnicalGate: TechnicalGate{
			Name:   "bmc0a_superposition_robustness_audit_gate",
			Status: techStatus,
			Reason: techReason,
		},
		RobustnessOutcome: outcome,
		PromotionGate: report.PromotionGate{
			Name:   "full_bmc_toy_gate",
			Status: report.StatusBlocked,
			Reason: "Friedmann residual and faithfulness review remain unpaid debt.",
		},
		EbpDebt: report.EbpDebt{
			NeedMap:                "partial",
			NeedInvariant:          "partial",
			NeedToyCheck:           "active",
			NeedNullModel:          "partial",
			NeedObstruction:        "active",
			NeedFaithfulnessReview: "contested",
		},
		Warnings: warnings,
	}

	return r, nil
}
