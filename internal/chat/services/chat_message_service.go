package chatservice

import (
	"context"
	"github.com/harshgupta9473/chatapp/internal/chat/dto"
	chatrepositories "github.com/harshgupta9473/chatapp/internal/chat/repositories"
	"github.com/harshgupta9473/chatapp/internal/chat/utils"
)

type ChatMessageService struct {
	ChatMessageRepo *chatrepositories.ChatMessageRepository
}

func NewChatMessageService() *ChatMessageService {
	chatMessageRepo := chatrepositories.NewChatMessageRepository()
	return &ChatMessageService{chatMessageRepo}
}

func (ch *ChatMessageService) SendMessage(ctx context.Context, message dto.ChatMessage) error {
	message.ChatID = utils.GenerateChatIDForUsers(message.SenderMobileNo, message.ReceiverMobileNo)
	ch.ChatMessageRepo.SaveMsgInDB(ctx, &message)

	return nil
}
