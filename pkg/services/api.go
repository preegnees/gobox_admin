package services

import (
	models "jwt/pkg/models"
	storage "jwt/pkg/storage"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IService interface {
	SignIn(models.SignIn) (ok bool, err error)
	SignUp(models.SignUp) (err error)
	SignOut(refreshToken string) (err error)
	Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error)
	SaveAppData(useraname string, appData []models.Tokens) (err error)
	GetAppData(username string) (appTokens []models.Tokens, err error)
}

type service struct {
	storage storage.IStorage
}

var _ IService = (*service)(nil)

func New(storage storage.IStorage) IService {
	return &service{
		storage: storage,
	}
}
