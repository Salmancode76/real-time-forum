package handlers

import (
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/models"
	"time"

	"github.com/gofrs/uuid"
)



func GetHome(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

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
func GetDashboard(w http.ResponseWriter, r *http.Request) {

    user := r.Context().Value(contextKeyUser)
    if user == nil {
	 //w.WriteHeader(http.StatusUnauthorized)
	//w.WriteHeader(http.StatusSeeOther)
	//http.Redirect(w, r, "/login", http.StatusSeeOther)
	
	http.Redirect(w, r, "/login", http.StatusSeeOther) // 303 Redirect

	//w.WriteHeader(http.StatusUnauthorized)
	//w.Write([]byte("Unauthorized. Please log in."))
//http.Redirect(w, r, "/login", http.StatusSeeOther)
        //SendResponse(w, "GetDashboard", "User not logged in", false, http.StatusUnauthorized)

		return
    } else {
        http.ServeFile(w, r, "./index.html")
    }
}

func GenerateSessionID() string {
	sessionID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	return sessionID.String()
}

func Cookies(w http.ResponseWriter, userID string) string {
	expiration := time.Now().Add(24 * time.Hour)
	sessionID := GenerateSessionID()

	sessionCookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // Set to true if using HTTPS
	}

	userIDCookie := &http.Cookie{
		Name:     "userID",
		Value:    userID,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
		Secure:   false, // Set to true if using HTTPS
	}

	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, userIDCookie)
	return sessionID
}
