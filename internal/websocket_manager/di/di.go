package diws

import (
	"context"
	"github.com/harshgupta9473/chatapp/internal/messages"
	websocket_service "github.com/harshgupta9473/chatapp/internal/websocket_manager"
	handlers "github.com/harshgupta9473/chatapp/internal/websocket_manager/handler"
	"github.com/harshgupta9473/chatapp/internal/websocket_manager/kafka"
	"log"
	"net/http"
)

type WebsocketServiceContainer struct {
	webSocketManager   *websocket_service.WebSocketConnectionManager
	kafkaProducer      *kafka.KafkaProducer
	kafkaConsumer      *messages.Consumer
	websocketHandler   *handlers.WebSocketHandler
	producerHandler    *handlers.ProducerRouterHandler
	consumerHandler    *handlers.ConsumerHandler
	chatserviceHandler *handlers.ChatServiceHandler
	messagerouter      *handlers.WebSocketMessageRouter
}

func initializedependencies() (*WebsocketServiceContainer, error) {
	websocketManager := websocket_service.NewWebSocketConnectionManager()
	kafkaconsumer, err := kafka.NewWebSocketKafkaConsumer()
	if err != nil {
		return nil, err
	}
	kafkaproducer, err := kafka.NewWebSocketKafkaProducer()
	if err != nil {
		return nil, err
	}
	chatServiceHandler := handlers.NewChatServiceHandler(kafkaproducer)
	consumerHandler := handlers.NewConsumerHandler(websocketManager)
	producerHandler, err := handlers.NewProducerRouterHandler()
	if err != nil {
		return nil, err
	}
	producerHandler.RegisterHandler(chatServiceHandler)
	websocketHandler := handlers.NewWebSocketHandler(producerHandler, websocketManager)
	messageRouter := handlers.NewMessageRouter(kafkaconsumer)
	messageRouter.RegisterWithConsumer(consumerHandler)
	return &WebsocketServiceContainer{
		webSocketManager:   websocketManager,
		kafkaProducer:      kafkaproducer,
		kafkaConsumer:      kafkaconsumer,
		websocketHandler:   websocketHandler,
		producerHandler:    producerHandler,
		consumerHandler:    consumerHandler,
		chatserviceHandler: chatServiceHandler,
		messagerouter:      messageRouter,
	}, nil
}

func InitializeWS() error {
	websocketContainer, err := initializedependencies()
	ctx := context.Background()
	if err != nil {
		return err
	}

	// Start Kafka consumer (non-blocking)
	go func() {
		err := websocketContainer.kafkaConsumer.Start(ctx)
		if err != nil {
			log.Printf("Kafka consumer error: %v\n", err)
		}
	}()

	// Register WebSocket handler with HTTP
	http.HandleFunc("/ws", websocketContainer.websocketHandler.WebsocketHandler())

	// Start HTTP server (non-blocking if used inside bigger app)
	go func() {
		log.Println("WebSocket server started at :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	return nil
}
