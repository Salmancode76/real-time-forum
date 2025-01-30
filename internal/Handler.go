package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

var data map[string]interface{}

func s_test(w http.ResponseWriter, r *http.Request) {
	// Check if the Accept header is set to application/json
	if r.Header.Get("Accept") == "application/json" {
		// If the Accept header is correct, return JSON data
		response := map[string]interface{}{
			"showRequestMade": true,
			"message":         "Yes, you can show content",
			"routeName":       "/s",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// Serve HTML for the /s route
		http.ServeFile(w, r, "./index.html")
	}
}
func getHome(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}

func postsign(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return
	}
	fmt.Println(data)

	response := map[string]interface{}{
		"status":  "success",
		"message": "Data received",
		"data":    data,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
	fmt.Println(response)

}

func postLogin(w http.ResponseWriter, r *http.Request) {
	err := json.NewDecoder(r.Body).Decode(&data)

	if err != nil {
		return
	}
	fmt.Println(data)

	type User struct {
		Uename   string `json:"uename"`
		Password string `json:"password"`
	}

	var user User

	if uename, ok := data["uename"].(string); ok {
		user.Uename = uename
	}
	if password, ok := data["password"].(string); ok {
		user.Password = password
	}

	response := map[string]interface{}{
		"status":  "success",
		"message": "Data received",
		"data":    user,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
	fmt.Println(response)

	fmt.Println(GenerateSessionID())
	Cookies(w, user.Uename)

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
		MaxAge:   900, // 15 minutes
		Expires:  expiration,
		HttpOnly: false,
		Path:     "/",
	}

	http.SetCookie(w, session_cookie)

	fmt.Println(expiration)
	fmt.Printf("New cookie set: %+v\n", session_cookie) // Log for debugging

}
