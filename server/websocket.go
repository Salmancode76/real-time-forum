package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Form(w http.ResponseWriter, r *http.Request) {

	fileName := "./HTML/index.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		log.Fatal("Erro is ", err)
		return
	}
	error := t.ExecuteTemplate(w, "index.html", nil)

	if error != nil {
		log.Fatal(error)
	}
}
func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	//form(w, r)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	//handle websocket message
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var myMessage MyMessage
		json.Unmarshal(data, &myMessage)
		//fmt.Println(myMessage)
		handleWebSocketMessage(conn, myMessage)

	}
}

func handleMessageMessage(conn *websocket.Conn, message MyMessage) {
	db := OpenDatabase()
	defer db.Close()
	To := GetUserID(db, message.To)
	From := GetUserID(db, message.From)
	fmt.Println(message.Text)
	AddMessageToHistory(From, To, message.Text)
}

func handleGetUsersMessage(conn *websocket.Conn) {
	db := OpenDatabase()
	defer db.Close()
	users := getAllUsers(db)
	//fmt.Println(users)
	var allUsers []ServerUser
	for _, i := range users {
		allUsers = append(allUsers, ServerUser{Name: i})
	}
	message := ServerMessage{Type: "users", Users: allUsers}
	conn.WriteJSON(message)
	//fmt.Println(message)
}
func handleGetChatHistoryMessage(conn *websocket.Conn, m MyMessage) {
	db := OpenDatabase()
	defer db.Close()
	To := GetUserID(db, m.To)
	From := GetUserID(db, m.From)

	messages := GetChatHistory(To, From, 0)

	message := ServerMessage{Type: "oldmessages", ChatHistory: messages}
	conn.WriteJSON(message)
	//fmt.Println(message)

}

// function to handle the messages received from the client
func handleWebSocketMessage(conn *websocket.Conn, message MyMessage) {
	//fmt.Println("Message received: ", message.Type)

	switch message.Type {
	case "message":
		handleMessageMessage(conn, message)
	case "get_users":
		handleGetUsersMessage(conn)
	case "get_chat_history":
		handleGetChatHistoryMessage(conn, message)

	}
}

// // function to handle the messages received from the client
// func handleWebSocketMessage(conn *websocket.Conn, message ServerMessage) {
// 	//fmt.Println("Message received: ", message.Type)

// 	switch message.Type {
// 	case "new_user":
// 		handleNewUserMessage(message)
// 	case "create_category":
// 		handleNewCategoryMessage(message)
// 	case "new_post":
// 		handleNewPostMessage(message)
// 	case "get_posts":
// 		handleGetPostsMessage(conn, message)
// 	case "get_chat_history":
// 		handleGetChatHistoryMessage(conn, message)
// 	case "message":
// 		handleMessageMessage(conn, message)
// 	case "login":
// 		handleLoginMessage(conn, message)
// 	case "loginResponse":
// 		handleLoginResponseMessage(conn, message)
// 	case "logout":
// 		handleLogoutMessage(conn, message)
// 	case "register":
// 		handleRegisterMessage(conn, message)
// 	case "registerResponse":
// 		handleRegisterResponseMessage(conn, message)
// 	case "get_categories":
// 		handleGetCategoriesMessage(conn, message)
// 	case "get_comments":
// 		handleGetCommentsMessage(conn, message)
// 	case "new_comment":
// 		handleNewCommentMessage(conn, message)
// 	case "get_users":
// 		handleGetUsersMessage(conn, message)
// 	case "get_offline_users":
// 		handleGetOfflineUsersMessage(conn, message)
// 	case "typing":
// 		handleTypingMessage(conn, message)
// 	case "stopTyping":
// 		handleStopTypingMessage(conn, message)
// 	case "postsByCategory":
// 		handleGetPostsForCategory(conn, message)
// 	}
// }
