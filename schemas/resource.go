package schemas

import "gorm.io/gorm"

type StudyMaterial struct {
	gorm.Model
	Name      string  `gorm:"not null; unique"`
	AssetID   uint    `gorm:"default: null"`
	Asset     Asset   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	SubjectID uint    `gorm:"default: null"`
	Subject   Subject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Subject struct {
	gorm.Model
	SubjectCode       string          `gorm:"not null; unique"`
	Name              string          `gorm:"not null"`
	SubjectCategoryID uint            `gorm:"default: null"`
	SubjectCategory   SubjectCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type SubjectCategory struct {
	gorm.Model
	Name string `gorm:"not null; unique"`
}
