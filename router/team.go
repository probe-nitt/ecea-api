package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/labstack/echo/v4"
)

func TeamRoutes(e *echo.Group, c controllers.TeamController) {
	team := e.Group("/team")

	team.POST("/add", c.AddMember)
}
