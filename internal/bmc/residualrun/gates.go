package residualrun

// ResidualRunGate represents a policy boundary check.
type ResidualRunGate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}
