package schemas

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	Name        string `gorm:"not null;unique"`
	AssetTypeID uint   `gorm:"not null;"`
	AssetType   AssetType
}

type AssetType struct {
	gorm.Model
	Name string `gorm:"not null;"`
}

type AssetTypes string

const (
	Image    AssetTypes = "IMAGE"
	Document AssetTypes = "DOCUMENT"
)
