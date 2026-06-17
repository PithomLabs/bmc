package friedmannspec

// FriedmannSpecGate represents a policy or check gate to enforce Sprint 6 EBP constraints.
type FriedmannSpecGate struct {
	Name   string `json:"name"`
	Status string `json:"status"` // pass, blocked, contested
	Reason string `json:"reason"`
}
