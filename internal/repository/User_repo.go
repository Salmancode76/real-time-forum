package repository

import (
	"database/sql"
	"errors"
	"log"
	"real-time-forum/internal/models/entities"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) GetUserByID(id int) (entities.UserData, error) {
	var user entities.UserData
	stmt := `SELECT * FROM User WHERE UserID = ?`
	row := u.DB.QueryRow(stmt, id)
	err := row.Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return user, errors.New("Error finding user")

	}

	return user, nil
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
		log.Printf("Error while inserting: %v\n", err)
		return err
	}

	return nil
}

func (u *UserModel) IsUnique(username, email string) (bool, bool, error) {
	var user entities.UserData
	usernameUnique := true
	emailUnique := true
	stmt := `SELECT UserID FROM User WHERE Username = ?`

	row := u.DB.QueryRow(stmt, username)
	err := row.Scan(&user.UserID)
	if err == nil {
		usernameUnique = false
		log.Print(err)
	} else if err != sql.ErrNoRows {
		return false, false, err
	}

	stmt = `SELECT UserID FROM User WHERE Email = ?`
	row = u.DB.QueryRow(stmt, email)
	err = row.Scan(&user.UserID)

	if err == nil {
		emailUnique = false
	} else if err != sql.ErrNoRows {
		return true, false, err
	}

	return usernameUnique, emailUnique, nil
}
func (u *UserModel) Auth(uename, password string) (entities.UserData, error) {
	var user entities.UserData
	stmt := `SELECT * FROM User WHERE Username = ?`
	row := u.DB.QueryRow(stmt, uename)
	err := row.Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		stmt = `SELECT * FROM User WHERE Email = ?`
		row := u.DB.QueryRow(stmt, uename)
		err = row.Scan(&user.UserID, &user.Username, &user.Age, &user.Gender, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	}
	if err != nil {
		return user, errors.New("Error finding user")
	}

	if err := comparePasswords(user.Password, password); err != nil {
		return user, errors.New("incorrect password")
	}

	return user, nil
}

func comparePasswords(hashedPassword, plainPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}
	return nil
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
