package schemas

type UserRegistration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserLogin struct {
	Username string
	Password string
}
