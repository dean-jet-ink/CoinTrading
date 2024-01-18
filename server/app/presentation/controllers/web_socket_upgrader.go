package controllers

import (
	"cointrading/app/config"
	"cointrading/app/domain/myerror"
	"cointrading/app/presentation/router"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketUpgrader struct {
	upgrader *websocket.Upgrader
}

var allowedOrigins = map[string]bool{
	"http://localhost:30000": true,
}

func NewWebSocketUpgrader() *WebSocketUpgrader {
	timeoutSec := config.WebSocketHandshakeTimeout()
	buffer := config.WebSocketBufferSize()

	return &WebSocketUpgrader{
		upgrader: &websocket.Upgrader{
			HandshakeTimeout: time.Duration(timeoutSec) * time.Second,
			ReadBufferSize:   buffer,
			WriteBufferSize:  buffer,
			CheckOrigin: func(r *http.Request) bool {
				origin := r.Header.Get("Origin")
				return allowedOrigins[origin]
			},
		},
	}
}

func (w *WebSocketUpgrader) Upgrade(c router.Context) (*websocket.Conn, error) {
	conn, err := w.upgrader.Upgrade(c.Writer(), c.Request(), nil)
	if err != nil {
		err = fmt.Errorf("%w: Failed to upgrade websocket connection: %v", myerror.ErrFailedToConnectNetwork, err)
		return nil, err
	}

	return conn, nil
}
