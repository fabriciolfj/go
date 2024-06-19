package main

import (
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
		return
	}

	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()

		if err != nil {
			break
		}

		if err := conn.WriteMessage(messageType, message); err != nil {
			break
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/ws/ echo", echoHandler)
}
