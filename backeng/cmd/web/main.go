package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx"
	"github.com/joho/godotenv"

	"yacode.malafeev.mikhail/internal/models"
	"yacode.malafeev.mikhail/internal/pkg/helpers"
)


type application struct{
    logger *slog.Logger
    rooms *models.RoomModel
}


func main() {

    // Создание своего логгера
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

    // Подключение(пинг) БД
    db, err := openDB()
    if err != nil{
        logger.Error(err.Error())
        os.Exit(1)
    }
    defer db.Close()


    app := &application{
        logger: logger,
        rooms: &models.RoomModel{DB : db},
    }

    logger.Info("WebSocket server started on :8080")
    
    err = http.ListenAndServe(":8080", app.routes())
    logger.Error(err.Error())
    os.Exit(1)
}

func openDB() (c *pgx.Conn, err error){

    _  = godotenv.Load();

    cfg := &pgx.ConnConfig{
        Host:      helpers.GetEnv("DB_HOST", ""),
        Port:     helpers.GetEnvAsInt("DB_PORT", 0),
        Database: helpers.GetEnv("DB_DATABASE", ""),
        User:     helpers.GetEnv("DB_USER", ""),
        Password: helpers.GetEnv("DB_PASSWORD", ""),
    }
    

    db, err := pgx.Connect(*cfg)
	if err != nil {
        db.Close()
        return nil, err
    }

    var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
        db.Close()
        return nil, err
    }
    return db, nil
}

