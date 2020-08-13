package types

// ChildOrderState is a type for childorderstate.
type ChildOrderState string

const (
	// ChildOrderStateActive is a childorderstate for ACTIVE.
	ChildOrderStateActive ChildOrderState = "ACTIVE"
	// ChildOrderStateCompleted is a childorderstate for COMPLETED.
	ChildOrderStateCompleted ChildOrderState = "COMPLETED"
	// ChildOrderStateCanceled is a childorderstate for CANCELED.
	ChildOrderStateCanceled ChildOrderState = "CANCELED"
	// ChildOrderStateExpired is a childorderstate for EXPIRED.
	ChildOrderStateExpired ChildOrderState = "EXPIRED"
	// ChildOrderStateRejected is a childorderstate for REJECTED.
	ChildOrderStateRejected ChildOrderState = "REJECTED"
)
