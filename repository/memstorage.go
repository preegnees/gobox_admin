package repository

import (
	"context"
	dto "core/dto"
)

func (r *repository) SaveUserMetadata(ctx context.Context, metadata *dto.DTOUserMetadata) (err error) {
	if err := r.memStorage.SaveUserMetadata(ctx, metadata); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserMetadata(ctx context.Context, username string) (metadata *dto.DTOUserMetadata, err error) {
	metadata, err = r.memStorage.GetUserMetadata(ctx, username)
	if err != nil {
		return nil, err
	}
	return metadata, nil
}
