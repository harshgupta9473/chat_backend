package handlers

import (
	"context"
	"encoding/json"
	"github.com/harshgupta9473/chatapp/internal/chat/dto"
	chatservice "github.com/harshgupta9473/chatapp/internal/chat/services"
	messagesdto "github.com/harshgupta9473/chatapp/internal/messages"
)

type ChatMessageHandler struct {
	ChatMessageService *chatservice.ChatMessageService
}

func (c *ChatMessageHandler) SendMessageHandler(ctx context.Context, msg *messagesdto.DomainMessage) error {
	var chatmsg dto.ChatMessage
	err := json.Unmarshal([]byte(msg.Payload), &chatmsg)
	if err != nil {
		return err
	}
	err = c.ChatMessageService.SendMessage(ctx, chatmsg)
	if err != nil {
		return err
	}
	return nil
}
