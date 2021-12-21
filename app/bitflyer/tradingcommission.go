package bitflyer

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// tradingCommissionPath is a path for tradingcommission endpoing.
	tradingCommissionPath = "/me/gettradingcommission"
)

// GetTradingCommission gets tradingcommission.
func (c *Client) GetTradingCommission(tradingCommissionParams *model.TradingCommissionParams) (*model.TradingCommissionResponse, error) {
	body, err := c.Do(http.MethodGet, tradingCommissionPath, tradingCommissionParams.MakeTradingCommissionParams(), nil)
	if err != nil {
		return nil, err
	}

	var tcr model.TradingCommissionResponse
	if err = json.Unmarshal(body, &tcr); err != nil {
		return nil, err
	}

	return &tcr, nil
}
