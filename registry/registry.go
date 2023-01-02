package registry

import (
	"github.com/ecea-nitt/ecea-server/controllers"
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

func (r *registry) NewAppController() controllers.AppController {
	return controllers.AppController{
		User: r.NewUserController(),
		Team: r.NewTeamController(),
		Seed: r.NewSeedController(),
	}
}
