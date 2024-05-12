package userReqs

type Register struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type Login struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type Update struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
	NewUsername string `json:"newUsername"`
}
type Del struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
