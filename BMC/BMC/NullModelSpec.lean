import BMC.ToyReport

structure BMCNullModelSpecReport where
  toyAnalysisOnly           : Bool
  finalTruthClaim           : Bool
  residualComputed          : Bool
  nullComparisonComputed    : Bool
  friedmannRecoveryClaim    : Bool
  fullBMCBlocked            : Bool
  nullModelDebtActive       : Bool
  clockChoiceDebtActive     : Bool
  classicalTargetDebtActive : Bool
  unitConventionDebtActive  : Bool
  signConventionDebtActive  : Bool
  normalizationDebtActive   : Bool
  faithfulnessContested     : Bool
  deriving Repr

def reportPassesBMC0AFriedmannNullSpecGate (r : BMCNullModelSpecReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.residualComputed &&
  not r.nullComparisonComputed &&
  not r.friedmannRecoveryClaim &&
  r.fullBMCBlocked &&
  r.nullModelDebtActive &&
  r.clockChoiceDebtActive &&
  r.classicalTargetDebtActive &&
  r.unitConventionDebtActive &&
  r.signConventionDebtActive &&
  r.normalizationDebtActive &&
  r.faithfulnessContested

def reportPassesFullBMCForNullSpec (r : BMCNullModelSpecReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.fullBMCBlocked

-- Policy safety theorems

theorem null_model_spec_requires_toy_only
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.left.left.left.left.left.left.left.left.left

theorem null_model_spec_blocks_final_truth
  (r : BMCNullModelSpecReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AFriedmannNullSpecGate r = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate, h]

theorem null_model_spec_forbids_residual_computation
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.residualComputed = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  have h_comp := h.left.left.left.left.left.left.left.left.left.left.right
  exact h_comp

theorem null_model_spec_forbids_null_comparison_result
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.nullComparisonComputed = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  have h_comp := h.left.left.left.left.left.left.left.left.left.right
  exact h_comp

theorem null_model_spec_forbids_friedmann_recovery_claim
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.friedmannRecoveryClaim = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  have h_claim := h.left.left.left.left.left.left.left.left.right
  exact h_claim

theorem null_model_spec_requires_full_bmc_blocked
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.left.left.left.left.right

theorem null_model_spec_requires_null_model_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.nullModelDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.left.left.left.right

theorem null_model_spec_requires_clock_choice_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.clockChoiceDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.left.left.right

theorem null_model_spec_requires_classical_target_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.classicalTargetDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.left.right

theorem null_model_spec_requires_unit_convention_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.unitConventionDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.left.right

theorem null_model_spec_requires_sign_convention_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.signConventionDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.left.right

theorem null_model_spec_requires_normalization_debt_active
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.normalizationDebtActive = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.left.right

theorem null_model_spec_requires_faithfulness_contested
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.faithfulnessContested = true := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  exact h.right

theorem null_model_spec_does_not_imply_friedmann_recovery
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  r.friedmannRecoveryClaim = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  have h_claim := h.left.left.left.left.left.left.left.left.right
  exact h_claim

theorem null_model_spec_does_not_imply_full_bmc
  (r : BMCNullModelSpecReport)
  (h : reportPassesBMC0AFriedmannNullSpecGate r = true) :
  reportPassesFullBMCForNullSpec r = false := by
  simp [reportPassesBMC0AFriedmannNullSpecGate] at h
  have h_block := h.left.left.left.left.left.left.left.right
  simp [reportPassesFullBMCForNullSpec, h_block]

-- Witness and checks

def sprint7NullModelSpecWitness : BMCNullModelSpecReport := {
  toyAnalysisOnly           := true,
  finalTruthClaim           := false,
  residualComputed          := false,
  nullComparisonComputed    := false,
  friedmannRecoveryClaim    := false,
  fullBMCBlocked            := true,
  nullModelDebtActive       := true,
  clockChoiceDebtActive     := true,
  classicalTargetDebtActive := true,
  unitConventionDebtActive  := true,
  signConventionDebtActive  := true,
  normalizationDebtActive   := true,
  faithfulnessContested     := true
}

theorem null_model_spec_witness_passes_gate :
  reportPassesBMC0AFriedmannNullSpecGate sprint7NullModelSpecWitness = true := by decide

theorem null_model_spec_witness_fails_full_bmc :
  reportPassesFullBMCForNullSpec sprint7NullModelSpecWitness = false := by decide
