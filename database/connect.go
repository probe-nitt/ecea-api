package database

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/probe-nitt/probe-server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB connect to db
func ConnectDB() {

	config := config.GetConfig()

	var err error

	// using strings.Builder to build the connection string
	var connectionString strings.Builder

	connectionString.WriteString(config.DBUsername)
	connectionString.WriteString(":")
	connectionString.WriteString(config.DBPassword)
	connectionString.WriteString("@tcp(")
	connectionString.WriteString(config.DBHost)
	connectionString.WriteString(":")
	connectionString.WriteString(config.DBPort)
	connectionString.WriteString(")/")
	connectionString.WriteString(config.DBName)
	connectionString.WriteString("?charset=utf8mb4&parseTime=True&loc=Local")

	// connect to db
	db, err = gorm.Open(mysql.Open(connectionString.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(color.RedString("Error connecting to database"))
	} else {
		fmt.Println(color.GreenString("Database connected"))
	}
}
