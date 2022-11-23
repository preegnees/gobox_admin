package storage

import (
	"context"
	"core/dto"
	"testing"
)

func TestSaveUser(t *testing.T) {
	c := ConfPostgres{
		Cxt:      context.TODO(),
		User:     "docker",
		Password: "docker",
		Host:     "localhost",
		Port:     5431,
		DB:       "docker",
	}
	p, err := New(&c)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	user := dto.DTOUser{
		Username:     "username",
		PasswordHash: "pass",
		UserRole: "u",
		Email: "email@",
	}
	if err := p.SaveUser(ctx, &user); err != nil {
		panic(err)
	}
}

func TestFindUserByUsername(t *testing.T) {
	c := ConfPostgres{
		Cxt:      context.TODO(),
		User:     "docker",
		Password: "docker",
		Host:     "localhost",
		Port:     5431,
		DB:       "docker",
	}
	p, err := New(&c)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	username := "roman"
	user, err := p.FindUserByUsername(ctx, username)
	if err != nil {
		panic(err)
	}
	if (*user).Username != username {
		panic("(*user).Username != username")
	}
}

func TestSaveToken(t *testing.T) {
	c := ConfPostgres{
		Cxt:      context.TODO(),
		User:     "docker",
		Password: "docker",
		Host:     "localhost",
		Port:     5431,
		DB:       "docker",
	}
	p, err := New(&c)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	token := dto.DTOTokens{
		Username: "roman",
		Token: "token1",
	}
	err = p.SaveToken(ctx, &token)
	if err != nil {
		panic(err)
	}
}

func TestRemoveToken(t *testing.T) {
	c := ConfPostgres{
		Cxt:      context.TODO(),
		User:     "docker",
		Password: "docker",
		Host:     "localhost",
		Port:     5431,
		DB:       "docker",
	}
	p, err := New(&c)
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	token := dto.DTOTokens{
		Username: "roman",
		Token: "token1",
	}
	err = p.SaveToken(ctx, &token)
	if err != nil {
		
	}
	err = p.RemoveToken(ctx, &token)
	if err != nil {
		panic(err)
	}
}

