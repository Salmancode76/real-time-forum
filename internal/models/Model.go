package models

import (
	"database/sql"
	"net/http"
	"real-time-forum/internal/repository"
)

type UserModel struct {
	DB *sql.DB
}

type Server struct {
	HTTP *http.Server
}

type App struct {
	Server *Server
	Users  *repository.UserModel
	DB     *sql.DB
}

type UserData struct {
	UserID    int    `json:"userId"`
	Username  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
