package repositories

import (
	"github.com/ecea-nitt/ecea-server/config"
	"github.com/ecea-nitt/ecea-server/models"
	"github.com/ecea-nitt/ecea-server/schemas"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type teamRepository struct {
	db *gorm.DB
}

type TeamRepository interface {
	FindAllMembers() ([]models.Members, error)
	FindMemberByRollNo(rollNo string) (schemas.Member, error)
	FindTeamByName(name string) (uint, error)
	FindRoleByName(name string) (uint, error)
	InsertAsset(name string) (uint, error)
	UpdateAsset(id uint, name string) error
	InsertMember(member *schemas.Member) error
	UpdateMember(member *schemas.Member) error
	DeleteMember(rollNo string) error
	FindMember(rollNo string) (models.Members, error)
}

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepository{db}
}

func (tr *teamRepository) FindMemberByRollNo(rollNo string) (schemas.Member, error) {
	var member schemas.Member
	err := tr.db.Preload(clause.Associations).Where("roll_no = ?", rollNo).First(&member).Error
	return member, err
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

func (tr *teamRepository) UpdateAsset(id uint, name string) error {
	return tr.db.Model(&schemas.Asset{}).Where("id = ?", id).Update("name", name).Error
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

func (tr *teamRepository) UpdateMember(member *schemas.Member) error {
	return tr.db.Model(&member).Where("id = ?", member.ID).Updates(&member).Error
}

func (tr *teamRepository) DeleteMember(rollNo string) error {
	return tr.db.Unscoped().Where("roll_no = ?", rollNo).Delete(&schemas.Member{}).Error
}

func (tr *teamRepository) FindMember(rollNo string) (models.Members, error) {
	var member models.Members
	err := tr.db.Table("members").Select(
		`members.name,members.roll_no,roles.name as role,teams.name as team,
		CONCAT(?::text,'/static/images','/',assets.name) as image_url`, config.Origin,
	).Joins(
		"JOIN assets on assets.id = members.asset_id").Joins(
		"JOIN teams on teams.id = members.team_id").Joins(
		"JOIN roles on roles.id = members.role_id").Where(
		"roll_no = ?", rollNo).First(
		&member).Error

	return member, err
}

func (tr *teamRepository) FindAllMembers() ([]models.Members, error) {
	var members []models.Members
	if err := tr.db.Table("members").Select(
		`members.name,members.roll_no,roles.name as role,teams.name as team,
		CONCAT(?::text,'/static/images','/',assets.name) as image_url`, config.Origin,
	).Joins(
		"JOIN assets on assets.id = members.asset_id").Joins(
		"JOIN teams on teams.id = members.team_id").Joins(
		"JOIN roles on roles.id = members.role_id").Scan(
		&members).Error; err != nil {
		return nil, err
	}

	return members, nil
}
