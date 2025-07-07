package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	messages "github.com/harshgupta9473/chatapp/internal/messages/kafka"
)

func NewChatConsumer() (*messages.Consumer, error) {
	kafkaConfig := &kafka.ConfigMap{
		"bootstrap.servers":         "localhost:9092",
		"group.id":                  "chat_service",
		"auto.offset.reset":         "earliest",
		"enable.auto.commit":        false,
		"max.poll.interval.ms":      300000, // 5 minutes
		"session.timeout.ms":        30000,  // 30 seconds
		"heartbeat.interval.ms":     10000,  // 10 seconds
		"fetch.min.bytes":           1,
		"fetch.max.bytes":           52428800, // 50MB
		"max.partition.fetch.bytes": 1048576,  // 1MB
	}
	consumer, err := messages.NewConsumer(kafkaConfig, []string{"chat_message"})
	if err != nil {
		return nil, err
	}
	return consumer, nil
}
