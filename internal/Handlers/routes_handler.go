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

func ViewPost (app * models.App)http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request) {
		var Post entities.Post
		id := r.URL.Query().Get("id")
		log.Println(id)
		Post,err :=app.Posts.FindPost(id)


		if err!=nil{
			log.Println("Error :", err,Post)
		}
				
		log.Print(Post)
		
	   if r.Header.Get("Accept") == "application/json" {
            posts := []entities.Post{Post}
            SendResponse(w, "Fetch Post", "", true, http.StatusOK, posts)
            return
        }
			http.ServeFile(w, r, "./index.html")
	
		}
}

func Logout(app *models.App) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(contextKeyUser).(entities.UserData)

		fmt.Println(user)
		if !ok{
            http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		tempID := strconv.Itoa(user.UserID)

		delete(app.Session,tempID)
		delete(app.UserID,user.Username)
		 w.WriteHeader(http.StatusNoContent) 
		

	}
}


func GetAllPosts(app *models.App) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        posts, err := app.Posts.FetchAllPost()
        if err != nil {
            log.Println("Can't get posts ", err)
            http.Error(w, "Failed to fetch posts", http.StatusInternalServerError)
            return
        }
		//fmt.Println(posts)
		SendResponse(w,"Fetch Posts","",true,http.StatusOK,posts)

    }
}


func GetHome(w http.ResponseWriter, r *http.Request) {

	http.ServeFile(w, r, "./index.html")

}


func CreatePost(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post entities.Post
		err:= json.NewDecoder(r.Body).Decode(&post)

		if err !=nil{
				http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		sCookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
			return
		}

		IDCookie, err := r.Cookie("userID")
		if err != nil {
			http.Error(w, "Unable to retrieve session", http.StatusInternalServerError)
			return
		}

		if app.Session[IDCookie.Value]!=sCookie.Value{
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return	
		}
		post.UserID = IDCookie.Value

		if err !=nil{
			log.Print("ID Invalid")
		}


		/*
				categories,err := app.Posts.GetAllCategories()

		for id,namr := range categories{
			fmt.Println(id,"   ", namr)
		}
		*/		
		app.Posts.Insert(post.UserID,post.Title,post.Content,post.Categories)
		
		//post.User = r.
		log.Printf("POSTS ", post)
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

		app.Session[id] = Cookies(w, id)
		app.UserID[id] = user.Username

		SendResponse(w, "Login", "User authenticated", true, http.StatusOK)

		for _, cookie := range w.Header()["Set-Cookie"] {
			log.Printf("Setting cookie: %s", cookie)
		}
	}
}