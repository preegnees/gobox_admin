package handlers

import (
	services "jwt/pkg/services"
	jwt "jwt/pkg/utils/jwt"

	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IHandlers interface {
	AuthSignIn(c echo.Context) error
	AuthSignUp(c echo.Context) error
	AuthSignOut(c echo.Context) error
	AuthRefresh(c echo.Context) error
	ApiSaveAppData(c echo.Context) error
	ApiGetAppData(c echo.Context) error
}

type handler struct {
	service services.IService
	jwt     jwt.IJWT
	
}

var _ IHandlers = (*handler)(nil)

func New(service services.IService, jwt jwt.IJWT) IHandlers {

	return &handler{
		service: service,
		jwt:     jwt,
	}
}
