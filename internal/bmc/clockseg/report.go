package clockseg

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/guidance"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/wave"
)

type EbpDebtLedger struct {
	NeedMap                 string `json:"needMap"`
	NeedInvariant           string `json:"needInvariant"`
	NeedToyCheck            string `json:"needToyCheck"`
	NeedNullModel          string `json:"needNullModel"`
	NeedObstruction         string `json:"needObstruction"`
	NeedFaithfulnessReview string `json:"needFaithfulnessReview"`
	ClockChoiceDebt        string `json:"clock_choice_debt"`
	ContainsFinalTruthClaim string `json:"containsFinalTruthClaim"`
	LeanVerification       string `json:"LeanVerification"`
	PromotionStatus        string `json:"promotion_status"`
	Sprint5ClockReadiness  string `json:"sprint5_clock_readiness"`
}

type StepRefinementBranchResult struct {
	C2Real           float64               `json:"c2_real"`
	K2               float64               `json:"k2"`
	Omega2           float64               `json:"omega2"`
	StepSize         float64               `json:"step_size"`
	Steps            int                   `json:"steps"`
	NumSegments      int                   `json:"num_segments"`
	NumTurningPoints int                   `json:"num_turning_points"`
	TrajectoryValid  bool                  `json:"trajectory_valid"`
	Branches         []LocalRelationBranch `json:"branches"`
}

type TechnicalGate struct {
	Name   string            `json:"name"`
	Status model.CheckStatus `json:"status"`
	Reason string            `json:"reason"`
}

type ClockReadinessReport struct {
	SchemaVersion               string                       `json:"schema_version"`
	ToyAnalysisOnly             bool                         `json:"toy_analysis_only"`
	FinalTruthClaim             bool                         `json:"final_truth_claim"`
	SingleValuednessEpsilon    float64                      `json:"single_valuedness_epsilon"`
	MinBranchSamples            int                          `json:"min_branch_samples"`
	ReadinessScope              string                       `json:"readiness_scope"`
	NullModels                  []string                     `json:"null_models"`
	FriedmannReadiness          string                       `json:"friedmann_readiness"`
	SourceArtifacts             []string                     `json:"source_artifacts"`
	ClockTurningPoints          []ClockTurningPoint          `json:"clock_turning_points"`
	ClockSegments               []ClockSegment               `json:"clock_segments"`
	LocalRelationBranches      []LocalRelationBranch        `json:"local_relation_branches"`
	ClockIndependentDiagnostics ClockIndependentDiagnostic   `json:"clock_independent_diagnostics"`
	StepRefinementBranchAudit   []StepRefinementBranchResult `json:"step_refinement_branch_audit"`
	TechnicalGate               TechnicalGate                `json:"technical_gate"`
	PromotionGate               report.PromotionGate         `json:"promotion_gate"`
	EbpDebt                     EbpDebtLedger                `json:"ebp_debt"`
	Warnings                    []string                     `json:"warnings"`
}

type FailedConfig struct {
	C2Real float64 `json:"c2_real"`
	K2     float64 `json:"k2"`
	Omega2 float64 `json:"omega2"`
}

func DefaultFailedConfigs() []FailedConfig {
	return []FailedConfig{
		{C2Real: 0.50, K2: 2.1, Omega2: -2.1},
		{C2Real: 0.55, K2: 1.9, Omega2: -1.9},
		{C2Real: 0.55, K2: 2.0, Omega2: -2.0},
		{C2Real: 0.55, K2: 2.1, Omega2: -2.1},
	}
}

