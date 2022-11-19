package storage

import (
	"context"

	models "jwt/pkg/models"
	memstorage "jwt/pkg/repository/memstorage"
	storage "jwt/pkg/repository/storage"
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

type repository struct{
	memstorage memstorage.IMemStorage
	storage storage.IStorage
}

var _ IStorage = (*repository)(nil)

func New(m memstorage.IMemStorage, s storage.IStorage) IStorage {
	return &repository{
		memstorage: m,
		storage: s,
	}
}
