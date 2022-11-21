package repository

import (
	"context"
	dto "core/dto"
)

func (r *repository) SaveUser(ctx context.Context, user *dto.DTOUser) (err error) {
	if err := r.storage.SaveUser(ctx, user); err != nil {
		return err
	}
	return nil
}	

func (r *repository) SaveToken(ctx context.Context, token *dto.DTOTokens) (err error) {
	if err := r.storage.SaveToken(ctx, token); err != nil {
		return err
	}
	return nil
}

func (r *repository) RemoveToken(ctx context.Context, token *dto.DTOTokens) (err error) {
	if err := r.storage.RemoveToken(ctx, token); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetTokens(ctx context.Context, username string) (tokens *[]dto.DTOTokens, err error) {
	tokens, err = r.storage.GetTokens(ctx, username)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}