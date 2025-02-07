package interfaces

import (
	"context"
	"gostart/internal/domain/entities"
)

type UserService interface {
	HashPassword(password string) (string, error)
	CreateUser(ctx context.Context, user *entities.User) error
	ValidatePassword(ctx context.Context, user *entities.User) bool
}
