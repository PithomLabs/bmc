import BMC.ToyReport

def reportPassesBMC0AControlGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.nodeDetection && -- added for structure compatibility
  checkPassed r.nodeContactFree && -- added for structure compatibility
  checkPassed r.qFiniteAwayFromNodes && -- added for structure compatibility
  checkPassed r.phaseGradientFinite && -- added for structure compatibility
  checkPassed r.classicalLimit &&
  checkDeferred r.friedmannResidual

def reportPassesFullBMCToyGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.nodeDetection &&
  checkPassed r.nodeContactFree &&
  checkPassed r.qFiniteAwayFromNodes &&
  checkPassed r.phaseGradientFinite &&
  checkPassed r.classicalLimit &&
  checkPassed r.friedmannResidual &&
  checkPassed r.faithfulness

def reportPassesBMC0ASuperpositionSafeGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.wdwResidual &&
  checkPassed r.trajectoryFinite &&
  checkPassed r.clockMonotonic &&
  checkPassed r.nodeDetection &&
  checkPassed r.nodeContactFree &&
  checkPassed r.qFiniteAwayFromNodes &&
  checkPassed r.phaseGradientFinite &&
  checkPassed r.classicalLimit &&
  checkDeferred r.friedmannResidual

def reportPassesBMC0ANodeDetectionGate (r : BMCReport) : Bool :=
  r.toyAnalysisOnly &&
  !r.finalTruthClaim &&
  checkPassed r.nodeDetection &&
  checkPassed r.wdwResidual &&
  checkDeferred r.friedmannResidual

-- Theorem obligations

theorem final_truth_blocks_control_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0AControlGate r = false := by
  simp [reportPassesBMC0AControlGate, h]

theorem final_truth_blocks_toy_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesFullBMCToyGate r = false := by
  simp [reportPassesFullBMCToyGate, h]

theorem final_truth_blocks_superposition_safe_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0ASuperpositionSafeGate r = false := by
  simp [reportPassesBMC0ASuperpositionSafeGate, h]

theorem final_truth_blocks_node_detection_gate
  (r : BMCReport)
  (h : r.finalTruthClaim = true) :
  reportPassesBMC0ANodeDetectionGate r = false := by
  simp [reportPassesBMC0ANodeDetectionGate, h]

theorem control_gate_requires_toy_only
  (r : BMCReport)
  (h : reportPassesBMC0AControlGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0AControlGate] at h
  exact h.left.left.left.left.left.left.left.left.left.left

theorem toy_gate_requires_toy_only
  (r : BMCReport)
  (h : reportPassesFullBMCToyGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesFullBMCToyGate] at h
  exact h.left.left.left.left.left.left.left.left.left.left.left

theorem superposition_safe_gate_requires_toy_only
  (r : BMCReport)
  (h : reportPassesBMC0ASuperpositionSafeGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0ASuperpositionSafeGate] at h
  exact h.left.left.left.left.left.left.left.left.left.left

theorem node_detection_gate_requires_toy_only
  (r : BMCReport)
  (h : reportPassesBMC0ANodeDetectionGate r = true) :
  r.toyAnalysisOnly = true := by
  simp [reportPassesBMC0ANodeDetectionGate] at h
  exact h.left.left.left.left

theorem faithfulness_required_for_full_gate
  (r : BMCReport)
  (h : reportPassesFullBMCToyGate r = true) :
  checkPassed r.faithfulness = true := by
  simp [reportPassesFullBMCToyGate] at h
  exact h.right

theorem friedmann_deferred_in_control_gate
  (r : BMCReport)
  (h : reportPassesBMC0AControlGate r = true) :
  checkDeferred r.friedmannResidual = true := by
  simp [reportPassesBMC0AControlGate] at h
  exact h.right

theorem superposition_safe_gate_requires_no_node_contact
  (r : BMCReport)
  (h : reportPassesBMC0ASuperpositionSafeGate r = true) :
  checkPassed r.nodeContactFree = true := by
  simp [reportPassesBMC0ASuperpositionSafeGate] at h
  exact h.left.left.left.left.right

-- Witnesses and check validations

def sprint2SafeWitness : BMCReport := {
  toyAnalysisOnly      := true,
  finalTruthClaim      := false,
  wdwResidual          := CheckStatus.pass,
  trajectoryFinite     := CheckStatus.pass,
  clockMonotonic       := CheckStatus.pass,
  nodeDetection        := CheckStatus.pass,
  nodeContactFree      := CheckStatus.pass,
  qFiniteAwayFromNodes := CheckStatus.pass,
  phaseGradientFinite  := CheckStatus.pass,
  classicalLimit       := CheckStatus.pass,
  friedmannResidual    := CheckStatus.deferred,
  faithfulness         := CheckStatus.contested
}

def sprint2NodeProbeWitness : BMCReport := {
  toyAnalysisOnly      := true,
  finalTruthClaim      := false,
  wdwResidual          := CheckStatus.pass,
  trajectoryFinite     := CheckStatus.fail,
  clockMonotonic       := CheckStatus.fail,
  nodeDetection        := CheckStatus.pass,
  nodeContactFree      := CheckStatus.fail,
  qFiniteAwayFromNodes := CheckStatus.contested,
  phaseGradientFinite  := CheckStatus.contested,
  classicalLimit       := CheckStatus.fail,
  friedmannResidual    := CheckStatus.deferred,
  faithfulness         := CheckStatus.contested
}

theorem sprint2_safe_witness_passes_superposition_safe :
  reportPassesBMC0ASuperpositionSafeGate sprint2SafeWitness = true := by decide

theorem sprint2_node_probe_witness_passes_node_detection :
  reportPassesBMC0ANodeDetectionGate sprint2NodeProbeWitness = true := by decide

theorem sprint2_node_probe_witness_fails_superposition_safe :
  reportPassesBMC0ASuperpositionSafeGate sprint2NodeProbeWitness = false := by decide
