package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"real-time-forum/internal/models"
	"time"

	"github.com/gorilla/websocket"
)

var allUsers []ServerUser
var NotUsers []ServerUser
var sockets = make(map[string]websocket.Conn)

var userSockets *map[string]websocket.Conn = &sockets

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development only)
	},
}

func HandleWebSocket(app *models.App, w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("userID")
	if err != nil {
		log.Println("No cookie")
		return
	} else {
		fmt.Println("cookie name =", cookie.Name)
		fmt.Println("cookie value =", cookie.Value)
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}

	(*userSockets)[cookie.Value] = *conn
	// Configure ping/pong handlers
	conn.SetPingHandler(func(string) error {
		log.Println("Received ping, sending pong")
		return conn.WriteControl(websocket.PongMessage, []byte{}, time.Now().Add(time.Second))
	})

	// Send periodic pings
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			select {
			case <-ticker.C:
				if err := conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Second)); err != nil {
					log.Printf("Ping failed: %v", err)

					return
				}
			case <-done:
				return
			}
		}
	}()

	// Message handling loop
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("Connection closed unexpectedly: %v", err)
				delete(app.UserID, cookie.Value)
				//fmt.Println(app.UserID)
			}
			return
		}
		log.Printf("Received: %s", message)
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Printf("Write failed: %v", err)
			return
		}

		var myMessage MyMessage
		json.Unmarshal(message, &myMessage)
		//fmt.Println(myMessage)
		// cookieValue := r.Header.Get("Cookie")
		// handleWebSocketConnection(conn, cookieValue)
		handleWebSocketMessage(app, conn, myMessage)

	}

}

func handleWebSocketMessage(app *models.App, conn *websocket.Conn, message MyMessage) {

	switch message.Type {
	case "message":
		handleMessageMessage(conn, message)
		notifyMassage(conn, message)
	case "get_users":
		handleGetFriends(conn, message.To)
		handleGetUsersMessage(conn)
		onlineusers(app, conn)
	case "get_chat_history":
		handleGetChatHistoryMessage(conn, message)
	case "read_message":
		SetRead(message.From, message.To)
	case "logout":
		logoutUser(message.From, app)
	default:
		log.Printf("Unsupported message type: %s", message.Type)

	}
}

func logoutUser(id string, app *models.App) {
	// fmt.Println("will remove log out user ====>",id)
	delete(app.UserID, id)
	// fmt.Println(app.UserID)

}

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return db
}

func handleGetFriends(conn *websocket.Conn, to string) {

	db := OpenDatabase()
	defer db.Close()
	users := getAllUsers(db)
	var frinds []ServerUser
	for _, i := range users {
		name := GetUserID(db, i)
		msg, read := GetLastMessage(db, name, to)
		if msg == "" {
			allUsers = append(allUsers, ServerUser{Name: i})
		} else {
			frinds = append(frinds, ServerUser{Name: i}) // Friends = append(Friends, User)
			if read == 0 {
				NotUsers = append(NotUsers, ServerUser{Name: i})
			}
		}
	}
	message := ServerMessage{Type: "frinds", Users: frinds}
	conn.WriteJSON(message)
	//fmt.Println(message)

	// db := OpenDatabase()
	// defer db.Close()
	// users := getFriends(db, to)

	// if users != nil {
	// 	fmt.Println("found users")
	// }
	// //fmt.Println(users)
}

func handleGetUsersMessage(conn *websocket.Conn) {
	message := ServerMessage{Type: "users", Users: allUsers}
	conn.WriteJSON(message)
	//fmt.Println(message)
	message = ServerMessage{Type: "notify", Users: NotUsers}
	conn.WriteJSON(message)

}

func onlineusers(app *models.App, conn *websocket.Conn) {

	message := ServerMessage{Type: "online", Online: app.UserID}
	conn.WriteJSON(message)
}

func handleMessageMessage(conn *websocket.Conn, message MyMessage) {
	db := OpenDatabase()
	defer db.Close()
	name := GetUserName(db, message.From)
	From := message.From
	To := GetUserID(db, message.To)
	// fmt.Println(message.Text)
	AddMessageToHistory(From, To, message.Text)
	recipientConn, ok := (*userSockets)[To]
	if ok {
		nmmessage := ServerMessage{Type: "PM", From: name, Message: message.Text}
		err := recipientConn.WriteJSON(nmmessage)
		//fmt.Println(nmmessage)
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			return
		}
	}
}

func handleGetChatHistoryMessage(conn *websocket.Conn, m MyMessage) {
	db := OpenDatabase()
	defer db.Close()
	To := m.To
	From := GetUserID(db, m.From)
	//fmt.Println(To)
	//fmt.Println(From)
	// fmt.Println("New message  ", m)

	messages := GetChatHistory(To, From, m.Set)

	message := ServerMessage{Type: "oldmessages", ChatHistory: messages}
	conn.WriteJSON(message)
	// fmt.Println(message)

}
func UpdateOfflineUsers(app *models.App, name string) {
	fmt.Println("updating all offilne users")
	message := ServerMessage{Type: "offline", To: name}
	for _, socket := range *userSockets {
		socket.WriteJSON(message)
	}
}

func UpdateOnlineUsers(app *models.App) {
	fmt.Println("updating all users")
	message := ServerMessage{Type: "online", Online: app.UserID}
	for _, socket := range *userSockets {
		socket.WriteJSON(message)
	}
}

func notifyMassage(conn *websocket.Conn, m MyMessage){
	var	PMnotify []ServerUser
	PMnotify = append(PMnotify, ServerUser{Name: m.From})
	recipientConn, ok := (*userSockets)[m.To]
	if ok {
		nmmessage := ServerMessage{Type: "notify", Users: PMnotify}
		err := recipientConn.WriteJSON(nmmessage)
	
		if err != nil {
			log.Println("Error writing to WebSocket:", err)
			return
		}
	}
}