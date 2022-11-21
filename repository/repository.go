package repository

import (
	"context"
	dto "core/dto"
	services "core/services"
)

//go:generate mockgen -source=repository.go -destination=mock/mock.go

type IMemStorage interface {
	SaveUserMetadata(ctx context.Context, metadata *dto.DTOUserMetadata) (err error)
	GetUserMetadata(ctx context.Context, fingerprint string) (metadata *dto.DTOUserMetadata, err error)
}

type IStorage interface {
	SaveUser(ctx context.Context, user *dto.DTOUser) (err error)
	SaveToken(ctx context.Context, token *dto.DTOTokens) (err error)
	RemoveToken(ctx context.Context, token *dto.DTOTokens) (err error)
	GetTokens(ctx context.Context, username string) (tokens *[]dto.DTOTokens, err error)
}

type repository struct {
	memStorage IMemStorage
	storage    IStorage
}

var _ services.IRepository = (*repository)(nil)

func New(m IMemStorage, s IStorage) *repository {
	return &repository{
		memStorage: m,
		storage:    s,
	}
}
