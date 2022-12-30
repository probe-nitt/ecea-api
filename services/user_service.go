package services

import (
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/schemas"
	"github.com/ecea-nitt/ecea-server/utils"
)

type userService struct {
	repo repositories.UserRepository
}

type UserService interface {
	RegisterUser(request models.RegisterRequest) (schemas.User, error)
	AddVerificationCode(user *schemas.User, code string) error
	CompleteVerification(code string) error
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo}
}

func (us *userService) RegisterUser(request models.RegisterRequest) (schemas.User, error) {
	var user schemas.User
	name, err := utils.NameValidator(request.Name)
	if err != nil {
		return user, err
	}

	var email string

	email, err = utils.EmailValidator(request.Email)
	if err != nil {
		return user, err
	}

	password := utils.PasswordHasher(request.Password)

	user = schemas.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err = us.repo.CreateUser(&user)

	return user, err
}

func (us *userService) AddVerificationCode(user *schemas.User, code string) error {
	return us.repo.UpdateVerificationCode(user, code)
}

func (us *userService) CompleteVerification(code string) error {
	return us.repo.UpdateVerification(code)
}
