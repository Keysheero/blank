package schemas

type UserRegistration struct {
	Name     string
	Age      int
	Username string
	Password string
}

type UserLogin struct {
	Username string
	Password string
}
