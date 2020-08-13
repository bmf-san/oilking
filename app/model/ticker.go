package model

import (
	"github.com/bmf-san/oilking/app/types"
)

// TickerResponse is a response for ticker.
type TickerResponse struct {
	ProductCode     string  `json:"product_code"`
	Timestamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBid         float64 `json:"best_bid"`
	BestAsk         float64 `json:"best_ask"`
	BestBidSize     float64 `json:"best_bid_size"`
	BestAskSize     float64 `json:"best_ask_size"`
	TotalBidDepth   float64 `json:"total_bid_depth"`
	TotalAskDepth   float64 `json:"total_ask_depth"`
	MarketBidSize   float64 `json:"market_bid_size"`
	MarketAskSize   float64 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float64 `json:"volume"`
	VolumeByProduct float64 `json:"volume_by_product"`
}

// GetMidPrice gets mid price.
func (t *TickerResponse) GetMidPrice() float64 {
	return (t.BestBid + t.BestAsk) / 2
}

// TickerParams is params for ticker.
type TickerParams struct {
	ProductCode types.ProductCode
}

// MakeTickerParams makes a tickerparams.
func (tp *TickerParams) MakeTickerParams() map[string]string {
	return map[string]string{"product_code": string(tp.ProductCode)}
}
