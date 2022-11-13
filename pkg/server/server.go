package server

import (
	mdl "jwt/pkg/models"
	hdl "jwt/pkg/handler"
	srv "jwt/pkg/service"
	srg "jwt/pkg/storage"

	"github.com/labstack/echo/v4"
)

type server struct {

}

func New() {

}

var _ mdl.IServer = (*server)(nil)

func (s *server) Run() error {

	service := srv.New()
	storage := srg.New()
	handler := hdl.New(service, storage)

	e := echo.New()
	gAuth := e.Group("/auth")
	gAuth.POST("/sign-in", handler.AuthSignIn)
	gAuth.POST("/sign-up", handler.AuthSignUp)
	gAuth.POST("/refresh", handler.AuthRefresh)

	gApi := e.Group("/api")
	gApi.POST("/", handler.ApiSave)
	return e.Start("localhost:80")
}