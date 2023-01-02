package middlewares

import (
	"net/http"

	"github.com/ecea-nitt/ecea-server/config"
	"github.com/labstack/echo/v4"
)

func Authorizer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if len(authHeader) == 0 {
			return Responder(c, http.StatusBadRequest, "please set Header Authorization")
		}
		if authHeader != config.Admin {
			return Responder(c, http.StatusUnauthorized, "unauthorized entry")
		}
		return next(c)
	}
}
