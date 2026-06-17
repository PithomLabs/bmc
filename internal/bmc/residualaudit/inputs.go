package residualaudit

type SourceArtifactRef struct {
	ArtifactID   string `json:"artifact_id"`
	ArtifactKind string `json:"artifact_kind"`
	Provenance   string `json:"provenance"`
	Path         string `json:"path,omitempty"`
	Status       string `json:"status"`
	Notes        string `json:"notes"`
}

func defaultSourceArtifacts() []SourceArtifactRef {
	return []SourceArtifactRef{
		{ArtifactID: "bmc0a_clock_readiness", ArtifactKind: "clock_readiness_report", Provenance: ProvenanceNotAvailable, Status: "not_available", Notes: "Clock readiness report not loaded."},
		{ArtifactID: "bmc0a_nullrun", ArtifactKind: "nullrun_report", Provenance: ProvenanceNotAvailable, Status: "not_available", Notes: "Null model run report not loaded."},
		{ArtifactID: "bmc0a_local_residual", ArtifactKind: "local_residual_report", Provenance: ProvenanceNotAvailable, Status: "not_available", Notes: "Local residual report not loaded."},
		{ArtifactID: "bmc0a_friedmann_spec", ArtifactKind: "friedmann_spec_report", Provenance: ProvenanceNotAvailable, Status: "not_available", Notes: "Friedmann spec report not loaded."},
		{ArtifactID: "bmc0a_prior_art_boundary", ArtifactKind: "prior_art_boundary_report", Provenance: ProvenanceNotAvailable, Status: "not_available", Notes: "Prior art boundary report not loaded."},
	}
}

func markSourceRead(sources []SourceArtifactRef, artifactID, path string) {
	for idx := range sources {
		if sources[idx].ArtifactID == artifactID {
			sources[idx].Provenance = ProvenanceFileRead
			sources[idx].Path = path
			sources[idx].Status = "available"
			sources[idx].Notes = "Successfully read source artifact."
			return
		}
	}
}

func markSourceUnavailable(sources []SourceArtifactRef, artifactID, path string, err error) {
	for idx := range sources {
		if sources[idx].ArtifactID == artifactID {
			sources[idx].Provenance = ProvenanceNotAvailable
			sources[idx].Path = path
			sources[idx].Status = "not_available"
			sources[idx].Notes = "Source artifact unavailable."
			if err != nil {
				sources[idx].Notes = "Source artifact unavailable: " + err.Error()
			}
			return
		}
	}
}
