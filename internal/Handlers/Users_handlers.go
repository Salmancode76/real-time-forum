package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"real-time-forum/internal/models"
	"real-time-forum/internal/models/entities"
	"strconv"
)

func GetHome(w http.ResponseWriter, r *http.Request) {

	if _, err := os.Stat("./index.html"); os.IsNotExist(err) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`500 Internal server error`))
		return
	}
	http.ServeFile(w, r, "./index.html")

}

func Chat(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		HandleWebSocket(app, w, r)
	}
}

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

		app.Session[id] = Cookies(w, id, user.Username)
		app.UserID[id] = user.Username
		UpdateOnlineUsers(app)

		SendResponse(w, "Login", "User authenticated", true, http.StatusOK)

		for _, cookie := range w.Header()["Set-Cookie"] {
			log.Printf("Setting cookie: %s", cookie)
		}
	}
}

func Logout(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(contextKeyUser).(entities.UserData)

		//fmt.Println(user)
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		tempID := strconv.Itoa(user.UserID)

		delete(app.Session, tempID)
		delete(app.UserID, string(user.UserID))
		w.WriteHeader(http.StatusNoContent)
		UpdateOfflineUsers(app,user.Username)

	}
}

func Lost404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte(`
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<meta name="viewport" content="width=device-width, initial-scale=1.0">
				<link rel="stylesheet" href="/static/styles/sign_up.css">
				<link rel="stylesheet" href="/static/styles/login.css"> 
				<link rel="stylesheet" href="/static/styles/nav.css"> 
				<link rel="stylesheet" href="/static/styles/create_post.css">
				<link rel="stylesheet" href="/static/styles/home.css">
				<link rel="stylesheet" href="/static/styles/view_post.css">
				<link rel="stylesheet" href="/static/styles/errors.css">



				<title>Community Forum</title>
			</head>
			<body>
				<nav id="nav">
					<ul>
						<li class="header">  <a href="/" onclick="navigateTo('/');">  Community Forum </a> </li>
						<div class="nav-links">
							<li><a <a href="/sign" onclick="navigateTo('/sign');"> Sign-Up</a></li>
							<li><a <a href="/login" onclick="navigateTo('/login');"> Login</a></li>

						</div>
						<img id="hamICON" src="/static/images/ham_menu.svg" alt="Menu">
					</ul>
				</nav>
				<div id="app">
				 
        	  </div>
				</div>
				<script type="module" src="./static/index.js"></script>
			</body>
			</html>
						
			
			
			`))
}
