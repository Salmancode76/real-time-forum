package handlers

import (
	"fmt"
	"net/http"
)

func authenticate(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sCookie, err := r.Cookie("session")
		if err != nil {
			next.ServeHTTP(w, r)
			return
			fmt.Println(sCookie)
		}
	})
}
