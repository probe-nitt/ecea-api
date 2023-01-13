package services

import (
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/repositories"
	"github.com/ecea-nitt/ecea-server/schemas"
)

type seederService struct {
	repo repositories.SeedRepository
}

type SeederService interface {
	AssetTypesSeeder() error
	TeamsSeeder() error
	RolesSeeder() error
	PodcastTypesSeeder() error
}

var teams = []schemas.Team{
	{
		Name: string(models.Webops),
	},
	{
		Name: string(models.Events),
	},
}

var roles = []schemas.Role{
	{
		Name: string(models.Chairperson),
	},
	{
		Name: string(models.OCBoy),
	},
	{
		Name: string(models.OCGirl),
	},
	{
		Name: string(models.Treasurer),
	},
	{
		Name: string(models.Head),
	},
	{
		Name: string(models.DeputyHead),
	},
	{
		Name: string(models.Manager),
	},
	{
		Name: string(models.Coordinator),
	},
}

var assetTypes = []schemas.AssetType{
	{
		Name: string(schemas.Image),
	},
	{
		Name: string(schemas.Document),
	},
}

var podcastTypes = []schemas.PodcastType{
	{
		Name: string(schemas.CareerPath),
	},
	{
		Name: string(schemas.GuestLecture),
	},
}

func NewSeeder(repo repositories.SeedRepository) SeederService {
	return &seederService{repo}
}

func (s *seederService) TeamsSeeder() error {
	return s.repo.InsertTeams(teams)
}

func (s *seederService) RolesSeeder() error {
	return s.repo.InsertRoles(roles)
}

func (s *seederService) AssetTypesSeeder() error {
	return s.repo.InsertAssetTypes(assetTypes)
}

func (s *seederService) PodcastTypesSeeder() error {
	return s.repo.InsertPodcastTypes(podcastTypes)
}
