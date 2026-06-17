inductive CheckStatus
| pass
| fail
| deferred
| contested
deriving DecidableEq, Repr

def checkPassed : CheckStatus -> Bool
| CheckStatus.pass => true
| _ => false

def checkDeferred : CheckStatus -> Bool
| CheckStatus.deferred => true
| _ => false

structure BMCReport where
  toyAnalysisOnly      : Bool
  finalTruthClaim      : Bool
  wdwResidual          : CheckStatus
  trajectoryFinite     : CheckStatus
  clockMonotonic       : CheckStatus
  nodeDetection        : CheckStatus
  nodeContactFree      : CheckStatus
  qFiniteAwayFromNodes : CheckStatus
  phaseGradientFinite  : CheckStatus
  classicalLimit       : CheckStatus
  friedmannResidual    : CheckStatus
  faithfulness         : CheckStatus
  deriving Repr
