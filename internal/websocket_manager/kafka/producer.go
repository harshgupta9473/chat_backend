package kafka

import (
	"context"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	messages "github.com/harshgupta9473/chatapp/internal/messages"
)

type KafkaProducer struct {
	producer *messages.Producer
}

func NewWebSocketKafkaProducer() (*KafkaProducer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":            "localhost:9092",
		"client.id":                    "websocket_service",
		"acks":                         "all",
		"delivery.timeout.ms":          5000,
		"queue.buffering.max.messages": 100000,
		"linger.ms":                    5,
		"batch.size":                   16384,
		"compression.type":             "snappy",
	}
	producer, err := messages.NewProducer(config)
	if err != nil {
		return nil, err
	}
	return &KafkaProducer{
		producer: producer,
	}, nil
}

func (p *KafkaProducer) PublishEvents(ctx context.Context, message *messages.DomainMessage) error {
	msg, err := messages.NewDomainMessage(
		message.Header.MobileNumber,
		message.Header.PacketName,
		"websocket_service",
		message.Header.DestinationService,
		message.Payload,
	)
	if err != nil {
		return err
	}
	return p.producer.PublishMessage(ctx, msg, "chat_req")
}
