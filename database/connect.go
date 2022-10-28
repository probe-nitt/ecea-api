package database

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/probe-nitt/probe-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectDB connect to db
func ConnectDB() {

	config := config.GetConfig()

	var err error

	// using strings.Builder to build the connection string
	var connString strings.Builder

	connString.WriteString("host=")
	connString.WriteString(config.DBHost)
	connString.WriteString(" user=")
	connString.WriteString(config.DBUsername)
	connString.WriteString(" password=")
	connString.WriteString(config.DBPassword)
	connString.WriteString(" dbname=")
	connString.WriteString(config.DBName)
	connString.WriteString(" port=")
	connString.WriteString(config.DBPort)

	// connecting to the database
	db, err = gorm.Open(postgres.Open(connString.String()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		fmt.Println(color.RedString("Error connecting to database", err))
	} else {
		fmt.Println(color.GreenString("Database connected"))
	}
}
