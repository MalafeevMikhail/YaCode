package models

import (
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgtype"
)


type Room struct {
	ID string
	Name string
	Created time.Time
    Code string  
    Language string
}

type RoomModel struct{
	DB *pgx.Conn
}

func(m *RoomModel) Create(name string)(string, error){
    stmt := `INSERT INTO Rooms(name, language) VALUES($1, $2) RETURNING id;`

    var id string

    err := m.DB.QueryRow(stmt, name, "go").Scan(&id)

    if err != nil {
        return "", err
    }

    return id, nil
}

func(m *RoomModel) Get(id string)(Room, error){
    stmt := `SELECT id, name, create_at, code, language FROM Rooms WHERE id = $1;`
    
    var room Room
    var code pgtype.Varchar

    err := m.DB.QueryRow(stmt, id).Scan(&room.ID, &room.Name, &room.Created, &code, &room.Language)
    
    if err != nil {
        if err == pgx.ErrNoRows {
            return Room{}, ErrNoRecod
        }
        return Room{}, err
    }
    
    if code.Status == pgtype.Present{
        room.Code = code.String
    }
    
    return room, nil

}