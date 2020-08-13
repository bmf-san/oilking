package types

// TimeInForce is a type for timeinforce.
type TimeInForce string

const (
	// TimeInForceGTC is a timeinforce for GTC.
	TimeInForceGTC TimeInForce = "GTC"
	// TimeInForceIOC is a timeinforce for IOC.
	TimeInForceIOC TimeInForce = "IOC"
	// TimeInForceFOK is a timeinforce for FOK.
	TimeInForceFOK TimeInForce = "FOK"
)
