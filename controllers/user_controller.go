package controllers

import (
	"log"
	"net/http"

	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/middlewares"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/services"
	"github.com/ecea-nitt/ecea-server/utils"
	"github.com/labstack/echo/v4"
)

type userController struct {
	us services.UserService
	ms services.MailService
}

type UserController interface {
	Register(c echo.Context) error
	VerifyEmail(c echo.Context) error
}

func NewUserController(us services.UserService, ms services.MailService) UserController {
	return &userController{us, ms}
}

// Register godoc
//
//	@Summary		Register an user
//	@Description	register an user
//	@Tags			user
//	@Accept					json
//	@Produce		json
//	@Param					user	body		models.RegisterRequest	true	"Add user"
//	@Success		200	{object}	models.RegisterResponse
//	@Failure		400	{object}	models.Error
//	@Router			/user/signup [post]
func (uc *userController) Register(c echo.Context) error {
	request := new(models.RegisterRequest)
	if err := c.Bind(request); err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Bad Request")
	}

	user, err := uc.us.RegisterUser(*request)

	if err != nil {
		log.Fatalln(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Register")
	}

	code := utils.GenerateVerificationCode()
	code = code + user.Email

	emailData := &models.EmailData{
		Name:    user.Name,
		URL:     config.Origin + "/v1/user/verifyemail/" + code,
		Subject: "user verification",
	}

	err = uc.us.AddVerificationCode(&user, code)

	if err != nil {
		log.Fatalln(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Register")
	}

	err = uc.ms.MailUser(user, *emailData)

	if err != nil {
		log.Fatalln(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Send Email")
	}

	return middlewares.Responder(c, http.StatusOK, "Check Your Email for verification mail")
}

// VerifyEmail godoc
//
//	@Summary		verify an user
//	@Description	verify an user by sending email
//	@Tags			user
//	@Accept					json
//	@Produce		json
//	@Param					verificationCode	path  string true "Verify user"
//	@Success		200	{object}	string
//	@Failure		400	{object}	models.Error
//	@Router			/user/verifyemail/{verificationCode} [post]
func (uc *userController) VerifyEmail(c echo.Context) error {

	code := c.Param("verificationCode")
	log.Println("VerificationCode", code)
	err := uc.us.CompleteVerification(code)
	if err != nil {
		log.Println(err)
		return middlewares.Responder(c, http.StatusBadRequest, "Unable to Verify")
	}
	return middlewares.Responder(c, http.StatusOK, "verified successfully")

}
