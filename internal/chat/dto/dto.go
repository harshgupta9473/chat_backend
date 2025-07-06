package dto

import (
	"time"
)

type ChatMessage struct {
	ID               int64
	ChatID           string
	SenderMobileNo   string
	ReceiverMobileNo string
	Message          string
	IsRead           bool
	CreatedAt        time.Time
}
