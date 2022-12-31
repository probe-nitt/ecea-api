package main

import (
	"github.com/ecea-nitt/ecea-server/config"
	_ "github.com/ecea-nitt/ecea-server/docs" //swagger cli
	"github.com/ecea-nitt/ecea-server/registry"
	"github.com/ecea-nitt/ecea-server/router"
	"github.com/labstack/echo/v4"
)

//	@title			Probe API
//	@version		1.0
//	@description	API for Probe Application.

//	@termsOfService	http://swagger.io/terms/

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

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
