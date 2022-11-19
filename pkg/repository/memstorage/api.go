package memstorage

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IMemStorage interface {
	SaveRefreshToken(ctx context.Context, token string) error
	CheckRefreshToken(ctx context.Context, token string) error
	DeleteRefreshToken(ctx context.Context, token string) error
}

type redis_ struct {
	client redis.Client
}

func New() (IMemStorage, error) {

	addr := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return nil, err
	}
	if addr == "" {
		return nil, errors.New("addr is empty")
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redis_{
		client: *rdb,
	}, nil
}
