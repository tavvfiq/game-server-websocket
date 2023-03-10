package event

import (
	"context"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/store"
	"github.com/tavvfiq/game-server-websocket/internal/pkg/utils"
)

func StateUpdateEventHandler(ctx context.Context, serverID string, player model.Player) error {
	currentState := store.GetPlayer(player.ID)
	newState := model.PlayerState{
		X:      currentState.State.X + player.State.DeltaX,
		Y:      currentState.State.Y + player.State.DeltaY,
		DeltaX: player.State.DeltaX,
		DeltaY: player.State.DeltaY,
	}
	currentState.State = newState
	store.UpdatePlayer(currentState)
	Ids := utils.FilterOutString(store.GetConnIDs(), player.ID)
	return broadcastMessage(STATE_UPDATE, serverID, Ids, currentState)
}
