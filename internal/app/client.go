package app

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"game-server-websocket/internal/pkg/event"
	"game-server-websocket/internal/pkg/model"
	"game-server-websocket/internal/pkg/store"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var addr = "localhost:8080"

func RunClient(ctx context.Context, serverID, playerID string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: addr, Path: "/ws", RawQuery: fmt.Sprintf("serverID=%s&playerID=%s", serverID, playerID)}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()
	player := model.Player{
		ID: playerID,
	}
	b, _ := json.Marshal(player)
	payload := model.SocketPayload{
		EventType: event.NEW_CONNECTION,
		Data:      b,
	}
	err = c.WriteJSON(payload)
	if err != nil {
		log.Fatal("error on connecting to the server: ", err)
		return
	}

	done := make(chan struct{})

	state := make(chan model.PlayerState)

	go func() {
		defer close(done)
		for {
			resp := model.SocketResponse{}
			err := c.ReadJSON(&resp)
			if err != nil {
				log.Println("read:", err)
				return
			}
			switch resp.EventType {
			case event.NEW_CONNECTION:
				log.Println("new player connected")
			case event.PLAYER_DISCONNECT:
				log.Println("a player disconnected")
			case event.STATE_UPDATE:
				player := model.Player{}
				json.Unmarshal(resp.Data, &player)
				store.UpdatePlayer(&player)
				log.Printf("%v", store.GetPlayers())
			}
		}
	}()

	go func() {
		for {
			reader := bufio.NewReader(os.Stdin)
			cmd, _ := reader.ReadString('\n')
			if strings.Contains(cmd, "move") {
				log.Println("move")
				delta := model.PlayerState{}
				segment := strings.Split(strings.Trim(cmd, "\n"), " ")
				direction := segment[1]
				value, err := strconv.ParseFloat(segment[2], 64)
				if err != nil {
					log.Println(err)
				}
				log.Println(direction, value)
				switch direction {
				case "up":
					delta.DeltaY = value
				case "down":
					delta.DeltaY = -value
				case "left":
					delta.DeltaX = -value
				case "right":
					delta.DeltaX = value
				}
				state <- delta
			}
		}
	}()

	for {
		select {
		case <-done:
			return
		case update := <-state:
			player := model.Player{
				ID:    playerID,
				State: update,
			}
			b, _ := json.Marshal(player)
			payload := model.SocketPayload{
				EventType: event.STATE_UPDATE,
				Data:      b,
			}
			log.Println(player)
			err := c.WriteJSON(payload)
			if err != nil {
				log.Printf("error on update player %s state: %v", playerID, err)
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}
}
