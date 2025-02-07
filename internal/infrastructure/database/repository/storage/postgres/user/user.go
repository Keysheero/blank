package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gostart/internal/application/dto"
	"gostart/internal/infrastructure/config"
	"gostart/internal/infrastructure/database/repository/storage/postgres"
	"gostart/internal/infrastructure/database/repository/storage/postgres/interfaces"
	"gostart/internal/infrastructure/logger"
)

type userRepository struct {
	postgres.Repository
}

func (ur *userRepository) SaveUser(ctx context.Context, user *dto.UserDTO) error {

	q := `INSERT INTO "users" (id, email, password_hash) 
          VALUES ($1, $2, $3)
          RETURNING id`
	tag, err := ur.Client.Exec(ctx, q, user.Id, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	ur.Logger.Info("New User has been created", tag.String())
	return nil
}

func NewUserRepository(client *pgxpool.Pool, logger logger.Logger, config *config.Config) interfaces.UserRepository {
	return &userRepository{
		Repository: postgres.Repository{
			Client: client,
			Logger: logger,
			Config: config,
		},
	}
}
