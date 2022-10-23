package routes

import (
	"github.com/labstack/echo/v4"
	user_controllers "github.com/probe-nitt/probe-server/controllers/user"
)

func UserRoutes(e *echo.Group) {
	user := e.Group("/user")

	user.POST("/signup", user_controllers.SignupUser)
	user.POST("/get", user_controllers.GetUser)
}
