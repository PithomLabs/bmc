import BMC.ToyReport

structure BMCClockFragilityReport where
  toyAnalysisOnly       : Bool
  finalTruthClaim       : Bool
  technicalGatePassed   : Bool
  friedmannDeferred     : Bool
  fullBMCBlocked        : Bool
  clockChoiceDebtActive : Bool
  deriving Repr

def reportPassesBMC0AClockFragilityDiagnosticGate (r : BMCClockFragilityReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.technicalGatePassed &&
  r.friedmannDeferred &&
  r.fullBMCBlocked &&
  r.clockChoiceDebtActive

def reportPassesFullBMCForClockFragility (r : BMCClockFragilityReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  r.technicalGatePassed &&
  !r.friedmannDeferred &&
  !r.fullBMCBlocked

-- Policy safety theorems

theorem clock_fragility_requires_toy_only
  (r : BMCClockFragilityReport)
  (h : reportPassesBMC0AClockFragilityDiagnosticGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0AClockFragilityDiagnosticGate] at h
  exact h.left.left.left.left.left

theorem clock_fragility_blocks_final_truth
  (r : BMCClockFragilityReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AClockFragilityDiagnosticGate r = false := by
  simp [reportPassesBMC0AClockFragilityDiagnosticGate, h]

theorem clock_fragility_requires_friedmann_deferred
  (r : BMCClockFragilityReport)
  (h : reportPassesBMC0AClockFragilityDiagnosticGate r = true) :
  r.friedmannDeferred = true := by
  simp [reportPassesBMC0AClockFragilityDiagnosticGate] at h
  exact h.left.left.right

theorem clock_fragility_does_not_imply_full_bmc
  (r : BMCClockFragilityReport)
  (h : reportPassesBMC0AClockFragilityDiagnosticGate r = true) :
  reportPassesFullBMCForClockFragility r = false := by
  simp [reportPassesBMC0AClockFragilityDiagnosticGate] at h
  have h_blocked := h.left.right
  simp [reportPassesFullBMCForClockFragility, h_blocked]

theorem clock_fragility_keeps_clock_choice_debt_active
  (r : BMCClockFragilityReport)
  (h : reportPassesBMC0AClockFragilityDiagnosticGate r = true) :
  r.clockChoiceDebtActive = true := by
  simp [reportPassesBMC0AClockFragilityDiagnosticGate] at h
  exact h.right

-- Witness and checks

def sprint4ClockFragilityWitness : BMCClockFragilityReport := {
  toyAnalysisOnly       := true,
  finalTruthClaim       := false,
  technicalGatePassed   := true,
  friedmannDeferred     := true,
  fullBMCBlocked        := true,
  clockChoiceDebtActive := true
}

theorem clock_fragility_witness_passes_gate :
  reportPassesBMC0AClockFragilityDiagnosticGate sprint4ClockFragilityWitness = true := by decide

theorem clock_fragility_witness_fails_full_bmc :
  reportPassesFullBMCForClockFragility sprint4ClockFragilityWitness = false := by decide
