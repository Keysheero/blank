package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"gostart/internal/infrastructure/config"
	"gostart/internal/infrastructure/logger"
	"log/slog"
)

type Repository struct {
	Client *pgxpool.Pool
	Logger logger.Logger
	Config *config.Config
}

func NewRepository(client *pgxpool.Pool, logger *slog.Logger) *Repository {
	return &Repository{Client: client, Logger: logger}
}
