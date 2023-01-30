package store

import (
	"sync"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
)

type _connections struct {
	mtx         sync.Mutex
	connections []*model.WebSocketConnection
}

var conns = _connections{}

func GetConnections() []*model.WebSocketConnection {
	return conns.connections
}

func AddConnections(conn *model.WebSocketConnection) error {
	conns.mtx.Lock()
	conns.connections = append(conns.connections, conn)
	conns.mtx.Unlock()
	return nil
}

func RemoveConnection(thisConn *model.WebSocketConnection) error {
	conns.mtx.Lock()
	newConn := make([]*model.WebSocketConnection, 0)
	for _, conn := range conns.connections {
		if conn != thisConn {
			newConn = append(newConn, conn)
		}
	}
	conns.connections = newConn
	conns.mtx.Unlock()
	return nil
}

func GetConnIDs() []string {
	IDs := []string{}
	for _, conn := range conns.connections {
		IDs = append(IDs, conn.PlayerID)
	}
	return IDs
}
