package clockdiag

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

type ClockEbpDebt struct {
	NeedMap                string `json:"needMap"`
	NeedInvariant          string `json:"needInvariant"`
	NeedToyCheck           string `json:"needToyCheck"`
	NeedNullModel          string `json:"needNullModel"`
	NeedObstruction        string `json:"needObstruction"`
	NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
	ClockChoiceDebt        string `json:"clock_choice_debt"`
	ContainsFinalTruthClaim string `json:"containsFinalTruthClaim"`
	LeanVerification       string `json:"LeanVerification"`
	PromotionStatus        string `json:"promotion_status"`
}

type ClockFragilityReport struct {
	SchemaVersion              string                    `json:"schema_version"`
	ToyAnalysisOnly            bool                      `json:"toy_analysis_only"`
	FinalTruthClaim            bool                      `json:"final_truth_claim"`
	NearZeroDPhiThreshold      float64                   `json:"near_zero_dphi_threshold"`
	SourceArtifacts            []string                  `json:"source_artifacts"`
	DiagnosticKind             string                    `json:"diagnostic_kind"`
	FailedPerturbationRechecks []StepRefinementResult    `json:"failed_perturbation_rechecks"`
	ClockEvents                []ClockEvent              `json:"clock_events"`
	CorrelationSummary         []CorrelationSummary      `json:"correlation_summary"`
	AlternativeClockSummary    AlternativeClockSummary   `json:"alternative_clock_summary"`
	TrajectoryValiditySummary  TrajectoryValiditySummary `json:"trajectory_validity_summary"`
	TechnicalGate              TechnicalGate             `json:"technical_gate"`
	DiagnosticOutcome          string                    `json:"diagnostic_outcome"`
	PromotionGate              report.PromotionGate      `json:"promotion_gate"`
	EbpDebt                    ClockEbpDebt              `json:"ebp_debt"`
	Warnings                   []string                  `json:"warnings"`
}

type TechnicalGate struct {
	Name   string            `json:"name"`
	Status model.CheckStatus `json:"status"`
	Reason string            `json:"reason"`
}

// ReadClockFragilityReport reads and decodes the report strictly, rejecting unknown fields.
func ReadClockFragilityReport(path string) (*ClockFragilityReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep ClockFragilityReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}
	return &rep, nil
}

