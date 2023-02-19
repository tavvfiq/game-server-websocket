package event

import (
	"context"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/store"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/utils"
)

func NewPlayerEventHandler(ctx context.Context, serverID string, player model.Player) error {
	store.AddPlayer(player)
	IDs := utils.FilterOutString(store.GetConnIDs(), player.ID)
	go broadcastMessage(NEW_CONNECTION, serverID, IDs, player)
	go broadcastMessage(SYNC_STATE, serverID, []string{player.ID}, store.GetPlayers())
	return nil
}
