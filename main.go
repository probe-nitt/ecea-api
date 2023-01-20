package main

import (
	"github.com/ecea-nitt/ecea-server/config"
	_ "github.com/ecea-nitt/ecea-server/docs" //swagger cli
	"github.com/ecea-nitt/ecea-server/registry"
	"github.com/ecea-nitt/ecea-server/router"
	"github.com/labstack/echo/v4"
)

//	@title			Probe Admin
//	@version		1.0
//	@description	Admin Panel for Probe Application.

//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Probe Webops
//	@contact.email	probe.eceanitt@gmail.com

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				Description for what is this security definition being used

func main() {

	// Load config
	config.InitApp()

	// Register app
	reg := registry.NewRegistry(config.GetDB())

	appController := reg.NewAppController()

	// Seed database
	//appController.Seed.SeedDB()

	// Create and Setup Echo Server
	server := echo.New()

	// Create Router
	router.NewRouter(server, appController)

	// Start Server
	server.Logger.Fatal(server.Start(":" + config.ServerPort))
}
