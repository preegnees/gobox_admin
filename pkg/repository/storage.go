package storage

import (
	"context"

	models "jwt/pkg/models"
)

func (s *storage) SaveUser(ctx context.Context, signUp models.SignUp) (err error) {
	return nil
}
func (s *storage) CheckUser(ctx context.Context, username string, password string) (ok bool, err error) {
	return false, nil
}
func (s *storage) SaveAppData(ctx context.Context, username string, appTokens []models.Tokens) (err error) {
	return nil
}
func (s *storage) GetAppData(ctx context.Context, username string) (appTokens []models.Tokens, err error) {
	return []models.Tokens{}, nil
}

func (s *storage) SaveRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.SaveRefreshToken(ctx, refreshToken)
}
func (s *storage) CheckRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.CheckRefreshToken(ctx, refreshToken)
}
func (s *storage) DeleteRefreshToken(ctx context.Context, refreshToken string) (err error) {
	return s.memstorage.DeleteRefreshToken(ctx, refreshToken)
}
