package chatrepositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/harshgupta9473/chatapp/internal/chat/dto"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type ChatMessageRepository struct {
	DB *sql.DB
}

func NewChatMessageRepository() (*ChatMessageRepository, error) {
	repo := &ChatMessageRepository{}
	connStr := "host=localhost port=5432 user=postgres password=supersecurepass dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	repo.DB = db
	return repo, nil
}

func (ch *ChatMessageRepository) SaveMsgInDB(ctx context.Context, chat *dto.ChatMessage) error {
	query := `
		INSERT INTO messages (chat_id, sender_mobile, receiver_mobile, message, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING created_at
	`

	err := ch.DB.QueryRowContext(
		ctx,
		query,
		chat.ChatID,
		chat.SenderMobileNo,
		chat.ReceiverMobileNo,
		chat.Message,
		time.Now(),
	).Scan(&chat.CreatedAt)

	if err != nil {
		return err
	}
	return nil
}

func (ch *ChatMessageRepository) GetMessagesByChatID(ctx context.Context, chatID uuid.UUID, limit, offset int) ([]*dto.ChatMessage, error) {
	query := `
		SELECT id, chat_id, sender_mobile, receiver_mobile, message, is_read, created_at
		FROM messages
		WHERE chat_id = $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`

	rows, err := ch.DB.QueryContext(ctx, query, chatID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*dto.ChatMessage
	for rows.Next() {
		var m dto.ChatMessage
		err := rows.Scan(&m.ID, &m.ChatID, &m.SenderMobileNo, &m.ReceiverMobileNo, &m.Message, &m.IsRead, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}

func (ch *ChatMessageRepository) GetUnreadMessages(ctx context.Context, receiverMobile string, since time.Time) ([]*dto.ChatMessage, error) {
	query := `
		SELECT id, chat_id, sender_mobile, receiver_mobile, message, is_read, created_at
		FROM messages
		WHERE receiver_mobile = $1 AND created_at > $2 AND is_read = false
		ORDER BY created_at ASC
	`

	rows, err := ch.DB.QueryContext(ctx, query, receiverMobile, since)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []*dto.ChatMessage
	for rows.Next() {
		var m dto.ChatMessage
		err := rows.Scan(&m.ID, &m.ChatID, &m.SenderMobileNo, &m.ReceiverMobileNo, &m.Message, &m.IsRead, &m.CreatedAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, &m)
	}
	return messages, nil
}

func (ch *ChatMessageRepository) MarkAllMessagesAsRead(ctx context.Context, receiverMobile, senderMobile string) error {
	query := `
		UPDATE messages
		SET is_read = true
		WHERE receiver_mobile = $1 AND sender_mobile = $2 AND is_read = false
	`

	result, err := ch.DB.ExecContext(ctx, query, receiverMobile, senderMobile)
	if err != nil {
		return fmt.Errorf("failed to mark messages as read for receiver %s from sender %s: %w", receiverMobile, senderMobile, err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("Marked %d messages as read for receiver %s from sender %s", rowsAffected, receiverMobile, senderMobile)
	return nil
}
