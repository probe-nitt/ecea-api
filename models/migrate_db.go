package models

import (
	"github.com/probe-nitt/probe-server/database"
)

func MigrateDB() {
	db := database.GetDB()

	for _, model := range []interface{}{
		User{},
	} {
		if err := db.AutoMigrate(&model); err != nil {
			panic(err)
		}
	}
}
