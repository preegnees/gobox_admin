package storage

import (
	mdl "jwt/pkg/models"
)

type storage struct {}

var _ mdl.IStorage = (*storage)(nil)

func New() mdl.IStorage {
	return &storage{}
}

func (s *storage) CheckUser(mdl.Ctx, mdl.Email, mdl.Password) error {
	return nil
}
func (s *storage) SaveRefreshToken(mdl.Ctx, mdl.Email, mdl.RefreshToken) error {
	return nil
}
func (s *storage) SaveAppTokens(mdl.Ctx, mdl.Email, mdl.AppTokens) error {
	return nil
}
func (s *storage) GiveAppTokens(mdl.Ctx, mdl.Email) (mdl.AppTokens, error) {
	return mdl.AppTokens{}, nil
}