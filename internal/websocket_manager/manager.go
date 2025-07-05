package websocket_manager

import (
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

type WebSocketConnectionManager struct {
	Connections map[string]Connection
	mu          sync.Mutex
}

func NewWebSocketConnectionManager() *WebSocketConnectionManager {
	return &WebSocketConnectionManager{
		Connections: make(map[string]Connection),
	}
}

func (m *WebSocketConnectionManager) AddConnection(mobileNo string, conn *websocket.Conn) {
	m.mu.Lock()
	defer m.mu.Unlock()
	newConnection := NewWSConnection(conn)
	m.Connections[mobileNo] = newConnection
}

func (m *WebSocketConnectionManager) RemoveConnection(mobileNo string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.Connections, mobileNo)
}

func (m *WebSocketConnectionManager) GetConnection(mobileNo string) (Connection, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	conn, ok := m.Connections[mobileNo]
	if !ok {
		return nil, fmt.Errorf("connection not found for mobileNo: %s", mobileNo)
	}
	return conn, nil
}
