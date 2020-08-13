package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bmf-san/oilking/app/api"
	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/types"
)

// Example codes for api client.
func main() {
	// TODO: Replace with test code later.
	l := logger.NewLogger(logger.LevelInfo, time.FixedZone("Asia/Tokyo", 9*60*60))
	f, err := l.SetOutput()
	if err != nil {
		l.Error(logger.Entry{
			Message: err.Error(),
		})
		os.Exit(1)
	}
	defer f.Close()
	ak := os.Getenv("BITFLYER_API_KEY")
	as := os.Getenv("BITFLYER_API_SECRET")
	c := api.NewClient(ak, as, l)

	// ticker
	tp := &model.TickerParams{
		ProductCode: types.ProductCodeBTCJPY,
	}
	tresp, err := c.GetTicker(tp)
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", tresp)

	// getbalance
	bresp, err := c.GetBalance()
	if err != nil {
		l.Error(err.Error())
	}
	for k, v := range bresp {
		fmt.Printf("%#v:%#v\n", k, v)
	}

	// getcollateral
	cresp, err := c.GetCollateral()
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", cresp)

	// getchat
	cp := &model.ChatParams{
		FromDate: time.Now().AddDate(0, 0, -2),
	}
	chresp, err := c.GetChat(cp)
	if err != nil {
		l.Error(err.Error())
	}
	for k, v := range chresp {
		fmt.Printf("%#v:%#v\n", k, v)
	}

	// sendchildorder
	scor := &model.SendChildOrderRequest{
		ProductCode:     string(types.ProductCodeBTCJPY),
		ChildOrderType:  string(types.ChildOrderTypeLimit),
		Side:            string(types.SideBuy),
		Price:           30000,
		Size:            0.1,
		MinuteToExpires: 30 * 24 * 60,
		TimeInForce:     string(types.TimeInForceGTC),
	}
	scoresp, err := c.SendChildOrder(scor)
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", scoresp)

	// cancelchildorder
	ccor := &model.CancelChildOrderRequest{
		ProductCode:  string(types.ProductCodeBTCJPY),
		ChildOrderID: "test",
	}
	err = c.CancelChildOrder(ccor)
	if err != nil {
		l.Error(err.Error())
	}

	// cancelallchildorders
	cacor := &model.CancelAllChildOrdersRequest{
		ProductCode: string(types.ProductCodeBTCJPY),
	}
	err = c.CancelAllChildOrders(cacor)
	if err != nil {
		l.Error(err.Error())
	}

	// getchildorders
	op := &model.ChildOrderParams{
		ProductCode: types.ProductCodeBTCJPY,
	}
	or, err := c.GetChildOrders(op)
	if err != nil {
		l.Error(err.Error())
	}
	for k, v := range or {
		fmt.Printf("%#v:%#v\n", k, v)
	}

	// getexecutions
	ep := &model.ExecutionParams{
		ProductCode: types.ProductCodeBTCJPY,
	}
	er, err := c.GetExecutions(ep)
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", er)

	// getbalancehistory
	bhp := &model.BalanceHistoryParams{
		CurrencyCode: types.CurrencyCodeJPY,
	}
	bhr, err := c.GetBalanceHistory(bhp)
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", bhr)

	// getpositions
	pp := &model.PositionParams{
		ProductCode: types.ProductCodeBTCJPY,
	}
	pr, err := c.GetPositions(pp)
	if err != nil {
		l.Error(err.Error())
	}
	for k, v := range pr {
		fmt.Printf("%#v:%#v\n", k, v)
	}

	// gettradingcommission
	tcp := &model.TradingCommissionParams{
		ProductCode: types.ProductCodeBTCJPY,
	}
	tcr, err := c.GetTradingCommission(tcp)
	if err != nil {
		l.Error(err.Error())
	}
	fmt.Printf("%#v\n", tcr)
}
