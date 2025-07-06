package kafka

import (
	messages "github.com/harshgupta9473/chatapp/internal/messages/kafka"
)

type KafkaChatConsumer struct {
	Consumer *messages.Consumer
}

func SetUpChatKafkaConsumer() {
	consumer, err := messages.NewConsumer()
	if err != nil {
		panic(err)
	}
	kafkachatConsumer := &KafkaChatConsumer{}
	kafkachatConsumer.Consumer = consumer

}
