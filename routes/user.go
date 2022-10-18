package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
