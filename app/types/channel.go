package types

// Channel is a type for channel.
type Channel string

const (
	// ChannelTickerBTCJPY is a channel for lightninng_ticker_BTC_JPY.
	ChannelTickerBTCJPY Channel = "lightning_ticker_BTC_JPY"
	// ChannelTickerFXBTCJPY is a channel for lightninng_ticker_FXx_BTC_JPY.
	ChannelTickerFXBTCJPY Channel = "lightning_ticker_FX_BTC_JPY"
)
