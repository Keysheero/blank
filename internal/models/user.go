package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Age      int
	Username string
	Password string
}

type Data struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Username string
	Password string
}
