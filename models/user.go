package models

import (
	"errors"

	"github.com/chethanbhat/go-rest-api/db"
	"github.com/chethanbhat/go-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}
	userID, err := result.LastInsertId()

	u.ID = userID

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users where email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var hashedPassword string

	err := row.Scan(&u.ID, &hashedPassword)

	if err != nil {
		return err
	}

	err = utils.ComparePassword(hashedPassword, u.Password)

	if err != nil {
		return errors.New("invalid credentials")
	}

	return nil

}
