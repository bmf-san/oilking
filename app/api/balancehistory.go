package api

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getBalanceHistory is a path for getbalancehistory endpoint.
	getBalanceHistory = "/me/getbalancehistory"
)

// GetBalanceHistory gets balancehistory.
func (c *Client) GetBalanceHistory(balanceHistoryParams *model.BalanceHistoryParams) (*model.BalanceHistoryResponse, error) {
	body, err := c.Do(http.MethodGet, getBalanceHistory, balanceHistoryParams.MakeBalanceHistoryParams(), nil)
	if err != nil {
		return nil, err
	}

	var bhr model.BalanceHistoryResponse
	if err = json.Unmarshal(body, &bhr); err != nil {
		return nil, err
	}

	return &bhr, nil
}
