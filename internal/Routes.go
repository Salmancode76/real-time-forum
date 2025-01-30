package server

import "net/http"

func Routes() http.Handler {

	mux := http.NewServeMux()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// WebSocket route
	mux.HandleFunc("/ws", handleWebSocket)

	// Other routes
	mux.HandleFunc("/s", s_test)

	mux.HandleFunc("GET /sign", getHome)

	mux.HandleFunc("POST /sign", postsign)

	mux.HandleFunc("GET /login", getHome)

	mux.HandleFunc("POST /login", postLogin)

	// Serve HTML for the root route
	mux.HandleFunc("/", getHome)

	return mux
}
