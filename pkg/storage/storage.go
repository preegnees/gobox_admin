package storage

import (
	"context"
	
	models "jwt/pkg/models"
)

func (s *storage) CheckUser(ctx context.Context, username int, password string) (err error) {
	return nil
}
func (s *storage) SaveRefreshToken(ctx context.Context, username int, refreshToken string) (err error) {
	return nil
}
func (s *storage) SaveAppTokens(ctx context.Context, username int, appTokens []models.AppData) (err error) {
	return nil
}
func (s *storage) GiveAppTokens(ctx context.Context, username int) (appTokens []models.AppData, err error) {
	return []models.AppData{}, nil
}
