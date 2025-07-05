package kafka

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer() (*Producer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            "cfg.BootstrapServers",
		"client.id":                    " cfg.ClientID",
		"acks":                         "all",
		"delivery.timeout.ms":          5000,
		"queue.buffering.max.messages": 100000,
		"linger.ms":                    5,
		"batch.size":                   16384,
		"compression.type":             "snappy",
	})
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: p,
	}, nil
}

func (p *Producer) PublishMessage(ctx context.Context, message string, topic string) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}
	// publish msg to kafka
	kafkamsg := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Key:   []byte(message),
		Value: data,
		Headers: []kafka.Header{
			{
				Key:   "packet_name",
				Value: []byte(message),
			},
			{
				Key:   "source_service",
				Value: []byte(message),
			},
			{
				Key:   "destination_service",
				Value: []byte(message),
			},
		},
		Timestamp: time.Now(),
	}
	err = p.producer.Produce(kafkamsg, nil)
	if err != nil {
		return err
	}
	return nil
}
