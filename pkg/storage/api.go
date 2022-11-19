package storage

import (
	"context"
	
	models "jwt/pkg/models"
)

type IStorage interface {
	CheckUser(ctx context.Context, username int, password string) (err error)
	SaveRefreshToken(ctx context.Context, username int, refreshToken string) (err error)
	SaveAppTokens(ctx context.Context, username int, appTokens []models.AppData) (err error)
	GiveAppTokens(ctx context.Context, username int) (appTokens []models.AppData, err error)
}

type storage struct{}

var _ IStorage = (*storage)(nil)

func New() IStorage {
	return &storage{}
}
