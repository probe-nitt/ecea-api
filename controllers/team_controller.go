package controllers

import (
	"log"
	"net/http"

	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/services"
	"github.com/labstack/echo/v4"
)

type teamController struct {
	ts services.TeamService
}

type TeamController interface {
	AddMember(c echo.Context) error
}

func NewTeamController(ts services.TeamService) TeamController {
	return &teamController{ts}
}

// AddMember godoc
//
//	@Summary		Add a team member
//	@Description	adds a new member to Database
//	@Tags			team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					team	formData	models.MemberRequest	true 	"Add member"
//	@Param					image	formData	file	true	"member image"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//	@Router			/v1/team/add [post]
func (tc *teamController) AddMember(c echo.Context) error {

	request := new(models.MemberRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	file, err := c.FormFile("image")
	if err != nil {
		log.Println(err)
		return err
	}

	err = tc.ts.CreateTeamMember(*request, file)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusConflict, "Conflict")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}
