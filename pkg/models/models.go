package models

type SignIn struct {
	Username int    `json:"username"`
	Password string `json:"password"`
}

type SignUp struct {
	Username  int    `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	EmailCode int    `json:"email_code"`
}

type AppToken struct {
	Token string `json:"token"`
}

type SaveAppTokens struct {
	Username int        `json:"username"`
	Tokens   []AppToken `json:"app_tokens"`
}

type GiveAppTokens struct {
	Username int `json:"useraname"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	AppToken `json:"app_tokens"`
}
