package postgres

import (
	"golang.org/x/crypto/bcrypt"
	user "gostart/internal/models"
	"gostart/internal/schemas"
)

type Model interface {
	user.User | schemas.UserRegistration
}

func GetHashedPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}
