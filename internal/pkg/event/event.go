package event

import (
	"encoding/json"
	"game-server-websocket/internal/pkg/model"
	"game-server-websocket/internal/pkg/store"
)

func broadcastMessage(eventType, serverID, playerID string, payload interface{}) error {
	var errs error
	for _, conn := range store.GetConnections() {
		if conn.ServerID != serverID || conn.PlayerID == playerID {
			continue
		}
		b, _ := json.Marshal(payload)
		err := conn.WriteJSON(model.SocketResponse{
			EventType: eventType,
			Data:      b,
		})
		if err != nil {
			errs = err
		}
	}
	return errs
}
