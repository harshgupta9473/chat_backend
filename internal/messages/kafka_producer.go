package messages

import (
	"context"
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"time"
)

type Producer struct {
	producer *kafka.Producer
}

func NewProducer(kafkaConfig *kafka.ConfigMap) (*Producer, error) {
	p, err := kafka.NewProducer(kafkaConfig)
	if err != nil {
		return nil, err
	}
	return &Producer{
		producer: p,
	}, nil
}

func (p *Producer) PublishMessage(ctx context.Context, message *DomainMessage, topic string) error {
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
		Key:   []byte(message.Header.MobileNumber),
		Value: data,
		Headers: []kafka.Header{
			{
				Key:   "packet_name",
				Value: []byte(message.Header.PacketName),
			},
			{
				Key:   "source_service",
				Value: []byte(message.Header.SourceService),
			},
			{
				Key:   "destination_service",
				Value: []byte(message.Header.DestinationService),
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
