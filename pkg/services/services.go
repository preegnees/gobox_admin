package services

import (
	"context"
	models "jwt/pkg/models"
	"time"
)

func (s *service) SignUp(user models.SignUp) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// тут что то нужно сделать с кодом почты
	if err := s.storage.SaveUser(ctx, user); err != nil {
		return err
	}
	return nil
}
func (s *service) SignIn(user models.SignIn) (ok bool, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	ok, err = s.storage.CheckUser(ctx, user.Username, user.Password)
	if err != nil {
		return false, err
	}
	if ok {
		return true, nil
	} else {
		return false, nil
	}
}
func (s *service) SignOut(refreshToken string) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.storage.DeleteRefreshToken(ctx, refreshToken); err != nil {
		return err
	}
	return nil
}
func (s *service) SaveRefreshToken(refreshToken string) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.storage.SaveRefreshToken(ctx, refreshToken); err != nil {
		return err
	}
	return nil
}
func (s *service) SaveAppData(useraname string, tokens []models.Tokens) (err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.storage.SaveAppData(ctx, useraname, tokens); err != nil {
		return err
	}
	return nil
}
func (s *service) GetAppData(useraname string) (tokens []models.Tokens, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tokens, err = s.storage.GetAppData(ctx, useraname)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}
