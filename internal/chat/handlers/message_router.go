package handlers

import (
	messages "github.com/harshgupta9473/chatapp/internal/messages/kafka"
)

type ChatMessageRouter struct {
	Consumer           *messages.Consumer
	ChatMessageHandler *ChatMessageHandler
}

func NewMessageRouter(
	consumer *messages.Consumer,
	chathandler *ChatMessageHandler,
) *ChatMessageRouter {
	return &ChatMessageRouter{
		Consumer:           consumer,
		ChatMessageHandler: chathandler,
	}
}

func (m *ChatMessageRouter) RegisterWithConsumer() {
	//m.Consumer.RegisterHandler()
}
