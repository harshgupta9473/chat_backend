package handlers

import (
	"github.com/harshgupta9473/chatapp/internal/messages"
)

type WebSocketMessageRouter struct {
	Consumer *messages.Consumer
}

func NewMessageRouter(
	consumer *messages.Consumer,
) *WebSocketMessageRouter {
	return &WebSocketMessageRouter{
		Consumer: consumer,
	}
}

func (m *WebSocketMessageRouter) RegisterWithConsumer() {
	//	m.Consumer.RegisterHandler("sm", )
}
