package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true }, // Allow all connections
}
var clients = make(map[*websocket.Conn]bool) // Track active clients

func handleConnections(w http.ResponseWriter, r *http.Request) {
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
	
    defer ws.Close()

	clients[ws] = true

    for {
        _, msg, err := ws.ReadMessage()
        if err != nil {
            fmt.Println("read error:", err)
            delete(clients, ws)
            break
        }

        for client := range clients {
            if err := client.WriteMessage(websocket.TextMessage, msg); err != nil {
                fmt.Println("broadcast error:", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

func main() {
    http.HandleFunc("/ws", handleConnections)

    fmt.Println("WebSocket server started on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("ListenAndServe:", err)
    }
}