package messages

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"sync"
	"time"
)

type MessageHandler func(ctx context.Context, msg *DomainMessage) error

type Consumer struct {
	consumer *kafka.Consumer
	handlers map[string]MessageHandler
	topics   []string
	ctx      context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	stop     chan struct{}
}

func NewConsumer(kafkaConfig *kafka.ConfigMap, topics []string) (*Consumer, error) {
	c, err := kafka.NewConsumer(kafkaConfig)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Consumer{
		consumer: c,
		handlers: make(map[string]MessageHandler),
		topics:   topics,
		ctx:      ctx,
		cancel:   cancel,
		stop:     make(chan struct{}),
	}, nil
}

func (c *Consumer) RegisterHandler(packetType string, handler MessageHandler) {
	c.handlers[packetType] = handler
}

// starts consuming msg
func (c *Consumer) Start(ctx context.Context) error {
	if len(c.topics) == 0 {
		return errors.New("no topics specified")
	}
	if err := c.consumer.SubscribeTopics(c.topics, nil); err != nil {
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
				log.Println(msgCopy)

			}()

		}
	}
}

func (c *Consumer) processMessage(ctx context.Context, msg *kafka.Message) error {
	// Parse the domain message
	var domainMsg DomainMessage
	if err := json.Unmarshal(msg.Value, &domainMsg); err != nil {
		return err
	}

	// Get packet name from message or header
	packetName := domainMsg.Header.PacketName
	if packetName == "" {
		log.Println("WARNING: no packet name provided")
		return nil
	}

	// Find handler for this packet type
	handler, exists := c.handlers[packetName]
	if !exists {
		log.Printf("Received message with no handler for packet '%s'", packetName)
		return nil
	}

	// Create context with timeout
	msgCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Call the handler
	return handler(msgCtx, &domainMsg)
}
