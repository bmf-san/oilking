package trade

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/bmf-san/oilking/app/api"
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

// Trade is a trade.
type Trade struct {
	apiClient *api.Client
	logger    *logger.Logger
}

// NewTrade creates a new trade.
func NewTrade(c *api.Client, l *logger.Logger) *Trade {
	return &Trade{
		apiClient: c,
		logger:    l,
	}
}

// Run runs actions for trading.
// Rule:
// - Always has only one order.
// - Always has only one position.
// - Discard all orders at the end.
// - Handle both buy and sell orders and positions.
func (t *Trade) Run(tr *model.TickerResponse) {
	time.Sleep(5 * time.Second)

	if t.hasOrders() {
		return
	}

	if t.hasPositions() {
		p := t.getPosition()
		switch p.Side {
		case string(types.SideBuy):
			t.sell(math.Trunc(p.Price * 1.001))
		case string(types.SideSell):
			t.buy(math.Trunc(p.Price * 0.999))
		default:
			t.logger.Error(logger.Entry{
				Message: errors.New("Something was wrong").Error(),
			})
		}
		return
	}

	if tr.BestBidSize <= tr.BestAskSize {
		t.sell(math.Trunc(tr.BestBid))
		return
	}

	t.buy(math.Trunc(tr.BestAsk))
	return
}

// buy gets a buy position.
func (t *Trade) buy(p float64) {
	scor := &model.SendChildOrderRequest{
		ProductCode:     string(targetProductCode),
		ChildOrderType:  string(types.ChildOrderTypeLimit),
		Side:            string(types.SideBuy),
		Price:           p,
		Size:            targetSize,
		MinuteToExpires: 1,
		TimeInForce:     string(types.TimeInForceGTC),
	}
	_, err := t.apiClient.SendChildOrder(scor)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	} else {
		t.logger.Info(logger.TradingLogEntry{
			Label:  logger.LabelChildOrder,
			Action: "send a child order",
		})
	}
}

// sell gets a sell position.
func (t *Trade) sell(p float64) {
	scor := &model.SendChildOrderRequest{
		ProductCode:     string(targetProductCode),
		ChildOrderType:  string(types.ChildOrderTypeLimit),
		Side:            string(types.SideSell),
		Price:           p,
		Size:            targetSize,
		MinuteToExpires: 1,
		TimeInForce:     string(types.TimeInForceGTC),
	}
	_, err := t.apiClient.SendChildOrder(scor)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	} else {
		t.logger.Info(logger.TradingLogEntry{
			Label:  logger.LabelChildOrder,
			Action: "send a child order",
		})
	}
}

// hasPositions determines if there are positions.
func (t *Trade) hasPositions() bool {
	pp := &model.PositionParams{
		ProductCode: targetProductCode,
	}
	pr, err := t.apiClient.GetPositions(pp)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}

	return len(pr) > 0
}

// hasOrders determines if there are orders.
func (t *Trade) hasOrders() bool {
	op := &model.ChildOrderParams{
		ProductCode:     targetProductCode,
		Count:           1,
		ChildOrderState: types.ChildOrderStateActive,
	}
	or, err := t.apiClient.GetChildOrders(op)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}

	return len(or) > 0
}

// getPosition gets a position.
func (t *Trade) getPosition() *model.PositionResponse {
	pp := &model.PositionParams{
		ProductCode: targetProductCode,
	}
	pr, err := t.apiClient.GetPositions(pp)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}
	if len(pr) == 0 {
		return nil
	}
	var p *model.PositionResponse
	for _, v := range pr {
		p = v
	}
	return p
}

// StopSafety runs actions for StopSafety at the end of runnning.
func (t *Trade) StopSafety() {
	// Leave no orders when the bot is out of service.
	cacor := &model.CancelAllChildOrdersRequest{
		ProductCode: string(targetProductCode),
	}
	err := t.apiClient.CancelAllChildOrders(cacor)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	} else {
		t.logger.Info(logger.TradingLogEntry{
			Label:  logger.LabelChildOrder,
			Action: "all orders have been cancelled",
		})
	}

	// Check the remaining open interest
	pp := &model.PositionParams{
		ProductCode: targetProductCode,
	}
	pr, err := t.apiClient.GetPositions(pp)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	} else {
		msg := "no positions"
		if len(pr) > 0 {
			msg = fmt.Sprintf("get %d positions", len(pr))
		}
		t.logger.Info(logger.TradingLogEntry{
			Label:  logger.LabelPosition,
			Action: msg,
		})
	}
}
