package model

// CollateralResponse is a response for collateral.
type CollateralResponse struct {
	Collateral        float64 `json:"collateral"`
	OpenPositionPnl   float64 `json:"open_position_pnl"`
	RequireCollateral float64 `json:"require_collateral"`
	KeepRate          float64 `json:"keep_rate"`
}
