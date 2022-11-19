package middlewares

import (
	jwt "jwt/pkg/utils/jwt"

	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=api.go -destination=mock/mock.go

type IMiddleware interface {
	CheckJWT(next echo.HandlerFunc) echo.HandlerFunc
}

type middlewares struct {
	jsonWebToken jwt.IJWT
}

func New(jwt jwt.IJWT) IMiddleware {
	return &middlewares{
		jsonWebToken: jwt,
	}
}
