package bitflyer

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getTickerPath is a path for ticker endpoint.
	getTickerPath = "/getticker"
)

// GetTicker gets ticker.
func (c *Client) GetTicker(tickerParams *model.TickerParams) (*model.TickerResponse, error) {
	body, err := c.Do(http.MethodGet, getTickerPath, tickerParams.MakeTickerParams(), nil)
	if err != nil {
		return nil, err
	}

	var t model.TickerResponse
	if err = json.Unmarshal(body, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
