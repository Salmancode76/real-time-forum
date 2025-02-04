package handlers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

func GetHome(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}

func GenerateSessionID() string {
	sessionID, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}

	return sessionID.String()
}

func Cookies(w http.ResponseWriter, uename string) {

	expiration := time.Now().Add(24 * time.Hour)
	sessionID := GenerateSessionID()

	session_cookie := &http.Cookie{
		Name:     "session",
		Value:    sessionID,
		MaxAge:   900,
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}

	http.SetCookie(w, session_cookie)

	fmt.Println(expiration)
	fmt.Printf("New cookie set: %+v\n", session_cookie)

}
