package router

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerRoutes(e *echo.Group) {
	origin := config.Origin
	url := echoSwagger.URL(origin + "/v1/admin/doc.json")
	e.GET("/admin/*", echoSwagger.EchoWrapHandler(url))
	e.File("/admin/index.html", "templates/html/swagger.html")
	//e.File("/admin/swagger-ui.css", "templates/css/swagger-ui.css")
}
