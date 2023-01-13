package controllers

import (
	"log"
	"net/http"

	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/services"
	"github.com/labstack/echo/v4"
)

type podcastController struct {
	ps services.PodcastService
}

type PodcastController interface {
	CreatePodcast(c echo.Context) error
}

func NewPodcastController(ps services.PodcastService) PodcastController {
	return &podcastController{ps}
}

// AddMember godoc
//
//	@Summary		Create Podcast
//	@Description	Adds a new podcast to the database
//	@Tags			Podcast
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	formData	string	true 	"Enter name"
//	@Param					type	formData	models.PodcastType	true 	"Choose a type"
//	@Param					description	formData	string	true 	"Enter description"
//	@Param					mediaURL	formData	string	true 	"Enter Media URL"
//	@Param					image	formData	file	true	"Upload Thumbnail"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Router			/v1/podcast/create [post]
func (pc *podcastController) CreatePodcast(c echo.Context) error {
	request := new(models.PodcastRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	file, err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		return err
	}

	err = pc.ps.CreatePodcast(*request, file)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}
