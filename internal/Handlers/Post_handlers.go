package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"real-time-forum/internal/models"
	"real-time-forum/internal/models/entities"
)

var data map[string]interface{}

func CreateComment (app * models.App) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var comment entities.Comment
		json.NewDecoder(r.Body).Decode(&comment)

			sCookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unable to retrieve Cookie", http.StatusUnauthorized)
			SendResponse(w, "Valdate Cooke","Unable to retrieve Cookie",false, http.StatusUnauthorized)

			return
		}

		IDCookie, err := r.Cookie("userID")
		if err != nil {
			http.Error(w, "Unable to retrieve Cookie", http.StatusUnauthorized)
		SendResponse(w, "Valdate Cooke","Unable to retrieve Cookie",false, http.StatusUnauthorized)

			return
		}

		if app.Session[IDCookie.Value] != sCookie.Value {
			http.Error(w, "User not logged in", http.StatusUnauthorized)
			SendResponse(w, "Valdate Cooke","User not logged in",false, http.StatusUnauthorized)

			return
		}

		if app.Session[IDCookie.Value]!=sCookie.Value{
		http.Error(w, "User not logged in", http.StatusUnauthorized)
		return	
		}
		comment.UserID = IDCookie.Value


		app.Posts.InsertComment(comment.UserID,comment.Comment,comment.Postid)
	}
}


func CreatePost(app *models.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post entities.Post
		err := json.NewDecoder(r.Body).Decode(&post)

		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}
		sCookie, err := r.Cookie("session")
		if err != nil {
			http.Error(w, "Unable to retrieve Cookie", http.StatusUnauthorized)

			return
		}

		IDCookie, err := r.Cookie("userID")
		if err != nil {
			http.Error(w, "Unable to retrieve Cookie", http.StatusUnauthorized)

			return
		}

		if app.Session[IDCookie.Value] != sCookie.Value {
			http.Error(w, "User not logged in", http.StatusUnauthorized)
			return
		}
		post.UserID = IDCookie.Value

		if err != nil {
			log.Print("ID Invalid")
		}

		/*
					categories,err := app.Posts.GetAllCategories()

			for id,namr := range categories{
				fmt.Println(id,"   ", namr)
			}
		*/
		app.Posts.Insert(post.UserID, post.Title, post.Content, post.Categories)

		//post.User = r.
		log.Printf("POSTS ", post)
	}
}



func ViewPost (app * models.App)http.HandlerFunc{
		return func(w http.ResponseWriter, r *http.Request) {
		var Post entities.Post
		id := r.URL.Query().Get("id")
		log.Println(id)
		Post,err :=app.Posts.FindPost(id)


		if err!=nil{
			http.Redirect(w,r,"/p",200)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
				
	//	log.Print(Post)
				var Comments []entities.Comment
		Comments,err = app.Posts.GetPostComment(id)
		if err!=nil{
			log.Println(err)
		}
		Post.Comments = Comments
		log.Println("Comments ",Comments)
	
	   if r.Header.Get("Accept") == "application/json" {
            posts := []entities.Post{Post}
            SendResponse(w, "Fetch Post", "", true, http.StatusOK, posts)
            return
        }
			http.ServeFile(w, r, "./index.html")
	



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


