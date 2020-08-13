package model

import (
	"strconv"

	"github.com/bmf-san/oilking/app/types"
)

// ExecutionResponse is a response for an execution.
type ExecutionResponse struct {
	ID                     int     `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	Side                   string  `json:"side"`
	Price                  float64 `json:"price"`
	Size                   float64 `json:"size"`
	Commission             float64 `json:"commission"`
	ExecDate               string  `json:"exec_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
}

// ExecutionParams is params for executions.
type ExecutionParams struct {
	ProductCode            types.ProductCode
	Count                  int
	Before                 int
	After                  int
	ChildOrderID           string
	ChildOrderAcceptanceID string
}

// MakeExecutionParams makes an executionsparams.
func (ep *ExecutionParams) MakeExecutionParams() map[string]string {
	return map[string]string{
		"product_code":              string(ep.ProductCode),
		"count":                     strconv.Itoa(ep.Count),
		"before":                    strconv.Itoa(ep.Before),
		"after":                     strconv.Itoa(ep.After),
		"child_order_id":            ep.ChildOrderID,
		"child_order_acceptance_id": ep.ChildOrderAcceptanceID,
	}
}