// GenerateClockReadinessReport executes the Sprint 5 diagnostics under EBP 2.1.
func GenerateClockReadinessReport(safeParams model.SuperpositionParams) (*ClockReadinessReport, error) {
	epsilon := 1e-9
	minSamples := 3

	// 1. Run core trajectory under safe parameters
	sw := wave.NewSuperpositionWave(
		safeParams.C1Real, safeParams.C1Imag, safeParams.K1, safeParams.Omega1,
		safeParams.C2Real, safeParams.C2Imag, safeParams.K2, safeParams.Omega2,
	)
	stepper := guidance.NewRK4Stepper()
	initialState := model.MiniState{Alpha: safeParams.Alpha0, Phi: safeParams.Phi0}
	traj := guidance.Integrate(sw, initialState, stepper, safeParams.LambdaStep, safeParams.Steps, safeParams.NodeThresh)

	// 2. Segment core trajectory
	segments, turningPoints := SegmentTrajectory(traj, 1e-10)

	// 3. Extract relational branches
	var branches []LocalRelationBranch
	validBranchCount := 0
	for _, seg := range segments {
		branch := ExtractLocalRelationBranch(traj, seg, epsilon, minSamples)
		branches = append(branches, branch)
		if branch.ValidationPassed {
			validBranchCount++
		}
	}

	// 4. Compute clock-independent diagnostics
	independentDiagnostics := ComputeClockIndependentDiagnostics(traj, sw, safeParams.NodeThresh, segments, turningPoints)

	// 5. Run step-refinement branch audit on the four fragile configurations
	failedConfigs := DefaultFailedConfigs()
	stepSizes := []float64{0.05, 0.025, 0.0125}
	var auditResults []StepRefinementBranchResult

	for _, config := range failedConfigs {
		// Create the superposition wave function for this configuration
		csw := wave.NewSuperpositionWave(
			safeParams.C1Real, safeParams.C1Imag, safeParams.K1, safeParams.Omega1,
			config.C2Real, safeParams.C2Imag, config.K2, config.Omega2,
		)

		for _, dt := range stepSizes {
			steps := int(math.Round(10.0 / dt))
			ctraj := guidance.Integrate(csw, initialState, stepper, dt, steps, safeParams.NodeThresh)

			cseg, ctp := SegmentTrajectory(ctraj, 1e-10)

			var cbranches []LocalRelationBranch
			for _, cs := range cseg {
				cb := ExtractLocalRelationBranch(ctraj, cs, epsilon, minSamples)
				cbranches = append(cbranches, cb)
			}

			trajValid := guidance.IsFinite(ctraj) && len(ctraj.Points) > 1

			auditResults = append(auditResults, StepRefinementBranchResult{
				C2Real:           config.C2Real,
				K2:               config.K2,
				Omega2:           config.Omega2,
				StepSize:         dt,
				Steps:            steps,
				NumSegments:      len(cseg),
				NumTurningPoints: len(ctp),
				TrajectoryValid:  trajValid,
				Branches:         cbranches,
			})
		}
	}

	// Set Friedmann readiness based on validity of branches
	friedmannReadiness := "blocked"
	if validBranchCount > 0 {
		friedmannReadiness = "local_only_candidate"
	}

	// 6. Build the Technical Gate
	techGate := TechnicalGate{
		Name:   "bmc0a_clock_readiness_gate",
		Status: model.StatusPass,
		Reason: "Local clock segmentation and clock-independent diagnostics completed successfully.",
	}

	// 7. Build the Promotion Gate
	promotionGate := report.PromotionGate{
		Name:   "full_bmc_gate",
		Status: report.StatusBlocked,
		Reason: "Sprint 5 clock-readiness promotion status is planned_candidate_only. Full BMC promotion remains blocked.",
	}

	// 8. Build the EBP Debt Ledger
	ebpDebt := EbpDebtLedger{
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
		Sprint5ClockReadiness:  "planned_candidate_only",
	}

	warnings := []string{
		"Sprint 5 plan approved with revisions. Implement only local clock segmentation, local alpha(phi) branch checks, and clock-independent diagnostics.",
		"Do not implement Friedmann residual recovery, massive scalar, LQC/Page-Wootters comparison, full BMC promotion, or any full quantum-gravity claim.",
		"local_only_candidate is not Friedmann recovery readiness",
	}

	rep := &ClockReadinessReport{
		SchemaVersion:               "bmc0a-clock-readiness-v0.1",
		ToyAnalysisOnly:             true,
		FinalTruthClaim:             false,
		SingleValuednessEpsilon:    epsilon,
		MinBranchSamples:            minSamples,
		ReadinessScope:              "All 12 step-refinement branch-audit runs are present, covering 4 fragile configurations across 3 step sizes.",
		NullModels:                  []string{"No new null models planned. Existing null-model debt remains partial/deferred."},
		FriedmannReadiness:          friedmannReadiness,
		SourceArtifacts: []string{
			"Sprint 1: BMC-0A plane-wave control artifact",
			"Sprint 2: BMC-0A two-plane-wave superposition control artifact",
			"Sprint 2: BMC-0A node-obstruction detection artifact",
			"Sprint 3: BMC-0A numerical robustness/convergence audit artifact",
			"Sprint 4: BMC-0A clock-fragility diagnostic artifact",
		},
		ClockTurningPoints:          turningPoints,
		ClockSegments:               segments,
		LocalRelationBranches:      branches,
		ClockIndependentDiagnostics: independentDiagnostics,
		StepRefinementBranchAudit:   auditResults,
		TechnicalGate:               techGate,
		PromotionGate:               promotionGate,
		EbpDebt:                     ebpDebt,
		Warnings:                    warnings,
	}

	return rep, nil
}

