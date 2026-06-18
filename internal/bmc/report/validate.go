package report

import (
	"fmt"

	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/wdw"
)

type ValidationSeverity string

const (
	ValidationInfo    ValidationSeverity = "info"
	ValidationWarning ValidationSeverity = "warning"
	ValidationFail    ValidationSeverity = "fail"
)

type ValidationError struct {
	Field    string             `json:"field"`
	Message  string             `json:"message"`
	Severity ValidationSeverity `json:"severity"`
}

// Validate performs structural and semantic checks on a Report.
func Validate(r *Report) []ValidationError {
	var errors []ValidationError

	// 1. Schema version
	if r.SchemaVersion != "bmc-report-v0.1" {
		errors = append(errors, ValidationError{
			Field:    "schema_version",
			Message:  fmt.Sprintf("unsupported schema version: %s (expected 'bmc-report-v0.1')", r.SchemaVersion),
			Severity: ValidationFail,
		})
	}

	// 2. Toy analysis assertion
	if !r.ToyAnalysisOnly {
		errors = append(errors, ValidationError{
			Field:    "toy_analysis_only",
			Message:  "EBP 2.1 requires toy_analysis_only to be true for BMC-0A",
			Severity: ValidationFail,
		})
	}

	// 3. Rejects Final Truth Claims
	if r.FinalTruthClaim {
		errors = append(errors, ValidationError{
			Field:    "final_truth_claim",
			Message:  "EBP 2.1 promotion blocker: final_truth_claim must be false",
			Severity: ValidationFail,
		})
	}

	// 4. Required Checks Presence & Status/Pass Consistency
	requiredChecks := []string{"wdw_residual", "trajectory", "clock_monotonicity", "quantum_potential", "classical_limit", "friedmann_residual"}
	if r.ModelID == "bmc0a_superposition" {
		requiredChecks = append(requiredChecks, "node_detection", "node_contact_free", "q_finite_away_from_nodes", "phase_gradient_finite")
	}

	for _, checkName := range requiredChecks {
		check, exists := r.Checks[checkName]
		if !exists {
			errors = append(errors, ValidationError{
				Field:    "checks",
				Message:  fmt.Sprintf("missing required check: %s", checkName),
				Severity: ValidationFail,
			})
			continue
		}

		// Consistency check: status = "pass" <=> pass = true
		if check.Status == model.StatusPass && !check.Pass {
			errors = append(errors, ValidationError{
				Field:    fmt.Sprintf("checks.%s", checkName),
				Message:  "inconsistent check status: status is 'pass' but pass boolean is false",
				Severity: ValidationFail,
			})
		}
		if check.Status != model.StatusPass && check.Pass {
			errors = append(errors, ValidationError{
				Field:    fmt.Sprintf("checks.%s", checkName),
				Message:  fmt.Sprintf("inconsistent check status: status is '%s' but pass boolean is true", check.Status),
				Severity: ValidationFail,
			})
		}

		// Validation checks for Numerical Residual fields (EBP 2.1 failure-detection constraints)
		if check.NumericalResidualStatus != nil {
			statusVal := *check.NumericalResidualStatus
			// Reject unknown status values
			if statusVal != wdw.NumericalResidualPass &&
				statusVal != wdw.NumericalResidualViolationDetected &&
				statusVal != wdw.NumericalResidualNotComputed &&
				statusVal != wdw.NumericalResidualError {
				errors = append(errors, ValidationError{
					Field:    fmt.Sprintf("checks.%s.numerical_residual_status", checkName),
					Message:  fmt.Sprintf("unknown numerical residual status: %s", statusVal),
					Severity: ValidationFail,
				})
			}

			// Reject reports that claim pass while numerical residual violation or error is present
			if (statusVal == wdw.NumericalResidualViolationDetected || statusVal == wdw.NumericalResidualError) && check.Status == model.StatusPass {
				errors = append(errors, ValidationError{
					Field:    fmt.Sprintf("checks.%s", checkName),
					Message:  fmt.Sprintf("check cannot pass when numerical residual status is '%s'", statusVal),
					Severity: ValidationFail,
				})
			}
		}

		if check.NumericalResidualAuthority != nil {
			authVal := *check.NumericalResidualAuthority
			// Reject unknown authority values
			if authVal != wdw.NumericalAuthorityDiagnostic &&
				authVal != wdw.NumericalAuthorityOracleOnly &&
				authVal != wdw.NumericalAuthorityNone {
				errors = append(errors, ValidationError{
					Field:    fmt.Sprintf("checks.%s.numerical_residual_authority", checkName),
					Message:  fmt.Sprintf("unknown numerical residual authority: %s", authVal),
					Severity: ValidationFail,
				})
			}
		}

		// Require plane-wave passing WdW check to have correct numerical status and authority
		if r.ModelID == "bmc0a_plane" && checkName == "wdw_residual" && check.Status == model.StatusPass {
			if check.NumericalResidualStatus == nil || *check.NumericalResidualStatus != wdw.NumericalResidualPass {
				errors = append(errors, ValidationError{
					Field:    fmt.Sprintf("checks.%s.numerical_residual_status", checkName),
					Message:  fmt.Sprintf("plane-wave passing WdW check requires numerical_residual_status to be '%s'", wdw.NumericalResidualPass),
					Severity: ValidationFail,
				})
			}
			if check.NumericalResidualAuthority == nil || *check.NumericalResidualAuthority != wdw.NumericalAuthorityDiagnostic {
				errors = append(errors, ValidationError{
					Field:    fmt.Sprintf("checks.%s.numerical_residual_authority", checkName),
					Message:  fmt.Sprintf("plane-wave passing WdW check requires numerical_residual_authority to be '%s'", wdw.NumericalAuthorityDiagnostic),
					Severity: ValidationFail,
				})
			}
		}
	}


	// 5. Technical Gate Status Consistency
	if r.TechnicalGate.Name == "bmc0a_plane_control_gate" {
		allTechChecksPass := true
		techChecks := []string{"wdw_residual", "trajectory", "clock_monotonicity", "quantum_potential", "classical_limit"}
		for _, tc := range techChecks {
			check, exists := r.Checks[tc]
			if exists && check.Status != model.StatusPass {
				allTechChecksPass = false
			}
		}

		expectedTechStatus := model.StatusPass
		if !allTechChecksPass {
			expectedTechStatus = model.StatusFail
		}

		if r.TechnicalGate.Status != expectedTechStatus {
			errors = append(errors, ValidationError{
				Field:    "technical_gate.status",
				Message:  fmt.Sprintf("technical gate status '%s' is inconsistent with plane wave control checks (expected '%s')", r.TechnicalGate.Status, expectedTechStatus),
				Severity: ValidationFail,
			})
		}
	} else if r.TechnicalGate.Name == "bmc0a_superposition_safe_gate" {
		allTechChecksPass := true
		techChecks := []string{
			"wdw_residual", "trajectory", "clock_monotonicity", "quantum_potential",
			"classical_limit", "node_detection", "node_contact_free",
			"q_finite_away_from_nodes", "phase_gradient_finite",
		}
		for _, tc := range techChecks {
			check, exists := r.Checks[tc]
			if exists && check.Status != model.StatusPass {
				allTechChecksPass = false
			}
		}

		expectedTechStatus := model.StatusPass
		if !allTechChecksPass {
			expectedTechStatus = model.StatusFail
		}

		if r.TechnicalGate.Status != expectedTechStatus {
			errors = append(errors, ValidationError{
				Field:    "technical_gate.status",
				Message:  fmt.Sprintf("technical gate status '%s' is inconsistent with superposition safe checks (expected '%s')", r.TechnicalGate.Status, expectedTechStatus),
				Severity: ValidationFail,
			})
		}
	} else if r.TechnicalGate.Name == "node_detection_validation_gate" {
		// Node detection validation gate passes if node_detection check is pass and node_contact_free is fail/blocker
		nodeDetectCheck, detectExists := r.Checks["node_detection"]
		nodeContactCheck, contactExists := r.Checks["node_contact_free"]

		expectedTechStatus := model.StatusFail
		if detectExists && contactExists && nodeDetectCheck.Status == model.StatusPass && nodeContactCheck.Status == model.StatusFail {
			expectedTechStatus = model.StatusPass
		}

		if r.TechnicalGate.Status != expectedTechStatus {
			errors = append(errors, ValidationError{
				Field:    "technical_gate.status",
				Message:  fmt.Sprintf("technical gate status '%s' is inconsistent with node detection validation (expected '%s')", r.TechnicalGate.Status, expectedTechStatus),
				Severity: ValidationFail,
			})
		}
	} else {
		errors = append(errors, ValidationError{
			Field:    "technical_gate.name",
			Message:  fmt.Sprintf("unknown technical gate name: %s", r.TechnicalGate.Name),
			Severity: ValidationFail,
		})
	}

	// 6. Promotion Gate Blocking Validation
	// If Friedmann is deferred or Faithfulness is contested, the promotion gate MUST be blocked
	friedmannCheck, fExists := r.Checks["friedmann_residual"]
	if fExists && (friedmannCheck.Status == model.StatusDeferred || r.Faithfulness.Status == model.StatusContested) {
		if r.PromotionGate.Status != StatusBlocked {
			errors = append(errors, ValidationError{
				Field:    "promotion_gate.status",
				Message:  "promotion gate must be 'blocked' while Friedmann residual is deferred or faithfulness is contested",
				Severity: ValidationFail,
			})
		}
	}

	// 7. Non-empty Warnings
	if len(r.Warnings) == 0 {
		errors = append(errors, ValidationError{
			Field:    "warnings",
			Message:  "warnings list must not be empty (EBP non-claim requirements)",
			Severity: ValidationFail,
		})
	}

	return errors
}
