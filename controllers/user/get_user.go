package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/probe-nitt/probe-server/database"
	"github.com/probe-nitt/probe-server/middlewares"
	"github.com/probe-nitt/probe-server/models"
	"github.com/probe-nitt/probe-server/utils"
)

type GetUserRequest struct {
	ID string `json:"id" validate:"required"`
}

type GetUserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func GetUser(c echo.Context) error {
	req := new(GetUserRequest)

	if err := utils.ValidateRequest(c, req); err != nil {
		return middlewares.SendResponse(c, http.StatusBadRequest, "Invalid request")
	}

	db := database.GetDB()

	var user models.User

	db.Where("id = ?", req.ID).First(&user)

	if user.ID == 0 {
		return middlewares.SendResponse(c, http.StatusBadRequest, "User not found")
	}

	res := GetUserResponse{
		Username: user.Username,
		Name:     user.Name,
	}

	return middlewares.SendResponse(c, http.StatusOK, res)
}
