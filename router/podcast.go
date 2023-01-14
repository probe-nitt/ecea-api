package router

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/labstack/echo/v4"
)

func PodcastRoutes(e *echo.Group, c controllers.PodcastController) {
	podcast := e.Group("/podcast")

	// Create
	podcast.POST("/create", c.CreatePodcast)

	// Read
	// podcast.GET("/getall", c.GetAllPodcasts)
	// podcast.GET("/getall/:type", c.GetPodcastByType)
	// podcast.GET("/get/:name", c.GetPodcastByName)

	// Update
	podcast.PUT("/edit/thumbnail", c.EditThumbnail)
	podcast.PUT("/edit/url", c.EditURL)
	podcast.PUT("/edit/description", c.EditDescription)

	// Delete
	podcast.DELETE("/delete/:name", c.DeletePodcast)
}
