package kafka

import (
	"context"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"sync"
	"time"
)

type MessageHandler func(ctx context.Context, msg string) error

type Consumer struct {
	consumer *kafka.Consumer
	handlers map[string]MessageHandler
	wg       sync.WaitGroup
	stop     chan struct{}
}

func NewConsumer() (*Consumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":         "cfg.BootstrapServers",
		"group.id":                  "cfg.GroupID",
		"auto.offset.reset":         "earliest",
		"enable.auto.commit":        false,
		"max.poll.interval.ms":      300000, // 5 minutes
		"session.timeout.ms":        30000,  // 30 seconds
		"heartbeat.interval.ms":     10000,  // 10 seconds
		"fetch.min.bytes":           1,
		"fetch.max.bytes":           52428800, // 50MB
		"max.partition.fetch.bytes": 1048576,  // 1MB
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{
		consumer: c,
		handlers: make(map[string]MessageHandler),
		stop:     make(chan struct{}),
	}, nil
}

func (c *Consumer) RegisterHandler(packetType string, handler MessageHandler) {
	c.handlers[packetType] = handler
}

// starts consuming msg
func (c *Consumer) Start(ctx context.Context, topics []string) error {
	if len(topics) == 0 {
		return errors.New("no topics specified")
	}
	if err := c.consumer.SubscribeTopics(topics, nil); err != nil {
		return err
	}
	c.wg.Add(1)
	go c.consumeMessage(ctx)
	return nil
}

func (c *Consumer) consumeMessage(ctx context.Context) {
	defer c.wg.Done()

	for {
		select {
		case <-ctx.Done():
			return
		case <-c.stop:
			return
		default:
			msg, err := c.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				if err.(kafka.Error).Code() == kafka.ErrTimedOut {
					continue
				}
				continue
			}
			msgCopy := msg
			go func() {
				//process message

			}()

		}
	}
}
