package configs

import (
	"os"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func Load() error {
	cfg = new(config)

	cfg.API = APIConfig{
		Port: getEnv("SERVER_PORT", "9000"),
	}

	cfg.DB = DBConfig{
		Host:     getEnv("DB_HOST", "db"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Pass:     getEnv("DB_PASS ", "secret"),
		Database: getEnv("DB_NAME", "api_todo"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
