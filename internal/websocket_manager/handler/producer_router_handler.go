package handlers

import (
	"context"
	messages2 "github.com/harshgupta9473/chatapp/internal/messages"
	"time"
)

type Handler interface {
	HandleMessage(ctx context.Context, msg *messages2.DomainMessage) error
}

type ProducerRouterHandler struct {
	handlers map[string]Handler
	ctx      context.Context
	cancel   context.CancelFunc
}

func NewProducerRouterHandler() (*ProducerRouterHandler, error) {
	handlers := make(map[string]Handler)
	ctx, cancel := context.WithCancel(context.Background())
	producrHandler := &ProducerRouterHandler{
		handlers: handlers,
		ctx:      ctx,
		cancel:   cancel,
	}
	return producrHandler, nil
}

func (h *ProducerRouterHandler) RegisterHandler(
	chatServiceHandler *ChatServiceHandler,
) {
	h.handlers["chat_service"] = chatServiceHandler
}

func (h *ProducerRouterHandler) HandleMessage(ctx context.Context, msg *messages2.DomainMessage) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return h.handlers[msg.Header.DestinationService].HandleMessage(h.ctx, msg)
}
