package database

import (
	"gorm.io/gorm"
)

// DB is the database
var db *gorm.DB

// GetDB returns the database
func GetDB() *gorm.DB {
	return db
}
