package priorart

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func helperDefaultReport(t *testing.T) *PriorArtBoundaryReport {
	t.Helper()
	return GenerateDefaultReport()
}

func TestPriorArtBoundaryReportValidation(t *testing.T) {
	rep := helperDefaultReport(t)
	errs := ValidateReport(rep, "{}")
	if len(errs) > 0 {
		t.Fatalf("Default report failed validation: %v", errs)
	}
}

func TestPriorArtBoundaryRejectsNoveltyClaimMade(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ScientificNoveltyClaimMade = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ScientificNoveltyClaimMade is true")
	}
}

func TestPriorArtBoundaryRejectsScientificNoveltyAllowed(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ScientificNoveltyClaimAllowed = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ScientificNoveltyClaimAllowed is true")
	}
}

func TestPriorArtBoundaryRejectsResidualComputed(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.ResidualComputed = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when ResidualComputed is true")
	}
}

func TestPriorArtBoundaryRejectsNullComparisonComputed(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.NullComparisonComputed = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when NullComparisonComputed is true")
	}
}

func TestPriorArtBoundaryRejectsRecoveryClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.RecoveryClaim = true
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when RecoveryClaim is true")
	}
}

func TestPriorArtBoundaryRequiresSeedSources(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.PriorArtSources = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when PriorArtSources is empty")
	}
}

func TestPriorArtBoundaryRequiresBoundaryClaims(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.BoundaryClaims = nil
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when BoundaryClaims is empty")
	}
}

func TestPriorArtBoundaryRejectsMissingRequiredBoundaryClaims(t *testing.T) {
	rep := helperDefaultReport(t)
	// Remove one required claim
	rep.BoundaryClaims = rep.BoundaryClaims[1:]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a required boundary claim is missing")
	}
}

func TestPriorArtBoundaryRejectsDuplicateSourceIDs(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.PriorArtSources = append(rep.PriorArtSources, rep.PriorArtSources[0])
	errs := ValidateReport(rep, "{}")
	foundDuplicateErr := false
	for _, e := range errs {
		if strings.Contains(e.Message, "duplicate source_id") || strings.Contains(e.Message, "duplicate source_id detected") {
			foundDuplicateErr = true
		}
	}
	if !foundDuplicateErr {
		t.Fatal("Expected duplicate source_id validation error")
	}
}

func TestPriorArtBoundaryRejectsDuplicateClaimIDs(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.BoundaryClaims = append(rep.BoundaryClaims, rep.BoundaryClaims[0])
	errs := ValidateReport(rep, "{}")
	foundDuplicateErr := false
	for _, e := range errs {
		if strings.Contains(e.Message, "duplicate claim_id") || strings.Contains(e.Message, "duplicate claim_id detected") {
			foundDuplicateErr = true
		}
	}
	if !foundDuplicateErr {
		t.Fatal("Expected duplicate claim_id validation error")
	}
}

func TestPriorArtBoundaryRejectsForbiddenBoundaryStatuses(t *testing.T) {
	for _, forbidden := range []string{"novel", "first_ever", "scientifically_original", "breakthrough", "proved_new", "validated_physics"} {
		rep := helperDefaultReport(t)
		rep.BoundaryClaims[0].BoundaryStatus = forbidden
		errs := ValidateReport(rep, "{}")
		if len(errs) == 0 {
			t.Fatalf("Expected failure for forbidden boundary status: %s", forbidden)
		}
	}
}

func TestPriorArtBoundaryRequiresAllGatesExactlyOnce(t *testing.T) {
	rep := helperDefaultReport(t)
	// Missing one gate
	gatesBackup := rep.Gates
	rep.Gates = rep.Gates[1:]
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a required gate is missing")
	}

	// Duplicated gate
	rep.Gates = append(gatesBackup, gatesBackup[0])
	errs = ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate is duplicated")
	}
}

func TestPriorArtBoundaryRejectsNonPassGate(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.Gates[0].Status = "fail"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when a gate status is not pass")
	}
}

func TestPriorArtBoundaryRejectsUnknownFields(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_unknown.json")

	raw := `{"schema_version":"bmc0a-prior-art-boundary-v0.1","extra_unknown_field":"some_val"}`
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err := ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with unknown fields")
	}
}

