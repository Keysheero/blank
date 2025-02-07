package user

import (
	"context"
	"gostart/internal/application/dto"
	"gostart/internal/application/usecases/interfaces"
	serviceInterfaces "gostart/internal/domain/services/interfaces"
	"gostart/internal/infrastructure/converters"
	"gostart/internal/infrastructure/logger"
)

type userUsecase struct {
	UserService serviceInterfaces.UserService
	logger      logger.Logger
}

func NewUserUsecase(us serviceInterfaces.UserService, logger logger.Logger) interfaces.UserUsecase {
	return &userUsecase{us, logger}
}

func (uuc *userUsecase) RegisterNewUser(user *dto.UserRegisterDTO) error {
	hashedPassword, err := uuc.UserService.HashPassword(user.Password)
	userEntity := converters.ToUserEntityFromRegisterDTO(user, hashedPassword)

	if err != nil && uuc.UserService.ValidatePassword(context.Background(), userEntity) {
		return err
	}
	err = uuc.UserService.CreateUser(context.Background(), userEntity)
	if err != nil {
		return err
	}
	return nil
}
