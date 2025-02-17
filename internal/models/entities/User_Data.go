package entities

import "time"

type UserData struct {
	UserID    int    `json:"userId"`
	Username  string `json:"username"`
	Age       string `json:"age"`
	Gender    string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
type Comment struct{
	ID 		string
	Postid string `json:"PostID"`
	Comment string `json:"content"`
	UserID string
	Username string
	Date       time.Time `json:"date"` 
}



type Post struct {
    ID         string    `json:"id"`
    UserID     string    `json:"userId"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    Date       time.Time `json:"date"`  
    Categories []string  `json:"categories"`
	//Category string `json:"categorie"`
	Comments 	[]Comment 
}
type Category struct{
	Id string
	Name string

}