func TestPriorArtBoundaryRejectsTrailingJSONTokens(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "report_trailing.json")

	// Valid JSON content followed by trailing content
	rep := helperDefaultReport(t)
	data, err := json.Marshal(rep)
	if err != nil {
		t.Fatal(err)
	}
	raw := string(data) + "   true" // trailing token
	if err := os.WriteFile(path, []byte(raw), 0644); err != nil {
		t.Fatal(err)
	}

	_, err = ReadReport(path)
	if err == nil {
		t.Fatal("Expected error when reading report with trailing tokens")
	}
}

func TestPriorArtBoundaryDeterministicJSON(t *testing.T) {
	rep1 := helperDefaultReport(t)
	rep2 := helperDefaultReport(t)

	data1, err1 := json.Marshal(rep1)
	data2, err2 := json.Marshal(rep2)
	if err1 != nil || err2 != nil {
		t.Fatalf("Marshal failed: %v %v", err1, err2)
	}

	if string(data1) != string(data2) {
		t.Fatal("Serialized reports are not identical (non-deterministic)")
	}
}

func TestPriorArtBoundaryForbiddenPhraseScan(t *testing.T) {
	phrase1 := "first" + " " + "ever"
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This is the "+phrase1+" claim.")
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected forbidden phrase scan to catch forbidden phrase in warnings")
	}

	// Validate error message does not contain the phrase
	for _, e := range errs {
		if strings.Contains(e.Message, phrase1) {
			t.Fatalf("Error message contains forbidden phrase: %s", e.Message)
		}
	}
}

func TestPriorArtBoundaryCLIRouting(t *testing.T) {
	rep := GenerateDefaultReport()
	if rep.SchemaVersion != "bmc0a-prior-art-boundary-v0.1" {
		t.Fatalf("Unexpected schema version: %s", rep.SchemaVersion)
	}
}

func TestPriorArtBoundaryUnknownProfileFails(t *testing.T) {
	rep := helperDefaultReport(t)
	rep.SchemaVersion = "invalid-schema-v9.9"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail for wrong schema version")
	}
}

// ======================== Sprint 8.1 Repair Tests ========================

func TestPriorArtBoundaryRejectsWrongStatusDirectionForPhysicsClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	// Modify a physics-ingredient claim to a workflow status
	for idx, clm := range rep.BoundaryClaims {
		if clm.ClaimID == "bmc_uses_bohmian_minisuperspace_trajectories" {
			rep.BoundaryClaims[idx].BoundaryStatus = BoundaryStatusWorkflowDistinctiveCandidate
		}
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when physics claim is reclassified as workflow distinctive candidate")
	}
}

func TestPriorArtBoundaryRejectsWrongStatusDirectionForWorkflowClaim(t *testing.T) {
	rep := helperDefaultReport(t)
	// Modify a workflow claim to established prior art
	for idx, clm := range rep.BoundaryClaims {
		if clm.ClaimID == "bmc_defines_null_model_scaffold" {
			rep.BoundaryClaims[idx].BoundaryStatus = BoundaryStatusEstablishedPriorArt
		}
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when workflow claim is classified as established prior art")
	}
}

func TestPriorArtBoundaryRejectsWrongStatusDirectionForNotClaimedRecovery(t *testing.T) {
	rep := helperDefaultReport(t)
	for idx, clm := range rep.BoundaryClaims {
		if clm.ClaimID == "bmc_does_not_claim_recovery" {
			rep.BoundaryClaims[idx].BoundaryStatus = BoundaryStatusEstablishedPriorArt
		}
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when recovery claim is not 'not_claimed'")
	}
}

func TestPriorArtBoundaryRejectsWrongStatusDirectionForNoveltyBoundary(t *testing.T) {
	rep := helperDefaultReport(t)
	for idx, clm := range rep.BoundaryClaims {
		if clm.ClaimID == "bmc_does_not_claim_scientific_novelty" {
			rep.BoundaryClaims[idx].BoundaryStatus = BoundaryStatusLikelyPriorArt
		}
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure when novelty claim is not 'blocked'")
	}
}

func TestPriorArtBoundaryForbiddenBoundaryStatusErrorIsPhraseSafe(t *testing.T) {
	rep := helperDefaultReport(t)
	forbiddenVal := "scientifically" + "_" + "original"
	for idx, clm := range rep.BoundaryClaims {
		if clm.ClaimID == "bmc_uses_bohmian_minisuperspace_trajectories" {
			rep.BoundaryClaims[idx].BoundaryStatus = forbiddenVal
		}
	}
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure")
	}
	for _, e := range errs {
		if strings.Contains(e.Message, forbiddenVal) {
			t.Fatalf("Error message leaked forbidden boundary status: %s", e.Message)
		}
	}
}

