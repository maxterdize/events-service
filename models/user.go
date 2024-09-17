package models

import (
	"database/sql"
	"errors"
	"events-service/db"
	"events-service/utils"
)

type User struct {
	ID       float64 `json:"-"`
	Email    string  `json:"email" binding:"required"`
	Password string  `json:"password" binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users (email, password) 
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	user.ID = float64(userId)

	return err
}

func (user *User) Authenticate() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, user.Email)

	var retrievedPassword string
	err := row.Scan(&user.ID, &retrievedPassword)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("invalid email or password")
		}
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid email or password")
	}

	return nil
}
