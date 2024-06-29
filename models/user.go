package models

import (
	"fmt"
	"github/rigel-developer/go-final-rest-api/db"
	"github/rigel-developer/go-final-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	insertUserSQL := `
	INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(insertUserSQL)
	if err != nil {
		return err

	}
	defer stmt.Close()

	//hash password
	hashedPassword, err := utils.HashPassword(u.Password)
	fmt.Println(hashedPassword)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateCredentials() error {
	row := db.DB.QueryRow("SELECT id, password FROM users WHERE email = ?", u.Email)

	var passwordHashed string
	err := row.Scan(&u.ID, &passwordHashed)
	if err != nil {
		return err
	}

	err = utils.ComparePassword(passwordHashed, u.Password)
	if err != nil {
		return err
	}

	return nil
}
