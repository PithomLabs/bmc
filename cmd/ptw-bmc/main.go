package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/PithomLabs/bmc/internal/bmc/audit"
	"github.com/PithomLabs/bmc/internal/bmc/clockdiag"
	"github.com/PithomLabs/bmc/internal/bmc/clockseg"
	"github.com/PithomLabs/bmc/internal/bmc/friedmannspec"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/nullspec"
	"github.com/PithomLabs/bmc/internal/bmc/priorart"
	"github.com/PithomLabs/bmc/internal/bmc/report"
	"github.com/PithomLabs/bmc/internal/bmc/residualrun"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	subcommand := os.Args[1]
	switch subcommand {
	case "run":
		runCmd()
	case "validate":
		validateCmd()
	case "summarize":
		summarizeCmd()
	case "audit":
		auditCmd()
	case "diagnose-clock":
		diagnoseClockCmd()
	case "segment-clock":
		segmentClockCmd()
	case "spec-friedmann":
		specFriedmannCmd()
	case "spec-nullmodels":
		specNullModelsCmd()
	case "prior-art-boundary":
		priorArtBoundaryCmd()
	case "run-nullmodels":
		runNullModelsCmd()
	case "run-residuals":
		runResidualsCmd()
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown subcommand '%s'\n", subcommand)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: ptw-bmc <subcommand> [flags]")
	fmt.Println("Subcommands:")
	fmt.Println("  run        Run the simulation and generate a JSON report")
	fmt.Println("  validate   Validate a generated JSON report schema and constraints")
	fmt.Println("  summarize  Print a human-readable summary of the report")
	fmt.Println("  audit      Run a numerical robustness and convergence audit")
	fmt.Println("  diagnose-clock Run a clock-monotonicity fragility investigation and diagnostic")
	fmt.Println("  segment-clock Run local clock segmentation and clock-independent readiness diagnostics")
	fmt.Println("  spec-friedmann Run Friedmann-residual specification and gate design")
	fmt.Println("  spec-nullmodels Run Null-Model specification and future comparison gate design")
	fmt.Println("  prior-art-boundary Run prior-art boundary and gate verification")
	fmt.Println("  run-nullmodels Run the null-model runner and generate comparison report")
	fmt.Println("  run-residuals Run the candidate local-branch residual runner")
}

