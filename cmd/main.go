package main

import (
	"database/sql"
	"log"
	"net/http"
	"real-time-forum/internal/models"
	"real-time-forum/internal/repository"
	"real-time-forum/internal/router"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//"./db/db.db"
	db, err := sql.Open("sqlite3", "./db/db.db")
	if err != nil {
		log.Fatal("Something bad happened while opening the database:", err)
	}

	app := &models.App{
		DB:      db,
		Users:   &repository.UserModel{DB: db},
		Posts :  &repository.PostModel{DB:db},
		Server:  &models.Server{},
		Session: make(map[string]string),
		UserID:  make(map[string]string),
	}
	GlobalApp := &router.GlobalApp{
		App: app,
	}
	s := models.Server{
		HTTP: &http.Server{
			Addr:         ":8080",
			Handler:      GlobalApp.Routes(),
			WriteTimeout: 10 * time.Second,
			ReadTimeout:  10 * time.Second,
		},
	}

	app.Server = &s

	log.Println("Starting server on", GlobalApp.App.Server.HTTP.Addr)
	if err := GlobalApp.App.Server.HTTP.ListenAndServe(); err != nil {
		log.Fatal("Server failed to start:", err)
	}

}
