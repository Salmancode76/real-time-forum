package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins (for development only)
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade failed: %v", err)
		return
	}
	defer conn.Close()

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
	To := message.To
	From := GetUserID(db, message.From)
	fmt.Println(message.Text)
	AddMessageToHistory(From, To, message.Text)
}

func handleGetChatHistoryMessage(conn *websocket.Conn, m MyMessage) {
	db := OpenDatabase()
	defer db.Close()

	To := m.To
	From := GetUserID(db, m.From)
	fmt.Println(To)
	fmt.Println(From)
	messages := GetChatHistory(To, From, 0)

	message := ServerMessage{Type: "oldmessages", ChatHistory: messages}
	conn.WriteJSON(message)
	fmt.Println(message)

}
