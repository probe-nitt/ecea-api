package repositories

import (
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
)

type seedRepository struct {
	db *gorm.DB
}

type SeedRepository interface {
	InsertTeams(teams []schemas.Team) error
	InsertRoles(roles []schemas.Role) error
	InsertAssetTypes(types []schemas.AssetType) error
}

func NewSeedRepository(db *gorm.DB) SeedRepository {
	return &seedRepository{db}
}

func (sr *seedRepository) InsertTeams(teams []schemas.Team) error {
	return sr.db.Create(&teams).Error

}

func (sr *seedRepository) InsertRoles(roles []schemas.Role) error {
	return sr.db.Create(&roles).Error
}

func (sr *seedRepository) InsertAssetTypes(types []schemas.AssetType) error {
	return sr.db.Create(&types).Error
}
