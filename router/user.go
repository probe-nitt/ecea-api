package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group, c controllers.UserController) {
	user := e.Group("/user")

	user.POST("/signup", c.Register)
	user.GET("/verifyemail/:verificationCode", c.VerifyEmail)
}
