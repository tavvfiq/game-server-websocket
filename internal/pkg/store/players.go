package store

import (
	"sync"

	"github.com/tavvfiq/game-server-websocket/internal/pkg/model"
)

type _players struct {
	mtx     sync.Mutex
	players map[string]*model.Player
}

var players = _players{
	players: make(map[string]*model.Player, 0),
}

func AddPlayer(player model.Player) {
	players.mtx.Lock()
	defer players.mtx.Unlock()
	players.players[player.ID] = &player
}

func RemovePlayer(playerID string) {
	players.mtx.Lock()
	defer players.mtx.Unlock()
	delete(players.players, playerID)
}

func UpdatePlayer(player *model.Player) {
	players.mtx.Lock()
	defer players.mtx.Unlock()
	players.players[player.ID] = player
}

func GetPlayer(playerID string) *model.Player {
	return players.players[playerID]
}

func GetPlayers() map[string]*model.Player {
	return players.players
}
