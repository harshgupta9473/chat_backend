package handlers

import (
	"context"
	messages2 "github.com/harshgupta9473/chatapp/internal/messages"
	"github.com/harshgupta9473/chatapp/internal/websocket_manager/kafka"
	"time"
)

type ProducerHandler struct {
	producer *kafka.KafkaProducer
}

func (h *ProducerHandler) HandleMessageSendReq(msg *messages2.DomainMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return h.producer.PublishEvents(ctx, msg)
}
