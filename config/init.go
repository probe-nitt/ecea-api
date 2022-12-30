package config

func InitApp() {
	LoadEnvironment()
	ConnectDB()
	MigrateDB()
}
