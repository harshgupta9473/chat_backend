package handlers

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	messages2 "github.com/harshgupta9473/chatapp/internal/messages"
	"github.com/harshgupta9473/chatapp/internal/websocket_manager"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	Upgrader         *websocket.Upgrader
	WebsocketManager *websocket_service.WebSocketConnectionManager
	ProducerHandler  *ProducerRouterHandler
}

func NewWebSocketHandler(
	producerhandler *ProducerRouterHandler,
	websocketManager *websocket_service.WebSocketConnectionManager,
) *WebSocketHandler {
	handler := &WebSocketHandler{
		Upgrader:         &upgrader,
		WebsocketManager: websocketManager,
		ProducerHandler:  producerhandler,
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
				data := <-conn.ReadMsg()
				var msg messages2.DomainMessage
				err := json.Unmarshal(data, &msg)
				if err != nil {
					log.Println(err)
					continue
				}
				go func() {
					ws.processmessage(&msg)
				}()
			}
		}()
	}
}

func (ws *WebSocketHandler) processmessage(msg *messages2.DomainMessage) {
	if msg.Header.DestinationService == "" {
		log.Println("invalid destination services")
		return
	}
	handler, ok := ws.ProducerHandler.handlers[msg.Header.DestinationService]
	if !ok {
		log.Println("invalid destination services no handler available")
		return
	}
	ctx := ws.ProducerHandler.ctx
	handler.HandleMessage(ctx, msg)
}
