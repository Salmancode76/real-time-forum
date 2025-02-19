package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)


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
		Secure:   false, 
	}

	userIDCookie := &http.Cookie{
		Name:     "userID",
		Value:    userID,
		Expires:  expiration,
		HttpOnly: true,
		Path:     "/",
		Secure:   false, 
	}

	http.SetCookie(w, sessionCookie)
	http.SetCookie(w, userIDCookie)
	return sessionID
}
