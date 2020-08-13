package api

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getCollateralPath is a path for collateral endpoint.
	getCollateralPath = "/me/getcollateral"
)

// GetCollateral gets collateral.
func (c *Client) GetCollateral() (*model.CollateralResponse, error) {
	body, err := c.Do(http.MethodGet, getCollateralPath, nil, nil)
	if err != nil {
		return nil, err
	}

	var co model.CollateralResponse
	if err = json.Unmarshal(body, &co); err != nil {
		return nil, err
	}

	return &co, nil
}
