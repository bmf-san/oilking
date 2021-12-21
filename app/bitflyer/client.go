package bitflyer

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"

	"github.com/bmf-san/oilking/app/logger"
	"github.com/bmf-san/oilking/app/model"
	"github.com/bmf-san/oilking/app/types"
	"github.com/gorilla/websocket"
)

const (
	httpAPIScheme      = "https"
	httpAPIHost        = "api.bitflyer.jp"
	httpAPIVersionPath = "/v1"
	realtimeAPIScheme  = "wss"
	realtimeAPIHost    = "ws.lightstream.bitflyer.com"
	realtimeAPIPath    = "/json-rpc"
	jsonRPCVersion     = "2.0"
)

// Client is a client.
type Client struct {
	httpclient *http.Client
	wsConn     *websocket.Conn
	apikey     string
	apisecret  string
	Logger     *logger.Logger
}

// NewClient creates a new client.
func NewClient(apikey string, apisecret string, logger *logger.Logger) *Client {
	return &Client{
		httpclient: &http.Client{
			Timeout: time.Duration(100 * time.Second),
		},
		apikey:    apikey,
		apisecret: apisecret,
		Logger:    logger,
	}
}

// It seems that it is necessary to change the sign when requesting with query and when requesting with body.
// Not all api clients are implemented and may need to be modified later.
func (c *Client) sign(text string) string {
	mac := hmac.New(sha256.New, []byte(c.apisecret))
	mac.Write([]byte(text))

	return hex.EncodeToString(mac.Sum(nil))
}

func (c *Client) certificationHeader(ts string, sign string) map[string]string {
	return map[string]string{
		"ACCESS-KEY":       c.apikey,
		"ACCESS-TIMESTAMP": ts,
		"ACCESS-SIGN":      sign,
		"Content-Type":     "application/json",
	}
}

// Do sends an HTTP request and returns an HTTP response.
func (c *Client) Do(method string, apipath string, query map[string]string, data []byte) ([]byte, error) {
	p := path.Join(httpAPIVersionPath, apipath)
	u := url.URL{Scheme: httpAPIScheme, Host: httpAPIHost, Path: p}

	req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	base := ts + method + p
	txt := base + string(data)
	if len(q.Encode()) > 0 {
		txt = base + "?" + q.Encode()
	}
	sign := c.sign(txt)
	for key, value := range c.certificationHeader(ts, sign) {
		req.Header.Add(key, value)
	}

	c.Logger.Info(logger.AccessLogEntry{
		Method: req.Method,
		URL:    req.URL.String(),
		Body:   string(data),
	})

	resp, err := c.httpclient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}

	rbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rbody, nil
}

// connWs gets websocket connection.
func (c *Client) connWs(w *model.Writer) {
	u := url.URL{Scheme: realtimeAPIScheme, Host: realtimeAPIHost, Path: realtimeAPIPath}
	p := u.String()

	conn, _, err := websocket.DefaultDialer.Dial(p, nil)
	if err != nil {
		c.Logger.Error(logger.Entry{
			Message: err.Error(),
		})
	}
	c.wsConn = conn

	j := &model.JSONRPC2{
		Version: jsonRPCVersion,
		Method:  string(w.Method),
		Params: &model.SubscribeParams{
			Channel: string(w.Channel),
		},
	}

	if err := conn.WriteJSON(j); err != nil {
		c.Logger.Error(logger.Entry{
			Message: "write:" + err.Error(),
		})
		return
	}

	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
}

// DoWs sends an websocket request and returns an websocket response.
func (c *Client) DoWs(w *model.Writer) {
LOOP:
	for {
		if c.wsConn == nil {
			c.connWs(w)
		}

		j := new(model.JSONRPC2)

		c.wsConn.SetWriteDeadline(time.Now().Add(10 * time.Second))

		if err := c.wsConn.ReadJSON(j); err != nil {
			c.Logger.Error(logger.Entry{
				Message: "read:" + err.Error(),
			})

			c.wsConn.Close()
			c.wsConn = nil
			continue LOOP
		}

		if j.Method == "channelMessage" {
			switch v := j.Params.(type) {
			case map[string]interface{}:
				for key, bin := range v {
					if key == "message" {
						mbin, err := json.Marshal(bin)
						if err != nil {
							c.Logger.Error(logger.Entry{
								Message: err.Error(),
							})
							continue LOOP
						}

						switch w.Channel {
						case types.ChannelTickerBTCJPY:
							var tres model.TickerResponse
							if err := json.Unmarshal(mbin, &tres); err != nil {
								c.Logger.Error(logger.Entry{
									Message: err.Error(),
								})
								continue LOOP
							}
							w.ChannelTicker <- tres
						case types.ChannelTickerFXBTCJPY:
							var tres model.TickerResponse
							if err := json.Unmarshal(mbin, &tres); err != nil {
								c.Logger.Error(logger.Entry{
									Message: err.Error(),
								})
								continue LOOP
							}
							w.ChannelTicker <- tres
						default:
							c.Logger.Error(logger.Entry{
								Message: errors.New("there is no channel for response").Error(),
							})
							continue LOOP
						}
					}
				}
			}
		}
	}
}
