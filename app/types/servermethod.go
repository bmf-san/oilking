package types

// ServerMethod is a type for servermethod.
type ServerMethod string

const (
	// ServerMethodAuth is a servermethod for auth.
	ServerMethodAuth ServerMethod = "auth"
	// ServerMethodSubscribe is a severmethod for subscribe.
	ServerMethodSubscribe ServerMethod = "subscribe"
	// ServerMethodUnsubscribe is a servermethod for unsubscribe.
	ServerMethodUnsubscribe ServerMethod = "unsubscribe"
)
