package dics

import (
	"context"
	"github.com/harshgupta9473/chatapp/internal/chat/handlers"
	"github.com/harshgupta9473/chatapp/internal/chat/kafka"
	chatrepositories "github.com/harshgupta9473/chatapp/internal/chat/repositories"
	chatservice "github.com/harshgupta9473/chatapp/internal/chat/services"
	"github.com/harshgupta9473/chatapp/internal/messages"
	"log"
)

type ChatServiceContainer struct {
	MessageRouterHandler *handlers.ChatMessageRouter
	ChatMessageHandler   *handlers.ChatMessageHandler
	KafkaConsumer        *messages.Consumer
	KafkaProducer        *kafka.KafkaProducer
	ChatMessageService   *chatservice.ChatMessageService
	ChatRepository       *chatrepositories.ChatMessageRepository
}

func initializedics() (*ChatServiceContainer, error) {
	var err error
	c := &ChatServiceContainer{}
	c.KafkaConsumer, err = kafka.NewChatConsumer()
	if err != nil {
		return nil, err
	}
	c.KafkaProducer, err = kafka.NewKafkaChatProducer()
	if err != nil {
		return nil, err
	}
	c.ChatRepository, err = chatrepositories.NewChatMessageRepository()
	if err != nil {
		return nil, err
	}
	c.ChatMessageService = chatservice.NewChatMessageService(c.ChatRepository, c.KafkaProducer)
	c.ChatMessageHandler = handlers.NewChatMessageHandler(c.ChatMessageService)
	c.MessageRouterHandler = handlers.NewMessageRouter(c.KafkaConsumer, c.ChatMessageHandler)
	return c, nil
}

func InitializeChatService() error {
	container, err := initializedics()
	if err != nil {
		return err
	}
	ctx := context.Background()

	go func() {
		err := container.MessageRouterHandler.Consumer.Start(ctx)
		if err != nil {
			log.Printf("Kafka consumer error: %v\n", err)
		}
	}()
	return nil
}
