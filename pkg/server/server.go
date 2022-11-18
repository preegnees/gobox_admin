package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	hs "jwt/pkg/handlers"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

type IServer interface {
	Run() error
}

type server struct {
	ctx      context.Context
	address  string
	handlers hs.IHandlers
}

func New(ctx context.Context, address string, handlers hs.IHandlers) IServer {
	return &server{
		ctx:      ctx,
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
	gAuth.POST("/sign-out", s.handlers.AuthSignOut)
	gAuth.POST("/refresh", s.handlers.AuthRefresh)

	gApi := e.Group("/api")
	gApi.POST("/save", s.handlers.ApiSaveAppTokens)
	gApi.GET("/give", s.handlers.ApiGiveAppTokens)

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
