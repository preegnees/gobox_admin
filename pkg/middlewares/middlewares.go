package middlewares

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func (m *middlewares) CheckJWT(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header

		if len(header["authorization"]) == 0 {
			return c.NoContent(http.StatusUnauthorized)
		}

		auth := header["authorization"][0]
		if auth == "" {
			return c.NoContent(http.StatusUnauthorized)
		}

		token := strings.Split(auth, " Bearer")[1]
		if token == "" {
			return c.NoContent(http.StatusUnauthorized)
		}

		if !m.jsonWebToken.CheckJwt(token) {
			c.NoContent(http.StatusUnauthorized)
		}

		return next(c)
	}
}
