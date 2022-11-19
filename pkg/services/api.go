package services

import (
	models "jwt/pkg/models"
	repository "jwt/pkg/repository"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IService interface {
	SignIn(models.SignIn) (ok bool, err error)
	SignUp(models.SignUp) (err error)
	SignOut(refreshToken string) (err error)
	SaveRefreshToken(refreshToken string) (err error)
	SaveAppData(useraname string, appData []models.Tokens) (err error)
	GetAppData(username string) (appTokens []models.Tokens, err error)
}

// TODO(определить тут интерфейс)
type service struct {
	storage repository.IStorage
}

var _ IService = (*service)(nil)

func New(repository repository.IStorage) IService {
	return &service{
		storage: repository,
	}
}
