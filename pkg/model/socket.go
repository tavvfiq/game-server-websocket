package model

import "github.com/gorilla/websocket"

type SocketData struct {
}

type SocketPayload struct {
	EventType string
	Data      []byte
}

type SocketResponse struct {
	EventType string
	Data      []byte
	Message   string
}

type WebSocketConnection struct {
	*websocket.Conn
	ServerID string
	PlayerID string
}
