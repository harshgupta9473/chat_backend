package kafka

import messages "github.com/harshgupta9473/chatapp/internal/messages/kafka"

type KafkaProducer struct {
	Producer *messages.Producer
}