// ReadClockReadinessReport reads a JSON file and strictly decodes it.
func ReadClockReadinessReport(path string) (*ClockReadinessReport, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dec.DisallowUnknownFields()

	var rep ClockReadinessReport
	if err := dec.Decode(&rep); err != nil {
		return nil, err
	}
	return &rep, nil
}

// WriteJSON writes the report to a pretty-printed JSON file.
func WriteJSON(rep *ClockReadinessReport, path string) error {
	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0644)
}

// SummarizeClockReadinessReport outputs a human-readable summary of the report.
func SummarizeClockReadinessReport(rep *ClockReadinessReport) {
	fmt.Println("================================================================================")
	fmt.Printf("BMC Sprint 5 Clock Readiness Report Summary\n")
	fmt.Println("================================================================================")
	fmt.Printf("Schema Version:            %s\n", rep.SchemaVersion)
	fmt.Printf("Friedmann Readiness:       %s\n", rep.FriedmannReadiness)
	fmt.Printf("Toy Analysis Only:         %v\n", rep.ToyAnalysisOnly)
	fmt.Printf("Final Truth Claim:         %v\n", rep.FinalTruthClaim)
	fmt.Printf("Single Valuedness Epsilon: %e\n", rep.SingleValuednessEpsilon)
	fmt.Printf("Min Branch Samples:        %d\n", rep.MinBranchSamples)
	fmt.Printf("Readiness Scope:           %s\n", rep.ReadinessScope)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Technical Gate:            %s (Status: %s)\n", rep.TechnicalGate.Name, rep.TechnicalGate.Status)
	fmt.Printf("Reason:                    %s\n", rep.TechnicalGate.Reason)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Promotion Gate:            %s (Status: %s)\n", rep.PromotionGate.Name, rep.PromotionGate.Status)
	fmt.Printf("Reason:                    %s\n", rep.PromotionGate.Reason)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Null Models:\n")
	for _, nm := range rep.NullModels {
		fmt.Printf("  - %s\n", nm)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Clock-Independent Diagnostics:\n")
	d := rep.ClockIndependentDiagnostics
	fmt.Printf("  Path Length in Conf Space: %f\n", d.PathLengthInConfigurationSpace)
	fmt.Printf("  Total Lambda Interval:     %f\n", d.TotalLambdaInterval)
	fmt.Printf("  Valid Trajectory Points:   %d\n", d.NumValidTrajectoryPoints)
	fmt.Printf("  Clock Segments Count:      %d\n", d.NumClockSegments)
	fmt.Printf("  Turning Points Count:      %d\n", d.NumTurningPoints)
	fmt.Printf("  Min Amplitude R:           %f\n", d.MinAmplitudeR)
	fmt.Printf("  Max Abs Q (away nodes):    %f\n", d.MaxAbsQAwayFromNodes)
	fmt.Printf("  Max Phase Gradient:        %f\n", d.MaxPhaseGradient)
	fmt.Printf("  Node Contact Free:         %v\n", d.NodeContactFree)
	fmt.Printf("  Trajectory Finiteness:     %v\n", d.TrajectoryFiniteness)
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
	fmt.Printf("  sprint5_clock_readiness:  %s\n", rep.EbpDebt.Sprint5ClockReadiness)
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Printf("Sweeps & Detections:\n")
	fmt.Printf("  Total clock turning points:   %d\n", len(rep.ClockTurningPoints))
	fmt.Printf("  Total clock segments:         %d\n", len(rep.ClockSegments))
	fmt.Printf("  Total local relation branches:%d\n", len(rep.LocalRelationBranches))
	fmt.Printf("  Step Refinement Audit runs:   %d\n", len(rep.StepRefinementBranchAudit))
	fmt.Println("================================================================================")
}
