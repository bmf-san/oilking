package bitflyer

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getChildOrdersPath is a path for getchildorders endpoint.
	getChildOrdersPath = "/me/getchildorders"
	// cancelAllChildOrdersPath is a path for cancelallchildorders endpoint.
	cancelAllChildOrdersPath = "/me/cancelallchildorders"
	// cancelchildorder is a path for cancelchildorder endpoint.
	cancelChildOrderPath = "/me/cancelchildorder"
	// sendChildOrderPath is a path for sendchildorder endpoint.
	sendChildOrderPath = "/me/sendchildorder"
)

// GetChildOrders gets child orders.
func (c *Client) GetChildOrders(childOrderParams *model.ChildOrderParams) ([]*model.ChildOrderResponse, error) {
	body, err := c.Do(http.MethodGet, getChildOrdersPath, childOrderParams.MakeChildOrderParams(), nil)
	if err != nil {
		return nil, err
	}

	var o []*model.ChildOrderResponse
	if err = json.Unmarshal(body, &o); err != nil {
		return nil, err
	}

	return o, nil
}

// CancelAllChildOrders cancels all child orders.
func (c *Client) CancelAllChildOrders(cancelAllChildOrdersRequest *model.CancelAllChildOrdersRequest) error {
	data, err := json.Marshal(cancelAllChildOrdersRequest)
	if err != nil {
		return err
	}

	_, err = c.Do(http.MethodPost, cancelAllChildOrdersPath, nil, data)
	if err != nil {
		return err
	}

	return nil
}

// CancelChildOrder cancels a child order.
func (c *Client) CancelChildOrder(cancelChildOrderRequest *model.CancelChildOrderRequest) error {
	data, err := json.Marshal(cancelChildOrderRequest)
	if err != nil {
		return err
	}

	_, err = c.Do(http.MethodPost, cancelChildOrderPath, nil, data)
	if err != nil {
		return err
	}

	return nil
}

// SendChildOrder sends a child order.
func (c *Client) SendChildOrder(sendChildOrderRequest *model.SendChildOrderRequest) (*model.SendChildChildOrderResponse, error) {
	data, err := json.Marshal(sendChildOrderRequest)
	if err != nil {
		return nil, err
	}

	body, err := c.Do(http.MethodPost, sendChildOrderPath, nil, data)
	if err != nil {
		return nil, err
	}

	var sr model.SendChildChildOrderResponse
	if err = json.Unmarshal(body, &sr); err != nil {
		return nil, err
	}

	return &sr, nil
}
