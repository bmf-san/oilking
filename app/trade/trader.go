package trade

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/bmf-san/oilking/app/bitflyer"
	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/types"
)

const (
	// targetProductCode is a types for trading.
	targetProductCode = types.ProductCodeFXBTCJPY
	// targetSize is a size for buy or sell.
	targetSize = 0.01
)

// Trader is a trader.
type Trader struct {
	Handler *Handler
	Logger  *logger.Logger
}

// NewTrader creates a new trader.
func NewTrader(c *bitflyer.Client, l *logger.Logger) *Trader {
	return &Trader{
		Handler: NewHandler(c, l),
		Logger:  l,
	}
}

// Run runs trading..
func (t *Trader) Run(tr *model.TickerResponse) {
	time.Sleep(1 * time.Second)

	if len(t.Handler.GetOrders()) > 0 {
		t.Handler.CancelAllChildOrders()
		return
	}

	psn := t.Handler.GetPositions()
	if len(psn) > 0 {
		p := psn[0]
		switch p.Side {
		case string(types.SideBuy):
			t.Handler.SendChildOrder(types.SideBuy, math.Trunc(p.Price*1.0002))
		case string(types.SideSell):
			t.Handler.SendChildOrder(types.SideSell, math.Trunc(p.Price*0.9998))
		default:
			t.Logger.Error(logger.Entry{
				Message: errors.New("Something was wrong").Error(),
			})
		}
		return
	}
	if tr.BestBidSize <= tr.BestAskSize {
		t.Handler.SendChildOrder(types.SideSell, math.Trunc(tr.BestBid))
		return
	}
	t.Handler.SendChildOrder(types.SideBuy, math.Trunc(tr.BestAsk))
}

// StopSafety runs actions for StopSafety at the end of runnning.
func (t *Trader) StopSafety() {
	t.Handler.CancelAllChildOrders()
	pr := t.Handler.GetPositions()
	msg := "no positions"
	if len(pr) > 0 {
		msg = fmt.Sprintf("get %d positions", len(pr))
	}
	t.Logger.Info(logger.TradingLogEntry{
		Label:  logger.LabelPosition,
		Action: msg,
	})
}
