package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/services"
)

func (r *registry) NewUserController() controllers.UserController {
	return controllers.NewUserController(r.NewUserService(), r.NewMailService())
}

func (r *registry) NewUserService() services.UserService {
	return services.NewUserService(r.NewUserRepository())
}

func (r *registry) NewUserRepository() repositories.UserRepository {
	return repositories.NewUserRepository(r.db)
}
