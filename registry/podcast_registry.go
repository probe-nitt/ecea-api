package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/services"
)

func (r *registry) NewPodcastController() controllers.PodcastController {
	return controllers.NewPodcastController(r.NewPodcastService())
}

func (r *registry) NewPodcastService() services.PodcastService {
	return services.NewPodcastService(r.NewPodcastRepository())
}

func (r *registry) NewPodcastRepository() repositories.PodcastRepository {
	return repositories.NewPodcastRepository(r.db)
}
