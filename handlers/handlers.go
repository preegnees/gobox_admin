package handlers

import (
	dto "core/dto"
)

//go:generate mockgen -source=handlers.go -destination=mock/mock.go

type IServices interface {
	ConfirmEmail(email string) (err error)
	SendEmailCode(email string) (err error)

	SignUp(user *dto.DTOUser) (err error)
	SignIn(user *dto.DTOUser) (err error)
	SignOut(metadata *dto.DTOUserMetadata) (err error)

	SaveToken(token *dto.DTOTokens) (err error)
	RemoveToken(token *dto.DTOTokens) (err error)
	GetTokens(username string) (tokens *[]dto.DTOTokens, err error)
}

type handlers struct {
	services IServices
}

func New(s IServices) *handlers {
	return &handlers{
		services: s,
	}
}
