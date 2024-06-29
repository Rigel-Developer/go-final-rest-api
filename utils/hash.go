package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hashedPassword, plainTextPassword string) error {
	fmt.Println(hashedPassword, plainTextPassword)
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainTextPassword))
}
