package main

import (
	"database/sql"
	"log"
	"net/http"
	internal "real-time-forum/Internal"
	"real-time-forum/Internal/Handlers"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db/db.db")
	if err != nil {
		log.Fatal("Something bad happened while opening the database:", err)
	}

	app := &internal.App{
		DB:    db,
		Users: &Handlers.UserModel{DB: db},
	}

	s := internal.Server{
		HTTP: &http.Server{
			Addr:         ":8080",
			Handler:      app.Routes(),
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}

	app.Server = &s

	app.Server.RunServer()
}
