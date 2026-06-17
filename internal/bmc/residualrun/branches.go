package residualrun

// LocalBranchEligibility represents the eligibility diagnostics for a candidate local branch.
type LocalBranchEligibility struct {
	BranchID                  string `json:"branch_id"`
	SourceArtifact            string `json:"source_artifact"`
	Eligible                  bool   `json:"eligible"`
	EligibilityStatus         string `json:"eligibility_status"`
	Reason                    string `json:"reason"`
	NodeContactFree          bool   `json:"node_contact_free"`
	TrajectoryFinite          bool   `json:"trajectory_finite"`
	LocalClockStatus          string `json:"local_clock_status"`
	DerivativeReadinessStatus string `json:"derivative_readiness_status"`
}
