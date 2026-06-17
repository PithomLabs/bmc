structure BMCResidualAuditReport where
  toyAnalysisOnly         : Bool
  finalTruthClaim         : Bool
  recoveryClaim           : Bool
  scientificNoveltyClaim  : Bool
  bmcBeatsNullModelsClaim : Bool
  fullBMCBlocked          : Bool
  faithfulnessContested   : Bool
  deriving Repr

def reportPassesResidualAuditGate (r : BMCResidualAuditReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.recoveryClaim &&
  not r.scientificNoveltyClaim &&
  not r.bmcBeatsNullModelsClaim &&
  r.fullBMCBlocked &&
  r.faithfulnessContested

theorem residualaudit_forbids_recovery_claim
  (r : BMCResidualAuditReport)
  (h : reportPassesResidualAuditGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesResidualAuditGate] at h
  exact h.left.left.left.left.right

theorem residualaudit_forbids_bmc_beats_nulls_claim
  (r : BMCResidualAuditReport)
  (h : reportPassesResidualAuditGate r = true) :
  r.bmcBeatsNullModelsClaim = false := by
  simp [reportPassesResidualAuditGate] at h
  exact h.left.left.right

theorem residualaudit_requires_full_bmc_blocked
  (r : BMCResidualAuditReport)
  (h : reportPassesResidualAuditGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesResidualAuditGate] at h
  exact h.left.right

theorem residualaudit_does_not_imply_residual_success
  (r : BMCResidualAuditReport)
  (h : reportPassesResidualAuditGate r = true) :
  r.bmcBeatsNullModelsClaim = false := by
  simp [reportPassesResidualAuditGate] at h
  exact h.left.left.right

def sprint11ResidualAuditWitness : BMCResidualAuditReport := {
  toyAnalysisOnly         := true,
  finalTruthClaim         := false,
  recoveryClaim           := false,
  scientificNoveltyClaim  := false,
  bmcBeatsNullModelsClaim := false,
  fullBMCBlocked          := true,
  faithfulnessContested   := true
}

theorem residualaudit_witness_passes_gate :
  reportPassesResidualAuditGate sprint11ResidualAuditWitness = true := by decide
