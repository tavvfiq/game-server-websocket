/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"game-server-websocket/internal/app"
	"game-server-websocket/internal/pkg/model"
	"game-server-websocket/internal/pkg/store"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var READ_BUFFER_SIZE = 1024
var WRITE_BUFFER_SIZE = 1024

var upgrader = websocket.Upgrader{
	ReadBufferSize:  READ_BUFFER_SIZE,
	WriteBufferSize: WRITE_BUFFER_SIZE,
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("pong"))
		})

		http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
			currentGorillaConn, err := upgrader.Upgrade(w, r, r.Header)
			if err != nil {
				http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
			}

			serverID := r.URL.Query().Get("serverID")
			playerID := r.URL.Query().Get("playerID")
			currentConn := model.WebSocketConnection{Conn: currentGorillaConn, ServerID: serverID, PlayerID: playerID}
			log.Printf("new player %s connected on server %s", playerID, serverID)
			store.AddConnections(&currentConn)
			go app.HandleIO(r.Context(), &currentConn)
		})

		fmt.Println("Server starting at :8080")
		http.ListenAndServe(":8080", nil)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
