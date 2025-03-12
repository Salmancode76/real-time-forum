package router

import (
	"net/http"
	handlers "real-time-forum/internal/Handlers"
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

	mux.HandleFunc("GET /chat", handlers.GetHome)

	mux.HandleFunc("Post /chat", handlers.Chat(app.App))

	mux.HandleFunc("/s", handlers.S_test)

	mux.HandleFunc("GET /sign", handlers.GetHome)

	mux.HandleFunc("POST /sign", handlers.PostSign(app.App))

	mux.HandleFunc("GET /login", handlers.GetHome)

	mux.HandleFunc("POST /login", handlers.PostLogin(app.App))

	mux.HandleFunc("/createPost", handlers.GetHome)

	mux.HandleFunc("POST /createPost", handlers.CreatePost(app.App))

	mux.HandleFunc("POST /createComment", handlers.CreateComment(app.App))

	mux.HandleFunc("/api/posts", handlers.GetAllPosts(app.App))

	mux.HandleFunc("/logout", handlers.Logout(app.App))

	mux.HandleFunc("/post", handlers.ViewPost(app.App))

	mux.Handle("/auth-check", handlers.Authorized(app.App))

	mux.Handle("/fetchPost", handlers.FetchCategory(app.App))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			handlers.Lost404(w, r)

		} else {
			handlers.GetHome(w, r)
		}
	})

	return handlers.MiddleWare(mux, app.App)
}
