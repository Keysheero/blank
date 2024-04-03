package handlers

import (
	"gostart/internal/config"
	"gostart/internal/database/postgres"
	"log/slog"
)

type Handler struct {
	Repo   *postgres.UserRepository
	Logger *slog.Logger
	Config *config.Config
}

func NewHandler(repo *postgres.UserRepository, logger *slog.Logger) *Handler {
	return &Handler{
		Repo:   repo,
		Logger: logger,
	}
}
