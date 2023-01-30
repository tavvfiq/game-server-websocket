package event

import (
	"context"
	"fmt"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/store"
)

func PlayerDisconnectedEventHandler(ctx context.Context, serverID string, playerID string) error {
	store.RemovePlayer(playerID)
	message := fmt.Sprintf("player %s disconnected from server", playerID)
	return broadcastMessage(PLAYER_DISCONNECT, serverID, playerID, message)
}
