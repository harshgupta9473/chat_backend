package main

import (
	dics "github.com/harshgupta9473/chatapp/internal/chat/di"
	diws "github.com/harshgupta9473/chatapp/internal/websocket_manager/di"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	errChan := make(chan error, 2)

	go func() {
		errChan <- dics.InitializeChatService()
	}()

	// WebSocket Manager init
	go func() {
		errChan <- diws.InitializeWS()
	}()

	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			log.Fatalf(" Initialization error: %v", err)
			return
		}
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)
	log.Println("Server is running. Press Ctrl+C to exit.")
	<-sigs

}
