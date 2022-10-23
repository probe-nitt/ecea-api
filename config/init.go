package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

// Initialize config
func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(color.RedString("Error loading .env file"))
	}

	stage := os.Getenv("STAGE")

	// Load JSON config
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(color.RedString("Error loading config.json file"))
	}
	defer configFile.Close()

	// Parse JSON config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&allConfigurations)

	if err != nil {
		fmt.Println(color.RedString("Error decoding config.json file"))
	}

	// Set current config
	switch stage {
	case "dev":
		currentConfig = allConfigurations.Dev
	case "docker":
		currentConfig = allConfigurations.Docker
	case "prod":
		currentConfig = allConfigurations.Prod
	default:
		fmt.Println(color.RedString("Error setting current config"))
	}
}
