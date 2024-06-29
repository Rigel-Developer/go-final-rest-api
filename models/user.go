package models

import (
	"github/rigel-developer/go-final-rest-api/db"
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

	result, err := stmt.Exec(u.Email, u.Password)

	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()

	if err != nil {
		return err
	}

	return nil
}
