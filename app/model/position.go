package model

import (
	"github.com/bmf-san/oilking/app/types"
)

// PositionResponse is a response for an position.
type PositionResponse struct {
	ProductCode         string  `json:"product_code"`
	Side                string  `json:"side"`
	Price               float64 `json:"price"`
	Size                float64 `json:"size"`
	Commission          float64 `json:"commission"`
	SwapPointAccumulate float64 `json:"swap_point_accumulate"`
	RequireCollateral   float64 `json:"require_collateral"`
	OpenDate            string  `json:"open_date"`
	Leverage            float64 `json:"leverage"`
	Pnl                 float64 `json:"pnl"`
	Sfd                 float64 `json:"sfd"`
}

// PositionParams is params for positions.
type PositionParams struct {
	ProductCode types.ProductCode
}

// MakePositionParams makes an positionsparams.
func (pp *PositionParams) MakePositionParams() map[string]string {
	return map[string]string{
		"product_code": string(pp.ProductCode),
	}
}
