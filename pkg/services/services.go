package services

import (
	models "jwt/pkg/models"
)

func (s *service) SignUp(models.SignUp) (err error) {
	return nil
}
func (s *service) SignIn(models.SignIn) (ok bool, err error) {
	return false, nil
}
func (s *service) SignOut(refreshToken string) (err error) {
	return nil
}
func (s *service) Refresh(refreshToken string) (newRefreshToken string, newAccessToken string, err error) {
	return "", "", nil
}
func (s *service) SaveAppData(useraname string, tokens []models.Tokens) (err error) {
	return nil
}
func (s *service) GetAppData(useraname string) (tokens []models.Tokens, err error) {
	return []models.Tokens{}, nil
}
