package app

import (
	"context"
	"encoding/json"
	"fmt"
	"game-server-websocket/internal/pkg/event"
	"game-server-websocket/internal/pkg/model"
	"game-server-websocket/internal/pkg/store"
	"log"
	"strings"
)

func HandleIO(ctx context.Context, currentConn *model.WebSocketConnection) {
	for {
		defer func() {
			if r := recover(); r != nil {
				log.Println("ERROR", fmt.Sprintf("%v", r))
			}
		}()
		payload := model.SocketPayload{}
		err := currentConn.ReadJSON(&payload)
		if err != nil {
			if strings.Contains(err.Error(), "websocket: close") {
				log.Println("disconnect")
				event.PlayerDisconnectedEventHandler(ctx, currentConn.ServerID, currentConn.PlayerID)
				store.RemoveConnection(currentConn)
				return
			}
			log.Println("ERROR", err.Error())
			return
		}
		player := model.Player{}
		json.Unmarshal(payload.Data, &player)
		log.Println(payload.EventType)
		switch payload.EventType {
		case event.NEW_CONNECTION:
			event.NewPlayerEventHandler(ctx, currentConn.ServerID, player)
		case event.STATE_UPDATE:
			event.StateUpdateEventHandler(ctx, currentConn.ServerID, player)
		}
	}
}
