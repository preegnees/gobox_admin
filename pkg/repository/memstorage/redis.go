package memstorage

import (
	"context"
)

func (r *redis_) SaveRefreshToken(ctx context.Context, token string) error {

	return r.client.Set(ctx, token, "", 0).Err()
}

func (r *redis_) CheckRefreshToken(ctx context.Context, token string) error {

	return r.client.Get(ctx, token).Err()
}

func (r *redis_) DeleteRefreshToken(ctx context.Context, token string) error {

	return r.client.Del(ctx, token).Err()
}