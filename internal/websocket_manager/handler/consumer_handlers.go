package handlers

import (
	"encoding/json"
	messages2 "github.com/harshgupta9473/chatapp/internal/messages"
	websocket_service "github.com/harshgupta9473/chatapp/internal/websocket_manager"
)

type ConsumerHandler struct {
	websocketManager *websocket_service.WebSocketConnectionManager
}

func (c *ConsumerHandler) SendMessageToUser(msg *messages2.DomainMessage) error {
	conn, err := c.websocketManager.GetConnection(msg.Header.MobileNumber)
	if err != nil {
		return err
	}
	msgbytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	conn.WriteMsg(msgbytes)
	return nil
}
