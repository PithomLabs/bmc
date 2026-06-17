import BMC.ToyReport

structure BMCFriedmannSpecReport where
  toyAnalysisOnly           : Bool
  finalTruthClaim           : Bool
  residualComputed          : Bool
  friedmannRecoveryClaim    : Bool
  fullBMCBlocked            : Bool
  clockChoiceDebtActive     : Bool
  classicalTargetDebtActive : Bool
  unitConventionDebtActive  : Bool
  signConventionDebtActive  : Bool
  normalizationDebtActive   : Bool
  nullModelDebtActive       : Bool
  faithfulnessContested     : Bool
  deriving Repr

def reportPassesBMC0AFriedmannSpecGate (r : BMCFriedmannSpecReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.residualComputed &&
  not r.friedmannRecoveryClaim &&
  r.fullBMCBlocked &&
  r.clockChoiceDebtActive &&
  r.classicalTargetDebtActive &&
  r.unitConventionDebtActive &&
  r.signConventionDebtActive &&
  r.normalizationDebtActive &&
  r.nullModelDebtActive &&
  r.faithfulnessContested

def reportPassesFullBMCForFriedmannSpec (r : BMCFriedmannSpecReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.fullBMCBlocked

-- Policy safety theorems

theorem friedmann_spec_requires_toy_only
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.left.left.left.left.left.left.left.left

theorem friedmann_spec_blocks_final_truth
  (r : BMCFriedmannSpecReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AFriedmannSpecGate r = false := by
  simp [reportPassesBMC0AFriedmannSpecGate, h]

theorem friedmann_spec_forbids_residual_computation
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.residualComputed = false := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  have h_comp := h.left.left.left.left.left.left.left.left.left.right
  exact h_comp

theorem friedmann_spec_forbids_recovery_claim
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.friedmannRecoveryClaim = false := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  have h_claim := h.left.left.left.left.left.left.left.left.right
  exact h_claim

theorem friedmann_spec_requires_full_bmc_blocked
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.left.left.left.left.right

theorem friedmann_spec_requires_clock_choice_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.clockChoiceDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.left.left.left.right

theorem friedmann_spec_requires_classical_target_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.classicalTargetDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.left.left.right

theorem friedmann_spec_requires_unit_convention_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.unitConventionDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.left.right

theorem friedmann_spec_requires_sign_convention_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.signConventionDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.left.right

theorem friedmann_spec_requires_normalization_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.normalizationDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.left.right

theorem friedmann_spec_requires_null_model_debt_active
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.nullModelDebtActive = true := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  exact h.left.right

theorem friedmann_spec_does_not_imply_friedmann_recovery
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  r.friedmannRecoveryClaim = false := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  have h_claim := h.left.left.left.left.left.left.left.left.right
  exact h_claim

theorem friedmann_spec_does_not_imply_full_bmc
  (r : BMCFriedmannSpecReport)
  (h : reportPassesBMC0AFriedmannSpecGate r = true) :
  reportPassesFullBMCForFriedmannSpec r = false := by
  simp [reportPassesBMC0AFriedmannSpecGate] at h
  have h_block := h.left.left.left.left.left.left.left.right
  simp [reportPassesFullBMCForFriedmannSpec, h_block]

-- Witness and checks

def sprint6FriedmannSpecWitness : BMCFriedmannSpecReport := {
  toyAnalysisOnly           := true,
  finalTruthClaim           := false,
  residualComputed          := false,
  friedmannRecoveryClaim    := false,
  fullBMCBlocked            := true,
  clockChoiceDebtActive     := true,
  classicalTargetDebtActive := true,
  unitConventionDebtActive  := true,
  signConventionDebtActive  := true,
  normalizationDebtActive   := true,
  nullModelDebtActive       := true,
  faithfulnessContested     := true
}

theorem friedmann_spec_witness_passes_gate :
  reportPassesBMC0AFriedmannSpecGate sprint6FriedmannSpecWitness = true := by decide

theorem friedmann_spec_witness_fails_full_bmc :
  reportPassesFullBMCForFriedmannSpec sprint6FriedmannSpecWitness = false := by decide
