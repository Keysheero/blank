package postgres

import (
	"golang.org/x/crypto/bcrypt"
	"gostart/internal/http-server/schemas"
	user "gostart/internal/models"
)

type Model interface {
	user.User | schemas.UserRegistration
}

func GetHashedPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}
