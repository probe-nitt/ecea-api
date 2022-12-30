package main

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/registry"
	"github.com/ecea-nitt/ecea-server/router"
	"github.com/labstack/echo/v4"
)

func main() {

	// Load config
	config.InitApp()

	// Register app
	reg := registry.NewRegistry(config.GetDB())

	// Create and Setup Echo Server
	server := echo.New()

	// Create Router
	router.NewRouter(server, reg.NewAppController())

	// Start Server
	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
