package models

import (
	"time"

	"github.com/jackc/pgx"
)


type Room struct {
	ID string
	Name string
	Created time.Time
}

type RoomModel struct{
	DB *pgx.Conn
}


func(m *RoomModel) Get(id string)(Room, error){
    stmt := `SELECT id, name, create_at FROM Rooms WHERE id = $1`
    
    var room Room
    
    err := m.DB.QueryRow( stmt, id).Scan(&room.ID, &room.Name, &room.Created)
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return Room{}, ErrNoRecod
        }
        return Room{}, err
    }
    
    return room, nil

}