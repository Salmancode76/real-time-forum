package internal

import (
	"net/http"
)

func (app *App) Routes() http.Handler {

	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/ws", handleWebSocket)

	mux.HandleFunc("/s", s_test)

	mux.HandleFunc("GET /sign", getHome)

	mux.HandleFunc("POST /sign", app.postsign)

	mux.HandleFunc("GET /login", getHome)

	mux.HandleFunc("POST /login", app.postLogin)

	mux.HandleFunc("/", getHome)

	return mux
}
