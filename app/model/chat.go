package model

import "time"

// ChatResponse is a response for a chat.
type ChatResponse struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"`
}

// ChatParams is params for chat.
type ChatParams struct {
	FromDate time.Time
}

// MakeChatParams makes a chatparams.
func (cp *ChatParams) MakeChatParams() map[string]string {
	return map[string]string{"from_date": cp.FromDate.Format("2006-01-02")}
}
