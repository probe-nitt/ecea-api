package schemas

import (
	"gorm.io/gorm"
)

type Member struct {
	gorm.Model
	Name    string `gorm:"not null;"`
	RollNo  string `gorm:"not null;unique"`
	RoleID  uint   `gorm:"not null;"`
	Role    Role
	TeamID  uint `gorm:"default:null;"`
	Team    Team
	AssetID uint `gorm:"default:null;"`
	Asset   Asset
}

type Role struct {
	gorm.Model
	Name string `gorm:"not null;"`
}

type Team struct {
	gorm.Model
	Name string `gorm:"not null;"`
}
