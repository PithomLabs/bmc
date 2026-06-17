package residualaudit

import (
	"encoding/json"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/PithomLabs/bmc/internal/bmc/nullrun"
	"github.com/PithomLabs/bmc/internal/bmc/residualrun"
)

func helperAuditReport(t *testing.T) *ResidualAuditReport {
	t.Helper()
	rep := GenerateBlockedDefaultReport()
	for idx := range rep.SourceArtifacts {
		rep.SourceArtifacts[idx].Provenance = ProvenanceFileRead
		rep.SourceArtifacts[idx].Path = "out/mock.json"
		rep.SourceArtifacts[idx].Status = "available"
	}
	rep.ResidualAuditComputed = true
	rep.InterpretationStatus = InterpretComparisonStructurallyHonest
	rep.ComparisonAudits = []ResidualComparisonAudit{{
		AuditID:              "audit_bmc0a_residual_vs_null_v0_1",
		SourceComparisonID:   "bmc0a-residual-vs-null-v0.1",
		AuditComputed:        true,
		AuditStatus:          AuditStatusComparisonAudited,
		AuditProvenance:      ProvenanceDerivedFromFileRead,
		TargetResidualIDs:    []string{"candidate_residual_branch_0"},
		NullModelIDs:         []string{"constant_phase_control"},
		MetricsAudited:       []string{"mean_abs_residual"},
		Findings:             []string{"Comparison references computed diagnostics."},
		InterpretationStatus: InterpretComparisonStructurallyHonest,
		Notes:                "Comparison integrity audit is structural only.",
	}}
	base := 0.1
	pert := 0.10000001
	abs := math.Abs(pert - base)
	rel := abs / base
	rep.StabilityAudits = []ResidualStabilityAudit{{
		StabilityID:           "stability_branch_0_alpha_point_perturbation",
		BranchID:              "branch_0",
		StabilityComputed:     true,
		PerturbationKind:      PerturbAlphaPoint,
		PerturbationMagnitude: 1e-6,
		BaselineMetric:        "mean_abs_residual",
		BaselineValue:         &base,
		PerturbedValue:        &pert,
		AbsoluteDelta:         &abs,
		RelativeDelta:         &rel,
		StabilityStatus:       StabilityStable,
		Notes:                 "Deterministic candidate-only stability diagnostic.",
	}}
	return rep
}

func helperRaw(t *testing.T, rep *ResidualAuditReport) string {
	t.Helper()
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}

func TestResidualAuditReportValidation(t *testing.T) {
	rep := helperAuditReport(t)
	if errs := ValidateReport(rep, helperRaw(t, rep)); len(errs) > 0 {
		t.Fatalf("Expected audit report to validate, got %v", errs)
	}
}

func TestResidualAuditRejectsRecoveryClaim(t *testing.T) {
	rep := helperAuditReport(t)
	rep.RecoveryClaim = true
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected recovery claim rejection")
	}
}

func TestResidualAuditRejectsScientificNoveltyClaim(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ScientificNoveltyClaimMade = true
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected scientific novelty claim rejection")
	}
}

func TestResidualAuditRejectsBMCBeatsNullModelsClaim(t *testing.T) {
	rep := helperAuditReport(t)
	rep.BmcBeatsNullModelsClaim = true
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected superiority claim rejection")
	}
}

func TestResidualAuditRequiresFullBMCBlocked(t *testing.T) {
	rep := helperAuditReport(t)
	rep.FullBmcToyGate = "open"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected full BMC gate rejection")
	}
}

func TestResidualAuditRequiresLocalOnlyBoundary(t *testing.T) {
	rep := helperAuditReport(t)
	rep.LocalBranchOnly = false
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected local-only boundary rejection")
	}
}

func TestResidualAuditRejectsGlobalCosmologyClaim(t *testing.T) {
	rep := helperAuditReport(t)
	rep.GlobalCosmologyClaim = true
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected global-cosmology boundary rejection")
	}
}

func TestResidualAuditRequiresSourceArtifacts(t *testing.T) {
	rep := helperAuditReport(t)
	rep.SourceArtifacts = rep.SourceArtifacts[:4]
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected missing source artifact rejection")
	}
}

func TestResidualAuditRejectsDuplicateSourceArtifact(t *testing.T) {
	rep := helperAuditReport(t)
	rep.SourceArtifacts = append(rep.SourceArtifacts, rep.SourceArtifacts[0])
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected duplicate source artifact rejection")
	}
}

