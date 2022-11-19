package models

type SignIn struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type SignUp struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	Role      string `json:"role" validate:"required"`
	Email     string `json:"email" validate:"required"`
	EmailCode int    `json:"email_code" validate:"required"`
}

type Tokens struct {
	Token  string `json:"token"`
	Action int    `json:"action"` // 0 default, add 1, remove 2
}

type AppData struct {
	Username string   `json:"username" validate:"required"`
	Tokens   []Tokens `json:"tokens"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Tokens   `json:"tokens"`
}
