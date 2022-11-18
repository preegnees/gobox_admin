package services

import (
	se "jwt/pkg/storage"
	ms "jwt/pkg/models"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IService interface {
	SignIn(ms.SignIn) (ok bool, err error)
	SignUp(username string, password string, email string, emailCode int) (err error)
	Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error)
	SaveAppTokens(useraname string, appTokens []ms.AppToken) (err error)
	GiveAppTokens(username string) (appTokens []ms.AppToken, err error)
}

type service struct {
	storage se.IStorage
}

var _ IService = (*service)(nil)

func New(storage se.IStorage) IService {
	return &service{
		storage: storage,
	}
}