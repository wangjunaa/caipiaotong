package request

type UserRegister struct {
	Username string
	Password string
	Phone    string
}

type UserLogin struct {
	Phone    string
	Password string
}
