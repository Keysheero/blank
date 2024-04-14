package postgres

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"gostart/internal/config"
	"gostart/internal/http-server/schemas"
	user2 "gostart/internal/models"
	"log/slog"
)

type UserRepository struct {
	Repository
}

func (ur *UserRepository) NewUser(ctx context.Context, model schemas.UserRegistration) error {
	id := uuid.New()
	attendance := 0
	status := "LFK"
	ur.Logger.Info("server has got model", model.Email, model.Password)
	hashedPassword := GetHashedPassword(model.Password)
	q := `INSERT INTO "student" (id, email, password, status, attendance) 
			  VALUES ($1, $2, $3, $4, $5)
			  RETURNING id`
	tag, err := ur.Client.Exec(ctx, q, id, model.Email, hashedPassword, status, attendance)

	if err != nil {
		ur.Logger.Warn("some problems  with creating new user", "Error", err)
		return err
	}
	ur.Logger.Info("New User has been created", tag.String())
	return nil
}

func (ur *UserRepository) GetUserByEmail(ctx context.Context, email string) (user2.Student, error) {
	var user user2.Student
	q := `SELECT id, email, password, status, attendance FROM "student" WHERE email = $1`
	row := ur.Client.QueryRow(ctx, q, email)
	if err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Status, &user.Attendance); err != nil {
		if err == pgx.ErrNoRows {
			return user, fmt.Errorf("user with email %s not found", email)
		}
		return user, err
	}
	return user, nil
}

//
//func (ur *UserRepository) ValidateUser(ctx context.Context, username string, password string) (user.User, bool) {
//	q := `SELECT * FROM "student"
//          WHERE username = $1 AND password = $2;`
//	var res user.User
//	if err := ur.Client.QueryRow(ctx, q, username, password).Scan(&res.Username, res.Password); err != nil {
//		ur.Logger.Warn("Error while executing validate repository function", "ERROR", err)
//	}
//	if res.Username == username && ValidatePassword(password, res.Password) {
//		return res, true
//	}
//	return res, true
//}

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
