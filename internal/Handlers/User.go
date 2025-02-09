package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/models"
	"real-time-forum/internal/models/entities"
	"strconv"
)

var data map[string]interface{}



// PostSign is the handler for the POS
func PostSign(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.UserData
		var UsernameUnique bool
		var EmailUnique bool

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		log.Printf("Received user: %+v", user)

		// Process the age field to ensure it's an integer
		intAge, err := strconv.Atoi(user.Age)
		if err != nil {
			SendResponse(w, "Sign up", "Invalid age", false, http.StatusBadRequest)
			return
		}
		UsernameUnique, EmailUnique, err = app.Users.IsUnique(user.Username, user.Email)

		fmt.Println(UsernameUnique, EmailUnique)
		if !UsernameUnique {
			SendResponse(w, "Sign up", "Username is NOT UNIQUE!", false, http.StatusBadRequest)
			return
		}
		if !EmailUnique {
			SendResponse(w, "Sign up", "Email is NOT UNIQUE!", false, http.StatusBadRequest)
			return
		}
		if err != nil {
			SendResponse(w, "Sign up", "Failed to check uniqueness", false, http.StatusBadRequest)
			return
		}

		err = app.Users.Insert(user.Username, user.Email, user.Password, user.Gender, user.FirstName, user.LastName, intAge)
		if err != nil {
			SendResponse(w, "Sign up", "Failed to insert user", false, http.StatusBadRequest)

			return
		}
		SendResponse(w, "Sign up", "User inserted successfully", true, http.StatusOK)

	}
}

func PostLogin(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			SendResponse(w, "Login", "Invalid data", false, http.StatusBadRequest)
			return
		}
		fmt.Println(data)

		var user entities.UserData

		user, err = app.Users.Auth(data["uename"].(string), data["password"].(string))
		if err != nil {
			SendResponse(w, "Login", err.Error(), false, http.StatusBadRequest)
			return
		}

		id := strconv.Itoa(user.UserID)

		app.Session[id] = Cookies(w, id)
		app.UserID[id] = user.Username

		SendResponse(w, "Login", "User authenticated", true, http.StatusOK)

		for _, cookie := range w.Header()["Set-Cookie"] {
			log.Printf("Setting cookie: %s", cookie)
		}
	}
}