package handlers

import (
	"net/http"
	"time"

	models "jwt/pkg/models"

	"github.com/labstack/echo/v4"
)

func (h *handler) AuthSignIn(c echo.Context) error {

	signIn := models.SignIn{}
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

	cooke := new(http.Cookie)
	cooke.Name = "refresh_token"
	cooke.Value = rt
	cooke.HttpOnly = true
	c.SetCookie(cooke)

	header := c.Response().Header()
	header.Add("Authorization", "Bearer "+at)

	return c.NoContent(http.StatusOK)
}

func (h *handler) AuthSignUp(c echo.Context) error {

	signUp := models.SignUp{}
	if err := c.Bind(&signUp); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(&signUp); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := h.service.SignUp(signUp)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) AuthSignOut(c echo.Context) error {

	rToken, err := c.Cookie("refresh_token")
	if err != nil || rToken.Value == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	err = h.service.SignOut(rToken.Value)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	cooke := new(http.Cookie)
	cooke.Name = rToken.Name
	cooke.Value = ""
	cooke.Expires = time.Unix(0, 0)
	cooke.HttpOnly = true
	c.SetCookie(cooke)

	return c.NoContent(http.StatusOK)
}

// не протестировано
func (h *handler) AuthRefresh(c echo.Context) error {

	rToken, err := c.Cookie("refresh_token")
	if err != nil || rToken.Value == "" {
		return c.NoContent(http.StatusBadRequest)
	}

	// убрать Refresh пихать в redis // тут где то создать токены и сохарнить
	err = h.service.SaveRefreshToken(rToken.Value)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	cooke := new(http.Cookie)
	cooke.Name = rToken.Name
	// cooke.Value = rt
	cooke.HttpOnly = true
	c.SetCookie(cooke)

	// header := c.Response().Header()
	// header.Add("Authorization", "Bearer "+at)

	return c.NoContent(http.StatusOK)
}

func (h *handler) ApiSaveAppData(c echo.Context) error {

	appData := models.AppData{}
	if err := c.Bind(&appData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(&appData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	err := h.service.SaveAppData(appData.Username, appData.Tokens)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusOK)
}

func (h *handler) ApiGetAppData(c echo.Context) error {

	appData := models.AppData{}
	if err := c.Bind(&appData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(&appData); err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	tokens, err := h.service.GetAppData(appData.Username)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	appData.Tokens = tokens

	return c.JSON(http.StatusOK, appData)
}
