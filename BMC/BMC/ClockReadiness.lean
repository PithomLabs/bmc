import BMC.ToyReport

structure BMCClockReadinessReport where
  toyAnalysisOnly                 : Bool
  finalTruthClaim                 : Bool
  technicalGatePassed             : Bool
  friedmannDeferred               : Bool
  fullBMCBlocked                  : Bool
  clockChoiceDebtActive           : Bool
  faithfulnessContested           : Bool
  humanFaithfulnessReviewRequired : Bool
  hasValidLocalBranches           : Bool
  friedmannReadiness              : String
  deriving Repr

def reportPassesBMC0AClockReadinessGate (r : BMCClockReadinessReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  r.technicalGatePassed &&
  r.friedmannDeferred &&
  r.fullBMCBlocked &&
  r.clockChoiceDebtActive &&
  r.faithfulnessContested &&
  r.humanFaithfulnessReviewRequired &&
  (if r.friedmannReadiness == "local_only_candidate" then r.hasValidLocalBranches else true) &&
  (r.friedmannReadiness == "blocked" || r.friedmannReadiness == "local_only_candidate" || r.friedmannReadiness == "contested")

def reportPassesFullBMCForClockReadiness (r : BMCClockReadinessReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  r.technicalGatePassed &&
  not r.friedmannDeferred &&
  not r.fullBMCBlocked

-- Policy safety theorems

theorem clock_readiness_requires_toy_only
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  exact h.left.left.left.left.left.left.left.left.left

theorem clock_readiness_blocks_final_truth
  (r : BMCClockReadinessReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AClockReadinessGate r = false := by
  simp [reportPassesBMC0AClockReadinessGate, h]

theorem clock_readiness_requires_friedmann_deferred
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true) :
  r.friedmannDeferred = true := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  exact h.left.left.left.left.left.left.right

theorem clock_readiness_does_not_imply_full_bmc
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true) :
  reportPassesFullBMCForClockReadiness r = false := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  have h_block := h.left.left.left.left.left.right
  simp [reportPassesFullBMCForClockReadiness]
  intro _ _ _ _
  exact h_block

theorem clock_readiness_keeps_clock_choice_debt_active
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true) :
  r.clockChoiceDebtActive = true := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  exact h.left.left.left.left.right

theorem clock_readiness_local_candidate_does_not_mean_friedmann_recovered
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true)
  (_ : r.friedmannReadiness = "local_only_candidate") :
  r.friedmannDeferred = true ∧ r.fullBMCBlocked = true := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  exact ⟨h.left.left.left.left.left.left.right, h.left.left.left.left.left.right⟩

theorem clock_readiness_requires_faithfulness_contested_or_human_review
  (r : BMCClockReadinessReport)
  (h : reportPassesBMC0AClockReadinessGate r = true) :
  r.faithfulnessContested = true ∧ r.humanFaithfulnessReviewRequired = true := by
  simp [reportPassesBMC0AClockReadinessGate] at h
  exact ⟨h.left.left.left.right, h.left.left.right⟩

-- Witness and checks

def sprint5ClockReadinessWitness : BMCClockReadinessReport := {
  toyAnalysisOnly                 := true,
  finalTruthClaim                 := false,
  technicalGatePassed             := true,
  friedmannDeferred               := true,
  fullBMCBlocked                  := true,
  clockChoiceDebtActive           := true,
  hasValidLocalBranches           := true,
  friedmannReadiness              := "local_only_candidate",
  faithfulnessContested           := true,
  humanFaithfulnessReviewRequired := true
}

theorem clock_readiness_witness_passes_gate :
  reportPassesBMC0AClockReadinessGate sprint5ClockReadinessWitness = true := by decide

theorem clock_readiness_witness_fails_full_bmc :
  reportPassesFullBMCForClockReadiness sprint5ClockReadinessWitness = false := by decide
