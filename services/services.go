package services

import (
	"context"
	dto "core/dto"
	handlers "core/handlers"
)

//go:generate mockgen -source=services.go -destination=mock/mock.go

type IRepository interface {
	SaveUserMetadata(ctx context.Context, metadata *dto.DTOUserMetadata) (err error)
	GetUserMetadata(ctx context.Context, fingerprint string) (metadata *dto.DTOUserMetadata, err error)

	SaveUser(ctx context.Context, user *dto.DTOUser) (err error)
	
	SaveToken(ctx context.Context, token *dto.DTOTokens) (err error)
	RemoveToken(ctx context.Context, token *dto.DTOTokens) (err error)
	GetTokens(ctx context.Context, username string) (tokens *[]dto.DTOTokens, err error)
}

type services struct {
	repository IRepository
}

var _ handlers.IServices = (*services)(nil)

func New(r IRepository) *services {
	return &services{
		repository: r,
	}
}