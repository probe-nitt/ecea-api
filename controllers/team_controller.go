package controllers

import (
	"log"
	"net/http"

	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/services"
	"github.com/ecea-nitt/ecea-server/utils"
	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

type teamController struct {
	ts services.TeamService
}

type TeamController interface {
	AddMember(c echo.Context) error
	EditMemberImage(c echo.Context) error
	EditMemberName(c echo.Context) error
	EditMemberRole(c echo.Context) error
	EditMemberTeam(c echo.Context) error
	DeleteMember(c echo.Context) error
	GetAllMembers(c echo.Context) error
	GetMember(c echo.Context) error
}

func NewTeamController(ts services.TeamService) TeamController {
	return &teamController{ts}
}

// GetMember godoc
//
//	@Summary		Get a team member
//	@Description	Fetches a member and remove form Database
//	@Tags			Team
//	@Accept					json
//	@Produce		json
//	@Param					rollnumber	path string	true 	"Get member"
//	@Success		200	{object}    models.Members
//	@Failure		400	{object}	models.Error
//
// @Router			/v1/team/get/{rollnumber} [get]
func (tc *teamController) GetMember(c echo.Context) error {

	param := c.Param("rollnumber")

	rollNumber, err := utils.NumericValidator(param)

	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, rollNumber)
	}

	res, err := tc.ts.GetTeamMember(rollNumber)

	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Error Occurred")
	}

	return middlewares.Responder(c, http.StatusOK, &res)
}

// GetAllMembers godoc
//
//	@Summary		Get all team members
//	@Description	Fetches all the team members from Database
//	@Tags			Team
//	@Accept					json
//	@Produce		json
//	@Success		200	{object}    []models.Members
//	@Failure		400	{object}	models.Error
//	@Router			/v1/team/getall [get]
func (tc *teamController) GetAllMembers(c echo.Context) error {

	res, err := tc.ts.GetAllTeamMember()
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Conflict")
	}
	return middlewares.Responder(c, http.StatusOK, res)
}

// AddMember godoc
//
//	@Summary		Add a team member
//	@Description	Creates a new member and adds to Database
//	@Tags			Team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	formData	string	true 	"Enter name"
//	@Param					rollnumber	formData	string	true 	"Enter roll no"
//	@Param					role	formData	models.MemberRoles	true 	"Choose a role"
//	@Param					team	formData	models.MemberTeams	true 	"Choose a team"
//	@Param					image	formData	file	true	"Upload Image"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/team/add [post]
func (tc *teamController) AddMember(c echo.Context) error {

	request := new(models.MemberRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	log.Println(request.Team)
	log.Println(request.Role)
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

// EditMemberName godoc
//
//	@Summary		Edit a team member's name
//	@Description	Edits a member and updates to Database
//	@Tags			Team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	formData	string	true 	"Edit name"
//	@Param					rollnumber	formData	string	true 	"Enter roll no"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//	@Security 		ApiKeyAuth
//	@Router			/v1/team/edit/name [put]
func (tc *teamController) EditMemberName(c echo.Context) error {

	request := new(models.MemberRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err := tc.ts.EditTeamMemberName(*request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// EditMemberImage godoc
//
//	@Summary		Edit a team member's image
//	@Description	Edits a member and updates to Database
//	@Tags			Team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					rollnumber	formData	string	true 	"Enter roll no"
//	@Param					image	formData	file	true	"Edit Image"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//	@Security 		ApiKeyAuth
//	@Router			/v1/team/edit/image [put]
func (tc *teamController) EditMemberImage(c echo.Context) error {

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

	err = tc.ts.EditTeamMemberImage(*request, file)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// EditMemberRole godoc
//
//	@Summary		Edit a team member's role
//	@Description	Edits a member and updates to Database
//	@Tags			Team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					rollnumber	formData	string	true 	"Enter roll no"
//	@Param					role	formData	models.MemberRoles	true	"Change role"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//	@Security 		ApiKeyAuth
//	@Router			/v1/team/edit/role [put]
func (tc *teamController) EditMemberRole(c echo.Context) error {

	request := new(models.MemberRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err := tc.ts.EditTeamMemberRole(*request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// EditMemberTeam godoc
//
//	@Summary		Edit a team member's team
//	@Description	Edits a member and updates to Database
//	@Tags			Team
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					rollnumber	formData	string	true 	"Enter roll no"
//	@Param					team	formData	models.MemberTeams true	"Change team"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//	@Security 		ApiKeyAuth
//	@Router			/v1/team/edit/team [put]
func (tc *teamController) EditMemberTeam(c echo.Context) error {

	request := new(models.MemberRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	err := tc.ts.EditTeamMemberTeam(*request)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// DeleteMember godoc
//
//	@Summary		Delete a team member
//	@Description	Deletes a member and remove form Database
//	@Tags			Team
//	@Accept					json
//	@Produce		json
//	@Param					rollnumber	path string	true 	"Delete member"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
// @Security 		ApiKeyAuth
// @Router			/v1/team/delete/{rollnumber} [delete]
func (tc *teamController) DeleteMember(c echo.Context) error {

	param := c.Param("rollnumber")

	rollNumber, err := utils.NumericValidator(param)

	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusBadRequest, rollNumber)
	}

	err = tc.ts.RemoveTeamMember(rollNumber)

	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Error Occurred")
	}

	return middlewares.Responder(c, http.StatusOK, "Success")
}
