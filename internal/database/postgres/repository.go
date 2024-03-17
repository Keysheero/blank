package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gostart/internal/database"
	"log/slog"
)

type RepositoryInterface[T database.ModelConstraint] interface {
	Create(ctx context.Context, model T) error
	GetByID(ctx context.Context, id string) (T, error)
}

type Repository struct {
	Client *pgxpool.Pool
	Logger *slog.Logger
}

func NewRepository(client *pgxpool.Pool, logger *slog.Logger) *Repository {
	return &Repository{Client: client, Logger: logger}
}
