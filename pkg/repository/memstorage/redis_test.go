package memstorage

import (
	"context"
	"os"
	"testing"
)

func TestRedis(t *testing.T) {

	os.Setenv("REDIS_HOST", "localhost:6379")
	os.Setenv("REDIS_PASSWORD", "")
	os.Setenv("REDIS_DB", "1")


	r, err := New()
	if err != nil {
		panic(err)
	}

	TOKEN := "token"
	if err := r.SaveRefreshToken(context.TODO(), TOKEN); err != nil {
		panic(err)
	}

	if err := r.CheckRefreshToken(context.TODO(), TOKEN); err != nil {
		panic(err)
	}

	if err := r.DeleteRefreshToken(context.TODO(), TOKEN); err != nil {
		panic(err)
	}
}