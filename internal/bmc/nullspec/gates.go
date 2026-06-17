package nullspec

// NullSpecGate represents an EBP policy safety gate for Sprint 7.
type NullSpecGate struct {
	Name   string `json:"name"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}
