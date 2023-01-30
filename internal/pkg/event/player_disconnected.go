package event

import (
	"context"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/store"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/utils"
)

func PlayerDisconnectedEventHandler(ctx context.Context, serverID string, playerID string) error {
	store.RemovePlayer(playerID)
	p := model.Player{
		ID: playerID,
	}
	IDs := utils.FilterOutString(store.GetConnIDs(), playerID)
	return broadcastMessage(PLAYER_DISCONNECT, serverID, IDs, p)
}
