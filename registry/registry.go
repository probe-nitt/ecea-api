package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
	"github.com/ecea-nitt/ecea-server/services"
	"gorm.io/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppController() controllers.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewMailService() services.MailService {
	return services.NewMailService()
}

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		User: r.NewUserController(),
	}
}
