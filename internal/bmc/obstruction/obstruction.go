package obstruction

import (
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// List of all 9 EBP 2.1 obstruction names required by the project specifications.
const (
	NodeObstruction                 = "node_obstruction"
	PhaseUnwrapObstruction          = "phase_unwrap_obstruction"
	NonfiniteQObstruction           = "nonfinite_q_obstruction"
	ClockNonmonotonicityObstruction = "clock_nonmonotonicity_obstruction"
	WdWResidualObstruction          = "wdw_residual_obstruction"
	ClassicalLimitFailure           = "classical_limit_failure"
	LapseOrTimeInterpretationDebt   = "lapse_or_time_interpretation_debt"
	MeasureProblemDeferred          = "measure_problem_deferred"
	FullQgOverclaimBlocker          = "full_qg_overclaim_blocker"
)

// NewDeferred creates an obstruction that represents an honestly deferred debt item.
func NewDeferred(name, reason string) model.Obstruction {
	return model.Obstruction{
		Name:        name,
		Applies:     false,
		Severity:    model.SeverityInfo,
		Evidence:    reason,
		Consequence: "Deferred to a future sprint.",
		Status:      model.StatusDeferred,
	}
}
