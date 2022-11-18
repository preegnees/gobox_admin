package models

type SignIn struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type SignUp struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	EmailCode int    `json:"email_code"`
}

type AppToken struct {
	Token string `json:"token"`
}

type SaveAppTokens struct {
	Username string     `json:"username"`
	Tokens   []AppToken `json:"app_tokens"`
}

type GiveAppTokens struct {
	Username string `json:"useraname"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	AppToken `json:"app_tokens"`
}

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
