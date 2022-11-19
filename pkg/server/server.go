package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	handlers "jwt/pkg/handlers"
	middlewares "jwt/pkg/middlewares"
	validator_ "jwt/pkg/utils/validator"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/sync/errgroup"
)

type IServer interface {
	Run() error
}

type server struct {
	ctx         context.Context
	address     string
	handlers    handlers.IHandlers
	middlewares middlewares.IMiddleware
}

func New(ctx context.Context, address string, handlers handlers.IHandlers, middlewares middlewares.IMiddleware) IServer {
	return &server{
		ctx:         ctx,
		address:     address,
		handlers:    handlers,
		middlewares: middlewares,
	}
}

var _ IServer = (*server)(nil)

func (s *server) Run() error {

	e := echo.New()

	e.Validator = &validator_.CustomValidator{
		Validator: validator.New(),
	}

	e.Use(middleware.Logger())

	gAuth := e.Group("/auth")
	gAuth.POST("/sign-in", s.handlers.AuthSignIn)
	gAuth.POST("/sign-up", s.handlers.AuthSignUp)
	gAuth.POST("/sign-out", s.handlers.AuthSignOut)
	gAuth.POST("/refresh", s.handlers.AuthRefresh)

	gApi := e.Group("/api")
	gApi.Use(s.middlewares.CheckJWT)
	gApi.POST("/save", s.handlers.ApiSaveAppData)
	gApi.GET("/get", s.handlers.ApiGetAppData)

	g := new(errgroup.Group)

	g.Go(func() error {
		if err := e.Start(s.address); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			return err
		}
		return nil
	})

	select {
	case <-s.ctx.Done():
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			return err
		}
		if err := g.Wait(); err != nil {
			return err
		}
		return nil
	}
}
