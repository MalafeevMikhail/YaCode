package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func(app *application) roomCreatePost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	name := r.PostForm.Get("name")
	if name == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

}


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