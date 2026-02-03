package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true }, // Allow all connections
}

var clients = make(map[*websocket.Conn]bool) // Track active clients

func(app *application) routes() http.Handler {

	mux := http.NewServeMux()


	mux.HandleFunc("POST /create-room",  app.roomCreatePost)
    mux.HandleFunc("/ws", handleConnections)

	return mux
}