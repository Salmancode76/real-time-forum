package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) Insert(username, email, password, gender, fname, lname string, age int) error {
	hashed, err := hashPassword(password)
	if err != nil {
		return err
	}
	stmt := `INSERT INTO User (
                     Username,
                     Age,
                     Gender,
                     First_Name,
                     Last_Name,
                     Email,
                     Password
                 )
                 VALUES (
                     ?,
                     ?,
                     ?,
                     ?,
                     ?,
                     ?,
                     ?
                 );`

	_, err = u.DB.Exec(stmt, username, age, gender, fname, lname, email, hashed)

	if err != nil {
		fmt.Printf("Error while inserting: %v\n", err)
		return err
	}

	return nil

}

func hashPassword(password string) (string, error) {
	// Create a SHA-256 hash object
	hash := sha256.New()

	// Write the password to the hash object
	_, err := hash.Write([]byte(password))
	if err != nil {

		return "", err
	}

	// Get the hashed bytes
	hashedBytes := hash.Sum(nil)

	// Convert the hashed bytes to a hex string
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString, nil
}
