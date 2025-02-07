package converters

import (
	"github.com/google/uuid"
	"gostart/internal/application/dto"
	"gostart/internal/domain/entities"
)

func ToUserDTO(user *entities.User) *dto.UserDTO {
	return &dto.UserDTO{
		Id:           user.ID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func ToUserEntity(user *dto.UserDTO) *entities.User {
	return &entities.User{
		ID:           uuid.New(),
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}
}

func ToUserEntityFromRegisterDTO(userRegisterDTO *dto.UserRegisterDTO, hashedPassword string) *entities.User {
	return &entities.User{
		ID:           uuid.New(),
		Email:        userRegisterDTO.Email,
		PasswordHash: hashedPassword,
	}
}
