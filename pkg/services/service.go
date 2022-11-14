package service

import (
	ms "jwt/pkg/models"
	se "jwt/pkg/storage"
)

type IService interface {
	SignUp(username int, password string, email string, emailCode int) (err error)
	SignIn(username int, password string) (newRefreshToken string, newAccessToken string, err error)
	Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error)
	SaveAppTokens(useraname int, appTokens []ms.AppToken) (err error)
	GiveAppTokens(username int) (appTokens []ms.AppToken, err error)
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

func (s *service) SignUp(username int, password string, email string, emailCode int) (err error) {
	return nil
}
func (s *service) SignIn(username int, password string) (newRefreshToken string, newAccessToken string, err error) {
	return "", "", nil
}
func (s *service) Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error) {
	return "", "", nil
}
func (s *service) SaveAppTokens(useraname int, tokens []ms.AppToken) (err error) {
	return nil
}
func (s *service) GiveAppTokens(useraname int) (tokens []ms.AppToken, err error) {
	return []ms.AppToken{}, nil
}
