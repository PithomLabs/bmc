package priorart

// Allowed SourceKind values
const (
	SourceKindReview  = "review"
	SourceKindPaper   = "paper"
	SourceKindBook    = "book"
	SourceKindUnknown = "unknown"
)

// Allowed ReviewStatus values
const (
	ReviewStatusSeedUnreviewed      = "seed_unreviewed"
	ReviewStatusAbstractReviewed    = "abstract_reviewed"
	ReviewStatusSkimReviewed        = "skim_reviewed"
	ReviewStatusHumanReviewRequired = "human_review_required"
)

// Allowed BoundaryStatus values
const (
	BoundaryStatusEstablishedPriorArt          = "established_prior_art"
	BoundaryStatusLikelyPriorArt               = "likely_prior_art"
	BoundaryStatusImplementationVariant        = "implementation_variant"
	BoundaryStatusWorkflowDistinctiveCandidate = "workflow_distinctive_candidate"
	BoundaryStatusUnknownRequiresReview        = "unknown_requires_review"
	BoundaryStatusBlocked                      = "blocked"
	BoundaryStatusNotClaimed                   = "not_claimed"
)
