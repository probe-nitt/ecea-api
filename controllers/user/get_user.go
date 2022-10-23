package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/probe-nitt/probe-server/database"
	"github.com/probe-nitt/probe-server/models"
	"github.com/probe-nitt/probe-server/utils"
)

type GetUserRequest struct {
	Id string `json:"id" validate:"required"`
}

type GetUserResponse struct {
	Username string `json:"username"`
	Name     string `json:"name"`
}

func GetUser (c echo.Context) error {
	req := new(GetUserRequest)

	if err := utils.ValidateRequest(c, req); err != nil {
		return utils.SendResponse(c, http.StatusBadRequest, "Invalid request")
	}

	db := database.GetDB()

	var user models.User

	db.Where("id = ?", req.Id).First(&user)

	if user.ID == 0 {
		return utils.SendResponse(c, http.StatusBadRequest, "User not found")
	}

	res := GetUserResponse{
		Username: user.Username,
		Name:     user.Name,
	}

	return utils.SendResponse(c, http.StatusOK, res)
}