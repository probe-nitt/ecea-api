package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/labstack/echo/v4"
)

func TeamRoutes(e *echo.Group, c controllers.TeamController) {
	team := e.Group("/team")

	team.POST("/add", middlewares.Authorizer(c.AddMember))
	team.PUT("/edit/image", middlewares.Authorizer(c.EditMemberImage))
	team.PUT("/edit/name", middlewares.Authorizer(c.EditMemberName))
	team.PUT("/edit/role", middlewares.Authorizer(c.EditMemberRole))
	team.PUT("/edit/team", middlewares.Authorizer(c.EditMemberTeam))
	team.GET("/getall", c.GetAllMembers)
	team.GET("/get/:rollnumber", c.GetMember)
	team.DELETE("/delete/:rollnumber", middlewares.Authorizer(c.DeleteMember))

}
