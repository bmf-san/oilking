package model

// JSONRPC2 is a jsonrpc2
type JSONRPC2 struct {
	Version string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Result  interface{} `json:"result,omitempty"`
	ID      *int        `json:"id,omitempty"`
}

// SubscribeParams is a subscribeparams.
type SubscribeParams struct {
	Channel string `json:"channel"`
}
