package storage

import (
	"context"

	models "jwt/pkg/models"
)

func (s *repository) SaveUser(ctx context.Context, signUp models.SignUp) (err error) {
	return nil
}
func (s *repository) CheckUser(ctx context.Context, username string, password string) (ok bool, err error) {
	return false, nil
}
func (s *repository) SaveAppData(ctx context.Context, username string, appTokens []models.Tokens) (err error) {
	return nil
}
func (s *repository) GetAppData(ctx context.Context, username string) (appTokens []models.Tokens, err error) {
	return []models.Tokens{}, nil
}

func (s *repository) SaveRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.SaveRefreshToken(ctx, refreshToken)
}
func (s *repository) CheckRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.CheckRefreshToken(ctx, refreshToken)
}
func (s *repository) DeleteRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.DeleteRefreshToken(ctx, refreshToken)
}
