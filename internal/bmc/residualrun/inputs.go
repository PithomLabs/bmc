package residualrun

// SourceArtifactRef represents a detailed provenance log for each required input artifact.
type SourceArtifactRef struct {
	ArtifactID   string `json:"artifact_id"`
	ArtifactKind string `json:"artifact_kind"`
	Provenance   string `json:"provenance"` // file_read, source_artifact_summary, not_available
	Path         string `json:"path,omitempty"`
	Status       string `json:"status"`
	Notes        string `json:"notes"`
}
