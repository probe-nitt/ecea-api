package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(e *echo.Echo, c controllers.AppController) {
	e.Use(middleware.CORS())
	e.Use(middlewares.Logger(e))
	e.Static("/static", "static")

	api := e.Group("/v1")

	UserRoutes(api, c.User)
	TeamRoutes(api, c.Team)
	StudyMaterialRoutes(api, c.StudyMaterial)
	SwaggerRoutes(api)
}
