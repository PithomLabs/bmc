package nullspec

// Allowed status values for null models, metric contracts, and comparison contracts
const (
	StatusPlanned  = "planned"
	StatusDeferred = "deferred"
	StatusBlocked  = "blocked"
)

// Allowed input availability status values
const (
	AvailabilityAvailable = "available"
	AvailabilityPlanned   = "planned"
	AvailabilityDeferred  = "deferred"
	AvailabilityBlocked   = "blocked"
)

// Allowed spec scope values
const (
	SpecScopeNullModelOnly = "null_model_specification_only"
)
