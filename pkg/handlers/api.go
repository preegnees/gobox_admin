package handlers

import (
	ss "jwt/pkg/services"
	jt "jwt/pkg/utils/jwt"

	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IHandlers interface {
	AuthSignIn(c echo.Context) error
	AuthSignUp(c echo.Context) error
	AuthSignOut(c echo.Context) error
	AuthRefresh(c echo.Context) error
	ApiSaveAppTokens(c echo.Context) error
	ApiGiveAppTokens(c echo.Context) error
}

type handler struct {
	service ss.IService
	jwt     jt.IJWT
}

var _ IHandlers = (*handler)(nil)

func New(service ss.IService, jwt jt.IJWT) IHandlers {

	return &handler{
		service: service,
		jwt:     jwt,
	}
}
