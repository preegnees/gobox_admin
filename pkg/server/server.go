package server

import (
	hs "jwt/pkg/handlers"

	"github.com/labstack/echo/v4"
)

type IServer interface {
	Run() error
}

type server struct {
	address  string
	handlers hs.IHandlers
}

func New(address string, handlers hs.IHandlers) IServer {
	return &server{
		address:  address,
		handlers: handlers,
	}
}

var _ IServer = (*server)(nil)

func (s *server) Run() error {

	e := echo.New()
	gAuth := e.Group("/auth")
	gAuth.POST("/sign-in", s.handlers.AuthSignIn)
	gAuth.POST("/sign-up", s.handlers.AuthSignUp)
	gAuth.POST("/refresh", s.handlers.AuthRefresh)

	gApi := e.Group("/api")
	gApi.POST("/save", s.handlers.ApiSaveAppTokens)
	gApi.GET("/give", s.handlers.ApiGiveAppTokens)

	return e.Start(s.address)
}
