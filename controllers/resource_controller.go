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

type studyMaterialController struct {
	rs services.StudyMaterialService
}

type StudyMaterialController interface {
	AddStudyMaterial(c echo.Context) error
	GetCategoryStudyMaterials(c echo.Context) error
	GetStudyMaterial(c echo.Context) error
	UpdateStudyMaterialSubject(c echo.Context) error
	UpdateStudyMaterialURL(c echo.Context) error
	DeleteStudyMaterial(c echo.Context) error
}

func NewStudyMaterialController(rs services.StudyMaterialService) StudyMaterialController {
	return &studyMaterialController{rs}
}

// AddStudyMaterial godoc
//
//	@Summary		Add a study material
//	@Description	Creates a new study material and adds to Database
//	@Tags			StudyMaterial
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	        formData	string                  true 	"Enter material name"
//	@Param					subject	        formData	string	                true 	"Enter subject name"
//	@Param					code	    	formData	string	                true 	"Enter subject code"
//	@Param					category		formData	models.SubjectCategory	true 	"Choose a category"
//	@Param					document     	formData	file	                true	"Upload Document"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router		/v1/studymaterial/add [post]
func (rc *studyMaterialController) AddStudyMaterial(c echo.Context) error {
	request := new(models.StudyMaterialRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return err
	}
	file, err := c.FormFile("document")
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(request)
	err = rc.rs.CreateNewStudyMaterial(*request, file)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusConflict, "Conflict")
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// GetStudyMaterial godoc
//
//	@Summary		Get a study material
//	@Description	Fetches a study material and removes form database
//	@Tags			StudyMaterial
//	@Accept					json
//	@Produce		json
//	@Param					name	path string	true 	"Get study material"
//	@Success		200	{object}    models.StudyMaterials
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router			/v1/studymaterial/get/{name} [get]
func (rc *studyMaterialController) GetStudyMaterial(c echo.Context) error {
	rname := c.Param("name")
	res, err := rc.rs.GetStudyMaterial(rname)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Error Occurred")
	}

	return middlewares.Responder(c, http.StatusOK, &res)
}

// GetCategoryStudyMaterials godoc
//
//	@Summary		Get study materials from a category
//	@Description	Fetches all study materials belonging to a category from Database
//	@Tags			StudyMaterial
//	@Accept					json
//	@Produce		json
//	@Param					category	path models.SubjectCategory	true 	"Get study material from category"
//	@Success		200	{object}    []models.StudyMaterials
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router			/v1/studymaterial/getcat/{category} [get]
func (rc *studyMaterialController) GetCategoryStudyMaterials(c echo.Context) error {
	/*
		request := new(models.StudyMaterialRequest)
		if err := c.Bind(request); err != nil {
			return err
		}
	*/
	res, err := rc.rs.GetCategoryStudyMaterials(c.Param("category"))
	log.Println(c.Param("category"))
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Conflict")
	}
	return middlewares.Responder(c, http.StatusOK, res)
}

// UpdateStudyMaterialSubject godoc
//
//	@Summary		Updates a studymaterial's subject name
//	@Description	Updates a subject name and updates to Database
//	@Tags			StudyMaterial
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					subject	    formData	string	true 	"Edit name"
//	@Param					subjectCode	formData	string	true 	"Enter subject code"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router			/v1/studymaterial/edit/subject [put]
func (rc *studyMaterialController) UpdateStudyMaterialSubject(c echo.Context) error {

	request := new(models.StudyMaterialRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	/*
		subject := c.Param("subject")
		subjectCode := c.Param("subjectCode")
		if err := c.Bind(request); err != nil {
			log.Println(err)
			return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
		}
	*/
	log.Println(request.Subject, request.SubjectCode)
	err := rc.rs.EditStudyMaterialSubject(request.Subject, request.SubjectCode)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// UpdateStudyMaterialURL godoc
//
//	@Summary		Update a study material's document
//	@Description	Update a document and update it on Database
//	@Tags			StudyMaterial
//	@Accept					multipart/form-data
//	@Produce		json
//	@Param					name	    formData	string	true 	"Enter the name of material"
//	@Param					document	formData	file	true	"Edit Document"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router			/v1/studymaterial/edit/document [put]
func (rc *studyMaterialController) UpdateStudyMaterialURL(c echo.Context) error {

	request := new(models.StudyMaterialRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}
	/*
		name := c.Param("name")
		request := new(models.StudyMaterialRequest)
		if err := c.Bind(request); err != nil {
			log.Println(err)
			return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
		}
	*/
	file, err := c.FormFile("document")
	if err != nil {
		log.Println(err)
		return err
	}

	err = rc.rs.EditStudyMaterialURL(request.Name, file)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, err)
	}
	return middlewares.Responder(c, http.StatusOK, "Success")
}

// DeleteStudyMaterial godoc
//
//	@Summary		Deletes a study material
//	@Description	Deletes a study material along with its document and remove form Database
//	@Tags			StudyMaterial
//	@Accept					json
//	@Produce		json
//	@Param					name	path string	true 	"Delete study material"
//	@Success		200	{object}    string
//	@Failure		400	{object}	models.Error
//
//	@Security 		ApiKeyAuth
//	@Router			/v1/studymaterial/delete/{name} [delete]
func (rc *studyMaterialController) DeleteStudyMaterial(c echo.Context) error {
	rname := c.Param("name")

	err := rc.rs.RemoveStudyMaterial(rname)
	log.Println(rname)
	if err != nil {
		log.Println(color.RedString(err.Error()))
		return middlewares.Responder(c, http.StatusConflict, "Error Occurred")
	}

	return middlewares.Responder(c, http.StatusOK, "Success")
}
