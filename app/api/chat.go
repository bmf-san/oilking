package api

import (
	"encoding/json"
	"net/http"

	"github.com/bmf-san/oilking/app/model"
)

const (
	// getChatsPath is a path for balance endpoint.
	getChatsPath = "/getchats"
)

// GetChat gets chat.
func (c *Client) GetChat(chatParams *model.ChatParams) ([]*model.ChatResponse, error) {
	body, err := c.Do(http.MethodGet, getChatsPath, chatParams.MakeChatParams(), nil)
	if err != nil {
		return nil, err
	}

	var ch []*model.ChatResponse
	if err = json.Unmarshal(body, &ch); err != nil {
		return nil, err
	}

	return ch, nil
}
