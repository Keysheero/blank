package user

import (
	"context"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gostart/internal/domain/entities"
	serviceInterfaces "gostart/internal/domain/services/interfaces"
	"gostart/internal/infrastructure/config"
	user2 "gostart/internal/infrastructure/converters"
	"gostart/internal/infrastructure/database/repository/storage/postgres/interfaces"
	"gostart/internal/infrastructure/logger"
	"time"
)

type userService struct {
	UserRepo interfaces.UserRepository
	Logger   logger.Logger
	Config   config.Config
}

func NewUserService(userRepo interfaces.UserRepository, logger logger.Logger) serviceInterfaces.UserService {
	return &userService{UserRepo: userRepo, Logger: logger}
}

func (us *userService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (us *userService) ValidatePassword(ctx context.Context, user *entities.User) bool {
	password := user.PasswordHash
	if len(password) < 8 {
		return false
	} else {
		return true
	}
}

func (us *userService) CreateUser(ctx context.Context, user *entities.User) error {
	userModel := user2.ToUserDTO(user)
	err := us.UserRepo.SaveUser(ctx, userModel)
	if err != nil {
		return err
	}
	return nil
}

func (us *userService) GenerateJwtToken(ctx context.Context, user *entities.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     us.Config.Auth.AccessTokenTTL,
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(us.Config.Auth.Secret)
}

func (us *userService) ValidateJwtToken(ctx context.Context, tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(us.Config.Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
