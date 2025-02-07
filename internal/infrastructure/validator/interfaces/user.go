package validator

type UserValidator interface {
	ValidateEmail(email string) error
	ValidatePassword(password string) error
}
