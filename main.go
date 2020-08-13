package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bmf-san/bitbankgo/asset"
	"github.com/bmf-san/bitbankgo/candlestick"
	"github.com/bmf-san/bitbankgo/client"
	"github.com/bmf-san/bitbankgo/depth"
	"github.com/bmf-san/bitbankgo/order"
	"github.com/bmf-san/bitbankgo/setting"
	"github.com/bmf-san/bitbankgo/status"
	"github.com/bmf-san/bitbankgo/ticker"
	"github.com/bmf-san/bitbankgo/trade"
	"github.com/bmf-san/bitbankgo/transaction"
	"github.com/bmf-san/bitbankgo/types"
	"github.com/bmf-san/bitbankgo/withdrawal"
)

// GET /{pair}/ticker
func getTicker(client *client.Client) {
	ticker := &ticker.Ticker{
		Client: client,
	}
	s, err := ticker.Get(types.BtcJpy)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(s)
}

// GET /{pair}/depth
func getDepth(client *client.Client) {
	depth := &depth.Depth{
		Client: client,
	}
	d, err := depth.Get(types.BtcJpy)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(d)
}

// GET /{pair}/transactions/{YYYYMMDD}
func getTransaction(client *client.Client) {
	tc := &transaction.Transaction{
		Client: client,
	}
	t, err := tc.Get(types.BtcJpy, time.Now().AddDate(0, 0, -2))
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(t)
}

// GET /{pair}/candlestick/{candle-type}/{YYYYMMDD|YYYY}
func getCandlestick(client *client.Client) {
	cc := &candlestick.CandleStick{
		Client: client,
	}
	c, err := cc.Get(types.BtcJpy, types.OneMin, time.Now().AddDate(0, 0, -2))
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(c)
}

// GET /v1/user/assets
func getAsset(client *client.Client) {
	ac := &asset.Asset{
		Client: client,
	}
	a, err := ac.Get()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(a)
}

// GET /v1/user/spot/order
func getOrder(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.GetOrder(order.GetOrderParams{
		Pair:    types.BtcJpy,
		OrderID: 11004961431,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// POST /v1/user/spot/order
func postOrder(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.PostOrder(order.PostOrderParams{
		Pair:   types.BtcJpy,
		Amount: "1",
		Price:  "900",
		Side:   "buy",
		Type:   "limit",
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// POST /v1/user/spot/cancel_order
func postCancelOrder(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.PostCancelOrder(order.PostCancelOrderParams{
		Pair:    types.BtcJpy,
		OrderID: 11003711501,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// POST /v1/user/spot/cancel_orders
func postCancelOrders(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.PostCancelOrders(order.PostCancelOrdersParams{
		Pair:     types.BtcJpy,
		OrderIDs: []int{11004377787},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// GET /v1/user/spot/orders_info.
func getOrdersInfo(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.GetOrdersInfo(order.GetOrdersInfoParams{
		Pair:     types.BtcJpy,
		OrderIDs: []int{11004567037},
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// GET /v1/user/spot/active_orders
func getActiveOrders(client *client.Client) {
	oc := &order.Order{
		Client: client,
	}
	o, err := oc.GetActiveOrders(order.GetActiveOrdersParams{
		Pair:  types.BtcJpy,
		Count: 2,
		// FromID: 11004961431,
		// EndID:  11004567037,
		// Since:  1597499629,
		// End:    1597586029,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(o)
}

// GET /v1/user/spot/trade_history
func getTradeHistory(client *client.Client) {
	tc := &trade.Trade{
		Client: client,
	}
	t, err := tc.GetTradeHistory(trade.GetTradeHistoryParams{
		Pair: types.BtcJpy,
		// Count: 2,
		// OrderID: 11004961431,
		// Since:  11004567037,
		// Eend:  1597499629,
		// Order: asc,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(t)

}

// GET /v1/user/withdrawal_account
func getWithdrawalAccount(client *client.Client) {
	wc := &withdrawal.Withdrawal{
		Client: client,
	}
	w, err := wc.GetWithdrawalAccount(withdrawal.GetWithdrawalAccountParams{
		Asset: types.Jpy,
	})
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(w)
}

// GET /v1/spot/status
func getStatus(client *client.Client) {
	sc := &status.Status{
		Client: client,
	}
	s, err := sc.GetStatus()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(s)
}

// GET /v1/spot/pairs
func getSetting(client *client.Client) {
	sc := &setting.Setting{
		Client: client,
	}
	s, err := sc.GetSetting()
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Println(s)
}

func main() {
	apiKey := os.Getenv("BITBANK_API_KEY")
	apiSecret := os.Getenv("BITBANK_API_SECRET")
	client := client.New(apiKey, apiSecret)

	// NOTE: Write a your trading logic by using a bitbank api.
	getTicker(client)
	// getDepth(client)
	// getTransaction(client)
	// getCandlestick(client)
	// getAsset(client)
	// getOrder(client)
	// postOrder(client)
	// postCancelOrder(client)
	// postCancelOrders(client)
	// getOrdersInfo(client)
	// getActiveOrders(client)
	// getTradeHistory(client)
	// getWithdrawalAccount(client)
	// getStatus(client)
	// getSetting(client)
}
