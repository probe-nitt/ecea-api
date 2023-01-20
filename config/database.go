package config

import (
	"fmt"

	"github.com/ecea-nitt/ecea-server/schemas"
	"github.com/fatih/color"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is the database
var db *gorm.DB

// ConnectDB connect to db
func ConnectDB() {

	var er error
	var dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		DbHost, DbUser, DbPassword, DbName, DbPort)

	db, er = gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})
	if er != nil {
		fmt.Println(color.RedString("Error connecting to database"))
	} else {
		fmt.Println(color.GreenString("Database connected"))
	}
}

// GetDB returns the database
func GetDB() *gorm.DB {
	return db
}

// MigrateDB migrates schemas the database
func MigrateDB() {
	db := GetDB()

	for _, schema := range []interface{}{
		&schemas.User{},
		&schemas.AssetType{},
		&schemas.Asset{},
		&schemas.Team{},
		&schemas.Role{},
		&schemas.Member{},
		&schemas.Subject{},
		&schemas.StudyMaterial{},
		&schemas.SubjectCategory{},
	} {
		if err := db.AutoMigrate(&schema); err != nil {
			panic(err)
		}
	}
	fmt.Println(color.BlueString("Migration Success"))
}
