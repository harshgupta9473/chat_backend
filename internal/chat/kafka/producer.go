package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	messages "github.com/harshgupta9473/chatapp/internal/messages/kafka"
)

type KafkaProducer struct {
	producer *messages.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
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

func (p *KafkaProducer) SendMessage(msg *kafka.Message) error {

}
