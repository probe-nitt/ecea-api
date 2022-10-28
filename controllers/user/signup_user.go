package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/probe-nitt/probe-server/database"
	"github.com/probe-nitt/probe-server/middlewares"
	"github.com/probe-nitt/probe-server/models"
	"github.com/probe-nitt/probe-server/utils"
)

type SignupRequest struct {
	Username string `json:"username" validate:"required"`
	Name     string `json:"name" validate:"required"`
	RollNo   string `json:"roll_no" validate:"required"`
}

func SignupUser(c echo.Context) error {
	var req SignupRequest

	if err := utils.ValidateRequest(c, req); err != nil {
		return middlewares.SendResponse(c, http.StatusBadRequest, "Invalid request")
	}

	newUser := models.User{
		Username: req.Username,
		Name:     req.Name,
		RollNo:   req.RollNo,
	}

	db := database.GetDB()

	if err := db.Create(&newUser).Error; err != nil {
		return middlewares.SendResponse(c, http.StatusBadRequest, "User already exists")
	}

	return middlewares.SendResponse(c, http.StatusOK, "User created successfully")
}
