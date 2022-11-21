package storage

import (
	"context"
	"core/dto"
	"testing"
)

func TestSaveUser(t *testing.T) {
	c := ConfPostgres{
		Cxt: context.TODO(),
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     5431,
		DB:       "postgres",
	}
	p, err := New(&c)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	user := dto.DTOUser{
		Username: "username",
		Password: "pass",
	}
	if err := p.SaveUser(ctx, &user); err != nil {
		panic(err)
	}
}