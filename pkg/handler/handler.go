package handler

import (
	"net/http"

	mdl "jwt/pkg/models"

	"github.com/labstack/echo/v4"
)

type handler struct {
	service mdl.IService
	storage mdl.IStorage
}

func New(service mdl.IService, storage mdl.IStorage) handler {
	return handler{
		service: service,
		storage: storage,
	}
}

func (h *handler) AuthSignIn(c echo.Context) error {
	signIn := mdl.SignIn{}
	if err := c.Bind(&signIn); err != nil {
		return err
	}
	rt, at, err := h.service.SignIn(signIn.Email, signIn.Password)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"refresh_token": string(rt), "access_token": string(at)})
}

func (h *handler) AuthSignUp(c echo.Context) error {
	signUp := mdl.SignUp{}
	if err := c.Bind(&signUp); err != nil {
		return err
	}
	err := h.service.SignUp(signUp.Email, signUp.Password)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) AuthRefresh(c echo.Context) error {
	rft := mdl.RefreshToken("test")
	rt, at, err := h.service.Refresh(rft)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{"refresh_token": string(rt), "access_token": string(at)})
}

func (h *handler) ApiSave(c echo.Context) error {
	// signUp := mdl.SignUp{}
	// if err := c.Bind(&signUp); err != nil {
	// 	return err
	// }
	ts := mdl.AppTokens{}
	err := h.service.SaveAppTokens(ts)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) ApiGive(c echo.Context) error {
	// signUp := mdl.SignUp{}
	// if err := c.Bind(&signUp); err != nil {
	// 	return err
	// }
	at := mdl.AccessToken("test")
	err := h.service.GiveAppTokens(at)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusOK)
}



