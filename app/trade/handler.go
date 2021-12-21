package trade

import (
	"github.com/bmf-san/oilking/app/bitflyer"
	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/types"
)

// Handler is a handler.
type Handler struct {
	apiClient *bitflyer.Client
	logger    *logger.Logger
}

func NewHandler(c *bitflyer.Client, l *logger.Logger) *Handler {
	return &Handler{
		apiClient: c,
		logger:    l,
	}
}

// SendChildOrder sends a child order.
func (h *Handler) SendChildOrder(side types.Side, price float64) {
	targetProductCode := types.ProductCodeFXBTCJPY
	targetSize := 0.01

	scor := &model.SendChildOrderRequest{
		ProductCode:     string(targetProductCode),
		ChildOrderType:  string(types.ChildOrderTypeLimit),
		Side:            string(side),
		Price:           price,
		Size:            targetSize,
		MinuteToExpires: 1,
		TimeInForce:     string(types.TimeInForceGTC),
	}
	_, err := h.apiClient.SendChildOrder(scor)
	if err != nil {
		h.logger.Error(logger.Entry{
			Message: err.Error(),
		})
		return
	}
	h.logger.Info(logger.TradingLogEntry{
		Label:  logger.LabelChildOrder,
		Action: "send a child order",
	})
}

// GetOrders gets orders.
func (h *Handler) GetOrders() []*model.ChildOrderResponse {
	op := &model.ChildOrderParams{
		ProductCode:     targetProductCode,
		Count:           1,
		ChildOrderState: types.ChildOrderStateActive,
	}
	or, err := h.apiClient.GetChildOrders(op)
	if err != nil {
		h.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}
	return or
}

// GetPositions gets positions.
func (h *Handler) GetPositions() []*model.PositionResponse {
	pp := &model.PositionParams{
		ProductCode: targetProductCode,
	}
	pr, err := h.apiClient.GetPositions(pp)
	if err != nil {
		h.logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}
	if len(pr) == 0 {
		return nil
	}
	return pr
}

// CancelAllChildOrders cancels all child orders.
func (h *Handler) CancelAllChildOrders() {
	cacor := &model.CancelAllChildOrdersRequest{
		ProductCode: string(targetProductCode),
	}
	err := h.apiClient.CancelAllChildOrders(cacor)
	if err != nil {
		h.logger.Error(logger.Entry{
			Message: err.Error(),
		})
		return
	}
	h.logger.Info(logger.TradingLogEntry{
		Label:  logger.LabelChildOrder,
		Action: "all orders have been cancelled",
	})
}