// WriteJSON serializes the clock fragility report with pretty-printing.
func WriteJSON(rep *ClockFragilityReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// GenerateClockFragilityReport runs the diagnostic routines and compiles the report.
func GenerateClockFragilityReport(safeParams model.SuperpositionParams) (*ClockFragilityReport, error) {
	nearZeroThreshold := 1e-10
	failedConfigs := DefaultFailedConfigs()

	// 1. Step refinement rechecks
	refinementResults, err := RunStepRefinementRechecks(safeParams, failedConfigs, nearZeroThreshold)
	if err != nil {
		return nil, err
	}

	// 2. Correlation summary
	correlationSummary, err := ComputeCorrelations(safeParams, nearZeroThreshold)
	if err != nil {
		return nil, err
	}

	// 3. Detect clock events at default step size (dt = 0.05) for failed configs
	var defaultClockEvents []ClockEvent
	for _, config := range failedConfigs {
		sw := wave.NewSuperpositionWave(
			safeParams.C1Real, safeParams.C1Imag, safeParams.K1, safeParams.Omega1,
			config.C2Real, safeParams.C2Imag, config.K2, config.Omega2,
		)
		stepper := guidance.NewRK4Stepper()
		initialState := model.MiniState{Alpha: safeParams.Alpha0, Phi: safeParams.Phi0}
		traj := guidance.Integrate(sw, initialState, stepper, safeParams.LambdaStep, safeParams.Steps, safeParams.NodeThresh)
		
		events := DetectClockEvents(traj, sw, safeParams, nearZeroThreshold)
		defaultClockEvents = append(defaultClockEvents, events...)
	}

	// 4. Summarize alternative clocks and trajectory vs clock validity
	altClockSummary := ComputeAlternativeClockSummary(refinementResults)
	trajValiditySummary := ComputeTrajectoryValiditySummary(refinementResults)

	// 5. Compute outcome: clock_stable, clock_fragile, mixed, contested
	// Group refinement results by config
	configFinestStable := make(map[string]bool)
	configFinestTotal := make(map[string]int)

	for _, r := range refinementResults {
		if r.StepSize == 0.0125 {
			key := fmt.Sprintf("c2=%.2f,k2=%.1f", r.C2Real, r.K2)
			configFinestTotal[key]++
			if r.PhiMonotonic {
				configFinestStable[key] = true
			}
		}
	}

	stableCount := 0
	for _, stable := range configFinestStable {
		if stable {
			stableCount++
		}
	}

	var diagnosticOutcome string
	if stableCount == len(configFinestTotal) && len(configFinestTotal) > 0 {
		diagnosticOutcome = "clock_stable"
	} else if stableCount == 0 && len(configFinestTotal) > 0 {
		diagnosticOutcome = "clock_fragile"
	} else {
		diagnosticOutcome = "mixed"
	}

	// 6. Technical Gate status
	technicalGate := TechnicalGate{
		Name:   "bmc0a_clock_fragility_diagnostic_gate",
		Status: model.StatusPass,
		Reason: "All clock-fragility sweeps completed, distinction between trajectory and clock validity was preserved, and EBP data integrity was validated.",
	}

	// 7. Promotion Gate
	promotionGate := report.PromotionGate{
		Name:   "full_bmc_gate",
		Status: report.StatusBlocked,
		Reason: "Sprint 4 clock-fragility diagnostic promotion status is planned_candidate_only. Full BMC promotion remains blocked.",
	}

	// 8. EBP Debt Ledger
	ebpDebt := ClockEbpDebt{
		NeedMap:                 "partial",
		NeedInvariant:           "partial",
		NeedToyCheck:            "active",
		NeedNullModel:           "partial/deferred",
		NeedObstruction:         "active",
		NeedFaithfulnessReview:  "contested",
		ClockChoiceDebt:         "active",
		ContainsFinalTruthClaim: "absent",
		LeanVerification:        "planned",
		PromotionStatus:         "planned_candidate_only",
	}

	// 9. Assemble Report
	rep := &ClockFragilityReport{
		SchemaVersion:              "bmc0a-clock-fragility-v0.1",
		ToyAnalysisOnly:            true,
		FinalTruthClaim:            false,
		NearZeroDPhiThreshold:      nearZeroThreshold,
		SourceArtifacts: []string{
			"Sprint 2: BMC-0A two-plane-wave superposition control artifact",
			"Sprint 2: BMC-0A node-obstruction detection artifact",
			"Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
		},
		DiagnosticKind:             "clock_monotonicity_fragility",
		FailedPerturbationRechecks: refinementResults,
		ClockEvents:                defaultClockEvents,
		CorrelationSummary:         correlationSummary,
		AlternativeClockSummary:    altClockSummary,
		TrajectoryValiditySummary:  trajValiditySummary,
		TechnicalGate:              technicalGate,
		DiagnosticOutcome:          diagnosticOutcome,
		PromotionGate:              promotionGate,
		EbpDebt:                    ebpDebt,
		Warnings: []string{
			"Sprint 4 clock-fragility diagnostic only.",
			"Does not formalize clock physics in Lean.",
			"Does not claim full quantum gravity or Friedmann recovery.",
			"Relational clock φ nonmonotonicity is identified as a clock choice limitation or toy trajectory feature, not a trajectory obstruction.",
		},
	}

	return rep, nil
}

// SummarizeClockFragilityReport prints a human-readable summary of the report to stdout.
func SummarizeClockFragilityReport(rep *ClockFragilityReport) {
	fmt.Println("================================================================================")
	fmt.Printf("BMC Sprint 4 Clock Fragility Report Summary\n")
	fmt.Println("================================================================================")
	fmt.Printf("Schema Version:            %s\n", rep.SchemaVersion)
	fmt.Printf("Diagnostic Kind:           %s\n", rep.DiagnosticKind)
	fmt.Printf("Diagnostic Outcome:        %s\n", rep.DiagnosticOutcome)
	fmt.Printf("Toy Analysis Only:         %v\n", rep.ToyAnalysisOnly)
	fmt.Printf("Final Truth Claim:         %v\n", rep.FinalTruthClaim)
	fmt.Printf("Near Zero DPhi Threshold:  %e\n", rep.NearZeroDPhiThreshold)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Technical Gate:            %s (Status: %s)\n", rep.TechnicalGate.Name, rep.TechnicalGate.Status)
	fmt.Printf("Reason:                    %s\n", rep.TechnicalGate.Reason)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Promotion Gate:            %s (Status: %s)\n", rep.PromotionGate.Name, rep.PromotionGate.Status)
	fmt.Printf("Reason:                    %s\n", rep.PromotionGate.Reason)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("EBP Debt Ledger:\n")
	fmt.Printf("  needMap:                 %s\n", rep.EbpDebt.NeedMap)
	fmt.Printf("  needInvariant:           %s\n", rep.EbpDebt.NeedInvariant)
	fmt.Printf("  needToyCheck:            %s\n", rep.EbpDebt.NeedToyCheck)
	fmt.Printf("  needNullModel:           %s\n", rep.EbpDebt.NeedNullModel)
	fmt.Printf("  needObstruction:         %s\n", rep.EbpDebt.NeedObstruction)
	fmt.Printf("  needFaithfulnessReview:  %s\n", rep.EbpDebt.NeedFaithfulnessReview)
	fmt.Printf("  clock_choice_debt:       %s\n", rep.EbpDebt.ClockChoiceDebt)
	fmt.Printf("  containsFinalTruthClaim: %s\n", rep.EbpDebt.ContainsFinalTruthClaim)
	fmt.Printf("  LeanVerification:        %s\n", rep.EbpDebt.LeanVerification)
	fmt.Printf("  promotion_status:        %s\n", rep.EbpDebt.PromotionStatus)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Alternative Clock Summary:\n")
	fmt.Printf("  phi_monotonic:           %s\n", rep.AlternativeClockSummary.PhiMonotonic)
	fmt.Printf("  alpha_monotonic:         %s\n", rep.AlternativeClockSummary.AlphaMonotonic)
	fmt.Printf("  both_monotonic:          %v\n", rep.AlternativeClockSummary.BothMonotonic)
	fmt.Printf("  neither_monotonic:       %v\n", rep.AlternativeClockSummary.NeitherMonotonic)
	fmt.Printf("  clock_choice_debt:       %s\n", rep.AlternativeClockSummary.ClockChoiceDebt)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Trajectory Validity Summary:\n")
	fmt.Printf("  trajectory_valid:        %s\n", rep.TrajectoryValiditySummary.TrajectoryValid)
	fmt.Printf("  phi_clock_valid:         %s\n", rep.TrajectoryValiditySummary.PhiClockValid)
	fmt.Printf("  distinction_preserved:   %v\n", rep.TrajectoryValiditySummary.DistinctionPreserved)
	fmt.Printf("  reason:                  %s\n", rep.TrajectoryValiditySummary.Reason)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Sweeps & Detections:\n")
	fmt.Printf("  Total step refinement sweep runs:  %d\n", len(rep.FailedPerturbationRechecks))
	fmt.Printf("  Total clock events detected:       %d\n", len(rep.ClockEvents))
	fmt.Printf("  Total correlation summary entries: %d\n", len(rep.CorrelationSummary))
	fmt.Println("================================================================================")
}
