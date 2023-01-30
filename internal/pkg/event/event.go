package event

import (
	"encoding/json"
	"strings"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/store"
)

func broadcastMessage(eventType, serverID string, IDs []string, payload interface{}) error {
	var errs error
	joined := strings.Join(IDs, ".")
	for _, conn := range store.GetConnections() {
		if conn.ServerID != serverID || !strings.Contains(joined, conn.PlayerID) {
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
