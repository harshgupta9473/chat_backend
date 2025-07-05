package websocket_manager

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func websocket_handler(websocketManger *WebSocketConnectionManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mobilnoStr := r.URL.Query().Get("mobilno")
		if mobilnoStr != "" {

		}

		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {

		}
		websocketManger.AddConnection(mobilnoStr, ws)
		conn, err := websocketManger.GetConnection(mobilnoStr)
		if err != nil {

		}

		go func() {
			for {
				msg := <-conn.ReadMsg()

			}
		}()
	}
}