func TestResidualAuditRejectsUnknownSourceArtifact(t *testing.T) {
	rep := helperAuditReport(t)
	rep.SourceArtifacts[0].ArtifactID = "unknown_source"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected unknown source artifact rejection")
	}
}

func TestResidualAuditRejectsFileReadWithoutPath(t *testing.T) {
	rep := helperAuditReport(t)
	rep.SourceArtifacts[0].Path = ""
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected file_read without path rejection")
	}
}

func TestResidualAuditRejectsDecorativeComparisonAudit(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].MetricsAudited = []string{}
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected decorative comparison audit rejection")
	}
}

func TestResidualAuditRejectsForbiddenAuditStatus(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].AuditStatus = "winner"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected forbidden audit status rejection")
	}
}

func TestResidualAuditRejectsForbiddenAuditProvenance(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].AuditProvenance = "fabricated"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected forbidden audit provenance rejection")
	}
}

func TestResidualAuditRejectsForbiddenInterpretationStatus(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].InterpretationStatus = "passed"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected forbidden interpretation status rejection")
	}
}

func TestResidualAuditRejectsNonfiniteStabilityMetrics(t *testing.T) {
	rep := helperAuditReport(t)
	v := math.NaN()
	rep.StabilityAudits[0].BaselineValue = &v
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected nonfinite stability metric rejection")
	}
}

func TestResidualAuditRejectsNegativePerturbationMagnitude(t *testing.T) {
	rep := helperAuditReport(t)
	rep.StabilityAudits[0].PerturbationMagnitude = -1
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected negative perturbation rejection")
	}
}

func TestResidualAuditRejectsNegativeAbsoluteDelta(t *testing.T) {
	rep := helperAuditReport(t)
	v := -1.0
	rep.StabilityAudits[0].AbsoluteDelta = &v
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected negative absolute_delta rejection")
	}
}

func TestResidualAuditRejectsNegativeRelativeDelta(t *testing.T) {
	rep := helperAuditReport(t)
	v := -1.0
	rep.StabilityAudits[0].RelativeDelta = &v
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected negative relative_delta rejection")
	}
}

func TestResidualAuditRejectsComputedStabilityWithNotComputedStatus(t *testing.T) {
	rep := helperAuditReport(t)
	rep.StabilityAudits[0].StabilityStatus = StabilityNotComputed
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected computed stability with not_computed status rejection")
	}
}

func TestResidualAuditRejectsNoncomputedStabilityMarkedStable(t *testing.T) {
	rep := helperAuditReport(t)
	rep.StabilityAudits[0].StabilityComputed = false
	rep.StabilityAudits[0].StabilityStatus = StabilityStable
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected noncomputed stability marked stable rejection")
	}
}

func TestResidualAuditRejectsNoncomputedStabilityMarkedSensitive(t *testing.T) {
	rep := helperAuditReport(t)
	rep.StabilityAudits[0].StabilityComputed = false
	rep.StabilityAudits[0].StabilityStatus = StabilitySensitive
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected noncomputed stability marked sensitive rejection")
	}
}

func TestResidualAuditRequiresAllGatesExactlyOnce(t *testing.T) {
	rep := helperAuditReport(t)
	rep.Gates = rep.Gates[:len(rep.Gates)-1]
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected missing gate rejection")
	}
	rep = helperAuditReport(t)
	rep.Gates = append(rep.Gates, rep.Gates[0])
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected duplicate gate rejection")
	}
}

func TestResidualAuditRejectsNonPassGate(t *testing.T) {
	rep := helperAuditReport(t)
	rep.Gates[0].Status = "fail"
	if len(ValidateReport(rep, "{}")) == 0 {
		t.Fatal("Expected non-pass gate rejection")
	}
}

func TestResidualAuditRejectsUnknownFields(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	if err := os.WriteFile(path, []byte(`{"schema_version":"bmc0a-residual-audit-v0.1","unknown":true}`), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadReport(path); err == nil {
		t.Fatal("Expected strict decoder to reject unknown field")
	}
}

func TestResidualAuditRejectsTrailingJSONTokens(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "bad.json")
	rep := helperAuditReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(data, []byte(` {}`)...), 0644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadReport(path); err == nil {
		t.Fatal("Expected strict decoder to reject trailing JSON tokens")
	}
}

