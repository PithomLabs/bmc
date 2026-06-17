package residualrun

// CandidateResidualMetrics carries candidate local-branch residual metrics.
type CandidateResidualMetrics struct {
	NumEvaluationPoints     int      `json:"num_evaluation_points"`
	NumFiniteResidualPoints int      `json:"num_finite_residual_points"`
	MeanAbsResidual         *float64 `json:"mean_abs_residual"`
	MaxAbsResidual          *float64 `json:"max_abs_residual"`
	RmsResidual             *float64 `json:"rms_residual"`
	ResidualFinite          bool     `json:"residual_finite"`
	DiagnosticWarnings      []string `json:"diagnostic_warnings"`
}

type ResidualInputPoint struct {
	BranchID               string   `json:"branch_id"`
	PointIndex             int      `json:"point_index"`
	Lambda                 *float64 `json:"lambda"`
	Alpha                  *float64 `json:"alpha"`
	Phi                    *float64 `json:"phi"`
	CandidateLeftHandSide  *float64 `json:"candidate_left_hand_side"`
	CandidateRightHandSide *float64 `json:"candidate_right_hand_side"`
	InputProvenance        string   `json:"input_provenance"`
}

// CandidateResidualDiagnostic evaluates the candidate residual status and metrics on a local branch.
type CandidateResidualDiagnostic struct {
	BranchID            string                   `json:"branch_id"`
	ResidualID          string                   `json:"residual_id"`
	ResidualComputed    bool                     `json:"residual_computed"`
	ResidualStatus      string                   `json:"residual_status"`
	ResidualProvenance  string                   `json:"residual_provenance"`
	Metrics             CandidateResidualMetrics `json:"metrics"`
	BlockedReason       string                   `json:"blocked_reason,omitempty"`
	Notes               string                   `json:"notes"`
	ResidualInputPoints []ResidualInputPoint     `json:"residual_input_points"`
}
