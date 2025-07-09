package handlers

import (
	"context"
	messages2 "github.com/harshgupta9473/chatapp/internal/messages"
	kafka2 "github.com/harshgupta9473/chatapp/internal/websocket_manager/kafka"

	"time"
)

type ChatServiceHandler struct {
	producer *kafka2.KafkaProducer
}

func NewChatServiceHandler(producer *kafka2.KafkaProducer) *ChatServiceHandler {
	return &ChatServiceHandler{
		producer: producer,
	}
}

func (c *ChatServiceHandler) HandleMessage(ctx context.Context, msg *messages2.DomainMessage) error {
	switch msg.Header.PacketName {
	case "sm":
		return c.handleMessageSendReq(ctx, msg)

	}
	return nil
}

func (h *ChatServiceHandler) handleMessageSendReq(ctx context.Context, msg *messages2.DomainMessage) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	return h.producer.PublishEvents(ctx, msg)
}
