package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/models"
	"strconv"
)
type contextKey string

const contextKeyUser = contextKey("user")

func MiddleWare(next http.Handler, app *models.App) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		SCookie, err := r.Cookie("session")
		if err != nil {
			log.Println("No cookie")
			next.ServeHTTP(w, r)
			return
		}

		//log.Println("cookie", SCookie.Value)
		IDcookie, err := r.Cookie("userID")
		if err != nil {
			log.Println("No ID cookie")
			next.ServeHTTP(w, r)
			return
		}


        //log.Println(app.Session[IDcookie.Value])

		if app.Session[IDcookie.Value] != SCookie.Value {
            log.Println("Not a valid session")
			next.ServeHTTP(w, r)
			return
		}

		id, err := strconv.Atoi(IDcookie.Value)
		if err != nil {
			log.Println("Error converting cookie to int:", err)
			next.ServeHTTP(w, r)
			return
		}
		user, err := app.Users.GetUserByID(id)
		if err != nil {
			log.Println("Error retrieving user:", err)
			next.ServeHTTP(w, r)
			return
		}
		
		ctx := context.WithValue(r.Context(), contextKeyUser, user)
				//log.Println(user)
		next.ServeHTTP(w, r.WithContext(ctx))


	})
}

func Authorized(app *models.App) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
    user := r.Context().Value(contextKeyUser)
	if user == nil {

		w.WriteHeader(http.StatusSeeOther)

		fmt.Println("Unauthorized")


		return

	} else {
		fmt.Println("Authorized")
		http.Redirect(w, r, "/", http.StatusOK) 
		
	}

}
}