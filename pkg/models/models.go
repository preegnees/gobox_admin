package models

import (
	"context"
)

type Password string
type Email string
type RefreshToken string
type AccessToken string
type AppToken struct {
	Token string
}
type AppTokens []AppToken
type Ctx context.Context

type IService interface {
	SignUp(Email, Password) error
	SignIn(Email, Password) (RefreshToken, AccessToken, error)
	Refresh(RefreshToken) (RefreshToken, AccessToken, error)
	SaveAppTokens(AppTokens) error
	GiveAppTokens(AccessToken) error
}

type IStorage interface {
	CheckUser(Ctx, Email, Password) error
	SaveRefreshToken(Ctx, Email, RefreshToken) error
	SaveAppTokens(Ctx, Email, AppTokens) error
	GiveAppTokens(Ctx, Email) (AppTokens, error)
}

type IServer interface {
	Run() error
}

type SignIn struct {
	Email    `json:"email"`
	Password `json:"password"`
}

type SignUp struct {
	Email    `json:"email"`
	Password `json:"password"`
}
