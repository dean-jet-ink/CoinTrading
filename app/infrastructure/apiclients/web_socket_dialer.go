package apiclients

import (
	"cointrading/app/config"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketDialer struct {
	dialer *websocket.Dialer
}

func NewWebSocketDialer() *WebSocketDialer {
	timeoutSec := config.WebSocketHandshakeTimeout()
	buffer := config.WebSocketBufferSize()

	return &WebSocketDialer{
		dialer: &websocket.Dialer{
			HandshakeTimeout: time.Duration(timeoutSec) * time.Second,
			ReadBufferSize:   buffer,
			WriteBufferSize:  buffer,
			Proxy:            http.ProxyFromEnvironment,
		},
	}
}

func (w *WebSocketDialer) Dial(urlStr string, requestHeader http.Header) (*websocket.Conn, *http.Response, error) {
	return w.dialer.Dial(urlStr, requestHeader)
}

type JsonRPC2 struct {
	Version string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	Result  any    `json:"result,omitempty"`
	Id      *int   `json:"id,omitempty"`
}
