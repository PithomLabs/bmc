structure BMCNullRunReport where
  toyAnalysisOnly               : Bool
  finalTruthClaim               : Bool
  residualComputed              : Bool
  nullDiagnosticsComputed       : Bool
  targetNullComparisonComputed  : Bool
  recoveryClaim                 : Bool
  scientificNoveltyClaim        : Bool
  fullBMCBlocked                : Bool
  faithfulnessContested         : Bool
  deriving Repr

def reportPassesNullRunGate (r : BMCNullRunReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.residualComputed &&
  r.nullDiagnosticsComputed &&
  r.targetNullComparisonComputed &&
  not r.recoveryClaim &&
  not r.scientificNoveltyClaim &&
  r.fullBMCBlocked &&
  r.faithfulnessContested

-- Policy safety theorems

theorem nullrun_forbids_residual_computation
  (r : BMCNullRunReport)
  (h : reportPassesNullRunGate r = true) :
  r.residualComputed = false := by
  simp [reportPassesNullRunGate] at h
  exact h.left.left.left.left.left.left.right

theorem nullrun_forbids_recovery_claim
  (r : BMCNullRunReport)
  (h : reportPassesNullRunGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesNullRunGate] at h
  exact h.left.left.left.right

theorem nullrun_requires_full_bmc_blocked
  (r : BMCNullRunReport)
  (h : reportPassesNullRunGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesNullRunGate] at h
  exact h.left.right

theorem nullrun_does_not_imply_bmc_beats_nulls
  (r : BMCNullRunReport)
  (h : reportPassesNullRunGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesNullRunGate] at h
  exact h.left.left.left.right

-- Witness and checks

def sprint9NullRunWitness : BMCNullRunReport := {
  toyAnalysisOnly               := true,
  finalTruthClaim               := false,
  residualComputed              := false,
  nullDiagnosticsComputed       := true,
  targetNullComparisonComputed  := true,
  recoveryClaim                 := false,
  scientificNoveltyClaim        := false,
  fullBMCBlocked                := true,
  faithfulnessContested         := true
}

theorem nullrun_witness_passes_gate :
  reportPassesNullRunGate sprint9NullRunWitness = true := by decide
