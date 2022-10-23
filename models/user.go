package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"default:null;"`
	Name     string `gorm:"default:null;"`
	RollNo   string `gorm:"not null;unique"`
}
