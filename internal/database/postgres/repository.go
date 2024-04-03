package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"gostart/internal/config"
	"log/slog"
)

type RepositoryInterface[T Model] interface {
	NewUser(ctx context.Context, model T) error
	GetByID(ctx context.Context, id uuid.UUID) (T, error)
	GetByName(ctx context.Context, name string) (T, error)
}

type Repository struct {
	Client *pgxpool.Pool
	Logger *slog.Logger
	Config *config.Config
}

func NewRepository(client *pgxpool.Pool, logger *slog.Logger) *Repository {
	return &Repository{Client: client, Logger: logger}
}
