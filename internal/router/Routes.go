package router

import (
	"net/http"
	"real-time-forum/internal/handlers"
	"real-time-forum/internal/models"
)

type GlobalApp struct {
	*models.App
}

func (app *GlobalApp) Routes() http.Handler {

	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/ws", handlers.HandleWebSocket)

	mux.HandleFunc("/s", handlers.S_test)

	mux.HandleFunc("GET /sign", handlers.GetHome)

	mux.HandleFunc("POST /sign", handlers.PostSign(app.Users))

	mux.HandleFunc("GET /login", handlers.GetHome)

	mux.HandleFunc("POST /login", handlers.PostLogin(app.Users))

	mux.HandleFunc("/", handlers.GetHome)

	return mux
}