func runCmd() {
	runFlags := flag.NewFlagSet("run", flag.ExitOnError)
	profileOpt := runFlags.String("profile", "bmc0a-plane", "Model profile to run (bmc0a-plane, bmc0a-superposition-safe, bmc0a-superposition-node-probe)")
	outOpt := runFlags.String("out", "", "Output path for the generated JSON report (required)")
	finalTruthOpt := runFlags.Bool("final-truth-claim", false, "Assert final truth claim (will fail validation)")

	if len(os.Args) < 3 {
		runFlags.Usage()
		os.Exit(1)
	}

	if err := runFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		runFlags.Usage()
		os.Exit(1)
	}

	var rep *report.Report
	var err error

	switch *profileOpt {
	case "bmc0a-plane":
		params := model.DefaultPlaneWaveParams()
		rep, err = report.Generate(params, *finalTruthOpt)
	case "bmc0a-superposition-safe":
		params := model.DefaultSuperpositionSafeParams()
		rep, err = report.GenerateSuperposition(params, *finalTruthOpt)
	case "bmc0a-superposition-node-probe":
		params := model.DefaultSuperpositionNodeProbeParams()
		rep, err = report.GenerateSuperposition(params, *finalTruthOpt)
	default:
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-plane', 'bmc0a-superposition-safe', 'bmc0a-superposition-node-probe')\n", *profileOpt)
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating report: %v\n", err)
		os.Exit(1)
	}

	if err := report.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func validateCmd() {
	validateFlags := flag.NewFlagSet("validate", flag.ExitOnError)
	reportOpt := validateFlags.String("report", "", "Path to the JSON report file (required)")

	if len(os.Args) < 3 {
		validateFlags.Usage()
		os.Exit(1)
	}

	if err := validateFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *reportOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --report is required")
		validateFlags.Usage()
		os.Exit(1)
	}

	data, err := os.ReadFile(*reportOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var helper struct {
		SchemaVersion string `json:"schema_version"`
	}
	_ = json.Unmarshal(data, &helper)

	if helper.SchemaVersion == "bmc0a-superposition-robustness-v0.1" {
		rep, err := audit.ReadRobustnessReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading robustness report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		errors := audit.ValidateRobustnessReport(rep)
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Robustness Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Robustness Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-clock-fragility-v0.1" {
		rep, err := clockdiag.ReadClockFragilityReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading clock fragility report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		errors := clockdiag.ValidateClockFragilityReport(rep)
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Clock Fragility Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Clock Fragility Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-clock-readiness-v0.1" {
		rep, err := clockseg.ReadClockReadinessReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading clock readiness report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		errors := clockseg.ValidateClockReadinessReport(rep)
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Clock Readiness Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Clock Readiness Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-friedmann-spec-v0.1" {
		rep, err := friedmannspec.ReadFriedmannSpecReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading Friedmann spec report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		errors := friedmannspec.ValidateFriedmannSpecReport(rep)
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Friedmann Spec Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Friedmann Spec Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-nullmodel-spec-v0.1" {
		rep, err := nullspec.ReadNullModelSpecReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading null model spec report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		errors := nullspec.ValidateNullModelSpecReport(rep)
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Null Model Spec Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Null Model Spec Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-prior-art-boundary-v0.1" {
		rep, err := priorart.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading prior art boundary report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		data, err := os.ReadFile(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading prior art boundary report file: %v\n", err)
			os.Exit(1)
		}
		errors := priorart.ValidateReport(rep, string(data))
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Prior-Art Boundary Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Prior-Art Boundary Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-nullrun-v0.1" {
		rep, err := nullrun.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading null model run report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		data, err := os.ReadFile(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading null model run report file: %v\n", err)
			os.Exit(1)
		}
		errors := nullrun.ValidateReport(rep, string(data))
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Null Model Run Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Null Model Run Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	if helper.SchemaVersion == "bmc0a-local-residual-v0.1" {
		rep, err := residualrun.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading residual run report (strict decoding): %v\n", err)
			os.Exit(1)
		}
		data, err := os.ReadFile(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading residual run report file: %v\n", err)
			os.Exit(1)
		}
		errors := residualrun.ValidateReport(rep, string(data))
		if len(errors) > 0 {
			fmt.Fprintln(os.Stderr, "Candidate Local-Branch Residual Report Validation FAILED:")
			for _, valErr := range errors {
				fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
			}
			os.Exit(1)
		}
		fmt.Println("Candidate Local-Branch Residual Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
		return
	}

	rep, err := readReport(*reportOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading report: %v\n", err)
		os.Exit(1)
	}

	errors := report.Validate(rep)
	if len(errors) > 0 {
		fmt.Fprintln(os.Stderr, "Report Validation FAILED:")
		for _, valErr := range errors {
			fmt.Fprintf(os.Stderr, "  - [%s] Field '%s': %s\n", valErr.Severity, valErr.Field, valErr.Message)
		}
		os.Exit(1)
	}

	fmt.Println("Report Validation PASSED: Schema and EBP 2.1 promotion constraints satisfied.")
}

func summarizeCmd() {
	summarizeFlags := flag.NewFlagSet("summarize", flag.ExitOnError)
	reportOpt := summarizeFlags.String("report", "", "Path to the JSON report file (required)")

	if len(os.Args) < 3 {
		summarizeFlags.Usage()
		os.Exit(1)
	}

	if err := summarizeFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *reportOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --report is required")
		summarizeFlags.Usage()
		os.Exit(1)
	}

	data, err := os.ReadFile(*reportOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	var helper struct {
		SchemaVersion string `json:"schema_version"`
	}
	_ = json.Unmarshal(data, &helper)

	if helper.SchemaVersion == "bmc0a-superposition-robustness-v0.1" {
		rep, err := audit.ReadRobustnessReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading robustness report: %v\n", err)
			os.Exit(1)
		}
		audit.SummarizeRobustnessReport(rep)
		return
	}

	if helper.SchemaVersion == "bmc0a-clock-fragility-v0.1" {
		rep, err := clockdiag.ReadClockFragilityReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading clock fragility report: %v\n", err)
			os.Exit(1)
		}
		clockdiag.SummarizeClockFragilityReport(rep)
		return
	}

	if helper.SchemaVersion == "bmc0a-clock-readiness-v0.1" {
		rep, err := clockseg.ReadClockReadinessReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading clock readiness report: %v\n", err)
			os.Exit(1)
		}
		clockseg.SummarizeClockReadinessReport(rep)
		return
	}

	if helper.SchemaVersion == "bmc0a-friedmann-spec-v0.1" {
		rep, err := friedmannspec.ReadFriedmannSpecReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading Friedmann spec report: %v\n", err)
			os.Exit(1)
		}
		friedmannspec.SummarizeFriedmannSpecReport(rep)
		return
	}

	if helper.SchemaVersion == "bmc0a-nullmodel-spec-v0.1" {
		rep, err := nullspec.ReadNullModelSpecReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading null model spec report: %v\n", err)
			os.Exit(1)
		}
		nullspec.SummarizeNullModelSpecReport(rep)
		return
	}

	if helper.SchemaVersion == "bmc0a-prior-art-boundary-v0.1" {
		rep, err := priorart.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading prior art boundary report: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("BMC Sprint 8-Lite Prior-Art Boundary Summary")
		fmt.Printf("Schema Version: %s\n", rep.SchemaVersion)
		fmt.Printf("Scope: %s\n", rep.Scope)
		fmt.Printf("Scientific Novelty Claim Made: %v\n", rep.ScientificNoveltyClaimMade)
		fmt.Printf("Scientific Novelty Claim Allowed: %v\n", rep.ScientificNoveltyClaimAllowed)
		fmt.Printf("Workflow Distinctiveness: %s\n", rep.WorkflowDistinctivenessStatus)
		fmt.Printf("Residual Computed: %v\n", rep.ResidualComputed)
		fmt.Printf("Null Comparison Computed: %v\n", rep.NullComparisonComputed)
		fmt.Printf("Recovery Claim: %v\n", rep.RecoveryClaim)
		fmt.Printf("Prior-Art Sources Seeded: %d\n", len(rep.PriorArtSources))
		fmt.Printf("Boundary Claims Declared: %d\n", len(rep.BoundaryClaims))
		fmt.Printf("Full BMC: %s\n", rep.FullBmcToyGate)
		fmt.Printf("Promotion Status: %s\n", rep.EbpDebt.PromotionStatus)
		return
	}

	if helper.SchemaVersion == "bmc0a-nullrun-v0.1" {
		rep, err := nullrun.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading null model run report: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("BMC Sprint 9 Null-Model Runner Summary")
		fmt.Printf("Schema Version: %s\n", rep.SchemaVersion)
		fmt.Printf("Scope: %s\n", rep.Scope)
		fmt.Printf("Residual Computed: %v\n", rep.ResidualComputed)
		fmt.Printf("Null Diagnostics Computed: %v\n", rep.NullDiagnosticsComputed)
		fmt.Printf("Target/Null Comparison Computed: %v\n", rep.TargetNullComparisonComputed)
		fmt.Printf("Recovery Claim: %v\n", rep.RecoveryClaim)
		fmt.Printf("Scientific Novelty Claim Made: %v\n", rep.ScientificNoveltyClaimMade)
		fmt.Printf("Full BMC: %s\n", rep.FullBmcToyGate)
		numWithDiag := 0
		numBlocked := 0
		numDeferred := 0
		for _, run := range rep.NullModelRuns {
			if run.RunStatus == "diagnostics_generated" {
				numWithDiag++
			} else if run.RunStatus == "blocked" {
				numBlocked++
			} else if run.RunStatus == "deferred" {
				numDeferred++
			}
		}
		totalAccounted := numWithDiag + numBlocked + numDeferred
		fmt.Printf("Null Models Registered: %d\n", len(rep.NullModelRuns))
		fmt.Printf("Null Models With Diagnostics: %d\n", numWithDiag)
		fmt.Printf("Null Models Blocked: %d\n", numBlocked)
		fmt.Printf("Null Models Deferred: %d\n", numDeferred)
		fmt.Printf("Null Models Accounted For: %d/%d\n", totalAccounted, len(rep.NullModelRuns))
		fmt.Printf("Interpretation Status: %s\n", rep.InterpretationStatus)
		fmt.Printf("Promotion Status: %s\n", rep.EbpDebt.PromotionStatus)
		return
	}

	if helper.SchemaVersion == "bmc0a-local-residual-v0.1" {
		rep, err := residualrun.ReadReport(*reportOpt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading residual run report: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("BMC Sprint 10 Candidate Local-Branch Residual Summary")
		fmt.Printf("Schema Version: %s\n", rep.SchemaVersion)
		fmt.Printf("Scope: %s\n", rep.Scope)
		fmt.Printf("Candidate Residual Computed: %v\n", rep.CandidateResidualComputed)
		fmt.Printf("Recovery Claim: %v\n", rep.ResidualRecoveryClaim)
		fmt.Printf("Scientific Novelty Claim Made: %v\n", rep.ScientificNoveltyClaimMade)
		fmt.Printf("BMC Beats Null Models Claim: %v\n", rep.BmcBeatsNullModelsClaim)
		fmt.Printf("Full BMC: %s\n", rep.FullBmcToyGate)
		eligibleCount := 0
		for _, b := range rep.LocalBranchEligibility {
			if b.Eligible {
				eligibleCount++
			}
		}
		fmt.Printf("Eligible Local Branches: %d\n", eligibleCount)
		numComputed := 0
		numBlocked := 0
		for _, rd := range rep.CandidateResidualDiagnostics {
			if rd.ResidualComputed {
				numComputed++
			} else {
				numBlocked++
			}
		}
		fmt.Printf("Computed Candidate Residual Diagnostics: %d\n", numComputed)
		fmt.Printf("Blocked Candidate Residual Diagnostics: %d\n", numBlocked)
		fmt.Printf("Total Candidate Residual Diagnostics: %d\n", len(rep.CandidateResidualDiagnostics))
		fmt.Printf("Residual/Null Comparisons: %d\n", len(rep.ResidualNullComparisons))
		fmt.Printf("Interpretation Status: %s\n", rep.InterpretationStatus)
		fmt.Printf("Promotion Status: %s\n", rep.EbpDebt.PromotionStatus)
		return
	}

	rep, err := readReport(*reportOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading report: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("================================================================================")
	fmt.Printf("Bohmian Minisuperspace Cosmology (BMC-0.1) - Report Summary\n")
	fmt.Printf("Model ID:       %s\n", rep.ModelID)
	fmt.Printf("Schema Version: %s\n", rep.SchemaVersion)
	fmt.Printf("Toy-Only:       %t\n", rep.ToyAnalysisOnly)
	fmt.Printf("Final-Truth:    %t (Must be false)\n", rep.FinalTruthClaim)
	fmt.Println("================================================================================")
	fmt.Printf("Technical Gate:            %s (%s)\n", rep.TechnicalGate.Status, rep.TechnicalGate.Name)
	fmt.Printf("Promotion Gate (Full):     %s\n", rep.PromotionGate.Status)
	if rep.PromotionGate.Status == report.StatusBlocked {
		fmt.Printf("  Reason: %s\n", rep.PromotionGate.Reason)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	if rep.Parameters.PlaneWave != nil {
		fmt.Println("Parameters (Plane Wave):")
		p := rep.Parameters.PlaneWave
		fmt.Printf("  k:           %f\n", p.K)
		fmt.Printf("  omega:       %f\n", p.Omega)
		fmt.Printf("  alpha0:      %f\n", p.Alpha0)
		fmt.Printf("  phi0:        %f\n", p.Phi0)
	} else if rep.Parameters.Superposition != nil {
		fmt.Println("Parameters (Superposition):")
		p := rep.Parameters.Superposition
		fmt.Printf("  c1:          (%f, %f)\n", p.C1Real, p.C1Imag)
		fmt.Printf("  k1:          %f\n", p.K1)
		fmt.Printf("  omega1:      %f\n", p.Omega1)
		fmt.Printf("  c2:          (%f, %f)\n", p.C2Real, p.C2Imag)
		fmt.Printf("  k2:          %f\n", p.K2)
		fmt.Printf("  omega2:      %f\n", p.Omega2)
		fmt.Printf("  alpha0:      %f\n", p.Alpha0)
		fmt.Printf("  phi0:        %f\n", p.Phi0)
		fmt.Printf("  node_thresh: %e\n", p.NodeThresh)
		fmt.Printf("  max_p_grad:  %f\n", p.MaxPhaseGrad)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Checks:")
	for name, check := range rep.Checks {
		passStr := "FAIL"
		if check.Pass {
			passStr = "PASS"
		}
		fmt.Printf("  - %-25s [%-9s] Status: %s - %s\n", name, passStr, check.Status, check.Reason)
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("Active Obstructions:")
	hasActiveObstructions := false
	for _, obs := range rep.Obstructions {
		if obs.Applies {
			hasActiveObstructions = true
			fmt.Printf("  - %s (%s): %s\n", obs.Name, obs.Severity, obs.Evidence)
		}
	}
	if !hasActiveObstructions {
		fmt.Println("  No active obstructions detected.")
	}
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println("EBP Warnings (Mandatory Non-Claims):")
	for _, warning := range rep.Warnings {
		fmt.Printf("  * %s\n", warning)
	}
	fmt.Println("================================================================================")
}

func auditCmd() {
	auditFlags := flag.NewFlagSet("audit", flag.ExitOnError)
	profileOpt := auditFlags.String("profile", "bmc0a-superposition-robustness", "Audit profile to run (bmc0a-superposition-robustness)")
	outOpt := auditFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		auditFlags.Usage()
		os.Exit(1)
	}

	if err := auditFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		auditFlags.Usage()
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-superposition-robustness" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-superposition-robustness')\n", *profileOpt)
		os.Exit(1)
	}

	safeParams := model.DefaultSuperpositionSafeParams()
	nodeProbeParams := model.DefaultSuperpositionNodeProbeParams()

	rep, err := audit.GenerateAuditReport(safeParams, nodeProbeParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating audit report: %v\n", err)
		os.Exit(1)
	}

	if err := audit.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran audit profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func diagnoseClockCmd() {
	diagFlags := flag.NewFlagSet("diagnose-clock", flag.ExitOnError)
	profileOpt := diagFlags.String("profile", "bmc0a-clock-fragility", "Diagnostic profile to run (bmc0a-clock-fragility)")
	outOpt := diagFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		diagFlags.Usage()
		os.Exit(1)
	}

	if err := diagFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		diagFlags.Usage()
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-clock-fragility" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-clock-fragility')\n", *profileOpt)
		os.Exit(1)
	}

	safeParams := model.DefaultSuperpositionSafeParams()

	rep, err := clockdiag.GenerateClockFragilityReport(safeParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating clock fragility report: %v\n", err)
		os.Exit(1)
	}

	if err := clockdiag.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran clock fragility diagnostic profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func readReport(path string) (*report.Report, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rep report.Report
	if err := json.Unmarshal(data, &rep); err != nil {
		return nil, err
	}

	return &rep, nil
}

