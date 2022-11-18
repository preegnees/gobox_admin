package handlers

import (
	"net/http"

	ms "jwt/pkg/models"

	"github.com/labstack/echo/v4"
)

func (h *handler) AuthSignIn(c echo.Context) error {

	signIn := ms.SignIn{}
	if err := c.Bind(&signIn); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(&signIn); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	ok, err := h.service.SignIn(signIn)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	if !ok {
		return c.NoContent(http.StatusUnauthorized)
	}

	at, rt, err := h.jwt.GenerateJWT(signIn.Username, signIn.Role)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, ms.AuthTokens{
		AccessToken:  at,
		RefreshToken: rt,
	})
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

func (h *handler) AuthSignOut(c echo.Context) error {
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
