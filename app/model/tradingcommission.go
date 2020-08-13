package model

import (
	"github.com/bmf-san/oilking/app/types"
)

// TradingCommissionResponse is a response for tradingcommission.
type TradingCommissionResponse struct {
	CommissionRate float64 `json:"commission_rate"`
}

// TradingCommissionParams is params for ticker.
type TradingCommissionParams struct {
	ProductCode types.ProductCode
}

// MakeTradingCommissionParams makes a tradingcommissionparams.
func (tp *TradingCommissionParams) MakeTradingCommissionParams() map[string]string {
	return map[string]string{"product_code": string(tp.ProductCode)}
}
