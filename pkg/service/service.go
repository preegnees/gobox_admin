package service

import (
	mdl "jwt/pkg/models"
)

type service struct {}

var _ mdl.IService = (*service)(nil)

func New() mdl.IService {
	return &service{}
}

func (s *service) SignUp(mdl.Email, mdl.Password) error {
	return nil
}
func (s *service) SignIn(mdl.Email, mdl.Password) (mdl.RefreshToken, mdl.AccessToken, error){
	return "", "", nil
}
func (s *service) Refresh(mdl.RefreshToken) (mdl.RefreshToken, mdl.AccessToken, error){
	return "", "", nil
}
func (s *service) SaveAppTokens(mdl.AppTokens) error{
	return nil
}
func (s *service) GiveAppTokens(mdl.AccessToken) error{
	return nil
}