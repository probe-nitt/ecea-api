package controllers

import (
	"log"

	"github.com/ecea-nitt/ecea-server/services"
	"github.com/fatih/color"
)

type AppController struct {
	User          interface{ UserController }
	Team          interface{ TeamController }
	Seed          interface{ SeedController }
	StudyMaterial interface{ StudyMaterialController }
}

type seedController struct {
	seed services.SeederService
}

type SeedController interface {
	SeedDB() error
}

func NewSeedController(seed services.SeederService) SeedController {
	return &seedController{seed}
}

func (s *seedController) SeedDB() error {
	err := s.seed.AssetTypesSeeder()
	if err != nil {
		log.Panic(err)
		return err
	}
	err = s.seed.TeamsSeeder()
	if err != nil {
		log.Panic(err)
		return err
	}
	err = s.seed.RolesSeeder()
	if err != nil {
		log.Panic(err)
		return err
	}
	err = s.seed.StudyMaterialSeeder()
	if err != nil {
		log.Panic(err)
		return err
	}
	log.Println(color.GreenString(" Seeded Successfully"))
	return nil
}
