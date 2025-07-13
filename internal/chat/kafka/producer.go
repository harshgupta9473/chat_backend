package kafka

import (
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	messages "github.com/harshgupta9473/chatapp/internal/messages"
)

type KafkaProducer struct {
	producer *messages.Producer
}

type EventEmitter interface {
	Emit(ctx context.Context, data interface{}, packetName string, mobilenumber string) error
}

func NewKafkaChatProducer() (*KafkaProducer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":            "localhost:9092",
		"client.id":                    "chat_service",
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

func (p *KafkaProducer) PublishEvents(ctx context.Context, data interface{}, packetName string, mobilenumber string) error {
	if packetName == "" {
		return errors.New("invalid packet name")
	}
	if packetName == "cm" {
		msg, err := messages.NewDomainMessage(
			mobilenumber,
			packetName,
			"chat_service",
			"websocket_service",
			data,
		)
		if err != nil {
			return err
		}
		return p.producer.PublishMessage(ctx, msg, "chat_res")
	}
	return nil
}

func (p *KafkaProducer) Emit(ctx context.Context, data interface{}, packetName string, mobilenumber string) error {
	return p.PublishEvents(ctx, data, packetName, mobilenumber)
}
