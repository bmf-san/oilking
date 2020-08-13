package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bmf-san/oilking/app/api"
	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/trade"
	"github.com/bmf-san/oilking/app/types"
	"github.com/go-redis/redis/v7"
)

func main() {
	l := logger.NewLogger(logger.LevelInfo, time.FixedZone("Asia/Tokyo", 9*60*60))
	ak := os.Getenv("BITFLYER_API_KEY")
	as := os.Getenv("BITFLYER_API_SECRET")
	c := api.NewClient(ak, as, l)
	// TODO: I haven't planned to use it yet, so I commented it out.
	// db := database.NewDB()
	// d, err := db.Conn()
	// if err != nil {
	// 	l.Error(logger.Entry{
	// 		Message: err.Error(),
	// 	})
	// 	os.Exit(1)
	// }
	// defer d.Close()
	t := trade.NewTrade(c, l, &redis.Client{})
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
