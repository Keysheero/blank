package user

import (
	usecaseinterfaces "gostart/internal/application/usecases/interfaces"
	"gostart/internal/infrastructure/logger"
	"gostart/internal/infrastructure/validator/interfaces"
	"gostart/internal/transport/http-server/handlers"
	"gostart/internal/transport/http-server/handlers/interfaces"
	"gostart/internal/validator/interfaces"
)

type userHandler struct {
	handlers.Handler
}

func NewUserHandler(uuc usecaseinterfaces.UserUsecase, logger logger.Logger, uv validator.validator) interfaces.UserHandler {
	return userHandler{handlers.Handler{
		Logger: logger,
		UUC:    uuc,
		UV:     uv,
	}}
}
