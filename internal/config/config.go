package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName     string
	Environment string
	ServerPort  string
	DatabaseDSN string
	ApiKey      string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		AppName:     getEnv("APPNAME", "analytics"),
		Environment: getEnv("Environment", "development"),
		ServerPort:  getEnv("ServerPort", "8080"),
		DatabaseDSN: getEnv("DatabaseDSN", ""),
		ApiKey:      getEnv("API_KEY", ""),
	}
}

func getEnv(key, falllback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return falllback
}
