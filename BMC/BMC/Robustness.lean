import BMC.ToyReport

structure BMCRobustnessReport where
  toyAnalysisOnly      : Bool
  finalTruthClaim      : Bool
  technicalGateStatus  : CheckStatus
  promotionGateStatus  : CheckStatus
  friedmannResidual    : CheckStatus
  faithfulness         : CheckStatus
  deriving Repr

def reportPassesBMC0ARobustnessAuditGate (r : BMCRobustnessReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.technicalGateStatus &&
  checkDeferred r.friedmannResidual

def reportPassesFullBMCForRobustness (r : BMCRobustnessReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.technicalGateStatus &&
  checkPassed r.friedmannResidual &&
  checkPassed r.faithfulness

-- Policy safety theorems

theorem robustness_audit_requires_toy_only
  (r : BMCRobustnessReport)
  (h : reportPassesBMC0ARobustnessAuditGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0ARobustnessAuditGate] at h
  exact h.left.left.left

theorem robustness_audit_blocks_final_truth
  (r : BMCRobustnessReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0ARobustnessAuditGate r = false := by
  simp [reportPassesBMC0ARobustnessAuditGate, h]

theorem robustness_audit_requires_friedmann_deferred
  (r : BMCRobustnessReport)
  (h : reportPassesBMC0ARobustnessAuditGate r = true) :
  checkDeferred r.friedmannResidual = true := by
  simp [reportPassesBMC0ARobustnessAuditGate] at h
  exact h.right

theorem robustness_audit_does_not_imply_full_bmc
  (r : BMCRobustnessReport)
  (h : reportPassesBMC0ARobustnessAuditGate r = true) :
  reportPassesFullBMCForRobustness r = false := by
  simp [reportPassesBMC0ARobustnessAuditGate] at h
  have h_def := h.right
  cases hr : r.friedmannResidual
  · -- CheckStatus.pass
    simp [checkDeferred, hr] at h_def
  · -- CheckStatus.fail
    simp [reportPassesFullBMCForRobustness, checkPassed, hr]
  · -- CheckStatus.deferred
    simp [reportPassesFullBMCForRobustness, checkPassed, hr]
  · -- CheckStatus.contested
    simp [reportPassesFullBMCForRobustness, checkPassed, hr]

-- Robustness witness validation

def sprint3RobustnessWitness : BMCRobustnessReport := {
  toyAnalysisOnly      := true,
  finalTruthClaim      := false,
  technicalGateStatus  := CheckStatus.pass,
  promotionGateStatus  := CheckStatus.contested,
  friedmannResidual    := CheckStatus.deferred,
  faithfulness         := CheckStatus.contested
}

theorem robustness_witness_passes_audit :
  reportPassesBMC0ARobustnessAuditGate sprint3RobustnessWitness = true := by decide

theorem robustness_witness_fails_full_bmc :
  reportPassesFullBMCForRobustness sprint3RobustnessWitness = false := by decide
