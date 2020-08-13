package model

// BalanceResponse is a response for a balance.
type BalanceResponse struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
	Available    float64 `json:"available"`
}
