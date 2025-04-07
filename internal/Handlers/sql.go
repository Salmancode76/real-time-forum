package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Open the database

//closeing the data base

func CloseDB(db *sql.DB) {
	db.Close()
}
func SetRead(from string, to string) {

	db := OpenDatabase()
	user := GetUserID(db, from)
	defer CloseDB(db)
	_, err := db.Exec("UPDATE messages SET is_read = 1 WHERE from_id = ? AND to_id = ?", user, to)
	if err != nil {
		fmt.Printf("Server >> Error setting message as read: %s", err)
	}

}

func AddMessageToHistory(fromUSer string, toUser string, messageText string) {
	isread := 0
	db := OpenDatabase()
	defer CloseDB(db)
	//inserting data into table
	_, err := db.Exec("INSERT INTO messages (from_id,to_id,message,time,is_read) VALUES (?,?,?,?,?)", fromUSer, toUser, messageText, time.Now().Format("2006-01-02 15:04:05"), isread)
	if err != nil {
		// Handle error
		fmt.Printf("Server >> Error adding message to database: %s ", err)
	}
}

func GetChatHistory(user string, from string, offset int) []Message {
	db := OpenDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT from_id, to_id, is_read, message, time FROM messages WHERE (from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?) ORDER BY time DESC LIMIT 10 OFFSET ?", user, from, from, user, offset)
	//rows, err := db.Query("SELECT from_id, to_id, is_read, message, time FROM messages ")

	if err != nil {
		fmt.Printf("Server >> Error getting chat history: %s", err)
	}

	messages := []Message{}
	for rows.Next() {
		var fromUser, toUser string
		var isread int
		var message string
		var time string
		err = rows.Scan(&fromUser, &toUser, &isread, &message, &time)
		if err != nil {
			fmt.Printf("Server >> Error reading chat history: %s", err)
		}

		toUser = GetUsernameFromId(db, toUser)
		fromUser = GetUsernameFromId(db, fromUser)

		msg := Message{
			From:      fromUser,
			To:        toUser,
			Read:      isread,
			Text:      message,
			CreatedAt: time,
		}
		messages = append(messages, msg)
	}

	return messages
}

// Get username depending on userID
func GetUsernameFromId(db *sql.DB, id string) string {
	// Prepare the SQL query to retrieve the user ID based on the username
	query := "SELECT Username FROM User WHERE UserID = ?"

	// Execute the query and retrieve the user ID
	var username string
	err := db.QueryRow(query, id).Scan(&username)
	if err != nil {
		fmt.Printf("Server >> Error getting user ID: %s", err)
	}

	return username
}

func GetUserName(db *sql.DB, From string) string {
	query := "SELECT Username FROM User WHERE UserID = ?"

	// Execute the query and retrieve the user ID
	var username string
	err := db.QueryRow(query, From).Scan(&username)
	if err != nil {
		fmt.Printf("Server >> Error getting user ID: %s", err)
	}

	return username
}

func GetUserID(db *sql.DB, username string) string {
	// Prepare the SQL query to retrieve the user ID based on the username
	query := "SELECT UserID FROM User WHERE Username = ?"
	// Execute the query and retrieve the user ID
	var userID string
	err := db.QueryRow(query, username).Scan(&userID)
	if err != nil {
		fmt.Printf("Server >> Error getting user ID: %s", err)
	}

	return userID
}

func getAllUsers(db *sql.DB) []string {

	query := "SELECT Username FROM User"
	var names []string
	rows, err := db.Query(query)
	if err != nil {
		fmt.Printf("Server >> Error getting all users: %s", err)
	}
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			fmt.Printf("Server >> Error getting all users: %s", err)
		}
		names = append(names, name)
	}
	return names
}

// func getFriends(db *sql.DB, to string) []string {

// 	query := "SELECT Username FROM User WHERE UserID != ?"
// 	var names []string
// 	rows, err := db.Query(query,to)
// 	if err != nil {
// 		fmt.Printf("Server >> Error getting all users: %s", err)
// 	}
// 	for rows.Next() {
// 		var name string
// 		err = rows.Scan(&name)
// 		if err != nil {
// 			fmt.Printf("Server >> Error getting all users: %s", err)
// 		}
// 		names = append(names, name)
// 	}
// 	return names
// }

func GetLastMessage(db *sql.DB, senderId string, receiverId string) (string,int) {

	var message string
	var read int
	fmt.Println(senderId)
	err := db.QueryRow(`
    SELECT  message, is_read FROM messages
    WHERE (from_id = ? AND to_id = ?) OR (from_id = ? AND to_id = ?)
    ORDER BY messageID DESC
    LIMIT 1`,
		senderId, receiverId, receiverId, senderId,
	).Scan(&message,&read)
	if err != nil {
		if err == sql.ErrNoRows {
			return "",1
		} else {
			log.Fatal(err)
		}
	}

	return message ,read
}

// func GetOtherUsersData(UserId int, activeUsers []int) []User {
// 	var Friends []User
// 	var Stranger []User

// 	// Query the database for post data
// 	rows, err := Accountsdb.Query("SELECT id, Username, UserImg FROM accounts WHERE id != ?", UserId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()

// 	// Iterate over the rows and populate the Posts slice
// 	for rows.Next() {
// 		var User User
// 		err := rows.Scan(&User.UserId, &User.Username, &User.ProfileImg)
// 		if err != nil {
// 			log.Println("Error scanning row:", err)
// 			continue
// 		}
// 		User.MessageID, User.LastMessage = GetLastMessage(UserId, User.UserId)
// 		if IsUserActive(User.UserId, activeUsers) {
// 			User.Status = "Online"
// 		} else {
// 			User.Status = "Offline"
// 		}
// 		User.Notifications = GetNumOfNotifications(User.UserId, UserId)

// 		if User.LastMessage == "" {
// 			Stranger = append(Stranger, User)
// 		} else {
// 			Friends = append(Friends, User)
// 		}
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	Users := CombineUsers(Friends, Stranger)

// 	return Users
// }
