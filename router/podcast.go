package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/labstack/echo/v4"
)

func PodcastRoutes(e *echo.Group, c controllers.PodcastController) {
	podcast := e.Group("/podcast")

	// Create
	podcast.POST("/create", middlewares.Authorizer(c.CreatePodcast))

	// Read
	podcast.GET("/get/all", c.GetAllPodcasts)
	podcast.GET("/getall/:type", c.GetPodcastByType)
	podcast.GET("/get/:episodeNo/:type", c.GetPodcast)

	// Update
	podcast.PUT("/edit/thumbnail", middlewares.Authorizer(c.EditThumbnail))
	podcast.PUT("/edit/url", middlewares.Authorizer(c.EditURL))
	podcast.PUT("/edit/description", middlewares.Authorizer(c.EditDescription))

	// Delete
	podcast.DELETE("/delete", middlewares.Authorizer(c.DeletePodcast))
}
