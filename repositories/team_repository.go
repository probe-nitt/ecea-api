package repositories

import (
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
)

type teamRepository struct {
	db *gorm.DB
}

type TeamRepository interface {
	FindTeamByName(name string) (uint, error)
	FindRoleByName(name string) (uint, error)
	InsertAsset(name string) (uint, error)
	InsertMember(member *schemas.Member) error
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db}
}

func (tr *teamRepository) FindTeamByName(name string) (uint, error) {
	var team schemas.Team
	res := tr.db.Where("name = ?", name).First(&team)
	if res.Error != nil {
		return 0, res.Error
	}
	return team.ID, nil
}

func (tr *teamRepository) FindRoleByName(name string) (uint, error) {
	var role schemas.Role
	res := tr.db.Where("name = ?", name).First(&role)
	if res.Error != nil {
		return 0, res.Error
	}
	return role.ID, nil
}

func (tr *teamRepository) InsertAsset(name string) (uint, error) {
	var assetType schemas.AssetType
	if err := tr.db.Where("name = ?", schemas.Image).First(&assetType).Error; err != nil {
		return 0, err
	}
	asset := schemas.Asset{
		Name:        name,
		AssetTypeID: assetType.ID,
	}
	if err := tr.db.Create(&asset).Error; err != nil {
		return 0, err
	}
	return asset.ID, nil
}

func (tr *teamRepository) InsertMember(member *schemas.Member) error {
	return tr.db.Create(member).Error
}
