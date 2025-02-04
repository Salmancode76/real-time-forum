package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"real-time-forum/internal/models/entities"
	"real-time-forum/internal/repository"
	"strconv"
)

var data map[string]interface{}

// PostSign is the handler for the POST /sign route
func PostSign(userModel *repository.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user entities.UserData
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		log.Printf("Received user: %+v", user)

		// Process the age field to ensure it's an integer
		intAge, err := strconv.Atoi(user.Age)
		if err != nil {
			http.Error(w, "Invalid age", http.StatusBadRequest)
			return
		}

		err = userModel.Insert(user.Username, user.Email, user.Password, user.Gender, user.FirstName, user.LastName, intAge)
		if err != nil {
			http.Error(w, "Failed to insert user", http.StatusInternalServerError)
			return
		}

		// Send the response
		response := map[string]interface{}{
			"status":  "success",
			"message": "User created successfully",
			"data":    user,
			"":        "",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func PostLogin(userModel *repository.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := json.NewDecoder(r.Body).Decode(&data)

		if err != nil {
			return
		}
		fmt.Println(data)

		type Userz struct {
			Uename   string `json:"uename"`
			Password string `json:"password"`
		}

		var user Userz

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
}
