package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/probe-nitt/probe-server/config"
	"github.com/probe-nitt/probe-server/routes"
)

func main() {

	server := echo.New()
	config.InitLogger(server)
	server.Use(middleware.CORS())
	routes.Init(server)
	server.Logger.Fatal(server.Start(":" + config.GetConfig().ServerPort))
}
