package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"gostart/internal/infrastructure/config"
	"gostart/internal/infrastructure/logger"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, conData config.DatabaseConfig, logger logger.Logger) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conData.Username,
		conData.Password,
		conData.Host,
		conData.Port,
		conData.Name,
	)

	for attempt := 1; attempt <= conData.MaxAttempts; attempt++ {
		pool, err = pgxpool.New(ctx, dsn)
		if err != nil {
			logger.Error("Error with pool connection", err)
		} else {
			break
		}
		time.Sleep(time.Duration(attempt) * time.Second)
	}
	return

}
