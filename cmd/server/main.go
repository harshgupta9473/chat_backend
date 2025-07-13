package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	dics "github.com/harshgupta9473/chatapp/internal/chat/di"
	dius "github.com/harshgupta9473/chatapp/internal/userservice/di"
	diws "github.com/harshgupta9473/chatapp/internal/websocket_manager/di"
)

func main() {
	errChan := make(chan error, 3)

	// Chat service
	go func() {
		log.Println("Initializing Chat Service...")
		errChan <- dics.InitializeChatService()
	}()

	// WebSocket Manager
	go func() {
		log.Println("Initializing WebSocket Service...")
		errChan <- diws.InitializeWS()
	}()

	// User Service
	go func() {
		log.Println("Initializing User Service...")
		errChan <- dius.InitializeUserService()
	}()

	// Wait for all 3 services to initialize
	for i := 0; i < 3; i++ {
		if err := <-errChan; err != nil {
			log.Fatalf("Initialization error: %v", err)
			return
		}
	}

	log.Println("All services initialized successfully.")
	log.Println(" Server is running. Press Ctrl+C to exit.")

	// Wait for OS interrupt to gracefully shut down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	<-sigs
	log.Println("Server shutdown signal received.")
}