func TestResidualAuditDeterministicJSON(t *testing.T) {
	rep1 := helperAuditReport(t)
	rep2 := helperAuditReport(t)
	data1, err := json.MarshalIndent(rep1, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	data2, err := json.MarshalIndent(rep2, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if string(data1) != string(data2) {
		t.Fatal("Expected deterministic JSON bytes")
	}
}

func TestResidualAuditForbiddenPhraseScan(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].Notes = "winner"
	if len(ValidateReport(rep, helperRaw(t, rep))) == 0 {
		t.Fatal("Expected forbidden phrase rejection")
	}
}

func TestResidualAuditForbiddenPhraseErrorsArePhraseSafe(t *testing.T) {
	rep := helperAuditReport(t)
	rep.ComparisonAudits[0].Notes = "winner"
	errs := ValidateReport(rep, helperRaw(t, rep))
	if len(errs) == 0 {
		t.Fatal("Expected forbidden phrase rejection")
	}
	for _, err := range errs {
		if strings.Contains(strings.ToLower(err.Message), "winner") {
			t.Fatalf("Validation error leaked forbidden phrase: %s", err.Message)
		}
	}
}

func TestResidualAuditCLIRoutingRunsGenerateValidateSummarize(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping CLI build in short mode")
	}
	dir := t.TempDir()
	bin := buildTestCLI(t, dir)
	outPath := filepath.Join(dir, "audit.json")

	genCmd := exec.Command(bin, "audit-residuals", "--profile", "bmc0a-residual-audit", "--out", outPath)
	if out, err := genCmd.CombinedOutput(); err != nil {
		t.Fatalf("audit-residuals failed: %v\n%s", err, out)
	}
	validateCmd := exec.Command(bin, "validate", "--report", outPath)
	if out, err := validateCmd.CombinedOutput(); err != nil {
		t.Fatalf("validate failed: %v\n%s", err, out)
	}
	summaryCmd := exec.Command(bin, "summarize", "--report", outPath)
	out, err := summaryCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("summarize failed: %v\n%s", err, out)
	}
	if !strings.Contains(string(out), "BMC Sprint 11 Residual/Null Comparison Audit Summary") {
		t.Fatalf("summary did not route to residual audit summary:\n%s", out)
	}
}

func TestResidualAuditUnknownProfileFailsAtCLI(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping CLI build in short mode")
	}
	dir := t.TempDir()
	bin := buildTestCLI(t, dir)
	cmd := exec.Command(bin, "audit-residuals", "--profile", "unknown-profile", "--out", filepath.Join(dir, "out.json"))
	if err := cmd.Run(); err == nil {
		t.Fatal("Expected unknown profile to fail")
	}
}

func TestResidualAuditGeneratedFromInputs(t *testing.T) {
	residualRep := helperResidualRunReport()
	nullRep := nullrun.GenerateDefaultReport()
	sources := defaultSourceArtifacts()
	for idx := range sources {
		sources[idx].Provenance = ProvenanceFileRead
		sources[idx].Path = "out/mock.json"
		sources[idx].Status = "available"
	}
	rep := RunAuditFromInputs(residualRep, nullRep, sources)
	if !rep.ResidualAuditComputed {
		t.Fatal("Expected generated audit to be computed")
	}
	if rep.InterpretationStatus != InterpretComparisonStructurallyHonest && rep.InterpretationStatus != InterpretComparisonStabilityMixed {
		t.Fatalf("Unexpected interpretation status: %s", rep.InterpretationStatus)
	}
	if len(rep.ComparisonAudits) == 0 || len(rep.StabilityAudits) == 0 {
		t.Fatal("Expected comparison and stability audit records")
	}
}

