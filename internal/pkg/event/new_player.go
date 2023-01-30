package event

import (
	"context"
	"fmt"
	"game-server-websocket/internal/pkg/model"
	"game-server-websocket/internal/pkg/store"
)

func NewPlayerEventHandler(ctx context.Context, serverID string, player model.Player) error {
	store.AddPlayer(player)
	message := fmt.Sprintf("new player %s connected to server", player.ID)
	return broadcastMessage(NEW_CONNECTION, serverID, player.ID, message)
	// go broadcastMessage(STATE_UPDATE, serverID, "", store.GetPlayers())
	// return nil
}