func segmentClockCmd() {
	segFlags := flag.NewFlagSet("segment-clock", flag.ExitOnError)
	profileOpt := segFlags.String("profile", "bmc0a-clock-readiness", "Profile to segment (bmc0a-clock-readiness)")
	outOpt := segFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		segFlags.Usage()
		os.Exit(1)
	}

	if err := segFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		segFlags.Usage()
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-clock-readiness" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-clock-readiness')\n", *profileOpt)
		os.Exit(1)
	}

	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := clockseg.GenerateClockReadinessReport(safeParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating clock readiness report: %v\n", err)
		os.Exit(1)
	}

	if err := clockseg.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran clock readiness profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func specFriedmannCmd() {
	specFlags := flag.NewFlagSet("spec-friedmann", flag.ExitOnError)
	profileOpt := specFlags.String("profile", "bmc0a-friedmann-spec", "Profile to specify (bmc0a-friedmann-spec)")
	outOpt := specFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		specFlags.Usage()
		os.Exit(1)
	}

	if err := specFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		specFlags.Usage()
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-friedmann-spec" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-friedmann-spec')\n", *profileOpt)
		os.Exit(1)
	}

	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := friedmannspec.GenerateFriedmannSpecReport(safeParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating Friedmann spec report: %v\n", err)
		os.Exit(1)
	}

	if err := friedmannspec.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran Friedmann spec profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func specNullModelsCmd() {
	specFlags := flag.NewFlagSet("spec-nullmodels", flag.ExitOnError)
	profileOpt := specFlags.String("profile", "bmc0a-nullmodel-spec", "Profile to specify (bmc0a-nullmodel-spec)")
	outOpt := specFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		specFlags.Usage()
		os.Exit(1)
	}

	if err := specFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		specFlags.Usage()
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-nullmodel-spec" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-nullmodel-spec')\n", *profileOpt)
		os.Exit(1)
	}

	safeParams := model.DefaultSuperpositionSafeParams()
	rep, err := nullspec.GenerateNullModelSpecReport(safeParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error generating null model spec report: %v\n", err)
		os.Exit(1)
	}

	if err := nullspec.WriteJSON(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing report to %s: %v\n", *outOpt, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran Null Model spec profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func priorArtBoundaryCmd() {
	cmdFlags := flag.NewFlagSet("prior-art-boundary", flag.ExitOnError)
	profileOpt := cmdFlags.String("profile", "", "Profile to run (required)")
	outOpt := cmdFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		cmdFlags.Usage()
		os.Exit(1)
	}

	if err := cmdFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *profileOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --profile is required")
		os.Exit(1)
	}
	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-prior-art-boundary" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-prior-art-boundary')\n", *profileOpt)
		os.Exit(1)
	}

	rep := priorart.GenerateDefaultReport()
	if err := priorart.WriteReport(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing prior-art boundary report: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran prior-art boundary profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func runNullModelsCmd() {
	cmdFlags := flag.NewFlagSet("run-nullmodels", flag.ExitOnError)
	profileOpt := cmdFlags.String("profile", "", "Profile to run (required)")
	outOpt := cmdFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		cmdFlags.Usage()
		os.Exit(1)
	}

	if err := cmdFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *profileOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --profile is required")
		os.Exit(1)
	}
	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-nullrun" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-nullrun')\n", *profileOpt)
		os.Exit(1)
	}

	runs, err := nullrun.RunNullModels(*profileOpt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error running null models: %v\n", err)
		os.Exit(1)
	}

	rep := nullrun.GenerateDefaultReport()
	rep.NullModelRuns = runs

	if err := nullrun.WriteReport(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing null run report: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran null models profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

func findWorkspaceRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", fmt.Errorf("could not find workspace root (go.mod)")
}

