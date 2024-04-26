package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPass     string
	DBAddr     string
	DBName     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load(".env")
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "localhost"),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPass:     getEnv("DB_PASS", "root"),
		DBAddr:     getEnv("DB_ADDR", "localhost:3306"),
		DBName:     getEnv("DB_NAME", "simple_api"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
