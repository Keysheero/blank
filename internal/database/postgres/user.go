package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	user "gostart/internal/models"
	"gostart/internal/schemas"
	"log/slog"
)

type UserRepository struct {
	Repository
}

func (ur *UserRepository) Create(ctx context.Context, model schemas.UserRegistration) error {
	id := uuid.New()

	hashedPassword := GetHashedPassword(model.Password)
	q := `INSERT INTO "user" (id, name, age, username, password) 
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	_, err := ur.Client.Exec(ctx, q, id, model.Name, model.Age, model.Username, hashedPassword)
	if err != nil {
		ur.Logger.Warn("some problems with function(CREATE) execution", "Error", err)
		return err
	}

	return nil
}

func (ur *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (user.User, error) {
	q := `SELECT * FROM "user"
		  WHERE id = $1`
	var res user.User
	if err := ur.Client.QueryRow(ctx, q, id).Scan(&res.ID, &res.Name, &res.Age, &res.Username, &res.Password); err != nil {
		ur.Logger.Warn("Error while executing GetByID func", "ERROR:", err)
		return res, err
	}
	return res, nil
}

func (ur *UserRepository) GetByName(ctx context.Context, name string) (user.User, error) {
	q := `SELECT * FROM "user"
		WHERE name = $1`
	var res user.User
	if err := ur.Client.QueryRow(ctx, q, name).Scan(&res.ID, &res.Name, &res.Age); err != nil {
		ur.Logger.Warn("Error while executing GetByName func", "ERROR", err)
		return res, nil
	}
	return res, nil
}

func NewUserRepository(client *pgxpool.Pool, logger *slog.Logger) *UserRepository {
	return &UserRepository{
		Repository{
			Client: client,
			Logger: logger,
		},
	}
}
