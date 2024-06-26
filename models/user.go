package models

import (
	"errors"
	"eventapi/database"
	"eventapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId
	return err

}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := database.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}
