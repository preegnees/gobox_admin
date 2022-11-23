package models

type User struct {
	UserId       int
	Username     string
	PasswordHash string
	UserRole     string
	Email        string
}

type Tokens struct {
	Username string
	TempToken    string
}