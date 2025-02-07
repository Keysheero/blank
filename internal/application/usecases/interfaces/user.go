package interfaces

import (
	"gostart/internal/application/dto"
)

type UserUsecase interface {
	RegisterNewUser(user *dto.UserRegisterDTO) error
}
