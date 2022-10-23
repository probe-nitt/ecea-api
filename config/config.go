package config

type Config struct {
	// prod, dev, docker
	Stage      string
	ServerPort string

	// Database
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

// All configurations
var allConfigurations = struct {

	// Configuration for environment : dev
	Dev Config

	// Configuration for environment : docker
	Docker Config

	// Configuration for environment : prod
	Prod Config
}{}

// Current config
var currentConfig Config

// Function to get current config
func GetConfig() Config {
	return currentConfig
}
