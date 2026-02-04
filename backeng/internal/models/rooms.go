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

func(m *RoomModel) Create(name string)(string, error){
    stmt := `INSERT INTO Rooms(name) VALUES($1) RETURNING id;`

    var id string

    err := m.DB.QueryRow(stmt, name).Scan(&id)

    if err != nil {
        return "", err
    }

    return id, nil
}

func(m *RoomModel) Get(id string)(Room, error){
    stmt := `SELECT id, name, create_at FROM Rooms WHERE id = $1;`
    
    var room Room
    
    err := m.DB.QueryRow(stmt, id).Scan(&room.ID, &room.Name, &room.Created)
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return Room{}, ErrNoRecod
        }
        return Room{}, err
    }
    
    return room, nil

}

func (m *RoomModel) CheckRoom(id string) (bool, error) {
    stmt := "SELECT id, name, create_at FROM Rooms WHERE id = $1;"


    var room Room

    err := m.DB.QueryRow(stmt, id).Scan(&room.ID, &room.Name, &room.Created)

    if err != nil{
        if err == pgx.ErrNoRows {
            return false, ErrNoRecod
        }
        return false, err
    }

    return true, nil
}