structure BMCPriorArtBoundaryReport where
  toyAnalysisOnly               : Bool
  finalTruthClaim               : Bool
  scientificNoveltyClaimMade    : Bool
  scientificNoveltyClaimAllowed : Bool
  residualComputed              : Bool
  nullComparisonComputed        : Bool
  recoveryClaim                 : Bool
  fullBMCBlocked                : Bool
  deriving Repr

def reportPassesPriorArtBoundaryGate (r : BMCPriorArtBoundaryReport) : Bool :=
  r.toyAnalysisOnly &&
  not r.finalTruthClaim &&
  not r.scientificNoveltyClaimMade &&
  not r.scientificNoveltyClaimAllowed &&
  not r.residualComputed &&
  not r.nullComparisonComputed &&
  not r.recoveryClaim &&
  r.fullBMCBlocked

-- Policy safety theorems

theorem prior_art_requires_toy_only
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.left.left.left.left.left.left.left

theorem prior_art_blocks_scientific_novelty
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.scientificNoveltyClaimMade = false := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.left.left.left.left.left.right

theorem prior_art_blocks_residual_computed
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.residualComputed = false := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.left.left.left.right

theorem prior_art_blocks_null_comparison_computed
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.nullComparisonComputed = false := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.left.left.right

theorem prior_art_blocks_recovery_claim
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.recoveryClaim = false := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.left.right

theorem prior_art_blocks_full_bmc
  (r : BMCPriorArtBoundaryReport)
  (h : reportPassesPriorArtBoundaryGate r = true) :
  r.fullBMCBlocked = true := by
  simp [reportPassesPriorArtBoundaryGate] at h
  exact h.right

-- Witness and checks

def sprint8PriorArtBoundaryWitness : BMCPriorArtBoundaryReport := {
  toyAnalysisOnly               := true,
  finalTruthClaim               := false,
  scientificNoveltyClaimMade    := false,
  scientificNoveltyClaimAllowed := false,
  residualComputed              := false,
  nullComparisonComputed        := false,
  recoveryClaim                 := false,
  fullBMCBlocked                := true
}

theorem prior_art_witness_passes_gate :
  reportPassesPriorArtBoundaryGate sprint8PriorArtBoundaryWitness = true := by decide
