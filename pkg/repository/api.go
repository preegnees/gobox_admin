package storage

import (
	"context"

	models "jwt/pkg/models"
	memstorage "jwt/pkg/storage/memstorage"
)

type IStorage interface {
	SaveUser(ctx context.Context, signUp models.SignUp) (err error)
	CheckUser(ctx context.Context, username string, password string) (ok bool, err error)
	CheckRefreshToken(ctx context.Context, refreshToken string) (err error)
	SaveRefreshToken(ctx context.Context, refreshToken string) (err error)
	DeleteRefreshToken(ctx context.Context, refreshToken string) (err error)
	SaveAppData(ctx context.Context, username string, appTokens []models.Tokens) (err error)
	GetAppData(ctx context.Context, username string) (appTokens []models.Tokens, err error)
}

type storage struct{
	memstorage memstorage.IRedis
}

var _ IStorage = (*storage)(nil)

func New(m memstorage.IRedis) IStorage {
	return &storage{
		memstorage: m,
	}
}
