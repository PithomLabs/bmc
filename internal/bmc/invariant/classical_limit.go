package invariant

import (
	"github.com/PithomLabs/bmc/internal/bmc/model"
)

// CheckClassicalLimit validates if the quantum potential is small enough to indicate classical-limit recovery.
func CheckClassicalLimit(maxAbsQ float64, tolerance float64) model.CheckResult {
	pass := maxAbsQ <= tolerance

	var status model.CheckStatus
	var reason string
	if pass {
		status = model.StatusPass
		reason = "Plane-wave control has Q approximately zero."
	} else {
		status = model.StatusFail
		reason = "Classical limit check failed: quantum potential is non-negligible."
	}

	return model.CheckResult{
		Status: status,
		Pass:   pass,
		Reason: reason,
	}
}
