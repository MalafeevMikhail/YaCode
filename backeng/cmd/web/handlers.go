package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"yacode.malafeev.mikhail/internal/models"
)

func(app *application) roomCreatePost(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        Name string `json:"name"`
    }

    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestData); err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    
    name := requestData.Name
	if name == "" {
		app.clientError(w, http.StatusBadRequest)
		return
	}

    id, err := app.rooms.Create(name)
    if err != nil {
        app.serverError(w, r, err)
    }

    response := map[string]string{
        "id": id,
    }
    
    json.NewEncoder(w).Encode(response)
}

func (app *application) checkRoomPost(w http.ResponseWriter, r *http.Request) {
    var requestData struct {
        RoomId string `json:"roomId"`
    }
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&requestData); err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    id := requestData.RoomId
    if id == "" {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    if err := uuid.Validate(id); err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }
    response := map[string]bool{
        "isValid" : false,
    }
    isValid, err := app.rooms.CheckRoom(id)
    if err != nil {
        if errors.Is(err, models.ErrNoRecod){
            response["isValid"]= false
        }else{
            app.serverError(w, r, err)
            return
        }
    }else{
        response["isValid"] = isValid
    }
    
    json.NewEncoder(w).Encode(response)
}


func(app *application) roomGet(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if id == "" {
        http.NotFound(w,r)
        return
    }

    if err := uuid.Validate(id); err != nil {
        app.clientError(w, http.StatusBadRequest)
        return
    }

    room, err := app.rooms.Get(id)
    if err != nil{
        if errors.Is(err, models.ErrNoRecod){
            http.NotFound(w, r)
        }else{
            app.serverError(w, r, err)
        }
        return
    }

    json.NewEncoder(w).Encode(room)
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