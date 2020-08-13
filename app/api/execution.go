package api

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getExecutionsPath is a path for getexecutions endpoint.
	getExecutionsPath = "/me/getexecutions"
)

// GetExecutions gets executions.
func (c *Client) GetExecutions(executionParams *model.ExecutionParams) (*model.ExecutionResponse, error) {
	body, err := c.Do(http.MethodGet, getExecutionsPath, executionParams.MakeExecutionParams(), nil)
	if err != nil {
		return nil, err
	}

	var e model.ExecutionResponse
	if err = json.Unmarshal(body, &e); err != nil {
		return nil, err
	}

	return &e, nil
}
