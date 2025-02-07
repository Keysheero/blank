package dto

import (
	"github.com/google/uuid"
)

type UserDTO struct {
	Id           uuid.UUID
	Email        string
	PasswordHash string
}

type UserRegisterDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
