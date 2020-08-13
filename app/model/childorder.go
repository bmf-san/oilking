package model

import (
	"strconv"

	"github.com/bmf-san/oilking/app/types"
)

// ChildOrderResponse is a response for a child order.
type ChildOrderResponse struct {
	ID                     int     `json:"id"`
	ChildOrderID           string  `json:"child_order_id"`
	ProductCode            string  `json:"product_code"`
	Side                   string  `json:"side"`
	ChildOrderType         string  `json:"child_order_type"`
	Price                  float64 `json:"price"`
	AveragePrice           float64 `json:"average_price"`
	Size                   float64 `json:"size"`
	ChildOrderState        string  `json:"child_order_state"`
	ExpireDate             string  `json:"expire_date"`
	ChildOrderDate         string  `json:"child_order_date"`
	ChildOrderAcceptanceID string  `json:"child_order_acceptance_id"`
	OutstandingSize        float64 `json:"outstanding_size"`
	CancelSize             float64 `json:"cancel_size"`
	ExecutedSize           float64 `json:"executed_size"`
	TotalCommission        float64 `json:"total_commission"`
}

// ChildOrderParams is params for child order.
type ChildOrderParams struct {
	ProductCode            types.ProductCode
	Count                  int
	Before                 int
	After                  int
	ChildOrderState        types.ChildOrderState
	ChildOrderID           string
	ChildOrderAcceptanceID string
	ParentOrderID          string
}

// MakeChildOrderParams makes a childorderparams.
func (op *ChildOrderParams) MakeChildOrderParams() map[string]string {
	return map[string]string{
		"product_code":              string(op.ProductCode),
		"count":                     strconv.Itoa(op.Count),
		"before":                    strconv.Itoa(op.Before),
		"after":                     strconv.Itoa(op.After),
		"child_order_state":         string(op.ChildOrderState),
		"child_order_id":            op.ChildOrderID,
		"child_order_acceptance_id": op.ChildOrderAcceptanceID,
		"parent_order_id":           op.ParentOrderID,
	}
}
