package controllers

import (
	"log"
	"net/http"

	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/services"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

type podcastController struct {
	ps services.PodcastService
}

type PodcastController interface {
	CreatePodcast(c echo.Context) error
	EditThumbnail(c echo.Context) error
	EditURL(c echo.Context) error
	EditDescription(c echo.Context) error
	DeletePodcast(c echo.Context) error
}

func NewPodcastController(ps services.PodcastService) PodcastController {
	return &podcastController{ps}
}

// CreatePodcast godoc
//
//	@Summary		Create Podcast
//	@Description	Adds a new podcast to the database
//	@Tags			Podcast
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	formData	string	true 	"Enter name"
//	@Param					episodeNo	formData	uint	true 	"Enter episode number"
//	@Param					type	formData	models.PodcastType	true 	"Choose a type"
//	@Param					description	formData	string	true 	"Enter description"
//	@Param					mediaURL	formData	string	true 	"Enter Media URL"
//	@Param					image	formData	file	true	"Upload Thumbnail"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/podcast/create [post]
func (pc *podcastController) CreatePodcast(c echo.Context) error {
	request := new(models.PodcastRequest)
	if err := c.Bind(request); err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	file, err := c.FormFile("image")
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return err
	}

	err = pc.ps.CreatePodcast(*request, file)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}

// EditThumbnail godoc
//
//	@Summary		Edit Thumbnail
//	@Description	Edits the thumbnail of a podcast
//	@Tags			Podcast
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					episodeNo	formData	uint	true 	"Enter episode number"
//
// @Param 					type 	formData	models.PodcastType	true 	"Choose a type"
//
//	@Param					image	formData	file	true	"Upload Thumbnail"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/podcast/edit/thumbnail [put]
func (pc *podcastController) EditThumbnail(c echo.Context) error {
	request := new(models.PodcastRequest)
	if err := c.Bind(request); err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	file, err := c.FormFile("image")
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return err
	}

	err = pc.ps.EditThumbnail(*request, file)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}

// EditURL godoc
//
//	@Summary		Edit URL
//	@Description	Edits the media url of a podcast
//	@Tags			Podcast
//	@Accept		multipart/form-data
//	@Produce		json
//	@Param					episodeNo	formData	uint	true 	"Enter episode number"
//
// @Param 					type 	formData	models.PodcastType	true 	"Choose a type"
//
//	@Param					mediaURL	formData	string	true 	"Enter Media URL"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/podcast/edit/url [put]
func (pc *podcastController) EditURL(c echo.Context) error {
	request := new(models.PodcastRequest)
	if err := c.Bind(request); err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err := pc.ps.EditURL(*request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}

// EditDescription godoc
//
//	@Summary		Edit Description
//	@Description	Edits the description of a podcast
//	@Tags			Podcast
//	@Accept		multipart/form-data
//	@Produce		json
//	@Param					episodeNo	formData	uint	true 	"Enter episode number"
//
// @Param 					type 	formData	models.PodcastType	true 	"Choose a type"
//
//	@Param					description	formData	string	true 	"Enter description"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/podcast/edit/description [put]
func (pc *podcastController) EditDescription(c echo.Context) error {
	request := new(models.PodcastRequest)
	if err := c.Bind(request); err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err := pc.ps.EditDescription(*request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}

// DeletePodcast godoc
//
//	@Summary		Delete Podcast
//	@Description	Deletes a podcast
//	@Tags			Podcast
//	@Accept		multipart/form-data
//	@Produce		json
//	@Param					episodeNo	formData	uint	true 	"Enter episode number"
//
// @Param 					type 	formData	models.PodcastType	true 	"Choose a type"
//
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/podcast/delete [delete]
func (pc *podcastController) DeletePodcast(c echo.Context) error {
	request := new(models.PodcastRequest)

	if err := c.Bind(request); err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
	}

	err := pc.ps.DeletePodcast(*request)

	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}

	return middlewares.Responder(c, http.StatusOK, http.StatusText(http.StatusOK))
}
