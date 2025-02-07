package validator

import (
	"errors"
	"gostart/internal/infrastructure/validator/interfaces"
	"regexp"
)

type userValidator struct {
}

func NewUserValidator() validator.UserValidator {
	return &userValidator{}
}

func (uv *userValidator) ValidateEmail(email string) error {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(regex)

	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}

func (uv *userValidator) ValidatePassword(password string) error {
	if len(password) >= 8 && len(password) <= 24 {
		return nil
	} else {
		return errors.New("invalid password format. it should be more than 8 and less than 24 long")
	}
}
