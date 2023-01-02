package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"

	"github.com/joho/godotenv"
)

var DbHost string
var DbUser string
var DbPassword string
var DbName string
var DbPort string
var ServerPort string
var Origin string
var JwtSecret string
var MailPort int
var MailHost string
var MailUser string
var MailPassword string
var Admin string

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env"))
	}

	Admin = os.Getenv("ADMIN")
	DbHost = os.Getenv("DB_HOST")
	DbUser = os.Getenv("POSTGRES_USER")
	DbPassword = os.Getenv("POSTGRES_PASSWORD")
	DbName = os.Getenv("POSTGRES_DB")
	DbPort = os.Getenv("POSTGRES_PORT")
	ServerPort = os.Getenv("SERVER_EXTERNAL_PORT")
	Origin = os.Getenv("ORIGIN")
	JwtSecret = os.Getenv("JWT_SECRET")
	MailHost = os.Getenv("MAIL_HOST")
	MailUser = os.Getenv("MAIL_USER")
	MailPassword = os.Getenv("MAIL_PASSWORD")
	MailPort, err = strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic(err)
	}
}
