package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Logger(server *echo.Echo) echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${status} | ${method} ${uri} \t | ${latency_human}\n",
		Output: server.Logger.Output(),
	})
}