func TestPriorArtBoundaryForbiddenReviewStatusErrorIsPhraseSafe(t *testing.T) {
	rep := helperDefaultReport(t)
	forbiddenVal := "confirms" + "_" + "novelty"
	rep.PriorArtSources[0].ReviewStatus = forbiddenVal
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected failure")
	}
	for _, e := range errs {
		if strings.Contains(e.Message, forbiddenVal) {
			t.Fatalf("Error message leaked forbidden review status: %s", e.Message)
		}
	}
}

func TestPriorArtBoundaryForbiddenPhraseErrorIsPhraseSafe(t *testing.T) {
	rep := helperDefaultReport(t)
	forbiddenVal := "first" + " " + "ever"
	rep.Warnings = append(rep.Warnings, "This is "+forbiddenVal)
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected failure")
	}
	for _, e := range errs {
		if strings.Contains(e.Message, forbiddenVal) {
			t.Fatalf("Error message leaked forbidden phrase: %s", e.Message)
		}
	}
}

func TestPriorArtBoundaryDeclaresDebtVocabulary(t *testing.T) {
	rep := helperDefaultReport(t)
	if rep.EbpDebtVocabulary != "ptw_runtime_debt_status_v0.1" {
		t.Fatalf("Unexpected debt vocabulary: %s", rep.EbpDebtVocabulary)
	}

	rep.EbpDebtVocabulary = "wrong_vocab"
	errs := ValidateReport(rep, "{}")
	if len(errs) == 0 {
		t.Fatal("Expected validation to fail when ebp_debt_vocabulary is wrong")
	}
}

func TestPriorArtBoundaryCLIRoutingRunsGenerateValidateSummarize(t *testing.T) {
	// Compile test binary
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test")

	// 1. Generate prior-art boundary note
	genCmd := exec.Command("./ptw-bmc-test", "prior-art-boundary", "--profile", "bmc0a-prior-art-boundary", "--out", "out_test.json")
	genCmd.Dir = "."
	if out, err := genCmd.CombinedOutput(); err != nil {
		t.Fatalf("Generation failed: %v, output: %s", err, string(out))
	}
	defer os.Remove("./out_test.json")

	// 2. Validate
	valCmd := exec.Command("./ptw-bmc-test", "validate", "--report", "out_test.json")
	valCmd.Dir = "."
	valOut, err := valCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Validation failed: %v, output: %s", err, string(valOut))
	}
	if !strings.Contains(string(valOut), "PASSED") {
		t.Fatalf("Unexpected validation output: %s", string(valOut))
	}

	// 3. Summarize
	sumCmd := exec.Command("./ptw-bmc-test", "summarize", "--report", "out_test.json")
	sumCmd.Dir = "."
	sumOut, err := sumCmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Summarization failed: %v, output: %s", err, string(sumOut))
	}
	if !strings.Contains(string(sumOut), "Prior-Art Sources Seeded: 5") {
		t.Fatalf("Unexpected summary output: %s", string(sumOut))
	}
}

func TestPriorArtBoundaryUnknownProfileFailsAtCLI(t *testing.T) {
	// Compile test binary
	cmd := exec.Command("go", "build", "-buildvcs=false", "-o", "ptw-bmc-test2", "../../../cmd/ptw-bmc")
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("Failed to build ptw-bmc CLI test binary: %v, output: %s", err, string(out))
	}
	defer os.Remove("./ptw-bmc-test2")

	// Run with unknown profile
	failCmd := exec.Command("./ptw-bmc-test2", "prior-art-boundary", "--profile", "unknown-profile", "--out", "out_test2.json")
	failCmd.Dir = "."
	out, err := failCmd.CombinedOutput()
	if err == nil {
		t.Fatal("Expected CLI to exit with failure for unknown profile")
	}
	if !strings.Contains(string(out), "is not supported") {
		t.Fatalf("Unexpected CLI error message: %s", string(out))
	}
}

func TestPriorArtBoundaryForbiddenPhraseScanIsCaseInsensitive(t *testing.T) {
	phraseMixed := "FiRsT" + " " + "EvEr"
	rep := helperDefaultReport(t)
	rep.Warnings = append(rep.Warnings, "This is "+phraseMixed)
	errs := ValidateReport(rep, "")
	if len(errs) == 0 {
		t.Fatal("Expected case-insensitive phrase scan to catch mixed case forbidden phrase")
	}
}
