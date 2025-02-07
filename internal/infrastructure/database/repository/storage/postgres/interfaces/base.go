package interfaces

import (
	"context"
	"gostart/internal/application/dto"
)

type UserRepository interface {
	SaveUser(ctx context.Context, user *dto.UserDTO) error
}