func TestResidualAuditStabilityRecomputesAfterPerturbation(t *testing.T) {
	diag := helperResidualRunReport().CandidateResidualDiagnostics[0]

	alphaAudit := computeStabilityAudit(diag, PerturbAlphaPoint, 0.1, "mean_abs_residual")
	if !alphaAudit.StabilityComputed {
		t.Fatalf("Expected alpha perturbation audit to compute: %+v", alphaAudit)
	}
	baseline, ok := recomputeMetricFromResidualInputPoints(diag.ResidualInputPoints, "mean_abs_residual")
	if !ok {
		t.Fatal("Expected baseline metric recomputation")
	}
	copied := copyResidualInputPoints(diag.ResidualInputPoints)
	if !perturbAlphaPoint(copied, 0.1) {
		t.Fatal("Expected alpha perturbation to succeed")
	}
	expected, ok := recomputeMetricFromResidualInputPoints(copied, "mean_abs_residual")
	if !ok {
		t.Fatal("Expected perturbed metric recomputation")
	}
	decorative := baseline * (1 + 0.1)
	if alphaAudit.PerturbedValue == nil {
		t.Fatal("Expected perturbed value")
	}
	if math.Abs(*alphaAudit.PerturbedValue-expected) > 1e-12 {
		t.Fatalf("Expected recomputed perturbed metric %v, got %v", expected, *alphaAudit.PerturbedValue)
	}
	if math.Abs(expected-decorative) < 1e-12 {
		t.Fatalf("Test fixture is not adversarial: recomputed metric equals decorative baseline scaling")
	}
	if math.Abs(*alphaAudit.PerturbedValue-decorative) < 1e-12 {
		t.Fatalf("Perturbed metric used decorative baseline scaling: got %v", *alphaAudit.PerturbedValue)
	}

	phiAudit := computeStabilityAudit(diag, PerturbPhiPoint, 0.1, "max_abs_residual")
	if !phiAudit.StabilityComputed || phiAudit.PerturbedValue == nil {
		t.Fatalf("Expected phi perturbation max audit to compute: %+v", phiAudit)
	}
	lambdaAudit := computeStabilityAudit(diag, PerturbLambdaSpacing, 0.1, "rms_residual")
	if !lambdaAudit.StabilityComputed || lambdaAudit.PerturbedValue == nil {
		t.Fatalf("Expected lambda perturbation rms audit to compute: %+v", lambdaAudit)
	}
	diag.ResidualInputPoints[0].Lambda = nil
	diag.ResidualInputPoints[1].Lambda = nil
	lambdaNoCoordinateAudit := computeStabilityAudit(diag, PerturbLambdaSpacing, 0.1, "mean_abs_residual")
	if !lambdaNoCoordinateAudit.StabilityComputed || lambdaNoCoordinateAudit.PerturbedValue == nil {
		t.Fatalf("Expected interval-proxy lambda perturbation to compute without coordinate lambda: %+v", lambdaNoCoordinateAudit)
	}
}

func helperResidualRunReport() *residualrun.ResidualRunReport {
	rep := residualrun.GenerateBlockedDefaultReport()
	rep.CandidateResidualComputed = true
	rep.InterpretationStatus = residualrun.InterpretDiagComparisonOnly
	rep.ResidualNullComparisons = []residualrun.ResidualNullComparison{{
		ComparisonID:         "bmc0a-residual-vs-null-v0.1",
		TargetResidualIDs:    []string{"candidate_residual_branch_0"},
		NullModelIDs:         []string{"constant_phase_control"},
		MetricsCompared:      []string{"mean_abs_residual"},
		ComparisonComputed:   true,
		InterpretationStatus: residualrun.InterpretDiagComparisonOnly,
		Reason:               "Candidate-only diagnostic comparison.",
	}}
	l0, l1 := 0.1, 0.2
	a0, a1 := 1.0, 1.1
	p0, p1 := 0.2, 0.3
	lhs0, rhs0 := 1.0, 0.25
	lhs1, rhs1 := 1.2, 0.35
	mean, maxv, rms := 0.8, 0.85, 0.81
	rep.CandidateResidualDiagnostics[0] = residualrun.CandidateResidualDiagnostic{
		BranchID:           "branch_0",
		ResidualID:         "candidate_residual_branch_0",
		ResidualComputed:   true,
		ResidualStatus:     residualrun.ResidualStatusGenerated,
		ResidualProvenance: residualrun.ProvenanceComputed,
		Metrics: residualrun.CandidateResidualMetrics{
			NumEvaluationPoints:     2,
			NumFiniteResidualPoints: 2,
			MeanAbsResidual:         &mean,
			MaxAbsResidual:          &maxv,
			RmsResidual:             &rms,
			ResidualFinite:          true,
			DiagnosticWarnings:      []string{},
		},
		Notes: "Candidate residual diagnostic.",
		ResidualInputPoints: []residualrun.ResidualInputPoint{
			{BranchID: "branch_0", PointIndex: 0, Lambda: &l0, Alpha: &a0, Phi: &p0, CandidateLeftHandSide: &lhs0, CandidateRightHandSide: &rhs0, InputProvenance: "derived_from_file_read"},
			{BranchID: "branch_0", PointIndex: 1, Lambda: &l1, Alpha: &a1, Phi: &p1, CandidateLeftHandSide: &lhs1, CandidateRightHandSide: &rhs1, InputProvenance: "derived_from_file_read"},
		},
	}
	return rep
}

func buildTestCLI(t *testing.T, dir string) string {
	t.Helper()
	bin := filepath.Join(dir, "ptw-bmc-test")
	build := exec.Command("go", "build", "-buildvcs=false", "-o", bin, "../../../cmd/ptw-bmc")
	build.Env = append(os.Environ(), "GOCACHE=/tmp/go-build-cache")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("build failed: %v\n%s", err, out)
	}
	return bin
}
