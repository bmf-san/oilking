package model

// SendChildChildOrderResponse is a response forr sendchildorder.
type SendChildChildOrderResponse struct {
	ChildOrderAcceptanceID string `json:"child_order_acceptance_id"`
}

// SendChildOrderRequest is a request for sendchildorder.
type SendChildOrderRequest struct {
	ProductCode     string  `json:"product_code"`
	ChildOrderType  string  `json:"child_order_type"`
	Side            string  `json:"side"`
	Price           float64 `json:"price"`
	Size            float64 `json:"size"`
	MinuteToExpires int     `json:"minute_to_expire"`
	TimeInForce     string  `json:"time_in_force"`
}
