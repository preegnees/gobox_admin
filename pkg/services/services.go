package services

import (
	ms "jwt/pkg/models"
)

func (s *service) SignUp(username string, password string, email string, emailCode int) (err error) {
	return nil
}
func (s *service) SignIn(ms.SignIn) (ok bool, err error) {
	return false, nil
}
func (s *service) Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error) {
	return "", "", nil
}
func (s *service) SaveAppTokens(useraname string, tokens []ms.AppToken) (err error) {
	return nil
}
func (s *service) GiveAppTokens(useraname string) (tokens []ms.AppToken, err error) {
	return []ms.AppToken{}, nil
}
