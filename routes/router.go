package routes

import (
	"github.com/labstack/echo/v4"
)

func Init(e *echo.Echo) {
	api := e.Group("/api")

	UserRoutes(api)
}
