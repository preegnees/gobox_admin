package handler

import (
	"net/http"

	ms "jwt/pkg/models"
	ss "jwt/pkg/services"

	"github.com/labstack/echo/v4"
)

//go:generate mockgen -source=handler.go -destination=mock/mock.go

type IHandlers interface {
	AuthSignIn(c echo.Context) error
	AuthSignUp(c echo.Context) error
	AuthRefresh(c echo.Context) error
	ApiSaveAppTokens(c echo.Context) error
	ApiGiveAppTokens(c echo.Context) error
}

type handler struct {
	service ss.IService
}

var _ IHandlers = (*handler)(nil)

func New(service ss.IService) IHandlers {
	return &handler{
		service: service,
	}
}

func (h *handler) AuthSignIn(c echo.Context) error {
	signIn := ms.SignIn{}
	if err := c.Bind(&signIn); err != nil {
		return err
	}
	rt, at, err := h.service.SignIn(signIn.Username, signIn.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"refresh_token": rt, "access_token": at})
}

func (h *handler) AuthSignUp(c echo.Context) error {
	signUp := ms.SignUp{}
	if err := c.Bind(&signUp); err != nil {
		return err
	}
	err := h.service.SignUp(signUp.Username, signUp.Password, signUp.Email, signUp.EmailCode)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) AuthRefresh(c echo.Context) error {
	rft := "test"
	rt, at, err := h.service.Refresh(rft)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"refresh_token": rt, "access_token": at})
}

func (h *handler) ApiSaveAppTokens(c echo.Context) error {
	sat := ms.SaveAppTokens{}
	err := h.service.SaveAppTokens(sat.Username, sat.Tokens)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) ApiGiveAppTokens(c echo.Context) error {
	at := ms.GiveAppTokens{}
	ts, err := h.service.GiveAppTokens(at.Username)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string][]ms.AppToken{"tokens": ts})
}
