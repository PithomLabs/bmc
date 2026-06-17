package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/PithomLabs/bmc/internal/bmc/audit"
	"github.com/PithomLabs/bmc/internal/bmc/clockdiag"
	"github.com/PithomLabs/bmc/internal/bmc/clockseg"
	"github.com/PithomLabs/bmc/internal/bmc/friedmannspec"
	"github.com/PithomLabs/bmc/internal/bmc/model"
	"github.com/PithomLabs/bmc/internal/bmc/report"
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
