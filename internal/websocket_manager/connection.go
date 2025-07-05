package websocket_manager

import "github.com/gorilla/websocket"

type Connection interface {
	Start()
	ReadMsg() <-chan []byte
	WriteMsg(msg []byte)
	Close()
}

type WSConnection struct {
	Conn      *websocket.Conn
	readChan  chan []byte
	writeChan chan []byte
}

func NewWSConnection(conn *websocket.Conn) Connection {
	return &WSConnection{
		Conn:      conn,
		readChan:  make(chan []byte),
		writeChan: make(chan []byte),
	}
}

func (conn *WSConnection) Start() {
	go conn.read()
	go conn.write()
}

func (conn *WSConnection) ReadMsg() <-chan []byte {
	return conn.readChan
}

func (conn *WSConnection) WriteMsg(msg []byte) {
	conn.writeChan <- msg
}

func (conn *WSConnection) Close() {
	conn.Conn.Close()
}

func (conn *WSConnection) read() {
	for {
		_, msg, err := conn.Conn.ReadMessage()
		if err != nil {
			return
		}
		conn.readChan <- msg
	}
}

func (conn *WSConnection) write() {
	for {
		msg := <-conn.writeChan
		err := conn.Conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			return
		}
	}
}

func (conn *WSConnection) close() {
	conn.Conn.Close()
}
