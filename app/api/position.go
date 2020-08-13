package api

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getPositionsPath is a path for getpositions endpoint.
	getPositionsPath = "/me/getpositions"
)

// GetPositions gets positions.
func (c *Client) GetPositions(positionParams *model.PositionParams) ([]*model.PositionResponse, error) {
	body, err := c.Do(http.MethodGet, getPositionsPath, positionParams.MakePositionParams(), nil)
	if err != nil {
		return nil, err
	}

	var pr []*model.PositionResponse
	if err = json.Unmarshal(body, &pr); err != nil {
		return nil, err
	}

	return pr, nil
}
