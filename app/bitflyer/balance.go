package bitflyer

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// balancePath is a path for balance endpoint.
	balancePath = "/me/getbalance"
)

// GetBalance gets balance.
func (c *Client) GetBalance() ([]*model.BalanceResponse, error) {
	body, err := c.Do(http.MethodGet, balancePath, nil, nil)
	if err != nil {
		return nil, err
	}

	var b []*model.BalanceResponse
	if err = json.Unmarshal(body, &b); err != nil {
		return nil, err
	}

	return b, nil
}
