package model

import (
	"github.com/bmf-san/oilking/app/types"
)

// Writer is a writer for websocket.
type Writer struct {
	Channel       types.Channel
	Method        types.ServerMethod
	ChannelTicker chan TickerResponse
}
