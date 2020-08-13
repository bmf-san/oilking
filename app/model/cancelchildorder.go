package model

// CancelChildOrderRequest is a request for cancelchildorder
type CancelChildOrderRequest struct {
	ProductCode  string `json:"product_code"`
	ChildOrderID string `json:"child_order_id"`
}
