package storage

import (
	"context"
	ms "jwt/pkg/models"
)

// TODO(в мидваре нужно превратить username и пароль в число)
type IStorage interface {
	CheckUser(ctx context.Context, username int, password string) (err error)
	SaveRefreshToken(ctx context.Context, username int, refreshToken string) (err error)
	SaveAppTokens(ctx context.Context, username int, appTokens []ms.AppToken) (err error)
	GiveAppTokens(ctx context.Context, username int) (appTokens []ms.AppToken, err error)
}

type storage struct{}

var _ IStorage = (*storage)(nil)

func New() IStorage {
	return &storage{}
}

func (s *storage) CheckUser(ctx context.Context, username int, password string) (err error) {
	return nil
}
func (s *storage) SaveRefreshToken(ctx context.Context, username int, refreshToken string) (err error) {
	return nil
}
func (s *storage) SaveAppTokens(ctx context.Context, username int, appTokens []ms.AppToken) (err error) {
	return nil
}
func (s *storage) GiveAppTokens(ctx context.Context, username int) (appTokens []ms.AppToken, err error) {
	return []ms.AppToken{}, nil
}