func runResidualsCmd() {
	cmdFlags := flag.NewFlagSet("run-residuals", flag.ExitOnError)
	profileOpt := cmdFlags.String("profile", "", "Profile to run (required)")
	outOpt := cmdFlags.String("out", "", "Output path for the generated JSON report (required)")

	if len(os.Args) < 3 {
		cmdFlags.Usage()
		os.Exit(1)
	}

	if err := cmdFlags.Parse(os.Args[2:]); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(1)
	}

	if *profileOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --profile is required")
		os.Exit(1)
	}
	if *outOpt == "" {
		fmt.Fprintln(os.Stderr, "Error: --out is required")
		os.Exit(1)
	}

	if *profileOpt != "bmc0a-local-residual" {
		fmt.Fprintf(os.Stderr, "Error: profile '%s' is not supported (use 'bmc0a-local-residual')\n", *profileOpt)
		os.Exit(1)
	}

	root, err := findWorkspaceRoot()
	var rep *residualrun.ResidualRunReport
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not locate workspace root: %v. Generating default blocked report.\n", err)
		rep = residualrun.GenerateBlockedDefaultReport()
	} else {
		clockPath := filepath.Join(root, "out", "bmc0a_clock_readiness.json")
		friedmannPath := filepath.Join(root, "out", "bmc0a_friedmann_spec.json")
		nullrunPath := filepath.Join(root, "out", "bmc0a_nullrun.json")
		priorartPath := filepath.Join(root, "out", "bmc0a_prior_art_boundary.json")

		rep, err = residualrun.RunResidualsFromFiles(clockPath, friedmannPath, nullrunPath, priorartPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error running residuals: %v\n", err)
			os.Exit(1)
		}
	}

	if err := residualrun.WriteReport(rep, *outOpt); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing residual run report: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully ran residual profile '%s' and generated report: %s\n", *profileOpt, *outOpt)
}

