package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bmf-san/oilking/app/bitflyer"
	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/trade"
	"github.com/bmf-san/oilking/app/types"
)

func main() {
	l := logger.NewLogger(logger.LevelInfo, time.FixedZone("Asia/Tokyo", 9*60*60))
	ak := os.Getenv("BITFLYER_API_KEY")
	as := os.Getenv("BITFLYER_API_SECRET")
	c := bitflyer.NewClient(ak, as, l)
	t := trade.NewTrader(c, l)
	ch := make(chan model.TickerResponse)
	w := &model.Writer{
		Channel:       types.ChannelTickerFXBTCJPY,
		Method:        types.ServerMethodSubscribe,
		ChannelTicker: ch,
	}

	l.Info(logger.Entry{
		Message: "[START] Oilking starts running...",
	})

	go func() {
		go c.DoWs(w)
		for c := range ch {
			t.Run(&c)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-sig

	t.StopSafety()

	l.Info(logger.Entry{
		Message: "[END] Oilking stops running",
	})
}
