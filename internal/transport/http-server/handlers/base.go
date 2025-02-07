package handlers

import (
	"gostart/internal/application/usecases/interfaces"
	"gostart/internal/infrastructure/logger"
	"gostart/internal/infrastructure/validator/interfaces"
)

type Handler struct {
	Logger logger.Logger
	UUC    interfaces.UserUsecase
	UV     validator.validator
}
