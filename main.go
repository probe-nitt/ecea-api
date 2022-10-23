package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/probe-nitt/probe-server/config"
	"github.com/probe-nitt/probe-server/database"
	"github.com/probe-nitt/probe-server/models"
	"github.com/probe-nitt/probe-server/routes"
	"github.com/probe-nitt/probe-server/utils"
)

func main() {

	// Load config
	config.InitConfig()

	// Connect to database
	database.ConnectDB()

	// Migrate database
	models.MigrateDB()

	// Create and Setup Echo Server
	server := echo.New()
	utils.InitLogger(server)
	utils.InitValidator(server)
	server.Use(middleware.CORS())

	// Routes
	routes.Init(server)

	// Start server
	server.Logger.Fatal(server.Start(":" + config.GetConfig().ServerPort))
}
