package types

// ChildOrderType is a type for childordertype.
type ChildOrderType string

const (
	// ChildOrderTypeLimit is a childordertype for LIMIT.
	ChildOrderTypeLimit ChildOrderType = "LIMIT"
	// ChildOrderTypeMarket is a childordertype for MARKET.
	ChildOrderTypeMarket ChildOrderType = "MARKET"
)
