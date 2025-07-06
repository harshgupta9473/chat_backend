package websocket_manager

import (
	"github.com/gorilla/websocket"
	"github.com/harshgupta9473/chatapp/internal/messages/kafka"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	Upgrader         *websocket.Upgrader
	WebsocketManager *WebSocketConnectionManager
	KafkaProducer    *messages.Producer
	KafkaConsumer    *messages.Consumer
}

func NewWebSocketHandler() *WebSocketHandler {
	var err error
	handler := &WebSocketHandler{
		Upgrader:         &upgrader,
		WebsocketManager: NewWebSocketConnectionManager(),
	}
	handler.KafkaProducer, err = messages.NewProducer()
	if err != nil {
		return &WebSocketHandler{}
	}
	handler.KafkaConsumer, err = messages.NewConsumer()
	if err != nil {
		return &WebSocketHandler{}
	}
	return handler
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func (ws *WebSocketHandler) WebsocketHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mobilnoStr := r.URL.Query().Get("mobilno")
		if mobilnoStr != "" {

		}

		wsconn, err := ws.Upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err)
			return
		}

		ws.WebsocketManager.AddConnection(mobilnoStr, wsconn)
		conn, err := ws.WebsocketManager.GetConnection(mobilnoStr)
		if err != nil {

		}

		go func() {
			for {
				msg := <-conn.ReadMsg()
				log.Println(msg)
			}
		}()
	}
}
