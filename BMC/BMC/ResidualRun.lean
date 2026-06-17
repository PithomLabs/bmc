structure BMCResidualRunReport where
  toyAnalysisOnly               : Bool
  finalTruthClaim               : Bool
  candidateResidualComputed     : Bool
  recoveryClaim                 : Bool
  scientificNoveltyClaim        : Bool
  bmcBeatsNullModelsClaim       : Bool
  fullBMCBlocked                : Bool
  faithfulnessContested         : Bool
  blockedByNoEligibleBranch     : Bool
  deriving Repr

def reportPassesResidualRunGate (r : BMCResidualRunReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  (r.candidateResidualComputed == not r.blockedByNoEligibleBranch) &&
  not r.recoveryClaim &&
  not r.scientificNoveltyClaim &&
  not r.bmcBeatsNullModelsClaim &&
  r.fullBMCBlocked &&
  r.faithfulnessContested

-- Policy safety theorems

theorem residualrun_forbids_recovery_claim
  (r : BMCResidualRunReport)
  (h : reportPassesResidualRunGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesResidualRunGate] at h
  exact h.left.left.left.left.right

theorem residualrun_forbids_bmc_beats_nulls_claim
  (r : BMCResidualRunReport)
  (h : reportPassesResidualRunGate r = true) :
  r.bmcBeatsNullModelsClaim = false := by
  simp [reportPassesResidualRunGate] at h
  exact h.left.left.right

theorem residualrun_requires_full_bmc_blocked
  (r : BMCResidualRunReport)
  (h : reportPassesResidualRunGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesResidualRunGate] at h
  exact h.left.right

theorem residualrun_does_not_imply_classical_limit
  (r : BMCResidualRunReport)
  (h : reportPassesResidualRunGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesResidualRunGate] at h
  exact h.left.left.left.left.right

-- Witness and checks

def sprint10ResidualRunWitness : BMCResidualRunReport := {
  toyAnalysisOnly               := true,
  finalTruthClaim               := false,
  candidateResidualComputed     := true,
  recoveryClaim                 := false,
  scientificNoveltyClaim        := false,
  bmcBeatsNullModelsClaim       := false,
  fullBMCBlocked                := true,
  faithfulnessContested         := true,
  blockedByNoEligibleBranch     := false
}

theorem residualrun_witness_passes_gate :
  reportPassesResidualRunGate sprint10ResidualRunWitness = true := by decide

def sprint10ResidualRunBlockedWitness : BMCResidualRunReport := {
  toyAnalysisOnly               := true,
  finalTruthClaim               := false,
  candidateResidualComputed     := false,
  recoveryClaim                 := false,
  scientificNoveltyClaim        := false,
  bmcBeatsNullModelsClaim       := false,
  fullBMCBlocked                := true,
  faithfulnessContested         := true,
  blockedByNoEligibleBranch     := true
}

theorem residualrun_blocked_witness_passes_gate :
  reportPassesResidualRunGate sprint10ResidualRunBlockedWitness = true := by decide

