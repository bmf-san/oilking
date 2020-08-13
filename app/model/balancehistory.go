package model

import (
	"strconv"

	"github.com/bmf-san/oilking/app/types"
)

// BalanceHistoryResponse is a response for  a balancehistory.
type BalanceHistoryResponse struct {
	ID           int     `json:"id"`
	TradeDate    string  `json:"trade_date"`
	EventDate    string  `json:"event_date"`
	ProductCode  string  `json:"product_code"`
	CurrencyCode string  `json:"currency_code"`
	TradeType    string  `json:"trade_type"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Quantity     float64 `json:"quantity"`
	Commission   float64 `json:"commission"`
	Balance      float64 `json:"balance"`
	OrderID      string  `json:"order_id"`
}

// BalanceHistoryParams is params for balancehistory.
type BalanceHistoryParams struct {
	CurrencyCode types.CurrencyCode
	Count        int
	Before       int
	After        int
}

// MakeBalanceHistoryParams makes a balancehistoryparams.
func (bhp *BalanceHistoryParams) MakeBalanceHistoryParams() map[string]string {
	return map[string]string{
		"currency_code": string(bhp.CurrencyCode),
		"count":         strconv.Itoa(bhp.Count),
		"before":        strconv.Itoa(bhp.Before),
		"after":         strconv.Itoa(bhp.After),
	}
}
