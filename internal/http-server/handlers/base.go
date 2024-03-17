package handlers

import (
	"gostart/internal/database/postgres"
	"log/slog"
)

type Handler struct {
	Repo   *postgres.UserRepository
	Logger *slog.Logger
}

func NewHandler(repo *postgres.UserRepository, logger *slog.Logger) *Handler {
	return &Handler{
		Repo:   repo,
		Logger: logger,
	}
}

func NewUserHandler(repo *postgres.UserRepository, logger *slog.Logger) *UserHandler {
	return &UserHandler{Handler: Handler{
		Repo:   repo,
		Logger: logger,
	}}
}
