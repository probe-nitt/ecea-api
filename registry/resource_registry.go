package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/services"
)

func (r *registry) NewStudyMaterialController() controllers.StudyMaterialController {
	return controllers.NewStudyMaterialController(r.NewStudyMaterialService())
}

func (r *registry) NewStudyMaterialService() services.StudyMaterialService {
	return services.NewStudyMaterialService(r.NewStudyMaterialRepository())
}

func (r *registry) NewStudyMaterialRepository() repositories.StudyMaterialRepository {
	return repositories.NewStudyMaterialRepository(r.db)
}
