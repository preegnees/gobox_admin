package memstorage

import (
	dto "core/dto"
	models "core/models"
	repository "core/repository"

	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type redis_ struct {
	db *redis.Client
}

var _ repository.IMemStorage = (*redis_)(nil)

type ConfRedis struct {
	Addr     string
	Password string
	DB       int
}

func New(cnf *ConfRedis) *redis_ {
	if cnf == nil {
		cnf = &ConfRedis{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}

	}
	return &redis_{
		db: redis.NewClient(&redis.Options{
			Addr:     cnf.Addr,
			Password: cnf.Password,
			DB:       cnf.DB,
		}),
	}
}

func (r *redis_) SaveUserMetadata(ctx context.Context, metadata *dto.DTOUserMetadata) (err error) {

	jdata, err := json.Marshal(models.UserMetadata{
		Username:     metadata.Username,
		RefreshToken: metadata.RefreshToken,
	})
	if err != nil {
		return err
	}

	if err := r.db.Set(ctx, metadata.Fingerprint, jdata, 0).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redis_) GetUserMetadata(ctx context.Context, fingerprint string) (metadata *dto.DTOUserMetadata, err error) {

	m, err := r.db.Get(ctx, fingerprint).Result()
	if err != nil {
		return nil, err
	}

	um := models.UserMetadata{}
	if err := json.Unmarshal([]byte(m), &um); err != nil {
		return nil, err
	}
	return &dto.DTOUserMetadata{
		Fingerprint:  fingerprint,
		RefreshToken: um.RefreshToken,
		Username:     um.Username,
	}, nil
}
