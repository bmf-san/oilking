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
	"github.com/go-redis/redis/v7"
)

const (
	// NOTE: These values ​​can be parameterized later.
	// targetProductCode is a types for trading.
	targetProductCode = types.ProductCodeFXBTCJPY
	// targetSize is a size for buy or sell.
	targetSize = 0.01

	// TODO: Will be defined later in the logger.
	// NOTE: These values user for logging. It ​​may be used when implementing reporting services.
	// NOTE: It seems better to prepare the action name as well. ex. [ACTION]
	// labelPosition is a label for position.
	labelPosition = "[Position]"
	// labelOrder is a label for order.
	labelOrder = "[Order]"
	// labelCollateral is a label for collateral.
	labelCollateral = "[Collateral]"
)

// Trade is a trade.
type Trade struct {
	apiClient *api.Client
	logger    *logger.Logger
	dbClient  *redis.Client
}

// NewTrade creates a new trade.
func NewTrade(c *api.Client, l *logger.Logger, db *redis.Client) *Trade {
	return &Trade{
		apiClient: c,
		logger:    l,
		dbClient:  db,
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
		t.logger.Info(logger.Entry{
			Message: fmt.Sprintf("%s send a child order", labelOrder),
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
		t.logger.Info(logger.Entry{
			Message: fmt.Sprintf("%s send a child order", labelOrder),
		})
	}
}

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
		t.logger.Info(logger.Entry{
			Message: fmt.Sprintf("%s All orders have been cancelled.", labelOrder),
		})
	}

	// I want to close the position by counter-trading, but for the time being I will make an error instead of an alert.
	pp := &model.PositionParams{
		ProductCode: targetProductCode,
	}
	pr, err := t.apiClient.GetPositions(pp)
	if err != nil {
		t.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	} else {
		t.logger.Info(logger.Entry{
			Message: fmt.Sprintf("%s %d positions", labelPosition, len(pr)),
		})
	}
}
