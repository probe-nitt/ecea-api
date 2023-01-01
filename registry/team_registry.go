package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/services"
)

func (r *registry) NewTeamController() controllers.TeamController {
	return controllers.NewTeamController(r.NewTeamService())
}

func (r *registry) NewTeamService() services.TeamService {
	return services.NewTeamService(r.NewTeamRepository())
}

func (r *registry) NewTeamRepository() repositories.TeamRepository {
	return repositories.NewTeamRepository(r.db)
}
