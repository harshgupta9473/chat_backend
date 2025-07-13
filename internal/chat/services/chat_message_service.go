package chatservice

import (
	"context"
	"github.com/harshgupta9473/chatapp/internal/chat/dto"
	"github.com/harshgupta9473/chatapp/internal/chat/kafka"
	chatrepositories "github.com/harshgupta9473/chatapp/internal/chat/repositories"
	"github.com/harshgupta9473/chatapp/internal/chat/utils"
)

type ChatMessageService struct {
	ChatMessageRepo *chatrepositories.ChatMessageRepository
	Emitter         kafka.EventEmitter
}

func NewChatMessageService(chatMessageRepo *chatrepositories.ChatMessageRepository, emitter kafka.EventEmitter) *ChatMessageService {
	return &ChatMessageService{chatMessageRepo, emitter}
}

func (ch *ChatMessageService) SendMessage(ctx context.Context, message dto.ChatMessage) error {
	message.ChatID = utils.GenerateChatIDForUsers(message.SenderMobileNo, message.ReceiverMobileNo)
	ch.ChatMessageRepo.SaveMsgInDB(ctx, &message)
	return nil
}
