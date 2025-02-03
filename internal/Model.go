package internal

import (
	"database/sql"
	"net/http"
	"real-time-forum/internal/handlers"
)

type User struct {
	Username  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type App struct {
	Server *Server
	Users  *handlers.UserModel
	DB     *sql.DB
}
type Server struct {
	HTTP *http.Server
}
