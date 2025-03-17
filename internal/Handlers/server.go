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
var userSockets = make(map[string]*websocket.Conn)
// var clients = make(map[*websocket.Conn]*Session)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development only)
	},
}

func HandleWebSocket(app *models.App, w http.ResponseWriter, r *http.Request) {
	
	cookie, err := r.Cookie("userID")
	

	fmt.Println("cookie name =",cookie.Name)
	fmt.Println("cookie value =",cookie.Value)
	
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()
	userSockets[cookie.Value]=conn
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
		handleWebSocketMessage(conn, myMessage)

	}
}

func handleWebSocketMessage(conn *websocket.Conn, message MyMessage) {
	//fmt.Println("Message received: ", message.Type)

	switch message.Type {
	case "message":
		handleMessageMessage(conn, message)

	case "get_users":
		fmt.Println("will get users now=======>>>>>")
		handleGetUsersMessage(conn)
	case "get_chat_history":
		handleGetChatHistoryMessage(conn, message)
	case "read_message":
		SetRead(message.From, message.To)
	 default:
        log.Printf("Unsupported message type: %s", message.Type)

	}
}

func OpenDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "db/db.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}
	return db
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
func handleMessageMessage(conn *websocket.Conn, message MyMessage) {
	db := OpenDatabase()
	defer db.Close()
	name:=GetUserName(db,message.From)
	From := message.From
	To := GetUserID(db, message.To)
	fmt.Println(message.Text)
	AddMessageToHistory(From, To, message.Text)
	recipientConn, ok := userSockets[To]
			if ok {
				nmmessage := ServerMessage{Type: "PM",From: name ,Message: message.Text}
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
		fmt.Println("New message  ",m)

	messages := GetChatHistory(To, From, m.Set)

	message := ServerMessage{Type: "oldmessages", ChatHistory: messages}
	conn.WriteJSON(message)
	fmt.Println(message)

}

// func handleWebSocketConnection(conn *websocket.Conn, cookieValue string) {
// 	// SessionsMutex is a mutex to lock the sessions map when adding or removing sessions from it
// 	var sessionsMutex sync.Mutex
// 	fmt.Println("New WebSocket Connection", conn.RemoteAddr().String())

// 	if cookieValue != "" {
// 		for _, v := range LoggedInUsers {
// 			if v.Cookie == cookieValue[14:] {
// 				fmt.Println("Server >> User is logged in with cookie value: ", cookieValue)
// 				fmt.Println("Server >> Updating connection ID for user: ", v.Username)
// 				v.WebSocketConn = conn.RemoteAddr().String()
// 			}
// 		}
// 	} else {
// 		fmt.Println("Server >> User is not logged in")
// 		conn.WriteJSON(ServerMessage{Type: "status", Data: map[string]string{"refresh": "true"}})
// 	}

// 	// Create a new session for the WebSocket client
// 	session := Session{
// 		WebSocketConn: conn.RemoteAddr().String(),
// 		UserID:        0,
// 	}

// 	// Add the session to the clients map
// 	// Note: you need to synchronize access to the map using a mutex
// 	sessionsMutex.Lock()
// 	clients[conn] = &session
// 	sessionsMutex.Unlock()

// }
