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

// User tabel
// вынести в сервисы, убрать json ()
// apimodel {createUserResp, create ...}, (entity, db_model + db_tag)

type CreateUserReq struct { // DTO
	Username string  // with json tag
	Password string
	Email    string
}

// DTO - приходили и выходили только те данные, которые используются (почитать)
// api -> DTO (глобальная папка) -> service

type UserEntity struct {
	id           int64  //, guid,
	Username     string `json:"username" validate:"required"`
	PasswordHash string `json:"password_hash" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Role         string `json:"role" validate:"required"`
	// Tokens   `json:"tokens"`
}

// Tokens tabel
type TokensEntity struct {
	Username string `json:"username" validate:"required"`
	Token    string `json:"token" validate:"required"`
}
