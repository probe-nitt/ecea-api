package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/services"
)

func (r *registry) NewSeedController() controllers.SeedController {
	return controllers.NewSeedController(r.NewSeeder())
}

func (r *registry) NewSeeder() services.SeederService {
	return services.NewSeeder(r.NewSeedRepository())
}

func (r *registry) NewSeedRepository() repositories.SeedRepository {
	return repositories.NewSeedRepository(r.db)
}
