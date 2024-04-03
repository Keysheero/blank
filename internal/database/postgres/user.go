package postgres

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"gostart/internal/config"
	"gostart/internal/http-server/schemas"
	user "gostart/internal/models"
	"log/slog"
)

type UserRepository struct {
	Repository
}

func (ur *UserRepository) NewUser(ctx context.Context, model schemas.UserRegistration) error {
	id := uuid.New()

	hashedPassword := GetHashedPassword(model.Password)
	q := `INSERT INTO "user" (id, name, age, username, password) 
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	_, err := ur.Client.Exec(ctx, q, id, model.Name, model.Age, model.Username, hashedPassword)
	if err != nil {
		ur.Logger.Warn("some with creating new user", "Error", err)
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

func (ur *UserRepository) ValidateUser(ctx context.Context, username string, password string) (user.User, bool) {
	q := `SELECT * FROM "user" 
          WHERE username = $1 AND password = $2;`
	var res user.User
	if err := ur.Client.QueryRow(ctx, q, username, password).Scan(&res.Username, res.Password); err != nil {
		ur.Logger.Warn("Error while executing validate repository function", "ERROR", err)
	}
	if res.Username == username && ValidatePassword(password, res.Password) {
		return res, true
	}
	return res, true
}

func NewUserRepository(client *pgxpool.Pool, logger *slog.Logger, config *config.Config) *UserRepository {
	return &UserRepository{
		Repository{
			Client: client,
			Logger: logger,
			Config: config,
		},
	}
}

func ValidatePassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
