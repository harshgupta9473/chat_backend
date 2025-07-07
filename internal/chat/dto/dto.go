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

type UndeliveredMessage struct {
	ID             int        `db:"id"`
	MessageID      int        `db:"message_id"` // FK to messages.id
	ReceiverMobile string     `db:"receiver_mobile"`
	CreatedAt      time.Time  `db:"created_at"`
	DeletedAt      *time.Time `db:"deleted_at"` // Soft delete: null = active
}
