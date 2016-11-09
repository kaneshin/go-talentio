package talentio

// Common statuses of candidate.
const (
	StatusOngoing      = "ongoing"
	StatusReject       = "reject"
	StatusFail         = "fail"
	StatusPass         = "pass"
	StatusPoolActive   = "poolActive"
	StatusPoolInactive = "poolInactive"
)

// Common types of the state between candidate and recruiter.
const (
	TypeContact   = "contact"
	TypeResume    = "resume"
	TypeInterview = "interview"
)

// Sorting keys.
const (
	SortRegisteredAtAscKey  = "registeredAt"
	SortRegisteredAtDescKey = "-registeredAt"
)
